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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 副本集个数，1为单副本集实例，大于1为分片集群实例，最大不超过10
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// 每个副本集内从节点个数，目前只支持从节点数为2
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// MongoDB引擎版本，值包括：MONGO_2、MONGO_3_MMAP、MONGO_3_WT 、MONGO_3_ROCKS和MONGO_36_WT
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 实例类型，GIO：高IO版；TGIO：高IO万兆
	Machine *string `json:"Machine,omitempty" name:"Machine"`

	// 实例数量，默认值为1, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 可用区信息，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例角色，支持值包括：MASTER-表示主实例，DR-表示灾备实例，RO-表示只读实例
	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 数据是否加密，当且仅当引擎版本为MONGO_3_ROCKS，可以选择加密
	Encrypt *uint64 `json:"Encrypt,omitempty" name:"Encrypt"`

	// 私有网络ID，如果不传则默认选择基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 项目ID，不填为默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 安全组参数
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`
}

func (r *CreateDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 每个副本集内从节点个数
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 版本号，当前仅支持 MONGO_3_WT
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// 机器类型，GIO：高IO版；TGIO：高IO万兆
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// 实例数量，默认值为1, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例所属区域名称，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 时长，购买月数
	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 实例密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 项目ID，不填为默认项目
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 安全组参数
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`

	// 私有网络ID，如果不传则默认选择基础网络
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
}

func (r *CreateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TerminateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID，表示注销实例成功
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TerminateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 升级后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`
}

func (r *UpgradeDBInstanceHourRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceHourRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceHourResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceHourResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceHourResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 升级后的内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 升级后的硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 升级后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
