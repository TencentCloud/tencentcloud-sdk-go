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
}

type Placement struct {

	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属的可用区ID。该参数也可以通过调用 DescribeZones 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
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

type VPCSettings struct {

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}
