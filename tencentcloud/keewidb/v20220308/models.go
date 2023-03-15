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

package v20220308

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要绑定的安全组 ID，类似sg-efil7***。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID，格式如：kee-c1nl9***，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要绑定的安全组 ID，类似sg-efil7***。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID，格式如：kee-c1nl9***，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type BackupInfo struct {
	// 备份开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份 ID。
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// 备份类型。<ul><li>1：手动备份，指根据业务运维排障需求，立即执行备份任务的操作。</li> <li>0：自动备份，指根据自动备份策略定时自动发起的备份任务。</li></ul>
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份的备注信息.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 备份状态。  <ul><li>1：备份任务被其它流程锁定。</li><li>2：备份正常，没有被任何流程锁定。</li> <li>-1：备份已过期。</li><li>3：备份正在被导出。</li> <li>4：备份导出成功。</li></ul>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 备份是否被锁定。<ul><li>0：未被锁定。</li><li>1：已被锁定。</li></ul>
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`
}

type BinlogInfo struct {
	// 备份开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份 ID。
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// 备份文件名。
	Filename *string `json:"Filename,omitempty" name:"Filename"`

	// 备份文件大小，单位：Byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
}

// Predefined struct for user
type ChangeInstanceMasterRequestParams struct {
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 副本节点 ID。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

type ChangeInstanceMasterRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 副本节点 ID。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *ChangeInstanceMasterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceMasterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeInstanceMasterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeInstanceMasterResponseParams struct {
	// 异步任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChangeInstanceMasterResponse struct {
	*tchttp.BaseResponse
	Response *ChangeInstanceMasterResponseParams `json:"Response"`
}

func (r *ChangeInstanceMasterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeInstanceMasterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CleanUpInstanceRequestParams struct {
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CleanUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubhg****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例访问密码。
	// 实例为免密访问，则无需设置该参数。
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubhg****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例访问密码。
	// 实例为免密访问，则无需设置该参数。
	Password *string `json:"Password,omitempty" name:"Password"`
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
	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateBackupManuallyRequestParams struct {
	// 待操作的实例 ID，可通过 DescribeInstance接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 本次备份的备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 备份文件保存天数。0代表指定默认保留时间
	StorageDays *int64 `json:"StorageDays,omitempty" name:"StorageDays"`
}

type CreateBackupManuallyRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例 ID，可通过 DescribeInstance接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 本次备份的备注信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 备份文件保存天数。0代表指定默认保留时间
	StorageDays *int64 `json:"StorageDays,omitempty" name:"StorageDays"`
}

func (r *CreateBackupManuallyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupManuallyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Remark")
	delete(f, "StorageDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupManuallyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupManuallyResponseParams struct {
	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupManuallyResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupManuallyResponseParams `json:"Response"`
}

func (r *CreateBackupManuallyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupManuallyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesRequestParams struct {
	// 产品版本。
	// 14：当前仅支持混合存储版。
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// 私有网络唯一ID。
	// 请登录控制台在私有网络列表查询，如：vpc-azlk3***。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络所属子网唯一ID。
	// 请登录控制台在私有网络列表查询，如：subnet-8abje***。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 计费模式。<ul><li>0：按量计费。</li><li>1：包年包月。</li></ul>
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例数量，单次最大购买数量以查询产品售卖规格返回的数量为准。
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 选择包年包月计费模式（BillingMode 设置为1）时，您需要选择购买实例的时长。单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。按量计费（BillingMode 设置为0）实例该参数设置为1即可。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 分片数量，支持选择3、5、6、8、9、10、12、15、16、18、20、21、24、25、27、30、32、33、35、36、39、40、42、45、48、50、51、54、55、56、57、60、63、64分片。
	ShardNum *int64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// 副本数。当前仅支持设置1个副本节点，即每一个分片仅包含1个主节点与1个副本节点，数据主从实时热备。
	ReplicasNum *int64 `json:"ReplicasNum,omitempty" name:"ReplicasNum"`

	// 计算cpu核心数。
	MachineCpu *int64 `json:"MachineCpu,omitempty" name:"MachineCpu"`

	// 实例内存容量，单位：GB。
	// KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。
	MachineMemory *int64 `json:"MachineMemory,omitempty" name:"MachineMemory"`

	// 实例所属的可用区ID。<ul><li>具体取值，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106)获取。</li><li>参数<b>ZoneId</b>和<b>ZoneName</b>至少配置其中一个。</li></u>
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例所属的可用区名称。<ul><li>具体取值，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106)获取。</li><li>参数<b>ZoneId</b>和<b>ZoneName</b>至少配置其中一个。</li></u>
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 创建实例的名称。
	// 仅支持长度小于60的中文、英文或者数字，短划线"-"、下划线"_"。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 指明创建的实例是否需要支持免密访问。<ul><li>true：免密实例。</li><li>false：非免密实例，默认为非免密实例。此时，需要设置访问密码。</li></ul>
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// 实例访问密码。<ul><li>当参数<b>NoAuth</b>为<b>true</b>时，Password为无需设置，否则Password为必填参数。</li>
	// <li>密码复杂度要求：<ul><li>8-30个字符。</li><li>至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种。</li><li>不能以"/"开头。</li></ul></li></ul>
	Password *string `json:"Password,omitempty" name:"Password"`

	// 自定义端口。默认为6379，范围[1024,65535]。
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// 包年包月计费的续费模式。<ul><li>0：默认状态，指手动续费。</li><li>1：自动续费。</li><li>2：到期不再续费。</ul>
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 给实例设置安全组 ID 数组。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`

	// 给实例绑定标签。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 混合存储版，单分片持久化内存容量，单位：GB。
	// KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// 每个分片硬盘的容量。单位：GB。
	// 每一缓存分片容量，对应的磁盘容量范围不同。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 项目id，取值以用户账户>用户账户相关接口查询>项目列表返回的projectId为准。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 产品版本。
	// 14：当前仅支持混合存储版。
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// 私有网络唯一ID。
	// 请登录控制台在私有网络列表查询，如：vpc-azlk3***。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络所属子网唯一ID。
	// 请登录控制台在私有网络列表查询，如：subnet-8abje***。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 计费模式。<ul><li>0：按量计费。</li><li>1：包年包月。</li></ul>
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例数量，单次最大购买数量以查询产品售卖规格返回的数量为准。
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 选择包年包月计费模式（BillingMode 设置为1）时，您需要选择购买实例的时长。单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。按量计费（BillingMode 设置为0）实例该参数设置为1即可。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 分片数量，支持选择3、5、6、8、9、10、12、15、16、18、20、21、24、25、27、30、32、33、35、36、39、40、42、45、48、50、51、54、55、56、57、60、63、64分片。
	ShardNum *int64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// 副本数。当前仅支持设置1个副本节点，即每一个分片仅包含1个主节点与1个副本节点，数据主从实时热备。
	ReplicasNum *int64 `json:"ReplicasNum,omitempty" name:"ReplicasNum"`

	// 计算cpu核心数。
	MachineCpu *int64 `json:"MachineCpu,omitempty" name:"MachineCpu"`

	// 实例内存容量，单位：GB。
	// KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。
	MachineMemory *int64 `json:"MachineMemory,omitempty" name:"MachineMemory"`

	// 实例所属的可用区ID。<ul><li>具体取值，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106)获取。</li><li>参数<b>ZoneId</b>和<b>ZoneName</b>至少配置其中一个。</li></u>
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例所属的可用区名称。<ul><li>具体取值，请参见[地域和可用区](https://cloud.tencent.com/document/product/239/4106)获取。</li><li>参数<b>ZoneId</b>和<b>ZoneName</b>至少配置其中一个。</li></u>
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 创建实例的名称。
	// 仅支持长度小于60的中文、英文或者数字，短划线"-"、下划线"_"。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 指明创建的实例是否需要支持免密访问。<ul><li>true：免密实例。</li><li>false：非免密实例，默认为非免密实例。此时，需要设置访问密码。</li></ul>
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// 实例访问密码。<ul><li>当参数<b>NoAuth</b>为<b>true</b>时，Password为无需设置，否则Password为必填参数。</li>
	// <li>密码复杂度要求：<ul><li>8-30个字符。</li><li>至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种。</li><li>不能以"/"开头。</li></ul></li></ul>
	Password *string `json:"Password,omitempty" name:"Password"`

	// 自定义端口。默认为6379，范围[1024,65535]。
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// 包年包月计费的续费模式。<ul><li>0：默认状态，指手动续费。</li><li>1：自动续费。</li><li>2：到期不再续费。</ul>
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 给实例设置安全组 ID 数组。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`

	// 给实例绑定标签。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 混合存储版，单分片持久化内存容量，单位：GB。
	// KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// 每个分片硬盘的容量。单位：GB。
	// 每一缓存分片容量，对应的磁盘容量范围不同。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 项目id，取值以用户账户>用户账户相关接口查询>项目列表返回的projectId为准。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
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
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "BillingMode")
	delete(f, "GoodsNum")
	delete(f, "Period")
	delete(f, "ShardNum")
	delete(f, "ReplicasNum")
	delete(f, "MachineCpu")
	delete(f, "MachineMemory")
	delete(f, "ZoneId")
	delete(f, "ZoneName")
	delete(f, "InstanceName")
	delete(f, "NoAuth")
	delete(f, "Password")
	delete(f, "VPort")
	delete(f, "AutoRenew")
	delete(f, "SecurityGroupIdList")
	delete(f, "ResourceTags")
	delete(f, "MemSize")
	delete(f, "DiskSize")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstancesResponseParams struct {
	// 交易 ID。
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 实例 ID 。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeAutoBackupConfigRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 自动备份的周期。包括：Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 自动备份时间段。
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// 全量备份文件保存天数。
	BackupStorageDays *int64 `json:"BackupStorageDays,omitempty" name:"BackupStorageDays"`

	// 增量备份文件保存天数。
	BinlogStorageDays *int64 `json:"BinlogStorageDays,omitempty" name:"BinlogStorageDays"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeDBSecurityGroupsRequestParams struct {
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例ID，格式如：kee-c1nl9***。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例ID，格式如：kee-c1nl9***。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// 安全组生效内网地址。
	VIP *string `json:"VIP,omitempty" name:"VIP"`

	// 安全组生效内网端口。
	VPort *string `json:"VPort,omitempty" name:"VPort"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceBackupsRequestParams struct {
	// 待操作的实例ID，可通过 DescribeInstance 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出的备份列表大小，即每页输出的备份文件的数量，默认值20，取值范围为[1,100]。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 备份列表分页偏移量，取Limit整数倍。
	// 计算公式为offset=limit*(页码-1)。例如 limit=10，第1页offset就为0，第2页offset就为10，依次类推。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询备份文件的开始时间，格式如：2017-02-08 16:46:34。查询实例在 [BeginTime, EndTime] 时间段内的备份列表。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询备份文件的结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内的备份列表。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份任务状态。<ul><li>1：备份在流程中。</li><li>2：备份正常。</li><li>3：备份转RDB文件处理中。</li><li>4：已完成RDB转换。</li><li>-1：备份已过期。</li><li>-2：备份已删除。</li></ul>
	Status []*int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的实例ID，可通过 DescribeInstance 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出的备份列表大小，即每页输出的备份文件的数量，默认值20，取值范围为[1,100]。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 备份列表分页偏移量，取Limit整数倍。
	// 计算公式为offset=limit*(页码-1)。例如 limit=10，第1页offset就为0，第2页offset就为10，依次类推。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询备份文件的开始时间，格式如：2017-02-08 16:46:34。查询实例在 [BeginTime, EndTime] 时间段内的备份列表。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询备份文件的结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内的备份列表。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份任务状态。<ul><li>1：备份在流程中。</li><li>2：备份正常。</li><li>3：备份转RDB文件处理中。</li><li>4：已完成RDB转换。</li><li>-1：备份已过期。</li><li>-2：备份已删除。</li></ul>
	Status []*int64 `json:"Status,omitempty" name:"Status"`
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
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBackupsResponseParams struct {
	// 备份文件总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 废弃字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSet []*BinlogInfo `json:"BackupSet,omitempty" name:"BackupSet"`

	// 实例备份信息列表。
	BackupRecord []*BackupInfo `json:"BackupRecord,omitempty" name:"BackupRecord"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceBinlogsRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出备份列表大小，默认大小20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间，格式如：2017-02-08 16:46:34。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeInstanceBinlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出备份列表大小，默认大小20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间，格式如：2017-02-08 16:46:34。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeInstanceBinlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBinlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceBinlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceBinlogsResponseParams struct {
	// 备份总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例的备份信息数组
	BackupSet []*BinlogInfo `json:"BackupSet,omitempty" name:"BackupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceBinlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceBinlogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceBinlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceBinlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceDealDetailRequestParams struct {
	// 订单交易ID数组，即 [CreateInstances](https://cloud.tencent.com/document/api/1520/86207) 的输出参数DealId。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest
	
	// 订单交易ID数组，即 [CreateInstances](https://cloud.tencent.com/document/api/1520/86207) 的输出参数DealId。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds"`
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
	// 订单详细信息
	DealDetails []*TradeDealDetail `json:"DealDetails,omitempty" name:"DealDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceNodeInfoRequestParams struct {
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出的节点信息大小。默认为 20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出的节点信息大小。默认为 20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	// Proxy 节点数量。
	ProxyCount *int64 `json:"ProxyCount,omitempty" name:"ProxyCount"`

	// Proxy 节点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proxy []*ProxyNodeInfo `json:"Proxy,omitempty" name:"Proxy"`

	// Redis 节点数量。该参数仅为产品兼容性而保留，并不具有实际意义，可忽略。
	RedisCount *int64 `json:"RedisCount,omitempty" name:"RedisCount"`

	// Redis 节点信息。该参数仅为产品兼容性而保留，并不具有实际意义，可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Redis []*RedisNodeInfo `json:"Redis,omitempty" name:"Redis"`

	// Tendis 节点数量。该参数仅为产品兼容性而保留，并不具有实际意义，可忽略。
	TendisCount *int64 `json:"TendisCount,omitempty" name:"TendisCount"`

	// Tendis 节点信息。该参数仅为产品兼容性而保留，并不具有实际意义，可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tendis []*InstanceNodeInfo `json:"Tendis,omitempty" name:"Tendis"`

	// KeewiDB 节点数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeeWiDBCount *int64 `json:"KeeWiDBCount,omitempty" name:"KeeWiDBCount"`

	// KeewiDB 节点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeeWiDB []*InstanceNodeInfo `json:"KeeWiDB,omitempty" name:"KeeWiDB"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出的参数列表大小。默认为 20，最多输出100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 每页输出的参数列表大小。默认为 20，最多输出100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	// 修改历史记录总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 修改历史记录信息。
	InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitempty" name:"InstanceParamHistory"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 实例参数总数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例枚举类型参数数组。
	InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitempty" name:"InstanceEnumParam"`

	// 实例整型参数数组。
	InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitempty" name:"InstanceIntegerParam"`

	// 实例字符型参数数组。
	InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitempty" name:"InstanceTextParam"`

	// 实例多选项型参数数组。
	InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitempty" name:"InstanceMultiParam"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeInstanceReplicasRequestParams struct {
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceReplicasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceReplicasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceReplicasResponseParams struct {
	// 实例所有节点数量，包括主节点、副本节点。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例节点信息。
	ReplicaGroups []*ReplicaGroup `json:"ReplicaGroups,omitempty" name:"ReplicaGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceReplicasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceReplicasResponseParams `json:"Response"`
}

func (r *DescribeInstanceReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 每页输出的实例列表的大小，即每页输出的实例数量，默认值20，取值范围为[1,1000]。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。
	// 计算公式为offset=limit*(页码-1)。例如 limit=10，第1页offset就为0，第2页offset就为10，依次类推。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 排序依据。枚举范围如下所示。 <ul><li>projectId：实例按照项目ID排序。</li><li>createtime：实例按照创建时间排序。</li><li>instancename：实例按照实例名称排序。</li><li>type：实例按照类型排序。</li><li>curDeadline：实例按照到期时间排序。</li></ul>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式。<ul><li>1：倒序。默认为倒序。</li><li>0：顺序。</li></ul>
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 私有网络ID数组。数组下标从0开始，如果不传则默认选择基础网络，如：47525
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 子网ID数组，数组下标从0开始，如：56854
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 项目ID 组成的数组，数组下标从0开始
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 查找关键字，可输入实例的ID或者实例名称。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 私有网络ID数组，数组下标从0开始，如果不传则默认选择基础网络，如：vpc-sad23jfdfk
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// 子网ID数组，数组下标从0开始，如：subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// 实例状态。<ul><li>0：待初始化。</li><li>1：流程中。</li><li>2：运行中。</li><li>-2：已隔离。</li><li>-3：待删除。</li></ul>
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 包年包月计费的续费模式。<ul><li>0：默认状态，指手动续费。</li><li>1：自动续费。</li><li>2：到期不再续费。</ul>
	AutoRenew []*int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 计费模式。<ul><li>postpaid：按量计费。</li><li>prepaid：包年包月。</li></ul>
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例类型。<ul><li>13：标准版。</li><li>14：集群版。</li></ul>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 搜索关键词：支持实例 ID、实例名称、私有网络IP地址。
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// 内部参数，用户可忽略。
	TypeList []*int64 `json:"TypeList,omitempty" name:"TypeList"`

	// 内部参数，用户可忽略。
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`

	// 根据标签的 Key 和 Value 筛选资源。该参数不配置或者数组设置为空值，则不根据标签进行过滤。
	InstanceTags *InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags"`

	// 根据标签的 Key 筛选资源，该参数不配置或者数组设置为空值，则不根据标签Key进行过滤。
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 每页输出的实例列表的大小，即每页输出的实例数量，默认值20，取值范围为[1,1000]。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，取Limit整数倍。
	// 计算公式为offset=limit*(页码-1)。例如 limit=10，第1页offset就为0，第2页offset就为10，依次类推。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 排序依据。枚举范围如下所示。 <ul><li>projectId：实例按照项目ID排序。</li><li>createtime：实例按照创建时间排序。</li><li>instancename：实例按照实例名称排序。</li><li>type：实例按照类型排序。</li><li>curDeadline：实例按照到期时间排序。</li></ul>
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式。<ul><li>1：倒序。默认为倒序。</li><li>0：顺序。</li></ul>
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 私有网络ID数组。数组下标从0开始，如果不传则默认选择基础网络，如：47525
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds"`

	// 子网ID数组，数组下标从0开始，如：56854
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`

	// 项目ID 组成的数组，数组下标从0开始
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 查找关键字，可输入实例的ID或者实例名称。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 私有网络ID数组，数组下标从0开始，如果不传则默认选择基础网络，如：vpc-sad23jfdfk
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// 子网ID数组，数组下标从0开始，如：subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// 实例状态。<ul><li>0：待初始化。</li><li>1：流程中。</li><li>2：运行中。</li><li>-2：已隔离。</li><li>-3：待删除。</li></ul>
	Status []*int64 `json:"Status,omitempty" name:"Status"`

	// 包年包月计费的续费模式。<ul><li>0：默认状态，指手动续费。</li><li>1：自动续费。</li><li>2：到期不再续费。</ul>
	AutoRenew []*int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 计费模式。<ul><li>postpaid：按量计费。</li><li>prepaid：包年包月。</li></ul>
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例类型。<ul><li>13：标准版。</li><li>14：集群版。</li></ul>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 搜索关键词：支持实例 ID、实例名称、私有网络IP地址。
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// 内部参数，用户可忽略。
	TypeList []*int64 `json:"TypeList,omitempty" name:"TypeList"`

	// 内部参数，用户可忽略。
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`

	// 根据标签的 Key 和 Value 筛选资源。该参数不配置或者数组设置为空值，则不根据标签进行过滤。
	InstanceTags *InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags"`

	// 根据标签的 Key 筛选资源，该参数不配置或者数组设置为空值，则不根据标签Key进行过滤。
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`
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
	delete(f, "ProjectIds")
	delete(f, "SearchKey")
	delete(f, "InstanceName")
	delete(f, "UniqVpcIds")
	delete(f, "UniqSubnetIds")
	delete(f, "Status")
	delete(f, "AutoRenew")
	delete(f, "BillingMode")
	delete(f, "Type")
	delete(f, "SearchKeys")
	delete(f, "TypeList")
	delete(f, "MonitorVersion")
	delete(f, "InstanceTags")
	delete(f, "TagKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例详细信息列表
	InstanceSet []*InstanceInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubhg***。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubhg***。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 维护时间窗起始时间，如：03:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 维护时间窗结束时间，如：06:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 地域售卖信息
	RegionSet []*RegionConf `json:"RegionSet,omitempty" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeProjectSecurityGroupsRequestParams struct {
	// 数据库引擎名称。该产品固定为 keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 项目 ID。
	// 登录 [账号中心](https://console.cloud.tencent.com/developer)，在<b>项目管理</b>中可获取项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页安全组的数量限制。默认为 20，最多输出100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键词，支持根据安全组 ID 或者安全组名称搜索。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称。该产品固定为 keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 项目 ID。
	// 登录 [账号中心](https://console.cloud.tencent.com/developer)，在<b>项目管理</b>中可获取项目 ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页安全组的数量限制。默认为 20，最多输出100条。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索关键词，支持根据安全组 ID 或者安全组名称搜索。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
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
	Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`

	// 符合条件的安全组总数量。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubhgouj
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢查询平均执行时间阈值。<ul><li>单位：毫秒。</li><li>取值范围：10、20、30、40、50。</li></ul>
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// 每个页面大小，即每个页面输出慢日志的数量。取值范围为：10、20、30、40、50，默认为 20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页面偏移量，取Limit整数倍，计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeProxySlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubhgouj
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢查询平均执行时间阈值。<ul><li>单位：毫秒。</li><li>取值范围：10、20、30、40、50。</li></ul>
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// 每个页面大小，即每个页面输出慢日志的数量。取值范围为：10、20、30、40、50，默认为 20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页面偏移量，取Limit整数倍，计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 慢查询详情。
	InstanceProxySlowLogDetail []*InstanceProxySlowlogDetail `json:"InstanceProxySlowLogDetail,omitempty" name:"InstanceProxySlowLogDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTaskInfoRequestParams struct {
	// 任务 ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
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
	// 任务状态。<ul><li>preparing：待执行。</li><li>running：执行中。</li><li>succeed：成功。</li><li>failed：失败。</li><li>error：执行出错。</li></ul>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 实例的ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务信息，错误时显示错误信息。执行中与成功则为空值。
	TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 每页输出的任务列表大小。默认为 20，最多输出100条。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset：分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 项目ID。
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 任务类型。可设置为：FLOW_CREATE、FLOW_SETPWD、FLOW_CLOSE等。
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// 起始时间。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 终止时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态。
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 每页输出的任务列表大小。默认为 20，最多输出100条。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Offset：分页偏移量，取Limit整数倍。计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 项目ID。
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 任务类型。可设置为：FLOW_CREATE、FLOW_SETPWD、FLOW_CLOSE等。
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// 起始时间。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 终止时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态。
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskListResponseParams struct {
	// 任务总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务详细信息列表。
	Tasks []*TaskInfoDetail `json:"Tasks,omitempty" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢查询平均执行时间阈值。<ul><li>单位：毫秒。</li><li>取值范围：10、20、30、40、50。</li></ul>
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// 每个页面大小，即每个页面输出慢日志的数量。取值范围为：10、20、30、40、50。默认为 20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页面偏移量，取Limit整数倍，计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTendisSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢查询平均执行时间阈值。<ul><li>单位：毫秒。</li><li>取值范围：10、20、30、40、50。</li></ul>
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// 每个页面大小，即每个页面输出慢日志的数量。取值范围为：10、20、30、40、50。默认为 20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 页面偏移量，取Limit整数倍，计算公式：offset=limit*(页码-1)。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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
	// 慢查询详情。
	TendisSlowLogDetail []*TendisSlowLogDetail `json:"TendisSlowLogDetail,omitempty" name:"TendisSlowLogDetail"`

	// 慢查询总数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DestroyPostpaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DestroyPrepaidInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 交易ID。
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DisassociateSecurityGroupsRequestParams struct {
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要绑定的安全组 ID，类似sg-efil****。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID，格式如：kee-c1nl****，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要绑定的安全组 ID，类似sg-efil****。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例 ID，格式如：kee-c1nl****，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Inbound struct {
	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type InstanceEnumParam struct {
	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：enum
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数可取值
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceInfo struct {
	// 实例名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户的Appid。
	Appid *int64 `json:"Appid,omitempty" name:"Appid"`

	// 项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域ID。<ul><li>1：广州。</li><li>4：上海。</li><li>8：北京。</li></ul>
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区 ID。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// VPC 网络 ID， 如：75101。该参数当前暂保留，可忽略。
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 实例当前状态。<ul><li>0：待初始化。</li><li>1：实例在流程中。</li><li>2：实例运行中。</li><li>-2：实例已隔离。</li><li>-3：实例待删除。</li></ul>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// VPC 网络下子网 ID， 如：46315。该参数当前暂保留，可忽略。
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例 VIP。
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 实例端口号。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例创建时间。
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// 实例持久内存总容量大小，单位：MB。
	Size *float64 `json:"Size,omitempty" name:"Size"`

	// 实例类型。<ul><li>13：标准版。</li><li>14：集群版。</li></ul>
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 实例是否设置自动续费标识。<ul><li>1：设置自动续费。</li><li>0：未设置自动续费。</li></ul>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 实例到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 存储引擎。
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 产品类型。<ul><li>standalone ：标准版。</li><li>cluster ：集群版。</li></ul>
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`

	// VPC 网络 ID， 如：vpc-fk33jsf4****。
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// VPC 网络下子网 ID，如：subnet-fd3j6l3****。
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 计费模式。<ul><li>0：按量计费。</li><li>1：包年包月。</li></ul>
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例运行状态描述：如”实例运行中“。
	InstanceTitle *string `json:"InstanceTitle,omitempty" name:"InstanceTitle"`

	// 计划下线时间。
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 流程中的实例，返回子状态。
	SubStatus *int64 `json:"SubStatus,omitempty" name:"SubStatus"`

	// 反亲和性标签
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 分片大小。
	RedisShardSize *int64 `json:"RedisShardSize,omitempty" name:"RedisShardSize"`

	// 分片数量。
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// 副本数量。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// 计费 ID。
	PriceId *int64 `json:"PriceId,omitempty" name:"PriceId"`

	// 隔离时间。
	CloseTime *string `json:"CloseTime,omitempty" name:"CloseTime"`

	// 从节点读取权重。
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitempty" name:"SlaveReadWeight"`

	// 实例关联的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags"`

	// 项目名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 是否为免密实例；<ul><li>true：免密实例。</li><li>false：非免密实例。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// 客户端连接数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimit *int64 `json:"ClientLimit,omitempty" name:"ClientLimit"`

	// DTS状态（内部参数，用户可忽略）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsStatus *int64 `json:"DtsStatus,omitempty" name:"DtsStatus"`

	// 分片带宽上限，单位 MB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetLimit *int64 `json:"NetLimit,omitempty" name:"NetLimit"`

	// 免密实例标识（内部参数，用户可忽略）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PasswordFree *int64 `json:"PasswordFree,omitempty" name:"PasswordFree"`

	// 实例只读标识（内部参数，用户可忽略）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// 内部参数，用户可忽略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitempty" name:"RemainBandwidthDuration"`

	// 实例的磁盘容量大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 监控版本。<ul><li>1m：分钟粒度监控。</li><li>5s：5秒粒度监控。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`

	// 客户端最大连接数可设置的最小值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimitMin *int64 `json:"ClientLimitMin,omitempty" name:"ClientLimitMin"`

	// 客户端最大连接数可设置的最大值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimitMax *int64 `json:"ClientLimitMax,omitempty" name:"ClientLimitMax"`

	// 实例的节点详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeSet []*NodeInfo `json:"NodeSet,omitempty" name:"NodeSet"`

	// 实例所在的地域信息，比如ap-guangzhou。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例内存容量，单位：GB。KeeWiDB 内存容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MachineMemory *int64 `json:"MachineMemory,omitempty" name:"MachineMemory"`

	// 单分片磁盘大小，单位：MB
	DiskShardSize *int64 `json:"DiskShardSize,omitempty" name:"DiskShardSize"`

	// 3
	DiskShardNum *int64 `json:"DiskShardNum,omitempty" name:"DiskShardNum"`

	// 1
	DiskReplicasNum *int64 `json:"DiskReplicasNum,omitempty" name:"DiskReplicasNum"`
}

type InstanceIntegerParam struct {
	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：integer
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 参数最大值
	Max *string `json:"Max,omitempty" name:"Max"`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 参数单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type InstanceMultiParam struct {
	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：multi
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数说明
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceNodeInfo struct {
	// 节点ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点角色
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
}

type InstanceParam struct {
	// 设置参数的名字
	Key *string `json:"Key,omitempty" name:"Key"`

	// 设置参数的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceParamHistory struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 修改前值
	PreValue *string `json:"PreValue,omitempty" name:"PreValue"`

	// 修改后值
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// 状态：1-参数配置修改中；2-参数配置修改成功；3-参数配置修改失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type InstanceProxySlowlogDetail struct {
	// 慢查询耗时
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 客户端地址
	Client *string `json:"Client,omitempty" name:"Client"`

	// 命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 详细命令行信息
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// 执行时间
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`
}

type InstanceTagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type InstanceTextParam struct {
	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：text
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数可取值
	TextValue []*string `json:"TextValue,omitempty" name:"TextValue"`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type KeeWiDBNode struct {
	// 节点的序列ID。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点的状态。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 节点角色。
	Role *string `json:"Role,omitempty" name:"Role"`
}

// Predefined struct for user
type ModifyAutoBackupConfigRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份周期。可设置为 Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday，该参数暂不支持修改、
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 备份任务执行时间段。
	// 可设置的格式为一个整点到下一个整点。例如：00:00-01:00、01:00-02:00、21:00-22:00、23:00-00:00等。
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份周期。可设置为 Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday，该参数暂不支持修改、
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 备份任务执行时间段。
	// 可设置的格式为一个整点到下一个整点。例如：00:00-01:00、01:00-02:00、21:00-22:00、23:00-00:00等。
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoBackupConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoBackupConfigResponseParams struct {
	// 增量备份文件保存天数。
	BackupStorageDays *int64 `json:"BackupStorageDays,omitempty" name:"BackupStorageDays"`

	// 全量备份文件保存天数。
	BinlogStorageDays *int64 `json:"BinlogStorageDays,omitempty" name:"BinlogStorageDays"`

	// 备份时间段。
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// 备份周期。Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyConnectionConfigRequestParams struct {
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 单分片附加带宽，取值范围[0,512]，单位：MB。
	// <ul><li>开启副本只读时，实例总带宽  = 单分片附加带宽 * 分片数 + 标准带宽 * 分片数 * Max ([只读副本数量, 1])，标准架构的分片数 = 1。</li><li>没有开启副本只读时，实例总带宽 = 单分片附加带宽 * 分片数 + 标准带宽 * 分片数，标准架构的分片数 = 1。</li></ul>
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 单分片的总连接数。
	// <ul>默认为10000，整个实例的最大连接数为单个分片的最大连接数 x 分片数量。标准架构分片数量为1。
	// <li>关闭副本只读：每个分片的最大连接数的取值范围为[10000,40000]。</li><li>开启副本只读：每个分片的最大连接数的取值范围为 [10000,10000 x (副本数 + 3)]。</li></ul>
	ClientLimit *int64 `json:"ClientLimit,omitempty" name:"ClientLimit"`
}

type ModifyConnectionConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 单分片附加带宽，取值范围[0,512]，单位：MB。
	// <ul><li>开启副本只读时，实例总带宽  = 单分片附加带宽 * 分片数 + 标准带宽 * 分片数 * Max ([只读副本数量, 1])，标准架构的分片数 = 1。</li><li>没有开启副本只读时，实例总带宽 = 单分片附加带宽 * 分片数 + 标准带宽 * 分片数，标准架构的分片数 = 1。</li></ul>
	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`

	// 单分片的总连接数。
	// <ul>默认为10000，整个实例的最大连接数为单个分片的最大连接数 x 分片数量。标准架构分片数量为1。
	// <li>关闭副本只读：每个分片的最大连接数的取值范围为[10000,40000]。</li><li>开启副本只读：每个分片的最大连接数的取值范围为 [10000,10000 x (副本数 + 3)]。</li></ul>
	ClientLimit *int64 `json:"ClientLimit,omitempty" name:"ClientLimit"`
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
	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要修改的安全组ID列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 实例ID，格式如：kee-c1nl****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库引擎名称：keewidb。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要修改的安全组ID列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 实例ID，格式如：kee-c1nl****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyInstanceParamsRequestParams struct {
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例修改的参数列表。
	InstanceParams []*InstanceParam `json:"InstanceParams,omitempty" name:"InstanceParams"`
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例修改的参数列表。
	InstanceParams []*InstanceParam `json:"InstanceParams,omitempty" name:"InstanceParams"`
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
	// 修改是否成功。<ul><li>true：修改成功。</li><li>false：修改失败。</li></ul>
	Changed *bool `json:"Changed,omitempty" name:"Changed"`

	// 任务 ID。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyInstanceRequestParams struct {
	// 修改实例操作。<ul><li>rename：表示实例重命名。</li><li>modifyProject：修改实例所属项目。</li><li>modifyAutoRenew：修改实例续费模式。</li></ul>
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例 ID 数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例的新名称。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// 实例新的项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 包年包月计费的续费模式。<b>InstanceIds</b>数组和<b>AutoRenews</b>数组中的修改值对应。<ul><li>0：默认状态，指手动续费。</li><li>1：自动续费。</li><li>2：到期不再续费。</ul>
	AutoRenews []*int64 `json:"AutoRenews,omitempty" name:"AutoRenews"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 修改实例操作。<ul><li>rename：表示实例重命名。</li><li>modifyProject：修改实例所属项目。</li><li>modifyAutoRenew：修改实例续费模式。</li></ul>
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例 ID 数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 实例的新名称。
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// 实例新的项目 ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 包年包月计费的续费模式。<b>InstanceIds</b>数组和<b>AutoRenews</b>数组中的修改值对应。<ul><li>0：默认状态，指手动续费。</li><li>1：自动续费。</li><li>2：到期不再续费。</ul>
	AutoRenews []*int64 `json:"AutoRenews,omitempty" name:"AutoRenews"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维护时间窗起始时间，如：03:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 维护时间窗结束时间，如：06:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维护时间窗起始时间，如：03:00。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 维护时间窗结束时间，如：06:00。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	// 执行结果。<ul><li>success：修改成功。 </li> <li>failed：修改失败。</li></ul>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 操作类型。<ul><li>changeVip：修改实例私有网络。</li><li>changeVpc：修改实例私有网络所属子网。</li><li>changeBaseToVpc：基础网络转为私有网络。</li></ul>
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 修改后的 VIP 地址。
	// 当参数<b>Operation</b>为<b>changeVip</b>时，需设置实例修改后的 VIP 地址。该参数不配置，则自动分配。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 修改后的私有网络 ID。
	// 当参数<b>Operation</b>为<b>changeVip</b>或者为<b>changeBaseToVpc</b>时，务必设置实例修改后的私有网络 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 修改后的所属子网 ID。
	// 当参数<b>Operation</b>为<b>changeVpc</b>或者为<b>changeBaseToVpc</b>时，务必设置实例修改后的子网 ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 原 VIP 保留时长。
	// 单位：天。取值范围：0、1、2、3、7、15。
	Recycle *int64 `json:"Recycle,omitempty" name:"Recycle"`
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，如：kee-6ubh****。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 操作类型。<ul><li>changeVip：修改实例私有网络。</li><li>changeVpc：修改实例私有网络所属子网。</li><li>changeBaseToVpc：基础网络转为私有网络。</li></ul>
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 修改后的 VIP 地址。
	// 当参数<b>Operation</b>为<b>changeVip</b>时，需设置实例修改后的 VIP 地址。该参数不配置，则自动分配。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 修改后的私有网络 ID。
	// 当参数<b>Operation</b>为<b>changeVip</b>或者为<b>changeBaseToVpc</b>时，务必设置实例修改后的私有网络 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 修改后的所属子网 ID。
	// 当参数<b>Operation</b>为<b>changeVpc</b>或者为<b>changeBaseToVpc</b>时，务必设置实例修改后的子网 ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 原 VIP 保留时长。
	// 单位：天。取值范围：0、1、2、3、7、15。
	Recycle *int64 `json:"Recycle,omitempty" name:"Recycle"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkConfigResponseParams struct {
	// 执行状态。<ul><li>true：执行成功。</li><li>false：执行失败。</li></ul>
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 修改后的子网 ID。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 修改后的私有网络 ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 修改后的 VIP 地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type NodeInfo struct {
	// 节点类型，0 为主节点，1 为副本节点
	NodeType *int64 `json:"NodeType,omitempty" name:"NodeType"`

	// 主节点或者副本节点的ID，创建时不需要传递此参数。
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// 主节点或者副本节点的可用区ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 主节点或者副本节点的可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type Outbound struct {
	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ProductConf struct {
	// 产品类型，13-KeewiDB标准架构，14-KeewiDB集群架构
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// KeewiDB标准架构，KeewiDB集群架构
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 购买时的最小数量
	MinBuyNum *int64 `json:"MinBuyNum,omitempty" name:"MinBuyNum"`

	// 购买时的最大数量
	MaxBuyNum *int64 `json:"MaxBuyNum,omitempty" name:"MaxBuyNum"`

	// 产品是否售罄
	Saleout *bool `json:"Saleout,omitempty" name:"Saleout"`

	// 产品引擎，keewidb
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 兼容版本，Redis-2.8，Redis-3.2，Redis-4.0
	Version *string `json:"Version,omitempty" name:"Version"`

	// 副本数量
	ReplicaNum []*string `json:"ReplicaNum,omitempty" name:"ReplicaNum"`

	// 支持的计费模式，1-包年包月，0-按量计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type ProxyNodeInfo struct {
	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

type RedisNodeInfo struct {
	// 节点ID
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点角色
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// 分片ID
	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type RegionConf struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域简称
	RegionShortName *string `json:"RegionShortName,omitempty" name:"RegionShortName"`

	// 地域所在大区名称
	Area *string `json:"Area,omitempty" name:"Area"`

	// 可用区信息
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitempty" name:"ZoneSet"`
}

// Predefined struct for user
type RenewInstanceRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 购买时长。单位：月。取值为 [1,2,3,4,5,6,7,8,9,10,11,12,24,36,48,60]。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 购买时长。单位：月。取值为 [1,2,3,4,5,6,7,8,9,10,11,12,24,36,48,60]。
	Period *uint64 `json:"Period,omitempty" name:"Period"`
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
	delete(f, "InstanceId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewInstanceResponseParams struct {
	// 交易 ID。
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 节点 ID。
	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`

	// 节点组的名称，主节点为空。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 节点的可用区ID，比如ap-guangzhou-1。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 节点组角色。<ul><li>master：为主节点。</li><li>replica：为副本节点。</li></ul>
	Role *string `json:"Role,omitempty" name:"Role"`

	// 节点组节点列表。
	KeeWiDBNodes []*KeeWiDBNode `json:"KeeWiDBNodes,omitempty" name:"KeeWiDBNodes"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设置新密码。<ul><li>当参数<b>NoAuth</b>设置为<b>true</b>，切换为免密实例时，可不设置该参数。</li><li>密码复杂度要求：<ul><li>长度8 - 30位, 推荐使用12位以上的密码。</li><li>不能以"/"开头。</li>
	// <li>至少包含以下两项：<ul><li>小写字母a - z</li><li>大写字母A - Z</li><li>数字0 - 9</li><li>()~!@#$%^&*-+=_|{}[]:;<>,.?/</li></ul></li></ul></li></ul>
	Password *string `json:"Password,omitempty" name:"Password"`

	// 标识实例是否切换免密认证。<ul><li>false：由免密码认证方式切换为密码认证实例。默认为false。</li><li>true：由密码认证方式切换为免密码认证的方式。</li></ul>
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 设置新密码。<ul><li>当参数<b>NoAuth</b>设置为<b>true</b>，切换为免密实例时，可不设置该参数。</li><li>密码复杂度要求：<ul><li>长度8 - 30位, 推荐使用12位以上的密码。</li><li>不能以"/"开头。</li>
	// <li>至少包含以下两项：<ul><li>小写字母a - z</li><li>大写字母A - Z</li><li>数字0 - 9</li><li>()~!@#$%^&*-+=_|{}[]:;<>,.?/</li></ul></li></ul></li></ul>
	Password *string `json:"Password,omitempty" name:"Password"`

	// 标识实例是否切换免密认证。<ul><li>false：由免密码认证方式切换为密码认证实例。默认为false。</li><li>true：由密码认证方式切换为免密码认证的方式。</li></ul>
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
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
	// 任务 ID。
	// <b>说明：</b>修改密码时的任务ID，如果切换免密访问或者非免密码实例，则无需关注此返回值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type ResourceTag struct {
	// 标签 Key。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签 Value。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type SecurityGroup struct {
	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// 出站规则。
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`

	// 入站规则。
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`
}

// Predefined struct for user
type StartUpInstanceRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type StartUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartUpInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *StartUpInstanceResponseParams `json:"Response"`
}

func (r *StartUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskInfoDetail struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 操作者用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`
}

type TendisSlowLogDetail struct {
	// 执行时间
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`

	// 慢查询耗时（毫秒）
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 详细命令行信息
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// 节点ID
	Node *string `json:"Node,omitempty" name:"Node"`
}

type TradeDealDetail struct {
	// 订单号ID，调用云API时使用此ID	
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 长订单ID，反馈订单问题给官方客服使用此ID	
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 订单关联的实例数
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 创建用户uin
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 订单创建时间
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 订单超时时间
	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`

	// 订单完成时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 订单状态 1：未支付 2:已支付，未发货 3:发货中 4:发货成功 5:发货失败 6:已退款 7:已关闭订单 8:订单过期 9:订单已失效 10:产品已失效 11:代付拒绝 12:支付中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 订单状态描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 订单实际总价，单位：分
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置变更后，每个分片持久化内存容量，单位：GB。
	// <ul><li>KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// CPU 核数。
	MachineCpu *int64 `json:"MachineCpu,omitempty" name:"MachineCpu"`

	// 实例内存容量，单位：GB。
	// <ul><li>KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	MachineMemory *int64 `json:"MachineMemory,omitempty" name:"MachineMemory"`

	// 配置变更后，分片数量。
	// <ul><li>增加后分片的数量务必为增加之前数量的整数倍。分片数量支持选择3、5、6、8、9、10、12、15、16、18、20、21、24、25、27、30、32、33、35、36、39、40、42、45、48、50、51、54、55、56、57、60、63、64分片。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	ShardNum *int64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// 配置变更后，每个分片硬盘的容量。单位：GB。
	// <ul><li>每一缓存分片容量，对应的磁盘容量范围不同。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置变更后，每个分片持久化内存容量，单位：GB。
	// <ul><li>KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// CPU 核数。
	MachineCpu *int64 `json:"MachineCpu,omitempty" name:"MachineCpu"`

	// 实例内存容量，单位：GB。
	// <ul><li>KeeWiDB 内存容量<b>MachineMemory</b>与持久内存容量<b>MemSize</b>为固定搭配，即2GB内存，固定分配8GB的持久内存，不可选择。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	MachineMemory *int64 `json:"MachineMemory,omitempty" name:"MachineMemory"`

	// 配置变更后，分片数量。
	// <ul><li>增加后分片的数量务必为增加之前数量的整数倍。分片数量支持选择3、5、6、8、9、10、12、15、16、18、20、21、24、25、27、30、32、33、35、36、39、40、42、45、48、50、51、54、55、56、57、60、63、64分片。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	ShardNum *int64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// 配置变更后，每个分片硬盘的容量。单位：GB。
	// <ul><li>每一缓存分片容量，对应的磁盘容量范围不同。具体信息，请参见[产品规格](https://cloud.tencent.com/document/product/1520/80808)。</li><li>变更实例内存、持久化内存与磁盘、变更实例的分片数量，每次只能变更一项。</li></ul>
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
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
	delete(f, "MachineCpu")
	delete(f, "MachineMemory")
	delete(f, "ShardNum")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// 交易ID。
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type ZoneCapacityConf struct {
	// 可用区ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区是否售罄
	IsSaleout *bool `json:"IsSaleout,omitempty" name:"IsSaleout"`

	// 是否为默认可用区
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 网络类型：basenet -- 基础网络；vpcnet -- VPC网络
	NetWorkType []*string `json:"NetWorkType,omitempty" name:"NetWorkType"`

	// 产品规格等信息
	ProductSet []*ProductConf `json:"ProductSet,omitempty" name:"ProductSet"`

	// Int类型可用区ID
	OldZoneId *int64 `json:"OldZoneId,omitempty" name:"OldZoneId"`
}