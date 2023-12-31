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

type ChargeProperties struct {
	// 计费类型，“PREPAID” 预付费，“POSTPAID_BY_HOUR” 后付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 是否自动续费，1表示自动续费开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 计费时间长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 计费时间单位，“m”表示月等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`
}

type ClusterConfigsInfoFromEMR struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 配置文件对应的相关属性信息
	FileConf *string `json:"FileConf,omitnil" name:"FileConf"`

	// 配置文件对应的其他属性信息
	KeyConf *string `json:"KeyConf,omitnil" name:"KeyConf"`

	// 配置文件的内容，base64编码
	OriParam *string `json:"OriParam,omitnil" name:"OriParam"`

	// 用于表示当前配置文件是不是有过修改后没有重启，提醒用户需要重启
	NeedRestart *int64 `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// 配置文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`

	// 配置文件kv值
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: FileKeyValues is deprecated.
	FileKeyValues *string `json:"FileKeyValues,omitnil" name:"FileKeyValues"`

	// 配置文件kv值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileKeyValuesNew []*ConfigKeyValue `json:"FileKeyValuesNew,omitnil" name:"FileKeyValuesNew"`
}

type ConfigKeyValue struct {
	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyName *string `json:"KeyName,omitnil" name:"KeyName"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 1-只读，2-可修改但不可删除，3-可删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Display *int64 `json:"Display,omitnil" name:"Display"`

	// 0不支持 1支持热更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportHotUpdate *int64 `json:"SupportHotUpdate,omitnil" name:"SupportHotUpdate"`
}

// Predefined struct for user
type CreateInstanceNewRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// FE规格
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil" name:"FeSpec"`

	// BE规格
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil" name:"BeSpec"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil" name:"HaFlag"`

	// 用户VPCID
	UserVPCId *string `json:"UserVPCId,omitnil" name:"UserVPCId"`

	// 用户子网ID
	UserSubnetId *string `json:"UserSubnetId,omitnil" name:"UserSubnetId"`

	// 产品版本号
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 付费类型
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil" name:"ChargeProperties"`

	// 实例名字
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 数据库密码
	DorisUserPwd *string `json:"DorisUserPwd,omitnil" name:"DorisUserPwd"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 高可用类型：
	// 0：非高可用（只有1个FE，FeSpec.CreateInstanceSpec.Count=1），
	// 1：读高可用（至少需部署3个FE，FeSpec.CreateInstanceSpec.Count>=3，且为奇数），
	// 2：读写高可用（至少需部署5个FE，FeSpec.CreateInstanceSpec.Count>=5，且为奇数）。
	HaType *int64 `json:"HaType,omitnil" name:"HaType"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	CaseSensitive *int64 `json:"CaseSensitive,omitnil" name:"CaseSensitive"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// FE规格
	FeSpec *CreateInstanceSpec `json:"FeSpec,omitnil" name:"FeSpec"`

	// BE规格
	BeSpec *CreateInstanceSpec `json:"BeSpec,omitnil" name:"BeSpec"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil" name:"HaFlag"`

	// 用户VPCID
	UserVPCId *string `json:"UserVPCId,omitnil" name:"UserVPCId"`

	// 用户子网ID
	UserSubnetId *string `json:"UserSubnetId,omitnil" name:"UserSubnetId"`

	// 产品版本号
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 付费类型
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil" name:"ChargeProperties"`

	// 实例名字
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 数据库密码
	DorisUserPwd *string `json:"DorisUserPwd,omitnil" name:"DorisUserPwd"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 高可用类型：
	// 0：非高可用（只有1个FE，FeSpec.CreateInstanceSpec.Count=1），
	// 1：读高可用（至少需部署3个FE，FeSpec.CreateInstanceSpec.Count>=3，且为奇数），
	// 2：读写高可用（至少需部署5个FE，FeSpec.CreateInstanceSpec.Count>=5，且为奇数）。
	HaType *int64 `json:"HaType,omitnil" name:"HaType"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	CaseSensitive *int64 `json:"CaseSensitive,omitnil" name:"CaseSensitive"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 数量
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type DataBaseAuditRecord struct {
	// 查询用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	OsUser *string `json:"OsUser,omitnil" name:"OsUser"`

	// 查询ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialQueryId *string `json:"InitialQueryId,omitnil" name:"InitialQueryId"`

	// SQL语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sql *string `json:"Sql,omitnil" name:"Sql"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryStartTime *string `json:"QueryStartTime,omitnil" name:"QueryStartTime"`

	// 执行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationMs *int64 `json:"DurationMs,omitnil" name:"DurationMs"`

	// 读取行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadRows *int64 `json:"ReadRows,omitnil" name:"ReadRows"`

	// 读取字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultRows *int64 `json:"ResultRows,omitnil" name:"ResultRows"`

	// 结果字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultBytes *uint64 `json:"ResultBytes,omitnil" name:"ResultBytes"`

	// 内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemoryUsage *int64 `json:"MemoryUsage,omitnil" name:"MemoryUsage"`

	// 初始查询IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitialAddress *string `json:"InitialAddress,omitnil" name:"InitialAddress"`

	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// sql类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlType *string `json:"SqlType,omitnil" name:"SqlType"`

	// catalog名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil" name:"Catalog"`
}

// Predefined struct for user
type DescribeClusterConfigsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	//  0 公有云查询；1青鹅查询，青鹅查询显示所有需要展示的
	ConfigType *int64 `json:"ConfigType,omitnil" name:"ConfigType"`

	// 模糊搜索关键字文件
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 0集群维度 1节点维度
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil" name:"ClusterConfigType"`

	// eth0的ip地址
	IPAddress *string `json:"IPAddress,omitnil" name:"IPAddress"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	//  0 公有云查询；1青鹅查询，青鹅查询显示所有需要展示的
	ConfigType *int64 `json:"ConfigType,omitnil" name:"ConfigType"`

	// 模糊搜索关键字文件
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 0集群维度 1节点维度
	ClusterConfigType *int64 `json:"ClusterConfigType,omitnil" name:"ClusterConfigType"`

	// eth0的ip地址
	IPAddress *string `json:"IPAddress,omitnil" name:"IPAddress"`
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
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil" name:"ClusterConfList"`

	// 返回当前内核版本 如果不存在则返回空字符串
	BuildVersion *string `json:"BuildVersion,omitnil" name:"BuildVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil" name:"Sql"`

	// 用户 多选
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 数据库 多选
	DbNames []*string `json:"DbNames,omitnil" name:"DbNames"`

	// sql类型 多选
	SqlTypes []*string `json:"SqlTypes,omitnil" name:"SqlTypes"`

	// catalog名称 （多选）
	Catalogs []*string `json:"Catalogs,omitnil" name:"Catalogs"`
}

type DescribeDatabaseAuditDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil" name:"Sql"`

	// 用户 多选
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 数据库 多选
	DbNames []*string `json:"DbNames,omitnil" name:"DbNames"`

	// sql类型 多选
	SqlTypes []*string `json:"SqlTypes,omitnil" name:"SqlTypes"`

	// catalog名称 （多选）
	Catalogs []*string `json:"Catalogs,omitnil" name:"Catalogs"`
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
	CosUrl *string `json:"CosUrl,omitnil" name:"CosUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil" name:"Sql"`

	// 用户 （多选）
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 数据库 （多选）
	DbNames []*string `json:"DbNames,omitnil" name:"DbNames"`

	// sql类型 （多选）
	SqlTypes []*string `json:"SqlTypes,omitnil" name:"SqlTypes"`

	// catalog名称（多选）
	Catalogs []*string `json:"Catalogs,omitnil" name:"Catalogs"`
}

type DescribeDatabaseAuditRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 排序参数
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 用户
	User *string `json:"User,omitnil" name:"User"`

	// 数据库
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// sql类型
	SqlType *string `json:"SqlType,omitnil" name:"SqlType"`

	// sql语句
	Sql *string `json:"Sql,omitnil" name:"Sql"`

	// 用户 （多选）
	Users []*string `json:"Users,omitnil" name:"Users"`

	// 数据库 （多选）
	DbNames []*string `json:"DbNames,omitnil" name:"DbNames"`

	// sql类型 （多选）
	SqlTypes []*string `json:"SqlTypes,omitnil" name:"SqlTypes"`

	// catalog名称（多选）
	Catalogs []*string `json:"Catalogs,omitnil" name:"Catalogs"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 记录列表
	SlowQueryRecords *DataBaseAuditRecord `json:"SlowQueryRecords,omitnil" name:"SlowQueryRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeInstanceNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
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
	BeNodes []*string `json:"BeNodes,omitnil" name:"BeNodes"`

	// Fe节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeNodes []*string `json:"FeNodes,omitnil" name:"FeNodes"`

	// Fe master节点
	FeMaster *string `json:"FeMaster,omitnil" name:"FeMaster"`

	// Be节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeNodeInfos []*NodeInfo `json:"BeNodeInfos,omitnil" name:"BeNodeInfos"`

	// Fe节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeNodeInfos []*NodeInfo `json:"FeNodeInfos,omitnil" name:"FeNodeInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeInstanceStateRequestParams struct {
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例名称
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitnil" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *float64 `json:"FlowProgress,omitnil" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil" name:"FlowMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type DescribeSlowQueryRecordsDownloadRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil" name:"DurationMs"`
}

type DescribeSlowQueryRecordsDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil" name:"DurationMs"`
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
	CosUrl *string `json:"CosUrl,omitnil" name:"CosUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil" name:"DurationMs"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil" name:"DbName"`

	// 是否是查询，0：否， 1：是
	IsQuery *int64 `json:"IsQuery,omitnil" name:"IsQuery"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil" name:"CatalogName"`
}

type DescribeSlowQueryRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 慢查询时间
	QueryDurationMs *int64 `json:"QueryDurationMs,omitnil" name:"QueryDurationMs"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 分页
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 排序参数
	DurationMs *string `json:"DurationMs,omitnil" name:"DurationMs"`

	// 数据库名称
	DbName []*string `json:"DbName,omitnil" name:"DbName"`

	// 是否是查询，0：否， 1：是
	IsQuery *int64 `json:"IsQuery,omitnil" name:"IsQuery"`

	// catalog名称
	CatalogName []*string `json:"CatalogName,omitnil" name:"CatalogName"`
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
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 记录列表
	SlowQueryRecords []*SlowQueryRecord `json:"SlowQueryRecords,omitnil" name:"SlowQueryRecords"`

	// 所有数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBNameList []*string `json:"DBNameList,omitnil" name:"DBNameList"`

	// 所有catalog名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogNameList []*string `json:"CatalogNameList,omitnil" name:"CatalogNameList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

	// 判断审计日志表是否有catalog字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IfExistCatalog is deprecated.
	IfExistCatalog *int64 `json:"IfExistCatalog,omitnil" name:"IfExistCatalog"`

	// 页面特性，用于前端屏蔽一些页面入口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Characteristic []*string `json:"Characteristic,omitnil" name:"Characteristic"`

	// 超时时间 单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestartTimeout *string `json:"RestartTimeout,omitnil" name:"RestartTimeout"`

	// 内核优雅重启超时时间，如果为-1说明未设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GraceShutdownWaitSeconds *string `json:"GraceShutdownWaitSeconds,omitnil" name:"GraceShutdownWaitSeconds"`

	// 表名大小写是否敏感，0：敏感；1：不敏感，以小写进行比较；2：不敏感，表名改为以小写存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *int64 `json:"CaseSensitive,omitnil" name:"CaseSensitive"`
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

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 新修改的实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type NodeInfo struct {
	// 用户IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 节点状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`
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

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil" name:"Type"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil" name:"Type"`

	// 云盘大小
	DiskSize *uint64 `json:"DiskSize,omitnil" name:"DiskSize"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 流程ID
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil" name:"ConfigName"`

	// 每次重启的批次
	BatchSize *int64 `json:"BatchSize,omitnil" name:"BatchSize"`

	// 重启节点
	NodeList []*string `json:"NodeList,omitnil" name:"NodeList"`

	// false表示非滚动重启，true表示滚动重启
	RollingRestart *bool `json:"RollingRestart,omitnil" name:"RollingRestart"`
}

type RestartClusterForNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 配置文件名称
	ConfigName *string `json:"ConfigName,omitnil" name:"ConfigName"`

	// 每次重启的批次
	BatchSize *int64 `json:"BatchSize,omitnil" name:"BatchSize"`

	// 重启节点
	NodeList []*string `json:"NodeList,omitnil" name:"NodeList"`

	// false表示非滚动重启，true表示滚动重启
	RollingRestart *bool `json:"RollingRestart,omitnil" name:"RollingRestart"`
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
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil" name:"Type"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// 扩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil" name:"HaType"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil" name:"Type"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// 扩容后集群高可用类型：0：非高可用，1：读高可用，2：读写高可用。
	HaType *int64 `json:"HaType,omitnil" name:"HaType"`
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
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点规格
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点规格
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 角色（MATER/CORE），MASTER 对应 FE，CORE对应BE
	Type *string `json:"Type,omitnil" name:"Type"`
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
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil" name:"AllValue"`
}

type SlowQueryRecord struct {
	// 查询用户
	OsUser *string `json:"OsUser,omitnil" name:"OsUser"`

	// 查询ID
	InitialQueryId *string `json:"InitialQueryId,omitnil" name:"InitialQueryId"`

	// SQL语句
	Sql *string `json:"Sql,omitnil" name:"Sql"`

	// 开始时间
	QueryStartTime *string `json:"QueryStartTime,omitnil" name:"QueryStartTime"`

	// 执行耗时
	DurationMs *int64 `json:"DurationMs,omitnil" name:"DurationMs"`

	// 读取行数
	ReadRows *int64 `json:"ReadRows,omitnil" name:"ReadRows"`

	// 读取字节数
	ResultRows *int64 `json:"ResultRows,omitnil" name:"ResultRows"`

	// 结果字节数
	ResultBytes *uint64 `json:"ResultBytes,omitnil" name:"ResultBytes"`

	// 内存
	MemoryUsage *int64 `json:"MemoryUsage,omitnil" name:"MemoryUsage"`

	// 初始查询IP
	InitialAddress *string `json:"InitialAddress,omitnil" name:"InitialAddress"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 是否是查询，0：否，1：查询语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsQuery *int64 `json:"IsQuery,omitnil" name:"IsQuery"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}