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

package v20180416

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AutoScaleDiskInfo struct {
	// 节点类型 hotData,warmData
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 0:百分比扩容;1:绝对值扩容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleType *uint64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// 触发阈值,单位%,例如80%
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// 触发持续时间,单位分钟,例如60
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 每次扩容比例,单位%,例如20%
	// 注意：此字段可能返回 null，表示取不到有效值。
	PercentSize *uint64 `json:"PercentSize,omitnil,omitempty" name:"PercentSize"`

	// 绝对值扩容,单位GB,例如100GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixSize *uint64 `json:"FixSize,omitnil,omitempty" name:"FixSize"`

	// 扩容上限,单位GB,例如500GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type BackingIndexMetaField struct {
	// 后备索引名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 后备索引状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// 后备索引存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// 后备索引当前生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexPhrase *string `json:"IndexPhrase,omitnil,omitempty" name:"IndexPhrase"`

	// 后备索引创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`
}

// Predefined struct for user
type CheckMigrateIndexMetaDataRequestParams struct {
	// 索引 id
	ServerlessId *string `json:"ServerlessId,omitnil,omitempty" name:"ServerlessId"`

	// 快照名
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// Cos桶名
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// BasePath路径
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 云上集群名
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`

	// 普通索引名列表
	CommonIndexArr []*string `json:"CommonIndexArr,omitnil,omitempty" name:"CommonIndexArr"`

	// 自治索引名列表
	DataStreamArr []*string `json:"DataStreamArr,omitnil,omitempty" name:"DataStreamArr"`
}

type CheckMigrateIndexMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 索引 id
	ServerlessId *string `json:"ServerlessId,omitnil,omitempty" name:"ServerlessId"`

	// 快照名
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// Cos桶名
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// BasePath路径
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 云上集群名
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`

	// 普通索引名列表
	CommonIndexArr []*string `json:"CommonIndexArr,omitnil,omitempty" name:"CommonIndexArr"`

	// 自治索引名列表
	DataStreamArr []*string `json:"DataStreamArr,omitnil,omitempty" name:"DataStreamArr"`
}

func (r *CheckMigrateIndexMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMigrateIndexMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServerlessId")
	delete(f, "Snapshot")
	delete(f, "CosBucket")
	delete(f, "BasePath")
	delete(f, "ClusterInstanceId")
	delete(f, "CommonIndexArr")
	delete(f, "DataStreamArr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckMigrateIndexMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckMigrateIndexMetaDataResponseParams struct {
	// 不存在于目标索引时间字段相同的字段
	MappingTimeFieldCheckFailedIndexArr []*string `json:"MappingTimeFieldCheckFailedIndexArr,omitnil,omitempty" name:"MappingTimeFieldCheckFailedIndexArr"`

	// @timestamp不为date类型，与目标索引时间字段冲突
	MappingTimeTypeCheckFailedIndexArr []*string `json:"MappingTimeTypeCheckFailedIndexArr,omitnil,omitempty" name:"MappingTimeTypeCheckFailedIndexArr"`

	// 索引的创建时间不在 serverless的存储周期内
	SettingCheckFailedIndexArr []*string `json:"SettingCheckFailedIndexArr,omitnil,omitempty" name:"SettingCheckFailedIndexArr"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckMigrateIndexMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *CheckMigrateIndexMetaDataResponseParams `json:"Response"`
}

func (r *CheckMigrateIndexMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckMigrateIndexMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterView struct {
	// 集群健康状态
	Health *float64 `json:"Health,omitnil,omitempty" name:"Health"`

	// 集群是否可见
	Visible *float64 `json:"Visible,omitnil,omitempty" name:"Visible"`

	// 集群是否熔断
	Break *float64 `json:"Break,omitnil,omitempty" name:"Break"`

	// 平均磁盘使用率
	AvgDiskUsage *float64 `json:"AvgDiskUsage,omitnil,omitempty" name:"AvgDiskUsage"`

	// 平均内存使用率
	AvgMemUsage *float64 `json:"AvgMemUsage,omitnil,omitempty" name:"AvgMemUsage"`

	// 平均cpu使用率
	AvgCpuUsage *float64 `json:"AvgCpuUsage,omitnil,omitempty" name:"AvgCpuUsage"`

	// 集群总存储大小
	TotalDiskSize *uint64 `json:"TotalDiskSize,omitnil,omitempty" name:"TotalDiskSize"`

	// 客户端请求节点
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`

	// 在线节点数
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 总节点数
	TotalNodeNum *int64 `json:"TotalNodeNum,omitnil,omitempty" name:"TotalNodeNum"`

	// 数据节点数
	DataNodeNum *int64 `json:"DataNodeNum,omitnil,omitempty" name:"DataNodeNum"`

	// 索引数
	IndexNum *int64 `json:"IndexNum,omitnil,omitempty" name:"IndexNum"`

	// 文档数
	DocNum *int64 `json:"DocNum,omitnil,omitempty" name:"DocNum"`

	// 磁盘已使用字节数
	DiskUsedInBytes *int64 `json:"DiskUsedInBytes,omitnil,omitempty" name:"DiskUsedInBytes"`

	// 分片个数
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 主分片个数
	PrimaryShardNum *int64 `json:"PrimaryShardNum,omitnil,omitempty" name:"PrimaryShardNum"`

	// 迁移中的分片个数
	RelocatingShardNum *int64 `json:"RelocatingShardNum,omitnil,omitempty" name:"RelocatingShardNum"`

	// 初始化中的分片个数
	InitializingShardNum *int64 `json:"InitializingShardNum,omitnil,omitempty" name:"InitializingShardNum"`

	// 未分配的分片个数
	UnassignedShardNum *int64 `json:"UnassignedShardNum,omitnil,omitempty" name:"UnassignedShardNum"`

	// 企业版COS存储容量大小，单位GB
	TotalCosStorage *int64 `json:"TotalCosStorage,omitnil,omitempty" name:"TotalCosStorage"`

	// 企业版集群可搜索快照cos存放的bucket名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchableSnapshotCosBucket *string `json:"SearchableSnapshotCosBucket,omitnil,omitempty" name:"SearchableSnapshotCosBucket"`

	// 企业版集群可搜索快照cos所属appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchableSnapshotCosAppId *string `json:"SearchableSnapshotCosAppId,omitnil,omitempty" name:"SearchableSnapshotCosAppId"`
}

type CommonIndexInfo struct {
	// 普通索引名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 分片状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsShardComplete *int64 `json:"IsShardComplete,omitnil,omitempty" name:"IsShardComplete"`
}

type CosBackup struct {
	// 是否开启cos自动备份
	IsAutoBackup *bool `json:"IsAutoBackup,omitnil,omitempty" name:"IsAutoBackup"`

	// 自动备份执行时间（精确到小时）, e.g. "22:00"
	BackupTime *string `json:"BackupTime,omitnil,omitempty" name:"BackupTime"`

	// 0 腾讯云仓库; 1 客户仓库
	EsRepositoryType *uint64 `json:"EsRepositoryType,omitnil,omitempty" name:"EsRepositoryType"`

	// 客户快照仓库名称
	UserEsRepository *string `json:"UserEsRepository,omitnil,omitempty" name:"UserEsRepository"`

	// 快照存储周期 单位天
	StorageDuration *uint64 `json:"StorageDuration,omitnil,omitempty" name:"StorageDuration"`

	// 自动备份频率单位小时
	AutoBackupInterval *uint64 `json:"AutoBackupInterval,omitnil,omitempty" name:"AutoBackupInterval"`
}

type CosSnapShotInfo struct {
	// cos 桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// base path
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 快照名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 快照版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 普通[{"DataStreamName":"ilm-history-5","Is索引信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonIndexArr []*CommonIndexInfo `json:"CommonIndexArr,omitnil,omitempty" name:"CommonIndexArr"`

	// 自治索引信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataStreamArr []*DataStreamInfo `json:"DataStreamArr,omitnil,omitempty" name:"DataStreamArr"`
}

// Predefined struct for user
type CreateClusterSnapshotRequestParams struct {
	// 实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 索引名称
	Indices *string `json:"Indices,omitnil,omitempty" name:"Indices"`
}

type CreateClusterSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 索引名称
	Indices *string `json:"Indices,omitnil,omitempty" name:"Indices"`
}

func (r *CreateClusterSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SnapshotName")
	delete(f, "Indices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterSnapshotResponseParams struct {
	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterSnapshotResponseParams `json:"Response"`
}

func (r *CreateClusterSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosMigrateToServerlessInstanceRequestParams struct {
	// 快照名
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// 索引 id
	ServerlessId *string `json:"ServerlessId,omitnil,omitempty" name:"ServerlessId"`

	// cos 桶名
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// BasePath 路径
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 云上集群 id
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`

	// 待迁移普通索引名列表
	CommonIndexArr []*string `json:"CommonIndexArr,omitnil,omitempty" name:"CommonIndexArr"`

	// 待迁移自治索引名列表
	DataStreamArr []*string `json:"DataStreamArr,omitnil,omitempty" name:"DataStreamArr"`
}

type CreateCosMigrateToServerlessInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 快照名
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// 索引 id
	ServerlessId *string `json:"ServerlessId,omitnil,omitempty" name:"ServerlessId"`

	// cos 桶名
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// BasePath 路径
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 云上集群 id
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`

	// 待迁移普通索引名列表
	CommonIndexArr []*string `json:"CommonIndexArr,omitnil,omitempty" name:"CommonIndexArr"`

	// 待迁移自治索引名列表
	DataStreamArr []*string `json:"DataStreamArr,omitnil,omitempty" name:"DataStreamArr"`
}

func (r *CreateCosMigrateToServerlessInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosMigrateToServerlessInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Snapshot")
	delete(f, "ServerlessId")
	delete(f, "CosBucket")
	delete(f, "BasePath")
	delete(f, "ClusterInstanceId")
	delete(f, "CommonIndexArr")
	delete(f, "DataStreamArr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCosMigrateToServerlessInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCosMigrateToServerlessInstanceResponseParams struct {
	// 迁移 taskid
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCosMigrateToServerlessInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCosMigrateToServerlessInstanceResponseParams `json:"Response"`
}

func (r *CreateCosMigrateToServerlessInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCosMigrateToServerlessInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexRequestParams struct {
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 创建的索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 创建的索引名
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 【必填】创建的索引元数据JSON，如mappings、settings
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 创建的索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 创建的索引名
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 【必填】创建的索引元数据JSON，如mappings、settings
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "IndexMetaJson")
	delete(f, "Username")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *CreateIndexResponseParams `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例版本（支持"5.6.4"、"6.4.3"、"6.8.2"、"7.5.1"、"7.10.1"）
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 访问密码（密码需8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 已废弃请使用NodeInfoList
	// 节点数量（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 计费类型<li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li>默认值POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月购买时长（单位由参数TimeUnit决定）
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// 自动续费标识<li>RENEW_FLAG_AUTO：自动续费</li><li>RENEW_FLAG_MANUAL：不自动续费，用户手动续费</li>ChargeType为PREPAID时需要设置，如不传递该参数，普通用户默认不自动续费，SVIP用户自动续费
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 已废弃请使用NodeInfoList
	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 已废弃请使用NodeInfoList
	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高性能云硬盘</li><li> CLOUD_HSSD：增强型SSD云硬盘</li><li> CLOUD_BSSD：通用型SSD云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 已废弃请使用NodeInfoList
	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 计费时长单位（ChargeType为PREPAID时需要设置，默认值为“m”，表示月，当前只支持“m”）
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 已废弃请使用NodeInfoList
	// 是否创建专用主节点<li>true：开启专用主节点</li><li>false：不开启专用主节点</li>默认值false
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// 已废弃请使用NodeInfoList
	// 专用主节点个数（只支持3个和5个，EnableDedicatedMaster为true时该值必传）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// 已废弃请使用NodeInfoList
	// 专用主节点类型（EnableDedicatedMaster为true时必传）<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点磁盘大小（单位GB，非必传，若传递则必须为50，暂不支持自定义）
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// 集群配置文件中的ClusterName（系统默认配置为实例ID，暂不支持自定义）
	ClusterNameInConf *string `json:"ClusterNameInConf,omitnil,omitempty" name:"ClusterNameInConf"`

	// 集群部署方式<li>0：单可用区部署</li><li>1：多可用区部署，北京、上海、上海金融、广州、南京、香港、新加坡、法兰克福（白名单控制）</li>默认为0
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 多可用区部署时可用区的详细信息(DeployMode为1时必传)
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 节点信息列表， 用于描述集群各类节点的规格信息如节点类型，节点个数，节点规格，磁盘类型，磁盘大小等
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// 场景化模板类型 0：不启用 1：通用 2：日志 3：搜索
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// 可视化节点配置
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// 创建https集群，默认是http
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 可维护时间段
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// 是否开启存算分离
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// 是否开启essd 增强型云盘
	DiskEnhance *uint64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// 是否开启智能巡检
	EnableDiagnose *bool `json:"EnableDiagnose,omitnil,omitempty" name:"EnableDiagnose"`

	// cdcId，使用cdc子网时传递
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 置放群组亲和度，范围[0,10]，0表示不开启
	DisasterRecoverGroupAffinity *uint64 `json:"DisasterRecoverGroupAffinity,omitnil,omitempty" name:"DisasterRecoverGroupAffinity"`

	// 子产品ID枚举值： 开源版："sp_es_io2"， 基础版："sp_es_basic"，白金版："sp_es_platinum"，企业版："sp_es_enterprise"，CDC白金版："sp_es_cdc_platinum"，日志增强版："sp_es_enlogging"，tsearch："sp_tsearch_io2"，logstash："sp_es_logstash" ，可以为空，为空的时候后台取LicenseType映射该字段
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 读写分离模式：0-不开启，1-本地读写分离，2-远端读写分离
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// 置放群组是否开启异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组开启异步任务的可维护时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 自动扩盘参数列表
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// 是否开启kibana公网访问，不传默认开启
	EnableKibanaPublicAccess *string `json:"EnableKibanaPublicAccess,omitnil,omitempty" name:"EnableKibanaPublicAccess"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例版本（支持"5.6.4"、"6.4.3"、"6.8.2"、"7.5.1"、"7.10.1"）
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 访问密码（密码需8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 已废弃请使用NodeInfoList
	// 节点数量（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 计费类型<li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li>默认值POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月购买时长（单位由参数TimeUnit决定）
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// 自动续费标识<li>RENEW_FLAG_AUTO：自动续费</li><li>RENEW_FLAG_MANUAL：不自动续费，用户手动续费</li>ChargeType为PREPAID时需要设置，如不传递该参数，普通用户默认不自动续费，SVIP用户自动续费
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 已废弃请使用NodeInfoList
	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 已废弃请使用NodeInfoList
	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高性能云硬盘</li><li> CLOUD_HSSD：增强型SSD云硬盘</li><li> CLOUD_BSSD：通用型SSD云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 已废弃请使用NodeInfoList
	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 计费时长单位（ChargeType为PREPAID时需要设置，默认值为“m”，表示月，当前只支持“m”）
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 已废弃请使用NodeInfoList
	// 是否创建专用主节点<li>true：开启专用主节点</li><li>false：不开启专用主节点</li>默认值false
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// 已废弃请使用NodeInfoList
	// 专用主节点个数（只支持3个和5个，EnableDedicatedMaster为true时该值必传）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// 已废弃请使用NodeInfoList
	// 专用主节点类型（EnableDedicatedMaster为true时必传）<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点磁盘大小（单位GB，非必传，若传递则必须为50，暂不支持自定义）
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// 集群配置文件中的ClusterName（系统默认配置为实例ID，暂不支持自定义）
	ClusterNameInConf *string `json:"ClusterNameInConf,omitnil,omitempty" name:"ClusterNameInConf"`

	// 集群部署方式<li>0：单可用区部署</li><li>1：多可用区部署，北京、上海、上海金融、广州、南京、香港、新加坡、法兰克福（白名单控制）</li>默认为0
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 多可用区部署时可用区的详细信息(DeployMode为1时必传)
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 节点信息列表， 用于描述集群各类节点的规格信息如节点类型，节点个数，节点规格，磁盘类型，磁盘大小等
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// 场景化模板类型 0：不启用 1：通用 2：日志 3：搜索
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// 可视化节点配置
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// 创建https集群，默认是http
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 可维护时间段
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// 是否开启存算分离
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// 是否开启essd 增强型云盘
	DiskEnhance *uint64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// 是否开启智能巡检
	EnableDiagnose *bool `json:"EnableDiagnose,omitnil,omitempty" name:"EnableDiagnose"`

	// cdcId，使用cdc子网时传递
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// 置放群组亲和度，范围[0,10]，0表示不开启
	DisasterRecoverGroupAffinity *uint64 `json:"DisasterRecoverGroupAffinity,omitnil,omitempty" name:"DisasterRecoverGroupAffinity"`

	// 子产品ID枚举值： 开源版："sp_es_io2"， 基础版："sp_es_basic"，白金版："sp_es_platinum"，企业版："sp_es_enterprise"，CDC白金版："sp_es_cdc_platinum"，日志增强版："sp_es_enlogging"，tsearch："sp_tsearch_io2"，logstash："sp_es_logstash" ，可以为空，为空的时候后台取LicenseType映射该字段
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 读写分离模式：0-不开启，1-本地读写分离，2-远端读写分离
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// 置放群组是否开启异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组开启异步任务的可维护时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 自动扩盘参数列表
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// 是否开启kibana公网访问，不传默认开启
	EnableKibanaPublicAccess *string `json:"EnableKibanaPublicAccess,omitnil,omitempty" name:"EnableKibanaPublicAccess"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "EsVersion")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "InstanceName")
	delete(f, "NodeNum")
	delete(f, "ChargeType")
	delete(f, "ChargePeriod")
	delete(f, "RenewFlag")
	delete(f, "NodeType")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "TimeUnit")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "EnableDedicatedMaster")
	delete(f, "MasterNodeNum")
	delete(f, "MasterNodeType")
	delete(f, "MasterNodeDiskSize")
	delete(f, "ClusterNameInConf")
	delete(f, "DeployMode")
	delete(f, "MultiZoneInfo")
	delete(f, "LicenseType")
	delete(f, "NodeInfoList")
	delete(f, "TagList")
	delete(f, "BasicSecurityType")
	delete(f, "SceneType")
	delete(f, "WebNodeTypeInfo")
	delete(f, "Protocol")
	delete(f, "OperationDuration")
	delete(f, "EnableHybridStorage")
	delete(f, "DiskEnhance")
	delete(f, "EnableDiagnose")
	delete(f, "CdcId")
	delete(f, "DisasterRecoverGroupAffinity")
	delete(f, "SubProductCode")
	delete(f, "ReadWriteMode")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	delete(f, "AutoScaleDiskInfoList")
	delete(f, "EnableKibanaPublicAccess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogstashInstanceRequestParams struct {
	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例版本（支持"6.8.13"、"7.10.1"）
	LogstashVersion *string `json:"LogstashVersion,omitnil,omitempty" name:"LogstashVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 节点数量（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 计费类型<li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li>默认值POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月购买时长（单位由参数TimeUnit决定）
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// 计费时长单位（ChargeType为PREPAID时需要设置，默认值为“m”，表示月，当前只支持“m”）
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 自动续费标识<li>RENEW_FLAG_AUTO：自动续费</li><li>RENEW_FLAG_MANUAL：不自动续费，用户手动续费</li>ChargeType为PREPAID时需要设置，如不传递该参数，普通用户默认不自动续费，SVIP用户自动续费
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 节点规格<li>LOGSTASH.S1.SMALL2：1核2G</li><li>LOGSTASH.S1.MEDIUM4：2核4G</li><li>LOGSTASH.S1.MEDIUM8：2核8G</li><li>LOGSTASH.S1.LARGE16：4核16G</li><li>LOGSTASH.S1.2XLARGE32：8核32G</li><li>LOGSTASH.S1.4XLARGE32：16核32G</li><li>LOGSTASH.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高硬能云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// License类型<li>oss：开源版</li><li>xpack：xpack版</li>默认值xpack
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 可维护时间段
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// 多可用区部署时可用区的详细信息
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// 部署模式，0：单可用区、1：多可用区
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`
}

type CreateLogstashInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例版本（支持"6.8.13"、"7.10.1"）
	LogstashVersion *string `json:"LogstashVersion,omitnil,omitempty" name:"LogstashVersion"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 节点数量（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 计费类型<li>PREPAID：预付费，即包年包月</li><li>POSTPAID_BY_HOUR：按小时后付费</li>默认值POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月购买时长（单位由参数TimeUnit决定）
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// 计费时长单位（ChargeType为PREPAID时需要设置，默认值为“m”，表示月，当前只支持“m”）
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 自动续费标识<li>RENEW_FLAG_AUTO：自动续费</li><li>RENEW_FLAG_MANUAL：不自动续费，用户手动续费</li>ChargeType为PREPAID时需要设置，如不传递该参数，普通用户默认不自动续费，SVIP用户自动续费
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 节点规格<li>LOGSTASH.S1.SMALL2：1核2G</li><li>LOGSTASH.S1.MEDIUM4：2核4G</li><li>LOGSTASH.S1.MEDIUM8：2核8G</li><li>LOGSTASH.S1.LARGE16：4核16G</li><li>LOGSTASH.S1.2XLARGE32：8核32G</li><li>LOGSTASH.S1.4XLARGE32：16核32G</li><li>LOGSTASH.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高硬能云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// License类型<li>oss：开源版</li><li>xpack：xpack版</li>默认值xpack
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 可维护时间段
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// 多可用区部署时可用区的详细信息
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// 部署模式，0：单可用区、1：多可用区
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`
}

func (r *CreateLogstashInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogstashInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "Zone")
	delete(f, "LogstashVersion")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "NodeNum")
	delete(f, "ChargeType")
	delete(f, "ChargePeriod")
	delete(f, "TimeUnit")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "RenewFlag")
	delete(f, "NodeType")
	delete(f, "DiskType")
	delete(f, "DiskSize")
	delete(f, "LicenseType")
	delete(f, "TagList")
	delete(f, "OperationDuration")
	delete(f, "MultiZoneInfo")
	delete(f, "DeployMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLogstashInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLogstashInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLogstashInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateLogstashInstanceResponseParams `json:"Response"`
}

func (r *CreateLogstashInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLogstashInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessInstanceRequestParams struct {
	// 索引名，需以-AppId结尾
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 创建的索引元数据JSON，如mappings、settings
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// 创建索引的空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 创建索引的用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 创建索引的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 创建数据接入
	ServerlessDi *ServerlessDi `json:"ServerlessDi,omitnil,omitempty" name:"ServerlessDi"`

	// 是否自行添加白名单ip
	AutoGetIp *uint64 `json:"AutoGetIp,omitnil,omitempty" name:"AutoGetIp"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// kibana公网白名单
	KibanaWhiteIpList []*string `json:"KibanaWhiteIpList,omitnil,omitempty" name:"KibanaWhiteIpList"`
}

type CreateServerlessInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 索引名，需以-AppId结尾
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 创建的索引元数据JSON，如mappings、settings
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// 创建索引的空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 创建索引的用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 创建索引的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 创建数据接入
	ServerlessDi *ServerlessDi `json:"ServerlessDi,omitnil,omitempty" name:"ServerlessDi"`

	// 是否自行添加白名单ip
	AutoGetIp *uint64 `json:"AutoGetIp,omitnil,omitempty" name:"AutoGetIp"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// kibana公网白名单
	KibanaWhiteIpList []*string `json:"KibanaWhiteIpList,omitnil,omitempty" name:"KibanaWhiteIpList"`
}

func (r *CreateServerlessInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IndexName")
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IndexMetaJson")
	delete(f, "SpaceId")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "ServerlessDi")
	delete(f, "AutoGetIp")
	delete(f, "TagList")
	delete(f, "KibanaWhiteIpList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerlessInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServerlessInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateServerlessInstanceResponseParams `json:"Response"`
}

func (r *CreateServerlessInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessSpaceV2RequestParams struct {
	// vpc信息
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 索引空间名
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 空间名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 白名单列表
	KibanaWhiteIpList []*string `json:"KibanaWhiteIpList,omitnil,omitempty" name:"KibanaWhiteIpList"`

	// 空间id
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type CreateServerlessSpaceV2Request struct {
	*tchttp.BaseRequest
	
	// vpc信息
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 索引空间名
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 空间名称
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 白名单列表
	KibanaWhiteIpList []*string `json:"KibanaWhiteIpList,omitnil,omitempty" name:"KibanaWhiteIpList"`

	// 空间id
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *CreateServerlessSpaceV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessSpaceV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcInfo")
	delete(f, "SpaceName")
	delete(f, "Zone")
	delete(f, "KibanaWhiteIpList")
	delete(f, "ZoneId")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServerlessSpaceV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServerlessSpaceV2ResponseParams struct {
	// 空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServerlessSpaceV2Response struct {
	*tchttp.BaseResponse
	Response *CreateServerlessSpaceV2ResponseParams `json:"Response"`
}

func (r *CreateServerlessSpaceV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServerlessSpaceV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataStreamInfo struct {
	// 自治索引名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataStreamName *string `json:"DataStreamName,omitnil,omitempty" name:"DataStreamName"`

	// 分片状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsShardComplete *int64 `json:"IsShardComplete,omitnil,omitempty" name:"IsShardComplete"`
}

// Predefined struct for user
type DeleteClusterSnapshotRequestParams struct {
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 集群快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

type DeleteClusterSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 集群快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

func (r *DeleteClusterSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RepositoryName")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterSnapshotResponseParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterSnapshotResponseParams `json:"Response"`
}

func (r *DeleteClusterSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexRequestParams struct {
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 删除的索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 删除的索引名
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 后备索引名
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 删除的索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 删除的索引名
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 后备索引名
	BackingIndexName *string `json:"BackingIndexName,omitnil,omitempty" name:"BackingIndexName"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "BackingIndexName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIndexResponseParams `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogstashInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteLogstashInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteLogstashInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogstashInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogstashInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogstashInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogstashInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogstashInstanceResponseParams `json:"Response"`
}

func (r *DeleteLogstashInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogstashInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogstashPipelinesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 管道ID列表
	PipelineIds []*string `json:"PipelineIds,omitnil,omitempty" name:"PipelineIds"`
}

type DeleteLogstashPipelinesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 管道ID列表
	PipelineIds []*string `json:"PipelineIds,omitnil,omitempty" name:"PipelineIds"`
}

func (r *DeleteLogstashPipelinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogstashPipelinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PipelineIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLogstashPipelinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLogstashPipelinesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLogstashPipelinesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLogstashPipelinesResponseParams `json:"Response"`
}

func (r *DeleteLogstashPipelinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLogstashPipelinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessInstanceRequestParams struct {
	// serverless实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteServerlessInstanceRequest struct {
	*tchttp.BaseRequest
	
	// serverless实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteServerlessInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServerlessInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServerlessInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServerlessInstanceResponseParams `json:"Response"`
}

func (r *DeleteServerlessInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessSpaceUserRequestParams struct {
	// 空间的ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 创建索引的用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

type DeleteServerlessSpaceUserRequest struct {
	*tchttp.BaseRequest
	
	// 空间的ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 创建索引的用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

func (r *DeleteServerlessSpaceUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessSpaceUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Username")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServerlessSpaceUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServerlessSpaceUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServerlessSpaceUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServerlessSpaceUserResponseParams `json:"Response"`
}

func (r *DeleteServerlessSpaceUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServerlessSpaceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSnapshotRequestParams struct {
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 集群快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

type DescribeClusterSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 集群快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`
}

func (r *DescribeClusterSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RepositoryName")
	delete(f, "SnapshotName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterSnapshotResponseParams struct {
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 快照备份详情列表
	Snapshots []*Snapshots `json:"Snapshots,omitnil,omitempty" name:"Snapshots"`

	// 快照仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterSnapshotResponseParams `json:"Response"`
}

func (r *DescribeClusterSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagnoseRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 报告日期，格式20210301
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 报告返回份数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDiagnoseRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 报告日期，格式20210301
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 报告返回份数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDiagnoseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagnoseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagnoseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagnoseResponseParams struct {
	// 诊断报告个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 诊断报告列表
	DiagnoseResults []*DiagnoseResult `json:"DiagnoseResults,omitnil,omitempty" name:"DiagnoseResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiagnoseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagnoseResponseParams `json:"Response"`
}

func (r *DescribeDiagnoseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagnoseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListRequestParams struct {
	// 索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 索引名，若填空则获取所有索引
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 分页起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页展示数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持索引名：IndexName、索引存储量：IndexStorage、索引创建时间：IndexCreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤索引状态
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// 排序顺序，支持asc、desc，默认为desc 数据格式"asc","desc"
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeIndexListRequest struct {
	*tchttp.BaseRequest
	
	// 索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 索引名，若填空则获取所有索引
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 分页起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页展示数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持索引名：IndexName、索引存储量：IndexStorage、索引创建时间：IndexCreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤索引状态
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// 排序顺序，支持asc、desc，默认为desc 数据格式"asc","desc"
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeIndexListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IndexType")
	delete(f, "InstanceId")
	delete(f, "IndexName")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "IndexStatusList")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexListResponseParams struct {
	// 索引元数据字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexMetaFields []*IndexMetaField `json:"IndexMetaFields,omitnil,omitempty" name:"IndexMetaFields"`

	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexListResponseParams `json:"Response"`
}

func (r *DescribeIndexListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexMetaRequestParams struct {
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 索引名，若填空则获取所有索引
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type DescribeIndexMetaRequest struct {
	*tchttp.BaseRequest
	
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 索引名，若填空则获取所有索引
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *DescribeIndexMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "Username")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexMetaResponseParams struct {
	// 索引元数据字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexMetaField *IndexMetaField `json:"IndexMetaField,omitnil,omitempty" name:"IndexMetaField"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexMetaResponseParams `json:"Response"`
}

func (r *DescribeIndexMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型，默认值为1
	// <li>1, 主日志</li>
	// <li>2, 搜索慢日志</li>
	// <li>3, 索引慢日志</li>
	// <li>4, GC日志</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 搜索词，支持LUCENE语法，如 level:WARN、ip:1.1.1.1、message:test-index等
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 日志开始时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志结束时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值, 默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为100，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 时间排序方式，默认值为0
	// <li>0, 降序</li>
	// <li>1, 升序</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型，默认值为1
	// <li>1, 主日志</li>
	// <li>2, 搜索慢日志</li>
	// <li>3, 索引慢日志</li>
	// <li>4, GC日志</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 搜索词，支持LUCENE语法，如 level:WARN、ip:1.1.1.1、message:test-index等
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 日志开始时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志结束时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值, 默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为100，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 时间排序方式，默认值为0
	// <li>0, 降序</li>
	// <li>1, 升序</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "SearchKey")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsResponseParams struct {
	// 返回的日志条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志详细信息列表
	InstanceLogList []*InstanceLog `json:"InstanceLogList,omitnil,omitempty" name:"InstanceLogList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间, e.g. "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间, e.g. "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间, e.g. "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间, e.g. "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceOperationsResponseParams struct {
	// 操作记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 操作记录
	Operations []*Operation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceOperationsResponseParams `json:"Response"`
}

func (r *DescribeInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancePluginInfo struct {
	// 插件名
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// 插件版本
	PluginVersion *string `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// 插件描述
	PluginDesc *string `json:"PluginDesc,omitnil,omitempty" name:"PluginDesc"`

	// 插件状态：-2 已卸载 -1 卸载中 0 安装中 1 已安装
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 插件是否可卸载
	Removable *bool `json:"Removable,omitnil,omitempty" name:"Removable"`

	// 0：系统插件
	PluginType *int64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`

	// 插件变更时间
	PluginUpdateTime *string `json:"PluginUpdateTime,omitnil,omitempty" name:"PluginUpdateTime"`
}

// Predefined struct for user
type DescribeInstancePluginListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段<li>1：插件名 pluginName</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式<li>0：升序 asc</li><li>1：降序 desc</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 0：系统插件
	PluginType *int64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

type DescribeInstancePluginListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段<li>1：插件名 pluginName</li>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 排序方式<li>0：升序 asc</li><li>1：降序 desc</li>
	OrderByType *string `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 0：系统插件
	PluginType *int64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

func (r *DescribeInstancePluginListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancePluginListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "PluginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancePluginListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancePluginListResponseParams struct {
	// 返回的插件个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 插件信息列表
	PluginList []*DescribeInstancePluginInfo `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstancePluginListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancePluginListResponseParams `json:"Response"`
}

func (r *DescribeInstancePluginListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancePluginListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 集群实例所属可用区，不传则默认所有可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段<li>1：实例ID</li><li>2：实例名称</li><li>3：可用区</li><li>4：创建时间</li>若orderByKey未传递则按创建时间降序排序
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// 排序方式<li>0：升序</li><li>1：降序</li>若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 私有网络vip列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 可用区列表
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// 健康状态筛列表:0表示绿色，1表示黄色，2表示红色,-1表示未知
	HealthStatus []*int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Vpc列表 筛选项
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// cdc集群id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例所属可用区，不传则默认所有可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 集群实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段<li>1：实例ID</li><li>2：实例名称</li><li>3：可用区</li><li>4：创建时间</li>若orderByKey未传递则按创建时间降序排序
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// 排序方式<li>0：升序</li><li>1：降序</li>若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// 节点标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 私有网络vip列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 可用区列表
	ZoneList []*string `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// 健康状态筛列表:0表示绿色，1表示黄色，2表示红色,-1表示未知
	HealthStatus []*int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Vpc列表 筛选项
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// cdc集群id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
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
	delete(f, "Zone")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByKey")
	delete(f, "OrderByType")
	delete(f, "TagList")
	delete(f, "IpList")
	delete(f, "ZoneList")
	delete(f, "HealthStatus")
	delete(f, "VpcIds")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 返回的实例个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表
	InstanceList []*InstanceInfo `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

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

// Predefined struct for user
type DescribeLogstashInstanceLogsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型，默认值为1
	// <li>1, 主日志</li>
	// <li>2, 慢日志</li>
	// <li>3, GC日志</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 搜索词，支持LUCENE语法，如 level:WARN、ip:1.1.1.1、message:test-index等
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 日志开始时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志结束时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值, 默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为100，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 时间排序方式，默认值为0
	// <li>0, 降序</li>
	// <li>1, 升序</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

type DescribeLogstashInstanceLogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志类型，默认值为1
	// <li>1, 主日志</li>
	// <li>2, 慢日志</li>
	// <li>3, GC日志</li>
	LogType *uint64 `json:"LogType,omitnil,omitempty" name:"LogType"`

	// 搜索词，支持LUCENE语法，如 level:WARN、ip:1.1.1.1、message:test-index等
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 日志开始时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 日志结束时间，格式为YYYY-MM-DD HH:MM:SS, 如2019-01-22 20:15:53
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值, 默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值为100，最大值100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 时间排序方式，默认值为0
	// <li>0, 降序</li>
	// <li>1, 升序</li>
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`
}

func (r *DescribeLogstashInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashInstanceLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LogType")
	delete(f, "SearchKey")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogstashInstanceLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashInstanceLogsResponseParams struct {
	// 返回的日志条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志详细信息列表
	InstanceLogList []*InstanceLog `json:"InstanceLogList,omitnil,omitempty" name:"InstanceLogList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogstashInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogstashInstanceLogsResponseParams `json:"Response"`
}

func (r *DescribeLogstashInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashInstanceOperationsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间, e.g. "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间, e.g. "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLogstashInstanceOperationsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 起始时间, e.g. "2019-03-07 16:30:39"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间, e.g. "2019-03-30 20:18:03"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页起始值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeLogstashInstanceOperationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashInstanceOperationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogstashInstanceOperationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashInstanceOperationsResponseParams struct {
	// 操作记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 操作记录
	Operations []*Operation `json:"Operations,omitnil,omitempty" name:"Operations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogstashInstanceOperationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogstashInstanceOperationsResponseParams `json:"Response"`
}

func (r *DescribeLogstashInstanceOperationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashInstanceOperationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashInstancesRequestParams struct {
	// 实例所属可用区，不传则默认所有可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段<li>1：实例ID</li><li>2：实例名称</li><li>3：可用区</li><li>4：创建时间</li>若orderKey未传递则按创建时间降序排序
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// 排序方式<li>0：升序</li><li>1：降序</li>若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// VpcId 筛选项
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type DescribeLogstashInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所属可用区，不传则默认所有可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 分页起始值, 默认值0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小，默认值20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段<li>1：实例ID</li><li>2：实例名称</li><li>3：可用区</li><li>4：创建时间</li>若orderKey未传递则按创建时间降序排序
	OrderByKey *uint64 `json:"OrderByKey,omitnil,omitempty" name:"OrderByKey"`

	// 排序方式<li>0：升序</li><li>1：降序</li>若传递了orderByKey未传递orderByType, 则默认升序
	OrderByType *uint64 `json:"OrderByType,omitnil,omitempty" name:"OrderByType"`

	// VpcId 筛选项
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 标签信息列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *DescribeLogstashInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstanceIds")
	delete(f, "InstanceNames")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderByKey")
	delete(f, "OrderByType")
	delete(f, "VpcIds")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogstashInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashInstancesResponseParams struct {
	// 返回的实例个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详细信息列表
	InstanceList []*LogstashInstanceInfo `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogstashInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogstashInstancesResponseParams `json:"Response"`
}

func (r *DescribeLogstashInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashPipelinesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeLogstashPipelinesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeLogstashPipelinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashPipelinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogstashPipelinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogstashPipelinesResponseParams struct {
	// 管道总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 管道列表
	LogstashPipelineList []*LogstashPipelineInfo `json:"LogstashPipelineList,omitnil,omitempty" name:"LogstashPipelineList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogstashPipelinesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogstashPipelinesResponseParams `json:"Response"`
}

func (r *DescribeLogstashPipelinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogstashPipelinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessInstancesRequestParams struct {
	// 索引集群ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 索引名
	IndexNames []*string `json:"IndexNames,omitnil,omitempty" name:"IndexNames"`

	// 分页起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页展示数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持索引名：IndexName、索引存储量：IndexStorage、索引创建时间：IndexCreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤索引状态
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// 排序顺序，支持asc、desc，默认为desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 索引空间ID列表
	SpaceIds []*string `json:"SpaceIds,omitnil,omitempty" name:"SpaceIds"`

	// 数据链路数据源类型
	DiSourceTypes []*string `json:"DiSourceTypes,omitnil,omitempty" name:"DiSourceTypes"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type DescribeServerlessInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 索引集群ID
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 索引名
	IndexNames []*string `json:"IndexNames,omitnil,omitempty" name:"IndexNames"`

	// 分页起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页展示数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持索引名：IndexName、索引存储量：IndexStorage、索引创建时间：IndexCreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 过滤索引状态
	IndexStatusList []*string `json:"IndexStatusList,omitnil,omitempty" name:"IndexStatusList"`

	// 排序顺序，支持asc、desc，默认为desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 索引空间ID列表
	SpaceIds []*string `json:"SpaceIds,omitnil,omitempty" name:"SpaceIds"`

	// 数据链路数据源类型
	DiSourceTypes []*string `json:"DiSourceTypes,omitnil,omitempty" name:"DiSourceTypes"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *DescribeServerlessInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "IndexNames")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "IndexStatusList")
	delete(f, "Order")
	delete(f, "SpaceIds")
	delete(f, "DiSourceTypes")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessInstancesResponseParams struct {
	// 索引元数据字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexMetaFields []*ServerlessIndexMetaField `json:"IndexMetaFields,omitnil,omitempty" name:"IndexMetaFields"`

	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessInstancesResponseParams `json:"Response"`
}

func (r *DescribeServerlessInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessMetricsRequestParams struct {
	// space空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// index索引id
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// 指标类型，暂时只支持Storage(存储大小),AllMetric(所有存储指标：索引流量、存储大小、文档数量、读请求和写请求)
	MetricType []*string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 时间长度类型DurationType(1: 3小时, 2: 昨天1天,3: 今日0点到现在)
	DurationType *int64 `json:"DurationType,omitnil,omitempty" name:"DurationType"`

	// 索引数据
	BatchIndexList []*string `json:"BatchIndexList,omitnil,omitempty" name:"BatchIndexList"`
}

type DescribeServerlessMetricsRequest struct {
	*tchttp.BaseRequest
	
	// space空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// index索引id
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// 指标类型，暂时只支持Storage(存储大小),AllMetric(所有存储指标：索引流量、存储大小、文档数量、读请求和写请求)
	MetricType []*string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// 时间长度类型DurationType(1: 3小时, 2: 昨天1天,3: 今日0点到现在)
	DurationType *int64 `json:"DurationType,omitnil,omitempty" name:"DurationType"`

	// 索引数据
	BatchIndexList []*string `json:"BatchIndexList,omitnil,omitempty" name:"BatchIndexList"`
}

func (r *DescribeServerlessMetricsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessMetricsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "IndexId")
	delete(f, "MetricType")
	delete(f, "DurationType")
	delete(f, "BatchIndexList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessMetricsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessMetricsResponseParams struct {
	// storage指标值，单位byte
	Storage *float64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// IndexTraffic指标值，单位byte
	IndexTraffic *float64 `json:"IndexTraffic,omitnil,omitempty" name:"IndexTraffic"`

	// 读请求数，单位次数
	ReadReqTimes *int64 `json:"ReadReqTimes,omitnil,omitempty" name:"ReadReqTimes"`

	// 写请求数，单位次数
	WriteReqTimes *int64 `json:"WriteReqTimes,omitnil,omitempty" name:"WriteReqTimes"`

	// 文档数量，单位个数
	DocCount *int64 `json:"DocCount,omitnil,omitempty" name:"DocCount"`

	// 指标数据数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricMapList []*MetricMapByIndexId `json:"MetricMapList,omitnil,omitempty" name:"MetricMapList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessMetricsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessMetricsResponseParams `json:"Response"`
}

func (r *DescribeServerlessMetricsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessSpaceUserRequestParams struct {
	// 空间的ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 游标
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户名列表过滤
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// 用户类型
	UserTypes []*int64 `json:"UserTypes,omitnil,omitempty" name:"UserTypes"`

	// 权限类型
	PrivilegeTypes []*int64 `json:"PrivilegeTypes,omitnil,omitempty" name:"PrivilegeTypes"`
}

type DescribeServerlessSpaceUserRequest struct {
	*tchttp.BaseRequest
	
	// 空间的ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 游标
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户名列表过滤
	UserNames []*string `json:"UserNames,omitnil,omitempty" name:"UserNames"`

	// 用户类型
	UserTypes []*int64 `json:"UserTypes,omitnil,omitempty" name:"UserTypes"`

	// 权限类型
	PrivilegeTypes []*int64 `json:"PrivilegeTypes,omitnil,omitempty" name:"PrivilegeTypes"`
}

func (r *DescribeServerlessSpaceUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessSpaceUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UserNames")
	delete(f, "UserTypes")
	delete(f, "PrivilegeTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessSpaceUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessSpaceUserResponseParams struct {
	// 用户数组
	ServerlessSpaceUsers []*ServerlessSpaceUser `json:"ServerlessSpaceUsers,omitnil,omitempty" name:"ServerlessSpaceUsers"`

	// 用户总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessSpaceUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessSpaceUserResponseParams `json:"Response"`
}

func (r *DescribeServerlessSpaceUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessSpaceUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessSpacesRequestParams struct {
	// 过滤的空间ID
	SpaceIds []*string `json:"SpaceIds,omitnil,omitempty" name:"SpaceIds"`

	// 过滤的空间名
	SpaceNames []*string `json:"SpaceNames,omitnil,omitempty" name:"SpaceNames"`

	// 排序顺序，支持升序asc、降序desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持空间创建时间SpaceCreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// vpcId信息数组
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 分页起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type DescribeServerlessSpacesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤的空间ID
	SpaceIds []*string `json:"SpaceIds,omitnil,omitempty" name:"SpaceIds"`

	// 过滤的空间名
	SpaceNames []*string `json:"SpaceNames,omitnil,omitempty" name:"SpaceNames"`

	// 排序顺序，支持升序asc、降序desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持空间创建时间SpaceCreateTime
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// vpcId信息数组
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 分页起始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *DescribeServerlessSpacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessSpacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceIds")
	delete(f, "SpaceNames")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "VpcIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServerlessSpacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServerlessSpacesResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Serverless空间信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerlessSpaces []*ServerlessSpace `json:"ServerlessSpaces,omitnil,omitempty" name:"ServerlessSpaces"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServerlessSpacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServerlessSpacesResponseParams `json:"Response"`
}

func (r *DescribeServerlessSpacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServerlessSpacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceKibanaToolsRequestParams struct {
	// space的ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DescribeSpaceKibanaToolsRequest struct {
	*tchttp.BaseRequest
	
	// space的ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DescribeSpaceKibanaToolsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceKibanaToolsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceKibanaToolsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceKibanaToolsResponseParams struct {
	// 该token用于登录内嵌kibana
	KibanaToken *string `json:"KibanaToken,omitnil,omitempty" name:"KibanaToken"`

	// token的过期时间
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceKibanaToolsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceKibanaToolsResponseParams `json:"Response"`
}

func (r *DescribeSpaceKibanaToolsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceKibanaToolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCosSnapshotListRequestParams struct {
	// cos桶名
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// bucket 桶下的备份路径
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 云上集群迁移集群名
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`
}

type DescribeUserCosSnapshotListRequest struct {
	*tchttp.BaseRequest
	
	// cos桶名
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// bucket 桶下的备份路径
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 云上集群迁移集群名
	ClusterInstanceId *string `json:"ClusterInstanceId,omitnil,omitempty" name:"ClusterInstanceId"`
}

func (r *DescribeUserCosSnapshotListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCosSnapshotListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CosBucket")
	delete(f, "BasePath")
	delete(f, "ClusterInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCosSnapshotListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCosSnapshotListResponseParams struct {
	// cos 快照信息列表
	CosSnapshotInfoList []*CosSnapShotInfo `json:"CosSnapshotInfoList,omitnil,omitempty" name:"CosSnapshotInfoList"`

	// cos 快照数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserCosSnapshotListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCosSnapshotListResponseParams `json:"Response"`
}

func (r *DescribeUserCosSnapshotListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCosSnapshotListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeViewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsResponseParams struct {
	// 集群维度视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterView *ClusterView `json:"ClusterView,omitnil,omitempty" name:"ClusterView"`

	// 节点维度视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodesView []*NodeView `json:"NodesView,omitnil,omitempty" name:"NodesView"`

	// Kibana维度视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanasView []*KibanaView `json:"KibanasView,omitnil,omitempty" name:"KibanasView"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeViewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeViewsResponseParams `json:"Response"`
}

func (r *DescribeViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiData struct {
	// 数据接入id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiId *string `json:"DiId,omitnil,omitempty" name:"DiId"`

	// 数据接入创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据接入状态，0处理中，1正常，-2删除中，-3已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// cvm数据源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiDataSourceCvm *DiDataSourceCvm `json:"DiDataSourceCvm,omitnil,omitempty" name:"DiDataSourceCvm"`

	// tke数据源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiDataSourceTke *DiDataSourceTke `json:"DiDataSourceTke,omitnil,omitempty" name:"DiDataSourceTke"`

	// serverless目的端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiDataSinkServerless *DiDataSinkServerless `json:"DiDataSinkServerless,omitnil,omitempty" name:"DiDataSinkServerless"`

	// 数据接入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiDataSourceType *string `json:"DiDataSourceType,omitnil,omitempty" name:"DiDataSourceType"`
}

type DiDataSinkServerless struct {
	// serverless实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerlessId *string `json:"ServerlessId,omitnil,omitempty" name:"ServerlessId"`
}

type DiDataSourceCvm struct {
	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 采集路径列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogPaths []*string `json:"LogPaths,omitnil,omitempty" name:"LogPaths"`

	// cvm实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstances []*DiDataSourceCvmInstance `json:"CvmInstances,omitnil,omitempty" name:"CvmInstances"`

	// 采集器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorId *string `json:"CollectorId,omitnil,omitempty" name:"CollectorId"`
}

type DiDataSourceCvmInstance struct {
	// cvm实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`
}

type DiDataSourceTke struct {
	// vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tke实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TkeId *string `json:"TkeId,omitnil,omitempty" name:"TkeId"`

	// 采集器id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorId *string `json:"CollectorId,omitnil,omitempty" name:"CollectorId"`

	// 采集源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorName *string `json:"CollectorName,omitnil,omitempty" name:"CollectorName"`

	// 采集器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorType *string `json:"CollectorType,omitnil,omitempty" name:"CollectorType"`

	// 采集器版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorVersion *string `json:"CollectorVersion,omitnil,omitempty" name:"CollectorVersion"`

	// tke包含的命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeNamespaces []*string `json:"IncludeNamespaces,omitnil,omitempty" name:"IncludeNamespaces"`

	// tke不包含的命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeNamespaces []*string `json:"ExcludeNamespaces,omitnil,omitempty" name:"ExcludeNamespaces"`

	// tke pod标签名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodLabelKeys []*string `json:"PodLabelKeys,omitnil,omitempty" name:"PodLabelKeys"`

	// tke pod标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	PodLabelValues []*string `json:"PodLabelValues,omitnil,omitempty" name:"PodLabelValues"`

	// tke容器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// tke采集器beat配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// /
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// TKE 日志采集路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputPath *string `json:"InputPath,omitnil,omitempty" name:"InputPath"`
}

type DiSourceCvm struct {
	// 数据源所属vpc id，创建后不允许修改
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// cvm列表
	CvmIds []*string `json:"CvmIds,omitnil,omitempty" name:"CvmIds"`

	// 采集路径列表，支持通配符
	LogPaths []*string `json:"LogPaths,omitnil,omitempty" name:"LogPaths"`
}

type DiSourceTke struct {
	// 数据源所属vpc id，创建后不允许修改
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// tke实例id，创建后不允许修改
	TkeId *string `json:"TkeId,omitnil,omitempty" name:"TkeId"`

	// tke包含的命名空间
	IncludeNamespaces []*string `json:"IncludeNamespaces,omitnil,omitempty" name:"IncludeNamespaces"`

	// tke不包含的命名空间
	ExcludeNamespaces []*string `json:"ExcludeNamespaces,omitnil,omitempty" name:"ExcludeNamespaces"`

	// tke容器名称
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`

	// tke日志内容过滤
	LogFilters *string `json:"LogFilters,omitnil,omitempty" name:"LogFilters"`

	// tke beats配置项
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// tke pod标签
	PodLabel []*DiSourceTkePodLabel `json:"PodLabel,omitnil,omitempty" name:"PodLabel"`

	// /
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// tke 日志采集路径
	InputPath *string `json:"InputPath,omitnil,omitempty" name:"InputPath"`
}

type DiSourceTkePodLabel struct {
	// 标签key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DiagnoseInstanceRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要触发的诊断项
	DiagnoseJobs []*string `json:"DiagnoseJobs,omitnil,omitempty" name:"DiagnoseJobs"`

	// 需要诊断的索引，支持通配符
	DiagnoseIndices *string `json:"DiagnoseIndices,omitnil,omitempty" name:"DiagnoseIndices"`
}

type DiagnoseInstanceRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要触发的诊断项
	DiagnoseJobs []*string `json:"DiagnoseJobs,omitnil,omitempty" name:"DiagnoseJobs"`

	// 需要诊断的索引，支持通配符
	DiagnoseIndices *string `json:"DiagnoseIndices,omitnil,omitempty" name:"DiagnoseIndices"`
}

func (r *DiagnoseInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DiagnoseInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DiagnoseJobs")
	delete(f, "DiagnoseIndices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DiagnoseInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DiagnoseInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DiagnoseInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DiagnoseInstanceResponseParams `json:"Response"`
}

func (r *DiagnoseInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DiagnoseInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagnoseJobMeta struct {
	// 智能运维诊断项英文名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 智能运维诊断项中文名
	JobZhName *string `json:"JobZhName,omitnil,omitempty" name:"JobZhName"`

	// 智能运维诊断项描述
	JobDescription *string `json:"JobDescription,omitnil,omitempty" name:"JobDescription"`
}

type DiagnoseJobResult struct {
	// 诊断项名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// 诊断项状态：-2失败，-1待重试，0运行中，1成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 诊断项得分
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 诊断摘要
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 诊断建议
	Advise *string `json:"Advise,omitnil,omitempty" name:"Advise"`

	// 诊断详情
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 诊断指标详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricDetails []*MetricDetail `json:"MetricDetails,omitnil,omitempty" name:"MetricDetails"`

	// 诊断日志详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogDetails []*LogDetail `json:"LogDetails,omitnil,omitempty" name:"LogDetails"`

	// 诊断配置详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SettingDetails []*SettingDetail `json:"SettingDetails,omitnil,omitempty" name:"SettingDetails"`
}

type DiagnoseResult struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 诊断报告ID
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 诊断触发时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 诊断是否完成
	Completed *bool `json:"Completed,omitnil,omitempty" name:"Completed"`

	// 诊断总得分
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 诊断类型，2 定时诊断，3 客户手动触发诊断
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 诊断参数，如诊断时间，诊断索引等
	JobParam *JobParam `json:"JobParam,omitnil,omitempty" name:"JobParam"`

	// 诊断项结果列表
	JobResults []*DiagnoseJobResult `json:"JobResults,omitnil,omitempty" name:"JobResults"`
}

type DictInfo struct {
	// 词典键值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 词典名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 词典大小，单位B
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type Dimension struct {
	// 智能运维指标维度Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 智能运维指标维度值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type EnableScheduleOperationDuration struct {
	// 支持开启异步任务的日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Periods []*string `json:"Periods,omitnil,omitempty" name:"Periods"`

	// 支持开启异步的开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// 支持开启异步的结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 支持开启异步的时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type EsAcl struct {
	// kibana访问黑名单
	BlackIpList []*string `json:"BlackIpList,omitnil,omitempty" name:"BlackIpList"`

	// kibana访问白名单
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

type EsConfigSetInfo struct {
	// 配置组类型，如ldap,ad等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// "{\"order\":0,\"url\":\"ldap://10.0.1.72:389\",\"bind_dn\":\"cn=admin,dc=tencent,dc=com\",\"user_search.base_dn\":\"dc=tencent,dc=com\",\"user_search.filter\":\"(cn={0})\",\"group_search.base_dn\":\"dc=tencent,dc=com\"}"
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`
}

type EsDictionaryInfo struct {
	// 启用词词典列表
	MainDict []*DictInfo `json:"MainDict,omitnil,omitempty" name:"MainDict"`

	// 停用词词典列表
	Stopwords []*DictInfo `json:"Stopwords,omitnil,omitempty" name:"Stopwords"`

	// QQ分词词典列表
	QQDict []*DictInfo `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// 同义词词典列表
	Synonym []*DictInfo `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// 更新词典类型
	UpdateType *string `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// ansj启用词词典列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnsjMain []*DictInfo `json:"AnsjMain,omitnil,omitempty" name:"AnsjMain"`

	// ansj停用词词典列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnsjStop []*DictInfo `json:"AnsjStop,omitnil,omitempty" name:"AnsjStop"`

	// ansj歧义词库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnsjAmbiguity []*DictInfo `json:"AnsjAmbiguity,omitnil,omitempty" name:"AnsjAmbiguity"`

	// ansj同义词词典列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnsjSynonyms []*DictInfo `json:"AnsjSynonyms,omitnil,omitempty" name:"AnsjSynonyms"`
}

type EsPublicAcl struct {
	// 访问黑名单
	BlackIpList []*string `json:"BlackIpList,omitnil,omitempty" name:"BlackIpList"`

	// 访问白名单
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

type Failures struct {
	// 备份失败的索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 快照失败的分片号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardId *int64 `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// 快照失败的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 快照失败的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type GetDiagnoseSettingsRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type GetDiagnoseSettingsRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *GetDiagnoseSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDiagnoseSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDiagnoseSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDiagnoseSettingsResponseParams struct {
	// 智能运维诊断项和元信息
	DiagnoseJobMetas []*DiagnoseJobMeta `json:"DiagnoseJobMetas,omitnil,omitempty" name:"DiagnoseJobMetas"`

	// 0：开启智能运维；-1：关闭智能运维
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 智能运维每天定时巡检时间
	CronTime *string `json:"CronTime,omitnil,omitempty" name:"CronTime"`

	// 智能运维当天已手动触发诊断次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 智能运维每天最大可手动触发次数
	MaxCount *int64 `json:"MaxCount,omitnil,omitempty" name:"MaxCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDiagnoseSettingsResponse struct {
	*tchttp.BaseResponse
	Response *GetDiagnoseSettingsResponseParams `json:"Response"`
}

func (r *GetDiagnoseSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDiagnoseSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRequestTargetNodeTypesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type GetRequestTargetNodeTypesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *GetRequestTargetNodeTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRequestTargetNodeTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRequestTargetNodeTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRequestTargetNodeTypesResponseParams struct {
	// 接收请求的目标节点类型列表
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRequestTargetNodeTypesResponse struct {
	*tchttp.BaseResponse
	Response *GetRequestTargetNodeTypesResponseParams `json:"Response"`
}

func (r *GetRequestTargetNodeTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRequestTargetNodeTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GpuInfo struct {
	// Gpu块数
	GpuCount *int64 `json:"GpuCount,omitnil,omitempty" name:"GpuCount"`

	// Gpu类型
	GpuType *string `json:"GpuType,omitnil,omitempty" name:"GpuType"`
}

type IndexMetaField struct {
	// 索引类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 索引名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 索引元数据JSON
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// 索引状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`

	// 索引存储大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// 索引创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`

	// 后备索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackingIndices []*BackingIndexMetaField `json:"BackingIndices,omitnil,omitempty" name:"BackingIndices"`

	// 索引所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 索引所属集群名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 索引所属集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 索引生命周期字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexPolicyField *IndexPolicyField `json:"IndexPolicyField,omitnil,omitempty" name:"IndexPolicyField"`

	// 索引自治字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexOptionsField *IndexOptionsField `json:"IndexOptionsField,omitnil,omitempty" name:"IndexOptionsField"`

	// 索引配置字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexSettingsField *IndexSettingsField `json:"IndexSettingsField,omitnil,omitempty" name:"IndexSettingsField"`

	// 索引别名字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexAliasesField []*string `json:"IndexAliasesField,omitnil,omitempty" name:"IndexAliasesField"`

	// 索引所属集群APP ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 索引文档数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexDocs *uint64 `json:"IndexDocs,omitnil,omitempty" name:"IndexDocs"`
}

type IndexOptionsField struct {
	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireMaxAge *string `json:"ExpireMaxAge,omitnil,omitempty" name:"ExpireMaxAge"`

	// 过期大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireMaxSize *string `json:"ExpireMaxSize,omitnil,omitempty" name:"ExpireMaxSize"`

	// 滚动周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RolloverMaxAge *string `json:"RolloverMaxAge,omitnil,omitempty" name:"RolloverMaxAge"`

	// 是否开启动态滚动
	// 注意：此字段可能返回 null，表示取不到有效值。
	RolloverDynamic *string `json:"RolloverDynamic,omitnil,omitempty" name:"RolloverDynamic"`

	// 是否开启动态分片
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardNumDynamic *string `json:"ShardNumDynamic,omitnil,omitempty" name:"ShardNumDynamic"`

	// 时间分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimestampField *string `json:"TimestampField,omitnil,omitempty" name:"TimestampField"`

	// 写入模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	WriteMode *string `json:"WriteMode,omitnil,omitempty" name:"WriteMode"`
}

type IndexPolicyField struct {
	// 是否开启warm阶段
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmEnable *string `json:"WarmEnable,omitnil,omitempty" name:"WarmEnable"`

	// warm阶段转入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmMinAge *string `json:"WarmMinAge,omitnil,omitempty" name:"WarmMinAge"`

	// 是否开启cold阶段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdEnable *string `json:"ColdEnable,omitnil,omitempty" name:"ColdEnable"`

	// cold阶段转入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdMinAge *string `json:"ColdMinAge,omitnil,omitempty" name:"ColdMinAge"`

	// 是否开启frozen阶段
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenEnable *string `json:"FrozenEnable,omitnil,omitempty" name:"FrozenEnable"`

	// frozen阶段转入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenMinAge *string `json:"FrozenMinAge,omitnil,omitempty" name:"FrozenMinAge"`

	// /
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdAction *string `json:"ColdAction,omitnil,omitempty" name:"ColdAction"`
}

type IndexSettingsField struct {
	// 索引主分片数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfShards *string `json:"NumberOfShards,omitnil,omitempty" name:"NumberOfShards"`

	// 索引副本分片数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfReplicas *string `json:"NumberOfReplicas,omitnil,omitempty" name:"NumberOfReplicas"`

	// 索引刷新频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefreshInterval *string `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`
}

// Predefined struct for user
type InquirePriceRenewInstanceRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InquirePriceRenewInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *InquirePriceRenewInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRenewInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRenewInstanceResponseParams struct {
	// 刊例价，即集群原始价格
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折后价
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 折扣，如65折
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 货币，如CNY代表人民币
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceRenewInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRenewInstanceResponseParams `json:"Response"`
}

func (r *InquirePriceRenewInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRenewInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallInstanceModelRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户上传到cos的模型地址，单次请求限制一个。cos文件为压缩文件，格式包括：zip、tgz和tar.gz
	UsrCosModelUrlList []*string `json:"UsrCosModelUrlList,omitnil,omitempty" name:"UsrCosModelUrlList"`

	// 客户指定安装的模型名称，可为空，默认为模型文件名
	ModelNames []*string `json:"ModelNames,omitnil,omitempty" name:"ModelNames"`

	// 模型使用的任务类型，包括：fill_mask, ner, question_answering, text_classification, text_embedding, text_expansion, text_similarity和zero_shot_classification，默认为text_embedding
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`
}

type InstallInstanceModelRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户上传到cos的模型地址，单次请求限制一个。cos文件为压缩文件，格式包括：zip、tgz和tar.gz
	UsrCosModelUrlList []*string `json:"UsrCosModelUrlList,omitnil,omitempty" name:"UsrCosModelUrlList"`

	// 客户指定安装的模型名称，可为空，默认为模型文件名
	ModelNames []*string `json:"ModelNames,omitnil,omitempty" name:"ModelNames"`

	// 模型使用的任务类型，包括：fill_mask, ner, question_answering, text_classification, text_embedding, text_expansion, text_similarity和zero_shot_classification，默认为text_embedding
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`
}

func (r *InstallInstanceModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallInstanceModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UsrCosModelUrlList")
	delete(f, "ModelNames")
	delete(f, "TaskTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InstallInstanceModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InstallInstanceModelResponseParams struct {
	// 发起异步流程的flowId
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 调用接口的错误信息
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InstallInstanceModelResponse struct {
	*tchttp.BaseResponse
	Response *InstallInstanceModelResponseParams `json:"Response"`
}

func (r *InstallInstanceModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InstallInstanceModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 用户ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 实例所属VPC的UID
	VpcUid *string `json:"VpcUid,omitnil,omitempty" name:"VpcUid"`

	// 实例所属子网的UID
	SubnetUid *string `json:"SubnetUid,omitnil,omitempty" name:"SubnetUid"`

	// 实例状态，0:处理中,1:正常,-1:停止,-2:销毁中,-3:已销毁, -4:隔离中,2:创建集群时初始化中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 自动续费标识。取值范围：
	// RENEW_FLAG_AUTO：自动续费  
	// RENEW_FLAG_MANUAL：不自动续费
	// 默认取值：
	// RENEW_FLAG_DEFAULT：不自动续费
	// 若该参数指定为 RENEW_FLAG_AUTO，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 实例计费模式。取值范围：  PREPAID：表示预付费，即包年包月  POSTPAID_BY_HOUR：表示后付费，即按量计费  CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月购买时长,单位:月
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点个数
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 节点CPU核数
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 节点内存大小，单位GB
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 节点磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 节点磁盘大小，单位GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// ES域名
	EsDomain *string `json:"EsDomain,omitnil,omitempty" name:"EsDomain"`

	// ES VIP
	EsVip *string `json:"EsVip,omitnil,omitempty" name:"EsVip"`

	// ES端口
	EsPort *uint64 `json:"EsPort,omitnil,omitempty" name:"EsPort"`

	// Kibana访问url
	KibanaUrl *string `json:"KibanaUrl,omitnil,omitempty" name:"KibanaUrl"`

	// ES版本号
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// ES配置项
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// Kibana访问控制配置
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例最后修改操作时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 实例到期时间
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 实例类型（实例类型标识，当前只有1,2两种）
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// Ik分词器配置
	IkConfig *EsDictionaryInfo `json:"IkConfig,omitnil,omitempty" name:"IkConfig"`

	// 专用主节点配置
	MasterNodeInfo *MasterNodeInfo `json:"MasterNodeInfo,omitnil,omitempty" name:"MasterNodeInfo"`

	// cos自动备份配置
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// 是否允许cos自动备份
	AllowCosBackup *bool `json:"AllowCosBackup,omitnil,omitempty" name:"AllowCosBackup"`

	// 实例拥有的标签列表
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 是否为冷热集群<li>true: 冷热集群</li><li>false: 非冷热集群</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableHotWarmMode *bool `json:"EnableHotWarmMode,omitnil,omitempty" name:"EnableHotWarmMode"`

	// 温节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmNodeType *string `json:"WarmNodeType,omitnil,omitempty" name:"WarmNodeType"`

	// 温节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmNodeNum *uint64 `json:"WarmNodeNum,omitnil,omitempty" name:"WarmNodeNum"`

	// 温节点CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmCpuNum *uint64 `json:"WarmCpuNum,omitnil,omitempty" name:"WarmCpuNum"`

	// 温节点内存内存大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmMemSize *uint64 `json:"WarmMemSize,omitnil,omitempty" name:"WarmMemSize"`

	// 温节点磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmDiskType *string `json:"WarmDiskType,omitnil,omitempty" name:"WarmDiskType"`

	// 温节点磁盘大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	WarmDiskSize *uint64 `json:"WarmDiskSize,omitnil,omitempty" name:"WarmDiskSize"`

	// 集群节点信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// Es公网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsPublicUrl *string `json:"EsPublicUrl,omitnil,omitempty" name:"EsPublicUrl"`

	// 多可用区网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// 部署模式<li>0：单可用区</li><li>1：多可用区，北京、上海、上海金融、广州、南京、香港、新加坡、法兰克福（白名单控制）</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// ES公网访问状态<li>OPEN：开启</li><li>CLOSE：关闭</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// ES公网访问控制配置
	EsPublicAcl *EsAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// Kibana内网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitnil,omitempty" name:"KibanaPrivateUrl"`

	// Kibana公网访问状态<li>OPEN：开启</li><li>CLOSE：关闭</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态<li>OPEN：开启</li><li>CLOSE：关闭</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *uint64 `json:"SecurityType,omitnil,omitempty" name:"SecurityType"`

	// 场景化模板类型：0、不开启；1、通用场景；2、日志场景；3、搜索场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// Kibana节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaNodeInfo *KibanaNodeInfo `json:"KibanaNodeInfo,omitnil,omitempty" name:"KibanaNodeInfo"`

	// 可视化节点配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// JDK类型，oracle或kona
	// 注意：此字段可能返回 null，表示取不到有效值。
	Jdk *string `json:"Jdk,omitnil,omitempty" name:"Jdk"`

	// 集群网络通讯协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 安全组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// 冷节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdNodeType *string `json:"ColdNodeType,omitnil,omitempty" name:"ColdNodeType"`

	// 冷节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdNodeNum *uint64 `json:"ColdNodeNum,omitnil,omitempty" name:"ColdNodeNum"`

	// 冷节点CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdCpuNum *uint64 `json:"ColdCpuNum,omitnil,omitempty" name:"ColdCpuNum"`

	// 冷节点内存大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdMemSize *uint64 `json:"ColdMemSize,omitnil,omitempty" name:"ColdMemSize"`

	// 冷节点磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdDiskType *string `json:"ColdDiskType,omitnil,omitempty" name:"ColdDiskType"`

	// 冷节点磁盘大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColdDiskSize *uint64 `json:"ColdDiskSize,omitnil,omitempty" name:"ColdDiskSize"`

	// 冻节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenNodeType *string `json:"FrozenNodeType,omitnil,omitempty" name:"FrozenNodeType"`

	// 冻节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenNodeNum *uint64 `json:"FrozenNodeNum,omitnil,omitempty" name:"FrozenNodeNum"`

	// 冻节点CPU核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenCpuNum *uint64 `json:"FrozenCpuNum,omitnil,omitempty" name:"FrozenCpuNum"`

	// 冻节点内存大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenMemSize *uint64 `json:"FrozenMemSize,omitnil,omitempty" name:"FrozenMemSize"`

	// 冻节点磁盘类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenDiskType *string `json:"FrozenDiskType,omitnil,omitempty" name:"FrozenDiskType"`

	// 冻节点磁盘大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenDiskSize *uint64 `json:"FrozenDiskSize,omitnil,omitempty" name:"FrozenDiskSize"`

	// 集群健康状态 -1 未知；0 Green; 1 Yellow; 2 Red
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthStatus *int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// https集群内网url
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsPrivateUrl *string `json:"EsPrivateUrl,omitnil,omitempty" name:"EsPrivateUrl"`

	// https集群内网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsPrivateDomain *string `json:"EsPrivateDomain,omitnil,omitempty" name:"EsPrivateDomain"`

	// 集群的配置组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsConfigSets []*EsConfigSetInfo `json:"EsConfigSets,omitnil,omitempty" name:"EsConfigSets"`

	// 集群可维护时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// web节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionalWebServiceInfos []*OptionalWebServiceInfo `json:"OptionalWebServiceInfos,omitnil,omitempty" name:"OptionalWebServiceInfos"`

	// 自治索引开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoIndexEnabled *bool `json:"AutoIndexEnabled,omitnil,omitempty" name:"AutoIndexEnabled"`

	// 是否支持存储计算分离
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableHybridStorage *bool `json:"EnableHybridStorage,omitnil,omitempty" name:"EnableHybridStorage"`

	// 流程进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessPercent *float64 `json:"ProcessPercent,omitnil,omitempty" name:"ProcessPercent"`

	// Kibana的alerting外网告警策略<li>OPEN：开启</li><li>CLOSE：关闭</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`

	// 本月是否有内核可以更新：false-无，true-有
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasKernelUpgrade *bool `json:"HasKernelUpgrade,omitnil,omitempty" name:"HasKernelUpgrade"`

	// cdcId，使用cdc子网时传递
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// kibana内网vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateVip *string `json:"KibanaPrivateVip,omitnil,omitempty" name:"KibanaPrivateVip"`

	// 自定义kibana内网url
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomKibanaPrivateUrl *string `json:"CustomKibanaPrivateUrl,omitnil,omitempty" name:"CustomKibanaPrivateUrl"`

	// 节点出站访问详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutboundPublicAcls []*OutboundPublicAcl `json:"OutboundPublicAcls,omitnil,omitempty" name:"OutboundPublicAcls"`

	// 网络连接方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetConnectScheme *string `json:"NetConnectScheme,omitnil,omitempty" name:"NetConnectScheme"`

	// 置放群组相关参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisasterRecoverGroupAffinity *uint64 `json:"DisasterRecoverGroupAffinity,omitnil,omitempty" name:"DisasterRecoverGroupAffinity"`

	// 子产品ID枚举值： 开源版："sp_es_io2"， 基础版："sp_es_basic"，白金版："sp_es_platinum"，企业版："sp_es_enterprise"，CDC白金版："sp_es_cdc_platinum"，日志增强版："sp_es_enlogging"，tsearch："sp_tsearch_io2"，logstash："sp_es_logstash" ，可以为空，为空的时候后台取LicenseType映射该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 存算分离cos用量，单位M
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketStorageSize *uint64 `json:"CosBucketStorageSize,omitnil,omitempty" name:"CosBucketStorageSize"`

	// 读写分离模式：0-不开启，1-本地读写分离，2-远端读写分离
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// 是否有置放群组异步调度任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 异步调度任务的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 开启集群保护：OPEN-开启，CLOSE-关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDestroyProtection *string `json:"EnableDestroyProtection,omitnil,omitempty" name:"EnableDestroyProtection"`

	// kibana内网访问地址
	ShowKibanaIpPort *string `json:"ShowKibanaIpPort,omitnil,omitempty" name:"ShowKibanaIpPort"`

	// 是否为CDZLite可用区
	IsCdzLite *bool `json:"IsCdzLite,omitnil,omitempty" name:"IsCdzLite"`
}

type InstanceLog struct {
	// 日志时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志级别
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 集群节点ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 日志内容
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 集群节点ID
	NodeID *string `json:"NodeID,omitnil,omitempty" name:"NodeID"`
}

type JobParam struct {
	// 诊断项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Jobs []*string `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// 诊断索引
	Indices *string `json:"Indices,omitnil,omitempty" name:"Indices"`

	// 历史诊断时间
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`
}

type KeyValue struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KibanaNodeInfo struct {
	// Kibana节点规格
	KibanaNodeType *string `json:"KibanaNodeType,omitnil,omitempty" name:"KibanaNodeType"`

	// Kibana节点个数
	KibanaNodeNum *uint64 `json:"KibanaNodeNum,omitnil,omitempty" name:"KibanaNodeNum"`

	// Kibana节点CPU数
	KibanaNodeCpuNum *uint64 `json:"KibanaNodeCpuNum,omitnil,omitempty" name:"KibanaNodeCpuNum"`

	// Kibana节点内存GB
	KibanaNodeMemSize *uint64 `json:"KibanaNodeMemSize,omitnil,omitempty" name:"KibanaNodeMemSize"`

	// Kibana节点磁盘类型
	KibanaNodeDiskType *string `json:"KibanaNodeDiskType,omitnil,omitempty" name:"KibanaNodeDiskType"`

	// Kibana节点磁盘大小
	KibanaNodeDiskSize *uint64 `json:"KibanaNodeDiskSize,omitnil,omitempty" name:"KibanaNodeDiskSize"`
}

type KibanaPublicAcl struct {
	// kibana访问白名单
	WhiteIpList []*string `json:"WhiteIpList,omitnil,omitempty" name:"WhiteIpList"`
}

type KibanaView struct {
	// Kibana节点IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点总磁盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘使用率
	DiskUsage *float64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 节点内存大小
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 内存使用率
	MemUsage *float64 `json:"MemUsage,omitnil,omitempty" name:"MemUsage"`

	// 节点cpu个数
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// cpu使用率
	CpuUsage *float64 `json:"CpuUsage,omitnil,omitempty" name:"CpuUsage"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// ts-0noqayxu-az6-hot-03222010-0
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type LocalDiskInfo struct {
	// 本地盘类型<li>LOCAL_SATA：大数据型</li><li>NVME_SSD：高IO型</li>
	LocalDiskType *string `json:"LocalDiskType,omitnil,omitempty" name:"LocalDiskType"`

	// 本地盘单盘大小
	LocalDiskSize *uint64 `json:"LocalDiskSize,omitnil,omitempty" name:"LocalDiskSize"`

	// 本地盘块数
	LocalDiskCount *uint64 `json:"LocalDiskCount,omitnil,omitempty" name:"LocalDiskCount"`
}

type LogDetail struct {
	// 日志异常名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 日志异常处理建议
	Advise *string `json:"Advise,omitnil,omitempty" name:"Advise"`

	// 日志异常名出现次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type LogstashBindedES struct {
	// ES集群ID
	ESInstanceId *string `json:"ESInstanceId,omitnil,omitempty" name:"ESInstanceId"`

	// ES集群用户名
	ESUserName *string `json:"ESUserName,omitnil,omitempty" name:"ESUserName"`

	// ES集群密码
	ESPassword *string `json:"ESPassword,omitnil,omitempty" name:"ESPassword"`
}

type LogstashExtendedFile struct {
	// 扩展文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 扩展文件大小，单位B
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

type LogstashInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 用户ID
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 实例所属VPC的ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 实例所属子网的ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例状态，0:处理中,1:正常,-1停止,-2:销毁中,-3:已销毁
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例计费模式。取值范围：  PREPAID：表示预付费，即包年包月  POSTPAID_BY_HOUR：表示后付费，即按量计费  CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 包年包月购买时长,单位:月
	ChargePeriod *uint64 `json:"ChargePeriod,omitnil,omitempty" name:"ChargePeriod"`

	// 自动续费标识。取值范围：  NOTIFY_AND_AUTO_RENEW：通知过期且自动续费  NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费  DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费  默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 节点规格<li>LOGSTASH.S1.SMALL2：1核2G</li><li>LOGSTASH.S1.MEDIUM4：2核4G</li><li>LOGSTASH.S1.MEDIUM8：2核8G</li><li>LOGSTASH.S1.LARGE16：4核16G</li><li>LOGSTASH.S1.2XLARGE32：8核32G</li><li>LOGSTASH.S1.4XLARGE32：16核32G</li><li>LOGSTASH.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点个数
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 节点磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 节点磁盘大小，单位GB
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// Logstash版本号
	LogstashVersion *string `json:"LogstashVersion,omitnil,omitempty" name:"LogstashVersion"`

	// License类型<li>oss：开源版</li><li>xpack：基础版</li>默认值xpack
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例最后修改操作时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 实例到期时间
	Deadline *string `json:"Deadline,omitnil,omitempty" name:"Deadline"`

	// 实例节点类型
	Nodes []*LogstashNodeInfo `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 实例绑定的ES集群ID
	BindedESInstanceId *string `json:"BindedESInstanceId,omitnil,omitempty" name:"BindedESInstanceId"`

	// 实例的YML配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	YMLConfig *string `json:"YMLConfig,omitnil,omitempty" name:"YMLConfig"`

	// 扩展文件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendedFiles []*LogstashExtendedFile `json:"ExtendedFiles,omitnil,omitempty" name:"ExtendedFiles"`

	// 可维护时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationDuration *OperationDuration `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// CPU数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 实例标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *uint64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 部署模式，0：单可用区、1：多可用区
	DeployMode *uint64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 多可用区部署时可用区的详细信息
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`
}

type LogstashNodeInfo struct {
	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 节点所在zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type LogstashPipeline struct {
	// 管道ID
	PipelineId *string `json:"PipelineId,omitnil,omitempty" name:"PipelineId"`

	// 管道描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PipelineDesc *string `json:"PipelineDesc,omitnil,omitempty" name:"PipelineDesc"`

	// 管道配置内容
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 管道的Worker数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workers *uint64 `json:"Workers,omitnil,omitempty" name:"Workers"`

	// 管道批处理大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchSize *uint64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 管道批处理延迟
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchDelay *uint64 `json:"BatchDelay,omitnil,omitempty" name:"BatchDelay"`

	// 管道缓冲队列类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 管道缓冲队列大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueMaxBytes *string `json:"QueueMaxBytes,omitnil,omitempty" name:"QueueMaxBytes"`

	// 管道缓冲队列检查点写入数
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueCheckPointWrites *uint64 `json:"QueueCheckPointWrites,omitnil,omitempty" name:"QueueCheckPointWrites"`
}

type LogstashPipelineInfo struct {
	// 管道ID
	PipelineId *string `json:"PipelineId,omitnil,omitempty" name:"PipelineId"`

	// 管道描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PipelineDesc *string `json:"PipelineDesc,omitnil,omitempty" name:"PipelineDesc"`

	// 管道配置内容
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 管道状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 管道的Worker数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workers *uint64 `json:"Workers,omitnil,omitempty" name:"Workers"`

	// 管道批处理大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchSize *uint64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 管道批处理延迟
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchDelay *uint64 `json:"BatchDelay,omitnil,omitempty" name:"BatchDelay"`

	// 管道缓冲队列类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 管道缓冲队列大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueMaxBytes *string `json:"QueueMaxBytes,omitnil,omitempty" name:"QueueMaxBytes"`

	// 管道缓冲队列检查点写入数
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueCheckPointWrites *uint64 `json:"QueueCheckPointWrites,omitnil,omitempty" name:"QueueCheckPointWrites"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type MasterNodeInfo struct {
	// 是否启用了专用主节点
	EnableDedicatedMaster *bool `json:"EnableDedicatedMaster,omitnil,omitempty" name:"EnableDedicatedMaster"`

	// 专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// 专用主节点个数
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// 专用主节点CPU核数
	MasterNodeCpuNum *uint64 `json:"MasterNodeCpuNum,omitnil,omitempty" name:"MasterNodeCpuNum"`

	// 专用主节点内存大小，单位GB
	MasterNodeMemSize *uint64 `json:"MasterNodeMemSize,omitnil,omitempty" name:"MasterNodeMemSize"`

	// 专用主节点磁盘大小，单位GB
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// 专用主节点磁盘类型
	MasterNodeDiskType *string `json:"MasterNodeDiskType,omitnil,omitempty" name:"MasterNodeDiskType"`
}

type Metric struct {
	// 指标维度族
	Dimensions []*Dimension `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 指标值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type MetricAllData struct {
	// 索引流量
	IndexTraffic *float64 `json:"IndexTraffic,omitnil,omitempty" name:"IndexTraffic"`

	// 存储大小
	Storage *float64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 读请求次数
	ReadReqTimes *int64 `json:"ReadReqTimes,omitnil,omitempty" name:"ReadReqTimes"`

	// 写请求次数
	WriteReqTimes *int64 `json:"WriteReqTimes,omitnil,omitempty" name:"WriteReqTimes"`

	// 文档数量
	DocCount *int64 `json:"DocCount,omitnil,omitempty" name:"DocCount"`
}

type MetricDetail struct {
	// 指标详情名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 指标详情值
	Metrics []*Metric `json:"Metrics,omitnil,omitempty" name:"Metrics"`
}

type MetricMapByIndexId struct {
	// 实例id
	IndexId *string `json:"IndexId,omitnil,omitempty" name:"IndexId"`

	// 指标数据
	MetricAllData *MetricAllData `json:"MetricAllData,omitnil,omitempty" name:"MetricAllData"`
}

// Predefined struct for user
type ModifyEsVipSecurityGroupRequestParams struct {
	// es集群的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 安全组id列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

type ModifyEsVipSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// es集群的实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 安全组id列表
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`
}

func (r *ModifyEsVipSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEsVipSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEsVipSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEsVipSecurityGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEsVipSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEsVipSecurityGroupResponseParams `json:"Response"`
}

func (r *ModifyEsVipSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEsVipSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {
	// 节点数量
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点类型<li>hotData: 热数据节点</li>
	// <li>warmData: 冷数据节点</li>
	// <li>dedicatedMaster: 专用主节点</li>
	// 默认值为hotData
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点磁盘类型<li>CLOUD_SSD：SSD云硬盘</li><li>CLOUD_PREMIUM：高硬能云硬盘</li>默认值CLOUD_SSD
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 节点磁盘容量（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 节点本地盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalDiskInfo *LocalDiskInfo `json:"LocalDiskInfo,omitnil,omitempty" name:"LocalDiskInfo"`

	// 节点磁盘块数
	DiskCount *uint64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 节点磁盘是否加密 0: 不加密，1: 加密；默认不加密
	DiskEncrypt *uint64 `json:"DiskEncrypt,omitnil,omitempty" name:"DiskEncrypt"`

	// cpu数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	CpuNum *uint64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// 内存大小，单位GB
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// /
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskEnhance *int64 `json:"DiskEnhance,omitnil,omitempty" name:"DiskEnhance"`

	// 节点Gpu信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GpuInfo *GpuInfo `json:"GpuInfo,omitnil,omitempty" name:"GpuInfo"`
}

type NodeView struct {
	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点IP
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// 节点是否可见
	Visible *float64 `json:"Visible,omitnil,omitempty" name:"Visible"`

	// 是否熔断
	Break *float64 `json:"Break,omitnil,omitempty" name:"Break"`

	// 节点总磁盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘使用率
	DiskUsage *float64 `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// 节点内存大小，单位GB
	MemSize *int64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 内存使用率
	MemUsage *float64 `json:"MemUsage,omitnil,omitempty" name:"MemUsage"`

	// 节点cpu个数
	CpuNum *int64 `json:"CpuNum,omitnil,omitempty" name:"CpuNum"`

	// cpu使用率
	CpuUsage *float64 `json:"CpuUsage,omitnil,omitempty" name:"CpuUsage"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 节点角色
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 节点HTTP IP
	NodeHttpIp *string `json:"NodeHttpIp,omitnil,omitempty" name:"NodeHttpIp"`

	// JVM内存使用率
	JvmMemUsage *float64 `json:"JvmMemUsage,omitnil,omitempty" name:"JvmMemUsage"`

	// 节点分片数
	ShardNum *int64 `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 节点上磁盘ID列表
	DiskIds []*string `json:"DiskIds,omitnil,omitempty" name:"DiskIds"`

	// 是否为隐藏可用区
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`

	// 是否充当协调节点的角色
	IsCoordinationNode *bool `json:"IsCoordinationNode,omitnil,omitempty" name:"IsCoordinationNode"`

	// CVM运行状态
	CVMStatus *string `json:"CVMStatus,omitnil,omitempty" name:"CVMStatus"`

	// cvm绑定的置放群组的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVMDisasterRecoverGroupId *string `json:"CVMDisasterRecoverGroupId,omitnil,omitempty" name:"CVMDisasterRecoverGroupId"`

	// cvm绑定置放群组的状态。2: 已绑定；1: 绑定中；0: 未开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVMDisasterRecoverGroupStatus *int64 `json:"CVMDisasterRecoverGroupStatus,omitnil,omitempty" name:"CVMDisasterRecoverGroupStatus"`
}

type Operation struct {
	// 操作唯一id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 操作开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 操作类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 操作详情
	Detail *OperationDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 操作结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 流程任务信息
	Tasks []*TaskDetail `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 操作进度
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 回滚标记， 0未回滚 ，1回滚中，2已回滚
	RollbackTag *int64 `json:"RollbackTag,omitnil,omitempty" name:"RollbackTag"`

	// 操作者Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 自动扩容标识：0-非自动，1-自动
	AutoScaleTag *uint64 `json:"AutoScaleTag,omitnil,omitempty" name:"AutoScaleTag"`
}

type OperationDetail struct {
	// 实例原始配置信息
	OldInfo []*KeyValue `json:"OldInfo,omitnil,omitempty" name:"OldInfo"`

	// 实例更新后配置信息
	NewInfo []*KeyValue `json:"NewInfo,omitnil,omitempty" name:"NewInfo"`
}

type OperationDuration struct {
	// 维护周期，表示周一到周日，可取值[0, 6]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Periods []*uint64 `json:"Periods,omitnil,omitempty" name:"Periods"`

	// 维护开始时间
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// 维护结束时间
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 时区，以UTC形式表示
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`
}

type OperationDurationUpdated struct {
	// 维护周期，表示周一到周日，可取值[0, 6]
	Periods []*uint64 `json:"Periods,omitnil,omitempty" name:"Periods"`

	// 维护开始时间
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// 维护结束时间
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 时区，以UTC形式表示
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// ES集群ID数组
	MoreInstances []*string `json:"MoreInstances,omitnil,omitempty" name:"MoreInstances"`
}

type OptionalWebServiceInfo struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 公网url
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicUrl *string `json:"PublicUrl,omitnil,omitempty" name:"PublicUrl"`

	// 内网url
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateUrl *string `json:"PrivateUrl,omitnil,omitempty" name:"PrivateUrl"`

	// 公网访问权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// 内网访问权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateAccess *string `json:"PrivateAccess,omitnil,omitempty" name:"PrivateAccess"`

	// 版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// web服务内网vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateVip *string `json:"PrivateVip,omitnil,omitempty" name:"PrivateVip"`

	// 自定义cerebro内网url
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomPrivateUrl *string `json:"CustomPrivateUrl,omitnil,omitempty" name:"CustomPrivateUrl"`
}

type OutboundPublicAcl struct {
	// 允许节点出站访问的节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 允许节点出站访问的白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhiteHostList []*string `json:"WhiteHostList,omitnil,omitempty" name:"WhiteHostList"`
}

type ProcessDetail struct {
	// 已完成数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Completed *int64 `json:"Completed,omitnil,omitempty" name:"Completed"`

	// 剩余数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remain *int64 `json:"Remain,omitnil,omitempty" name:"Remain"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 任务类型：
	// 60：重启型任务
	// 70：分片迁移型任务
	// 80：节点变配任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

// Predefined struct for user
type RestartInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否强制重启<li>true：强制重启</li><li>false：不强制重启</li>默认false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// 重启模式：0 滚动重启； 1 全量重启
	RestartMode *int64 `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// 重启时选择是否升级内核patch版本
	UpgradeKernel *bool `json:"UpgradeKernel,omitnil,omitempty" name:"UpgradeKernel"`
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否强制重启<li>true：强制重启</li><li>false：不强制重启</li>默认false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// 重启模式：0 滚动重启； 1 全量重启
	RestartMode *int64 `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// 重启时选择是否升级内核patch版本
	UpgradeKernel *bool `json:"UpgradeKernel,omitnil,omitempty" name:"UpgradeKernel"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ForceRestart")
	delete(f, "RestartMode")
	delete(f, "UpgradeKernel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartInstanceResponseParams `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartKibanaRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type RestartKibanaRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *RestartKibanaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartKibanaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartKibanaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartKibanaResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartKibanaResponse struct {
	*tchttp.BaseResponse
	Response *RestartKibanaResponseParams `json:"Response"`
}

func (r *RestartKibanaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartKibanaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartLogstashInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重启类型，0全量重启，1滚动重启
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type RestartLogstashInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重启类型，0全量重启，1滚动重启
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *RestartLogstashInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartLogstashInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartLogstashInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartLogstashInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartLogstashInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartLogstashInstanceResponseParams `json:"Response"`
}

func (r *RestartLogstashInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartLogstashInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点名称列表
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 是否强制重启
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// 可选重启模式"in-place","blue-green"，分别表示重启，蓝绿重启；默认值为"in-place"
	RestartMode *string `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// 节点状态，在蓝绿模式中使用；离线节点蓝绿有风险
	IsOffline *bool `json:"IsOffline,omitnil,omitempty" name:"IsOffline"`

	// cvm延迟上架时间
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// 分片迁移并发数
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// 分片迁移并发速度
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// 是否开启置放群组异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组异步任务时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 事件id列表
	EventTypeIds []*string `json:"EventTypeIds,omitnil,omitempty" name:"EventTypeIds"`
}

type RestartNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点名称列表
	NodeNames []*string `json:"NodeNames,omitnil,omitempty" name:"NodeNames"`

	// 是否强制重启
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// 可选重启模式"in-place","blue-green"，分别表示重启，蓝绿重启；默认值为"in-place"
	RestartMode *string `json:"RestartMode,omitnil,omitempty" name:"RestartMode"`

	// 节点状态，在蓝绿模式中使用；离线节点蓝绿有风险
	IsOffline *bool `json:"IsOffline,omitnil,omitempty" name:"IsOffline"`

	// cvm延迟上架时间
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// 分片迁移并发数
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// 分片迁移并发速度
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// 是否开启置放群组异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组异步任务时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 事件id列表
	EventTypeIds []*string `json:"EventTypeIds,omitnil,omitempty" name:"EventTypeIds"`
}

func (r *RestartNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeNames")
	delete(f, "ForceRestart")
	delete(f, "RestartMode")
	delete(f, "IsOffline")
	delete(f, "CvmDelayOnlineTime")
	delete(f, "ShardAllocationConcurrents")
	delete(f, "ShardAllocationBytes")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	delete(f, "EventTypeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartNodesResponse struct {
	*tchttp.BaseResponse
	Response *RestartNodesResponseParams `json:"Response"`
}

func (r *RestartNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreClusterSnapshotRequestParams struct {
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 集群快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 目标集群实例Id，格式：es-xxxx，如果是恢复到本地，则和InstanceId一致
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// elastic用户名对应的密码信息
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 要在所有恢复的索引中添加或更改的设置的逗号分隔列表。使用此参数可以在恢复快照时覆盖索引设置。
	IndexSettings *string `json:"IndexSettings,omitnil,omitempty" name:"IndexSettings"`

	// 不应从快照还原的以逗号分隔的索引设置列表。
	IncludeGlobalState []*string `json:"IncludeGlobalState,omitnil,omitempty" name:"IncludeGlobalState"`

	// 需要恢复的索引名称，非必填，为空则表示恢复所有
	// 
	// 支持传多个索引名称
	Indices *string `json:"Indices,omitnil,omitempty" name:"Indices"`

	// 如果为 false，则如果快照中包含的一个或多个索引没有所有主分片可用，则整个恢复操作将失败。默认为 false,
	// 
	// 如果为 true，则允许恢复具有不可用分片的索引的部分快照。只有成功包含在快照中的分片才会被恢复。所有丢失的碎片将被重新创建为空
	Partial *string `json:"Partial,omitnil,omitempty" name:"Partial"`
}

type RestoreClusterSnapshotRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id，格式：es-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 仓库名称
	RepositoryName *string `json:"RepositoryName,omitnil,omitempty" name:"RepositoryName"`

	// 集群快照名称
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 目标集群实例Id，格式：es-xxxx，如果是恢复到本地，则和InstanceId一致
	TargetInstanceId *string `json:"TargetInstanceId,omitnil,omitempty" name:"TargetInstanceId"`

	// elastic用户名对应的密码信息
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 要在所有恢复的索引中添加或更改的设置的逗号分隔列表。使用此参数可以在恢复快照时覆盖索引设置。
	IndexSettings *string `json:"IndexSettings,omitnil,omitempty" name:"IndexSettings"`

	// 不应从快照还原的以逗号分隔的索引设置列表。
	IncludeGlobalState []*string `json:"IncludeGlobalState,omitnil,omitempty" name:"IncludeGlobalState"`

	// 需要恢复的索引名称，非必填，为空则表示恢复所有
	// 
	// 支持传多个索引名称
	Indices *string `json:"Indices,omitnil,omitempty" name:"Indices"`

	// 如果为 false，则如果快照中包含的一个或多个索引没有所有主分片可用，则整个恢复操作将失败。默认为 false,
	// 
	// 如果为 true，则允许恢复具有不可用分片的索引的部分快照。只有成功包含在快照中的分片才会被恢复。所有丢失的碎片将被重新创建为空
	Partial *string `json:"Partial,omitnil,omitempty" name:"Partial"`
}

func (r *RestoreClusterSnapshotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreClusterSnapshotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RepositoryName")
	delete(f, "SnapshotName")
	delete(f, "TargetInstanceId")
	delete(f, "Password")
	delete(f, "IndexSettings")
	delete(f, "IncludeGlobalState")
	delete(f, "Indices")
	delete(f, "Partial")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreClusterSnapshotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreClusterSnapshotResponseParams struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestoreClusterSnapshotResponse struct {
	*tchttp.BaseResponse
	Response *RestoreClusterSnapshotResponseParams `json:"Response"`
}

func (r *RestoreClusterSnapshotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreClusterSnapshotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveAndDeployLogstashPipelineRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例管道信息
	Pipeline *LogstashPipeline `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// 操作类型<li>1：只保存</li><li>2：保存并部署</li>
	OpType *uint64 `json:"OpType,omitnil,omitempty" name:"OpType"`
}

type SaveAndDeployLogstashPipelineRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例管道信息
	Pipeline *LogstashPipeline `json:"Pipeline,omitnil,omitempty" name:"Pipeline"`

	// 操作类型<li>1：只保存</li><li>2：保存并部署</li>
	OpType *uint64 `json:"OpType,omitnil,omitempty" name:"OpType"`
}

func (r *SaveAndDeployLogstashPipelineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveAndDeployLogstashPipelineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Pipeline")
	delete(f, "OpType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveAndDeployLogstashPipelineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveAndDeployLogstashPipelineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveAndDeployLogstashPipelineResponse struct {
	*tchttp.BaseResponse
	Response *SaveAndDeployLogstashPipelineResponseParams `json:"Response"`
}

func (r *SaveAndDeployLogstashPipelineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveAndDeployLogstashPipelineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerlessDi struct {
	// 数据链路采集源类型，如cvm_collector、tke_collector
	DiSourceType *string `json:"DiSourceType,omitnil,omitempty" name:"DiSourceType"`

	// cvm数据源
	DiSourceCvm *DiSourceCvm `json:"DiSourceCvm,omitnil,omitempty" name:"DiSourceCvm"`

	// tke数据源
	DiSourceTke *DiSourceTke `json:"DiSourceTke,omitnil,omitempty" name:"DiSourceTke"`
}

type ServerlessIndexMetaField struct {
	// 索引所属集群APP ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 索引名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 索引元数据JSON
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexMetaJson *string `json:"IndexMetaJson,omitnil,omitempty" name:"IndexMetaJson"`

	// 索引文档数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexDocs *int64 `json:"IndexDocs,omitnil,omitempty" name:"IndexDocs"`

	// 索引存储大小，单位Byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStorage *int64 `json:"IndexStorage,omitnil,omitempty" name:"IndexStorage"`

	// 索引创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexCreateTime *string `json:"IndexCreateTime,omitnil,omitempty" name:"IndexCreateTime"`

	// 索引实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 索引自治字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexOptionsField *ServerlessIndexOptionsField `json:"IndexOptionsField,omitnil,omitempty" name:"IndexOptionsField"`

	// 索引配置字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexSettingsField *ServerlessIndexSettingsField `json:"IndexSettingsField,omitnil,omitempty" name:"IndexSettingsField"`

	// 索引所属连接相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexNetworkField *ServerlessIndexNetworkField `json:"IndexNetworkField,omitnil,omitempty" name:"IndexNetworkField"`

	// Kibana公网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaUrl *string `json:"KibanaUrl,omitnil,omitempty" name:"KibanaUrl"`

	// Kibana内网域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitnil,omitempty" name:"KibanaPrivateUrl"`

	// 索引内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexAccessUrl *string `json:"IndexAccessUrl,omitnil,omitempty" name:"IndexAccessUrl"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 索引空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 索引空间名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// 存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *int64 `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 标签信息
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 索引流量，单位byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexTraffic *float64 `json:"IndexTraffic,omitnil,omitempty" name:"IndexTraffic"`
}

type ServerlessIndexNetworkField struct {
	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// vpc唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcUid *string `json:"VpcUid,omitnil,omitempty" name:"VpcUid"`

	// 子网唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetUid *string `json:"SubnetUid,omitnil,omitempty" name:"SubnetUid"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type ServerlessIndexOptionsField struct {
	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireMaxAge *string `json:"ExpireMaxAge,omitnil,omitempty" name:"ExpireMaxAge"`

	// 时间分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimestampField *string `json:"TimestampField,omitnil,omitempty" name:"TimestampField"`

	// 标准存储时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	SinkCycleAge *string `json:"SinkCycleAge,omitnil,omitempty" name:"SinkCycleAge"`

	// 标准存储时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardStorageAge *string `json:"StandardStorageAge,omitnil,omitempty" name:"StandardStorageAge"`
}

type ServerlessIndexSettingsField struct {
	// 索引主分片数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfShards *string `json:"NumberOfShards,omitnil,omitempty" name:"NumberOfShards"`

	// 索引刷新频率
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefreshInterval *string `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`

	// 自定义参数
	CustomSetting *string `json:"CustomSetting,omitnil,omitempty" name:"CustomSetting"`
}

type ServerlessSpace struct {
	// Serverless索引空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Serverless索引空间名
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// Serverless索引空间状态，0正常，-1已删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建日期
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 空间内索引数量
	IndexCount *int64 `json:"IndexCount,omitnil,omitempty" name:"IndexCount"`

	// kibana公网uri
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaUrl *string `json:"KibanaUrl,omitnil,omitempty" name:"KibanaUrl"`

	// kibana内网url
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPrivateUrl *string `json:"KibanaPrivateUrl,omitnil,omitempty" name:"KibanaPrivateUrl"`

	// 空间内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexAccessUrl *string `json:"IndexAccessUrl,omitnil,omitempty" name:"IndexAccessUrl"`

	// 空间白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaPublicAcl *EsAcl `json:"KibanaPublicAcl,omitnil,omitempty" name:"KibanaPublicAcl"`

	// 空间检索分析域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaEmbedUrl *string `json:"KibanaEmbedUrl,omitnil,omitempty" name:"KibanaEmbedUrl"`

	// 数据联路
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiDataList *DiData `json:"DiDataList,omitnil,omitempty" name:"DiDataList"`

	// 空间vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// kibana公网开关，0关闭，1开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableKibanaPublicAccess *int64 `json:"EnableKibanaPublicAccess,omitnil,omitempty" name:"EnableKibanaPublicAccess"`

	// kibana内网开关，0关闭，1开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableKibanaPrivateAccess *int64 `json:"EnableKibanaPrivateAccess,omitnil,omitempty" name:"EnableKibanaPrivateAccess"`

	// 空间所属appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// //默认en， 可选zh-CN
	// 注意：此字段可能返回 null，表示取不到有效值。
	KibanaLanguage *string `json:"KibanaLanguage,omitnil,omitempty" name:"KibanaLanguage"`

	// 0
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *int64 `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 空间标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type ServerlessSpaceUser struct {
	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 用户密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 有权限的索引数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndicesScope []*string `json:"IndicesScope,omitnil,omitempty" name:"IndicesScope"`

	// 权限类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeType *int64 `json:"PrivilegeType,omitnil,omitempty" name:"PrivilegeType"`
}

type SettingDetail struct {
	// 配置key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 配置当前值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 配置处理建议
	Advise *string `json:"Advise,omitnil,omitempty" name:"Advise"`
}

type Snapshots struct {
	// 快照名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotName *string `json:"SnapshotName,omitnil,omitempty" name:"SnapshotName"`

	// 快照Uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 仓库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repository *string `json:"Repository,omitnil,omitempty" name:"Repository"`

	// 该快照所属集群的版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 备份的索引列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Indices []*string `json:"Indices,omitnil,omitempty" name:"Indices"`

	// 备份的datastream列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataStreams []*string `json:"DataStreams,omitnil,omitempty" name:"DataStreams"`

	// 备份的状态
	// 
	// FAILED            备份失败
	// 
	// IN_PROGRESS 备份执行中
	// 
	// PARTIAL          备份部分成功，部分失败，备份失败的索引和原因会在Failures字段中展示
	// 
	// SUCCESS     备份成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 快照备份的开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 快照备份的结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 快照备份的耗时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationInMillis *int64 `json:"DurationInMillis,omitnil,omitempty" name:"DurationInMillis"`

	// 备份的总分片数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalShards *int64 `json:"TotalShards,omitnil,omitempty" name:"TotalShards"`

	// 备份失败的分片数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedShards *int64 `json:"FailedShards,omitnil,omitempty" name:"FailedShards"`

	// 备份成功的分片数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessfulShards *int64 `json:"SuccessfulShards,omitnil,omitempty" name:"SuccessfulShards"`

	// 备份失败的索引分片和失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Failures []*Failures `json:"Failures,omitnil,omitempty" name:"Failures"`

	// 是否用户备份
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserBackUp *string `json:"UserBackUp,omitnil,omitempty" name:"UserBackUp"`
}

// Predefined struct for user
type StartLogstashPipelinesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 管道ID列表
	PipelineIds []*string `json:"PipelineIds,omitnil,omitempty" name:"PipelineIds"`
}

type StartLogstashPipelinesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 管道ID列表
	PipelineIds []*string `json:"PipelineIds,omitnil,omitempty" name:"PipelineIds"`
}

func (r *StartLogstashPipelinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLogstashPipelinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PipelineIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartLogstashPipelinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartLogstashPipelinesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartLogstashPipelinesResponse struct {
	*tchttp.BaseResponse
	Response *StartLogstashPipelinesResponseParams `json:"Response"`
}

func (r *StartLogstashPipelinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLogstashPipelinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLogstashPipelinesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 管道ID列表
	PipelineIds []*string `json:"PipelineIds,omitnil,omitempty" name:"PipelineIds"`
}

type StopLogstashPipelinesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 管道ID列表
	PipelineIds []*string `json:"PipelineIds,omitnil,omitempty" name:"PipelineIds"`
}

func (r *StopLogstashPipelinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLogstashPipelinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PipelineIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLogstashPipelinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLogstashPipelinesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopLogstashPipelinesResponse struct {
	*tchttp.BaseResponse
	Response *StopLogstashPipelinesResponseParams `json:"Response"`
}

func (r *StopLogstashPipelinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLogstashPipelinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubTaskDetail struct {
	// 子任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子任务结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 子任务错误信息
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 子任务类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 子任务状态，0处理中 1成功 -1失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 升级检查失败的索引名
	FailedIndices []*string `json:"FailedIndices,omitnil,omitempty" name:"FailedIndices"`

	// 子任务结束时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 子任务等级，1警告 2失败
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskDetail struct {
	// 任务名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务进度
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务完成时间
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 子任务
	SubTasks []*SubTaskDetail `json:"SubTasks,omitnil,omitempty" name:"SubTasks"`

	// 任务花费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElapsedTime *int64 `json:"ElapsedTime,omitnil,omitempty" name:"ElapsedTime"`

	// 任务进度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessInfo *ProcessDetail `json:"ProcessInfo,omitnil,omitempty" name:"ProcessInfo"`
}

// Predefined struct for user
type UpdateDiagnoseSettingsRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0：开启智能运维；-1：关闭智能运维
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 智能运维每天定时巡检时间，时间格式为HH:00:00，例如15:00:00
	CronTime *string `json:"CronTime,omitnil,omitempty" name:"CronTime"`
}

type UpdateDiagnoseSettingsRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0：开启智能运维；-1：关闭智能运维
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 智能运维每天定时巡检时间，时间格式为HH:00:00，例如15:00:00
	CronTime *string `json:"CronTime,omitnil,omitempty" name:"CronTime"`
}

func (r *UpdateDiagnoseSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDiagnoseSettingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Status")
	delete(f, "CronTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDiagnoseSettingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDiagnoseSettingsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDiagnoseSettingsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDiagnoseSettingsResponseParams `json:"Response"`
}

func (r *UpdateDiagnoseSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDiagnoseSettingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDictionariesRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 安装时填IK分词主词典COS地址，删除时填词典名如test.dic
	IkMainDicts []*string `json:"IkMainDicts,omitnil,omitempty" name:"IkMainDicts"`

	// 安装时填IK分词停用词词典COS地址，删除时填词典名如test.dic
	IkStopwords []*string `json:"IkStopwords,omitnil,omitempty" name:"IkStopwords"`

	// 安装时填同义词词典COS地址，删除时填词典名如test.dic
	Synonym []*string `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// 安装时填QQ分词词典COS地址，删除时填词典名如test.dic
	QQDict []*string `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// 0：安装；1：删除。默认值0
	UpdateType *int64 `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// 是否强制重启集群。默认值false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpdateDictionariesRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 安装时填IK分词主词典COS地址，删除时填词典名如test.dic
	IkMainDicts []*string `json:"IkMainDicts,omitnil,omitempty" name:"IkMainDicts"`

	// 安装时填IK分词停用词词典COS地址，删除时填词典名如test.dic
	IkStopwords []*string `json:"IkStopwords,omitnil,omitempty" name:"IkStopwords"`

	// 安装时填同义词词典COS地址，删除时填词典名如test.dic
	Synonym []*string `json:"Synonym,omitnil,omitempty" name:"Synonym"`

	// 安装时填QQ分词词典COS地址，删除时填词典名如test.dic
	QQDict []*string `json:"QQDict,omitnil,omitempty" name:"QQDict"`

	// 0：安装；1：删除。默认值0
	UpdateType *int64 `json:"UpdateType,omitnil,omitempty" name:"UpdateType"`

	// 是否强制重启集群。默认值false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

func (r *UpdateDictionariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDictionariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IkMainDicts")
	delete(f, "IkStopwords")
	delete(f, "Synonym")
	delete(f, "QQDict")
	delete(f, "UpdateType")
	delete(f, "ForceRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDictionariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDictionariesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDictionariesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDictionariesResponseParams `json:"Response"`
}

func (r *UpdateDictionariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDictionariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIndexRequestParams struct {
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 更新的索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 更新的索引名
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 更新的索引元数据JSON，如mappings、settings
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否滚动后备索引
	RolloverBackingIndex *bool `json:"RolloverBackingIndex,omitnil,omitempty" name:"RolloverBackingIndex"`
}

type UpdateIndexRequest struct {
	*tchttp.BaseRequest
	
	// ES集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 更新的索引类型。auto：自治索引；normal：普通索引
	IndexType *string `json:"IndexType,omitnil,omitempty" name:"IndexType"`

	// 更新的索引名
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 更新的索引元数据JSON，如mappings、settings
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// 集群访问用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 集群访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否滚动后备索引
	RolloverBackingIndex *bool `json:"RolloverBackingIndex,omitnil,omitempty" name:"RolloverBackingIndex"`
}

func (r *UpdateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IndexType")
	delete(f, "IndexName")
	delete(f, "UpdateMetaJson")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "RolloverBackingIndex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIndexResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateIndexResponse struct {
	*tchttp.BaseResponse
	Response *UpdateIndexResponseParams `json:"Response"`
}

func (r *UpdateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 已废弃请使用NodeInfoList
	// 节点个数（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// ES配置项（JSON格式字符串）
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// 默认用户elastic的密码（8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 可视化组件（Kibana、Cerebro）的公网访问策略
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// 已废弃请使用NodeInfoList
	// 磁盘大小（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 已废弃请使用NodeInfoList
	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点个数（只支持3个或5个）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// 已废弃请使用NodeInfoList
	// 专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点磁盘大小（单位GB系统默认配置为50GB,暂不支持自定义）
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// 更新配置时是否强制重启<li>true强制重启</li><li>false不强制重启</li>当前仅更新EsConfig时需要设置，默认值为false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// COS自动备份信息
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// 节点信息列表，可以只传递要更新的节点及其对应的规格信息。支持的操作包括<li>修改一种节点的个数</li><li>修改一种节点的节点规格及磁盘大小</li><li>增加一种节点类型（需要同时指定该节点的类型，个数，规格，磁盘等信息）</li>上述操作一次只能进行一种，且磁盘类型不支持修改
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// ES集群公网访问状态
	// OPEN 开启
	// CLOSE 关闭
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// 公网访问控制列表
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// Kibana公网访问状态
	// OPEN 开启
	// CLOSE 关闭
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态
	// OPEN 开启
	// CLOSE 关闭
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// ES 6.8及以上版本基础版开启或关闭用户认证
	BasicSecurityType *int64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Kibana内网端口
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitnil,omitempty" name:"KibanaPrivatePort"`

	// 0: 蓝绿变更方式扩容，集群不重启 （默认） 1: 磁盘解挂载扩容，集群滚动重启
	ScaleType *int64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// 多可用区部署
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// 场景化模板类型 -1：不启用 1：通用 2：日志 3：搜索
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana配置项（JSON格式字符串）
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// 可视化节点配置
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// 切换到新网络架构
	SwitchPrivateLink *string `json:"SwitchPrivateLink,omitnil,omitempty" name:"SwitchPrivateLink"`

	// 启用Cerebro
	EnableCerebro *bool `json:"EnableCerebro,omitnil,omitempty" name:"EnableCerebro"`

	// Cerebro公网访问状态
	// OPEN 开启
	// CLOSE 关闭
	CerebroPublicAccess *string `json:"CerebroPublicAccess,omitnil,omitempty" name:"CerebroPublicAccess"`

	// Cerebro内网访问状态
	// OPEN 开启
	// CLOSE 关闭
	CerebroPrivateAccess *string `json:"CerebroPrivateAccess,omitnil,omitempty" name:"CerebroPrivateAccess"`

	// 新增或修改的配置组信息
	EsConfigSet *EsConfigSetInfo `json:"EsConfigSet,omitnil,omitempty" name:"EsConfigSet"`

	// 可维护时间段
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// 是否开启Alerting 外网告警输出：
	// OPEN 开启
	// CLOSE 关闭
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`

	// kibana内网自定义域名
	KibanaPrivateDomain *string `json:"KibanaPrivateDomain,omitnil,omitempty" name:"KibanaPrivateDomain"`

	// cerebro内网自定义域名
	CerebroPrivateDomain *string `json:"CerebroPrivateDomain,omitnil,omitempty" name:"CerebroPrivateDomain"`

	// 变更为https集群，默认是http
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 节点出站访问详细信息
	OutboundPublicAcls []*OutboundPublicAcl `json:"OutboundPublicAcls,omitnil,omitempty" name:"OutboundPublicAcls"`

	// 节点出站访问操作
	// OPEN 开启
	// CLOSE 关闭
	OutboundPublicAccess *string `json:"OutboundPublicAccess,omitnil,omitempty" name:"OutboundPublicAccess"`

	// cvm延迟上架参数
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// 分片迁移并发数
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// 分片迁移并发速度
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// 读写分离模式：-1-不开启，1-本地读写分离，2-远端读写分离
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// 是否开启置放群组异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组异步任务可维护时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 开启集群保护：OPEN-开启，CLOSE-关闭
	EnableDestroyProtection *string `json:"EnableDestroyProtection,omitnil,omitempty" name:"EnableDestroyProtection"`

	// 自动扩盘参数
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// 自动扩盘删除参数
	AutoScaleDiskDeleteNodeTypeList []*string `json:"AutoScaleDiskDeleteNodeTypeList,omitnil,omitempty" name:"AutoScaleDiskDeleteNodeTypeList"`
}

type UpdateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称（1-50 个英文、汉字、数字、连接线-或下划线_）
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 已废弃请使用NodeInfoList
	// 节点个数（2-50个）
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// ES配置项（JSON格式字符串）
	EsConfig *string `json:"EsConfig,omitnil,omitempty" name:"EsConfig"`

	// 默认用户elastic的密码（8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号）
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 可视化组件（Kibana、Cerebro）的公网访问策略
	EsAcl *EsAcl `json:"EsAcl,omitnil,omitempty" name:"EsAcl"`

	// 已废弃请使用NodeInfoList
	// 磁盘大小（单位GB）
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 已废弃请使用NodeInfoList
	// 节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点个数（只支持3个或5个）
	MasterNodeNum *uint64 `json:"MasterNodeNum,omitnil,omitempty" name:"MasterNodeNum"`

	// 已废弃请使用NodeInfoList
	// 专用主节点规格<li>ES.S1.SMALL2：1核2G</li><li>ES.S1.MEDIUM4：2核4G</li><li>ES.S1.MEDIUM8：2核8G</li><li>ES.S1.LARGE16：4核16G</li><li>ES.S1.2XLARGE32：8核32G</li><li>ES.S1.4XLARGE32：16核32G</li><li>ES.S1.4XLARGE64：16核64G</li>
	MasterNodeType *string `json:"MasterNodeType,omitnil,omitempty" name:"MasterNodeType"`

	// 已废弃请使用NodeInfoList
	// 专用主节点磁盘大小（单位GB系统默认配置为50GB,暂不支持自定义）
	MasterNodeDiskSize *uint64 `json:"MasterNodeDiskSize,omitnil,omitempty" name:"MasterNodeDiskSize"`

	// 更新配置时是否强制重启<li>true强制重启</li><li>false不强制重启</li>当前仅更新EsConfig时需要设置，默认值为false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// COS自动备份信息
	CosBackup *CosBackup `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// 节点信息列表，可以只传递要更新的节点及其对应的规格信息。支持的操作包括<li>修改一种节点的个数</li><li>修改一种节点的节点规格及磁盘大小</li><li>增加一种节点类型（需要同时指定该节点的类型，个数，规格，磁盘等信息）</li>上述操作一次只能进行一种，且磁盘类型不支持修改
	NodeInfoList []*NodeInfo `json:"NodeInfoList,omitnil,omitempty" name:"NodeInfoList"`

	// ES集群公网访问状态
	// OPEN 开启
	// CLOSE 关闭
	PublicAccess *string `json:"PublicAccess,omitnil,omitempty" name:"PublicAccess"`

	// 公网访问控制列表
	EsPublicAcl *EsPublicAcl `json:"EsPublicAcl,omitnil,omitempty" name:"EsPublicAcl"`

	// Kibana公网访问状态
	// OPEN 开启
	// CLOSE 关闭
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// Kibana内网访问状态
	// OPEN 开启
	// CLOSE 关闭
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// ES 6.8及以上版本基础版开启或关闭用户认证
	BasicSecurityType *int64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// Kibana内网端口
	KibanaPrivatePort *uint64 `json:"KibanaPrivatePort,omitnil,omitempty" name:"KibanaPrivatePort"`

	// 0: 蓝绿变更方式扩容，集群不重启 （默认） 1: 磁盘解挂载扩容，集群滚动重启
	ScaleType *int64 `json:"ScaleType,omitnil,omitempty" name:"ScaleType"`

	// 多可用区部署
	MultiZoneInfo []*ZoneDetail `json:"MultiZoneInfo,omitnil,omitempty" name:"MultiZoneInfo"`

	// 场景化模板类型 -1：不启用 1：通用 2：日志 3：搜索
	SceneType *int64 `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// Kibana配置项（JSON格式字符串）
	KibanaConfig *string `json:"KibanaConfig,omitnil,omitempty" name:"KibanaConfig"`

	// 可视化节点配置
	WebNodeTypeInfo *WebNodeTypeInfo `json:"WebNodeTypeInfo,omitnil,omitempty" name:"WebNodeTypeInfo"`

	// 切换到新网络架构
	SwitchPrivateLink *string `json:"SwitchPrivateLink,omitnil,omitempty" name:"SwitchPrivateLink"`

	// 启用Cerebro
	EnableCerebro *bool `json:"EnableCerebro,omitnil,omitempty" name:"EnableCerebro"`

	// Cerebro公网访问状态
	// OPEN 开启
	// CLOSE 关闭
	CerebroPublicAccess *string `json:"CerebroPublicAccess,omitnil,omitempty" name:"CerebroPublicAccess"`

	// Cerebro内网访问状态
	// OPEN 开启
	// CLOSE 关闭
	CerebroPrivateAccess *string `json:"CerebroPrivateAccess,omitnil,omitempty" name:"CerebroPrivateAccess"`

	// 新增或修改的配置组信息
	EsConfigSet *EsConfigSetInfo `json:"EsConfigSet,omitnil,omitempty" name:"EsConfigSet"`

	// 可维护时间段
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`

	// 是否开启Alerting 外网告警输出：
	// OPEN 开启
	// CLOSE 关闭
	KibanaAlteringPublicAccess *string `json:"KibanaAlteringPublicAccess,omitnil,omitempty" name:"KibanaAlteringPublicAccess"`

	// kibana内网自定义域名
	KibanaPrivateDomain *string `json:"KibanaPrivateDomain,omitnil,omitempty" name:"KibanaPrivateDomain"`

	// cerebro内网自定义域名
	CerebroPrivateDomain *string `json:"CerebroPrivateDomain,omitnil,omitempty" name:"CerebroPrivateDomain"`

	// 变更为https集群，默认是http
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 节点出站访问详细信息
	OutboundPublicAcls []*OutboundPublicAcl `json:"OutboundPublicAcls,omitnil,omitempty" name:"OutboundPublicAcls"`

	// 节点出站访问操作
	// OPEN 开启
	// CLOSE 关闭
	OutboundPublicAccess *string `json:"OutboundPublicAccess,omitnil,omitempty" name:"OutboundPublicAccess"`

	// cvm延迟上架参数
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// 分片迁移并发数
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// 分片迁移并发速度
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// 读写分离模式：-1-不开启，1-本地读写分离，2-远端读写分离
	ReadWriteMode *int64 `json:"ReadWriteMode,omitnil,omitempty" name:"ReadWriteMode"`

	// 是否开启置放群组异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组异步任务可维护时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`

	// 开启集群保护：OPEN-开启，CLOSE-关闭
	EnableDestroyProtection *string `json:"EnableDestroyProtection,omitnil,omitempty" name:"EnableDestroyProtection"`

	// 自动扩盘参数
	AutoScaleDiskInfoList []*AutoScaleDiskInfo `json:"AutoScaleDiskInfoList,omitnil,omitempty" name:"AutoScaleDiskInfoList"`

	// 自动扩盘删除参数
	AutoScaleDiskDeleteNodeTypeList []*string `json:"AutoScaleDiskDeleteNodeTypeList,omitnil,omitempty" name:"AutoScaleDiskDeleteNodeTypeList"`
}

func (r *UpdateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "NodeNum")
	delete(f, "EsConfig")
	delete(f, "Password")
	delete(f, "EsAcl")
	delete(f, "DiskSize")
	delete(f, "NodeType")
	delete(f, "MasterNodeNum")
	delete(f, "MasterNodeType")
	delete(f, "MasterNodeDiskSize")
	delete(f, "ForceRestart")
	delete(f, "CosBackup")
	delete(f, "NodeInfoList")
	delete(f, "PublicAccess")
	delete(f, "EsPublicAcl")
	delete(f, "KibanaPublicAccess")
	delete(f, "KibanaPrivateAccess")
	delete(f, "BasicSecurityType")
	delete(f, "KibanaPrivatePort")
	delete(f, "ScaleType")
	delete(f, "MultiZoneInfo")
	delete(f, "SceneType")
	delete(f, "KibanaConfig")
	delete(f, "WebNodeTypeInfo")
	delete(f, "SwitchPrivateLink")
	delete(f, "EnableCerebro")
	delete(f, "CerebroPublicAccess")
	delete(f, "CerebroPrivateAccess")
	delete(f, "EsConfigSet")
	delete(f, "OperationDuration")
	delete(f, "KibanaAlteringPublicAccess")
	delete(f, "KibanaPrivateDomain")
	delete(f, "CerebroPrivateDomain")
	delete(f, "Protocol")
	delete(f, "OutboundPublicAcls")
	delete(f, "OutboundPublicAccess")
	delete(f, "CvmDelayOnlineTime")
	delete(f, "ShardAllocationConcurrents")
	delete(f, "ShardAllocationBytes")
	delete(f, "ReadWriteMode")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	delete(f, "EnableDestroyProtection")
	delete(f, "AutoScaleDiskInfoList")
	delete(f, "AutoScaleDiskDeleteNodeTypeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInstanceResponseParams struct {
	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateInstanceResponseParams `json:"Response"`
}

func (r *UpdateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateJdkRequestParams struct {
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Jdk类型，支持kona和oracle
	Jdk *string `json:"Jdk,omitnil,omitempty" name:"Jdk"`

	// Gc类型，支持g1和cms
	Gc *string `json:"Gc,omitnil,omitempty" name:"Gc"`

	// 是否强制重启
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpdateJdkRequest struct {
	*tchttp.BaseRequest
	
	// ES实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Jdk类型，支持kona和oracle
	Jdk *string `json:"Jdk,omitnil,omitempty" name:"Jdk"`

	// Gc类型，支持g1和cms
	Gc *string `json:"Gc,omitnil,omitempty" name:"Gc"`

	// 是否强制重启
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

func (r *UpdateJdkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateJdkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Jdk")
	delete(f, "Gc")
	delete(f, "ForceRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateJdkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateJdkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateJdkResponse struct {
	*tchttp.BaseResponse
	Response *UpdateJdkResponseParams `json:"Response"`
}

func (r *UpdateJdkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateJdkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLogstashInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例节点数量
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例YML配置
	YMLConfig *string `json:"YMLConfig,omitnil,omitempty" name:"YMLConfig"`

	// 实例绑定的ES集群信息
	BindedES *LogstashBindedES `json:"BindedES,omitnil,omitempty" name:"BindedES"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 扩展文件列表
	ExtendedFiles []*LogstashExtendedFile `json:"ExtendedFiles,omitnil,omitempty" name:"ExtendedFiles"`

	// 实例规格
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点磁盘容量
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 可维护时间段
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`
}

type UpdateLogstashInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例节点数量
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 实例YML配置
	YMLConfig *string `json:"YMLConfig,omitnil,omitempty" name:"YMLConfig"`

	// 实例绑定的ES集群信息
	BindedES *LogstashBindedES `json:"BindedES,omitnil,omitempty" name:"BindedES"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 扩展文件列表
	ExtendedFiles []*LogstashExtendedFile `json:"ExtendedFiles,omitnil,omitempty" name:"ExtendedFiles"`

	// 实例规格
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点磁盘容量
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 可维护时间段
	OperationDuration *OperationDurationUpdated `json:"OperationDuration,omitnil,omitempty" name:"OperationDuration"`
}

func (r *UpdateLogstashInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLogstashInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NodeNum")
	delete(f, "YMLConfig")
	delete(f, "BindedES")
	delete(f, "InstanceName")
	delete(f, "ExtendedFiles")
	delete(f, "NodeType")
	delete(f, "DiskSize")
	delete(f, "OperationDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateLogstashInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLogstashInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateLogstashInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateLogstashInstanceResponseParams `json:"Response"`
}

func (r *UpdateLogstashInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLogstashInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLogstashPipelineDescRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例管道ID
	PipelineId *string `json:"PipelineId,omitnil,omitempty" name:"PipelineId"`

	// 管道描述信息
	PipelineDesc *string `json:"PipelineDesc,omitnil,omitempty" name:"PipelineDesc"`
}

type UpdateLogstashPipelineDescRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例管道ID
	PipelineId *string `json:"PipelineId,omitnil,omitempty" name:"PipelineId"`

	// 管道描述信息
	PipelineDesc *string `json:"PipelineDesc,omitnil,omitempty" name:"PipelineDesc"`
}

func (r *UpdateLogstashPipelineDescRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLogstashPipelineDescRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PipelineId")
	delete(f, "PipelineDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateLogstashPipelineDescRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateLogstashPipelineDescResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateLogstashPipelineDescResponse struct {
	*tchttp.BaseResponse
	Response *UpdateLogstashPipelineDescResponseParams `json:"Response"`
}

func (r *UpdateLogstashPipelineDescResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateLogstashPipelineDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePluginsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要安装的插件名列表
	InstallPluginList []*string `json:"InstallPluginList,omitnil,omitempty" name:"InstallPluginList"`

	// 需要卸载的插件名列表
	RemovePluginList []*string `json:"RemovePluginList,omitnil,omitempty" name:"RemovePluginList"`

	// 是否强制重启，默认值false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// 是否重新安装，默认值false
	ForceUpdate *bool `json:"ForceUpdate,omitnil,omitempty" name:"ForceUpdate"`

	// 0：系统插件
	PluginType *uint64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

type UpdatePluginsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要安装的插件名列表
	InstallPluginList []*string `json:"InstallPluginList,omitnil,omitempty" name:"InstallPluginList"`

	// 需要卸载的插件名列表
	RemovePluginList []*string `json:"RemovePluginList,omitnil,omitempty" name:"RemovePluginList"`

	// 是否强制重启，默认值false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`

	// 是否重新安装，默认值false
	ForceUpdate *bool `json:"ForceUpdate,omitnil,omitempty" name:"ForceUpdate"`

	// 0：系统插件
	PluginType *uint64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`
}

func (r *UpdatePluginsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePluginsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstallPluginList")
	delete(f, "RemovePluginList")
	delete(f, "ForceRestart")
	delete(f, "ForceUpdate")
	delete(f, "PluginType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePluginsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePluginsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePluginsResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePluginsResponseParams `json:"Response"`
}

func (r *UpdatePluginsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePluginsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRequestTargetNodeTypesRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接收请求的目标节点类型列表
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`
}

type UpdateRequestTargetNodeTypesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接收请求的目标节点类型列表
	TargetNodeTypes []*string `json:"TargetNodeTypes,omitnil,omitempty" name:"TargetNodeTypes"`
}

func (r *UpdateRequestTargetNodeTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRequestTargetNodeTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TargetNodeTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRequestTargetNodeTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRequestTargetNodeTypesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRequestTargetNodeTypesResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRequestTargetNodeTypesResponseParams `json:"Response"`
}

func (r *UpdateRequestTargetNodeTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRequestTargetNodeTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServerlessInstanceRequestParams struct {
	// Serveless实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 更新的索引元数据JSON，如mappings、settings
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// 实例访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 开启kibana内网访问，如OPEN、CLOSE
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// 开启kibana外网访问，如OPEN、CLOSE
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// 访问控制列表
	KibanaPublicAcl *KibanaPublicAcl `json:"KibanaPublicAcl,omitnil,omitempty" name:"KibanaPublicAcl"`
}

type UpdateServerlessInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Serveless实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 更新的索引元数据JSON，如mappings、settings
	UpdateMetaJson *string `json:"UpdateMetaJson,omitnil,omitempty" name:"UpdateMetaJson"`

	// 实例访问密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 开启kibana内网访问，如OPEN、CLOSE
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// 开启kibana外网访问，如OPEN、CLOSE
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// 访问控制列表
	KibanaPublicAcl *KibanaPublicAcl `json:"KibanaPublicAcl,omitnil,omitempty" name:"KibanaPublicAcl"`
}

func (r *UpdateServerlessInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServerlessInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UpdateMetaJson")
	delete(f, "Password")
	delete(f, "KibanaPrivateAccess")
	delete(f, "KibanaPublicAccess")
	delete(f, "KibanaPublicAcl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateServerlessInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServerlessInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateServerlessInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateServerlessInstanceResponseParams `json:"Response"`
}

func (r *UpdateServerlessInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServerlessInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServerlessSpaceRequestParams struct {
	// Serveless索引空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Serveless索引空间名
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// kibana内网开关
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// kibana公网开关
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// kibana公网白名单
	KibanaPublicAcl *EsAcl `json:"KibanaPublicAcl,omitnil,omitempty" name:"KibanaPublicAcl"`
}

type UpdateServerlessSpaceRequest struct {
	*tchttp.BaseRequest
	
	// Serveless索引空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// Serveless索引空间名
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// kibana内网开关
	KibanaPrivateAccess *string `json:"KibanaPrivateAccess,omitnil,omitempty" name:"KibanaPrivateAccess"`

	// kibana公网开关
	KibanaPublicAccess *string `json:"KibanaPublicAccess,omitnil,omitempty" name:"KibanaPublicAccess"`

	// kibana公网白名单
	KibanaPublicAcl *EsAcl `json:"KibanaPublicAcl,omitnil,omitempty" name:"KibanaPublicAcl"`
}

func (r *UpdateServerlessSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServerlessSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "SpaceName")
	delete(f, "KibanaPrivateAccess")
	delete(f, "KibanaPublicAccess")
	delete(f, "KibanaPublicAcl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateServerlessSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateServerlessSpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateServerlessSpaceResponse struct {
	*tchttp.BaseResponse
	Response *UpdateServerlessSpaceResponseParams `json:"Response"`
}

func (r *UpdateServerlessSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateServerlessSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标ES版本，支持：”6.4.3“, "6.8.2"，"7.5.1", "7.10.1", "7.14.2"
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// 是否只做升级检查，默认值为false
	CheckOnly *bool `json:"CheckOnly,omitnil,omitempty" name:"CheckOnly"`

	// 目标商业特性版本：<li>oss 开源版</li><li>basic 基础版</li>当前仅在5.6.4升级6.x版本时使用，默认值为basic
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// 升级方式：<li>scale 蓝绿变更</li><li>restart 滚动重启</li>默认值为scale
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// 升级版本前是否对集群进行备份，默认不备份
	CosBackup *bool `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// 滚动模式时，是否跳过检查，进行强制重启。默认值为false
	SkipCheckForceRestart *bool `json:"SkipCheckForceRestart,omitnil,omitempty" name:"SkipCheckForceRestart"`

	// cvm延迟上架参数
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// 分片迁移并发数
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// 分片迁移并发速度
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// 是否开启置放群组异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组异步任务时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`
}

type UpgradeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 目标ES版本，支持：”6.4.3“, "6.8.2"，"7.5.1", "7.10.1", "7.14.2"
	EsVersion *string `json:"EsVersion,omitnil,omitempty" name:"EsVersion"`

	// 是否只做升级检查，默认值为false
	CheckOnly *bool `json:"CheckOnly,omitnil,omitempty" name:"CheckOnly"`

	// 目标商业特性版本：<li>oss 开源版</li><li>basic 基础版</li>当前仅在5.6.4升级6.x版本时使用，默认值为basic
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// 升级方式：<li>scale 蓝绿变更</li><li>restart 滚动重启</li>默认值为scale
	UpgradeMode *string `json:"UpgradeMode,omitnil,omitempty" name:"UpgradeMode"`

	// 升级版本前是否对集群进行备份，默认不备份
	CosBackup *bool `json:"CosBackup,omitnil,omitempty" name:"CosBackup"`

	// 滚动模式时，是否跳过检查，进行强制重启。默认值为false
	SkipCheckForceRestart *bool `json:"SkipCheckForceRestart,omitnil,omitempty" name:"SkipCheckForceRestart"`

	// cvm延迟上架参数
	CvmDelayOnlineTime *uint64 `json:"CvmDelayOnlineTime,omitnil,omitempty" name:"CvmDelayOnlineTime"`

	// 分片迁移并发数
	ShardAllocationConcurrents *uint64 `json:"ShardAllocationConcurrents,omitnil,omitempty" name:"ShardAllocationConcurrents"`

	// 分片迁移并发速度
	ShardAllocationBytes *uint64 `json:"ShardAllocationBytes,omitnil,omitempty" name:"ShardAllocationBytes"`

	// 是否开启置放群组异步任务
	EnableScheduleRecoverGroup *bool `json:"EnableScheduleRecoverGroup,omitnil,omitempty" name:"EnableScheduleRecoverGroup"`

	// 置放群组异步任务时间段
	EnableScheduleOperationDuration *EnableScheduleOperationDuration `json:"EnableScheduleOperationDuration,omitnil,omitempty" name:"EnableScheduleOperationDuration"`
}

func (r *UpgradeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EsVersion")
	delete(f, "CheckOnly")
	delete(f, "LicenseType")
	delete(f, "BasicSecurityType")
	delete(f, "UpgradeMode")
	delete(f, "CosBackup")
	delete(f, "SkipCheckForceRestart")
	delete(f, "CvmDelayOnlineTime")
	delete(f, "ShardAllocationConcurrents")
	delete(f, "ShardAllocationBytes")
	delete(f, "EnableScheduleRecoverGroup")
	delete(f, "EnableScheduleOperationDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeInstanceResponseParams `json:"Response"`
}

func (r *UpgradeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLicenseRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li><li>不传参时会默认维持原状，传参时需注意只能由不开启变为开启，无法由开启变为不开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// 是否强制重启<li>true强制重启</li><li>false不强制重启</li> 默认值false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

type UpgradeLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// License类型<li>oss：开源版</li><li>basic：基础版</li><li>platinum：白金版</li>默认值platinum
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 是否自动使用代金券<li>0：不自动使用</li><li>1：自动使用</li>默认值0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表（目前仅支持指定一张代金券）
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 6.8（及以上版本）基础版是否开启xpack security认证<li>1：不开启</li><li>2：开启</li><li>不传参时会默认维持原状，传参时需注意只能由不开启变为开启，无法由开启变为不开启</li>
	BasicSecurityType *uint64 `json:"BasicSecurityType,omitnil,omitempty" name:"BasicSecurityType"`

	// 是否强制重启<li>true强制重启</li><li>false不强制重启</li> 默认值false
	ForceRestart *bool `json:"ForceRestart,omitnil,omitempty" name:"ForceRestart"`
}

func (r *UpgradeLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "LicenseType")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "BasicSecurityType")
	delete(f, "ForceRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLicenseResponseParams struct {
	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeLicenseResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeLicenseResponseParams `json:"Response"`
}

func (r *UpgradeLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {
	// vpcId信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// SubnetId信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// VpcUid信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcUid *uint64 `json:"VpcUid,omitnil,omitempty" name:"VpcUid"`

	// SubnetUid 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetUid *uint64 `json:"SubnetUid,omitnil,omitempty" name:"SubnetUid"`

	// 可用ip数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitnil,omitempty" name:"AvailableIpAddressCount"`
}

type WebNodeTypeInfo struct {
	// 可视化节点个数，固定为1
	NodeNum *uint64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 可视化节点规格
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`
}

type ZoneDetail struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否为隐藏可用区
	Hidden *bool `json:"Hidden,omitnil,omitempty" name:"Hidden"`
}