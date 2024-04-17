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

package v20190719

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AutoSnapshotPolicyInfo struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略ID
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照策略创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 关联的文件系统个数
	FileSystemNums *uint64 `json:"FileSystemNums,omitnil,omitempty" name:"FileSystemNums"`

	// 快照定期备份在一星期哪一天，该参数与DayOfMonth,IntervalDays互斥
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 是否激活定期快照功能,1代表已激活，0代表未激活
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 下一次触发快照时间
	NextActiveTime *string `json:"NextActiveTime,omitnil,omitempty" name:"NextActiveTime"`

	// 快照策略状态，1代表快照策略状态正常。这里只有一种状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 帐号ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 保留时间
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 文件系统信息
	FileSystems []*FileSystemByPolicy `json:"FileSystems,omitnil,omitempty" name:"FileSystems"`

	// 快照定期备份在一个月的某个时间；该参数与DayOfWeek,IntervalDays互斥
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 快照定期间隔天数，1-365 天；该参数与DayOfMonth,DayOfWeek互斥
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`

	// 跨地域复制的快照保留时间，单位天
	CrossRegionsAliveDays *uint64 `json:"CrossRegionsAliveDays,omitnil,omitempty" name:"CrossRegionsAliveDays"`
}

type AvailableProtoStatus struct {
	// 售卖状态。可选值有 sale_out 售罄、saling可售、no_saling不可销售
	SaleStatus *string `json:"SaleStatus,omitnil,omitempty" name:"SaleStatus"`

	// 协议类型。可选值有 NFS、CIFS、TURBO
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type AvailableRegion struct {
	// 区域名称，如“ap-beijing”
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 区域名称，如“bj”
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 区域可用情况，当区域内至少有一个可用区处于可售状态时，取值为AVAILABLE，否则为UNAVAILABLE
	RegionStatus *string `json:"RegionStatus,omitnil,omitempty" name:"RegionStatus"`

	// 可用区数组
	Zones []*AvailableZone `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 区域中文名称，如“广州”
	RegionCnName *string `json:"RegionCnName,omitnil,omitempty" name:"RegionCnName"`
}

type AvailableType struct {
	// 协议与售卖详情
	Protocols []*AvailableProtoStatus `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// 存储类型。返回值中 SD 为通用标准型存储， HP为通用性能型存储， TB为Turbo标准型， TP 为Turbo性能型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否支持预付费。返回值中 true 为支持、false 为不支持
	Prepayment *bool `json:"Prepayment,omitnil,omitempty" name:"Prepayment"`
}

type AvailableZone struct {
	// 可用区名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区中文名称
	ZoneCnName *string `json:"ZoneCnName,omitnil,omitempty" name:"ZoneCnName"`

	// Type数组
	Types []*AvailableType `json:"Types,omitnil,omitempty" name:"Types"`

	// 可用区中英文名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

// Predefined struct for user
type BindAutoSnapshotPolicyRequestParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 文件系统列表
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 文件系统列表
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`
}

func (r *BindAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "FileSystemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *BindAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *BindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BucketInfo struct {
	// 桶名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 桶所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type CreateAutoSnapshotPolicyRequestParams struct {
	// 快照重复时间点,0-23
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照重复日期，星期一到星期日。 1代表星期一、7代表星期天
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照保留时长，单位天
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 快照按月重复，每月1-31号，选择一天，每月将在这一天自动创建快照。
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照重复时间点,0-23
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照重复日期，星期一到星期日。 1代表星期一、7代表星期天
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照保留时长，单位天
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 快照按月重复，每月1-31号，选择一天，每月将在这一天自动创建快照。
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

func (r *CreateAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Hour")
	delete(f, "PolicyName")
	delete(f, "DayOfWeek")
	delete(f, "AliveDays")
	delete(f, "DayOfMonth")
	delete(f, "IntervalDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *CreateAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsFileSystemRequestParams struct {
	// 可用区名称，例如ap-beijing-1，请参考 [概览](https://cloud.tencent.com/document/product/582/13225) 文档中的地域与可用区列表
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 网络类型，可选值为 VPC，CCN；其中 VPC 为私有网络， CCN 为云联网。通用标准型/性能型请选择VPC，Turbo标准型/性能型请选择CCN。
	NetInterface *string `json:"NetInterface,omitnil,omitempty" name:"NetInterface"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统协议类型， 值为 NFS、CIFS、TURBO ; 若留空则默认为 NFS协议，turbo系列必须选择turbo，不支持NFS、CIFS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值为 SD ；其中 SD 为通用标准型存储， HP为通用性能型存储， TB为Turbo标准型， TP 为Turbo性能型。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 私有网络（VPC） ID，若网络类型选择的是VPC，该字段为必填。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID，若网络类型选择的是VPC，该字段为必填。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP，Turbo系列当前不支持指定
	MountIP *string `json:"MountIP,omitnil,omitempty" name:"MountIP"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。用于保证请求幂等性的字符串失效时间为2小时。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云联网ID， 若网络类型选择的是CCN，该字段为必填
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 云联网中CFS使用的网段， 若网络类型选择的是Ccn，该字段为必填，且不能和Ccn中已经绑定的网段冲突
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 文件系统容量，turbo系列必填，单位为GiB。 turbo标准型单位GB，起售20TiB，即20480 GiB；扩容步长20TiB，即20480 GiB。turbo性能型起售10TiB，即10240 GiB；扩容步长10TiB，10240 GiB。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 定期快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 是否开启默认扩容，仅Turbo类型文件存储支持
	EnableAutoScaleUp *bool `json:"EnableAutoScaleUp,omitnil,omitempty" name:"EnableAutoScaleUp"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称，例如ap-beijing-1，请参考 [概览](https://cloud.tencent.com/document/product/582/13225) 文档中的地域与可用区列表
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 网络类型，可选值为 VPC，CCN；其中 VPC 为私有网络， CCN 为云联网。通用标准型/性能型请选择VPC，Turbo标准型/性能型请选择CCN。
	NetInterface *string `json:"NetInterface,omitnil,omitempty" name:"NetInterface"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统协议类型， 值为 NFS、CIFS、TURBO ; 若留空则默认为 NFS协议，turbo系列必须选择turbo，不支持NFS、CIFS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值为 SD ；其中 SD 为通用标准型存储， HP为通用性能型存储， TB为Turbo标准型， TP 为Turbo性能型。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 私有网络（VPC） ID，若网络类型选择的是VPC，该字段为必填。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID，若网络类型选择的是VPC，该字段为必填。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP，Turbo系列当前不支持指定
	MountIP *string `json:"MountIP,omitnil,omitempty" name:"MountIP"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。用于保证请求幂等性的字符串失效时间为2小时。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云联网ID， 若网络类型选择的是CCN，该字段为必填
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 云联网中CFS使用的网段， 若网络类型选择的是Ccn，该字段为必填，且不能和Ccn中已经绑定的网段冲突
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 文件系统容量，turbo系列必填，单位为GiB。 turbo标准型单位GB，起售20TiB，即20480 GiB；扩容步长20TiB，即20480 GiB。turbo性能型起售10TiB，即10240 GiB；扩容步长10TiB，10240 GiB。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 定期快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 是否开启默认扩容，仅Turbo类型文件存储支持
	EnableAutoScaleUp *bool `json:"EnableAutoScaleUp,omitnil,omitempty" name:"EnableAutoScaleUp"`
}

func (r *CreateCfsFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "NetInterface")
	delete(f, "PGroupId")
	delete(f, "Protocol")
	delete(f, "StorageType")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "MountIP")
	delete(f, "FsName")
	delete(f, "ResourceTags")
	delete(f, "ClientToken")
	delete(f, "CcnId")
	delete(f, "CidrBlock")
	delete(f, "Capacity")
	delete(f, "SnapshotId")
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "EnableAutoScaleUp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsFileSystemResponseParams struct {
	// 文件系统创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 用户自定义文件系统名称
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统状态，可能出现状态包括：“creating”  创建中, “create_failed” 创建失败, “available” 可用, “unserviced” 不可用, “upgrading” 升级中， “deleting” 删除中。
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// 文件系统已使用容量大小，单位为 Byte
	SizeByte *uint64 `json:"SizeByte,omitnil,omitempty" name:"SizeByte"`

	// 可用区 ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统是否加密
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsFileSystemResponseParams `json:"Response"`
}

func (r *CreateCfsFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsPGroupRequestParams struct {
	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

type CreateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

func (r *CreateCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DescInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsPGroupResponseParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限组描述信息
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`

	// 已经与该权限组绑定的文件系统个数
	BindCfsNum *int64 `json:"BindCfsNum,omitnil,omitempty" name:"BindCfsNum"`

	// 权限组创建时间
	CDate *string `json:"CDate,omitnil,omitempty" name:"CDate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsPGroupResponseParams `json:"Response"`
}

func (r *CreateCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsRuleRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 读写权限, 值为 RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为 all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 读写权限, 值为 RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为 all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`
}

func (r *CreateCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "AuthClientIp")
	delete(f, "Priority")
	delete(f, "RWPermission")
	delete(f, "UserPermission")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsRuleResponseParams struct {
	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 客户端 IP
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsRuleResponseParams `json:"Response"`
}

func (r *CreateCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsSnapshotRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type CreateCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

func (r *CreateCfsSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SnapshotName")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsSnapshotResponseParams struct {
	// 文件系统快照id
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCfsSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateCfsSnapshotResponseParams `json:"Response"`
}

func (r *CreateCfsSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCfsSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationTaskRequestParams struct {
	// 迁移任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 迁移方式标志位，默认为0。0: 桶迁移；1: 清单迁移
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// 迁移模式，默认为0。0: 全量迁移
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// 数据源账号的SecretId
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// 数据源账号的SecretKey
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`

	// 文件系统实例Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// 同名文件迁移时覆盖策略，默认为0。0: 最后修改时间优先；1: 全覆盖；2: 不覆盖
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 数据源服务商。COS: 腾讯云COS，OSS: 阿里云OSS，OBS:华为云OBS
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// 数据源桶名称，名称和地址至少有一个
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 数据源桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 数据源桶地址，名称和地址至少有一个
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 清单地址，迁移方式为清单迁移时必填
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// 目标文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 源桶路径，默认为/
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

type CreateMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 迁移方式标志位，默认为0。0: 桶迁移；1: 清单迁移
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// 迁移模式，默认为0。0: 全量迁移
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// 数据源账号的SecretId
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// 数据源账号的SecretKey
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`

	// 文件系统实例Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// 同名文件迁移时覆盖策略，默认为0。0: 最后修改时间优先；1: 全覆盖；2: 不覆盖
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 数据源服务商。COS: 腾讯云COS，OSS: 阿里云OSS，OBS:华为云OBS
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// 数据源桶名称，名称和地址至少有一个
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 数据源桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 数据源桶地址，名称和地址至少有一个
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 清单地址，迁移方式为清单迁移时必填
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// 目标文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 源桶路径，默认为/
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

func (r *CreateMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "MigrationType")
	delete(f, "MigrationMode")
	delete(f, "SrcSecretId")
	delete(f, "SrcSecretKey")
	delete(f, "FileSystemId")
	delete(f, "FsPath")
	delete(f, "CoverType")
	delete(f, "SrcService")
	delete(f, "BucketName")
	delete(f, "BucketRegion")
	delete(f, "BucketAddress")
	delete(f, "ListAddress")
	delete(f, "FsName")
	delete(f, "BucketPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationTaskResponseParams struct {
	// 迁移任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrationTaskResponseParams `json:"Response"`
}

func (r *CreateMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPolicyRequestParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type DeleteAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

func (r *DeleteAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *DeleteAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsFileSystemRequestParams struct {
	// 文件系统 ID。说明，进行删除文件系统操作前需要先调用 DeleteMountTarget 接口删除该文件系统的挂载点，否则会删除失败。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID。说明，进行删除文件系统操作前需要先调用 DeleteMountTarget 接口删除该文件系统的挂载点，否则会删除失败。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DeleteCfsFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsFileSystemResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsFileSystemResponseParams `json:"Response"`
}

func (r *DeleteCfsFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsPGroupRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

type DeleteCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

func (r *DeleteCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsPGroupResponseParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 用户 ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsPGroupResponseParams `json:"Response"`
}

func (r *DeleteCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsRuleRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsRuleResponseParams struct {
	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsRuleResponseParams `json:"Response"`
}

func (r *DeleteCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsSnapshotRequestParams struct {
	// 文件系统快照id
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 需要删除的文件文件系统快照ID 列表，快照ID，跟ID列表至少填一项
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

type DeleteCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统快照id
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 需要删除的文件文件系统快照ID 列表，快照ID，跟ID列表至少填一项
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

func (r *DeleteCfsSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCfsSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCfsSnapshotResponseParams struct {
	// 文件系统ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCfsSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCfsSnapshotResponseParams `json:"Response"`
}

func (r *DeleteCfsSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCfsSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrationTaskRequestParams struct {
	// 迁移任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrationTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMigrationTaskResponseParams `json:"Response"`
}

func (r *DeleteMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountTargetRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 挂载点 ID
	MountTargetId *string `json:"MountTargetId,omitnil,omitempty" name:"MountTargetId"`
}

type DeleteMountTargetRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 挂载点 ID
	MountTargetId *string `json:"MountTargetId,omitnil,omitempty" name:"MountTargetId"`
}

func (r *DeleteMountTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "MountTargetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMountTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountTargetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMountTargetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMountTargetResponseParams `json:"Response"`
}

func (r *DeleteMountTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserQuotaRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteUserQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "UserType")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserQuotaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserQuotaResponseParams `json:"Response"`
}

func (r *DeleteUserQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesRequestParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面长
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 升序，降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面长
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 升序，降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`
}

func (r *DescribeAutoSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "OrderField")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoSnapshotPoliciesResponseParams struct {
	// 快照策略总个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 快照策略信息
	AutoSnapshotPolicies []*AutoSnapshotPolicyInfo `json:"AutoSnapshotPolicies,omitnil,omitempty" name:"AutoSnapshotPolicies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAutoSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableZoneInfoRequestParams struct {

}

type DescribeAvailableZoneInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailableZoneInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableZoneInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailableZoneInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailableZoneInfoResponseParams struct {
	// 各可用区的资源售卖情况以及支持的存储类型、存储协议等信息
	RegionZones []*AvailableRegion `json:"RegionZones,omitnil,omitempty" name:"RegionZones"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAvailableZoneInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailableZoneInfoResponseParams `json:"Response"`
}

func (r *DescribeAvailableZoneInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailableZoneInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBucketListRequestParams struct {
	// 数据源服务商。COS: 腾讯云COS，OSS: 阿里云OSS，OBS:华为云OBS
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// 数据源账号的SecretId
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// 数据源账号的SecretKey
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`
}

type DescribeBucketListRequest struct {
	*tchttp.BaseRequest
	
	// 数据源服务商。COS: 腾讯云COS，OSS: 阿里云OSS，OBS:华为云OBS
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// 数据源账号的SecretId
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// 数据源账号的SecretKey
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`
}

func (r *DescribeBucketListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBucketListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcService")
	delete(f, "SrcSecretId")
	delete(f, "SrcSecretKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBucketListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBucketListResponseParams struct {
	// 桶的数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 桶列表
	BucketList []*BucketInfo `json:"BucketList,omitnil,omitempty" name:"BucketList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBucketListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBucketListResponseParams `json:"Response"`
}

func (r *DescribeBucketListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBucketListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemClientsRequestParams struct {
	// 文件系统 ID。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCfsFileSystemClientsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeCfsFileSystemClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemClientsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemClientsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemClientsResponseParams struct {
	// 客户端列表
	ClientList []*FileSystemClient `json:"ClientList,omitnil,omitempty" name:"ClientList"`

	// 文件系统总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsFileSystemClientsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsFileSystemClientsResponseParams `json:"Response"`
}

func (r *DescribeCfsFileSystemClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemsRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 私有网络（VPC） ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户自定义名称
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`
}

type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 私有网络（VPC） ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户自定义名称
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`
}

func (r *DescribeCfsFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreationToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemsResponseParams struct {
	// 文件系统信息
	FileSystems []*FileSystemInfo `json:"FileSystems,omitnil,omitempty" name:"FileSystems"`

	// 文件系统总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsFileSystemsResponseParams `json:"Response"`
}

func (r *DescribeCfsFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsPGroupsRequestParams struct {

}

type DescribeCfsPGroupsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfsPGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsPGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsPGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsPGroupsResponseParams struct {
	// 权限组信息列表
	PGroupList []*PGroupInfo `json:"PGroupList,omitnil,omitempty" name:"PGroupList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsPGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsPGroupsResponseParams `json:"Response"`
}

func (r *DescribeCfsPGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsPGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsRulesRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

type DescribeCfsRulesRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

func (r *DescribeCfsRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsRulesResponseParams struct {
	// 权限组规则列表
	RuleList []*PGroupRuleInfo `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsRulesResponseParams `json:"Response"`
}

func (r *DescribeCfsRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsServiceStatusRequestParams struct {

}

type DescribeCfsServiceStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfsServiceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsServiceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsServiceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsServiceStatusResponseParams struct {
	// 该用户当前 CFS 服务的状态，none 为未开通，creating 为开通中，created 为已开通
	CfsServiceStatus *string `json:"CfsServiceStatus,omitnil,omitempty" name:"CfsServiceStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsServiceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsServiceStatusResponseParams `json:"Response"`
}

func (r *DescribeCfsServiceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotOverviewRequestParams struct {

}

type DescribeCfsSnapshotOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfsSnapshotOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsSnapshotOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotOverviewResponseParams struct {
	// 统计信息
	StatisticsList []*SnapshotStatistics `json:"StatisticsList,omitnil,omitempty" name:"StatisticsList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsSnapshotOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsSnapshotOverviewResponseParams `json:"Response"`
}

func (r *DescribeCfsSnapshotOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotsRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 分页起始位置，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面长度，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。
	// <br>SnapshotId - Array of String - 是否必填：否 -（过滤条件）按快照ID过滤。
	// <br>SnapshotName - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。
	// <br>FileSystemId - Array of String - 是否必填：否 -（过滤条件）按文件系统ID过滤。
	// <br>FsName - Array of String - 是否必填：否 -（过滤条件）按文件系统名过滤。
	// <br>Status - Array of String - 是否必填：否 -（过滤条件）按按照快照状态过滤。(creating：表示创建中 | available：表示可用。| rollbacking：表示回滚。| rollbacking_new：表示由快照创建新文件系统中。
	// <br>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。
	// <br>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序取值
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序 升序或者降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeCfsSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 分页起始位置，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面长度，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件。
	// <br>SnapshotId - Array of String - 是否必填：否 -（过滤条件）按快照ID过滤。
	// <br>SnapshotName - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。
	// <br>FileSystemId - Array of String - 是否必填：否 -（过滤条件）按文件系统ID过滤。
	// <br>FsName - Array of String - 是否必填：否 -（过滤条件）按文件系统名过滤。
	// <br>Status - Array of String - 是否必填：否 -（过滤条件）按按照快照状态过滤。(creating：表示创建中 | available：表示可用。| rollbacking：表示回滚。| rollbacking_new：表示由快照创建新文件系统中。
	// <br>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。
	// <br>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序取值
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序 升序或者降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeCfsSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SnapshotId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsSnapshotsResponseParams struct {
	// 总个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 快照信息描述
	Snapshots []*SnapshotInfo `json:"Snapshots,omitnil,omitempty" name:"Snapshots"`

	// 快照列表快照汇总
	TotalSize *uint64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfsSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfsSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeCfsSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfsSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTasksRequestParams struct {
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <br><li> taskId
	// 
	// 按照【迁移任务id】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> taskName
	// 
	// 按照【迁移任务名字】进行模糊搜索过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMigrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <br><li> taskId
	// 
	// 按照【迁移任务id】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> taskName
	// 
	// 按照【迁移任务名字】进行模糊搜索过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为100。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMigrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTasksResponseParams struct {
	// 迁移任务的数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 迁移任务详情
	MigrationTasks []*MigrationTaskInfo `json:"MigrationTasks,omitnil,omitempty" name:"MigrationTasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationTasksResponseParams `json:"Response"`
}

func (r *DescribeMigrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountTargetsRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeMountTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountTargetsResponseParams struct {
	// 挂载点详情
	MountTargets []*MountInfo `json:"MountTargets,omitnil,omitempty" name:"MountTargets"`

	// 挂载点数量
	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitnil,omitempty" name:"NumberOfMountTargets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountTargetsResponseParams `json:"Response"`
}

func (r *DescribeMountTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOperationLogsRequestParams struct {
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeSnapshotOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeSnapshotOperationLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOperationLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotOperationLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotOperationLogsResponseParams struct {
	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 操作日志
	SnapshotOperates []*SnapshotOperateLog `json:"SnapshotOperates,omitnil,omitempty" name:"SnapshotOperates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSnapshotOperationLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotOperationLogsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotOperationLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotOperationLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotaRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 过滤条件。
	// <br><li>UserType - Array of String - 是否必填：否 -（过滤条件）按配额类型过滤。(Uid| Gid )
	// <br><li>UserId - Array of String - 是否必填：否 -（过滤条件）按UID/GID过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，可填范围为大于0的整数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 过滤条件。
	// <br><li>UserType - Array of String - 是否必填：否 -（过滤条件）按配额类型过滤。(Uid| Gid )
	// <br><li>UserId - Array of String - 是否必填：否 -（过滤条件）按UID/GID过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，可填范围为大于0的整数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeUserQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserQuotaResponseParams struct {
	// UserQuota条目总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// UserQuota条目
	UserQuotaInfo []*UserQuota `json:"UserQuotaInfo,omitnil,omitempty" name:"UserQuotaInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserQuotaResponseParams `json:"Response"`
}

func (r *DescribeUserQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystemByPolicy struct {
	// 文件系统名称
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统大小，单位Byte
	SizeByte *uint64 `json:"SizeByte,omitnil,omitempty" name:"SizeByte"`

	// 存储类型，HP：通用性能型；SD：通用标准型；TP:turbo性能型；TB：turbo标准型；THP：吞吐型
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 快照总大小，单位GiB
	TotalSnapshotSize *uint64 `json:"TotalSnapshotSize,omitnil,omitempty" name:"TotalSnapshotSize"`

	// 文件系统创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 文件系统所在区ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type FileSystemClient struct {
	// 文件系统IP地址
	CfsVip *string `json:"CfsVip,omitnil,omitempty" name:"CfsVip"`

	// 客户端IP地址
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 文件系统所属VPCID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区名称，例如ap-beijing-1，请参考 概览文档中的地域与可用区列表
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 该文件系统被挂载到客户端上的路径信息
	MountDirectory *string `json:"MountDirectory,omitnil,omitempty" name:"MountDirectory"`
}

type FileSystemInfo struct {
	// 创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 用户自定义名称
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统状态。取值范围：
	// - creating:创建中
	// - mounting:挂载中
	// - create_failed:创建失败
	// - available:可使用
	// - unserviced:停服中
	// - upgrading:升级中
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// 文件系统已使用容量。单位：Byte
	SizeByte *uint64 `json:"SizeByte,omitnil,omitempty" name:"SizeByte"`

	// 文件系统空间限制。单位:GiB
	SizeLimit *uint64 `json:"SizeLimit,omitnil,omitempty" name:"SizeLimit"`

	// 区域 ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 区域名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 文件系统协议类型, 支持 NFS,CIFS,TURBO
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 存储类型，HP：通用性能型；SD：通用标准型；TP:turbo性能型；TB：turbo标准型；THP：吞吐型
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 文件系统绑定的预付费存储包
	StorageResourcePkg *string `json:"StorageResourcePkg,omitnil,omitempty" name:"StorageResourcePkg"`

	// 文件系统绑定的预付费带宽包（暂未支持）
	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitnil,omitempty" name:"BandwidthResourcePkg"`

	// 文件系统绑定权限组信息
	PGroup *PGroup `json:"PGroup,omitnil,omitempty" name:"PGroup"`

	// 用户自定义名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统是否加密,true：代表加密，false：非加密
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// 加密所使用的密钥，可以为密钥的 ID 或者 ARN
	KmsKeyId *string `json:"KmsKeyId,omitnil,omitempty" name:"KmsKeyId"`

	// 应用ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 文件系统吞吐上限，吞吐上限是根据文件系统当前已使用存储量、绑定的存储资源包以及吞吐资源包一同确定. 单位MiB/s
	BandwidthLimit *float64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`

	// 文件系统容量规格上限
	// 单位:GiB
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 文件系统标签列表
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 文件系统生命周期管理状态
	// NotAvailable：不可用
	// Available:可用
	TieringState *string `json:"TieringState,omitnil,omitempty" name:"TieringState"`

	// 分层存储详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TieringDetail *TieringDetailInfo `json:"TieringDetail,omitnil,omitempty" name:"TieringDetail"`
}

type Filter struct {
	// 值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type MigrationTaskInfo struct {
	// 迁移任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 迁移任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 迁移方式标志位，默认为0。0: 桶迁移；1: 清单迁移
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// 迁移模式，默认为0。0: 全量迁移
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// 数据源桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 数据源桶地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 数据源桶地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 清单地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// 文件系统实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统实例Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// 同名文件迁移时覆盖策略，默认为0。0: 最后修改时间优先；1: 全覆盖；2: 不覆盖
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 完成/终止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 迁移状态。0: 已完成；1: 进行中；2: 已终止
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileTotalCount *uint64 `json:"FileTotalCount,omitnil,omitempty" name:"FileTotalCount"`

	// 已迁移文件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMigratedCount *uint64 `json:"FileMigratedCount,omitnil,omitempty" name:"FileMigratedCount"`

	// 迁移失败文件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileFailedCount *uint64 `json:"FileFailedCount,omitnil,omitempty" name:"FileFailedCount"`

	// 文件容量，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileTotalSize *int64 `json:"FileTotalSize,omitnil,omitempty" name:"FileTotalSize"`

	// 已迁移文件容量，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMigratedSize *int64 `json:"FileMigratedSize,omitnil,omitempty" name:"FileMigratedSize"`

	// 迁移失败文件容量，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileFailedSize *int64 `json:"FileFailedSize,omitnil,omitempty" name:"FileFailedSize"`

	// 全部清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileTotalList *string `json:"FileTotalList,omitnil,omitempty" name:"FileTotalList"`

	// 已完成文件清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileCompletedList *string `json:"FileCompletedList,omitnil,omitempty" name:"FileCompletedList"`

	// 失败文件清单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileFailedList *string `json:"FileFailedList,omitnil,omitempty" name:"FileFailedList"`

	// 源桶路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`
}

// Predefined struct for user
type ModifyFileSystemAutoScaleUpRuleRequestParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容阈值，范围[10-90]
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// 扩容后目标阈值,范围[10-90],该值要小于ScaleUpThreshold
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// 规则状态0:关闭，1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyFileSystemAutoScaleUpRuleRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容阈值，范围[10-90]
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// 扩容后目标阈值,范围[10-90],该值要小于ScaleUpThreshold
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// 规则状态0:关闭，1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyFileSystemAutoScaleUpRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemAutoScaleUpRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "ScaleUpThreshold")
	delete(f, "TargetThreshold")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFileSystemAutoScaleUpRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemAutoScaleUpRuleResponseParams struct {
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 规则状态0:关闭，1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 扩容阈值,范围[10-90]
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// 扩容后达到阈值,范围[10-90]
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFileSystemAutoScaleUpRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFileSystemAutoScaleUpRuleResponseParams `json:"Response"`
}

func (r *ModifyFileSystemAutoScaleUpRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemAutoScaleUpRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountInfo struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 挂载点 ID
	MountTargetId *string `json:"MountTargetId,omitnil,omitempty" name:"MountTargetId"`

	// 挂载点 IP
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`

	// 挂载根目录
	FSID *string `json:"FSID,omitnil,omitempty" name:"FSID"`

	// 挂载点状态
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// 网络类型
	NetworkInterface *string `json:"NetworkInterface,omitnil,omitempty" name:"NetworkInterface"`

	// 私有网络 ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 子网 Id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// CFS Turbo使用的云联网ID
	CcnID *string `json:"CcnID,omitnil,omitempty" name:"CcnID"`

	// 云联网中CFS Turbo使用的网段
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`
}

type PGroup struct {
	// 权限组ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type PGroupInfo struct {
	// 权限组ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`

	// 创建时间
	CDate *string `json:"CDate,omitnil,omitempty" name:"CDate"`

	// 关联文件系统个数
	BindCfsNum *int64 `json:"BindCfsNum,omitnil,omitempty" name:"BindCfsNum"`
}

type PGroupRuleInfo struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 允许访问的客户端IP
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限, ro为只读，rw为读写
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 规则优先级，1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

// Predefined struct for user
type ScaleUpFileSystemRequestParams struct {
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容的目标容量（单位GiB）
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`
}

type ScaleUpFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容的目标容量（单位GiB）
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`
}

func (r *ScaleUpFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "TargetCapacity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpFileSystemResponseParams struct {
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容的目标容量（单位GiB）
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpFileSystemResponseParams `json:"Response"`
}

func (r *ScaleUpFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserQuotaRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitnil,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个
	FileHardLimit *uint64 `json:"FileHardLimit,omitnil,omitempty" name:"FileHardLimit"`
}

type SetUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitnil,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个
	FileHardLimit *uint64 `json:"FileHardLimit,omitnil,omitempty" name:"FileHardLimit"`
}

func (r *SetUserQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "UserType")
	delete(f, "UserId")
	delete(f, "CapacityHardLimit")
	delete(f, "FileHardLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetUserQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserQuotaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetUserQuotaResponse struct {
	*tchttp.BaseResponse
	Response *SetUserQuotaResponseParams `json:"Response"`
}

func (r *SetUserQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetUserQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignUpCfsServiceRequestParams struct {

}

type SignUpCfsServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *SignUpCfsServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignUpCfsServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignUpCfsServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignUpCfsServiceResponseParams struct {
	// 该用户当前 CFS 服务的状态，creating 是开通中，created 是已开通
	CfsServiceStatus *string `json:"CfsServiceStatus,omitnil,omitempty" name:"CfsServiceStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SignUpCfsServiceResponse struct {
	*tchttp.BaseResponse
	Response *SignUpCfsServiceResponseParams `json:"Response"`
}

func (r *SignUpCfsServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignUpCfsServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotInfo struct {
	// 创建快照时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 快照状态，createing-创建中；available-运行中；deleting-删除中；rollbacking-new 创建新文件系统中；create-failed 创建失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照大小
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 保留时长天
	AliveDay *uint64 `json:"AliveDay,omitnil,omitempty" name:"AliveDay"`

	// 快照进度百分比，1表示1%
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 帐号ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 快照删除时间
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 快照标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 快照类型, general为通用系列快照，turbo为Turbo系列快照
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotType *string `json:"SnapshotType,omitnil,omitempty" name:"SnapshotType"`

	// 实际快照时间，反应快照对应文件系统某个时刻的数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotTime *string `json:"SnapshotTime,omitnil,omitempty" name:"SnapshotTime"`
}

type SnapshotOperateLog struct {
	// 操作类型
	// CreateCfsSnapshot：创建快照
	// DeleteCfsSnapshot：删除快照
	// CreateCfsFileSystem：创建文件系统
	// UpdateCfsSnapshotAttribute：更新快照
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 操作时间
	ActionTime *string `json:"ActionTime,omitnil,omitempty" name:"ActionTime"`

	// 操作名称
	// CreateCfsSnapshot
	// DeleteCfsSnapshot
	// CreateCfsFileSystem
	// UpdateCfsSnapshotAttribute
	ActionName *string `json:"ActionName,omitnil,omitempty" name:"ActionName"`

	// 操作者uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 1-任务进行中；2-任务成功；3-任务失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`
}

type SnapshotStatistics struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 快照总个数
	SnapshotNumber *uint64 `json:"SnapshotNumber,omitnil,omitempty" name:"SnapshotNumber"`

	// 快照总容量
	SnapshotSize *uint64 `json:"SnapshotSize,omitnil,omitempty" name:"SnapshotSize"`
}

// Predefined struct for user
type StopMigrationTaskRequestParams struct {
	// 迁移任务名称
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务名称
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrationTaskResponseParams struct {
	// 迁移任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 迁移状态。0: 已完成；1: 进行中；2: 已终止
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrationTaskResponseParams `json:"Response"`
}

func (r *StopMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TieringDetailInfo struct {
	// 低频存储容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TieringSizeInBytes *int64 `json:"TieringSizeInBytes,omitnil,omitempty" name:"TieringSizeInBytes"`
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyRequestParams struct {
	// 需要解绑的文件系统ID列表，用"," 分割
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`

	// 解绑的快照ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 需要解绑的文件系统ID列表，用"," 分割
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`

	// 解绑的快照ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

func (r *UnbindAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemIds")
	delete(f, "AutoSnapshotPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UnbindAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *UnbindAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAutoSnapshotPolicyRequestParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照定期备份，按照星期一到星期日。 1代表星期一，7代表星期日
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 快照保留日期
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 是否激活定期快照功能；1代表激活，0代表未激活
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 定期快照在每月的第几天创建快照，该参数与DayOfWeek互斥
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数定期执行快照，该参数与DayOfWeek,DayOfMonth 互斥
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type UpdateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照定期备份，按照星期一到星期日。 1代表星期一，7代表星期日
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 快照保留日期
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 是否激活定期快照功能；1代表激活，0代表未激活
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 定期快照在每月的第几天创建快照，该参数与DayOfWeek互斥
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数定期执行快照，该参数与DayOfWeek,DayOfMonth 互斥
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

func (r *UpdateAutoSnapshotPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAutoSnapshotPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoSnapshotPolicyId")
	delete(f, "PolicyName")
	delete(f, "DayOfWeek")
	delete(f, "Hour")
	delete(f, "AliveDays")
	delete(f, "IsActivated")
	delete(f, "DayOfMonth")
	delete(f, "IntervalDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAutoSnapshotPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAutoSnapshotPolicyResponseParams `json:"Response"`
}

func (r *UpdateAutoSnapshotPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAutoSnapshotPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemNameRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`
}

func (r *UpdateCfsFileSystemNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FsName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsFileSystemNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemNameResponseParams struct {
	// 用户自定义文件系统名称
	CreationToken *string `json:"CreationToken,omitnil,omitempty" name:"CreationToken"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsFileSystemNameResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsFileSystemNameResponseParams `json:"Response"`
}

func (r *UpdateCfsFileSystemNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemPGroupRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsFileSystemPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemPGroupResponseParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsFileSystemPGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsFileSystemPGroupResponseParams `json:"Response"`
}

func (r *UpdateCfsFileSystemPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemSizeLimitRequestParams struct {
	// 文件系统容量限制大小，输入范围0-1073741824, 单位为GB；其中输入值为0时，表示不限制文件系统容量。
	FsLimit *uint64 `json:"FsLimit,omitnil,omitempty" name:"FsLimit"`

	// 文件系统ID，目前仅支持标准型文件系统。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemSizeLimitRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统容量限制大小，输入范围0-1073741824, 单位为GB；其中输入值为0时，表示不限制文件系统容量。
	FsLimit *uint64 `json:"FsLimit,omitnil,omitempty" name:"FsLimit"`

	// 文件系统ID，目前仅支持标准型文件系统。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *UpdateCfsFileSystemSizeLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemSizeLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FsLimit")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsFileSystemSizeLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsFileSystemSizeLimitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsFileSystemSizeLimitResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsFileSystemSizeLimitResponseParams `json:"Response"`
}

func (r *UpdateCfsFileSystemSizeLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsFileSystemSizeLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsPGroupRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符。 Name和Descinfo不能同时为空
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

type UpdateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符。 Name和Descinfo不能同时为空
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

func (r *UpdateCfsPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "Name")
	delete(f, "DescInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsPGroupResponseParams struct {
	// 权限组ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsPGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsPGroupResponseParams `json:"Response"`
}

func (r *UpdateCfsPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsRuleRequestParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限, 值为RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type UpdateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限, 值为RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

func (r *UpdateCfsRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PGroupId")
	delete(f, "RuleId")
	delete(f, "AuthClientIp")
	delete(f, "RWPermission")
	delete(f, "UserPermission")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsRuleResponseParams struct {
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 允许访问的客户端 IP 或者 IP 段
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsRuleResponseParams `json:"Response"`
}

func (r *UpdateCfsRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsSnapshotAttributeRequestParams struct {
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 文件系统快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 文件系统快照保留天数
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`
}

type UpdateCfsSnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 文件系统快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 文件系统快照保留天数
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`
}

func (r *UpdateCfsSnapshotAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsSnapshotAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotId")
	delete(f, "SnapshotName")
	delete(f, "AliveDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCfsSnapshotAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCfsSnapshotAttributeResponseParams struct {
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCfsSnapshotAttributeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCfsSnapshotAttributeResponseParams `json:"Response"`
}

func (r *UpdateCfsSnapshotAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCfsSnapshotAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFileSystemBandwidthLimitRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统带宽，仅吞吐型可填。单位MiB/s，最小为1GiB/s，最大200GiB/s。
	BandwidthLimit *int64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`
}

type UpdateFileSystemBandwidthLimitRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统带宽，仅吞吐型可填。单位MiB/s，最小为1GiB/s，最大200GiB/s。
	BandwidthLimit *int64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`
}

func (r *UpdateFileSystemBandwidthLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFileSystemBandwidthLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "BandwidthLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateFileSystemBandwidthLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateFileSystemBandwidthLimitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateFileSystemBandwidthLimitResponse struct {
	*tchttp.BaseResponse
	Response *UpdateFileSystemBandwidthLimitResponseParams `json:"Response"`
}

func (r *UpdateFileSystemBandwidthLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateFileSystemBandwidthLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserQuota struct {
	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitnil,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个
	FileHardLimit *uint64 `json:"FileHardLimit,omitnil,omitempty" name:"FileHardLimit"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 容量使用，单位GiB
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapacityUsed *uint64 `json:"CapacityUsed,omitnil,omitempty" name:"CapacityUsed"`

	// 文件使用个数，单位个
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUsed *uint64 `json:"FileUsed,omitnil,omitempty" name:"FileUsed"`
}