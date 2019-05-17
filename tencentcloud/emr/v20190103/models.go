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

package v20190103

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type COSSettings struct {

	// 日志存储在COS上的路径
	LogOnCosPath *string `json:"LogOnCosPath,omitempty" name:"LogOnCosPath"`

	// COS SecretId
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// COS SecrectKey
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`
}

type ClusterInfoResult struct {

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCnt *uint64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

	// 集群信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*ClusterInstanceInfo `json:"ClusterList,omitempty" name:"ClusterList" list`
}

type ClusterInstanceInfo struct {

	// clusterId
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 集群名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群地域
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 用户APPID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Addtime *string `json:"Addtime,omitempty" name:"Addtime"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Runtime *string `json:"Runtime,omitempty" name:"Runtime"`

	// 集群配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *EMRProductConfigSettings `json:"Config,omitempty" name:"Config"`

	// 集群IP
	MasterIp *string `json:"MasterIp,omitempty" name:"MasterIp"`

	// 集群版本
	EmrVersion *string `json:"EmrVersion,omitempty" name:"EmrVersion"`

	// 集群计费类型
	ChargeType *uint64 `json:"ChargeType,omitempty" name:"ChargeType"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *uint64 `json:"ProductId,omitempty" name:"ProductId"`

	// VPC设置参数
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`

	// 软件列表
	Software []*string `json:"Software,omitempty" name:"Software" list`

	// 资源描述
	ResourceSpec *ResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// 支持HA
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 计费类型
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 集群位置信息
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 登录配置
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 客户端Token
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// COS设置参数
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`

	// 安全组ID
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 预执行脚本设置
	PreExecutedFileSettings *PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

	// 自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 是否需要外网Ip。支持填NEED_MASTER_WAN，不支持使用NOT_NEED_MASTER_WAN，默认使用NEED_MASTER_WAN
	NeedMasterWan *string `json:"NeedMasterWan,omitempty" name:"NeedMasterWan"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建实例结果信息
		Result *CreateInstanceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceResult struct {

	// 客户端TOKEN
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 集群名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 订单列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 查询列表,  如果不填写，返回该AppId下所有实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 查询偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果限制，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例数量
		Result *ClusterInfoResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EMRProductConfigSettings struct {

	// 集群软件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftInfo []*string `json:"SoftInfo,omitempty" name:"SoftInfo" list`

	// master节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterNodeSize *uint64 `json:"MasterNodeSize,omitempty" name:"MasterNodeSize"`

	// core节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreNodeSize *uint64 `json:"CoreNodeSize,omitempty" name:"CoreNodeSize"`

	// task节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNodeSize *uint64 `json:"TaskNodeSize,omitempty" name:"TaskNodeSize"`

	// common节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComNodeSize *uint64 `json:"ComNodeSize,omitempty" name:"ComNodeSize"`

	// master规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResourceSpec *NodeSpec `json:"MasterResourceSpec,omitempty" name:"MasterResourceSpec"`

	// core规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResourceSpec *NodeSpec `json:"CoreResourceSpec,omitempty" name:"CoreResourceSpec"`

	// task规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResourceSpec *NodeSpec `json:"TaskResourceSpec,omitempty" name:"TaskResourceSpec"`

	// common规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonResourceSpec *NodeSpec `json:"CommonResourceSpec,omitempty" name:"CommonResourceSpec"`

	// 是否使用COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	Oncos *bool `json:"Oncos,omitempty" name:"Oncos"`

	// COS配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 询价资源描述
	ResourceSpec *ResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

	// 货币种类
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 计费类型
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 是否支持HA， 1 支持，0 不支持
	SupportHA *uint64 `json:"SupportHA,omitempty" name:"SupportHA"`

	// 软件列表
	Software []*string `json:"Software,omitempty" name:"Software" list`

	// 位置信息
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// VPC信息
	VPCSettings *VPCSettings `json:"VPCSettings,omitempty" name:"VPCSettings"`
}

func (r *InquiryPriceCreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceCreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 询价结果
		Result *InquiryPriceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceCreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceCreateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceResult struct {

	// 原始价格
	OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

	// 折扣后价格
	DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

type InquiryPriceScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// 时间单位。s:按量用例单位。m:包年包月用例单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度。按量用例长度为3600。
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// Zone ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 计费类型
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 扩容Core节点个数
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// 扩容Task节点个数
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 货币种类
	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

func (r *InquiryPriceScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceScaleOutInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扩容价格
		Result *InquiryPriceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceScaleOutInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginSettings struct {

	// Password
	Password *string `json:"Password,omitempty" name:"Password"`

	// Public Key
	PublicKeyId *string `json:"PublicKeyId,omitempty" name:"PublicKeyId"`
}

type MultiDisk struct {

	// 云盘类型("CLOUD_PREMIUM","CLOUD_SSD","CLOUD_BASIC")的一种
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 云盘大小
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`
}

type NodeSpec struct {

	// 内存容量,单位为M
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPUCores *uint64 `json:"CPUCores,omitempty" name:"CPUCores"`

	// 数据盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 节点规格描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 系统盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootDiskVolume *uint64 `json:"RootDiskVolume,omitempty" name:"RootDiskVolume"`

	// 存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *uint64 `json:"StorageType,omitempty" name:"StorageType"`

	// 规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 多云盘参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks" list`
}

type Placement struct {

	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属的可用区ID。该参数也可以通过调用 DescribeZones 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type PreExecuteFileSettings struct {

	// 脚本在COS上路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 执行脚本参数
	Args []*string `json:"Args,omitempty" name:"Args" list`

	// COS的Bucket名称
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// COS的Region名称
	Region *string `json:"Region,omitempty" name:"Region"`

	// COS的Domain数据
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type ResourceSpec struct {

	// Common节点数量
	CommonCount *uint64 `json:"CommonCount,omitempty" name:"CommonCount"`

	// 描述Master节点资源
	MasterResourceSpec *NodeSpec `json:"MasterResourceSpec,omitempty" name:"MasterResourceSpec"`

	// 描述Core节点资源
	CoreResourceSpec *NodeSpec `json:"CoreResourceSpec,omitempty" name:"CoreResourceSpec"`

	// 描述Task节点资源
	TaskResourceSpec *NodeSpec `json:"TaskResourceSpec,omitempty" name:"TaskResourceSpec"`

	// Master节点数量
	MasterCount *uint64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// Core节点数量
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Task节点数量
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 描述Common节点资源
	CommonResourceSpec *NodeSpec `json:"CommonResourceSpec,omitempty" name:"CommonResourceSpec"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// Token
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 扩容实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 付费类型
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 预执行脚本设置
	PreExecutedFileSettings *PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings"`

	// 扩容Task节点数量
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 扩容Core节点数量
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 扩容结果
		Result *ScaleOutInstanceResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ScaleOutInstanceResult struct {

	// 客户端调用时传入的TOKEN
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 扩容实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 订单名称
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`
}

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest

	// 被销毁的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 退单描述
		Result *TerminateResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateResult struct {

	// 退单集群ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资源资源ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`
}

type TerminateTasksRequest struct {
	*tchttp.BaseRequest

	// 销毁节点所属实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 销毁节点ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`
}

func (r *TerminateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 退单结果
		Result *TerminateResult `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type VPCSettings struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
