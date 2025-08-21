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

package v20200915

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ActionAlterCkUserRequestParams struct {
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// api接口类型，
	// AddSystemUser新增用户，UpdateSystemUser，修改用户
	ApiType *string `json:"ApiType,omitnil,omitempty" name:"ApiType"`
}

type ActionAlterCkUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户信息
	UserInfo *CkUserAlterInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// api接口类型，
	// AddSystemUser新增用户，UpdateSystemUser，修改用户
	ApiType *string `json:"ApiType,omitnil,omitempty" name:"ApiType"`
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
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘容量，单位G
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 磁盘总数
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 描述
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`
}

type BackUpJobDisplay struct {
	// 备份任务id
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 备份任务名
	Snapshot *string `json:"Snapshot,omitnil,omitempty" name:"Snapshot"`

	// 任务类型(元数据),(数据)
	BackUpType *string `json:"BackUpType,omitnil,omitempty" name:"BackUpType"`

	// 备份数据量
	BackUpSize *int64 `json:"BackUpSize,omitnil,omitempty" name:"BackUpSize"`

	// 任务创建时间
	BackUpTime *string `json:"BackUpTime,omitnil,omitempty" name:"BackUpTime"`

	// 任务过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 任务状态
	JobStatus *string `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// 处理数据量
	ProcessSize *int64 `json:"ProcessSize,omitnil,omitempty" name:"ProcessSize"`

	// 错误原因
	ErrorReason *string `json:"ErrorReason,omitnil,omitempty" name:"ErrorReason"`
}

type BackupTableContent struct {
	// 数据库
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 表总字节数
	TotalBytes *int64 `json:"TotalBytes,omitnil,omitempty" name:"TotalBytes"`

	// 虚拟cluster
	VCluster *string `json:"VCluster,omitnil,omitempty" name:"VCluster"`

	// 表ip
	Ips *string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// zk路径
	ZooPath *string `json:"ZooPath,omitnil,omitempty" name:"ZooPath"`

	// cvm的ip地址
	Rip *string `json:"Rip,omitnil,omitempty" name:"Rip"`
}

type CNResource struct {
	// 资源id
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 集群的id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 用户appid
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 组件
	Component *string `json:"Component,omitnil,omitempty" name:"Component"`

	// 部署模式
	DeployMode *int64 `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 规格名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 资源id
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// 资源的状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有网络ip
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 核数
	CPU *uint64 `json:"CPU,omitnil,omitempty" name:"CPU"`

	// 内存
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 存储大小
	Storage *uint64 `json:"Storage,omitnil,omitempty" name:"Storage"`

	// 服务器ID
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 详细信息
	Details *string `json:"Details,omitnil,omitempty" name:"Details"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type Charge struct {
	// 计费类型，“PREPAID” 预付费，“POSTPAID_BY_HOUR” 后付费
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// PREPAID需要传递，是否自动续费，1表示自动续费开启
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 预付费需要传递，计费时间长度，多少个月
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type CkUserAlterInfo struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// base64加密后的密码
	PassWord *string `json:"PassWord,omitnil,omitempty" name:"PassWord"`

	// 描述
	Describe *string `json:"Describe,omitnil,omitempty" name:"Describe"`
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

	// 保存配置文件的路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type ClusterInfo struct {
	// vcluster名字
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 当前cluster的IP列表
	NodeIps []*string `json:"NodeIps,omitnil,omitempty" name:"NodeIps"`
}

type CnInstanceInfo struct {
	// ID值
	ID *int64 `json:"ID,omitnil,omitempty" name:"ID"`

	// cdwch-cn或者其他
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// cdwch-cn或者其他
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// Running
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 运行中
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 无
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitnil,omitempty" name:"InstanceStateInfo"`

	// -
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 无
	Resources []*CNResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type ConfigSubmitContext struct {
	// 配置文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件旧内容，base64编码
	OldConfValue *string `json:"OldConfValue,omitnil,omitempty" name:"OldConfValue"`

	// 配置文件新内容，base64编码
	NewConfValue *string `json:"NewConfValue,omitnil,omitempty" name:"NewConfValue"`

	// 保存配置文件的路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

// Predefined struct for user
type CreateBackUpScheduleRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略类型 meta(元数据)  data (表数据)
	ScheduleType *string `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 操作类型 create(创建) update(编辑修改)
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 保留天数 例如7
	RetainDays *int64 `json:"RetainDays,omitnil,omitempty" name:"RetainDays"`

	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔，例如 2 代表周二
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil,omitempty" name:"BackUpTables"`
}

type CreateBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略类型 meta(元数据)  data (表数据)
	ScheduleType *string `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 操作类型 create(创建) update(编辑修改)
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 保留天数 例如7
	RetainDays *int64 `json:"RetainDays,omitnil,omitempty" name:"RetainDays"`

	// 编辑时需要传
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// 选择的星期 逗号分隔，例如 2 代表周二
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 执行小时
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// 备份表列表
	BackUpTables []*BackupTableContent `json:"BackUpTables,omitnil,omitempty" name:"BackUpTables"`
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
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 系统版本
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 计费方式
	ChargeProperties *Charge `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据节点
	// SpecName从DescribeSpec接口中返回的DataSpec.Name获取
	DataSpec *NodeSpec `json:"DataSpec,omitnil,omitempty" name:"DataSpec"`

	// 标签列表（废弃）
	//
	// Deprecated: Tags is deprecated.
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 日志主题ID
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// COS桶名称
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// 是否是裸盘挂载，默认值 0 为 未挂载，1 为挂载。
	MountDiskType *int64 `json:"MountDiskType,omitnil,omitempty" name:"MountDiskType"`

	// 是否是ZK高可用
	HAZk *bool `json:"HAZk,omitnil,omitempty" name:"HAZk"`

	// ZK节点SpecName从DescribeSpec接口中返回的CommonSpec结构体的Name（ZK节点）获取
	CommonSpec *NodeSpec `json:"CommonSpec,omitnil,omitempty" name:"CommonSpec"`

	// 标签列表
	TagItems []*Tag `json:"TagItems,omitnil,omitempty" name:"TagItems"`

	// 副可用区信息
	SecondaryZoneInfo []*SecondaryZoneInfo `json:"SecondaryZoneInfo,omitnil,omitempty" name:"SecondaryZoneInfo"`

	// default账号登陆实例的密码。8-16个字符，至少包含大写字母、小写字母、数字和特殊字符!@#%^*中的三种，第一个字符不能为特殊字符
	CkDefaultUserPwd *string `json:"CkDefaultUserPwd,omitnil,omitempty" name:"CkDefaultUserPwd"`
}

type CreateInstanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 是否高可用
	HaFlag *bool `json:"HaFlag,omitnil,omitempty" name:"HaFlag"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil,omitempty" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 系统版本
	ProductVersion *string `json:"ProductVersion,omitnil,omitempty" name:"ProductVersion"`

	// 计费方式
	ChargeProperties *Charge `json:"ChargeProperties,omitnil,omitempty" name:"ChargeProperties"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据节点
	// SpecName从DescribeSpec接口中返回的DataSpec.Name获取
	DataSpec *NodeSpec `json:"DataSpec,omitnil,omitempty" name:"DataSpec"`

	// 标签列表（废弃）
	Tags *Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 日志主题ID
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// COS桶名称
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// 是否是裸盘挂载，默认值 0 为 未挂载，1 为挂载。
	MountDiskType *int64 `json:"MountDiskType,omitnil,omitempty" name:"MountDiskType"`

	// 是否是ZK高可用
	HAZk *bool `json:"HAZk,omitnil,omitempty" name:"HAZk"`

	// ZK节点SpecName从DescribeSpec接口中返回的CommonSpec结构体的Name（ZK节点）获取
	CommonSpec *NodeSpec `json:"CommonSpec,omitnil,omitempty" name:"CommonSpec"`

	// 标签列表
	TagItems []*Tag `json:"TagItems,omitnil,omitempty" name:"TagItems"`

	// 副可用区信息
	SecondaryZoneInfo []*SecondaryZoneInfo `json:"SecondaryZoneInfo,omitnil,omitempty" name:"SecondaryZoneInfo"`

	// default账号登陆实例的密码。8-16个字符，至少包含大写字母、小写字母、数字和特殊字符!@#%^*中的三种，第一个字符不能为特殊字符
	CkDefaultUserPwd *string `json:"CkDefaultUserPwd,omitnil,omitempty" name:"CkDefaultUserPwd"`
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
	delete(f, "TagItems")
	delete(f, "SecondaryZoneInfo")
	delete(f, "CkDefaultUserPwd")
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

type DatabasePrivilegeInfo struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 库表权限，SELECT、INSERT_ALL、ALTER、TRUNCATE、DROP_TABLE、CREATE_TABLE、DROP_DATABASE
	DatabasePrivileges []*string `json:"DatabasePrivileges,omitnil,omitempty" name:"DatabasePrivileges"`

	// 库下面的表权限
	TablePrivilegeList []*TablePrivilegeInfo `json:"TablePrivilegeList,omitnil,omitempty" name:"TablePrivilegeList"`
}

// Predefined struct for user
type DeleteBackUpDataRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// 是否删除所有数据
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
}

type DeleteBackUpDataRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`

	// 是否删除所有数据
	IsDeleteAll *bool `json:"IsDeleteAll,omitnil,omitempty" name:"IsDeleteAll"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type DescribeBackUpJobDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
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
	TableContents []*BackupTableContent `json:"TableContents,omitnil,omitempty" name:"TableContents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页号
	PageNum *int64 `json:"PageNum,omitnil,omitempty" name:"PageNum"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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
	BackUpJobs []*BackUpJobDisplay `json:"BackUpJobs,omitnil,omitempty" name:"BackUpJobs"`

	// 错误描述
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackUpScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	BackUpOpened *bool `json:"BackUpOpened,omitnil,omitempty" name:"BackUpOpened"`

	// 元数据备份策略
	MetaStrategy *ScheduleStrategy `json:"MetaStrategy,omitnil,omitempty" name:"MetaStrategy"`

	// 表数据备份策略
	DataStrategy *ScheduleStrategy `json:"DataStrategy,omitnil,omitempty" name:"DataStrategy"`

	// 备份表列表
	BackUpContents []*BackupTableContent `json:"BackUpContents,omitnil,omitempty" name:"BackUpContents"`

	// 备份的状态
	BackUpStatus *int64 `json:"BackUpStatus,omitnil,omitempty" name:"BackUpStatus"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeBackUpTablesRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	AvailableTables []*BackupTableContent `json:"AvailableTables,omitnil,omitempty" name:"AvailableTables"`

	// 错误描述
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCNInstancesRequestParams struct {
	// 搜索的集群id名称
	SearchInstanceID *string `json:"SearchInstanceID,omitnil,omitempty" name:"SearchInstanceID"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`

	// 集群类型，弹性版或自研数仓版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 组件名称列表
	Components []*string `json:"Components,omitnil,omitempty" name:"Components"`
}

type DescribeCNInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 搜索的集群id名称
	SearchInstanceID *string `json:"SearchInstanceID,omitnil,omitempty" name:"SearchInstanceID"`

	// 搜索的集群name
	SearchInstanceName *string `json:"SearchInstanceName,omitnil,omitempty" name:"SearchInstanceName"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索标签列表
	SearchTags []*SearchTags `json:"SearchTags,omitnil,omitempty" name:"SearchTags"`

	// 集群类型，弹性版或自研数仓版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 组件名称列表
	Components []*string `json:"Components,omitnil,omitempty" name:"Components"`
}

func (r *DescribeCNInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCNInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchInstanceID")
	delete(f, "SearchInstanceName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchTags")
	delete(f, "InstanceType")
	delete(f, "Components")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCNInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCNInstancesResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例数组
	InstancesList []*CnInstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCNInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCNInstancesResponseParams `json:"Response"`
}

func (r *DescribeCNInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCNInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCkSqlApisRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

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
	ApiType *string `json:"ApiType,omitnil,omitempty" name:"ApiType"`

	// 集群名称，当ApiType取值为GET_SYSTEM_USERS，GET_PRIVILEGE_USERS，GET_CLUSTER_DATABASES，GET_CLUSTER_TABLES 时，此参数必填
	Cluster *string `json:"Cluster,omitnil,omitempty" name:"Cluster"`

	// 用户名称，api与user相关的必填
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账户的类型
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type DescribeCkSqlApisRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

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
	ApiType *string `json:"ApiType,omitnil,omitempty" name:"ApiType"`

	// 集群名称，当ApiType取值为GET_SYSTEM_USERS，GET_PRIVILEGE_USERS，GET_CLUSTER_DATABASES，GET_CLUSTER_TABLES 时，此参数必填
	Cluster *string `json:"Cluster,omitnil,omitempty" name:"Cluster"`

	// 用户名称，api与user相关的必填
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账户的类型
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
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
	ReturnData *string `json:"ReturnData,omitnil,omitempty" name:"ReturnData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	ClusterConfList []*ClusterConfigsInfoFromEMR `json:"ClusterConfList,omitnil,omitempty" name:"ClusterConfList"`

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
type DescribeInstanceClustersRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceClustersRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Clusters []*ClusterInfo `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 搜索的配置项名称
	SearchConfigName *string `json:"SearchConfigName,omitnil,omitempty" name:"SearchConfigName"`
}

type DescribeInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 搜索的配置项名称
	SearchConfigName *string `json:"SearchConfigName,omitnil,omitempty" name:"SearchConfigName"`
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
	ConfigItems []*InstanceConfigInfo `json:"ConfigItems,omitnil,omitempty" name:"ConfigItems"`

	// 未配置的参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnConfigItems []*InstanceConfigInfo `json:"UnConfigItems,omitnil,omitempty" name:"UnConfigItems"`

	// 配置的多层级参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapConfigItems []*MapConfigItem `json:"MapConfigItems,omitnil,omitempty" name:"MapConfigItems"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceNodesRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群角色类型，“DATA” 为数据节点、“COMMON” 为 ZooKeeper 节点，默认为 "DATA" 数据节点。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`

	// 当true的时候返回所有节点，即Limit无限大
	ForceAll *bool `json:"ForceAll,omitnil,omitempty" name:"ForceAll"`
}

type DescribeInstanceNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群角色类型，“DATA” 为数据节点、“COMMON” 为 ZooKeeper 节点，默认为 "DATA" 数据节点。
	NodeRole *string `json:"NodeRole,omitnil,omitempty" name:"NodeRole"`

	// 分页参数，第一页为0，第二页为10
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页步长，默认为10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 展现策略，All时显示所有
	DisplayPolicy *string `json:"DisplayPolicy,omitnil,omitempty" name:"DisplayPolicy"`

	// 当true的时候返回所有节点，即Limit无限大
	ForceAll *bool `json:"ForceAll,omitnil,omitempty" name:"ForceAll"`
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
	delete(f, "ForceAll")
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

	// 是否是open api查询
	IsOpenApi *bool `json:"IsOpenApi,omitnil,omitempty" name:"IsOpenApi"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否是open api查询
	IsOpenApi *bool `json:"IsOpenApi,omitnil,omitempty" name:"IsOpenApi"`
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
type DescribeInstanceShardsRequestParams struct {
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceShardsRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	InstanceShardsList *string `json:"InstanceShardsList,omitnil,omitempty" name:"InstanceShardsList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 集群操作进度
	FlowProgress *float64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 当前步骤的名称，例如：”购买资源中“
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// 当前步骤的名称，例如：”购买资源中“
	ProcessSubName *string `json:"ProcessSubName,omitnil,omitempty" name:"ProcessSubName"`

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
type DescribeInstancesNewRequestParams struct {
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

	// 信息详细与否
	IsSimple *bool `json:"IsSimple,omitnil,omitempty" name:"IsSimple"`

	// vip列表
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`
}

type DescribeInstancesNewRequest struct {
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

	// 信息详细与否
	IsSimple *bool `json:"IsSimple,omitnil,omitempty" name:"IsSimple"`

	// vip列表
	Vips []*string `json:"Vips,omitnil,omitempty" name:"Vips"`
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
	delete(f, "Vips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesNewResponseParams struct {
	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例数组
	InstancesList []*InstanceInfo `json:"InstancesList,omitnil,omitempty" name:"InstancesList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitnil,omitempty" name:"IsElastic"`

	// 是否是购买页面需要的spec
	CaseType *int64 `json:"CaseType,omitnil,omitempty" name:"CaseType"`
}

type DescribeSpecRequest struct {
	*tchttp.BaseRequest
	
	// 地域信息，例如"ap-guangzhou-1"
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 计费类型，PREPAID 包年包月，POSTPAID_BY_HOUR 按量计费
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitnil,omitempty" name:"IsElastic"`

	// 是否是购买页面需要的spec
	CaseType *int64 `json:"CaseType,omitnil,omitempty" name:"CaseType"`
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
	delete(f, "CaseType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecResponseParams struct {
	// zookeeper节点规格描述
	CommonSpec []*ResourceSpec `json:"CommonSpec,omitnil,omitempty" name:"CommonSpec"`

	// 数据节点规格描述
	DataSpec []*ResourceSpec `json:"DataSpec,omitnil,omitempty" name:"DataSpec"`

	// 云盘列表
	AttachCBSSpec []*DiskSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DestroyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
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
	// 作业id
	FlowID *string `json:"FlowID,omitnil,omitempty" name:"FlowID"`

	// 集群id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

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

type DiskSpec struct {
	// 磁盘类型，例如“CLOUD_SSD", "LOCAL_SSD"等
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 磁盘类型说明，例如"云SSD", "本地SSD"等
	DiskDesc *string `json:"DiskDesc,omitnil,omitempty" name:"DiskDesc"`

	// 磁盘最小规格大小，单位G
	MinDiskSize *int64 `json:"MinDiskSize,omitnil,omitempty" name:"MinDiskSize"`

	// 磁盘最大规格大小，单位G
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// 磁盘数目
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`
}

type GroupInfo struct {
	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 分片变量名称
	ShardName *string `json:"ShardName,omitnil,omitempty" name:"ShardName"`

	// 副本变量名称
	ReplicaName *string `json:"ReplicaName,omitnil,omitempty" name:"ReplicaName"`
}

type InstanceConfigInfo struct {
	// 配置项名称
	ConfKey *string `json:"ConfKey,omitnil,omitempty" name:"ConfKey"`

	// 配置项内容
	ConfValue *string `json:"ConfValue,omitnil,omitempty" name:"ConfValue"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 是否需要重启
	NeedRestart *bool `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 是否可编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 配置项解释
	ConfDesc *string `json:"ConfDesc,omitnil,omitempty" name:"ConfDesc"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 规则名称类型
	ModifyRuleType *string `json:"ModifyRuleType,omitnil,omitempty" name:"ModifyRuleType"`

	// 规则名称内容
	ModifyRuleValue *string `json:"ModifyRuleValue,omitnil,omitempty" name:"ModifyRuleValue"`

	// 修改人的uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type InstanceConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitnil,omitempty" name:"ConfKey"`

	// value
	ConfValue *string `json:"ConfValue,omitnil,omitempty" name:"ConfValue"`
}

type InstanceDetail struct {
	// 告警策略是否可用
	EnableAlarmStrategy *bool `json:"EnableAlarmStrategy,omitnil,omitempty" name:"EnableAlarmStrategy"`
}

type InstanceInfo struct {
	// 集群实例ID, "cdw-xxxx" 字符串类型
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 状态,
	// Init 创建中; Serving 运行中； 
	// Deleted已销毁；Deleting 销毁中；
	// Modify 集群变更中；
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 地域, ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区， ap-guangzhou-3
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 私有网络名称
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网名称
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 付费类型，"hour", "prepay"
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 数据节点描述信息
	MasterSummary *NodesSummary `json:"MasterSummary,omitnil,omitempty" name:"MasterSummary"`

	// zookeeper节点描述信息
	CommonSummary *NodesSummary `json:"CommonSummary,omitnil,omitempty" name:"CommonSummary"`

	// 高可用,"true" "false"
	HA *string `json:"HA,omitnil,omitempty" name:"HA"`

	// 访问地址，例如 "10.0.0.1:9000"
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`

	// 记录ID，数值型
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// regionId, 表示地域
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 可用区说明，例如 "广州二区"
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// 错误流程说明信息
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 状态描述，例如“运行中”等
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 自动续费标记
	RenewFlag *bool `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 监控信息
	Monitor *string `json:"Monitor,omitnil,omitempty" name:"Monitor"`

	// 是否开通日志
	HasClsTopic *bool `json:"HasClsTopic,omitnil,omitempty" name:"HasClsTopic"`

	// 日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 日志集ID
	ClsLogSetId *string `json:"ClsLogSetId,omitnil,omitempty" name:"ClsLogSetId"`

	// 是否支持xml配置管理
	EnableXMLConfig *int64 `json:"EnableXMLConfig,omitnil,omitempty" name:"EnableXMLConfig"`

	// 区域
	RegionDesc *string `json:"RegionDesc,omitnil,omitempty" name:"RegionDesc"`

	// 弹性网卡地址
	Eip *string `json:"Eip,omitnil,omitempty" name:"Eip"`

	// 冷热分层系数
	CosMoveFactor *int64 `json:"CosMoveFactor,omitnil,omitempty" name:"CosMoveFactor"`

	// external/local/yunti
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 是否弹性ck
	IsElastic *bool `json:"IsElastic,omitnil,omitempty" name:"IsElastic"`

	// 集群详细状态
	InstanceStateInfo *InstanceStateInfo `json:"InstanceStateInfo,omitnil,omitempty" name:"InstanceStateInfo"`

	// ZK高可用
	HAZk *bool `json:"HAZk,omitnil,omitempty" name:"HAZk"`

	// 挂载盘,默认0:没有类型；1:裸盘;2:lvm
	MountDiskType *int64 `json:"MountDiskType,omitnil,omitempty" name:"MountDiskType"`

	// chproxy连接ip
	CHProxyVip *string `json:"CHProxyVip,omitnil,omitempty" name:"CHProxyVip"`

	// cos buket的名字
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// 是否可以挂载云盘
	CanAttachCbs *bool `json:"CanAttachCbs,omitnil,omitempty" name:"CanAttachCbs"`

	// 是否可以挂载云盘阵列
	CanAttachCbsLvm *bool `json:"CanAttachCbsLvm,omitnil,omitempty" name:"CanAttachCbsLvm"`

	// 是否可以挂载cos
	CanAttachCos *bool `json:"CanAttachCos,omitnil,omitempty" name:"CanAttachCos"`

	// 服务信息
	Components []*ServiceInfo `json:"Components,omitnil,omitempty" name:"Components"`

	// 可升级的内核版本
	UpgradeVersions *string `json:"UpgradeVersions,omitnil,omitempty" name:"UpgradeVersions"`

	// ex-index
	EsIndexId *string `json:"EsIndexId,omitnil,omitempty" name:"EsIndexId"`

	// username
	EsIndexUsername *string `json:"EsIndexUsername,omitnil,omitempty" name:"EsIndexUsername"`

	// password
	EsIndexPassword *string `json:"EsIndexPassword,omitnil,omitempty" name:"EsIndexPassword"`

	// true
	HasEsIndex *bool `json:"HasEsIndex,omitnil,omitempty" name:"HasEsIndex"`

	// true
	IsSecondaryZone *bool `json:"IsSecondaryZone,omitnil,omitempty" name:"IsSecondaryZone"`

	// desc
	SecondaryZoneInfo *string `json:"SecondaryZoneInfo,omitnil,omitempty" name:"SecondaryZoneInfo"`

	// 是否clickhouse-keeper
	ClickHouseKeeper *bool `json:"ClickHouseKeeper,omitnil,omitempty" name:"ClickHouseKeeper"`

	// 实例扩展信息
	Details *InstanceDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 安全组白名单
	IsWhiteSGs *bool `json:"IsWhiteSGs,omitnil,omitempty" name:"IsWhiteSGs"`

	// 绑定的安全组
	BindSGs []*string `json:"BindSGs,omitnil,omitempty" name:"BindSGs"`

	// 是否开启公网clb
	HasPublicCloudClb *bool `json:"HasPublicCloudClb,omitnil,omitempty" name:"HasPublicCloudClb"`

	// 可升级的zk版本
	UpgradeZkVersions *string `json:"UpgradeZkVersions,omitnil,omitempty" name:"UpgradeZkVersions"`
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
	Cluster *string `json:"Cluster,omitnil,omitempty" name:"Cluster"`

	// 节点所属的分组信息
	NodeGroups []*GroupInfo `json:"NodeGroups,omitnil,omitempty" name:"NodeGroups"`

	// VPC IP
	Rip *string `json:"Rip,omitnil,omitempty" name:"Rip"`

	// ture的时候表示该节点上部署了chproxy进程
	IsCHProxy *bool `json:"IsCHProxy,omitnil,omitempty" name:"IsCHProxy"`

	// 节点状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点uuid
	UUID *string `json:"UUID,omitnil,omitempty" name:"UUID"`

	// 区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 区描述
	ZoneDesc *string `json:"ZoneDesc,omitnil,omitempty" name:"ZoneDesc"`

	// 真实资源id
	RealResourceId *string `json:"RealResourceId,omitnil,omitempty" name:"RealResourceId"`
}

type InstanceStateInfo struct {
	// 集群状态，例如：Serving
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`

	// 集群操作创建时间
	FlowCreateTime *string `json:"FlowCreateTime,omitnil,omitempty" name:"FlowCreateTime"`

	// 集群操作名称
	FlowName *string `json:"FlowName,omitnil,omitempty" name:"FlowName"`

	// 集群操作进度
	FlowProgress *int64 `json:"FlowProgress,omitnil,omitempty" name:"FlowProgress"`

	// 集群状态描述，例如：运行中
	InstanceStateDesc *string `json:"InstanceStateDesc,omitnil,omitempty" name:"InstanceStateDesc"`

	// 集群流程错误信息，例如：“创建失败，资源不足”
	FlowMsg *string `json:"FlowMsg,omitnil,omitempty" name:"FlowMsg"`

	// 当前步骤的名称，例如：”购买资源中“
	ProcessName *string `json:"ProcessName,omitnil,omitempty" name:"ProcessName"`

	// 请求id
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 流程的二级名称
	ProcessSubName *string `json:"ProcessSubName,omitnil,omitempty" name:"ProcessSubName"`
}

type MapConfigItem struct {
	// key
	ConfKey *string `json:"ConfKey,omitnil,omitempty" name:"ConfKey"`

	// 列表
	Items []*InstanceConfigInfo `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ModifyClusterConfigsRequestParams struct {
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitnil,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyClusterConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，例如cdwch-xxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件修改信息
	ModifyConfContext []*ConfigSubmitContext `json:"ModifyConfContext,omitnil,omitempty" name:"ModifyConfContext"`

	// 修改原因
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil,omitempty" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil,omitempty" name:"UpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil,omitempty" name:"DeleteItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil,omitempty" name:"DelItems"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyInstanceKeyValConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 新增配置列表
	AddItems []*InstanceConfigItem `json:"AddItems,omitnil,omitempty" name:"AddItems"`

	// 更新配置列表
	UpdateItems []*InstanceConfigItem `json:"UpdateItems,omitnil,omitempty" name:"UpdateItems"`

	// 删除配置列表
	DeleteItems *InstanceConfigItem `json:"DeleteItems,omitnil,omitempty" name:"DeleteItems"`

	// 删除配置列表
	DelItems []*InstanceConfigItem `json:"DelItems,omitnil,omitempty" name:"DelItems"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// ID
	FlowId *int64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// cluster名称
	Cluster *string `json:"Cluster,omitnil,omitempty" name:"Cluster"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 是否所有数据库表
	AllDatabase *bool `json:"AllDatabase,omitnil,omitempty" name:"AllDatabase"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库表权限
	DatabasePrivilegeList []*DatabasePrivilegeInfo `json:"DatabasePrivilegeList,omitnil,omitempty" name:"DatabasePrivilegeList"`
}

type ModifyUserNewPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// cluster名称
	Cluster *string `json:"Cluster,omitnil,omitempty" name:"Cluster"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 是否所有数据库表
	AllDatabase *bool `json:"AllDatabase,omitnil,omitempty" name:"AllDatabase"`

	// 全局权限
	GlobalPrivileges []*string `json:"GlobalPrivileges,omitnil,omitempty" name:"GlobalPrivileges"`

	// 数据库表权限
	DatabasePrivilegeList []*DatabasePrivilegeInfo `json:"DatabasePrivilegeList,omitnil,omitempty" name:"DatabasePrivilegeList"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 云盘大小
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
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
	AttachCBSSpec *AttachCBSSpec `json:"AttachCBSSpec,omitnil,omitempty" name:"AttachCBSSpec"`

	// 子产品类型
	SubProductType *string `json:"SubProductType,omitnil,omitempty" name:"SubProductType"`

	// 规格对应的核数
	SpecCore *int64 `json:"SpecCore,omitnil,omitempty" name:"SpecCore"`

	// 规格对应的内存大小
	SpecMemory *int64 `json:"SpecMemory,omitnil,omitempty" name:"SpecMemory"`

	// 磁盘的数量
	DiskCount *int64 `json:"DiskCount,omitnil,omitempty" name:"DiskCount"`

	// 磁盘的最大大小
	MaxDiskSize *int64 `json:"MaxDiskSize,omitnil,omitempty" name:"MaxDiskSize"`

	// 是否为加密云盘
	Encrypt *int64 `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`
}

// Predefined struct for user
type OpenBackUpRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`
}

type OpenBackUpRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// OPEN 或者CLOSE
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 桶名字
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
}

type RecoverBackUpJobRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务id
	BackUpJobId *int64 `json:"BackUpJobId,omitnil,omitempty" name:"BackUpJobId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 磁盘扩容后容量，不能小于原有用量。clickhouse最小200，且为100的整数倍。 zk最小100，且为10的整数倍；
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type ResizeDiskRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 磁盘扩容后容量，不能小于原有用量。clickhouse最小200，且为100的整数倍。 zk最小100，且为10的整数倍；
	DiskSize *int64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
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
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

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

type ResourceSpec struct {
	// 规格名称，例如“SCH1"
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// cpu核数
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存大小，单位G
	Mem *int64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 分类标记，STANDARD/BIGDATA/HIGHIO分别表示标准型/大数据型/高IO
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 系统盘描述信息
	SystemDisk *DiskSpec `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 数据盘描述信息
	DataDisk *DiskSpec `json:"DataDisk,omitnil,omitempty" name:"DataDisk"`

	// 最大节点数目限制
	MaxNodeSize *int64 `json:"MaxNodeSize,omitnil,omitempty" name:"MaxNodeSize"`

	// 是否可用，false代表售罄
	Available *bool `json:"Available,omitnil,omitempty" name:"Available"`

	// 规格描述信息
	ComputeSpecDesc *string `json:"ComputeSpecDesc,omitnil,omitempty" name:"ComputeSpecDesc"`

	// 规格名
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 库存数
	InstanceQuota *int64 `json:"InstanceQuota,omitnil,omitempty" name:"InstanceQuota"`
}

// Predefined struct for user
type ScaleCNOutUpInstanceRequestParams struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// warehouse名称
	VirtualCluster *string `json:"VirtualCluster,omitnil,omitempty" name:"VirtualCluster"`

	// 子网id
	UserSubnetID *string `json:"UserSubnetID,omitnil,omitempty" name:"UserSubnetID"`

	// 新的warehouse的个数
	NewCount *int64 `json:"NewCount,omitnil,omitempty" name:"NewCount"`

	// 集群的规格2X-Small、X-Small、Small
	NewSpecName *string `json:"NewSpecName,omitnil,omitempty" name:"NewSpecName"`
}

type ScaleCNOutUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// warehouse名称
	VirtualCluster *string `json:"VirtualCluster,omitnil,omitempty" name:"VirtualCluster"`

	// 子网id
	UserSubnetID *string `json:"UserSubnetID,omitnil,omitempty" name:"UserSubnetID"`

	// 新的warehouse的个数
	NewCount *int64 `json:"NewCount,omitnil,omitempty" name:"NewCount"`

	// 集群的规格2X-Small、X-Small、Small
	NewSpecName *string `json:"NewSpecName,omitnil,omitempty" name:"NewSpecName"`
}

func (r *ScaleCNOutUpInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleCNOutUpInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualCluster")
	delete(f, "UserSubnetID")
	delete(f, "NewCount")
	delete(f, "NewSpecName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleCNOutUpInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleCNOutUpInstanceResponseParams struct {
	// 流程ID
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleCNOutUpInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ScaleCNOutUpInstanceResponseParams `json:"Response"`
}

func (r *ScaleCNOutUpInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleCNOutUpInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstanceRequestParams struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 调整clickhouse节点数量
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// v_cluster分组，	
	// 新增扩容节点将加入到已选择的v_cluster分组中，提交同步VIP生效.
	ScaleOutCluster *string `json:"ScaleOutCluster,omitnil,omitempty" name:"ScaleOutCluster"`

	// 子网剩余ip数量，用于判断当前实例子网剩余ip数是否能扩容。需要根据实际填写
	UserSubnetIPNum *int64 `json:"UserSubnetIPNum,omitnil,omitempty" name:"UserSubnetIPNum"`

	// 同步元数据节点IP （uip），扩容的时候必填
	ScaleOutNodeIp *string `json:"ScaleOutNodeIp,omitnil,omitempty" name:"ScaleOutNodeIp"`

	// 缩容节点shard的节点IP （uip），其中ha集群需要主副节点ip都传入以逗号分隔，缩容的时候必填
	ReduceShardInfo []*string `json:"ReduceShardInfo,omitnil,omitempty" name:"ReduceShardInfo"`
}

type ScaleOutInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 调整clickhouse节点数量
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// v_cluster分组，	
	// 新增扩容节点将加入到已选择的v_cluster分组中，提交同步VIP生效.
	ScaleOutCluster *string `json:"ScaleOutCluster,omitnil,omitempty" name:"ScaleOutCluster"`

	// 子网剩余ip数量，用于判断当前实例子网剩余ip数是否能扩容。需要根据实际填写
	UserSubnetIPNum *int64 `json:"UserSubnetIPNum,omitnil,omitempty" name:"UserSubnetIPNum"`

	// 同步元数据节点IP （uip），扩容的时候必填
	ScaleOutNodeIp *string `json:"ScaleOutNodeIp,omitnil,omitempty" name:"ScaleOutNodeIp"`

	// 缩容节点shard的节点IP （uip），其中ha集群需要主副节点ip都传入以逗号分隔，缩容的时候必填
	ReduceShardInfo []*string `json:"ReduceShardInfo,omitnil,omitempty" name:"ReduceShardInfo"`
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
	FlowId *string `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 实例ID
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
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// clickhouse节点规格。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 是否滚动重启，false为不滚动重启，true为滚动重启
	ScaleUpEnableRolling *bool `json:"ScaleUpEnableRolling,omitnil,omitempty" name:"ScaleUpEnableRolling"`
}

type ScaleUpInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 节点类型，DATA：clickhouse节点，COMMON：为zookeeper节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// clickhouse节点规格。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 是否滚动重启，false为不滚动重启，true为滚动重启
	ScaleUpEnableRolling *bool `json:"ScaleUpEnableRolling,omitnil,omitempty" name:"ScaleUpEnableRolling"`
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

type ScheduleStrategy struct {
	// 备份桶名称
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`

	// 备份保留天数
	RetainDays *int64 `json:"RetainDays,omitnil,omitempty" name:"RetainDays"`

	// 备份的天
	WeekDays *string `json:"WeekDays,omitnil,omitempty" name:"WeekDays"`

	// 备份小时
	ExecuteHour *int64 `json:"ExecuteHour,omitnil,omitempty" name:"ExecuteHour"`

	// 策略id
	ScheduleId *int64 `json:"ScheduleId,omitnil,omitempty" name:"ScheduleId"`

	// 下次备份时间
	NextBackupTime *string `json:"NextBackupTime,omitnil,omitempty" name:"NextBackupTime"`
}

type SearchTags struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 1表示只输入标签的键，没有输入值；0表示输入键时且输入值
	AllValue *int64 `json:"AllValue,omitnil,omitempty" name:"AllValue"`
}

type SecondaryZoneInfo struct {
	// 副可用区
	SecondaryZone *string `json:"SecondaryZone,omitnil,omitempty" name:"SecondaryZone"`

	// 可用区可用的子网id
	SecondarySubnet *string `json:"SecondarySubnet,omitnil,omitempty" name:"SecondarySubnet"`

	// 可用区可用的子网可用ip的数量
	UserIpNum *string `json:"UserIpNum,omitnil,omitempty" name:"UserIpNum"`

	// 可用区可用的子网可用ip的数量
	SecondaryUserSubnetIPNum *int64 `json:"SecondaryUserSubnetIPNum,omitnil,omitempty" name:"SecondaryUserSubnetIPNum"`
}

type ServiceInfo struct {
	// 服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 服务的版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type TablePrivilegeInfo struct {
	// 表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表权限列表 SELECT、INSERT_ALL、ALTER、TRUNCATE、DROP_TABLE 查询、插入、设置、清空表、删除表
	TablePrivileges []*string `json:"TablePrivileges,omitnil,omitempty" name:"TablePrivileges"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}