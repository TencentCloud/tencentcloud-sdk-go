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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ActionAlterCkUserRequestParams struct {
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// api接口类型，
	// AddSystemUser新增用户，UpdateSystemUser，修改用户
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
}

type ActionAlterCkUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// api接口类型，
	// AddSystemUser新增用户，UpdateSystemUser，修改用户
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`
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
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘容量，单位G
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 磁盘总数
	DiskCount *int64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 描述
	DiskDesc *string `json:"DiskDesc,omitempty" name:"DiskDesc"`
}

type BackupTableContent struct {
	// 数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// 表
	Table *string `json:"Table,omitempty" name:"Table"`

	// 表总字节数
	TotalBytes *int64 `json:"TotalBytes,omitempty" name:"TotalBytes"`

	// 虚拟cluster
	VCluster *string `json:"VCluster,omitempty" name:"VCluster"`

	// 表ip
	Ips *string `json:"Ips,omitempty" name:"Ips"`
}

type Charge struct {
	// 计费类型，“PREPAID” 预付费，“POSTPAID_BY_HOUR” 后付费
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// PREPAID需要传递，是否自动续费，1表示自动续费开启
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 预付费需要传递，计费时间长度，多少个月
	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

type CkUserAlterInfo struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	PassWord *string `json:"PassWord,omitempty" name:"PassWord"`

	// 描述
	Describe *string `json:"Describe,omitempty" name:"Describe"`
}

type ClusterConfigsInfoFromEMR struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 配置文件对应的相关属性信息
	FileConf *string `json:"FileConf,omitempty" name:"FileConf"`

	// 配置文件对应的其他属性信息
	KeyConf *string `json:"KeyConf,omitempty" name:"KeyConf"`

	// 配置文件的内容，base64编码
	OriParam *string `json:"OriParam,omitempty" name:"OriParam"`

	// 用于表示当前配置文件是不是有过修改后没有重启，提醒用户需要重启
	NeedRestart *int64 `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 保存配置文件的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
}

type ClusterInfo struct {
	// vcluster名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 当前cluster的IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeIps []*string `json:"NodeIps,omitempty" name:"NodeIps"`
}

type ConfigSubmitContext struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 配置文件旧内容，base64编码
	OldConfValue *string `json:"OldConfValue,omitempty" name:"OldConfValue"`

	// 配置文件新内容，base64编码
	NewConfValue *string `json:"NewConfValue,omitempty" name:"NewConfValue"`

	// 保存配置文件的路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`
}

// Predefined struct for user
type CreateBackUpScheduleRequestParams struct {
	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔，例如 2 代表周二
	WeekDays *string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitempty" name:"BackUpTables"`
}

type CreateBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔，例如 2 代表周二
	WeekDays *string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitempty" name:"BackUpTables"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitempty" name:"HaFlag"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitempty" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitempty" name:"UserSubnetId"`

	// 版本
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`

	// 计费方式
	ChargeProperties *Charge `json:"ChargeProperties,omitempty" name:"ChargeProperties"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 数据节点
	DataSpec *NodeSpec `json:"DataSpec,omitempty" name:"DataSpec"`

	// 标签列表
	Tags *Tag `json:"Tags,omitempty" name:"Tags"`

	// 日志主题ID
	ClsLogSetId *string `json:"ClsLogSetId,omitempty" name:"ClsLogSetId"`

	// COS桶名称
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 是否是裸盘挂载
	MountDiskType *int64 `json:"MountDiskType,omitempty" name:"MountDiskType"`

	// 是否是ZK高可用
	HAZk *bool `json:"HAZk,omitempty" name:"HAZk"`

	// ZK节点
	CommonSpec *NodeSpec `json:"CommonSpec,omitempty" name:"CommonSpec"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitempty" name:"HaFlag"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitempty" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitempty" name:"UserSubnetId"`

	// 版本
	ProductVersion *string `json:"ProductVersion,omitempty" name:"ProductVersion"`

	// 计费方式
	ChargeProperties *Charge `json:"ChargeProperties,omitempty" name:"ChargeProperties"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 数据节点
	DataSpec *NodeSpec `json:"DataSpec,omitempty" name:"DataSpec"`

	// 标签列表
	Tags *Tag `json:"Tags,omitempty" name:"Tags"`

	// 日志主题ID
	ClsLogSetId *string `json:"ClsLogSetId,omitempty" name:"ClsLogSetId"`

	// COS桶名称
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 是否是裸盘挂载
	MountDiskType *int64 `json:"MountDiskType,omitempty" name:"MountDiskType"`

	// 是否是ZK高可用
	HAZk *bool `json:"HAZk,omitempty" name:"HAZk"`

	// ZK节点
	CommonSpec *NodeSpec `json:"CommonSpec,omitempty" name:"CommonSpec"`
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
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeBackUpScheduleRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	BackUpOpened *bool `json:"BackUpOpened,omitempty" name:"BackUpOpened"`

	// 元数据备份策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaStrategy *ScheduleStrategy `json:"MetaStrategy,omitempty" name:"MetaStrategy"`

	// 表数据备份策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataStrategy *ScheduleStrategy `json:"DataStrategy,omitempty" name:"DataStrategy"`

	// 备份表列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackUpContents []*BackupTableContent `json:"BackUpContents,omitempty" name:"BackUpContents"`

	// 备份的状态
	BackUpStatus *int64 `json:"BackUpStatus,omitempty" name:"BackUpStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCkSqlApisRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

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
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 集群名称，GET_SYSTEM_USERS，GET_PRIVILEGE_USERS，GET_CLUSTER_DATABASES，GET_CLUSTER_TABLES 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 用户名称，api与user相关的必填
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DescribeCkSqlApisRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

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
	ApiType *string `json:"ApiType,omitempty" name:"ApiType"`

	// 集群名称，GET_SYSTEM_USERS，GET_PRIVILEGE_USERS，GET_CLUSTER_DATABASES，GET_CLUSTER_TABLES 必填
	Cluster *string `json:"Cluster,omitempty" name:"Cluster"`

	// 用户名称，api与user相关的必填
	UserName *string `json:"UserName,omitempty" name:"UserName"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCkSqlApisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkSqlApisResponseParams struct {
	// 返回的查询数据，大部分情况是list，也可能是bool
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnData *string `json:"ReturnData,omitempty" name:"ReturnData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitempty" name:"ClusterConfList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceClustersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 搜索的配置项名称
	SearchConfigName *string `json:"SearchConfigName,omitempty" name:"SearchConfigName"`
}

type DescribeInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 搜索的配置项名称
	SearchConfigName *string `json:"SearchConfigName,omitempty" name:"SearchConfigName"`
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
	ConfigItems []*InstanceConfigInfo `json:"ConfigItems,omitempty" name:"ConfigItems"`

	// 未配置的参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnConfigItems []*InstanceConfigInfo `json:"UnConfigItems,omitempty" name:"UnConfigItems"`

	// 配置的多层级参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapConfigItems []*MapConfigItem `json:"MapConfigItems,omitempty" name:"MapConfigItems"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	InstanceInfo *InstanceInfo `json:"InstanceInfo,omitempty" name:"InstanceInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	InstanceShardsList *string `json:"InstanceShardsList,omitempty" name:"InstanceShardsList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSpecRequestParams struct {
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitempty" name:"IsElastic"`
}

type DescribeSpecRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitempty" name:"IsElastic"`
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
	CommonSpec []*ResourceSpec `json:"CommonSpec,omitempty" name:"CommonSpec"`

	// 数据节点规格描述
	DataSpec []*ResourceSpec `json:"DataSpec,omitempty" name:"DataSpec"`

	// 云盘列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec []*DiskSpec `json:"AttachCBSSpec,omitempty" name:"AttachCBSSpec"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	FlowID *string `json:"FlowID,omitempty" name:"FlowID"`

	// 集群id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘类型说明，例如"云SSD", "本地SSD"等
	DiskDesc *string `json:"DiskDesc,omitempty" name:"DiskDesc"`

	// 磁盘最小规格大小，单位G
	MinDiskSize *int64 `json:"MinDiskSize,omitempty" name:"MinDiskSize"`

	// 磁盘最大规格大小，单位G
	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// 磁盘数目
	DiskCount *int64 `json:"DiskCount,omitempty" name:"DiskCount"`
}

type InstanceConfigInfo struct {
	// 配置项名称
	ConfKey *string `json:"ConfKey,omitempty" name:"ConfKey"`

	// 配置项内容
	ConfValue *string `json:"ConfValue,omitempty" name:"ConfValue"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`

	// 是否需要重启
	NeedRestart *bool `json:"NeedRestart,omitempty" name:"NeedRestart"`

	// 是否可编辑
	Editable *bool `json:"Editable,omitempty" name:"Editable"`

	// 配置项解释
	ConfDesc *string `json:"ConfDesc,omitempty" name:"ConfDesc"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 规则名称类型
	ModifyRuleType *string `json:"ModifyRuleType,omitempty" name:"ModifyRuleType"`

	// 规则名称内容
	ModifyRuleValue *string `json:"ModifyRuleValue,omitempty" name:"ModifyRuleValue"`

	// 修改人的uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type InstanceConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitempty" name:"ConfKey"`

	// value
	ConfValue *string `json:"ConfValue,omitempty" name:"ConfValue"`
}

type InstanceInfo struct {
	// 集群实例ID, "cdw-xxxx" 字符串类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 集群实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 状态,
	// Init 创建中; Serving 运行中； 
	// Deleted已销毁；Deleting 销毁中；
	// Modify 集群变更中；
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 地域, ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 可用区， ap-guangzhou-3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 付费类型，"hour", "prepay"
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 数据节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSummary *NodesSummary `json:"MasterSummary,omitempty" name:"MasterSummary"`

	// zookeeper节点描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonSummary *NodesSummary `json:"CommonSummary,omitempty" name:"CommonSummary"`

	// 高可用，“true" "false"
	// 注意：此字段可能返回 null，表示取不到有效值。
	HA *string `json:"HA,omitempty" name:"HA"`

	// 访问地址，例如 "10.0.0.1:9000"
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfo *string `json:"AccessInfo,omitempty" name:"AccessInfo"`

	// 记录ID，数值型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// regionId, 表示地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 可用区说明，例如 "广州二区"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneDesc *string `json:"ZoneDesc,omitempty" name:"ZoneDesc"`

	// 错误流程说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitempty" name:"FlowMsg"`

	// 状态描述，例如“运行中”等
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 自动续费标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *bool `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 监控信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Monitor *string `json:"Monitor,omitempty" name:"Monitor"`

	// 是否开通日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasClsTopic *bool `json:"HasClsTopic,omitempty" name:"HasClsTopic"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsTopicId *string `json:"ClsTopicId,omitempty" name:"ClsTopicId"`

	// 日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClsLogSetId *string `json:"ClsLogSetId,omitempty" name:"ClsLogSetId"`

	// 是否支持xml配置管理
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitempty" name:"EnableXMLConfig"`

	// 区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDesc *string `json:"RegionDesc,omitempty" name:"RegionDesc"`

	// 弹性网卡地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 冷热分层系数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosMoveFactor *int64 `json:"CosMoveFactor,omitempty" name:"CosMoveFactor"`

	// external/local/yunti
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 是否弹性ck
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsElastic *bool `json:"IsElastic,omitempty" name:"IsElastic"`

	// 集群详细状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitempty" name:"InstanceStateInfo"`

	// ZK高可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	HAZk *bool `json:"HAZk,omitempty" name:"HAZk"`

	// 挂载盘,默认0:没有类型；1:裸盘;2:lvm
	// 注意：此字段可能返回 null，表示取不到有效值。
	MountDiskType *int64 `json:"MountDiskType,omitempty" name:"MountDiskType"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	CHProxyVip *string `json:"CHProxyVip,omitempty" name:"CHProxyVip"`

	// cos buket的名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 是否可以挂载云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbs *bool `json:"CanAttachCbs,omitempty" name:"CanAttachCbs"`

	// 是否可以挂载云盘阵列
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCbsLvm *bool `json:"CanAttachCbsLvm,omitempty" name:"CanAttachCbsLvm"`

	// 是否可以挂载cos
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAttachCos *bool `json:"CanAttachCos,omitempty" name:"CanAttachCos"`

	// 服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Components []*ServiceInfo `json:"Components,omitempty" name:"Components"`

	// 可升级的内核版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpgradeVersions *string `json:"UpgradeVersions,omitempty" name:"UpgradeVersions"`
}

type InstanceStateInfo struct {
	// 集群状态，例如：Serving
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`

	// 集群操作创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowCreateTime *string `json:"FlowCreateTime,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 集群操作进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowProgress *int64 `json:"FlowProgress,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStateDesc *string `json:"InstanceStateDesc,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowMsg *string `json:"FlowMsg,omitempty" name:"FlowMsg"`

	// 当前步骤的名称，例如：”购买资源中“
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessName *string `json:"ProcessName,omitempty" name:"ProcessName"`

	// 请求id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MapConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitempty" name:"ConfKey"`

	// 列表
	Items []*InstanceConfigInfo `json:"Items,omitempty" name:"Items"`
}

// Predefined struct for user
type ModifyClusterConfigsRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitempty" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitempty" name:"UpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitempty" name:"DeleteItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitempty" name:"DelItems"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitempty" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitempty" name:"UpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitempty" name:"DeleteItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitempty" name:"DelItems"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

}

type ModifyUserNewPrivilegeRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserNewPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserNewPrivilegeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 数量
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 云盘大小
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type NodesSummary struct {
	// 机型，如 S1
	Spec *string `json:"Spec,omitempty" name:"Spec"`

	// 节点数目
	NodeSize *int64 `json:"NodeSize,omitempty" name:"NodeSize"`

	// cpu核数，单位个
	Core *int64 `json:"Core,omitempty" name:"Core"`

	// 内存大小，单位G
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 磁盘大小，单位G
	Disk *int64 `json:"Disk,omitempty" name:"Disk"`

	// 磁盘类型
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 磁盘描述
	DiskDesc *string `json:"DiskDesc,omitempty" name:"DiskDesc"`

	// 挂载云盘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitempty" name:"AttachCBSSpec"`

	// 子产品类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductType *string `json:"SubProductType,omitempty" name:"SubProductType"`

	// 规格对应的核数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecCore *int64 `json:"SpecCore,omitempty" name:"SpecCore"`

	// 规格对应的内存大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecMemory *int64 `json:"SpecMemory,omitempty" name:"SpecMemory"`

	// 磁盘的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskCount *int64 `json:"DiskCount,omitempty" name:"DiskCount"`

	// 磁盘的最大大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`

	// 是否为加密云盘
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *int64 `json:"Encrypt,omitempty" name:"Encrypt"`
}

// Predefined struct for user
type OpenBackUpRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`
}

type OpenBackUpRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ResizeDiskRequestParams struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitempty" name:"Type"`

	// 磁盘扩容后容量，不能小于原有用量。clickhouse最小200，且为100的整数倍。 zk最小100，且为10的整数倍；
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitempty" name:"Type"`

	// 磁盘扩容后容量，不能小于原有用量。clickhouse最小200，且为100的整数倍。 zk最小100，且为10的整数倍；
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
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
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存大小，单位G
	Mem *int64 `json:"Mem,omitempty" name:"Mem"`

	// 分类标记，STANDARD/BIGDATA/HIGHIO分别表示标准型/大数据型/高IO
	Type *string `json:"Type,omitempty" name:"Type"`

	// 系统盘描述信息
	SystemDisk *DiskSpec `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 数据盘描述信息
	DataDisk *DiskSpec `json:"DataDisk,omitempty" name:"DataDisk"`

	// 最大节点数目限制
	MaxNodeSize *int64 `json:"MaxNodeSize,omitempty" name:"MaxNodeSize"`

	// 是否可用，false代表售罄
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *bool `json:"Available,omitempty" name:"Available"`

	// 规格描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeSpecDesc *string `json:"ComputeSpecDesc,omitempty" name:"ComputeSpecDesc"`

	// 规格名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 库存数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceQuota *int64 `json:"InstanceQuota,omitempty" name:"InstanceQuota"`
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitempty" name:"Type"`

	// 调整clickhouse节点数量
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// v_cluster分组，	
	// 新增扩容节点将加入到已选择的v_cluster分组中，提交同步VIP生效.
	ScaleOutCluster *string `json:"ScaleOutCluster,omitempty" name:"ScaleOutCluster"`

	// 子网剩余ip数量，用于判断当前实例子网剩余ip数是否能扩容。需要根据实际填写
	UserSubnetIPNum *int64 `json:"UserSubnetIPNum,omitempty" name:"UserSubnetIPNum"`

	// 同步元数据节点IP （uip），扩容的时候必填
	ScaleOutNodeIp *string `json:"ScaleOutNodeIp,omitempty" name:"ScaleOutNodeIp"`

	// 缩容节点shard的节点IP （uip），其中ha集群需要主副节点ip都传入以逗号分隔，缩容的时候必填
	ReduceShardInfo []*string `json:"ReduceShardInfo,omitempty" name:"ReduceShardInfo"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitempty" name:"Type"`

	// 调整clickhouse节点数量
	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`

	// v_cluster分组，	
	// 新增扩容节点将加入到已选择的v_cluster分组中，提交同步VIP生效.
	ScaleOutCluster *string `json:"ScaleOutCluster,omitempty" name:"ScaleOutCluster"`

	// 子网剩余ip数量，用于判断当前实例子网剩余ip数是否能扩容。需要根据实际填写
	UserSubnetIPNum *int64 `json:"UserSubnetIPNum,omitempty" name:"UserSubnetIPNum"`

	// 同步元数据节点IP （uip），扩容的时候必填
	ScaleOutNodeIp *string `json:"ScaleOutNodeIp,omitempty" name:"ScaleOutNodeIp"`

	// 缩容节点shard的节点IP （uip），其中ha集群需要主副节点ip都传入以逗号分隔，缩容的时候必填
	ReduceShardInfo []*string `json:"ReduceShardInfo,omitempty" name:"ReduceShardInfo"`
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
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitempty" name:"Type"`

	// clickhouse节点规格。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 是否滚动重启，false为不滚动重启，true为滚动重启
	ScaleUpEnableRolling *bool `json:"ScaleUpEnableRolling,omitempty" name:"ScaleUpEnableRolling"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitempty" name:"Type"`

	// clickhouse节点规格。
	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`

	// 是否滚动重启，false为不滚动重启，true为滚动重启
	ScaleUpEnableRolling *bool `json:"ScaleUpEnableRolling,omitempty" name:"ScaleUpEnableRolling"`
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
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 备份保留天数
	RetainDays *int64 `json:"RetainDays,omitempty" name:"RetainDays"`

	// 备份的天
	WeekDays *string `json:"WeekDays,omitempty" name:"WeekDays"`

	// 备份小时
	ExecuteHour *int64 `json:"ExecuteHour,omitempty" name:"ExecuteHour"`

	// 策略id
	ScheduleId *int64 `json:"ScheduleId,omitempty" name:"ScheduleId"`
}

type ServiceInfo struct {
	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}