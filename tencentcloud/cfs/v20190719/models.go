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

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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