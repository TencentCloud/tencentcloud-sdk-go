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

package v20180525

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AcquireClusterAdminRoleRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type AcquireClusterAdminRoleRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *AcquireClusterAdminRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireClusterAdminRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcquireClusterAdminRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcquireClusterAdminRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcquireClusterAdminRoleResponse struct {
	*tchttp.BaseResponse
	Response *AcquireClusterAdminRoleResponseParams `json:"Response"`
}

func (r *AcquireClusterAdminRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcquireClusterAdminRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterCIDRRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 增加的ClusterCIDR
	ClusterCIDRs []*string `json:"ClusterCIDRs,omitnil,omitempty" name:"ClusterCIDRs"`

	// 是否忽略ClusterCIDR与VPC路由表的冲突，默认false，为true时忽略冲突
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitnil,omitempty" name:"IgnoreClusterCIDRConflict"`
}

type AddClusterCIDRRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 增加的ClusterCIDR
	ClusterCIDRs []*string `json:"ClusterCIDRs,omitnil,omitempty" name:"ClusterCIDRs"`

	// 是否忽略ClusterCIDR与VPC路由表的冲突，默认false，为true时忽略冲突
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitnil,omitempty" name:"IgnoreClusterCIDRConflict"`
}

func (r *AddClusterCIDRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterCIDRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterCIDRs")
	delete(f, "IgnoreClusterCIDRConflict")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterCIDRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterCIDRResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddClusterCIDRResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterCIDRResponseParams `json:"Response"`
}

func (r *AddClusterCIDRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterCIDRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddExistedInstancesRequestParams struct {
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例列表，不支持竞价实例（请登录 [CVM控制台](https://console.cloud.tencent.com/cvm) 获取待添加节点ID ）
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例额外需要设置参数信息(默认值)
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口HostName一致)
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 节点池选项
	NodePool *NodePoolOption `json:"NodePool,omitnil,omitempty" name:"NodePool"`

	// 校验规则相关选项，可配置跳过某些校验规则。目前支持GlobalRouteCIDRCheck（跳过GlobalRouter的相关校验），VpcCniCIDRCheck（跳过VpcCni相关校验）
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitnil,omitempty" name:"SkipValidateOptions"`

	// 参数InstanceAdvancedSettingsOverride数组用于定制化地配置各台instance，与InstanceIds顺序对应。当传入InstanceAdvancedSettingsOverrides数组时，将覆盖默认参数InstanceAdvancedSettings；当没有传入参数InstanceAdvancedSettingsOverrides时，InstanceAdvancedSettings参数对每台instance生效。参数InstanceAdvancedSettingsOverride数组的长度应与InstanceIds数组一致；当长度大于InstanceIds数组长度时将报错；当长度小于InstanceIds数组时，没有对应配置的instance将使用默认配置。
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitnil,omitempty" name:"InstanceAdvancedSettingsOverrides"`

	// 节点镜像
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type AddExistedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 实例列表，不支持竞价实例（请登录 [CVM控制台](https://console.cloud.tencent.com/cvm) 获取待添加节点ID ）
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例额外需要设置参数信息(默认值)
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口HostName一致)
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 节点池选项
	NodePool *NodePoolOption `json:"NodePool,omitnil,omitempty" name:"NodePool"`

	// 校验规则相关选项，可配置跳过某些校验规则。目前支持GlobalRouteCIDRCheck（跳过GlobalRouter的相关校验），VpcCniCIDRCheck（跳过VpcCni相关校验）
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitnil,omitempty" name:"SkipValidateOptions"`

	// 参数InstanceAdvancedSettingsOverride数组用于定制化地配置各台instance，与InstanceIds顺序对应。当传入InstanceAdvancedSettingsOverrides数组时，将覆盖默认参数InstanceAdvancedSettings；当没有传入参数InstanceAdvancedSettingsOverrides时，InstanceAdvancedSettings参数对每台instance生效。参数InstanceAdvancedSettingsOverride数组的长度应与InstanceIds数组一致；当长度大于InstanceIds数组长度时将报错；当长度小于InstanceIds数组时，没有对应配置的instance将使用默认配置。
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitnil,omitempty" name:"InstanceAdvancedSettingsOverrides"`

	// 节点镜像
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddExistedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "EnhancedService")
	delete(f, "LoginSettings")
	delete(f, "HostName")
	delete(f, "SecurityGroupIds")
	delete(f, "NodePool")
	delete(f, "SkipValidateOptions")
	delete(f, "InstanceAdvancedSettingsOverrides")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddExistedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddExistedInstancesResponseParams struct {
	// 失败的节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 成功的节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitnil,omitempty" name:"SuccInstanceIds"`

	// 超时未返回出来节点的ID(可能失败，也可能成功)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitnil,omitempty" name:"TimeoutInstanceIds"`

	// 失败的节点的失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReasons []*string `json:"FailedReasons,omitnil,omitempty" name:"FailedReasons"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AddExistedInstancesResponseParams `json:"Response"`
}

func (r *AddExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddExistedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodeToNodePoolRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点id，获取参考 https://cloud.tencent.com/document/product/213/15728
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AddNodeToNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点id，获取参考 https://cloud.tencent.com/document/product/213/15728
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AddNodeToNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodeToNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNodeToNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodeToNodePoolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNodeToNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *AddNodeToNodePoolResponseParams `json:"Response"`
}

func (r *AddNodeToNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodeToNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddVpcCniSubnetsRequestParams struct {
	// 集群 ID，请[登录控制台](https://console.cloud.tencent.com/tke2/cluster)在集群列表复制ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 为集群容器网络增加的子网列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 集群所属的VPC的ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 是否同步添加 vpc 网段到 ip-masq-agent-config 的 NonMasqueradeCIDRs 字段，默认 false 会同步添加
	SkipAddingNonMasqueradeCIDRs *bool `json:"SkipAddingNonMasqueradeCIDRs,omitnil,omitempty" name:"SkipAddingNonMasqueradeCIDRs"`
}

type AddVpcCniSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID，请[登录控制台](https://console.cloud.tencent.com/tke2/cluster)在集群列表复制ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 为集群容器网络增加的子网列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 集群所属的VPC的ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 是否同步添加 vpc 网段到 ip-masq-agent-config 的 NonMasqueradeCIDRs 字段，默认 false 会同步添加
	SkipAddingNonMasqueradeCIDRs *bool `json:"SkipAddingNonMasqueradeCIDRs,omitnil,omitempty" name:"SkipAddingNonMasqueradeCIDRs"`
}

func (r *AddVpcCniSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddVpcCniSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubnetIds")
	delete(f, "VpcId")
	delete(f, "SkipAddingNonMasqueradeCIDRs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddVpcCniSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddVpcCniSubnetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddVpcCniSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *AddVpcCniSubnetsResponseParams `json:"Response"`
}

func (r *AddVpcCniSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddVpcCniSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Addon struct {
	// addon名称
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`

	// addon的版本
	AddonVersion *string `json:"AddonVersion,omitnil,omitempty" name:"AddonVersion"`

	// addon的参数，是一个json格式的base64转码后的字符串
	RawValues *string `json:"RawValues,omitnil,omitempty" name:"RawValues"`

	// addon的状态：
	// - Installing：安装中
	// - Upgrading：升级中
	// - Terminating：删除中
	// - Succeeded：安装/升级成功
	// - InstallFailed：安装失败
	// - UpgradFailed：升级失败
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// addon失败的原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// addon的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AnnotationValue struct {
	// 注释键
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 注释值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AppChart struct {
	// chart名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// chart的标签
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// chart的版本
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`
}

type AutoScalingGroupRange struct {
	// 伸缩组最小实例数
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 伸缩组最大实例数
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type AutoUpgradeClusterLevel struct {
	// 是否开启自动变配集群等级
	IsAutoUpgrade *bool `json:"IsAutoUpgrade,omitnil,omitempty" name:"IsAutoUpgrade"`
}

type AutoscalingAdded struct {
	// 正在加入中的节点数量
	Joining *int64 `json:"Joining,omitnil,omitempty" name:"Joining"`

	// 初始化中的节点数量
	Initializing *int64 `json:"Initializing,omitnil,omitempty" name:"Initializing"`

	// 正常的节点数量
	Normal *int64 `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 节点总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type AvailableExtraArgs struct {
	// kube-apiserver可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeAPIServer []*Flag `json:"KubeAPIServer,omitnil,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeControllerManager []*Flag `json:"KubeControllerManager,omitnil,omitempty" name:"KubeControllerManager"`

	// kube-scheduler可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeScheduler []*Flag `json:"KubeScheduler,omitnil,omitempty" name:"KubeScheduler"`

	// kubelet可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kubelet []*Flag `json:"Kubelet,omitnil,omitempty" name:"Kubelet"`
}

type BackupStorageLocation struct {
	// 备份仓库名称	
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 存储仓库所属地域，比如COS广州(ap-guangzhou)	
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 存储服务提供方，默认腾讯云	
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 对象存储桶名称，如果是COS必须是tke-backup-前缀开头	
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 对象存储桶路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 存储仓库状态
	// - Available: 可用
	// - Unavailable: 不可用
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 详细状态信息	
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 最后一次检查时间	
	LastValidationTime *string `json:"LastValidationTime,omitnil,omitempty" name:"LastValidationTime"`
}

type CUDNN struct {
	// cuDNN的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// cuDNN的版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// cuDNN的Dev名字
	DevName *string `json:"DevName,omitnil,omitempty" name:"DevName"`

	// cuDNN的Doc名字
	DocName *string `json:"DocName,omitnil,omitempty" name:"DocName"`
}

// Predefined struct for user
type CancelClusterReleaseRequestParams struct {
	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型，支持传入 tke(标准集群), eks(Serverless集群), external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type CancelClusterReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型，支持传入 tke(标准集群), eks(Serverless集群), external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *CancelClusterReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelClusterReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelClusterReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelClusterReleaseResponseParams struct {
	// 应用信息
	Release *PendingRelease `json:"Release,omitnil,omitempty" name:"Release"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelClusterReleaseResponse struct {
	*tchttp.BaseResponse
	Response *CancelClusterReleaseResponseParams `json:"Response"`
}

func (r *CancelClusterReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelClusterReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelUpgradePlanRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 升级计划ID
	PlanID *int64 `json:"PlanID,omitnil,omitempty" name:"PlanID"`
}

type CancelUpgradePlanRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 升级计划ID
	PlanID *int64 `json:"PlanID,omitnil,omitempty" name:"PlanID"`
}

func (r *CancelUpgradePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelUpgradePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "PlanID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelUpgradePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelUpgradePlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelUpgradePlanResponse struct {
	*tchttp.BaseResponse
	Response *CancelUpgradePlanResponseParams `json:"Response"`
}

func (r *CancelUpgradePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelUpgradePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Capabilities struct {
	// 启用安全能力项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Add []*string `json:"Add,omitnil,omitempty" name:"Add"`

	// 禁用安全能力向列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Drop []*string `json:"Drop,omitnil,omitempty" name:"Drop"`
}

type CbsVolume struct {
	// 腾讯云cbs盘Id
	CbsDiskId *string `json:"CbsDiskId,omitnil,omitempty" name:"CbsDiskId"`

	// cbs volume 数据卷名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type CheckEdgeClusterCIDRRequestParams struct {
	// 集群的vpc-id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群的pod CIDR
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群的service CIDR
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`
}

type CheckEdgeClusterCIDRRequest struct {
	*tchttp.BaseRequest
	
	// 集群的vpc-id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群的pod CIDR
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群的service CIDR
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`
}

func (r *CheckEdgeClusterCIDRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckEdgeClusterCIDRRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "PodCIDR")
	delete(f, "ServiceCIDR")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckEdgeClusterCIDRRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckEdgeClusterCIDRResponseParams struct {
	// 返回码，具体如下
	// -1 内部错误
	// 0 没冲突
	// 1 vpc 和 serviceCIDR 冲突
	// 2 vpc 和 podCIDR 冲突
	// 3 serviceCIDR  和 podCIDR 冲突
	ConflictCode *int64 `json:"ConflictCode,omitnil,omitempty" name:"ConflictCode"`

	// CIDR冲突描述信息。
	ConflictMsg *string `json:"ConflictMsg,omitnil,omitempty" name:"ConflictMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckEdgeClusterCIDRResponse struct {
	*tchttp.BaseResponse
	Response *CheckEdgeClusterCIDRResponseParams `json:"Response"`
}

func (r *CheckEdgeClusterCIDRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckEdgeClusterCIDRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstancesUpgradeAbleRequestParams struct {
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点列表，空为全部节点
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 升级类型，枚举值：reset(重装升级，支持大版本和小版本)，hot(原地滚动小版本升级)，major(原地滚动大版本升级)
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 偏移量，默认为0。关于 Offset 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件列表，Name 可选值为ip、instanceId、hostname、label
	// 
	// - Name 为 ip 时，Value 为待查询节点 ip
	// - Name 为 hostname 时，Value 为待查询节点名称
	// - Name 为 label 时，Value 为待查询 K8S 节点 label
	// - Name 为 instanceId 时，Value 为 待查询节点 id
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CheckInstancesUpgradeAbleRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点列表，空为全部节点
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 升级类型，枚举值：reset(重装升级，支持大版本和小版本)，hot(原地滚动小版本升级)，major(原地滚动大版本升级)
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 偏移量，默认为0。关于 Offset 的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件列表，Name 可选值为ip、instanceId、hostname、label
	// 
	// - Name 为 ip 时，Value 为待查询节点 ip
	// - Name 为 hostname 时，Value 为待查询节点名称
	// - Name 为 label 时，Value 为待查询 K8S 节点 label
	// - Name 为 instanceId 时，Value 为 待查询节点 id
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CheckInstancesUpgradeAbleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstancesUpgradeAbleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "UpgradeType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckInstancesUpgradeAbleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckInstancesUpgradeAbleResponseParams struct {
	// 集群master当前小版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 集群master对应的大版本目前最新小版本
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// 可升级节点列表
	UpgradeAbleInstances []*UpgradeAbleInstancesItem `json:"UpgradeAbleInstances,omitnil,omitempty" name:"UpgradeAbleInstances"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 不可升级原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnavailableVersionReason []*UnavailableReason `json:"UnavailableVersionReason,omitnil,omitempty" name:"UnavailableVersionReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckInstancesUpgradeAbleResponse struct {
	*tchttp.BaseResponse
	Response *CheckInstancesUpgradeAbleResponseParams `json:"Response"`
}

func (r *CheckInstancesUpgradeAbleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckInstancesUpgradeAbleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDescription *string `json:"ClusterDescription,omitnil,omitempty" name:"ClusterDescription"`

	// 集群版本（默认值为1.10.5）
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 集群系统。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，默认取值为ubuntu16.04.1 LTSx86_64
	ClusterOs *string `json:"ClusterOs,omitnil,omitempty" name:"ClusterOs"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群网络相关参数
	ClusterNetworkSettings *ClusterNetworkSettings `json:"ClusterNetworkSettings,omitnil,omitempty" name:"ClusterNetworkSettings"`

	// 集群当前node数量
	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitnil,omitempty" name:"ClusterNodeNum"`

	// 集群所属的项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签描述列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 集群状态 (Trading 集群开通中,Creating 创建中,Running 运行中,Deleting 删除中,Idling 闲置中,Recovering 唤醒中,Scaling 规模调整中,Upgrading 升级中,WaittingForConnect 等待注册,Trading 集群开通中,Isolated 欠费隔离中,Pause 集群升级暂停,NodeUpgrading 节点升级中,RuntimeUpgrading 节点运行时升级中,MasterScaling Master扩缩容中,ClusterLevelUpgrading 调整规格中,ResourceIsolate 隔离中,ResourceIsolated 已隔离,ResourceReverse 冲正中,Abnormal 异常)
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// 集群属性(包括集群不同属性的MAP，属性字段包括NodeNameType (lan-ip模式和hostname 模式，默认无lan-ip模式))
	Property *string `json:"Property,omitnil,omitempty" name:"Property"`

	// 集群当前master数量
	ClusterMaterNodeNum *uint64 `json:"ClusterMaterNodeNum,omitnil,omitempty" name:"ClusterMaterNodeNum"`

	// 集群使用镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// OsCustomizeType 系统定制类型
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// 集群运行环境docker或container
	ContainerRuntime *string `json:"ContainerRuntime,omitnil,omitempty" name:"ContainerRuntime"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 集群删除保护开关，打开：true，关闭：false
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 集群是否开启第三方节点支持，开启：true，关闭：false
	EnableExternalNode *bool `json:"EnableExternalNode,omitnil,omitempty" name:"EnableExternalNode"`

	// 集群等级，针对托管集群生效
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 自动变配集群等级，针对托管集群生效。开启：true，关闭：false
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 是否开启QGPU共享，开启：true，关闭：false
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitnil,omitempty" name:"QGPUShareEnable"`

	// 运行时版本
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// 集群当前etcd数量
	ClusterEtcdNodeNum *uint64 `json:"ClusterEtcdNodeNum,omitnil,omitempty" name:"ClusterEtcdNodeNum"`

	// 本地专用集群Id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type ClusterAdvancedSettings struct {
	// 是否启用集群节点自动扩缩容(创建集群流程不支持开启此功能)
	AsEnabled *bool `json:"AsEnabled,omitnil,omitempty" name:"AsEnabled"`

	// 是否开启审计开关
	AuditEnabled *bool `json:"AuditEnabled,omitnil,omitempty" name:"AuditEnabled"`

	// 审计日志上传到的topic
	AuditLogTopicId *string `json:"AuditLogTopicId,omitnil,omitempty" name:"AuditLogTopicId"`

	// 审计日志上传到的logset日志集
	AuditLogsetId *string `json:"AuditLogsetId,omitnil,omitempty" name:"AuditLogsetId"`

	// 自定义模式下的基础pod数量
	BasePodNumber *int64 `json:"BasePodNumber,omitnil,omitempty" name:"BasePodNumber"`

	// 启用 CiliumMode 的模式，空值表示不启用，“clusterIP” 表示启用 Cilium 支持 ClusterIP
	CiliumMode *string `json:"CiliumMode,omitnil,omitempty" name:"CiliumMode"`

	// 集群使用的runtime类型，包括"docker"和"containerd"两种类型，默认为"docker"
	ContainerRuntime *string `json:"ContainerRuntime,omitnil,omitempty" name:"ContainerRuntime"`

	// 是否启用 DataPlaneV2（cilium替代kube-proxy） 
	DataPlaneV2 *bool `json:"DataPlaneV2,omitnil,omitempty" name:"DataPlaneV2"`

	// 是否启用集群删除保护
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 是否开节点podCIDR大小的自定义模式
	EnableCustomizedPodCIDR *bool `json:"EnableCustomizedPodCIDR,omitnil,omitempty" name:"EnableCustomizedPodCIDR"`

	// 元数据拆分存储Etcd配置
	EtcdOverrideConfigs []*EtcdOverrideConfig `json:"EtcdOverrideConfigs,omitnil,omitempty" name:"EtcdOverrideConfigs"`

	// 集群自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// 是否启用IPVS
	IPVS *bool `json:"IPVS,omitnil,omitempty" name:"IPVS"`

	// 集群VPC-CNI模式下是否是双栈集群，默认false，表明非双栈集群。
	IsDualStack *bool `json:"IsDualStack,omitnil,omitempty" name:"IsDualStack"`

	// 集群VPC-CNI模式是否为非固定IP，默认: FALSE 固定IP。
	IsNonStaticIpMode *bool `json:"IsNonStaticIpMode,omitnil,omitempty" name:"IsNonStaticIpMode"`

	// 集群的网络代理模型，目前tke集群支持的网络代理模式有三种：iptables,ipvs,ipvs-bpf，此参数仅在使用ipvs-bpf模式时使用，三种网络模式的参数设置关系如下：
	// iptables模式：IPVS和KubeProxyMode都不设置
	// ipvs模式: 设置IPVS为true, KubeProxyMode不设置
	// ipvs-bpf模式: 设置KubeProxyMode为kube-proxy-bpf
	// 使用ipvs-bpf的网络模式需要满足以下条件：
	// 1. 集群版本必须为1.14及以上；
	// 2. 系统镜像必须是: Tencent Linux 2.4；
	KubeProxyMode *string `json:"KubeProxyMode,omitnil,omitempty" name:"KubeProxyMode"`

	// 集群网络类型，默认为GR。
	// - GR: 全局路由
	// - VPC-CNI: VPC-CNI模式
	// - CiliumOverlay: CiliumOverlay模式
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 集群中节点NodeName类型（包括 hostname,lan-ip两种形式，默认为lan-ip。如果开启了hostname模式，创建节点时需要设置HostName参数，并且InstanceName需要和HostName一致）
	NodeNameType *string `json:"NodeNameType,omitnil,omitempty" name:"NodeNameType"`

	// 是否开启QGPU共享
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitnil,omitempty" name:"QGPUShareEnable"`

	// 运行时版本
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// 区分共享网卡多IP模式和独立网卡模式，共享网卡多 IP 模式填写"tke-route-eni"，独立网卡模式填写"tke-direct-eni"，默认为共享网卡模式
	VpcCniType *string `json:"VpcCniType,omitnil,omitempty" name:"VpcCniType"`
}

type ClusterAsGroup struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组状态(开启 enabled 开启中 enabling 关闭 disabled 关闭中 disabling 更新中 updating 删除中 deleting 开启缩容中 scaleDownEnabling 关闭缩容中 scaleDownDisabling)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点是否设置成不可调度
	IsUnschedulable *bool `json:"IsUnschedulable,omitnil,omitempty" name:"IsUnschedulable"`

	// 伸缩组的label列表
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type ClusterAsGroupAttribute struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 是否开启
	AutoScalingGroupEnabled *bool `json:"AutoScalingGroupEnabled,omitnil,omitempty" name:"AutoScalingGroupEnabled"`

	// 伸缩组最大最小实例数
	AutoScalingGroupRange *AutoScalingGroupRange `json:"AutoScalingGroupRange,omitnil,omitempty" name:"AutoScalingGroupRange"`
}

type ClusterAsGroupOption struct {
	// 是否开启缩容
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsScaleDownEnabled *bool `json:"IsScaleDownEnabled,omitnil,omitempty" name:"IsScaleDownEnabled"`

	// 多伸缩组情况下扩容选择算法(random 随机选择，most-pods 最多类型的Pod least-waste 最少的资源浪费，默认为random)
	Expander *string `json:"Expander,omitnil,omitempty" name:"Expander"`

	// 最大并发缩容数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxEmptyBulkDelete *int64 `json:"MaxEmptyBulkDelete,omitnil,omitempty" name:"MaxEmptyBulkDelete"`

	// 集群扩容后多少分钟开始判断缩容（默认为10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownDelay *int64 `json:"ScaleDownDelay,omitnil,omitempty" name:"ScaleDownDelay"`

	// 节点连续空闲多少分钟后被缩容（默认为 10分钟）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownUnneededTime *int64 `json:"ScaleDownUnneededTime,omitnil,omitempty" name:"ScaleDownUnneededTime"`

	// 节点资源使用量低于多少(百分比)时认为空闲(默认: 50(百分比))
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownUtilizationThreshold *int64 `json:"ScaleDownUtilizationThreshold,omitnil,omitempty" name:"ScaleDownUtilizationThreshold"`

	// 含有本地存储Pod的节点是否不缩容(默认： true)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipNodesWithLocalStorage *bool `json:"SkipNodesWithLocalStorage,omitnil,omitempty" name:"SkipNodesWithLocalStorage"`

	// 含有kube-system namespace下非DaemonSet管理的Pod的节点是否不缩容 (默认： true)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipNodesWithSystemPods *bool `json:"SkipNodesWithSystemPods,omitnil,omitempty" name:"SkipNodesWithSystemPods"`

	// 计算资源使用量时是否默认忽略DaemonSet的实例(默认值: False，不忽略)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreDaemonSetsUtilization *bool `json:"IgnoreDaemonSetsUtilization,omitnil,omitempty" name:"IgnoreDaemonSetsUtilization"`

	// CA做健康性判断的个数，默认3，即超过OkTotalUnreadyCount个数后，CA会进行健康性判断。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OkTotalUnreadyCount *int64 `json:"OkTotalUnreadyCount,omitnil,omitempty" name:"OkTotalUnreadyCount"`

	// 未就绪节点的最大百分比，此后CA会停止操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTotalUnreadyPercentage *int64 `json:"MaxTotalUnreadyPercentage,omitnil,omitempty" name:"MaxTotalUnreadyPercentage"`

	// 表示未准备就绪的节点在有资格进行缩减之前应该停留多少分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDownUnreadyTime *int64 `json:"ScaleDownUnreadyTime,omitnil,omitempty" name:"ScaleDownUnreadyTime"`

	// CA删除未在Kubernetes中注册的节点之前等待的分钟数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnregisteredNodeRemovalTime *int64 `json:"UnregisteredNodeRemovalTime,omitnil,omitempty" name:"UnregisteredNodeRemovalTime"`
}

type ClusterBasicSettings struct {
	// 集群操作系统，支持设置公共镜像(字段传相应镜像Name)和自定义镜像(字段传相应镜像ID)，详情参考：https://cloud.tencent.com/document/product/457/68289
	ClusterOs *string `json:"ClusterOs,omitnil,omitempty" name:"ClusterOs"`

	// 集群版本,默认值为1.10.5
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDescription *string `json:"ClusterDescription,omitnil,omitempty" name:"ClusterDescription"`

	// 私有网络ID，形如vpc-e55paxnt。创建托管空集群时必传。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群内新增资源所属项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到集群实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// 是否开启节点的默认安全组(默认: 否，Alpha特性)
	NeedWorkSecurityGroup *bool `json:"NeedWorkSecurityGroup,omitnil,omitempty" name:"NeedWorkSecurityGroup"`

	// 控制面子网信息，仅在以下场景使用时要求必填。
	// - 容器网络插件为CiliumOverlay时，TKE会从该子网获取2个IP用来创建内网负载均衡。
	// - 创建支持CDC的托管集群，且网络插件为VPC-CNI时，要求预留至少12个IP。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 集群等级，针对托管集群生效
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 自动变配集群等级，针对托管集群生效
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`
}

type ClusterCIDRSettings struct {
	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。
	ClusterCIDR *string `json:"ClusterCIDR,omitnil,omitempty" name:"ClusterCIDR"`

	// 是否忽略 ClusterCIDR 冲突错误, 默认不忽略
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitnil,omitempty" name:"IgnoreClusterCIDRConflict"`

	// 集群中每个Node上最大的Pod数量。取值范围16～256。不为2的幂值时会向上取最接近的2的幂值。
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitnil,omitempty" name:"MaxNodePodNum"`

	// 集群最大的service数量。取值范围32～32768，不为2的幂值时会向上取最接近的2的幂值。默认值256
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitnil,omitempty" name:"MaxClusterServiceNum"`

	// 用于分配集群服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// VPC-CNI网络模式下，弹性网卡的子网Id。
	EniSubnetIds []*string `json:"EniSubnetIds,omitnil,omitempty" name:"EniSubnetIds"`

	// VPC-CNI网络模式下，弹性网卡IP的回收时间，取值范围[300,15768000)
	ClaimExpiredSeconds *int64 `json:"ClaimExpiredSeconds,omitnil,omitempty" name:"ClaimExpiredSeconds"`

	// 是否忽略 ServiceCIDR 冲突错误, 仅在 VPC-CNI 模式生效，默认不忽略
	IgnoreServiceCIDRConflict *bool `json:"IgnoreServiceCIDRConflict,omitnil,omitempty" name:"IgnoreServiceCIDRConflict"`
}

type ClusterCondition struct {
	// 集群创建过程类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 集群创建过程状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 最后一次探测到该状态的时间
	LastProbeTime *string `json:"LastProbeTime,omitnil,omitempty" name:"LastProbeTime"`

	// 最后一次转换到该过程的时间
	LastTransitionTime *string `json:"LastTransitionTime,omitnil,omitempty" name:"LastTransitionTime"`

	// 转换到该过程的简明原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 转换到该过程的更多信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ClusterCredential struct {
	// CA 根证书
	CACert *string `json:"CACert,omitnil,omitempty" name:"CACert"`

	// 认证用的Token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type ClusterExtraArgs struct {
	// etcd自定义参数，只支持独立集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	Etcd []*string `json:"Etcd,omitnil,omitempty" name:"Etcd"`

	// kube-apiserver自定义参数，参数格式为["k1=v1", "k1=v2"]， 例如["max-requests-inflight=500","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeAPIServer []*string `json:"KubeAPIServer,omitnil,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeControllerManager []*string `json:"KubeControllerManager,omitnil,omitempty" name:"KubeControllerManager"`

	// kube-scheduler自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	KubeScheduler []*string `json:"KubeScheduler,omitnil,omitempty" name:"KubeScheduler"`
}

type ClusterInternalLB struct {
	// 是否开启内网访问LB
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 内网访问LB关联的子网Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ClusterLevelAttribute struct {
	// 集群等级
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 等级名称
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// Pod数量
	PodCount *uint64 `json:"PodCount,omitnil,omitempty" name:"PodCount"`

	// Configmap数量
	ConfigMapCount *uint64 `json:"ConfigMapCount,omitnil,omitempty" name:"ConfigMapCount"`

	// ReplicaSets数量
	RSCount *uint64 `json:"RSCount,omitnil,omitempty" name:"RSCount"`

	// CRD数量
	CRDCount *uint64 `json:"CRDCount,omitnil,omitempty" name:"CRDCount"`

	// 是否启用
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 其他资源数量
	OtherCount *uint64 `json:"OtherCount,omitnil,omitempty" name:"OtherCount"`
}

type ClusterLevelChangeRecord struct {
	// 记录ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 变配状态：trading 发货中,upgrading 变配中,success 变配成功,failed 变配失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 变配前规模
	OldLevel *string `json:"OldLevel,omitnil,omitempty" name:"OldLevel"`

	// 变配后规模
	NewLevel *string `json:"NewLevel,omitnil,omitempty" name:"NewLevel"`

	// 变配触发类型：manual 手动,auto 自动
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 开始时间
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// 结束时间
	EndedAt *string `json:"EndedAt,omitnil,omitempty" name:"EndedAt"`
}

type ClusterNetworkSettings struct {
	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突
	ClusterCIDR *string `json:"ClusterCIDR,omitnil,omitempty" name:"ClusterCIDR"`

	// 是否忽略 ClusterCIDR 冲突错误, 默认不忽略
	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitnil,omitempty" name:"IgnoreClusterCIDRConflict"`

	// 集群中每个Node上最大的Pod数量(默认为256)
	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitnil,omitempty" name:"MaxNodePodNum"`

	// 集群最大的service数量(默认为256)
	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitnil,omitempty" name:"MaxClusterServiceNum"`

	// 是否启用IPVS(默认不开启)
	Ipvs *bool `json:"Ipvs,omitnil,omitempty" name:"Ipvs"`

	// 集群的VPCID（如果创建空集群，为必传值，否则自动设置为和集群的节点保持一致）
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 网络插件是否启用CNI(默认开启)
	Cni *bool `json:"Cni,omitnil,omitempty" name:"Cni"`

	// service的网络模式，当前参数只适用于ipvs+bpf模式
	KubeProxyMode *string `json:"KubeProxyMode,omitnil,omitempty" name:"KubeProxyMode"`

	// 用于分配service的IP range，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 集群关联的容器子网
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnets []*string `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// 是否忽略 ServiceCIDR 冲突错误, 仅在 VPC-CNI 模式生效，默认不忽略
	IgnoreServiceCIDRConflict *bool `json:"IgnoreServiceCIDRConflict,omitnil,omitempty" name:"IgnoreServiceCIDRConflict"`

	// 集群VPC-CNI模式是否为非双栈集群，默认false，非双栈。
	IsDualStack *bool `json:"IsDualStack,omitnil,omitempty" name:"IsDualStack"`

	// 用于分配service的IP range，由系统自动分配
	Ipv6ServiceCIDR *string `json:"Ipv6ServiceCIDR,omitnil,omitempty" name:"Ipv6ServiceCIDR"`

	// 集群Cilium Mode配置
	// - clusterIP
	CiliumMode *string `json:"CiliumMode,omitnil,omitempty" name:"CiliumMode"`

	// 控制面子网信息，仅在以下场景返回。
	// - 容器网络插件为CiliumOverlay。
	// - 支持CDC的托管集群，且网络插件为VPC-CNI。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否启用了 DataPlaneV2（cilium替代kube-proxy）
	DataPlaneV2 *bool `json:"DataPlaneV2,omitnil,omitempty" name:"DataPlaneV2"`
}

type ClusterProperty struct {
	// 节点hostname命名模式
	NodeNameType *string `json:"NodeNameType,omitnil,omitempty" name:"NodeNameType"`
}

type ClusterPublicLB struct {
	// 是否开启公网访问LB
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 允许访问的来源CIDR列表
	AllowFromCidrs []*string `json:"AllowFromCidrs,omitnil,omitempty" name:"AllowFromCidrs"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`

	// 外网访问相关的扩展参数，格式为json
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 新内外网功能，需要传递安全组
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`
}

type ClusterRollOutSequenceTag struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群发布序列标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type ClusterStatus struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群状态
	ClusterState *string `json:"ClusterState,omitnil,omitempty" name:"ClusterState"`

	// 集群下机器实例的状态
	ClusterInstanceState *string `json:"ClusterInstanceState,omitnil,omitempty" name:"ClusterInstanceState"`

	// 集群是否开启监控
	ClusterBMonitor *bool `json:"ClusterBMonitor,omitnil,omitempty" name:"ClusterBMonitor"`

	// 集群创建中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterInitNodeNum *int64 `json:"ClusterInitNodeNum,omitnil,omitempty" name:"ClusterInitNodeNum"`

	// 集群运行中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterRunningNodeNum *int64 `json:"ClusterRunningNodeNum,omitnil,omitempty" name:"ClusterRunningNodeNum"`

	// 集群异常的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterFailedNodeNum *int64 `json:"ClusterFailedNodeNum,omitnil,omitempty" name:"ClusterFailedNodeNum"`

	// 集群已关机的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterClosedNodeNum *int64 `json:"ClusterClosedNodeNum,omitnil,omitempty" name:"ClusterClosedNodeNum"`

	// 集群关机中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败
	ClusterClosingNodeNum *int64 `json:"ClusterClosingNodeNum,omitnil,omitempty" name:"ClusterClosingNodeNum"`

	// 集群是否开启删除保护
	ClusterDeletionProtection *bool `json:"ClusterDeletionProtection,omitnil,omitempty" name:"ClusterDeletionProtection"`

	// 集群是否可审计
	ClusterAuditEnabled *bool `json:"ClusterAuditEnabled,omitnil,omitempty" name:"ClusterAuditEnabled"`
}

type ClusterVersion struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群主版本号列表，例如1.18.4
	Versions []*string `json:"Versions,omitnil,omitempty" name:"Versions"`
}

type CommonName struct {
	// 子账户UIN
	SubaccountUin *string `json:"SubaccountUin,omitnil,omitempty" name:"SubaccountUin"`

	// 子账户客户端证书中的CommonName字段
	CN *string `json:"CN,omitnil,omitempty" name:"CN"`
}

type ComponentLogConfig struct {
	// 组件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 日志级别，对于支持动态调整日志级别的组件，开启日志时可指定该参数
	LogLevel *int64 `json:"LogLevel,omitnil,omitempty" name:"LogLevel"`

	// 日志集ID。如果不指定，则自动创建
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 日志主题ID。如果不指定，则自动创建
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// topic 所属region。 该参数可实现日志跨地域投递
	TopicRegion *string `json:"TopicRegion,omitnil,omitempty" name:"TopicRegion"`
}

type Container struct {
	// 镜像
	Image *string `json:"Image,omitnil,omitempty" name:"Image"`

	// 容器名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 容器启动参数
	Args []*string `json:"Args,omitnil,omitempty" name:"Args"`

	// 容器启动命令
	Commands []*string `json:"Commands,omitnil,omitempty" name:"Commands"`

	// CPU，制改容器最多可使用的核数，该值不可超过容器实例的总核数。单位：核。
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 当前状态
	CurrentState *ContainerState `json:"CurrentState,omitnil,omitempty" name:"CurrentState"`

	// 容器内操作系统的环境变量
	EnvironmentVars []*EnvironmentVariable `json:"EnvironmentVars,omitnil,omitempty" name:"EnvironmentVars"`

	// Gpu限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuLimit *uint64 `json:"GpuLimit,omitnil,omitempty" name:"GpuLimit"`

	// 存活探针
	// 注意：此字段可能返回 null，表示取不到有效值。
	LivenessProbe *LivenessOrReadinessProbe `json:"LivenessProbe,omitnil,omitempty" name:"LivenessProbe"`

	// 内存，限制该容器最多可使用的内存值，该值不可超过容器实例的总内存值。单位：GiB
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 就绪探针
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadinessProbe *LivenessOrReadinessProbe `json:"ReadinessProbe,omitnil,omitempty" name:"ReadinessProbe"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *uint64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 容器的安全上下文
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityContext *SecurityContext `json:"SecurityContext,omitnil,omitempty" name:"SecurityContext"`

	// 数据卷挂载信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeMounts []*VolumeMount `json:"VolumeMounts,omitnil,omitempty" name:"VolumeMounts"`

	// 容器工作目录
	WorkingDir *string `json:"WorkingDir,omitnil,omitempty" name:"WorkingDir"`
}

type ContainerState struct {
	// 容器运行退出码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExitCode *int64 `json:"ExitCode,omitnil,omitempty" name:"ExitCode"`

	// 容器运行结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 容器状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 容器状态 Reason
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 容器重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitnil,omitempty" name:"RestartCount"`

	// 容器运行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 容器状态：created, running, exited, unknown
	State *string `json:"State,omitnil,omitempty" name:"State"`
}

type ControllerStatus struct {
	// 控制器的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 控制器是否开启
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type CreateBackupStorageLocationRequestParams struct {
	// 存储仓库所属地域，比如COS广州(ap-guangzhou)
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 对象存储桶名称，如果是COS必须是tke-backup前缀开头，字符长度是19
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 备份仓库名称，字符长度为63
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 存储服务提供方，默认腾讯云
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 对象存储桶路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type CreateBackupStorageLocationRequest struct {
	*tchttp.BaseRequest
	
	// 存储仓库所属地域，比如COS广州(ap-guangzhou)
	StorageRegion *string `json:"StorageRegion,omitnil,omitempty" name:"StorageRegion"`

	// 对象存储桶名称，如果是COS必须是tke-backup前缀开头，字符长度是19
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 备份仓库名称，字符长度为63
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 存储服务提供方，默认腾讯云
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 对象存储桶路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

func (r *CreateBackupStorageLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupStorageLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StorageRegion")
	delete(f, "Bucket")
	delete(f, "Name")
	delete(f, "Provider")
	delete(f, "Path")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupStorageLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupStorageLocationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBackupStorageLocationResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupStorageLocationResponseParams `json:"Response"`
}

func (r *CreateBackupStorageLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupStorageLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSLogConfigRequestParams struct {
	// 日志采集配置的json表达
	LogConfig *string `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CLS日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 当前集群类型支持tke（标准集群）、eks（serverless集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type CreateCLSLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 日志采集配置的json表达
	LogConfig *string `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CLS日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 当前集群类型支持tke（标准集群）、eks（serverless集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *CreateCLSLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogConfig")
	delete(f, "ClusterId")
	delete(f, "LogsetId")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCLSLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCLSLogConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCLSLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateCLSLogConfigResponseParams `json:"Response"`
}

func (r *CreateCLSLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCLSLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群端口所在的子网ID  (仅在开启非外网访问时需要填，必须为集群所在VPC内的子网)。获取方式：https://cloud.tencent.com/document/product/215/15784
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`

	// 设置域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 使用的安全组，只有外网访问需要传递（开启外网访问且不使用已有clb时必传）。获取方式：https://cloud.tencent.com/document/api/215/15808
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 创建lb参数，只有外网访问需要设置，是一个json格式化后的字符串：{"InternetAccessible":{"InternetChargeType":"TRAFFIC_POSTPAID_BY_HOUR","InternetMaxBandwidthOut":200},"VipIsp":"","BandwidthPackageId":""}。
	// 各个参数意义：
	// InternetAccessible.InternetChargeType含义：TRAFFIC_POSTPAID_BY_HOUR按流量按小时后计费;BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费;InternetAccessible.BANDWIDTH_PACKAGE 按带宽包计费。
	// InternetMaxBandwidthOut含义：最大出带宽，单位Mbps，范围支持0到2048，默认值10。
	// VipIsp含义：CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通，如果不指定本参数，则默认使用BGP。可通过 DescribeSingleIsp 接口查询一个地域所支持的Isp。如果指定运营商，则网络计费式只能使用按带宽包计费BANDWIDTH_PACKAGE。
	// BandwidthPackageId含义：带宽包ID，指定此参数时，网络计费方式InternetAccessible.InternetChargeType只支持按带宽包计费BANDWIDTH_PACKAGE。
	ExtensiveParameters *string `json:"ExtensiveParameters,omitnil,omitempty" name:"ExtensiveParameters"`

	// 使用已有clb开启内网或外网访问
	ExistedLoadBalancerId *string `json:"ExistedLoadBalancerId,omitnil,omitempty" name:"ExistedLoadBalancerId"`
}

type CreateClusterEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群端口所在的子网ID  (仅在开启非外网访问时需要填，必须为集群所在VPC内的子网)。获取方式：https://cloud.tencent.com/document/product/215/15784
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`

	// 设置域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 使用的安全组，只有外网访问需要传递（开启外网访问且不使用已有clb时必传）。获取方式：https://cloud.tencent.com/document/api/215/15808
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 创建lb参数，只有外网访问需要设置，是一个json格式化后的字符串：{"InternetAccessible":{"InternetChargeType":"TRAFFIC_POSTPAID_BY_HOUR","InternetMaxBandwidthOut":200},"VipIsp":"","BandwidthPackageId":""}。
	// 各个参数意义：
	// InternetAccessible.InternetChargeType含义：TRAFFIC_POSTPAID_BY_HOUR按流量按小时后计费;BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费;InternetAccessible.BANDWIDTH_PACKAGE 按带宽包计费。
	// InternetMaxBandwidthOut含义：最大出带宽，单位Mbps，范围支持0到2048，默认值10。
	// VipIsp含义：CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通，如果不指定本参数，则默认使用BGP。可通过 DescribeSingleIsp 接口查询一个地域所支持的Isp。如果指定运营商，则网络计费式只能使用按带宽包计费BANDWIDTH_PACKAGE。
	// BandwidthPackageId含义：带宽包ID，指定此参数时，网络计费方式InternetAccessible.InternetChargeType只支持按带宽包计费BANDWIDTH_PACKAGE。
	ExtensiveParameters *string `json:"ExtensiveParameters,omitnil,omitempty" name:"ExtensiveParameters"`

	// 使用已有clb开启内网或外网访问
	ExistedLoadBalancerId *string `json:"ExistedLoadBalancerId,omitnil,omitempty" name:"ExistedLoadBalancerId"`
}

func (r *CreateClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubnetId")
	delete(f, "IsExtranet")
	delete(f, "Domain")
	delete(f, "SecurityGroup")
	delete(f, "ExtensiveParameters")
	delete(f, "ExistedLoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterEndpointResponseParams `json:"Response"`
}

func (r *CreateClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointVipRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`
}

type CreateClusterEndpointVipRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`
}

func (r *CreateClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SecurityPolicies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterEndpointVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterEndpointVipResponseParams struct {
	// 请求任务的FlowId
	RequestFlowId *int64 `json:"RequestFlowId,omitnil,omitempty" name:"RequestFlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterEndpointVipResponseParams `json:"Response"`
}

func (r *CreateClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterEndpointVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterInstancesRequestParams struct {
	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CVM创建透传参数，json化字符串格式，如需要保证扩展集群节点请求幂等性需要在此参数添加ClientToken字段，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口。
	RunInstancePara *string `json:"RunInstancePara,omitnil,omitempty" name:"RunInstancePara"`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 校验规则相关选项，可配置跳过某些校验规则。目前支持GlobalRouteCIDRCheck（跳过GlobalRouter的相关校验），VpcCniCIDRCheck（跳过VpcCni相关校验）
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitnil,omitempty" name:"SkipValidateOptions"`
}

type CreateClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CVM创建透传参数，json化字符串格式，如需要保证扩展集群节点请求幂等性需要在此参数添加ClientToken字段，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口。
	RunInstancePara *string `json:"RunInstancePara,omitnil,omitempty" name:"RunInstancePara"`

	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 校验规则相关选项，可配置跳过某些校验规则。目前支持GlobalRouteCIDRCheck（跳过GlobalRouter的相关校验），VpcCniCIDRCheck（跳过VpcCni相关校验）
	SkipValidateOptions []*string `json:"SkipValidateOptions,omitnil,omitempty" name:"SkipValidateOptions"`
}

func (r *CreateClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RunInstancePara")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "SkipValidateOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterInstancesResponseParams struct {
	// 节点实例ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitnil,omitempty" name:"InstanceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterInstancesResponseParams `json:"Response"`
}

func (r *CreateClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterMaintenanceWindowAndExclusionsRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

type CreateClusterMaintenanceWindowAndExclusionsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

func (r *CreateClusterMaintenanceWindowAndExclusionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterMaintenanceWindowAndExclusionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "MaintenanceTime")
	delete(f, "Duration")
	delete(f, "DayOfWeek")
	delete(f, "Exclusions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterMaintenanceWindowAndExclusionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterMaintenanceWindowAndExclusionsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterMaintenanceWindowAndExclusionsResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterMaintenanceWindowAndExclusionsResponseParams `json:"Response"`
}

func (r *CreateClusterMaintenanceWindowAndExclusionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterMaintenanceWindowAndExclusionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterNodePoolRequestParams struct {
	// cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// AutoScalingGroupPara AS组参数，参考 https://cloud.tencent.com/document/product/377/20440
	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitnil,omitempty" name:"AutoScalingGroupPara"`

	// LaunchConfigurePara 运行参数，参考 https://cloud.tencent.com/document/product/377/20447
	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitnil,omitempty" name:"LaunchConfigurePara"`

	// InstanceAdvancedSettings
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 是否启用自动伸缩
	EnableAutoscale *bool `json:"EnableAutoscale,omitnil,omitempty" name:"EnableAutoscale"`

	// 节点池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Labels标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Taints互斥
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点Annotation 列表
	Annotations []*AnnotationValue `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 节点池纬度运行时类型及版本
	ContainerRuntime *string `json:"ContainerRuntime,omitnil,omitempty" name:"ContainerRuntime"`

	// 运行时版本
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// 节点池os，当为自定义镜像时，传镜像id；否则为公共镜像的osName
	NodePoolOs *string `json:"NodePoolOs,omitnil,omitempty" name:"NodePoolOs"`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 删除保护开关
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type CreateClusterNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// cluster id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// AutoScalingGroupPara AS组参数，参考 https://cloud.tencent.com/document/product/377/20440
	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitnil,omitempty" name:"AutoScalingGroupPara"`

	// LaunchConfigurePara 运行参数，参考 https://cloud.tencent.com/document/product/377/20447
	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitnil,omitempty" name:"LaunchConfigurePara"`

	// InstanceAdvancedSettings
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 是否启用自动伸缩
	EnableAutoscale *bool `json:"EnableAutoscale,omitnil,omitempty" name:"EnableAutoscale"`

	// 节点池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Labels标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Taints互斥
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点Annotation 列表
	Annotations []*AnnotationValue `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 节点池纬度运行时类型及版本
	ContainerRuntime *string `json:"ContainerRuntime,omitnil,omitempty" name:"ContainerRuntime"`

	// 运行时版本
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// 节点池os，当为自定义镜像时，传镜像id；否则为公共镜像的osName
	NodePoolOs *string `json:"NodePoolOs,omitnil,omitempty" name:"NodePoolOs"`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 删除保护开关
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *CreateClusterNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoScalingGroupPara")
	delete(f, "LaunchConfigurePara")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "EnableAutoscale")
	delete(f, "Name")
	delete(f, "Labels")
	delete(f, "Taints")
	delete(f, "Annotations")
	delete(f, "ContainerRuntime")
	delete(f, "RuntimeVersion")
	delete(f, "NodePoolOs")
	delete(f, "OsCustomizeType")
	delete(f, "Tags")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterNodePoolResponseParams struct {
	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterNodePoolResponseParams `json:"Response"`
}

func (r *CreateClusterNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterReleaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称，最长63个字符，只能包含小写字母、数字及分隔符“-”，且必须以小写字母开头，数字或小写字母结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间，从集群详情命名空间获取
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 制品名称(从应用市场获取)或从第三方repo 安装chart时，制品压缩包下载地址, 不支持重定向类型chart 地址，结尾为*.tgz
	Chart *string `json:"Chart,omitnil,omitempty" name:"Chart"`

	// 自定义参数
	Values *ReleaseValues `json:"Values,omitnil,omitempty" name:"Values"`

	// 制品来源，范围：tke-market 或 other默认值：tke-market。
	ChartFrom *string `json:"ChartFrom,omitnil,omitempty" name:"ChartFrom"`

	// 制品版本
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`

	// 制品仓库URL地址
	ChartRepoURL *string `json:"ChartRepoURL,omitnil,omitempty" name:"ChartRepoURL"`

	// 制品访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 制品访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 制品命名空间，ChartFrom为tke-market时ChartNamespace不为空，值为DescribeProducts接口反馈的Namespace
	ChartNamespace *string `json:"ChartNamespace,omitnil,omitempty" name:"ChartNamespace"`

	// 集群类型，支持传 tke, eks, tkeedge, external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type CreateClusterReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称，最长63个字符，只能包含小写字母、数字及分隔符“-”，且必须以小写字母开头，数字或小写字母结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间，从集群详情命名空间获取
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 制品名称(从应用市场获取)或从第三方repo 安装chart时，制品压缩包下载地址, 不支持重定向类型chart 地址，结尾为*.tgz
	Chart *string `json:"Chart,omitnil,omitempty" name:"Chart"`

	// 自定义参数
	Values *ReleaseValues `json:"Values,omitnil,omitempty" name:"Values"`

	// 制品来源，范围：tke-market 或 other默认值：tke-market。
	ChartFrom *string `json:"ChartFrom,omitnil,omitempty" name:"ChartFrom"`

	// 制品版本
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`

	// 制品仓库URL地址
	ChartRepoURL *string `json:"ChartRepoURL,omitnil,omitempty" name:"ChartRepoURL"`

	// 制品访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 制品访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 制品命名空间，ChartFrom为tke-market时ChartNamespace不为空，值为DescribeProducts接口反馈的Namespace
	ChartNamespace *string `json:"ChartNamespace,omitnil,omitempty" name:"ChartNamespace"`

	// 集群类型，支持传 tke, eks, tkeedge, external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *CreateClusterReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "Chart")
	delete(f, "Values")
	delete(f, "ChartFrom")
	delete(f, "ChartVersion")
	delete(f, "ChartRepoURL")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "ChartNamespace")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterReleaseResponseParams struct {
	// 应用详情
	Release *PendingRelease `json:"Release,omitnil,omitempty" name:"Release"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterReleaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterReleaseResponseParams `json:"Response"`
}

func (r *CreateClusterReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群容器网络配置信息
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitnil,omitempty" name:"ClusterCIDRSettings"`

	// CVM创建透传参数，json化字符串格式，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口。总机型(包括地域)数量不超过10个，相同机型(地域)购买多台机器可以通过设置参数中RunInstances中InstanceCount来实现。
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitnil,omitempty" name:"RunInstancesForNode"`

	// 集群的基本配置信息
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitnil,omitempty" name:"ClusterBasicSettings"`

	// 集群高级配置信息
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitnil,omitempty" name:"ClusterAdvancedSettings"`

	// 节点高级配置信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 已存在实例的配置信息。所有实例必须在同一个VPC中，最大数量不超过100，不支持添加竞价实例。
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitnil,omitempty" name:"ExistedInstancesForNode"`

	// CVM类型和其对应的数据盘挂载配置信息
	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitnil,omitempty" name:"InstanceDataDiskMountSettings"`

	// 需要安装的扩展组件信息
	ExtensionAddons []*ExtensionAddon `json:"ExtensionAddons,omitnil,omitempty" name:"ExtensionAddons"`

	// 本地专用集群Id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 屏蔽安装指定Addon组件，填写相应的AddonName
	DisableAddons []*string `json:"DisableAddons,omitnil,omitempty" name:"DisableAddons"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群容器网络配置信息
	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitnil,omitempty" name:"ClusterCIDRSettings"`

	// CVM创建透传参数，json化字符串格式，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口。总机型(包括地域)数量不超过10个，相同机型(地域)购买多台机器可以通过设置参数中RunInstances中InstanceCount来实现。
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitnil,omitempty" name:"RunInstancesForNode"`

	// 集群的基本配置信息
	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitnil,omitempty" name:"ClusterBasicSettings"`

	// 集群高级配置信息
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitnil,omitempty" name:"ClusterAdvancedSettings"`

	// 节点高级配置信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 已存在实例的配置信息。所有实例必须在同一个VPC中，最大数量不超过100，不支持添加竞价实例。
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitnil,omitempty" name:"ExistedInstancesForNode"`

	// CVM类型和其对应的数据盘挂载配置信息
	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitnil,omitempty" name:"InstanceDataDiskMountSettings"`

	// 需要安装的扩展组件信息
	ExtensionAddons []*ExtensionAddon `json:"ExtensionAddons,omitnil,omitempty" name:"ExtensionAddons"`

	// 本地专用集群Id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 屏蔽安装指定Addon组件，填写相应的AddonName
	DisableAddons []*string `json:"DisableAddons,omitnil,omitempty" name:"DisableAddons"`
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
	delete(f, "ClusterType")
	delete(f, "ClusterCIDRSettings")
	delete(f, "RunInstancesForNode")
	delete(f, "ClusterBasicSettings")
	delete(f, "ClusterAdvancedSettings")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "ExistedInstancesForNode")
	delete(f, "InstanceDataDiskMountSettings")
	delete(f, "ExtensionAddons")
	delete(f, "CdcId")
	delete(f, "DisableAddons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateClusterRouteRequestParams struct {
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 目的节点的 PodCIDR
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// 下一跳地址，即目的节点的内网 IP 地址
	GatewayIp *string `json:"GatewayIp,omitnil,omitempty" name:"GatewayIp"`
}

type CreateClusterRouteRequest struct {
	*tchttp.BaseRequest
	
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 目的节点的 PodCIDR
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// 下一跳地址，即目的节点的内网 IP 地址
	GatewayIp *string `json:"GatewayIp,omitnil,omitempty" name:"GatewayIp"`
}

func (r *CreateClusterRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "DestinationCidrBlock")
	delete(f, "GatewayIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterRouteResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterRouteResponseParams `json:"Response"`
}

func (r *CreateClusterRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRouteTableRequestParams struct {
	// 路由表名称，一般为集群ID
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 路由表CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitnil,omitempty" name:"RouteTableCidrBlock"`

	// 路由表绑定的VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 是否忽略CIDR与 vpc 路由表的冲突， 0 表示不忽略，1表示忽略。默认不忽略
	IgnoreClusterCidrConflict *int64 `json:"IgnoreClusterCidrConflict,omitnil,omitempty" name:"IgnoreClusterCidrConflict"`
}

type CreateClusterRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// 路由表名称，一般为集群ID
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 路由表CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitnil,omitempty" name:"RouteTableCidrBlock"`

	// 路由表绑定的VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 是否忽略CIDR与 vpc 路由表的冲突， 0 表示不忽略，1表示忽略。默认不忽略
	IgnoreClusterCidrConflict *int64 `json:"IgnoreClusterCidrConflict,omitnil,omitempty" name:"IgnoreClusterCidrConflict"`
}

func (r *CreateClusterRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "RouteTableCidrBlock")
	delete(f, "VpcId")
	delete(f, "IgnoreClusterCidrConflict")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterRouteTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterRouteTableResponseParams `json:"Response"`
}

func (r *CreateClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterVirtualNodePoolRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安全组ID列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 虚拟节点label
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 虚拟节点taint
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点列表
	VirtualNodes []*VirtualNodeSpec `json:"VirtualNodes,omitnil,omitempty" name:"VirtualNodes"`

	// 删除保护开关，默认关闭
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 节点池操作系统：
	// - linux（默认）
	// - windows
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`
}

type CreateClusterVirtualNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安全组ID列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 虚拟节点label
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 虚拟节点taint
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点列表
	VirtualNodes []*VirtualNodeSpec `json:"VirtualNodes,omitnil,omitempty" name:"VirtualNodes"`

	// 删除保护开关，默认关闭
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 节点池操作系统：
	// - linux（默认）
	// - windows
	OS *string `json:"OS,omitnil,omitempty" name:"OS"`
}

func (r *CreateClusterVirtualNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterVirtualNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "SecurityGroupIds")
	delete(f, "SubnetIds")
	delete(f, "Labels")
	delete(f, "Taints")
	delete(f, "VirtualNodes")
	delete(f, "DeletionProtection")
	delete(f, "OS")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterVirtualNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterVirtualNodePoolResponseParams struct {
	// 节点池ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterVirtualNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterVirtualNodePoolResponseParams `json:"Response"`
}

func (r *CreateClusterVirtualNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterVirtualNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterVirtualNodeRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 虚拟节点所属节点池，通过DescribeNodePools接口获取
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 虚拟节点所属子网，SubnetId、SubnetIds、VirtualNodes必选一个。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 虚拟节点子网ID列表，SubnetId、SubnetIds、VirtualNodes必选一个。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 虚拟节点列表，SubnetId、SubnetIds、VirtualNodes必选一个。
	VirtualNodes []*VirtualNodeSpec `json:"VirtualNodes,omitnil,omitempty" name:"VirtualNodes"`
}

type CreateClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 虚拟节点所属节点池，通过DescribeNodePools接口获取
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 虚拟节点所属子网，SubnetId、SubnetIds、VirtualNodes必选一个。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 虚拟节点子网ID列表，SubnetId、SubnetIds、VirtualNodes必选一个。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 虚拟节点列表，SubnetId、SubnetIds、VirtualNodes必选一个。
	VirtualNodes []*VirtualNodeSpec `json:"VirtualNodes,omitnil,omitempty" name:"VirtualNodes"`
}

func (r *CreateClusterVirtualNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterVirtualNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "SubnetId")
	delete(f, "SubnetIds")
	delete(f, "VirtualNodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterVirtualNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterVirtualNodeResponseParams struct {
	// 虚拟节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterVirtualNodeResponseParams `json:"Response"`
}

func (r *CreateClusterVirtualNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateECMInstancesRequestParams struct {
	// 集群id，边缘集群需要先开启公网访问才能添加ecm节点
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 边缘模块id
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 需要创建实例的可用区及创建数目及运营商的列表
	ZoneInstanceCountISPSet []*ECMZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitnil,omitempty" name:"ZoneInstanceCountISPSet"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 公网带宽
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 主机名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 增强服务，包括云镜和云监控
	EnhancedService *ECMEnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用户自定义脚本
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 实例扩展信息
	External *string `json:"External,omitnil,omitempty" name:"External"`

	// 实例所属安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type CreateECMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id，边缘集群需要先开启公网访问才能添加ecm节点
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 边缘模块id
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 需要创建实例的可用区及创建数目及运营商的列表
	ZoneInstanceCountISPSet []*ECMZoneInstanceCountISP `json:"ZoneInstanceCountISPSet,omitnil,omitempty" name:"ZoneInstanceCountISPSet"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 公网带宽
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 主机名称
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 增强服务，包括云镜和云监控
	EnhancedService *ECMEnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 用户自定义脚本
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 实例扩展信息
	External *string `json:"External,omitnil,omitempty" name:"External"`

	// 实例所属安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *CreateECMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateECMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "ModuleId")
	delete(f, "ZoneInstanceCountISPSet")
	delete(f, "Password")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "ImageId")
	delete(f, "InstanceName")
	delete(f, "HostName")
	delete(f, "EnhancedService")
	delete(f, "UserData")
	delete(f, "External")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateECMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateECMInstancesResponseParams struct {
	// ecm id 列表
	EcmIdSet []*string `json:"EcmIdSet,omitnil,omitempty" name:"EcmIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateECMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateECMInstancesResponseParams `json:"Response"`
}

func (r *CreateECMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateECMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEKSClusterRequestParams struct {
	// k8s版本号。可为1.18.4 1.20.6。
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// vpc 的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 子网Id 列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// Service CIDR 或 Serivce 所在子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// 集群自定义的Dns服务器信息
	DnsServers []*DnsServerConf `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// 扩展参数。须是map[string]string 的json 格式。
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 是否在用户集群内开启Dns。默认为true
	EnableVpcCoreDNS *bool `json:"EnableVpcCoreDNS,omitnil,omitempty" name:"EnableVpcCoreDNS"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到集群实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 子网信息列表
	SubnetInfos []*SubnetInfos `json:"SubnetInfos,omitnil,omitempty" name:"SubnetInfos"`
}

type CreateEKSClusterRequest struct {
	*tchttp.BaseRequest
	
	// k8s版本号。可为1.18.4 1.20.6。
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// vpc 的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 子网Id 列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// Service CIDR 或 Serivce 所在子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// 集群自定义的Dns服务器信息
	DnsServers []*DnsServerConf `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// 扩展参数。须是map[string]string 的json 格式。
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`

	// 是否在用户集群内开启Dns。默认为true
	EnableVpcCoreDNS *bool `json:"EnableVpcCoreDNS,omitnil,omitempty" name:"EnableVpcCoreDNS"`

	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到集群实例。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`

	// 子网信息列表
	SubnetInfos []*SubnetInfos `json:"SubnetInfos,omitnil,omitempty" name:"SubnetInfos"`
}

func (r *CreateEKSClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEKSClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "K8SVersion")
	delete(f, "VpcId")
	delete(f, "ClusterName")
	delete(f, "SubnetIds")
	delete(f, "ClusterDesc")
	delete(f, "ServiceSubnetId")
	delete(f, "DnsServers")
	delete(f, "ExtraParam")
	delete(f, "EnableVpcCoreDNS")
	delete(f, "TagSpecification")
	delete(f, "SubnetInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEKSClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEKSClusterResponseParams struct {
	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEKSClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateEKSClusterResponseParams `json:"Response"`
}

func (r *CreateEKSClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEKSClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEKSContainerInstancesRequestParams struct {
	// 容器组
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// EKS Container Instance容器实例名称
	EksCiName *string `json:"EksCiName,omitnil,omitempty" name:"EksCiName"`

	// 指定新创建实例所属于的安全组Id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例所属子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例所属VPC的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 内存，单位：GiB。可参考[资源规格](https://cloud.tencent.com/document/product/457/39808)文档
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CPU，单位：核。可参考[资源规格](https://cloud.tencent.com/document/product/457/39808)文档
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例重启策略： Always(总是重启)、Never(从不重启)、OnFailure(失败时重启)，默认：Always。
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 数据卷，包含NfsVolume数组和CbsVolume数组
	EksCiVolume *EksCiVolume `json:"EksCiVolume,omitnil,omitempty" name:"EksCiVolume"`

	// 实例副本数，默认为1
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// Init 容器
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 自定义DNS配置
	DnsConfig *DNSConfig `json:"DnsConfig,omitnil,omitempty" name:"DnsConfig"`

	// 用来绑定容器实例的已有EIP的列表。如传值，需要保证数值和Replicas相等。
	// 另外此参数和AutoCreateEipAttribute互斥。
	ExistedEipIds []*string `json:"ExistedEipIds,omitnil,omitempty" name:"ExistedEipIds"`

	// 自动创建EIP的可选参数。若传此参数，则会自动创建EIP。
	// 另外此参数和ExistedEipIds互斥
	AutoCreateEipAttribute *EipAttribute `json:"AutoCreateEipAttribute,omitnil,omitempty" name:"AutoCreateEipAttribute"`

	// 是否为容器实例自动创建EIP，默认为false。若传true，则此参数和ExistedEipIds互斥
	AutoCreateEip *bool `json:"AutoCreateEip,omitnil,omitempty" name:"AutoCreateEip"`

	// Pod 所需的 CPU 资源型号，如果不填写则默认不强制指定 CPU 类型。目前支持型号如下：
	// intel
	// amd
	// - 支持优先级顺序写法，如 “amd,intel” 表示优先创建 amd 资源 Pod，如果所选地域可用区 amd 资源不足，则会创建 intel 资源 Pod。
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 容器实例所需的 GPU 资源型号，目前支持型号如下：
	// 1/4\*V100
	// 1/2\*V100
	// V100
	// 1/4\*T4
	// 1/2\*T4
	// T4
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// Pod 所需的 GPU 数量，如填写，请确保为支持的规格。默认单位为卡，无需再次注明。
	GpuCount *uint64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// 为容器实例关联 CAM 角色，value 填写 CAM 角色名称，容器实例可获取该 CAM 角色包含的权限策略，方便 容器实例 内的程序进行如购买资源、读写存储等云资源操作。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`
}

type CreateEKSContainerInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 容器组
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// EKS Container Instance容器实例名称
	EksCiName *string `json:"EksCiName,omitnil,omitempty" name:"EksCiName"`

	// 指定新创建实例所属于的安全组Id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例所属子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例所属VPC的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 内存，单位：GiB。可参考[资源规格](https://cloud.tencent.com/document/product/457/39808)文档
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// CPU，单位：核。可参考[资源规格](https://cloud.tencent.com/document/product/457/39808)文档
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例重启策略： Always(总是重启)、Never(从不重启)、OnFailure(失败时重启)，默认：Always。
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 数据卷，包含NfsVolume数组和CbsVolume数组
	EksCiVolume *EksCiVolume `json:"EksCiVolume,omitnil,omitempty" name:"EksCiVolume"`

	// 实例副本数，默认为1
	Replicas *int64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// Init 容器
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 自定义DNS配置
	DnsConfig *DNSConfig `json:"DnsConfig,omitnil,omitempty" name:"DnsConfig"`

	// 用来绑定容器实例的已有EIP的列表。如传值，需要保证数值和Replicas相等。
	// 另外此参数和AutoCreateEipAttribute互斥。
	ExistedEipIds []*string `json:"ExistedEipIds,omitnil,omitempty" name:"ExistedEipIds"`

	// 自动创建EIP的可选参数。若传此参数，则会自动创建EIP。
	// 另外此参数和ExistedEipIds互斥
	AutoCreateEipAttribute *EipAttribute `json:"AutoCreateEipAttribute,omitnil,omitempty" name:"AutoCreateEipAttribute"`

	// 是否为容器实例自动创建EIP，默认为false。若传true，则此参数和ExistedEipIds互斥
	AutoCreateEip *bool `json:"AutoCreateEip,omitnil,omitempty" name:"AutoCreateEip"`

	// Pod 所需的 CPU 资源型号，如果不填写则默认不强制指定 CPU 类型。目前支持型号如下：
	// intel
	// amd
	// - 支持优先级顺序写法，如 “amd,intel” 表示优先创建 amd 资源 Pod，如果所选地域可用区 amd 资源不足，则会创建 intel 资源 Pod。
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 容器实例所需的 GPU 资源型号，目前支持型号如下：
	// 1/4\*V100
	// 1/2\*V100
	// V100
	// 1/4\*T4
	// 1/2\*T4
	// T4
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// Pod 所需的 GPU 数量，如填写，请确保为支持的规格。默认单位为卡，无需再次注明。
	GpuCount *uint64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// 为容器实例关联 CAM 角色，value 填写 CAM 角色名称，容器实例可获取该 CAM 角色包含的权限策略，方便 容器实例 内的程序进行如购买资源、读写存储等云资源操作。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`
}

func (r *CreateEKSContainerInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEKSContainerInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Containers")
	delete(f, "EksCiName")
	delete(f, "SecurityGroupIds")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "Memory")
	delete(f, "Cpu")
	delete(f, "RestartPolicy")
	delete(f, "ImageRegistryCredentials")
	delete(f, "EksCiVolume")
	delete(f, "Replicas")
	delete(f, "InitContainers")
	delete(f, "DnsConfig")
	delete(f, "ExistedEipIds")
	delete(f, "AutoCreateEipAttribute")
	delete(f, "AutoCreateEip")
	delete(f, "CpuType")
	delete(f, "GpuType")
	delete(f, "GpuCount")
	delete(f, "CamRoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEKSContainerInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEKSContainerInstancesResponseParams struct {
	// EKS Container Instance Id集合，格式为eksci-xxx，是容器实例的唯一标识。
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEKSContainerInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateEKSContainerInstancesResponseParams `json:"Response"`
}

func (r *CreateEKSContainerInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEKSContainerInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeCVMInstancesRequestParams struct {
	// 集群id，边缘集群需要先开启公网访问才能添加cvm节点
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// CVM创建透传参数，json化字符串格式，如需要保证扩展集群节点请求幂等性需要在此参数添加ClientToken字段，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口。
	RunInstancePara *string `json:"RunInstancePara,omitnil,omitempty" name:"RunInstancePara"`

	// CVM所属Region
	CvmRegion *string `json:"CvmRegion,omitnil,omitempty" name:"CvmRegion"`

	// CVM数量
	CvmCount *int64 `json:"CvmCount,omitnil,omitempty" name:"CvmCount"`

	// 实例扩展信息
	External *string `json:"External,omitnil,omitempty" name:"External"`

	// 用户自定义脚本
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 是否开启弹性网卡功能
	EnableEni *bool `json:"EnableEni,omitnil,omitempty" name:"EnableEni"`
}

type CreateEdgeCVMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id，边缘集群需要先开启公网访问才能添加cvm节点
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// CVM创建透传参数，json化字符串格式，如需要保证扩展集群节点请求幂等性需要在此参数添加ClientToken字段，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口。
	RunInstancePara *string `json:"RunInstancePara,omitnil,omitempty" name:"RunInstancePara"`

	// CVM所属Region
	CvmRegion *string `json:"CvmRegion,omitnil,omitempty" name:"CvmRegion"`

	// CVM数量
	CvmCount *int64 `json:"CvmCount,omitnil,omitempty" name:"CvmCount"`

	// 实例扩展信息
	External *string `json:"External,omitnil,omitempty" name:"External"`

	// 用户自定义脚本
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 是否开启弹性网卡功能
	EnableEni *bool `json:"EnableEni,omitnil,omitempty" name:"EnableEni"`
}

func (r *CreateEdgeCVMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeCVMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "RunInstancePara")
	delete(f, "CvmRegion")
	delete(f, "CvmCount")
	delete(f, "External")
	delete(f, "UserScript")
	delete(f, "EnableEni")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeCVMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeCVMInstancesResponseParams struct {
	// cvm id 列表
	CvmIdSet []*string `json:"CvmIdSet,omitnil,omitempty" name:"CvmIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeCVMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeCVMInstancesResponseParams `json:"Response"`
}

func (r *CreateEdgeCVMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeCVMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeLogConfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 日志采集配置的json表达
	LogConfig *string `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// CLS日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type CreateEdgeLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 日志采集配置的json表达
	LogConfig *string `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// CLS日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

func (r *CreateEdgeLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LogConfig")
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeLogConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeLogConfigResponseParams `json:"Response"`
}

func (r *CreateEdgeLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEksLogConfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 日志采集配置的json表达
	LogConfig *string `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

type CreateEksLogConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 日志采集配置的json表达
	LogConfig *string `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`
}

func (r *CreateEksLogConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEksLogConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LogConfig")
	delete(f, "LogsetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEksLogConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEksLogConfigResponseParams struct {
	// 日志采集topicid
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEksLogConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateEksLogConfigResponseParams `json:"Response"`
}

func (r *CreateEksLogConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEksLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalMaintenanceWindowAndExclusionsRequestParams struct {
	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长(小时)
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 地域
	TargetRegions []*string `json:"TargetRegions,omitnil,omitempty" name:"TargetRegions"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

type CreateGlobalMaintenanceWindowAndExclusionsRequest struct {
	*tchttp.BaseRequest
	
	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长(小时)
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 地域
	TargetRegions []*string `json:"TargetRegions,omitnil,omitempty" name:"TargetRegions"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

func (r *CreateGlobalMaintenanceWindowAndExclusionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalMaintenanceWindowAndExclusionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaintenanceTime")
	delete(f, "Duration")
	delete(f, "DayOfWeek")
	delete(f, "TargetRegions")
	delete(f, "Exclusions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlobalMaintenanceWindowAndExclusionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlobalMaintenanceWindowAndExclusionsResponseParams struct {
	// 维护窗口ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlobalMaintenanceWindowAndExclusionsResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlobalMaintenanceWindowAndExclusionsResponseParams `json:"Response"`
}

func (r *CreateGlobalMaintenanceWindowAndExclusionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlobalMaintenanceWindowAndExclusionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageCacheRequestParams struct {
	// 用于制作镜像缓存的容器镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// 实例所属子网 ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例所属 VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 镜像缓存名称
	ImageCacheName *string `json:"ImageCacheName,omitnil,omitempty" name:"ImageCacheName"`

	// 安全组 ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 用来绑定容器实例的已有EIP
	ExistedEipId *string `json:"ExistedEipId,omitnil,omitempty" name:"ExistedEipId"`

	// 是否为容器实例自动创建EIP，默认为false。若传true，则此参数和ExistedEipIds互斥
	AutoCreateEip *bool `json:"AutoCreateEip,omitnil,omitempty" name:"AutoCreateEip"`

	// 自动创建EIP的可选参数。若传此参数，则会自动创建EIP。
	// 另外此参数和ExistedEipIds互斥
	AutoCreateEipAttribute *EipAttribute `json:"AutoCreateEipAttribute,omitnil,omitempty" name:"AutoCreateEipAttribute"`

	// 镜像缓存的大小。默认为20 GiB。取值范围参考[云硬盘类型](https://cloud.tencent.com/document/product/362/2353)中的高性能云盘类型的大小限制。
	ImageCacheSize *uint64 `json:"ImageCacheSize,omitnil,omitempty" name:"ImageCacheSize"`

	// 镜像缓存保留时间天数，过期将会自动清理，默认为0，永不过期。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 指定拉取镜像仓库的镜像时不校验证书。如["harbor.example.com"]。
	RegistrySkipVerifyList []*string `json:"RegistrySkipVerifyList,omitnil,omitempty" name:"RegistrySkipVerifyList"`

	// 指定拉取镜像仓库的镜像时使用 HTTP 协议。如["harbor.example.com"]。
	RegistryHttpEndPointList []*string `json:"RegistryHttpEndPointList,omitnil,omitempty" name:"RegistryHttpEndPointList"`

	// 自定义制作镜像缓存过程中容器实例的宿主机上的 DNS。如：
	// "nameserver 4.4.4.4\nnameserver 8.8.8.8"
	ResolveConfig *string `json:"ResolveConfig,omitnil,omitempty" name:"ResolveConfig"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateImageCacheRequest struct {
	*tchttp.BaseRequest
	
	// 用于制作镜像缓存的容器镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// 实例所属子网 ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例所属 VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 镜像缓存名称
	ImageCacheName *string `json:"ImageCacheName,omitnil,omitempty" name:"ImageCacheName"`

	// 安全组 ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 用来绑定容器实例的已有EIP
	ExistedEipId *string `json:"ExistedEipId,omitnil,omitempty" name:"ExistedEipId"`

	// 是否为容器实例自动创建EIP，默认为false。若传true，则此参数和ExistedEipIds互斥
	AutoCreateEip *bool `json:"AutoCreateEip,omitnil,omitempty" name:"AutoCreateEip"`

	// 自动创建EIP的可选参数。若传此参数，则会自动创建EIP。
	// 另外此参数和ExistedEipIds互斥
	AutoCreateEipAttribute *EipAttribute `json:"AutoCreateEipAttribute,omitnil,omitempty" name:"AutoCreateEipAttribute"`

	// 镜像缓存的大小。默认为20 GiB。取值范围参考[云硬盘类型](https://cloud.tencent.com/document/product/362/2353)中的高性能云盘类型的大小限制。
	ImageCacheSize *uint64 `json:"ImageCacheSize,omitnil,omitempty" name:"ImageCacheSize"`

	// 镜像缓存保留时间天数，过期将会自动清理，默认为0，永不过期。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 指定拉取镜像仓库的镜像时不校验证书。如["harbor.example.com"]。
	RegistrySkipVerifyList []*string `json:"RegistrySkipVerifyList,omitnil,omitempty" name:"RegistrySkipVerifyList"`

	// 指定拉取镜像仓库的镜像时使用 HTTP 协议。如["harbor.example.com"]。
	RegistryHttpEndPointList []*string `json:"RegistryHttpEndPointList,omitnil,omitempty" name:"RegistryHttpEndPointList"`

	// 自定义制作镜像缓存过程中容器实例的宿主机上的 DNS。如：
	// "nameserver 4.4.4.4\nnameserver 8.8.8.8"
	ResolveConfig *string `json:"ResolveConfig,omitnil,omitempty" name:"ResolveConfig"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Images")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "ImageCacheName")
	delete(f, "SecurityGroupIds")
	delete(f, "ImageRegistryCredentials")
	delete(f, "ExistedEipId")
	delete(f, "AutoCreateEip")
	delete(f, "AutoCreateEipAttribute")
	delete(f, "ImageCacheSize")
	delete(f, "RetentionDays")
	delete(f, "RegistrySkipVerifyList")
	delete(f, "RegistryHttpEndPointList")
	delete(f, "ResolveConfig")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImageCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImageCacheResponseParams struct {
	// 镜像缓存Id
	ImageCacheId *string `json:"ImageCacheId,omitnil,omitempty" name:"ImageCacheId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImageCacheResponse struct {
	*tchttp.BaseResponse
	Response *CreateImageCacheResponseParams `json:"Response"`
}

func (r *CreateImageCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

type CreatePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

func (r *CreatePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertPolicyResponseParams struct {
	// 告警id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *CreatePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertRuleRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

type CreatePrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

func (r *CreatePrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusAlertRuleResponseParams struct {
	// 告警id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *CreatePrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusClusterAgentRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// agent列表
	Agents []*PrometheusClusterAgentBasic `json:"Agents,omitnil,omitempty" name:"Agents"`
}

type CreatePrometheusClusterAgentRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// agent列表
	Agents []*PrometheusClusterAgentBasic `json:"Agents,omitnil,omitempty" name:"Agents"`
}

func (r *CreatePrometheusClusterAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusClusterAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Agents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusClusterAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusClusterAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusClusterAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusClusterAgentResponseParams `json:"Response"`
}

func (r *CreatePrometheusClusterAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusClusterAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Probe 配置
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`
}

type CreatePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Probe 配置
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`
}

func (r *CreatePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	delete(f, "Probes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusConfigResponseParams `json:"Response"`
}

func (r *CreatePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusDashboardRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 面板组名称
	DashboardName *string `json:"DashboardName,omitnil,omitempty" name:"DashboardName"`

	// 面板列表
	// 每一项是一个grafana dashboard的json定义
	Contents []*string `json:"Contents,omitnil,omitempty" name:"Contents"`
}

type CreatePrometheusDashboardRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 面板组名称
	DashboardName *string `json:"DashboardName,omitnil,omitempty" name:"DashboardName"`

	// 面板列表
	// 每一项是一个grafana dashboard的json定义
	Contents []*string `json:"Contents,omitnil,omitempty" name:"Contents"`
}

func (r *CreatePrometheusDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DashboardName")
	delete(f, "Contents")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusDashboardResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusDashboardResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusDashboardResponseParams `json:"Response"`
}

func (r *CreatePrometheusDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusGlobalNotificationRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

type CreatePrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

func (r *CreatePrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Notification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusGlobalNotificationResponseParams struct {
	// 全局告警通知渠道ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *CreatePrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// yaml的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CreatePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// yaml的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *CreatePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *CreatePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTempRequestParams struct {
	// 模板设置
	Template *PrometheusTemp `json:"Template,omitnil,omitempty" name:"Template"`
}

type CreatePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模板设置
	Template *PrometheusTemp `json:"Template,omitnil,omitempty" name:"Template"`
}

func (r *CreatePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTempResponseParams struct {
	// 模板Id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusTempResponseParams `json:"Response"`
}

func (r *CreatePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTemplateRequestParams struct {
	// 模板设置
	Template *PrometheusTemplate `json:"Template,omitnil,omitempty" name:"Template"`
}

type CreatePrometheusTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板设置
	Template *PrometheusTemplate `json:"Template,omitnil,omitempty" name:"Template"`
}

func (r *CreatePrometheusTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrometheusTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrometheusTemplateResponseParams struct {
	// 模板Id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePrometheusTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrometheusTemplateResponseParams `json:"Response"`
}

func (r *CreatePrometheusTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrometheusTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReservedInstancesRequestParams struct {
	// 预留券实例规格。
	ReservedInstanceSpec *ReservedInstanceSpec `json:"ReservedInstanceSpec,omitnil,omitempty" name:"ReservedInstanceSpec"`

	// 购买实例数量，一次最大购买数量为300。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 预留券名称，名称不得超过60个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。
	// true：只预检，不会创建实例。默认值为：false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type CreateReservedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 预留券实例规格。
	ReservedInstanceSpec *ReservedInstanceSpec `json:"ReservedInstanceSpec,omitnil,omitempty" name:"ReservedInstanceSpec"`

	// 购买实例数量，一次最大购买数量为300。
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 预留券名称，名称不得超过60个字符。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。
	// true：只预检，不会创建实例。默认值为：false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *CreateReservedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReservedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReservedInstanceSpec")
	delete(f, "InstanceCount")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceName")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReservedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReservedInstancesResponseParams struct {
	// 预留券实例 ID。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReservedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateReservedInstancesResponseParams `json:"Response"`
}

func (r *CreateReservedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReservedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRollOutSequenceRequestParams struct {
	// 发布序列名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布序列步骤
	SequenceFlows []*SequenceFlow `json:"SequenceFlows,omitnil,omitempty" name:"SequenceFlows"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CreateRollOutSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 发布序列名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布序列步骤
	SequenceFlows []*SequenceFlow `json:"SequenceFlows,omitnil,omitempty" name:"SequenceFlows"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *CreateRollOutSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRollOutSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SequenceFlows")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRollOutSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRollOutSequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRollOutSequenceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRollOutSequenceResponseParams `json:"Response"`
}

func (r *CreateRollOutSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRollOutSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTKEEdgeClusterRequestParams struct {
	// k8s版本号
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// vpc 的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群pod cidr
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群service cidr
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群高级设置
	ClusterAdvancedSettings *EdgeClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitnil,omitempty" name:"ClusterAdvancedSettings"`

	// 节点上最大Pod数量
	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitnil,omitempty" name:"MaxNodePodNum"`

	// 边缘计算集群公网访问LB信息
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 集群的级别
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 集群是否支持自动升配
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 集群计费方式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 边缘集群版本，此版本区别于k8s版本，是整个集群各组件版本集合
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`

	// 边缘组件镜像仓库前缀
	RegistryPrefix *string `json:"RegistryPrefix,omitnil,omitempty" name:"RegistryPrefix"`

	// 集群绑定的云标签
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type CreateTKEEdgeClusterRequest struct {
	*tchttp.BaseRequest
	
	// k8s版本号
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// vpc 的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群pod cidr
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群service cidr
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群高级设置
	ClusterAdvancedSettings *EdgeClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitnil,omitempty" name:"ClusterAdvancedSettings"`

	// 节点上最大Pod数量
	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitnil,omitempty" name:"MaxNodePodNum"`

	// 边缘计算集群公网访问LB信息
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 集群的级别
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 集群是否支持自动升配
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 集群计费方式
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 边缘集群版本，此版本区别于k8s版本，是整个集群各组件版本集合
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`

	// 边缘组件镜像仓库前缀
	RegistryPrefix *string `json:"RegistryPrefix,omitnil,omitempty" name:"RegistryPrefix"`

	// 集群绑定的云标签
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

func (r *CreateTKEEdgeClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTKEEdgeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "K8SVersion")
	delete(f, "VpcId")
	delete(f, "ClusterName")
	delete(f, "PodCIDR")
	delete(f, "ServiceCIDR")
	delete(f, "ClusterDesc")
	delete(f, "ClusterAdvancedSettings")
	delete(f, "MaxNodePodNum")
	delete(f, "PublicLB")
	delete(f, "ClusterLevel")
	delete(f, "AutoUpgradeClusterLevel")
	delete(f, "ChargeType")
	delete(f, "EdgeVersion")
	delete(f, "RegistryPrefix")
	delete(f, "TagSpecification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTKEEdgeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTKEEdgeClusterResponseParams struct {
	// 边缘计算集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTKEEdgeClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateTKEEdgeClusterResponseParams `json:"Response"`
}

func (r *CreateTKEEdgeClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTKEEdgeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomDriver struct {
	// 自定义GPU驱动地址链接
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`
}

type DNSConfig struct {
	// DNS 服务器IP地址列表
	Nameservers []*string `json:"Nameservers,omitnil,omitempty" name:"Nameservers"`

	// 对象选项列表，每个对象由name和value（可选）构成
	Options []*DNSConfigOption `json:"Options,omitnil,omitempty" name:"Options"`

	// DNS搜索域列表
	Searches []*string `json:"Searches,omitnil,omitempty" name:"Searches"`
}

type DNSConfigOption struct {
	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DataDisk struct {
	// 云盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 文件系统(ext3/ext4/xfs)
	FileSystem *string `json:"FileSystem,omitnil,omitempty" name:"FileSystem"`

	// 云盘大小(G）
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 是否自动化格式盘并挂载
	AutoFormatAndMount *bool `json:"AutoFormatAndMount,omitnil,omitempty" name:"AutoFormatAndMount"`

	// 挂载目录
	MountTarget *string `json:"MountTarget,omitnil,omitempty" name:"MountTarget"`

	// 挂载设备名或分区名，当且仅当添加已有节点时需要
	DiskPartition *string `json:"DiskPartition,omitnil,omitempty" name:"DiskPartition"`
}

// Predefined struct for user
type DeleteAddonRequestParams struct {
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`
}

type DeleteAddonRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`
}

func (r *DeleteAddonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AddonName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddonResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAddonResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddonResponseParams `json:"Response"`
}

func (r *DeleteAddonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupStorageLocationRequestParams struct {
	// 备份仓库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteBackupStorageLocationRequest struct {
	*tchttp.BaseRequest
	
	// 备份仓库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteBackupStorageLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupStorageLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupStorageLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupStorageLocationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBackupStorageLocationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupStorageLocationResponseParams `json:"Response"`
}

func (r *DeleteBackupStorageLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupStorageLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterAsGroupsRequestParams struct {
	// 集群ID，通过[DescribeClusters](https://cloud.tencent.com/document/api/457/31862)接口获取。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群伸缩组ID的列表
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// 是否保留伸缩组中的节点(默认值： false(不保留))
	KeepInstance *bool `json:"KeepInstance,omitnil,omitempty" name:"KeepInstance"`
}

type DeleteClusterAsGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过[DescribeClusters](https://cloud.tencent.com/document/api/457/31862)接口获取。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群伸缩组ID的列表
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// 是否保留伸缩组中的节点(默认值： false(不保留))
	KeepInstance *bool `json:"KeepInstance,omitnil,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterAsGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoScalingGroupIds")
	delete(f, "KeepInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterAsGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterAsGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterAsGroupsResponseParams `json:"Response"`
}

func (r *DeleteClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`
}

type DeleteClusterEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`
}

func (r *DeleteClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterEndpointResponseParams `json:"Response"`
}

func (r *DeleteClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointVipRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteClusterEndpointVipRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterEndpointVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterEndpointVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterEndpointVipResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterEndpointVipResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterEndpointVipResponseParams `json:"Response"`
}

func (r *DeleteClusterEndpointVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterEndpointVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterInstancesRequestParams struct {
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主机 InstanceId 列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitnil,omitempty" name:"InstanceDeleteMode"`

	// 是否强制删除(当节点在初始化时，可以指定参数为TRUE)
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`

	// 集群删除时资源的删除策略，目前支持CBS（默认保留CBS）
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitnil,omitempty" name:"ResourceDeleteOptions"`
}

type DeleteClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主机 InstanceId 列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitnil,omitempty" name:"InstanceDeleteMode"`

	// 是否强制删除(当节点在初始化时，可以指定参数为TRUE)
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`

	// 集群删除时资源的删除策略，目前支持CBS（默认保留CBS）
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitnil,omitempty" name:"ResourceDeleteOptions"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "InstanceDeleteMode")
	delete(f, "ForceDelete")
	delete(f, "ResourceDeleteOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterInstancesResponseParams struct {
	// 删除成功的实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitnil,omitempty" name:"SuccInstanceIds"`

	// 删除失败的实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 未匹配到的实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotFoundInstanceIds []*string `json:"NotFoundInstanceIds,omitnil,omitempty" name:"NotFoundInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterInstancesResponseParams `json:"Response"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterMaintenanceWindowAndExclusionRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`
}

type DeleteClusterMaintenanceWindowAndExclusionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`
}

func (r *DeleteClusterMaintenanceWindowAndExclusionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterMaintenanceWindowAndExclusionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterMaintenanceWindowAndExclusionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterMaintenanceWindowAndExclusionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterMaintenanceWindowAndExclusionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterMaintenanceWindowAndExclusionResponseParams `json:"Response"`
}

func (r *DeleteClusterMaintenanceWindowAndExclusionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterMaintenanceWindowAndExclusionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterNodePoolRequestParams struct {
	// 节点池对应的 ClusterId
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要删除的节点池 Id 列表
	NodePoolIds []*string `json:"NodePoolIds,omitnil,omitempty" name:"NodePoolIds"`

	// 删除节点池时是否保留节点池内节点(节点仍然会被移出集群，但对应的实例不会被销毁)
	KeepInstance *bool `json:"KeepInstance,omitnil,omitempty" name:"KeepInstance"`
}

type DeleteClusterNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 节点池对应的 ClusterId
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要删除的节点池 Id 列表
	NodePoolIds []*string `json:"NodePoolIds,omitnil,omitempty" name:"NodePoolIds"`

	// 删除节点池时是否保留节点池内节点(节点仍然会被移出集群，但对应的实例不会被销毁)
	KeepInstance *bool `json:"KeepInstance,omitnil,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolIds")
	delete(f, "KeepInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterNodePoolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterNodePoolResponseParams `json:"Response"`
}

func (r *DeleteClusterNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitnil,omitempty" name:"InstanceDeleteMode"`

	// 集群删除时资源的删除策略，目前支持CBS（默认保留CBS）
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitnil,omitempty" name:"ResourceDeleteOptions"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitnil,omitempty" name:"InstanceDeleteMode"`

	// 集群删除时资源的删除策略，目前支持CBS（默认保留CBS）
	ResourceDeleteOptions []*ResourceDeleteOption `json:"ResourceDeleteOptions,omitnil,omitempty" name:"ResourceDeleteOptions"`
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
	delete(f, "InstanceDeleteMode")
	delete(f, "ResourceDeleteOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteClusterRouteRequestParams struct {
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 下一跳地址。
	GatewayIp *string `json:"GatewayIp,omitnil,omitempty" name:"GatewayIp"`

	// 目的端CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`
}

type DeleteClusterRouteRequest struct {
	*tchttp.BaseRequest
	
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 下一跳地址。
	GatewayIp *string `json:"GatewayIp,omitnil,omitempty" name:"GatewayIp"`

	// 目的端CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`
}

func (r *DeleteClusterRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "GatewayIp")
	delete(f, "DestinationCidrBlock")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterRouteResponseParams `json:"Response"`
}

func (r *DeleteClusterRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteTableRequestParams struct {
	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`
}

type DeleteClusterRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// 路由表名称
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`
}

func (r *DeleteClusterRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRouteTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterRouteTableResponseParams `json:"Response"`
}

func (r *DeleteClusterRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterVirtualNodePoolRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID，通过DescribeNodePools接口获取
	NodePoolIds []*string `json:"NodePoolIds,omitnil,omitempty" name:"NodePoolIds"`

	// 是否强制删除，在超级节点上有pod的情况下，如果选择非强制删除，则删除会失败
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteClusterVirtualNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID，通过DescribeNodePools接口获取
	NodePoolIds []*string `json:"NodePoolIds,omitnil,omitempty" name:"NodePoolIds"`

	// 是否强制删除，在超级节点上有pod的情况下，如果选择非强制删除，则删除会失败
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *DeleteClusterVirtualNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterVirtualNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolIds")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterVirtualNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterVirtualNodePoolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterVirtualNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterVirtualNodePoolResponseParams `json:"Response"`
}

func (r *DeleteClusterVirtualNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterVirtualNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterVirtualNodeRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 虚拟节点ID列表
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 是否强制删除：如果虚拟节点上有运行中Pod，则非强制删除状态下不会进行删除
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 虚拟节点ID列表
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 是否强制删除：如果虚拟节点上有运行中Pod，则非强制删除状态下不会进行删除
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *DeleteClusterVirtualNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterVirtualNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeNames")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterVirtualNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterVirtualNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterVirtualNodeResponseParams `json:"Response"`
}

func (r *DeleteClusterVirtualNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteECMInstancesRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// ecm id集合
	EcmIdSet []*string `json:"EcmIdSet,omitnil,omitempty" name:"EcmIdSet"`
}

type DeleteECMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// ecm id集合
	EcmIdSet []*string `json:"EcmIdSet,omitnil,omitempty" name:"EcmIdSet"`
}

func (r *DeleteECMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteECMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "EcmIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteECMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteECMInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteECMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteECMInstancesResponseParams `json:"Response"`
}

func (r *DeleteECMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteECMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEKSClusterRequestParams struct {
	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteEKSClusterRequest struct {
	*tchttp.BaseRequest
	
	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteEKSClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEKSClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEKSClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEKSClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEKSClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEKSClusterResponseParams `json:"Response"`
}

func (r *DeleteEKSClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEKSClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEKSContainerInstancesRequestParams struct {
	// 需要删除的EksCi的Id。 最大数量不超过20
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`

	// 是否释放为EksCi自动创建的Eip
	ReleaseAutoCreatedEip *bool `json:"ReleaseAutoCreatedEip,omitnil,omitempty" name:"ReleaseAutoCreatedEip"`
}

type DeleteEKSContainerInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的EksCi的Id。 最大数量不超过20
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`

	// 是否释放为EksCi自动创建的Eip
	ReleaseAutoCreatedEip *bool `json:"ReleaseAutoCreatedEip,omitnil,omitempty" name:"ReleaseAutoCreatedEip"`
}

func (r *DeleteEKSContainerInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEKSContainerInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EksCiIds")
	delete(f, "ReleaseAutoCreatedEip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEKSContainerInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEKSContainerInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEKSContainerInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEKSContainerInstancesResponseParams `json:"Response"`
}

func (r *DeleteEKSContainerInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEKSContainerInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeCVMInstancesRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// cvm id集合
	CvmIdSet []*string `json:"CvmIdSet,omitnil,omitempty" name:"CvmIdSet"`
}

type DeleteEdgeCVMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// cvm id集合
	CvmIdSet []*string `json:"CvmIdSet,omitnil,omitempty" name:"CvmIdSet"`
}

func (r *DeleteEdgeCVMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeCVMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "CvmIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeCVMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeCVMInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeCVMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeCVMInstancesResponseParams `json:"Response"`
}

func (r *DeleteEdgeCVMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeCVMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeClusterInstancesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 待删除实例ID数组
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DeleteEdgeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 待删除实例ID数组
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DeleteEdgeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeClusterInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeClusterInstancesResponseParams `json:"Response"`
}

func (r *DeleteEdgeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalMaintenanceWindowAndExclusionRequestParams struct {
	// 全局维护时间窗口ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteGlobalMaintenanceWindowAndExclusionRequest struct {
	*tchttp.BaseRequest
	
	// 全局维护时间窗口ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteGlobalMaintenanceWindowAndExclusionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalMaintenanceWindowAndExclusionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlobalMaintenanceWindowAndExclusionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlobalMaintenanceWindowAndExclusionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlobalMaintenanceWindowAndExclusionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlobalMaintenanceWindowAndExclusionResponseParams `json:"Response"`
}

func (r *DeleteGlobalMaintenanceWindowAndExclusionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlobalMaintenanceWindowAndExclusionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageCachesRequestParams struct {
	// 镜像缓存ID数组
	ImageCacheIds []*string `json:"ImageCacheIds,omitnil,omitempty" name:"ImageCacheIds"`
}

type DeleteImageCachesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像缓存ID数组
	ImageCacheIds []*string `json:"ImageCacheIds,omitnil,omitempty" name:"ImageCacheIds"`
}

func (r *DeleteImageCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageCachesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageCacheIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteImageCachesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteImageCachesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteImageCachesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteImageCachesResponseParams `json:"Response"`
}

func (r *DeleteImageCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteImageCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogConfigsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 待删除采集规则名称，多个采集规则使用","分隔
	LogConfigNames *string `json:"LogConfigNames,omitnil,omitempty" name:"LogConfigNames"`

	// 集群集群类型, tke/eks 默认为 tke 集群
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DeleteLogConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 待删除采集规则名称，多个采集规则使用","分隔
	LogConfigNames *string `json:"LogConfigNames,omitnil,omitempty" name:"LogConfigNames"`

	// 集群集群类型, tke/eks 默认为 tke 集群
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DeleteLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LogConfigNames")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogConfigsResponseParams struct {
	// 删除采集规则遇到错误时返回错误原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogConfigsResponseParams `json:"Response"`
}

func (r *DeleteLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警策略id列表
	AlertIds []*string `json:"AlertIds,omitnil,omitempty" name:"AlertIds"`

	// 告警策略名称
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DeletePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警策略id列表
	AlertIds []*string `json:"AlertIds,omitnil,omitempty" name:"AlertIds"`

	// 告警策略名称
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DeletePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertIds")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *DeletePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertRuleRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警规则id列表
	AlertIds []*string `json:"AlertIds,omitnil,omitempty" name:"AlertIds"`
}

type DeletePrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警规则id列表
	AlertIds []*string `json:"AlertIds,omitnil,omitempty" name:"AlertIds"`
}

func (r *DeletePrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusAlertRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *DeletePrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusClusterAgentRequestParams struct {
	// agent列表
	Agents []*PrometheusAgentInfo `json:"Agents,omitnil,omitempty" name:"Agents"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 在7天可回收期间，强制解除绑定
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeletePrometheusClusterAgentRequest struct {
	*tchttp.BaseRequest
	
	// agent列表
	Agents []*PrometheusAgentInfo `json:"Agents,omitnil,omitempty" name:"Agents"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 在7天可回收期间，强制解除绑定
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *DeletePrometheusClusterAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusClusterAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Agents")
	delete(f, "InstanceId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusClusterAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusClusterAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusClusterAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusClusterAgentResponseParams `json:"Response"`
}

func (r *DeletePrometheusClusterAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusClusterAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 要删除的ServiceMonitor名字列表
	ServiceMonitors []*string `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// 要删除的PodMonitor名字列表
	PodMonitors []*string `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 要删除的RawJobs名字列表
	RawJobs []*string `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 要删除的Probe名字列表
	Probes []*string `json:"Probes,omitnil,omitempty" name:"Probes"`
}

type DeletePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 要删除的ServiceMonitor名字列表
	ServiceMonitors []*string `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// 要删除的PodMonitor名字列表
	PodMonitors []*string `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 要删除的RawJobs名字列表
	RawJobs []*string `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 要删除的Probe名字列表
	Probes []*string `json:"Probes,omitnil,omitempty" name:"Probes"`
}

func (r *DeletePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	delete(f, "Probes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusConfigResponseParams `json:"Response"`
}

func (r *DeletePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 聚合规则列表
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DeletePrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 聚合规则列表
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DeletePrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *DeletePrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempRequestParams struct {
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeletePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeletePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTempResponseParams `json:"Response"`
}

func (r *DeletePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempSyncRequestParams struct {
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 取消同步的对象列表
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type DeletePrometheusTempSyncRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 取消同步的对象列表
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *DeletePrometheusTempSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTempSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTempSyncResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusTempSyncResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTempSyncResponseParams `json:"Response"`
}

func (r *DeletePrometheusTempSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTempSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTemplateRequestParams struct {
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeletePrometheusTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeletePrometheusTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTemplateResponseParams `json:"Response"`
}

func (r *DeletePrometheusTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTemplateSyncRequestParams struct {
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 取消同步的对象列表
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type DeletePrometheusTemplateSyncRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 取消同步的对象列表
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *DeletePrometheusTemplateSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTemplateSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePrometheusTemplateSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePrometheusTemplateSyncResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePrometheusTemplateSyncResponse struct {
	*tchttp.BaseResponse
	Response *DeletePrometheusTemplateSyncResponseParams `json:"Response"`
}

func (r *DeletePrometheusTemplateSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePrometheusTemplateSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReservedInstancesRequestParams struct {
	// 预留券实例ID。可通过 [DescribeReservedInstances](https://cloud.tencent.com/document/product/457/99162) 接口返回值中的ReservedInstanceId获取。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`
}

type DeleteReservedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 预留券实例ID。可通过 [DescribeReservedInstances](https://cloud.tencent.com/document/product/457/99162) 接口返回值中的ReservedInstanceId获取。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`
}

func (r *DeleteReservedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReservedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReservedInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReservedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReservedInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReservedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReservedInstancesResponseParams `json:"Response"`
}

func (r *DeleteReservedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReservedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRollOutSequenceRequestParams struct {
	// 发布序列ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type DeleteRollOutSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 发布序列ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

func (r *DeleteRollOutSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRollOutSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRollOutSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRollOutSequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRollOutSequenceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRollOutSequenceResponseParams `json:"Response"`
}

func (r *DeleteRollOutSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRollOutSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTKEEdgeClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteTKEEdgeClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteTKEEdgeClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTKEEdgeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTKEEdgeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTKEEdgeClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTKEEdgeClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTKEEdgeClusterResponseParams `json:"Response"`
}

func (r *DeleteTKEEdgeClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTKEEdgeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddonRequestParams struct {
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称（不传时会返回集群下全部的addon）
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`
}

type DescribeAddonRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称（不传时会返回集群下全部的addon）
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`
}

func (r *DescribeAddonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AddonName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddonResponseParams struct {
	// addon列表
	Addons []*Addon `json:"Addons,omitnil,omitempty" name:"Addons"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddonResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddonResponseParams `json:"Response"`
}

func (r *DescribeAddonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddonValuesRequestParams struct {
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`
}

type DescribeAddonValuesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`
}

func (r *DescribeAddonValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddonValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AddonName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddonValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddonValuesResponseParams struct {
	// 参数列表，如果addon已安装，会使用已设置的参数和chart里的默认参数做渲染，是一个json格式的字符串，未安装addon时返回为空值。
	Values *string `json:"Values,omitnil,omitempty" name:"Values"`

	// addon支持的参数列表，值为chart的默认值，是一个json格式的字符串。
	DefaultValues *string `json:"DefaultValues,omitnil,omitempty" name:"DefaultValues"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddonValuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddonValuesResponseParams `json:"Response"`
}

func (r *DescribeAddonValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddonValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableClusterVersionRequestParams struct {
	// 集群 Id。若只查询某个集群可升级的版本，需填写此项，与ClusterIds 参数二选一。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群 Id 列表。若查询多个集群可升级的版本，需填写此项，与ClusterId 参数二选一。
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeAvailableClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群 Id。若只查询某个集群可升级的版本，需填写此项，与ClusterIds 参数二选一。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群 Id 列表。若查询多个集群可升级的版本，需填写此项，与ClusterId 参数二选一。
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *DescribeAvailableClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableClusterVersionResponseParams struct {
	// 可升级的集群版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Versions []*string `json:"Versions,omitnil,omitempty" name:"Versions"`

	// 集群信息
	Clusters []*ClusterVersion `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableClusterVersionResponseParams `json:"Response"`
}

func (r *DescribeAvailableClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableTKEEdgeVersionRequestParams struct {
	// 填写ClusterId获取当前集群各个组件版本和最新版本
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeAvailableTKEEdgeVersionRequest struct {
	*tchttp.BaseRequest
	
	// 填写ClusterId获取当前集群各个组件版本和最新版本
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeAvailableTKEEdgeVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableTKEEdgeVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableTKEEdgeVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableTKEEdgeVersionResponseParams struct {
	// 版本列表
	Versions []*string `json:"Versions,omitnil,omitempty" name:"Versions"`

	// 边缘集群最新版本
	EdgeVersionLatest *string `json:"EdgeVersionLatest,omitnil,omitempty" name:"EdgeVersionLatest"`

	// 边缘集群当前版本
	EdgeVersionCurrent *string `json:"EdgeVersionCurrent,omitnil,omitempty" name:"EdgeVersionCurrent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableTKEEdgeVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableTKEEdgeVersionResponseParams `json:"Response"`
}

func (r *DescribeAvailableTKEEdgeVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableTKEEdgeVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupStorageLocationsRequestParams struct {
	// 多个备份仓库名称，如果不填写，默认返回当前地域所有存储仓库名称
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeBackupStorageLocationsRequest struct {
	*tchttp.BaseRequest
	
	// 多个备份仓库名称，如果不填写，默认返回当前地域所有存储仓库名称
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DescribeBackupStorageLocationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupStorageLocationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupStorageLocationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupStorageLocationsResponseParams struct {
	// 详细备份仓库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupStorageLocationSet []*BackupStorageLocation `json:"BackupStorageLocationSet,omitnil,omitempty" name:"BackupStorageLocationSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupStorageLocationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupStorageLocationsResponseParams `json:"Response"`
}

func (r *DescribeBackupStorageLocationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupStorageLocationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchModifyTagsStatusRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBatchModifyTagsStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeBatchModifyTagsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchModifyTagsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchModifyTagsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchModifyTagsStatusResponseParams struct {
	// 失败资源列表
	FailedResources []*FailedResource `json:"FailedResources,omitnil,omitempty" name:"FailedResources"`

	// 任务状态：
	// - running 运行中
	// - failed 失败
	// - done 成功
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否同步集群内子资源标签
	SyncSubresource *bool `json:"SyncSubresource,omitnil,omitempty" name:"SyncSubresource"`

	// 集群标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchModifyTagsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchModifyTagsStatusResponseParams `json:"Response"`
}

func (r *DescribeBatchModifyTagsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchModifyTagsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupOptionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterAsGroupOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAsGroupOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAsGroupOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupOptionResponseParams struct {
	// 集群弹性伸缩属性
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitnil,omitempty" name:"ClusterAsGroupOption"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAsGroupOptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAsGroupOptionResponseParams `json:"Response"`
}

func (r *DescribeClusterAsGroupOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 伸缩组ID列表，如果为空，表示拉取集群关联的所有伸缩组。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeClusterAsGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 伸缩组ID列表，如果为空，表示拉取集群关联的所有伸缩组。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeClusterAsGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AutoScalingGroupIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAsGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAsGroupsResponseParams struct {
	// 集群关联的伸缩组总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群关联的伸缩组列表
	ClusterAsGroupSet []*ClusterAsGroup `json:"ClusterAsGroupSet,omitnil,omitempty" name:"ClusterAsGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAsGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAsGroupsResponseParams `json:"Response"`
}

func (r *DescribeClusterAsGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAuthenticationOptionsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterAuthenticationOptionsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAuthenticationOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAuthenticationOptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAuthenticationOptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAuthenticationOptionsResponseParams struct {
	// ServiceAccount认证配置
	ServiceAccounts *ServiceAccountAuthenticationOptions `json:"ServiceAccounts,omitnil,omitempty" name:"ServiceAccounts"`

	// 最近一次修改操作结果，返回值可能为：Updating，Success，Failed，TimeOut
	LatestOperationState *string `json:"LatestOperationState,omitnil,omitempty" name:"LatestOperationState"`

	// OIDC认证配置
	OIDCConfig *OIDCConfigAuthenticationOptions `json:"OIDCConfig,omitnil,omitempty" name:"OIDCConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAuthenticationOptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAuthenticationOptionsResponseParams `json:"Response"`
}

func (r *DescribeClusterAuthenticationOptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAuthenticationOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAvailableExtraArgsRequestParams struct {
	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 集群类型(MANAGED_CLUSTER或INDEPENDENT_CLUSTER)
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeClusterAvailableExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 集群类型(MANAGED_CLUSTER或INDEPENDENT_CLUSTER)
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribeClusterAvailableExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAvailableExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterVersion")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterAvailableExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterAvailableExtraArgsResponseParams struct {
	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 可用的自定义参数
	AvailableExtraArgs *AvailableExtraArgs `json:"AvailableExtraArgs,omitnil,omitempty" name:"AvailableExtraArgs"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterAvailableExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterAvailableExtraArgsResponseParams `json:"Response"`
}

func (r *DescribeClusterAvailableExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterAvailableExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterCommonNamesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 子账户列表，不可超出最大值50
	SubaccountUins []*string `json:"SubaccountUins,omitnil,omitempty" name:"SubaccountUins"`

	// 角色ID列表，不可超出最大值50
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

type DescribeClusterCommonNamesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 子账户列表，不可超出最大值50
	SubaccountUins []*string `json:"SubaccountUins,omitnil,omitempty" name:"SubaccountUins"`

	// 角色ID列表，不可超出最大值50
	RoleIds []*string `json:"RoleIds,omitnil,omitempty" name:"RoleIds"`
}

func (r *DescribeClusterCommonNamesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterCommonNamesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubaccountUins")
	delete(f, "RoleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterCommonNamesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterCommonNamesResponseParams struct {
	// 子账户Uin与其客户端证书的CN字段映射
	CommonNames []*CommonName `json:"CommonNames,omitnil,omitempty" name:"CommonNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterCommonNamesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterCommonNamesResponseParams `json:"Response"`
}

func (r *DescribeClusterCommonNamesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterCommonNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterControllersRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterControllersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterControllersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterControllersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterControllersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterControllersResponseParams struct {
	// 描述集群中各个控制器的状态
	ControllerStatusSet []*ControllerStatus `json:"ControllerStatusSet,omitnil,omitempty" name:"ControllerStatusSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterControllersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterControllersResponseParams `json:"Response"`
}

func (r *DescribeClusterControllersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterControllersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointStatusRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`
}

type DescribeClusterEndpointStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`
}

func (r *DescribeClusterEndpointStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterEndpointStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointStatusResponseParams struct {
	// 查询集群访问端口状态（Created 开启成功，Creating 开启中，NotFound 未开启，CreateFailed 开启失败）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 开启访问入口失败信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterEndpointStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterEndpointStatusResponseParams `json:"Response"`
}

func (r *DescribeClusterEndpointStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointVipStatusRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterEndpointVipStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointVipStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointVipStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterEndpointVipStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointVipStatusResponseParams struct {
	// 端口操作状态 (Creating 创建中  CreateFailed 创建失败 Created 创建完成 Deleting 删除中 DeletedFailed 删除失败 Deleted 已删除 NotFound 未发现操作 )
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作失败的原因
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterEndpointVipStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterEndpointVipStatusResponseParams `json:"Response"`
}

func (r *DescribeClusterEndpointVipStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointVipStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterEndpointsResponseParams struct {
	// 集群APIServer的CA证书
	CertificationAuthority *string `json:"CertificationAuthority,omitnil,omitempty" name:"CertificationAuthority"`

	// 集群APIServer的外网访问地址
	ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitnil,omitempty" name:"ClusterExternalEndpoint"`

	// 集群APIServer的内网访问地址
	ClusterIntranetEndpoint *string `json:"ClusterIntranetEndpoint,omitnil,omitempty" name:"ClusterIntranetEndpoint"`

	// 集群APIServer的域名
	ClusterDomain *string `json:"ClusterDomain,omitnil,omitempty" name:"ClusterDomain"`

	// 集群APIServer的外网访问ACL列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterExternalACL []*string `json:"ClusterExternalACL,omitnil,omitempty" name:"ClusterExternalACL"`

	// 外网域名
	ClusterExternalDomain *string `json:"ClusterExternalDomain,omitnil,omitempty" name:"ClusterExternalDomain"`

	// 内网域名
	ClusterIntranetDomain *string `json:"ClusterIntranetDomain,omitnil,omitempty" name:"ClusterIntranetDomain"`

	// 外网安全组
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 内网访问所属子网
	ClusterIntranetSubnetId *string `json:"ClusterIntranetSubnetId,omitnil,omitempty" name:"ClusterIntranetSubnetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterEndpointsResponseParams `json:"Response"`
}

func (r *DescribeClusterEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterExtraArgsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterExtraArgsResponseParams struct {
	// 集群自定义参数
	ClusterExtraArgs *ClusterExtraArgs `json:"ClusterExtraArgs,omitnil,omitempty" name:"ClusterExtraArgs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterExtraArgsResponseParams `json:"Response"`
}

func (r *DescribeClusterExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInspectionResultsOverviewRequestParams struct {
	// Array of String	目标集群列表，为空查询用户所有集群
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 聚合字段信息，概览结果按照 GroupBy 信息聚合后返回，可选参数：
	// catalogue.first：按一级分类聚合
	// catalogue.second：按二级分类聚合
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`
}

type DescribeClusterInspectionResultsOverviewRequest struct {
	*tchttp.BaseRequest
	
	// Array of String	目标集群列表，为空查询用户所有集群
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 聚合字段信息，概览结果按照 GroupBy 信息聚合后返回，可选参数：
	// catalogue.first：按一级分类聚合
	// catalogue.second：按二级分类聚合
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`
}

func (r *DescribeClusterInspectionResultsOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInspectionResultsOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "GroupBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInspectionResultsOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInspectionResultsOverviewResponseParams struct {
	// 诊断结果统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statistics []*KubeJarvisStateStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 诊断结果概览
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnostics []*KubeJarvisStateDiagnosticOverview `json:"Diagnostics,omitnil,omitempty" name:"Diagnostics"`

	// 集群诊断结果概览
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectionOverview []*KubeJarvisStateInspectionOverview `json:"InspectionOverview,omitnil,omitempty" name:"InspectionOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInspectionResultsOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInspectionResultsOverviewResponseParams `json:"Response"`
}

func (r *DescribeClusterInspectionResultsOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInspectionResultsOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesRequestParams struct {
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 需要获取的节点实例 ID 列表。如果为空，表示拉取集群下所有节点实例。普通节点用 CVM ID（如 ins-1cghhuuu ），原生节点用节点池内机器 ID （如 np-1ade4uid-0a2dc ）。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 节点角色，默认为WORKER类型
	// 
	// - WOKRER 集群内常规计算节点
	// - MASTER 独立集群控制面组件节点
	// - ETCD 独立集群 etcd 存储节点
	// - MASTER_ETCD 独立集群控制面、etcd 共同部署的 master 节点
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 过滤条件列表；Name的可选值为nodepool-id、nodepool-instance-type；Name为nodepool-id表示根据节点池id过滤机器，Value的值为具体的节点池id，Name为nodepool-instance-type表示节点加入节点池的方式，Value的值为MANUALLY_ADDED（手动加入节点池）、AUTOSCALING_ADDED（伸缩组扩容方式加入节点池）、ALL（手动加入节点池 和 伸缩组扩容方式加入节点池）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 需要获取的节点实例 ID 列表。如果为空，表示拉取集群下所有节点实例。普通节点用 CVM ID（如 ins-1cghhuuu ），原生节点用节点池内机器 ID （如 np-1ade4uid-0a2dc ）。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 节点角色，默认为WORKER类型
	// 
	// - WOKRER 集群内常规计算节点
	// - MASTER 独立集群控制面组件节点
	// - ETCD 独立集群 etcd 存储节点
	// - MASTER_ETCD 独立集群控制面、etcd 共同部署的 master 节点
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 过滤条件列表；Name的可选值为nodepool-id、nodepool-instance-type；Name为nodepool-id表示根据节点池id过滤机器，Value的值为具体的节点池id，Name为nodepool-instance-type表示节点加入节点池的方式，Value的值为MANUALLY_ADDED（手动加入节点池）、AUTOSCALING_ADDED（伸缩组扩容方式加入节点池）、ALL（手动加入节点池 和 伸缩组扩容方式加入节点池）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	delete(f, "InstanceRole")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterInstancesResponseParams struct {
	// 集群中实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群中实例列表
	InstanceSet []*Instance `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterInstancesResponseParams `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterKubeconfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 默认false 获取内网，是否获取外网访问的kubeconfig
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`
}

type DescribeClusterKubeconfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 默认false 获取内网，是否获取外网访问的kubeconfig
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`
}

func (r *DescribeClusterKubeconfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterKubeconfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterKubeconfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterKubeconfigResponseParams struct {
	// 子账户kubeconfig文件，可用于直接访问集群kube-apiserver（入参IsExtranet为false，返回内网访问的kubeconfig，没开内网的情况下server会是一个默认域名；入参IsExtranet为true，返回外网的kubeconfig，没开外网的情况下server会是一个默认域名。默认域名默认不可达，需要自行处理）
	Kubeconfig *string `json:"Kubeconfig,omitnil,omitempty" name:"Kubeconfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterKubeconfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterKubeconfigResponseParams `json:"Response"`
}

func (r *DescribeClusterKubeconfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelAttributeRequestParams struct {
	// 集群ID，变配时使用
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`
}

type DescribeClusterLevelAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，变配时使用
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`
}

func (r *DescribeClusterLevelAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterLevelAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelAttributeResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群规模
	Items []*ClusterLevelAttribute `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterLevelAttributeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterLevelAttributeResponseParams `json:"Response"`
}

func (r *DescribeClusterLevelAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelChangeRecordsRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 结束时间
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeClusterLevelChangeRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 结束时间
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeClusterLevelChangeRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelChangeRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "StartAt")
	delete(f, "EndAt")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterLevelChangeRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterLevelChangeRecordsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群规模
	Items []*ClusterLevelChangeRecord `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterLevelChangeRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterLevelChangeRecordsResponseParams `json:"Response"`
}

func (r *DescribeClusterLevelChangeRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterLevelChangeRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterMaintenanceWindowAndExclusionsRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤项
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClusterMaintenanceWindowAndExclusionsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤项
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeClusterMaintenanceWindowAndExclusionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterMaintenanceWindowAndExclusionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterMaintenanceWindowAndExclusionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterMaintenanceWindowAndExclusionsResponseParams struct {
	// 维护时间窗口和排除项
	MaintenanceWindowAndExclusions []*MaintenanceWindowAndExclusion `json:"MaintenanceWindowAndExclusions,omitnil,omitempty" name:"MaintenanceWindowAndExclusions"`

	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterMaintenanceWindowAndExclusionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterMaintenanceWindowAndExclusionsResponseParams `json:"Response"`
}

func (r *DescribeClusterMaintenanceWindowAndExclusionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterMaintenanceWindowAndExclusionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolDetailRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`
}

type DescribeClusterNodePoolDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`
}

func (r *DescribeClusterNodePoolDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodePoolDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolDetailResponseParams struct {
	// 节点池详情
	NodePool *NodePool `json:"NodePool,omitnil,omitempty" name:"NodePool"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterNodePoolDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodePoolDetailResponseParams `json:"Response"`
}

func (r *DescribeClusterNodePoolDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolsRequestParams struct {
	// ClusterId（集群id）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// · "Name":"NodePoolsName","Values": ["test"]
	//     按照【节点池名】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  "Name":"NodePoolsId","Values": ["np-d2mb2zb"]
	//     按照【节点池id】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  "Name":"Tags","Values": ["product:tke"]
	//     按照【标签键值对】进行过滤。
	//     类型：String
	//     必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClusterNodePoolsRequest struct {
	*tchttp.BaseRequest
	
	// ClusterId（集群id）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// · "Name":"NodePoolsName","Values": ["test"]
	//     按照【节点池名】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  "Name":"NodePoolsId","Values": ["np-d2mb2zb"]
	//     按照【节点池id】进行过滤。
	//     类型：String
	//     必选：否
	// 
	// ·  "Name":"Tags","Values": ["product:tke"]
	//     按照【标签键值对】进行过滤。
	//     类型：String
	//     必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeClusterNodePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNodePoolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNodePoolsResponseParams struct {
	// NodePools（节点池列表）
	NodePoolSet []*NodePool `json:"NodePoolSet,omitnil,omitempty" name:"NodePoolSet"`

	// 资源总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterNodePoolsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNodePoolsResponseParams `json:"Response"`
}

func (r *DescribeClusterNodePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNodePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPendingReleasesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 返回数量限制，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 集群类型，支持传入tke（标准集群），eks（Serverless集群)，external（注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeClusterPendingReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 返回数量限制，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 集群类型，支持传入tke（标准集群），eks（Serverless集群)，external（注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribeClusterPendingReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPendingReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPendingReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPendingReleasesResponseParams struct {
	// 正在安装中应用列表
	ReleaseSet []*PendingRelease `json:"ReleaseSet,omitnil,omitempty" name:"ReleaseSet"`

	// 每页返回数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterPendingReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPendingReleasesResponseParams `json:"Response"`
}

func (r *DescribeClusterPendingReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPendingReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReleaseDetailsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用所在命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群类型，传入 tke(标准集群), eks(Serverless集群), external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeClusterReleaseDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用所在命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群类型，传入 tke(标准集群), eks(Serverless集群), external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribeClusterReleaseDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReleaseDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterReleaseDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReleaseDetailsResponseParams struct {
	// 应用详情
	Release *ReleaseDetails `json:"Release,omitnil,omitempty" name:"Release"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterReleaseDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterReleaseDetailsResponseParams `json:"Response"`
}

func (r *DescribeClusterReleaseDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReleaseDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReleaseHistoryRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用所在命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群类型，传入 tke(标准集群), eks(Serverless集群), external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeClusterReleaseHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用所在命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群类型，传入 tke(标准集群), eks(Serverless集群), external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribeClusterReleaseHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReleaseHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterReleaseHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReleaseHistoryResponseParams struct {
	// 已安装应用版本历史
	ReleaseHistorySet []*ReleaseHistory `json:"ReleaseHistorySet,omitnil,omitempty" name:"ReleaseHistorySet"`

	// 总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterReleaseHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterReleaseHistoryResponseParams `json:"Response"`
}

func (r *DescribeClusterReleaseHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReleaseHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReleasesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 每页数量限制，默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 集群类型，传入 tke(标准集群)，eks(Serverless集群)，external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// helm Release 安装的namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// helm Release 的名字
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// helm Chart 的名字
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`
}

type DescribeClusterReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 每页数量限制，默认值为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 集群类型，传入 tke(标准集群)，eks(Serverless集群)，external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// helm Release 安装的namespace
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// helm Release 的名字
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// helm Chart 的名字
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`
}

func (r *DescribeClusterReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ClusterType")
	delete(f, "Namespace")
	delete(f, "ReleaseName")
	delete(f, "ChartName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterReleasesResponseParams struct {
	// 数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 已安装应用列表
	ReleaseSet []*Release `json:"ReleaseSet,omitnil,omitempty" name:"ReleaseSet"`

	// 已安装应用总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterReleasesResponseParams `json:"Response"`
}

func (r *DescribeClusterReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRollOutSequenceTagsRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤项
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClusterRollOutSequenceTagsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤项
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeClusterRollOutSequenceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRollOutSequenceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRollOutSequenceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRollOutSequenceTagsResponseParams struct {
	// 集群发布序列标签
	ClusterTags []*ClusterRollOutSequenceTag `json:"ClusterTags,omitnil,omitempty" name:"ClusterTags"`

	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterRollOutSequenceTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRollOutSequenceTagsResponseParams `json:"Response"`
}

func (r *DescribeClusterRollOutSequenceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRollOutSequenceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRouteTablesRequestParams struct {

}

type DescribeClusterRouteTablesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClusterRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRouteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRouteTablesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群路由表对象。
	RouteTableSet []*RouteTableInfo `json:"RouteTableSet,omitnil,omitempty" name:"RouteTableSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRouteTablesResponseParams `json:"Response"`
}

func (r *DescribeClusterRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRoutesRequestParams struct {
	// 路由表名称。与集群 ID 一致，可以到[集群控制台](https://console.cloud.tencent.com/tke2)进行复制。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 过滤条件,当前只支持按照单个条件GatewayIP进行过滤（可选）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClusterRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表名称。与集群 ID 一致，可以到[集群控制台](https://console.cloud.tencent.com/tke2)进行复制。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 过滤条件,当前只支持按照单个条件GatewayIP进行过滤（可选）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeClusterRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterRoutesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群路由对象。
	RouteSet []*RouteInfo `json:"RouteSet,omitnil,omitempty" name:"RouteSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterRoutesResponseParams `json:"Response"`
}

func (r *DescribeClusterRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSecurityRequestParams struct {
	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterSecurityRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterSecurityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSecurityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterSecurityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSecurityResponseParams struct {
	// 集群的账号名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 集群的访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 集群访问CA证书
	CertificationAuthority *string `json:"CertificationAuthority,omitnil,omitempty" name:"CertificationAuthority"`

	// 集群访问的地址
	ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitnil,omitempty" name:"ClusterExternalEndpoint"`

	// 集群访问的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 集群Endpoint地址
	PgwEndpoint *string `json:"PgwEndpoint,omitnil,omitempty" name:"PgwEndpoint"`

	// 集群访问策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityPolicy []*string `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`

	// 集群Kubeconfig文件
	Kubeconfig *string `json:"Kubeconfig,omitnil,omitempty" name:"Kubeconfig"`

	// 集群JnsGw的访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	JnsGwEndpoint *string `json:"JnsGwEndpoint,omitnil,omitempty" name:"JnsGwEndpoint"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterSecurityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterSecurityResponseParams `json:"Response"`
}

func (r *DescribeClusterSecurityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSecurityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStatusRequestParams struct {
	// 集群ID列表，不传默认拉取所有集群
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeClusterStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表，不传默认拉取所有集群
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStatusResponseParams struct {
	// 集群状态列表
	ClusterStatusSet []*ClusterStatus `json:"ClusterStatusSet,omitnil,omitempty" name:"ClusterStatusSet"`

	// 集群个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterStatusResponseParams `json:"Response"`
}

func (r *DescribeClusterStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterVirtualNodePoolsRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterVirtualNodePoolsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterVirtualNodePoolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterVirtualNodePoolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterVirtualNodePoolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterVirtualNodePoolsResponseParams struct {
	// 节点池总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 虚拟节点池列表
	NodePoolSet []*VirtualNodePool `json:"NodePoolSet,omitnil,omitempty" name:"NodePoolSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterVirtualNodePoolsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterVirtualNodePoolsResponseParams `json:"Response"`
}

func (r *DescribeClusterVirtualNodePoolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterVirtualNodePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterVirtualNodeRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID，通过DescribeNodePools接口获取
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点名称，可搜索DescribeClusterVirtualNode接口节点
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`
}

type DescribeClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID，通过DescribeNodePools接口获取
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点名称，可搜索DescribeClusterVirtualNode接口节点
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`
}

func (r *DescribeClusterVirtualNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterVirtualNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "NodeNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterVirtualNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterVirtualNodeResponseParams struct {
	// 节点列表
	Nodes []*VirtualNode `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 节点总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterVirtualNodeResponseParams `json:"Response"`
}

func (r *DescribeClusterVirtualNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// · "Name":"ClusterName","Values": ["test"] 按照【集群名】进行过滤。 类型：String 必选：否 · "Name":"ClusterType","Values": ["MANAGED_CLUSTER"] 按照【集群类型】进行过滤。 类型：String 必选：否 · "Name":"ClusterStatus","Values": ["Running"] 按照【集群状态】进行过滤。 类型：String 必选：否 · "Name":"vpc-id","Values": ["vpc-2wds9k9p"] 按照【VPC】进行过滤。 类型：String 必选：否 · "Name":"tag-key","Values": ["testKey"] 按照【标签键】进行过滤。 类型：String 必选：否 · "Name":"tag-value","Values": ["testValue"] 按照【标签值】进行过滤。 类型：String 必选：否 · "Name":"Tags","Values": ["product:tke"] 按照【标签键值对】进行过滤。 类型：String 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// · "Name":"ClusterName","Values": ["test"] 按照【集群名】进行过滤。 类型：String 必选：否 · "Name":"ClusterType","Values": ["MANAGED_CLUSTER"] 按照【集群类型】进行过滤。 类型：String 必选：否 · "Name":"ClusterStatus","Values": ["Running"] 按照【集群状态】进行过滤。 类型：String 必选：否 · "Name":"vpc-id","Values": ["vpc-2wds9k9p"] 按照【VPC】进行过滤。 类型：String 必选：否 · "Name":"tag-key","Values": ["testKey"] 按照【标签键】进行过滤。 类型：String 必选：否 · "Name":"tag-value","Values": ["testValue"] 按照【标签值】进行过滤。 类型：String 必选：否 · "Name":"Tags","Values": ["product:tke"] 按照【标签键值对】进行过滤。 类型：String 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
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
	delete(f, "Filters")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群总个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群信息列表
	Clusters []*Cluster `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeControlPlaneLogsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型。当前只支持tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeControlPlaneLogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型。当前只支持tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribeControlPlaneLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeControlPlaneLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeControlPlaneLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeControlPlaneLogsResponseParams struct {
	// 日志采集配置查询结果
	Details []*ComponentLogConfig `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeControlPlaneLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeControlPlaneLogsResponseParams `json:"Response"`
}

func (r *DescribeControlPlaneLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeControlPlaneLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeECMInstancesRequestParams struct {
	// 集群id
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 过滤条件
	// 仅支持ecm-id过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeECMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 过滤条件
	// 仅支持ecm-id过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeECMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeECMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeECMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeECMInstancesResponseParams struct {
	// 返回的实例相关信息列表的长度
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的实例相关信息列表
	InstanceInfoSet []*string `json:"InstanceInfoSet,omitnil,omitempty" name:"InstanceInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeECMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeECMInstancesResponseParams `json:"Response"`
}

func (r *DescribeECMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeECMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSClusterCredentialRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeEKSClusterCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeEKSClusterCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSClusterCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEKSClusterCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSClusterCredentialResponseParams struct {
	// 集群的接入地址信息
	Addresses []*IPAddress `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// 集群的认证信息（token只有请求是主账号才返回，子账户请使用返回的kubeconfig）
	Credential *ClusterCredential `json:"Credential,omitnil,omitempty" name:"Credential"`

	// 集群的公网访问信息
	PublicLB *ClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 集群的内网访问信息
	InternalLB *ClusterInternalLB `json:"InternalLB,omitnil,omitempty" name:"InternalLB"`

	// 标记是否新的内外网功能
	ProxyLB *bool `json:"ProxyLB,omitnil,omitempty" name:"ProxyLB"`

	// 连接用户集群k8s 的Config
	Kubeconfig *string `json:"Kubeconfig,omitnil,omitempty" name:"Kubeconfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEKSClusterCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEKSClusterCredentialResponseParams `json:"Response"`
}

func (r *DescribeEKSClusterCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSClusterCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSClustersRequestParams struct {
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件,当前只支持按照单个条件ClusterName进行过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEKSClustersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件,当前只支持按照单个条件ClusterName进行过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEKSClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEKSClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSClustersResponseParams struct {
	// 集群总个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群信息列表
	Clusters []*EksCluster `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEKSClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEKSClustersResponseParams `json:"Response"`
}

func (r *DescribeEKSClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSContainerInstanceEventRequestParams struct {
	// 容器实例id
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 最大事件数量。默认为50，最大取值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeEKSContainerInstanceEventRequest struct {
	*tchttp.BaseRequest
	
	// 容器实例id
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 最大事件数量。默认为50，最大取值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeEKSContainerInstanceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSContainerInstanceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EksCiId")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEKSContainerInstanceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSContainerInstanceEventResponseParams struct {
	// 事件集合
	Events []*Event `json:"Events,omitnil,omitempty" name:"Events"`

	// 容器实例id
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEKSContainerInstanceEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEKSContainerInstanceEventResponseParams `json:"Response"`
}

func (r *DescribeEKSContainerInstanceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSContainerInstanceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSContainerInstanceRegionsRequestParams struct {

}

type DescribeEKSContainerInstanceRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEKSContainerInstanceRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSContainerInstanceRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEKSContainerInstanceRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSContainerInstanceRegionsResponseParams struct {
	// EKS Container Instance支持的地域信息
	Regions []*EksCiRegionInfo `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEKSContainerInstanceRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEKSContainerInstanceRegionsResponseParams `json:"Response"`
}

func (r *DescribeEKSContainerInstanceRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSContainerInstanceRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSContainerInstancesRequestParams struct {
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，可条件：
	// (1)实例名称
	// KeyName: eks-ci-name
	// 类型：String
	// 
	// (2)实例状态
	// KeyName: status
	// 类型：String
	// 可选值："Pending", "Running", "Succeeded", "Failed"
	// 
	// (3)内网ip
	// KeyName: private-ip
	// 类型：String
	// 
	// (4)EIP地址
	// KeyName: eip-address
	// 类型：String
	// 
	// (5)VpcId
	// KeyName: vpc-id
	// 类型：String
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 容器实例 ID 数组
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`
}

type DescribeEKSContainerInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，可条件：
	// (1)实例名称
	// KeyName: eks-ci-name
	// 类型：String
	// 
	// (2)实例状态
	// KeyName: status
	// 类型：String
	// 可选值："Pending", "Running", "Succeeded", "Failed"
	// 
	// (3)内网ip
	// KeyName: private-ip
	// 类型：String
	// 
	// (4)EIP地址
	// KeyName: eip-address
	// 类型：String
	// 
	// (5)VpcId
	// KeyName: vpc-id
	// 类型：String
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 容器实例 ID 数组
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`
}

func (r *DescribeEKSContainerInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSContainerInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "EksCiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEKSContainerInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEKSContainerInstancesResponseParams struct {
	// 容器组总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 容器组列表
	EksCis []*EksCi `json:"EksCis,omitnil,omitempty" name:"EksCis"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEKSContainerInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEKSContainerInstancesResponseParams `json:"Response"`
}

func (r *DescribeEKSContainerInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEKSContainerInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAvailableExtraArgsRequestParams struct {
	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`
}

type DescribeEdgeAvailableExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`
}

func (r *DescribeEdgeAvailableExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAvailableExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeAvailableExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAvailableExtraArgsResponseParams struct {
	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 可用的自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableExtraArgs *EdgeAvailableExtraArgs `json:"AvailableExtraArgs,omitnil,omitempty" name:"AvailableExtraArgs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeAvailableExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeAvailableExtraArgsResponseParams `json:"Response"`
}

func (r *DescribeEdgeAvailableExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAvailableExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeCVMInstancesRequestParams struct {
	// 集群id
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 过滤条件
	// 仅支持cvm-id过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEdgeCVMInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 过滤条件
	// 仅支持cvm-id过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEdgeCVMInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeCVMInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeCVMInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeCVMInstancesResponseParams struct {
	// 返回的实例相关信息列表的长度
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的实例相关信息列表
	InstanceInfoSet []*string `json:"InstanceInfoSet,omitnil,omitempty" name:"InstanceInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeCVMInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeCVMInstancesResponseParams `json:"Response"`
}

func (r *DescribeEdgeCVMInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeCVMInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterExtraArgsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeEdgeClusterExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeEdgeClusterExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeClusterExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterExtraArgsResponseParams struct {
	// 集群自定义参数
	ClusterExtraArgs *EdgeClusterExtraArgs `json:"ClusterExtraArgs,omitnil,omitempty" name:"ClusterExtraArgs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeClusterExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeClusterExtraArgsResponseParams `json:"Response"`
}

func (r *DescribeEdgeClusterExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterInstancesRequestParams struct {
	// 集群id
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 查询总数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，仅支持NodeName过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEdgeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 查询总数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，仅支持NodeName过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeEdgeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterInstancesResponseParams struct {
	// 该集群总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 节点信息集合
	InstanceInfoSet *string `json:"InstanceInfoSet,omitnil,omitempty" name:"InstanceInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeClusterInstancesResponseParams `json:"Response"`
}

func (r *DescribeEdgeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterUpgradeInfoRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 要升级到的TKEEdge版本
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`
}

type DescribeEdgeClusterUpgradeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 要升级到的TKEEdge版本
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`
}

func (r *DescribeEdgeClusterUpgradeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterUpgradeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EdgeVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeClusterUpgradeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeClusterUpgradeInfoResponseParams struct {
	// 可升级的集群组件及其版本
	ComponentVersion *string `json:"ComponentVersion,omitnil,omitempty" name:"ComponentVersion"`

	// 边缘集群当前版本
	EdgeVersionCurrent *string `json:"EdgeVersionCurrent,omitnil,omitempty" name:"EdgeVersionCurrent"`

	// 边缘组件镜像仓库地址前缀，包含域名和命名空间
	RegistryPrefix *string `json:"RegistryPrefix,omitnil,omitempty" name:"RegistryPrefix"`

	// 集群升级状态，可能值：running、updating、failed
	ClusterUpgradeStatus *string `json:"ClusterUpgradeStatus,omitnil,omitempty" name:"ClusterUpgradeStatus"`

	// 集群升级中状态或者失败原因
	ClusterUpgradeStatusReason *string `json:"ClusterUpgradeStatusReason,omitnil,omitempty" name:"ClusterUpgradeStatusReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeClusterUpgradeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeClusterUpgradeInfoResponseParams `json:"Response"`
}

func (r *DescribeEdgeClusterUpgradeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeClusterUpgradeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeLogSwitchesRequestParams struct {
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

type DescribeEdgeLogSwitchesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`
}

func (r *DescribeEdgeLogSwitchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeLogSwitchesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeLogSwitchesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeLogSwitchesResponseParams struct {
	// 集群日志开关集合
	SwitchSet []*string `json:"SwitchSet,omitnil,omitempty" name:"SwitchSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeLogSwitchesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeLogSwitchesResponseParams `json:"Response"`
}

func (r *DescribeEdgeLogSwitchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeLogSwitchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEksContainerInstanceLogRequestParams struct {
	// Eks Container Instance Id，即容器实例Id
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 容器名称，单容器的实例可选填。如果为多容器实例，请指定容器名称。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// 返回最新日志行数，默认500，最大2000。日志内容最大返回 1M 数据。
	Tail *uint64 `json:"Tail,omitnil,omitempty" name:"Tail"`

	// UTC时间，RFC3339标准
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 是否是查上一个容器（如果容器退出重启了）
	Previous *bool `json:"Previous,omitnil,omitempty" name:"Previous"`

	// 查询最近多少秒内的日志
	SinceSeconds *uint64 `json:"SinceSeconds,omitnil,omitempty" name:"SinceSeconds"`

	// 日志总大小限制
	LimitBytes *uint64 `json:"LimitBytes,omitnil,omitempty" name:"LimitBytes"`
}

type DescribeEksContainerInstanceLogRequest struct {
	*tchttp.BaseRequest
	
	// Eks Container Instance Id，即容器实例Id
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 容器名称，单容器的实例可选填。如果为多容器实例，请指定容器名称。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// 返回最新日志行数，默认500，最大2000。日志内容最大返回 1M 数据。
	Tail *uint64 `json:"Tail,omitnil,omitempty" name:"Tail"`

	// UTC时间，RFC3339标准
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 是否是查上一个容器（如果容器退出重启了）
	Previous *bool `json:"Previous,omitnil,omitempty" name:"Previous"`

	// 查询最近多少秒内的日志
	SinceSeconds *uint64 `json:"SinceSeconds,omitnil,omitempty" name:"SinceSeconds"`

	// 日志总大小限制
	LimitBytes *uint64 `json:"LimitBytes,omitnil,omitempty" name:"LimitBytes"`
}

func (r *DescribeEksContainerInstanceLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEksContainerInstanceLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EksCiId")
	delete(f, "ContainerName")
	delete(f, "Tail")
	delete(f, "StartTime")
	delete(f, "Previous")
	delete(f, "SinceSeconds")
	delete(f, "LimitBytes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEksContainerInstanceLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEksContainerInstanceLogResponseParams struct {
	// 容器名称
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// 日志内容
	LogContent *string `json:"LogContent,omitnil,omitempty" name:"LogContent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEksContainerInstanceLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEksContainerInstanceLogResponseParams `json:"Response"`
}

func (r *DescribeEksContainerInstanceLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEksContainerInstanceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnableVpcCniProgressRequestParams struct {
	// 开启vpc-cni的集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeEnableVpcCniProgressRequest struct {
	*tchttp.BaseRequest
	
	// 开启vpc-cni的集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeEnableVpcCniProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnableVpcCniProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnableVpcCniProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnableVpcCniProgressResponseParams struct {
	// 任务进度的描述：Running/Succeed/Failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 当任务进度为Failed时，对任务状态的进一步描述，例如IPAMD组件安装失败
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnableVpcCniProgressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnableVpcCniProgressResponseParams `json:"Response"`
}

func (r *DescribeEnableVpcCniProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnableVpcCniProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptionStatusRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeEncryptionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeEncryptionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEncryptionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEncryptionStatusResponseParams struct {
	// 加密状态
	// - Opening: 开启中
	// - Opened: 已开启
	// - Closing: 关闭中
	// - Closed: 已关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 加密错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEncryptionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEncryptionStatusResponseParams `json:"Response"`
}

func (r *DescribeEncryptionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEncryptionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExistedInstancesRequestParams struct {
	// 集群 ID，请填写查询集群列表 接口中返回的 ClusterId 字段（仅通过ClusterId获取需要过滤条件中的VPCID。节点状态比较时会使用该地域下所有集群中的节点进行比较。参数不支持同时指定InstanceIds和ClusterId。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 按照一个或者多个实例ID查询。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API简介的id.N一节）。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤条件,字段和详见[CVM查询实例](https://cloud.tencent.com/document/api/213/15728)如果设置了ClusterId，会附加集群的VPCID作为查询字段，在此情况下如果在Filter中指定了"vpc-id"作为过滤字段，指定的VPCID必须与集群的VPCID相同。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 实例IP进行过滤(同时支持内网IP和外网IP)
	VagueIpAddress *string `json:"VagueIpAddress,omitnil,omitempty" name:"VagueIpAddress"`

	// 实例名称进行过滤
	VagueInstanceName *string `json:"VagueInstanceName,omitnil,omitempty" name:"VagueInstanceName"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据多个实例IP进行过滤
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`
}

type DescribeExistedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID，请填写查询集群列表 接口中返回的 ClusterId 字段（仅通过ClusterId获取需要过滤条件中的VPCID。节点状态比较时会使用该地域下所有集群中的节点进行比较。参数不支持同时指定InstanceIds和ClusterId。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 按照一个或者多个实例ID查询。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API简介的id.N一节）。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤条件,字段和详见[CVM查询实例](https://cloud.tencent.com/document/api/213/15728)如果设置了ClusterId，会附加集群的VPCID作为查询字段，在此情况下如果在Filter中指定了"vpc-id"作为过滤字段，指定的VPCID必须与集群的VPCID相同。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 实例IP进行过滤(同时支持内网IP和外网IP)
	VagueIpAddress *string `json:"VagueIpAddress,omitnil,omitempty" name:"VagueIpAddress"`

	// 实例名称进行过滤
	VagueInstanceName *string `json:"VagueInstanceName,omitnil,omitempty" name:"VagueInstanceName"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据多个实例IP进行过滤
	IpAddresses []*string `json:"IpAddresses,omitnil,omitempty" name:"IpAddresses"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExistedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "VagueIpAddress")
	delete(f, "VagueInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExistedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExistedInstancesResponseParams struct {
	// 已经存在的实例信息数组。
	ExistedInstanceSet []*ExistedInstance `json:"ExistedInstanceSet,omitnil,omitempty" name:"ExistedInstanceSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExistedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExistedInstancesResponseParams `json:"Response"`
}

func (r *DescribeExistedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExistedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalNodeSupportConfigRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeExternalNodeSupportConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeExternalNodeSupportConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalNodeSupportConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExternalNodeSupportConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExternalNodeSupportConfigResponseParams struct {
	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。
	ClusterCIDR *string `json:"ClusterCIDR,omitnil,omitempty" name:"ClusterCIDR"`

	// 集群网络插件类型，支持：CiliumBGP、CiliumVXLan
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否开启第三方节点专线连接支持
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 节点所属交换机的BGP AS 号
	AS *string `json:"AS,omitnil,omitempty" name:"AS"`

	// 节点所属交换机的交换机 IP
	SwitchIP *string `json:"SwitchIP,omitnil,omitempty" name:"SwitchIP"`

	// 开启第三方节点池状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 如果开启失败原因
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 内网访问地址
	Master *string `json:"Master,omitnil,omitempty" name:"Master"`

	// 镜像仓库代理地址
	Proxy *string `json:"Proxy,omitnil,omitempty" name:"Proxy"`

	// 用于记录开启第三方节点的过程进行到哪一步了
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress []*Step `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 是否开启第三方节点公网连接支持
	EnabledPublicConnect *bool `json:"EnabledPublicConnect,omitnil,omitempty" name:"EnabledPublicConnect"`

	// 注册节点公网版公网连接地址
	PublicConnectUrl *string `json:"PublicConnectUrl,omitnil,omitempty" name:"PublicConnectUrl"`

	// 注册节点公网版自定义域名
	PublicCustomDomain *string `json:"PublicCustomDomain,omitnil,omitempty" name:"PublicCustomDomain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExternalNodeSupportConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExternalNodeSupportConfigResponseParams `json:"Response"`
}

func (r *DescribeExternalNodeSupportConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExternalNodeSupportConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalMaintenanceWindowAndExclusionsRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选项，支持按照以下字段过滤：
	// - "ID"：维护窗口ID，如：10
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeGlobalMaintenanceWindowAndExclusionsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选项，支持按照以下字段过滤：
	// - "ID"：维护窗口ID，如：10
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeGlobalMaintenanceWindowAndExclusionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalMaintenanceWindowAndExclusionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalMaintenanceWindowAndExclusionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalMaintenanceWindowAndExclusionsResponseParams struct {
	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 维护时间窗口
	MaintenanceWindowAndExclusions []*GlobalMaintenanceWindowAndExclusion `json:"MaintenanceWindowAndExclusions,omitnil,omitempty" name:"MaintenanceWindowAndExclusions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalMaintenanceWindowAndExclusionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalMaintenanceWindowAndExclusionsResponseParams `json:"Response"`
}

func (r *DescribeGlobalMaintenanceWindowAndExclusionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalMaintenanceWindowAndExclusionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPAMDRequestParams struct {
	// 集群 ID，请[登录控制台](https://console.cloud.tencent.com/tke2/cluster)在集群列表复制ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeIPAMDRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID，请[登录控制台](https://console.cloud.tencent.com/tke2/cluster)在集群列表复制ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeIPAMDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPAMDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPAMDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPAMDResponseParams struct {
	// 是否安装了eniipamd组件
	EnableIPAMD *bool `json:"EnableIPAMD,omitnil,omitempty" name:"EnableIPAMD"`

	// 是否开启自定义podcidr，默认为false，已安装eniipamd组件才意义
	EnableCustomizedPodCidr *bool `json:"EnableCustomizedPodCidr,omitnil,omitempty" name:"EnableCustomizedPodCidr"`

	// 是否不开启vpccni模式，默认为false，已安装eniipamd组件才意义
	DisableVpcCniMode *bool `json:"DisableVpcCniMode,omitnil,omitempty" name:"DisableVpcCniMode"`

	// 组件状态，已安装eniipamd组件才会有值
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 错误信息，已安装eniipamd组件且状态为非running才会有错误信息
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 子网信息，已安装eniipamd组件才会有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 固定ip回收时间，已安装eniipamd组件才会有值
	ClaimExpiredDuration *string `json:"ClaimExpiredDuration,omitnil,omitempty" name:"ClaimExpiredDuration"`

	// 是否开启了中继网卡模式
	EnableTrunkingENI *bool `json:"EnableTrunkingENI,omitnil,omitempty" name:"EnableTrunkingENI"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIPAMDResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPAMDResponseParams `json:"Response"`
}

func (r *DescribeIPAMDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPAMDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageCachesRequestParams struct {
	// 镜像缓存Id数组
	ImageCacheIds []*string `json:"ImageCacheIds,omitnil,omitempty" name:"ImageCacheIds"`

	// 镜像缓存名称数组
	ImageCacheNames []*string `json:"ImageCacheNames,omitnil,omitempty" name:"ImageCacheNames"`

	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，可选条件：
	// (1)实例名称
	// KeyName: image-cache-name
	// 类型：String
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeImageCachesRequest struct {
	*tchttp.BaseRequest
	
	// 镜像缓存Id数组
	ImageCacheIds []*string `json:"ImageCacheIds,omitnil,omitempty" name:"ImageCacheIds"`

	// 镜像缓存名称数组
	ImageCacheNames []*string `json:"ImageCacheNames,omitnil,omitempty" name:"ImageCacheNames"`

	// 限定此次返回资源的数量。如果不设定，默认返回20，最大不能超过50
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，可选条件：
	// (1)实例名称
	// KeyName: image-cache-name
	// 类型：String
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeImageCachesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageCachesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageCacheIds")
	delete(f, "ImageCacheNames")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImageCachesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImageCachesResponseParams struct {
	// 镜像缓存总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像缓存信息列表
	ImageCaches []*ImageCache `json:"ImageCaches,omitnil,omitempty" name:"ImageCaches"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImageCachesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImageCachesResponseParams `json:"Response"`
}

func (r *DescribeImageCachesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImageCachesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesRequestParams struct {

}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeImagesResponseParams struct {
	// 镜像数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 镜像信息列表
	ImageInstanceSet []*ImageInstance `json:"ImageInstanceSet,omitnil,omitempty" name:"ImageInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeImagesResponseParams `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogConfigsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 当前集群类型支持tke、eks。默认为tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 按照采集规则名称查找，多个采集规则使用 "," 分隔。
	LogConfigNames *string `json:"LogConfigNames,omitnil,omitempty" name:"LogConfigNames"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLogConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 当前集群类型支持tke、eks。默认为tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 按照采集规则名称查找，多个采集规则使用 "," 分隔。
	LogConfigNames *string `json:"LogConfigNames,omitnil,omitempty" name:"LogConfigNames"`

	// 偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLogConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	delete(f, "LogConfigNames")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogConfigsResponseParams struct {
	// 分页查找时返回采集规则总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 指定采集规则名称查找，部分失败时返回失败采集规则名称及最后一个失败原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 采集规则查询结果
	LogConfigs *string `json:"LogConfigs,omitnil,omitempty" name:"LogConfigs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogConfigsResponseParams `json:"Response"`
}

func (r *DescribeLogConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSwitchesRequestParams struct {
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 集群类型，tke 或eks
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribeLogSwitchesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 集群类型，tke 或eks
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribeLogSwitchesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSwitchesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogSwitchesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogSwitchesResponseParams struct {
	// 集群日志开关集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchSet []*Switch `json:"SwitchSet,omitnil,omitempty" name:"SwitchSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogSwitchesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogSwitchesResponseParams `json:"Response"`
}

func (r *DescribeLogSwitchesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogSwitchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMasterComponentRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// master组件名称，支持kube-apiserver、kube-scheduler、kube-controller-manager
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`
}

type DescribeMasterComponentRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// master组件名称，支持kube-apiserver、kube-scheduler、kube-controller-manager
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`
}

func (r *DescribeMasterComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMasterComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Component")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMasterComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMasterComponentResponseParams struct {
	// master组件名称
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// master组件状态，三种状态：running、updating、shutdown
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMasterComponentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMasterComponentResponseParams `json:"Response"`
}

func (r *DescribeMasterComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMasterComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOSImagesRequestParams struct {

}

type DescribeOSImagesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOSImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOSImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOSImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOSImagesResponseParams struct {
	// 镜像信息列表
	OSImageSeriesSet []*OSImage `json:"OSImageSeriesSet,omitnil,omitempty" name:"OSImageSeriesSet"`

	// 镜像数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOSImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOSImagesResponseParams `json:"Response"`
}

func (r *DescribeOSImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOSImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPolicyListRequestParams struct {
	// 集群ID，请从容器服务集群列表获取（https://console.cloud.tencent.com/tke2/cluster）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 策略分类 基线：baseline 优选：priority 可选：optional
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

type DescribeOpenPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，请从容器服务集群列表获取（https://console.cloud.tencent.com/tke2/cluster）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 策略分类 基线：baseline 优选：priority 可选：optional
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

func (r *DescribeOpenPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Category")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpenPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpenPolicyListResponseParams struct {
	// 策略信息列表
	OpenPolicyInfoList []*OpenPolicyInfo `json:"OpenPolicyInfoList,omitnil,omitempty" name:"OpenPolicyInfoList"`

	// 集群内是否安装了gatekeeper addon
	GatekeeperStatus *int64 `json:"GatekeeperStatus,omitnil,omitempty" name:"GatekeeperStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOpenPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpenPolicyListResponseParams `json:"Response"`
}

func (r *DescribeOpenPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpenPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodChargeInfoRequestParams struct {
	// 集群 ID。TKE 集群可通过 [DescribeClusters](https://cloud.tencent.com/document/api/457/31862) 接口返回值中的ClusterId获取。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Pod的Uid，可以通过Uids 来批量查询，也可以通过 Namespace 和 Name 来查询某个 Pod 的计费信息。Uids 不传时，Namespace 和 Name 必须同时传。
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`
}

type DescribePodChargeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID。TKE 集群可通过 [DescribeClusters](https://cloud.tencent.com/document/api/457/31862) 接口返回值中的ClusterId获取。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Pod的Uid，可以通过Uids 来批量查询，也可以通过 Namespace 和 Name 来查询某个 Pod 的计费信息。Uids 不传时，Namespace 和 Name 必须同时传。
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`
}

func (r *DescribePodChargeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodChargeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "Name")
	delete(f, "Uids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePodChargeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodChargeInfoResponseParams struct {
	// Pod计费信息
	ChargeInfoSet []*PodChargeInfo `json:"ChargeInfoSet,omitnil,omitempty" name:"ChargeInfoSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePodChargeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePodChargeInfoResponseParams `json:"Response"`
}

func (r *DescribePodChargeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodChargeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodDeductionRateRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type DescribePodDeductionRateRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

func (r *DescribePodDeductionRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodDeductionRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "ClusterId")
	delete(f, "NodeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePodDeductionRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodDeductionRateResponseParams struct {
	// 各个规格的 可被预留券抵扣的Pod 抵扣率
	PodDeductionRateSet []*PodDeductionRate `json:"PodDeductionRateSet,omitnil,omitempty" name:"PodDeductionRateSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePodDeductionRateResponse struct {
	*tchttp.BaseResponse
	Response *DescribePodDeductionRateResponseParams `json:"Response"`
}

func (r *DescribePodDeductionRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodDeductionRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodsBySpecRequestParams struct {
	// 核数
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位：GiB
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 卡数，有0.25、0.5、1、2、4和8
	GpuNum *string `json:"GpuNum,omitnil,omitempty" name:"GpuNum"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// pod-type
	// 按照**【Pod 类型**】进行过滤。资源类型：intel、amd、v100、t4、a10\*gnv4、a10\*gnv4v等。
	// 类型：String
	// 必选：否
	// 
	// pod-deduct
	// 按照**【上个周期抵扣的Pod**】进行过滤。Values可不设置。
	// 必选：否
	// 
	// pod-not-deduct
	// 按照**【上个周期未抵扣的Pod**】进行过滤。Values可不设置。
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePodsBySpecRequest struct {
	*tchttp.BaseRequest
	
	// 核数
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位：GiB
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 卡数，有0.25、0.5、1、2、4和8
	GpuNum *string `json:"GpuNum,omitnil,omitempty" name:"GpuNum"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// pod-type
	// 按照**【Pod 类型**】进行过滤。资源类型：intel、amd、v100、t4、a10\*gnv4、a10\*gnv4v等。
	// 类型：String
	// 必选：否
	// 
	// pod-deduct
	// 按照**【上个周期抵扣的Pod**】进行过滤。Values可不设置。
	// 必选：否
	// 
	// pod-not-deduct
	// 按照**【上个周期未抵扣的Pod**】进行过滤。Values可不设置。
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePodsBySpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodsBySpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "GpuNum")
	delete(f, "Zone")
	delete(f, "ClusterId")
	delete(f, "NodeName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePodsBySpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePodsBySpecResponseParams struct {
	// Pod 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Pod 节点信息
	PodSet []*PodNodeInfo `json:"PodSet,omitnil,omitempty" name:"PodSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePodsBySpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribePodsBySpecResponseParams `json:"Response"`
}

func (r *DescribePodsBySpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePodsBySpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostNodeResourcesRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	//  节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type DescribePostNodeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	//  节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

func (r *DescribePostNodeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostNodeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePostNodeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePostNodeResourcesResponseParams struct {
	// Pod详情
	PodSet []*SuperNodeResource `json:"PodSet,omitnil,omitempty" name:"PodSet"`

	// 预留券详情
	ReservedInstanceSet []*SuperNodeResource `json:"ReservedInstanceSet,omitnil,omitempty" name:"ReservedInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePostNodeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePostNodeResourcesResponseParams `json:"Response"`
}

func (r *DescribePostNodeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePostNodeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentInstancesRequestParams struct {
	// 集群id
	// 可以是tke, eks, edge的集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribePrometheusAgentInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	// 可以是tke, eks, edge的集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribePrometheusAgentInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAgentInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentInstancesResponseParams struct {
	// 关联该集群的实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*string `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAgentInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAgentInstancesResponseParams `json:"Response"`
}

func (r *DescribePrometheusAgentInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentsRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusAgentsRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAgentsResponseParams struct {
	// 被关联集群信息
	Agents []*PrometheusAgentOverview `json:"Agents,omitnil,omitempty" name:"Agents"`

	// 被关联集群总量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAgentsResponseParams `json:"Response"`
}

func (r *DescribePrometheusAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertHistoryRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// label集合
	Labels *string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 分片
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分片
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusAlertHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// label集合
	Labels *string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 分片
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分片
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusAlertHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RuleName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Labels")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAlertHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertHistoryResponseParams struct {
	// 告警历史
	Items []*PrometheusAlertHistoryItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAlertHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAlertHistoryResponseParams `json:"Response"`
}

func (r *DescribePrometheusAlertHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤
	// 支持ID，Name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤
	// 支持ID，Name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertPolicyResponseParams struct {
	// 告警详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertRules []*PrometheusAlertPolicyItem `json:"AlertRules,omitnil,omitempty" name:"AlertRules"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *DescribePrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertRuleRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤
	// 支持ID，Name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤
	// 支持ID，Name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusAlertRuleResponseParams struct {
	// 告警详情
	AlertRules []*PrometheusAlertRuleDetail `json:"AlertRules,omitnil,omitempty" name:"AlertRules"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *DescribePrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusClusterAgentsRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusClusterAgentsRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusClusterAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusClusterAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusClusterAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusClusterAgentsResponseParams struct {
	// 被关联集群信息
	Agents []*PrometheusAgentOverview `json:"Agents,omitnil,omitempty" name:"Agents"`

	// 被关联集群总量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusClusterAgentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusClusterAgentsResponseParams `json:"Response"`
}

func (r *DescribePrometheusClusterAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusClusterAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type DescribePrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *DescribePrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusConfigResponseParams struct {
	// 全局配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// ServiceMonitor配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitor配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 原生Job
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Probe配置
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusConfigResponseParams `json:"Response"`
}

func (r *DescribePrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalConfigRequestParams struct {
	// 实例级别抓取配置
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否禁用统计
	DisableStatistics *bool `json:"DisableStatistics,omitnil,omitempty" name:"DisableStatistics"`
}

type DescribePrometheusGlobalConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例级别抓取配置
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否禁用统计
	DisableStatistics *bool `json:"DisableStatistics,omitnil,omitempty" name:"DisableStatistics"`
}

func (r *DescribePrometheusGlobalConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DisableStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusGlobalConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalConfigResponseParams struct {
	// 配置内容
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// ServiceMonitors列表以及对应targets信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitors列表以及对应targets信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// RawJobs列表以及对应targets信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusGlobalConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusGlobalConfigResponseParams `json:"Response"`
}

func (r *DescribePrometheusGlobalConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalNotificationRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusGlobalNotificationResponseParams struct {
	// 全局告警通知渠道
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *DescribePrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceInitStatusRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceInitStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceInitStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceInitStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceInitStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceInitStatusResponseParams struct {
	// 实例初始化状态，取值：
	// uninitialized 未初始化 
	// initializing 初始化中
	// running 初始化完成，运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 初始化任务步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*TaskStepInfo `json:"Steps,omitnil,omitempty" name:"Steps"`

	// 实例eks集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EksClusterId *string `json:"EksClusterId,omitnil,omitempty" name:"EksClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceInitStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceInitStatusResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceInitStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceInitStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribePrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribePrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstanceResponseParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 私有网络id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// cos桶名称
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// 数据查询地址
	QueryAddress *string `json:"QueryAddress,omitnil,omitempty" name:"QueryAddress"`

	// 实例中grafana相关的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grafana *PrometheusGrafanaInfo `json:"Grafana,omitnil,omitempty" name:"Grafana"`

	// 用户自定义alertmanager
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertManagerUrl *string `json:"AlertManagerUrl,omitnil,omitempty" name:"AlertManagerUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstanceResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesOverviewRequestParams struct {
	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤实例，目前支持：
	// ID: 通过实例ID来过滤 
	// Name: 通过实例名称来过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusInstancesOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤实例，目前支持：
	// ID: 通过实例ID来过滤 
	// Name: 通过实例名称来过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusInstancesOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusInstancesOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusInstancesOverviewResponseParams struct {
	// 实例列表
	Instances []*PrometheusInstancesOverview `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusInstancesOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusInstancesOverviewResponseParams `json:"Response"`
}

func (r *DescribePrometheusInstancesOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusInstancesOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusOverviewsRequestParams struct {
	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤实例，目前支持：
	// ID: 通过实例ID来过滤 
	// Name: 通过实例名称来过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusOverviewsRequest struct {
	*tchttp.BaseRequest
	
	// 用于分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 用于分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤实例，目前支持：
	// ID: 通过实例ID来过滤 
	// Name: 通过实例名称来过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusOverviewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusOverviewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusOverviewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusOverviewsResponseParams struct {
	// 实例列表
	Instances []*PrometheusInstanceOverview `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusOverviewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusOverviewsResponseParams `json:"Response"`
}

func (r *DescribePrometheusOverviewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusOverviewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRulesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusRecordRulesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusRecordRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusRecordRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusRecordRulesResponseParams struct {
	// 聚合规则
	Records []*PrometheusRecordRuleYamlItem `json:"Records,omitnil,omitempty" name:"Records"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusRecordRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusRecordRulesResponseParams `json:"Response"`
}

func (r *DescribePrometheusRecordRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusRecordRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTargetsRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 过滤条件，当前支持
	// Name=state
	// Value=up, down, unknown
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePrometheusTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 过滤条件，当前支持
	// Name=state
	// Value=up, down, unknown
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePrometheusTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTargetsResponseParams struct {
	// 所有Job的targets信息
	Jobs []*PrometheusJobTargets `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTargetsResponseParams `json:"Response"`
}

func (r *DescribePrometheusTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempRequestParams struct {
	// 模糊过滤条件，支持
	// Level 按模板级别过滤
	// Name 按名称过滤
	// Describe 按描述过滤
	// ID 按templateId过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 总数限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模糊过滤条件，支持
	// Level 按模板级别过滤
	// Name 按名称过滤
	// Describe 按描述过滤
	// ID 按templateId过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 总数限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempResponseParams struct {
	// 模板列表
	Templates []*PrometheusTemp `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTempResponseParams `json:"Response"`
}

func (r *DescribePrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempSyncRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribePrometheusTempSyncRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribePrometheusTempSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTempSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTempSyncResponseParams struct {
	// 同步目标详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTempSyncResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTempSyncResponseParams `json:"Response"`
}

func (r *DescribePrometheusTempSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTempSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTemplateSyncRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribePrometheusTemplateSyncRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribePrometheusTemplateSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTemplateSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTemplateSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTemplateSyncResponseParams struct {
	// 同步目标详情
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTemplateSyncResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTemplateSyncResponseParams `json:"Response"`
}

func (r *DescribePrometheusTemplateSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTemplateSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTemplatesRequestParams struct {
	// 模糊过滤条件，支持
	// Level 按模板级别过滤
	// Name 按名称过滤
	// Describe 按描述过滤
	// ID 按templateId过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 总数限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribePrometheusTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 模糊过滤条件，支持
	// Level 按模板级别过滤
	// Name 按名称过滤
	// Describe 按描述过滤
	// ID 按templateId过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 总数限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribePrometheusTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrometheusTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrometheusTemplatesResponseParams struct {
	// 模板列表
	Templates []*PrometheusTemplate `json:"Templates,omitnil,omitempty" name:"Templates"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePrometheusTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrometheusTemplatesResponseParams `json:"Response"`
}

func (r *DescribePrometheusTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrometheusTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRIUtilizationDetailRequestParams struct {
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// reserved-instance-id
	// 按照**【预留实例ID**】进行过滤。预留实例ID形如：eksri-xxxxxxxx。
	// 类型：String
	// 必选：否
	// 
	// begin-time
	// 按照**【抵扣开始时间**】进行过滤。形如：2023-06-28 15:27:40。
	// 类型：String
	// 必选：否
	// 
	// end-time
	// 按照**【抵扣结束时间**】进行过滤。形如：2023-06-28 15:27:40。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRIUtilizationDetailRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// reserved-instance-id
	// 按照**【预留实例ID**】进行过滤。预留实例ID形如：eksri-xxxxxxxx。
	// 类型：String
	// 必选：否
	// 
	// begin-time
	// 按照**【抵扣开始时间**】进行过滤。形如：2023-06-28 15:27:40。
	// 类型：String
	// 必选：否
	// 
	// end-time
	// 按照**【抵扣结束时间**】进行过滤。形如：2023-06-28 15:27:40。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRIUtilizationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRIUtilizationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRIUtilizationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRIUtilizationDetailResponseParams struct {
	// 总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 详情。
	RIUtilizationDetailSet []*RIUtilizationDetail `json:"RIUtilizationDetailSet,omitnil,omitempty" name:"RIUtilizationDetailSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRIUtilizationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRIUtilizationDetailResponseParams `json:"Response"`
}

func (r *DescribeRIUtilizationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRIUtilizationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 地域的数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 地域列表
	RegionInstanceSet []*RegionInstance `json:"RegionInstanceSet,omitnil,omitempty" name:"RegionInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstanceUtilizationRateRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type DescribeReservedInstanceUtilizationRateRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

func (r *DescribeReservedInstanceUtilizationRateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstanceUtilizationRateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "ClusterId")
	delete(f, "NodeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedInstanceUtilizationRateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstanceUtilizationRateResponseParams struct {
	// 预留券使用率
	UtilizationRateSet []*ReservedInstanceUtilizationRate `json:"UtilizationRateSet,omitnil,omitempty" name:"UtilizationRateSet"`

	// 按量计费的 Pod 总数
	PodNum *uint64 `json:"PodNum,omitnil,omitempty" name:"PodNum"`

	//  Pod 被预留券抵扣的抵扣率
	PodRate *float64 `json:"PodRate,omitnil,omitempty" name:"PodRate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReservedInstanceUtilizationRateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedInstanceUtilizationRateResponseParams `json:"Response"`
}

func (r *DescribeReservedInstanceUtilizationRateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstanceUtilizationRateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesRequestParams struct {
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// status
	// 按照**【状态**】进行过滤。状态：Creating：创建中、Active：生效中、Expired：已过期、Refunded：已退还。
	// 类型：String
	// 必选：否
	// 
	// resource-type
	// 按照**【资源类型**】进行过滤。资源类型：common、amd、v100、t4、a10\*gnv4、a10\*gnv4v等，common表示通用类型。
	// 类型：String
	// 必选：否
	// 
	// cpu
	// 按照**【核数**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// memory
	// 按照**【内存**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// gpu
	// 按照**【GPU卡数**】进行过滤，取值有0.25、0.5、1、2、4等。
	// 类型：String
	// 必选：否
	// 
	// cluster-id
	// 按照**【集群ID**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// node-name
	// 按照**【节点名称**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// scope
	// 按照**【可用区**】进行过滤。比如：ap-guangzhou-2，为空字符串表示地域抵扣范围。如果只过滤可用区抵扣范围，需要同时将cluster-id、node-name设置为空字符串。
	// 类型：String
	// 必选：否
	// 
	// reserved-instance-id
	// 按照**【预留实例ID**】进行过滤。预留实例ID形如：eksri-xxxxxxxx。
	// 类型：String
	// 必选：否
	// 
	// reserved-instance-name
	// 按照**【预留实例名称**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// reserved-instance-deduct
	// 按照**【上个周期抵扣的预留券**】进行过滤。Values可不设置。
	// 必选：否
	// 
	// reserved-instance-not-deduct
	// 按照**【上个周期未抵扣的预留券**】进行过滤。Values可不设置。
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段。支持CreatedAt、ActiveAt、ExpireAt。默认值CreatedAt。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方法。顺序：ASC，倒序：DESC。默认值DESC。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeReservedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// status
	// 按照**【状态**】进行过滤。状态：Creating：创建中、Active：生效中、Expired：已过期、Refunded：已退还。
	// 类型：String
	// 必选：否
	// 
	// resource-type
	// 按照**【资源类型**】进行过滤。资源类型：common、amd、v100、t4、a10\*gnv4、a10\*gnv4v等，common表示通用类型。
	// 类型：String
	// 必选：否
	// 
	// cpu
	// 按照**【核数**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// memory
	// 按照**【内存**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// gpu
	// 按照**【GPU卡数**】进行过滤，取值有0.25、0.5、1、2、4等。
	// 类型：String
	// 必选：否
	// 
	// cluster-id
	// 按照**【集群ID**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// node-name
	// 按照**【节点名称**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// scope
	// 按照**【可用区**】进行过滤。比如：ap-guangzhou-2，为空字符串表示地域抵扣范围。如果只过滤可用区抵扣范围，需要同时将cluster-id、node-name设置为空字符串。
	// 类型：String
	// 必选：否
	// 
	// reserved-instance-id
	// 按照**【预留实例ID**】进行过滤。预留实例ID形如：eksri-xxxxxxxx。
	// 类型：String
	// 必选：否
	// 
	// reserved-instance-name
	// 按照**【预留实例名称**】进行过滤。
	// 类型：String
	// 必选：否
	// 
	// reserved-instance-deduct
	// 按照**【上个周期抵扣的预留券**】进行过滤。Values可不设置。
	// 必选：否
	// 
	// reserved-instance-not-deduct
	// 按照**【上个周期未抵扣的预留券**】进行过滤。Values可不设置。
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段。支持CreatedAt、ActiveAt、ExpireAt。默认值CreatedAt。
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序方法。顺序：ASC，倒序：DESC。默认值DESC。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeReservedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReservedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReservedInstancesResponseParams struct {
	// 总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 预留实例列表。
	ReservedInstanceSet []*ReservedInstance `json:"ReservedInstanceSet,omitnil,omitempty" name:"ReservedInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReservedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReservedInstancesResponseParams `json:"Response"`
}

func (r *DescribeReservedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReservedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeResourceUsageRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeResourceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageResponseParams struct {
	// CRD使用量
	CRDUsage *ResourceUsage `json:"CRDUsage,omitnil,omitempty" name:"CRDUsage"`

	// Pod使用量，单位：个数
	PodUsage *uint64 `json:"PodUsage,omitnil,omitempty" name:"PodUsage"`

	// ReplicaSet使用量，单位：个数
	RSUsage *uint64 `json:"RSUsage,omitnil,omitempty" name:"RSUsage"`

	// ConfigMap使用量，单位：个数
	ConfigMapUsage *uint64 `json:"ConfigMapUsage,omitnil,omitempty" name:"ConfigMapUsage"`

	// 其他资源使用量
	OtherUsage *ResourceUsage `json:"OtherUsage,omitnil,omitempty" name:"OtherUsage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceUsageResponseParams `json:"Response"`
}

func (r *DescribeResourceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollOutSequencesRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRollOutSequencesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRollOutSequencesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollOutSequencesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollOutSequencesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollOutSequencesResponseParams struct {
	// 发布序列
	Sequences []*RollOutSequence `json:"Sequences,omitnil,omitempty" name:"Sequences"`

	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRollOutSequencesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollOutSequencesResponseParams `json:"Response"`
}

func (r *DescribeRollOutSequencesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollOutSequencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTableConflictsRequestParams struct {
	// 路由表CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitnil,omitempty" name:"RouteTableCidrBlock"`

	// 路由表绑定的VPC，请到 [VPC 控制台](https://console.cloud.tencent.com/vpc/vpc)复制 VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DescribeRouteTableConflictsRequest struct {
	*tchttp.BaseRequest
	
	// 路由表CIDR
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitnil,omitempty" name:"RouteTableCidrBlock"`

	// 路由表绑定的VPC，请到 [VPC 控制台](https://console.cloud.tencent.com/vpc/vpc)复制 VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *DescribeRouteTableConflictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTableConflictsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableCidrBlock")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTableConflictsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTableConflictsResponseParams struct {
	// 路由表是否冲突。
	HasConflict *bool `json:"HasConflict,omitnil,omitempty" name:"HasConflict"`

	// 路由表冲突列表。
	RouteTableConflictSet []*RouteTableConflict `json:"RouteTableConflictSet,omitnil,omitempty" name:"RouteTableConflictSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRouteTableConflictsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteTableConflictsResponseParams `json:"Response"`
}

func (r *DescribeRouteTableConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTableConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedRuntimeRequestParams struct {
	// K8S版本
	K8sVersion *string `json:"K8sVersion,omitnil,omitempty" name:"K8sVersion"`
}

type DescribeSupportedRuntimeRequest struct {
	*tchttp.BaseRequest
	
	// K8S版本
	K8sVersion *string `json:"K8sVersion,omitnil,omitempty" name:"K8sVersion"`
}

func (r *DescribeSupportedRuntimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedRuntimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "K8sVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedRuntimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedRuntimeResponseParams struct {
	// 可选运行时列表
	OptionalRuntimes []*OptionalRuntimes `json:"OptionalRuntimes,omitnil,omitempty" name:"OptionalRuntimes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSupportedRuntimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportedRuntimeResponseParams `json:"Response"`
}

func (r *DescribeSupportedRuntimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedRuntimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterCredentialRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeTKEEdgeClusterCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeTKEEdgeClusterCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeClusterCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterCredentialResponseParams struct {
	// 集群的接入地址信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Addresses []*IPAddress `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// 集群的认证信息
	Credential *ClusterCredential `json:"Credential,omitnil,omitempty" name:"Credential"`

	// 集群的公网访问信息
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 集群的内网访问信息
	InternalLB *EdgeClusterInternalLB `json:"InternalLB,omitnil,omitempty" name:"InternalLB"`

	// 集群的CoreDns部署信息
	CoreDns *string `json:"CoreDns,omitnil,omitempty" name:"CoreDns"`

	// 集群的健康检查多地域部署信息
	HealthRegion *string `json:"HealthRegion,omitnil,omitempty" name:"HealthRegion"`

	// 集群的健康检查部署信息
	Health *string `json:"Health,omitnil,omitempty" name:"Health"`

	// 是否部署GridDaemon以支持headless service
	GridDaemon *string `json:"GridDaemon,omitnil,omitempty" name:"GridDaemon"`

	// 公网访问kins集群
	UnitCluster *string `json:"UnitCluster,omitnil,omitempty" name:"UnitCluster"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeClusterCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeClusterCredentialResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeClusterCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterStatusRequestParams struct {
	// 边缘计算容器集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeTKEEdgeClusterStatusRequest struct {
	*tchttp.BaseRequest
	
	// 边缘计算容器集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeTKEEdgeClusterStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeClusterStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClusterStatusResponseParams struct {
	// 集群当前状态
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 集群过程数组
	Conditions []*ClusterCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeClusterStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeClusterStatusResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeClusterStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClustersRequestParams struct {
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件,当前只支持按照ClusterName和云标签进行过滤,云标签过滤格式Tags:["key1:value1","key2:value2"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTKEEdgeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 偏移量,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件,当前只支持按照ClusterName和云标签进行过滤,云标签过滤格式Tags:["key1:value1","key2:value2"]
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTKEEdgeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeClustersResponseParams struct {
	// 集群总个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群信息列表
	Clusters []*EdgeCluster `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeClustersResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeExternalKubeconfigRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeTKEEdgeExternalKubeconfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeTKEEdgeExternalKubeconfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeExternalKubeconfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeExternalKubeconfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeExternalKubeconfigResponseParams struct {
	// kubeconfig文件内容
	Kubeconfig *string `json:"Kubeconfig,omitnil,omitempty" name:"Kubeconfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeExternalKubeconfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeExternalKubeconfigResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeExternalKubeconfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeExternalKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeScriptRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 网卡名,指定边缘节点上kubelet向apiserver注册使用的网卡
	Interface *string `json:"Interface,omitnil,omitempty" name:"Interface"`

	// 节点名字
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// json格式的节点配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 可以下载某个历史版本的edgectl脚本，默认下载最新版本，edgectl版本信息可以在脚本里查看
	ScriptVersion *string `json:"ScriptVersion,omitnil,omitempty" name:"ScriptVersion"`
}

type DescribeTKEEdgeScriptRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 网卡名,指定边缘节点上kubelet向apiserver注册使用的网卡
	Interface *string `json:"Interface,omitnil,omitempty" name:"Interface"`

	// 节点名字
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// json格式的节点配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 可以下载某个历史版本的edgectl脚本，默认下载最新版本，edgectl版本信息可以在脚本里查看
	ScriptVersion *string `json:"ScriptVersion,omitnil,omitempty" name:"ScriptVersion"`
}

func (r *DescribeTKEEdgeScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Interface")
	delete(f, "NodeName")
	delete(f, "Config")
	delete(f, "ScriptVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTKEEdgeScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTKEEdgeScriptResponseParams struct {
	// 下载链接
	Link *string `json:"Link,omitnil,omitempty" name:"Link"`

	// 下载需要的token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 下载命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// edgectl脚本版本，默认拉取最新版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptVersion *string `json:"ScriptVersion,omitnil,omitempty" name:"ScriptVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTKEEdgeScriptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTKEEdgeScriptResponseParams `json:"Response"`
}

func (r *DescribeTKEEdgeScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTKEEdgeScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 根据filter做过滤，支持ClusterId（取值示例：cls-xxxx）、TaskType（任务类型，取值示例：add_cluster_cidr、node_upgrade、node_upgrade_ctl等）其中任务类型必传
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 表示最新的任务条目，此值为true的话，输出任务列表中只会有最新的一条
	Latest *bool `json:"Latest,omitnil,omitempty" name:"Latest"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 根据filter做过滤，支持ClusterId（取值示例：cls-xxxx）、TaskType（任务类型，取值示例：add_cluster_cidr、node_upgrade、node_upgrade_ctl等）其中任务类型必传
	Filter []*Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 表示最新的任务条目，此值为true的话，输出任务列表中只会有最新的一条
	Latest *bool `json:"Latest,omitnil,omitempty" name:"Latest"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Latest")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务步骤信息
	Tasks []*Task `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpgradeTaskDetailRequestParams struct {
	// 升级任务ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUpgradeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 升级任务ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUpgradeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpgradeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpgradeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpgradeTaskDetailResponseParams struct {
	// 升级计划
	UpgradePlans []*UpgradePlan `json:"UpgradePlans,omitnil,omitempty" name:"UpgradePlans"`

	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUpgradeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpgradeTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeUpgradeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpgradeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpgradeTasksRequestParams struct {
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUpgradeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 最大输出条目数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUpgradeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpgradeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpgradeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpgradeTasksResponseParams struct {
	// 升级任务
	UpgradeTasks []*UpgradeTask `json:"UpgradeTasks,omitnil,omitempty" name:"UpgradeTasks"`

	// 总条目数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUpgradeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpgradeTasksResponseParams `json:"Response"`
}

func (r *DescribeUpgradeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpgradeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionsRequestParams struct {

}

type DescribeVersionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVersionsResponseParams struct {
	// 版本数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 版本列表
	VersionInstanceSet []*VersionInstance `json:"VersionInstanceSet,omitnil,omitempty" name:"VersionInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVersionsResponseParams `json:"Response"`
}

func (r *DescribeVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcCniPodLimitsRequestParams struct {
	// 查询的机型所在可用区，如：ap-guangzhou-3，默认为空，即不按可用区过滤信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 查询的实例机型系列信息，如：S5，默认为空，即不按机型系列过滤信息
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 查询的实例机型信息，如：S5.LARGE8，默认为空，即不按机型过滤信息
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type DescribeVpcCniPodLimitsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的机型所在可用区，如：ap-guangzhou-3，默认为空，即不按可用区过滤信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 查询的实例机型系列信息，如：S5，默认为空，即不按机型系列过滤信息
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 查询的实例机型信息，如：S5.LARGE8，默认为空，即不按机型过滤信息
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *DescribeVpcCniPodLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcCniPodLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstanceFamily")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcCniPodLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcCniPodLimitsResponseParams struct {
	// 机型数据数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 机型信息及其可支持的最大VPC-CNI模式Pod数量信息
	PodLimitsInstanceSet []*PodLimitsInstance `json:"PodLimitsInstanceSet,omitnil,omitempty" name:"PodLimitsInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcCniPodLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcCniPodLimitsResponseParams `json:"Response"`
}

func (r *DescribeVpcCniPodLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcCniPodLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClusterAuditRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 取值为true代表关闭集群审计时删除默认创建的日志集和主题，false代表不删除
	DeleteLogSetAndTopic *bool `json:"DeleteLogSetAndTopic,omitnil,omitempty" name:"DeleteLogSetAndTopic"`
}

type DisableClusterAuditRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 取值为true代表关闭集群审计时删除默认创建的日志集和主题，false代表不删除
	DeleteLogSetAndTopic *bool `json:"DeleteLogSetAndTopic,omitnil,omitempty" name:"DeleteLogSetAndTopic"`
}

func (r *DisableClusterAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClusterAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DeleteLogSetAndTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableClusterAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClusterAuditResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableClusterAuditResponse struct {
	*tchttp.BaseResponse
	Response *DisableClusterAuditResponseParams `json:"Response"`
}

func (r *DisableClusterAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClusterAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClusterDeletionProtectionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DisableClusterDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DisableClusterDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClusterDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableClusterDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableClusterDeletionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableClusterDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DisableClusterDeletionProtectionResponseParams `json:"Response"`
}

func (r *DisableClusterDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableClusterDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableControlPlaneLogsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型。当前只支持tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 组件名称列表，目前支持的组件有：cluster-autoscaler、kapenter
	ComponentNames []*string `json:"ComponentNames,omitnil,omitempty" name:"ComponentNames"`

	// 是否删除日志集和topic。 如果日志集和topic被其他采集规则使用，则不会被删除
	DeleteLogSetAndTopic *bool `json:"DeleteLogSetAndTopic,omitnil,omitempty" name:"DeleteLogSetAndTopic"`
}

type DisableControlPlaneLogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型。当前只支持tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 组件名称列表，目前支持的组件有：cluster-autoscaler、kapenter
	ComponentNames []*string `json:"ComponentNames,omitnil,omitempty" name:"ComponentNames"`

	// 是否删除日志集和topic。 如果日志集和topic被其他采集规则使用，则不会被删除
	DeleteLogSetAndTopic *bool `json:"DeleteLogSetAndTopic,omitnil,omitempty" name:"DeleteLogSetAndTopic"`
}

func (r *DisableControlPlaneLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableControlPlaneLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	delete(f, "ComponentNames")
	delete(f, "DeleteLogSetAndTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableControlPlaneLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableControlPlaneLogsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableControlPlaneLogsResponse struct {
	*tchttp.BaseResponse
	Response *DisableControlPlaneLogsResponseParams `json:"Response"`
}

func (r *DisableControlPlaneLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableControlPlaneLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableEncryptionProtectionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DisableEncryptionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DisableEncryptionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableEncryptionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableEncryptionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableEncryptionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableEncryptionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DisableEncryptionProtectionResponseParams `json:"Response"`
}

func (r *DisableEncryptionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableEncryptionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableEventPersistenceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 取值为true代表关闭集群审计时删除默认创建的日志集和主题，false代表不删除
	DeleteLogSetAndTopic *bool `json:"DeleteLogSetAndTopic,omitnil,omitempty" name:"DeleteLogSetAndTopic"`
}

type DisableEventPersistenceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 取值为true代表关闭集群审计时删除默认创建的日志集和主题，false代表不删除
	DeleteLogSetAndTopic *bool `json:"DeleteLogSetAndTopic,omitnil,omitempty" name:"DeleteLogSetAndTopic"`
}

func (r *DisableEventPersistenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableEventPersistenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DeleteLogSetAndTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableEventPersistenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableEventPersistenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableEventPersistenceResponse struct {
	*tchttp.BaseResponse
	Response *DisableEventPersistenceResponseParams `json:"Response"`
}

func (r *DisableEventPersistenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableEventPersistenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableVpcCniNetworkTypeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DisableVpcCniNetworkTypeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DisableVpcCniNetworkTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableVpcCniNetworkTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableVpcCniNetworkTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableVpcCniNetworkTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableVpcCniNetworkTypeResponse struct {
	*tchttp.BaseResponse
	Response *DisableVpcCniNetworkTypeResponseParams `json:"Response"`
}

func (r *DisableVpcCniNetworkTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableVpcCniNetworkTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsServerConf struct {
	// 域名。空字符串表示所有域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// dns 服务器地址列表。地址格式 ip:port
	DnsServers []*string `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`
}

// Predefined struct for user
type DrainClusterVirtualNodeRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type DrainClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

func (r *DrainClusterVirtualNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DrainClusterVirtualNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DrainClusterVirtualNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DrainClusterVirtualNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DrainClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse
	Response *DrainClusterVirtualNodeResponseParams `json:"Response"`
}

func (r *DrainClusterVirtualNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DrainClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriverVersion struct {
	// GPU驱动或者CUDA的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// GPU驱动或者CUDA的版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type ECMEnhancedService struct {
	// 是否开启云监控服务
	SecurityService *ECMRunMonitorServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// 是否开启云镜服务
	MonitorService *ECMRunSecurityServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`
}

type ECMRunMonitorServiceEnabled struct {
	// 是否开启
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ECMRunSecurityServiceEnabled struct {
	// 是否开启
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 云镜版本：0 基础版，1 专业版
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`
}

type ECMZoneInstanceCountISP struct {
	// 创建实例的可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 在当前可用区欲创建的实例数目
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 运营商
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`
}

type EdgeArgsFlag struct {
	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数描述
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 参数默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数可选范围（目前包含range和in两种，"[]"代表range，如"[1, 5]"表示参数必须>=1且 <=5, "()"代表in， 如"('aa', 'bb')"表示参数只能为字符串'aa'或者'bb'，该参数为空表示不校验）
	Constraint *string `json:"Constraint,omitnil,omitempty" name:"Constraint"`
}

type EdgeAvailableExtraArgs struct {
	// kube-apiserver可用的自定义参数
	KubeAPIServer []*EdgeArgsFlag `json:"KubeAPIServer,omitnil,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager可用的自定义参数
	KubeControllerManager []*EdgeArgsFlag `json:"KubeControllerManager,omitnil,omitempty" name:"KubeControllerManager"`

	// kube-scheduler可用的自定义参数
	KubeScheduler []*EdgeArgsFlag `json:"KubeScheduler,omitnil,omitempty" name:"KubeScheduler"`

	// kubelet可用的自定义参数
	Kubelet []*EdgeArgsFlag `json:"Kubelet,omitnil,omitempty" name:"Kubelet"`
}

type EdgeCluster struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 集群pod cidr
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 集群 service cidr
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// k8s 版本号
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// 集群状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 边缘集群版本
	EdgeClusterVersion *string `json:"EdgeClusterVersion,omitnil,omitempty" name:"EdgeClusterVersion"`

	// 节点最大Pod数
	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitnil,omitempty" name:"MaxNodePodNum"`

	// 集群高级设置
	ClusterAdvancedSettings *EdgeClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitnil,omitempty" name:"ClusterAdvancedSettings"`

	// 边缘容器集群级别
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 是否支持自动提升集群配置
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 集群付费模式，支持POSTPAID_BY_HOUR或者PREPAID
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 边缘集群组件的版本
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`

	// 集群绑定的云标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification *TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

type EdgeClusterAdvancedSettings struct {
	// 集群自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraArgs *EdgeClusterExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// 运行时类型，支持"docker"和"containerd"，默认为docker
	Runtime *string `json:"Runtime,omitnil,omitempty" name:"Runtime"`

	// 集群kube-proxy转发模式，支持"iptables"和"ipvs"，默认为iptables
	ProxyMode *string `json:"ProxyMode,omitnil,omitempty" name:"ProxyMode"`
}

type EdgeClusterExtraArgs struct {
	// kube-apiserver自定义参数，参数格式为["k1=v1", "k1=v2"]， 例如["max-requests-inflight=500","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"]
	KubeAPIServer []*string `json:"KubeAPIServer,omitnil,omitempty" name:"KubeAPIServer"`

	// kube-controller-manager自定义参数
	KubeControllerManager []*string `json:"KubeControllerManager,omitnil,omitempty" name:"KubeControllerManager"`

	// kube-scheduler自定义参数
	KubeScheduler []*string `json:"KubeScheduler,omitnil,omitempty" name:"KubeScheduler"`
}

type EdgeClusterInternalLB struct {
	// 是否开启内网访问LB
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 内网访问LB关联的子网Id
	SubnetId []*string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type EdgeClusterPublicLB struct {
	// 是否开启公网访问LB
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 允许访问的公网cidr
	AllowFromCidrs []*string `json:"AllowFromCidrs,omitnil,omitempty" name:"AllowFromCidrs"`
}

type EipAttribute struct {
	// 容器实例删除后，EIP是否释放。
	// Never表示不释放，其他任意值（包括空字符串）表示释放。
	DeletePolicy *string `json:"DeletePolicy,omitnil,omitempty" name:"DeletePolicy"`

	// EIP线路类型。默认值：BGP。
	// 已开通静态单线IP白名单的用户，可选值：
	// CMCC：中国移动
	// CTCC：中国电信
	// CUCC：中国联通
	// 注意：仅部分地域支持静态单线IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil,omitempty" name:"InternetServiceProvider"`

	// EIP出带宽上限，单位：Mbps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`
}

type EksCi struct {
	// 自动为用户创建的EipId
	AutoCreatedEipId *string `json:"AutoCreatedEipId,omitnil,omitempty" name:"AutoCreatedEipId"`

	// 为容器实例关联 CAM 角色，value 填写 CAM 角色名称，容器实例可获取该 CAM 角色包含的权限策略，方便 容器实例 内的程序进行如购买资源、读写存储等云资源操作。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 容器列表
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// CPU大小
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// CPU类型
	CpuType *string `json:"CpuType,omitnil,omitempty" name:"CpuType"`

	// 接到请求后的系统创建时间。
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 容器实例绑定的Eip地址，注意可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddress *string `json:"EipAddress,omitnil,omitempty" name:"EipAddress"`

	// EKS Container Instance Id
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// EKS Container Instance Name
	EksCiName *string `json:"EksCiName,omitnil,omitempty" name:"EksCiName"`

	// 数据卷信息
	EksCiVolume *EksCiVolume `json:"EksCiVolume,omitnil,omitempty" name:"EksCiVolume"`

	// GPU卡数量
	GpuCount *uint64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// GPU类型。如无使用GPU则不返回
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`

	// 初始化容器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 内存大小
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 容器状态是否持久化
	// 注意：此字段可能返回 null，表示取不到有效值。
	PersistStatus *bool `json:"PersistStatus,omitnil,omitempty" name:"PersistStatus"`

	// 内网ip地址
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 容器组的重启策略
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 容器组运行的安全上下文
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityContext *SecurityContext `json:"SecurityContext,omitnil,omitempty" name:"SecurityContext"`

	// 安全组ID
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 返回容器组创建状态：Pending，Running，Succeeded，Failed。其中：
	// Failed （运行失败）指的容器组退出，RestartPolilcy为Never， 有容器exitCode非0；
	// Succeeded（运行成功）指的是容器组退出了，RestartPolicy为Never或onFailure，所有容器exitCode都为0；
	// Failed和Succeeded这两种状态都会停止运行，停止计费。
	// Pending是创建中，Running是 运行中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例所属子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 容器全部成功退出后的时间
	SucceededTime *string `json:"SucceededTime,omitnil,omitempty" name:"SucceededTime"`

	// 实例所属VPC的Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type EksCiRegionInfo struct {
	// 地域别名，形如gz
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 地域名，形如ap-guangzhou
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type EksCiVolume struct {
	// Cbs Volume
	// 注意：此字段可能返回 null，表示取不到有效值。
	CbsVolumes []*CbsVolume `json:"CbsVolumes,omitnil,omitempty" name:"CbsVolumes"`

	// Nfs Volume
	// 注意：此字段可能返回 null，表示取不到有效值。
	NfsVolumes []*NfsVolume `json:"NfsVolumes,omitnil,omitempty" name:"NfsVolumes"`
}

type EksCluster struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// Vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// k8s 版本号
	K8SVersion *string `json:"K8SVersion,omitnil,omitempty" name:"K8SVersion"`

	// 集群状态(running 运行中，initializing 初始化中，Failed 异常，Idling 空闲中，Activating 激活中，Terminating 删除中)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Service 子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// 集群的自定义dns 服务器信息
	DnsServers []*DnsServerConf `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// 将来删除集群时是否要删除cbs。默认为 FALSE
	NeedDeleteCbs *bool `json:"NeedDeleteCbs,omitnil,omitempty" name:"NeedDeleteCbs"`

	// 是否在用户集群内开启Dns。默认为TRUE
	EnableVpcCoreDNS *bool `json:"EnableVpcCoreDNS,omitnil,omitempty" name:"EnableVpcCoreDNS"`

	// 标签描述列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSpecification []*TagSpecification `json:"TagSpecification,omitnil,omitempty" name:"TagSpecification"`
}

// Predefined struct for user
type EnableClusterAuditRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CLS日志集ID，可以通过cls接口或者控制台获取
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// CLS日志主题ID，可以通过cls接口或者控制台获取
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// topic所在region，默认为集群当前region
	TopicRegion *string `json:"TopicRegion,omitnil,omitempty" name:"TopicRegion"`
}

type EnableClusterAuditRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// CLS日志集ID，可以通过cls接口或者控制台获取
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// CLS日志主题ID，可以通过cls接口或者控制台获取
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// topic所在region，默认为集群当前region
	TopicRegion *string `json:"TopicRegion,omitnil,omitempty" name:"TopicRegion"`
}

func (r *EnableClusterAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClusterAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "TopicRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableClusterAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClusterAuditResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableClusterAuditResponse struct {
	*tchttp.BaseResponse
	Response *EnableClusterAuditResponseParams `json:"Response"`
}

func (r *EnableClusterAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClusterAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClusterDeletionProtectionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type EnableClusterDeletionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *EnableClusterDeletionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClusterDeletionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableClusterDeletionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableClusterDeletionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableClusterDeletionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *EnableClusterDeletionProtectionResponseParams `json:"Response"`
}

func (r *EnableClusterDeletionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableClusterDeletionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableControlPlaneLogsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型。当前只支持tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 组件名称列表，目前支持的组件有：cluster-autoscaler、kapenter
	Components []*ComponentLogConfig `json:"Components,omitnil,omitempty" name:"Components"`
}

type EnableControlPlaneLogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型。当前只支持tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 组件名称列表，目前支持的组件有：cluster-autoscaler、kapenter
	Components []*ComponentLogConfig `json:"Components,omitnil,omitempty" name:"Components"`
}

func (r *EnableControlPlaneLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableControlPlaneLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterType")
	delete(f, "Components")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableControlPlaneLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableControlPlaneLogsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableControlPlaneLogsResponse struct {
	*tchttp.BaseResponse
	Response *EnableControlPlaneLogsResponseParams `json:"Response"`
}

func (r *EnableControlPlaneLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableControlPlaneLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableEncryptionProtectionRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// kms加密配置
	KMSConfiguration *KMSConfiguration `json:"KMSConfiguration,omitnil,omitempty" name:"KMSConfiguration"`
}

type EnableEncryptionProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// kms加密配置
	KMSConfiguration *KMSConfiguration `json:"KMSConfiguration,omitnil,omitempty" name:"KMSConfiguration"`
}

func (r *EnableEncryptionProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableEncryptionProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "KMSConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableEncryptionProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableEncryptionProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableEncryptionProtectionResponse struct {
	*tchttp.BaseResponse
	Response *EnableEncryptionProtectionResponseParams `json:"Response"`
}

func (r *EnableEncryptionProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableEncryptionProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableEventPersistenceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// cls服务的logsetID，通过cls接口或者控制台获取
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// cls服务的topicID，通过cls接口或者控制台获取
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// topic所在地域，默认为集群所在地域
	TopicRegion *string `json:"TopicRegion,omitnil,omitempty" name:"TopicRegion"`
}

type EnableEventPersistenceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// cls服务的logsetID，通过cls接口或者控制台获取
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// cls服务的topicID，通过cls接口或者控制台获取
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// topic所在地域，默认为集群所在地域
	TopicRegion *string `json:"TopicRegion,omitnil,omitempty" name:"TopicRegion"`
}

func (r *EnableEventPersistenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableEventPersistenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "TopicRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableEventPersistenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableEventPersistenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableEventPersistenceResponse struct {
	*tchttp.BaseResponse
	Response *EnableEventPersistenceResponseParams `json:"Response"`
}

func (r *EnableEventPersistenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableEventPersistenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcCniNetworkTypeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 开启vpc-cni的模式，tke-route-eni开启的是策略路由模式，tke-direct-eni开启的是独立网卡模式
	VpcCniType *string `json:"VpcCniType,omitnil,omitempty" name:"VpcCniType"`

	// 是否开启固定IP模式
	EnableStaticIp *bool `json:"EnableStaticIp,omitnil,omitempty" name:"EnableStaticIp"`

	// 使用的容器子网
	Subnets []*string `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// 在固定IP模式下，Pod销毁后退还IP的时间，传参必须大于300；不传默认IP永不销毁。
	ExpiredSeconds *uint64 `json:"ExpiredSeconds,omitnil,omitempty" name:"ExpiredSeconds"`

	// 是否同步添加 vpc 网段到 ip-masq-agent-config 的 NonMasqueradeCIDRs 字段，默认 false 会同步添加
	SkipAddingNonMasqueradeCIDRs *bool `json:"SkipAddingNonMasqueradeCIDRs,omitnil,omitempty" name:"SkipAddingNonMasqueradeCIDRs"`
}

type EnableVpcCniNetworkTypeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 开启vpc-cni的模式，tke-route-eni开启的是策略路由模式，tke-direct-eni开启的是独立网卡模式
	VpcCniType *string `json:"VpcCniType,omitnil,omitempty" name:"VpcCniType"`

	// 是否开启固定IP模式
	EnableStaticIp *bool `json:"EnableStaticIp,omitnil,omitempty" name:"EnableStaticIp"`

	// 使用的容器子网
	Subnets []*string `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// 在固定IP模式下，Pod销毁后退还IP的时间，传参必须大于300；不传默认IP永不销毁。
	ExpiredSeconds *uint64 `json:"ExpiredSeconds,omitnil,omitempty" name:"ExpiredSeconds"`

	// 是否同步添加 vpc 网段到 ip-masq-agent-config 的 NonMasqueradeCIDRs 字段，默认 false 会同步添加
	SkipAddingNonMasqueradeCIDRs *bool `json:"SkipAddingNonMasqueradeCIDRs,omitnil,omitempty" name:"SkipAddingNonMasqueradeCIDRs"`
}

func (r *EnableVpcCniNetworkTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcCniNetworkTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VpcCniType")
	delete(f, "EnableStaticIp")
	delete(f, "Subnets")
	delete(f, "ExpiredSeconds")
	delete(f, "SkipAddingNonMasqueradeCIDRs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableVpcCniNetworkTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcCniNetworkTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableVpcCniNetworkTypeResponse struct {
	*tchttp.BaseResponse
	Response *EnableVpcCniNetworkTypeResponseParams `json:"Response"`
}

func (r *EnableVpcCniNetworkTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcCniNetworkTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`

	// 开启云自动化助手服务（TencentCloud Automation Tools，TAT）。若不指定该参数，则公共镜像默认开启云自动化助手服务，其他镜像默认不开启云自动化助手服务。
	AutomationService *RunAutomationServiceEnabled `json:"AutomationService,omitnil,omitempty" name:"AutomationService"`
}

type EnvironmentVariable struct {
	// key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// val
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EtcdOverrideConfig struct {
	// k8s资源，支持核心资源，控制类资源，配置及敏感资源
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type Event struct {
	// pod名称
	PodName *string `json:"PodName,omitnil,omitempty" name:"PodName"`

	// 事件原因内容
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 事件类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件出现次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 事件第一次出现时间
	FirstTimestamp *string `json:"FirstTimestamp,omitnil,omitempty" name:"FirstTimestamp"`

	// 事件最后一次出现时间
	LastTimestamp *string `json:"LastTimestamp,omitnil,omitempty" name:"LastTimestamp"`

	// 事件内容
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type Exec struct {
	// 容器内检测的命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	Commands []*string `json:"Commands,omitnil,omitempty" name:"Commands"`
}

type ExistedInstance struct {
	// 实例是否支持加入集群(TRUE 可以加入 FALSE 不能加入)。
	Usable *bool `json:"Usable,omitnil,omitempty" name:"Usable"`

	// 实例不支持加入的原因。
	UnusableReason *string `json:"UnusableReason,omitnil,omitempty" name:"UnusableReason"`

	// 实例已经所在的集群ID。
	AlreadyInCluster *string `json:"AlreadyInCluster,omitnil,omitempty" name:"AlreadyInCluster"`

	// 实例ID形如：ins-e55paxnt。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例主网卡的内网IP列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 实例主网卡的公网IP列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 实例的CPU核数，单位：核。
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 实例内存容量，单位：GB。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 操作系统名称。
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 伸缩组ID
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitnil,omitempty" name:"AutoscalingGroupId"`

	// 实例计费模式。取值范围： PREPAID：表示预付费，即包年包月 POSTPAID_BY_HOUR：表示后付费，即按量计费 CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的IPv6地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Addresses []*string `json:"IPv6Addresses,omitnil,omitempty" name:"IPv6Addresses"`
}

type ExistedInstancesForNode struct {
	// 节点角色，取值:MASTER_ETCD, WORKER。MASTER_ETCD只有在创建 INDEPENDENT_CLUSTER 独立集群时需要指定。MASTER_ETCD节点数量为3～7，建议为奇数。MASTER_ETCD最小配置为4C8G。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 已存在实例的重装参数
	ExistedInstancesPara *ExistedInstancesPara `json:"ExistedInstancesPara,omitnil,omitempty" name:"ExistedInstancesPara"`

	// 节点高级设置，会覆盖集群级别设置的InstanceAdvancedSettings（当前只对节点自定义参数ExtraArgs生效）
	InstanceAdvancedSettingsOverride *InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverride,omitnil,omitempty" name:"InstanceAdvancedSettingsOverride"`

	// 自定义模式集群，可指定每个节点的pod数量
	DesiredPodNumbers []*int64 `json:"DesiredPodNumbers,omitnil,omitempty" name:"DesiredPodNumbers"`
}

type ExistedInstancesPara struct {
	// 集群ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例额外需要设置参数信息（目前后端暂不支持此字段，我们将在未来的版本中实现）
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口HostName一致)
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`
}

type ExtensionAddon struct {
	// 扩展组件名称
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`

	// 扩展组件信息(扩展组件资源对象的json字符串描述)
	AddonParam *string `json:"AddonParam,omitnil,omitempty" name:"AddonParam"`
}

type FailedResource struct {
	// 资源六段式
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 执行失败的原因
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Flag struct {
	// 参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数描述
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 参数默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数可选范围（目前包含range和in两种，"[]"代表range，如"[1, 5]"表示参数必须>=1且 <=5, "()"代表in， 如"('aa', 'bb')"表示参数只能为字符串'aa'或者'bb'，该参数为空表示不校验）
	Constraint *string `json:"Constraint,omitnil,omitempty" name:"Constraint"`
}

// Predefined struct for user
type ForwardTKEEdgeApplicationRequestV3RequestParams struct {
	// 请求集群addon的访问
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 请求集群addon的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 请求集群addon后允许接收的数据格式
	Accept *string `json:"Accept,omitnil,omitempty" name:"Accept"`

	// 请求集群addon的数据格式
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 请求集群addon的数据
	RequestBody *string `json:"RequestBody,omitnil,omitempty" name:"RequestBody"`

	// 集群名称，例如cls-1234abcd
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 是否编码请求内容
	EncodedBody *string `json:"EncodedBody,omitnil,omitempty" name:"EncodedBody"`
}

type ForwardTKEEdgeApplicationRequestV3Request struct {
	*tchttp.BaseRequest
	
	// 请求集群addon的访问
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 请求集群addon的路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 请求集群addon后允许接收的数据格式
	Accept *string `json:"Accept,omitnil,omitempty" name:"Accept"`

	// 请求集群addon的数据格式
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// 请求集群addon的数据
	RequestBody *string `json:"RequestBody,omitnil,omitempty" name:"RequestBody"`

	// 集群名称，例如cls-1234abcd
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 是否编码请求内容
	EncodedBody *string `json:"EncodedBody,omitnil,omitempty" name:"EncodedBody"`
}

func (r *ForwardTKEEdgeApplicationRequestV3Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForwardTKEEdgeApplicationRequestV3Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Method")
	delete(f, "Path")
	delete(f, "Accept")
	delete(f, "ContentType")
	delete(f, "RequestBody")
	delete(f, "ClusterName")
	delete(f, "EncodedBody")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForwardTKEEdgeApplicationRequestV3Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForwardTKEEdgeApplicationRequestV3ResponseParams struct {
	// 请求集群addon后返回的数据
	ResponseBody *string `json:"ResponseBody,omitnil,omitempty" name:"ResponseBody"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ForwardTKEEdgeApplicationRequestV3Response struct {
	*tchttp.BaseResponse
	Response *ForwardTKEEdgeApplicationRequestV3ResponseParams `json:"Response"`
}

func (r *ForwardTKEEdgeApplicationRequestV3Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForwardTKEEdgeApplicationRequestV3Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GPUArgs struct {
	// CUDA版本信息
	CUDA *DriverVersion `json:"CUDA,omitnil,omitempty" name:"CUDA"`

	// cuDNN版本信息
	CUDNN *CUDNN `json:"CUDNN,omitnil,omitempty" name:"CUDNN"`

	// 自定义GPU驱动信息
	CustomDriver *CustomDriver `json:"CustomDriver,omitnil,omitempty" name:"CustomDriver"`

	// GPU驱动版本信息
	Driver *DriverVersion `json:"Driver,omitnil,omitempty" name:"Driver"`

	// 是否启用MIG特性
	MIGEnable *bool `json:"MIGEnable,omitnil,omitempty" name:"MIGEnable"`
}

// Predefined struct for user
type GetClusterLevelPriceRequestParams struct {
	// 集群规格，托管集群询价，集群等级：L20、L50、L100、L200、L500、L1000、L3000、L5000
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

type GetClusterLevelPriceRequest struct {
	*tchttp.BaseRequest
	
	// 集群规格，托管集群询价，集群等级：L20、L50、L100、L200、L500、L1000、L3000、L5000
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

func (r *GetClusterLevelPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterLevelPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetClusterLevelPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetClusterLevelPriceResponseParams struct {
	// 询价结果，单位：分，打折后
	Cost *uint64 `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 询价结果，单位：分，折扣前
	TotalCost *uint64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 总的折扣，100表示100%不打折
	Policy *float64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetClusterLevelPriceResponse struct {
	*tchttp.BaseResponse
	Response *GetClusterLevelPriceResponseParams `json:"Response"`
}

func (r *GetClusterLevelPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetClusterLevelPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMostSuitableImageCacheRequestParams struct {
	// 容器镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`
}

type GetMostSuitableImageCacheRequest struct {
	*tchttp.BaseRequest
	
	// 容器镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`
}

func (r *GetMostSuitableImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMostSuitableImageCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Images")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMostSuitableImageCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMostSuitableImageCacheResponseParams struct {
	// 是否有匹配的镜像缓存
	Found *bool `json:"Found,omitnil,omitempty" name:"Found"`

	// 匹配的镜像缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageCache *ImageCache `json:"ImageCache,omitnil,omitempty" name:"ImageCache"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMostSuitableImageCacheResponse struct {
	*tchttp.BaseResponse
	Response *GetMostSuitableImageCacheResponseParams `json:"Response"`
}

func (r *GetMostSuitableImageCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMostSuitableImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTkeAppChartListRequestParams struct {
	// app类型，取值log,scheduler,network,storage,monitor,dns,image,other,invisible
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// app支持的操作系统，取值arm32、arm64、amd64
	Arch *string `json:"Arch,omitnil,omitempty" name:"Arch"`

	// 集群类型，取值tke、eks
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type GetTkeAppChartListRequest struct {
	*tchttp.BaseRequest
	
	// app类型，取值log,scheduler,network,storage,monitor,dns,image,other,invisible
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// app支持的操作系统，取值arm32、arm64、amd64
	Arch *string `json:"Arch,omitnil,omitempty" name:"Arch"`

	// 集群类型，取值tke、eks
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *GetTkeAppChartListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTkeAppChartListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Kind")
	delete(f, "Arch")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTkeAppChartListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTkeAppChartListResponseParams struct {
	// 所支持的chart列表
	AppCharts []*AppChart `json:"AppCharts,omitnil,omitempty" name:"AppCharts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTkeAppChartListResponse struct {
	*tchttp.BaseResponse
	Response *GetTkeAppChartListResponseParams `json:"Response"`
}

func (r *GetTkeAppChartListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTkeAppChartListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeInstanceProgressRequestParams struct {
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 最多获取多少个节点的进度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 从第几个节点开始获取进度
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type GetUpgradeInstanceProgressRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 最多获取多少个节点的进度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 从第几个节点开始获取进度
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *GetUpgradeInstanceProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeInstanceProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUpgradeInstanceProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUpgradeInstanceProgressResponseParams struct {
	// 升级节点总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 已升级节点总数
	Done *int64 `json:"Done,omitnil,omitempty" name:"Done"`

	// 升级任务生命周期
	// 
	// pending 还未开始
	// process 运行中
	// paused 已停止
	// pauing 正在停止
	// done  已完成
	// timeout 已超时
	// aborted 已取消
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 各节点升级进度详情
	Instances []*InstanceUpgradeProgressItem `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 集群当前状态
	ClusterStatus *InstanceUpgradeClusterStatus `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUpgradeInstanceProgressResponse struct {
	*tchttp.BaseResponse
	Response *GetUpgradeInstanceProgressResponseParams `json:"Response"`
}

func (r *GetUpgradeInstanceProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUpgradeInstanceProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlobalMaintenanceWindowAndExclusion struct {
	// 地域
	TargetRegions []*string `json:"TargetRegions,omitnil,omitempty" name:"TargetRegions"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`

	// 维护窗口ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type HttpGet struct {
	// HttpGet检测的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// HttpGet检测的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// HTTP or HTTPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`
}

type IPAddress struct {
	// Ip 地址的类型。可为 advertise, public 等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Ip 地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 网络端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type ImageCache struct {
	// 镜像缓存Id
	ImageCacheId *string `json:"ImageCacheId,omitnil,omitempty" name:"ImageCacheId"`

	// 镜像缓存名称
	ImageCacheName *string `json:"ImageCacheName,omitnil,omitempty" name:"ImageCacheName"`

	// 镜像缓存大小。单位：GiB
	ImageCacheSize *uint64 `json:"ImageCacheSize,omitnil,omitempty" name:"ImageCacheSize"`

	// 镜像缓存包含的镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 到期时间
	ExpireDateTime *string `json:"ExpireDateTime,omitnil,omitempty" name:"ExpireDateTime"`

	// 镜像缓存事件信息
	Events []*ImageCacheEvent `json:"Events,omitnil,omitempty" name:"Events"`

	// 最新一次匹配到镜像缓存的时间
	LastMatchedTime *string `json:"LastMatchedTime,omitnil,omitempty" name:"LastMatchedTime"`

	// 镜像缓存对应的快照Id
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 镜像缓存状态，可能取值：
	// Pending：创建中
	// Ready：创建完成
	// Failed：创建失败
	// Updating：更新中
	// UpdateFailed：更新失败
	// 只有状态为Ready时，才能正常使用镜像缓存
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 镜像缓存保留时间天数，过期将会自动清理，默认为0，永不过期。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 镜像拉取凭证
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ImageCacheEvent struct {
	// 镜像缓存Id
	ImageCacheId *string `json:"ImageCacheId,omitnil,omitempty" name:"ImageCacheId"`

	// 事件类型, Normal或者Warning
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件原因简述
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 事件原因详述
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 事件第一次出现时间
	FirstTimestamp *string `json:"FirstTimestamp,omitnil,omitempty" name:"FirstTimestamp"`

	// 事件最后一次出现时间
	LastTimestamp *string `json:"LastTimestamp,omitnil,omitempty" name:"LastTimestamp"`
}

type ImageInstance struct {
	// 镜像别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 操作系统名称
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`
}

type ImageRegistryCredential struct {
	// 镜像仓库地址
	Server *string `json:"Server,omitnil,omitempty" name:"Server"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// ImageRegistryCredential的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type InstallAddonRequestParams struct {
	// 集群ID（仅支持标准tke集群），请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`

	// addon版本（不传默认安装最新版本）
	AddonVersion *string `json:"AddonVersion,omitnil,omitempty" name:"AddonVersion"`

	// addon的参数，是一个json格式的base64转码后的字符串（addon参数由DescribeAddonValues获取）
	RawValues *string `json:"RawValues,omitnil,omitempty" name:"RawValues"`

	// 是否仅做安装检查，设置为true时仅做检查，不会安装组件。默认值为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type InstallAddonRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（仅支持标准tke集群），请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`

	// addon版本（不传默认安装最新版本）
	AddonVersion *string `json:"AddonVersion,omitnil,omitempty" name:"AddonVersion"`

	// addon的参数，是一个json格式的base64转码后的字符串（addon参数由DescribeAddonValues获取）
	RawValues *string `json:"RawValues,omitnil,omitempty" name:"RawValues"`

	// 是否仅做安装检查，设置为true时仅做检查，不会安装组件。默认值为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *InstallAddonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAddonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AddonName")
	delete(f, "AddonVersion")
	delete(f, "RawValues")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallAddonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallAddonResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallAddonResponse struct {
	*tchttp.BaseResponse
	Response *InstallAddonResponseParams `json:"Response"`
}

func (r *InstallAddonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallAddonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallEdgeLogAgentRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type InstallEdgeLogAgentRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *InstallEdgeLogAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallEdgeLogAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallEdgeLogAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallEdgeLogAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallEdgeLogAgentResponse struct {
	*tchttp.BaseResponse
	Response *InstallEdgeLogAgentResponseParams `json:"Response"`
}

func (r *InstallEdgeLogAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallEdgeLogAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallLogAgentRequestParams struct {
	// TKE集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// kubelet根目录
	KubeletRootDir *string `json:"KubeletRootDir,omitnil,omitempty" name:"KubeletRootDir"`

	// 集群类型 tke/eks，默认tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type InstallLogAgentRequest struct {
	*tchttp.BaseRequest
	
	// TKE集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// kubelet根目录
	KubeletRootDir *string `json:"KubeletRootDir,omitnil,omitempty" name:"KubeletRootDir"`

	// 集群类型 tke/eks，默认tke
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *InstallLogAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallLogAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "KubeletRootDir")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallLogAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallLogAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallLogAgentResponse struct {
	*tchttp.BaseResponse
	Response *InstallLogAgentResponseParams `json:"Response"`
}

func (r *InstallLogAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallLogAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例异常(或者处于初始化中)的原因
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`

	// 实例的状态（running 运行中，initializing 初始化中，failed 异常）
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 实例是否封锁状态
	DrainStatus *string `json:"DrainStatus,omitnil,omitempty" name:"DrainStatus"`

	// 节点配置
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 添加时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 节点内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIP *string `json:"LanIP,omitnil,omitempty" name:"LanIP"`

	// 资源池ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 自动伸缩组ID
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitnil,omitempty" name:"AutoscalingGroupId"`
}

type InstanceAdvancedSettings struct {
	// 该节点属于podCIDR大小自定义模式时，可指定节点上运行的pod数量上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredPodNumber *int64 `json:"DesiredPodNumber,omitnil,omitempty" name:"DesiredPodNumber"`

	// GPU驱动相关参数,相关的GPU参数获取:https://cloud.tencent.com/document/api/213/15715
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUArgs *GPUArgs `json:"GPUArgs,omitnil,omitempty" name:"GPUArgs"`

	// base64 编码的用户脚本，在初始化节点之前执行，目前只对添加已有节点生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreStartUserScript *string `json:"PreStartUserScript,omitnil,omitempty" name:"PreStartUserScript"`

	// 节点污点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 数据盘挂载点, 默认不挂载数据盘. 已格式化的 ext3，ext4，xfs 文件系统的数据盘将直接挂载，其他文件系统或未格式化的数据盘将自动格式化为ext4 (tlinux系统格式化成xfs)并挂载，请注意备份数据! 无数据盘或有多块数据盘的云主机此设置不生效。
	// 注意：多盘场景请使用下方的DataDisks数据结构，设置对应的云盘类型、云盘大小、挂载路径、是否格式化等信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountTarget *string `json:"MountTarget,omitnil,omitempty" name:"MountTarget"`

	// dockerd --graph 指定值。若未指定此参数，将使用内置默认路径 /var/lib/docker 作为存储根目录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DockerGraphPath *string `json:"DockerGraphPath,omitnil,omitempty" name:"DockerGraphPath"`

	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行后执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看, 如果要求节点需要在进行初始化完成后才可加入调度, 可配合 unschedulable 参数使用, 在 userScript 最后初始化完成后, 添加 kubectl uncordon nodename --kubeconfig=/root/.kube/config 命令使节点加入调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.
	Unschedulable *int64 `json:"Unschedulable,omitnil,omitempty" name:"Unschedulable"`

	// 节点Label数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 多盘数据盘挂载信息：新建节点时请确保购买CVM的参数传递了购买多个数据盘的信息，如CreateClusterInstances API的RunInstancesPara下的DataDisks也需要设置购买多个数据盘, 具体可以参考CreateClusterInstances接口的添加集群节点(多块数据盘)样例；添加已有节点时，请确保填写的分区信息在节点上真实存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 节点相关的自定义参数信息
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceDataDiskMountSetting struct {
	// CVM实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 数据盘挂载信息
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// CVM实例所属可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type InstanceExtraArgs struct {
	// kubelet自定义参数，参数格式为["k1=v1", "k1=v2"]， 例如["root-dir=/var/lib/kubelet","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kubelet []*string `json:"Kubelet,omitnil,omitempty" name:"Kubelet"`
}

type InstanceUpgradeClusterStatus struct {
	// pod总数
	PodTotal *int64 `json:"PodTotal,omitnil,omitempty" name:"PodTotal"`

	// NotReady pod总数
	NotReadyPod *int64 `json:"NotReadyPod,omitnil,omitempty" name:"NotReadyPod"`
}

type InstanceUpgradePreCheckResult struct {
	// 检查是否通过
	CheckPass *bool `json:"CheckPass,omitnil,omitempty" name:"CheckPass"`

	// 检查项数组
	Items []*InstanceUpgradePreCheckResultItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 本节点独立pod列表
	SinglePods []*string `json:"SinglePods,omitnil,omitempty" name:"SinglePods"`
}

type InstanceUpgradePreCheckResultItem struct {
	// 工作负载的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 工作负载类型
	WorkLoadKind *string `json:"WorkLoadKind,omitnil,omitempty" name:"WorkLoadKind"`

	// 工作负载名称
	WorkLoadName *string `json:"WorkLoadName,omitnil,omitempty" name:"WorkLoadName"`

	// 驱逐节点前工作负载running的pod数目
	Before *uint64 `json:"Before,omitnil,omitempty" name:"Before"`

	// 驱逐节点后工作负载running的pod数目
	After *uint64 `json:"After,omitnil,omitempty" name:"After"`

	// 工作负载在本节点上的pod列表
	Pods []*string `json:"Pods,omitnil,omitempty" name:"Pods"`
}

type InstanceUpgradeProgressItem struct {
	// 节点instanceID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 任务生命周期
	// process 运行中
	// paused 已停止
	// pauing 正在停止
	// done  已完成
	// timeout 已超时
	// aborted 已取消
	// pending 还未开始
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 升级开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 升级结束时间
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 升级前检查结果
	CheckResult *InstanceUpgradePreCheckResult `json:"CheckResult,omitnil,omitempty" name:"CheckResult"`

	// 升级步骤详情
	Detail []*TaskStepInfo `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type KMSConfiguration struct {
	// kms id，可以在密钥管理控制台获取
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// kms 地域
	KmsRegion *string `json:"KmsRegion,omitnil,omitempty" name:"KmsRegion"`
}

type KubeJarvisStateCatalogue struct {
	// 目录级别，支持参数：
	// first：一级目录
	// second：二级目录
	CatalogueLevel *string `json:"CatalogueLevel,omitnil,omitempty" name:"CatalogueLevel"`

	// 目录名
	CatalogueName *string `json:"CatalogueName,omitnil,omitempty" name:"CatalogueName"`
}

type KubeJarvisStateDiagnostic struct {
	// 诊断开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 诊断结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 诊断目录
	Catalogues []*KubeJarvisStateCatalogue `json:"Catalogues,omitnil,omitempty" name:"Catalogues"`

	// 诊断类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 诊断名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 诊断描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 诊断结果列表
	Results []*KubeJarvisStateResultsItem `json:"Results,omitnil,omitempty" name:"Results"`

	// 诊断结果统计
	Statistics []*KubeJarvisStateStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`
}

type KubeJarvisStateDiagnosticOverview struct {
	// 诊断目录
	Catalogues []*KubeJarvisStateCatalogue `json:"Catalogues,omitnil,omitempty" name:"Catalogues"`

	// 诊断结果统计
	Statistics []*KubeJarvisStateStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`
}

type KubeJarvisStateInspectionOverview struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 诊断结果统计
	Statistics []*KubeJarvisStateStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 诊断结果详情
	Diagnostics []*KubeJarvisStateDiagnosticOverview `json:"Diagnostics,omitnil,omitempty" name:"Diagnostics"`
}

type KubeJarvisStateInspectionResult struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 诊断开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 诊断结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 诊断结果统计
	Statistics []*KubeJarvisStateStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 诊断结果详情
	Diagnostics []*KubeJarvisStateDiagnostic `json:"Diagnostics,omitnil,omitempty" name:"Diagnostics"`

	// 查询巡检报告相关报错
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type KubeJarvisStateInspectionResultsItem struct {
	// 巡检结果名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 诊断结果统计
	Statistics []*KubeJarvisStateStatistic `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 诊断数据开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 诊断数据结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type KubeJarvisStateResultObjInfo struct {
	// 对象属性名称
	PropertyName *string `json:"PropertyName,omitnil,omitempty" name:"PropertyName"`

	// 对象属性值
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`
}

type KubeJarvisStateResultsItem struct {
	// 诊断结果级别
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 诊断对象名
	ObjName *string `json:"ObjName,omitnil,omitempty" name:"ObjName"`

	// 诊断对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjInfo []*KubeJarvisStateResultObjInfo `json:"ObjInfo,omitnil,omitempty" name:"ObjInfo"`

	// 诊断项标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 诊断项描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 诊断建议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proposal *string `json:"Proposal,omitnil,omitempty" name:"Proposal"`

	// 诊断建议文档链接
	ProposalDocUrl *string `json:"ProposalDocUrl,omitnil,omitempty" name:"ProposalDocUrl"`

	// 诊断建议文档名称
	ProposalDocName *string `json:"ProposalDocName,omitnil,omitempty" name:"ProposalDocName"`
}

type KubeJarvisStateStatistic struct {
	// 诊断结果的健康水平，健康水平取值：
	// serious：高风险
	// risk：中风险
	// warn：低风险
	// good：健康
	// failed：诊断流程异常
	HealthyLevel *string `json:"HealthyLevel,omitnil,omitempty" name:"HealthyLevel"`

	// 诊断结果的统计
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type Label struct {
	// map表中的Name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// map表中的Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ListClusterInspectionResultsItemsRequestParams struct {
	// 目标集群ID
	// 
	// 取值可参考：
	// [查询TKE集群列表](https://cloud.tencent.com/document/api/457/31862)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 查询历史结果的开始时间，Unix时间戳
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询历史结果的结束时间，默认当前距离开始时间3天，Unix时间戳
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ListClusterInspectionResultsItemsRequest struct {
	*tchttp.BaseRequest
	
	// 目标集群ID
	// 
	// 取值可参考：
	// [查询TKE集群列表](https://cloud.tencent.com/document/api/457/31862)
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 查询历史结果的开始时间，Unix时间戳
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询历史结果的结束时间，默认当前距离开始时间3天，Unix时间戳
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ListClusterInspectionResultsItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClusterInspectionResultsItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListClusterInspectionResultsItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClusterInspectionResultsItemsResponseParams struct {
	// 巡检结果历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectionResultsItems []*KubeJarvisStateInspectionResultsItem `json:"InspectionResultsItems,omitnil,omitempty" name:"InspectionResultsItems"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListClusterInspectionResultsItemsResponse struct {
	*tchttp.BaseResponse
	Response *ListClusterInspectionResultsItemsResponseParams `json:"Response"`
}

func (r *ListClusterInspectionResultsItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClusterInspectionResultsItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClusterInspectionResultsRequestParams struct {
	// 目标集群列表，为空查询用户所有集群
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 隐藏的字段信息，为了减少无效的字段返回，隐藏字段不会在返回值中返回。可选值：results
	Hide []*string `json:"Hide,omitnil,omitempty" name:"Hide"`

	// 指定查询结果的报告名称，默认查询最新的每个集群只查询最新的一条
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ListClusterInspectionResultsRequest struct {
	*tchttp.BaseRequest
	
	// 目标集群列表，为空查询用户所有集群
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 隐藏的字段信息，为了减少无效的字段返回，隐藏字段不会在返回值中返回。可选值：results
	Hide []*string `json:"Hide,omitnil,omitempty" name:"Hide"`

	// 指定查询结果的报告名称，默认查询最新的每个集群只查询最新的一条
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *ListClusterInspectionResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClusterInspectionResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Hide")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListClusterInspectionResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListClusterInspectionResultsResponseParams struct {
	// 集群诊断结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InspectionResults []*KubeJarvisStateInspectionResult `json:"InspectionResults,omitnil,omitempty" name:"InspectionResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListClusterInspectionResultsResponse struct {
	*tchttp.BaseResponse
	Response *ListClusterInspectionResultsResponseParams `json:"Response"`
}

func (r *ListClusterInspectionResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListClusterInspectionResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LivenessOrReadinessProbe struct {
	// 探针参数
	Probe *Probe `json:"Probe,omitnil,omitempty" name:"Probe"`

	// 容器内检测命令参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exec *Exec `json:"Exec,omitnil,omitempty" name:"Exec"`

	// HttpGet检测参数
	HttpGet *HttpGet `json:"HttpGet,omitnil,omitempty" name:"HttpGet"`

	// TcpSocket检测的端口参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TcpSocket *TcpSocket `json:"TcpSocket,omitnil,omitempty" name:"TcpSocket"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<li>Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。</li><li>Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]中的特殊符号。</li>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口[DescribeKeyPairs](https://cloud.tencent.com/document/api/213/15699)获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为true。取值范围：<li>true：表示保持镜像的登录设置</li><li>false：表示不保持镜像的登录设置</li>默认取值：false。
	KeepImageLogin *string `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

type MaintenanceExclusion struct {
	// 维护排除项名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 维护排除项开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 维护排除项结束时间
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 维护排除项ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

type MaintenanceWindowAndExclusion struct {
	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群版本
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

type ManuallyAdded struct {
	// 加入中的节点数量
	Joining *int64 `json:"Joining,omitnil,omitempty" name:"Joining"`

	// 初始化中的节点数量
	Initializing *int64 `json:"Initializing,omitnil,omitempty" name:"Initializing"`

	// 正常的节点数量
	Normal *int64 `json:"Normal,omitnil,omitempty" name:"Normal"`

	// 节点总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type ModifyClusterAsGroupAttributeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群关联的伸缩组属性
	ClusterAsGroupAttribute *ClusterAsGroupAttribute `json:"ClusterAsGroupAttribute,omitnil,omitempty" name:"ClusterAsGroupAttribute"`
}

type ModifyClusterAsGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群关联的伸缩组属性
	ClusterAsGroupAttribute *ClusterAsGroupAttribute `json:"ClusterAsGroupAttribute,omitnil,omitempty" name:"ClusterAsGroupAttribute"`
}

func (r *ModifyClusterAsGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterAsGroupAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAsGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAsGroupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterAsGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAsGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyClusterAsGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAsGroupOptionAttributeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群弹性伸缩属性
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitnil,omitempty" name:"ClusterAsGroupOption"`
}

type ModifyClusterAsGroupOptionAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群弹性伸缩属性
	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitnil,omitempty" name:"ClusterAsGroupOption"`
}

func (r *ModifyClusterAsGroupOptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupOptionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterAsGroupOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAsGroupOptionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAsGroupOptionAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterAsGroupOptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAsGroupOptionAttributeResponseParams `json:"Response"`
}

func (r *ModifyClusterAsGroupOptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAsGroupOptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAttributeRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群所属项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群名称,字符长度50
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群等级，等级类型：L20、L50、L100、L200、L500、L1000、L3000、L5000
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 自动变配集群等级
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 是否开启QGPU共享
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitnil,omitempty" name:"QGPUShareEnable"`

	// 集群属性
	ClusterProperty *ClusterProperty `json:"ClusterProperty,omitnil,omitempty" name:"ClusterProperty"`
}

type ModifyClusterAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群所属项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群名称,字符长度50
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群等级，等级类型：L20、L50、L100、L200、L500、L1000、L3000、L5000
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 自动变配集群等级
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 是否开启QGPU共享
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitnil,omitempty" name:"QGPUShareEnable"`

	// 集群属性
	ClusterProperty *ClusterProperty `json:"ClusterProperty,omitnil,omitempty" name:"ClusterProperty"`
}

func (r *ModifyClusterAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProjectId")
	delete(f, "ClusterName")
	delete(f, "ClusterDesc")
	delete(f, "ClusterLevel")
	delete(f, "AutoUpgradeClusterLevel")
	delete(f, "QGPUShareEnable")
	delete(f, "ClusterProperty")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAttributeResponseParams struct {
	// 集群所属项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群描述
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 集群等级
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`

	// 自动变配集群等级
	AutoUpgradeClusterLevel *AutoUpgradeClusterLevel `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 是否开启QGPU共享
	QGPUShareEnable *bool `json:"QGPUShareEnable,omitnil,omitempty" name:"QGPUShareEnable"`

	// 集群属性
	ClusterProperty *ClusterProperty `json:"ClusterProperty,omitnil,omitempty" name:"ClusterProperty"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAttributeResponseParams `json:"Response"`
}

func (r *ModifyClusterAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAuthenticationOptionsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ServiceAccount认证配置
	ServiceAccounts *ServiceAccountAuthenticationOptions `json:"ServiceAccounts,omitnil,omitempty" name:"ServiceAccounts"`

	// OIDC认证配置
	OIDCConfig *OIDCConfigAuthenticationOptions `json:"OIDCConfig,omitnil,omitempty" name:"OIDCConfig"`
}

type ModifyClusterAuthenticationOptionsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ServiceAccount认证配置
	ServiceAccounts *ServiceAccountAuthenticationOptions `json:"ServiceAccounts,omitnil,omitempty" name:"ServiceAccounts"`

	// OIDC认证配置
	OIDCConfig *OIDCConfigAuthenticationOptions `json:"OIDCConfig,omitnil,omitempty" name:"OIDCConfig"`
}

func (r *ModifyClusterAuthenticationOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAuthenticationOptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ServiceAccounts")
	delete(f, "OIDCConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterAuthenticationOptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterAuthenticationOptionsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterAuthenticationOptionsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterAuthenticationOptionsResponseParams `json:"Response"`
}

func (r *ModifyClusterAuthenticationOptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterAuthenticationOptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterEndpointSPRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`

	// 修改外网访问安全组
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`
}

type ModifyClusterEndpointSPRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)
	SecurityPolicies []*string `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`

	// 修改外网访问安全组
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`
}

func (r *ModifyClusterEndpointSPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterEndpointSPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SecurityPolicies")
	delete(f, "SecurityGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterEndpointSPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterEndpointSPResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterEndpointSPResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterEndpointSPResponseParams `json:"Response"`
}

func (r *ModifyClusterEndpointSPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterEndpointSPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterExtraArgsRequestParams struct {
	// 目标集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群自定义参数
	ClusterExtraArgs *ClusterExtraArgs `json:"ClusterExtraArgs,omitnil,omitempty" name:"ClusterExtraArgs"`
}

type ModifyClusterExtraArgsRequest struct {
	*tchttp.BaseRequest
	
	// 目标集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群自定义参数
	ClusterExtraArgs *ClusterExtraArgs `json:"ClusterExtraArgs,omitnil,omitempty" name:"ClusterExtraArgs"`
}

func (r *ModifyClusterExtraArgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterExtraArgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterExtraArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterExtraArgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterExtraArgsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterExtraArgsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterExtraArgsResponseParams `json:"Response"`
}

func (r *ModifyClusterExtraArgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterExtraArgsTaskStateRequestParams struct {
	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 操作类型：
	// abort 取消并回退任务
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

type ModifyClusterExtraArgsTaskStateRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 操作类型：
	// abort 取消并回退任务
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`
}

func (r *ModifyClusterExtraArgsTaskStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterExtraArgsTaskStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Operation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterExtraArgsTaskStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterExtraArgsTaskStateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterExtraArgsTaskStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterExtraArgsTaskStateResponseParams `json:"Response"`
}

func (r *ModifyClusterExtraArgsTaskStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterExtraArgsTaskStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterImageRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 指定有效的镜像ID，格式形如img-e55paxnt。可通过登录控制台查询，也可调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715)，取返回信息中的ImageId字段。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type ModifyClusterImageRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 指定有效的镜像ID，格式形如img-e55paxnt。可通过登录控制台查询，也可调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715)，取返回信息中的ImageId字段。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

func (r *ModifyClusterImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ImageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterImageResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterImageResponseParams `json:"Response"`
}

func (r *ModifyClusterImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterMaintenanceWindowAndExclusionsRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

type ModifyClusterMaintenanceWindowAndExclusionsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

func (r *ModifyClusterMaintenanceWindowAndExclusionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterMaintenanceWindowAndExclusionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "MaintenanceTime")
	delete(f, "Duration")
	delete(f, "DayOfWeek")
	delete(f, "Exclusions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterMaintenanceWindowAndExclusionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterMaintenanceWindowAndExclusionsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterMaintenanceWindowAndExclusionsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterMaintenanceWindowAndExclusionsResponseParams `json:"Response"`
}

func (r *ModifyClusterMaintenanceWindowAndExclusionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterMaintenanceWindowAndExclusionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNodePoolRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 名称，最长63个字符，只能包含小写字母、数字及分隔符“_”，且必须以小写字母开头，数字或小写字母结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 最大节点数
	MaxNodesNum *int64 `json:"MaxNodesNum,omitnil,omitempty" name:"MaxNodesNum"`

	// 最小节点数
	MinNodesNum *int64 `json:"MinNodesNum,omitnil,omitempty" name:"MinNodesNum"`

	// 标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 污点
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点 Annotation 列表
	Annotations []*AnnotationValue `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 是否开启伸缩
	EnableAutoscale *bool `json:"EnableAutoscale,omitnil,omitempty" name:"EnableAutoscale"`

	// 操作系统名称
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// GPU驱动版本，CUDA版本，cuDNN版本以及是否启用MIG特性
	GPUArgs *GPUArgs `json:"GPUArgs,omitnil,omitempty" name:"GPUArgs"`

	// base64编码后的自定义脚本
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 更新label和taint时忽略存量节点
	IgnoreExistedNode *bool `json:"IgnoreExistedNode,omitnil,omitempty" name:"IgnoreExistedNode"`

	// 节点自定义参数
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.
	Unschedulable *int64 `json:"Unschedulable,omitnil,omitempty" name:"Unschedulable"`

	// 删除保护开关
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// dockerd --graph 指定值, 默认为 /var/lib/docker
	DockerGraphPath *string `json:"DockerGraphPath,omitnil,omitempty" name:"DockerGraphPath"`

	// base64编码后的自定义脚本
	PreStartUserScript *string `json:"PreStartUserScript,omitnil,omitempty" name:"PreStartUserScript"`
}

type ModifyClusterNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 名称，最长63个字符，只能包含小写字母、数字及分隔符“_”，且必须以小写字母开头，数字或小写字母结尾
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 最大节点数
	MaxNodesNum *int64 `json:"MaxNodesNum,omitnil,omitempty" name:"MaxNodesNum"`

	// 最小节点数
	MinNodesNum *int64 `json:"MinNodesNum,omitnil,omitempty" name:"MinNodesNum"`

	// 标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 污点
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点 Annotation 列表
	Annotations []*AnnotationValue `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 是否开启伸缩
	EnableAutoscale *bool `json:"EnableAutoscale,omitnil,omitempty" name:"EnableAutoscale"`

	// 操作系统名称
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// GPU驱动版本，CUDA版本，cuDNN版本以及是否启用MIG特性
	GPUArgs *GPUArgs `json:"GPUArgs,omitnil,omitempty" name:"GPUArgs"`

	// base64编码后的自定义脚本
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 更新label和taint时忽略存量节点
	IgnoreExistedNode *bool `json:"IgnoreExistedNode,omitnil,omitempty" name:"IgnoreExistedNode"`

	// 节点自定义参数
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.
	Unschedulable *int64 `json:"Unschedulable,omitnil,omitempty" name:"Unschedulable"`

	// 删除保护开关
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// dockerd --graph 指定值, 默认为 /var/lib/docker
	DockerGraphPath *string `json:"DockerGraphPath,omitnil,omitempty" name:"DockerGraphPath"`

	// base64编码后的自定义脚本
	PreStartUserScript *string `json:"PreStartUserScript,omitnil,omitempty" name:"PreStartUserScript"`
}

func (r *ModifyClusterNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "Name")
	delete(f, "MaxNodesNum")
	delete(f, "MinNodesNum")
	delete(f, "Labels")
	delete(f, "Taints")
	delete(f, "Annotations")
	delete(f, "EnableAutoscale")
	delete(f, "OsName")
	delete(f, "OsCustomizeType")
	delete(f, "GPUArgs")
	delete(f, "UserScript")
	delete(f, "IgnoreExistedNode")
	delete(f, "ExtraArgs")
	delete(f, "Tags")
	delete(f, "Unschedulable")
	delete(f, "DeletionProtection")
	delete(f, "DockerGraphPath")
	delete(f, "PreStartUserScript")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNodePoolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterNodePoolResponseParams `json:"Response"`
}

func (r *ModifyClusterNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterRollOutSequenceTagsRequestParams struct {
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 集群发布序列标签（为空时表示移除集群标签）
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyClusterRollOutSequenceTagsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 集群发布序列标签（为空时表示移除集群标签）
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyClusterRollOutSequenceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRollOutSequenceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterID")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterRollOutSequenceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterRollOutSequenceTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterRollOutSequenceTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterRollOutSequenceTagsResponseParams `json:"Response"`
}

func (r *ModifyClusterRollOutSequenceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRollOutSequenceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterRuntimeConfigRequestParams struct {
	// 集群ID，必填
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 运行时版本需依据指定的Kubernetes版本进行设置。典型情况为，在升级过程中因版本冲突而需要调整运行时版本时。
	DstK8SVersion *string `json:"DstK8SVersion,omitnil,omitempty" name:"DstK8SVersion"`

	// 需要修改集群运行时填写
	ClusterRuntimeConfig *RuntimeConfig `json:"ClusterRuntimeConfig,omitnil,omitempty" name:"ClusterRuntimeConfig"`

	// 需要修改节点池运行时，填需要修改的部分
	NodePoolRuntimeConfig []*NodePoolRuntime `json:"NodePoolRuntimeConfig,omitnil,omitempty" name:"NodePoolRuntimeConfig"`
}

type ModifyClusterRuntimeConfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，必填
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 运行时版本需依据指定的Kubernetes版本进行设置。典型情况为，在升级过程中因版本冲突而需要调整运行时版本时。
	DstK8SVersion *string `json:"DstK8SVersion,omitnil,omitempty" name:"DstK8SVersion"`

	// 需要修改集群运行时填写
	ClusterRuntimeConfig *RuntimeConfig `json:"ClusterRuntimeConfig,omitnil,omitempty" name:"ClusterRuntimeConfig"`

	// 需要修改节点池运行时，填需要修改的部分
	NodePoolRuntimeConfig []*NodePoolRuntime `json:"NodePoolRuntimeConfig,omitnil,omitempty" name:"NodePoolRuntimeConfig"`
}

func (r *ModifyClusterRuntimeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRuntimeConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DstK8SVersion")
	delete(f, "ClusterRuntimeConfig")
	delete(f, "NodePoolRuntimeConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterRuntimeConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterRuntimeConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterRuntimeConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterRuntimeConfigResponseParams `json:"Response"`
}

func (r *ModifyClusterRuntimeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRuntimeConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterTagsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群标签:
	// [{"TagKey":"env","TagValue":"dev"}]}]
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否同步集群内子资源标签
	SyncSubresource *bool `json:"SyncSubresource,omitnil,omitempty" name:"SyncSubresource"`
}

type ModifyClusterTagsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群标签:
	// [{"TagKey":"env","TagValue":"dev"}]}]
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否同步集群内子资源标签
	SyncSubresource *bool `json:"SyncSubresource,omitnil,omitempty" name:"SyncSubresource"`
}

func (r *ModifyClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Tags")
	delete(f, "SyncSubresource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterTagsResponseParams struct {
	// 集群标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterTagsResponseParams `json:"Response"`
}

func (r *ModifyClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterVirtualNodePoolRequestParams struct {
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID，通过DescribeNodePools接口获取
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点池名称，必须修改至少一个参数
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安全组ID列表，必须修改至少一个参数
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 虚拟节点label，必须修改至少一个参数
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 虚拟节点taint，必须修改至少一个参数
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 删除保护开关，必须修改至少一个参数
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

type ModifyClusterVirtualNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，通过DescribeClusters接口获取
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池ID，通过DescribeNodePools接口获取
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点池名称，必须修改至少一个参数
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安全组ID列表，必须修改至少一个参数
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 虚拟节点label，必须修改至少一个参数
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 虚拟节点taint，必须修改至少一个参数
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 删除保护开关，必须修改至少一个参数
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`
}

func (r *ModifyClusterVirtualNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterVirtualNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "Name")
	delete(f, "SecurityGroupIds")
	delete(f, "Labels")
	delete(f, "Taints")
	delete(f, "DeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterVirtualNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterVirtualNodePoolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterVirtualNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterVirtualNodePoolResponseParams `json:"Response"`
}

func (r *ModifyClusterVirtualNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterVirtualNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalMaintenanceWindowAndExclusionsRequestParams struct {
	// 维护窗口ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 地域
	TargetRegions []*string `json:"TargetRegions,omitnil,omitempty" name:"TargetRegions"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

type ModifyGlobalMaintenanceWindowAndExclusionsRequest struct {
	*tchttp.BaseRequest
	
	// 维护窗口ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 地域
	TargetRegions []*string `json:"TargetRegions,omitnil,omitempty" name:"TargetRegions"`

	// 维护开始时间
	MaintenanceTime *string `json:"MaintenanceTime,omitnil,omitempty" name:"MaintenanceTime"`

	// 维护时长（小时）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 维护周期（一周中的哪几天）
	DayOfWeek []*string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 维护排除项
	Exclusions []*MaintenanceExclusion `json:"Exclusions,omitnil,omitempty" name:"Exclusions"`
}

func (r *ModifyGlobalMaintenanceWindowAndExclusionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalMaintenanceWindowAndExclusionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "TargetRegions")
	delete(f, "MaintenanceTime")
	delete(f, "Duration")
	delete(f, "DayOfWeek")
	delete(f, "Exclusions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlobalMaintenanceWindowAndExclusionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlobalMaintenanceWindowAndExclusionsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlobalMaintenanceWindowAndExclusionsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlobalMaintenanceWindowAndExclusionsResponseParams `json:"Response"`
}

func (r *ModifyGlobalMaintenanceWindowAndExclusionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlobalMaintenanceWindowAndExclusionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMasterComponentRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// master组件名称，支持kube-apiserver、kube-scheduler、kube-controller-manager
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 停机或恢复，值只能为：shutdown或restore
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 为true时，不真正执行停机或恢复操作	
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type ModifyMasterComponentRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// master组件名称，支持kube-apiserver、kube-scheduler、kube-controller-manager
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 停机或恢复，值只能为：shutdown或restore
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 为true时，不真正执行停机或恢复操作	
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *ModifyMasterComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMasterComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Component")
	delete(f, "Operation")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMasterComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMasterComponentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMasterComponentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMasterComponentResponseParams `json:"Response"`
}

func (r *ModifyMasterComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMasterComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodePoolDesiredCapacityAboutAsgRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点池所关联的伸缩组的期望实例数
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`
}

type ModifyNodePoolDesiredCapacityAboutAsgRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点池所关联的伸缩组的期望实例数
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`
}

func (r *ModifyNodePoolDesiredCapacityAboutAsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodePoolDesiredCapacityAboutAsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "DesiredCapacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNodePoolDesiredCapacityAboutAsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodePoolDesiredCapacityAboutAsgResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNodePoolDesiredCapacityAboutAsgResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNodePoolDesiredCapacityAboutAsgResponseParams `json:"Response"`
}

func (r *ModifyNodePoolDesiredCapacityAboutAsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodePoolDesiredCapacityAboutAsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodePoolInstanceTypesRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 机型列表，参考 https://cloud.tencent.com/document/product/213/11518，主实例机型不支持修改
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`
}

type ModifyNodePoolInstanceTypesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 机型列表，参考 https://cloud.tencent.com/document/product/213/11518，主实例机型不支持修改
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`
}

func (r *ModifyNodePoolInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodePoolInstanceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNodePoolInstanceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNodePoolInstanceTypesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNodePoolInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNodePoolInstanceTypesResponseParams `json:"Response"`
}

func (r *ModifyNodePoolInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNodePoolInstanceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOpenPolicyListRequestParams struct {
	// 集群ID，请从容器服务集群控制台获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 修改的策略列表，目前仅支持修改EnforcementAction字段
	OpenPolicyInfoList []*OpenPolicySwitch `json:"OpenPolicyInfoList,omitnil,omitempty" name:"OpenPolicyInfoList"`

	// 策略分类 基线：baseline 优选：priority 可选：optional
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

type ModifyOpenPolicyListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，请从容器服务集群控制台获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 修改的策略列表，目前仅支持修改EnforcementAction字段
	OpenPolicyInfoList []*OpenPolicySwitch `json:"OpenPolicyInfoList,omitnil,omitempty" name:"OpenPolicyInfoList"`

	// 策略分类 基线：baseline 优选：priority 可选：optional
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

func (r *ModifyOpenPolicyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOpenPolicyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OpenPolicyInfoList")
	delete(f, "Category")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOpenPolicyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOpenPolicyListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOpenPolicyListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOpenPolicyListResponseParams `json:"Response"`
}

func (r *ModifyOpenPolicyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOpenPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAgentExternalLabelsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新的external_labels
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`
}

type ModifyPrometheusAgentExternalLabelsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新的external_labels
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`
}

func (r *ModifyPrometheusAgentExternalLabelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAgentExternalLabelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterId")
	delete(f, "ExternalLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAgentExternalLabelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAgentExternalLabelsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusAgentExternalLabelsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAgentExternalLabelsResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAgentExternalLabelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAgentExternalLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertPolicyRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

type ModifyPrometheusAlertPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertPolicyItem `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

func (r *ModifyPrometheusAlertPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAlertPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusAlertPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAlertPolicyResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAlertPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertRuleRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

type ModifyPrometheusAlertRuleRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警配置
	AlertRule *PrometheusAlertRuleDetail `json:"AlertRule,omitnil,omitempty" name:"AlertRule"`
}

func (r *ModifyPrometheusAlertRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AlertRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusAlertRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusAlertRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusAlertRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusAlertRuleResponseParams `json:"Response"`
}

func (r *ModifyPrometheusAlertRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusAlertRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusConfigRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Probes 配置
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`
}

type ModifyPrometheusConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ServiceMonitors配置
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// PodMonitors配置
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// prometheus原生Job配置
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// Probes 配置
	Probes []*PrometheusConfigItem `json:"Probes,omitnil,omitempty" name:"Probes"`
}

func (r *ModifyPrometheusConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterType")
	delete(f, "ClusterId")
	delete(f, "ServiceMonitors")
	delete(f, "PodMonitors")
	delete(f, "RawJobs")
	delete(f, "Probes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusConfigResponseParams `json:"Response"`
}

func (r *ModifyPrometheusConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusGlobalNotificationRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

type ModifyPrometheusGlobalNotificationRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 告警通知渠道
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`
}

func (r *ModifyPrometheusGlobalNotificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusGlobalNotificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Notification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusGlobalNotificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusGlobalNotificationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusGlobalNotificationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusGlobalNotificationResponseParams `json:"Response"`
}

func (r *ModifyPrometheusGlobalNotificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusGlobalNotificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusRecordRuleYamlRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 聚合实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 新的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ModifyPrometheusRecordRuleYamlRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 聚合实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 新的内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ModifyPrometheusRecordRuleYamlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusRecordRuleYamlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusRecordRuleYamlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusRecordRuleYamlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusRecordRuleYamlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusRecordRuleYamlResponseParams `json:"Response"`
}

func (r *ModifyPrometheusRecordRuleYamlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusRecordRuleYamlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTempRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 修改内容
	Template *PrometheusTempModify `json:"Template,omitnil,omitempty" name:"Template"`
}

type ModifyPrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 修改内容
	Template *PrometheusTempModify `json:"Template,omitnil,omitempty" name:"Template"`
}

func (r *ModifyPrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTempResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusTempResponseParams `json:"Response"`
}

func (r *ModifyPrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTemplateRequestParams struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 修改内容
	Template *PrometheusTemplateModify `json:"Template,omitnil,omitempty" name:"Template"`
}

type ModifyPrometheusTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 修改内容
	Template *PrometheusTemplateModify `json:"Template,omitnil,omitempty" name:"Template"`
}

func (r *ModifyPrometheusTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Template")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrometheusTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrometheusTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPrometheusTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrometheusTemplateResponseParams `json:"Response"`
}

func (r *ModifyPrometheusTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrometheusTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReservedInstanceScopeRequestParams struct {
	// 预留券实例ID。可通过 [DescribeReservedInstances](https://cloud.tencent.com/document/product/457/99162) 接口返回值中的ReservedInstanceId获取。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`

	// 预留券抵扣范围信息
	ReservedInstanceScope *ReservedInstanceScope `json:"ReservedInstanceScope,omitnil,omitempty" name:"ReservedInstanceScope"`
}

type ModifyReservedInstanceScopeRequest struct {
	*tchttp.BaseRequest
	
	// 预留券实例ID。可通过 [DescribeReservedInstances](https://cloud.tencent.com/document/product/457/99162) 接口返回值中的ReservedInstanceId获取。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`

	// 预留券抵扣范围信息
	ReservedInstanceScope *ReservedInstanceScope `json:"ReservedInstanceScope,omitnil,omitempty" name:"ReservedInstanceScope"`
}

func (r *ModifyReservedInstanceScopeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReservedInstanceScopeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReservedInstanceIds")
	delete(f, "ReservedInstanceScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReservedInstanceScopeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReservedInstanceScopeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReservedInstanceScopeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReservedInstanceScopeResponseParams `json:"Response"`
}

func (r *ModifyReservedInstanceScopeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReservedInstanceScopeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRollOutSequenceRequestParams struct {
	// 发布序列ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 发布序列名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布序列步骤
	SequenceFlows []*SequenceFlow `json:"SequenceFlows,omitnil,omitempty" name:"SequenceFlows"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ModifyRollOutSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 发布序列ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 发布序列名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布序列步骤
	SequenceFlows []*SequenceFlow `json:"SequenceFlows,omitnil,omitempty" name:"SequenceFlows"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ModifyRollOutSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRollOutSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "Name")
	delete(f, "SequenceFlows")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRollOutSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRollOutSequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRollOutSequenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRollOutSequenceResponseParams `json:"Response"`
}

func (r *ModifyRollOutSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRollOutSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NfsVolume struct {
	// nfs volume 数据卷名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// NFS 数据卷路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// NFS 服务器地址
	Server *string `json:"Server,omitnil,omitempty" name:"Server"`

	// 默认为 False
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type NodeCountSummary struct {
	// 手动管理的节点
	ManuallyAdded *ManuallyAdded `json:"ManuallyAdded,omitnil,omitempty" name:"ManuallyAdded"`

	// 自动管理的节点
	AutoscalingAdded *AutoscalingAdded `json:"AutoscalingAdded,omitnil,omitempty" name:"AutoscalingAdded"`
}

type NodePool struct {
	// NodePoolId 资源池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// Name 资源池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// ClusterInstanceId 集群实例id
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`

	// LifeState 状态，当前节点池生命周期状态包括：creating，normal，updating，deleting，deleted
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// LaunchConfigurationId 配置
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// AutoscalingGroupId 分组id
	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitnil,omitempty" name:"AutoscalingGroupId"`

	// Labels 标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// Taints 污点标记
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`

	// 节点 Annotation 列表
	Annotations []*AnnotationValue `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// NodeCountSummary 节点列表
	NodeCountSummary *NodeCountSummary `json:"NodeCountSummary,omitnil,omitempty" name:"NodeCountSummary"`

	// 状态信息
	AutoscalingGroupStatus *string `json:"AutoscalingGroupStatus,omitnil,omitempty" name:"AutoscalingGroupStatus"`

	// 最大节点数量
	MaxNodesNum *int64 `json:"MaxNodesNum,omitnil,omitempty" name:"MaxNodesNum"`

	// 最小节点数量
	MinNodesNum *int64 `json:"MinNodesNum,omitnil,omitempty" name:"MinNodesNum"`

	// 期望的节点数量
	DesiredNodesNum *int64 `json:"DesiredNodesNum,omitnil,omitempty" name:"DesiredNodesNum"`

	// 运行时描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitnil,omitempty" name:"RuntimeConfig"`

	// 节点池osName
	NodePoolOs *string `json:"NodePoolOs,omitnil,omitempty" name:"NodePoolOs"`

	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// 镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 集群属于节点podCIDR大小自定义模式时，节点池需要带上pod数量属性
	DesiredPodNum *int64 `json:"DesiredPodNum,omitnil,omitempty" name:"DesiredPodNum"`

	// 用户自定义脚本
	UserScript *string `json:"UserScript,omitnil,omitempty" name:"UserScript"`

	// 资源标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 删除保护开关
	DeletionProtection *bool `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 节点配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// GPU驱动相关参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GPUArgs *GPUArgs `json:"GPUArgs,omitnil,omitempty" name:"GPUArgs"`

	// dockerd --graph 指定值, 默认为 /var/lib/docker
	DockerGraphPath *string `json:"DockerGraphPath,omitnil,omitempty" name:"DockerGraphPath"`

	// 多盘数据盘挂载信息：新建节点时请确保购买CVM的参数传递了购买多个数据盘的信息，如CreateClusterInstances API的RunInstancesPara下的DataDisks也需要设置购买多个数据盘, 具体可以参考CreateClusterInstances接口的添加集群节点(多块数据盘)样例；添加已有节点时，请确保填写的分区信息在节点上真实存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 是否不可调度
	Unschedulable *int64 `json:"Unschedulable,omitnil,omitempty" name:"Unschedulable"`

	// 用户自定义脚本,在UserScript前执行
	PreStartUserScript *string `json:"PreStartUserScript,omitnil,omitempty" name:"PreStartUserScript"`
}

type NodePoolOption struct {
	// 是否加入节点池
	AddToNodePool *bool `json:"AddToNodePool,omitnil,omitempty" name:"AddToNodePool"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 是否继承节点池相关配置
	InheritConfigurationFromNodePool *bool `json:"InheritConfigurationFromNodePool,omitnil,omitempty" name:"InheritConfigurationFromNodePool"`
}

type NodePoolRuntime struct {
	// 节点池ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 运行时类型，参考：https://cloud.tencent.com/document/api/457/105241
	RuntimeType *string `json:"RuntimeType,omitnil,omitempty" name:"RuntimeType"`

	// 运行时版本，参考：https://cloud.tencent.com/document/api/457/105241
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// 节点池名称，限制 255 字符
	NodePoolName *string `json:"NodePoolName,omitnil,omitempty" name:"NodePoolName"`
}

type OIDCConfigAuthenticationOptions struct {
	// 创建身份提供商
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCreateOIDCConfig *bool `json:"AutoCreateOIDCConfig,omitnil,omitempty" name:"AutoCreateOIDCConfig"`

	// 创建身份提供商的ClientId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCreateClientId []*string `json:"AutoCreateClientId,omitnil,omitempty" name:"AutoCreateClientId"`

	// 创建PodIdentityWebhook组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoInstallPodIdentityWebhookAddon *bool `json:"AutoInstallPodIdentityWebhookAddon,omitnil,omitempty" name:"AutoInstallPodIdentityWebhookAddon"`
}

type OSImage struct {
	// os聚合名称
	SeriesName *string `json:"SeriesName,omitnil,omitempty" name:"SeriesName"`

	// os别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// os名称
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 操作系统类型(分为定制和非定制，取值分别为:DOCKER_CUSTOMIZE、GENERAL)
	OsCustomizeType *string `json:"OsCustomizeType,omitnil,omitempty" name:"OsCustomizeType"`

	// os是否下线(online表示在线,offline表示下线)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 镜像id
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`
}

type OpenConstraintInfo struct {
	// 策略实例关联事件数
	EventNums *uint64 `json:"EventNums,omitnil,omitempty" name:"EventNums"`

	// 策略实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例yaml详情base64编码
	YamlDetail *string `json:"YamlDetail,omitnil,omitempty" name:"YamlDetail"`
}

type OpenPolicyInfo struct {
	// 策略分类：cluster集群策略、node节点策略、namespace命名空间策略、configuration配置相关策略、compute计算资源策略、storage存储资源策略、network网络资源策略
	PolicyCategory *string `json:"PolicyCategory,omitnil,omitempty" name:"PolicyCategory"`

	// 策略中文名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略描述
	PolicyDesc *string `json:"PolicyDesc,omitnil,omitempty" name:"PolicyDesc"`

	// 策略运行模式：dryrun空跑不生效，deny拦截生效
	EnforcementAction *string `json:"EnforcementAction,omitnil,omitempty" name:"EnforcementAction"`

	// 关联的事件数量(最近7d)
	EventNums *uint64 `json:"EventNums,omitnil,omitempty" name:"EventNums"`

	// 策略英文名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略模板类型，支持的类型如下：
	// 优选策略：
	// blocknamespacedeletion：存在pod的命名空间不允许删除
	// blockcrddeletion：存在cr的crd不允许删除
	// blockmountablevolumetype：禁止挂载指定的volume类型
	// disallowalwayspullimage：禁止镜像拉取策略使用Always
	// tkeallowedrepos：容器镜像来源限制
	// blockunknowndaemonset：禁止未知的DaemonSet部署
	// blockpvdeletion：PV处于绑定状态则不允许删除
	// corednsprotect：CoreDNS组件删除保护
	// blockschedulablenodedelete：非封锁状态的Node不允许删除
	// resourcesdeletionprotection：资源删除保护
	// tkeenirequest：弹性网卡资源配置限制
	// blockworkloadcrossversionupgrade：工作负载镜像版本升级策略管控
	// blockserviceaccountgranthighprivilegepermission：ServiceAccount权限管控
	// blockclusteripserviceexist：不允许Service为ClusterIP类型
	// blockinternetaccess：禁止公网访问
	// assign：禁止访问Metadata Server
	// blockhostnetworkpod：禁止创建HostNetwork类型Pod
	// 
	// 可选策略：
	// blockvolumemountpath：禁止容器挂载指定的目录
	// k8sallowedrepos：容器镜像必须以指定字符串列表中的字符串开头
	// k8sblockendpointeditdefaultrole：禁止默认ClusterRole修改Endpoints
	// k8sblockloadbalancer：不允许Service为LoadBalancer类型
	// k8sblocknodeport：不允许Service为NodePort类型
	// k8sblockwildcardingress：禁止ingress配置空白或通配符类型的hostname
	// k8scontainerlimits：限制容器必须设置CPU和内存Limit
	// k8scontainerratios：限制CPU和内存的Request与Limit的最大比率
	// k8scontainerrequests：限制CPU和内存的Request必须设置且小于配置的最大值
	// k8srequiredresources：必须配置内存的Limit，CPU和内存的Request
	// k8sdisallowanonymous：不允许将白名单以外的ClusterRole和Role关联到system:anonymous User和system:unauthenticated Group
	// k8sdisallowedtags：约束容器镜像tag
	// k8sexternalips：限制服务externalIP仅为允许的IP地址列表
	// k8simagedigests：容器镜像必须包含digest
	// noupdateserviceaccount：拒绝白名单外的资源更新ServiceAccount
	// k8sreplicalimits：要求具有spec.replicas字段的对象（Deployments、ReplicaSets等）在定义的范围内
	// k8srequiredannotations：要求资源包含指定的annotations，其值与提供的正则表达式匹配
	// k8srequiredlabels：要求资源包含指定的标签，其值与提供的正则表达式匹配
	// k8srequiredprobes：要求Pod具有Readiness或Liveness Probe
	// k8spspautomountserviceaccounttokenpod：约束容器不能设置automountServiceAccountToken为true
	// k8spspallowprivilegeescalationcontainer：约束PodSecurityPolicy中的allowPrivilegeEscalation字段为false
	// k8spspapparmor：约束AppArmor字段列表
	// k8spspcapabilities：限制PodSecurityPolicy中的allowedCapabilities和requiredDropCapabilities字段
	// k8spspflexvolumes：约束PodSecurityPolicy中的allowedFlexVolumes字段类型
	// k8spspforbiddensysctls：约束PodSecurityPolicy中的sysctls字段不能使用的name
	// k8spspfsgroup：控制PodSecurityPolicy中的fsGroup字段在限制范围内
	// k8spsphostfilesystem：约束PodSecurityPolicy中的hostPath字段的参数
	// k8spsphostnamespace：限制PodSecurityPolicy中的hostPID和hostIPC字段
	// k8spsphostnetworkingports：约束PodSecurityPolicy中的hostNetwork和hostPorts字段
	// k8spspprivilegedcontainer：禁止PodSecurityPolicy中的privileged字段为true
	// k8spspprocmount：约束PodSecurityPolicy中的allowedProcMountTypes字段
	// k8spspreadonlyrootfilesystem：约束PodSecurityPolicy中的readOnlyRootFilesystem字段
	// k8spspseccomp：约束PodSecurityPolicy上的seccomp.security.alpha.kubernetes.io/allowedProfileNames注解
	// k8spspselinuxv2：约束Pod定义SELinux配置的允许列表
	// k8spspallowedusers：约束PodSecurityPolicy中的runAsUser、runAsGroup、supplementalGroups和fsGroup字段
	// k8spspvolumetypes：约束PodSecurityPolicy中的volumes字段类型
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 策略开关状态：open打开，close关闭
	EnabledStatus *string `json:"EnabledStatus,omitnil,omitempty" name:"EnabledStatus"`

	// 策略的实例的yaml示例base64编码
	ConstraintYamlExample *string `json:"ConstraintYamlExample,omitnil,omitempty" name:"ConstraintYamlExample"`

	// 策略关联的实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenConstraintInfoList []*OpenConstraintInfo `json:"OpenConstraintInfoList,omitnil,omitempty" name:"OpenConstraintInfoList"`
}

type OpenPolicySwitch struct {
	// 策略运行模式：dryrun空跑不生效，deny拦截生效
	EnforcementAction *string `json:"EnforcementAction,omitnil,omitempty" name:"EnforcementAction"`

	// 策略英文名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略模板类型，可选值：
	// 优选策略：
	// blocknamespacedeletion：存在pod的命名空间不允许删除
	// blockcrddeletion：存在cr的crd不允许删除
	// blockmountablevolumetype：禁止挂载指定的volume类型
	// disallowalwayspullimage：禁止镜像拉取策略使用Always
	// tkeallowedrepos：容器镜像来源限制
	// blockunknowndaemonset：禁止未知的DaemonSet部署
	// blockpvdeletion：PV处于绑定状态则不允许删除
	// corednsprotect：CoreDNS组件删除保护
	// blockschedulablenodedelete：非封锁状态的Node不允许删除
	// resourcesdeletionprotection：资源删除保护
	// tkeenirequest：弹性网卡资源配置限制
	// blockworkloadcrossversionupgrade：工作负载镜像版本升级策略管控
	// blockserviceaccountgranthighprivilegepermission：ServiceAccount权限管控
	// blockclusteripserviceexist：不允许Service为ClusterIP类型
	// blockinternetaccess：禁止公网访问
	// assign：禁止访问Metadata Server
	// blockhostnetworkpod：禁止创建HostNetwork类型Pod
	// 
	// 可选策略：
	// blockvolumemountpath：禁止容器挂载指定的目录
	// k8sallowedrepos：容器镜像必须以指定字符串列表中的字符串开头
	// k8sblockendpointeditdefaultrole：禁止默认ClusterRole修改Endpoints
	// k8sblockloadbalancer：不允许Service为LoadBalancer类型
	// k8sblocknodeport：不允许Service为NodePort类型
	// k8sblockwildcardingress：禁止ingress配置空白或通配符类型的hostname
	// k8scontainerlimits：限制容器必须设置CPU和内存Limit
	// k8scontainerratios：限制CPU和内存的Request与Limit的最大比率
	// k8scontainerrequests：限制CPU和内存的Request必须设置且小于配置的最大值
	// k8srequiredresources：必须配置内存的Limit，CPU和内存的Request
	// k8sdisallowanonymous：不允许将白名单以外的ClusterRole和Role关联到system:anonymous User和system:unauthenticated Group
	// k8sdisallowedtags：约束容器镜像tag
	// k8sexternalips：限制服务externalIP仅为允许的IP地址列表
	// k8simagedigests：容器镜像必须包含digest
	// noupdateserviceaccount：拒绝白名单外的资源更新ServiceAccount
	// k8sreplicalimits：要求具有spec.replicas字段的对象（Deployments、ReplicaSets等）在定义的范围内
	// k8srequiredannotations：要求资源包含指定的annotations，其值与提供的正则表达式匹配
	// k8srequiredlabels：要求资源包含指定的标签，其值与提供的正则表达式匹配
	// k8srequiredprobes：要求Pod具有Readiness或Liveness Probe
	// k8spspautomountserviceaccounttokenpod：约束容器不能设置automountServiceAccountToken为true
	// k8spspallowprivilegeescalationcontainer：约束PodSecurityPolicy中的allowPrivilegeEscalation字段为false
	// k8spspapparmor：约束AppArmor字段列表
	// k8spspcapabilities：限制PodSecurityPolicy中的allowedCapabilities和requiredDropCapabilities字段
	// k8spspflexvolumes：约束PodSecurityPolicy中的allowedFlexVolumes字段类型
	// k8spspforbiddensysctls：约束PodSecurityPolicy中的sysctls字段不能使用的name
	// k8spspfsgroup：控制PodSecurityPolicy中的fsGroup字段在限制范围内
	// k8spsphostfilesystem：约束PodSecurityPolicy中的hostPath字段的参数
	// k8spsphostnamespace：限制PodSecurityPolicy中的hostPID和hostIPC字段
	// k8spsphostnetworkingports：约束PodSecurityPolicy中的hostNetwork和hostPorts字段
	// k8spspprivilegedcontainer：禁止PodSecurityPolicy中的privileged字段为true
	// k8spspprocmount：约束PodSecurityPolicy中的allowedProcMountTypes字段
	// k8spspreadonlyrootfilesystem：约束PodSecurityPolicy中的readOnlyRootFilesystem字段
	// k8spspseccomp：约束PodSecurityPolicy上的seccomp.security.alpha.kubernetes.io/allowedProfileNames注解
	// k8spspselinuxv2：约束Pod定义SELinux配置的允许列表
	// k8spspallowedusers：约束PodSecurityPolicy中的runAsUser、runAsGroup、supplementalGroups和fsGroup字段
	// k8spspvolumetypes：约束PodSecurityPolicy中的volumes字段类型
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 策略开关状态：open打开，close关闭
	EnabledStatus *string `json:"EnabledStatus,omitnil,omitempty" name:"EnabledStatus"`

	// 策略关联的实例列表
	OpenConstraintInfoList []*OpenConstraintInfo `json:"OpenConstraintInfoList,omitnil,omitempty" name:"OpenConstraintInfoList"`
}

type OptionalRuntimes struct {
	// 运行时类型
	RuntimeType *string `json:"RuntimeType,omitnil,omitempty" name:"RuntimeType"`

	// 运行时版本列表
	RuntimeVersions []*string `json:"RuntimeVersions,omitnil,omitempty" name:"RuntimeVersions"`

	// 该类型的默认运行时版本
	DefaultVersion *string `json:"DefaultVersion,omitnil,omitempty" name:"DefaultVersion"`
}

type PendingRelease struct {
	// 应用状态详情
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 应用ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 应用状态，参考 Helm 发布状态。
	// 可选值及其释义如下：
	// • ​​unknown​​: 状态未知
	// • ​​deployed​​: 已成功部署
	// • ​​uninstalled​​: 已卸载
	// • ​​superseded​​: 已被新版本替代
	// • ​​failed​​: 部署失败
	// • ​​uninstalling​​: 正在卸载中
	// • ​​pending-install​​: 等待安装/安装进行中
	// • ​​pending-upgrade​​: 等待升级/升级进行中
	// • ​​pending-rollback​​: 等待回滚/回滚进行中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`
}

type PodChargeInfo struct {
	// Pod计费开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Pod的Uid
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// Pod的CPU
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Pod的内存
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// Pod类型：intel、amd、v100、t4、a10\*gnv4、a10\*gnv4v等。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Pod是GPU时，表示GPU卡数
	Gpu *string `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 计费类型
	// PREPAID：Pod调度到包月超级节点
	// POSTPAID_BY_HOUR：按量计费
	// RESERVED_INSTANCE：上个周期被预留券抵扣
	// SPOT：竞价实例
	// TPOD：特惠实例
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// Pod名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type PodDeductionRate struct {
	// Pod的 CPU
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// Pod 的内存，单位：GiB
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	//  Pod 的类型， intel，amd，windows-common，windows-amd，sa4，sa5，s7，s8，t4，v100，l20，l40，a10\*gnv4，a10\*gnv4v，a10\*pnv4
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	//  Pod 的 GPU 卡数，Pod 类型为 GPU 时有效。
	GpuNum *string `json:"GpuNum,omitnil,omitempty" name:"GpuNum"`

	// 这种规格的 Pod总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 这种规格的 Pod被预留券抵扣的数量
	DeductionNum *uint64 `json:"DeductionNum,omitnil,omitempty" name:"DeductionNum"`
}

type PodLimitsByType struct {
	// TKE共享网卡非固定IP模式可支持的Pod数量
	TKERouteENINonStaticIP *int64 `json:"TKERouteENINonStaticIP,omitnil,omitempty" name:"TKERouteENINonStaticIP"`

	// TKE共享网卡固定IP模式可支持的Pod数量
	TKERouteENIStaticIP *int64 `json:"TKERouteENIStaticIP,omitnil,omitempty" name:"TKERouteENIStaticIP"`

	// TKE独立网卡模式可支持的Pod数量
	TKEDirectENI *int64 `json:"TKEDirectENI,omitnil,omitempty" name:"TKEDirectENI"`

	// TKE中继网卡模式可支持的Pod数量
	TKESubENI *int64 `json:"TKESubENI,omitnil,omitempty" name:"TKESubENI"`
}

type PodLimitsInstance struct {
	// 机型所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 机型所属机型族
	InstanceFamily *string `json:"InstanceFamily,omitnil,omitempty" name:"InstanceFamily"`

	// 实例机型名称
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 机型可支持的最大VPC-CNI模式Pod数量信息
	PodLimits *PodLimitsByType `json:"PodLimits,omitnil,omitempty" name:"PodLimits"`
}

type PodNodeInfo struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	//  Pod 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type Probe struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureThreshold *int64 `json:"FailureThreshold,omitnil,omitempty" name:"FailureThreshold"`

	// Number of seconds after the container has started before liveness probes are initiated.
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil,omitempty" name:"PeriodSeconds"`

	// Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessThreshold *int64 `json:"SuccessThreshold,omitnil,omitempty" name:"SuccessThreshold"`

	// Number of seconds after which the probe times out.
	// Defaults to 1 second. Minimum value is 1.
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutSeconds *int64 `json:"TimeoutSeconds,omitnil,omitempty" name:"TimeoutSeconds"`
}

type PrometheusAgentInfo struct {
	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 集群所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type PrometheusAgentOverview struct {
	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// agent状态
	// normal = 正常
	// abnormal = 异常
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 额外labels
	// 本集群的所有指标都会带上这几个label
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`

	// 集群所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群所在VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 记录关联等操作的失败信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`
}

type PrometheusAlertHistoryItem struct {
	// 告警名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 告警开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 告警内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 告警状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 触发的规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleItem *string `json:"RuleItem,omitnil,omitempty" name:"RuleItem"`

	// 告警渠道的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 告警渠道的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type PrometheusAlertManagerConfig struct {
	// alertmanager url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// alertmanager部署所在集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// alertmanager部署所在集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type PrometheusAlertPolicyItem struct {
	// 策略名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则列表
	Rules []*PrometheusAlertRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 告警策略 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 如果该告警来自模板下发，则TemplateId为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 告警渠道，模板中使用可能返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notification *PrometheusNotificationItem `json:"Notification,omitnil,omitempty" name:"Notification"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 如果告警策略来源于用户集群CRD资源定义，则ClusterId为所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type PrometheusAlertRule struct {
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// prometheus语句
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 额外标签
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 告警发送模板
	Template *string `json:"Template,omitnil,omitempty" name:"Template"`

	// 持续时间
	For *string `json:"For,omitnil,omitempty" name:"For"`

	// 该条规则的描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 参考prometheus rule中的annotations
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotations []*Label `json:"Annotations,omitnil,omitempty" name:"Annotations"`

	// 告警规则状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleState *int64 `json:"RuleState,omitnil,omitempty" name:"RuleState"`
}

type PrometheusAlertRuleDetail struct {
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则列表
	Rules []*PrometheusAlertRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 最后修改时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 告警渠道
	Notification *PrometheusNotification `json:"Notification,omitnil,omitempty" name:"Notification"`

	// 告警 id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 如果该告警来至模板下发，则TemplateId为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 计算周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`
}

type PrometheusClusterAgentBasic struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否开启公网CLB
	EnableExternal *bool `json:"EnableExternal,omitnil,omitempty" name:"EnableExternal"`

	// 集群内部署组件的pod配置
	InClusterPodConfig *PrometheusClusterAgentPodConfig `json:"InClusterPodConfig,omitnil,omitempty" name:"InClusterPodConfig"`

	// 该集群采集的所有指标都会带上这些labels
	ExternalLabels []*Label `json:"ExternalLabels,omitnil,omitempty" name:"ExternalLabels"`

	// 是否安装默认采集配置
	NotInstallBasicScrape *bool `json:"NotInstallBasicScrape,omitnil,omitempty" name:"NotInstallBasicScrape"`

	// 是否采集指标，true代表drop所有指标，false代表采集默认指标
	NotScrape *bool `json:"NotScrape,omitnil,omitempty" name:"NotScrape"`
}

type PrometheusClusterAgentPodConfig struct {
	// 是否使用HostNetWork
	HostNet *bool `json:"HostNet,omitnil,omitempty" name:"HostNet"`

	// 指定pod运行节点
	NodeSelector []*Label `json:"NodeSelector,omitnil,omitempty" name:"NodeSelector"`

	// 容忍污点
	Tolerations []*Toleration `json:"Tolerations,omitnil,omitempty" name:"Tolerations"`
}

type PrometheusConfigItem struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置内容
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 用于出参，如果该配置来至模板，则为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type PrometheusGrafanaInfo struct {
	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 域名，只有开启外网访问才有效果
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 内网地址，或者外网地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 是否开启了外网访问
	// close = 未开启外网访问
	// opening = 正在开启外网访问
	// open  = 已开启外网访问
	Internet *string `json:"Internet,omitnil,omitempty" name:"Internet"`

	// grafana管理员用户名
	AdminUser *string `json:"AdminUser,omitnil,omitempty" name:"AdminUser"`
}

type PrometheusInstanceOverview struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 实例子网Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例当前的状态
	// prepare_env = 初始化环境
	// install_suit = 安装组件
	// running = 运行中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// COS桶存储
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// grafana默认地址，如果开启外网访问得为域名，否则为内网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaURL *string `json:"GrafanaURL,omitnil,omitempty" name:"GrafanaURL"`

	// 关联集群总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundTotal *uint64 `json:"BoundTotal,omitnil,omitempty" name:"BoundTotal"`

	// 运行正常的集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundNormal *uint64 `json:"BoundNormal,omitnil,omitempty" name:"BoundNormal"`
}

type PrometheusInstancesOverview struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 运行状态（1:正在创建；2:运行中；3:异常；4:重启中；5:销毁中； 6:已停机； 7: 已删除）
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 计费状态（1:正常；2:过期; 3:销毁; 4:分配中; 5:分配失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeStatus *int64 `json:"ChargeStatus,omitnil,omitempty" name:"ChargeStatus"`

	// 是否开启 Grafana（0:不开启，1:开启）
	EnableGrafana *int64 `json:"EnableGrafana,omitnil,omitempty" name:"EnableGrafana"`

	// Grafana 面板 URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	GrafanaURL *string `json:"GrafanaURL,omitnil,omitempty" name:"GrafanaURL"`

	// 实例付费类型（1:试用版；2:预付费）
	InstanceChargeType *int64 `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 存储周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`

	// 购买的实例过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标记(0:不自动续费；1:开启自动续费；2:禁止自动续费；-1:无效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 绑定集群总数
	BoundTotal *int64 `json:"BoundTotal,omitnil,omitempty" name:"BoundTotal"`

	// 绑定集群正常状态总数
	BoundNormal *int64 `json:"BoundNormal,omitnil,omitempty" name:"BoundNormal"`
}

type PrometheusJobTargets struct {
	// 该Job的targets列表
	Targets []*PrometheusTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// job的名称
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// targets总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 健康的target总数
	Up *uint64 `json:"Up,omitnil,omitempty" name:"Up"`
}

type PrometheusNotification struct {
	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 收敛时间
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// 生效起始时间
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// 生效结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 告警通知方式。目前有SMS、EMAIL、CALL、WECHAT方式。
	// 分别代表：短信、邮件、电话、微信
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// 告警接收组（用户组）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverGroups []*uint64 `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// 电话告警顺序。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNotifyOrder []*uint64 `json:"PhoneNotifyOrder,omitnil,omitempty" name:"PhoneNotifyOrder"`

	// 电话告警次数。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitnil,omitempty" name:"PhoneCircleTimes"`

	// 电话告警轮内间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitnil,omitempty" name:"PhoneInnerInterval"`

	// 电话告警轮外间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitnil,omitempty" name:"PhoneCircleInterval"`

	// 电话告警触达通知
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneArriveNotice *bool `json:"PhoneArriveNotice,omitnil,omitempty" name:"PhoneArriveNotice"`

	// 通道类型，默认为amp，支持以下
	// amp
	// webhook
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 如果Type为webhook, 则该字段为必填项
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebHook *string `json:"WebHook,omitnil,omitempty" name:"WebHook"`
}

type PrometheusNotificationItem struct {
	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 通道类型，默认为amp，支持以下
	// amp
	// webhook
	// alertmanager
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 如果Type为webhook, 则该字段为必填项
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebHook *string `json:"WebHook,omitnil,omitempty" name:"WebHook"`

	// 如果Type为alertmanager, 则该字段为必填项
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertManager *PrometheusAlertManagerConfig `json:"AlertManager,omitnil,omitempty" name:"AlertManager"`

	// 收敛时间
	RepeatInterval *string `json:"RepeatInterval,omitnil,omitempty" name:"RepeatInterval"`

	// 生效起始时间
	TimeRangeStart *string `json:"TimeRangeStart,omitnil,omitempty" name:"TimeRangeStart"`

	// 生效结束时间
	TimeRangeEnd *string `json:"TimeRangeEnd,omitnil,omitempty" name:"TimeRangeEnd"`

	// 告警通知方式。目前有SMS、EMAIL、CALL、WECHAT方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay []*string `json:"NotifyWay,omitnil,omitempty" name:"NotifyWay"`

	// 告警接收组（用户组）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverGroups []*string `json:"ReceiverGroups,omitnil,omitempty" name:"ReceiverGroups"`

	// 电话告警顺序。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNotifyOrder []*uint64 `json:"PhoneNotifyOrder,omitnil,omitempty" name:"PhoneNotifyOrder"`

	// 电话告警次数。
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitnil,omitempty" name:"PhoneCircleTimes"`

	// 电话告警轮内间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitnil,omitempty" name:"PhoneInnerInterval"`

	// 电话告警轮外间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitnil,omitempty" name:"PhoneCircleInterval"`

	// 电话告警触达通知
	// 注：NotifyWay选择CALL，采用该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneArriveNotice *bool `json:"PhoneArriveNotice,omitnil,omitempty" name:"PhoneArriveNotice"`
}

type PrometheusRecordRuleYamlItem struct {
	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 最近更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Yaml内容
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 如果该聚合规则来至模板，则TemplateId为模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 该聚合规则如果来源于用户集群crd资源定义，则ClusterId为所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type PrometheusTarget struct {
	// 抓取目标的URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// target当前状态,当前支持
	// up = 健康
	// down = 不健康
	// unknown = 未知
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// target的元label
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 上一次抓取的时间
	LastScrape *string `json:"LastScrape,omitnil,omitempty" name:"LastScrape"`

	// 上一次抓取的耗时，单位是s
	ScrapeDuration *float64 `json:"ScrapeDuration,omitnil,omitempty" name:"ScrapeDuration"`

	// 上一次抓取如果错误，该字段存储错误信息
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type PrometheusTemp struct {
	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板维度，支持以下类型
	// instance 实例级别
	// cluster 集群级别
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 当Level为instance时有效，
	// 模板中的聚合规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitnil,omitempty" name:"RecordRules"`

	// 当Level为cluster时有效，
	// 模板中的ServiceMonitor规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// 当Level为cluster时有效，
	// 模板中的PodMonitors规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 当Level为cluster时有效，
	// 模板中的RawJobs规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 模板的ID, 用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 最近更新时间，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 当前版本，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 是否系统提供的默认模板，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 当Level为instance时有效，
	// 模板中的告警配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertDetailRules []*PrometheusAlertPolicyItem `json:"AlertDetailRules,omitnil,omitempty" name:"AlertDetailRules"`

	// 关联实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetsTotal *int64 `json:"TargetsTotal,omitnil,omitempty" name:"TargetsTotal"`
}

type PrometheusTempModify struct {
	// 修改名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 修改描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 当Level为cluster时有效，
	// 模板中的ServiceMonitor规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// 当Level为cluster时有效，
	// 模板中的PodMonitors规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 当Level为cluster时有效，
	// 模板中的RawJobs规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 当Level为instance时有效，
	// 模板中的聚合规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitnil,omitempty" name:"RecordRules"`

	// 修改内容，只有当模板类型是Alert时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertDetailRules []*PrometheusAlertPolicyItem `json:"AlertDetailRules,omitnil,omitempty" name:"AlertDetailRules"`
}

type PrometheusTemplate struct {
	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板维度，支持以下类型
	// instance 实例级别
	// cluster 集群级别
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 当Level为instance时有效，
	// 模板中的告警配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertRules []*PrometheusAlertRule `json:"AlertRules,omitnil,omitempty" name:"AlertRules"`

	// 当Level为instance时有效，
	// 模板中的聚合规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitnil,omitempty" name:"RecordRules"`

	// 当Level为cluster时有效，
	// 模板中的ServiceMonitor规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// 当Level为cluster时有效，
	// 模板中的PodMonitors规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 当Level为cluster时有效，
	// 模板中的RawJobs规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 模板的ID, 用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 最近更新时间，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 当前版本，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 是否系统提供的默认模板，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 当Level为instance时有效，
	// 模板中的告警配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertDetailRules []*PrometheusAlertRuleDetail `json:"AlertDetailRules,omitnil,omitempty" name:"AlertDetailRules"`
}

type PrometheusTemplateModify struct {
	// 修改名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 修改描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`

	// 修改内容，只有当模板类型是Alert时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertRules []*PrometheusAlertRule `json:"AlertRules,omitnil,omitempty" name:"AlertRules"`

	// 当Level为instance时有效，
	// 模板中的聚合规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordRules []*PrometheusConfigItem `json:"RecordRules,omitnil,omitempty" name:"RecordRules"`

	// 当Level为cluster时有效，
	// 模板中的ServiceMonitor规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceMonitors []*PrometheusConfigItem `json:"ServiceMonitors,omitnil,omitempty" name:"ServiceMonitors"`

	// 当Level为cluster时有效，
	// 模板中的PodMonitors规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodMonitors []*PrometheusConfigItem `json:"PodMonitors,omitnil,omitempty" name:"PodMonitors"`

	// 当Level为cluster时有效，
	// 模板中的RawJobs规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawJobs []*PrometheusConfigItem `json:"RawJobs,omitnil,omitempty" name:"RawJobs"`

	// 修改内容，只有当模板类型是Alert时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlertDetailRules []*PrometheusAlertRuleDetail `json:"AlertDetailRules,omitnil,omitempty" name:"AlertDetailRules"`
}

type PrometheusTemplateSyncTarget struct {
	// 目标所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 目标实例
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群id，只有当采集模板的Level为cluster的时候需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 最后一次同步时间， 用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncTime *string `json:"SyncTime,omitnil,omitempty" name:"SyncTime"`

	// 当前使用的模板版本，用于出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 集群类型，只有当采集模板的Level为cluster的时候需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 用于出参，实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用于出参，集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type RIUtilizationDetail struct {
	// 预留券ID
	ReservedInstanceId *string `json:"ReservedInstanceId,omitnil,omitempty" name:"ReservedInstanceId"`

	// Pod唯一ID
	EksId *string `json:"EksId,omitnil,omitempty" name:"EksId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Pod的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Pod的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 工作负载类型，如 deployment、statefulset和pod等。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 工作负载名称
	KindName *string `json:"KindName,omitnil,omitempty" name:"KindName"`

	// Pod的uid
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用量开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 用量结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 抵扣资源所属产品
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type RegionInstance struct {
	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域状态\n- alluser: 所有用户可见
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域特性开关(按照JSON的形式返回所有属性)
	FeatureGates *string `json:"FeatureGates,omitnil,omitempty" name:"FeatureGates"`

	// 地域简称
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 地域白名单
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type Release struct {
	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 应用当前版本
	Revision *string `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 应用状态，参考 Helm 发布状态。
	// 可选值及其释义如下：
	// • ​​unknown​​: 状态未知
	// • ​​deployed​​: 已成功部署
	// • ​​uninstalled​​: 已卸载
	// • ​​superseded​​: 已被新版本替代
	// • ​​failed​​: 部署失败
	// • ​​uninstalling​​: 正在卸载中
	// • ​​pending-install​​: 等待安装/安装进行中
	// • ​​pending-upgrade​​: 等待升级/升级进行中
	// • ​​pending-rollback​​: 等待回滚/回滚进行中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 制品名称
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`

	// 制品版本
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`

	// 制品应用版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ReleaseDetails struct {
	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用所在命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 应用当前版本
	Version *uint64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 应用状态，参考 Helm 发布状态。 
	// 可选值及其释义如下：
	// • ​​unknown​​: 状态未知 
	// • ​​deployed​​: 已成功部署 
	// • ​​uninstalled​​: 已卸载 
	// • ​​superseded​​: 已被新版本替代 
	// • ​​failed​​: 部署失败 
	// • ​​uninstalling​​: 正在卸载中
	// • ​​pending-install​​: 等待安装/安装进行中
	// • ​​pending-upgrade​​: 等待升级/升级进行中 
	// • ​​pending-rollback​​: 等待回滚/回滚进行中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 应用提示
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 用户自定义参数
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 应用资源详情
	Manifest *string `json:"Manifest,omitnil,omitempty" name:"Manifest"`

	// 应用制品版本
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`

	// 应用制品名称
	ChartName *string `json:"ChartName,omitnil,omitempty" name:"ChartName"`

	// 应用制品描述
	ChartDescription *string `json:"ChartDescription,omitnil,omitempty" name:"ChartDescription"`

	// 应用制品app版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 应用首次部署时间
	FirstDeployedTime *string `json:"FirstDeployedTime,omitnil,omitempty" name:"FirstDeployedTime"`

	// 应用最近部署时间
	LastDeployedTime *string `json:"LastDeployedTime,omitnil,omitempty" name:"LastDeployedTime"`

	// 应用参数
	ComputedValues *string `json:"ComputedValues,omitnil,omitempty" name:"ComputedValues"`

	// chart 的来源， tke-market, others
	ChartFrom *string `json:"ChartFrom,omitnil,omitempty" name:"ChartFrom"`

	// 第三方chart 的安装地址
	ChartURL *string `json:"ChartURL,omitnil,omitempty" name:"ChartURL"`

	// 通过chart 创建的资源
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type ReleaseHistory struct {
	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 应用版本
	Revision *uint64 `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 应用状态，参考 Helm 发布状态。 
	// 可选值及其释义如下：
	// • ​​unknown​​: 状态未知 
	// • ​​deployed​​: 已成功部署 
	// • ​​uninstalled​​: 已卸载 
	// • ​​superseded​​: 已被新版本替代 
	// • ​​failed​​: 部署失败 
	// • ​​uninstalling​​: 正在卸载中
	// • ​​pending-install​​: 等待安装/安装进行中
	// • ​​pending-upgrade​​: 等待升级/升级进行中 
	// • ​​pending-rollback​​: 等待回滚/回滚进行中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 应用制品名称
	Chart *string `json:"Chart,omitnil,omitempty" name:"Chart"`

	// 应用制品版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 应用更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ReleaseValues struct {
	// 自定义参数原始值
	RawOriginal *string `json:"RawOriginal,omitnil,omitempty" name:"RawOriginal"`

	// 自定义参数值类型
	ValuesType *string `json:"ValuesType,omitnil,omitempty" name:"ValuesType"`
}

// Predefined struct for user
type RemoveNodeFromNodePoolRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点id列表，一次最多支持100台
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type RemoveNodeFromNodePoolRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点id列表，一次最多支持100台
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *RemoveNodeFromNodePoolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNodeFromNodePoolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveNodeFromNodePoolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNodeFromNodePoolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveNodeFromNodePoolResponse struct {
	*tchttp.BaseResponse
	Response *RemoveNodeFromNodePoolResponseParams `json:"Response"`
}

func (r *RemoveNodeFromNodePoolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNodeFromNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewReservedInstancesRequestParams struct {
	// 预留券实例ID。可通过 [DescribeReservedInstances](https://cloud.tencent.com/document/product/457/99162) 接口返回值中的ReservedInstanceId获取，每次请求实例的上限为100。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type RenewReservedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 预留券实例ID。可通过 [DescribeReservedInstances](https://cloud.tencent.com/document/product/457/99162) 接口返回值中的ReservedInstanceId获取，每次请求实例的上限为100。
	ReservedInstanceIds []*string `json:"ReservedInstanceIds,omitnil,omitempty" name:"ReservedInstanceIds"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

func (r *RenewReservedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewReservedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ReservedInstanceIds")
	delete(f, "InstanceChargePrepaid")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewReservedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewReservedInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewReservedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RenewReservedInstancesResponseParams `json:"Response"`
}

func (r *RenewReservedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewReservedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReservedInstance struct {
	// 预留实例ID
	ReservedInstanceId *string `json:"ReservedInstanceId,omitnil,omitempty" name:"ReservedInstanceId"`

	// 预留实例名称
	ReservedInstanceName *string `json:"ReservedInstanceName,omitnil,omitempty" name:"ReservedInstanceName"`

	// 预留券状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 有效期，单位：月
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 抵扣资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源核数
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 资源内存，单位：Gi
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 预留券的范围，默认值region。
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 生效时间
	ActiveAt *string `json:"ActiveAt,omitnil,omitempty" name:"ActiveAt"`

	// 过期时间
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// GPU卡数
	GpuCount *string `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// 自动续费标记
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	//  上个周期预留券的抵扣状态，Deduct：已抵扣、NotDeduct：未抵扣
	DeductStatus *string `json:"DeductStatus,omitnil,omitempty" name:"DeductStatus"`
}

type ReservedInstanceScope struct {
	// 抵扣范围，取值：Region、Zone 和 Node
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type ReservedInstanceSpec struct {
	// 资源类型：common、amd、v100、t4、a10\*gnv4、a10\*gnv4v、a10\*pnv4、windows-common、windows-amd，common表示通用类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 核数
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位：GiB
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// GPU卡数，当Type为GPU类型时设置。
	Gpu *float64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`
}

type ReservedInstanceUtilizationRate struct {
	// 使用率
	Rate *float64 `json:"Rate,omitnil,omitempty" name:"Rate"`

	// 预留券数量
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 核数
	CPU *float64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存，单位：GiB
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	//  预留券类型, common：CPU通用，amd：AMD专用，windows-common: Windows容器 CPU通用，windows-amd：Windows容器 AMD专用，sa4，sa5，s7，s8，t4，v100，l20，l40，a10\*gnv4，a10\*gnv4v，a10\*pnv4
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// GPU 卡数
	GpuNum *string `json:"GpuNum,omitnil,omitempty" name:"GpuNum"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点 ID
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// Pod 数量
	PodNum *uint64 `json:"PodNum,omitnil,omitempty" name:"PodNum"`
}

type ResourceDeleteOption struct {
	// 资源类型，例如CBS、CLB、CVM
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 集群删除时CBS资源的删除模式：terminate（销毁），retain （保留）。其他资源默认为销毁。
	DeleteMode *string `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`

	// 是否跳过开启删除保护的资源，默认false，设置为true时不清理开启了删除保护的资源，clb有终端节点的情况也属于开了删除保护。
	SkipDeletionProtection *bool `json:"SkipDeletionProtection,omitnil,omitempty" name:"SkipDeletionProtection"`
}

type ResourceUsage struct {
	// 资源类型，参考k8s 官方资源
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源使用量，单位：个数
	Usage *uint64 `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 资源使用详情
	Details []*ResourceUsageDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type ResourceUsageDetail struct {
	// 资源名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源使用量
	Usage *uint64 `json:"Usage,omitnil,omitempty" name:"Usage"`
}

// Predefined struct for user
type RestartEKSContainerInstancesRequestParams struct {
	// EKS instance ids
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`
}

type RestartEKSContainerInstancesRequest struct {
	*tchttp.BaseRequest
	
	// EKS instance ids
	EksCiIds []*string `json:"EksCiIds,omitnil,omitempty" name:"EksCiIds"`
}

func (r *RestartEKSContainerInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartEKSContainerInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EksCiIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartEKSContainerInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartEKSContainerInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartEKSContainerInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RestartEKSContainerInstancesResponseParams `json:"Response"`
}

func (r *RestartEKSContainerInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartEKSContainerInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollOutSequence struct {
	// 发布序列名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布序列步骤
	SequenceFlows []*SequenceFlow `json:"SequenceFlows,omitnil,omitempty" name:"SequenceFlows"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 发布序列ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`
}

// Predefined struct for user
type RollbackClusterReleaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 回滚版本号
	Revision *int64 `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 集群类型，传入 tke(标准集群)，eks(Serverless集群)，external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type RollbackClusterReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 回滚版本号
	Revision *int64 `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 集群类型，传入 tke(标准集群)，eks(Serverless集群)，external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *RollbackClusterReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackClusterReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "Revision")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackClusterReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackClusterReleaseResponseParams struct {
	// 应用详情
	Release *PendingRelease `json:"Release,omitnil,omitempty" name:"Release"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackClusterReleaseResponse struct {
	*tchttp.BaseResponse
	Response *RollbackClusterReleaseResponseParams `json:"Response"`
}

func (r *RollbackClusterReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackClusterReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteInfo struct {
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 目的端CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil,omitempty" name:"DestinationCidrBlock"`

	// 下一跳地址。
	GatewayIp *string `json:"GatewayIp,omitnil,omitempty" name:"GatewayIp"`
}

type RouteTableConflict struct {
	// 路由表类型。
	// 枚举值：CcsCluster、Vpc、VpcRouteTable、CcsClusterRouteTable
	RouteTableType *string `json:"RouteTableType,omitnil,omitempty" name:"RouteTableType"`

	// 路由表CIDR。
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitnil,omitempty" name:"RouteTableCidrBlock"`

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 路由表ID。
	RouteTableId *string `json:"RouteTableId,omitnil,omitempty" name:"RouteTableId"`
}

type RouteTableInfo struct {
	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil,omitempty" name:"RouteTableName"`

	// 路由表CIDR。
	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitnil,omitempty" name:"RouteTableCidrBlock"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type RunAutomationServiceEnabled struct {
	// 是否开启云自动化助手。取值范围：<br><li>true：表示开启云自动化助手服务<br><li>false：表示不开启云自动化助手服务<br><br>默认取值：false。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunInstancesForNode struct {
	// 节点角色，取值:MASTER_ETCD, WORKER。MASTER_ETCD只有在创建 INDEPENDENT_CLUSTER 独立集群时需要指定。MASTER_ETCD节点数量为3～7，建议为奇数。MASTER_ETCD节点最小配置为4C8G。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// CVM创建透传参数，json化字符串格式，详见[CVM创建实例](https://cloud.tencent.com/document/product/213/15730)接口，传入公共参数外的其他参数即可，其中ImageId会替换为TKE集群OS对应的镜像。
	RunInstancesPara []*string `json:"RunInstancesPara,omitnil,omitempty" name:"RunInstancesPara"`

	// 节点高级设置，该参数会覆盖集群级别设置的InstanceAdvancedSettings，和上边的RunInstancesPara按照顺序一一对应（当前只对节点自定义参数ExtraArgs生效）。
	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitnil,omitempty" name:"InstanceAdvancedSettingsOverrides"`
}

type RunMonitorServiceEnabled struct {
	// 是否开启[云监控](/document/product/248)服务。取值范围：<br><li>true：表示开启云监控服务</li><li>false：表示不开启云监控服务</li><br>默认取值：true。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type RunPrometheusInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子网ID，默认使用实例所用子网初始化，也可通过该参数传递新的子网ID初始化
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type RunPrometheusInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子网ID，默认使用实例所用子网初始化，也可通过该参数传递新的子网ID初始化
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

func (r *RunPrometheusInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunPrometheusInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunPrometheusInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunPrometheusInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunPrometheusInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RunPrometheusInstanceResponseParams `json:"Response"`
}

func (r *RunPrometheusInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunPrometheusInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSecurityServiceEnabled struct {
	// 是否开启[云安全](/document/product/296)服务。取值范围：<br><li>true：表示开启云安全服务<br><li>false：表示不开启云安全服务<br><br>默认取值：true。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RuntimeConfig struct {
	// 运行时类型，支持的类型有 docker、containerd
	RuntimeType *string `json:"RuntimeType,omitnil,omitempty" name:"RuntimeType"`

	// 运行时版本，参考：https://cloud.tencent.com/document/api/457/105241
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`
}

// Predefined struct for user
type ScaleInClusterMasterRequestParams struct {
	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// master缩容选项
	ScaleInMasters []*ScaleInMaster `json:"ScaleInMasters,omitnil,omitempty" name:"ScaleInMasters"`
}

type ScaleInClusterMasterRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// master缩容选项
	ScaleInMasters []*ScaleInMaster `json:"ScaleInMasters,omitnil,omitempty" name:"ScaleInMasters"`
}

func (r *ScaleInClusterMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleInClusterMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ScaleInMasters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleInClusterMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleInClusterMasterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleInClusterMasterResponse struct {
	*tchttp.BaseResponse
	Response *ScaleInClusterMasterResponseParams `json:"Response"`
}

func (r *ScaleInClusterMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleInClusterMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleInMaster struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 缩容的实例角色：MASTER,ETCD,MASTER_ETCD
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）
	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitnil,omitempty" name:"InstanceDeleteMode"`
}

// Predefined struct for user
type ScaleOutClusterMasterRequestParams struct {
	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新建节点参数
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitnil,omitempty" name:"RunInstancesForNode"`

	// 添加已有节点相关参数
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitnil,omitempty" name:"ExistedInstancesForNode"`

	// 实例高级设置
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 集群master组件自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`
}

type ScaleOutClusterMasterRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 新建节点参数
	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitnil,omitempty" name:"RunInstancesForNode"`

	// 添加已有节点相关参数
	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitnil,omitempty" name:"ExistedInstancesForNode"`

	// 实例高级设置
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 集群master组件自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`
}

func (r *ScaleOutClusterMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "RunInstancesForNode")
	delete(f, "ExistedInstancesForNode")
	delete(f, "InstanceAdvancedSettings")
	delete(f, "ExtraArgs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutClusterMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutClusterMasterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutClusterMasterResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutClusterMasterResponseParams `json:"Response"`
}

func (r *ScaleOutClusterMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutClusterMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityContext struct {
	// 安全能力清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Capabilities *Capabilities `json:"Capabilities,omitnil,omitempty" name:"Capabilities"`
}

type SequenceFlow struct {
	// 发布序列步骤标签
	Tags []*SequenceTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 等待时间（秒）
	SoakTime *int64 `json:"SoakTime,omitnil,omitempty" name:"SoakTime"`
}

type SequenceTag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value []*string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ServiceAccountAuthenticationOptions struct {
	// 使用TKE默认issuer和jwksuri
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseTKEDefault *bool `json:"UseTKEDefault,omitnil,omitempty" name:"UseTKEDefault"`

	// service-account-issuer
	// 注意：此字段可能返回 null，表示取不到有效值。
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// service-account-jwks-uri
	// 注意：此字段可能返回 null，表示取不到有效值。
	JWKSURI *string `json:"JWKSURI,omitnil,omitempty" name:"JWKSURI"`

	// 如果为true，则会自动创建允许匿名用户访问'/.well-known/openid-configuration'和/openid/v1/jwks的rbac规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCreateDiscoveryAnonymousAuth *bool `json:"AutoCreateDiscoveryAnonymousAuth,omitnil,omitempty" name:"AutoCreateDiscoveryAnonymousAuth"`
}

// Predefined struct for user
type SetNodePoolNodeProtectionRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点id
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 节点是否需要移出保护
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`
}

type SetNodePoolNodeProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 节点池id
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 节点id
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 节点是否需要移出保护
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`
}

func (r *SetNodePoolNodeProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNodePoolNodeProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodePoolId")
	delete(f, "InstanceIds")
	delete(f, "ProtectedFromScaleIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNodePoolNodeProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNodePoolNodeProtectionResponseParams struct {
	// 成功设置的节点id
	SucceedInstanceIds []*string `json:"SucceedInstanceIds,omitnil,omitempty" name:"SucceedInstanceIds"`

	// 没有成功设置的节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitnil,omitempty" name:"FailedInstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetNodePoolNodeProtectionResponse struct {
	*tchttp.BaseResponse
	Response *SetNodePoolNodeProtectionResponseParams `json:"Response"`
}

func (r *SetNodePoolNodeProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNodePoolNodeProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Step struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 结束时间
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 当前状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 执行信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type SubnetInfos struct {
	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 安全组id
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// 系统，默认linux
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 硬件架构，默认amd64
	Arch *string `json:"Arch,omitnil,omitempty" name:"Arch"`
}

type SuperNodeResource struct {
	// 节点名称，此字段在出参中有效。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点上的资源总数，QuotaType为 exact 时，表示创建指定规格的精确配额数量。
	Num *uint64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 节点上的总核数，QuotaType为 exact 时表示指定规格的核数。
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 节点上的总内存数，QuotaType为 exact 时表示指定规格的内存，单位：Gi
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 节点上的总 GPU 卡数，QuotaType为 exact 时表示指定规格的GPU卡数。
	Gpu *float64 `json:"Gpu,omitnil,omitempty" name:"Gpu"`

	// 节点资源的配额类型，exact表示精确配额，fuzzy 表示模糊配额。
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 配额的计费类型，PREPAID表示包月，POSTPAID_BY_HOUR表示按量。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// QuotaType为 exact 时，此字段有效，表示精确配额的资源类型。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 置放群组 ID，QuotaType为 exact 时有效，表示购买的精确配额需满足置放群组。可通过 [DescribeDisasterRecoverGroups](https://cloud.tencent.com/document/product/213/17810) 接口返回值中的DisasterRecoverGroupId获取。
	DisasterRecoverGroupId *string `json:"DisasterRecoverGroupId,omitnil,omitempty" name:"DisasterRecoverGroupId"`

	// 定价类型
	PriceType *string `json:"PriceType,omitnil,omitempty" name:"PriceType"`
}

type Switch struct {
	// 审计开关的详细信息
	Audit *SwitchInfo `json:"Audit,omitnil,omitempty" name:"Audit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 事件开关的详细信息
	Event *SwitchInfo `json:"Event,omitnil,omitempty" name:"Event"`

	// 普通日志的详细信息
	Log *SwitchInfo `json:"Log,omitnil,omitempty" name:"Log"`

	// master 日志详细信息
	MasterLog *SwitchInfo `json:"MasterLog,omitnil,omitempty" name:"MasterLog"`
}

// Predefined struct for user
type SwitchClusterEndpointRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`

	// 切换回滚至非直连
	Rollback *bool `json:"Rollback,omitnil,omitempty" name:"Rollback"`
}

type SwitchClusterEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）
	IsExtranet *bool `json:"IsExtranet,omitnil,omitempty" name:"IsExtranet"`

	// 切换回滚至非直连
	Rollback *bool `json:"Rollback,omitnil,omitempty" name:"Rollback"`
}

func (r *SwitchClusterEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IsExtranet")
	delete(f, "Rollback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchClusterEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchClusterEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchClusterEndpointResponse struct {
	*tchttp.BaseResponse
	Response *SwitchClusterEndpointResponseParams `json:"Response"`
}

func (r *SwitchClusterEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchInfo struct {
	// 开启标识符 true代表开启
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 获取日志状态失败时，返回错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// CLS日志集ID
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题状态，opened表示已开启，opening开启中，closed表示已关闭，closing 表示关闭中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CLS日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// CLS日志主题所属region
	TopicRegion *string `json:"TopicRegion,omitnil,omitempty" name:"TopicRegion"`

	// 是否可升级
	UpgradeAble *bool `json:"UpgradeAble,omitnil,omitempty" name:"UpgradeAble"`

	// 当前log-agent版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

// Predefined struct for user
type SyncPrometheusTempRequestParams struct {
	// 实例id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 同步目标
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type SyncPrometheusTempRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 同步目标
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *SyncPrometheusTempRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTempRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncPrometheusTempRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPrometheusTempResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncPrometheusTempResponse struct {
	*tchttp.BaseResponse
	Response *SyncPrometheusTempResponseParams `json:"Response"`
}

func (r *SyncPrometheusTempResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTempResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPrometheusTemplateRequestParams struct {
	// 实例id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 同步目标
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type SyncPrometheusTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 同步目标
	Targets []*PrometheusTemplateSyncTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *SyncPrometheusTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncPrometheusTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncPrometheusTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncPrometheusTemplateResponse struct {
	*tchttp.BaseResponse
	Response *SyncPrometheusTemplateResponseParams `json:"Response"`
}

func (r *SyncPrometheusTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncPrometheusTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TagSpecification struct {
	// 标签绑定的资源类型，当前支持类型："cluster"
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 标签对列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type Taint struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// Effect
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`
}

type Task struct {
	// 任务状态（process(运行中)、pause(暂停)、pausing(暂停中)、paused(已暂停)、done(已完成)、abort(中止)、aborted(已中止)、resume(重新执行)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 任务目标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObj *string `json:"TargetObj,omitnil,omitempty" name:"TargetObj"`

	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Param *string `json:"Param,omitnil,omitempty" name:"Param"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastError *string `json:"LastError,omitnil,omitempty" name:"LastError"`

	// 任务所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 任务更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 创建任务唯一请求ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskID *string `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type TaskStepInfo struct {
	// 步骤名称
	Step *string `json:"Step,omitnil,omitempty" name:"Step"`

	// 生命周期
	// pending : 步骤未开始
	// running: 步骤执行中
	// success: 步骤成功完成
	// failed: 步骤失败
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 步骤开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 步骤结束时间
	EndAt *string `json:"EndAt,omitnil,omitempty" name:"EndAt"`

	// 若步骤生命周期为failed,则此字段显示错误信息
	FailedMsg *string `json:"FailedMsg,omitnil,omitempty" name:"FailedMsg"`
}

type TcpSocket struct {
	// TcpSocket检测的端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type Toleration struct {
	// 容忍应用到的 taint key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 键与值的关系
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 要匹配的污点效果
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`
}

type UnavailableReason struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

// Predefined struct for user
type UninstallClusterReleaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群类型，传入 tke(标准集群)， eks(Serverless集群)，external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type UninstallClusterReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 集群类型，传入 tke(标准集群)， eks(Serverless集群)，external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *UninstallClusterReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallClusterReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallClusterReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallClusterReleaseResponseParams struct {
	// 应用详情
	Release *PendingRelease `json:"Release,omitnil,omitempty" name:"Release"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UninstallClusterReleaseResponse struct {
	*tchttp.BaseResponse
	Response *UninstallClusterReleaseResponseParams `json:"Response"`
}

func (r *UninstallClusterReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallClusterReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallEdgeLogAgentRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type UninstallEdgeLogAgentRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *UninstallEdgeLogAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallEdgeLogAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallEdgeLogAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallEdgeLogAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UninstallEdgeLogAgentResponse struct {
	*tchttp.BaseResponse
	Response *UninstallEdgeLogAgentResponseParams `json:"Response"`
}

func (r *UninstallEdgeLogAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallEdgeLogAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallLogAgentRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type UninstallLogAgentRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *UninstallLogAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallLogAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UninstallLogAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UninstallLogAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UninstallLogAgentResponse struct {
	*tchttp.BaseResponse
	Response *UninstallLogAgentResponseParams `json:"Response"`
}

func (r *UninstallLogAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UninstallLogAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAddonRequestParams struct {
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`

	// addon版本（不传默认不更新，不传AddonVersion时RawValues必传）
	AddonVersion *string `json:"AddonVersion,omitnil,omitempty" name:"AddonVersion"`

	// addon的参数，是一个json格式的base64转码后的字符串（addon参数由DescribeAddonValues获取，不传RawValues时AddonVersion必传）。
	RawValues *string `json:"RawValues,omitnil,omitempty" name:"RawValues"`

	// addon参数的更新策略，支持replace和merge两种策略，默认值为merge，兼容旧版本API。replace：使用新RawValues全量替换addon原RawValues，merge：根据新RawValues新增或更新addon原RawValues中对应参数。
	UpdateStrategy *string `json:"UpdateStrategy,omitnil,omitempty" name:"UpdateStrategy"`

	// 是否仅做更新检查，设置为true时仅做检查，不会更新组件。默认值为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type UpdateAddonRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，请从容器服务控制台集群列表中获取（https://console.cloud.tencent.com/tke2/cluster）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// addon名称，请从容器服务控制台进入集群详情页后，在组件管理页面中获取。
	AddonName *string `json:"AddonName,omitnil,omitempty" name:"AddonName"`

	// addon版本（不传默认不更新，不传AddonVersion时RawValues必传）
	AddonVersion *string `json:"AddonVersion,omitnil,omitempty" name:"AddonVersion"`

	// addon的参数，是一个json格式的base64转码后的字符串（addon参数由DescribeAddonValues获取，不传RawValues时AddonVersion必传）。
	RawValues *string `json:"RawValues,omitnil,omitempty" name:"RawValues"`

	// addon参数的更新策略，支持replace和merge两种策略，默认值为merge，兼容旧版本API。replace：使用新RawValues全量替换addon原RawValues，merge：根据新RawValues新增或更新addon原RawValues中对应参数。
	UpdateStrategy *string `json:"UpdateStrategy,omitnil,omitempty" name:"UpdateStrategy"`

	// 是否仅做更新检查，设置为true时仅做检查，不会更新组件。默认值为 false。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *UpdateAddonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAddonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "AddonName")
	delete(f, "AddonVersion")
	delete(f, "RawValues")
	delete(f, "UpdateStrategy")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAddonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAddonResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAddonResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAddonResponseParams `json:"Response"`
}

func (r *UpdateAddonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAddonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterKubeconfigRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 子账户Uin列表，传空默认为调用此接口的SubUin
	SubAccounts []*string `json:"SubAccounts,omitnil,omitempty" name:"SubAccounts"`
}

type UpdateClusterKubeconfigRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 子账户Uin列表，传空默认为调用此接口的SubUin
	SubAccounts []*string `json:"SubAccounts,omitnil,omitempty" name:"SubAccounts"`
}

func (r *UpdateClusterKubeconfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterKubeconfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SubAccounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateClusterKubeconfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterKubeconfigResponseParams struct {
	// 已更新的子账户Uin列表
	UpdatedSubAccounts []*string `json:"UpdatedSubAccounts,omitnil,omitempty" name:"UpdatedSubAccounts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateClusterKubeconfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateClusterKubeconfigResponseParams `json:"Response"`
}

func (r *UpdateClusterKubeconfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterVersionRequestParams struct {
	// 集群 Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要升级到的版本
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// 集群自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// 可容忍的最大不可用pod数目。默认0
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitnil,omitempty" name:"MaxNotReadyPercent"`

	// 是否跳过预检查阶段，默认false
	SkipPreCheck *bool `json:"SkipPreCheck,omitnil,omitempty" name:"SkipPreCheck"`
}

type UpdateClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群 Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要升级到的版本
	DstVersion *string `json:"DstVersion,omitnil,omitempty" name:"DstVersion"`

	// 集群自定义参数
	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitnil,omitempty" name:"ExtraArgs"`

	// 可容忍的最大不可用pod数目。默认0
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitnil,omitempty" name:"MaxNotReadyPercent"`

	// 是否跳过预检查阶段，默认false
	SkipPreCheck *bool `json:"SkipPreCheck,omitnil,omitempty" name:"SkipPreCheck"`
}

func (r *UpdateClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "DstVersion")
	delete(f, "ExtraArgs")
	delete(f, "MaxNotReadyPercent")
	delete(f, "SkipPreCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateClusterVersionResponseParams `json:"Response"`
}

func (r *UpdateClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEKSClusterRequestParams struct {
	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 弹性集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 弹性集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 子网Id 列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 弹性容器集群公网访问LB信息
	PublicLB *ClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 弹性容器集群内网访问LB信息
	InternalLB *ClusterInternalLB `json:"InternalLB,omitnil,omitempty" name:"InternalLB"`

	// Service 子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// 集群自定义的dns 服务器信息
	DnsServers []*DnsServerConf `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// 是否清空自定义dns 服务器设置。为1 表示 是。其他表示 否。
	ClearDnsServer *string `json:"ClearDnsServer,omitnil,omitempty" name:"ClearDnsServer"`

	// 将来删除集群时是否要删除cbs。默认为 FALSE
	NeedDeleteCbs *bool `json:"NeedDeleteCbs,omitnil,omitempty" name:"NeedDeleteCbs"`

	// 标记是否是新的内外网。默认为false
	ProxyLB *bool `json:"ProxyLB,omitnil,omitempty" name:"ProxyLB"`

	// 扩展参数。须是map[string]string 的json 格式。
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`
}

type UpdateEKSClusterRequest struct {
	*tchttp.BaseRequest
	
	// 弹性集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 弹性集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 弹性集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 子网Id 列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 弹性容器集群公网访问LB信息
	PublicLB *ClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 弹性容器集群内网访问LB信息
	InternalLB *ClusterInternalLB `json:"InternalLB,omitnil,omitempty" name:"InternalLB"`

	// Service 子网Id
	ServiceSubnetId *string `json:"ServiceSubnetId,omitnil,omitempty" name:"ServiceSubnetId"`

	// 集群自定义的dns 服务器信息
	DnsServers []*DnsServerConf `json:"DnsServers,omitnil,omitempty" name:"DnsServers"`

	// 是否清空自定义dns 服务器设置。为1 表示 是。其他表示 否。
	ClearDnsServer *string `json:"ClearDnsServer,omitnil,omitempty" name:"ClearDnsServer"`

	// 将来删除集群时是否要删除cbs。默认为 FALSE
	NeedDeleteCbs *bool `json:"NeedDeleteCbs,omitnil,omitempty" name:"NeedDeleteCbs"`

	// 标记是否是新的内外网。默认为false
	ProxyLB *bool `json:"ProxyLB,omitnil,omitempty" name:"ProxyLB"`

	// 扩展参数。须是map[string]string 的json 格式。
	ExtraParam *string `json:"ExtraParam,omitnil,omitempty" name:"ExtraParam"`
}

func (r *UpdateEKSClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEKSClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "ClusterDesc")
	delete(f, "SubnetIds")
	delete(f, "PublicLB")
	delete(f, "InternalLB")
	delete(f, "ServiceSubnetId")
	delete(f, "DnsServers")
	delete(f, "ClearDnsServer")
	delete(f, "NeedDeleteCbs")
	delete(f, "ProxyLB")
	delete(f, "ExtraParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEKSClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEKSClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEKSClusterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEKSClusterResponseParams `json:"Response"`
}

func (r *UpdateEKSClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEKSClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEKSContainerInstanceRequestParams struct {
	// 容器实例 ID
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 实例重启策略： Always(总是重启)、Never(从不重启)、OnFailure(失败时重启)
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 数据卷，包含NfsVolume数组和CbsVolume数组
	EksCiVolume *EksCiVolume `json:"EksCiVolume,omitnil,omitempty" name:"EksCiVolume"`

	// 容器组
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// Init 容器组
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 容器实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`
}

type UpdateEKSContainerInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 容器实例 ID
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 实例重启策略： Always(总是重启)、Never(从不重启)、OnFailure(失败时重启)
	RestartPolicy *string `json:"RestartPolicy,omitnil,omitempty" name:"RestartPolicy"`

	// 数据卷，包含NfsVolume数组和CbsVolume数组
	EksCiVolume *EksCiVolume `json:"EksCiVolume,omitnil,omitempty" name:"EksCiVolume"`

	// 容器组
	Containers []*Container `json:"Containers,omitnil,omitempty" name:"Containers"`

	// Init 容器组
	InitContainers []*Container `json:"InitContainers,omitnil,omitempty" name:"InitContainers"`

	// 容器实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`
}

func (r *UpdateEKSContainerInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEKSContainerInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EksCiId")
	delete(f, "RestartPolicy")
	delete(f, "EksCiVolume")
	delete(f, "Containers")
	delete(f, "InitContainers")
	delete(f, "Name")
	delete(f, "ImageRegistryCredentials")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEKSContainerInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEKSContainerInstanceResponseParams struct {
	// 容器实例 ID
	EksCiId *string `json:"EksCiId,omitnil,omitempty" name:"EksCiId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEKSContainerInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEKSContainerInstanceResponseParams `json:"Response"`
}

func (r *UpdateEKSContainerInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEKSContainerInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEdgeClusterVersionRequestParams struct {
	// 集群 Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要升级到的版本
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`

	// 自定义边缘组件镜像仓库前缀
	RegistryPrefix *string `json:"RegistryPrefix,omitnil,omitempty" name:"RegistryPrefix"`

	// 是否跳过预检查阶段
	SkipPreCheck *bool `json:"SkipPreCheck,omitnil,omitempty" name:"SkipPreCheck"`
}

type UpdateEdgeClusterVersionRequest struct {
	*tchttp.BaseRequest
	
	// 集群 Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 需要升级到的版本
	EdgeVersion *string `json:"EdgeVersion,omitnil,omitempty" name:"EdgeVersion"`

	// 自定义边缘组件镜像仓库前缀
	RegistryPrefix *string `json:"RegistryPrefix,omitnil,omitempty" name:"RegistryPrefix"`

	// 是否跳过预检查阶段
	SkipPreCheck *bool `json:"SkipPreCheck,omitnil,omitempty" name:"SkipPreCheck"`
}

func (r *UpdateEdgeClusterVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEdgeClusterVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EdgeVersion")
	delete(f, "RegistryPrefix")
	delete(f, "SkipPreCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEdgeClusterVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEdgeClusterVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEdgeClusterVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEdgeClusterVersionResponseParams `json:"Response"`
}

func (r *UpdateEdgeClusterVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEdgeClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageCacheRequestParams struct {
	// 镜像缓存ID
	ImageCacheId *string `json:"ImageCacheId,omitnil,omitempty" name:"ImageCacheId"`

	// 镜像缓存名称
	ImageCacheName *string `json:"ImageCacheName,omitnil,omitempty" name:"ImageCacheName"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 用于制作镜像缓存的容器镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// 镜像缓存的大小。默认为20 GiB。取值范围参考[云硬盘类型](https://cloud.tencent.com/document/product/362/2353)中的高性能云盘类型的大小限制。
	ImageCacheSize *uint64 `json:"ImageCacheSize,omitnil,omitempty" name:"ImageCacheSize"`

	// 镜像缓存保留时间天数，过期将会自动清理，默认为0，永不过期。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 安全组Id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type UpdateImageCacheRequest struct {
	*tchttp.BaseRequest
	
	// 镜像缓存ID
	ImageCacheId *string `json:"ImageCacheId,omitnil,omitempty" name:"ImageCacheId"`

	// 镜像缓存名称
	ImageCacheName *string `json:"ImageCacheName,omitnil,omitempty" name:"ImageCacheName"`

	// 镜像仓库凭证数组
	ImageRegistryCredentials []*ImageRegistryCredential `json:"ImageRegistryCredentials,omitnil,omitempty" name:"ImageRegistryCredentials"`

	// 用于制作镜像缓存的容器镜像列表
	Images []*string `json:"Images,omitnil,omitempty" name:"Images"`

	// 镜像缓存的大小。默认为20 GiB。取值范围参考[云硬盘类型](https://cloud.tencent.com/document/product/362/2353)中的高性能云盘类型的大小限制。
	ImageCacheSize *uint64 `json:"ImageCacheSize,omitnil,omitempty" name:"ImageCacheSize"`

	// 镜像缓存保留时间天数，过期将会自动清理，默认为0，永不过期。
	RetentionDays *uint64 `json:"RetentionDays,omitnil,omitempty" name:"RetentionDays"`

	// 安全组Id
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *UpdateImageCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageCacheId")
	delete(f, "ImageCacheName")
	delete(f, "ImageRegistryCredentials")
	delete(f, "Images")
	delete(f, "ImageCacheSize")
	delete(f, "RetentionDays")
	delete(f, "SecurityGroupIds")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateImageCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateImageCacheResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateImageCacheResponse struct {
	*tchttp.BaseResponse
	Response *UpdateImageCacheResponseParams `json:"Response"`
}

func (r *UpdateImageCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateImageCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTKEEdgeClusterRequestParams struct {
	// 边缘计算集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 边缘计算集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 边缘计算集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 边缘计算集群的pod cidr
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 边缘计算集群的service cidr
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 边缘计算集群公网访问LB信息
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 边缘计算集群内网访问LB信息
	InternalLB *EdgeClusterInternalLB `json:"InternalLB,omitnil,omitempty" name:"InternalLB"`

	// 边缘计算集群的CoreDns部署信息
	CoreDns *string `json:"CoreDns,omitnil,omitempty" name:"CoreDns"`

	// 边缘计算集群的健康检查多地域部署信息
	HealthRegion *string `json:"HealthRegion,omitnil,omitempty" name:"HealthRegion"`

	// 边缘计算集群的健康检查部署信息
	Health *string `json:"Health,omitnil,omitempty" name:"Health"`

	// 边缘计算集群的GridDaemon部署信息
	GridDaemon *string `json:"GridDaemon,omitnil,omitempty" name:"GridDaemon"`

	// 边缘集群开启自动升配
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 边缘集群的集群规模
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

type UpdateTKEEdgeClusterRequest struct {
	*tchttp.BaseRequest
	
	// 边缘计算集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 边缘计算集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 边缘计算集群描述信息
	ClusterDesc *string `json:"ClusterDesc,omitnil,omitempty" name:"ClusterDesc"`

	// 边缘计算集群的pod cidr
	PodCIDR *string `json:"PodCIDR,omitnil,omitempty" name:"PodCIDR"`

	// 边缘计算集群的service cidr
	ServiceCIDR *string `json:"ServiceCIDR,omitnil,omitempty" name:"ServiceCIDR"`

	// 边缘计算集群公网访问LB信息
	PublicLB *EdgeClusterPublicLB `json:"PublicLB,omitnil,omitempty" name:"PublicLB"`

	// 边缘计算集群内网访问LB信息
	InternalLB *EdgeClusterInternalLB `json:"InternalLB,omitnil,omitempty" name:"InternalLB"`

	// 边缘计算集群的CoreDns部署信息
	CoreDns *string `json:"CoreDns,omitnil,omitempty" name:"CoreDns"`

	// 边缘计算集群的健康检查多地域部署信息
	HealthRegion *string `json:"HealthRegion,omitnil,omitempty" name:"HealthRegion"`

	// 边缘计算集群的健康检查部署信息
	Health *string `json:"Health,omitnil,omitempty" name:"Health"`

	// 边缘计算集群的GridDaemon部署信息
	GridDaemon *string `json:"GridDaemon,omitnil,omitempty" name:"GridDaemon"`

	// 边缘集群开启自动升配
	AutoUpgradeClusterLevel *bool `json:"AutoUpgradeClusterLevel,omitnil,omitempty" name:"AutoUpgradeClusterLevel"`

	// 边缘集群的集群规模
	ClusterLevel *string `json:"ClusterLevel,omitnil,omitempty" name:"ClusterLevel"`
}

func (r *UpdateTKEEdgeClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTKEEdgeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "ClusterDesc")
	delete(f, "PodCIDR")
	delete(f, "ServiceCIDR")
	delete(f, "PublicLB")
	delete(f, "InternalLB")
	delete(f, "CoreDns")
	delete(f, "HealthRegion")
	delete(f, "Health")
	delete(f, "GridDaemon")
	delete(f, "AutoUpgradeClusterLevel")
	delete(f, "ClusterLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTKEEdgeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTKEEdgeClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTKEEdgeClusterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTKEEdgeClusterResponseParams `json:"Response"`
}

func (r *UpdateTKEEdgeClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTKEEdgeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAbleInstancesItem struct {
	// 节点Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点的当前版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 当前版本的最新小版本
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// RuntimeVersion
	RuntimeVersion *string `json:"RuntimeVersion,omitnil,omitempty" name:"RuntimeVersion"`

	// RuntimeLatestVersion
	RuntimeLatestVersion *string `json:"RuntimeLatestVersion,omitnil,omitempty" name:"RuntimeLatestVersion"`
}

// Predefined struct for user
type UpgradeClusterInstancesRequestParams struct {
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// create 表示开始一次升级任务
	// pause 表示停止任务
	// resume表示继续任务
	// abort表示终止任务
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 升级类型，只有Operation是create需要设置
	// reset 大版本重装升级
	// hot 小版本热升级
	// major 大版本原地升级
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 需要升级的节点列表，可以通过控制台或 [查询待升级节点接口
	// ](https://cloud.tencent.com/document/api/457/50366) 获取
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 当节点重新加入集群时候所使用的参数，参考添加已有节点接口
	ResetParam *UpgradeNodeResetParam `json:"ResetParam,omitnil,omitempty" name:"ResetParam"`

	// 是否忽略节点升级前检查，默认值 false
	SkipPreCheck *bool `json:"SkipPreCheck,omitnil,omitempty" name:"SkipPreCheck"`

	// 最大可容忍的不可用Pod比例
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitnil,omitempty" name:"MaxNotReadyPercent"`

	// 是否升级节点运行时，默认false不升级
	UpgradeRunTime *bool `json:"UpgradeRunTime,omitnil,omitempty" name:"UpgradeRunTime"`
}

type UpgradeClusterInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID（请登录 [TKE 控制台](https://console.cloud.tencent.com/tke2) 获取集群 ID ）
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// create 表示开始一次升级任务
	// pause 表示停止任务
	// resume表示继续任务
	// abort表示终止任务
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 升级类型，只有Operation是create需要设置
	// reset 大版本重装升级
	// hot 小版本热升级
	// major 大版本原地升级
	UpgradeType *string `json:"UpgradeType,omitnil,omitempty" name:"UpgradeType"`

	// 需要升级的节点列表，可以通过控制台或 [查询待升级节点接口
	// ](https://cloud.tencent.com/document/api/457/50366) 获取
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 当节点重新加入集群时候所使用的参数，参考添加已有节点接口
	ResetParam *UpgradeNodeResetParam `json:"ResetParam,omitnil,omitempty" name:"ResetParam"`

	// 是否忽略节点升级前检查，默认值 false
	SkipPreCheck *bool `json:"SkipPreCheck,omitnil,omitempty" name:"SkipPreCheck"`

	// 最大可容忍的不可用Pod比例
	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitnil,omitempty" name:"MaxNotReadyPercent"`

	// 是否升级节点运行时，默认false不升级
	UpgradeRunTime *bool `json:"UpgradeRunTime,omitnil,omitempty" name:"UpgradeRunTime"`
}

func (r *UpgradeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Operation")
	delete(f, "UpgradeType")
	delete(f, "InstanceIds")
	delete(f, "ResetParam")
	delete(f, "SkipPreCheck")
	delete(f, "MaxNotReadyPercent")
	delete(f, "UpgradeRunTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeClusterInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeClusterInstancesResponseParams `json:"Response"`
}

func (r *UpgradeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterReleaseRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 自定义的应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 制品名称或从第三方repo 安装chart时，制品压缩包下载地址, 不支持重定向类型chart 地址，结尾为*.tgz
	Chart *string `json:"Chart,omitnil,omitempty" name:"Chart"`

	// 自定义参数，覆盖chart 中values.yaml 中的参数
	Values *ReleaseValues `json:"Values,omitnil,omitempty" name:"Values"`

	// 制品来源，范围：tke-market 或 other 默认值：tke-market，示例值：tke-market
	ChartFrom *string `json:"ChartFrom,omitnil,omitempty" name:"ChartFrom"`

	// 制品版本( 从第三方安装时，不传这个参数）
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`

	// 制品仓库URL地址
	ChartRepoURL *string `json:"ChartRepoURL,omitnil,omitempty" name:"ChartRepoURL"`

	// 制品访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 制品访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 制品命名空间，ChartFrom为tke-market时ChartNamespace不为空，值为DescribeProducts接口反馈的Namespace
	ChartNamespace *string `json:"ChartNamespace,omitnil,omitempty" name:"ChartNamespace"`

	// 集群类型，支持传 tke, eks, tkeedge, external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

type UpgradeClusterReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 自定义的应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 制品名称或从第三方repo 安装chart时，制品压缩包下载地址, 不支持重定向类型chart 地址，结尾为*.tgz
	Chart *string `json:"Chart,omitnil,omitempty" name:"Chart"`

	// 自定义参数，覆盖chart 中values.yaml 中的参数
	Values *ReleaseValues `json:"Values,omitnil,omitempty" name:"Values"`

	// 制品来源，范围：tke-market 或 other 默认值：tke-market，示例值：tke-market
	ChartFrom *string `json:"ChartFrom,omitnil,omitempty" name:"ChartFrom"`

	// 制品版本( 从第三方安装时，不传这个参数）
	ChartVersion *string `json:"ChartVersion,omitnil,omitempty" name:"ChartVersion"`

	// 制品仓库URL地址
	ChartRepoURL *string `json:"ChartRepoURL,omitnil,omitempty" name:"ChartRepoURL"`

	// 制品访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 制品访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 制品命名空间，ChartFrom为tke-market时ChartNamespace不为空，值为DescribeProducts接口反馈的Namespace
	ChartNamespace *string `json:"ChartNamespace,omitnil,omitempty" name:"ChartNamespace"`

	// 集群类型，支持传 tke, eks, tkeedge, external(注册集群）
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`
}

func (r *UpgradeClusterReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "Chart")
	delete(f, "Values")
	delete(f, "ChartFrom")
	delete(f, "ChartVersion")
	delete(f, "ChartRepoURL")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "ChartNamespace")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeClusterReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeClusterReleaseResponseParams struct {
	// 应用详情
	Release *PendingRelease `json:"Release,omitnil,omitempty" name:"Release"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeClusterReleaseResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeClusterReleaseResponseParams `json:"Response"`
}

func (r *UpgradeClusterReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeClusterReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeNodeResetParam struct {
	// 实例额外需要设置参数信息
	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitnil,omitempty" name:"InstanceAdvancedSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type UpgradePlan struct {
	// 升级计划ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集群ID
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 预计开始时间
	PlanedStartAt *string `json:"PlanedStartAt,omitnil,omitempty" name:"PlanedStartAt"`

	// 升级开始时间
	UpgradeStartAt *string `json:"UpgradeStartAt,omitnil,omitempty" name:"UpgradeStartAt"`

	// 升级结束时间
	UpgradeEndAt *string `json:"UpgradeEndAt,omitnil,omitempty" name:"UpgradeEndAt"`

	// 升级状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type UpgradeTask struct {
	// 任务ID
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 组件名称
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 关联资源
	RelatedResources []*string `json:"RelatedResources,omitnil,omitempty" name:"RelatedResources"`

	// 升级影响
	UpgradeImpact *string `json:"UpgradeImpact,omitnil,omitempty" name:"UpgradeImpact"`

	// 预计开始时间
	PlanedStartAt *string `json:"PlanedStartAt,omitnil,omitempty" name:"PlanedStartAt"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type VersionInstance struct {
	// 版本名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 版本信息
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type VirtualNode struct {
	// 虚拟节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 虚拟节点所属子网
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 虚拟节点状态
	Phase *string `json:"Phase,omitnil,omitempty" name:"Phase"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type VirtualNodePool struct {
	// 节点池ID
	NodePoolId *string `json:"NodePoolId,omitnil,omitempty" name:"NodePoolId"`

	// 子网列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 节点池名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点池生命周期
	// - creating：创建中
	// - normal：正常
	// - updating：更新中
	LifeState *string `json:"LifeState,omitnil,omitempty" name:"LifeState"`

	// 虚拟节点label
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 虚拟节点taint
	// 注意：此字段可能返回 null，表示取不到有效值。
	Taints []*Taint `json:"Taints,omitnil,omitempty" name:"Taints"`
}

type VirtualNodeSpec struct {
	// 节点展示名称，建议不超过20个字符
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 腾讯云标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type VolumeMount struct {
	// 挂载路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountPath *string `json:"MountPath,omitnil,omitempty" name:"MountPath"`

	// volume名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 传播挂载方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountPropagation *string `json:"MountPropagation,omitnil,omitempty" name:"MountPropagation"`

	// 是否只读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 子路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubPath *string `json:"SubPath,omitnil,omitempty" name:"SubPath"`

	// 子路径表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubPathExpr *string `json:"SubPathExpr,omitnil,omitempty" name:"SubPathExpr"`
}