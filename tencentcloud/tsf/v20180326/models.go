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

package v20180326

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`
}

func (r *AddClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加云主机的返回列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *AddInstanceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInstanceResult struct {

	// 添加集群失败的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds" list`

	// 添加集群成功的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds" list`

	// 添加集群超时的节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds" list`
}

type AddInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 操作系统镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 重装系统密码设置
	Password *string `json:"Password,omitempty" name:"Password"`

	// 重装系统，关联密钥设置
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 安全组设置
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 云主机导入方式，虚拟机集群必填，容器集群不填写此字段，R：重装TSF系统镜像，M：手动安装agent
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`
}

func (r *AddInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 添加云主机是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplicationAttribute struct {

	// 总实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 运行实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 应用下部署组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupCount *int64 `json:"GroupCount,omitempty" name:"GroupCount"`
}

type ApplicationForPage struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 编程语言
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgLang *string `json:"ProgLang,omitempty" name:"ProgLang"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 应用资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`

	// Apigateway的serviceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApigatewayServiceId *string `json:"ApigatewayServiceId,omitempty" name:"ApigatewayServiceId"`
}

type Cluster struct {

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群所属私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群总CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalCpu *float64 `json:"ClusterTotalCpu,omitempty" name:"ClusterTotalCpu"`

	// 集群总内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTotalMem *float64 `json:"ClusterTotalMem,omitempty" name:"ClusterTotalMem"`

	// 集群已使用CPU，单位: 核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedCpu *float64 `json:"ClusterUsedCpu,omitempty" name:"ClusterUsedCpu"`

	// 集群已使用内存，单位: G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterUsedMem *float64 `json:"ClusterUsedMem,omitempty" name:"ClusterUsedMem"`

	// 集群机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 集群可用的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 集群正常状态的机器实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalInstanceCount *int64 `json:"NormalInstanceCount,omitempty" name:"NormalInstanceCount"`

	// 删除标记：true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群所属TSF地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfRegionName *string `json:"TsfRegionName,omitempty" name:"TsfRegionName"`

	// 集群所属TSF可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 集群所属TSF可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsfZoneName *string `json:"TsfZoneName,omitempty" name:"TsfZoneName"`

	// 集群不可删除的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlagReason *string `json:"DeleteFlagReason,omitempty" name:"DeleteFlagReason"`

	// 集群最大CPU限制，单位：核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitCpu *float64 `json:"ClusterLimitCpu,omitempty" name:"ClusterLimitCpu"`

	// 集群最大内存限制，单位：G
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterLimitMem *float64 `json:"ClusterLimitMem,omitempty" name:"ClusterLimitMem"`

	// 集群可用的服务实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunServiceInstanceCount *int64 `json:"RunServiceInstanceCount,omitempty" name:"RunServiceInstanceCount"`

	// 集群所属子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 返回给前端的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationInfo *OperationInfo `json:"OperationInfo,omitempty" name:"OperationInfo"`
}

type Config struct {

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项版本描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 删除标识，true：可以删除；false：不可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 配置项版本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersionCount *int64 `json:"ConfigVersionCount,omitempty" name:"ConfigVersionCount"`
}

type ConfigRelease struct {

	// 配置项发布ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

type ConfigReleaseLog struct {

	// 配置项发布日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseTime *string `json:"ReleaseTime,omitempty" name:"ReleaseTime"`

	// 发布描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`

	// 发布状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReleaseStatus *string `json:"ReleaseStatus,omitempty" name:"ReleaseStatus"`

	// 上次发布的配置项ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigId *string `json:"LastConfigId,omitempty" name:"LastConfigId"`

	// 上次发布的配置项名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigName *string `json:"LastConfigName,omitempty" name:"LastConfigName"`

	// 上次发布的配置项版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfigVersion *string `json:"LastConfigVersion,omitempty" name:"LastConfigVersion"`

	// 回滚标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollbackFlag *bool `json:"RollbackFlag,omitempty" name:"RollbackFlag"`
}

type ContainGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`
}

type ContainGroupResult struct {

	// 部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ContainGroup `json:"Content,omitempty" name:"Content" list`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ContainerGroupDetail struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 已启动实例总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentNum *int64 `json:"CurrentNum,omitempty" name:"CurrentNum"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 镜像server
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像名，如/tsf/nginx
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 负载均衡ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	LbIp *string `json:"LbIp,omitempty" name:"LbIp"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// Service ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIp *string `json:"ClusterIp,omitempty" name:"ClusterIp"`

	// NodePort端口，只有公网和NodePort访问方式才有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`

	// 最大分配的 CPU 核数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配的内存 MiB 数，对应 K8S limit
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 0:公网 1:集群内访问 2：NodePort
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *uint64 `json:"AccessType,omitempty" name:"AccessType"`

	// 更新方式：0:快速更新 1:滚动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 端口数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 环境变量数组对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Envs []*Env `json:"Envs,omitempty" name:"Envs" list`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// pod错误信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 初始分配的 CPU 核数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type CosCredentials struct {

	// 会话Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionToken *string `json:"SessionToken,omitempty" name:"SessionToken"`

	// 临时应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpAppId *string `json:"TmpAppId,omitempty" name:"TmpAppId"`

	// 临时调用者身份ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 所在域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type CosDownloadInfo struct {

	// 桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 鉴权信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Credentials *CosCredentials `json:"Credentials,omitempty" name:"Credentials"`
}

type CosUploadInfo struct {

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 目标地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 鉴权信息
	Credentials *CosCredentials `json:"Credentials,omitempty" name:"Credentials"`
}

type CreateApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型，V：虚拟机应用；C：容器应用；S：serverless应用
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型，M：service mesh应用；N：普通应用；G：网关应用
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用描述
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// 应用日志配置项，废弃参数
	ApplicationLogConfig *string `json:"ApplicationLogConfig,omitempty" name:"ApplicationLogConfig"`

	// 应用资源类型，废弃参数
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// 应用runtime类型
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`
}

func (r *CreateApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 分配给集群容器和服务IP的CIDR
	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`

	// 集群备注
	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`

	// 集群所属TSF地域
	TsfRegionId *string `json:"TsfRegionId,omitempty" name:"TsfRegionId"`

	// 集群所属TSF可用区
	TsfZoneId *string `json:"TsfZoneId,omitempty" name:"TsfZoneId"`

	// 私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项值类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`
}

func (r *CreateConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContainGroupRequest struct {
	*tchttp.BaseRequest

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 分组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组名称字段，长度1~60，字母或下划线开头，可包含字母数字下划线
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// 数组对象，见下方定义
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最大分配 CPU 核数，对应 K8S limit
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 最大分配内存 MiB 数，对应 K8S limit
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// 分组备注字段，长度应不大于200字符
	GroupComment *string `json:"GroupComment,omitempty" name:"GroupComment"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// 初始分配的 CPU 核数，对应 K8S request
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 初始分配的内存 MiB 数，对应 K8S request
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 部署组资源类型
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// agent 容器分配的 CPU 核数，对应 K8S 的 request
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent 容器最大的 CPU 核数，对应 K8S 的 limit
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent 容器分配的内存 MiB 数，对应 K8S 的 request
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent 容器最大的内存 MiB 数，对应 K8S 的 limit
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy 容器分配的 CPU 核数，对应 K8S 的 request
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy 容器最大的 CPU 核数，对应 K8S 的 limit
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy 容器分配的内存 MiB 数，对应 K8S 的 request
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy 容器最大的内存 MiB 数，对应 K8S 的 limit
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`
}

func (r *CreateContainGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContainGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContainGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回创建成功的部署组ID，返回null表示失败
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContainGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContainGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组所属的应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组所属命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组描述
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组资源类型
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// groupId， null表示创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务名称
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *CreateMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新增微服务是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间描述
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 命名空间资源类型(默认值为DEF)
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`

	// 是否是全局命名空间(默认是DEF，表示普通命名空间；GLOBAL表示全局命名空间)
	NamespaceType *string `json:"NamespaceType,omitempty" name:"NamespaceType"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功时为命名空间ID，失败为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项名称
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`

	// 配置项值，总是接收yaml格式的内容
	ConfigValue *string `json:"ConfigValue,omitempty" name:"ConfigValue"`

	// 配置项版本描述
	ConfigVersionDesc *string `json:"ConfigVersionDesc,omitempty" name:"ConfigVersionDesc"`

	// 配置项类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// Base64编码的配置项
	EncodeWithBase64 *bool `json:"EncodeWithBase64,omitempty" name:"EncodeWithBase64"`
}

func (r *CreatePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：创建成功；false：创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 分组名称字段，长度1~60，字母或下划线开头，可包含字母数字下划线
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 分组所属名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的部署组ID，返回null表示失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DeleteApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除应用操作是否成功。
	// true：操作成功。
	// false：操作失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：删除成功；false：删除失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除操作是否成功：
	// true：成功
	// false：失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除部署组操作是否成功。
	// true：操作成功。
	// false：操作失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTag struct {

	// 仓库名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 版本号:如V1
	TagName *string `json:"TagName,omitempty" name:"TagName"`
}

type DeleteImageTagsRequest struct {
	*tchttp.BaseRequest

	// 镜像版本数组
	ImageTags []*DeleteImageTag `json:"ImageTags,omitempty" name:"ImageTags" list`
}

func (r *DeleteImageTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteImageTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 批量删除操作是否成功。
	// true：成功。
	// false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImageTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteImageTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`
}

func (r *DeleteMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除微服务是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespaceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除命名空间是否成功。
	// true：删除成功。
	// false：删除失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePkgsRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 需要删除的程序包ID列表
	PkgIds []*string `json:"PkgIds,omitempty" name:"PkgIds" list`
}

func (r *DeletePkgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePkgsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePkgsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePkgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePkgsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeletePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：删除成功；false：删除失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// groupId，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，分组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 镜像server
	Server *string `json:"Server,omitempty" name:"Server"`

	// 镜像版本名称,如v1
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`

	// 旧版镜像名，如/tsf/nginx
	Reponame *string `json:"Reponame,omitempty" name:"Reponame"`

	// 业务容器最大的 CPU 核数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	CpuLimit *string `json:"CpuLimit,omitempty" name:"CpuLimit"`

	// 业务容器最大的内存 MiB 数，对应 K8S 的 limit；不填时默认为 request 的 2 倍
	MemLimit *string `json:"MemLimit,omitempty" name:"MemLimit"`

	// jvm参数
	JvmOpts *string `json:"JvmOpts,omitempty" name:"JvmOpts"`

	// 业务容器分配的 CPU 核数，对应 K8S 的 request
	CpuRequest *string `json:"CpuRequest,omitempty" name:"CpuRequest"`

	// 业务容器分配的内存 MiB 数，对应 K8S 的 request
	MemRequest *string `json:"MemRequest,omitempty" name:"MemRequest"`

	// 是否不立即启动
	DoNotStart *bool `json:"DoNotStart,omitempty" name:"DoNotStart"`

	// （优先使用）新版镜像名，如/tsf/nginx
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 滚动更新必填，更新间隔
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`

	// agent 容器分配的 CPU 核数，对应 K8S 的 request
	AgentCpuRequest *string `json:"AgentCpuRequest,omitempty" name:"AgentCpuRequest"`

	// agent 容器最大的 CPU 核数，对应 K8S 的 limit
	AgentCpuLimit *string `json:"AgentCpuLimit,omitempty" name:"AgentCpuLimit"`

	// agent 容器分配的内存 MiB 数，对应 K8S 的 request
	AgentMemRequest *string `json:"AgentMemRequest,omitempty" name:"AgentMemRequest"`

	// agent 容器最大的内存 MiB 数，对应 K8S 的 limit
	AgentMemLimit *string `json:"AgentMemLimit,omitempty" name:"AgentMemLimit"`

	// istioproxy 容器分配的 CPU 核数，对应 K8S 的 request
	IstioCpuRequest *string `json:"IstioCpuRequest,omitempty" name:"IstioCpuRequest"`

	// istioproxy 容器最大的 CPU 核数，对应 K8S 的 limit
	IstioCpuLimit *string `json:"IstioCpuLimit,omitempty" name:"IstioCpuLimit"`

	// istioproxy 容器分配的内存 MiB 数，对应 K8S 的 request
	IstioMemRequest *string `json:"IstioMemRequest,omitempty" name:"IstioMemRequest"`

	// istioproxy 容器最大的内存 MiB 数，对应 K8S 的 limit
	IstioMemLimit *string `json:"IstioMemLimit,omitempty" name:"IstioMemLimit"`
}

func (r *DeployContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署容器应用是否成功。
	// true：成功。
	// false：失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 部署组启动参数
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`
}

func (r *DeployGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 所需实例内存大小，取值为 1Gi 2Gi 4Gi 8Gi 16Gi，缺省为 1Gi，不传表示维持原态
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 要求最小实例数，取值范围 [1, 4]，缺省为 1，不传表示维持原态
	InstanceRequest *uint64 `json:"InstanceRequest,omitempty" name:"InstanceRequest"`

	// 部署组启动参数，不传表示维持原态
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`
}

func (r *DeployServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeployServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeployServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeployServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationAttributeRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用列表其它字段返回参数
		Result *ApplicationAttribute `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ApplicationForPage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用的微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList" list`
}

func (r *DescribeApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用分页列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageApplication `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeApplicationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群机器实例分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleaseLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的配置项发布历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfigReleaseLog `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleaseLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleasesRequest struct {
	*tchttp.BaseRequest

	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 部署组ID，不传入时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribeConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的配置发布信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfigRelease `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigReleasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *Config `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigSummaryRequest struct {
	*tchttp.BaseRequest

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字，模糊查询：应用名称，配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配置项分页对象
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest

	// 应用ID，不传入时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本，精确查询，不传入时查询全量
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页后的配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupDetailRequest struct {
	*tchttp.BaseRequest

	// 分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeContainerGroupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 容器部署组详情
		Result *ContainerGroupDetail `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerGroupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间 ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribeContainerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContainerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *ContainGroupResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContainerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContainerGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
}

func (r *DescribeDownloadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDownloadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// COS鉴权信息
		Result *CosDownloadInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDownloadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDownloadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInstancesRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署组机器信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机部署组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *VmGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组资源类型列表
	GroupResourceTypeList []*string `json:"GroupResourceTypeList,omitempty" name:"GroupResourceTypeList" list`
}

func (r *DescribeGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 虚拟机部署组分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageVmGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTagsRequest struct {
	*tchttp.BaseRequest

	// 应用Id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 不填和0:查询 1:不查询
	QueryImageIdFlag *int64 `json:"QueryImageIdFlag,omitempty" name:"QueryImageIdFlag"`

	// 可用于搜索的 tag 名字
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeImageTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeImageTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *ImageTagsResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImageTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeImageTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 微服务ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务详情实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageMsInstance `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroservicesRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 搜索字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序类型
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMicroservicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroservicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMicroservicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 微服务分页列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageMicroservice `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMicroservicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMicroservicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePkgsRequest struct {
	*tchttp.BaseRequest

	// 应用ID（只传入应用ID，返回该应用下所有软件包信息）
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 查询关键字（支持根据包ID，包名，包版本号搜索）
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 排序关键字（默认为"UploadTime"：上传时间）
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 升序：0/降序：1（默认降序）
	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`

	// 查询起始偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量限制
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePkgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePkgsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePkgsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询程序包信息列表
		Result *PkgList `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePkgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePkgsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePodInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所属groupId
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 偏移量，取值从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePodInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePodInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePodInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询的权限数据对象
		Result *GroupPodResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePodInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePodInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleaseLogsRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublicConfigReleaseLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleaseLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleaseLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页后的公共配置项发布历史列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfigReleaseLog `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigReleaseLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleaseLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleasesRequest struct {
	*tchttp.BaseRequest

	// 配置项名称，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 配置项ID，不传入时查询全量
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePublicConfigReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigReleasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 公共配置发布信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfigRelease `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigReleasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 需要查询的配置项ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 全局配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *Config `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigSummaryRequest struct {
	*tchttp.BaseRequest

	// 查询关键字，模糊查询：配置项名称，不传入时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublicConfigSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页的全局配置统计信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigsRequest struct {
	*tchttp.BaseRequest

	// 配置项ID，不传入时查询全量，高优先级
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 配置项ID列表，不传入时查询全量，低优先级
	ConfigIdList []*string `json:"ConfigIdList,omitempty" name:"ConfigIdList" list`

	// 配置项名称，精确查询，不传入时查询全量
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置项版本，精确查询，不传入时查询全量
	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribePublicConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePublicConfigsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分页后的全局配置项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageConfig `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePublicConfigsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReleasedConfigRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeReleasedConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReleasedConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReleasedConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 已发布的配置内容
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReleasedConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReleasedConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeServerlessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServerlessGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerlessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupsRequest struct {
	*tchttp.BaseRequest

	// 搜索字段，模糊搜索groupName字段
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 分组所属应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 排序字段，默认为 createTime字段，支持id， name， createTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，默认为1：倒序排序，0：正序，1：倒序
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 偏移量，取值从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分组所属名字空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 分组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeServerlessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeServerlessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据列表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *ServerlessGroupPage `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerlessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeServerlessGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleApplicationsRequest struct {
	*tchttp.BaseRequest

	// 应用ID列表
	ApplicationIdList []*string `json:"ApplicationIdList,omitempty" name:"ApplicationIdList" list`

	// 应用类型
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 微服务类型
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 资源类型数组
	ApplicationResourceTypeList []*string `json:"ApplicationResourceTypeList,omitempty" name:"ApplicationResourceTypeList" list`

	// 通过id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeSimpleApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleApplicationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 简单应用分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageSimpleApplication `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleApplicationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleClustersRequest struct {
	*tchttp.BaseRequest

	// 需要查询的集群ID列表，不填或不传入时查询所有内容
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList" list`

	// 需要查询的集群类型，不填或不传入时查询所有内容
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页个数，默认为20， 取值应为1~50
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 对id和name进行关键词过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeSimpleClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleClustersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TSF集群分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageCluster `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleClustersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleGroupsRequest struct {
	*tchttp.BaseRequest

	// 部署组ID列表，不填写时查询全量
	GroupIdList []*string `json:"GroupIdList,omitempty" name:"GroupIdList" list`

	// 应用ID，不填写时查询全量
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 集群ID，不填写时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间ID，不填写时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 部署组ID，不填写时查询全量
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 模糊查询，部署组名称，不填写时查询全量
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 部署组类型，精确过滤字段，M：service mesh, P：原生应用， M：网关应用
	AppMicroServiceType *string `json:"AppMicroServiceType,omitempty" name:"AppMicroServiceType"`
}

func (r *DescribeSimpleGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 简单部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageSimpleGroup `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleNamespacesRequest struct {
	*tchttp.BaseRequest

	// 命名空间ID列表，不传入时查询全量
	NamespaceIdList []*string `json:"NamespaceIdList,omitempty" name:"NamespaceIdList" list`

	// 集群ID，不传入时查询全量
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 每页条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 命名空间ID，不传入时查询全量
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 查询资源类型列表
	NamespaceResourceTypeList []*string `json:"NamespaceResourceTypeList,omitempty" name:"NamespaceResourceTypeList" list`

	// 通过id和name进行过滤
	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`

	// 查询的命名空间类型列表
	NamespaceTypeList []*string `json:"NamespaceTypeList,omitempty" name:"NamespaceTypeList" list`

	// 通过命名空间名精确过滤
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 通过是否是默认命名空间过滤，不传表示拉取全部命名空间。0：默认，命名空间。1：非默认命名空间
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`
}

func (r *DescribeSimpleNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleNamespacesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSimpleNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间分页列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TsfPageNamespace `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSimpleNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSimpleNamespacesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 程序包名
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包版本
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包类型
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包介绍
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`
}

func (r *DescribeUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// COS上传信息
		Result *CosUploadInfo `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUploadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Env struct {

	// 环境变量名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务端口
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ExpandGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 扩容的机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *ExpandGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExpandGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExpandGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExpandGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExpandGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupPod struct {

	// 实例名称(对应到kubernetes的pod名称)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 实例ID(对应到kubernetes的pod id)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodId *string `json:"PodId,omitempty" name:"PodId"`

	// 实例状态，请参考后面的实例以及容器的状态定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实例处于当前状态的原因，例如容器下载镜像失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 主机IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 实例中容器的重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`

	// 实例中已就绪容器的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadyCount *int64 `json:"ReadyCount,omitempty" name:"ReadyCount"`

	// 运行时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 实例启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 服务实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 机器实例可使用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 机器实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
}

type GroupPodResult struct {

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*GroupPod `json:"Content,omitempty" name:"Content" list`
}

type ImageTag struct {

	// 仓库名
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 版本名称
	TagName *string `json:"TagName,omitempty" name:"TagName"`

	// 版本ID
	TagId *string `json:"TagId,omitempty" name:"TagId"`

	// 镜像ID
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 镜像制作者
	Author *string `json:"Author,omitempty" name:"Author"`

	// CPU架构
	Architecture *string `json:"Architecture,omitempty" name:"Architecture"`

	// Docker客户端版本
	DockerVersion *string `json:"DockerVersion,omitempty" name:"DockerVersion"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	Os *string `json:"Os,omitempty" name:"Os"`

	// push时间
	PushTime *string `json:"PushTime,omitempty" name:"PushTime"`

	// 单位为字节
	SizeByte *int64 `json:"SizeByte,omitempty" name:"SizeByte"`
}

type ImageTagsResult struct {

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 仓库名,含命名空间,如tsf/ngin
	RepoName *string `json:"RepoName,omitempty" name:"RepoName"`

	// 镜像服务器地址
	Server *string `json:"Server,omitempty" name:"Server"`

	// 列表信息
	Content []*ImageTag `json:"Content,omitempty" name:"Content" list`
}

type Instance struct {

	// 机器实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 机器内网地址IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器外网地址IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// VM的状态 虚机：虚机的状态 容器：Pod所在虚机的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// VM的可使用状态 虚机：虚机是否能够作为资源使用 容器：虚机是否能够作为资源部署POD
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 服务下的服务实例的状态 虚机：应用是否可用 + Agent状态 容器：Pod状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 标识此instance是否已添加在tsf中
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountInTsf *int64 `json:"CountInTsf,omitempty" name:"CountInTsf"`

	// 机器所属部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器所属应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 机器所属应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 机器实例在CVM的创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCreatedTime *string `json:"InstanceCreatedTime,omitempty" name:"InstanceCreatedTime"`

	// 机器实例在CVM的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceExpiredTime *string `json:"InstanceExpiredTime,omitempty" name:"InstanceExpiredTime"`

	// 机器实例在CVM的计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 机器实例总CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalCpu *float64 `json:"InstanceTotalCpu,omitempty" name:"InstanceTotalCpu"`

	// 机器实例总内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalMem *float64 `json:"InstanceTotalMem,omitempty" name:"InstanceTotalMem"`

	// 机器实例使用的CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUsedCpu *float64 `json:"InstanceUsedCpu,omitempty" name:"InstanceUsedCpu"`

	// 机器实例使用的内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceUsedMem *float64 `json:"InstanceUsedMem,omitempty" name:"InstanceUsedMem"`

	// 机器实例Limit CPU信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLimitCpu *float64 `json:"InstanceLimitCpu,omitempty" name:"InstanceLimitCpu"`

	// 机器实例Limit 内存信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLimitMem *float64 `json:"InstanceLimitMem,omitempty" name:"InstanceLimitMem"`

	// 包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePkgVersion *string `json:"InstancePkgVersion,omitempty" name:"InstancePkgVersion"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 机器实例业务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestrictState *string `json:"RestrictState,omitempty" name:"RestrictState"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例执行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationState *int64 `json:"OperationState,omitempty" name:"OperationState"`

	// NamespaceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// InstanceZoneId
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceZoneId *string `json:"InstanceZoneId,omitempty" name:"InstanceZoneId"`

	// InstanceImportMode
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceImportMode *string `json:"InstanceImportMode,omitempty" name:"InstanceImportMode"`

	// ApplicationType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// ApplicationResourceType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// ServiceSidecarStatus
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceSidecarStatus *string `json:"ServiceSidecarStatus,omitempty" name:"ServiceSidecarStatus"`

	// GroupName
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// NamespaceName
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type Microservice struct {

	// 微服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceName *string `json:"MicroserviceName,omitempty" name:"MicroserviceName"`

	// 微服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 微服务的运行实例数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`
}

type ModifyContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 0:公网 1:集群内访问 2：NodePort
	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`

	// ProtocolPorts数组
	ProtocolPorts []*ProtocolPort `json:"ProtocolPorts,omitempty" name:"ProtocolPorts" list`

	// 更新方式：0:快速更新 1:滚动更新
	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`

	// 更新间隔,单位秒
	UpdateIvl *int64 `json:"UpdateIvl,omitempty" name:"UpdateIvl"`
}

func (r *ModifyContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新部署组是否成功。
	// true：成功。
	// false：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerReplicasRequest struct {
	*tchttp.BaseRequest

	// 部署组ID，部署组唯一标识
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 实例数量
	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
}

func (r *ModifyContainerReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerReplicasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContainerReplicasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果true：成功；false：失败；
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContainerReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContainerReplicasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroserviceRequest struct {
	*tchttp.BaseRequest

	// 微服务 ID
	MicroserviceId *string `json:"MicroserviceId,omitempty" name:"MicroserviceId"`

	// 微服务备注信息
	MicroserviceDesc *string `json:"MicroserviceDesc,omitempty" name:"MicroserviceDesc"`
}

func (r *ModifyMicroserviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMicroserviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMicroserviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改微服务详情是否成功。
	// true：操作成功。
	// false：操作失败。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMicroserviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMicroserviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUploadInfoRequest struct {
	*tchttp.BaseRequest

	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 调用DescribeUploadInfo接口时返回的软件包ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// COS返回上传结果（默认为0：成功，其他值表示失败）
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 程序包MD5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包大小（单位字节）
	Size *uint64 `json:"Size,omitempty" name:"Size"`
}

func (r *ModifyUploadInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUploadInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyUploadInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUploadInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyUploadInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MsInstance struct {

	// 机器实例ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 机器实例名称信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 服务运行的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 机器实例内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`

	// 机器实例外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 机器可用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceAvailableStatus *string `json:"InstanceAvailableStatus,omitempty" name:"InstanceAvailableStatus"`

	// 服务运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInstanceStatus *string `json:"ServiceInstanceStatus,omitempty" name:"ServiceInstanceStatus"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器TSF可用状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 健康检查URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" name:"HealthCheckUrl"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 应用程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationPackageVersion *string `json:"ApplicationPackageVersion,omitempty" name:"ApplicationPackageVersion"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`
}

type Namespace struct {

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceCode *string `json:"NamespaceCode,omitempty" name:"NamespaceCode"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 命名空间描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceDesc *string `json:"NamespaceDesc,omitempty" name:"NamespaceDesc"`

	// 默认命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *string `json:"IsDefault,omitempty" name:"IsDefault"`

	// 命名空间状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" name:"NamespaceStatus"`

	// 删除标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *bool `json:"DeleteFlag,omitempty" name:"DeleteFlag"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群数组，仅携带集群ID，集群名称，集群类型等基础信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*Cluster `json:"ClusterList,omitempty" name:"ClusterList" list`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceResourceType *string `json:"NamespaceResourceType,omitempty" name:"NamespaceResourceType"`

	// 命名空间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceType *string `json:"NamespaceType,omitempty" name:"NamespaceType"`
}

type OperationInfo struct {

	// 初始化按钮的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Init *OperationInfoDetail `json:"Init,omitempty" name:"Init"`

	// 添加实例按钮的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddInstance *OperationInfoDetail `json:"AddInstance,omitempty" name:"AddInstance"`

	// 销毁机器的控制信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destroy *OperationInfoDetail `json:"Destroy,omitempty" name:"Destroy"`
}

type OperationInfoDetail struct {

	// 不显示的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisabledReason *string `json:"DisabledReason,omitempty" name:"DisabledReason"`

	// 该按钮是否可点击
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 是否显示该按钮
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supported *bool `json:"Supported,omitempty" name:"Supported"`
}

type PkgInfo struct {

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 程序包类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgType *string `json:"PkgType,omitempty" name:"PkgType"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 程序包描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgDesc *string `json:"PkgDesc,omitempty" name:"PkgDesc"`

	// 上传时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadTime *string `json:"UploadTime,omitempty" name:"UploadTime"`

	// 程序包MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 程序包状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgPubStatus *int64 `json:"PkgPubStatus,omitempty" name:"PkgPubStatus"`
}

type PkgList struct {

	// 程序包总量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 程序包信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*PkgInfo `json:"Content,omitempty" name:"Content" list`
}

type ProtocolPort struct {

	// TCP UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 服务端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 容器端口
	TargetPort *int64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 主机端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePort *int64 `json:"NodePort,omitempty" name:"NodePort"`
}

type ReleaseConfigRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleaseConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：发布成功；false：发布失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleasePublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置ID
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 发布描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *ReleasePublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleasePublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleasePublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：发布成功；false：发布失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleasePublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleasePublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 云主机 ID 列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
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

		// 集群移除机器是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

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

type RevocationConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevocationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：回滚成功；false：回滚失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevocationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationPublicConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布ID
	ConfigReleaseId *string `json:"ConfigReleaseId,omitempty" name:"ConfigReleaseId"`
}

func (r *RevocationPublicConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationPublicConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RevocationPublicConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：撤销成功；false：撤销失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RevocationPublicConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RevocationPublicConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackConfigRequest struct {
	*tchttp.BaseRequest

	// 配置项发布历史ID
	ConfigReleaseLogId *string `json:"ConfigReleaseLogId,omitempty" name:"ConfigReleaseLogId"`

	// 回滚描述
	ReleaseDesc *string `json:"ReleaseDesc,omitempty" name:"ReleaseDesc"`
}

func (r *RollbackConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RollbackConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true：回滚成功；false：回滚失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RollbackConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RollbackConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ServerlessGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 服务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 程序包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgName *string `json:"PkgName,omitempty" name:"PkgName"`

	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// vpc ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// vpc 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 程序包版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgVersion *string `json:"PkgVersion,omitempty" name:"PkgVersion"`

	// 所需实例内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *string `json:"Memory,omitempty" name:"Memory"`

	// 要求最小实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceRequest *uint64 `json:"InstanceRequest,omitempty" name:"InstanceRequest"`

	// 部署组启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 部署组实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *uint64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName []*string `json:"ApplicationName,omitempty" name:"ApplicationName" list`
}

type ServerlessGroupPage struct {

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ServerlessGroup `json:"Content,omitempty" name:"Content" list`
}

type ShrinkGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ShrinkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShrinkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkInstancesRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 下线机器实例ID列表
	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList" list`
}

func (r *ShrinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShrinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShrinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ShrinkInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimpleApplication struct {

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// ApplicationDesc
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationDesc *string `json:"ApplicationDesc,omitempty" name:"ApplicationDesc"`

	// ProgLang
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgLang *string `json:"ProgLang,omitempty" name:"ProgLang"`

	// ApplicationResourceType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationResourceType *string `json:"ApplicationResourceType,omitempty" name:"ApplicationResourceType"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// UpdateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// ApigatewayServiceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApigatewayServiceId *string `json:"ApigatewayServiceId,omitempty" name:"ApigatewayServiceId"`

	// ApplicationRuntimeType
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationRuntimeType *string `json:"ApplicationRuntimeType,omitempty" name:"ApplicationRuntimeType"`
}

type SimpleGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppMicroServiceType *string `json:"AppMicroServiceType,omitempty" name:"AppMicroServiceType"`
}

type StartContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 启动操作是否成功。
	// true：启动成功
	// false：启动失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StartGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopContainerGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopContainerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopContainerGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopContainerGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 停止操作是否成功。
	// true：停止成功
	// flase：停止失败
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopContainerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopContainerGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGroupRequest struct {
	*tchttp.BaseRequest

	// 部署组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *StopGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *TaskId `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TaskId struct {

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type TsfPageApplication struct {

	// 应用总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 应用信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ApplicationForPage `json:"Content,omitempty" name:"Content" list`
}

type TsfPageCluster struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Cluster `json:"Content,omitempty" name:"Content" list`
}

type TsfPageConfig struct {

	// TsfPageConfig
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项列表
	Content []*Config `json:"Content,omitempty" name:"Content" list`
}

type TsfPageConfigRelease struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项发布信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigRelease `json:"Content,omitempty" name:"Content" list`
}

type TsfPageConfigReleaseLog struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 配置项发布日志数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*ConfigReleaseLog `json:"Content,omitempty" name:"Content" list`
}

type TsfPageInstance struct {

	// 机器实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 机器实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Instance `json:"Content,omitempty" name:"Content" list`
}

type TsfPageMicroservice struct {

	// 微服务总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Microservice `json:"Content,omitempty" name:"Content" list`
}

type TsfPageMsInstance struct {

	// 微服务实例总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 微服务实例列表内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*MsInstance `json:"Content,omitempty" name:"Content" list`
}

type TsfPageNamespace struct {

	// 命名空间总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 命名空间列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*Namespace `json:"Content,omitempty" name:"Content" list`
}

type TsfPageSimpleApplication struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 简单应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*SimpleApplication `json:"Content,omitempty" name:"Content" list`
}

type TsfPageSimpleGroup struct {

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 简单部署组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*SimpleGroup `json:"Content,omitempty" name:"Content" list`
}

type TsfPageVmGroup struct {

	// 虚拟机部署组总数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 虚拟机部署组列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*VmGroupSimple `json:"Content,omitempty" name:"Content" list`
}

type VmGroup struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 部署组状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupStatus *string `json:"GroupStatus,omitempty" name:"GroupStatus"`

	// 程序包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 程序包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 程序包版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageVersion *string `json:"PackageVersion,omitempty" name:"PackageVersion"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 部署组机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 部署组运行中机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunInstanceCount *int64 `json:"RunInstanceCount,omitempty" name:"RunInstanceCount"`

	// 部署组启动参数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 部署组停止机器数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffInstanceCount *int64 `json:"OffInstanceCount,omitempty" name:"OffInstanceCount"`

	// 部署组描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}

type VmGroupSimple struct {

	// 部署组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 部署组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 部署组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDesc *string `json:"GroupDesc,omitempty" name:"GroupDesc"`

	// 部署组更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 部署组启动参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupParameters *string `json:"StartupParameters,omitempty" name:"StartupParameters"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 部署组创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationName *string `json:"ApplicationName,omitempty" name:"ApplicationName"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 应用微服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MicroserviceType *string `json:"MicroserviceType,omitempty" name:"MicroserviceType"`

	// 部署组资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupResourceType *string `json:"GroupResourceType,omitempty" name:"GroupResourceType"`

	// 部署组更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" name:"UpdatedTime"`
}
