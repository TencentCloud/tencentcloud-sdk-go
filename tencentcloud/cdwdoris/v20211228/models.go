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

package v20211228

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AttachCBSSpec struct {
	// 节点磁盘类型，例如“CLOUD_SSD”\"CLOUD_PREMIUM"
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 磁盘容量，单位G
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 磁盘总数
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// 描述
	DiskDesc *string `json:"DiskDesc,omitnil" name:"DiskDesc"`
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 集群角色类型，默认为 "data"数据节点
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil" name:"DisplayPolicy"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 集群角色类型，默认为 "data"数据节点
	NodeRole *string `json:"NodeRole,omitnil" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil" name:"DisplayPolicy"`
}

func (r *DescribeInstanceNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeRole")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DisplayPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例节点总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNodesList []*InstanceNode `json:"InstanceNodesList,omitnil" name:"InstanceNodesList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例描述信息
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil" name:"InstanceInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil" name:"SearchTags"`
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
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例数组
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil" name:"InstancesList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type InstanceInfo struct {
	// 集群实例ID, "cdw-xxxx" 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 集群实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 状态,
	// Init 创建中; Serving 运行中； 
	// Deleted已销毁；Deleting 销毁中；
	// Modify 集群变更中；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`

	// 地域, ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 可用区， ap-guangzhou-3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 付费类型，"hour", "prepay"
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 数据节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSummary *NodesSummary `json:"MasterSummary,omitnil" name:"MasterSummary"`

	// zookeeper节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreSummary *NodesSummary `json:"CoreSummary,omitnil" name:"CoreSummary"`

	// 高可用，“true" "false"
	// 注意：此字段可能返回 null，表示取不到有效值。
	HA *string `json:"HA,omitnil" name:"HA"`

	// 高可用类型：
	// 0：非高可用
	// 1：读高可用
	// 2：读写高可用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaType *int64 `json:"HaType,omitnil" name:"HaType"`

	// 访问地址，例如 "10.0.0.1:9000"
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfo *string `json:"AccessInfo,omitnil" name:"AccessInfo"`

	// 记录ID，数值型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// regionId, 表示地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 可用区说明，例如 "广州二区"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneDesc *string `json:"ZoneDesc,omitnil" name:"ZoneDesc"`

	// 错误流程说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil" name:"FlowMsg"`

	// 状态描述，例如“运行中”等
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil" name:"StatusDesc"`

	// 自动续费标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *bool `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 监控信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Monitor *string `json:"Monitor,omitnil" name:"Monitor"`

	// 是否开通日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasClsTopic *bool `json:"HasClsTopic,omitnil" name:"HasClsTopic"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsTopicId *string `json:"ClsTopicId,omitnil" name:"ClsTopicId"`

	// 日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsLogSetId *string `json:"ClsLogSetId,omitnil" name:"ClsLogSetId"`

	// 是否支持xml配置管理
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitnil" name:"EnableXMLConfig"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDesc *string `json:"RegionDesc,omitnil" name:"RegionDesc"`

	// 弹性网卡地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// 冷热分层系数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosMoveFactor *int64 `json:"CosMoveFactor,omitnil" name:"CosMoveFactor"`

	// external/local/yunti
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil" name:"Kind"`

	// cos桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// cbs
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil" name:"CanAttachCbs"`

	// 小版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildVersion *string `json:"BuildVersion,omitnil" name:"BuildVersion"`

	// 组件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components *string `json:"Components,omitnil" name:"Components"`
}

type InstanceNode struct {
	// IP地址
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 机型，如 S1
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// cpu核数
	Core *int64 `json:"Core,omitnil" name:"Core"`

	// 内存大小
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 磁盘大小
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 所属clickhouse cluster名称
	Role *string `json:"Role,omitnil" name:"Role"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// rip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rip *string `json:"Rip,omitnil" name:"Rip"`

	// FE节点角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeRole *string `json:"FeRole,omitnil" name:"FeRole"`

	// UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUID *string `json:"UUID,omitnil" name:"UUID"`
}

type NodesSummary struct {
	// 机型，如 S1
	Spec *string `json:"Spec,omitnil" name:"Spec"`

	// 节点数目
	NodeSize *int64 `json:"NodeSize,omitnil" name:"NodeSize"`

	// cpu核数，单位个
	Core *int64 `json:"Core,omitnil" name:"Core"`

	// 内存大小，单位G
	Memory *int64 `json:"Memory,omitnil" name:"Memory"`

	// 磁盘大小，单位G
	Disk *int64 `json:"Disk,omitnil" name:"Disk"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 磁盘描述
	DiskDesc *string `json:"DiskDesc,omitnil" name:"DiskDesc"`

	// 挂载云盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitnil" name:"AttachCBSSpec"`

	// 子产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductType *string `json:"SubProductType,omitnil" name:"SubProductType"`

	// 规格核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCore *int64 `json:"SpecCore,omitnil" name:"SpecCore"`

	// 规格内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecMemory *int64 `json:"SpecMemory,omitnil" name:"SpecMemory"`

	// 磁盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// 是否加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *int64 `json:"Encrypt,omitnil" name:"Encrypt"`

	// 最大磁盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil" name:"MaxDiskSize"`
}

type SearchTags struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil" name:"AllValue"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}