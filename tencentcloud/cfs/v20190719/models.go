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

package v20190719

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ApplyPathLifecyclePolicyRequestParams struct {
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`

	// 生命周期管理策略关联目录的绝对路径列表
	Paths []*PathInfo `json:"Paths,omitnil,omitempty" name:"Paths"`
}

type ApplyPathLifecyclePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`

	// 生命周期管理策略关联目录的绝对路径列表
	Paths []*PathInfo `json:"Paths,omitnil,omitempty" name:"Paths"`
}

func (r *ApplyPathLifecyclePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyPathLifecyclePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecyclePolicyID")
	delete(f, "Paths")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyPathLifecyclePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyPathLifecyclePolicyResponseParams struct {
	// 有规则冲突时返回的已有冲突规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResults []*CheckResult `json:"CheckResults,omitnil,omitempty" name:"CheckResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyPathLifecyclePolicyResponse struct {
	*tchttp.BaseResponse
	Response *ApplyPathLifecyclePolicyResponseParams `json:"Response"`
}

func (r *ApplyPathLifecyclePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyPathLifecyclePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoScaleUpRule struct {
	// 自动扩容策略开启，关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群用量占比，到达这个值后开始扩容,范围[10-90]
	ScaleThreshold *uint64 `json:"ScaleThreshold,omitnil,omitempty" name:"ScaleThreshold"`

	// 扩容后使用量跟集群总量比例,范围[10-90]
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`
}

type AutoSnapshotPolicyInfo struct {
	// 快照策略ID
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称
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

	// 快照策略状态，available代表快照策略状态正常。这里只有一种状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号ID
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

	// 快照策略标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	// 需要解绑的文件系统ID列表，用"," 分割，文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 解绑的快照策略ID，可以通过[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208) 查询获取
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`
}

type BindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 需要解绑的文件系统ID列表，用"," 分割，文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 解绑的快照策略ID，可以通过[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208) 查询获取
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

type CheckResult struct {
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 目录绝对路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 生命周期管理策略关联的管理规则列表
	LifecycleRules []*LifecycleRule `json:"LifecycleRules,omitnil,omitempty" name:"LifecycleRules"`

	// 目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`
}

// Predefined struct for user
type CreateAccessCertRequestParams struct {
	// 证书描述，不超过64字符
	CertDesc *string `json:"CertDesc,omitnil,omitempty" name:"CertDesc"`
}

type CreateAccessCertRequest struct {
	*tchttp.BaseRequest
	
	// 证书描述，不超过64字符
	CertDesc *string `json:"CertDesc,omitnil,omitempty" name:"CertDesc"`
}

func (r *CreateAccessCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessCertResponseParams struct {
	// 凭证唯一标识
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessCertResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessCertResponseParams `json:"Response"`
}

func (r *CreateAccessCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoSnapshotPolicyRequestParams struct {
	// 快照重复时间点,0-23，小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 策略名称,限制64个字符数量仅支持输入中文、字母、数字、_或-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照重复日期，星期一到星期日。 1代表星期一、7代表星期天，与DayOfMonth，IntervalDays 三者选一
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照保留时长，单位天，默认永久0
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 快照按月重复，每月1-31号，选择一天，每月将在这一天自动创建快照；例如1 代表1号；与DayOfWeek，IntervalDays 三者选一
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数，与DayOfWeek，DayOfMonth 三者选一
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`

	// 快照策略标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type CreateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照重复时间点,0-23，小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 策略名称,限制64个字符数量仅支持输入中文、字母、数字、_或-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照重复日期，星期一到星期日。 1代表星期一、7代表星期天，与DayOfMonth，IntervalDays 三者选一
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照保留时长，单位天，默认永久0
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 快照按月重复，每月1-31号，选择一天，每月将在这一天自动创建快照；例如1 代表1号；与DayOfWeek，IntervalDays 三者选一
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数，与DayOfWeek，DayOfMonth 三者选一
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`

	// 快照策略标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
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
	delete(f, "ResourceTags")
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

	// 权限组 ID,pgroupbasic 是默认权限组，通过控制查询权限组列表接口获取[DescribeCfsPGroups](https://cloud.tencent.com/document/product/582/38157)
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统协议类型， 值为 NFS、CIFS、TURBO ; 若留空则默认为 NFS协议，turbo系列必须选择TURBO，不支持NFS、CIFS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值为 SD ；其中 SD 为通用标准型存储， HP为通用性能型存储， TB为Turbo标准型， TP 为Turbo性能型。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 私有网络（VPC） ID，若网络类型选择的是VPC，该字段为必填.通过查询私有网络接口获取，
	// [DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID，若网络类型选择的是VPC，该字段为必填。通过查询子网接口获取，
	// [DescribeSubnets](https://cloud.tencent.com/document/product/215/15784)
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP，Turbo系列当前不支持指定
	MountIP *string `json:"MountIP,omitnil,omitempty" name:"MountIP"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。用于保证请求幂等性的字符串失效时间为2小时。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云联网ID， 若网络类型选择的是CCN，该字段为必填;通过查询云联网列表接口获取，通过接口
	// [DescribeCcns](https://cloud.tencent.com/document/product/215/19199)
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 云联网中CFS使用的网段， 若网络类型选择的是Ccn，该字段为必填，且不能和Ccn中已经绑定的网段冲突
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 文件系统容量，turbo系列必填，单位为GiB。 turbo标准型单位GB，起售20TiB，即20480 GiB；扩容步长10TiB，即10240 GiB。turbo性能型起售10TiB，即10240 GiB；扩容步长10TiB，10240 GiB。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 文件系统快照ID，通过查询快照列表获取该参数，
	// [DescribeCfsSnapshots](https://cloud.tencent.com/document/product/582/80206)
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 定期快照策略ID，通过查询快照策略信息获取,
	// [DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/product/582/38157)
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 是否开启默认扩容，仅turbo类型文件存储支持
	EnableAutoScaleUp *bool `json:"EnableAutoScaleUp,omitnil,omitempty" name:"EnableAutoScaleUp"`

	// v1.5：创建普通版的通用文件系统；
	// v3.1：创建增强版的通用文件系统
	// 说明：增强版的通用系统需要开通白名单才能使用，如有需要请提交工单与我们联系。
	CfsVersion *string `json:"CfsVersion,omitnil,omitempty" name:"CfsVersion"`

	// turbo文件系统元数据属性
	// basic：创建标准型的元数据
	// enhanced：创建增强型的元数据
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`
}

type CreateCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 可用区名称，例如ap-beijing-1，请参考 [概览](https://cloud.tencent.com/document/product/582/13225) 文档中的地域与可用区列表
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 网络类型，可选值为 VPC，CCN；其中 VPC 为私有网络， CCN 为云联网。通用标准型/性能型请选择VPC，Turbo标准型/性能型请选择CCN。
	NetInterface *string `json:"NetInterface,omitnil,omitempty" name:"NetInterface"`

	// 权限组 ID,pgroupbasic 是默认权限组，通过控制查询权限组列表接口获取[DescribeCfsPGroups](https://cloud.tencent.com/document/product/582/38157)
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统协议类型， 值为 NFS、CIFS、TURBO ; 若留空则默认为 NFS协议，turbo系列必须选择TURBO，不支持NFS、CIFS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值为 SD ；其中 SD 为通用标准型存储， HP为通用性能型存储， TB为Turbo标准型， TP 为Turbo性能型。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 私有网络（VPC） ID，若网络类型选择的是VPC，该字段为必填.通过查询私有网络接口获取，
	// [DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID，若网络类型选择的是VPC，该字段为必填。通过查询子网接口获取，
	// [DescribeSubnets](https://cloud.tencent.com/document/product/215/15784)
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP，Turbo系列当前不支持指定
	MountIP *string `json:"MountIP,omitnil,omitempty" name:"MountIP"`

	// 用户自定义文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 文件系统标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。用于保证请求幂等性的字符串失效时间为2小时。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 云联网ID， 若网络类型选择的是CCN，该字段为必填;通过查询云联网列表接口获取，通过接口
	// [DescribeCcns](https://cloud.tencent.com/document/product/215/19199)
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 云联网中CFS使用的网段， 若网络类型选择的是Ccn，该字段为必填，且不能和Ccn中已经绑定的网段冲突
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 文件系统容量，turbo系列必填，单位为GiB。 turbo标准型单位GB，起售20TiB，即20480 GiB；扩容步长10TiB，即10240 GiB。turbo性能型起售10TiB，即10240 GiB；扩容步长10TiB，10240 GiB。
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// 文件系统快照ID，通过查询快照列表获取该参数，
	// [DescribeCfsSnapshots](https://cloud.tencent.com/document/product/582/80206)
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 定期快照策略ID，通过查询快照策略信息获取,
	// [DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/product/582/38157)
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 是否开启默认扩容，仅turbo类型文件存储支持
	EnableAutoScaleUp *bool `json:"EnableAutoScaleUp,omitnil,omitempty" name:"EnableAutoScaleUp"`

	// v1.5：创建普通版的通用文件系统；
	// v3.1：创建增强版的通用文件系统
	// 说明：增强版的通用系统需要开通白名单才能使用，如有需要请提交工单与我们联系。
	CfsVersion *string `json:"CfsVersion,omitnil,omitempty" name:"CfsVersion"`

	// turbo文件系统元数据属性
	// basic：创建标准型的元数据
	// enhanced：创建增强型的元数据
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`
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
	delete(f, "CfsVersion")
	delete(f, "MetaType")
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 读写权限, 值为 RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为 all_squash、no_all_squash、root_squash、no_root_squash。默认值为root_squash
	// all_squash：所有访问用户（含 root 用户）都会被映射为匿名用户或用户组。
	// no_all_squash：所有访问用户（含 root 用户）均保持原有的 UID/GID 信息。
	// root_squash：将来访的 root 用户映射为匿名用户或用户组，非 root 用户保持原有的 UID/GID 信息。
	// no_root_squash：与 no_all_squash 效果一致，所有访问用户（含 root 用户）均保持原有的 UID/GID 信息
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`
}

type CreateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 读写权限, 值为 RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为 all_squash、no_all_squash、root_squash、no_root_squash。默认值为root_squash
	// all_squash：所有访问用户（含 root 用户）都会被映射为匿名用户或用户组。
	// no_all_squash：所有访问用户（含 root 用户）均保持原有的 UID/GID 信息。
	// root_squash：将来访的 root 用户映射为匿名用户或用户组，非 root 用户保持原有的 UID/GID 信息。
	// no_root_squash：与 no_all_squash 效果一致，所有访问用户（含 root 用户）均保持原有的 UID/GID 信息
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
	// 文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照名称，支持不超过64字符长度，支持中文、数字、_、-
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照标签
	ResourceTags []*TagInfo `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type CreateCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照名称，支持不超过64字符长度，支持中文、数字、_、-
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
type CreateDataFlowRequestParams struct {
	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 源端数据类型；包含S3_COS，S3_L5 
	SourceStorageType *string `json:"SourceStorageType,omitnil,omitempty" name:"SourceStorageType"`

	// 源端存储地址
	SourceStorageAddress *string `json:"SourceStorageAddress,omitnil,omitempty" name:"SourceStorageAddress"`

	// 源端路径
	SourcePath *string `json:"SourcePath,omitnil,omitempty" name:"SourcePath"`

	// 文件系统内目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 密钥 ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 密钥 key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 数据流动名称；支持不超过64字符长度，支持中文、数字、_、-
	DataFlowName *string `json:"DataFlowName,omitnil,omitempty" name:"DataFlowName"`

	//  0：不开启自动更新  1：开启自动更新
	AutoRefresh *uint64 `json:"AutoRefresh,omitnil,omitempty" name:"AutoRefresh"`

	// KafkaConsumer 消费时使用的Topic参数
	UserKafkaTopic *string `json:"UserKafkaTopic,omitnil,omitempty" name:"UserKafkaTopic"`

	// 	服务地址 示例值：kafkaconsumer-ap-beijing.cls.tencentyun.com:9095
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Kafka消费用户名.示例值：name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Kafka消费用户密码。默认${SecretId}#${SecretKey}。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 源端数据类型；包含S3_COS，S3_L5 
	SourceStorageType *string `json:"SourceStorageType,omitnil,omitempty" name:"SourceStorageType"`

	// 源端存储地址
	SourceStorageAddress *string `json:"SourceStorageAddress,omitnil,omitempty" name:"SourceStorageAddress"`

	// 源端路径
	SourcePath *string `json:"SourcePath,omitnil,omitempty" name:"SourcePath"`

	// 文件系统内目标路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// 密钥 ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 密钥 key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 数据流动名称；支持不超过64字符长度，支持中文、数字、_、-
	DataFlowName *string `json:"DataFlowName,omitnil,omitempty" name:"DataFlowName"`

	//  0：不开启自动更新  1：开启自动更新
	AutoRefresh *uint64 `json:"AutoRefresh,omitnil,omitempty" name:"AutoRefresh"`

	// KafkaConsumer 消费时使用的Topic参数
	UserKafkaTopic *string `json:"UserKafkaTopic,omitnil,omitempty" name:"UserKafkaTopic"`

	// 	服务地址 示例值：kafkaconsumer-ap-beijing.cls.tencentyun.com:9095
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Kafka消费用户名.示例值：name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Kafka消费用户密码。默认${SecretId}#${SecretKey}。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *CreateDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "SourceStorageType")
	delete(f, "SourceStorageAddress")
	delete(f, "SourcePath")
	delete(f, "TargetPath")
	delete(f, "SecretId")
	delete(f, "SecretKey")
	delete(f, "DataFlowName")
	delete(f, "AutoRefresh")
	delete(f, "UserKafkaTopic")
	delete(f, "ServerAddr")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataFlowResponseParams struct {
	// 数据流动管理 ID
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataFlowResponseParams `json:"Response"`
}

func (r *CreateDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecycleDataTaskRequestParams struct {
	// 文件系统唯一 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 生命周期任务类型；archive：沉降；restore：预热；release：数据释放；metaload：元数据加载
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 需要沉降的路径或文件，仅支持传入1个路径，不允许为空。
	TaskPath *string `json:"TaskPath,omitnil,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 数据流动 ID ，该接口可以通过 DescribeDataFlow 查询
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`
}

type CreateLifecycleDataTaskRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统唯一 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 生命周期任务类型；archive：沉降；restore：预热；release：数据释放；metaload：元数据加载
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 需要沉降的路径或文件，仅支持传入1个路径，不允许为空。
	TaskPath *string `json:"TaskPath,omitnil,omitempty" name:"TaskPath"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 数据流动 ID ，该接口可以通过 DescribeDataFlow 查询
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`
}

func (r *CreateLifecycleDataTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecycleDataTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Type")
	delete(f, "TaskPath")
	delete(f, "TaskName")
	delete(f, "DataFlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifecycleDataTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecycleDataTaskResponseParams struct {
	// 任务 ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLifecycleDataTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifecycleDataTaskResponseParams `json:"Response"`
}

func (r *CreateLifecycleDataTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecycleDataTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecyclePolicyDownloadTaskRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 下载文件的类型，包含 FileSuccessList，FileTotalList，FileFailedList
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CreateLifecyclePolicyDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 下载文件的类型，包含 FileSuccessList，FileTotalList，FileFailedList
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *CreateLifecyclePolicyDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecyclePolicyDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifecyclePolicyDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecyclePolicyDownloadTaskResponseParams struct {
	// 下载路径
	DownloadAddress *string `json:"DownloadAddress,omitnil,omitempty" name:"DownloadAddress"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLifecyclePolicyDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifecyclePolicyDownloadTaskResponseParams `json:"Response"`
}

func (r *CreateLifecyclePolicyDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecyclePolicyDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecyclePolicyRequestParams struct {
	// 生命周期管理策略名称，中文/英文/数字/下划线/中划线的组合，不超过64个字符
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 生命周期管理策略关联的管理规则列表
	LifecycleRules []*LifecycleRule `json:"LifecycleRules,omitnil,omitempty" name:"LifecycleRules"`
}

type CreateLifecyclePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期管理策略名称，中文/英文/数字/下划线/中划线的组合，不超过64个字符
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 生命周期管理策略关联的管理规则列表
	LifecycleRules []*LifecycleRule `json:"LifecycleRules,omitnil,omitempty" name:"LifecycleRules"`
}

func (r *CreateLifecyclePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecyclePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecyclePolicyName")
	delete(f, "LifecycleRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifecyclePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecyclePolicyResponseParams struct {
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLifecyclePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifecyclePolicyResponseParams `json:"Response"`
}

func (r *CreateLifecyclePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecyclePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationTaskRequestParams struct {
	// 迁移任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 迁移方式标志位，默认为0。0：桶迁移；1：清单迁移
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// 迁移模式，默认为0。0: 全量迁移
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// 数据源账号的 SecretId
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// 数据源账号的 SecretKey
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`

	// 文件系统实例 ID，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// 同名文件迁移时覆盖策略，默认为0。0: 最后修改时间优先；1: 全覆盖；2: 不覆盖
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 数据源服务商。COS：腾讯云COS，OSS：阿里云OSS，OBS：华为云OBS
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// 数据源桶名称；桶迁移时，BucketName 和 BucketAddress 必填其一，清单迁移时无需填写此参数
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 数据源桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 数据源桶地址；桶迁移时，BucketName 和 BucketAddress 必填其一，清单迁移时无需填写此参数
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 清单地址，迁移方式为清单迁移时必填
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// 目标文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 源桶路径，默认为 /
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`

	// 迁移方向；0：对象存储迁移至文件系统，1：文件系统迁移至对象存储。默认为0
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type CreateMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 迁移方式标志位，默认为0。0：桶迁移；1：清单迁移
	MigrationType *uint64 `json:"MigrationType,omitnil,omitempty" name:"MigrationType"`

	// 迁移模式，默认为0。0: 全量迁移
	MigrationMode *uint64 `json:"MigrationMode,omitnil,omitempty" name:"MigrationMode"`

	// 数据源账号的 SecretId
	SrcSecretId *string `json:"SrcSecretId,omitnil,omitempty" name:"SrcSecretId"`

	// 数据源账号的 SecretKey
	SrcSecretKey *string `json:"SrcSecretKey,omitnil,omitempty" name:"SrcSecretKey"`

	// 文件系统实例 ID，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`

	// 同名文件迁移时覆盖策略，默认为0。0: 最后修改时间优先；1: 全覆盖；2: 不覆盖
	CoverType *uint64 `json:"CoverType,omitnil,omitempty" name:"CoverType"`

	// 数据源服务商。COS：腾讯云COS，OSS：阿里云OSS，OBS：华为云OBS
	SrcService *string `json:"SrcService,omitnil,omitempty" name:"SrcService"`

	// 数据源桶名称；桶迁移时，BucketName 和 BucketAddress 必填其一，清单迁移时无需填写此参数
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 数据源桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 数据源桶地址；桶迁移时，BucketName 和 BucketAddress 必填其一，清单迁移时无需填写此参数
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 清单地址，迁移方式为清单迁移时必填
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// 目标文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 源桶路径，默认为 /
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`

	// 迁移方向；0：对象存储迁移至文件系统，1：文件系统迁移至对象存储。默认为0
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
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
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationTaskResponseParams struct {
	// 迁移任务 ID
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

type DataFlowInfo struct {
	// 数据流动管理 ID
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 数据流动名称
	DataFlowName *string `json:"DataFlowName,omitnil,omitempty" name:"DataFlowName"`

	// 源端数据类型
	SourceStorageType *string `json:"SourceStorageType,omitnil,omitempty" name:"SourceStorageType"`

	// 源端存储地址
	SourceStorageAddress *string `json:"SourceStorageAddress,omitnil,omitempty" name:"SourceStorageAddress"`

	// 源端路径
	SourcePath *string `json:"SourcePath,omitnil,omitempty" name:"SourcePath"`

	// 目录路径
	TargetPath *string `json:"TargetPath,omitnil,omitempty" name:"TargetPath"`

	// available：已生效
	// pending：配置中
	// unavailable：失效
	// deleting：删除中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 0：不开启自动更新
	// 
	// 1：开启自动更新
	AutoRefresh *uint64 `json:"AutoRefresh,omitnil,omitempty" name:"AutoRefresh"`

	// KafkaConsumer 消费时使用的Topic参数
	UserKafkaTopic *string `json:"UserKafkaTopic,omitnil,omitempty" name:"UserKafkaTopic"`

	// 服务地址
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// Kafka消费用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 自动刷新的状态，available：已生效
	// pending：配置中
	// unavailable：失效
	AutoRefreshStatus *string `json:"AutoRefreshStatus,omitnil,omitempty" name:"AutoRefreshStatus"`

	// 自动刷新开启时间
	AutoRefreshTime *string `json:"AutoRefreshTime,omitnil,omitempty" name:"AutoRefreshTime"`
}

// Predefined struct for user
type DeleteAutoSnapshotPolicyRequestParams struct {
	// 快照策略ID，查询快照策略接口获取,[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208)
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type DeleteAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略ID，查询快照策略接口获取,[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208)
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
	// 文件系统 ID，通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取。说明，进行删除文件系统操作前需要先调用 DeleteMountTarget 接口删除该文件系统的挂载点，否则会删除失败。
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DeleteCfsFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID，通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取。说明，进行删除文件系统操作前需要先调用 DeleteMountTarget 接口删除该文件系统的挂载点，否则会删除失败。
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

type DeleteCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID，可通过[DescribeCfsRules](https://cloud.tencent.com/document/api/582/38156)接口获取
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID，可通过[DescribeCfsRules](https://cloud.tencent.com/document/api/582/38156)接口获取
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
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 需要删除的文件系统快照ID 列表，快照ID，跟ID列表至少填一项
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotIds []*string `json:"SnapshotIds,omitnil,omitempty" name:"SnapshotIds"`
}

type DeleteCfsSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 需要删除的文件系统快照ID 列表，快照ID，跟ID列表至少填一项
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
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
type DeleteDataFlowRequestParams struct {
	// 数据流动管理 ID
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DeleteDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// 数据流动管理 ID
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DeleteDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataFlowId")
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataFlowResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataFlowResponseParams `json:"Response"`
}

func (r *DeleteDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifecyclePolicyRequestParams struct {
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`
}

type DeleteLifecyclePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`
}

func (r *DeleteLifecyclePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifecyclePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecyclePolicyID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLifecyclePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifecyclePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLifecyclePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLifecyclePolicyResponseParams `json:"Response"`
}

func (r *DeleteLifecyclePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifecyclePolicyResponse) FromJsonString(s string) error {
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
type DeleteUserQuotaRequestParams struct {
	// 文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid（按用户ID限制）、Gid（按用户组ID限制）、Dir（按目录限制）
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息，和DirectoryPath参数，两者必须填写一个
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设置目录配额的目录的绝对路径，和UserId参数，两者必须填写一个
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`
}

type DeleteUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid（按用户ID限制）、Gid（按用户组ID限制）、Dir（按目录限制）
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息，和DirectoryPath参数，两者必须填写一个
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设置目录配额的目录的绝对路径，和UserId参数，两者必须填写一个
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`
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
	delete(f, "DirectoryPath")
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
	// 文件系统 ID，通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Offset 分页码，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，默认为10，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCfsFileSystemClientsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID，通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// Offset 分页码，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，默认为10，最大值为100
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

	// Offset 分页码,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，默认10
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

	// Offset 分页码,默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，默认10
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`
}

type DescribeCfsRulesRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
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
	// 文件系统 ID，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照 ID
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
	// <br>Status - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。状态分类：creating：创建中 | available：运行中 | deleting：删除中 | rollbacking_new：由快照创建新文件系统中 | create-failed：创建失败。
	// <br>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。
	// <br>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key 使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 按创建时间排序取值
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序；升序或者降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeCfsSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照 ID
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
	// <br>Status - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。状态分类：creating：创建中 | available：运行中 | deleting：删除中 | rollbacking_new：由快照创建新文件系统中 | create-failed：创建失败。
	// <br>tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。
	// <br>tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key 使用具体的标签键进行替换。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 按创建时间排序取值
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序；升序或者降序
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
type DescribeDataFlowRequestParams struct {
	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 数据流动 ID ，由创建数据流动返回
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 每次查询返回值个数，默认20；最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 文件系统版本；版本号：v1.5，v3.0，v3.1，v4.0
	CfsVersion *string `json:"CfsVersion,omitnil,omitempty" name:"CfsVersion"`
}

type DescribeDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 数据流动 ID ，由创建数据流动返回
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 每次查询返回值个数，默认20；最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 文件系统版本；版本号：v1.5，v3.0，v3.1，v4.0
	CfsVersion *string `json:"CfsVersion,omitnil,omitempty" name:"CfsVersion"`
}

func (r *DescribeDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "DataFlowId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CfsVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataFlowResponseParams struct {
	// 查询总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 无
	DataFlows []*DataFlowInfo `json:"DataFlows,omitnil,omitempty" name:"DataFlows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataFlowResponseParams `json:"Response"`
}

func (r *DescribeDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecycleDataTaskRequestParams struct {
	// 开始时间。须早于 EndTime ，仅支持查询最近3个月内的任务数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。须晚于 StartTime ，仅支持查询最近3个月内的任务数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Offset 分页码	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，TaskName，FileSystemId，Type
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 文件系统版本；v3.1: pcfs/hifs v4.0:Turbo
	CfsVersion *string `json:"CfsVersion,omitnil,omitempty" name:"CfsVersion"`
}

type DescribeLifecycleDataTaskRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。须早于 EndTime ，仅支持查询最近3个月内的任务数据。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。须晚于 StartTime ，仅支持查询最近3个月内的任务数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Offset 分页码	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件，TaskName，FileSystemId，Type
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 文件系统版本；v3.1: pcfs/hifs v4.0:Turbo
	CfsVersion *string `json:"CfsVersion,omitnil,omitempty" name:"CfsVersion"`
}

func (r *DescribeLifecycleDataTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecycleDataTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "CfsVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLifecycleDataTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecycleDataTaskResponseParams struct {
	// 任务数组
	LifecycleDataTask []*LifecycleDataTaskInfo `json:"LifecycleDataTask,omitnil,omitempty" name:"LifecycleDataTask"`

	// 查询结果总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLifecycleDataTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLifecycleDataTaskResponseParams `json:"Response"`
}

func (r *DescribeLifecycleDataTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecycleDataTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecyclePoliciesRequestParams struct {
	// 生命周期管理策略名称
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 每个分页包含的生命周期管理策略个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 列表的分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`
}

type DescribeLifecyclePoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期管理策略名称
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 每个分页包含的生命周期管理策略个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 列表的分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`
}

func (r *DescribeLifecyclePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecyclePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecyclePolicyName")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "FileSystemId")
	delete(f, "LifecyclePolicyID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLifecyclePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecyclePoliciesResponseParams struct {
	// 列表的分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每个分页包含的生命周期管理策略个数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 生命周期管理策略总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 生命周期管理策略列表
	LifecyclePolicies []*LifecyclePolicy `json:"LifecyclePolicies,omitnil,omitempty" name:"LifecyclePolicies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLifecyclePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLifecyclePoliciesResponseParams `json:"Response"`
}

func (r *DescribeLifecyclePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecyclePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTasksRequestParams struct {
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <br><li> taskId按照【迁移任务id】进行过滤。类型：String必选：否<br></li><br><li>  taskName按照【迁移任务名字】进行模糊搜索过滤。类型：String必选：否每次请求的Filters的上限为10，Filter.Values的上限为100。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMigrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <br><li> taskId按照【迁移任务id】进行过滤。类型：String必选：否<br></li><br><li>  taskName按照【迁移任务名字】进行模糊搜索过滤。类型：String必选：否每次请求的Filters的上限为10，Filter.Values的上限为100。</li>
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
	// 迁移任务的总数量
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
	// 文件系统 ID，[查询文件系统列表](https://cloud.tencent.com/document/api/582/38170)可以获得id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeMountTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID，[查询文件系统列表](https://cloud.tencent.com/document/api/582/38170)可以获得id
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
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 起始时间，格式“YYYY-MM-DD hh:mm:ss”
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式“YYYY-MM-DD hh:mm:ss”
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeSnapshotOperationLogsRequest struct {
	*tchttp.BaseRequest
	
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 起始时间，格式“YYYY-MM-DD hh:mm:ss”
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，格式“YYYY-MM-DD hh:mm:ss”
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
	// 文件系统 ID,通过[查询文件系统列表](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 过滤条件。
	// UserType - Array of String - 是否必填：否 -（过滤条件）按配额类型过滤。(Uid|Gid|Dir，分别对应用户，用户组，目录 )
	// UserId- Array of String - 是否必填：否 -（过滤条件）按用户id过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset 分页码，默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，可填范围为大于0的整数，默认值是10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID,通过[查询文件系统列表](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 过滤条件。
	// UserType - Array of String - 是否必填：否 -（过滤条件）按配额类型过滤。(Uid|Gid|Dir，分别对应用户，用户组，目录 )
	// UserId- Array of String - 是否必填：否 -（过滤条件）按用户id过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Offset 分页码，默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// Limit 页面大小，可填范围为大于0的整数，默认值是10
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

// Predefined struct for user
type DoDirectoryOperationRequestParams struct {
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// create：创建目录，等同于mkdir。
	// check：确认目录是否存在，等同于stat。
	// move：对文件/目录进行重命名，等同于mv。
	OpetationType *string `json:"OpetationType,omitnil,omitempty" name:"OpetationType"`

	// 目录的绝对路径  默认递归创建（即如果目录中有子目录不存在，则先创建出对应子目录）
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`

	// 创建目录的权限，若不传，默认为0755  若Operation Type为check，此值无实际意义
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// mv操作的目标目录名称；如果是turbo文件系统必须以/cfs/开头
	DestPath *string `json:"DestPath,omitnil,omitempty" name:"DestPath"`
}

type DoDirectoryOperationRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统Id
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// create：创建目录，等同于mkdir。
	// check：确认目录是否存在，等同于stat。
	// move：对文件/目录进行重命名，等同于mv。
	OpetationType *string `json:"OpetationType,omitnil,omitempty" name:"OpetationType"`

	// 目录的绝对路径  默认递归创建（即如果目录中有子目录不存在，则先创建出对应子目录）
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`

	// 创建目录的权限，若不传，默认为0755  若Operation Type为check，此值无实际意义
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// mv操作的目标目录名称；如果是turbo文件系统必须以/cfs/开头
	DestPath *string `json:"DestPath,omitnil,omitempty" name:"DestPath"`
}

func (r *DoDirectoryOperationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DoDirectoryOperationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "OpetationType")
	delete(f, "DirectoryPath")
	delete(f, "Mode")
	delete(f, "DestPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DoDirectoryOperationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DoDirectoryOperationResponseParams struct {
	// 1:成功  0:失败  创建目录的操作，1表示创建成功，0表示创建失败。  确认目录是否存在的操作，1表示目录存在，0表示目录不存在。  说明：创建目录操作若目录已存在，也会返回创建成功。
	Result *int64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DoDirectoryOperationResponse struct {
	*tchttp.BaseResponse
	Response *DoDirectoryOperationResponseParams `json:"Response"`
}

func (r *DoDirectoryOperationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DoDirectoryOperationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExstraPerformanceInfo struct {
	// fixed: 最终值固定
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 额外购买的CFS性能值，单位MB/s。
	Performance *int64 `json:"Performance,omitnil,omitempty" name:"Performance"`
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

	// 可用区名称，例如ap-beijing-1，参考[简介](https://cloud.tencent.com/document/api/582/38144)文档中的地域与可用区列表
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

	// 文件系统关联的快照策略
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 文件系统处理快照状态,snapping：快照中，normal：正常状态
	SnapStatus *string `json:"SnapStatus,omitnil,omitempty" name:"SnapStatus"`

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
	TieringDetail *TieringDetailInfo `json:"TieringDetail,omitnil,omitempty" name:"TieringDetail"`

	// 文件系统自动扩容策略
	AutoScaleUpRule *AutoScaleUpRule `json:"AutoScaleUpRule,omitnil,omitempty" name:"AutoScaleUpRule"`

	// 文件系统版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 额外性能信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExstraPerformanceInfo []*ExstraPerformanceInfo `json:"ExstraPerformanceInfo,omitnil,omitempty" name:"ExstraPerformanceInfo"`

	// basic：标准版元数据类型
	// enhanced：增项版元数据类型
	MetaType *string `json:"MetaType,omitnil,omitempty" name:"MetaType"`
}

type Filter struct {
	// 值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type LifecycleDataTaskInfo struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务状态.
	// init：未执行
	// running：执行中，finished：已完成
	// ,failed：失败
	// ,stopping：停止中,stopped：已停止
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务创建时间
	CreationTime *string `json:"CreationTime,omitnil,omitempty" name:"CreationTime"`

	// 任务结束时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 文件总数
	FileTotalCount *uint64 `json:"FileTotalCount,omitnil,omitempty" name:"FileTotalCount"`

	// 处理成功文件数量
	FileSuccessedCount *uint64 `json:"FileSuccessedCount,omitnil,omitempty" name:"FileSuccessedCount"`

	// 当前已经失败的文件数
	FileFailedCount *uint64 `json:"FileFailedCount,omitnil,omitempty" name:"FileFailedCount"`

	// 文件容量，单位Byte
	// 
	FileTotalSize *uint64 `json:"FileTotalSize,omitnil,omitempty" name:"FileTotalSize"`

	// 已处理完成的文件容量，单位Byte
	// 
	FileSuccessedSize *uint64 `json:"FileSuccessedSize,omitnil,omitempty" name:"FileSuccessedSize"`

	// 已处理失败文件容量，单位Byte
	FileFailedSize *uint64 `json:"FileFailedSize,omitnil,omitempty" name:"FileFailedSize"`

	// 总文件列表
	FileTotalList *string `json:"FileTotalList,omitnil,omitempty" name:"FileTotalList"`

	// 成功的文件列表
	FileSuccessedList *string `json:"FileSuccessedList,omitnil,omitempty" name:"FileSuccessedList"`

	// 失败文件的列表
	FileFailedList *string `json:"FileFailedList,omitnil,omitempty" name:"FileFailedList"`

	// FileSystemId
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务路径
	TaskPath *string `json:"TaskPath,omitnil,omitempty" name:"TaskPath"`

	// 任务类型,archive:表示沉降任务，restore：表示拉取任务
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据流动Id
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`
}

type LifecyclePolicy struct {
	// 生命周期管理策略创建的时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`

	// 生命周期管理策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 生命周期管理策略关联的管理规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleRules []*LifecycleRule `json:"LifecycleRules,omitnil,omitempty" name:"LifecycleRules"`

	// 生命周期管理策略关联目录的绝对路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paths []*PathInfo `json:"Paths,omitnil,omitempty" name:"Paths"`
}

type LifecycleRule struct {
	// 数据转储后的存储类型。其中：InfrequentAccess：低频介质存储；ColdStorage：冷存储。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 数据转储文件类型。其中，BIG_FILE：超大文件；STD_FILE：普通文件；SMALL_FILE：小文件；ALL：所有文件。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 数据转储行为。其中，Archive：沉降；Noarchive：不沉降。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 数据转储触发时间。由“DEFAULT_ATIME_”与“数字”组成，单位为天。当 Action 为 Noarchive，请保持为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 数据转储文件最大规格。其数值需使用“数字+单位”格式进行表示，单位支持K（KiB）、M（MiB）、G（GiB）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMaxSize *string `json:"FileMaxSize,omitnil,omitempty" name:"FileMaxSize"`

	// 数据转储文件最小规格。其数值需使用“数字+单位”格式进行表示，单位支持K（KiB）、M（MiB）、G（GiB）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileMinSize *string `json:"FileMinSize,omitnil,omitempty" name:"FileMinSize"`
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
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 数据源桶地域
	BucketRegion *string `json:"BucketRegion,omitnil,omitempty" name:"BucketRegion"`

	// 数据源桶地址
	BucketAddress *string `json:"BucketAddress,omitnil,omitempty" name:"BucketAddress"`

	// 清单地址
	ListAddress *string `json:"ListAddress,omitnil,omitempty" name:"ListAddress"`

	// 文件系统实例名称
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
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 迁移状态。0: 已完成；1: 创建中；2: 运行中；3: 终止中；4: 已终止；5: 创建失败；6: 运行失败；7: 结束中；8: 删除中；9: 等待中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件数量
	FileTotalCount *uint64 `json:"FileTotalCount,omitnil,omitempty" name:"FileTotalCount"`

	// 已迁移文件数量
	FileMigratedCount *uint64 `json:"FileMigratedCount,omitnil,omitempty" name:"FileMigratedCount"`

	// 迁移失败文件数量
	FileFailedCount *uint64 `json:"FileFailedCount,omitnil,omitempty" name:"FileFailedCount"`

	// 文件容量，单位Byte
	FileTotalSize *int64 `json:"FileTotalSize,omitnil,omitempty" name:"FileTotalSize"`

	// 已迁移文件容量，单位Byte
	FileMigratedSize *int64 `json:"FileMigratedSize,omitnil,omitempty" name:"FileMigratedSize"`

	// 迁移失败文件容量，单位Byte
	FileFailedSize *int64 `json:"FileFailedSize,omitnil,omitempty" name:"FileFailedSize"`

	// 全部清单
	FileTotalList *string `json:"FileTotalList,omitnil,omitempty" name:"FileTotalList"`

	// 已完成文件清单
	FileCompletedList *string `json:"FileCompletedList,omitnil,omitempty" name:"FileCompletedList"`

	// 失败文件清单
	FileFailedList *string `json:"FileFailedList,omitnil,omitempty" name:"FileFailedList"`

	// 源桶路径
	BucketPath *string `json:"BucketPath,omitnil,omitempty" name:"BucketPath"`

	// 迁移方向。0: 对象存储迁移至文件系统，1: 文件系统迁移至对象存储。默认 0
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

// Predefined struct for user
type ModifyDataFlowRequestParams struct {
	// 数据流动管理 ID ，通过查询数据流动接口获取
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 数据流动名称；支持不超过64字符长度，支持中文、数字、_、-
	DataFlowName *string `json:"DataFlowName,omitnil,omitempty" name:"DataFlowName"`

	// 密钥 ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 密钥 key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// KafkaConsumer 消费时使用的Topic参数
	UserKafkaTopic *string `json:"UserKafkaTopic,omitnil,omitempty" name:"UserKafkaTopic"`

	// 服务地址
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Kafka消费用户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 元数据增量更新开关；1开启，0关闭
	AutoRefresh *uint64 `json:"AutoRefresh,omitnil,omitempty" name:"AutoRefresh"`
}

type ModifyDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// 数据流动管理 ID ，通过查询数据流动接口获取
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 文件系统 ID ，通过查询文件系统 [DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170) 获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 数据流动名称；支持不超过64字符长度，支持中文、数字、_、-
	DataFlowName *string `json:"DataFlowName,omitnil,omitempty" name:"DataFlowName"`

	// 密钥 ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 密钥 key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// KafkaConsumer 消费时使用的Topic参数
	UserKafkaTopic *string `json:"UserKafkaTopic,omitnil,omitempty" name:"UserKafkaTopic"`

	// 服务地址
	ServerAddr *string `json:"ServerAddr,omitnil,omitempty" name:"ServerAddr"`

	// name
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// Kafka消费用户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 元数据增量更新开关；1开启，0关闭
	AutoRefresh *uint64 `json:"AutoRefresh,omitnil,omitempty" name:"AutoRefresh"`
}

func (r *ModifyDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataFlowId")
	delete(f, "FileSystemId")
	delete(f, "DataFlowName")
	delete(f, "SecretId")
	delete(f, "SecretKey")
	delete(f, "UserKafkaTopic")
	delete(f, "ServerAddr")
	delete(f, "UserName")
	delete(f, "Password")
	delete(f, "AutoRefresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataFlowResponseParams struct {
	// 数据流动管理 ID
	DataFlowId *string `json:"DataFlowId,omitnil,omitempty" name:"DataFlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataFlowResponseParams `json:"Response"`
}

func (r *ModifyDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemAutoScaleUpRuleRequestParams struct {
	// 文件系统 ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容阈值，范围[10-90]
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// 扩容后目标阈值，范围[10-90]，该值要小于 ScaleUpThreshold
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// 规则状态 0：关闭，1：开启；不传保留原状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyFileSystemAutoScaleUpRuleRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容阈值，范围[10-90]
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// 扩容后目标阈值，范围[10-90]，该值要小于 ScaleUpThreshold
	TargetThreshold *uint64 `json:"TargetThreshold,omitnil,omitempty" name:"TargetThreshold"`

	// 规则状态 0：关闭，1：开启；不传保留原状态
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
	// 文件系统 ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 规则状态 0：关闭，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 扩容阈值，范围[10-90]
	ScaleUpThreshold *uint64 `json:"ScaleUpThreshold,omitnil,omitempty" name:"ScaleUpThreshold"`

	// 扩容后达到阈值，范围[10-90]
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

// Predefined struct for user
type ModifyLifecyclePolicyRequestParams struct {
	// 生命周期管理策略名称，中文/英文/数字/下划线/中划线的组合，不超过64个字符
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 生命周期管理策略关联的管理规则列表
	LifecycleRules []*LifecycleRule `json:"LifecycleRules,omitnil,omitempty" name:"LifecycleRules"`

	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`
}

type ModifyLifecyclePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期管理策略名称，中文/英文/数字/下划线/中划线的组合，不超过64个字符
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitnil,omitempty" name:"LifecyclePolicyName"`

	// 生命周期管理策略关联的管理规则列表
	LifecycleRules []*LifecycleRule `json:"LifecycleRules,omitnil,omitempty" name:"LifecycleRules"`

	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`
}

func (r *ModifyLifecyclePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifecyclePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecyclePolicyName")
	delete(f, "LifecycleRules")
	delete(f, "LifecyclePolicyID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLifecyclePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifecyclePolicyResponseParams struct {
	// 生命周期管理策略ID
	LifecyclePolicyID *string `json:"LifecyclePolicyID,omitnil,omitempty" name:"LifecyclePolicyID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLifecyclePolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLifecyclePolicyResponseParams `json:"Response"`
}

func (r *ModifyLifecyclePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifecyclePolicyResponse) FromJsonString(s string) error {
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

	// 挂载点状态，包括creating：创建中；available：运行中；
	// deleting：删除中；
	// create_failed： 创建失败
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// 网络类型，包括VPC,CCN
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

	// all_squash：所有访问用户（含 root 用户）都会被映射为匿名用户或用户组。
	// no_all_squash：所有访问用户（含 root 用户）均保持原有的 UID/GID 信息。
	// root_squash：将来访的 root 用户映射为匿名用户或用户组，非 root 用户保持原有的 UID/GID 信息。
	// no_root_squash：与 no_all_squash 效果一致，所有访问用户（含 root 用户）均保持原有的 UID/GID 信息
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 规则优先级，1-100。 其中 1 为最高，100为最低
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type PathInfo struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 目录绝对路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

// Predefined struct for user
type ScaleUpFileSystemRequestParams struct {
	// 文件系统Id,该参数通过查询文件系统列表接口获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 扩容的目标容量（单位GiB）
	TargetCapacity *uint64 `json:"TargetCapacity,omitnil,omitempty" name:"TargetCapacity"`
}

type ScaleUpFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统Id,该参数通过查询文件系统列表接口获取
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
	// 文件系统 ID,通过[查询文件系统列表](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid，Dir，分别代表用户配额，用户组配额，目录配额
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB。设置范围10-10000000。
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitnil,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个。设置范围1000-100000000
	FileHardLimit *uint64 `json:"FileHardLimit,omitnil,omitempty" name:"FileHardLimit"`

	// 需设置目录配额的目录绝对路径，不同目录不可存在包含关系
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`
}

type SetUserQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID,通过[查询文件系统列表](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 指定配额类型，包括Uid、Gid，Dir，分别代表用户配额，用户组配额，目录配额
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 容量硬限制，单位GiB。设置范围10-10000000。
	CapacityHardLimit *uint64 `json:"CapacityHardLimit,omitnil,omitempty" name:"CapacityHardLimit"`

	// 文件硬限制，单位个。设置范围1000-100000000
	FileHardLimit *uint64 `json:"FileHardLimit,omitnil,omitempty" name:"FileHardLimit"`

	// 需设置目录配额的目录绝对路径，不同目录不可存在包含关系
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`
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
	delete(f, "DirectoryPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetUserQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetUserQuotaResponseParams struct {
	// UID/GID信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

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

	// 快照状态，creating-创建中；available-运行中；deleting-删除中；rollbacking-new 创建新文件系统中；create-failed 创建失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 快照大小
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 保留时长天
	AliveDay *uint64 `json:"AliveDay,omitnil,omitempty" name:"AliveDay"`

	// 快照进度百分比，1表示1% 范围1-100
	Percent *uint64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 账号ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 快照删除时间
	DeleteTime *string `json:"DeleteTime,omitnil,omitempty" name:"DeleteTime"`

	// 文件系统名称
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`

	// 快照标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 快照类型，general为通用系列快照，turbo为Turbo系列快照
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotType *string `json:"SnapshotType,omitnil,omitempty" name:"SnapshotType"`

	// 实际快照时间，反映快照对应文件系统某个时刻的数据。
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

	// 快照总容量，单位是MiB
	SnapshotSize *uint64 `json:"SnapshotSize,omitnil,omitempty" name:"SnapshotSize"`
}

// Predefined struct for user
type StopLifecycleDataTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopLifecycleDataTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *StopLifecycleDataTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLifecycleDataTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLifecycleDataTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLifecycleDataTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopLifecycleDataTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopLifecycleDataTaskResponseParams `json:"Response"`
}

func (r *StopLifecycleDataTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLifecycleDataTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrationTaskRequestParams struct {
	// 迁移任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type StopMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务Id
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

	// 迁移状态。0: 已完成；1: 创建中；2: 运行中；3: 终止中；4: 已终止；5: 创建失败；6: 运行失败；7: 结束中；8: 删除中；9: 等待中
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
	TieringSizeInBytes *int64 `json:"TieringSizeInBytes,omitnil,omitempty" name:"TieringSizeInBytes"`

	// 冷存储容量
	SecondaryTieringSizeInBytes *int64 `json:"SecondaryTieringSizeInBytes,omitnil,omitempty" name:"SecondaryTieringSizeInBytes"`
}

// Predefined struct for user
type UnbindAutoSnapshotPolicyRequestParams struct {
	// 需要解绑的文件系统ID列表，用"," 分割，文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`

	// 解绑的快照策略ID，可以通过[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208) 查询获取
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`
}

type UnbindAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 需要解绑的文件系统ID列表，用"," 分割，文件系统ID，通过查询文件系统列表获取；[DescribeCfsFileSystems](https://cloud.tencent.com/document/product/582/38170)
	FileSystemIds *string `json:"FileSystemIds,omitnil,omitempty" name:"FileSystemIds"`

	// 解绑的快照策略ID，可以通过[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208) 查询获取
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
	// 解绑的快照策略ID，可以通过[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208) 查询获取
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称，不超过64个字符
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照定期备份，按照星期一到星期日。 1代表星期一，7代表星期日，与DayOfMonth，IntervalDays 三者选一个
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 快照保留天数
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 是否激活定期快照功能；1代表激活，0代表未激活
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 定期快照在每月的第几天创建快照，该参数与DayOfWeek,IntervalDays 三者选一
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数定期执行快照，该参数与DayOfWeek,DayOfMonth 三者选一
	IntervalDays *uint64 `json:"IntervalDays,omitnil,omitempty" name:"IntervalDays"`
}

type UpdateAutoSnapshotPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 解绑的快照策略ID，可以通过[DescribeAutoSnapshotPolicies](https://cloud.tencent.com/document/api/582/80208) 查询获取
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitnil,omitempty" name:"AutoSnapshotPolicyId"`

	// 快照策略名称，不超过64个字符
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 快照定期备份，按照星期一到星期日。 1代表星期一，7代表星期日，与DayOfMonth，IntervalDays 三者选一个
	DayOfWeek *string `json:"DayOfWeek,omitnil,omitempty" name:"DayOfWeek"`

	// 快照定期备份在一天的哪一小时
	Hour *string `json:"Hour,omitnil,omitempty" name:"Hour"`

	// 快照保留天数
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`

	// 是否激活定期快照功能；1代表激活，0代表未激活
	IsActivated *uint64 `json:"IsActivated,omitnil,omitempty" name:"IsActivated"`

	// 定期快照在每月的第几天创建快照，该参数与DayOfWeek,IntervalDays 三者选一
	DayOfMonth *string `json:"DayOfMonth,omitnil,omitempty" name:"DayOfMonth"`

	// 间隔天数定期执行快照，该参数与DayOfWeek,DayOfMonth 三者选一
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
	// 文件系统 ID,通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称，64字节内的中文字母数字或者 _,-,与CreationToken 至少填一个
	FsName *string `json:"FsName,omitnil,omitempty" name:"FsName"`
}

type UpdateCfsFileSystemNameRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID,通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 用户自定义文件系统名称，64字节内的中文字母数字或者 _,-,与CreationToken 至少填一个
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统 ID，通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 文件系统 ID，通过[查询文件系统接口](https://cloud.tencent.com/document/api/582/38170)获取
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

	// 文件系统ID，目前仅支持标准型文件系统。该参数通过查询文件系统列表获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type UpdateCfsFileSystemSizeLimitRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统容量限制大小，输入范围0-1073741824, 单位为GB；其中输入值为0时，表示不限制文件系统容量。
	FsLimit *uint64 `json:"FsLimit,omitnil,omitempty" name:"FsLimit"`

	// 文件系统ID，目前仅支持标准型文件系统。该参数通过查询文件系统列表获取
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 权限组名称，1-64个字符且只能为中文，字母，数字，下划线或横线
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权限组描述信息，1-255个字符。 Name和Descinfo不能同时为空
	DescInfo *string `json:"DescInfo,omitnil,omitempty" name:"DescInfo"`
}

type UpdateCfsPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
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
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID，可通过[DescribeCfsRules](https://cloud.tencent.com/document/api/582/38156)接口获取
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限, 值为RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为all_squash、no_all_squash、root_squash、no_root_squash，默认值为root_squash
	// all_squash：所有访问用户（含 root 用户）都会被映射为匿名用户或用户组。
	// no_all_squash：所有访问用户（含 root 用户）均保持原有的 UID/GID 信息。
	// root_squash：将来访的 root 用户映射为匿名用户或用户组，非 root 用户保持原有的 UID/GID 信息。
	// no_root_squash：与 no_all_squash 效果一致，所有访问用户（含 root 用户）均保持原有的 UID/GID 信息
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低，默认值为100
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type UpdateCfsRuleRequest struct {
	*tchttp.BaseRequest
	
	// 权限组 ID，可通过[DescribeCfsPGroups接口](https://cloud.tencent.com/document/api/582/38157)获取
	PGroupId *string `json:"PGroupId,omitnil,omitempty" name:"PGroupId"`

	// 规则 ID，可通过[DescribeCfsRules](https://cloud.tencent.com/document/api/582/38156)接口获取
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 可以填写单个 IP 或者单个网段，例如 10.1.10.11 或者 10.10.1.0/24。默认来访地址为*表示允许所有。同时需要注意，此处需填写 CVM 的内网 IP。
	AuthClientIp *string `json:"AuthClientIp,omitnil,omitempty" name:"AuthClientIp"`

	// 读写权限, 值为RO、RW；其中 RO 为只读，RW 为读写，不填默认为只读
	RWPermission *string `json:"RWPermission,omitnil,omitempty" name:"RWPermission"`

	// 用户权限，值为all_squash、no_all_squash、root_squash、no_root_squash，默认值为root_squash
	// all_squash：所有访问用户（含 root 用户）都会被映射为匿名用户或用户组。
	// no_all_squash：所有访问用户（含 root 用户）均保持原有的 UID/GID 信息。
	// root_squash：将来访的 root 用户映射为匿名用户或用户组，非 root 用户保持原有的 UID/GID 信息。
	// no_root_squash：与 no_all_squash 效果一致，所有访问用户（含 root 用户）均保持原有的 UID/GID 信息
	UserPermission *string `json:"UserPermission,omitnil,omitempty" name:"UserPermission"`

	// 规则优先级，参数范围1-100。 其中 1 为最高，100为最低，默认值为100
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
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 文件系统快照名称，与AliveDays 必须填一个，快照名称，支持不超过64字符长度，支持中文、数字、_、-
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 文件系统快照保留天数，与SnapshotName必须填一个，如果原来是永久保留时间，不允许修改成短期有效期
	AliveDays *uint64 `json:"AliveDays,omitnil,omitempty" name:"AliveDays"`
}

type UpdateCfsSnapshotAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 快照ID，可以通过[DescribeCfsSnapshots](https://cloud.tencent.com/document/api/582/80206) 查询获取
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 文件系统快照名称，与AliveDays 必须填一个，快照名称，支持不超过64字符长度，支持中文、数字、_、-
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 文件系统快照保留天数，与SnapshotName必须填一个，如果原来是永久保留时间，不允许修改成短期有效期
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
	// 文件系统 ID,可通过[DescribeCfsFileSystems](https://cloud.tencent.com/document/api/582/38170)接口获取
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统带宽，仅吞吐型可填。单位MiB/s，最小为1GiB/s，最大200GiB/s。
	BandwidthLimit *int64 `json:"BandwidthLimit,omitnil,omitempty" name:"BandwidthLimit"`
}

type UpdateFileSystemBandwidthLimitRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统 ID,可通过[DescribeCfsFileSystems](https://cloud.tencent.com/document/api/582/38170)接口获取
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
	// 指定配额类型，包括Uid、Gid、Dir
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
	CapacityUsed *uint64 `json:"CapacityUsed,omitnil,omitempty" name:"CapacityUsed"`

	// 文件使用个数，单位个
	FileUsed *uint64 `json:"FileUsed,omitnil,omitempty" name:"FileUsed"`

	// 目录配额的目录绝对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	DirectoryPath *string `json:"DirectoryPath,omitnil,omitempty" name:"DirectoryPath"`

	// 配置规则状态，inavailable---配置中，available --已生效，deleting--删除中，deleted 已删除，failed--配置失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}