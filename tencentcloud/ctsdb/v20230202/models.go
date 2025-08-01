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

package v20230202

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Cluster struct {
	// 用户APPID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppID *uint64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountID *string `json:"AccountID,omitnil,omitempty" name:"AccountID"`

	// 自定义实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones *string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Networks is deprecated.
	Networks []*Network `json:"Networks,omitnil,omitempty" name:"Networks"`

	// 实例规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Spec is deprecated.
	Spec *Spec `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 实例状态：0：运行中,1：创建中 ,16：变配中,17：隔离中,18：待销毁,19：恢复中,20：关机 , 21：销毁中 ,22：已销毁 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例有效期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *Period `json:"Period,omitnil,omitempty" name:"Period"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 产品内部特性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenant *Tenant `json:"Tenant,omitnil,omitempty" name:"Tenant"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 安全组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Security []*string `json:"Security,omitnil,omitempty" name:"Security"`
}

type Database struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterID *string `json:"ClusterID,omitnil,omitempty" name:"ClusterID"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 降冷时间（天）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoolDownInDays *int64 `json:"CoolDownInDays,omitnil,omitempty" name:"CoolDownInDays"`

	// 数据保留时间（天）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionInDays *int64 `json:"RetentionInDays,omitnil,omitempty" name:"RetentionInDays"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 状态：0: 资源初始化中， 1: 资源创建中， 2: 正常状态， 3: 资源删除中， 4: 资源已删除， 5: 资源禁用中， 6: 资源已禁用， 7: 资源异常，需要人工操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 当前页数	
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 单页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询参数：支持通过实例ID（cluster_id）和实例名称（name）进行过滤查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序参数：支持通过创建时间字段（created_at）进行排序，Type可指定为DESC（降序）或ASC（升序）
	Orders []*Order `json:"Orders,omitnil,omitempty" name:"Orders"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 当前页数	
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 单页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询参数：支持通过实例ID（cluster_id）和实例名称（name）进行过滤查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序参数：支持通过创建时间字段（created_at）进行排序，Type可指定为DESC（降序）或ASC（升序）
	Orders []*Order `json:"Orders,omitnil,omitempty" name:"Orders"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "Orders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 当前条件下的总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clusters []*Cluster `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// 数据库参数
	Database *Database `json:"Database,omitnil,omitempty" name:"Database"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页面
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库参数
	Database *Database `json:"Database,omitnil,omitempty" name:"Database"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页面
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Database")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// 数据库列表
	Databases []*Database `json:"Databases,omitnil,omitempty" name:"Databases"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤参数
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤表达式
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// 参与过滤的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Network struct {
	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc subnet id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// vpc ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VIP *string `json:"VIP,omitnil,omitempty" name:"VIP"`

	// vpc port地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type Order struct {
	// 排序字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序方式
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Period struct {
	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type Spec struct {
	// 1：包年包月、2：按小时计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 请求单元，为0则表示走资源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestUnit *uint64 `json:"RequestUnit,omitnil,omitempty" name:"RequestUnit"`

	// CPU 核数最大限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuLimit *float64 `json:"CpuLimit,omitnil,omitempty" name:"CpuLimit"`

	// 内存 最大限制(Gi)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryLimit *float64 `json:"MemoryLimit,omitnil,omitempty" name:"MemoryLimit"`

	// 磁盘 最大限制(Gi)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskLimit *uint64 `json:"DiskLimit,omitnil,omitempty" name:"DiskLimit"`

	// 业务分片数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Shards *uint64 `json:"Shards,omitnil,omitempty" name:"Shards"`

	// 业务节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replicas *uint64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`
}

type Tag struct {
	// 键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Tenant struct {
	// 密码是否已加密
	IsPasswordEncrypted *bool `json:"IsPasswordEncrypted,omitnil,omitempty" name:"IsPasswordEncrypted"`
}