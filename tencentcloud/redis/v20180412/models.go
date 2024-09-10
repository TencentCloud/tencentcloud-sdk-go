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

package v20180412

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Account struct {
	// 实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账号名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 账号描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 读写权限策略。
	// - r：只读。
	// - w：只写。
	// - rw：读写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 只读路由策略。
	// - master：主节点。
	// - replication：从节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// 子账号状态.
	// - 1：账号变更中。
	// - 2：账号有效。
	// - 4：账号已删除。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type AddReplicationInstanceRequestParams struct {
	// 复制组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 给复制组添加的实例分配角色。<ul><li>rw：可读写。</li><li>r：只读。</li></ul>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

type AddReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 复制组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 给复制组添加的实例分配角色。<ul><li>rw：可读写。</li><li>r：只读。</li></ul>
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

func (r *AddReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "InstanceRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddReplicationInstanceResponseParams struct {
	// 异步流程ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *AddReplicationInstanceResponseParams `json:"Response"`
}

func (r *AddReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateWanAddressRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type AllocateWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *AllocateWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateWanAddressResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 开通外网的状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AllocateWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *AllocateWanAddressResponseParams `json:"Response"`
}

func (r *AllocateWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyParamsTemplateRequestParams struct {
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 应用的参数模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type ApplyParamsTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 应用的参数模板ID
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *ApplyParamsTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyParamsTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyParamsTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyParamsTemplateResponseParams struct {
	// 任务ID
	TaskIds []*int64 `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyParamsTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ApplyParamsTemplateResponseParams `json:"Response"`
}

func (r *ApplyParamsTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyParamsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 要绑定的安全组ID，类似sg-efil73jd。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 被绑定的实例ID，类似ins-lesecurk，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 要绑定的安全组ID，类似sg-efil73jd。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 被绑定的实例ID，类似ins-lesecurk，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableRegion struct {
	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableZones []*string `json:"AvailableZones,omitnil,omitempty" name:"AvailableZones"`
}

type BackupDownloadInfo struct {
	// 备份文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 备份文件大小，单位B，如果为0，表示无效。
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 备份文件外网下载地址。下载地址的有效时长为6小时，过期后请重新获取。
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 备份文件内网下载地址。下载地址的有效时长为6小时，过期后请重新获取。
	InnerDownloadUrl *string `json:"InnerDownloadUrl,omitnil,omitempty" name:"InnerDownloadUrl"`
}

type BackupLimitVpcItem struct {
	// 备份文件的下载地址对应VPC 所属的地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 备份文件下载地址的 VPC 列表。
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`
}

type BigKeyInfo struct {
	// 所属的database
	DB *int64 `json:"DB,omitnil,omitempty" name:"DB"`

	// 大Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 大小
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 数据时间戳
	Updatetime *int64 `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`
}

type BigKeyTypeInfo struct {
	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 大小
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 时间戳
	Updatetime *int64 `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`
}

type CDCResource struct {
	// 用户的Appid
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域id
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// redis独享集群id
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// 计费模式，1-包年包月，0-按量计费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 项目id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自动续费标识，0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 独享集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 实例创建时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例到期时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 集群状态：1-流程中，2-运行中，3-已隔离
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 基础管控资源包
	BaseBundles []*ResourceBundle `json:"BaseBundles,omitnil,omitempty" name:"BaseBundles"`

	// 资源包列表
	ResourceBundles []*ResourceBundle `json:"ResourceBundles,omitnil,omitempty" name:"ResourceBundles"`

	// 所属本地专有集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

// Predefined struct for user
type ChangeInstanceRoleRequestParams struct {
	// 复制组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例角色，rw可读写，r只读
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

type ChangeInstanceRoleRequest struct {
	*tchttp.BaseRequest
	
	// 复制组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例角色，rw可读写，r只读
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`
}

func (r *ChangeInstanceRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "InstanceRole")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeInstanceRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeInstanceRoleResponseParams struct {
	// 异步流程ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeInstanceRoleResponse struct {
	*tchttp.BaseResponse
	Response *ChangeInstanceRoleResponseParams `json:"Response"`
}

func (r *ChangeInstanceRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceRequestParams struct {
	// 复制组ID。创建复制组时，系统自动分配的 ID，是复制组的唯一标识。例如：crs-rpl-m3zt****，请登录[Redis 控制台](https://console.cloud.tencent.com/redis/replication)的全球复制组列表获取复制组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 指定待提升为主实例的只读实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 标识是否强制提主。
	// - true：强制提主。
	// - false：不强制提主。
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`
}

type ChangeMasterInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 复制组ID。创建复制组时，系统自动分配的 ID，是复制组的唯一标识。例如：crs-rpl-m3zt****，请登录[Redis 控制台](https://console.cloud.tencent.com/redis/replication)的全球复制组列表获取复制组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 指定待提升为主实例的只读实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 标识是否强制提主。
	// - true：强制提主。
	// - false：不强制提主。
	ForceSwitch *bool `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`
}

func (r *ChangeMasterInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMasterInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "ForceSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeMasterInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMasterInstanceResponseParams struct {
	// 异步流程ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeMasterInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ChangeMasterInstanceResponseParams `json:"Response"`
}

func (r *ChangeMasterInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMasterInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 副本节点组 ID，请通过接口[DescribeInstanceZoneInfo](https://cloud.tencent.com/document/product/239/50312)获取多 AZ备节点组的 ID 信息。单 AZ，则无需配置该参数。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type ChangeReplicaToMasterRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 副本节点组 ID，请通过接口[DescribeInstanceZoneInfo](https://cloud.tencent.com/document/product/239/50312)获取多 AZ备节点组的 ID 信息。单 AZ，则无需配置该参数。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *ChangeReplicaToMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeReplicaToMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeReplicaToMasterResponseParams struct {
	// 异步任务ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeReplicaToMasterResponse struct {
	*tchttp.BaseResponse
	Response *ChangeReplicaToMasterResponseParams `json:"Response"`
}

func (r *ChangeReplicaToMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeReplicaToMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CleanUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CleanUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CleanUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CleanUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CleanUpInstanceResponseParams `json:"Response"`
}

func (r *CleanUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CleanUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// redis的实例密码（免密实例不需要传密码，非免密实例必传）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// redis的实例密码（免密实例不需要传密码，非免密实例必传）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ClearInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearInstanceResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ClearInstanceResponseParams `json:"Response"`
}

func (r *ClearInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneInstancesRequestParams struct {
	// 指定待克隆的源实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 单次克隆实例的数量。
	// - 包年包月每次购买最大数量为100。
	// - 按量计费每次购买最大数量为30。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 克隆实例所属的可用区ID。当前所支持的可用区 ID，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106) 。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 付费方式。<ul><li>0：按量计费。</li><li>1：包年包月。</li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 购买实例时长。<ul><li>单位：月。</li><li>付费方式选择包年包月计费时，取值范围为[1,2,3,4,5,6,7,8,9,10,11,12,24,36,48,60]。</li><li>付费方式选择按量计费时，设置为1。</li></ul>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 安全组ID。请登录控制台，在<b>安全组</b>页面获取安全组 ID 信息。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// 克隆实例使用的备份ID。请通过接口[DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011)获取备份ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 配置克隆实例是否支持免密访问。开启 SSL 与外网均不支持免密访问。<ul><li>true：免密实例，</li><li>false：非免密实例。默认为非免密实例。</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// 配置克隆实例的私有网络ID。如果未配置该参数，默认选择基础网络。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 配置克隆实例所属私有网络的子网。基础网络时该参数无需配置。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 克隆实例的名称。<br>仅支持长度小于60的中文、英文或者数字，短划线"-"、下划线"_"。</br>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 克隆实例的访问密码。<ul><li>当输入参数<b>NoAuth</b>为<b>true</b>时，可不设置该参数。</li><li>当实例为Redis2.8、4.0和5.0时，其密码格式为：8-30个字符，至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种，不能以"/"开头；</li><li>当实例为CKV 3.2时，其密码格式为：8-30个字符，必须包含字母和数字，且不包含其他字符。</li></ul>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 自动续费标识。<ul><li>0：默认状态，手动续费。</li><li>1：自动续费。</li><li>2：不自动续费，到期自动隔离。</li></ul>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 用户自定义的端口，默认为6379，取值范围[1024,65535]。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 实例的节点信息。<ul><li>目前支持配置节点的类型（主节点或者副本节点），及其节点的可用区信息。具体信息，请参见[RedisNodeInfo](https://cloud.tencent.com/document/product/239/20022#RedisNodeInfo)。</li><li>单可用区部署可不配置该参数。</li></ul>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 项目 ID。登录[Redis 控制台](https://console.cloud.tencent.com/redis#/)，可在右上角的<b>账号中心</b> > <b>项目管理</b>中查找项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 克隆实例需绑定的标签。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 指定克隆实例相关的参数模板 ID。
	// - 若不配置该参数，则系统会依据所选择的兼容版本及架构，自动适配对应的默认模板。
	// - 请通过[DescribeParamTemplates](https://cloud.tencent.com/document/product/239/58750)接口，查询实例的参数模板列表，获取模板 ID 编号。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 指定克隆实例的告警策略 ID。请登录[腾讯云可观测平台控制台](https://console.cloud.tencent.com/monitor/alarm2/policy)，在 <b>告警管理</b> > <b>策略管理</b>页面获取策略 ID 信息。
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 克隆指定恢复数据的时间。
	// 仅支持已开通秒级备份的实例
	CloneTime *string `json:"CloneTime,omitnil,omitempty" name:"CloneTime"`
}

type CloneInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 指定待克隆的源实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 单次克隆实例的数量。
	// - 包年包月每次购买最大数量为100。
	// - 按量计费每次购买最大数量为30。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 克隆实例所属的可用区ID。当前所支持的可用区 ID，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106) 。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 付费方式。<ul><li>0：按量计费。</li><li>1：包年包月。</li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 购买实例时长。<ul><li>单位：月。</li><li>付费方式选择包年包月计费时，取值范围为[1,2,3,4,5,6,7,8,9,10,11,12,24,36,48,60]。</li><li>付费方式选择按量计费时，设置为1。</li></ul>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 安全组ID。请登录控制台，在<b>安全组</b>页面获取安全组 ID 信息。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// 克隆实例使用的备份ID。请通过接口[DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011)获取备份ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 配置克隆实例是否支持免密访问。开启 SSL 与外网均不支持免密访问。<ul><li>true：免密实例，</li><li>false：非免密实例。默认为非免密实例。</li></ul>
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// 配置克隆实例的私有网络ID。如果未配置该参数，默认选择基础网络。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 配置克隆实例所属私有网络的子网。基础网络时该参数无需配置。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 克隆实例的名称。<br>仅支持长度小于60的中文、英文或者数字，短划线"-"、下划线"_"。</br>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 克隆实例的访问密码。<ul><li>当输入参数<b>NoAuth</b>为<b>true</b>时，可不设置该参数。</li><li>当实例为Redis2.8、4.0和5.0时，其密码格式为：8-30个字符，至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种，不能以"/"开头；</li><li>当实例为CKV 3.2时，其密码格式为：8-30个字符，必须包含字母和数字，且不包含其他字符。</li></ul>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 自动续费标识。<ul><li>0：默认状态，手动续费。</li><li>1：自动续费。</li><li>2：不自动续费，到期自动隔离。</li></ul>
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 用户自定义的端口，默认为6379，取值范围[1024,65535]。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 实例的节点信息。<ul><li>目前支持配置节点的类型（主节点或者副本节点），及其节点的可用区信息。具体信息，请参见[RedisNodeInfo](https://cloud.tencent.com/document/product/239/20022#RedisNodeInfo)。</li><li>单可用区部署可不配置该参数。</li></ul>
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 项目 ID。登录[Redis 控制台](https://console.cloud.tencent.com/redis#/)，可在右上角的<b>账号中心</b> > <b>项目管理</b>中查找项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 克隆实例需绑定的标签。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 指定克隆实例相关的参数模板 ID。
	// - 若不配置该参数，则系统会依据所选择的兼容版本及架构，自动适配对应的默认模板。
	// - 请通过[DescribeParamTemplates](https://cloud.tencent.com/document/product/239/58750)接口，查询实例的参数模板列表，获取模板 ID 编号。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 指定克隆实例的告警策略 ID。请登录[腾讯云可观测平台控制台](https://console.cloud.tencent.com/monitor/alarm2/policy)，在 <b>告警管理</b> > <b>策略管理</b>页面获取策略 ID 信息。
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 克隆指定恢复数据的时间。
	// 仅支持已开通秒级备份的实例
	CloneTime *string `json:"CloneTime,omitnil,omitempty" name:"CloneTime"`
}

func (r *CloneInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GoodsNum")
	delete(f, "ZoneId")
	delete(f, "BillingMode")
	delete(f, "Period")
	delete(f, "SecurityGroupIdList")
	delete(f, "BackupId")
	delete(f, "NoAuth")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "InstanceName")
	delete(f, "Password")
	delete(f, "AutoRenew")
	delete(f, "VPort")
	delete(f, "NodeSet")
	delete(f, "ProjectId")
	delete(f, "ResourceTags")
	delete(f, "TemplateId")
	delete(f, "AlarmPolicyList")
	delete(f, "CloneTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneInstancesResponseParams struct {
	// 请求任务 ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 克隆实例的 ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CloneInstancesResponseParams `json:"Response"`
}

func (r *CloneInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseSSLRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseSSLResponseParams struct {
	// 任务ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseSSLResponse struct {
	*tchttp.BaseResponse
	Response *CloseSSLResponseParams `json:"Response"`
}

func (r *CloseSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommandTake struct {
	// 命令名。
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 耗时时长。单位：ms。
	Took *int64 `json:"Took,omitnil,omitempty" name:"Took"`
}

// Predefined struct for user
type CreateInstanceAccountRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义访问数据库的名称。
	// - 仅由字母、数字、下划线、中划线组成。
	// - 长度不能大于32位。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 设置自定义账号的密码。密码复杂度要求如下：
	// - 字符个数为[8,32]。
	// - 至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的两种。
	// - 不能以"/"开头。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 指定账号的读请求路由分发至主节点或副本节点。未开启副本只读，不支持选择副本节点。
	// - master：主节点
	// - replication：副本节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// 账户读写权限，支持选择只读与读写权限。
	// - r：只读
	// - rw: 读写权限
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 子账号描述信息，长度[0,64] 字节，支持中文。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 自定义访问数据库的名称。
	// - 仅由字母、数字、下划线、中划线组成。
	// - 长度不能大于32位。
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 设置自定义账号的密码。密码复杂度要求如下：
	// - 字符个数为[8,32]。
	// - 至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的两种。
	// - 不能以"/"开头。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 指定账号的读请求路由分发至主节点或副本节点。未开启副本只读，不支持选择副本节点。
	// - master：主节点
	// - replication：副本节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// 账户读写权限，支持选择只读与读写权限。
	// - r：只读
	// - rw: 读写权限
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 子账号描述信息，长度[0,64] 字节，支持中文。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "ReadonlyPolicy")
	delete(f, "Privilege")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceAccountResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceAccountResponseParams `json:"Response"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// 实例类型。
	// <ul><li>2：Redis 2.8 内存版（标准架构）。</li><li>3：CKV 3.2 内存版（标准架构）。</li><li>4：CKV 3.2 内存版（集群架构）。</li><li>6：Redis 4.0 内存版（标准架构）。</li><li>7：Redis 4.0 内存版（集群架构）。</li><li>8：Redis 5.0 内存版（标准架构）。</li><li>9：Redis 5.0 内存版（集群架构）。</li><li>15：Redis 6.2 内存版（标准架构）。</li><li>16：Redis 6.2 内存版（集群架构）。</li><li>17：Redis 7.0 内存版（标准架构）。</li><li>18：Redis 7.0 内存版（集群架构）。</li></ul>
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 内存容量，单位为MB， 数值需为1024的整数倍。具体规格，请通过 [DescribeProductInfo](https://cloud.tencent.com/document/api/239/30600) 接口查询全地域的售卖规格。
	// - **TypeId**为标准架构时，**MemSize**是实例总内存容量；
	// - **TypeId**为集群架构时，**MemSize**是单分片内存容量。
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 实例数量，单次购买实例数量。具体信息，请通过 [DescribeProductInfo](https://cloud.tencent.com/document/api/239/30600) 接口查询全地域的售卖规格。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 购买实例的时长。
	// - 若 **BillingMode**为**1**，即计费方式为包年包月时，需设置该参数，指定所购买实例的时长。单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	// - 若 **BillingMode**为**0**，即计费方式为按量计费时，该参数配置为1。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 计费方式。
	// - 0：按量计费。
	// - 1：包年包月。
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例所属的可用区ID，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 访问实例的密码。
	// - 当输入参数**NoAuth**为**true**时，指设置实例为免密码访问，Password可不用配置，否则Password为必填参数。
	// - 当实例类型**TypeId**为Redis 2.8 内存版标准架构、Redis 4.0、5.0、6.0内存版标准架构或集群架构时，其密码复杂度要求为：8-30个字符，至少包含小写字母、大写字母、数字和字符()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种，不能以"/"开头。
	// - 当实例类型**TypeId**为CKV 3.2 内存版标准架构或集群架构时，其密码复杂度为：8-30个字符，必须包含字母和数字，且 不包含其他字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 私有网络ID。如果不配置该参数则默认选择基础网络。请登录 [私有网络](https://console.cloud.tencent.com/vpc)控制台查询具体的ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络VPC的子网。基础网络下， 该参数无需配置。请登录 [私有网络](https://console.cloud.tencent.com/vpc)控制台查询子网列表获取具体的 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 项目 ID。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)，在右上角的账户信息菜单中，选择**项目管理**查询项目 ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自动续费标识。
	// - 0：默认状态（手动续费）。
	// - 1：自动续费。
	// - 2：到期不续费。
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 安全组 ID 数组。请通过[DescribeInstanceSecurityGroup](https://cloud.tencent.com/document/product/239/34447)接口获取实例的安全组 ID。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// 用户自定义的网络端口。默认为6379，范围为 [1024,65535]。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 实例分片数量。
	// - 标准版实例无需配置该参数。
	// - 集群版实例，分片数量范围为：[1、3、5、8、12、16、24、32、40、48、64、80、96、128]。
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 实例副本数量。
	// - Redis 内存版 4.0、5.0、6.2 标准架构和集群架构支持副本数量范围为[1,5]。
	// - Redis 2.8标准版、CKV标准版只支持1副本。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 标识实例是否需支持副本只读。
	// - Redis 2.8 标准版、CKV标准版不支持副本只读。
	// - 开启副本只读，实例将自动读写分离，写请求路由到主节点，读请求路由到副本节点。
	// - 如需开启副本只读，建议副本数量大于等于2。
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// 实例名称。命名要求：仅支持长度小于60的中文、英文或者数字，短划线"-"、下划线"_"。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 配置实例是否支持免密码访问。
	// - true：免密访问实例。
	// - false：非免密访问实例。默认为非免密方式，仅VPC网络的实例支持免密码访问。
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// 实例的节点信息，包含节点 ID、节点类型、节点可用区 ID等。具体信息，请参见[RedisNodeInfo ](https://cloud.tencent.com/document/product/239/20022)。
	// 目前支持传入节点的类型（主节点或者副本节点），节点的可用区。单可用区部署不需要传递此参数。
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 给实例设定标签。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 指定实例所属的可用区名称。具体信息，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 指定实例相关的参数模板 ID。
	// - 若不配置该参数，则系统会依据所选择的兼容版本及架构，自动适配对应的默认模板。
	// - 请通过[DescribeParamTemplates](https://cloud.tencent.com/document/product/239/58750)接口，查询实例的参数模板列表，获取模板 ID 编号。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 内部参数，标识创建实例是否需要检查。
	// - false ：默认值。发送正常请求，通过检查后直接创建实例。
	// - true：发送检查请求，不会创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 指定实例的产品版本。
	// - local：本地盘版。
	// - cloud：云盘版，
	// - cdc：独享集群版。如果不传默认发货为本地盘版本。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 独享集群 ID。当**ProductVersion**设置为**cdc**时，该参数必须设置。
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// 告警策略 ID 数组。
	// 
	// - 请登录[腾讯云可观测平台-告警管理-策略管理](https://console.cloud.tencent.com/monitor/alarm/policy)获取告警策略 ID。
	// - 若不配置该参数，则绑定默认告警策略。默认告警策略具体信息，请登录[腾讯云可观测平台-告警管理-策略管理](https://console.cloud.tencent.com/monitor/alarm/policy)查看。
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型。
	// <ul><li>2：Redis 2.8 内存版（标准架构）。</li><li>3：CKV 3.2 内存版（标准架构）。</li><li>4：CKV 3.2 内存版（集群架构）。</li><li>6：Redis 4.0 内存版（标准架构）。</li><li>7：Redis 4.0 内存版（集群架构）。</li><li>8：Redis 5.0 内存版（标准架构）。</li><li>9：Redis 5.0 内存版（集群架构）。</li><li>15：Redis 6.2 内存版（标准架构）。</li><li>16：Redis 6.2 内存版（集群架构）。</li><li>17：Redis 7.0 内存版（标准架构）。</li><li>18：Redis 7.0 内存版（集群架构）。</li></ul>
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 内存容量，单位为MB， 数值需为1024的整数倍。具体规格，请通过 [DescribeProductInfo](https://cloud.tencent.com/document/api/239/30600) 接口查询全地域的售卖规格。
	// - **TypeId**为标准架构时，**MemSize**是实例总内存容量；
	// - **TypeId**为集群架构时，**MemSize**是单分片内存容量。
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 实例数量，单次购买实例数量。具体信息，请通过 [DescribeProductInfo](https://cloud.tencent.com/document/api/239/30600) 接口查询全地域的售卖规格。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 购买实例的时长。
	// - 若 **BillingMode**为**1**，即计费方式为包年包月时，需设置该参数，指定所购买实例的时长。单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	// - 若 **BillingMode**为**0**，即计费方式为按量计费时，该参数配置为1。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 计费方式。
	// - 0：按量计费。
	// - 1：包年包月。
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例所属的可用区ID，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 访问实例的密码。
	// - 当输入参数**NoAuth**为**true**时，指设置实例为免密码访问，Password可不用配置，否则Password为必填参数。
	// - 当实例类型**TypeId**为Redis 2.8 内存版标准架构、Redis 4.0、5.0、6.0内存版标准架构或集群架构时，其密码复杂度要求为：8-30个字符，至少包含小写字母、大写字母、数字和字符()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种，不能以"/"开头。
	// - 当实例类型**TypeId**为CKV 3.2 内存版标准架构或集群架构时，其密码复杂度为：8-30个字符，必须包含字母和数字，且 不包含其他字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 私有网络ID。如果不配置该参数则默认选择基础网络。请登录 [私有网络](https://console.cloud.tencent.com/vpc)控制台查询具体的ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络VPC的子网。基础网络下， 该参数无需配置。请登录 [私有网络](https://console.cloud.tencent.com/vpc)控制台查询子网列表获取具体的 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 项目 ID。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)，在右上角的账户信息菜单中，选择**项目管理**查询项目 ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自动续费标识。
	// - 0：默认状态（手动续费）。
	// - 1：自动续费。
	// - 2：到期不续费。
	AutoRenew *uint64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 安全组 ID 数组。请通过[DescribeInstanceSecurityGroup](https://cloud.tencent.com/document/product/239/34447)接口获取实例的安全组 ID。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitnil,omitempty" name:"SecurityGroupIdList"`

	// 用户自定义的网络端口。默认为6379，范围为 [1024,65535]。
	VPort *uint64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 实例分片数量。
	// - 标准版实例无需配置该参数。
	// - 集群版实例，分片数量范围为：[1、3、5、8、12、16、24、32、40、48、64、80、96、128]。
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 实例副本数量。
	// - Redis 内存版 4.0、5.0、6.2 标准架构和集群架构支持副本数量范围为[1,5]。
	// - Redis 2.8标准版、CKV标准版只支持1副本。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 标识实例是否需支持副本只读。
	// - Redis 2.8 标准版、CKV标准版不支持副本只读。
	// - 开启副本只读，实例将自动读写分离，写请求路由到主节点，读请求路由到副本节点。
	// - 如需开启副本只读，建议副本数量大于等于2。
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// 实例名称。命名要求：仅支持长度小于60的中文、英文或者数字，短划线"-"、下划线"_"。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 配置实例是否支持免密码访问。
	// - true：免密访问实例。
	// - false：非免密访问实例。默认为非免密方式，仅VPC网络的实例支持免密码访问。
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// 实例的节点信息，包含节点 ID、节点类型、节点可用区 ID等。具体信息，请参见[RedisNodeInfo ](https://cloud.tencent.com/document/product/239/20022)。
	// 目前支持传入节点的类型（主节点或者副本节点），节点的可用区。单可用区部署不需要传递此参数。
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 给实例设定标签。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 指定实例所属的可用区名称。具体信息，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 指定实例相关的参数模板 ID。
	// - 若不配置该参数，则系统会依据所选择的兼容版本及架构，自动适配对应的默认模板。
	// - 请通过[DescribeParamTemplates](https://cloud.tencent.com/document/product/239/58750)接口，查询实例的参数模板列表，获取模板 ID 编号。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 内部参数，标识创建实例是否需要检查。
	// - false ：默认值。发送正常请求，通过检查后直接创建实例。
	// - true：发送检查请求，不会创建实例。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 指定实例的产品版本。
	// - local：本地盘版。
	// - cloud：云盘版，
	// - cdc：独享集群版。如果不传默认发货为本地盘版本。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 独享集群 ID。当**ProductVersion**设置为**cdc**时，该参数必须设置。
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// 告警策略 ID 数组。
	// 
	// - 请登录[腾讯云可观测平台-告警管理-策略管理](https://console.cloud.tencent.com/monitor/alarm/policy)获取告警策略 ID。
	// - 若不配置该参数，则绑定默认告警策略。默认告警策略具体信息，请登录[腾讯云可观测平台-告警管理-策略管理](https://console.cloud.tencent.com/monitor/alarm/policy)查看。
	AlarmPolicyList []*string `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "Password")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "AutoRenew")
	delete(f, "SecurityGroupIdList")
	delete(f, "VPort")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "InstanceName")
	delete(f, "NoAuth")
	delete(f, "NodeSet")
	delete(f, "ResourceTags")
	delete(f, "ZoneName")
	delete(f, "TemplateId")
	delete(f, "DryRun")
	delete(f, "ProductVersion")
	delete(f, "RedisClusterId")
	delete(f, "AlarmPolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 交易的ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 实例ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstancesResponseParams `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateRequestParams struct {
	// 参数模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 产品类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 源参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数列表。
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type CreateParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 产品类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 源参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数列表。
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *CreateParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ProductType")
	delete(f, "TemplateId")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateParamTemplateResponseParams struct {
	// 参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateParamTemplateResponseParams `json:"Response"`
}

func (r *CreateParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationGroupRequestParams struct {
	// 指定复制组中的主实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 复制组名称。名称只支持长度为2-64个字符的中文、英文、数字、下划线_、分隔符-。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// 指定复制组中的主实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 复制组名称。名称只支持长度为2-64个字符的中文、英文、数字、下划线_、分隔符-。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReplicationGroupResponseParams struct {
	// 异步流程ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateReplicationGroupResponseParams `json:"Response"`
}

func (r *CreateReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DelayDistribution struct {
	// 指延时分布阶梯，其与延时区间的对应关系如下所示。
	// - 1：[0ms,1ms]。
	// - 5： [1ms,5ms]。
	// - 10： [5ms,10ms]。
	// - 50： [10ms,50ms]。
	// - 200：[50ms,200ms]。
	// - -1： [200ms,∞]。
	Ladder *int64 `json:"Ladder,omitnil,omitempty" name:"Ladder"`

	// 延时处于当前分布阶梯的命令数量，单位：个。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 修改时间。
	Updatetime *int64 `json:"Updatetime,omitnil,omitempty" name:"Updatetime"`
}

// Predefined struct for user
type DeleteInstanceAccountRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子账号名称
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`
}

type DeleteInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子账号名称
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceAccountResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceAccountResponseParams `json:"Response"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateRequestParams struct {
	// 参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DeleteParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DeleteParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteParamTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteParamTemplateResponseParams `json:"Response"`
}

func (r *DeleteParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceRequestParams struct {
	// 复制组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据同步类型，true:需要数据强同步,false:不需要强同步，仅限删除主实例
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type DeleteReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 复制组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据同步类型，true:需要数据强同步,false:不需要强同步，仅限删除主实例
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

func (r *DeleteReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteReplicationInstanceResponseParams struct {
	// 异步任务ID
	TaskId *float64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteReplicationInstanceResponseParams `json:"Response"`
}

func (r *DeleteReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackupConfigRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoBackupConfigResponseParams struct {
	// 该参数因兼容性问题暂时保留，请忽略。
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`

	// 备份周期，默认为每天自动备份，Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 备份任务发起时间段。
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// 全量备份文件保存天数。默认为7天。如需保存更多天数，请[提交工单](https://console.cloud.tencent.com/workorder/category)申请。
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil,omitempty" name:"BackupStorageDays"`

	// 该参数不再使用，请忽略。
	BinlogStorageDays *int64 `json:"BinlogStorageDays,omitnil,omitempty" name:"BinlogStorageDays"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoBackupConfigResponseParams `json:"Response"`
}

func (r *DescribeAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDetailRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份 ID，可通过接口 [DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011) 返回的参数 **RedisBackupSet** 获取。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

type DescribeBackupDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份 ID，可通过接口 [DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011) 返回的参数 **RedisBackupSet** 获取。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`
}

func (r *DescribeBackupDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDetailResponseParams struct {
	// 备份 ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份方式。 
	// 
	// - 1：手动备份。
	// -  0：自动备份。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份状态。 
	// 
	// - 1：备份被其它流程锁定。
	// - 2：备份正常，没有被任何流程锁定。
	// - -1：备份已过期。
	// - 3：备份正在被导出。
	// - 4：备份导出成功。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 备份的备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 备份是否被锁定。
	// 
	// - 0：未被锁定。
	// - 1：已被锁定。
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 备份文件大小。单位：Byte。
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 实例类型。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 单分片内存规格大小，单位：MB。
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 分片数量。
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 副本数量。
	ReplicasNum *int64 `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDetailResponseParams `json:"Response"`
}

func (r *DescribeBackupDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionRequestParams struct {

}

type DescribeBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupDownloadRestrictionResponseParams struct {
	// 下载备份文件的网络限制类型：
	// 
	// - NoLimit：不限制，腾讯云内外网均可以下载备份文件。
	// -  LimitOnlyIntranet：仅腾讯云自动分配的内网地址可下载备份文件。
	// - Customize：指用户自定义的私有网络可下载备份文件。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持输入 In，表示自定义的**LimitVpc**可以下载备份文件。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// 标识自定义的 LimitIp 地址是否可下载备份文件。
	// 
	// - In: 自定义的 IP 地址可以下载。
	// - NotIn: 自定义的 IP 不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 自定义的可下载备份文件的 VPC ID。当参数**LimitType**为**Customize **时，显示该参数。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 自定义的可下载备份文件的 VPC IP 地址。当参数**LimitType**为**Customize **时，显示该参数。
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *DescribeBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUrlRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份 ID，可通过 [DescribeInstanceBackups ](https://cloud.tencent.com/document/product/239/20011)接口返回的参数 RedisBackupSet 获取。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 下载备份文件的网络限制类型，如果不配置该参数，则使用用户自定义的配置。
	// 
	// - NoLimit：不限制，腾讯云内外网均可以下载备份文件。
	// -  LimitOnlyIntranet：仅腾讯云自动分配的内网地址可下载备份文件。
	// - Customize：指用户自定义的私有网络可下载备份文件。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持输入 In，表示自定义的**LimitVpc**可以下载备份文件。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// 标识自定义的 LimitIp 地址是否可下载备份文件。
	// 
	// - In: 自定义的 IP 地址可以下载。默认为 In。
	// - NotIn: 自定义的 IP 不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 自定义的可下载备份文件的 VPC ID。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 自定义的可下载备份文件的 VPC IP 地址。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

type DescribeBackupUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份 ID，可通过 [DescribeInstanceBackups ](https://cloud.tencent.com/document/product/239/20011)接口返回的参数 RedisBackupSet 获取。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 下载备份文件的网络限制类型，如果不配置该参数，则使用用户自定义的配置。
	// 
	// - NoLimit：不限制，腾讯云内外网均可以下载备份文件。
	// -  LimitOnlyIntranet：仅腾讯云自动分配的内网地址可下载备份文件。
	// - Customize：指用户自定义的私有网络可下载备份文件。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持输入 In，表示自定义的**LimitVpc**可以下载备份文件。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// 标识自定义的 LimitIp 地址是否可下载备份文件。
	// 
	// - In: 自定义的 IP 地址可以下载。默认为 In。
	// - NotIn: 自定义的 IP 不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 自定义的可下载备份文件的 VPC ID。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 自定义的可下载备份文件的 VPC IP 地址。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

func (r *DescribeBackupUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUrlResponseParams struct {
	// 外网下载地址（6小时内链接有效），该字段正在逐步废弃中。
	DownloadUrl []*string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 内网下载地址（6小时内链接有效），该字段正在逐步废弃中。
	InnerDownloadUrl []*string `json:"InnerDownloadUrl,omitnil,omitempty" name:"InnerDownloadUrl"`

	// 文件名称，该字段正在逐步废弃中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Filenames []*string `json:"Filenames,omitnil,omitempty" name:"Filenames"`

	// 备份文件信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupInfos []*BackupDownloadInfo `json:"BackupInfos,omitnil,omitempty" name:"BackupInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBackupUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupUrlResponseParams `json:"Response"`
}

func (r *DescribeBackupUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthRangeRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBandwidthRangeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeBandwidthRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthRangeResponseParams struct {
	// 标准带宽。指购买实例时，系统为每个节点分配的带宽。
	BaseBandwidth *int64 `json:"BaseBandwidth,omitnil,omitempty" name:"BaseBandwidth"`

	// 指实例的附加带宽。标准带宽不满足需求的情况下，用户可自行增加的带宽。<ul><li>开启副本只读时，实例总带宽 = 附加带宽 * 分片数 + 标准带宽 * 分片数 * Max ([只读副本数量, 1])，标准架构的分片数 = 1。</li><li>没有开启副本只读时，实例总带宽 = 附加带宽 * 分片数 + 标准带宽 * 分片数，标准架构的分片数 = 1。</li></ul>
	AddBandwidth *int64 `json:"AddBandwidth,omitnil,omitempty" name:"AddBandwidth"`

	// 附加带宽设置下限。
	MinAddBandwidth *int64 `json:"MinAddBandwidth,omitnil,omitempty" name:"MinAddBandwidth"`

	// 附加带宽设置上限。
	MaxAddBandwidth *int64 `json:"MaxAddBandwidth,omitnil,omitempty" name:"MaxAddBandwidth"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBandwidthRangeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthRangeResponseParams `json:"Response"`
}

func (r *DescribeBandwidthRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonDBInstancesRequestParams struct {
	// vpc网络ID信息列表
	VpcIds []*int64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 子网ID信息列表
	SubnetIds []*int64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 计费类型过滤列表；0表示包年包月，1表示按量计费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例ID过滤信息列表，数组最大长度限制为100
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称过滤信息列表
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例状态信息过滤列表
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 实例vip信息列表
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// vpc网络ID信息列表
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// 子网统一ID列表
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// 数量限制，默认推荐100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCommonDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// vpc网络ID信息列表
	VpcIds []*int64 `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 子网ID信息列表
	SubnetIds []*int64 `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 计费类型过滤列表；0表示包年包月，1表示按量计费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例ID过滤信息列表，数组最大长度限制为100
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称过滤信息列表
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例状态信息过滤列表
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序字段
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 实例vip信息列表
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// vpc网络ID信息列表
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// 子网统一ID列表
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// 数量限制，默认推荐100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCommonDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "SubnetIds")
	delete(f, "PayMode")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "Status")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "Vips")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCommonDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCommonDBInstancesResponseParams struct {
	// 实例数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例信息
	InstanceDetails []*RedisCommonInstanceList `json:"InstanceDetails,omitnil,omitempty" name:"InstanceDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCommonDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCommonDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeCommonDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCommonDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组规则。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 实例内网IPv4地址。
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// 内网端口。
	VPort *string `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalReplicationAreaRequestParams struct {

}

type DescribeGlobalReplicationAreaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeGlobalReplicationAreaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalReplicationAreaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlobalReplicationAreaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlobalReplicationAreaResponseParams struct {
	// 可用地域信息
	AvailableRegions []*AvailableRegion `json:"AvailableRegions,omitnil,omitempty" name:"AvailableRegions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlobalReplicationAreaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlobalReplicationAreaResponseParams `json:"Response"`
}

func (r *DescribeGlobalReplicationAreaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlobalReplicationAreaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAccountRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceAccountResponseParams struct {
	// 账号详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 账号个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceAccountResponseParams `json:"Response"`
}

func (r *DescribeInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsRequestParams struct {
	// 每页输出的备份列表大小。默认大小为20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 待操作的实例ID，可通过 DescribeInstance 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式如：2017-02-08 16:46:34。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份任务的状态：
	// 1：备份在流程中。
	// 2：备份正常。
	// 3：备份转RDB文件处理中。
	// 4：已完成RDB转换。
	// -1：备份已过期。
	// -2：备份已删除。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例名称，支持根据实例名称模糊搜索。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 每页输出的备份列表大小。默认大小为20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 待操作的实例ID，可通过 DescribeInstance 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，格式如：2017-02-08 16:46:34。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份任务的状态：
	// 1：备份在流程中。
	// 2：备份正常。
	// 3：备份转RDB文件处理中。
	// 4：已完成RDB转换。
	// -1：备份已过期。
	// -2：备份已删除。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例名称，支持根据实例名称模糊搜索。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *DescribeInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Status")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsResponseParams struct {
	// 备份总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例的备份数组。
	BackupSet []*RedisBackupSet `json:"BackupSet,omitnil,omitempty" name:"BackupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceBackupsResponseParams `json:"Response"`
}

func (r *DescribeInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDTSInfoRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceDTSInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDTSInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDTSInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDTSInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDTSInfoResponseParams struct {
	// DTS任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// DTS任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 任务状态,取值为：1-创建中(Creating),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-撤销中（Stopping）,12-完成中（Completing）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 同步时延，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 断开时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CutDownTime *string `json:"CutDownTime,omitnil,omitempty" name:"CutDownTime"`

	// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *DescribeInstanceDTSInstanceInfo `json:"SrcInfo,omitnil,omitempty" name:"SrcInfo"`

	// 目标实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *DescribeInstanceDTSInstanceInfo `json:"DstInfo,omitnil,omitempty" name:"DstInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDTSInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDTSInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceDTSInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDTSInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInstanceInfo struct {
	// 地域 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 仓库ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetId *int64 `json:"SetId,omitnil,omitempty" name:"SetId"`

	// 可用区ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例访问地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DescribeInstanceDealDetailRequestParams struct {
	// 订单交易ID数组，即 [CreateInstances](https://cloud.tencent.com/document/api/239/20026) 的输出参数DealId。数组最大长度限制为10
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest
	
	// 订单交易ID数组，即 [CreateInstances](https://cloud.tencent.com/document/api/239/20026) 的输出参数DealId。数组最大长度限制为10
	DealIds []*string `json:"DealIds,omitnil,omitempty" name:"DealIds"`
}

func (r *DescribeInstanceDealDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDealDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceDealDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDealDetailResponseParams struct {
	// 订单详细信息。
	DealDetails []*TradeDealDetail `json:"DealDetails,omitnil,omitempty" name:"DealDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceDealDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceDealDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceDealDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceDealDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceEventsRequestParams struct {
	// 配置查询事件执行计划的起始日期。
	ExecutionStartDate *string `json:"ExecutionStartDate,omitnil,omitempty" name:"ExecutionStartDate"`

	// 配置查询事件执行计划的结束日期。
	ExecutionEndDate *string `json:"ExecutionEndDate,omitnil,omitempty" name:"ExecutionEndDate"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 输出每页显示事件的数量，默认：10。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 配置查询事件的输出页码，即支持根据PageNo（页码）与 PageSize （每页输出数量）查询某一页的事件。默认：1。
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 事件当前状态。
	// - Waiting：未到达执行日期或不在维护时间窗内的事件。
	// - Running：在维护时间窗内，正在执行维护的事件。
	// - Finished：已全部完成维护的事件。
	// - Canceled：已取消执行的事件。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件类型，当前仅支持配置实例迁移、资源腾挪、机房裁撤相关的运维操作。该参数仅支持配置为 **InstanceMigration**。
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// 配置查询事件等级。事件等级根据其影响严重程度和紧急程度进行分级，由重至轻依次为关键、重要、中等、一般。
	// - Critical：关键
	// - High：重要
	// - Middle：中等
	// - Low：一般
	Grades []*string `json:"Grades,omitnil,omitempty" name:"Grades"`
}

type DescribeInstanceEventsRequest struct {
	*tchttp.BaseRequest
	
	// 配置查询事件执行计划的起始日期。
	ExecutionStartDate *string `json:"ExecutionStartDate,omitnil,omitempty" name:"ExecutionStartDate"`

	// 配置查询事件执行计划的结束日期。
	ExecutionEndDate *string `json:"ExecutionEndDate,omitnil,omitempty" name:"ExecutionEndDate"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 输出每页显示事件的数量，默认：10。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 配置查询事件的输出页码，即支持根据PageNo（页码）与 PageSize （每页输出数量）查询某一页的事件。默认：1。
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 事件当前状态。
	// - Waiting：未到达执行日期或不在维护时间窗内的事件。
	// - Running：在维护时间窗内，正在执行维护的事件。
	// - Finished：已全部完成维护的事件。
	// - Canceled：已取消执行的事件。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件类型，当前仅支持配置实例迁移、资源腾挪、机房裁撤相关的运维操作。该参数仅支持配置为 **InstanceMigration**。
	EventTypes []*string `json:"EventTypes,omitnil,omitempty" name:"EventTypes"`

	// 配置查询事件等级。事件等级根据其影响严重程度和紧急程度进行分级，由重至轻依次为关键、重要、中等、一般。
	// - Critical：关键
	// - High：重要
	// - Middle：中等
	// - Low：一般
	Grades []*string `json:"Grades,omitnil,omitempty" name:"Grades"`
}

func (r *DescribeInstanceEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionStartDate")
	delete(f, "ExecutionEndDate")
	delete(f, "InstanceId")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Status")
	delete(f, "EventTypes")
	delete(f, "Grades")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceEventsResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例事件信息
	RedisInstanceEvents []*RedisInstanceEvent `json:"RedisInstanceEvents,omitnil,omitempty" name:"RedisInstanceEvents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceEventsResponseParams `json:"Response"`
}

func (r *DescribeInstanceEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogDeliveryRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceLogDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceLogDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogDeliveryResponseParams struct {
	// 实例慢日志投递信息。
	SlowLog *LogDeliveryInfo `json:"SlowLog,omitnil,omitempty" name:"SlowLog"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLogDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogDeliveryResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 请求类型：1——string类型，2——所有类型
	ReqType *int64 `json:"ReqType,omitnil,omitempty" name:"ReqType"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeInstanceMonitorBigKeyRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 请求类型：1——string类型，2——所有类型
	ReqType *int64 `json:"ReqType,omitnil,omitempty" name:"ReqType"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReqType")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyResponseParams struct {
	// 大Key详细信息
	Data []*BigKeyInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeyResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeySizeDistRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeInstanceMonitorBigKeySizeDistRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeySizeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeySizeDistResponseParams struct {
	// 大Key大小分布详情
	Data []*DelayDistribution `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeySizeDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeySizeDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeySizeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyTypeDistRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

type DescribeInstanceMonitorBigKeyTypeDistRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorBigKeyTypeDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorBigKeyTypeDistResponseParams struct {
	// 大Key类型分布详细信息
	Data []*BigKeyTypeInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorBigKeyTypeDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorBigKeyTypeDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorHotKeyRequestParams struct {
	// 指定实例 ID。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorHotKeyRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorHotKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorHotKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorHotKeyResponseParams struct {
	// 热 Key 详细信息。
	Data []*HotKeyInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorHotKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorHotKeyResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorHotKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorHotKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorSIPRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceMonitorSIPRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceMonitorSIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorSIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorSIPResponseParams struct {
	// 访问来源信息
	Data []*SourceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorSIPResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorSIPResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorSIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorSIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTookDistRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询时间日期。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorTookDistRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询时间日期。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTookDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTookDistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTookDistResponseParams struct {
	// 时延分布信息。
	Data []*DelayDistribution `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorTookDistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTookDistResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTookDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTookDistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorTopNCmdRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdResponseParams struct {
	// 访问命令信息
	Data []*SourceCommand `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorTopNCmdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTopNCmdResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdTookRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

type DescribeInstanceMonitorTopNCmdTookRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询时间范围。
	// - 1：实时。
	// - 2：近30分钟。
	// - 3：近6小时。
	// - 4：近24小时。
	SpanType *int64 `json:"SpanType,omitnil,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceMonitorTopNCmdTookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceMonitorTopNCmdTookResponseParams struct {
	// 耗时详细信息
	Data []*CommandTake `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceMonitorTopNCmdTookResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceMonitorTopNCmdTookResponseParams `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceMonitorTopNCmdTookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodeInfoRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 列表大小。每页输出的节点信息大小。默认为 20，最多输出1000条。该字段已不再使用，请忽略。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。该字段已不再使用，请忽略。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 列表大小。每页输出的节点信息大小。默认为 20，最多输出1000条。该字段已不再使用，请忽略。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。该字段已不再使用，请忽略。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodeInfoResponseParams struct {
	// Proxy节点数量。
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// Proxy节点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proxy []*ProxyNodes `json:"Proxy,omitnil,omitempty" name:"Proxy"`

	// Redis节点数量。
	RedisCount *int64 `json:"RedisCount,omitnil,omitempty" name:"RedisCount"`

	// Redis节点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Redis []*RedisNodes `json:"Redis,omitnil,omitempty" name:"Redis"`

	// 该参数不再使用，请忽略。
	TendisCount *int64 `json:"TendisCount,omitnil,omitempty" name:"TendisCount"`

	// 该参数不再使用，请忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tendis []*TendisNodes `json:"Tendis,omitnil,omitempty" name:"Tendis"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodeInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// 总的修改历史记录数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 修改历史记录信息。
	InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitnil,omitempty" name:"InstanceParamHistory"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 参数列表总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例枚举类型参数。
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitnil,omitempty" name:"InstanceEnumParam"`

	// 实例整型参数。
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitnil,omitempty" name:"InstanceIntegerParam"`

	// 实例字符型参数。
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitnil,omitempty" name:"InstanceTextParam"`

	// 实例多选项型参数。
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitnil,omitempty" name:"InstanceMultiParam"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSecurityGroupRequestParams struct {
	// 实例 ID 列表，数组长度限制[0,100]。例如：["crs-f2ho5rsz\n"]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表，数组长度限制[0,100]。例如：["crs-f2ho5rsz\n"]
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSecurityGroupResponseParams struct {
	// 实例安全组信息。
	InstanceSecurityGroupsDetail []*InstanceSecurityGroupDetail `json:"InstanceSecurityGroupsDetail,omitnil,omitempty" name:"InstanceSecurityGroupsDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否过滤掉从节信息。
	// - true；过滤从节点。
	// - false：不过滤。
	FilterSlave *bool `json:"FilterSlave,omitnil,omitempty" name:"FilterSlave"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否过滤掉从节信息。
	// - true；过滤从节点。
	// - false：不过滤。
	FilterSlave *bool `json:"FilterSlave,omitnil,omitempty" name:"FilterSlave"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterSlave")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsResponseParams struct {
	// 实例分片列表信息，包括：节点信息、节点ID、Key数量、使用容量、容量倾斜率等信息。
	InstanceShards []*InstanceClusterShard `json:"InstanceShards,omitnil,omitempty" name:"InstanceShards"`

	// 实例分片节点数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceShardsResponseParams `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecBandwidthRequestParams struct {

}

type DescribeInstanceSpecBandwidthRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeInstanceSpecBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSpecBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSpecBandwidthResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSpecBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSpecBandwidthResponseParams `json:"Response"`
}

func (r *DescribeInstanceSpecBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSpecBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSupportFeatureRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	// 示例值：crs-asdasdas
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 功能特性名称
	// - read-local-node-only 就近接入功能
	// - multi-account 多账号功能
	FeatureName *string `json:"FeatureName,omitnil,omitempty" name:"FeatureName"`
}

type DescribeInstanceSupportFeatureRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	// 示例值：crs-asdasdas
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 功能特性名称
	// - read-local-node-only 就近接入功能
	// - multi-account 多账号功能
	FeatureName *string `json:"FeatureName,omitnil,omitempty" name:"FeatureName"`
}

func (r *DescribeInstanceSupportFeatureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSupportFeatureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FeatureName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceSupportFeatureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceSupportFeatureResponseParams struct {
	// 是否支持
	Support *bool `json:"Support,omitnil,omitempty" name:"Support"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceSupportFeatureResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceSupportFeatureResponseParams `json:"Response"`
}

func (r *DescribeInstanceSupportFeatureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceSupportFeatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceZoneInfoRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceZoneInfoRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceZoneInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceZoneInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceZoneInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceZoneInfoResponseParams struct {
	// 实例节点组的个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例节点组列表。
	ReplicaGroups []*ReplicaGroup `json:"ReplicaGroups,omitnil,omitempty" name:"ReplicaGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceZoneInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceZoneInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceZoneInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceZoneInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 每页输出实例的数量，参数默认值20，最大值为1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例列表排序依据，枚举值如下所示：
	// - projectId：依据项目ID排序。
	// - createtime：依据实例创建时间排序。
	// - instancename：依据实例名称排序。
	// - type：依据实例类型排序。
	// - curDeadline：依据实例到期时间排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 实例排序方式，默认为倒序排序。
	// - 1：倒序。
	// - 0：顺序。
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 私有网络 ID 数组。如果不配置该参数或设置数组为空则默认选择基础网络。例如47525。该参数暂时保留，可忽略。请根据 UniqVpcIds 参数格式设置私有网络ID数组。
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 私有网络所属子网 ID 数组，例如：56854。该参数暂时保留，可忽略。请根据 UniqSubnetIds 参数格式设置私有网络子网 ID 数组。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 设置模糊查询关键字段，仅实例名称支持模糊查询。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目 ID 组成的数组。
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 私有网络 ID 数组。如果不配置该参数或者设置数组为空则默认选择基础网络，如：vpc-sad23jfdfk。
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// 私有网络所属子网 ID 数组，如：subnet-fdj24n34j2。
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// 地域 ID 数组，该参数已经弃用，可通过公共参数Region查询对应地域。
	RegionIds []*int64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 实例状态。
	// - 0：待初始化。
	// - 1：流程中。
	// - 2：运行中。
	// - -2：已隔离。
	// - -3：待删除。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例架构版本。
	// - 1：单机版。
	// - 2：主从版。
	// - 3：集群版。
	TypeVersion *int64 `json:"TypeVersion,omitnil,omitempty" name:"TypeVersion"`

	// 存储引擎信息。可设置为Redis-2.8、Redis-4.0、Redis-5.0、Redis-6.0 或者 CKV。
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 续费模式。
	// - 0：手动续费。
	// - 1：自动续费。
	// - 2：到期不再续费。
	AutoRenew []*int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 计费模式。
	// - postpaid：按量计费。
	// - prepaid：包年包月。
	BillingMode *string `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 3：CKV 3.2 内存版（标准架构）。
	// - 4：CKV 3.2 内存版（集群架构）。
	// - 5：Redis 2.8 内存版（单机）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 该参数为数组类型，支持配置实例名称、实例 ID、IP地址，其中实例名称为模糊匹配，实例 ID 和 IP 地址精确匹配。
	// 
	// - 数组中每一个元素取并集进行匹配查询。
	// - **InstanceId** 与 **SearchKeys** 同时配置，则取二者交集进行匹配查询。
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// 内部参数，用户可忽略。
	TypeList []*int64 `json:"TypeList,omitnil,omitempty" name:"TypeList"`

	// 内部参数，用户可忽略。
	MonitorVersion *string `json:"MonitorVersion,omitnil,omitempty" name:"MonitorVersion"`

	// 根据标签的 Key 和 Value 筛选资源。该参数不配置或者数组设置为空值，则不根据标签进行过滤。
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 根据标签的 Key 筛选资源，该参数不配置或者数组设置为空值，则不根据标签Key进行过滤。
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 实例的产品版本。如果该参数不配置或者数组设置为空值，则默认不依据此参数过滤实例。
	// - local：本地盘版。
	// - cdc：独享集群版。
	ProductVersions []*string `json:"ProductVersions,omitnil,omitempty" name:"ProductVersions"`

	// 批量查询指定的实例 ID，返回结果已 Limit 限制为主。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 可用区模式。
	// - singleaz：单可用区。
	// - multiaz：多可用区。
	AzMode *string `json:"AzMode,omitnil,omitempty" name:"AzMode"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 每页输出实例的数量，参数默认值20，最大值为1000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	// 
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例列表排序依据，枚举值如下所示：
	// - projectId：依据项目ID排序。
	// - createtime：依据实例创建时间排序。
	// - instancename：依据实例名称排序。
	// - type：依据实例类型排序。
	// - curDeadline：依据实例到期时间排序。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 实例排序方式，默认为倒序排序。
	// - 1：倒序。
	// - 0：顺序。
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 私有网络 ID 数组。如果不配置该参数或设置数组为空则默认选择基础网络。例如47525。该参数暂时保留，可忽略。请根据 UniqVpcIds 参数格式设置私有网络ID数组。
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 私有网络所属子网 ID 数组，例如：56854。该参数暂时保留，可忽略。请根据 UniqSubnetIds 参数格式设置私有网络子网 ID 数组。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 设置模糊查询关键字段，仅实例名称支持模糊查询。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目 ID 组成的数组。
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 私有网络 ID 数组。如果不配置该参数或者设置数组为空则默认选择基础网络，如：vpc-sad23jfdfk。
	UniqVpcIds []*string `json:"UniqVpcIds,omitnil,omitempty" name:"UniqVpcIds"`

	// 私有网络所属子网 ID 数组，如：subnet-fdj24n34j2。
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitnil,omitempty" name:"UniqSubnetIds"`

	// 地域 ID 数组，该参数已经弃用，可通过公共参数Region查询对应地域。
	RegionIds []*int64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 实例状态。
	// - 0：待初始化。
	// - 1：流程中。
	// - 2：运行中。
	// - -2：已隔离。
	// - -3：待删除。
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例架构版本。
	// - 1：单机版。
	// - 2：主从版。
	// - 3：集群版。
	TypeVersion *int64 `json:"TypeVersion,omitnil,omitempty" name:"TypeVersion"`

	// 存储引擎信息。可设置为Redis-2.8、Redis-4.0、Redis-5.0、Redis-6.0 或者 CKV。
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 续费模式。
	// - 0：手动续费。
	// - 1：自动续费。
	// - 2：到期不再续费。
	AutoRenew []*int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 计费模式。
	// - postpaid：按量计费。
	// - prepaid：包年包月。
	BillingMode *string `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 3：CKV 3.2 内存版（标准架构）。
	// - 4：CKV 3.2 内存版（集群架构）。
	// - 5：Redis 2.8 内存版（单机）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 该参数为数组类型，支持配置实例名称、实例 ID、IP地址，其中实例名称为模糊匹配，实例 ID 和 IP 地址精确匹配。
	// 
	// - 数组中每一个元素取并集进行匹配查询。
	// - **InstanceId** 与 **SearchKeys** 同时配置，则取二者交集进行匹配查询。
	SearchKeys []*string `json:"SearchKeys,omitnil,omitempty" name:"SearchKeys"`

	// 内部参数，用户可忽略。
	TypeList []*int64 `json:"TypeList,omitnil,omitempty" name:"TypeList"`

	// 内部参数，用户可忽略。
	MonitorVersion *string `json:"MonitorVersion,omitnil,omitempty" name:"MonitorVersion"`

	// 根据标签的 Key 和 Value 筛选资源。该参数不配置或者数组设置为空值，则不根据标签进行过滤。
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 根据标签的 Key 筛选资源，该参数不配置或者数组设置为空值，则不根据标签Key进行过滤。
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 实例的产品版本。如果该参数不配置或者数组设置为空值，则默认不依据此参数过滤实例。
	// - local：本地盘版。
	// - cdc：独享集群版。
	ProductVersions []*string `json:"ProductVersions,omitnil,omitempty" name:"ProductVersions"`

	// 批量查询指定的实例 ID，返回结果已 Limit 限制为主。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 可用区模式。
	// - singleaz：单可用区。
	// - multiaz：多可用区。
	AzMode *string `json:"AzMode,omitnil,omitempty" name:"AzMode"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "InstanceId")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "VpcIds")
	delete(f, "SubnetIds")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "InstanceName")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "RegionIds")
	delete(f, "Status")
	delete(f, "TypeVersion")
	delete(f, "EngineName")
	delete(f, "AutoRenew")
	delete(f, "BillingMode")
	delete(f, "Type")
	delete(f, "SearchKeys")
	delete(f, "TypeList")
	delete(f, "MonitorVersion")
	delete(f, "InstanceTags")
	delete(f, "TagKeys")
	delete(f, "ProductVersions")
	delete(f, "InstanceIds")
	delete(f, "AzMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表。
	InstanceSet []*InstanceSet `json:"InstanceSet,omitnil,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceWindowResponseParams struct {
	// 维护时间窗开始时间。取值范围为"00:00-23:00"的任意时间点，如03:24。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 维护时间窗结束时间。
	// - 取值范围为"00:00-23:00"的任意时间点，如：04:24。
	// - 维护时间持续时长最小为30分钟，最大为3小时。
	// - 结束时间务必是基于开始时间向后的时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintenanceWindowResponseParams `json:"Response"`
}

func (r *DescribeMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoRequestParams struct {
	// 指定查询的参数模板 ID。请通过接口[DescribeParamTemplates](https://cloud.tencent.com/document/product/239/58750)获取参数模板列表信息。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeParamTemplateInfoRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询的参数模板 ID。请通过接口[DescribeParamTemplates](https://cloud.tencent.com/document/product/239/58750)获取参数模板列表信息。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeParamTemplateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplateInfoResponseParams struct {
	// 参数模板的参数数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 产品类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 3：CKV 3.2 内存版（标准架构）。
	// - 4：CKV 3.2 内存版（集群架构）。
	// - 5：Redis 2.8 内存版（单机）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 参数模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数详情。包含：参数的名称，当前运行值，默认值，最大值、最小值、枚举值等信息。
	Items []*ParameterDetail `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplateInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplateInfoResponseParams `json:"Response"`
}

func (r *DescribeParamTemplateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesRequestParams struct {
	// 产品类型数组。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 3：CKV 3.2 内存版（标准架构）。
	// - 4：CKV 3.2 内存版（集群架构）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	ProductTypes []*int64 `json:"ProductTypes,omitnil,omitempty" name:"ProductTypes"`

	// 模板名称数组。数组最大长度限制为50
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// 模板ID数组。数组最大长度限制为50
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

type DescribeParamTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 产品类型数组。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 3：CKV 3.2 内存版（标准架构）。
	// - 4：CKV 3.2 内存版（集群架构）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	ProductTypes []*int64 `json:"ProductTypes,omitnil,omitempty" name:"ProductTypes"`

	// 模板名称数组。数组最大长度限制为50
	TemplateNames []*string `json:"TemplateNames,omitnil,omitempty" name:"TemplateNames"`

	// 模板ID数组。数组最大长度限制为50
	TemplateIds []*string `json:"TemplateIds,omitnil,omitempty" name:"TemplateIds"`
}

func (r *DescribeParamTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductTypes")
	delete(f, "TemplateNames")
	delete(f, "TemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeParamTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeParamTemplatesResponseParams struct {
	// 该用户的参数模板数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 参数模板详情。
	Items []*ParamTemplateInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeParamTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeParamTemplatesResponseParams `json:"Response"`
}

func (r *DescribeParamTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeParamTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductInfoRequestParams struct {

}

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductInfoResponseParams struct {
	// 地域售卖信息。
	RegionSet []*RegionConf `json:"RegionSet,omitnil,omitempty" name:"RegionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductInfoResponseParams `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupRequestParams struct {
	// 0:默认项目；-1 所有项目; >0: 特定项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组Id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

type DescribeProjectSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 0:默认项目；-1 所有项目; >0: 特定项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组Id
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeProjectSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupResponseParams struct {
	// 项目安全组
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitnil,omitempty" name:"SecurityGroupDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 偏移量，取值为Limit的整数倍。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取数量限制，默认 20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索条件，支持安全组 ID 或者安全组名称。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 偏移量，取值为Limit的整数倍。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取数量限制，默认 20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索条件，支持安全组 ID 或者安全组名称。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "ProjectId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// 安全组规则。
	Groups []*SecurityGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 符合条件的安全组总数量。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySlowLogRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询的开始时间。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 慢查询的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢查询阈值，单位：毫秒。
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// 分页大小。默认为 20，取值范围[20,1000]。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeProxySlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询的开始时间。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 慢查询的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢查询阈值，单位：毫秒。
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// 分页大小。默认为 20，取值范围[20,1000]。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeProxySlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySlowLogResponseParams struct {
	// 慢查询总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢查询详情。
	InstanceProxySlowLogDetail []*InstanceProxySlowlogDetail `json:"InstanceProxySlowLogDetail,omitnil,omitempty" name:"InstanceProxySlowLogDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySlowLogResponseParams `json:"Response"`
}

func (r *DescribeProxySlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClusterOverviewRequestParams struct {
	// 本地专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeRedisClusterOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 本地专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeRedisClusterOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClusterOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisClusterOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClusterOverviewResponseParams struct {
	// 资源包总数
	TotalBundle *int64 `json:"TotalBundle,omitnil,omitempty" name:"TotalBundle"`

	// 资源包总内存大小，单位：GB
	TotalMemory *int64 `json:"TotalMemory,omitnil,omitempty" name:"TotalMemory"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisClusterOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisClusterOverviewResponseParams `json:"Response"`
}

func (r *DescribeRedisClusterOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClusterOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClustersRequestParams struct {
	// Redis独享集群id
	RedisClusterIds []*string `json:"RedisClusterIds,omitnil,omitempty" name:"RedisClusterIds"`

	// 集群状态：1-流程中，2-运行中，3-已隔离
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目ID数组
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 续费模式：0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenewFlag []*int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Redis独享集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 搜索关键词：支持集群Id、集群名称
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 分页限制返回大小，不传则默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本地专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

type DescribeRedisClustersRequest struct {
	*tchttp.BaseRequest
	
	// Redis独享集群id
	RedisClusterIds []*string `json:"RedisClusterIds,omitnil,omitempty" name:"RedisClusterIds"`

	// 集群状态：1-流程中，2-运行中，3-已隔离
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目ID数组
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 续费模式：0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenewFlag []*int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// Redis独享集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 搜索关键词：支持集群Id、集群名称
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 分页限制返回大小，不传则默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 本地专用集群id
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`
}

func (r *DescribeRedisClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RedisClusterIds")
	delete(f, "Status")
	delete(f, "ProjectIds")
	delete(f, "AutoRenewFlag")
	delete(f, "ClusterName")
	delete(f, "SearchKey")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DedicatedClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisClustersResponseParams struct {
	// 集群总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// CDC集群资源列表
	Resources []*CDCResource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisClustersResponseParams `json:"Response"`
}

func (r *DescribeRedisClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupInstanceRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeReplicationGroupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeReplicationGroupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationGroupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupInstanceResponseParams struct {
	// AppID。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域数字编号。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 复制组字符串ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 复制组名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 实例复制组角色。
	// - r:  备实例
	// - rw: 主实例
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationGroupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationGroupInstanceResponseParams `json:"Response"`
}

func (r *DescribeReplicationGroupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupRequestParams struct {
	// 每页输出实例列表的大小，参数默认值20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定复制组 ID。例如：crs-rpl-m3zt****。请登录[Redis 控制台](https://console.cloud.tencent.com/redis/replication)的全球复制组列表获取复制组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 模糊查询的关键字，可以设置为复制组ID或复制组名称进行模糊查询。请登录[Redis 控制台](https://console.cloud.tencent.com/redis/replication)的全球复制组列表获取复制组 ID及名称。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// 每页输出实例列表的大小，参数默认值20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 指定复制组 ID。例如：crs-rpl-m3zt****。请登录[Redis 控制台](https://console.cloud.tencent.com/redis/replication)的全球复制组列表获取复制组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 模糊查询的关键字，可以设置为复制组ID或复制组名称进行模糊查询。请登录[Redis 控制台](https://console.cloud.tencent.com/redis/replication)的全球复制组列表获取复制组 ID及名称。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "GroupId")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReplicationGroupResponseParams struct {
	// 复制组数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 复制组信息。
	Groups []*Groups `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReplicationGroupResponseParams `json:"Response"`
}

func (r *DescribeReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeSSLStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeSSLStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSSLStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSSLStatusResponseParams struct {
	// SSL 证书下载地址。
	CertDownloadUrl *string `json:"CertDownloadUrl,omitnil,omitempty" name:"CertDownloadUrl"`

	// 证书下载链接到期时间。
	UrlExpiredTime *string `json:"UrlExpiredTime,omitnil,omitempty" name:"UrlExpiredTime"`

	// 标识实例开启 SSL 功能。
	// - true：开启 。
	// - false：关闭。
	SSLConfig *bool `json:"SSLConfig,omitnil,omitempty" name:"SSLConfig"`

	// 标识实例是否支持 SSL特性。
	// - true：支持。
	// - false：不支持。
	FeatureSupport *bool `json:"FeatureSupport,omitnil,omitempty" name:"FeatureSupport"`

	// 说明配置 SSL 的状态。
	// - 1: 配置中。
	// - 2：配置成功。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSSLStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSSLStatusResponseParams `json:"Response"`
}

func (r *DescribeSSLStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSSLStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 预查询慢日志的起始时间。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 预查询慢日志的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢查询平均执行时间阈值，单位：毫秒。
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// 每个页面展示的慢查询条数，默认值为20。取值范围：[20,1000]。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 慢查询条数的偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 节点所属角色。
	// - master：主节点。
	// - slave：从节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 预查询慢日志的起始时间。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 预查询慢日志的结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢查询平均执行时间阈值，单位：毫秒。
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// 每个页面展示的慢查询条数，默认值为20。取值范围：[20,1000]。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 慢查询条数的偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 节点所属角色。
	// - master：主节点。
	// - slave：从节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogResponseParams struct {
	// 慢查询总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 已废弃，该参数存在命名不规范问题，后续用参数InstanceSlowLogDetail取代。慢查询详情。
	InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitnil,omitempty" name:"InstanceSlowlogDetail"`

	// 慢查询详情。
	InstanceSlowLogDetail []*InstanceSlowlogDetail `json:"InstanceSlowLogDetail,omitnil,omitempty" name:"InstanceSlowLogDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogResponseParams `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoRequestParams struct {
	// 任务 ID。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInfoResponseParams struct {
	// 任务状态。
	// - preparing：待执行。
	// - running：执行中。
	// - succeed：成功。
	// - failed：失败。
	// - error：执行出错。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务类型。常见的类型包含：新建类型、配置变更、关闭实例、清空实例、重置密码、版本升级、备份实例、改变网络类型、实例可用区迁移、手动提主等。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实例的 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务执行返回的信息，执行错误时显示错误信息。执行中或执行成功则为空。
	TaskMessage *string `json:"TaskMessage,omitnil,omitempty" name:"TaskMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 每页输出的任务列表大小。默认为 20，最多输出100条。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 项目 ID。登录 [Redis 控制台](https://console.cloud.tencent.com/redis)，在右上角的账号信息下拉菜单中，选择**项目管理**，即可获取对应的项目 ID。
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 任务类型。
	// 
	// - FLOW_CREATE: "001"，新建实例。
	// - FLOW_RESIZE ： "002"，配置变更。
	// - FLOW_CLOSE："003"，关闭实例。
	// - FLOW_CLEAN： "004"，清空实例。
	// - FLOW_STARTUP："005"，实例启用。
	// - FLOW_DELETE："006"，删除实例。
	// - FLOW_SETPWD："007"，重置密码。
	// - FLOW_EXPORTBACKUP："009"，导出备份文件。
	// - FLOW_RESTOREBACKUP："010"，恢复备份。
	// - FLOW_BACKUPINSTANCE："012"，备份实例。
	// - FLOW_MIGRATEINSTANCE："013"，迁移实例。
	// - FLOW_DELBACKUP："014"，删除备份。
	// - FLOW_EXCHANGEINSTANCE： "016"，切换实例流程。
	// - FLOW_AUTOBACKUP："017"，自动备份实例。
	// - FLOW_MIGRATECHECK： "022"，迁移参数校验。
	// - FLOW_MIGRATETASK："023"，数据迁移中。
	// - FLOW_CLEANDB："025"，清空某个数据库。
	// - FLOW_CLONEBACKUP："026"，克隆备份。
	// - FLOW_CHANGEVIP： "027"，改变vip地址。
	// - FLOW_EXPORSHR ："028"，扩缩容。
	// - FLOW_ADDNODES："029"，加（减）节点。
	// - FLOW_CHANGENET："031"，改变网络类型。
	// - FLOW_MODIFYINSTACEREADONLY："033"，只读策略变更。
	// - FLOW_MODIFYINSTANCEPARAMS："034"，修改实例参数。
	// - FLOW_MODIFYINSTANCEPASSWORDFREE："035"，设置免密。
	// - FLOW_SWITCHINSTANCEVIP："036"，实例VIP切换。
	// - FLOW_MODIFYINSTANCEACCOUNT："037"，实例帐号变更。
	// - FLOW_MODIFYINSTANCEBANDWIDTH："038"，实例带宽变更。
	// - FLOW_ENABLEINSTANCE_REPLICATE："039"，开启副本只读。
	// - FLOW_DISABLEINSTANCE_REPLICATE："040"，关闭副本只读。
	// - FLOW_UpgradeArch："041"，实例架构升级，主从升集群。
	// - FLOW_DowngradeArch： "042"，实例架构降级，集群降主从。
	// - FLOW_UpgradeVersion： "043"，版本升级。
	// - FLOW_MODIFYCONNECTIONCONFIG："044"，带宽连接数调整，
	// - FLOW_CLEARNETWORK："045"，更换网络，
	// - FLOW_REMOVE_BACKUP_FILE："046"，删除备份。
	// - FLOW_UPGRADE_SUPPORT_MULTI_AZ："047"，升级实例支持多可用区。
	// - FLOW_SHUTDOWN_MASTER："048"，模拟故障。
	// - FLOW_CHANGE_REPLICA_TO_MASTER："049"，手动提主。
	// - FLOW_CODE_ADD_REPLICATION_INSTANCE："050"，新增复制组。
	// - FLOW_OPEN_WAN："052"，开通外网。
	// - FLOW_CLOSE_WAN："053"，关闭外网FLOW_UPDATE_WAN："054"，更新外网。
	// - FLOW_CODE_DELETE_REPLICATION_INSTANCE："055"，解绑复制组。
	// - FLOW_CODE_CHANGE_MASTER_INSTANCE："056"，复制组实例切主。
	// - FLOW_CODE_CHANGE_INSTANCE_ROLE： "057"，更改复制组实例角色。
	// - FLOW_MIGRATE_NODE："058"，迁移节点。
	// - FLOW_SWITCH_NODE："059"，切换节点。
	// - FLOW_UPGRADE_SMALL_VERSION："060"，升级 Redi s版本。
	// - FLOW_UPGRADE_PROXY_VERSION："061"，升级 Proxy 版本。
	// - FLOW_MODIFY_INSTANCE_NETWORK： "062"，实例修改网络。
	// - FLOW_MIGRATE_PROXY_NODE："063"，迁移proxy节点。
	// - FLOW_MIGRATION_INSTANCE_ZONE："066"，实例可用区迁移中。
	// - FLOW_UPGRADE_INSTANCE_CACHE_AND_PROXY： "067"，实例版本升级中。
	// - FLOW_MODIFY_PROXY_NUM："069"，加（减）Proxy 节点。
	// - FLOW_MODIFYBACKUPMOD："070"，变更实例备份模式。
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// 任务执行的起始时间。格式如：2021-12-30 00:00:00。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 任务运行的终止时间。格式如：2021-12-30 20:59:35
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 该参数为内部使用，请忽略。
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务执行状态。
	// - 0：任务初始化。
	// - 1：执行中。
	// - 2：完成。
	// - 4：失败。
	Result []*int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 该字段已废弃，使用OperateUin代替，请忽略。
	OperatorUin []*int64 `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 操作者账号 ID，UIN。
	OperateUin []*string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 每页输出的任务列表大小。默认为 20，最多输出100条。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 项目 ID。登录 [Redis 控制台](https://console.cloud.tencent.com/redis)，在右上角的账号信息下拉菜单中，选择**项目管理**，即可获取对应的项目 ID。
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 任务类型。
	// 
	// - FLOW_CREATE: "001"，新建实例。
	// - FLOW_RESIZE ： "002"，配置变更。
	// - FLOW_CLOSE："003"，关闭实例。
	// - FLOW_CLEAN： "004"，清空实例。
	// - FLOW_STARTUP："005"，实例启用。
	// - FLOW_DELETE："006"，删除实例。
	// - FLOW_SETPWD："007"，重置密码。
	// - FLOW_EXPORTBACKUP："009"，导出备份文件。
	// - FLOW_RESTOREBACKUP："010"，恢复备份。
	// - FLOW_BACKUPINSTANCE："012"，备份实例。
	// - FLOW_MIGRATEINSTANCE："013"，迁移实例。
	// - FLOW_DELBACKUP："014"，删除备份。
	// - FLOW_EXCHANGEINSTANCE： "016"，切换实例流程。
	// - FLOW_AUTOBACKUP："017"，自动备份实例。
	// - FLOW_MIGRATECHECK： "022"，迁移参数校验。
	// - FLOW_MIGRATETASK："023"，数据迁移中。
	// - FLOW_CLEANDB："025"，清空某个数据库。
	// - FLOW_CLONEBACKUP："026"，克隆备份。
	// - FLOW_CHANGEVIP： "027"，改变vip地址。
	// - FLOW_EXPORSHR ："028"，扩缩容。
	// - FLOW_ADDNODES："029"，加（减）节点。
	// - FLOW_CHANGENET："031"，改变网络类型。
	// - FLOW_MODIFYINSTACEREADONLY："033"，只读策略变更。
	// - FLOW_MODIFYINSTANCEPARAMS："034"，修改实例参数。
	// - FLOW_MODIFYINSTANCEPASSWORDFREE："035"，设置免密。
	// - FLOW_SWITCHINSTANCEVIP："036"，实例VIP切换。
	// - FLOW_MODIFYINSTANCEACCOUNT："037"，实例帐号变更。
	// - FLOW_MODIFYINSTANCEBANDWIDTH："038"，实例带宽变更。
	// - FLOW_ENABLEINSTANCE_REPLICATE："039"，开启副本只读。
	// - FLOW_DISABLEINSTANCE_REPLICATE："040"，关闭副本只读。
	// - FLOW_UpgradeArch："041"，实例架构升级，主从升集群。
	// - FLOW_DowngradeArch： "042"，实例架构降级，集群降主从。
	// - FLOW_UpgradeVersion： "043"，版本升级。
	// - FLOW_MODIFYCONNECTIONCONFIG："044"，带宽连接数调整，
	// - FLOW_CLEARNETWORK："045"，更换网络，
	// - FLOW_REMOVE_BACKUP_FILE："046"，删除备份。
	// - FLOW_UPGRADE_SUPPORT_MULTI_AZ："047"，升级实例支持多可用区。
	// - FLOW_SHUTDOWN_MASTER："048"，模拟故障。
	// - FLOW_CHANGE_REPLICA_TO_MASTER："049"，手动提主。
	// - FLOW_CODE_ADD_REPLICATION_INSTANCE："050"，新增复制组。
	// - FLOW_OPEN_WAN："052"，开通外网。
	// - FLOW_CLOSE_WAN："053"，关闭外网FLOW_UPDATE_WAN："054"，更新外网。
	// - FLOW_CODE_DELETE_REPLICATION_INSTANCE："055"，解绑复制组。
	// - FLOW_CODE_CHANGE_MASTER_INSTANCE："056"，复制组实例切主。
	// - FLOW_CODE_CHANGE_INSTANCE_ROLE： "057"，更改复制组实例角色。
	// - FLOW_MIGRATE_NODE："058"，迁移节点。
	// - FLOW_SWITCH_NODE："059"，切换节点。
	// - FLOW_UPGRADE_SMALL_VERSION："060"，升级 Redi s版本。
	// - FLOW_UPGRADE_PROXY_VERSION："061"，升级 Proxy 版本。
	// - FLOW_MODIFY_INSTANCE_NETWORK： "062"，实例修改网络。
	// - FLOW_MIGRATE_PROXY_NODE："063"，迁移proxy节点。
	// - FLOW_MIGRATION_INSTANCE_ZONE："066"，实例可用区迁移中。
	// - FLOW_UPGRADE_INSTANCE_CACHE_AND_PROXY： "067"，实例版本升级中。
	// - FLOW_MODIFY_PROXY_NUM："069"，加（减）Proxy 节点。
	// - FLOW_MODIFYBACKUPMOD："070"，变更实例备份模式。
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// 任务执行的起始时间。格式如：2021-12-30 00:00:00。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 任务运行的终止时间。格式如：2021-12-30 20:59:35
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 该参数为内部使用，请忽略。
	TaskStatus []*int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务执行状态。
	// - 0：任务初始化。
	// - 1：执行中。
	// - 2：完成。
	// - 4：失败。
	Result []*int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 该字段已废弃，使用OperateUin代替，请忽略。
	OperatorUin []*int64 `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 操作者账号 ID，UIN。
	OperateUin []*string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ProjectIds")
	delete(f, "TaskTypes")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TaskStatus")
	delete(f, "Result")
	delete(f, "OperatorUin")
	delete(f, "OperateUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// 任务总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务详细信息。
	Tasks []*TaskInfoDetail `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskListResponseParams `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTendisSlowLogRequestParams struct {
	// 实例Id：crs-ngvou0i1
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间：2019-09-08 12:12:41
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间：2019-09-09 12:12:41
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢查询阈值（毫秒）
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// 页面大小：默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTendisSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id：crs-ngvou0i1
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间：2019-09-08 12:12:41
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间：2019-09-09 12:12:41
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢查询阈值（毫秒）
	MinQueryTime *int64 `json:"MinQueryTime,omitnil,omitempty" name:"MinQueryTime"`

	// 页面大小：默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTendisSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTendisSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "MinQueryTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTendisSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTendisSlowLogResponseParams struct {
	// 慢查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢查询详情
	TendisSlowLogDetail []*TendisSlowLogDetail `json:"TendisSlowLogDetail,omitnil,omitempty" name:"TendisSlowLogDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTendisSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTendisSlowLogResponseParams `json:"Response"`
}

func (r *DescribeTendisSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTendisSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPostpaidInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyPostpaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyPostpaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPostpaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPostpaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPostpaidInstanceResponseParams struct {
	// 任务Id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPostpaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPostpaidInstanceResponseParams `json:"Response"`
}

func (r *DestroyPostpaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPostpaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrepaidInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyPrepaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyPrepaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrepaidInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyPrepaidInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyPrepaidInstanceResponseParams struct {
	// 订单Id
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyPrepaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyPrepaidInstanceResponseParams `json:"Response"`
}

func (r *DestroyPrepaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyPrepaidInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableReplicaReadonlyRequestParams struct {
	// 实例序号ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DisableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// 实例序号ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DisableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableReplicaReadonlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableReplicaReadonlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableReplicaReadonlyResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *DisableReplicaReadonlyResponseParams `json:"Response"`
}

func (r *DisableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableReplicaReadonlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例ID列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例ID列表，一个或者多个实例 ID 组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableReplicaReadonlyRequestParams struct {
	// 实例序号ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账号路由策略：填写master或者replication，表示路由主节点，从节点；不填路由策略默认为写主节点，读从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`
}

type EnableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest
	
	// 实例序号ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 账号路由策略：填写master或者replication，表示路由主节点，从节点；不填路由策略默认为写主节点，读从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`
}

func (r *EnableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableReplicaReadonlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadonlyPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableReplicaReadonlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableReplicaReadonlyResponseParams struct {
	// 错误：ERROR，正确OK（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *EnableReplicaReadonlyResponseParams `json:"Response"`
}

func (r *EnableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableReplicaReadonlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Groups struct {
	// 用户 APPID。APPID是与账号ID有唯一对应关系的应用 ID，部分腾讯云产品会使用此 APPID。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域ID 。
	// - 1：广州 
	// - 4：上海 
	// - 5：中国香港 
	// - 7：上海金融 
	// - 8：北京 
	// - 9：新加坡
	// - 11：深圳金融
	// - 15：美西（硅谷）
	// - 16：成都 
	// - 17：德国 
	// - 18：韩国 
	// - 19：重庆 
	// - 21：印度 
	// - 22：美东（弗吉尼亚）
	// - 23：泰国 
	// - 25：日本
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 复制组 ID。格式如：crs-rpl-deind****。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 复制组名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 复制组状态。
	// - 37：绑定复制组中。
	// - 38：复制组重连中。
	// - 51：解绑复制组中。
	// - 52：复制组实例切主中。
	// - 53：角色变更中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 复制组数量。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 复制组中的实例信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*Instances `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 备注信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type HotKeyInfo struct {
	// 热 Key 的名称。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Key 类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 某段时间内热 Key 的访问次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type Inbound struct {
	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type InquiryPriceCreateInstanceRequestParams struct {
	// 实例类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 内存容量，单位为MB， 数值需为1024的整数倍，具体规格以 [查询产品售卖规格](https://cloud.tencent.com/document/api/239/30600) 返回的规格为准。
	// TypeId为标准架构时，MemSize是实例总内存容量；TypeId为集群架构时，MemSize是单分片内存容量。
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 实例数量，单次购买实例数量以 [查询产品售卖规格](https://cloud.tencent.com/document/api/239/30600) 返回的规格为准。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 购买时长，在创建包年包月实例的时候需要填写，按量计费实例填1即可，单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 付费方式。
	// - 0：按量计费。
	// - 1：包年包月。
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例所属的可用区ID，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例分片数量。
	// Redis2.8标准架构、CKV标准架构、Redis4.0标准架构无需填写。
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 实例副本数量。
	// Redis2.8标准架构、CKV标准架构无需填写。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 是否支持副本只读。Redis2.8标准架构、CKV标准架构无需填写。
	// - true：无需支持副本只读。
	// - false：需支持。
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// 实例所属的可用区名称，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 部署方式。
	// - local：本地盘版，默认为 local。
	// - cloud：云盘版。
	// - cdc：独享集群版。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	TypeId *uint64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 内存容量，单位为MB， 数值需为1024的整数倍，具体规格以 [查询产品售卖规格](https://cloud.tencent.com/document/api/239/30600) 返回的规格为准。
	// TypeId为标准架构时，MemSize是实例总内存容量；TypeId为集群架构时，MemSize是单分片内存容量。
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 实例数量，单次购买实例数量以 [查询产品售卖规格](https://cloud.tencent.com/document/api/239/30600) 返回的规格为准。
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 购买时长，在创建包年包月实例的时候需要填写，按量计费实例填1即可，单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 付费方式。
	// - 0：按量计费。
	// - 1：包年包月。
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例所属的可用区ID，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 实例分片数量。
	// Redis2.8标准架构、CKV标准架构、Redis4.0标准架构无需填写。
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 实例副本数量。
	// Redis2.8标准架构、CKV标准架构无需填写。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 是否支持副本只读。Redis2.8标准架构、CKV标准架构无需填写。
	// - true：无需支持副本只读。
	// - false：需支持。
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitnil,omitempty" name:"ReplicasReadonly"`

	// 实例所属的可用区名称，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 部署方式。
	// - local：本地盘版，默认为 local。
	// - cloud：云盘版。
	// - cdc：独享集群版。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TypeId")
	delete(f, "MemSize")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "BillingMode")
	delete(f, "ZoneId")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "ReplicasReadonly")
	delete(f, "ZoneName")
	delete(f, "ProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateInstanceResponseParams struct {
	// 价格，单位：分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceRequestParams struct {
	// 包年包月实例的购买时长，单位：月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 包年包月实例的购买时长，单位：月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *InquiryPriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewInstanceResponseParams struct {
	// 价格，单位：分。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstanceRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片大小 单位 MB
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 分片数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 副本数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`
}

type InquiryPriceUpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分片大小 单位 MB
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 分片数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 副本数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`
}

func (r *InquiryPriceUpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeInstanceResponseParams struct {
	// 价格，单位：分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceClusterNode struct {
	// 节点组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例运行时节点 ID。
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// 集群角色。
	// - 0：master。
	// - 1：slave。
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点状态。
	// - 0：readwrite,。
	// - 1：read。
	// - 2：backup。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 服务状态。
	// 0-down。
	// 1-on
	Connected *int64 `json:"Connected,omitnil,omitempty" name:"Connected"`

	// 节点创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 节点下线时间。
	DownTime *string `json:"DownTime,omitnil,omitempty" name:"DownTime"`

	// 节点 Slot 分布区间。
	Slots *string `json:"Slots,omitnil,omitempty" name:"Slots"`

	// 节点 Key分布。
	Keys *int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 节点 QPS。分片节点每秒执行次数。单位：次/秒。
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 节点 QPS 倾斜度。
	QpsSlope *float64 `json:"QpsSlope,omitnil,omitempty" name:"QpsSlope"`

	// 节点存储。
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 节点存储倾斜度。
	StorageSlope *float64 `json:"StorageSlope,omitnil,omitempty" name:"StorageSlope"`
}

type InstanceClusterShard struct {
	// 分片节点名称。
	ShardName *string `json:"ShardName,omitnil,omitempty" name:"ShardName"`

	// 分片节点序号。
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 分片节点的角色。
	// - 0：主节点。
	// - 1：副本节点。
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// Key数量。
	Keys *int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Slot信息。
	Slots *string `json:"Slots,omitnil,omitempty" name:"Slots"`

	// 已使用容量。
	Storage *int64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 容量倾斜率。
	StorageSlope *float64 `json:"StorageSlope,omitnil,omitempty" name:"StorageSlope"`

	// 该字段因拼写不规范问题，建议使用RunId取代。含义：实例运行时节点 ID。
	Runid *string `json:"Runid,omitnil,omitempty" name:"Runid"`

	// 实例运行时节点 ID。
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// 服务状态。
	// - 0：down。
	// - 1：on。
	Connected *int64 `json:"Connected,omitnil,omitempty" name:"Connected"`
}

type InstanceEnumParam struct {
	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数类型，例如：Enum。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 参数值修改后是否需要重启。
	// - true：需要。
	// - false：不需要。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参数当前运行值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数说明。
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数可取的值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数修改状态。
	// - 1: 修改中。
	// - 2：修改完成。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceIntegerParam struct {
	// 参数名
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数类型：integer
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数最小值
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数最大值
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 参数单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type InstanceMultiParam struct {
	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数类型。例如：multi。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 参数修改后是否需要重启。
	// - true：需要。
	// - false：不需要。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 当前运行参数值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数说明。
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数枚举值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`

	// 参数修改的状态。
	// - 1：修改中。
	// - 2：修改完成。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceNode struct {
	// 实例 ID。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 节点详细信息。
	InstanceClusterNode []*InstanceClusterNode `json:"InstanceClusterNode,omitnil,omitempty" name:"InstanceClusterNode"`
}

type InstanceParam struct {
	// 设置参数的名称。例如timeout。当前支持自定义的参数，请参见<a href="https://cloud.tencent.com/document/product/239/49925">参数配置</a>。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 设置参数名称对应的运行值。例如timeout对应运行值可设置为120， 单位为秒（s）。指当客户端连接闲置时间达到120 s时，将关闭连接。更多参数取值信息，请参见<a href="https://cloud.tencent.com/document/product/239/49925">参数配置</a>。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InstanceParamHistory struct {
	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数修改之前的值。
	PreValue *string `json:"PreValue,omitnil,omitempty" name:"PreValue"`

	// 参数修改之后的值。
	NewValue *string `json:"NewValue,omitnil,omitempty" name:"NewValue"`

	// 参数配置状态。
	// - 1：参数配置修改中。
	// - 2：参数配置修改成功。
	// - 3：参数配置修改失败。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type InstanceProxySlowlogDetail struct {
	// 慢查询耗时时长。单位：毫秒。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 客户端地址。
	Client *string `json:"Client,omitnil,omitempty" name:"Client"`

	// 慢查询的命令。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 慢查询详细命令行信息。
	CommandLine *string `json:"CommandLine,omitnil,omitempty" name:"CommandLine"`

	// 执行时间。
	ExecuteTime *string `json:"ExecuteTime,omitnil,omitempty" name:"ExecuteTime"`
}

type InstanceSecurityGroupDetail struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 安全组信息，包括：安全组 ID、安全组名称、安全组出入站规则。
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitnil,omitempty" name:"SecurityGroupDetails"`
}

type InstanceSet struct {
	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户APPID。APPID是与账号ID有唯一对应关系的应用 ID，部分腾讯云产品会使用此 APPID。
	Appid *int64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// 项目 ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域 ID。<ul><li>1：广州。</li><li>4：上海。</li><li>5：中国香港。</li><li>7：上海金融。</li> <li>8：北京。</li> <li>9：新加坡。</li> <li>11：深圳金融。</li> <li>15：美西（硅谷）。</li><li>16：成都。</li><li>17：法兰克福。</li><li>18：首尔。</li><li>19：重庆。</li><li>21：孟买。</li><li>22：美东（弗吉尼亚）。</li><li>23：曼谷。</li><li>25：东京。</li></ul>
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域 ID。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// vpc网络 ID，例如75101。
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc网络下子网ID，如：46315。
	SubnetId *int64 `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例当前状态。<ul><li>0：待初始化。</li><li>1：实例在流程中。</li><li>2：实例运行中。</li><li>-2：实例已隔离。</li><li>-3：实例待删除。</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例 VIP。
	WanIp *string `json:"WanIp,omitnil,omitempty" name:"WanIp"`

	// 实例端口号。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例创建时间。格式如：2020-01-15 10:20:00。
	Createtime *string `json:"Createtime,omitnil,omitempty" name:"Createtime"`

	// 实例内存容量大小。单位：MB，1MB=1024KB。
	Size *float64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 该字段已废弃。请使用腾讯云可观测平台API 接口 [GetMonitorData](https://cloud.tencent.com/document/product/248/31014) 获取实例已使用的内存容量。
	SizeUsed *float64 `json:"SizeUsed,omitnil,omitempty" name:"SizeUsed"`

	// 实例类型。
	// - 2：Redis 2.8 内存版（标准架构）。
	// - 3：CKV 3.2 内存版（标准架构）。
	// - 4：CKV 3.2 内存版（集群架构）。
	// - 5：Redis 2.8 内存版（单机）。
	// - 6：Redis 4.0 内存版（标准架构）。
	// - 7：Redis 4.0 内存版（集群架构）。
	// - 8：Redis 5.0 内存版（标准架构）。
	// - 9：Redis 5.0 内存版（集群架构）。
	// - 15：Redis 6.2 内存版（标准架构）。
	// - 16：Redis 6.2 内存版（集群架构）。
	// - 17：Redis 7.0 内存版（标准架构）。
	// - 18：Redis 7.0 内存版（集群架构）。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例是否设置自动续费标识。<ul><li>1：设置自动续费。</li><li>0：未设置自动续费。</li></ul>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 包年包月计费实例到期的时间。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 引擎：社区版Redis、腾讯云CKV。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 产品类型。<ul><li>standalone：标准版。</li><li>cluster ：集群版。</li></ul>
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// vpc网络id，例如vpc-fk33jsf43kgv。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// vpc网络下子网id，例如：subnet-fd3j6l35mm0。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 计费模式。<ul><li>0：按量计费。</li><li>1：包年包月。</li></ul>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// 实例运行状态描述：如”实例运行中“。
	InstanceTitle *string `json:"InstanceTitle,omitnil,omitempty" name:"InstanceTitle"`

	// 已隔离实例默认下线时间。按量计费实例隔离后默认两小时后下线，包年包月默认7天后下线。格式如：2020-02-15 10:20:00。
	OfflineTime *string `json:"OfflineTime,omitnil,omitempty" name:"OfflineTime"`

	// 流程中的实例，返回子状态。
	SubStatus *int64 `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// 反亲和性标签。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例节点信息。
	InstanceNode []*InstanceNode `json:"InstanceNode,omitnil,omitempty" name:"InstanceNode"`

	// 分片大小。
	RedisShardSize *int64 `json:"RedisShardSize,omitnil,omitempty" name:"RedisShardSize"`

	// 分片数量。
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 副本数量。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 计费 ID。
	PriceId *int64 `json:"PriceId,omitnil,omitempty" name:"PriceId"`

	// 实例隔离开始的时间。
	CloseTime *string `json:"CloseTime,omitnil,omitempty" name:"CloseTime"`

	// 从节点读取权重。
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitnil,omitempty" name:"SlaveReadWeight"`

	// 实例关联的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 项目名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 是否为免密实例。<ul><li>true：免密实例。</li><li>false：非免密实例。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`

	// 客户端连接数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`

	// DTS状态（内部参数，用户可忽略）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsStatus *int64 `json:"DtsStatus,omitnil,omitempty" name:"DtsStatus"`

	// 分片带宽上限，单位MB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetLimit *int64 `json:"NetLimit,omitnil,omitempty" name:"NetLimit"`

	// 免密实例标识（内部参数，用户可忽略）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PasswordFree *int64 `json:"PasswordFree,omitnil,omitempty" name:"PasswordFree"`

	// 该参数存在命名不规范问题，建议用参数IPv6取代。内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// 内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6 *string `json:"IPv6,omitnil,omitempty" name:"IPv6"`

	// 实例只读标识（内部参数，用户可忽略）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *int64 `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`

	// 内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitnil,omitempty" name:"RemainBandwidthDuration"`

	// Redis实例请忽略该参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 监控版本。<ul><li>1m：1分钟粒度监控。目前该监控粒度已下线，具体信息，请参见[云数据库 Redis 1分钟粒度下线公告](https://cloud.tencent.com/document/product/239/80653)。</li><li>5s：5秒粒度监控。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorVersion *string `json:"MonitorVersion,omitnil,omitempty" name:"MonitorVersion"`

	// 客户端最大连接数可设置的最小值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimitMin *int64 `json:"ClientLimitMin,omitnil,omitempty" name:"ClientLimitMin"`

	// 客户端最大连接数可设置的最大值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimitMax *int64 `json:"ClientLimitMax,omitnil,omitempty" name:"ClientLimitMax"`

	// 实例的节点详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`

	// 实例所在的地域信息，比如ap-guangzhou。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 外网地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanAddress *string `json:"WanAddress,omitnil,omitempty" name:"WanAddress"`

	// 北极星服务地址，内部使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolarisServer *string `json:"PolarisServer,omitnil,omitempty" name:"PolarisServer"`

	// CDC Redis集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedisClusterId *string `json:"RedisClusterId,omitnil,omitempty" name:"RedisClusterId"`

	// CDC 集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil,omitempty" name:"DedicatedClusterId"`

	// 产品版本。<ul><li>local：本地盘。</li><li>cloud：云盘版。</li><li>cdc：CDC 集群版本。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 实例当前Proxy版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// 实例当前Cache小版本。如果实例加入全球复制组，显示全球复制的内核版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil,omitempty" name:"CurrentRedisVersion"`

	// 实例可升级Proxy版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil,omitempty" name:"UpgradeProxyVersion"`

	// 实例可升级Cache小版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil,omitempty" name:"UpgradeRedisVersion"`

	// 备份模式：- SecondLevelBackup   秒级备份- NormalLevelBackup    普通备份
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupMode *string `json:"BackupMode,omitnil,omitempty" name:"BackupMode"`
}

type InstanceSlowlogDetail struct {
	// 慢查询耗时
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 客户端地址
	Client *string `json:"Client,omitnil,omitempty" name:"Client"`

	// 命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 详细命令行信息
	CommandLine *string `json:"CommandLine,omitnil,omitempty" name:"CommandLine"`

	// 执行时间
	ExecuteTime *string `json:"ExecuteTime,omitnil,omitempty" name:"ExecuteTime"`

	// 节点ID
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`
}

type InstanceTagInfo struct {
	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type InstanceTextParam struct {
	// 参数名称。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数类型。例如：text。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 参数修改后是否需要重启。
	// - true：需要。
	// - false：不需要。
	NeedRestart *string `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 参数默认值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 参数当前运行值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 参数说明。
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 参数可取值。
	TextValue []*string `json:"TextValue,omitnil,omitempty" name:"TextValue"`

	// 参数修改状态。
	// - 1: 修改中。
	// - 2：修改完成。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type Instances struct {
	// 用户APPID。APPID是与账号ID有唯一对应关系的应用 ID，部分腾讯云产品会使用此 APPID。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域ID。<ul><li>1：广州。</li><li>4：上海。</li><li> 5：香港。</li>  <li> 7：上海金融。</li> <li> 8：北京。</li> <li> 9：新加坡。</li> <li> 11：深圳金融。</li> <li> 15：美西（硅谷）。</li> </ul>
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域 ID。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 副本数量。
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 分片数量。
	RedisShardNum *int64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 分片内存大小。
	RedisShardSize *int64 `json:"RedisShardSize,omitnil,omitempty" name:"RedisShardSize"`

	// 实例的磁盘大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 引擎：社区版Redis、腾讯云CKV。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 实例读写权限。<ul><li>rw：可读写。</li><li>r：只读。</li></ul>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 实例 VIP 地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 该参数存在命名不规范问题，建议用参数IPv6取代。内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip6 *string `json:"Vip6,omitnil,omitempty" name:"Vip6"`

	// 内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6 *string `json:"IPv6,omitnil,omitempty" name:"IPv6"`

	// VPC 网络ID，如：75101。
	VpcID *int64 `json:"VpcID,omitnil,omitempty" name:"VpcID"`

	// 实例端口。
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`

	// 实例状态。<ul><li>0：待初始化。</li><li>1：流程中。</li><li>2：运行中。</li><li>-2：已隔离。</li><li>-3：待删除。</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 仓库ID。
	GrocerySysId *int64 `json:"GrocerySysId,omitnil,omitempty" name:"GrocerySysId"`

	// 实例类型。
	// - 2：Redis 2.8内存版（标准架构）。
	// - 3：CKV 3.2内存版（标准架构）。
	// - 4：CKV 3.2内存版（集群架构）。
	// - 5：Redis 2.8内存版（单机）。
	// - 6：Redis 4.0内存版（标准架构）。
	// - 7：Redis 4.0内存版（集群架构）。
	// - 8：Redis 5.0内存版（标准架构）。
	// - 9：Redis 5.0内存版（集群架构）。
	// - 15：Redis 6.2内存版（标准架构）。
	// - 16：Redis 6.2内存版（集群架构）。
	ProductType *int64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 实例加入复制组的时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 复制组中实例更新的时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type KillMasterGroupRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该参数用于配置指定实例的访问密码。若为免密认证，该参数则无需配置。密码复杂度要求如下所示。
	// - 长度8-30位,推荐使用12位以上的密码
	// - 不能以"/"开头
	// - 至少包含小写字母a-z、大写字母A-Z、数字0-9及其 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/中的两项。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 分片集群的分片 ID。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

type KillMasterGroupRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该参数用于配置指定实例的访问密码。若为免密认证，该参数则无需配置。密码复杂度要求如下所示。
	// - 长度8-30位,推荐使用12位以上的密码
	// - 不能以"/"开头
	// - 至少包含小写字母a-z、大写字母A-Z、数字0-9及其 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/中的两项。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 分片集群的分片 ID。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

func (r *KillMasterGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMasterGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "ShardIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMasterGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillMasterGroupResponseParams struct {
	// 异步任务ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillMasterGroupResponse struct {
	*tchttp.BaseResponse
	Response *KillMasterGroupResponseParams `json:"Response"`
}

func (r *KillMasterGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMasterGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogDeliveryInfo struct {
	// 日志投递开启状态，开启：true，关闭：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 日志集ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 日志主题ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`
}

// Predefined struct for user
type ManualBackupInstanceRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 手动备份任务的备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 备份数据的保存天数。
	// - 单位：天；默认值为7天；取值范围：[0.1825]。如果超过 7天，请[提交工单](https://console.cloud.tencent.com/workorder/category)申请。
	// - 如果不配置该参数，默认与自动备份的保留时间一致。
	// - 如果未设置自动备份，默认为7天。
	StorageDays *int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`
}

type ManualBackupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 手动备份任务的备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 备份数据的保存天数。
	// - 单位：天；默认值为7天；取值范围：[0.1825]。如果超过 7天，请[提交工单](https://console.cloud.tencent.com/workorder/category)申请。
	// - 如果不配置该参数，默认与自动备份的保留时间一致。
	// - 如果未设置自动备份，默认为7天。
	StorageDays *int64 `json:"StorageDays,omitnil,omitempty" name:"StorageDays"`
}

func (r *ManualBackupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualBackupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Remark")
	delete(f, "StorageDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManualBackupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManualBackupInstanceResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManualBackupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ManualBackupInstanceResponseParams `json:"Response"`
}

func (r *ManualBackupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualBackupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例旧密码。
	OldPassword *string `json:"OldPassword,omitnil,omitempty" name:"OldPassword"`

	// 实例新密码。密码复杂度要求如下：
	// - 长度8 - 30位, 推荐使用12位以上的密码。
	// - 不能以"/"开头。
	// - 至少包含小写字母a - z、大写字母A - Z、数字0 - 9、特殊字符 ()~!@#$%^&*-+=_|{}[]:;<>,.?/中的两项。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ModfiyInstancePasswordRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例旧密码。
	OldPassword *string `json:"OldPassword,omitnil,omitempty" name:"OldPassword"`

	// 实例新密码。密码复杂度要求如下：
	// - 长度8 - 30位, 推荐使用12位以上的密码。
	// - 不能以"/"开头。
	// - 至少包含小写字母a - z、大写字母A - Z、数字0 - 9、特殊字符 ()~!@#$%^&*-+=_|{}[]:;<>,.?/中的两项。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *ModfiyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldPassword")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModfiyInstancePasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModfiyInstancePasswordResponseParams struct {
	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModfiyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModfiyInstancePasswordResponseParams `json:"Response"`
}

func (r *ModfiyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModfiyInstancePasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置自动备份周期。可设置为Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。该参数暂不支持修改。
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 备份时间段。可设置为每个整点。格式如：00:00-01:00, 01:00-02:00...... 23:00-00:00。
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// 自动备份类型。目前仅能配置为：1 ，指定时备份。
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设置自动备份周期。可设置为Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。该参数暂不支持修改。
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 备份时间段。可设置为每个整点。格式如：00:00-01:00, 01:00-02:00...... 23:00-00:00。
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// 自动备份类型。目前仅能配置为：1 ，指定时备份。
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`
}

func (r *ModifyAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "WeekDays")
	delete(f, "TimePeriod")
	delete(f, "AutoBackupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigResponseParams struct {
	// 自动备份类型。目前仅能配置为：1 ，指定时备份。
	AutoBackupType *int64 `json:"AutoBackupType,omitnil,omitempty" name:"AutoBackupType"`

	// 自动备份周期。取值为：Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
	WeekDays []*string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 自动定时备份时间段。格式如：00:00-01:00, 01:00-02:00...... 23:00-00:00。
	TimePeriod *string `json:"TimePeriod,omitnil,omitempty" name:"TimePeriod"`

	// 全量备份文件保存天数,单位：天。
	BackupStorageDays *int64 `json:"BackupStorageDays,omitnil,omitempty" name:"BackupStorageDays"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoBackupConfigResponseParams `json:"Response"`
}

func (r *ModifyAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoBackupConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionRequestParams struct {
	// 下载备份文件的网络限制类型：
	// 
	// - NoLimit：不限制，腾讯云内外网均可以下载备份文件。
	// -  LimitOnlyIntranet：仅腾讯云自动分配的内网地址可下载备份文件。
	// - Customize：指用户自定义的私有网络可下载备份文件。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持输入 In，表示自定义的**LimitVpc**可以下载备份文件。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// 标识自定义的 LimitIp 地址是否可下载备份文件。
	// 
	// - In: 自定义的 IP 地址可以下载。
	// - NotIn: 自定义的 IP 不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 自定义的可下载备份文件的 VPC ID。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 自定义的可下载备份文件的 VPC IP 地址。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

type ModifyBackupDownloadRestrictionRequest struct {
	*tchttp.BaseRequest
	
	// 下载备份文件的网络限制类型：
	// 
	// - NoLimit：不限制，腾讯云内外网均可以下载备份文件。
	// -  LimitOnlyIntranet：仅腾讯云自动分配的内网地址可下载备份文件。
	// - Customize：指用户自定义的私有网络可下载备份文件。
	LimitType *string `json:"LimitType,omitnil,omitempty" name:"LimitType"`

	// 该参数仅支持输入 In，表示自定义的**LimitVpc**可以下载备份文件。
	VpcComparisonSymbol *string `json:"VpcComparisonSymbol,omitnil,omitempty" name:"VpcComparisonSymbol"`

	// 标识自定义的 LimitIp 地址是否可下载备份文件。
	// 
	// - In: 自定义的 IP 地址可以下载。
	// - NotIn: 自定义的 IP 不可以下载。
	IpComparisonSymbol *string `json:"IpComparisonSymbol,omitnil,omitempty" name:"IpComparisonSymbol"`

	// 自定义的可下载备份文件的 VPC ID。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitVpc []*BackupLimitVpcItem `json:"LimitVpc,omitnil,omitempty" name:"LimitVpc"`

	// 自定义的可下载备份文件的 VPC IP 地址。当参数**LimitType**为**Customize **时，需配置该参数。
	LimitIp []*string `json:"LimitIp,omitnil,omitempty" name:"LimitIp"`
}

func (r *ModifyBackupDownloadRestrictionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitType")
	delete(f, "VpcComparisonSymbol")
	delete(f, "IpComparisonSymbol")
	delete(f, "LimitVpc")
	delete(f, "LimitIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupDownloadRestrictionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupDownloadRestrictionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBackupDownloadRestrictionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupDownloadRestrictionResponseParams `json:"Response"`
}

func (r *ModifyBackupDownloadRestrictionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupDownloadRestrictionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectionConfigRequestParams struct {
	// 实例的ID，长度在12-36之间。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 附加带宽，大于0，单位MB。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 单分片的总连接数。
	// 未开启副本只读时，下限为10000，上限为40000；
	// 开启副本只读时，下限为10000，上限为10000×(只读副本数+3)。
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`
}

type ModifyConnectionConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例的ID，长度在12-36之间。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 附加带宽，大于0，单位MB。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 单分片的总连接数。
	// 未开启副本只读时，下限为10000，上限为40000；
	// 开启副本只读时，下限为10000，上限为10000×(只读副本数+3)。
	ClientLimit *int64 `json:"ClientLimit,omitnil,omitempty" name:"ClientLimit"`
}

func (r *ModifyConnectionConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConnectionConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "ClientLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConnectionConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConnectionConfigResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConnectionConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConnectionConfigResponseParams `json:"Response"`
}

func (r *ModifyConnectionConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConnectionConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例 ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例 ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "SecurityGroupIds")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAccountRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子账号名称，如果要修改主账号，填root
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 子账号密码
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 子账号描述信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 路由策略：填写master或者replication，表示主节点或者从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// 子账号读写策略：填写r、w、rw，表示只读，只写，读写策略
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// true表示将主账号切换为免密账号，这里只适用于主账号，子账号不可免密。
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`
}

type ModifyInstanceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子账号名称，如果要修改主账号，填root
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 子账号密码
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// 子账号描述信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 路由策略：填写master或者replication，表示主节点或者从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitnil,omitempty" name:"ReadonlyPolicy"`

	// 子账号读写策略：填写r、w、rw，表示只读，只写，读写策略
	Privilege *string `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// true表示将主账号切换为免密账号，这里只适用于主账号，子账号不可免密。
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`
}

func (r *ModifyInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AccountName")
	delete(f, "AccountPassword")
	delete(f, "Remark")
	delete(f, "ReadonlyPolicy")
	delete(f, "Privilege")
	delete(f, "NoAuth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAccountResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAccountResponseParams `json:"Response"`
}

func (r *ModifyInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAvailabilityZonesRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****，请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 切换时间。
	// - 1：维护时间窗切换。
	// - 2：立即切换。
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// 实例的节点信息，包含节点 ID、节点类型、节点可用区 ID等。具体信息，请参见[RedisNodeInfo ](https://cloud.tencent.com/document/product/239/20022)。
	// 单可用区实例无需传NodeId，多可用区实例NodeId必传
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`
}

type ModifyInstanceAvailabilityZonesRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****，请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 切换时间。
	// - 1：维护时间窗切换。
	// - 2：立即切换。
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// 实例的节点信息，包含节点 ID、节点类型、节点可用区 ID等。具体信息，请参见[RedisNodeInfo ](https://cloud.tencent.com/document/product/239/20022)。
	// 单可用区实例无需传NodeId，多可用区实例NodeId必传
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`
}

func (r *ModifyInstanceAvailabilityZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SwitchOption")
	delete(f, "NodeSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceAvailabilityZonesResponseParams struct {
	// 任务ID。	
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceAvailabilityZonesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceAvailabilityZonesResponseParams `json:"Response"`
}

func (r *ModifyInstanceAvailabilityZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEventRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件 ID。请通过 DescribeInstanceEvents 获取需修改的事件 ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 修改事件执行的计划开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 修改事件计划执行的结束时间。开始时间配置之后，结束时间只能选择在开始时间之后的 30 分钟、1 小时、1.5 小时、2 小时和 3 小时之内。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 修改事件执行计划的开始日期。
	ExecutionDate *string `json:"ExecutionDate,omitnil,omitempty" name:"ExecutionDate"`

	// 修改事件的运行状态。该参数当前仅支持设置为 **Canceled**， 即取消执行当前事件。可通过 DescribeInstanceEvents 接口查询当前事件的运行状态与事件级别。
	// - 事件级别为Critical（关键）或 High（重要）类事件不支持取消。即严重的事件必须执行，不可取消。
	// - 仅运行状态为 Waiting （待执行的事件）的事件，才能执行取消操作。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyInstanceEventRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件 ID。请通过 DescribeInstanceEvents 获取需修改的事件 ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 修改事件执行的计划开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 修改事件计划执行的结束时间。开始时间配置之后，结束时间只能选择在开始时间之后的 30 分钟、1 小时、1.5 小时、2 小时和 3 小时之内。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 修改事件执行计划的开始日期。
	ExecutionDate *string `json:"ExecutionDate,omitnil,omitempty" name:"ExecutionDate"`

	// 修改事件的运行状态。该参数当前仅支持设置为 **Canceled**， 即取消执行当前事件。可通过 DescribeInstanceEvents 接口查询当前事件的运行状态与事件级别。
	// - 事件级别为Critical（关键）或 High（重要）类事件不支持取消。即严重的事件必须执行，不可取消。
	// - 仅运行状态为 Waiting （待执行的事件）的事件，才能执行取消操作。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyInstanceEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ExecutionDate")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEventResponseParams struct {
	// 事件 ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceEventResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceEventResponseParams `json:"Response"`
}

func (r *ModifyInstanceEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceLogDeliveryRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型。当前仅支持设置为slowlog，指慢查询日志。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 日志投递开启状态。
	// - true：开启。
	// - false：关闭。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 投递的日志集ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 投递的日志主题ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集名称。若**LogsetId**未指定具体的日志集ID，请配置该参数，设置日志集名称，系统会以设置的日志集名称自动创建新的日志集。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称，TopicId为空时必传，会自动创建新的日志主题。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集所在地域，不传默认使用实例所在地域。
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// 日志存储时间，默认为30天，可选范围1-3600天。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 创建日志主题时，是否创建索引。
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`
}

type ModifyInstanceLogDeliveryRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型。当前仅支持设置为slowlog，指慢查询日志。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 日志投递开启状态。
	// - true：开启。
	// - false：关闭。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 投递的日志集ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 投递的日志主题ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志集名称。若**LogsetId**未指定具体的日志集ID，请配置该参数，设置日志集名称，系统会以设置的日志集名称自动创建新的日志集。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志主题名称，TopicId为空时必传，会自动创建新的日志主题。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志集所在地域，不传默认使用实例所在地域。
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// 日志存储时间，默认为30天，可选范围1-3600天。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 创建日志主题时，是否创建索引。
	CreateIndex *bool `json:"CreateIndex,omitnil,omitempty" name:"CreateIndex"`
}

func (r *ModifyInstanceLogDeliveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceLogDeliveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "Enabled")
	delete(f, "LogsetId")
	delete(f, "TopicId")
	delete(f, "LogsetName")
	delete(f, "TopicName")
	delete(f, "LogRegion")
	delete(f, "Period")
	delete(f, "CreateIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceLogDeliveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceLogDeliveryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceLogDeliveryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceLogDeliveryResponseParams `json:"Response"`
}

func (r *ModifyInstanceLogDeliveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceLogDeliveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例修改的参数列表。
	InstanceParams []*InstanceParam `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例修改的参数列表。
	InstanceParams []*InstanceParam `json:"InstanceParams,omitnil,omitempty" name:"InstanceParams"`
}

func (r *ModifyInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamsResponseParams struct {
	// 说明修改参数配置是否成功。<br><li>true：指修改成功；</li><li>false：指修改失败。</li>
	Changed *bool `json:"Changed,omitnil,omitempty" name:"Changed"`

	// 任务ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamsResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceReadOnlyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例输入模式，0：读写 1：只读
	InputMode *string `json:"InputMode,omitnil,omitempty" name:"InputMode"`
}

type ModifyInstanceReadOnlyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例输入模式，0：读写 1：只读
	InputMode *string `json:"InputMode,omitnil,omitempty" name:"InputMode"`
}

func (r *ModifyInstanceReadOnlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InputMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceReadOnlyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceReadOnlyResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceReadOnlyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceReadOnlyResponseParams `json:"Response"`
}

func (r *ModifyInstanceReadOnlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceReadOnlyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 修改实例操作，如填写：rename-表示实例重命名；modifyProject-修改实例所属项目；modifyAutoRenew-修改实例续费标记
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 实例Id，每次请求的实例的上限为10。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例的新名称
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自动续费标识。0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenews []*int64 `json:"AutoRenews,omitnil,omitempty" name:"AutoRenews"`

	// 已经废弃
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 已经废弃
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 已经废弃
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 修改实例操作，如填写：rename-表示实例重命名；modifyProject-修改实例所属项目；modifyAutoRenew-修改实例续费标记
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 实例Id，每次请求的实例的上限为10。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例的新名称
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 自动续费标识。0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenews []*int64 `json:"AutoRenews,omitnil,omitempty" name:"AutoRenews"`

	// 已经废弃
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 已经废弃
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 已经废弃
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operation")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "ProjectId")
	delete(f, "AutoRenews")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "AutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗起始时间，如：17:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 维护时间窗结束时间，如：19:00
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维护时间窗起始时间，如：17:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 维护时间窗结束时间，如：19:00
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ModifyMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintenanceWindowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceWindowResponseParams struct {
	// 修改状态：success 或者 failed
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintenanceWindowResponseParams `json:"Response"`
}

func (r *ModifyMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceWindowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指预修改网络的类别，包括：
	// - changeVip：指切换私有网络，包含其内网IPv4地址及端口。
	// - changeVpc：指切换私有网络所属子网。
	// - changeBaseToVpc：指基础网络切换为私有网络。
	// - changeVPort：指仅修改实例网络端口。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 指实例私有网络内网 IPv4 地址。当**Operation**为**changeVip**时，需配置该参数。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 指修改后的私有网络 ID，当**Operation**为**changeVpc**或**changeBaseToVpc**时，需配置该参数。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 指修改后的私有网络所属子网 ID，当**Operation**为**changeVpc**或**changeBaseToVpc**时，需配置该参数。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 原内网 IPv4 地址保留时长。
	// - 单位：天。
	// - 取值范围：0、1、2、3、7、15。
	// 
	// **说明**：设置原地址保留时长需最新版SDK，否则原地址将立即释放，查看SDK版本，请参见 [SDK中心](https://cloud.tencent.com/document/sdk)。
	Recycle *int64 `json:"Recycle,omitnil,omitempty" name:"Recycle"`

	// 指修改后的网络端口。当**Operation**为**changeVPort**或**changeVip**时，需配置该参数。取值范围为[1024,65535]。
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指预修改网络的类别，包括：
	// - changeVip：指切换私有网络，包含其内网IPv4地址及端口。
	// - changeVpc：指切换私有网络所属子网。
	// - changeBaseToVpc：指基础网络切换为私有网络。
	// - changeVPort：指仅修改实例网络端口。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 指实例私有网络内网 IPv4 地址。当**Operation**为**changeVip**时，需配置该参数。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 指修改后的私有网络 ID，当**Operation**为**changeVpc**或**changeBaseToVpc**时，需配置该参数。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 指修改后的私有网络所属子网 ID，当**Operation**为**changeVpc**或**changeBaseToVpc**时，需配置该参数。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 原内网 IPv4 地址保留时长。
	// - 单位：天。
	// - 取值范围：0、1、2、3、7、15。
	// 
	// **说明**：设置原地址保留时长需最新版SDK，否则原地址将立即释放，查看SDK版本，请参见 [SDK中心](https://cloud.tencent.com/document/sdk)。
	Recycle *int64 `json:"Recycle,omitnil,omitempty" name:"Recycle"`

	// 指修改后的网络端口。当**Operation**为**changeVPort**或**changeVip**时，需配置该参数。取值范围为[1024,65535]。
	VPort *int64 `json:"VPort,omitnil,omitempty" name:"VPort"`
}

func (r *ModifyNetworkConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Operation")
	delete(f, "Vip")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Recycle")
	delete(f, "VPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigResponseParams struct {
	// 执行状态，请忽略该参数。
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 指实例新私有网络所属子网 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指实例新的私有网络ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 指实例新的内网 IPv4 地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 任务 ID。可获取**taskId**，通过接口 **DescribeTaskInfo **查询任务执行状态。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkConfigResponseParams `json:"Response"`
}

func (r *ModifyNetworkConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateRequestParams struct {
	// 源参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板修改后的新名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板修改后的新描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 修改后的新参数列表。
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type ModifyParamTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 源参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板修改后的新名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板修改后的新描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 修改后的新参数列表。
	ParamList []*InstanceParam `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

func (r *ModifyParamTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ParamList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyParamTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyParamTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyParamTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyParamTemplateResponseParams `json:"Response"`
}

func (r *ModifyParamTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyParamTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReplicationGroupRequestParams struct {
	// 复制组字符串ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 复制组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyReplicationGroupRequest struct {
	*tchttp.BaseRequest
	
	// 复制组字符串ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 复制组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyReplicationGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReplicationGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReplicationGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReplicationGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyReplicationGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReplicationGroupResponseParams `json:"Response"`
}

func (r *ModifyReplicationGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReplicationGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type OpenSSLRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *OpenSSLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenSSLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenSSLResponseParams struct {
	// 任务ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenSSLResponse struct {
	*tchttp.BaseResponse
	Response *OpenSSLResponseParams `json:"Response"`
}

func (r *OpenSSLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenSSLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {
	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitnil,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitnil,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitnil,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitnil,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type ParamTemplateInfo struct {
	// 参数模板 ID。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 参数模板名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数模板描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 实例类型。
	// - 2：Redis 2.8内存版（标准架构）。
	// - 3：CKV 3.2内存版（标准架构）。
	// - 4：CKV 3.2内存版（集群架构）。
	// - 5：Redis 2.8内存版（单机）。
	// - 6：Redis 4.0内存版（标准架构）。
	// - 7：Redis 4.0内存版（集群架构）。
	// - 8：Redis 5.0内存版（标准架构）。
	// - 9：Redis 5.0内存版（集群架构）。
	// - 15：Redis 6.2内存版（标准架构）。
	// - 16：Redis 6.2内存版（集群架构）。
	ProductType *uint64 `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

type ParameterDetail struct {
	// 参数名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数类型。
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 参数默认值。
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 参数描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数当前值。
	CurrentValue *string `json:"CurrentValue,omitnil,omitempty" name:"CurrentValue"`

	// 修改参数后，是否需要重启数据库以使参数生效。
	// - 0：不需要重启。
	// - 1：需要重启。
	NeedReboot *int64 `json:"NeedReboot,omitnil,omitempty" name:"NeedReboot"`

	// 参数允许的最大值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *string `json:"Max,omitnil,omitempty" name:"Max"`

	// 参数允许的最小值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *string `json:"Min,omitnil,omitempty" name:"Min"`

	// 参数可选枚举值。如果为非枚举参数，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnumValue []*string `json:"EnumValue,omitnil,omitempty" name:"EnumValue"`
}

type ProductConf struct {
	// 产品类型。
	// - 2：Redis 2.8内存版（标准架构）。
	// - 3：CKV 3.2内存版（标准架构）。
	// - 4：CKV 3.2内存版（集群架构）。
	// - 5：Redis 2.8内存版（单机）。
	// - 6：Redis 4.0内存版（标准架构）。
	// - 7：Redis 4.0内存版（集群架构）。
	// - 8：Redis 5.0内存版（标准架构）。
	// - 9：Redis 5.0内存版（集群架构）。
	// - 15：Redis 6.2内存版（标准架构）。
	// - 16：Redis 6.2内存版（集群架构）。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 产品名称。包括：Redis 主从版、CKV 主从版、CKV 集群版、Redis 单机版、Redis 集群版。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 购买时的最小数量。
	MinBuyNum *int64 `json:"MinBuyNum,omitnil,omitempty" name:"MinBuyNum"`

	// 购买时的最大数量。
	MaxBuyNum *int64 `json:"MaxBuyNum,omitnil,omitempty" name:"MaxBuyNum"`

	// 产品是否售罄。
	// - true：售罄。
	// - false：未售罄。
	Saleout *bool `json:"Saleout,omitnil,omitempty" name:"Saleout"`

	// 产品引擎。Redis 或者 CKV。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 兼容版本。包括：Redis-2.8、Redis-3.2、Redis-4.0、Redis-5.0、Redis-6.2。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 规格总大小，单位GB。
	TotalSize []*string `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 每个分片大小，单位GB。
	ShardSize []*string `json:"ShardSize,omitnil,omitempty" name:"ShardSize"`

	// 副本数量。
	ReplicaNum []*string `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// 分片数量。
	ShardNum []*string `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 支持的计费模式。
	// - 1：包年包月。
	// - 0：按量计费。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 该参数名因存在拼写不规范的问题，建议使用**EnableReplicaReadOnly**参数取代。其含义为是否支持副本只读。
	// - true：支持副本只读。
	// - false：不支持。
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitnil,omitempty" name:"EnableRepicaReadOnly"`

	// 是否支持副本只读。
	// - true：支持副本只读。
	// - false：不支持。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableReplicaReadOnly *bool `json:"EnableReplicaReadOnly,omitnil,omitempty" name:"EnableReplicaReadOnly"`
}

type ProxyNodes struct {
	// 节点 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 可用区 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type RedisBackupSet struct {
	// 备份开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备份任务ID。
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 备份类型。
	// - 1：凌晨系统发起的自动备份。
	// - 0：用户发起的手动备份。
	BackupType *string `json:"BackupType,omitnil,omitempty" name:"BackupType"`

	// 备份状态。 
	// - 1：备份被其它流程锁定。
	// - 2：备份正常，没有被任何流程锁定。
	// - -1：备份已过期。
	// - 3：备份正在被导出。
	// - 4：备份导出成功。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 备份的备注信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 备份是否被锁定。
	// - 0：未被锁定。
	// - 1：已被锁定。
	Locked *int64 `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 内部字段，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSize *int64 `json:"BackupSize,omitnil,omitempty" name:"BackupSize"`

	// 内部字段，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullBackup *int64 `json:"FullBackup,omitnil,omitempty" name:"FullBackup"`

	// 内部字段，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 本地备份所在地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 备份结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备份文件类型。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 备份文件过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type RedisCommonInstanceList struct {
	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户APPID。APPID是与账号ID有唯一对应关系的应用 ID，部分腾讯云产品会使用此 APPID。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 实例所属项目 ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例接入区域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例接入可用区。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络所属子网 ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例状态信息。
	// - 1-流程中。
	// - 2-运行中。
	// - -2-实例已隔离。
	// - -3-实例待回收。
	// - -4-实例已删除。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例私有网络 IP 地址。
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`

	// 实例网络端口。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例创建时间。
	Createtime *string `json:"Createtime,omitnil,omitempty" name:"Createtime"`

	// 计费类型。
	// - 0：按量计费。
	// - 1：包年包月。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 网络类型。
	// - 0：基础网络。
	// - 1：VPC 网络。
	NetType *int64 `json:"NetType,omitnil,omitempty" name:"NetType"`
}

type RedisInstanceEvent struct {
	// 事件 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 事件类型，当前仅支持配置实例迁移、资源腾挪、机房裁撤相关的运维操作。该参数仅支持配置为 **InstanceMigration**。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件等级根据其影响严重程度和紧急程度进行分级，由重至轻依次为关键、重要、中等、一般。
	// - Critical：关键
	// - High：重要
	// - Middle：中等
	// - Low：一般
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grade *string `json:"Grade,omitnil,omitempty" name:"Grade"`

	// 事件计划执行日期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionDate *string `json:"ExecutionDate,omitnil,omitempty" name:"ExecutionDate"`

	// 事件计划执行开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件计划执行结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 运维事件最迟执行的日期，即该事件必须在该日期之前完成，否则可能会对业务产生影响。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestExecutionDate *string `json:"LatestExecutionDate,omitnil,omitempty" name:"LatestExecutionDate"`

	// 事件当前状态。
	// - Waiting：未到达执行日期或不在维护时间窗内的事件。
	// - Running：在维护时间窗内，正在执行维护的事件。
	// - Finished：已全部完成维护的事件。
	// - Canceled：已取消执行的事件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件执行任务完成时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskEndTime *string `json:"TaskEndTime,omitnil,omitempty" name:"TaskEndTime"`

	// 事件影响信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectInfo *string `json:"EffectInfo,omitnil,omitempty" name:"EffectInfo"`

	// 事件最初计划执行日期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialExecutionDate *string `json:"InitialExecutionDate,omitnil,omitempty" name:"InitialExecutionDate"`
}

type RedisNode struct {
	// Redis 节点上 Key 的个数。
	Keys *int64 `json:"Keys,omitnil,omitempty" name:"Keys"`

	// Redis 节点 Slot 分布范围。例如：0-5460。
	Slot *string `json:"Slot,omitnil,omitempty" name:"Slot"`

	// 节点的序列 ID。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点的状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点角色。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type RedisNodeInfo struct {
	// 节点类型。<ul><li>0：为主节点。</li><li>1：为副本节点。</li></ul>
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 主节点或者副本节点的 ID。<ul><li>该参数用于创建 Redis 实例接口[CreateInstances](https://cloud.tencent.com/document/product/239/20026) 并不需要设置，而用于变更实例配置的接口 [UpgradeInstance](https://cloud.tencent.com/document/product/239/20013) 删除副本时才需要设置。</li><li>该参数可使用接口 [DescribeInstances](https://cloud.tencent.com/document/product/239/20018) 获取Integer类型的节点 ID。</li></ul>
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 主节点或者副本节点的可用区 ID。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 主节点或者副本节点的可用区名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type RedisNodes struct {
	// 节点 ID。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点角色。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分片 ID。
	ClusterId *int64 `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 可用区 ID。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type RegionConf struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地域简称
	RegionShortName *string `json:"RegionShortName,omitnil,omitempty" name:"RegionShortName"`

	// 地域所在大区名称
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 可用区信息
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`
}

// Predefined struct for user
type ReleaseWanAddressRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ReleaseWanAddressRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ReleaseWanAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseWanAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseWanAddressResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 关闭外网的状态
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseWanAddressResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseWanAddressResponseParams `json:"Response"`
}

func (r *ReleaseWanAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseWanAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationInstanceRequestParams struct {
	// 复制组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据同步类型，true:需要数据强同步,false:不需要强同步，仅限删除主实例
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type RemoveReplicationInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 复制组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据同步类型，true:需要数据强同步,false:不需要强同步，仅限删除主实例
	SyncType *bool `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

func (r *RemoveReplicationInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "InstanceId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveReplicationInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveReplicationInstanceResponseParams struct {
	// 异步任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveReplicationInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RemoveReplicationInstanceResponseParams `json:"Response"`
}

func (r *RemoveReplicationInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveReplicationInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// 购买时长，单位：月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 标识是否修改计费模式。<ul><li>当前实例计费模式为按量计费方式，预转换为包年包月而续费，请指定该参数为 <b>prepaid</b>。</li><li>当前实例计费模式为包年包月方式，可不设置该参数。</li></ul>
	ModifyPayMode *string `json:"ModifyPayMode,omitnil,omitempty" name:"ModifyPayMode"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 购买时长，单位：月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 标识是否修改计费模式。<ul><li>当前实例计费模式为按量计费方式，预转换为包年包月而续费，请指定该参数为 <b>prepaid</b>。</li><li>当前实例计费模式为包年包月方式，可不设置该参数。</li></ul>
	ModifyPayMode *string `json:"ModifyPayMode,omitnil,omitempty" name:"ModifyPayMode"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "InstanceId")
	delete(f, "ModifyPayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// 交易ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewInstanceResponseParams `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicaGroup struct {
	// 节点组 ID。
	GroupId *int64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 节点组的名称，主节点为空。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 节点的可用区ID，比如ap-guangzhou-1
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点组类型，master为主节点，replica为副本节点。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 节点组节点列表
	RedisNodes []*RedisNode `json:"RedisNodes,omitnil,omitempty" name:"RedisNodes"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// Redis实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重置的密码（切换为免密实例时，可不传；其他情况必传）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否切换免密实例，false-切换为非免密码实例，true-切换为免密码实例；默认false
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// Redis实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重置的密码（切换为免密实例时，可不传；其他情况必传）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否切换免密实例，false-切换为非免密码实例，true-切换为免密码实例；默认false
	NoAuth *bool `json:"NoAuth,omitnil,omitempty" name:"NoAuth"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Password")
	delete(f, "NoAuth")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// 任务ID（修改密码时的任务ID，如果时切换免密码或者非免密码实例，则无需关注此返回值）
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceBundle struct {
	// 资源包名称
	ResourceBundleName *string `json:"ResourceBundleName,omitnil,omitempty" name:"ResourceBundleName"`

	// 可售卖内存，单位：GB
	AvailableMemory *int64 `json:"AvailableMemory,omitnil,omitempty" name:"AvailableMemory"`

	// 资源包个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type ResourceTag struct {
	// 标签Key。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签 Key 对应的 Value。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestoreInstanceRequestParams struct {
	// 待操作的实例ID，可通过 DescribeInstances 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份ID，可通过 GetRedisBackupList 接口返回值中的 backupId 获取
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 实例密码，恢复实例时，需要校验实例密码（免密实例不需要传密码）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID，可通过 DescribeInstances 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备份ID，可通过 GetRedisBackupList 接口返回值中的 backupId 获取
	BackupId *string `json:"BackupId,omitnil,omitempty" name:"BackupId"`

	// 实例密码，恢复实例时，需要校验实例密码（免密实例不需要传密码）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreInstanceResponseParams struct {
	// 任务ID，可通过 DescribeTaskInfo 接口查询任务执行状态
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestoreInstanceResponseParams `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组备注。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// 出站规则。
	Outbound []*Outbound `json:"Outbound,omitnil,omitempty" name:"Outbound"`

	// 入站规则。
	Inbound []*Inbound `json:"Inbound,omitnil,omitempty" name:"Inbound"`
}

type SecurityGroupDetail struct {
	// 项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建安全组的时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 安全组 ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitnil,omitempty" name:"SecurityGroupName"`

	// 安全组标记。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitnil,omitempty" name:"SecurityGroupRemark"`

	// 安全组入站规则，即控制访问数据库的来源。
	InboundRule []*SecurityGroupsInboundAndOutbound `json:"InboundRule,omitnil,omitempty" name:"InboundRule"`

	// 安全组出站规则。
	OutboundRule []*SecurityGroupsInboundAndOutbound `json:"OutboundRule,omitnil,omitempty" name:"OutboundRule"`
}

type SecurityGroupsInboundAndOutbound struct {
	// 标识出入数据库的IP与端口是否被允许。
	// - ACCEPT：允许。
	// - DROP：不允许。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 出入数据库的IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 端口号。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议类型。
	Proto *string `json:"Proto,omitnil,omitempty" name:"Proto"`
}

type SourceCommand struct {
	// 命令名称。
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 执行次数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type SourceInfo struct {
	// 来源 IP 地址。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 客户端连接数量。
	Conn *int64 `json:"Conn,omitnil,omitempty" name:"Conn"`

	// 命令
	Cmd *int64 `json:"Cmd,omitnil,omitempty" name:"Cmd"`
}

// Predefined struct for user
type StartupInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type StartupInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *StartupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartupInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartupInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartupInstanceResponseParams struct {
	// 该字段已废弃，请通过查询实例接口获取到的状态来判断实例是否已解隔离
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartupInstanceResponseParams `json:"Response"`
}

func (r *StartupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartupInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchAccessNewInstanceRequestParams struct {
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	// 示例值：crs-asdasdas
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type SwitchAccessNewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 指定实例 ID。例如：crs-xjhsdj****。请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	// 示例值：crs-asdasdas
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *SwitchAccessNewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchAccessNewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchAccessNewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchAccessNewInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchAccessNewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *SwitchAccessNewInstanceResponseParams `json:"Response"`
}

func (r *SwitchAccessNewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchAccessNewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchInstanceVipRequestParams struct {
	// 源实例ID
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// 目标实例ID
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// 单位为秒。源实例与目标实例间DTS已断开时间，如果DTS断开时间大于TimeDelay，则不切换VIP，建议尽量根据业务设置一个可接受的值。
	TimeDelay *int64 `json:"TimeDelay,omitnil,omitempty" name:"TimeDelay"`

	// 在DTS断开的情况下是否强制切换。1：强制切换，0：不强制切换
	ForceSwitch *int64 `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// now: 立即切换，syncComplete：等待同步完成后切换
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`
}

type SwitchInstanceVipRequest struct {
	*tchttp.BaseRequest
	
	// 源实例ID
	SrcInstanceId *string `json:"SrcInstanceId,omitnil,omitempty" name:"SrcInstanceId"`

	// 目标实例ID
	DstInstanceId *string `json:"DstInstanceId,omitnil,omitempty" name:"DstInstanceId"`

	// 单位为秒。源实例与目标实例间DTS已断开时间，如果DTS断开时间大于TimeDelay，则不切换VIP，建议尽量根据业务设置一个可接受的值。
	TimeDelay *int64 `json:"TimeDelay,omitnil,omitempty" name:"TimeDelay"`

	// 在DTS断开的情况下是否强制切换。1：强制切换，0：不强制切换
	ForceSwitch *int64 `json:"ForceSwitch,omitnil,omitempty" name:"ForceSwitch"`

	// now: 立即切换，syncComplete：等待同步完成后切换
	SwitchTime *string `json:"SwitchTime,omitnil,omitempty" name:"SwitchTime"`
}

func (r *SwitchInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchInstanceVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcInstanceId")
	delete(f, "DstInstanceId")
	delete(f, "TimeDelay")
	delete(f, "ForceSwitch")
	delete(f, "SwitchTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchInstanceVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchInstanceVipResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *SwitchInstanceVipResponseParams `json:"Response"`
}

func (r *SwitchInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchInstanceVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例ProxyID
	ProxyID *string `json:"ProxyID,omitnil,omitempty" name:"ProxyID"`
}

type SwitchProxyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例ProxyID
	ProxyID *string `json:"ProxyID,omitnil,omitempty" name:"ProxyID"`
}

func (r *SwitchProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ProxyID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchProxyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchProxyResponse struct {
	*tchttp.BaseResponse
	Response *SwitchProxyResponseParams `json:"Response"`
}

func (r *SwitchProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskInfoDetail struct {
	// 任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务类型。
	// - FLOW_CREATE: "001"，新建实例
	// - FLOW_RESIZE ： "002"，配置变更
	// - FLOW_CLOSE："003"，关闭实例
	// - FLOW_CLEAN： "004"，清空实例
	// - FLOW_STARTUP："005"，实例启用。
	// - FLOW_DELETE："006"，删除实例。
	// - FLOW_SETPWD："007"，重置密码。
	// - FLOW_EXPORTBACKUP："009"，导出备份文件。
	// - FLOW_RESTOREBACKUP："010"，恢复备份。
	// - FLOW_BACKUPINSTANCE："012"，备份实例。
	// - FLOW_MIGRATEINSTANCE："013"，迁移实例。
	// - FLOW_DELBACKUP："014"，删除备份。
	// - FLOW_EXCHANGEINSTANCE： "016"，切换实例流程。
	// - FLOW_AUTOBACKUP："017"，自动备份实例。
	// - FLOW_MIGRATECHECK： "022"，迁移参数校验。
	// - FLOW_MIGRATETASK："023"，数据迁移中。
	// - FLOW_CLEANDB："025"，清空某个数据库。
	// - FLOW_CLONEBACKUP："026"，克隆备份。
	// - FLOW_CHANGEVIP： "027"，改变vip地址。
	// - FLOW_EXPORSHR ："028"，扩缩容。
	// - FLOW_ADDNODES："029"，加（减）节点。
	// - FLOW_CHANGENET："031"，改变网络类型。
	// - FLOW_MODIFYINSTACEREADONLY："033"，只读策略变更。
	// - FLOW_MODIFYINSTANCEPARAMS："034"，修改实例参数。
	// - FLOW_MODIFYINSTANCEPASSWORDFREE："035"，设置免密。
	// - FLOW_SWITCHINSTANCEVIP："036"，实例VIP切换。
	// - FLOW_MODIFYINSTANCEACCOUNT："037"，实例帐号变更。
	// - FLOW_MODIFYINSTANCEBANDWIDTH："038"，实例带宽变更。
	// - FLOW_ENABLEINSTANCE_REPLICATE："039"，开启副本只读。
	// - FLOW_DISABLEINSTANCE_REPLICATE："040"，关闭副本只读。
	// - FLOW_UpgradeArch："041"，实例架构升级，主从升集群。
	// - FLOW_DowngradeArch： "042"，实例架构降级，集群降主从。
	// - FLOW_UpgradeVersion： "043"，版本升级。
	// - FLOW_MODIFYCONNECTIONCONFIG："044"，带宽连接数调整。
	// - FLOW_CLEARNETWORK："045"，更换网络，
	// - FLOW_REMOVE_BACKUP_FILE："046"，删除备份。
	// - FLOW_UPGRADE_SUPPORT_MULTI_AZ："047"，升级实例支持多可用区。
	// - FLOW_SHUTDOWN_MASTER："048"，模拟故障。
	// - FLOW_CHANGE_REPLICA_TO_MASTER："049"，手动提主。
	// - FLOW_CODE_ADD_REPLICATION_INSTANCE："050"，新增复制组。
	// - FLOW_OPEN_WAN："052"，开通外网。
	// - FLOW_CLOSE_WAN："053"，关闭外网FLOW_UPDATE_WAN："054"，更新外网。
	// - FLOW_CODE_DELETE_REPLICATION_INSTANCE："055"，解绑复制组。
	// - FLOW_CODE_CHANGE_MASTER_INSTANCE："056"，复制组实例切主。
	// - FLOW_CODE_CHANGE_INSTANCE_ROLE： "057"，更改复制组实例角色。
	// - FLOW_MIGRATE_NODE："058"，迁移节点。
	// - FLOW_SWITCH_NODE："059"，切换节点。
	// - FLOW_UPGRADE_SMALL_VERSION："060"，升级 Redi s版本。
	// - FLOW_UPGRADE_PROXY_VERSION："061"，升级 Proxy 版本。
	// - FLOW_MODIFY_INSTANCE_NETWORK： "062"，实例修改网络。
	// - FLOW_MIGRATE_PROXY_NODE："063"，迁移proxy节点。
	// - FLOW_MIGRATION_INSTANCE_ZONE："066"，实例可用区迁移中。
	// - FLOW_UPGRADE_INSTANCE_CACHE_AND_PROXY： "067"，实例版本升级中。
	// - FLOW_MODIFY_PROXY_NUM："069"，加（减）Proxy 节点。
	// - FLOW_MODIFYBACKUPMOD："070"，变更实例备份模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 实例名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 项目 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务进度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务执行结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务执行状态。
	// 
	// 0：任务初始化。
	// 1：执行中。
	// 2：完成。
	// 4：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`
}

type TendisNodes struct {
	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点角色
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 可用区 ID。	
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type TendisSlowLogDetail struct {
	// 执行时间
	ExecuteTime *string `json:"ExecuteTime,omitnil,omitempty" name:"ExecuteTime"`

	// 慢查询耗时（毫秒）
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 详细命令行信息
	CommandLine *string `json:"CommandLine,omitnil,omitempty" name:"CommandLine"`

	// 节点ID
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`
}

type TradeDealDetail struct {
	// 订单号ID，调用云API时使用此ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 长订单ID，反馈订单问题给官方客服使用此ID
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 订单关联的实例数
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 创建用户uin
	Creater *string `json:"Creater,omitnil,omitempty" name:"Creater"`

	// 订单创建时间
	CreatTime *string `json:"CreatTime,omitnil,omitempty" name:"CreatTime"`

	// 订单超时时间
	OverdueTime *string `json:"OverdueTime,omitnil,omitempty" name:"OverdueTime"`

	// 订单完成时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 订单状态 1：未支付 2:已支付，未发货 3:发货中 4:发货成功 5:发货失败 6:已退款 7:已关闭订单 8:订单过期 9:订单已失效 10:产品已失效 11:代付拒绝 12:支付中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 订单状态描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 订单实际总价，单位：分
	Price *int64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// 待变更实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指实例每个分片内存变更后的大小。<ul><li>单位 MB。</li><li>每次只能修改参数MemSize、RedisShardNum和RedisReplicasNum其中的一个，不能同时修改。且修改其中一个参数时，其他两个参数需输入实例原有的配置规格。</li><li>缩容时，缩容后的规格务必要大于等于使用容量的1.3倍，否则将执行失败。</li></ul>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 指实例变更后的分片数量。<ul><li>标准架构不需要配置该参数，集群架构为必填参数。</li><li>集群架构，每次只能修改参数RedisShardNum、MemSize和RedisReplicasNum其中的一个，不能同时修改。且修改其中一个参数时，其他两个参数需输入实例原有的配置规格。</li></ul>
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 指实例变更后的副本数量。<ul><li>每次只能修改参数RedisReplicasNum、MemSize和RedisShardNum其中的一个，不能同时修改。且修改其中一个参数时，其他两个参数需输入实例原有的配置规格。</li><li>多AZ实例修改副本时必须要传入NodeSet。</li></ul>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 多AZ实例，增加副本时的节点信息，包括副本的 ID 编号及可用区信息。非多AZ实例不需要配置该参数。
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 待变更实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指实例每个分片内存变更后的大小。<ul><li>单位 MB。</li><li>每次只能修改参数MemSize、RedisShardNum和RedisReplicasNum其中的一个，不能同时修改。且修改其中一个参数时，其他两个参数需输入实例原有的配置规格。</li><li>缩容时，缩容后的规格务必要大于等于使用容量的1.3倍，否则将执行失败。</li></ul>
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 指实例变更后的分片数量。<ul><li>标准架构不需要配置该参数，集群架构为必填参数。</li><li>集群架构，每次只能修改参数RedisShardNum、MemSize和RedisReplicasNum其中的一个，不能同时修改。且修改其中一个参数时，其他两个参数需输入实例原有的配置规格。</li></ul>
	RedisShardNum *uint64 `json:"RedisShardNum,omitnil,omitempty" name:"RedisShardNum"`

	// 指实例变更后的副本数量。<ul><li>每次只能修改参数RedisReplicasNum、MemSize和RedisShardNum其中的一个，不能同时修改。且修改其中一个参数时，其他两个参数需输入实例原有的配置规格。</li><li>多AZ实例修改副本时必须要传入NodeSet。</li></ul>
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitnil,omitempty" name:"RedisReplicasNum"`

	// 多AZ实例，增加副本时的节点信息，包括副本的 ID 编号及可用区信息。非多AZ实例不需要配置该参数。
	NodeSet []*RedisNodeInfo `json:"NodeSet,omitnil,omitempty" name:"NodeSet"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MemSize")
	delete(f, "RedisShardNum")
	delete(f, "RedisReplicasNum")
	delete(f, "NodeSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// 订单ID。
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceVersionRequestParams struct {
	// 目标实例类型，同 [CreateInstances](https://cloud.tencent.com/document/api/239/20026) 的**TypeId**，即实例要变更的目标类型。
	// - Redis 4.0 及以上的版本，支持相同版本的实例从标准架构升级至集群架构，例如，支持 Redis 4.0 标准架构升级至 Redis 4.0 集群架构。
	// - 不支持跨版本架构升级，例如，Redis 4.0 标准架构升级至 Redis 5.0 集群架构。
	// - 不支持 Redis 2.8 版本升级架构。
	// - 不支持从集群架构降级至标准架构。
	TargetInstanceType *string `json:"TargetInstanceType,omitnil,omitempty" name:"TargetInstanceType"`

	// 切换时间。
	// - 1：维护时间窗切换。
	// - 2：立即切换。
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// 指定实例 ID。例如：crs-xjhsdj****，请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type UpgradeInstanceVersionRequest struct {
	*tchttp.BaseRequest
	
	// 目标实例类型，同 [CreateInstances](https://cloud.tencent.com/document/api/239/20026) 的**TypeId**，即实例要变更的目标类型。
	// - Redis 4.0 及以上的版本，支持相同版本的实例从标准架构升级至集群架构，例如，支持 Redis 4.0 标准架构升级至 Redis 4.0 集群架构。
	// - 不支持跨版本架构升级，例如，Redis 4.0 标准架构升级至 Redis 5.0 集群架构。
	// - 不支持 Redis 2.8 版本升级架构。
	// - 不支持从集群架构降级至标准架构。
	TargetInstanceType *string `json:"TargetInstanceType,omitnil,omitempty" name:"TargetInstanceType"`

	// 切换时间。
	// - 1：维护时间窗切换。
	// - 2：立即切换。
	SwitchOption *int64 `json:"SwitchOption,omitnil,omitempty" name:"SwitchOption"`

	// 指定实例 ID。例如：crs-xjhsdj****，请登录[Redis控制台](https://console.cloud.tencent.com/redis#/)在实例列表复制实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *UpgradeInstanceVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetInstanceType")
	delete(f, "SwitchOption")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceVersionResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceVersionResponseParams `json:"Response"`
}

func (r *UpgradeInstanceVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前proxy版本
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// 可升级的redis版本
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil,omitempty" name:"UpgradeProxyVersion"`

	// 1-立即升级   0-维护时间窗口升级
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

type UpgradeProxyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前proxy版本
	CurrentProxyVersion *string `json:"CurrentProxyVersion,omitnil,omitempty" name:"CurrentProxyVersion"`

	// 可升级的redis版本
	UpgradeProxyVersion *string `json:"UpgradeProxyVersion,omitnil,omitempty" name:"UpgradeProxyVersion"`

	// 1-立即升级   0-维护时间窗口升级
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

func (r *UpgradeProxyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CurrentProxyVersion")
	delete(f, "UpgradeProxyVersion")
	delete(f, "InstanceTypeUpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeProxyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeProxyVersionResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeProxyVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeProxyVersionResponseParams `json:"Response"`
}

func (r *UpgradeProxyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeProxyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeSmallVersionRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前redis版本
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil,omitempty" name:"CurrentRedisVersion"`

	// 可升级的redis版本
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil,omitempty" name:"UpgradeRedisVersion"`

	// 1-立即升级   0-维护时间窗口升级
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

type UpgradeSmallVersionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 当前redis版本
	CurrentRedisVersion *string `json:"CurrentRedisVersion,omitnil,omitempty" name:"CurrentRedisVersion"`

	// 可升级的redis版本
	UpgradeRedisVersion *string `json:"UpgradeRedisVersion,omitnil,omitempty" name:"UpgradeRedisVersion"`

	// 1-立即升级   0-维护时间窗口升级
	InstanceTypeUpgradeNow *int64 `json:"InstanceTypeUpgradeNow,omitnil,omitempty" name:"InstanceTypeUpgradeNow"`
}

func (r *UpgradeSmallVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeSmallVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CurrentRedisVersion")
	delete(f, "UpgradeRedisVersion")
	delete(f, "InstanceTypeUpgradeNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeSmallVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeSmallVersionResponseParams struct {
	// 异步流程ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeSmallVersionResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeSmallVersionResponseParams `json:"Response"`
}

func (r *UpgradeSmallVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeSmallVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeVersionToMultiAvailabilityZonesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级多可用区之后是否支持就近访问功能。
	// <ul><li>true：支持就近访问功能。升级过程，需同时升级 Proxy 版本和 Redis 内核小版本，涉及数据搬迁，可能会长达数小时。</li><li>false：无需支持就近访问功能。升级多可用区仅涉及管理元数据迁移，对服务没有影响，升级过程通常在3分钟内完成。</li></ul>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitnil,omitempty" name:"UpgradeProxyAndRedisServer"`
}

type UpgradeVersionToMultiAvailabilityZonesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级多可用区之后是否支持就近访问功能。
	// <ul><li>true：支持就近访问功能。升级过程，需同时升级 Proxy 版本和 Redis 内核小版本，涉及数据搬迁，可能会长达数小时。</li><li>false：无需支持就近访问功能。升级多可用区仅涉及管理元数据迁移，对服务没有影响，升级过程通常在3分钟内完成。</li></ul>
	UpgradeProxyAndRedisServer *bool `json:"UpgradeProxyAndRedisServer,omitnil,omitempty" name:"UpgradeProxyAndRedisServer"`
}

func (r *UpgradeVersionToMultiAvailabilityZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeVersionToMultiAvailabilityZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpgradeProxyAndRedisServer")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeVersionToMultiAvailabilityZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeVersionToMultiAvailabilityZonesResponseParams struct {
	// 任务ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeVersionToMultiAvailabilityZonesResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeVersionToMultiAvailabilityZonesResponseParams `json:"Response"`
}

func (r *UpgradeVersionToMultiAvailabilityZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeVersionToMultiAvailabilityZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneCapacityConf struct {
	// 可用区ID：如ap-guangzhou-3
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区名称。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区是否售罄。
	IsSaleout *bool `json:"IsSaleout,omitnil,omitempty" name:"IsSaleout"`

	// 是否为默认可用区。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 网络类型。
	// - basenet：基础网络。
	// - vpcnet -- VPC网络。
	NetWorkType []*string `json:"NetWorkType,omitnil,omitempty" name:"NetWorkType"`

	// 可用区内产品规格等信息
	ProductSet []*ProductConf `json:"ProductSet,omitnil,omitempty" name:"ProductSet"`

	// 可用区ID：如100003
	OldZoneId *int64 `json:"OldZoneId,omitnil,omitempty" name:"OldZoneId"`
}