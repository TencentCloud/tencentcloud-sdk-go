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

type AssignProjectRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *AssignProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssignProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的异步任务ID列表
		FlowIds []*uint64 `json:"FlowIds,omitempty" name:"FlowIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssignProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssignProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

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

	// MongoDB引擎版本，值包括MONGO_3_WT 、MONGO_3_ROCKS和MONGO_36_WT
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

		// 创建的实例ID列表
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

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

	// 版本号，当前支持 MONGO_3_WT、MONGO_3_ROCKS、MONGO_36_WT
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

		// 创建的实例ID列表
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

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

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例类型，取值范围：0-所有实例,1-正式实例，2-临时实例, 3-只读实例，-1-正式实例+只读+灾备实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 集群类型，取值范围：0-副本集实例，1-分片实例，-1-所有实例
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 实例状态，取值范围：0-待初始化，1-流程执行中，2-实例有效，-2-实例已过期
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// 私有网络的ID，基础网络则不传该参数
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络的子网ID，基础网络则不传该参数。入参设置该参数的同时，必须设置相应的VpcId
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 付费类型，取值范围：0-按量计费，1-包年包月，-1-按量计费+包年包月
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 单次请求返回的数量，最小值为1，最大值为100，默认值为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果集排序的字段，目前支持："ProjectId", "InstanceName", "CreateTime"，默认为升序排序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回结果集排序方式，目前支持："ASC"或者"DESC"
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的实例总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息
		InstanceDetails []*MongoDBInstanceDetail `json:"InstanceDetails,omitempty" name:"InstanceDetails" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 慢日志起始时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-01 10:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 慢日志终止时间，格式：yyyy-mm-dd hh:mm:ss，如：2019-06-02 12:00:00。查询起止时间间隔不能超过24小时，只允许查询最近7天内慢日志。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢日志执行时间阈值，返回执行时间超过该阈值的慢日志，单位为毫秒(ms)，最小为100毫秒。
	SlowMS *uint64 `json:"SlowMS,omitempty" name:"SlowMS"`

	// 偏移量，最小值为0，最大值为10000，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页大小，最小值为1，最大值为100，默认值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合查询条件的慢查询日志总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合查询条件的慢查询日志详情。
		SlowLogList []*string `json:"SlowLogList,omitempty" name:"SlowLogList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

func (r *DescribeSpecInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSpecInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例售卖规格信息列表
		SpecInfoList []*SpecificationInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpecInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSpecInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MongoDBInstance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`
}

type MongoDBInstanceDetail struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 付费类型，可能的返回值：1-包年包月；0-按量计费
	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 集群类型，可能的返回值：0-副本集实例，1-分片实例，
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 网络类型，可能的返回值：0-基础网络，1-私有网络
	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`

	// 私有网络的ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例状态，可能的返回值：0-待初始化，1-流程处理中，2-运行中，-2-实例已过期
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 端口号
	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例到期时间
	DeadLine *string `json:"DeadLine,omitempty" name:"DeadLine"`

	// 实例版本信息
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// 实例内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘规格，单位为MB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 实例CPU核心数
	CpuNum *uint64 `json:"CpuNum,omitempty" name:"CpuNum"`

	// 实例机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 实例从节点数
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// 实例分片数
	ReplicationSetNum *uint64 `json:"ReplicationSetNum,omitempty" name:"ReplicationSetNum"`

	// 实例自动续费标志，可能的返回值：0-手动续费，1-自动续费，2-确认不续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 已用容量，单位MB
	UsedVolume *uint64 `json:"UsedVolume,omitempty" name:"UsedVolume"`

	// 维护窗口起始时间
	MaintenanceStart *string `json:"MaintenanceStart,omitempty" name:"MaintenanceStart"`

	// 维护窗口结束时间
	MaintenanceEnd *string `json:"MaintenanceEnd,omitempty" name:"MaintenanceEnd"`

	// 分片信息
	ReplicaSets []*MongodbShardInfo `json:"ReplicaSets,omitempty" name:"ReplicaSets" list`

	// 只读实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadonlyInstances []*MongoDBInstance `json:"ReadonlyInstances,omitempty" name:"ReadonlyInstances" list`

	// 灾备实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandbyInstances []*MongoDBInstance `json:"StandbyInstances,omitempty" name:"StandbyInstances" list`

	// 临时实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloneInstances []*MongoDBInstance `json:"CloneInstances,omitempty" name:"CloneInstances" list`

	// 关联实例信息，对于正式实例，该字段表示它的临时实例信息；对于临时实例，则表示它的正式实例信息;如果为只读/灾备实例,则表示他的主实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedInstance *MongoDBInstance `json:"RelatedInstance,omitempty" name:"RelatedInstance"`

	// 实例标签信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 实例标记
	InstanceVer *uint64 `json:"InstanceVer,omitempty" name:"InstanceVer"`

	// 实例标记
	ClusterVer *uint64 `json:"ClusterVer,omitempty" name:"ClusterVer"`

	// 协议信息，可能的返回值：1-mongodb，2-dynamodb
	Protocol *uint64 `json:"Protocol,omitempty" name:"Protocol"`

	// 实例类型，可能的返回值，1-正式实例，2-临时实例，3-只读实例，4-灾备实例
	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例状态描述
	InstanceStatusDesc *string `json:"InstanceStatusDesc,omitempty" name:"InstanceStatusDesc"`

	// 实例对应的物理实例id，回档并替换过的实例有不同的InstanceId和RealInstanceId，从barad获取监控数据等场景下需要用物理id获取
	RealInstanceId *string `json:"RealInstanceId,omitempty" name:"RealInstanceId"`
}

type MongodbShardInfo struct {

	// 分片已使用容量
	UsedVolume *float64 `json:"UsedVolume,omitempty" name:"UsedVolume"`

	// 分片ID
	ReplicaSetId *string `json:"ReplicaSetId,omitempty" name:"ReplicaSetId"`

	// 分片名
	ReplicaSetName *string `json:"ReplicaSetName,omitempty" name:"ReplicaSetName"`

	// 分片内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 分片磁盘规格，单位为MB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 分片Oplog大小，单位为MB
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`

	// 分片从节点数
	SecondaryNum *uint64 `json:"SecondaryNum,omitempty" name:"SecondaryNum"`

	// 分片物理id
	RealReplicaSetId *string `json:"RealReplicaSetId,omitempty" name:"RealReplicaSetId"`
}

type RenameInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

func (r *RenameInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenameInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenameInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenameInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoRenewRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 续费选项，取值范围：0-手动续费，1-自动续费，2-确认不续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *SetAutoRenewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoRenewRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetAutoRenewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetAutoRenewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetAutoRenewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例账户名称
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 实例新密码，至少包含字母、数字和字符（!@#%^*()）中的两种，长度为8-16个字符
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *SetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的异步任务ID
		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SpecItem struct {

	// 规格信息标识
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 规格有效标志，取值：0-停止售卖，1-开放售卖
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 机器类型，取值：0-HIO，4-HIO10G
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// cpu核心数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存规格，单位为MB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 默认磁盘规格，单位MB
	DefaultStorage *uint64 `json:"DefaultStorage,omitempty" name:"DefaultStorage"`

	// 最大磁盘规格，单位MB
	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 最小磁盘规格，单位MB
	MinStorage *uint64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 可承载qps信息
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 连接数限制
	Conns *uint64 `json:"Conns,omitempty" name:"Conns"`

	// 实例mongodb版本信息
	MongoVersionCode *string `json:"MongoVersionCode,omitempty" name:"MongoVersionCode"`

	// 实例mongodb版本号
	MongoVersionValue *uint64 `json:"MongoVersionValue,omitempty" name:"MongoVersionValue"`

	// 实例mongodb版本号（短）
	Version *string `json:"Version,omitempty" name:"Version"`

	// 存储引擎
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// 集群类型，取值：1-分片集群，0-副本集集群
	ClusterType *uint64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 最小副本集从节点数
	MinNodeNum *uint64 `json:"MinNodeNum,omitempty" name:"MinNodeNum"`

	// 最大副本集从节点数
	MaxNodeNum *uint64 `json:"MaxNodeNum,omitempty" name:"MaxNodeNum"`

	// 最小分片数
	MinReplicateSetNum *uint64 `json:"MinReplicateSetNum,omitempty" name:"MinReplicateSetNum"`

	// 最大分片数
	MaxReplicateSetNum *uint64 `json:"MaxReplicateSetNum,omitempty" name:"MaxReplicateSetNum"`

	// 最小分片从节点数
	MinReplicateSetNodeNum *uint64 `json:"MinReplicateSetNodeNum,omitempty" name:"MinReplicateSetNodeNum"`

	// 最大分片从节点数
	MaxReplicateSetNodeNum *uint64 `json:"MaxReplicateSetNodeNum,omitempty" name:"MaxReplicateSetNodeNum"`
}

type SpecificationInfo struct {

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 售卖规格信息
	SpecItem []*SpecItem `json:"SpecItem,omitempty" name:"SpecItem" list`
}

type TagInfo struct {

	// 标签Key值
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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
