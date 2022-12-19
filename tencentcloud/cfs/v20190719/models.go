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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AutoSnapshotPolicyInfo struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略ID
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 快照策略创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 关联的文件系统个数
	FileSystemNums *uint64 `json:"FileSystemNums,omitempty" name:"FileSystemNums"`

	// 快照定期备份在一星期哪一天
	DayOfWeek *string `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitempty" name:"Hour"`

	// 是否激活定期快照功能
	IsActivated *uint64 `json:"IsActivated,omitempty" name:"IsActivated"`

	// 下一次触发快照时间
	NextActiveTime *string `json:"NextActiveTime,omitempty" name:"NextActiveTime"`

	// 快照策略状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 帐号ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 保留时间
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`

	// 地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 文件系统信息
	FileSystems []*FileSystemByPolicy `json:"FileSystems,omitempty" name:"FileSystems"`
}

type AvailableProtoStatus struct {
	// 售卖状态。可选值有 sale_out 售罄、saling可售、no_saling不可销售
	SaleStatus *string `json:"SaleStatus,omitempty" name:"SaleStatus"`

	// 协议类型。可选值有 NFS、CIFS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type AvailableRegion struct {
	// 区域名称，如“ap-beijing”
	Region *string `json:"Region,omitempty" name:"Region"`

	// 区域名称，如“bj”
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 区域可用情况，当区域内至少有一个可用区处于可售状态时，取值为AVAILABLE，否则为UNAVAILABLE
	RegionStatus *string `json:"RegionStatus,omitempty" name:"RegionStatus"`

	// 可用区数组
	Zones []*AvailableZone `json:"Zones,omitempty" name:"Zones"`

	// 区域中文名称，如“广州”
	RegionCnName *string `json:"RegionCnName,omitempty" name:"RegionCnName"`
}

type AvailableType struct {
	// 协议与售卖详情
	Protocols []*AvailableProtoStatus `json:"Protocols,omitempty" name:"Protocols"`

	// 存储类型。返回值中 SD 为标准型存储、HP 为性能型存储
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否支持预付费。返回值中 true 为支持、false 为不支持
	Prepayment *bool `json:"Prepayment,omitempty" name:"Prepayment"`
}

type AvailableZone struct {
	// 可用区名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区中文名称
	ZoneCnName *string `json:"ZoneCnName,omitempty" name:"ZoneCnName"`

	// Type数组
	Types []*AvailableType `json:"Types,omitempty" name:"Types"`

	// 可用区中英文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

// Predefined struct for user
type BindAutoSnapshotPolicyRequestParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 文件系统列表
	FileSystemIds *string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 文件系统列表
	FileSystemIds *string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`
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
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateAutoSnapshotPolicyRequestParams struct {
	// 快照重复日期，星期一到星期日
	DayOfWeek *string `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// 快照重复时间点
	Hour *string `json:"Hour,omitempty" name:"Hour"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 快照保留时长
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照重复日期，星期一到星期日
	DayOfWeek *string `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// 快照重复时间点
	Hour *string `json:"Hour,omitempty" name:"Hour"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 快照保留时长
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`
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
	delete(f, "DayOfWeek")
	delete(f, "Hour")
	delete(f, "PolicyName")
	delete(f, "AliveDays")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 网络类型，可选值为 VPC，BASIC，CCN；其中 VPC 为私有网络，BASIC 为基础网络, CCN 为云联网，Turbo系列当前必须选择云联网。目前基础网络已逐渐淘汰，不推荐使用。
	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`

	// 权限组 ID，通用标准型和性能型必填，turbo系列请填写pgroupbasic
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 文件系统协议类型， 值为 NFS、CIFS、TURBO ; 若留空则默认为 NFS协议，turbo系列必须选择turbo，不支持NFS、CIFS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值为 SD ；其中 SD 为通用标准型标准型存储， HP为通用性能型存储， TB为turbo标准型， TP 为turbo性能型。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 私有网络（VPC） ID，若网络类型选择的是VPC，该字段为必填。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID，若网络类型选择的是VPC，该字段为必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP，Turbo系列当前不支持指定
	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// 文件系统标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。用于保证请求幂等性的字符串失效时间为2小时。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 云联网ID， 若网络类型选择的是CCN，该字段为必填
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 云联网中CFS使用的网段， 若网络类型选择的是Ccn，该字段为必填，且不能和Ccn中已经绑定的网段冲突
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 文件系统容量，turbo系列必填，单位为GiB。 turbo标准型单位GB，起售40TiB，即40960 GiB；扩容步长20TiB，即20480 GiB。turbo性能型起售20TiB，即20480 GiB；扩容步长10TiB，10240 GiB。
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称，例如ap-beijing-1，请参考 [概览](https://cloud.tencent.com/document/product/582/13225) 文档中的地域与可用区列表
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 网络类型，可选值为 VPC，BASIC，CCN；其中 VPC 为私有网络，BASIC 为基础网络, CCN 为云联网，Turbo系列当前必须选择云联网。目前基础网络已逐渐淘汰，不推荐使用。
	NetInterface *string `json:"NetInterface,omitempty" name:"NetInterface"`

	// 权限组 ID，通用标准型和性能型必填，turbo系列请填写pgroupbasic
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 文件系统协议类型， 值为 NFS、CIFS、TURBO ; 若留空则默认为 NFS协议，turbo系列必须选择turbo，不支持NFS、CIFS
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值为 SD ；其中 SD 为通用标准型标准型存储， HP为通用性能型存储， TB为turbo标准型， TP 为turbo性能型。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 私有网络（VPC） ID，若网络类型选择的是VPC，该字段为必填。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID，若网络类型选择的是VPC，该字段为必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP，Turbo系列当前不支持指定
	MountIP *string `json:"MountIP,omitempty" name:"MountIP"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// 文件系统标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。用于保证请求幂等性的字符串失效时间为2小时。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 云联网ID， 若网络类型选择的是CCN，该字段为必填
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 云联网中CFS使用的网段， 若网络类型选择的是Ccn，该字段为必填，且不能和Ccn中已经绑定的网段冲突
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 文件系统容量，turbo系列必填，单位为GiB。 turbo标准型单位GB，起售40TiB，即40960 GiB；扩容步长20TiB，即20480 GiB。turbo性能型起售20TiB，即20480 GiB；扩容步长10TiB，10240 GiB。
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCfsFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCfsFileSystemResponseParams struct {
	// 文件系统创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 用户自定义文件系统名称
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 文件系统状态，可能出现状态包括：“creating”  创建中, “create_failed” 创建失败, “available” 可用, “unserviced” 不可用, “upgrading” 升级中， “deleting” 删除中。
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// 文件系统已使用容量大小，单位为 Byte
	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// 可用区 ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// 文件系统是否加密
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

type CreateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 权限组名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权限组描述信息
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// 已经与该权限组绑定的文件系统个数
	BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`

	// 权限组创建时间
	CDate *string `json:"CDate,omitempty" name:"CDate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 读写权限, 值为 RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限，值为 all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 读写权限, 值为 RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限，值为 all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`
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
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 客户端 IP
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 读写权限
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type CreateCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统id
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitempty" name:"ResourceTags"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteAutoSnapshotPolicyRequestParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
}

type DeleteAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
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
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID。说明，进行删除文件系统操作前需要先调用 DeleteMountTarget 接口删除该文件系统的挂载点，否则会删除失败。
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
}

type DeleteCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 用户 ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}

type DeleteCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
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
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 需要删除的文件文件系统快照ID 列表，快照ID，跟ID列表至少填一项
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
}

type DeleteCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统快照id
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 需要删除的文件文件系统快照ID 列表，快照ID，跟ID列表至少填一项
	SnapshotIds []*string `json:"SnapshotIds,omitempty" name:"SnapshotIds"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteMountTargetRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 挂载点 ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
}

type DeleteMountTargetRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 挂载点 ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DeleteUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 分页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页面长
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 升序，降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

type DescribeAutoSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 分页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页面长
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 升序，降序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 快照策略信息
	AutoSnapshotPolicies []*AutoSnapshotPolicyInfo `json:"AutoSnapshotPolicies,omitempty" name:"AutoSnapshotPolicies"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RegionZones []*AvailableRegion `json:"RegionZones,omitempty" name:"RegionZones"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCfsFileSystemClientsRequestParams struct {
	// 文件系统 ID。
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeCfsFileSystemClientsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID。
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemClientsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemClientsResponseParams struct {
	// 客户端列表
	ClientList []*FileSystemClient `json:"ClientList,omitempty" name:"ClientList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 私有网络（VPC） ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeCfsFileSystemsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 私有网络（VPC） ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网 ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfsFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfsFileSystemsResponseParams struct {
	// 文件系统信息
	FileSystems []*FileSystemInfo `json:"FileSystems,omitempty" name:"FileSystems"`

	// 文件系统总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupList []*PGroupInfo `json:"PGroupList,omitempty" name:"PGroupList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
}

type DescribeCfsRulesRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`
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
	RuleList []*PGroupRuleInfo `json:"RuleList,omitempty" name:"RuleList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StatisticsList []*SnapshotStatistics `json:"StatisticsList,omitempty" name:"StatisticsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页面长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序取值
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序 升序或者降序
	Order *string `json:"Order,omitempty" name:"Order"`
}

type DescribeCfsSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 分页起始位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 页面长度
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序取值
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序 升序或者降序
	Order *string `json:"Order,omitempty" name:"Order"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 快照信息描述
	Snapshots []*SnapshotInfo `json:"Snapshots,omitempty" name:"Snapshots"`

	// 快照列表快照汇总
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeMountTargetsRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	MountTargets []*MountInfo `json:"MountTargets,omitempty" name:"MountTargets"`

	// 挂载点数量
	NumberOfMountTargets *int64 `json:"NumberOfMountTargets,omitempty" name:"NumberOfMountTargets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeSnapshotOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 操作日志
	SnapshotOperates []*SnapshotOperateLog `json:"SnapshotOperates,omitempty" name:"SnapshotOperates"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 过滤条件。
	// <br><li>UserType - Array of String - 是否必填：否 -（过滤条件）按配额类型过滤。(Uid| Gid )
	// <br><li>UserId - Array of String - 是否必填：否 -（过滤条件）按UID/GID过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit 页面大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 过滤条件。
	// <br><li>UserType - Array of String - 是否必填：否 -（过滤条件）按配额类型过滤。(Uid| Gid )
	// <br><li>UserId - Array of String - 是否必填：否 -（过滤条件）按UID/GID过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset 分页码
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit 页面大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// UserQuota条目
	UserQuotaInfo []*UserQuota `json:"UserQuotaInfo,omitempty" name:"UserQuotaInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 文件系统大小
	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// 存储类型
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 快照总大小
	TotalSnapshotSize *uint64 `json:"TotalSnapshotSize,omitempty" name:"TotalSnapshotSize"`

	// 文件系统创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 文件系统所在区ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type FileSystemClient struct {
	// 文件系统IP地址
	CfsVip *string `json:"CfsVip,omitempty" name:"CfsVip"`

	// 客户端IP地址
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 文件系统所属VPCID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 可用区名称，例如ap-beijing-1，请参考 概览文档中的地域与可用区列表
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 该文件系统被挂载到客户端上的路径信息
	MountDirectory *string `json:"MountDirectory,omitempty" name:"MountDirectory"`
}

type FileSystemInfo struct {
	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 用户自定义名称
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 文件系统状态
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// 文件系统已使用容量
	SizeByte *uint64 `json:"SizeByte,omitempty" name:"SizeByte"`

	// 文件系统最大空间限制
	SizeLimit *uint64 `json:"SizeLimit,omitempty" name:"SizeLimit"`

	// 区域 ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 文件系统协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 文件系统存储类型
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 文件系统绑定的预付费存储包
	StorageResourcePkg *string `json:"StorageResourcePkg,omitempty" name:"StorageResourcePkg"`

	// 文件系统绑定的预付费带宽包（暂未支持）
	BandwidthResourcePkg *string `json:"BandwidthResourcePkg,omitempty" name:"BandwidthResourcePkg"`

	// 文件系统绑定权限组信息
	PGroup *PGroup `json:"PGroup,omitempty" name:"PGroup"`

	// 用户自定义名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// 文件系统是否加密
	Encrypted *bool `json:"Encrypted,omitempty" name:"Encrypted"`

	// 加密所使用的密钥，可以为密钥的 ID 或者 ARN
	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`

	// 应用ID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 文件系统吞吐上限，吞吐上限是根据文件系统当前已使用存储量、绑定的存储资源包以及吞吐资源包一同确定
	BandwidthLimit *float64 `json:"BandwidthLimit,omitempty" name:"BandwidthLimit"`

	// 文件系统总容量
	Capacity *uint64 `json:"Capacity,omitempty" name:"Capacity"`

	// 文件系统标签列表
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
}

type Filter struct {
	// 值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type MountInfo struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 挂载点 ID
	MountTargetId *string `json:"MountTargetId,omitempty" name:"MountTargetId"`

	// 挂载点 IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 挂载根目录
	FSID *string `json:"FSID,omitempty" name:"FSID"`

	// 挂载点状态
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// 网络类型
	NetworkInterface *string `json:"NetworkInterface,omitempty" name:"NetworkInterface"`

	// 私有网络 ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 子网 Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// CFS Turbo使用的云联网ID
	CcnID *string `json:"CcnID,omitempty" name:"CcnID"`

	// 云联网中CFS Turbo使用的网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`
}

type PGroup struct {
	// 权限组ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 权限组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type PGroupInfo struct {
	// 权限组ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 权限组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述信息
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// 创建时间
	CDate *string `json:"CDate,omitempty" name:"CDate"`

	// 关联文件系统个数
	BindCfsNum *int64 `json:"BindCfsNum,omitempty" name:"BindCfsNum"`
}

type PGroupRuleInfo struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 允许访问的客户端IP
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 读写权限, ro为只读，rw为读写
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// 规则优先级，1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

// Predefined struct for user
type SetUserQuotaRequestParams struct {
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个
	FileHardLimit *uint64 `json:"FileHardLimit,omitempty" name:"FileHardLimit"`
}

type SetUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个
	FileHardLimit *uint64 `json:"FileHardLimit,omitempty" name:"FileHardLimit"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CfsServiceStatus *string `json:"CfsServiceStatus,omitempty" name:"CfsServiceStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照ID
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 快照状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 快照大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 保留时长天
	AliveDay *uint64 `json:"AliveDay,omitempty" name:"AliveDay"`

	// 快照进度
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// 帐号ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 快照删除时间
	DeleteTime *string `json:"DeleteTime,omitempty" name:"DeleteTime"`

	// 文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// 快照标签
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
}

type SnapshotOperateLog struct {
	// 操作类型
	Action *string `json:"Action,omitempty" name:"Action"`

	// 操作时间
	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`

	// 操作名称
	ActionName *string `json:"ActionName,omitempty" name:"ActionName"`

	// 操作者
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 结果
	Result *uint64 `json:"Result,omitempty" name:"Result"`
}

type SnapshotStatistics struct {
	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 快照总个数
	SnapshotNumber *uint64 `json:"SnapshotNumber,omitempty" name:"SnapshotNumber"`

	// 快照总容量
	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyRequestParams struct {
	// 需要解绑的文件系统ID列表，用"," 分割
	FileSystemIds *string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`

	// 解绑的快照ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 需要解绑的文件系统ID列表，用"," 分割
	FileSystemIds *string `json:"FileSystemIds,omitempty" name:"FileSystemIds"`

	// 解绑的快照ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`
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
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 快照定期备份在一星期哪一天
	DayOfWeek *string `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitempty" name:"Hour"`

	// 快照保留日期
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`

	// 是否激活定期快照功能
	IsActivated *uint64 `json:"IsActivated,omitempty" name:"IsActivated"`
}

type UpdateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 快照定期备份在一星期哪一天
	DayOfWeek *string `json:"DayOfWeek,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitempty" name:"Hour"`

	// 快照保留日期
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`

	// 是否激活定期快照功能
	IsActivated *uint64 `json:"IsActivated,omitempty" name:"IsActivated"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAutoSnapshotPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAutoSnapshotPolicyResponseParams struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" name:"AutoSnapshotPolicyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`
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
	CreationToken *string `json:"CreationToken,omitempty" name:"CreationToken"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitempty" name:"FsName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	FsLimit *uint64 `json:"FsLimit,omitempty" name:"FsLimit"`

	// 文件系统ID，目前仅支持标准型文件系统。
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemSizeLimitRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统容量限制大小，输入范围0-1073741824, 单位为GB；其中输入值为0时，表示不限制文件系统容量。
	FsLimit *uint64 `json:"FsLimit,omitempty" name:"FsLimit"`

	// 文件系统ID，目前仅支持标准型文件系统。
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
}

type UpdateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 权限组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述信息
	DescInfo *string `json:"DescInfo,omitempty" name:"DescInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 读写权限, 值为RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限，值为all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
}

type UpdateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 读写权限, 值为RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限，值为all_squash、no_all_squash、root_squash、no_root_squash。其中all_squash为所有访问用户都会被映射为匿名用户或用户组；no_all_squash为访问用户会先与本机用户匹配，匹配失败后再映射为匿名用户或用户组；root_squash为将来访的root用户映射为匿名用户或用户组；no_root_squash为来访的root用户保持root帐号权限。不填默认为root_squash。
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`
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
	PGroupId *string `json:"PGroupId,omitempty" name:"PGroupId"`

	// 规则 ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 允许访问的客户端 IP 或者 IP 段
	AuthClientIp *string `json:"AuthClientIp,omitempty" name:"AuthClientIp"`

	// 读写权限
	RWPermission *string `json:"RWPermission,omitempty" name:"RWPermission"`

	// 用户权限
	UserPermission *string `json:"UserPermission,omitempty" name:"UserPermission"`

	// 优先级
	Priority *int64 `json:"Priority,omitempty" name:"Priority"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 文件系统快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 文件系统快照保留天数
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`
}

type UpdateCfsSnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统快照ID
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 文件系统快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 文件系统快照保留天数
	AliveDays *uint64 `json:"AliveDays,omitempty" name:"AliveDays"`
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
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type UserQuota struct {
	// 指定配额类型，包括Uid、Gid
	UserType *string `json:"UserType,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个
	FileHardLimit *uint64 `json:"FileHardLimit,omitempty" name:"FileHardLimit"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 容量使用，单位GiB
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapacityUsed *uint64 `json:"CapacityUsed,omitempty" name:"CapacityUsed"`

	// 文件使用个数，单位个
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUsed *uint64 `json:"FileUsed,omitempty" name:"FileUsed"`
}