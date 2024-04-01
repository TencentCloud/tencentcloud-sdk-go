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
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘容量，单位G
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘总数
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 描述
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`
}

type ChargeProperties struct {
	// 计费类型，“PREPAID” 预付费，“POSTPAID_BY_HOUR” 后付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 是否自动续费，1表示自动续费开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费时间长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 计费时间单位，“m”表示月等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`
}

type ClusterConfigsInfoFromEMR struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件对应的相关属性信息
	FileConf *string `json:"FileConf,omitnil,omitempty" name:"FileConf"`

	// 配置文件对应的其他属性信息
	KeyConf *string `json:"KeyConf,omitnil,omitempty" name:"KeyConf"`

	// 配置文件的内容，base64编码
	OriParam *string `json:"OriParam,omitnil,omitempty" name:"OriParam"`

	// 用于表示当前配置文件是不是有过修改后没有重启，提醒用户需要重启
	NeedRestart *int64 `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 配置文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 配置文件kv值
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: FileKeyValues is deprecated.
	FileKeyValues *string `json:"FileKeyValues,omitnil,omitempty" name:"FileKeyValues"`

	// 配置文件kv值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileKeyValuesNew []*ConfigKeyValue `json:"FileKeyValuesNew,omitnil,omitempty" name:"FileKeyValuesNew"`
}

type ConfigKeyValue struct {
	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 1-只读，2-可修改但不可删除，3-可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`

	// 0不支持 1支持热更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportHotUpdate *int64 `json:"SupportHotUpdate,omitnil,omitempty" name:"SupportHotUpdate"`
}

// Predefined struct for user
type CreateInstanceNewRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE规格
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE规格
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// 用户VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 用户子网ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 产品版本号
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 付费类型
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 实例名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据库密码
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 高可用类型：
	// 0：非高可用（只有1个FE，FeSpec.CreateInstanceSpec.Count=1），
	// 1：读高可用（至少需部署3个FE，FeSpec.CreateInstanceSpec.Count>=3，且为奇数），
	// 2：读写高可用（至少需部署5个FE，FeSpec.CreateInstanceSpec.Count>=5，且为奇数）。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 是否开启多可用区
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// 开启多可用区后，用户的所有可用区和子网信息
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// FE规格
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil,omitempty" name:"FeSpec"`

	// BE规格
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil,omitempty" name:"BeSpec"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// 用户VPCID
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 用户子网ID
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 产品版本号
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 付费类型
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 实例名字
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据库密码
	DorisUserPwd *string `json:"DorisUserPwd,omitnil,omitempty" name:"DorisUserPwd"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 高可用类型：
	// 0：非高可用（只有1个FE，FeSpec.CreateInstanceSpec.Count=1），
	// 1：读高可用（至少需部署3个FE，FeSpec.CreateInstanceSpec.Count>=3，且为奇数），
	// 2：读写高可用（至少需部署5个FE，FeSpec.CreateInstanceSpec.Count>=5，且为奇数）。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 是否开启多可用区
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// 开启多可用区后，用户的所有可用区和子网信息
	UserMultiZoneInfos *NetworkInfo `json:"UserMultiZoneInfos,omitnil,omitempty" name:"UserMultiZoneInfos"`
}

func (r *CreateInstanceNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "FeSpec")
	delete(f, "BeSpec")
	delete(f, "HaFlag")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ProductVersion")
	delete(f, "ChargeProperties")
	delete(f, "InstanceName")
	delete(f, "DorisUserPwd")
	delete(f, "Tags")
	delete(f, "HaType")
	delete(f, "CaseSensitive")
	delete(f, "EnableMultiZones")
	delete(f, "UserMultiZoneInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceNewResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceNewResponseParams `json:"Response"`
}

func (r *CreateInstanceNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceSpec struct {
	// 规格名字
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type DataBaseAuditRecord struct {
	// 查询用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// 查询ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// 执行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 读取行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// 读取字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// 结果字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// 初始查询IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// catalog名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`
}

// Predefined struct for user
type DescribeClusterConfigsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	//  0 公有云查询；1青鹅查询，青鹅查询显示所有需要展示的
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 模糊搜索关键字文件
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0集群维度 1节点维度
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0的ip地址
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	//  0 公有云查询；1青鹅查询，青鹅查询显示所有需要展示的
	ConfigType *int64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 模糊搜索关键字文件
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 0集群维度 1节点维度
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil,omitempty" name:"ClusterConfigType"`

	// eth0的ip地址
	IPAddress *string `json:"IPAddress,omitnil,omitempty" name:"IPAddress"`
}

func (r *DescribeClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigType")
	delete(f, "FileName")
	delete(f, "ClusterConfigType")
	delete(f, "IPAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsResponseParams struct {
	// 返回实例的配置文件相关的信息
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil,omitempty" name:"ClusterConfList"`

	// 返回当前内核版本 如果不存在则返回空字符串
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterConfigsResponseParams `json:"Response"`
}

func (r *DescribeClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 多选
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 多选
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 多选
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称 （多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 多选
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 多选
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 多选
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称 （多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditDownloadResponseParams struct {
	// 日志的cos地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditDownloadResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 （多选）
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 （多选）
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 （多选）
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称（多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

type DescribeDatabaseAuditRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 用户 （多选）
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 数据库 （多选）
	DbNames []*string `json:"DbNames,omitnil,omitempty" name:"DbNames"`

	// sql类型 （多选）
	SqlTypes []*string `json:"SqlTypes,omitnil,omitempty" name:"SqlTypes"`

	// catalog名称（多选）
	Catalogs []*string `json:"Catalogs,omitnil,omitempty" name:"Catalogs"`
}

func (r *DescribeDatabaseAuditRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "OrderType")
	delete(f, "User")
	delete(f, "DbName")
	delete(f, "SqlType")
	delete(f, "Sql")
	delete(f, "Users")
	delete(f, "DbNames")
	delete(f, "SqlTypes")
	delete(f, "Catalogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseAuditRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseAuditRecordsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	SlowQueryRecords *DataBaseAuditRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseAuditRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseAuditRecordsResponseParams `json:"Response"`
}

func (r *DescribeDatabaseAuditRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseAuditRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoRequestParams struct {
	// 集群id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeInstanceNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

func (r *DescribeInstanceNodesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceNodesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesInfoResponseParams struct {
	// Be节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeNodes []*string `json:"BeNodes,omitnil,omitempty" name:"BeNodes"`

	// Fe节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeNodes []*string `json:"FeNodes,omitnil,omitempty" name:"FeNodes"`

	// Fe master节点
	FeMaster *string `json:"FeMaster,omitnil,omitempty" name:"FeMaster"`

	// Be节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeNodeInfos []*NodeInfo `json:"BeNodeInfos,omitnil,omitempty" name:"BeNodeInfos"`

	// Fe节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeNodeInfos []*NodeInfo `json:"FeNodeInfos,omitnil,omitempty" name:"FeNodeInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceNodesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceNodesInfoResponseParams `json:"Response"`
}

func (r *DescribeInstanceNodesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceNodesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceNodesRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群角色类型，默认为 "data"数据节点
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群角色类型，默认为 "data"数据节点
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例节点总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNodesList []*InstanceNode `json:"InstanceNodesList,omitnil,omitempty" name:"InstanceNodesList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceStateRequestParams struct {
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceStateResponseParams struct {
	// 集群状态，例如：Serving
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceStateResponseParams `json:"Response"`
}

func (r *DescribeInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索的集群id名称
	SearchInstanceId *string `json:"SearchInstanceId,omitnil,omitempty" name:"SearchInstanceId"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例数组
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

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
type DescribeSlowQueryRecordsDownloadRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`
}

type DescribeSlowQueryRecordsDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`
}

func (r *DescribeSlowQueryRecordsDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DurationMs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadResponseParams struct {
	// cos地址
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsDownloadResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 是否是查询，0：否， 1：是
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeSlowQueryRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil,omitempty" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 是否是查询，0：否， 1：是
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeSlowQueryRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QueryDurationMs")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "DurationMs")
	delete(f, "DbName")
	delete(f, "IsQuery")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowQueryRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowQueryRecordsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	SlowQueryRecords []*SlowQueryRecord `json:"SlowQueryRecords,omitnil,omitempty" name:"SlowQueryRecords"`

	// 所有数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBNameList []*string `json:"DBNameList,omitnil,omitempty" name:"DBNameList"`

	// 所有catalog名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogNameList []*string `json:"CatalogNameList,omitnil,omitempty" name:"CatalogNameList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowQueryRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowQueryRecordsResponseParams `json:"Response"`
}

func (r *DescribeSlowQueryRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowQueryRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DestroyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceResponseParams `json:"Response"`
}

func (r *DestroyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {
	// 集群实例ID, "cdw-xxxx" 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 状态,
	// Init 创建中; Serving 运行中； 
	// Deleted已销毁；Deleting 销毁中；
	// Modify 集群变更中；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 地域, ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区， ap-guangzhou-3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 付费类型，"hour", "prepay"
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 数据节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSummary *NodesSummary `json:"MasterSummary,omitnil,omitempty" name:"MasterSummary"`

	// zookeeper节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoreSummary *NodesSummary `json:"CoreSummary,omitnil,omitempty" name:"CoreSummary"`

	// 高可用，“true" "false"
	// 注意：此字段可能返回 null，表示取不到有效值。
	HA *string `json:"HA,omitnil,omitempty" name:"HA"`

	// 高可用类型：
	// 0：非高可用
	// 1：读高可用
	// 2：读写高可用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`

	// 访问地址，例如 "10.0.0.1:9000"
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// 记录ID，数值型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// regionId, 表示地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 可用区说明，例如 "广州二区"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// 错误流程说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 状态描述，例如“运行中”等
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 自动续费标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Monitor *string `json:"Monitor,omitnil,omitempty" name:"Monitor"`

	// 是否开通日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasClsTopic *bool `json:"HasClsTopic,omitnil,omitempty" name:"HasClsTopic"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// 是否支持xml配置管理
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitnil,omitempty" name:"EnableXMLConfig"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// 弹性网卡地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eip *string `json:"Eip,omitnil,omitempty" name:"Eip"`

	// 冷热分层系数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosMoveFactor *int64 `json:"CosMoveFactor,omitnil,omitempty" name:"CosMoveFactor"`

	// external/local/yunti
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// cos桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// cbs
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil,omitempty" name:"CanAttachCbs"`

	// 小版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	BuildVersion *string `json:"BuildVersion,omitnil,omitempty" name:"BuildVersion"`

	// 组件信息
	// 注：这里返回类型实际为map[string]struct类型，并非显示的string类型，可以参考“示例值”进行数据的解析。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components *string `json:"Components,omitnil,omitempty" name:"Components"`

	// 判断审计日志表是否有catalog字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IfExistCatalog is deprecated.
	IfExistCatalog *int64 `json:"IfExistCatalog,omitnil,omitempty" name:"IfExistCatalog"`

	// 页面特性，用于前端屏蔽一些页面入口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Characteristic []*string `json:"Characteristic,omitnil,omitempty" name:"Characteristic"`

	// 超时时间 单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartTimeout *string `json:"RestartTimeout,omitnil,omitempty" name:"RestartTimeout"`

	// 内核优雅重启超时时间，如果为-1说明未设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GraceShutdownWaitSeconds *string `json:"GraceShutdownWaitSeconds,omitnil,omitempty" name:"GraceShutdownWaitSeconds"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *int64 `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 用户是否可以绑定安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWhiteSGs *bool `json:"IsWhiteSGs,omitnil,omitempty" name:"IsWhiteSGs"`

	// 已绑定的安全组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindSGs []*string `json:"BindSGs,omitnil,omitempty" name:"BindSGs"`

	// 是否为多可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableMultiZones *bool `json:"EnableMultiZones,omitnil,omitempty" name:"EnableMultiZones"`

	// 用户可用区和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserNetworkInfos *string `json:"UserNetworkInfos,omitnil,omitempty" name:"UserNetworkInfos"`
}

type InstanceNode struct {
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 机型，如 S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// cpu核数
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// 内存大小
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 所属clickhouse cluster名称
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// rip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rip *string `json:"Rip,omitnil,omitempty" name:"Rip"`

	// FE节点角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeRole *string `json:"FeRole,omitnil,omitempty" name:"FeRole"`

	// UUID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInfo struct {
	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 当前子网可用ip数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIpNum *int64 `json:"SubnetIpNum,omitnil,omitempty" name:"SubnetIpNum"`
}

type NodeInfo struct {
	// 用户IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 节点状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点角色名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 组件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentName *string `json:"ComponentName,omitnil,omitempty" name:"ComponentName"`

	// 节点角色
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 节点上次重启的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRestartTime *string `json:"LastRestartTime,omitnil,omitempty" name:"LastRestartTime"`

	// 节点所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type NodesSummary struct {
	// 机型，如 S1
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 节点数目
	NodeSize *int64 `json:"NodeSize,omitnil,omitempty" name:"NodeSize"`

	// cpu核数，单位个
	Core *int64 `json:"Core,omitnil,omitempty" name:"Core"`

	// 内存大小，单位G
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘大小，单位G
	Disk *int64 `json:"Disk,omitnil,omitempty" name:"Disk"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘描述
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// 挂载云盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// 子产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductType *string `json:"SubProductType,omitnil,omitempty" name:"SubProductType"`

	// 规格核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCore *int64 `json:"SpecCore,omitnil,omitempty" name:"SpecCore"`

	// 规格内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecMemory *int64 `json:"SpecMemory,omitnil,omitempty" name:"SpecMemory"`

	// 磁盘大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 是否加密
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 最大磁盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

func (r *ResizeDiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "DiskSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeDiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResizeDiskResponse struct {
	*tchttp.BaseResponse
	Response *ResizeDiskResponseParams `json:"Response"`
}

func (r *ResizeDiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeDiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 每次重启的批次
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 重启节点
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// false表示非滚动重启，true表示滚动重启
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

type RestartClusterForNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 每次重启的批次
	BatchSize *int64 `json:"BatchSize,omitnil,omitempty" name:"BatchSize"`

	// 重启节点
	NodeList []*string `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// false表示非滚动重启，true表示滚动重启
	RollingRestart *bool `json:"RollingRestart,omitnil,omitempty" name:"RollingRestart"`
}

func (r *RestartClusterForNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigName")
	delete(f, "BatchSize")
	delete(f, "NodeList")
	delete(f, "RollingRestart")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartClusterForNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartClusterForNodeResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartClusterForNodeResponse struct {
	*tchttp.BaseResponse
	Response *RestartClusterForNodeResponseParams `json:"Response"`
}

func (r *RestartClusterForNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartClusterForNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 扩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 扩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil,omitempty" name:"HaType"`
}

func (r *ScaleOutInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "NodeCount")
	delete(f, "HaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstanceResponseParams `json:"Response"`
}

func (r *ScaleOutInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点规格
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点规格
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ScaleUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SpecName")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchTags struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SlowQueryRecord struct {
	// 查询用户
	OsUser *string `json:"OsUser,omitnil,omitempty" name:"OsUser"`

	// 查询ID
	InitialQueryId *string `json:"InitialQueryId,omitnil,omitempty" name:"InitialQueryId"`

	// SQL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 开始时间
	QueryStartTime *string `json:"QueryStartTime,omitnil,omitempty" name:"QueryStartTime"`

	// 执行耗时
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 读取行数
	ReadRows *int64 `json:"ReadRows,omitnil,omitempty" name:"ReadRows"`

	// 读取字节数
	ResultRows *int64 `json:"ResultRows,omitnil,omitempty" name:"ResultRows"`

	// 结果字节数
	ResultBytes *uint64 `json:"ResultBytes,omitnil,omitempty" name:"ResultBytes"`

	// 内存
	MemoryUsage *int64 `json:"MemoryUsage,omitnil,omitempty" name:"MemoryUsage"`

	// 初始查询IP
	InitialAddress *string `json:"InitialAddress,omitnil,omitempty" name:"InitialAddress"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 是否是查询，0：否，1：查询语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQuery *int64 `json:"IsQuery,omitnil,omitempty" name:"IsQuery"`

	// ResultBytes的MB格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultBytesMB *float64 `json:"ResultBytesMB,omitnil,omitempty" name:"ResultBytesMB"`

	// MemoryUsage的MB表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryUsageMB *float64 `json:"MemoryUsageMB,omitnil,omitempty" name:"MemoryUsageMB"`

	// DurationMs的秒表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationSec *float64 `json:"DurationSec,omitnil,omitempty" name:"DurationSec"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}