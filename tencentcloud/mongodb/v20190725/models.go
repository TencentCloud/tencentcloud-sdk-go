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

package v20190725

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

type BackupFile struct {

	// 备份文件所属的副本集/分片ID
	ReplicateSetId *string `json:"ReplicateSetId,omitempty" name:"ReplicateSetId"`

	// 备份文件保存路径
	File *string `json:"File,omitempty" name:"File"`
}

type BackupInfo struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份方式，0-自动备份，1-手动备份
	BackupType *uint64 `json:"BackupType,omitempty" name:"BackupType"`

	// 备份名称
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 备份备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupDesc *string `json:"BackupDesc,omitempty" name:"BackupDesc"`

	// 备份文件大小，单位KB
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupSize *uint64 `json:"BackupSize,omitempty" name:"BackupSize"`

	// 备份开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备份状态，1-备份中，2-备份成功
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 备份方法，0-逻辑备份，1-物理备份
	BackupMethod *uint64 `json:"BackupMethod,omitempty" name:"BackupMethod"`
}

type ClientConnection struct {

	// 连接的客户端IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 对应客户端IP的连接数
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type CreateBackupDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 0-逻辑备份，1-物理备份
	BackupMethod *int64 `json:"BackupMethod,omitempty" name:"BackupMethod"`

	// 备份备注
	BackupRemark *string `json:"BackupRemark,omitempty" name:"BackupRemark"`
}

func (r *CreateBackupDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateBackupDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询备份流程的状态
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBackupDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDBInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 副本集个数，创建副本集实例时，该参数必须设置为1；创建分片实例时，具体参照查询云数据库的售卖规格返回参数
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// 每个副本集内节点个数，当前副本集节点数固定为3，分片从节点数可选，具体参照查询云数据库的售卖规格返回参数
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 版本号，具体支持的售卖版本请参照查询云数据库的售卖规格（DescribeSpecInfo）返回结果。参数与版本对应关系是MONGO_3_WT：MongoDB 3.2 WiredTiger存储引擎版本，MONGO_3_ROCKS：MongoDB 3.2 RocksDB存储引擎版本，MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// 机器类型，HIO：高IO型；HIO10G：高IO万兆
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// 实例数量，最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 可用区信息，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 私有网络ID，如果不设置该参数则默认选择基础网络
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，如果设置了 VpcId，则 SubnetId必填
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例密码，不设置该参数则需要在创建完成后通过设置密码接口初始化实例密码。密码必须是8-16位字符，且至少包含字母、数字和字符 !@#%^*() 中的两种
	Password *string `json:"Password,omitempty" name:"Password"`

	// 项目ID，不设置为默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例标签信息
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 1:正式实例,2:临时实例,3:只读实例，4：灾备实例
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// 父实例Id，当Clone为3或者4时，这个必须填
	Father *string `json:"Father,omitempty" name:"Father"`

	// 安全组
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

	// 每个副本集内节点个数，当前副本集节点数固定为3，分片从节点数可选，具体参照查询云数据库的售卖规格返回参数
	NodeNum *uint64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 实例内存大小，单位：GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 版本号，具体支持的售卖版本请参照查询云数据库的售卖规格（DescribeSpecInfo）返回结果。参数与版本对应关系是MONGO_3_WT：MongoDB 3.2 WiredTiger存储引擎版本，MONGO_3_ROCKS：MongoDB 3.2 RocksDB存储引擎版本，MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// 实例数量, 最小值1，最大值为10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例所属区域名称，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例时长，单位：月，可选值包括 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 机器类型，HIO：高IO型；HIO10G：高IO万兆型；STDS5：标准型
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群，STANDALONE-单节点
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 副本集个数，创建副本集实例时，该参数必须设置为1；创建分片实例时，具体参照查询云数据库的售卖规格返回参数；若为单节点实例，该参数设置为0
	ReplicateSetNum *uint64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`

	// 项目ID，不设置为默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络 ID，如果不传则默认选择基础网络，请使用 查询私有网络列表
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络下的子网 ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用 查询子网列表
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例密码，不设置该参数则需要在创建完成后通过设置密码接口初始化实例密码。密码必须是8-16位字符，且至少包含字母、数字和字符 !@#%^*() 中的两种
	Password *string `json:"Password,omitempty" name:"Password"`

	// 实例标签信息
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 自动续费标记，可选值为：0 - 不自动续费；1 - 自动续费。默认为不自动续费
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 是否自动选择代金券，可选值为：1 - 是；0 - 否； 默认为0
	AutoVoucher *uint64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 1:正式实例,2:临时实例,3:只读实例，4：灾备实例
	Clone *int64 `json:"Clone,omitempty" name:"Clone"`

	// 若是只读，灾备实例，Father必须填写，即主实例ID
	Father *string `json:"Father,omitempty" name:"Father"`

	// 安全组
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup" list`
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

type DBInstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DBInstancePrice struct {

	// 单价
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 原价
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 折扣加
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`
}

type DescribeAsyncRequestInfoRequest struct {
	*tchttp.BaseRequest

	// 异步请求Id
	AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`
}

func (r *DescribeAsyncRequestInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAsyncRequestInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态
		Status *string `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAsyncRequestInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAsyncRequestInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupAccessRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要获取下载授权的备份文件名
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *DescribeBackupAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例所属地域
		Region *string `json:"Region,omitempty" name:"Region"`

		// 备份文件所在存储桶
		Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

		// 备份文件的存储信息
		Files []*BackupFile `json:"Files,omitempty" name:"Files" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBackupAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClientConnectionsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询返回记录条数，默认为10000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认值为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClientConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClientConnectionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeClientConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 客户端连接信息，包括客户端IP和对应IP的连接数量。
		Clients []*ClientConnection `json:"Clients,omitempty" name:"Clients" list`

		// 满足条件的记录总条数，可用于分页查询。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClientConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeClientConnectionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBBackupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 备份列表
		BackupList []*BackupInfo `json:"BackupList,omitempty" name:"BackupList" list`

		// 备份总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBBackupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceDealRequest struct {
	*tchttp.BaseRequest

	// 订单ID，通过CreateDBInstance等接口返回
	DealId *string `json:"DealId,omitempty" name:"DealId"`
}

func (r *DescribeDBInstanceDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceDealRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBInstanceDealResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单状态，1：未支付，2：已支付，3：发货中，4：发货成功，5：发货失败，6：退款，7：订单关闭，8：超时未支付关闭。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 订单原价。
		OriginalPrice *float64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

		// 订单折扣价格。
		DiscountPrice *float64 `json:"DiscountPrice,omitempty" name:"DiscountPrice"`

		// 订单行为，purchase：新购，renew：续费，upgrade：升配，downgrade：降配，refund：退货退款。
		Action *string `json:"Action,omitempty" name:"Action"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBInstanceDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBInstanceDealResponse) FromJsonString(s string) error {
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

	// 项目 ID
	ProjectIds []*uint64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 搜索关键词，支持实例Id、实例名称、完整IP
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
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

		// 实例详细信息列表
		InstanceDetails []*InstanceDetail `json:"InstanceDetails,omitempty" name:"InstanceDetails" list`

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

type DescribeSlowLogPatternsRequest struct {
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

func (r *DescribeSlowLogPatternsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogPatternsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogPatternsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 慢日志统计信息总数
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// 慢日志统计信息
		SlowLogPatterns []*SlowLogPattern `json:"SlowLogPatterns,omitempty" name:"SlowLogPatterns" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogPatternsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogPatternsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsRequest struct {
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

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 慢日志总数
		Count *uint64 `json:"Count,omitempty" name:"Count"`

		// 慢日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SlowLogs []*string `json:"SlowLogs,omitempty" name:"SlowLogs" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSpecInfoRequest struct {
	*tchttp.BaseRequest

	// 待查询可用区
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

type FlushInstanceRouterConfigRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *FlushInstanceRouterConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlushInstanceRouterConfigRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FlushInstanceRouterConfigResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlushInstanceRouterConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *FlushInstanceRouterConfigResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例所属区域名称，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 每个副本集内节点个数，当前副本集节点数固定为3，分片从节点数可选，具体参照查询云数据库的售卖规格返回参数
	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`

	// 实例内存大小，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例硬盘大小，单位：GB
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`

	// 版本号，具体支持的售卖版本请参照查询云数据库的售卖规格（DescribeSpecInfo）返回结果。参数与版本对应关系是MONGO_3_WT：MongoDB 3.2 WiredTiger存储引擎版本，MONGO_3_ROCKS：MongoDB 3.2 RocksDB存储引擎版本，MONGO_36_WT：MongoDB 3.6 WiredTiger存储引擎版本，MONGO_40_WT：MongoDB 4.0 WiredTiger存储引擎版本
	MongoVersion *string `json:"MongoVersion,omitempty" name:"MongoVersion"`

	// 机器类型，HIO：高IO型；HIO10G：高IO万兆型；STDS5：标准型
	MachineCode *string `json:"MachineCode,omitempty" name:"MachineCode"`

	// 实例数量, 最小值1，最大值为10
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 实例时长，单位：月，可选值包括[1,2,3,4,5,6,7,8,9,10,11,12,24,36]
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 实例类型，REPLSET-副本集，SHARD-分片集群，STANDALONE-单节点
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 副本集个数，创建副本集实例时，该参数必须设置为1；创建分片实例时，具体参照查询云数据库的售卖规格返回参数；若为单节点实例，该参数设置为0
	ReplicateSetNum *int64 `json:"ReplicateSetNum,omitempty" name:"ReplicateSetNum"`
}

func (r *InquirePriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceCreateDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 价格
		Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceCreateDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 变更配置后实例内存大小，单位：GB。
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 变更配置后实例磁盘大小，单位：GB。
	Volume *int64 `json:"Volume,omitempty" name:"Volume"`
}

func (r *InquirePriceModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceModifyDBInstanceSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 价格。
		Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceModifyDBInstanceSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同，接口单次最多只支持5个实例进行操作。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 预付费模式（即包年包月）相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *InquirePriceRenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceRenewDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 价格
		Price *DBInstancePrice `json:"Price,omitempty" name:"Price"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceRenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InquirePriceRenewDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。默认为1。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceDetail struct {

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
	ReplicaSets []*ShardInfo `json:"ReplicaSets,omitempty" name:"ReplicaSets" list`

	// 只读实例信息
	ReadonlyInstances []*DBInstanceInfo `json:"ReadonlyInstances,omitempty" name:"ReadonlyInstances" list`

	// 灾备实例信息
	StandbyInstances []*DBInstanceInfo `json:"StandbyInstances,omitempty" name:"StandbyInstances" list`

	// 临时实例信息
	CloneInstances []*DBInstanceInfo `json:"CloneInstances,omitempty" name:"CloneInstances" list`

	// 关联实例信息，对于正式实例，该字段表示它的临时实例信息；对于临时实例，则表示它的正式实例信息;如果为只读/灾备实例,则表示他的主实例信息
	RelatedInstance *DBInstanceInfo `json:"RelatedInstance,omitempty" name:"RelatedInstance"`

	// 实例标签信息集合
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 实例版本标记
	InstanceVer *uint64 `json:"InstanceVer,omitempty" name:"InstanceVer"`

	// 实例版本标记
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

type IsolateDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IsolateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IsolateDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例配置变更后的内存大小，单位：GB。内存和磁盘必须同时升配或同时降配
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例配置变更后的硬盘大小，单位：GB。内存和磁盘必须同时升配或同时降配。降配时，新的磁盘参数必须大于已用磁盘容量的1.2倍
	Volume *uint64 `json:"Volume,omitempty" name:"Volume"`

	// 实例配置变更后oplog的大小，单位：GB，默认为磁盘空间的10%，允许设置的最小值为磁盘的10%，最大值为磁盘的90%
	OplogSize *uint64 `json:"OplogSize,omitempty" name:"OplogSize"`
}

func (r *ModifyDBInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSpecRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单ID
		DealId *string `json:"DealId,omitempty" name:"DealId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDBInstanceSpecResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedDBInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，格式如：cmgo-p8vnipr5。与云数据库控制台页面中显示的实例ID相同
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *OfflineIsolatedDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedDBInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type OfflineIsolatedDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OfflineIsolatedDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *OfflineIsolatedDBInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

type RenewDBInstancesRequest struct {
	*tchttp.BaseRequest

	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。包年包月实例该参数为必传参数。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *RenewDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDBInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RenewDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RenewDBInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetDBInstancePasswordRequest struct {
	*tchttp.BaseRequest

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例账号名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 新密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *ResetDBInstancePasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetDBInstancePasswordRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetDBInstancePasswordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 异步请求Id，用户查询该流程的运行状态
		AsyncRequestId *string `json:"AsyncRequestId,omitempty" name:"AsyncRequestId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetDBInstancePasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetDBInstancePasswordResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShardInfo struct {

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

type SlowLogPattern struct {

	// 慢日志模式
	Pattern *string `json:"Pattern,omitempty" name:"Pattern"`

	// 最大执行时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 平均执行时间
	AverageTime *uint64 `json:"AverageTime,omitempty" name:"AverageTime"`

	// 该模式慢日志条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`
}

type SpecItem struct {

	// 规格信息标识
	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`

	// 规格有效标志，取值：0-停止售卖，1-开放售卖
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 规格有效标志，取值：0-停止售卖，1-开放售卖
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

	// 机器类型，取值：0-HIO，4-HIO10G
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type SpecificationInfo struct {

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区信息
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 售卖规格信息
	SpecItems []*SpecItem `json:"SpecItems,omitempty" name:"SpecItems" list`
}

type TagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
