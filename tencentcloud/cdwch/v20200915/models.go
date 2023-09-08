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

package v20200915

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ActionAlterCkUserRequestParams struct {
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitnil" name:"UserInfo"`

	// api接口类型，
	// AddSystemUser新增用户，UpdateSystemUser，修改用户
	ApiType *string `json:"ApiType,omitnil" name:"ApiType"`
}

type ActionAlterCkUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitnil" name:"UserInfo"`

	// api接口类型，
	// AddSystemUser新增用户，UpdateSystemUser，修改用户
	ApiType *string `json:"ApiType,omitnil" name:"ApiType"`
}

func (r *ActionAlterCkUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActionAlterCkUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserInfo")
	delete(f, "ApiType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActionAlterCkUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActionAlterCkUserResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ActionAlterCkUserResponse struct {
	*tchttp.BaseResponse
	Response *ActionAlterCkUserResponseParams `json:"Response"`
}

func (r *ActionAlterCkUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActionAlterCkUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

type BackUpJobDisplay struct {
	// 备份任务id
	JobId *int64 `json:"JobId,omitnil" name:"JobId"`

	// 备份任务名
	Snapshot *string `json:"Snapshot,omitnil" name:"Snapshot"`

	// 任务类型(元数据),(数据)
	BackUpType *string `json:"BackUpType,omitnil" name:"BackUpType"`

	// 备份数据量
	BackUpSize *int64 `json:"BackUpSize,omitnil" name:"BackUpSize"`

	// 任务创建时间
	BackUpTime *string `json:"BackUpTime,omitnil" name:"BackUpTime"`

	// 任务过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 任务状态
	JobStatus *string `json:"JobStatus,omitnil" name:"JobStatus"`
}

type BackupTableContent struct {
	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Database *string `json:"Database,omitnil" name:"Database"`

	// 表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitnil" name:"Table"`

	// 表总字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalBytes *int64 `json:"TotalBytes,omitnil" name:"TotalBytes"`

	// 虚拟cluster
	// 注意：此字段可能返回 null，表示取不到有效值。
	VCluster *string `json:"VCluster,omitnil" name:"VCluster"`

	// 表ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ips *string `json:"Ips,omitnil" name:"Ips"`

	// zk路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZooPath *string `json:"ZooPath,omitnil" name:"ZooPath"`

	// cvm的ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rip *string `json:"Rip,omitnil" name:"Rip"`
}

type Charge struct {
	// 计费类型，“PREPAID” 预付费，“POSTPAID_BY_HOUR” 后付费
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// PREPAID需要传递，是否自动续费，1表示自动续费开启
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 预付费需要传递，计费时间长度，多少个月
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

type CkUserAlterInfo struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitnil" name:"PassWord"`

	// 描述
	Describe *string `json:"Describe,omitnil" name:"Describe"`
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

	// 保存配置文件的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`
}

type ClusterInfo struct {
	// vcluster名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// 当前cluster的IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIps []*string `json:"NodeIps,omitnil" name:"NodeIps"`
}

type ConfigSubmitContext struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 配置文件旧内容，base64编码
	OldConfValue *string `json:"OldConfValue,omitnil" name:"OldConfValue"`

	// 配置文件新内容，base64编码
	NewConfValue *string `json:"NewConfValue,omitnil" name:"NewConfValue"`

	// 保存配置文件的路径
	FilePath *string `json:"FilePath,omitnil" name:"FilePath"`
}

// Predefined struct for user
type CreateBackUpScheduleRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 策略类型 meta(元数据)  data (表数据)
	ScheduleType *string `json:"ScheduleType,omitnil" name:"ScheduleType"`

	// 操作类型 create(创建) update(编辑修改)
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 保留天数 例如7
	RetainDays *int64 `json:"RetainDays,omitnil" name:"RetainDays"`

	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 选择的星期 逗号分隔，例如 2 代表周二
	WeekDays *string `json:"WeekDays,omitnil" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitnil" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil" name:"BackUpTables"`
}

type CreateBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 策略类型 meta(元数据)  data (表数据)
	ScheduleType *string `json:"ScheduleType,omitnil" name:"ScheduleType"`

	// 操作类型 create(创建) update(编辑修改)
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 保留天数 例如7
	RetainDays *int64 `json:"RetainDays,omitnil" name:"RetainDays"`

	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 选择的星期 逗号分隔，例如 2 代表周二
	WeekDays *string `json:"WeekDays,omitnil" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitnil" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil" name:"BackUpTables"`
}

func (r *CreateBackUpScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ScheduleType")
	delete(f, "OperationType")
	delete(f, "RetainDays")
	delete(f, "ScheduleId")
	delete(f, "WeekDays")
	delete(f, "ExecuteHour")
	delete(f, "BackUpTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackUpScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackUpScheduleResponseParams struct {
	// 错误描述
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBackUpScheduleResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackUpScheduleResponseParams `json:"Response"`
}

func (r *CreateBackUpScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackUpScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewRequestParams struct {
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil" name:"HaFlag"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil" name:"UserSubnetId"`

	// 版本
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 计费方式
	ChargeProperties *Charge `json:"ChargeProperties,omitnil" name:"ChargeProperties"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 数据节点
	DataSpec *NodeSpec `json:"DataSpec,omitnil" name:"DataSpec"`

	// 标签列表
	Tags *Tag `json:"Tags,omitnil" name:"Tags"`

	// 日志主题ID
	ClsLogSetId *string `json:"ClsLogSetId,omitnil" name:"ClsLogSetId"`

	// COS桶名称
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// 是否是裸盘挂载
	MountDiskType *int64 `json:"MountDiskType,omitnil" name:"MountDiskType"`

	// 是否是ZK高可用
	HAZk *bool `json:"HAZk,omitnil" name:"HAZk"`

	// ZK节点
	CommonSpec *NodeSpec `json:"CommonSpec,omitnil" name:"CommonSpec"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil" name:"HaFlag"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil" name:"UserSubnetId"`

	// 版本
	ProductVersion *string `json:"ProductVersion,omitnil" name:"ProductVersion"`

	// 计费方式
	ChargeProperties *Charge `json:"ChargeProperties,omitnil" name:"ChargeProperties"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 数据节点
	DataSpec *NodeSpec `json:"DataSpec,omitnil" name:"DataSpec"`

	// 标签列表
	Tags *Tag `json:"Tags,omitnil" name:"Tags"`

	// 日志主题ID
	ClsLogSetId *string `json:"ClsLogSetId,omitnil" name:"ClsLogSetId"`

	// COS桶名称
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// 是否是裸盘挂载
	MountDiskType *int64 `json:"MountDiskType,omitnil" name:"MountDiskType"`

	// 是否是ZK高可用
	HAZk *bool `json:"HAZk,omitnil" name:"HAZk"`

	// ZK节点
	CommonSpec *NodeSpec `json:"CommonSpec,omitnil" name:"CommonSpec"`
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
	delete(f, "HaFlag")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ProductVersion")
	delete(f, "ChargeProperties")
	delete(f, "InstanceName")
	delete(f, "DataSpec")
	delete(f, "Tags")
	delete(f, "ClsLogSetId")
	delete(f, "CosBucketName")
	delete(f, "MountDiskType")
	delete(f, "HAZk")
	delete(f, "CommonSpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceNewResponseParams struct {
	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type DatabasePrivilegeInfo struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil" name:"DatabaseName"`

	// //库表权限，SELECT、INSERT_ALL、ALTER、TRUNCATE、DROP_TABLE、CREATE_TABLE、DROP_DATABASE
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabasePrivileges []*string `json:"DatabasePrivileges,omitnil" name:"DatabasePrivileges"`

	// // 库下面的表权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	TablePrivilegeList []*TablePrivilegeInfo `json:"TablePrivilegeList,omitnil" name:"TablePrivilegeList"`
}

// Predefined struct for user
type DeleteBackUpDataRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil" name:"BackUpJobId"`

	// 是否删除所有数据
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil" name:"IsDeleteAll"`
}

type DeleteBackUpDataRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil" name:"BackUpJobId"`

	// 是否删除所有数据
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil" name:"IsDeleteAll"`
}

func (r *DeleteBackUpDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackUpDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	delete(f, "IsDeleteAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackUpDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackUpDataResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBackUpDataResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackUpDataResponseParams `json:"Response"`
}

func (r *DeleteBackUpDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackUpDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobDetailRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil" name:"BackUpJobId"`
}

type DescribeBackUpJobDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil" name:"BackUpJobId"`
}

func (r *DescribeBackUpJobDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpJobDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobDetailResponseParams struct {
	// 备份表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableContents []*BackupTableContent `json:"TableContents,omitnil" name:"TableContents"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackUpJobDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpJobDetailResponseParams `json:"Response"`
}

func (r *DescribeBackUpJobDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页号
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 页号
	PageNum *int64 `json:"PageNum,omitnil" name:"PageNum"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeBackUpJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PageSize")
	delete(f, "PageNum")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpJobResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackUpJobs []*BackUpJobDisplay `json:"BackUpJobs,omitnil" name:"BackUpJobs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackUpJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpJobResponseParams `json:"Response"`
}

func (r *DescribeBackUpJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpScheduleRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeBackUpScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpScheduleResponseParams struct {
	// 备份是否开启
	BackUpOpened *bool `json:"BackUpOpened,omitnil" name:"BackUpOpened"`

	// 元数据备份策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaStrategy *ScheduleStrategy `json:"MetaStrategy,omitnil" name:"MetaStrategy"`

	// 表数据备份策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataStrategy *ScheduleStrategy `json:"DataStrategy,omitnil" name:"DataStrategy"`

	// 备份表列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackUpContents []*BackupTableContent `json:"BackUpContents,omitnil" name:"BackUpContents"`

	// 备份的状态
	BackUpStatus *int64 `json:"BackUpStatus,omitnil" name:"BackUpStatus"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackUpScheduleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpScheduleResponseParams `json:"Response"`
}

func (r *DescribeBackUpScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTablesRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeBackUpTablesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeBackUpTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackUpTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackUpTablesResponseParams struct {
	// 可备份表列表
	AvailableTables []*BackupTableContent `json:"AvailableTables,omitnil" name:"AvailableTables"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBackUpTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackUpTablesResponseParams `json:"Response"`
}

func (r *DescribeBackUpTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackUpTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkSqlApisRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// api接口名称,GetClusters:获取集群cluster列表
	// GetSystemUsers:获取系统用户列表
	// CheckNodeCluster: 检查节点是否隶属一个cluster
	// GetClusterDatabases: 获取一个cluster下的数据库列表
	// GetClusterTables: 获取一个cluster下的数据库表列表
	// GetPrivilegeUsers: 获取授权的用户列表
	// GET_USER_CLUSTER_PRIVILEGES:获取用户cluster下的权限   
	// GetUserClusterNewPrivileges:获取用户cluster下的权限 (新版）
	// RevokeClusterUser:解绑cluster用户
	// DeleteSystemUser:删除系统用户 —— 必须所有cluster先解绑
	// GetUserOptionMessages:获取用户配置备注信息
	// GET_USER_CONFIGS:获取用户配置列表  QUOTA、PROFILE、POLICY
	ApiType *string `json:"ApiType,omitnil" name:"ApiType"`

	// 集群名称，GET_SYSTEM_USERS，GET_PRIVILEGE_USERS，GET_CLUSTER_DATABASES，GET_CLUSTER_TABLES 必填
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 用户名称，api与user相关的必填
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 账户的类型
	UserType *string `json:"UserType,omitnil" name:"UserType"`
}

type DescribeCkSqlApisRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// api接口名称,GetClusters:获取集群cluster列表
	// GetSystemUsers:获取系统用户列表
	// CheckNodeCluster: 检查节点是否隶属一个cluster
	// GetClusterDatabases: 获取一个cluster下的数据库列表
	// GetClusterTables: 获取一个cluster下的数据库表列表
	// GetPrivilegeUsers: 获取授权的用户列表
	// GET_USER_CLUSTER_PRIVILEGES:获取用户cluster下的权限   
	// GetUserClusterNewPrivileges:获取用户cluster下的权限 (新版）
	// RevokeClusterUser:解绑cluster用户
	// DeleteSystemUser:删除系统用户 —— 必须所有cluster先解绑
	// GetUserOptionMessages:获取用户配置备注信息
	// GET_USER_CONFIGS:获取用户配置列表  QUOTA、PROFILE、POLICY
	ApiType *string `json:"ApiType,omitnil" name:"ApiType"`

	// 集群名称，GET_SYSTEM_USERS，GET_PRIVILEGE_USERS，GET_CLUSTER_DATABASES，GET_CLUSTER_TABLES 必填
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 用户名称，api与user相关的必填
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 账户的类型
	UserType *string `json:"UserType,omitnil" name:"UserType"`
}

func (r *DescribeCkSqlApisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkSqlApisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ApiType")
	delete(f, "Cluster")
	delete(f, "UserName")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkSqlApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkSqlApisResponseParams struct {
	// 返回的查询数据，大部分情况是list，也可能是bool
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnData *string `json:"ReturnData,omitnil" name:"ReturnData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCkSqlApisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCkSqlApisResponseParams `json:"Response"`
}

func (r *DescribeCkSqlApisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCkSqlApisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterConfigsResponseParams struct {
	// 返回实例的配置文件相关的信息
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil" name:"ClusterConfList"`

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
type DescribeInstanceClustersRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceClustersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceClustersResponseParams struct {
	// cluster列表
	Clusters []*ClusterInfo `json:"Clusters,omitnil" name:"Clusters"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceClustersResponseParams `json:"Response"`
}

func (r *DescribeInstanceClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceKeyValConfigsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 搜索的配置项名称
	SearchConfigName *string `json:"SearchConfigName,omitnil" name:"SearchConfigName"`
}

type DescribeInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 搜索的配置项名称
	SearchConfigName *string `json:"SearchConfigName,omitnil" name:"SearchConfigName"`
}

func (r *DescribeInstanceKeyValConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceKeyValConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchConfigName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceKeyValConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceKeyValConfigsResponseParams struct {
	// 参数列表
	ConfigItems []*InstanceConfigInfo `json:"ConfigItems,omitnil" name:"ConfigItems"`

	// 未配置的参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnConfigItems []*InstanceConfigInfo `json:"UnConfigItems,omitnil" name:"UnConfigItems"`

	// 配置的多层级参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapConfigItems []*MapConfigItem `json:"MapConfigItems,omitnil" name:"MapConfigItems"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceKeyValConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceKeyValConfigsResponseParams `json:"Response"`
}

func (r *DescribeInstanceKeyValConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceKeyValConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 是否是open api查询
	IsOpenApi *bool `json:"IsOpenApi,omitnil" name:"IsOpenApi"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 是否是open api查询
	IsOpenApi *bool `json:"IsOpenApi,omitnil" name:"IsOpenApi"`
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
	delete(f, "IsOpenApi")
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
type DescribeInstanceShardsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceShardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceShardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceShardsResponseParams struct {
	// 实例shard信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceShardsList *string `json:"InstanceShardsList,omitnil" name:"InstanceShardsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceShardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceShardsResponseParams `json:"Response"`
}

func (r *DescribeInstanceShardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceShardsResponse) FromJsonString(s string) error {
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
type DescribeInstancesNewRequestParams struct {
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

	// 信息详细与否
	IsSimple *bool `json:"IsSimple,omitnil" name:"IsSimple"`
}

type DescribeInstancesNewRequest struct {
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

	// 信息详细与否
	IsSimple *bool `json:"IsSimple,omitnil" name:"IsSimple"`
}

func (r *DescribeInstancesNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceId")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	delete(f, "IsSimple")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesNewResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例数组
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil" name:"InstancesList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesNewResponseParams `json:"Response"`
}

func (r *DescribeInstancesNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecRequestParams struct {
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitnil" name:"IsElastic"`
}

type DescribeSpecRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitnil" name:"IsElastic"`
}

func (r *DescribeSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "PayMode")
	delete(f, "IsElastic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecResponseParams struct {
	// zookeeper节点规格描述
	CommonSpec []*ResourceSpec `json:"CommonSpec,omitnil" name:"CommonSpec"`

	// 数据节点规格描述
	DataSpec []*ResourceSpec `json:"DataSpec,omitnil" name:"DataSpec"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec []*DiskSpec `json:"AttachCBSSpec,omitnil" name:"AttachCBSSpec"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecResponseParams `json:"Response"`
}

func (r *DescribeSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
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
	// 作业id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowID *string `json:"FlowID,omitnil" name:"FlowID"`

	// 集群id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

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

type DiskSpec struct {
	// 磁盘类型，例如“CLOUD_SSD", "LOCAL_SSD"等
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 磁盘类型说明，例如"云SSD", "本地SSD"等
	DiskDesc *string `json:"DiskDesc,omitnil" name:"DiskDesc"`

	// 磁盘最小规格大小，单位G
	MinDiskSize *int64 `json:"MinDiskSize,omitnil" name:"MinDiskSize"`

	// 磁盘最大规格大小，单位G
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil" name:"MaxDiskSize"`

	// 磁盘数目
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`
}

type InstanceConfigInfo struct {
	// 配置项名称
	ConfKey *string `json:"ConfKey,omitnil" name:"ConfKey"`

	// 配置项内容
	ConfValue *string `json:"ConfValue,omitnil" name:"ConfValue"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// 是否需要重启
	NeedRestart *bool `json:"NeedRestart,omitnil" name:"NeedRestart"`

	// 是否可编辑
	Editable *bool `json:"Editable,omitnil" name:"Editable"`

	// 配置项解释
	ConfDesc *string `json:"ConfDesc,omitnil" name:"ConfDesc"`

	// 文件名称
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 规则名称类型
	ModifyRuleType *string `json:"ModifyRuleType,omitnil" name:"ModifyRuleType"`

	// 规则名称内容
	ModifyRuleValue *string `json:"ModifyRuleValue,omitnil" name:"ModifyRuleValue"`

	// 修改人的uin
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type InstanceConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitnil" name:"ConfKey"`

	// value
	ConfValue *string `json:"ConfValue,omitnil" name:"ConfValue"`
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
	CommonSummary *NodesSummary `json:"CommonSummary,omitnil" name:"CommonSummary"`

	// 高可用，“true" "false"
	// 注意：此字段可能返回 null，表示取不到有效值。
	HA *string `json:"HA,omitnil" name:"HA"`

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

	// 是否弹性ck
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsElastic *bool `json:"IsElastic,omitnil" name:"IsElastic"`

	// 集群详细状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitnil" name:"InstanceStateInfo"`

	// ZK高可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	HAZk *bool `json:"HAZk,omitnil" name:"HAZk"`

	// 挂载盘,默认0:没有类型；1:裸盘;2:lvm
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountDiskType *int64 `json:"MountDiskType,omitnil" name:"MountDiskType"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	CHProxyVip *string `json:"CHProxyVip,omitnil" name:"CHProxyVip"`

	// cos buket的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// 是否可以挂载云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil" name:"CanAttachCbs"`

	// 是否可以挂载云盘阵列
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbsLvm *bool `json:"CanAttachCbsLvm,omitnil" name:"CanAttachCbsLvm"`

	// 是否可以挂载cos
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCos *bool `json:"CanAttachCos,omitnil" name:"CanAttachCos"`

	// 服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*ServiceInfo `json:"Components,omitnil" name:"Components"`

	// 可升级的内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeVersions *string `json:"UpgradeVersions,omitnil" name:"UpgradeVersions"`

	// ex-index
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsIndexId *string `json:"EsIndexId,omitnil" name:"EsIndexId"`

	// username
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsIndexUsername *string `json:"EsIndexUsername,omitnil" name:"EsIndexUsername"`

	// password
	// 注意：此字段可能返回 null，表示取不到有效值。
	EsIndexPassword *string `json:"EsIndexPassword,omitnil" name:"EsIndexPassword"`

	// true
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasEsIndex *bool `json:"HasEsIndex,omitnil" name:"HasEsIndex"`
}

type InstanceStateInfo struct {
	// 集群状态，例如：Serving
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitnil" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *int64 `json:"FlowProgress,omitnil" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitnil" name:"FlowMsg"`

	// 当前步骤的名称，例如：”购买资源中“
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessName *string `json:"ProcessName,omitnil" name:"ProcessName"`

	// 请求id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`

	// 流程的二级名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessSubName *string `json:"ProcessSubName,omitnil" name:"ProcessSubName"`
}

type MapConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitnil" name:"ConfKey"`

	// 列表
	Items []*InstanceConfigInfo `json:"Items,omitnil" name:"Items"`
}

// Predefined struct for user
type ModifyClusterConfigsRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitnil" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifyClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitnil" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifyClusterConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ModifyConfContext")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterConfigsResponseParams struct {
	// 流程相关信息
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyClusterConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterConfigsResponseParams `json:"Response"`
}

func (r *ModifyClusterConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceKeyValConfigsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil" name:"UpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil" name:"DeleteItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil" name:"DelItems"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifyInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil" name:"UpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil" name:"DeleteItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil" name:"DelItems"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifyInstanceKeyValConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceKeyValConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AddItems")
	delete(f, "UpdateItems")
	delete(f, "DeleteItems")
	delete(f, "DelItems")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceKeyValConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceKeyValConfigsResponseParams struct {
	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// ID
	FlowId *int64 `json:"FlowId,omitnil" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceKeyValConfigsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceKeyValConfigsResponseParams `json:"Response"`
}

func (r *ModifyInstanceKeyValConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceKeyValConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserNewPrivilegeRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// cluster名称
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 是否所有数据库表
	AllDatabase *bool `json:"AllDatabase,omitnil" name:"AllDatabase"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil" name:"GlobalPrivileges"`

	// 数据库表权限
	DatabasePrivilegeList []*DatabasePrivilegeInfo `json:"DatabasePrivilegeList,omitnil" name:"DatabasePrivilegeList"`
}

type ModifyUserNewPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// cluster名称
	Cluster *string `json:"Cluster,omitnil" name:"Cluster"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 是否所有数据库表
	AllDatabase *bool `json:"AllDatabase,omitnil" name:"AllDatabase"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil" name:"GlobalPrivileges"`

	// 数据库表权限
	DatabasePrivilegeList []*DatabasePrivilegeInfo `json:"DatabasePrivilegeList,omitnil" name:"DatabasePrivilegeList"`
}

func (r *ModifyUserNewPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserNewPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Cluster")
	delete(f, "UserName")
	delete(f, "AllDatabase")
	delete(f, "GlobalPrivileges")
	delete(f, "DatabasePrivilegeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserNewPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserNewPrivilegeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserNewPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserNewPrivilegeResponseParams `json:"Response"`
}

func (r *ModifyUserNewPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserNewPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeSpec struct {
	// 规格名称
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 数量
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 云盘大小
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
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

	// 子产品类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductType *string `json:"SubProductType,omitnil" name:"SubProductType"`

	// 规格对应的核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCore *int64 `json:"SpecCore,omitnil" name:"SpecCore"`

	// 规格对应的内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecMemory *int64 `json:"SpecMemory,omitnil" name:"SpecMemory"`

	// 磁盘的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`

	// 磁盘的最大大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil" name:"MaxDiskSize"`

	// 是否为加密云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *int64 `json:"Encrypt,omitnil" name:"Encrypt"`
}

// Predefined struct for user
type OpenBackUpRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`
}

type OpenBackUpRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`
}

func (r *OpenBackUpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenBackUpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OperationType")
	delete(f, "CosBucketName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenBackUpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenBackUpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type OpenBackUpResponse struct {
	*tchttp.BaseResponse
	Response *OpenBackUpResponseParams `json:"Response"`
}

func (r *OpenBackUpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenBackUpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverBackUpJobRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil" name:"BackUpJobId"`
}

type RecoverBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil" name:"BackUpJobId"`
}

func (r *RecoverBackUpJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverBackUpJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackUpJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverBackUpJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverBackUpJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecoverBackUpJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverBackUpJobResponseParams `json:"Response"`
}

func (r *RecoverBackUpJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverBackUpJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeDiskRequestParams struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil" name:"Type"`

	// 磁盘扩容后容量，不能小于原有用量。clickhouse最小200，且为100的整数倍。 zk最小100，且为10的整数倍；
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil" name:"Type"`

	// 磁盘扩容后容量，不能小于原有用量。clickhouse最小200，且为100的整数倍。 zk最小100，且为10的整数倍；
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`
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
	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type ResourceSpec struct {
	// 规格名称，例如“SCH1"
	Name *string `json:"Name,omitnil" name:"Name"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil" name:"Cpu"`

	// 内存大小，单位G
	Mem *int64 `json:"Mem,omitnil" name:"Mem"`

	// 分类标记，STANDARD/BIGDATA/HIGHIO分别表示标准型/大数据型/高IO
	Type *string `json:"Type,omitnil" name:"Type"`

	// 系统盘描述信息
	SystemDisk *DiskSpec `json:"SystemDisk,omitnil" name:"SystemDisk"`

	// 数据盘描述信息
	DataDisk *DiskSpec `json:"DataDisk,omitnil" name:"DataDisk"`

	// 最大节点数目限制
	MaxNodeSize *int64 `json:"MaxNodeSize,omitnil" name:"MaxNodeSize"`

	// 是否可用，false代表售罄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *bool `json:"Available,omitnil" name:"Available"`

	// 规格描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeSpecDesc *string `json:"ComputeSpecDesc,omitnil" name:"ComputeSpecDesc"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil" name:"DisplayName"`

	// 库存数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceQuota *int64 `json:"InstanceQuota,omitnil" name:"InstanceQuota"`
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil" name:"Type"`

	// 调整clickhouse节点数量
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// v_cluster分组，	
	// 新增扩容节点将加入到已选择的v_cluster分组中，提交同步VIP生效.
	ScaleOutCluster *string `json:"ScaleOutCluster,omitnil" name:"ScaleOutCluster"`

	// 子网剩余ip数量，用于判断当前实例子网剩余ip数是否能扩容。需要根据实际填写
	UserSubnetIPNum *int64 `json:"UserSubnetIPNum,omitnil" name:"UserSubnetIPNum"`

	// 同步元数据节点IP （uip），扩容的时候必填
	ScaleOutNodeIp *string `json:"ScaleOutNodeIp,omitnil" name:"ScaleOutNodeIp"`

	// 缩容节点shard的节点IP （uip），其中ha集群需要主副节点ip都传入以逗号分隔，缩容的时候必填
	ReduceShardInfo []*string `json:"ReduceShardInfo,omitnil" name:"ReduceShardInfo"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil" name:"Type"`

	// 调整clickhouse节点数量
	NodeCount *int64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// v_cluster分组，	
	// 新增扩容节点将加入到已选择的v_cluster分组中，提交同步VIP生效.
	ScaleOutCluster *string `json:"ScaleOutCluster,omitnil" name:"ScaleOutCluster"`

	// 子网剩余ip数量，用于判断当前实例子网剩余ip数是否能扩容。需要根据实际填写
	UserSubnetIPNum *int64 `json:"UserSubnetIPNum,omitnil" name:"UserSubnetIPNum"`

	// 同步元数据节点IP （uip），扩容的时候必填
	ScaleOutNodeIp *string `json:"ScaleOutNodeIp,omitnil" name:"ScaleOutNodeIp"`

	// 缩容节点shard的节点IP （uip），其中ha集群需要主副节点ip都传入以逗号分隔，缩容的时候必填
	ReduceShardInfo []*string `json:"ReduceShardInfo,omitnil" name:"ReduceShardInfo"`
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
	delete(f, "ScaleOutCluster")
	delete(f, "UserSubnetIPNum")
	delete(f, "ScaleOutNodeIp")
	delete(f, "ReduceShardInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceResponseParams struct {
	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil" name:"Type"`

	// clickhouse节点规格。
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 是否滚动重启，false为不滚动重启，true为滚动重启
	ScaleUpEnableRolling *bool `json:"ScaleUpEnableRolling,omitnil" name:"ScaleUpEnableRolling"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil" name:"Type"`

	// clickhouse节点规格。
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 是否滚动重启，false为不滚动重启，true为滚动重启
	ScaleUpEnableRolling *bool `json:"ScaleUpEnableRolling,omitnil" name:"ScaleUpEnableRolling"`
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
	delete(f, "Type")
	delete(f, "SpecName")
	delete(f, "ScaleUpEnableRolling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleUpInstanceResponseParams struct {
	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
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

type ScheduleStrategy struct {
	// 备份桶列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil" name:"CosBucketName"`

	// 备份保留天数
	RetainDays *int64 `json:"RetainDays,omitnil" name:"RetainDays"`

	// 备份的天
	WeekDays *string `json:"WeekDays,omitnil" name:"WeekDays"`

	// 备份小时
	ExecuteHour *int64 `json:"ExecuteHour,omitnil" name:"ExecuteHour"`

	// 策略id
	ScheduleId *int64 `json:"ScheduleId,omitnil" name:"ScheduleId"`

	// 下次备份时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextBackupTime *string `json:"NextBackupTime,omitnil" name:"NextBackupTime"`
}

type SearchTags struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil" name:"AllValue"`
}

type ServiceInfo struct {
	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 服务的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`
}

type TablePrivilegeInfo struct {
	// 表名称
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 表权限列表 SELECT、INSERT_ALL、ALTER、TRUNCATE、DROP_TABLE 查询、插入、设置、清空表、删除表
	TablePrivileges []*string `json:"TablePrivileges,omitnil" name:"TablePrivileges"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}