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

	// COS SecretId
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// COS SecrectKey
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// 日志存储在COS上的路径
	LogOnCosPath *string `json:"LogOnCosPath,omitempty" name:"LogOnCosPath"`
}

type ClusterInstancesInfo struct {

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ftitle *string `json:"Ftitle,omitempty" name:"Ftitle"`

	// 集群名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 地区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 用户APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 用户UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 集群VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 添加时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 已经运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunTime *string `json:"RunTime,omitempty" name:"RunTime"`

	// 集群产品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *EmrProductConfigOutter `json:"Config,omitempty" name:"Config"`

	// 主节点外网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterIp *string `json:"MasterIp,omitempty" name:"MasterIp"`

	// EMR版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmrVersion *string `json:"EmrVersion,omitempty" name:"EmrVersion"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`

	// 交易版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeVersion *int64 `json:"TradeVersion,omitempty" name:"TradeVersion"`

	// 资源订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceOrderId *int64 `json:"ResourceOrderId,omitempty" name:"ResourceOrderId"`

	// 是否计费集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTradeCluster *int64 `json:"IsTradeCluster,omitempty" name:"IsTradeCluster"`

	// 集群错误状态告警信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmInfo *string `json:"AlarmInfo,omitempty" name:"AlarmInfo"`
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
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

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

	// COS设置参数
	COSSettings *COSSettings `json:"COSSettings,omitempty" name:"COSSettings"`

	// 安全组ID
	SgId *string `json:"SgId,omitempty" name:"SgId"`

	// 预执行脚本设置
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings" list`

	// 自动续费
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 客户端Token
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否需要外网Ip。支持填NEED_MASTER_WAN，不支持使用NOT_NEED_MASTER_WAN，默认使用NEED_MASTER_WAN
	NeedMasterWan *string `json:"NeedMasterWan,omitempty" name:"NeedMasterWan"`

	// 是否需要开启外网远程登录，即22号端口，在SgId不为空时，该选项无效
	RemoteLoginAtCreate *int64 `json:"RemoteLoginAtCreate,omitempty" name:"RemoteLoginAtCreate"`

	// 是否开启安全集群，0表示不开启，非0表示开启
	CheckSecurity *int64 `json:"CheckSecurity,omitempty" name:"CheckSecurity"`

	// 访问外部文件系统
	ExtendFsField *string `json:"ExtendFsField,omitempty" name:"ExtendFsField"`
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

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群展示策略，该字段取值根据所选页面不同输入不同，集群列表页：clusterList，集群监控：monitorManage，云硬件管理：cloudHardwareManage，组件管理页：componentManage
	DisplayStrategy *string `json:"DisplayStrategy,omitempty" name:"DisplayStrategy"`

	// 查询列表,  如果不填写，返回该AppId下所有实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 查询偏移量，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果限制，默认值10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 项目列表，默认值-1
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 排序字段，当前支持以下排序字段：clusterId、addTime、status
	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`

	// 排序方法，0降序，1升序
	Asc *int64 `json:"Asc,omitempty" name:"Asc"`
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
		TotalCnt *int64 `json:"TotalCnt,omitempty" name:"TotalCnt"`

		// 集群实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterList []*ClusterInstancesInfo `json:"ClusterList,omitempty" name:"ClusterList" list`

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

type EmrProductConfigOutter struct {

	// 软件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftInfo []*string `json:"SoftInfo,omitempty" name:"SoftInfo" list`

	// Master节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterNodeSize *int64 `json:"MasterNodeSize,omitempty" name:"MasterNodeSize"`

	// Core节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreNodeSize *int64 `json:"CoreNodeSize,omitempty" name:"CoreNodeSize"`

	// Task节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNodeSize *int64 `json:"TaskNodeSize,omitempty" name:"TaskNodeSize"`

	// Common节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComNodeSize *int64 `json:"ComNodeSize,omitempty" name:"ComNodeSize"`

	// Master节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterResource *OutterResource `json:"MasterResource,omitempty" name:"MasterResource"`

	// Core节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreResource *OutterResource `json:"CoreResource,omitempty" name:"CoreResource"`

	// Task节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskResource *OutterResource `json:"TaskResource,omitempty" name:"TaskResource"`

	// Common节点资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComResource *OutterResource `json:"ComResource,omitempty" name:"ComResource"`

	// 是否使用COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnCos *bool `json:"OnCos,omitempty" name:"OnCos"`

	// 收费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *int64 `json:"ChargeType,omitempty" name:"ChargeType"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 询价资源描述
	ResourceSpec *NewResourceSpec `json:"ResourceSpec,omitempty" name:"ResourceSpec"`

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

		// 刊例价
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 时间单位，"s","m"
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 时间数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

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

type InquiryPriceRenewInstanceRequest struct {
	*tchttp.BaseRequest

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 资源ID列表
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

	// 位置信息
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 计费模式，0表示按量，1表示包年报月，此处只能为包年包月
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 时间单位，默认为m
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币种类
	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

func (r *InquiryPriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刊例价
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 时间单位，"s","m"
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 时间数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceRenewInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

		// 刊例价
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *string `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *string `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
		Unit *string `json:"Unit,omitempty" name:"Unit"`

		// 询价配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		PriceSpec *PriceResource `json:"PriceSpec,omitempty" name:"PriceSpec"`

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

type InquiryPriceUpdateInstanceRequest struct {
	*tchttp.BaseRequest

	// 时间单位。s:按量用例单位。m:包年包月用例单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度。按量用例长度为3600。
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 变配参数
	UpdateSpec *UpdateInstanceSettings `json:"UpdateSpec,omitempty" name:"UpdateSpec"`

	// 计费类型
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 位置信息
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 货币种类
	Currency *string `json:"Currency,omitempty" name:"Currency"`
}

func (r *InquiryPriceUpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpdateInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 刊例价
	// 注意：此字段可能返回 null，表示取不到有效值。
		OriginalCost *float64 `json:"OriginalCost,omitempty" name:"OriginalCost"`

		// 折扣价格
	// 注意：此字段可能返回 null，表示取不到有效值。
		DiscountCost *float64 `json:"DiscountCost,omitempty" name:"DiscountCost"`

		// 时间单位，"s","m"
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

		// 时间数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpdateInstanceResponse) FromJsonString(s string) error {
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

	// 该类型云盘个数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type NewResourceSpec struct {

	// 描述Master节点资源
	MasterResourceSpec *Resource `json:"MasterResourceSpec,omitempty" name:"MasterResourceSpec"`

	// 描述Core节点资源
	CoreResourceSpec *Resource `json:"CoreResourceSpec,omitempty" name:"CoreResourceSpec"`

	// 描述Task节点资源
	TaskResourceSpec *Resource `json:"TaskResourceSpec,omitempty" name:"TaskResourceSpec"`

	// Master节点数量
	MasterCount *int64 `json:"MasterCount,omitempty" name:"MasterCount"`

	// Core节点数量
	CoreCount *int64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// Task节点数量
	TaskCount *int64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 描述Common节点资源
	CommonResourceSpec *Resource `json:"CommonResourceSpec,omitempty" name:"CommonResourceSpec"`

	// Common节点数量
	CommonCount *int64 `json:"CommonCount,omitempty" name:"CommonCount"`
}

type OutterResource struct {

	// 规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// CPU个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Placement struct {

	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属的可用区，例如ap-guangzhou-1。该参数也可以通过调用 DescribeZones 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type PreExecuteFileSettings struct {

	// 脚本在COS上路径，已废弃
	Path *string `json:"Path,omitempty" name:"Path"`

	// 执行脚本参数
	Args []*string `json:"Args,omitempty" name:"Args" list`

	// COS的Bucket名称，已废弃
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// COS的Region名称，已废弃
	Region *string `json:"Region,omitempty" name:"Region"`

	// COS的Domain数据，已废弃
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 执行顺序
	RunOrder *int64 `json:"RunOrder,omitempty" name:"RunOrder"`

	// resourceAfter 或 clusterAfter
	WhenRun *string `json:"WhenRun,omitempty" name:"WhenRun"`

	// 脚本文件名，已废弃
	CosFileName *string `json:"CosFileName,omitempty" name:"CosFileName"`

	// 脚本的cos地址
	CosFileURI *string `json:"CosFileURI,omitempty" name:"CosFileURI"`

	// cos的SecretId
	CosSecretId *string `json:"CosSecretId,omitempty" name:"CosSecretId"`

	// Cos的SecretKey
	CosSecretKey *string `json:"CosSecretKey,omitempty" name:"CosSecretKey"`

	// cos的appid，已废弃
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}

type PriceResource struct {

	// 需要的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *uint64 `json:"StorageType,omitempty" name:"StorageType"`

	// 硬盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// 核心数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 硬盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks" list`

	// 磁盘数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCnt *int64 `json:"DiskCnt,omitempty" name:"DiskCnt"`
}

type Resource struct {

	// 节点规格描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitempty" name:"StorageType"`

	// 磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 内存容量,单位为M
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitempty" name:"MemSize"`

	// CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 数据盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 系统盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RootSize *int64 `json:"RootSize,omitempty" name:"RootSize"`

	// 云盘列表，当数据盘为一块云盘时，直接使用DiskType和DiskSize参数，超出部分使用MultiDisks
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiDisks []*MultiDisk `json:"MultiDisks,omitempty" name:"MultiDisks" list`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 时间长度
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 扩容实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 付费类型
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// Token
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 预执行脚本设置
	PreExecutedFileSettings []*PreExecuteFileSettings `json:"PreExecutedFileSettings,omitempty" name:"PreExecutedFileSettings" list`

	// 扩容Task节点数量
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`

	// 扩容Core节点数量
	CoreCount *uint64 `json:"CoreCount,omitempty" name:"CoreCount"`

	// 扩容时不需要安装的进程
	UnNecessaryNodeList []*uint64 `json:"UnNecessaryNodeList,omitempty" name:"UnNecessaryNodeList" list`
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

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
		DealNames []*string `json:"DealNames,omitempty" name:"DealNames" list`

		// token
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

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

type TerminateInstanceRequest struct {
	*tchttp.BaseRequest

	// 被销毁的实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 销毁节点ID
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`
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

type UpdateInstanceSettings struct {

	// 内存容量，单位为G
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// CPU核数
	CPUCores *uint64 `json:"CPUCores,omitempty" name:"CPUCores"`

	// 机器资源ID（EMR测资源标识）
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type VPCSettings struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
