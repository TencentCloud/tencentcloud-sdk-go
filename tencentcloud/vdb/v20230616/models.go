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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 实例ID数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称，支持模糊搜索。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例模糊搜索字段。
	InstanceKeys []*string `json:"InstanceKeys,omitnil,omitempty" name:"InstanceKeys"`

	// 根据状态获取实例， 为空则获取全部非隔离和非下线的实例。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 按照引擎筛选实例。
	EngineNames []*string `json:"EngineNames,omitnil,omitempty" name:"EngineNames"`

	// 按照版本筛选实例。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 按照创建时间筛选实例。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 按照可用区筛选实例。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 排序字段。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 查询开始位置。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照标签筛选实例
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID数组。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称，支持模糊搜索。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 实例模糊搜索字段。
	InstanceKeys []*string `json:"InstanceKeys,omitnil,omitempty" name:"InstanceKeys"`

	// 根据状态获取实例， 为空则获取全部非隔离和非下线的实例。
	Status []*string `json:"Status,omitnil,omitempty" name:"Status"`

	// 按照引擎筛选实例。
	EngineNames []*string `json:"EngineNames,omitnil,omitempty" name:"EngineNames"`

	// 按照版本筛选实例。
	EngineVersions []*string `json:"EngineVersions,omitnil,omitempty" name:"EngineVersions"`

	// 按照创建时间筛选实例。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 按照可用区筛选实例。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 排序字段。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 查询开始位置。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 列表查询数量。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照标签筛选实例
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
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
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "InstanceKeys")
	delete(f, "Status")
	delete(f, "EngineNames")
	delete(f, "EngineVersions")
	delete(f, "CreateAt")
	delete(f, "Zones")
	delete(f, "OrderBy")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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

type InstanceInfo struct {
	// 实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例自定义名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户APPID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 地域。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 产品。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 网络信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Networks []*Network `json:"Networks,omitnil,omitempty" name:"Networks"`

	// 分片信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardNum *uint64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 副本数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicaNum *uint64 `json:"ReplicaNum,omitnil,omitempty" name:"ReplicaNum"`

	// CPU.
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *float64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Disk *uint64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 健康得分。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthScore *float64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 异常告警。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warning *int64 `json:"Warning,omitnil,omitempty" name:"Warning"`

	// 所属项目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Project *string `json:"Project,omitnil,omitempty" name:"Project"`

	// 所属标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 资源状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 引擎名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 引擎版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 计费模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 差异化扩展信息, json格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extend *string `json:"Extend,omitnil,omitempty" name:"Extend"`

	// 过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// 是否不过期(永久)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNoExpired *bool `json:"IsNoExpired,omitnil,omitempty" name:"IsNoExpired"`

	// 外网地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WanAddress *string `json:"WanAddress,omitnil,omitempty" name:"WanAddress"`
}

type Network struct {
	// VpcId(VPC网络下有效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网Id(VPC网络下有效)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 内网访问IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网访问Port。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}