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

package v20180412

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Account struct {

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账号名称（如果是主账号，名称为root）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 账号描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 读写策略：r-只读，w-只写，rw-读写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// 路由策略：master-主节点，replication-从节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`

	// 子账号状态：1-账号变更中，2-账号有效，-4-账号已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb 等。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要绑定的安全组ID，类似sg-efil73jd。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 被绑定的实例ID，类似ins-lesecurk，支持指定多个实例。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BigKeyInfo struct {

	// 所属的database
	DB *int64 `json:"DB,omitempty" name:"DB"`

	// 大Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 大小
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 数据时间戳
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type BigKeyTypeInfo struct {

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 大小
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 时间戳
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type CleanUpInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CleanUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CleanUpInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CleanUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CleanUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CleanUpInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// redis的实例密码（免密实例不需要传密码，非免密实例必传）
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ClearInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommandTake struct {

	// 命令
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// 耗时
	Took *int64 `json:"Took,omitempty" name:"Took"`
}

type CreateInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子账号名称
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 子账号密码
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 路由策略：填写master或者replication，表示主节点或者从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`

	// 读写策略：填写r、rw，表示只读、读写
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// 子账号描述信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所属的可用区ID，可参考[地域和可用区](https://cloud.tencent.com/document/product/239/4106)  。
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例类型：2 – Redis2.8内存版（标准架构），3 – CKV 3.2内存版(标准架构)，4 – CKV 3.2内存版(集群架构)，6 – Redis4.0内存版（标准架构），7 – Redis4.0内存版（集群架构），8 – Redis5.0内存版（标准架构），9 – Redis5.0内存版（集群架构）。
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// 实例容量，单位MB， 数值需为1024的整数倍，取值大小以 [查询产品售卖规格](https://cloud.tencent.com/document/api/239/30600) 返回的规格为准。
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 实例数量，单次购买实例数量以 [查询产品售卖规格](https://cloud.tencent.com/document/api/239/30600) 返回的规格为准。
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 购买时长，在创建包年包月实例的时候需要填写，按量计费实例填1即可，单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 付费方式:0-按量计费，1-包年包月。
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例密码，8-30个字符，至少包含小写字母、大写字母、数字和字符 ()`~!@#$%^&*-+=_|{}[]:;<>,.?/ 中的2种，不能以"/"开头。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 私有网络ID，如果不传则默认选择基础网络，请使用私有网络列表查询，如：vpc-sad23jfdfk。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 基础网络下， subnetId无效； vpc子网下，取值以查询子网列表，如：subnet-fdj24n34j2。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 项目id，取值以用户账户>用户账户相关接口查询>项目列表返回的projectId为准。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 自动续费标识。0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费。
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 安全组id数组。
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList" list`

	// 用户自定义的端口 不填则默认为6379，范围[1024,65535]。
	VPort *uint64 `json:"VPort,omitempty" name:"VPort"`

	// 实例分片数量，购买标准版实例不需要填写，集群版分片数量范围[3,5,8,12,16,24,32,64,96,128]。
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// 实例副本数量，Redis 2.8标准版、CKV标准版只支持1副本，4.0、5.0标准版和集群版支持1-5个副本。
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// 是否支持副本只读，Redis 2.8标准版、CKV标准版不支持副本只读，开启副本只读，实例将自动读写分离，写请求路由到主节点，读请求路由到副本节点，如需开启副本只读建议副本数>=2。
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`

	// 实例名称，长度小于60的中文/英文/数字/"-"/"_"。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 是否支持免密，true-免密实例，false-非免密实例，默认为非免密实例，仅VPC网络的实例支持免密码访问。
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易的ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 实例ID(该字段灰度中，部分地域不可见)
		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DelayDistribution struct {

	// 分布阶梯
	Ladder *int64 `json:"Ladder,omitempty" name:"Ladder"`

	// 大小
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 修改时间
	Updatetime *int64 `json:"Updatetime,omitempty" name:"Updatetime"`
}

type DeleteInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子账号名称
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`
}

func (r *DeleteInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 备份类型。自动备份类型： 1 “定时回档”
		AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`

		// Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
		WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays" list`

		// 时间段。
		TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupUrlRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份ID，通过DescribeInstanceBackups接口可查
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`
}

func (r *DescribeBackupUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 外网下载地址（6小时）
		DownloadUrl []*string `json:"DownloadUrl,omitempty" name:"DownloadUrl" list`

		// 内网下载地址（6小时）
		InnerDownloadUrl []*string `json:"InnerDownloadUrl,omitempty" name:"InnerDownloadUrl" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称，本接口取值：redis。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组规则
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 账号详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Accounts []*Account `json:"Accounts,omitempty" name:"Accounts" list`

		// 账号个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceBackupsRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID，可通过 DescribeInstance 接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例列表大小，默认大小20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 开始时间，格式如：2017-02-08 16:46:34。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 1：备份在流程中，2：备份正常，3：备份转RDB文件处理中，4：已完成RDB转换，-1：备份已过期，-2：备份已删除。
	Status []*int64 `json:"Status,omitempty" name:"Status" list`
}

func (r *DescribeInstanceBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 备份总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例的备份数组
		BackupSet []*RedisBackupSet `json:"BackupSet,omitempty" name:"BackupSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInfoRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDTSInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDTSInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// DTS任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// DTS任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		JobName *string `json:"JobName,omitempty" name:"JobName"`

		// 任务状态,取值为：1-创建中(Creating),3-校验中(Checking)4-校验通过(CheckPass),5-校验不通过（CheckNotPass）,7-任务运行(Running),8-准备完成（ReadyComplete）,9-任务成功（Success）,10-任务失败（Failed）,11-撤销中（Stopping）,12-完成中（Completing）
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

		// 同步时延，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
		Offset *int64 `json:"Offset,omitempty" name:"Offset"`

		// 断开时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CutDownTime *string `json:"CutDownTime,omitempty" name:"CutDownTime"`

		// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		SrcInfo *DescribeInstanceDTSInstanceInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

		// 目标实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		DstInfo *DescribeInstanceDTSInstanceInfo `json:"DstInfo,omitempty" name:"DstInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDTSInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDTSInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDTSInstanceInfo struct {

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 仓库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeInstanceDealDetailRequest struct {
	*tchttp.BaseRequest

	// 订单交易ID数组，即 [CreateInstances](https://cloud.tencent.com/document/api/239/20026) 的输出参数DealId。
	DealIds []*string `json:"DealIds,omitempty" name:"DealIds" list`
}

func (r *DescribeInstanceDealDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDealDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDealDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单详细信息
		DealDetails []*TradeDealDetail `json:"DealDetails,omitempty" name:"DealDetails" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDealDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceDealDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 请求类型：1——string类型，2——所有类型
	ReqType *int64 `json:"ReqType,omitempty" name:"ReqType"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 大Key详细信息
		Data []*BigKeyInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeySizeDistRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeySizeDistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeySizeDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 大Key大小分布详情
		Data []*DelayDistribution `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeySizeDistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyTypeDistRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitempty" name:"Date"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyTypeDistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorBigKeyTypeDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 大Key类型分布详细信息
		Data []*BigKeyTypeInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorBigKeyTypeDistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorHotKeyRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间范围：1——实时，2——近30分钟，3——近6小时，4——近24小时
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorHotKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorHotKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorHotKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 热Key详细信息
		Data []*HotKeyInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorHotKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorHotKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorSIPRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceMonitorSIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorSIPRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorSIPResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 访问来源信息
		Data []*SourceInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorSIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorSIPResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTookDistRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间；例如："20190219"
	Date *string `json:"Date,omitempty" name:"Date"`

	// 时间范围：1——实时，2——近30分钟，3——近6小时，4——近24小时
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTookDistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTookDistRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTookDistResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 时延分布信息
		Data []*DelayDistribution `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTookDistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTookDistResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间范围：1——实时，2——近30分钟，3——近6小时，4——近24小时
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 访问命令信息
		Data []*SourceCommand `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdTookRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间范围：1——实时，2——近30分钟，3——近6小时，4——近24小时
	SpanType *int64 `json:"SpanType,omitempty" name:"SpanType"`
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdTookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceMonitorTopNCmdTookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 耗时详细信息
		Data []*CommandTake `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceMonitorTopNCmdTookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总的修改历史记录数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 修改历史记录信息。
		InstanceParamHistory []*InstanceParamHistory `json:"InstanceParamHistory,omitempty" name:"InstanceParamHistory" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例参数个数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例枚举类型参数
		InstanceEnumParam []*InstanceEnumParam `json:"InstanceEnumParam,omitempty" name:"InstanceEnumParam" list`

		// 实例整型参数
		InstanceIntegerParam []*InstanceIntegerParam `json:"InstanceIntegerParam,omitempty" name:"InstanceIntegerParam" list`

		// 实例字符型参数
		InstanceTextParam []*InstanceTextParam `json:"InstanceTextParam,omitempty" name:"InstanceTextParam" list`

		// 实例多选项型参数
		InstanceMultiParam []*InstanceMultiParam `json:"InstanceMultiParam,omitempty" name:"InstanceMultiParam" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 实例列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DescribeInstanceSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例安全组信息
		InstanceSecurityGroupsDetail []*InstanceSecurityGroupDetail `json:"InstanceSecurityGroupsDetail,omitempty" name:"InstanceSecurityGroupsDetail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 是否过滤掉从节信息
	FilterSlave *bool `json:"FilterSlave,omitempty" name:"FilterSlave"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例分片列表信息
		InstanceShards []*InstanceClusterShard `json:"InstanceShards,omitempty" name:"InstanceShards" list`

		// 实例分片节点总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例列表的大小，参数默认值20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 实例Id，如：crs-6ubhgouj
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 枚举范围： projectId,createtime,instancename,type,curDeadline
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1倒序，0顺序，默认倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 私有网络ID数组，数组下标从0开始，如果不传则默认选择基础网络，如：47525
	VpcIds []*string `json:"VpcIds,omitempty" name:"VpcIds" list`

	// 子网ID数组，数组下标从0开始，如：56854
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 项目ID 组成的数组，数组下标从0开始
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 查找实例的ID。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 私有网络ID数组，数组下标从0开始，如果不传则默认选择基础网络，如：vpc-sad23jfdfk
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds" list`

	// 子网ID数组，数组下标从0开始，如：subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds" list`

	// 地域ID，已经弃用，可通过公共参数Region查询对应地域
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds" list`

	// 实例状态：0-待初始化，1-流程中，2-运行中，-2-已隔离，-3-待删除
	Status []*int64 `json:"Status,omitempty" name:"Status" list`

	// 类型版本：1-单机版,2-主从版,3-集群版
	TypeVersion *int64 `json:"TypeVersion,omitempty" name:"TypeVersion"`

	// 引擎信息：Redis-2.8，Redis-4.0，CKV
	EngineName *string `json:"EngineName,omitempty" name:"EngineName"`

	// 续费模式：0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenew []*int64 `json:"AutoRenew,omitempty" name:"AutoRenew" list`

	// 计费模式：postpaid-按量计费；prepaid-包年包月
	BillingMode *string `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例类型：1-Redis老集群版；2-Redis 2.8主从版；3-CKV主从版；4-CKV集群版；5-Redis 2.8单机版；6-Redis 4.0主从版；7-Redis 4.0集群版；8 – Redis5.0主从版，9 – Redis5.0集群版，
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 搜索关键词：支持实例Id、实例名称、完整IP
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys" list`

	// 内部参数，用户可忽略
	TypeList []*int64 `json:"TypeList,omitempty" name:"TypeList" list`

	// 内部参数，用户可忽略
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`
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

		// 实例数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 实例详细信息列表
		InstanceSet []*InstanceSet `json:"InstanceSet,omitempty" name:"InstanceSet" list`

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

type DescribeMaintenanceWindowRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMaintenanceWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 维护时间窗起始时间，如：17:00
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 维护时间窗结束时间，如：19:00
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMaintenanceWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProductInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProductInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域售卖信息
		RegionSet []*RegionConf `json:"RegionSet,omitempty" name:"RegionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProductInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupRequest struct {
	*tchttp.BaseRequest

	// 0:默认项目；-1 所有项目; >0: 特定项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 安全组Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

func (r *DescribeProjectSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目安全组
		SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitempty" name:"SecurityGroupDetails" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb
	Product *string `json:"Product,omitempty" name:"Product"`

	// 项目Id。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取数量限制。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索条件，支持安全组id或者安全组名称。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组规则。
		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups" list`

		// 符合条件的安全组总数量。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProxySlowLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢查询阈值（单位：毫秒）
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// 页面大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeProxySlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProxySlowLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProxySlowLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 慢查询总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 慢查询详情
		InstanceProxySlowLogDetail []*InstanceProxySlowlogDetail `json:"InstanceProxySlowLogDetail,omitempty" name:"InstanceProxySlowLogDetail" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProxySlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProxySlowLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 慢查询阈值（单位：微秒）
	MinQueryTime *int64 `json:"MinQueryTime,omitempty" name:"MinQueryTime"`

	// 页面大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
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

		// 慢查询总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 慢查询详情
		InstanceSlowlogDetail []*InstanceSlowlogDetail `json:"InstanceSlowlogDetail,omitempty" name:"InstanceSlowlogDetail" list`

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

type DescribeTaskInfoRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务状态preparing:待执行，running：执行中，succeed：成功，failed：失败，error 执行出错
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 任务类型
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 实例的ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 任务信息，错误时显示错误信息。执行中与成功则为空
		TaskMessage *string `json:"TaskMessage,omitempty" name:"TaskMessage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskListRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 分页大小
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍（自动向下取整）
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 项目Id
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 任务类型
	TaskTypes []*string `json:"TaskTypes,omitempty" name:"TaskTypes" list`

	// 起始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 终止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态
	TaskStatus []*int64 `json:"TaskStatus,omitempty" name:"TaskStatus" list`
}

func (r *DescribeTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务详细信息
		Tasks []*TaskInfoDetail `json:"Tasks,omitempty" name:"Tasks" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPostpaidInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyPostpaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPostpaidInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPostpaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务Id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyPostpaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPostpaidInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPrepaidInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DestroyPrepaidInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPrepaidInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DestroyPrepaidInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单Id
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DestroyPrepaidInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DestroyPrepaidInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest

	// 实例序号ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DisableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableReplicaReadonlyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 失败:ERROR，成功:OK
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableReplicaReadonlyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb 等。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 安全组Id。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例ID列表，一个或者多个实例Id组成的数组。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableReplicaReadonlyRequest struct {
	*tchttp.BaseRequest

	// 实例序号ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账号路由策略：填写master或者replication，表示路由主节点，从节点；不填路由策略默认为写主节点，读从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`
}

func (r *EnableReplicaReadonlyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableReplicaReadonlyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableReplicaReadonlyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 错误：ERROR，正确OK。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableReplicaReadonlyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableReplicaReadonlyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type HotKeyInfo struct {

	// 热Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数量
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type Inbound struct {

	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type InquiryPriceCreateInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例所属的可用区id
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例类型：2 – Redis2.8主从版，3 – Redis3.2主从版(CKV主从版)，4 – Redis3.2集群版(CKV集群版)，5-Redis2.8单机版，6 – Redis4.0主从版，7 – Redis4.0集群版，
	TypeId *uint64 `json:"TypeId,omitempty" name:"TypeId"`

	// 实例容量，单位MB， 取值大小以 查询售卖规格接口返回的规格为准
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 实例数量，单次购买实例数量以 查询售卖规格接口返回的规格为准
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 购买时长，在创建包年包月实例的时候需要填写，按量计费实例填1即可，单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 付费方式:0-按量计费，1-包年包月。
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例分片数量，Redis2.8主从版、CKV主从版和Redis2.8单机版、Redis4.0主从版不需要填写
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// 实例副本数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// 是否支持副本只读，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	ReplicasReadonly *bool `json:"ReplicasReadonly,omitempty" name:"ReplicasReadonly"`
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

		// 价格，单位：分
	// 注意：此字段可能返回 null，表示取不到有效值。
		Price *float64 `json:"Price,omitempty" name:"Price"`

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

	// 购买时长，单位：月
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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

		// 价格，单位：分
	// 注意：此字段可能返回 null，表示取不到有效值。
		Price *float64 `json:"Price,omitempty" name:"Price"`

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

type InquiryPriceUpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片大小 单位 MB
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 分片数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// 副本数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`
}

func (r *InquiryPriceUpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquiryPriceUpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 价格，单位：分
	// 注意：此字段可能返回 null，表示取不到有效值。
		Price *float64 `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquiryPriceUpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquiryPriceUpgradeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceClusterNode struct {

	// 节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例运行时节点Id
	RunId *string `json:"RunId,omitempty" name:"RunId"`

	// 集群角色：0-master；1-slave
	Role *int64 `json:"Role,omitempty" name:"Role"`

	// 节点状态：0-readwrite, 1-read, 2-backup
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 服务状态：0-down；1-on
	Connected *int64 `json:"Connected,omitempty" name:"Connected"`

	// 节点创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 节点下线时间
	DownTime *string `json:"DownTime,omitempty" name:"DownTime"`

	// 节点slot分布
	Slots *string `json:"Slots,omitempty" name:"Slots"`

	// 节点key分布
	Keys *int64 `json:"Keys,omitempty" name:"Keys"`

	// 节点qps
	Qps *int64 `json:"Qps,omitempty" name:"Qps"`

	// 节点qps倾斜度
	QpsSlope *float64 `json:"QpsSlope,omitempty" name:"QpsSlope"`

	// 节点存储
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 节点存储倾斜度
	StorageSlope *float64 `json:"StorageSlope,omitempty" name:"StorageSlope"`
}

type InstanceClusterShard struct {

	// 分片节点名称
	ShardName *string `json:"ShardName,omitempty" name:"ShardName"`

	// 分片节点Id
	ShardId *string `json:"ShardId,omitempty" name:"ShardId"`

	// 角色
	Role *int64 `json:"Role,omitempty" name:"Role"`

	// Key数量
	Keys *int64 `json:"Keys,omitempty" name:"Keys"`

	// slot信息
	Slots *string `json:"Slots,omitempty" name:"Slots"`

	// 使用容量
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 容量倾斜率
	StorageSlope *float64 `json:"StorageSlope,omitempty" name:"StorageSlope"`

	// 实例运行时节点Id
	Runid *string `json:"Runid,omitempty" name:"Runid"`

	// 服务状态：0-down；1-on
	Connected *int64 `json:"Connected,omitempty" name:"Connected"`
}

type InstanceEnumParam struct {

	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：enum
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数可取值
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue" list`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceIntegerParam struct {

	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：integer
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数最小值
	Min *string `json:"Min,omitempty" name:"Min"`

	// 参数最大值
	Max *string `json:"Max,omitempty" name:"Max"`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceMultiParam struct {

	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：multi
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数说明
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue" list`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type InstanceNode struct {

	// Id
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 节点详细信息
	InstanceClusterNode []*InstanceClusterNode `json:"InstanceClusterNode,omitempty" name:"InstanceClusterNode" list`
}

type InstanceParam struct {

	// 设置参数的名字
	Key *string `json:"Key,omitempty" name:"Key"`

	// 设置参数的值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InstanceParamHistory struct {

	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 修改前值
	PreValue *string `json:"PreValue,omitempty" name:"PreValue"`

	// 修改后值
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// 状态：1-参数配置修改中；2-参数配置修改成功；3-参数配置修改失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type InstanceProxySlowlogDetail struct {

	// 慢查询耗时
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 客户端地址
	Client *string `json:"Client,omitempty" name:"Client"`

	// 命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 详细命令行信息
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// 执行时间
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`
}

type InstanceSecurityGroupDetail struct {

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 安全组信息
	SecurityGroupDetails []*SecurityGroupDetail `json:"SecurityGroupDetails,omitempty" name:"SecurityGroupDetails" list`
}

type InstanceSet struct {

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户的Appid
	Appid *int64 `json:"Appid,omitempty" name:"Appid"`

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域id 1--广州 4--上海 5-- 中国香港 6--多伦多 7--上海金融 8--北京 9-- 新加坡 11--深圳金融 15--美西（硅谷）16--成都 17--德国 18--韩国 19--重庆 21--印度 22--美东（弗吉尼亚）23--泰国 24--俄罗斯 25--日本
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 区域id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// vpc网络id 如：75101
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// vpc网络下子网id 如：46315
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例当前状态，0：待初始化；1：实例在流程中；2：实例运行中；-2：实例已隔离；-3：实例待删除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例vip
	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`

	// 实例端口号
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 实例创建时间
	Createtime *string `json:"Createtime,omitempty" name:"Createtime"`

	// 实例容量大小，单位：MB
	Size *float64 `json:"Size,omitempty" name:"Size"`

	// 该字段已废弃
	SizeUsed *float64 `json:"SizeUsed,omitempty" name:"SizeUsed"`

	// 实例类型：1 – Redis2.8内存版（集群架构），2 – Redis2.8内存版（标准架构），3 – CKV 3.2内存版(标准架构)，4 – CKV 3.2内存版(集群架构)，5 – Redis2.8内存版（单机），6 – Redis4.0内存版（标准架构），7 – Redis4.0内存版（集群架构），8 – Redis5.0内存版（标准架构），9 – Redis5.0内存版（集群架构）
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 实例是否设置自动续费标识，1：设置自动续费；0：未设置自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 实例到期时间
	DeadlineTime *string `json:"DeadlineTime,omitempty" name:"DeadlineTime"`

	// 引擎：社区版Redis、腾讯云CKV
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 产品类型：standalone – 标准版，cluster – 集群版
	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`

	// vpc网络id 如：vpc-fk33jsf43kgv
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// vpc网络下子网id 如：subnet-fd3j6l35mm0
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 计费模式：0-按量计费，1-包年包月
	BillingMode *int64 `json:"BillingMode,omitempty" name:"BillingMode"`

	// 实例运行状态描述：如”实例运行中“
	InstanceTitle *string `json:"InstanceTitle,omitempty" name:"InstanceTitle"`

	// 计划下线时间
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 流程中的实例，返回子状态
	SubStatus *int64 `json:"SubStatus,omitempty" name:"SubStatus"`

	// 反亲和性标签
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 实例节点信息
	InstanceNode []*InstanceNode `json:"InstanceNode,omitempty" name:"InstanceNode" list`

	// 分片大小
	RedisShardSize *int64 `json:"RedisShardSize,omitempty" name:"RedisShardSize"`

	// 分片数量
	RedisShardNum *int64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// 副本数量
	RedisReplicasNum *int64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`

	// 计费Id
	PriceId *int64 `json:"PriceId,omitempty" name:"PriceId"`

	// 隔离时间
	CloseTime *string `json:"CloseTime,omitempty" name:"CloseTime"`

	// 从节点读取权重
	SlaveReadWeight *int64 `json:"SlaveReadWeight,omitempty" name:"SlaveReadWeight"`

	// 实例关联的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTags []*InstanceTagInfo `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 是否为免密实例，true-免密实例；false-非免密实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`

	// 客户端连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientLimit *int64 `json:"ClientLimit,omitempty" name:"ClientLimit"`

	// DTS状态（内部参数，用户可忽略）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DtsStatus *int64 `json:"DtsStatus,omitempty" name:"DtsStatus"`

	// 分片带宽上限，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetLimit *int64 `json:"NetLimit,omitempty" name:"NetLimit"`

	// 免密实例标识（内部参数，用户可忽略）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PasswordFree *int64 `json:"PasswordFree,omitempty" name:"PasswordFree"`

	// 实例只读标识（内部参数，用户可忽略）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadOnly *int64 `json:"ReadOnly,omitempty" name:"ReadOnly"`

	// 内部参数，用户可忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip6 *string `json:"Vip6,omitempty" name:"Vip6"`

	// 内部参数，用户可忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainBandwidthDuration *string `json:"RemainBandwidthDuration,omitempty" name:"RemainBandwidthDuration"`

	// Tendis实例的磁盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 监控版本: 1m-分钟粒度监控，5s-5秒粒度监控
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorVersion *string `json:"MonitorVersion,omitempty" name:"MonitorVersion"`
}

type InstanceSlowlogDetail struct {

	// 慢查询耗时
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 客户端地址
	Client *string `json:"Client,omitempty" name:"Client"`

	// 命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 详细命令行信息
	CommandLine *string `json:"CommandLine,omitempty" name:"CommandLine"`

	// 执行时间
	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`
}

type InstanceTagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type InstanceTextParam struct {

	// 参数名
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数类型：text
	ValueType *string `json:"ValueType,omitempty" name:"ValueType"`

	// 修改后是否需要重启：true，false
	NeedRestart *string `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 当前运行参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 参数说明
	Tips *string `json:"Tips,omitempty" name:"Tips"`

	// 参数可取值
	TextValue []*string `json:"TextValue,omitempty" name:"TextValue" list`

	// 参数状态, 1: 修改中， 2：修改完成
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ManualBackupInstanceRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID，可通过 DescribeInstance接口返回值中的 InstanceId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份的备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ManualBackupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManualBackupInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ManualBackupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ManualBackupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ManualBackupInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModfiyInstancePasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 实例新密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ModfiyInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModfiyInstancePasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModfiyInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModfiyInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModfiyInstancePasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 日期 Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday
	WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays" list`

	// 时间段 00:00-01:00, 01:00-02:00...... 23:00-00:00
	TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

	// 自动备份类型： 1 “定时回档”
	AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`
}

func (r *ModifyAutoBackupConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoBackupConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoBackupConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自动备份类型： 1 “定时回档”
		AutoBackupType *int64 `json:"AutoBackupType,omitempty" name:"AutoBackupType"`

		// 日期Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
		WeekDays []*string `json:"WeekDays,omitempty" name:"WeekDays" list`

		// 时间段 00:00-01:00, 01:00-02:00...... 23:00-00:00
		TimePeriod *string `json:"TimePeriod,omitempty" name:"TimePeriod"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoBackupConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoBackupConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：mariadb,cdb,cynosdb,dcdb,redis,mongodb 等。
	Product *string `json:"Product,omitempty" name:"Product"`

	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子账号名称，如果要修改主账号，填root
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 子账号密码
	AccountPassword *string `json:"AccountPassword,omitempty" name:"AccountPassword"`

	// 子账号描述信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 子账号路由策略：填写master或者slave，表示路由主节点，从节点
	ReadonlyPolicy []*string `json:"ReadonlyPolicy,omitempty" name:"ReadonlyPolicy" list`

	// 子账号读写策略：填写r、w、rw，表示只读，只写，读写策略
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// true表示将主账号切换为免密账号，这里只适用于主账号，子账号不可免密
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

func (r *ModifyInstanceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAccountRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamsRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例修改的参数列表
	InstanceParams []*InstanceParam `json:"InstanceParams,omitempty" name:"InstanceParams" list`
}

func (r *ModifyInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改是否成功。
		Changed *bool `json:"Changed,omitempty" name:"Changed"`

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceParamsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest

	// 修改实例操作，如填写：rename-表示实例重命名；modifyProject-修改实例所属项目；modifyAutoRenew-修改实例续费标记
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 实例Id
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例的新名称
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 自动续费标识。0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenews []*int64 `json:"AutoRenews,omitempty" name:"AutoRenews" list`

	// 已经废弃
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 已经废弃
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 已经废弃
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMaintenanceWindowRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 维护时间窗起始时间，如：17:00
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 维护时间窗结束时间，如：19:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *ModifyMaintenanceWindowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMaintenanceWindowRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMaintenanceWindowResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改状态：success 或者 failed
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMaintenanceWindowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMaintenanceWindowResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 操作类型：changeVip——修改实例VIP；changeVpc——修改实例子网；changeBaseToVpc——基础网络转VPC网络
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// VIP地址，changeVip的时候填写，不填则默认分配
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 私有网络ID，changeVpc、changeBaseToVpc的时候需要提供
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID，changeVpc、changeBaseToVpc的时候需要提供
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *ModifyNetworkConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNetworkConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 执行状态：true|false
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 子网ID
		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

		// 私有网络ID
		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

		// VIP地址
		Vip *string `json:"Vip,omitempty" name:"Vip"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNetworkConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNetworkConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// 策略，ACCEPT或者DROP。
	Action *string `json:"Action,omitempty" name:"Action"`

	// 地址组id代表的地址集合。
	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`

	// 来源Ip或Ip段，例如192.168.0.0/16。
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 描述。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 网络协议，支持udp、tcp等。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 端口。
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 服务组id代表的协议和端口集合。
	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`

	// 安全组id代表的地址集合。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type ProductConf struct {

	// 产品类型，2-Redis主从版，3-CKV主从版，4-CKV集群版，5-Redis单机版，7-Redis集群版
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 产品名称，Redis主从版，CKV主从版，CKV集群版，Redis单机版，Redis集群版
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 购买时的最小数量
	MinBuyNum *int64 `json:"MinBuyNum,omitempty" name:"MinBuyNum"`

	// 购买时的最大数量
	MaxBuyNum *int64 `json:"MaxBuyNum,omitempty" name:"MaxBuyNum"`

	// 产品是否售罄
	Saleout *bool `json:"Saleout,omitempty" name:"Saleout"`

	// 产品引擎，腾讯云CKV或者社区版Redis
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 兼容版本，Redis-2.8，Redis-3.2，Redis-4.0
	Version *string `json:"Version,omitempty" name:"Version"`

	// 规格总大小，单位G
	TotalSize []*string `json:"TotalSize,omitempty" name:"TotalSize" list`

	// 每个分片大小，单位G
	ShardSize []*string `json:"ShardSize,omitempty" name:"ShardSize" list`

	// 副本数量
	ReplicaNum []*string `json:"ReplicaNum,omitempty" name:"ReplicaNum" list`

	// 分片数量
	ShardNum []*string `json:"ShardNum,omitempty" name:"ShardNum" list`

	// 支持的计费模式，1-包年包月，0-按量计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 是否支持副本只读
	EnableRepicaReadOnly *bool `json:"EnableRepicaReadOnly,omitempty" name:"EnableRepicaReadOnly"`
}

type RedisBackupSet struct {

	// 开始备份的时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份ID
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// 备份类型。 manualBackupInstance：用户发起的手动备份； systemBackupInstance：凌晨系统发起的备份
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份状态。  1:"备份被其它流程锁定";  2:"备份正常，没有被任何流程锁定";  -1:"备份已过期"； 3:"备份正在被导出";  4:"备份导出成功"
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 备份的备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 备份是否被锁定，0：未被锁定；1：已被锁定
	Locked *int64 `json:"Locked,omitempty" name:"Locked"`
}

type RegionConf struct {

	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域简称
	RegionShortName *string `json:"RegionShortName,omitempty" name:"RegionShortName"`

	// 地域所在大区名称
	Area *string `json:"Area,omitempty" name:"Area"`

	// 可用区信息
	ZoneSet []*ZoneCapacityConf `json:"ZoneSet,omitempty" name:"ZoneSet" list`
}

type RenewInstanceRequest struct {
	*tchttp.BaseRequest

	// 购买时长，单位：月
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 交易ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest

	// Redis实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 重置的密码（切换为免密实例时，可不传；其他情况必传）
	Password *string `json:"Password,omitempty" name:"Password"`

	// 是否切换免密实例，false-切换为非免密码实例，true-切换为免密码实例；默认false
	NoAuth *bool `json:"NoAuth,omitempty" name:"NoAuth"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID（修改密码时的任务ID，如果时切换免密码或者非免密码实例，则无需关注此返回值）
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetPasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest

	// 待操作的实例ID，可通过 DescribeRedis 接口返回值中的 redisId 获取。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份ID，可通过 GetRedisBackupList 接口返回值中的 backupId 获取
	BackupId *string `json:"BackupId,omitempty" name:"BackupId"`

	// 实例密码，恢复实例时，需要校验实例密码（免密实例不需要传密码）
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestoreInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID，可通过 DescribeTaskInfo 接口查询任务执行状态
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RestoreInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目ID。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称。
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注。
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// 出站规则。
	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound" list`

	// 入站规则。
	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound" list`
}

type SecurityGroupDetail struct {

	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 安全组Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组标记
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`

	// 安全组入站规则
	InboundRule []*SecurityGroupsInboundAndOutbound `json:"InboundRule,omitempty" name:"InboundRule" list`

	// 安全组出站规则
	OutboundRule []*SecurityGroupsInboundAndOutbound `json:"OutboundRule,omitempty" name:"OutboundRule" list`
}

type SecurityGroupsInboundAndOutbound struct {

	// 执行动作
	Action *string `json:"Action,omitempty" name:"Action"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 端口号
	Port *string `json:"Port,omitempty" name:"Port"`

	// 协议类型
	Proto *string `json:"Proto,omitempty" name:"Proto"`
}

type SourceCommand struct {

	// 命令
	Cmd *string `json:"Cmd,omitempty" name:"Cmd"`

	// 执行次数
	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type SourceInfo struct {

	// 来源IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 连接数
	Conn *int64 `json:"Conn,omitempty" name:"Conn"`

	// 命令
	Cmd *int64 `json:"Cmd,omitempty" name:"Cmd"`
}

type StartupInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *StartupInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartupInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartupInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务id
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartupInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartupInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchInstanceVipRequest struct {
	*tchttp.BaseRequest

	// 源实例ID
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// 目标实例ID
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// 单位为秒。源实例与目标实例间DTS已断开时间，如果DTS断开时间大于TimeDelay，则不切换VIP，建议尽量根据业务设置一个可接受的值。
	TimeDelay *int64 `json:"TimeDelay,omitempty" name:"TimeDelay"`

	// 在DTS断开的情况下是否强制切换。1：强制切换，0：不强制切换
	ForceSwitch *int64 `json:"ForceSwitch,omitempty" name:"ForceSwitch"`

	// now: 立即切换，syncComplete：等待同步完成后切换
	SwitchTime *string `json:"SwitchTime,omitempty" name:"SwitchTime"`
}

func (r *SwitchInstanceVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchInstanceVipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwitchInstanceVipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchInstanceVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwitchInstanceVipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TaskInfoDetail struct {

	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *float64 `json:"Progress,omitempty" name:"Progress"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *int64 `json:"Result,omitempty" name:"Result"`
}

type TradeDealDetail struct {

	// 订单号ID，调用云API时使用此ID
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 长订单ID，反馈订单问题给官方客服使用此ID
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 可用区id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 订单关联的实例数
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 创建用户uin
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 订单创建时间
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 订单超时时间
	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`

	// 订单完成时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 订单状态 1：未支付 2:已支付，未发货 3:发货中 4:发货成功 5:发货失败 6:已退款 7:已关闭订单 8:订单过期 9:订单已失效 10:产品已失效 11:代付拒绝 12:支付中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 订单状态描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 订单实际总价，单位：分
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 实例ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分片大小 单位 MB
	MemSize *uint64 `json:"MemSize,omitempty" name:"MemSize"`

	// 分片数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisShardNum *uint64 `json:"RedisShardNum,omitempty" name:"RedisShardNum"`

	// 副本数量，Redis2.8主从版、CKV主从版和Redis2.8单机版不需要填写
	RedisReplicasNum *uint64 `json:"RedisReplicasNum,omitempty" name:"RedisReplicasNum"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceVersionRequest struct {
	*tchttp.BaseRequest

	// 目标实例类型，同 [CreateInstances](https://cloud.tencent.com/document/api/239/20026) 的Type，即实例要变更的目标类型
	TargetInstanceType *string `json:"TargetInstanceType,omitempty" name:"TargetInstanceType"`

	// 切换模式：1-维护时间窗切换，2-立即切换
	SwitchOption *int64 `json:"SwitchOption,omitempty" name:"SwitchOption"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UpgradeInstanceVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeInstanceVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeInstanceVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeInstanceVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ZoneCapacityConf struct {

	// 可用区ID：如ap-guangzhou-3
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区是否售罄
	IsSaleout *bool `json:"IsSaleout,omitempty" name:"IsSaleout"`

	// 是否为默认可用区
	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`

	// 网络类型：basenet -- 基础网络；vpcnet -- VPC网络
	NetWorkType []*string `json:"NetWorkType,omitempty" name:"NetWorkType" list`

	// 可用区内产品规格等信息
	ProductSet []*ProductConf `json:"ProductSet,omitempty" name:"ProductSet" list`

	// 可用区ID：如100003
	OldZoneId *int64 `json:"OldZoneId,omitempty" name:"OldZoneId"`
}
