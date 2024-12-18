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

package v20180408

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AssignProjectRequestParams struct {
	// 实例 ID 列表。格式如：cmgo-p8vn****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 项目ID。项目 ID 具有唯一性，请[登录 MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在右上角的账户信息的下拉菜单中，选择**项目管理**，即可获取项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type AssignProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 列表。格式如：cmgo-p8vn****，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 项目ID。项目 ID 具有唯一性，请[登录 MongoDB 控制台](https://console.cloud.tencent.com/mongodb)，在右上角的账户信息的下拉菜单中，选择**项目管理**，即可获取项目ID。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *AssignProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignProjectResponseParams struct {
	// 返回的异步任务ID列表。
	FlowIds []*uint64 `json:"FlowIds,omitnil,omitempty" name:"FlowIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssignProjectResponse struct {
	*tchttp.BaseResponse
	Response *AssignProjectResponseParams `json:"Response"`
}

func (r *AssignProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientConnection struct {
	// 连接的客户端IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 对应客户端IP的连接数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type CreateDBInstanceHourRequestParams struct {
	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 副本集个数，1为单副本集实例，大于1为分片集群实例，最大不超过10
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// 每个副本集内从节点个数，目前只支持从节点数为2
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// MongoDB引擎版本，值包括MONGO_3_WT 、MONGO_3_ROCKS和MONGO_36_WT
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 实例类型，GIO：高IO版；TGIO：高IO万兆
	Machine *string `json:"Machine,omitnil,omitempty" name:"Machine"`

	// 实例数量，默认值为1, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例角色，支持值包括：MASTER-表示主实例，DR-表示灾备实例，RO-表示只读实例
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 数据是否加密，当且仅当引擎版本为MONGO_3_ROCKS，可以选择加密
	Encrypt *uint64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 私有网络ID，如果不传则默认选择基础网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 项目ID，不填为默认项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组参数
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 副本集个数，1为单副本集实例，大于1为分片集群实例，最大不超过10
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitnil,omitempty" name:"ReplicateSetNum"`

	// 每个副本集内从节点个数，目前只支持从节点数为2
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// MongoDB引擎版本，值包括MONGO_3_WT 、MONGO_3_ROCKS和MONGO_36_WT
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 实例类型，GIO：高IO版；TGIO：高IO万兆
	Machine *string `json:"Machine,omitnil,omitempty" name:"Machine"`

	// 实例数量，默认值为1, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 可用区信息，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例角色，支持值包括：MASTER-表示主实例，DR-表示灾备实例，RO-表示只读实例
	InstanceRole *string `json:"InstanceRole,omitnil,omitempty" name:"InstanceRole"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 数据是否加密，当且仅当引擎版本为MONGO_3_ROCKS，可以选择加密
	Encrypt *uint64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 私有网络ID，如果不传则默认选择基础网络
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 项目ID，不填为默认项目
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组参数
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "ReplicateSetNum")
	delete(f, "SecondaryNum")
	delete(f, "EngineVersion")
	delete(f, "Machine")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "InstanceRole")
	delete(f, "InstanceType")
	delete(f, "Encrypt")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "SecurityGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceHourResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 创建的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceHourResponseParams `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceRequestParams struct {
	// 每个副本集内从节点个数
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 指版本信息。具体支持的版本信息 ，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。 - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。 - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。 - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。 - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。 - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。 - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 机器类型，HIO10G：高IO万兆。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例数量，默认值为1, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例所属区域名称，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 时长，购买月数
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 项目ID，不填为默认项目
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组参数
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 私有网络ID，如果不传则默认选择基础网络
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群，默认为REPLSET
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 每个副本集内从节点个数
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 指版本信息。具体支持的版本信息 ，请通过接口 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567) 获取。 - MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本。 - MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本。 - MONGO_42_WT：MongoDB 4.2 WiredTiger存储引擎版本。 - MONGO_44_WT：MongoDB 4.4 WiredTiger存储引擎版本。 - MONGO_50_WT：MongoDB 5.0 WiredTiger存储引擎版本。 - MONGO_60_WT：MongoDB 6.0 WiredTiger存储引擎版本。
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 机器类型，HIO10G：高IO万兆。
	MachineCode *string `json:"MachineCode,omitnil,omitempty" name:"MachineCode"`

	// 实例数量，默认值为1, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 实例所属区域名称，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 时长，购买月数
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 实例密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 项目ID，不填为默认项目
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 安全组参数
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 私有网络ID，如果不传则默认选择基础网络
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群，默认为REPLSET
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecondaryNum")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "MongoVersion")
	delete(f, "MachineCode")
	delete(f, "GoodsNum")
	delete(f, "Zone")
	delete(f, "TimeSpan")
	delete(f, "Password")
	delete(f, "ProjectId")
	delete(f, "SecurityGroup")
	delete(f, "UniqVpcId")
	delete(f, "UniqSubnetId")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstanceResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 创建的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstanceResponseParams `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeClientConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeClientConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientConnectionsResponseParams struct {
	// 客户端连接信息，包括客户端IP和对应IP的连接数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clients []*ClientConnection `json:"Clients,omitnil,omitempty" name:"Clients"`

	// 连接数总结
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientConnectionsResponseParams `json:"Response"`
}

func (r *DescribeClientConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 实例ID列表，格式如：cmgo-p8vn****。与云数据库控制台页面中显示的实例ID相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例类型，取值范围：
	// <ul><li>0： 所有实例</li><li>1： 正式实例</li><li>2： 临时实例</li><li>3： 只读实例</li><li>-1： 正式实例+只读+灾备实例</li></ul>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群类型，取值范围： 
	// <ul><li>0： 副本集实例</li><li>1： 正式实例</li> <li>-1： 所有实例</li></ul>
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 实例状态，取值范围： 
	// <ul><li>0： 待初始化</li><li>1： 流程执行中</li> <li>2： 有效实例</li><li>-2： 已过期实例</li></ul>
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络的ID，基础网络则不传该参数。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网ID，基础网络则不传该参数。入参设置该参数的同时，必须设置相应的VpcId。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 付费类型，取值范围：
	// <ul><li>0： 按量计费</li><li>1：包年包月</li><li>-1： 按量计费+包年包月</li></ul>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 单次请求返回的数量，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回结果集排序的字段，目前支持： 
	// <ul><li>ProjectId： 按照项目ID排序</li><li>InstanceName：按照实例名称排序</li><li>CreateTime： 根据创建时间排序</li></ul>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果集排序方式，目前支持："ASC"或者"DESC"。
	// <ul><li>ASC： 顺序取值</li><li>DESC：倒序取值</li></ul>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，格式如：cmgo-p8vn****。与云数据库控制台页面中显示的实例ID相同。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例类型，取值范围：
	// <ul><li>0： 所有实例</li><li>1： 正式实例</li><li>2： 临时实例</li><li>3： 只读实例</li><li>-1： 正式实例+只读+灾备实例</li></ul>
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群类型，取值范围： 
	// <ul><li>0： 副本集实例</li><li>1： 正式实例</li> <li>-1： 所有实例</li></ul>
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 实例状态，取值范围： 
	// <ul><li>0： 待初始化</li><li>1： 流程执行中</li> <li>2： 有效实例</li><li>-2： 已过期实例</li></ul>
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络的ID，基础网络则不传该参数。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网ID，基础网络则不传该参数。入参设置该参数的同时，必须设置相应的VpcId。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 付费类型，取值范围：
	// <ul><li>0： 按量计费</li><li>1：包年包月</li><li>-1： 按量计费+包年包月</li></ul>
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 单次请求返回的数量，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回结果集排序的字段，目前支持： 
	// <ul><li>ProjectId： 按照项目ID排序</li><li>InstanceName：按照实例名称排序</li><li>CreateTime： 根据创建时间排序</li></ul>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回结果集排序方式，目前支持："ASC"或者"DESC"。
	// <ul><li>ASC： 顺序取值</li><li>DESC：倒序取值</li></ul>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "InstanceType")
	delete(f, "ClusterType")
	delete(f, "Status")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "PayMode")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 符合查询条件的实例总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息
	InstanceDetails []*MongoDBInstanceDetail `json:"InstanceDetails,omitnil,omitempty" name:"InstanceDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitnil,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SlowMS")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogResponseParams struct {
	// 符合查询条件的慢查询日志总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合查询条件的慢查询日志详情。
	SlowLogList []*string `json:"SlowLogList,omitnil,omitempty" name:"SlowLogList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogResponseParams `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *DescribeSpecInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecInfoResponseParams struct {
	// 实例售卖规格信息列表
	SpecInfoList []*SpecificationInfo `json:"SpecInfoList,omitnil,omitempty" name:"SpecInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpecInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecInfoResponseParams `json:"Response"`
}

func (r *DescribeSpecInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MongoDBInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type MongoDBInstanceDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 付费类型，可能的返回值：1-包年包月；0-按量计费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 集群类型，可能的返回值：0-副本集实例，1-分片实例，
	ClusterType *uint64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 网络类型，可能的返回值：0-基础网络，1-私有网络
	NetType *uint64 `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 私有网络的ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络的子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例状态，可能的返回值：0-待初始化，1-流程处理中，2-运行中，-2-实例已过期
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 端口号
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例到期时间
	DeadLine *string `json:"DeadLine,omitnil,omitempty" name:"DeadLine"`

	// 实例版本信息
	MongoVersion *string `json:"MongoVersion,omitnil,omitempty" name:"MongoVersion"`

	// 实例内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例磁盘规格，单位为MB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 实例CPU核心数
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 实例机器类型
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// 实例从节点数
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// 实例分片数
	ReplicationSetNum *uint64 `json:"ReplicationSetNum,omitnil,omitempty" name:"ReplicationSetNum"`

	// 实例自动续费标志，可能的返回值：0-手动续费，1-自动续费，2-确认不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 已用容量，单位MB
	UsedVolume *uint64 `json:"UsedVolume,omitnil,omitempty" name:"UsedVolume"`

	// 维护窗口起始时间
	MaintenanceStart *string `json:"MaintenanceStart,omitnil,omitempty" name:"MaintenanceStart"`

	// 维护窗口结束时间
	MaintenanceEnd *string `json:"MaintenanceEnd,omitnil,omitempty" name:"MaintenanceEnd"`

	// 分片信息
	ReplicaSets []*MongodbShardInfo `json:"ReplicaSets,omitnil,omitempty" name:"ReplicaSets"`

	// 只读实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadonlyInstances []*MongoDBInstance `json:"ReadonlyInstances,omitnil,omitempty" name:"ReadonlyInstances"`

	// 灾备实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbyInstances []*MongoDBInstance `json:"StandbyInstances,omitnil,omitempty" name:"StandbyInstances"`

	// 临时实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneInstances []*MongoDBInstance `json:"CloneInstances,omitnil,omitempty" name:"CloneInstances"`

	// 关联实例信息，对于正式实例，该字段表示它的临时实例信息；对于临时实例，则表示它的正式实例信息;如果为只读/灾备实例,则表示他的主实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedInstance *MongoDBInstance `json:"RelatedInstance,omitnil,omitempty" name:"RelatedInstance"`

	// 实例标签信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 实例标记
	InstanceVer *uint64 `json:"InstanceVer,omitnil,omitempty" name:"InstanceVer"`

	// 实例标记
	ClusterVer *uint64 `json:"ClusterVer,omitnil,omitempty" name:"ClusterVer"`

	// 协议信息，可能的返回值：1-mongodb，2-dynamodb
	Protocol *uint64 `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 实例类型，可能的返回值，1-正式实例，2-临时实例，3-只读实例，4-灾备实例
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态描述
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitnil,omitempty" name:"InstanceStatusDesc"`

	// 实例对应的物理实例ID，回档并替换过的实例有不同的InstanceId和RealInstanceId，从barad获取监控数据等场景下需要用物理id获取
	RealInstanceId *string `json:"RealInstanceId,omitnil,omitempty" name:"RealInstanceId"`
}

type MongodbShardInfo struct {
	// 分片已使用容量
	UsedVolume *float64 `json:"UsedVolume,omitnil,omitempty" name:"UsedVolume"`

	// 分片ID
	ReplicaSetId *string `json:"ReplicaSetId,omitnil,omitempty" name:"ReplicaSetId"`

	// 分片名
	ReplicaSetName *string `json:"ReplicaSetName,omitnil,omitempty" name:"ReplicaSetName"`

	// 分片内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 分片磁盘规格，单位为MB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 分片Oplog大小，单位为MB
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`

	// 分片从节点数
	SecondaryNum *uint64 `json:"SecondaryNum,omitnil,omitempty" name:"SecondaryNum"`

	// 分片物理ID
	RealReplicaSetId *string `json:"RealReplicaSetId,omitnil,omitempty" name:"RealReplicaSetId"`
}

// Predefined struct for user
type RenameInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例自定义名称
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

type RenameInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例自定义名称
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenameInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenameInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenameInstanceResponseParams `json:"Response"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAutoRenewRequestParams struct {
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 配置自动续费标识。
	// - 0：手动续费。
	// - 1：自动续费。
	// - 2：确认不续费。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type SetAutoRenewRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 配置自动续费标识。
	// - 0：手动续费。
	// - 1：自动续费。
	// - 2：确认不续费。
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAutoRenewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAutoRenewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *SetAutoRenewResponseParams `json:"Response"`
}

func (r *SetAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoRenewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPasswordRequestParams struct {
	// 实例ID，格式如：cmgo-p8vn****。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例账户名。初始化实例密码，本参数传mongouser。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 指定账户的新密码， 密码格式为8-32个字符长度，至少包含字母、数字和字符（!@#%^*()_）中的两种
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type SetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vn****。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例账户名。初始化实例密码，本参数传mongouser。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 指定账户的新密码， 密码格式为8-32个字符长度，至少包含字母、数字和字符（!@#%^*()_）中的两种
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *SetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPasswordResponseParams struct {
	// 返回的异步任务ID
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *SetPasswordResponseParams `json:"Response"`
}

func (r *SetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpecItem struct {
	// 规格信息标识
	SpecCode *string `json:"SpecCode,omitnil,omitempty" name:"SpecCode"`

	// 规格有效标志，取值：0-停止售卖，1-开放售卖
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 机器类型，取值：0-HIO，4-HIO10G
	MachineType *string `json:"MachineType,omitnil,omitempty" name:"MachineType"`

	// cpu核心数
	Cpu *uint64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 默认磁盘规格，单位MB
	DefaultStorage *uint64 `json:"DefaultStorage,omitnil,omitempty" name:"DefaultStorage"`

	// 最大磁盘规格，单位MB
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 最小磁盘规格，单位MB
	MinStorage *uint64 `json:"MinStorage,omitnil,omitempty" name:"MinStorage"`

	// 可承载qps信息
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 连接数限制
	Conns *uint64 `json:"Conns,omitnil,omitempty" name:"Conns"`

	// 实例mongodb版本信息
	MongoVersionCode *string `json:"MongoVersionCode,omitnil,omitempty" name:"MongoVersionCode"`

	// 实例mongodb版本号
	MongoVersionValue *uint64 `json:"MongoVersionValue,omitnil,omitempty" name:"MongoVersionValue"`

	// 实例mongodb版本号（短）
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 存储引擎
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 集群类型，取值：1-分片集群，0-副本集集群
	ClusterType *uint64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 最小副本集从节点数
	MinNodeNum *uint64 `json:"MinNodeNum,omitnil,omitempty" name:"MinNodeNum"`

	// 最大副本集从节点数
	MaxNodeNum *uint64 `json:"MaxNodeNum,omitnil,omitempty" name:"MaxNodeNum"`

	// 最小分片数
	MinReplicateSetNum *uint64 `json:"MinReplicateSetNum,omitnil,omitempty" name:"MinReplicateSetNum"`

	// 最大分片数
	MaxReplicateSetNum *uint64 `json:"MaxReplicateSetNum,omitnil,omitempty" name:"MaxReplicateSetNum"`

	// 最小分片从节点数
	MinReplicateSetNodeNum *uint64 `json:"MinReplicateSetNodeNum,omitnil,omitempty" name:"MinReplicateSetNodeNum"`

	// 最大分片从节点数
	MaxReplicateSetNodeNum *uint64 `json:"MaxReplicateSetNodeNum,omitnil,omitempty" name:"MaxReplicateSetNodeNum"`
}

type SpecificationInfo struct {
	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 售卖规格信息
	SpecItems []*SpecItem `json:"SpecItems,omitnil,omitempty" name:"SpecItems"`
}

type TagInfo struct {
	// 标签Key值
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

// Predefined struct for user
type TerminateDBInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDBInstanceResponseParams struct {
	// 订单ID，表示注销实例成功。
	AsyncRequestId *string `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDBInstanceResponseParams `json:"Response"`
}

func (r *TerminateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceHourRequestParams struct {
	// 实例ID，格式如：cmgo-p8vnipr5
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 升级后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`
}

type UpgradeDBInstanceHourRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-p8vnipr5
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 升级后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`
}

func (r *UpgradeDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceHourRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "OplogSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceHourRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceHourResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceHourResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
	// 实例ID，格式如：cmgo-iga0****。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 升级后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：cmgo-iga0****。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 升级后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitnil,omitempty" name:"OplogSize"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Volume")
	delete(f, "OplogSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// 订单ID
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}