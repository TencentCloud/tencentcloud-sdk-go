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

package v20180328

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccountCreateInfo struct {
	// 实例用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 实例密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// DB权限列表
	DBPrivileges []*DBPrivilege `json:"DBPrivileges,omitempty" name:"DBPrivileges"`

	// 账号备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否为管理员账户，当值为true 等价于基础版AccountType=L0，高可用AccountType=L1，当值为false，等价于AccountType=L3
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`

	// win-windows鉴权,sql-sqlserver鉴权，不填默认值为sql-sqlserver鉴权
	Authentication *string `json:"Authentication,omitempty" name:"Authentication"`

	// 账号类型，IsAdmin的扩展字段。 L0-超级权限(基础版独有),L1-高级权限,L2-特殊权限,L3-普通权限，默认L3
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`
}

type AccountDetail struct {
	// 账户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 账户备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 账户创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 账户状态，1-创建中，2-正常，3-修改中，4-密码重置中，-1-删除中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 账户更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 密码更新时间
	PassTime *string `json:"PassTime,omitempty" name:"PassTime"`

	// 账户内部状态，正常为enable
	InternalStatus *string `json:"InternalStatus,omitempty" name:"InternalStatus"`

	// 该账户对相关db的读写权限信息
	Dbs []*DBPrivilege `json:"Dbs,omitempty" name:"Dbs"`

	// 是否为管理员账户
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`

	// win-windows鉴权,sql-sqlserver鉴权
	Authentication *string `json:"Authentication,omitempty" name:"Authentication"`

	// win-windows鉴权账户需要host
	Host *string `json:"Host,omitempty" name:"Host"`

	// 账号类型。L0-超级权限(基础版独有),L1-高级权限,L2-特殊权限,L3-普通权限
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`
}

type AccountPassword struct {
	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type AccountPrivilege struct {
	// 数据库用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 数据库权限。ReadWrite表示可读写，ReadOnly表示只读,Delete表示删除DB对该账户的权限，DBOwner所有者
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`

	// 账户名称，L0-超级权限(基础版独有),L1-高级权限,L2-特殊权限,L3-普通权限
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`
}

type AccountPrivilegeModifyInfo struct {
	// 数据库用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 账号权限变更信息
	DBPrivileges []*DBPrivilegeModifyInfo `json:"DBPrivileges,omitempty" name:"DBPrivileges"`

	// 是否为管理员账户,当值为true 等价于基础版AccountType=L0，高可用AccountType=L1，当值为false时，表示删除管理员权限，默认false
	IsAdmin *bool `json:"IsAdmin,omitempty" name:"IsAdmin"`

	// 账号类型，IsAdmin字段的扩展字段。 L0-超级权限(基础版独有),L1-高级权限,L2-特殊权限,L3-普通权限，默认L3
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`
}

type AccountRemark struct {
	// 账户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 对应账户新的备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

// Predefined struct for user
type AssociateSecurityGroupsRequestParams struct {
	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例ID 列表，一个或者多个实例ID组成的数组。多个实例必须是同一个地域，同一个可用区，同一个项目下的。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type AssociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例ID 列表，一个或者多个实例ID组成的数组。多个实例必须是同一个地域，同一个可用区，同一个项目下的。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *AssociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backup struct {
	// 文件名，对于单库备份文件不返回此值；单库备份文件通过DescribeBackupFiles接口获取文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小，单位 KB，对于单库备份文件不返回此值；单库备份文件通过DescribeBackupFiles接口获取文件大小
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 内网下载地址，对于单库备份文件不返回此值；单库备份文件通过DescribeBackupFiles接口获取下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址，对于单库备份文件不返回此值；单库备份文件通过DescribeBackupFiles接口获取下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// 备份文件唯一标识，RestoreInstance接口会用到该字段，对于单库备份文件不返回此值；单库备份文件通过DescribeBackupFiles接口获取可回档的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 备份文件状态（0-创建中；1-成功；2-失败）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 多库备份时的DB列表
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// 备份策略（0-实例备份；1-多库备份）
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 备份方式，0-定时备份；1-手动临时备份；2-定期备份
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`

	// 备份任务名称，可自定义
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 聚合Id，对于打包备份文件不返回此值。通过此值调用DescribeBackupFiles接口，获取单库备份文件的详细信息
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 备份文件形式（pkg-打包备份文件，single-单库备份文件）
	BackupFormat *string `json:"BackupFormat,omitempty" name:"BackupFormat"`

	// 实例当前地域Code
	Region *string `json:"Region,omitempty" name:"Region"`

	// 跨地域备份的目的地域下载链接
	CrossBackupAddr []*CrossBackupAddr `json:"CrossBackupAddr,omitempty" name:"CrossBackupAddr"`

	// 跨地域备份的目标地域和备份状态
	CrossBackupStatus []*CrossRegionStatus `json:"CrossBackupStatus,omitempty" name:"CrossBackupStatus"`
}

type BackupFile struct {
	// 备份文件唯一标识
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 备份文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件大小(K)
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 备份文件的库的名称
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// 下载地址
	DownloadLink *string `json:"DownloadLink,omitempty" name:"DownloadLink"`

	// 当前实例地域码
	Region *string `json:"Region,omitempty" name:"Region"`

	// 备份的跨地域region和所对应的下载地址
	CrossBackupAddr []*CrossBackupAddr `json:"CrossBackupAddr,omitempty" name:"CrossBackupAddr"`
}

type BusinessIntelligenceFile struct {
	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件的COS_URL
	FileURL *string `json:"FileURL,omitempty" name:"FileURL"`

	// 文件在服务器上的路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 文件大小，单位时Byte
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件md5值
	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`

	// 部署文件状态 1-初始化待部署 2-部署中 3-部署成功 4-部署失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 文件创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 文件部署开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 文件部署结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 报错信息返回
	Message *string `json:"Message,omitempty" name:"Message"`

	// 商业智能实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 动作相关信息
	Action *FileAction `json:"Action,omitempty" name:"Action"`
}

// Predefined struct for user
type CloneDBRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照ReNameRestoreDatabase中的库进行克隆，并重命名，新库名称必须指定
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

type CloneDBRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照ReNameRestoreDatabase中的库进行克隆，并重命名，新库名称必须指定
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *CloneDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneDBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneDBResponseParams struct {
	// 异步流程任务ID，使用FlowId调用DescribeFlowStatus接口获取任务执行状态
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloneDBResponse struct {
	*tchttp.BaseResponse
	Response *CloneDBResponseParams `json:"Response"`
}

func (r *CloneDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneDBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseInterCommunicationRequestParams struct {
	// 关闭互通的实例ID集合
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type CloseInterCommunicationRequest struct {
	*tchttp.BaseRequest
	
	// 关闭互通的实例ID集合
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *CloseInterCommunicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseInterCommunicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseInterCommunicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseInterCommunicationResponseParams struct {
	// 实例和异步流程ID
	InterInstanceFlowSet []*InterInstanceFlow `json:"InterInstanceFlowSet,omitempty" name:"InterInstanceFlowSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CloseInterCommunicationResponse struct {
	*tchttp.BaseResponse
	Response *CloseInterCommunicationResponseParams `json:"Response"`
}

func (r *CloseInterCommunicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseInterCommunicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteExpansionRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CompleteExpansionRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CompleteExpansionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteExpansionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteExpansionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteExpansionResponseParams struct {
	// 流程ID，可通过接口DescribeFlowStatus查询立即切换升级任务的状态。
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompleteExpansionResponse struct {
	*tchttp.BaseResponse
	Response *CompleteExpansionResponseParams `json:"Response"`
}

func (r *CompleteExpansionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteExpansionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrationRequestParams struct {
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type CompleteMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *CompleteMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrationResponseParams struct {
	// 完成迁移流程发起后，返回的流程id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompleteMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CompleteMigrationResponseParams `json:"Response"`
}

func (r *CompleteMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosUploadBackupFile struct {
	// 备份名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 备份大小
	Size *int64 `json:"Size,omitempty" name:"Size"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库实例账户信息
	Accounts []*AccountCreateInfo `json:"Accounts,omitempty" name:"Accounts"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库实例账户信息
	Accounts []*AccountCreateInfo `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移任务恢复类型，FULL-全量备份恢复，FULL_LOG-全量备份+事务日志恢复，FULL_DIFF-全量备份+差异备份恢复
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// 备份上传类型，COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，需要用户上传。
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 任务名称
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// UploadType是COS_URL时这里填URL，COS_UPLOAD这里填备份文件的名字。只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

type CreateBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移任务恢复类型，FULL-全量备份恢复，FULL_LOG-全量备份+事务日志恢复，FULL_DIFF-全量备份+差异备份恢复
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// 备份上传类型，COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，需要用户上传。
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 任务名称
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// UploadType是COS_URL时这里填URL，COS_UPLOAD这里填备份文件的名字。只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

func (r *CreateBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RecoveryType")
	delete(f, "UploadType")
	delete(f, "MigrationName")
	delete(f, "BackupFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupMigrationResponseParams struct {
	// 备份导入任务ID
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupMigrationResponseParams `json:"Response"`
}

func (r *CreateBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// 备份策略(0-实例备份 1-多库备份)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 需要备份库名的列表(多库备份才填写)
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 实例ID，形如mssql-i1z41iwd
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份名称，若不填则自动生成“实例ID_备份开始时间戳”
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// 备份策略(0-实例备份 1-多库备份)
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 需要备份库名的列表(多库备份才填写)
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 实例ID，形如mssql-i1z41iwd
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份名称，若不填则自动生成“实例ID_备份开始时间戳”
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`
}

func (r *CreateBackupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Strategy")
	delete(f, "DBNames")
	delete(f, "InstanceId")
	delete(f, "BackupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// 异步任务ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse
	Response *CreateBackupResponseParams `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBasicDBInstancesRequestParams struct {
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例的CPU核心数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例内存大小，单位GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// VPC子网ID，形如subnet-bdoe83fa
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买实例的宿主机类型, CLOUD_PREMIUM-虚拟机高性能云盘，CLOUD_SSD-虚拟机SSD云盘
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 付费模式，取值支持 PREPAID（预付费），POSTPAID（后付费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 本次购买几个实例，默认值为1。取值不超过10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// sqlserver版本，目前只支持：2008R2（SQL Server 2008 Enterprise），2012SP3（SQL Server 2012 Enterprise），2016SP1（SQL Server 2016 Enterprise），201602（SQL Server 2016 Standard），2017（SQL Server 2017 Enterprise），201202（SQL Server 2012 Standard），201402（SQL Server 2014 Standard），2014SP2（SQL Server 2014 Enterprise），201702（SQL Server 2017 Standard）版本。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息。不填，默认为版本2008R2。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 自动续费标志：0-正常续费  1-自动续费，默认为1自动续费。只在购买预付费实例时有效。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 可维护时间窗配置，以周为单位，表示周几允许维护，1-7分别代表周一到周末
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 可维护时间窗配置，每天可维护的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可维护时间窗配置，持续时间，单位：小时
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

type CreateBasicDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例的CPU核心数
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例内存大小，单位GB
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`

	// VPC子网ID，形如subnet-bdoe83fa
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买实例的宿主机类型, CLOUD_PREMIUM-虚拟机高性能云盘，CLOUD_SSD-虚拟机SSD云盘
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 付费模式，取值支持 PREPAID（预付费），POSTPAID（后付费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 本次购买几个实例，默认值为1。取值不超过10
	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// sqlserver版本，目前只支持：2008R2（SQL Server 2008 Enterprise），2012SP3（SQL Server 2012 Enterprise），2016SP1（SQL Server 2016 Enterprise），201602（SQL Server 2016 Standard），2017（SQL Server 2017 Enterprise），201202（SQL Server 2012 Standard），201402（SQL Server 2014 Standard），2014SP2（SQL Server 2014 Enterprise），201702（SQL Server 2017 Standard）版本。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息。不填，默认为版本2008R2。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 自动续费标志：0-正常续费  1-自动续费，默认为1自动续费。只在购买预付费实例时有效。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 可维护时间窗配置，以周为单位，表示周几允许维护，1-7分别代表周一到周末
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 可维护时间窗配置，每天可维护的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可维护时间窗配置，持续时间，单位：小时
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

func (r *CreateBasicDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Cpu")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "MachineType")
	delete(f, "InstanceChargeType")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "DBVersion")
	delete(f, "Period")
	delete(f, "SecurityGroupList")
	delete(f, "AutoRenewFlag")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBasicDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBasicDBInstancesResponseParams struct {
	// 订单名称
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBasicDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateBasicDBInstancesResponseParams `json:"Response"`
}

func (r *CreateBasicDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBasicDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessDBInstancesRequestParams struct {
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 预购买实例的CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 购买实例的宿主机类型，CLOUD_PREMIUM-虚拟机高性能云盘，CLOUD_SSD-虚拟机SSD云盘
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 本次购买几个实例，默认值为1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 商业智能服务器版本，目前只支持：201603（SQL Server 2016 Integration Services），201703（SQL Server 2017 Integration Services），201903（SQL Server 2019 Integration Services）版本。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息。不填，默认为版本201903。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 可维护时间窗配置，以周为单位，表示周几允许维护，1-7分别代表周一到周末
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 可维护时间窗配置，每天可维护的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可维护时间窗配置，持续时间，单位：小时
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type CreateBusinessDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 预购买实例的CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 购买实例的宿主机类型，CLOUD_PREMIUM-虚拟机高性能云盘，CLOUD_SSD-虚拟机SSD云盘
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 本次购买几个实例，默认值为1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 商业智能服务器版本，目前只支持：201603（SQL Server 2016 Integration Services），201703（SQL Server 2017 Integration Services），201903（SQL Server 2019 Integration Services）版本。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息。不填，默认为版本201903。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 可维护时间窗配置，以周为单位，表示周几允许维护，1-7分别代表周一到周末
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 可维护时间窗配置，每天可维护的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可维护时间窗配置，持续时间，单位：小时
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

func (r *CreateBusinessDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Cpu")
	delete(f, "MachineType")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "DBVersion")
	delete(f, "SecurityGroupList")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBusinessDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessDBInstancesResponseParams struct {
	// 订单名称
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBusinessDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateBusinessDBInstancesResponseParams `json:"Response"`
}

func (r *CreateBusinessDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessIntelligenceFileRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// COS_URL
	FileURL *string `json:"FileURL,omitempty" name:"FileURL"`

	// 文件类型 FLAT-作为数据源的平面文件， SSIS-ssis项目包
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateBusinessIntelligenceFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// COS_URL
	FileURL *string `json:"FileURL,omitempty" name:"FileURL"`

	// 文件类型 FLAT-作为数据源的平面文件， SSIS-ssis项目包
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateBusinessIntelligenceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessIntelligenceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileURL")
	delete(f, "FileType")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBusinessIntelligenceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBusinessIntelligenceFileResponseParams struct {
	// 文件名称
	FileTaskId *string `json:"FileTaskId,omitempty" name:"FileTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateBusinessIntelligenceFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateBusinessIntelligenceFileResponseParams `json:"Response"`
}

func (r *CreateBusinessIntelligenceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBusinessIntelligenceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesRequestParams struct {
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 付费模式，取值支持 PREPAID（预付费），POSTPAID（后付费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 本次购买几个实例，默认值为1。取值不超过10
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// sqlserver版本，目前所有支持的版本有：2008R2 (SQL Server 2008 R2 Enterprise)，2012SP3 (SQL Server 2012 Enterprise)，201202 (SQL Server 2012 Standard)，2014SP2 (SQL Server 2014 Enterprise)，201402 (SQL Server 2014 Standard)，2016SP1 (SQL Server 2016 Enterprise)，201602 (SQL Server 2016 Standard)，2017 (SQL Server 2017 Enterprise)，201702 (SQL Server 2017 Standard)，2019 (SQL Server 2019 Enterprise)，201902 (SQL Server 2019 Standard)。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息。不填，默认为版本2008R2。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 自动续费标志：0-正常续费  1-自动续费，默认为1自动续费。只在购买预付费实例时有效。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 可维护时间窗配置，以周为单位，表示周几允许维护，1-7分别代表周一到周末
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 可维护时间窗配置，每天可维护的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可维护时间窗配置，持续时间，单位：小时
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// 购买高可用实例的类型：DUAL-双机高可用  CLUSTER-集群，默认值为DUAL
	HAType *string `json:"HAType,omitempty" name:"HAType"`

	// 是否跨可用区部署，默认值为false
	MultiZones *bool `json:"MultiZones,omitempty" name:"MultiZones"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

type CreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 付费模式，取值支持 PREPAID（预付费），POSTPAID（后付费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 本次购买几个实例，默认值为1。取值不超过10
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// sqlserver版本，目前所有支持的版本有：2008R2 (SQL Server 2008 R2 Enterprise)，2012SP3 (SQL Server 2012 Enterprise)，201202 (SQL Server 2012 Standard)，2014SP2 (SQL Server 2014 Enterprise)，201402 (SQL Server 2014 Standard)，2016SP1 (SQL Server 2016 Enterprise)，201602 (SQL Server 2016 Standard)，2017 (SQL Server 2017 Enterprise)，201702 (SQL Server 2017 Standard)，2019 (SQL Server 2019 Enterprise)，201902 (SQL Server 2019 Standard)。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息。不填，默认为版本2008R2。
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 自动续费标志：0-正常续费  1-自动续费，默认为1自动续费。只在购买预付费实例时有效。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 可维护时间窗配置，以周为单位，表示周几允许维护，1-7分别代表周一到周末
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 可维护时间窗配置，每天可维护的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可维护时间窗配置，持续时间，单位：小时
	Span *int64 `json:"Span,omitempty" name:"Span"`

	// 购买高可用实例的类型：DUAL-双机高可用  CLUSTER-集群，默认值为DUAL
	HAType *string `json:"HAType,omitempty" name:"HAType"`

	// 是否跨可用区部署，默认值为false
	MultiZones *bool `json:"MultiZones,omitempty" name:"MultiZones"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

func (r *CreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "InstanceChargeType")
	delete(f, "ProjectId")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "DBVersion")
	delete(f, "AutoRenewFlag")
	delete(f, "SecurityGroupList")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	delete(f, "HAType")
	delete(f, "MultiZones")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBInstancesResponseParams struct {
	// 订单名称
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 订单名称数组
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBInstancesResponseParams `json:"Response"`
}

func (r *CreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库创建信息
	DBs []*DBCreateInfo `json:"DBs,omitempty" name:"DBs"`
}

type CreateDBRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库创建信息
	DBs []*DBCreateInfo `json:"DBs,omitempty" name:"DBs"`
}

func (r *CreateDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDBResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBResponseParams `json:"Response"`
}

func (r *CreateDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIncrementalMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量备份文件，全量备份任务UploadType是COS_URL时这里填URL，是COS_UPLOAD这里填备份文件的名字；只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// 是否需要恢复，NO-不需要，YES-需要，默认不需要
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`
}

type CreateIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量备份文件，全量备份任务UploadType是COS_URL时这里填URL，是COS_UPLOAD这里填备份文件的名字；只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// 是否需要恢复，NO-不需要，YES-需要，默认不需要
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`
}

func (r *CreateIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "BackupFiles")
	delete(f, "IsRecovery")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIncrementalMigrationResponseParams struct {
	// 增量备份导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateIncrementalMigrationResponseParams `json:"Response"`
}

func (r *CreateIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationRequestParams struct {
	// 迁移任务的名称
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 迁移类型（1:结构迁移 2:数据迁移 3:增量同步）
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// 迁移源的类型 1:TencentDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式）
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 迁移源
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// 迁移目标
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// 迁移DB对象 ，离线迁移不使用（SourceType=4或SourceType=5）。
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`

	// 按照ReNameRestoreDatabase中的库进行恢复，并重命名，不填则按照默认方式命名恢复的库，且恢复所有的库。SourceType=5的情况下有效。
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

type CreateMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务的名称
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 迁移类型（1:结构迁移 2:数据迁移 3:增量同步）
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// 迁移源的类型 1:TencentDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式）
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 迁移源
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// 迁移目标
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// 迁移DB对象 ，离线迁移不使用（SourceType=4或SourceType=5）。
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`

	// 按照ReNameRestoreDatabase中的库进行恢复，并重命名，不填则按照默认方式命名恢复的库，且恢复所有的库。SourceType=5的情况下有效。
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *CreateMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateName")
	delete(f, "MigrateType")
	delete(f, "SourceType")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "MigrateDBSet")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationResponseParams struct {
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateMigrationResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrationResponseParams `json:"Response"`
}

func (r *CreateMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePublishSubscribeRequestParams struct {
	// 发布实例ID，形如mssql-j8kv137v
	PublishInstanceId *string `json:"PublishInstanceId,omitempty" name:"PublishInstanceId"`

	// 订阅实例ID，形如mssql-j8kv137v
	SubscribeInstanceId *string `json:"SubscribeInstanceId,omitempty" name:"SubscribeInstanceId"`

	// 数据库的订阅发布关系集合
	DatabaseTupleSet []*DatabaseTuple `json:"DatabaseTupleSet,omitempty" name:"DatabaseTupleSet"`

	// 发布订阅的名称，默认值为：default_name
	PublishSubscribeName *string `json:"PublishSubscribeName,omitempty" name:"PublishSubscribeName"`
}

type CreatePublishSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 发布实例ID，形如mssql-j8kv137v
	PublishInstanceId *string `json:"PublishInstanceId,omitempty" name:"PublishInstanceId"`

	// 订阅实例ID，形如mssql-j8kv137v
	SubscribeInstanceId *string `json:"SubscribeInstanceId,omitempty" name:"SubscribeInstanceId"`

	// 数据库的订阅发布关系集合
	DatabaseTupleSet []*DatabaseTuple `json:"DatabaseTupleSet,omitempty" name:"DatabaseTupleSet"`

	// 发布订阅的名称，默认值为：default_name
	PublishSubscribeName *string `json:"PublishSubscribeName,omitempty" name:"PublishSubscribeName"`
}

func (r *CreatePublishSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePublishSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PublishInstanceId")
	delete(f, "SubscribeInstanceId")
	delete(f, "DatabaseTupleSet")
	delete(f, "PublishSubscribeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePublishSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePublishSubscribeResponseParams struct {
	// 流程ID，可通过接口DescribeFlowStatus查询立即切换升级任务的状态。
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePublishSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreatePublishSubscribeResponseParams `json:"Response"`
}

func (r *CreatePublishSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePublishSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstancesRequestParams struct {
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 只读组类型选项，1-按照一个实例一个只读组的方式发货，2-新建只读组后发货，所有实例都在这个只读组下面， 3-发货的所有实例都在已有的只读组下面
	ReadOnlyGroupType *int64 `json:"ReadOnlyGroupType,omitempty" name:"ReadOnlyGroupType"`

	// 实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 0-默认不升级主实例，1-强制升级主实例完成ro部署；主实例为非集群版时需要填1，强制升级为集群版。填1 说明您已同意将主实例升级到集群版实例。
	ReadOnlyGroupForcedUpgrade *int64 `json:"ReadOnlyGroupForcedUpgrade,omitempty" name:"ReadOnlyGroupForcedUpgrade"`

	// ReadOnlyGroupType=3时必填,已存在的只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// ReadOnlyGroupType=2时必填，新建的只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// ReadOnlyGroupType=2时必填，新建的只读组是否开启延迟剔除功能，1-开启，0-关闭。当只读副本与主实例延迟大于阈值后，自动剔除。
	ReadOnlyGroupIsOfflineDelay *int64 `json:"ReadOnlyGroupIsOfflineDelay,omitempty" name:"ReadOnlyGroupIsOfflineDelay"`

	// ReadOnlyGroupType=2 且 ReadOnlyGroupIsOfflineDelay=1时必填，新建的只读组延迟剔除的阈值。
	ReadOnlyGroupMaxDelayTime *int64 `json:"ReadOnlyGroupMaxDelayTime,omitempty" name:"ReadOnlyGroupMaxDelayTime"`

	// ReadOnlyGroupType=2 且 ReadOnlyGroupIsOfflineDelay=1时必填，新建的只读组延迟剔除后至少保留只读副本的个数。
	ReadOnlyGroupMinInGroup *int64 `json:"ReadOnlyGroupMinInGroup,omitempty" name:"ReadOnlyGroupMinInGroup"`

	// 付费模式，取值支持 PREPAID（预付费），POSTPAID（后付费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 本次购买几个只读实例，默认值为1。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

type CreateReadOnlyDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 只读组类型选项，1-按照一个实例一个只读组的方式发货，2-新建只读组后发货，所有实例都在这个只读组下面， 3-发货的所有实例都在已有的只读组下面
	ReadOnlyGroupType *int64 `json:"ReadOnlyGroupType,omitempty" name:"ReadOnlyGroupType"`

	// 实例内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例磁盘大小，单位GB
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 0-默认不升级主实例，1-强制升级主实例完成ro部署；主实例为非集群版时需要填1，强制升级为集群版。填1 说明您已同意将主实例升级到集群版实例。
	ReadOnlyGroupForcedUpgrade *int64 `json:"ReadOnlyGroupForcedUpgrade,omitempty" name:"ReadOnlyGroupForcedUpgrade"`

	// ReadOnlyGroupType=3时必填,已存在的只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// ReadOnlyGroupType=2时必填，新建的只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// ReadOnlyGroupType=2时必填，新建的只读组是否开启延迟剔除功能，1-开启，0-关闭。当只读副本与主实例延迟大于阈值后，自动剔除。
	ReadOnlyGroupIsOfflineDelay *int64 `json:"ReadOnlyGroupIsOfflineDelay,omitempty" name:"ReadOnlyGroupIsOfflineDelay"`

	// ReadOnlyGroupType=2 且 ReadOnlyGroupIsOfflineDelay=1时必填，新建的只读组延迟剔除的阈值。
	ReadOnlyGroupMaxDelayTime *int64 `json:"ReadOnlyGroupMaxDelayTime,omitempty" name:"ReadOnlyGroupMaxDelayTime"`

	// ReadOnlyGroupType=2 且 ReadOnlyGroupIsOfflineDelay=1时必填，新建的只读组延迟剔除后至少保留只读副本的个数。
	ReadOnlyGroupMinInGroup *int64 `json:"ReadOnlyGroupMinInGroup,omitempty" name:"ReadOnlyGroupMinInGroup"`

	// 付费模式，取值支持 PREPAID（预付费），POSTPAID（后付费）。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 本次购买几个只读实例，默认值为1。
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 安全组列表，填写形如sg-xxx的安全组ID
	SecurityGroupList []*string `json:"SecurityGroupList,omitempty" name:"SecurityGroupList"`

	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 新建实例绑定的标签集合
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
}

func (r *CreateReadOnlyDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Zone")
	delete(f, "ReadOnlyGroupType")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "ReadOnlyGroupForcedUpgrade")
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "ReadOnlyGroupIsOfflineDelay")
	delete(f, "ReadOnlyGroupMaxDelayTime")
	delete(f, "ReadOnlyGroupMinInGroup")
	delete(f, "InstanceChargeType")
	delete(f, "GoodsNum")
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "Period")
	delete(f, "SecurityGroupList")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "ResourceTags")
	delete(f, "Collation")
	delete(f, "TimeZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReadOnlyDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReadOnlyDBInstancesResponseParams struct {
	// 订单名称数组
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateReadOnlyDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateReadOnlyDBInstancesResponseParams `json:"Response"`
}

func (r *CreateReadOnlyDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReadOnlyDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossBackupAddr struct {
	// 跨地域备份目标地域
	CrossRegion *string `json:"CrossRegion,omitempty" name:"CrossRegion"`

	// 跨地域备份内网下载地址
	CrossInternalAddr *string `json:"CrossInternalAddr,omitempty" name:"CrossInternalAddr"`

	// 跨地域备份外网下载地址
	CrossExternalAddr *string `json:"CrossExternalAddr,omitempty" name:"CrossExternalAddr"`
}

type CrossRegionStatus struct {
	// 跨地域备份目标地域
	CrossRegion *string `json:"CrossRegion,omitempty" name:"CrossRegion"`

	// 备份跨地域的同步状态 0-创建中；1-成功；2-失败；4-同步中
	CrossStatus *int64 `json:"CrossStatus,omitempty" name:"CrossStatus"`
}

type DBCreateInfo struct {
	// 数据库名
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// 字符集。可通过接口DescribeDBCharsets查到支持的字符集，不填默认为Chinese_PRC_CI_AS。
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 数据库账号权限信息
	Accounts []*AccountPrivilege `json:"Accounts,omitempty" name:"Accounts"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DBDetail struct {
	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字符集
	Charset *string `json:"Charset,omitempty" name:"Charset"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 数据库创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库状态。1--创建中， 2--运行中， 3--修改中，-1--删除中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 数据库账号权限信息
	Accounts []*AccountPrivilege `json:"Accounts,omitempty" name:"Accounts"`

	// 内部状态。ONLINE表示运行中
	InternalStatus *string `json:"InternalStatus,omitempty" name:"InternalStatus"`
}

type DBInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例所在项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所在地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 实例所在可用区ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 实例所在私有网络ID，基础网络时为 0
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 实例所在私有网络子网ID，基础网络时为 0
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例状态。取值范围： <li>1：申请中</li> <li>2：运行中</li> <li>3：受限运行中 (主备切换中)</li> <li>4：已隔离</li> <li>5：回收中</li> <li>6：已回收</li> <li>7：任务执行中 (实例做备份、回档等操作)</li> <li>8：已下线</li> <li>9：实例扩容中</li> <li>10：实例迁移中</li> <li>11：只读</li> <li>12：重启中</li>  <li>13：实例修改中且待切换</li> <li>14：订阅发布创建中</li> <li>15：订阅发布修改中</li> <li>16：实例修改中且切换中</li> <li>17：创建RO副本中</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例访问IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例访问端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 实例创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 实例计费开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例计费结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例隔离时间
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 实例内存大小，单位G
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例已经使用存储空间大小，单位G
	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`

	// 实例存储空间大小，单位G
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例版本
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 实例续费标记，0-正常续费，1-自动续费，2-到期不续费
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 实例高可用， 1-双机高可用，2-单机，3-跨可用区，4-集群跨可用区，5-集群，9-自研机房
	Model *int64 `json:"Model,omitempty" name:"Model"`

	// 实例所在地域名称，如 ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所在可用区名称，如 ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 备份时间点
	BackupTime *string `json:"BackupTime,omitempty" name:"BackupTime"`

	// 实例付费模式， 0-按量计费，1-包年包月
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例唯一UID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 实例cpu核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 实例版本代号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 物理机代号
	Type *string `json:"Type,omitempty" name:"Type"`

	// 计费ID
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 实例所属VPC的唯一字符串ID，格式如：vpc-xxx，基础网络时为空字符串
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 实例所属子网的唯一字符串ID，格式如： subnet-xxx，基础网络时为空字符串
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 实例隔离操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateOperator *string `json:"IsolateOperator,omitempty" name:"IsolateOperator"`

	// 发布订阅标识，SUB-订阅实例，PUB-发布实例，空值-没有发布订阅的普通实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubFlag *string `json:"SubFlag,omitempty" name:"SubFlag"`

	// 只读标识，RO-只读实例，MASTER-有RO实例的主实例，空值-没有只读组的非RO实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ROFlag *string `json:"ROFlag,omitempty" name:"ROFlag"`

	// 容灾类型，MIRROR-镜像，ALWAYSON-AlwaysOn, SINGLE-单例
	// 注意：此字段可能返回 null，表示取不到有效值。
	HAFlag *string `json:"HAFlag,omitempty" name:"HAFlag"`

	// 实例绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTags []*ResourceTag `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 备份模式，master_pkg-主节点打包备份(默认) ；master_no_pkg-主节点不打包备份；slave_pkg-从节点打包备份(always on集群有效)；slave_no_pkg-从节点不打包备份(always on集群有效)；只读副本对该值无效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupModel *string `json:"BackupModel,omitempty" name:"BackupModel"`

	// 实例备份信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNote *string `json:"InstanceNote,omitempty" name:"InstanceNote"`

	// 备份周期
	BackupCycle []*int64 `json:"BackupCycle,omitempty" name:"BackupCycle"`

	// 备份周期类型，[daily、weekly、monthly]
	BackupCycleType *string `json:"BackupCycleType,omitempty" name:"BackupCycleType"`

	// 数据(日志)备份保留时间
	BackupSaveDays *int64 `json:"BackupSaveDays,omitempty" name:"BackupSaveDays"`

	// 实例类型 HA-高可用 RO-只读实例 SI-基础版 BI-商业智能服务
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 跨地域备份目的地域，如果为空，则表示未开启跨地域备份
	CrossRegions []*string `json:"CrossRegions,omitempty" name:"CrossRegions"`

	// 跨地域备份状态 enable-开启，disable-关闭
	CrossBackupEnabled *string `json:"CrossBackupEnabled,omitempty" name:"CrossBackupEnabled"`

	// 跨地域备份保留天数，则默认7天
	CrossBackupSaveDays *uint64 `json:"CrossBackupSaveDays,omitempty" name:"CrossBackupSaveDays"`

	// 外网地址域名
	DnsPodDomain *string `json:"DnsPodDomain,omitempty" name:"DnsPodDomain"`

	// 外网端口号
	TgwWanVPort *int64 `json:"TgwWanVPort,omitempty" name:"TgwWanVPort"`

	// 系统字符集排序规则，默认：Chinese_PRC_CI_AS
	Collation *string `json:"Collation,omitempty" name:"Collation"`

	// 系统时区，默认：China Standard Time
	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`

	// 是否跨AZ
	IsDrZone *bool `json:"IsDrZone,omitempty" name:"IsDrZone"`

	// 备可用区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZones *SlaveZones `json:"SlaveZones,omitempty" name:"SlaveZones"`
}

type DBPrivilege struct {
	// 数据库名
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// 数据库权限，ReadWrite表示可读写，ReadOnly表示只读，DBOwner所有者
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type DBPrivilegeModifyInfo struct {
	// 数据库名
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// 权限变更信息。ReadWrite表示可读写，ReadOnly表示只读，Delete表示删除账号对该DB的权限，DBOwner所有者
	Privilege *string `json:"Privilege,omitempty" name:"Privilege"`
}

type DBRemark struct {
	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DBRenameRes struct {
	// 新数据库名称
	NewName *string `json:"NewName,omitempty" name:"NewName"`

	// 老数据库名称
	OldName *string `json:"OldName,omitempty" name:"OldName"`
}

type DatabaseTuple struct {
	// 发布数据库名称
	PublishDatabase *string `json:"PublishDatabase,omitempty" name:"PublishDatabase"`

	// 订阅数据库名称
	SubscribeDatabase *string `json:"SubscribeDatabase,omitempty" name:"SubscribeDatabase"`
}

type DatabaseTupleStatus struct {
	// 发布数据库名称
	PublishDatabase *string `json:"PublishDatabase,omitempty" name:"PublishDatabase"`

	// 订阅数据库名称
	SubscribeDatabase *string `json:"SubscribeDatabase,omitempty" name:"SubscribeDatabase"`

	// 最近一次同步时间
	LastSyncTime *string `json:"LastSyncTime,omitempty" name:"LastSyncTime"`

	// 数据库之间的发布订阅状态 running，success，fail，unknow
	Status *string `json:"Status,omitempty" name:"Status"`
}

type DbNormalDetail struct {
	// 是否已订阅 0：否 1：是
	IsSubscribed *string `json:"IsSubscribed,omitempty" name:"IsSubscribed"`

	// 数据库排序规则
	CollationName *string `json:"CollationName,omitempty" name:"CollationName"`

	// 开启CT之后是否自动清理 0：否 1：是
	IsAutoCleanupOn *string `json:"IsAutoCleanupOn,omitempty" name:"IsAutoCleanupOn"`

	// 是否已启用代理  0：否 1：是
	IsBrokerEnabled *string `json:"IsBrokerEnabled,omitempty" name:"IsBrokerEnabled"`

	// 是否已开启/关闭CDC 0：关闭 1：开启
	IsCdcEnabled *string `json:"IsCdcEnabled,omitempty" name:"IsCdcEnabled"`

	// 是否已启用/ 禁用CT 0：禁用 1：启用
	IsDbChainingOn *string `json:"IsDbChainingOn,omitempty" name:"IsDbChainingOn"`

	// 是否加密 0：否 1：是
	IsEncrypted *string `json:"IsEncrypted,omitempty" name:"IsEncrypted"`

	// 是否全文启用 0：否 1：是
	IsFulltextEnabled *string `json:"IsFulltextEnabled,omitempty" name:"IsFulltextEnabled"`

	// 是否是镜像 0：否 1：是
	IsMirroring *string `json:"IsMirroring,omitempty" name:"IsMirroring"`

	// 是否已发布 0：否 1：是
	IsPublished *string `json:"IsPublished,omitempty" name:"IsPublished"`

	// 是否开启快照 0：否 1：是
	IsReadCommittedSnapshotOn *string `json:"IsReadCommittedSnapshotOn,omitempty" name:"IsReadCommittedSnapshotOn"`

	// 是否可信任 0：否 1：是
	IsTrustworthyOn *string `json:"IsTrustworthyOn,omitempty" name:"IsTrustworthyOn"`

	// 镜像状态
	MirroringState *string `json:"MirroringState,omitempty" name:"MirroringState"`

	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 恢复模式
	RecoveryModelDesc *string `json:"RecoveryModelDesc,omitempty" name:"RecoveryModelDesc"`

	// 保留天数
	RetentionPeriod *string `json:"RetentionPeriod,omitempty" name:"RetentionPeriod"`

	// 数据库状态
	StateDesc *string `json:"StateDesc,omitempty" name:"StateDesc"`

	// 用户类型
	UserAccessDesc *string `json:"UserAccessDesc,omitempty" name:"UserAccessDesc"`

	// 数据库创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DbRollbackTimeInfo struct {
	// 数据库名称
	DBName *string `json:"DBName,omitempty" name:"DBName"`

	// 可回档开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 可回档结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DealInfo struct {
	// 订单名
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 商品数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 关联的流程 ID，可用于查询流程执行状态
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 只有创建实例的订单会填充该字段，表示该订单创建的实例的 ID。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 所属账号
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 实例付费类型
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例用户名数组
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例用户名数组
	UserNames []*string `json:"UserNames,omitempty" name:"UserNames"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupMigrationRequestParams struct {
	// 目标实例ID，由DescribeBackupMigration接口返回
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由DescribeBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

type DeleteBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 目标实例ID，由DescribeBackupMigration接口返回
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由DescribeBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

func (r *DeleteBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBackupMigrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBackupMigrationResponseParams `json:"Response"`
}

func (r *DeleteBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBusinessIntelligenceFileRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 文件名称集合
	FileNameSet []*string `json:"FileNameSet,omitempty" name:"FileNameSet"`
}

type DeleteBusinessIntelligenceFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 文件名称集合
	FileNameSet []*string `json:"FileNameSet,omitempty" name:"FileNameSet"`
}

func (r *DeleteBusinessIntelligenceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBusinessIntelligenceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileNameSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBusinessIntelligenceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBusinessIntelligenceFileResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteBusinessIntelligenceFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBusinessIntelligenceFileResponseParams `json:"Response"`
}

func (r *DeleteBusinessIntelligenceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBusinessIntelligenceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBInstanceRequestParams struct {
	// 实例ID，格式如：mssql-3l3fgqn7 或 mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DeleteDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssql-3l3fgqn7 或 mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBInstanceResponseParams `json:"Response"`
}

func (r *DeleteDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBRequestParams struct {
	// 实例ID，形如mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名数组
	Names []*string `json:"Names,omitempty" name:"Names"`
}

type DeleteDBRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名数组
	Names []*string `json:"Names,omitempty" name:"Names"`
}

func (r *DeleteDBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDBResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBResponseParams `json:"Response"`
}

func (r *DeleteDBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIncrementalMigrationRequestParams struct {
	// 目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量备份导入任务ID，由CreateIncrementalMigration接口返回
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

type DeleteIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量备份导入任务ID，由CreateIncrementalMigration接口返回
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DeleteIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIncrementalMigrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIncrementalMigrationResponseParams `json:"Response"`
}

func (r *DeleteIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrationRequestParams struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type DeleteMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *DeleteMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMigrationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMigrationResponseParams `json:"Response"`
}

func (r *DeleteMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePublishSubscribeRequestParams struct {
	// 发布订阅ID，可通过DescribePublishSubscribe接口获得
	PublishSubscribeId *uint64 `json:"PublishSubscribeId,omitempty" name:"PublishSubscribeId"`

	// 待删除的数据库的订阅发布关系集合
	DatabaseTupleSet []*DatabaseTuple `json:"DatabaseTupleSet,omitempty" name:"DatabaseTupleSet"`
}

type DeletePublishSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 发布订阅ID，可通过DescribePublishSubscribe接口获得
	PublishSubscribeId *uint64 `json:"PublishSubscribeId,omitempty" name:"PublishSubscribeId"`

	// 待删除的数据库的订阅发布关系集合
	DatabaseTupleSet []*DatabaseTuple `json:"DatabaseTupleSet,omitempty" name:"DatabaseTupleSet"`
}

func (r *DeletePublishSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePublishSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PublishSubscribeId")
	delete(f, "DatabaseTupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePublishSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePublishSubscribeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePublishSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *DeletePublishSubscribeResponseParams `json:"Response"`
}

func (r *DeletePublishSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePublishSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 账号名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// createTime,updateTime,passTime" note:"排序字段，默认按照账号创建时间倒序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序规则（desc-降序，asc-升序），默认desc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 账号名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// createTime,updateTime,passTime" note:"排序字段，默认按照账号创建时间倒序
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序规则（desc-降序，asc-升序），默认desc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Name")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountsResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账户信息列表
	Accounts []*AccountDetail `json:"Accounts,omitempty" name:"Accounts"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountsResponseParams `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupByFlowIdRequestParams struct {
	// 实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 创建备份流程ID，可通过 [CreateBackup](https://cloud.tencent.com/document/product/238/19946) 接口获取
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeBackupByFlowIdRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 创建备份流程ID，可通过 [CreateBackup](https://cloud.tencent.com/document/product/238/19946) 接口获取
	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeBackupByFlowIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupByFlowIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupByFlowIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupByFlowIdResponseParams struct {
	// 备份文件唯一标识，RestoreInstance接口会用到该字段，对于单库备份文件只返回第一条记录的备份文件唯一标识；单库备份文件需要通过DescribeBackupFiles接口获取全部记录的可回档的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 文件名，对于单库备份文件只返回第一条记录的文件名；单库备份文件需要通过DescribeBackupFiles接口获取全部记录的文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 备份任务名称，可自定义
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 备份开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 备份结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 文件大小，单位 KB，对于单库备份文件只返回第一条记录的文件大小；单库备份文件需要通过DescribeBackupFiles接口获取全部记录的文件大小
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 备份策略，0-实例备份；1-多库备份；实例状态是0-创建中时，该字段为默认值0，无实际意义
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 备份文件状态，0-创建中；1-成功；2-失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 备份方式，0-定时备份；1-手动临时备份；实例状态是0-创建中时，该字段为默认值0，无实际意义
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`

	// DB列表，对于单库备份文件只返回第一条记录包含的库名；单库备份文件需要通过DescribeBackupFiles接口获取全部记录的库名。
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// 内网下载地址，对于单库备份文件只返回第一条记录的内网下载地址；单库备份文件需要通过DescribeBackupFiles接口获取全部记录的下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址，对于单库备份文件只返回第一条记录的外网下载地址；单库备份文件需要通过DescribeBackupFiles接口获取全部记录的下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// 聚合Id，对于打包备份文件不返回此值。通过此值调用DescribeBackupFiles接口，获取单库备份文件的详细信息
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupByFlowIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupByFlowIdResponseParams `json:"Response"`
}

func (r *DescribeBackupByFlowIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupByFlowIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupCommandRequestParams struct {
	// 备份文件类型，FULL-全量备份，FULL_LOG-全量备份需要日志增量，FULL_DIFF-全量备份需要差异增量，LOG-日志备份，DIFF-差异备份
	BackupFileType *string `json:"BackupFileType,omitempty" name:"BackupFileType"`

	// 数据库名称
	DataBaseName *string `json:"DataBaseName,omitempty" name:"DataBaseName"`

	// 是否需要恢复，NO-不需要，YES-需要
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// 备份文件保存的路径；如果不填则默认在D:\\
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`
}

type DescribeBackupCommandRequest struct {
	*tchttp.BaseRequest
	
	// 备份文件类型，FULL-全量备份，FULL_LOG-全量备份需要日志增量，FULL_DIFF-全量备份需要差异增量，LOG-日志备份，DIFF-差异备份
	BackupFileType *string `json:"BackupFileType,omitempty" name:"BackupFileType"`

	// 数据库名称
	DataBaseName *string `json:"DataBaseName,omitempty" name:"DataBaseName"`

	// 是否需要恢复，NO-不需要，YES-需要
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// 备份文件保存的路径；如果不填则默认在D:\\
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`
}

func (r *DescribeBackupCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupFileType")
	delete(f, "DataBaseName")
	delete(f, "IsRecovery")
	delete(f, "LocalPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupCommandResponseParams struct {
	// 创建备份命令
	Command *string `json:"Command,omitempty" name:"Command"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupCommandResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupCommandResponseParams `json:"Response"`
}

func (r *DescribeBackupCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesRequestParams struct {
	// 实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 聚合ID, 可通过接口DescribeBackups获取
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照备份的库名称筛选，不填则不筛选此项
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 列表项排序，目前只按照备份大小排序（desc-降序，asc-升序），默认desc
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

type DescribeBackupFilesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 聚合ID, 可通过接口DescribeBackups获取
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照备份的库名称筛选，不填则不筛选此项
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 列表项排序，目前只按照备份大小排序（desc-降序，asc-升序），默认desc
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
}

func (r *DescribeBackupFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GroupId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "DatabaseName")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupFilesResponseParams struct {
	// 备份总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 备份文件列表详情
	BackupFiles []*BackupFile `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupFilesResponseParams `json:"Response"`
}

func (r *DescribeBackupFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 导入任务名称
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// 备份文件名称
	BackupFileName *string `json:"BackupFileName,omitempty" name:"BackupFileName"`

	// 导入任务状态集合
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 导入任务恢复类型，FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，用户上传
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 分页，页大小，默认值：100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，页数，默认值：0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，name；createTime；startTime；endTime，默认按照createTime递增排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc-递减排序，asc-递增排序。默认按照asc排序，且在OrderBy为有效值时，本参数有效
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 导入任务名称
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// 备份文件名称
	BackupFileName *string `json:"BackupFileName,omitempty" name:"BackupFileName"`

	// 导入任务状态集合
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 导入任务恢复类型，FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，用户上传
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 分页，页大小，默认值：100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，页数，默认值：0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，name；createTime；startTime；endTime，默认按照createTime递增排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc-递减排序，asc-递增排序。默认按照asc排序，且在OrderBy为有效值时，本参数有效
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "MigrationName")
	delete(f, "BackupFileName")
	delete(f, "StatusSet")
	delete(f, "RecoveryType")
	delete(f, "UploadType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupMigrationResponseParams struct {
	// 迁移任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 迁移任务集合
	BackupMigrationSet []*Migration `json:"BackupMigrationSet,omitempty" name:"BackupMigrationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupMigrationResponseParams `json:"Response"`
}

func (r *DescribeBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUploadSizeRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

type DescribeBackupUploadSizeRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DescribeBackupUploadSizeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUploadSizeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupUploadSizeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupUploadSizeResponseParams struct {
	// 已上传的备份的信息
	CosUploadBackupFileSet []*CosUploadBackupFile `json:"CosUploadBackupFileSet,omitempty" name:"CosUploadBackupFileSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupUploadSizeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupUploadSizeResponseParams `json:"Response"`
}

func (r *DescribeBackupUploadSizeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupUploadSizeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsRequestParams struct {
	// 开始时间(yyyy-MM-dd HH:mm:ss)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间(yyyy-MM-dd HH:mm:ss)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照备份名称筛选，不填则不筛选此项
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 按照备份策略筛选，0-实例备份，1-多库备份，不填则不筛选此项
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 按照备份方式筛选，0-后台自动定时备份，1-用户手动临时备份，2-定期备份，不填则不筛选此项
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`

	// 按照备份ID筛选，不填则不筛选此项
	BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

	// 按照备份的库名称筛选，不填则不筛选此项
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 是否分组查询，默认是0，单库备份情况下 0-兼容老方式不分组，1-单库备份分组后展示
	Group *int64 `json:"Group,omitempty" name:"Group"`

	// 备份类型，1-数据备份，2-日志备份，默认值为1
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 按照备份文件形式筛选，pkg-打包备份文件，single-单库备份文件
	BackupFormat *string `json:"BackupFormat,omitempty" name:"BackupFormat"`
}

type DescribeBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间(yyyy-MM-dd HH:mm:ss)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间(yyyy-MM-dd HH:mm:ss)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按照备份名称筛选，不填则不筛选此项
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 按照备份策略筛选，0-实例备份，1-多库备份，不填则不筛选此项
	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`

	// 按照备份方式筛选，0-后台自动定时备份，1-用户手动临时备份，2-定期备份，不填则不筛选此项
	BackupWay *int64 `json:"BackupWay,omitempty" name:"BackupWay"`

	// 按照备份ID筛选，不填则不筛选此项
	BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

	// 按照备份的库名称筛选，不填则不筛选此项
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 是否分组查询，默认是0，单库备份情况下 0-兼容老方式不分组，1-单库备份分组后展示
	Group *int64 `json:"Group,omitempty" name:"Group"`

	// 备份类型，1-数据备份，2-日志备份，默认值为1
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 按照备份文件形式筛选，pkg-打包备份文件，single-单库备份文件
	BackupFormat *string `json:"BackupFormat,omitempty" name:"BackupFormat"`
}

func (r *DescribeBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BackupName")
	delete(f, "Strategy")
	delete(f, "BackupWay")
	delete(f, "BackupId")
	delete(f, "DatabaseName")
	delete(f, "Group")
	delete(f, "Type")
	delete(f, "BackupFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBackupsResponseParams struct {
	// 备份总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 备份列表详情
	Backups []*Backup `json:"Backups,omitempty" name:"Backups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBackupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBackupsResponseParams `json:"Response"`
}

func (r *DescribeBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessIntelligenceFileRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 迁移任务状态集合,1-初始化待部署 2-部署中 3-部署成功 4-部署失败
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 文件类型 FLAT-平面文件，SSIS商业智能服务项目文件
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 分页，页大小，范围1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页,页数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，可选值file_name,create_time,start_time
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeBusinessIntelligenceFileRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 迁移任务状态集合,1-初始化待部署 2-部署中 3-部署成功 4-部署失败
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 文件类型 FLAT-平面文件，SSIS商业智能服务项目文件
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 分页，页大小，范围1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页,页数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，可选值file_name,create_time,start_time
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc,asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeBusinessIntelligenceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessIntelligenceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FileName")
	delete(f, "StatusSet")
	delete(f, "FileType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBusinessIntelligenceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBusinessIntelligenceFileResponseParams struct {
	// 文件部署任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 文件部署任务集合
	BackupMigrationSet []*BusinessIntelligenceFile `json:"BackupMigrationSet,omitempty" name:"BackupMigrationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBusinessIntelligenceFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBusinessIntelligenceFileResponseParams `json:"Response"`
}

func (r *DescribeBusinessIntelligenceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBusinessIntelligenceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossRegionZoneRequestParams struct {
	// 实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeCrossRegionZoneRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeCrossRegionZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossRegionZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossRegionZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossRegionZoneResponseParams struct {
	// 备机所在地域的字符串ID，形如：ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 备机所在可用区的字符串ID，形如：ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCrossRegionZoneResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossRegionZoneResponseParams `json:"Response"`
}

func (r *DescribeCrossRegionZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossRegionZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCharsetsRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBCharsetsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBCharsetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCharsetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBCharsetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBCharsetsResponseParams struct {
	// 数据库字符集列表
	DatabaseCharsets []*string `json:"DatabaseCharsets,omitempty" name:"DatabaseCharsets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBCharsetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBCharsetsResponseParams `json:"Response"`
}

func (r *DescribeDBCharsetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBCharsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInterRequestParams struct {
	// 分页，页大小，范围是1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按照实例ID筛选
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照状态筛选 1-互通IP打开中；2-互通IP已经打开；3-加入到互通组中；4-已加入到互通组；5-互通IP回收中；6-互通IP已回收；7-从互通组移除中；8-已从互通组中移除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例版本代号列表
	VersionSet []*string `json:"VersionSet,omitempty" name:"VersionSet"`

	// 实例所在可用区，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 分页，页数，默认是0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDBInstanceInterRequest struct {
	*tchttp.BaseRequest
	
	// 分页，页大小，范围是1-100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 按照实例ID筛选
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 按照状态筛选 1-互通IP打开中；2-互通IP已经打开；3-加入到互通组中；4-已加入到互通组；5-互通IP回收中；6-互通IP已回收；7-从互通组移除中；8-已从互通组中移除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例版本代号列表
	VersionSet []*string `json:"VersionSet,omitempty" name:"VersionSet"`

	// 实例所在可用区，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 分页，页数，默认是0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBInstanceInterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Status")
	delete(f, "VersionSet")
	delete(f, "Zone")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstanceInterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstanceInterResponseParams struct {
	// 互通组内总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 互通组内实例信息详情
	InterInstanceSet []*InterInstance `json:"InterInstanceSet,omitempty" name:"InterInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstanceInterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstanceInterResponseParams `json:"Response"`
}

func (r *DescribeDBInstanceInterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstanceInterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesRequestParams struct {
	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例状态。取值范围：
	// <li>1：申请中</li>
	// <li>2：运行中</li>
	// <li>3：受限运行中 (主备切换中)</li>
	// <li>4：已隔离</li>
	// <li>5：回收中</li>
	// <li>6：已回收</li>
	// <li>7：任务执行中 (实例做备份、回档等操作)</li>
	// <li>8：已下线</li>
	// <li>9：实例扩容中</li>
	// <li>10：实例迁移中</li>
	// <li>11：只读</li>
	// <li>12：重启中</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 一个或者多个实例ID。实例ID，格式如：mssql-si2823jyl
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 付费类型检索 1-包年包月，0-按量计费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例所属VPC的唯一字符串ID，格式如：vpc-xxx，传空字符串(“”)则按照基础网络筛选。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 实例所属子网的唯一字符串ID，格式如： subnet-xxx，传空字符串(“”)则按照基础网络筛选。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例内网地址列表，格式如：172.1.0.12
	VipSet []*string `json:"VipSet,omitempty" name:"VipSet"`

	// 实例名称列表，模糊查询
	InstanceNameSet []*string `json:"InstanceNameSet,omitempty" name:"InstanceNameSet"`

	// 实例版本代号列表，格式如：2008R2，2012SP3等
	VersionSet []*string `json:"VersionSet,omitempty" name:"VersionSet"`

	// 实例可用区，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例标签列表
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 模糊查询关键字，支持实例id、实例名、内网ip
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 实例唯一Uid列表
	UidSet []*string `json:"UidSet,omitempty" name:"UidSet"`

	// 实例类型 HA-高可用 RO-只读实例 SI-基础版 BI-商业智能服务
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DescribeDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例状态。取值范围：
	// <li>1：申请中</li>
	// <li>2：运行中</li>
	// <li>3：受限运行中 (主备切换中)</li>
	// <li>4：已隔离</li>
	// <li>5：回收中</li>
	// <li>6：已回收</li>
	// <li>7：任务执行中 (实例做备份、回档等操作)</li>
	// <li>8：已下线</li>
	// <li>9：实例扩容中</li>
	// <li>10：实例迁移中</li>
	// <li>11：只读</li>
	// <li>12：重启中</li>
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 一个或者多个实例ID。实例ID，格式如：mssql-si2823jyl
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 付费类型检索 1-包年包月，0-按量计费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 实例所属VPC的唯一字符串ID，格式如：vpc-xxx，传空字符串(“”)则按照基础网络筛选。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 实例所属子网的唯一字符串ID，格式如： subnet-xxx，传空字符串(“”)则按照基础网络筛选。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例内网地址列表，格式如：172.1.0.12
	VipSet []*string `json:"VipSet,omitempty" name:"VipSet"`

	// 实例名称列表，模糊查询
	InstanceNameSet []*string `json:"InstanceNameSet,omitempty" name:"InstanceNameSet"`

	// 实例版本代号列表，格式如：2008R2，2012SP3等
	VersionSet []*string `json:"VersionSet,omitempty" name:"VersionSet"`

	// 实例可用区，格式如：ap-guangzhou-2
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例标签列表
	TagKeys []*string `json:"TagKeys,omitempty" name:"TagKeys"`

	// 模糊查询关键字，支持实例id、实例名、内网ip
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 实例唯一Uid列表
	UidSet []*string `json:"UidSet,omitempty" name:"UidSet"`

	// 实例类型 HA-高可用 RO-只读实例 SI-基础版 BI-商业智能服务
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *DescribeDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIdSet")
	delete(f, "PayMode")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "VipSet")
	delete(f, "InstanceNameSet")
	delete(f, "VersionSet")
	delete(f, "Zone")
	delete(f, "TagKeys")
	delete(f, "SearchKey")
	delete(f, "UidSet")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBInstancesResponseParams struct {
	// 符合条件的实例总数。分页返回的话，这个值指的是所有符合条件的实例的个数，而非当前根据Limit和Offset值返回的实例个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例列表
	DBInstances []*DBInstance `json:"DBInstances,omitempty" name:"DBInstances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsRequestParams struct {
	// 实例ID，格式如：mssql-c1nl9rpv或者mssqlro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssql-c1nl9rpv或者mssqlro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSecurityGroupsResponseParams struct {
	// 安全组详情。
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsNormalRequestParams struct {
	// 实例ID，形如mssql-7vfv3rk3
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeDBsNormalRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-7vfv3rk3
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBsNormalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsNormalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBsNormalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsNormalResponseParams struct {
	// 表示当前实例下的数据库总个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回数据库的详细配置信息，例如：数据库是否开启CDC、CT等
	DBList []*DbNormalDetail `json:"DBList,omitempty" name:"DBList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBsNormalResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBsNormalResponseParams `json:"Response"`
}

func (r *DescribeDBsNormalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsNormalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsRequestParams struct {
	// 实例ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序规则（desc-降序，asc-升序），默认desc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeDBsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数据库名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序规则（desc-降序，asc-升序），默认desc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeDBsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Name")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBsResponseParams struct {
	// 数据库数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例数据库列表
	DBInstances []*InstanceDBDetail `json:"DBInstances,omitempty" name:"DBInstances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDBsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBsResponseParams `json:"Response"`
}

func (r *DescribeDBsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowStatusRequestParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeFlowStatusRequest struct {
	*tchttp.BaseRequest
	
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeFlowStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowStatusResponseParams struct {
	// 流程状态，0：成功，1：失败，2：运行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlowStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowStatusResponseParams `json:"Response"`
}

func (r *DescribeFlowStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementalMigrationRequestParams struct {
	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件名称
	BackupFileName *string `json:"BackupFileName,omitempty" name:"BackupFileName"`

	// 导入任务状态集合
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 分页，页大小，默认值：100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，页数，默认值：0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，name；createTime；startTime；endTime，默认按照createTime递增排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc-递减排序，asc-递增排序。默认按照asc排序，且在OrderBy为有效值时，本参数有效
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 增量备份导入任务ID，由CreateIncrementalMigration接口返回
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

type DescribeIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件名称
	BackupFileName *string `json:"BackupFileName,omitempty" name:"BackupFileName"`

	// 导入任务状态集合
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 分页，页大小，默认值：100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，页数，默认值：0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，name；createTime；startTime；endTime，默认按照createTime递增排序。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，desc-递减排序，asc-递增排序。默认按照asc排序，且在OrderBy为有效值时，本参数有效
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`

	// 增量备份导入任务ID，由CreateIncrementalMigration接口返回
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DescribeIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BackupMigrationId")
	delete(f, "InstanceId")
	delete(f, "BackupFileName")
	delete(f, "StatusSet")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIncrementalMigrationResponseParams struct {
	// 增量导入任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 增量导入任务集合
	IncrementalMigrationSet []*Migration `json:"IncrementalMigrationSet,omitempty" name:"IncrementalMigrationSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIncrementalMigrationResponseParams `json:"Response"`
}

func (r *DescribeIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsRequestParams struct {
	// 实例 ID，格式如：mssql-dj5i29c5n，与云数据库控制台页面中显示的实例 ID 相同，可使用 DescribeDBInstances 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页，页数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页，页大小，默认20，最大不超过100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeInstanceParamRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：mssql-dj5i29c5n，与云数据库控制台页面中显示的实例 ID 相同，可使用 DescribeDBInstances 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 分页，页数，默认0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页，页大小，默认20，最大不超过100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstanceParamRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamRecordsResponseParams struct {
	// 符合条件的记录数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数修改记录
	Items []*ParamRecord `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceParamRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamRecordsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsRequestParams struct {
	// 实例 ID，格式如：mssql-dj5i29c5n，与云数据库控制台页面中显示的实例 ID 相同，可使用 DescribeDBInstances 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeInstanceParamsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：mssql-dj5i29c5n，与云数据库控制台页面中显示的实例 ID 相同，可使用 DescribeDBInstances 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceParamsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceParamsResponseParams struct {
	// 实例的参数总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 参数详情
	Items []*ParameterDetail `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceParamsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceParamsResponseParams `json:"Response"`
}

func (r *DescribeInstanceParamsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceSpanRequestParams struct {
	// 实例ID，形如mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeMaintenanceSpanRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeMaintenanceSpanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceSpanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMaintenanceSpanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMaintenanceSpanResponseParams struct {
	// 以周为单位，表示周几允许维护，例如：[1,2,3,4,5,6,7]表示周一到周日均为可维护日。
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 每天可维护的开始时间，例如：10:24标识可维护时间窗10点24分开始。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 每天可维护的持续时间，单位是h，例如：1 表示从可维护的开始时间起持续1小时。
	Span *uint64 `json:"Span,omitempty" name:"Span"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMaintenanceSpanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMaintenanceSpanResponseParams `json:"Response"`
}

func (r *DescribeMaintenanceSpanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMaintenanceSpanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDatabasesRequestParams struct {
	// 迁移源实例的ID，格式如：mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移源实例用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 迁移源实例密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type DescribeMigrationDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 迁移源实例的ID，格式如：mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移源实例用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 迁移源实例密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *DescribeMigrationDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "UserName")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDatabasesResponseParams struct {
	// 数据库数量
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 数据库名称数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateDBSet []*string `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMigrationDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationDatabasesResponseParams `json:"Response"`
}

func (r *DescribeMigrationDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDetailRequestParams struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *DescribeMigrationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDetailResponseParams struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// 迁移任务名称
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 迁移任务所属的用户ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 迁移任务所属的地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 迁移源的类型 1:TencentDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式）
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// 迁移任务的创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 迁移任务的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 迁移任务的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 迁移任务的状态（1:初始化,4:迁移中,5.迁移失败,6.迁移成功）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 迁移任务当前进度
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 迁移类型（1:结构迁移 2:数据迁移 3:增量同步）
	MigrateType *int64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// 迁移源
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// 迁移目标
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// 迁移DB对象 ，离线迁移（SourceType=4或SourceType=5）不使用。
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMigrationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationDetailResponseParams `json:"Response"`
}

func (r *DescribeMigrationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationsRequestParams struct {
	// 状态集合。只要符合集合中某一状态的迁移任务，就会查出来
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 迁移任务的名称，模糊匹配
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果按照关键字排序，可选值为name、createTime、startTime，endTime，status
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，可选值为desc、asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

type DescribeMigrationsRequest struct {
	*tchttp.BaseRequest
	
	// 状态集合。只要符合集合中某一状态的迁移任务，就会查出来
	StatusSet []*int64 `json:"StatusSet,omitempty" name:"StatusSet"`

	// 迁移任务的名称，模糊匹配
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果按照关键字排序，可选值为name、createTime、startTime，endTime，status
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，可选值为desc、asc
	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
}

func (r *DescribeMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StatusSet")
	delete(f, "MigrateName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	delete(f, "OrderByType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationsResponseParams struct {
	// 查询结果的总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询结果的列表
	MigrateTaskSet []*MigrateTask `json:"MigrateTaskSet,omitempty" name:"MigrateTaskSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationsResponseParams `json:"Response"`
}

func (r *DescribeMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersRequestParams struct {
	// 订单数组。发货时会返回订单名字，利用该订单名字调用DescribeOrders接口查询发货情况
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeOrdersRequest struct {
	*tchttp.BaseRequest
	
	// 订单数组。发货时会返回订单名字，利用该订单名字调用DescribeOrders接口查询发货情况
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

func (r *DescribeOrdersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrdersResponseParams struct {
	// 订单信息数组
	Deals []*DealInfo `json:"Deals,omitempty" name:"Deals"`

	// 返回多少个订单的信息
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrdersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrdersResponseParams `json:"Response"`
}

func (r *DescribeOrdersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductConfigRequestParams struct {
	// 可用区英文ID，形如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 购买实例的类型 HA-高可用型(包括双机高可用，alwaysOn集群)，RO-只读副本型，SI-基础版本型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

type DescribeProductConfigRequest struct {
	*tchttp.BaseRequest
	
	// 可用区英文ID，形如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 购买实例的类型 HA-高可用型(包括双机高可用，alwaysOn集群)，RO-只读副本型，SI-基础版本型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
}

func (r *DescribeProductConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "InstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductConfigResponseParams struct {
	// 规格信息数组
	SpecInfoList []*SpecInfo `json:"SpecInfoList,omitempty" name:"SpecInfoList"`

	// 返回总共多少条数据
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProductConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductConfigResponseParams `json:"Response"`
}

func (r *DescribeProductConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsRequestParams struct {
	// 项目ID，可通过控制台项目管理中查看
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeProjectSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID，可通过控制台项目管理中查看
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectSecurityGroupsResponseParams struct {
	// 安全组详情。
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitempty" name:"SecurityGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeProjectSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublishSubscribeRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 订阅/发布实例ID，与InstanceId是发布实例还是订阅实例有关；当InstanceId为发布实例时，本字段按照订阅实例ID做筛选；当InstanceId为订阅实例时，本字段按照发布实例ID做筛选；
	PubOrSubInstanceId *string `json:"PubOrSubInstanceId,omitempty" name:"PubOrSubInstanceId"`

	// 订阅/发布实例内网IP，与InstanceId是发布实例还是订阅实例有关；当InstanceId为发布实例时，本字段按照订阅实例内网IP做筛选；当InstanceId为订阅实例时，本字段按照发布实例内网IP做筛选；
	PubOrSubInstanceIp *string `json:"PubOrSubInstanceIp,omitempty" name:"PubOrSubInstanceIp"`

	// 订阅发布ID，用于筛选
	PublishSubscribeId *uint64 `json:"PublishSubscribeId,omitempty" name:"PublishSubscribeId"`

	// 订阅发布名字，用于筛选
	PublishSubscribeName *string `json:"PublishSubscribeName,omitempty" name:"PublishSubscribeName"`

	// 发布库名字，用于筛选
	PublishDBName *string `json:"PublishDBName,omitempty" name:"PublishDBName"`

	// 订阅库名字，用于筛选
	SubscribeDBName *string `json:"SubscribeDBName,omitempty" name:"SubscribeDBName"`

	// 分页，页数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页，页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePublishSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 订阅/发布实例ID，与InstanceId是发布实例还是订阅实例有关；当InstanceId为发布实例时，本字段按照订阅实例ID做筛选；当InstanceId为订阅实例时，本字段按照发布实例ID做筛选；
	PubOrSubInstanceId *string `json:"PubOrSubInstanceId,omitempty" name:"PubOrSubInstanceId"`

	// 订阅/发布实例内网IP，与InstanceId是发布实例还是订阅实例有关；当InstanceId为发布实例时，本字段按照订阅实例内网IP做筛选；当InstanceId为订阅实例时，本字段按照发布实例内网IP做筛选；
	PubOrSubInstanceIp *string `json:"PubOrSubInstanceIp,omitempty" name:"PubOrSubInstanceIp"`

	// 订阅发布ID，用于筛选
	PublishSubscribeId *uint64 `json:"PublishSubscribeId,omitempty" name:"PublishSubscribeId"`

	// 订阅发布名字，用于筛选
	PublishSubscribeName *string `json:"PublishSubscribeName,omitempty" name:"PublishSubscribeName"`

	// 发布库名字，用于筛选
	PublishDBName *string `json:"PublishDBName,omitempty" name:"PublishDBName"`

	// 订阅库名字，用于筛选
	SubscribeDBName *string `json:"SubscribeDBName,omitempty" name:"SubscribeDBName"`

	// 分页，页数
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页，页大小
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePublishSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublishSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PubOrSubInstanceId")
	delete(f, "PubOrSubInstanceIp")
	delete(f, "PublishSubscribeId")
	delete(f, "PublishSubscribeName")
	delete(f, "PublishDBName")
	delete(f, "SubscribeDBName")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublishSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublishSubscribeResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 发布订阅列表
	PublishSubscribeSet []*PublishSubscribe `json:"PublishSubscribeSet,omitempty" name:"PublishSubscribeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePublishSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublishSubscribeResponseParams `json:"Response"`
}

func (r *DescribePublishSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublishSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupByReadOnlyInstanceRequestParams struct {
	// 实例ID，格式如：mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeReadOnlyGroupByReadOnlyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeReadOnlyGroupByReadOnlyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupByReadOnlyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReadOnlyGroupByReadOnlyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupByReadOnlyInstanceResponseParams struct {
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 只读组的地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 只读组的可用区ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 是否启动超时剔除功能 ,0-不开启剔除功能，1-开启剔除功能
	IsOfflineDelay *int64 `json:"IsOfflineDelay,omitempty" name:"IsOfflineDelay"`

	// 启动超时剔除功能后，使用的超时阈值，单位是秒
	ReadOnlyMaxDelayTime *int64 `json:"ReadOnlyMaxDelayTime,omitempty" name:"ReadOnlyMaxDelayTime"`

	// 启动超时剔除功能后，只读组至少保留的只读副本数
	MinReadOnlyInGroup *int64 `json:"MinReadOnlyInGroup,omitempty" name:"MinReadOnlyInGroup"`

	// 只读组vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 只读组vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 只读组在私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 只读组在私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 主实例ID，形如mssql-sgeshe3th
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// 主实例的地域ID
	MasterRegionId *string `json:"MasterRegionId,omitempty" name:"MasterRegionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReadOnlyGroupByReadOnlyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReadOnlyGroupByReadOnlyInstanceResponseParams `json:"Response"`
}

func (r *DescribeReadOnlyGroupByReadOnlyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupByReadOnlyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupDetailsRequestParams struct {
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读组ID，格式如：mssqlrg-3l3fgqn7
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type DescribeReadOnlyGroupDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读组ID，格式如：mssqlrg-3l3fgqn7
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *DescribeReadOnlyGroupDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReadOnlyGroupDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupDetailsResponseParams struct {
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 只读组的地域ID，与主实例相同
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 只读组的可用区ID，与主实例相同
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 是否启动超时剔除功能，0-不开启剔除功能，1-开启剔除功能
	IsOfflineDelay *int64 `json:"IsOfflineDelay,omitempty" name:"IsOfflineDelay"`

	// 启动超时剔除功能后，使用的超时阈值
	ReadOnlyMaxDelayTime *int64 `json:"ReadOnlyMaxDelayTime,omitempty" name:"ReadOnlyMaxDelayTime"`

	// 启动超时剔除功能后，至少只读组保留的只读副本数
	MinReadOnlyInGroup *int64 `json:"MinReadOnlyInGroup,omitempty" name:"MinReadOnlyInGroup"`

	// 只读组vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 只读组vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 只读组私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 只读组私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 只读实例副本集合
	ReadOnlyInstanceSet []*ReadOnlyInstance `json:"ReadOnlyInstanceSet,omitempty" name:"ReadOnlyInstanceSet"`

	// 只读组状态: 1-申请成功运行中，5-申请中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 主实例ID，形如mssql-sgeshe3th
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReadOnlyGroupDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReadOnlyGroupDetailsResponseParams `json:"Response"`
}

func (r *DescribeReadOnlyGroupDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupListRequestParams struct {
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeReadOnlyGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeReadOnlyGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReadOnlyGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReadOnlyGroupListResponseParams struct {
	// 只读组列表
	ReadOnlyGroupSet []*ReadOnlyGroup `json:"ReadOnlyGroupSet,omitempty" name:"ReadOnlyGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeReadOnlyGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReadOnlyGroupListResponseParams `json:"Response"`
}

func (r *DescribeReadOnlyGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReadOnlyGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsRequestParams struct {

}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRegionsResponseParams struct {
	// 返回地域信息总的条目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 地域信息数组
	RegionSet []*RegionInfo `json:"RegionSet,omitempty" name:"RegionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRegionsResponseParams `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要查询的数据库列表
	DBs []*string `json:"DBs,omitempty" name:"DBs"`
}

type DescribeRollbackTimeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 需要查询的数据库列表
	DBs []*string `json:"DBs,omitempty" name:"DBs"`
}

func (r *DescribeRollbackTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRollbackTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRollbackTimeResponseParams struct {
	// 数据库可回档实例信息
	Details []*DbRollbackTimeInfo `json:"Details,omitempty" name:"Details"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRollbackTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRollbackTimeResponseParams `json:"Response"`
}

func (r *DescribeRollbackTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRollbackTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowlogsRequestParams struct {
	// 实例ID，形如mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeSlowlogsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页返回，页编号，默认值为第0页
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowlogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowlogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowlogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowlogsResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 慢查询日志信息列表
	Slowlogs []*SlowlogInfo `json:"Slowlogs,omitempty" name:"Slowlogs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSlowlogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowlogsResponseParams `json:"Response"`
}

func (r *DescribeSlowlogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowlogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadBackupInfoRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

type DescribeUploadBackupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

func (r *DescribeUploadBackupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadBackupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadBackupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadBackupInfoResponseParams struct {
	// 存储桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 存储桶地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 存储路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 临时密钥ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥Key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时密钥Token
	XCosSecurityToken *string `json:"XCosSecurityToken,omitempty" name:"XCosSecurityToken"`

	// 临时密钥开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 临时密钥到期时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUploadBackupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadBackupInfoResponseParams `json:"Response"`
}

func (r *DescribeUploadBackupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadBackupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadIncrementalInfoRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

type DescribeUploadIncrementalInfoRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *DescribeUploadIncrementalInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadIncrementalInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUploadIncrementalInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUploadIncrementalInfoResponseParams struct {
	// 存储桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 存储桶地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 存储路径
	Path *string `json:"Path,omitempty" name:"Path"`

	// 临时密钥ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥Key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时密钥Token
	XCosSecurityToken *string `json:"XCosSecurityToken,omitempty" name:"XCosSecurityToken"`

	// 临时密钥开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 临时密钥到期时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUploadIncrementalInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUploadIncrementalInfoResponseParams `json:"Response"`
}

func (r *DescribeUploadIncrementalInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUploadIncrementalInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {

}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 返回多少个可用区信息
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可用区数组
	ZoneSet []*ZoneInfo `json:"ZoneSet,omitempty" name:"ZoneSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsRequestParams struct {
	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例ID 列表，一个或者多个实例ID组成的数组。多个实例必须是同一个地域，同一个可用区，同一个项目下的。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type DisassociateSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例ID 列表，一个或者多个实例ID组成的数组。多个实例必须是同一个地域，同一个可用区，同一个项目下的。
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *DisassociateSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisassociateSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileAction struct {
	// 支持的所有操作，值包括：view(查看列表) remark(修改备注)，deploy(部署)，delete(删除文件)
	AllAction []*string `json:"AllAction,omitempty" name:"AllAction"`

	// 当前状态允许的操作，AllAction的子集,为空表示禁止所有操作
	AllowedAction []*string `json:"AllowedAction,omitempty" name:"AllowedAction"`
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesRequestParams struct {
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 内存大小，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例容量大小，单位：GB。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 计费类型，取值支持 PREPAID，POSTPAID。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 购买时长，单位：月。取值为1到48，默认为1
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 一次性购买的实例数量。取值1-100，默认取值为1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// sqlserver版本，目前只支持：2008R2（SQL Server 2008 Enterprise），2012SP3（SQL Server 2012 Enterprise），2016SP1（SQL Server 2016 Enterprise），201602（SQL Server 2016 Standard）2017（SQL Server 2017 Enterprise）版本。默认为2008R2版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 预购买实例的CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 购买实例的类型 HA-高可用型(包括双机高可用，alwaysOn集群)，RO-只读副本，SI-基础版，默认取值HA
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 购买实例的宿主机类型，PM-物理机, CLOUD_PREMIUM-虚拟机高性能云盘，CLOUD_SSD-虚拟机SSD云盘，默认取值PM
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

type InquiryPriceCreateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 内存大小，单位：GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例容量大小，单位：GB。
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 计费类型，取值支持 PREPAID，POSTPAID。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 购买时长，单位：月。取值为1到48，默认为1
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 一次性购买的实例数量。取值1-100，默认取值为1
	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// sqlserver版本，目前只支持：2008R2（SQL Server 2008 Enterprise），2012SP3（SQL Server 2012 Enterprise），2016SP1（SQL Server 2016 Enterprise），201602（SQL Server 2016 Standard）2017（SQL Server 2017 Enterprise）版本。默认为2008R2版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 预购买实例的CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 购买实例的类型 HA-高可用型(包括双机高可用，alwaysOn集群)，RO-只读副本，SI-基础版，默认取值HA
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 购买实例的宿主机类型，PM-物理机, CLOUD_PREMIUM-虚拟机高性能云盘，CLOUD_SSD-虚拟机SSD云盘，默认取值PM
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *InquiryPriceCreateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "InstanceChargeType")
	delete(f, "Period")
	delete(f, "GoodsNum")
	delete(f, "DBVersion")
	delete(f, "Cpu")
	delete(f, "InstanceType")
	delete(f, "MachineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateDBInstancesResponseParams struct {
	// 未打折前价格，其值除以100表示多少钱。例如10010表示100.10元
	OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 实际需要支付的价格，其值除以100表示多少钱。例如10010表示100.10元
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceCreateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateDBInstancesResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDBInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费周期。按月续费最多48个月。默认查询续费一个月的价格
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 续费周期单位。month表示按月续费，当前只支持按月付费查询价格
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

type InquiryPriceRenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费周期。按月续费最多48个月。默认查询续费一个月的价格
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 续费周期单位。month表示按月续费，当前只支持按月付费查询价格
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
}

func (r *InquiryPriceRenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Period")
	delete(f, "TimeUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewDBInstanceResponseParams struct {
	// 未打折的原价，其值除以100表示最终的价格。例如10094表示100.94元
	OriginalPrice *uint64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 实际需要支付价格，其值除以100表示最终的价格。例如10094表示100.94元
	Price *uint64 `json:"Price,omitempty" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceRenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewDBInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceRequestParams struct {
	// 实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例升级后的内存大小，单位GB，其值不能比当前实例内存小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例升级后的磁盘大小，单位GB，其值不能比当前实例磁盘小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例升级后的CPU核心数，其值不能比当前实例CPU小
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

type InquiryPriceUpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例升级后的内存大小，单位GB，其值不能比当前实例内存小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例升级后的磁盘大小，单位GB，其值不能比当前实例磁盘小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 实例升级后的CPU核心数，其值不能比当前实例CPU小
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
}

func (r *InquiryPriceUpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "Cpu")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceUpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceUpgradeDBInstanceResponseParams struct {
	// 未打折的原价，其值除以100表示最终的价格。例如10094表示100.94元
	OriginalPrice *int64 `json:"OriginalPrice,omitempty" name:"OriginalPrice"`

	// 实际需要支付价格，其值除以100表示最终的价格。例如10094表示100.94元
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquiryPriceUpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceUpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *InquiryPriceUpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceUpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceDBDetail struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库信息列表
	DBDetails []*DBDetail `json:"DBDetails,omitempty" name:"DBDetails"`
}

type InstanceRenewInfo struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例续费标记。0：正常续费，1：自动续费，2：到期不续
	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InterInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例互通IP，用于加入互通组后访问
	InterVip *string `json:"InterVip,omitempty" name:"InterVip"`

	// 实例互通端口，用于加入互通组后访问
	InterPort *int64 `json:"InterPort,omitempty" name:"InterPort"`

	// 实例互通状态，1 -互通ipprot打开中 2 -互通ipprot已经打开 3 -已经打开互通ip的实例加入到互通组中 4 -已经打开互通ip的实例已加入到互通组 5 -互通ipprot回收中 6 -互通ipprot已回收 7 -已回收的实例从互通组中移除中 8 -已回收的实例从互通组中已经移除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例所在地域名称，如 ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例所在可用区名称，如 ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 实例版本代号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 实例版本
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 实例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例访问IP
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例访问端口
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type InterInstanceFlow struct {
	// 实例ID，例如：mssql-sdf32n1d
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例开通或者关闭互通组的流程ID，FlowId小于0-开通或者关闭失败，反之则成功。
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type MigrateDB struct {
	// 迁移数据库的名称
	DBName *string `json:"DBName,omitempty" name:"DBName"`
}

type MigrateDetail struct {
	// 当前环节的名称
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 当前环节的进度（单位是%）
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
}

type MigrateSource struct {
	// 迁移源实例的ID，MigrateType=1(TencentDB for SQLServers)时使用，格式如：mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移源Cvm的ID，MigrateType=2(云服务器自建SQLServer数据库)时使用
	CvmId *string `json:"CvmId,omitempty" name:"CvmId"`

	// 迁移源Cvm的Vpc网络标识，MigrateType=2(云服务器自建SQLServer数据库)时使用，格式如：vpc-6ys9ont9
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 迁移源Cvm的Vpc下的子网标识，MigrateType=2(云服务器自建SQLServer数据库)时使用，格式如：subnet-h9extioi
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 用户名，MigrateType=1或MigrateType=2使用
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 密码，MigrateType=1或MigrateType=2使用
	Password *string `json:"Password,omitempty" name:"Password"`

	// 迁移源Cvm自建库的内网IP，MigrateType=2(云服务器自建SQLServer数据库)时使用
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 迁移源Cvm自建库的端口号，MigrateType=2(云服务器自建SQLServer数据库)时使用
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 离线迁移的源备份地址，MigrateType=4或MigrateType=5使用
	Url []*string `json:"Url,omitempty" name:"Url"`

	// 离线迁移的源备份密码，MigrateType=4或MigrateType=5使用
	UrlPassword *string `json:"UrlPassword,omitempty" name:"UrlPassword"`
}

type MigrateTarget struct {
	// 迁移目标实例的ID，格式如：mssql-si2823jyl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移目标实例的用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 迁移目标实例的密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type MigrateTask struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// 迁移任务名称
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 迁移任务所属的用户ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 迁移任务所属的地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 迁移源的类型 1:TencentDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式）
	SourceType *int64 `json:"SourceType,omitempty" name:"SourceType"`

	// 迁移任务的创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 迁移任务的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 迁移任务的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 迁移任务的状态（1:初始化,4:迁移中,5.迁移失败,6.迁移成功,7已中止,8已删除,9中止中,10完成中,11中止失败,12完成失败）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 是否迁移任务经过检查（0:未校验,1:校验成功,2:校验失败,3:校验中）
	CheckFlag *uint64 `json:"CheckFlag,omitempty" name:"CheckFlag"`

	// 迁移任务当前进度（单位%）
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 迁移任务进度细节
	MigrateDetail *MigrateDetail `json:"MigrateDetail,omitempty" name:"MigrateDetail"`
}

type Migration struct {
	// 备份导入任务ID 或 增量导入任务ID
	MigrationId *string `json:"MigrationId,omitempty" name:"MigrationId"`

	// 备份导入名称，增量导入任务该字段为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// 应用ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 迁移目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 迁移任务恢复类型
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// 备份用户上传类型，COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，用户上传
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// 备份文件列表，UploadType确定，COS_URL则保存URL，COS_UPLOAD则保存备份名称
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// 迁移任务状态，2-创建完成，7-全量导入中，8-等待增量，9-导入成功，10-导入失败，12-增量导入中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 迁移任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 迁移任务开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 迁移任务结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 说明信息
	Message *string `json:"Message,omitempty" name:"Message"`

	// 迁移细节
	Detail *MigrationDetail `json:"Detail,omitempty" name:"Detail"`

	// 当前状态允许的操作
	Action *MigrationAction `json:"Action,omitempty" name:"Action"`

	// 是否是最终恢复，全量导入任务该字段为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// 重命名的数据库名称集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBRename []*DBRenameRes `json:"DBRename,omitempty" name:"DBRename"`
}

type MigrationAction struct {
	// 支持的所有操作，值包括：view(查看任务) ，modify(修改任务)， start(启动任务)，incremental(创建增量任务)，delete(删除任务)，upload(获取上传权限)。
	AllAction []*string `json:"AllAction,omitempty" name:"AllAction"`

	// 当前状态允许的操作，AllAction的子集,为空表示禁止所有操作
	AllowedAction []*string `json:"AllowedAction,omitempty" name:"AllowedAction"`
}

type MigrationDetail struct {
	// 总步骤数
	StepAll *int64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前步骤
	StepNow *int64 `json:"StepNow,omitempty" name:"StepNow"`

	// 总进度,如："30"表示30%
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 步骤信息，null表示还未开始迁移
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo []*MigrationStep `json:"StepInfo,omitempty" name:"StepInfo"`
}

type MigrationStep struct {
	// 步骤序列
	StepNo *int64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤展现名称
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 英文ID标识
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// 步骤状态:0-默认值,1-成功,2-失败,3-执行中,4-未执行
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyAccountPrivilegeRequestParams struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账号权限变更信息
	Accounts []*AccountPrivilegeModifyInfo `json:"Accounts,omitempty" name:"Accounts"`
}

type ModifyAccountPrivilegeRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 账号权限变更信息
	Accounts []*AccountPrivilegeModifyInfo `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ModifyAccountPrivilegeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountPrivilegeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountPrivilegeResponseParams struct {
	// 异步任务流程ID
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountPrivilegeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountPrivilegeResponseParams `json:"Response"`
}

func (r *ModifyAccountPrivilegeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountPrivilegeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountRemarkRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改备注的账户信息
	Accounts []*AccountRemark `json:"Accounts,omitempty" name:"Accounts"`
}

type ModifyAccountRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改备注的账户信息
	Accounts []*AccountRemark `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ModifyAccountRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccountRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccountRemarkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAccountRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccountRemarkResponseParams `json:"Response"`
}

func (r *ModifyAccountRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccountRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 任务名称
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// 迁移任务恢复类型，FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，用户上传
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// UploadType是COS_URL时这里时URL，COS_UPLOAD这里填备份文件的名字；只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// 需要重命名的数据库名称集合
	DBRename []*RenameRestoreDatabase `json:"DBRename,omitempty" name:"DBRename"`
}

type ModifyBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 任务名称
	MigrationName *string `json:"MigrationName,omitempty" name:"MigrationName"`

	// 迁移任务恢复类型，FULL,FULL_LOG,FULL_DIFF
	RecoveryType *string `json:"RecoveryType,omitempty" name:"RecoveryType"`

	// COS_URL-备份放在用户的对象存储上，提供URL。COS_UPLOAD-备份放在业务的对象存储上，用户上传
	UploadType *string `json:"UploadType,omitempty" name:"UploadType"`

	// UploadType是COS_URL时这里时URL，COS_UPLOAD这里填备份文件的名字；只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`

	// 需要重命名的数据库名称集合
	DBRename []*RenameRestoreDatabase `json:"DBRename,omitempty" name:"DBRename"`
}

func (r *ModifyBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "MigrationName")
	delete(f, "RecoveryType")
	delete(f, "UploadType")
	delete(f, "BackupFiles")
	delete(f, "DBRename")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupMigrationResponseParams struct {
	// 备份导入任务ID
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupMigrationResponseParams `json:"Response"`
}

func (r *ModifyBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameRequestParams struct {
	// 实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改的备份名称
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 要修改名称的备份ID，可通过 [DescribeBackups](https://cloud.tencent.com/document/product/238/19943)  接口获取。
	BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备份任务组ID，在单库备份文件模式下，可通过[DescribeBackups](https://cloud.tencent.com/document/product/238/19943) 接口获得。
	//  BackupId 和 GroupId 同时存在，按照BackupId进行修改。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type ModifyBackupNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 修改的备份名称
	BackupName *string `json:"BackupName,omitempty" name:"BackupName"`

	// 要修改名称的备份ID，可通过 [DescribeBackups](https://cloud.tencent.com/document/product/238/19943)  接口获取。
	BackupId *uint64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备份任务组ID，在单库备份文件模式下，可通过[DescribeBackups](https://cloud.tencent.com/document/product/238/19943) 接口获得。
	//  BackupId 和 GroupId 同时存在，按照BackupId进行修改。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyBackupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupName")
	delete(f, "BackupId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupNameResponseParams `json:"Response"`
}

func (r *ModifyBackupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupStrategyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份类型，当length(BackupDay) <=7 && length(BackupDay) >=2时，取值为weekly，当length(BackupDay)=1时，取值daily，默认daily
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份时间点，取值为0-23的整数
	BackupTime *uint64 `json:"BackupTime,omitempty" name:"BackupTime"`

	// BackupType取值为daily时，表示备份间隔天数。当前取值只能为1
	BackupDay *uint64 `json:"BackupDay,omitempty" name:"BackupDay"`

	// 备份模式，master_pkg-主节点上打包备份文件；master_no_pkg-主节点单库备份文件；slave_pkg-从节点上打包备份文件；slave_no_pkg-从节点上单库备份文件，从节点上备份只有在always on容灾模式下支持。
	BackupModel *string `json:"BackupModel,omitempty" name:"BackupModel"`

	// BackupType取值为weekly时，表示每周的星期N做备份。（如果数据备份保留时间<7天，则取值[1,2,3,4,5,6,7]。如果数据备份保留时间>=7天，则备份周期取值至少是一周的任意2天）
	BackupCycle []*uint64 `json:"BackupCycle,omitempty" name:"BackupCycle"`

	// 数据(日志)备份保留时间，取值[3-1830]天，默认7天
	BackupSaveDays *uint64 `json:"BackupSaveDays,omitempty" name:"BackupSaveDays"`

	// 定期备份状态 enable-开启，disable-关闭，默认关闭
	RegularBackupEnable *string `json:"RegularBackupEnable,omitempty" name:"RegularBackupEnable"`

	// 定期备份保留天数 [90 - 3650]天，默认365天
	RegularBackupSaveDays *uint64 `json:"RegularBackupSaveDays,omitempty" name:"RegularBackupSaveDays"`

	// 定期备份策略 years-每年，quarters-每季度，months-每月，默认months
	RegularBackupStrategy *string `json:"RegularBackupStrategy,omitempty" name:"RegularBackupStrategy"`

	// 定期备份保留个数，默认1个
	RegularBackupCounts *uint64 `json:"RegularBackupCounts,omitempty" name:"RegularBackupCounts"`

	// 定期备份开始日期，格式-YYYY-MM-DD 默认当前日期
	RegularBackupStartTime *string `json:"RegularBackupStartTime,omitempty" name:"RegularBackupStartTime"`
}

type ModifyBackupStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份类型，当length(BackupDay) <=7 && length(BackupDay) >=2时，取值为weekly，当length(BackupDay)=1时，取值daily，默认daily
	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`

	// 备份时间点，取值为0-23的整数
	BackupTime *uint64 `json:"BackupTime,omitempty" name:"BackupTime"`

	// BackupType取值为daily时，表示备份间隔天数。当前取值只能为1
	BackupDay *uint64 `json:"BackupDay,omitempty" name:"BackupDay"`

	// 备份模式，master_pkg-主节点上打包备份文件；master_no_pkg-主节点单库备份文件；slave_pkg-从节点上打包备份文件；slave_no_pkg-从节点上单库备份文件，从节点上备份只有在always on容灾模式下支持。
	BackupModel *string `json:"BackupModel,omitempty" name:"BackupModel"`

	// BackupType取值为weekly时，表示每周的星期N做备份。（如果数据备份保留时间<7天，则取值[1,2,3,4,5,6,7]。如果数据备份保留时间>=7天，则备份周期取值至少是一周的任意2天）
	BackupCycle []*uint64 `json:"BackupCycle,omitempty" name:"BackupCycle"`

	// 数据(日志)备份保留时间，取值[3-1830]天，默认7天
	BackupSaveDays *uint64 `json:"BackupSaveDays,omitempty" name:"BackupSaveDays"`

	// 定期备份状态 enable-开启，disable-关闭，默认关闭
	RegularBackupEnable *string `json:"RegularBackupEnable,omitempty" name:"RegularBackupEnable"`

	// 定期备份保留天数 [90 - 3650]天，默认365天
	RegularBackupSaveDays *uint64 `json:"RegularBackupSaveDays,omitempty" name:"RegularBackupSaveDays"`

	// 定期备份策略 years-每年，quarters-每季度，months-每月，默认months
	RegularBackupStrategy *string `json:"RegularBackupStrategy,omitempty" name:"RegularBackupStrategy"`

	// 定期备份保留个数，默认1个
	RegularBackupCounts *uint64 `json:"RegularBackupCounts,omitempty" name:"RegularBackupCounts"`

	// 定期备份开始日期，格式-YYYY-MM-DD 默认当前日期
	RegularBackupStartTime *string `json:"RegularBackupStartTime,omitempty" name:"RegularBackupStartTime"`
}

func (r *ModifyBackupStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupType")
	delete(f, "BackupTime")
	delete(f, "BackupDay")
	delete(f, "BackupModel")
	delete(f, "BackupCycle")
	delete(f, "BackupSaveDays")
	delete(f, "RegularBackupEnable")
	delete(f, "RegularBackupSaveDays")
	delete(f, "RegularBackupStrategy")
	delete(f, "RegularBackupCounts")
	delete(f, "RegularBackupStartTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBackupStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBackupStrategyResponseParams struct {
	// 返回错误码
	Errno *int64 `json:"Errno,omitempty" name:"Errno"`

	// 返回错误信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyBackupStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBackupStrategyResponseParams `json:"Response"`
}

func (r *ModifyBackupStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBackupStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameRequestParams struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新的数据库实例名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ModifyDBInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新的数据库实例名字
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyDBInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新VPC网络Id
	NewVpcId *string `json:"NewVpcId,omitempty" name:"NewVpcId"`

	// 新子网Id
	NewSubnetId *string `json:"NewSubnetId,omitempty" name:"NewSubnetId"`

	// 原vip保留时长，单位小时，默认为0，代表立即回收，最大为168小时
	OldIpRetainTime *int64 `json:"OldIpRetainTime,omitempty" name:"OldIpRetainTime"`

	// 指定VIP地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

type ModifyDBInstanceNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 新VPC网络Id
	NewVpcId *string `json:"NewVpcId,omitempty" name:"NewVpcId"`

	// 新子网Id
	NewSubnetId *string `json:"NewSubnetId,omitempty" name:"NewSubnetId"`

	// 原vip保留时长，单位小时，默认为0，代表立即回收，最大为168小时
	OldIpRetainTime *int64 `json:"OldIpRetainTime,omitempty" name:"OldIpRetainTime"`

	// 指定VIP地址
	Vip *string `json:"Vip,omitempty" name:"Vip"`
}

func (r *ModifyDBInstanceNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "NewVpcId")
	delete(f, "NewSubnetId")
	delete(f, "OldIpRetainTime")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceNetworkResponseParams struct {
	// 实例转网流程id，可通过[DescribeFlowStatus](https://cloud.tencent.com/document/product/238/19967)接口查询流程状态
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceNetworkResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceProjectRequestParams struct {
	// 实例ID数组，形如mssql-j8kv137v
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 项目ID，为0的话表示默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyDBInstanceProjectRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID数组，形如mssql-j8kv137v
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`

	// 项目ID，为0的话表示默认项目
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyDBInstanceProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceProjectResponseParams struct {
	// 修改成功的实例个数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceProjectResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceRenewFlagRequestParams struct {
	// 实例续费状态标记信息
	RenewFlags []*InstanceRenewInfo `json:"RenewFlags,omitempty" name:"RenewFlags"`
}

type ModifyDBInstanceRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例续费状态标记信息
	RenewFlags []*InstanceRenewInfo `json:"RenewFlags,omitempty" name:"RenewFlags"`
}

func (r *ModifyDBInstanceRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RenewFlags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceRenewFlagResponseParams struct {
	// 修改成功的个数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsRequestParams struct {
	// 实例 ID，格式如：mssql-c1nl9rpv 或者 mssqlro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，格式如：mssql-c1nl9rpv 或者 mssqlro-c1nl9rpv，与云数据库控制台页面中显示的实例 ID 相同。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 要修改的安全组 ID 列表，一个或者多个安全组 ID 组成的数组。
	SecurityGroupIdSet []*string `json:"SecurityGroupIdSet,omitempty" name:"SecurityGroupIdSet"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SecurityGroupIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBInstanceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBInstanceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBInstanceSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBNameRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 旧数据库名
	OldDBName *string `json:"OldDBName,omitempty" name:"OldDBName"`

	// 新数据库名
	NewDBName *string `json:"NewDBName,omitempty" name:"NewDBName"`
}

type ModifyDBNameRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 旧数据库名
	OldDBName *string `json:"OldDBName,omitempty" name:"OldDBName"`

	// 新数据库名
	NewDBName *string `json:"NewDBName,omitempty" name:"NewDBName"`
}

func (r *ModifyDBNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "OldDBName")
	delete(f, "NewDBName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBNameResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBNameResponseParams `json:"Response"`
}

func (r *ModifyDBNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBRemarkRequestParams struct {
	// 实例ID，形如mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名称及备注数组，每个元素包含数据库名和对应的备注
	DBRemarks []*DBRemark `json:"DBRemarks,omitempty" name:"DBRemarks"`
}

type ModifyDBRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-rljoi3bf
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 数据库名称及备注数组，每个元素包含数据库名和对应的备注
	DBRemarks []*DBRemark `json:"DBRemarks,omitempty" name:"DBRemarks"`
}

func (r *ModifyDBRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DBRemarks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDBRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDBRemarkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDBRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDBRemarkResponseParams `json:"Response"`
}

func (r *ModifyDBRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDBRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCDCRequestParams struct {
	// 数据库名数组
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 开启、关闭数据库CDC功能 enable；开启，disable：关闭
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ModifyDatabaseCDCRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名数组
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 开启、关闭数据库CDC功能 enable；开启，disable：关闭
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDatabaseCDCRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCDCRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBNames")
	delete(f, "ModifyType")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseCDCRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCDCResponseParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDatabaseCDCResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseCDCResponseParams `json:"Response"`
}

func (r *ModifyDatabaseCDCResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCDCResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCTRequestParams struct {
	// 数据库名数组
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 启用、禁用数据库CT功能 enable；启用，disable：禁用
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 启用CT时额外保留天数，默认保留3天，最小3天，最大30天
	ChangeRetentionDay *int64 `json:"ChangeRetentionDay,omitempty" name:"ChangeRetentionDay"`
}

type ModifyDatabaseCTRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名数组
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 启用、禁用数据库CT功能 enable；启用，disable：禁用
	ModifyType *string `json:"ModifyType,omitempty" name:"ModifyType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 启用CT时额外保留天数，默认保留3天，最小3天，最大30天
	ChangeRetentionDay *int64 `json:"ChangeRetentionDay,omitempty" name:"ChangeRetentionDay"`
}

func (r *ModifyDatabaseCTRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCTRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBNames")
	delete(f, "ModifyType")
	delete(f, "InstanceId")
	delete(f, "ChangeRetentionDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseCTRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseCTResponseParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDatabaseCTResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseCTResponseParams `json:"Response"`
}

func (r *ModifyDatabaseCTResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseCTResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseMdfRequestParams struct {
	// 数据库名数组
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type ModifyDatabaseMdfRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名数组
	DBNames []*string `json:"DBNames,omitempty" name:"DBNames"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDatabaseMdfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseMdfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DBNames")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseMdfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseMdfResponseParams struct {
	// 流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDatabaseMdfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseMdfResponseParams `json:"Response"`
}

func (r *ModifyDatabaseMdfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseMdfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIncrementalMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量导入任务ID，由CreateIncrementalMigration接口返回
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

	// 是否需要恢复，NO-不需要，YES-需要，默认不修改增量备份导入任务是否需要恢复的属性。
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// UploadType是COS_URL时这里时URL，COS_UPLOAD这里填备份文件的名字；只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

type ModifyIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量导入任务ID，由CreateIncrementalMigration接口返回
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

	// 是否需要恢复，NO-不需要，YES-需要，默认不修改增量备份导入任务是否需要恢复的属性。
	IsRecovery *string `json:"IsRecovery,omitempty" name:"IsRecovery"`

	// UploadType是COS_URL时这里时URL，COS_UPLOAD这里填备份文件的名字；只支持1个备份文件，但1个备份文件内可包含多个库
	BackupFiles []*string `json:"BackupFiles,omitempty" name:"BackupFiles"`
}

func (r *ModifyIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	delete(f, "IsRecovery")
	delete(f, "BackupFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIncrementalMigrationResponseParams struct {
	// 增量备份导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIncrementalMigrationResponseParams `json:"Response"`
}

func (r *ModifyIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamRequestParams struct {
	// 实例短 ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改的值。<b>注意</b>：如果修改的参数需要<b>重启</b>实例，那么您的实例将会在执行修改时<b>重启</b>。您可以通过DescribeInstanceParams接口查询修改参数时是否会重启实例，以免导致您的实例不符合预期重启。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行。
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

type ModifyInstanceParamRequest struct {
	*tchttp.BaseRequest
	
	// 实例短 ID 列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`

	// 要修改的参数列表。每一个元素是 Name 和 CurrentValue 的组合。Name 是参数名，CurrentValue 是要修改的值。<b>注意</b>：如果修改的参数需要<b>重启</b>实例，那么您的实例将会在执行修改时<b>重启</b>。您可以通过DescribeInstanceParams接口查询修改参数时是否会重启实例，以免导致您的实例不符合预期重启。
	ParamList []*Parameter `json:"ParamList,omitempty" name:"ParamList"`

	// 执行参数调整任务的方式，默认为 0。支持值包括：0 - 立刻执行，1 - 时间窗执行。
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *ModifyInstanceParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "ParamList")
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceParamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyInstanceParamResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceParamResponseParams `json:"Response"`
}

func (r *ModifyInstanceParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceSpanRequestParams struct {
	// 实例ID，形如mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 以周为单位，表示允许周几维护，例如：[1,2,3,4,5,6,7]表示周一到周日均为可维护日，本参数不填，则不修改此值。
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 每天可维护的开始时间，例如：10:24标识可维护时间窗10点24分开始，本参数不填，则不修改此值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 每天可维护的持续时间，单位是h，例如：1 表示从可维护的开始时间起持续1小时，本参数不填，则不修改此值。
	Span *uint64 `json:"Span,omitempty" name:"Span"`
}

type ModifyMaintenanceSpanRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-k8voqdlz
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 以周为单位，表示允许周几维护，例如：[1,2,3,4,5,6,7]表示周一到周日均为可维护日，本参数不填，则不修改此值。
	Weekly []*int64 `json:"Weekly,omitempty" name:"Weekly"`

	// 每天可维护的开始时间，例如：10:24标识可维护时间窗10点24分开始，本参数不填，则不修改此值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 每天可维护的持续时间，单位是h，例如：1 表示从可维护的开始时间起持续1小时，本参数不填，则不修改此值。
	Span *uint64 `json:"Span,omitempty" name:"Span"`
}

func (r *ModifyMaintenanceSpanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceSpanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Weekly")
	delete(f, "StartTime")
	delete(f, "Span")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMaintenanceSpanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMaintenanceSpanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMaintenanceSpanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMaintenanceSpanResponseParams `json:"Response"`
}

func (r *ModifyMaintenanceSpanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMaintenanceSpanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationRequestParams struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// 新的迁移任务的名称，若不填则不修改
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 新的迁移类型（1:结构迁移 2:数据迁移 3:增量同步），若不填则不修改
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// 迁移源的类型 1:TencentDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式），若不填则不修改
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 迁移源，若不填则不修改
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// 迁移目标，若不填则不修改
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// 迁移DB对象 ，离线迁移（SourceType=4或SourceType=5）不使用，若不填则不修改
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`
}

type ModifyMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// 新的迁移任务的名称，若不填则不修改
	MigrateName *string `json:"MigrateName,omitempty" name:"MigrateName"`

	// 新的迁移类型（1:结构迁移 2:数据迁移 3:增量同步），若不填则不修改
	MigrateType *uint64 `json:"MigrateType,omitempty" name:"MigrateType"`

	// 迁移源的类型 1:TencentDB for SQLServer 2:云服务器自建SQLServer数据库 4:SQLServer备份还原 5:SQLServer备份还原（COS方式），若不填则不修改
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 迁移源，若不填则不修改
	Source *MigrateSource `json:"Source,omitempty" name:"Source"`

	// 迁移目标，若不填则不修改
	Target *MigrateTarget `json:"Target,omitempty" name:"Target"`

	// 迁移DB对象 ，离线迁移（SourceType=4或SourceType=5）不使用，若不填则不修改
	MigrateDBSet []*MigrateDB `json:"MigrateDBSet,omitempty" name:"MigrateDBSet"`
}

func (r *ModifyMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	delete(f, "MigrateName")
	delete(f, "MigrateType")
	delete(f, "SourceType")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "MigrateDBSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationResponseParams struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyMigrationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrationResponseParams `json:"Response"`
}

func (r *ModifyMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublishSubscribeNameRequestParams struct {
	// 发布订阅ID
	PublishSubscribeId *uint64 `json:"PublishSubscribeId,omitempty" name:"PublishSubscribeId"`

	// 待修改的发布订阅名称
	PublishSubscribeName *string `json:"PublishSubscribeName,omitempty" name:"PublishSubscribeName"`
}

type ModifyPublishSubscribeNameRequest struct {
	*tchttp.BaseRequest
	
	// 发布订阅ID
	PublishSubscribeId *uint64 `json:"PublishSubscribeId,omitempty" name:"PublishSubscribeId"`

	// 待修改的发布订阅名称
	PublishSubscribeName *string `json:"PublishSubscribeName,omitempty" name:"PublishSubscribeName"`
}

func (r *ModifyPublishSubscribeNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublishSubscribeNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PublishSubscribeId")
	delete(f, "PublishSubscribeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPublishSubscribeNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublishSubscribeNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPublishSubscribeNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPublishSubscribeNameResponseParams `json:"Response"`
}

func (r *ModifyPublishSubscribeNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublishSubscribeNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyGroupDetailsRequestParams struct {
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称，不填此参数，则不修改
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 是否启动超时剔除功能,0-不开启剔除功能，1-开启剔除功能，不填此参数，则不修改
	IsOfflineDelay *int64 `json:"IsOfflineDelay,omitempty" name:"IsOfflineDelay"`

	// 启动超时剔除功能后，使用的超时阈值，不填此参数，则不修改
	ReadOnlyMaxDelayTime *int64 `json:"ReadOnlyMaxDelayTime,omitempty" name:"ReadOnlyMaxDelayTime"`

	// 启动超时剔除功能后，只读组至少保留的只读副本数，不填此参数，则不修改
	MinReadOnlyInGroup *int64 `json:"MinReadOnlyInGroup,omitempty" name:"MinReadOnlyInGroup"`

	// 只读组实例权重修改集合，不填此参数，则不修改
	WeightPairs []*ReadOnlyInstanceWeightPair `json:"WeightPairs,omitempty" name:"WeightPairs"`

	// 0-用户自定义权重（根据WeightPairs调整）,1-系统自动分配权重(WeightPairs无效)， 默认为0
	AutoWeight *int64 `json:"AutoWeight,omitempty" name:"AutoWeight"`

	// 0-不重新均衡负载，1-重新均衡负载，默认为0
	BalanceWeight *int64 `json:"BalanceWeight,omitempty" name:"BalanceWeight"`
}

type ModifyReadOnlyGroupDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 主实例ID，格式如：mssql-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称，不填此参数，则不修改
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 是否启动超时剔除功能,0-不开启剔除功能，1-开启剔除功能，不填此参数，则不修改
	IsOfflineDelay *int64 `json:"IsOfflineDelay,omitempty" name:"IsOfflineDelay"`

	// 启动超时剔除功能后，使用的超时阈值，不填此参数，则不修改
	ReadOnlyMaxDelayTime *int64 `json:"ReadOnlyMaxDelayTime,omitempty" name:"ReadOnlyMaxDelayTime"`

	// 启动超时剔除功能后，只读组至少保留的只读副本数，不填此参数，则不修改
	MinReadOnlyInGroup *int64 `json:"MinReadOnlyInGroup,omitempty" name:"MinReadOnlyInGroup"`

	// 只读组实例权重修改集合，不填此参数，则不修改
	WeightPairs []*ReadOnlyInstanceWeightPair `json:"WeightPairs,omitempty" name:"WeightPairs"`

	// 0-用户自定义权重（根据WeightPairs调整）,1-系统自动分配权重(WeightPairs无效)， 默认为0
	AutoWeight *int64 `json:"AutoWeight,omitempty" name:"AutoWeight"`

	// 0-不重新均衡负载，1-重新均衡负载，默认为0
	BalanceWeight *int64 `json:"BalanceWeight,omitempty" name:"BalanceWeight"`
}

func (r *ModifyReadOnlyGroupDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadOnlyGroupId")
	delete(f, "ReadOnlyGroupName")
	delete(f, "IsOfflineDelay")
	delete(f, "ReadOnlyMaxDelayTime")
	delete(f, "MinReadOnlyInGroup")
	delete(f, "WeightPairs")
	delete(f, "AutoWeight")
	delete(f, "BalanceWeight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyReadOnlyGroupDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyReadOnlyGroupDetailsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyReadOnlyGroupDetailsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyReadOnlyGroupDetailsResponseParams `json:"Response"`
}

func (r *ModifyReadOnlyGroupDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyReadOnlyGroupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenInterCommunicationRequestParams struct {
	// 打开互通组的实例ID集合
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type OpenInterCommunicationRequest struct {
	*tchttp.BaseRequest
	
	// 打开互通组的实例ID集合
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *OpenInterCommunicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenInterCommunicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenInterCommunicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenInterCommunicationResponseParams struct {
	// 实例和异步流程ID
	InterInstanceFlowSet []*InterInstanceFlow `json:"InterInstanceFlowSet,omitempty" name:"InterInstanceFlowSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type OpenInterCommunicationResponse struct {
	*tchttp.BaseResponse
	Response *OpenInterCommunicationResponseParams `json:"Response"`
}

func (r *OpenInterCommunicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenInterCommunicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamRecord struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 参数名称
	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`

	// 参数修改前的值
	OldValue *string `json:"OldValue,omitempty" name:"OldValue"`

	// 参数修改后的值
	NewValue *string `json:"NewValue,omitempty" name:"NewValue"`

	// 参数修改状态，1-初始化等待被执行，2-执行成功，3-执行失败，4-参数修改中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type Parameter struct {
	// 参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
}

type ParameterDetail struct {
	// 参数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 参数类型，integer-整型，enum-枚举型
	ParamType *string `json:"ParamType,omitempty" name:"ParamType"`

	// 参数默认值
	Default *string `json:"Default,omitempty" name:"Default"`

	// 参数描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 参数当前值
	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`

	// 修改参数后，是否需要重启数据库以使参数生效，0-不需要重启，1-需要重启
	NeedReboot *int64 `json:"NeedReboot,omitempty" name:"NeedReboot"`

	// 参数允许的最大值
	Max *int64 `json:"Max,omitempty" name:"Max"`

	// 参数允许的最小值
	Min *int64 `json:"Min,omitempty" name:"Min"`

	// 参数允许的枚举类型
	EnumValue []*string `json:"EnumValue,omitempty" name:"EnumValue"`

	// 参数状态 0-状态正常 1-在修改中
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type PublishSubscribe struct {
	// 发布订阅ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 发布订阅名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 发布实例ID
	PublishInstanceId *string `json:"PublishInstanceId,omitempty" name:"PublishInstanceId"`

	// 发布实例名称
	PublishInstanceName *string `json:"PublishInstanceName,omitempty" name:"PublishInstanceName"`

	// 发布实例IP
	PublishInstanceIp *string `json:"PublishInstanceIp,omitempty" name:"PublishInstanceIp"`

	// 订阅实例ID
	SubscribeInstanceId *string `json:"SubscribeInstanceId,omitempty" name:"SubscribeInstanceId"`

	// 订阅实例名称
	SubscribeInstanceName *string `json:"SubscribeInstanceName,omitempty" name:"SubscribeInstanceName"`

	// 订阅实例IP
	SubscribeInstanceIp *string `json:"SubscribeInstanceIp,omitempty" name:"SubscribeInstanceIp"`

	// 数据库的订阅发布关系集合
	DatabaseTupleSet []*DatabaseTupleStatus `json:"DatabaseTupleSet,omitempty" name:"DatabaseTupleSet"`
}

// Predefined struct for user
type QueryMigrationCheckProcessRequestParams struct {
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type QueryMigrationCheckProcessRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *QueryMigrationCheckProcessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMigrationCheckProcessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryMigrationCheckProcessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryMigrationCheckProcessResponseParams struct {
	// 总步骤数量
	TotalStep *int64 `json:"TotalStep,omitempty" name:"TotalStep"`

	// 当前步骤编号，从1开始
	CurrentStep *int64 `json:"CurrentStep,omitempty" name:"CurrentStep"`

	// 所有步骤详情
	StepDetails []*StepDetail `json:"StepDetails,omitempty" name:"StepDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryMigrationCheckProcessResponse struct {
	*tchttp.BaseResponse
	Response *QueryMigrationCheckProcessResponseParams `json:"Response"`
}

func (r *QueryMigrationCheckProcessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryMigrationCheckProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReadOnlyGroup struct {
	// 只读组ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`

	// 只读组名称
	ReadOnlyGroupName *string `json:"ReadOnlyGroupName,omitempty" name:"ReadOnlyGroupName"`

	// 只读组的地域ID，与主实例相同
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 只读组的可用区ID，与主实例相同
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 是否启动超时剔除功能，0-不开启剔除功能，1-开启剔除功能
	IsOfflineDelay *int64 `json:"IsOfflineDelay,omitempty" name:"IsOfflineDelay"`

	// 启动超时剔除功能后，使用的超时阈值
	ReadOnlyMaxDelayTime *int64 `json:"ReadOnlyMaxDelayTime,omitempty" name:"ReadOnlyMaxDelayTime"`

	// 启动超时剔除功能后，只读组至少保留的只读副本数
	MinReadOnlyInGroup *int64 `json:"MinReadOnlyInGroup,omitempty" name:"MinReadOnlyInGroup"`

	// 只读组vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 只读组vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 只读组私有网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 只读组私有网络子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 只读组状态: 1-申请成功运行中，5-申请中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 主实例ID，形如mssql-sgeshe3th
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" name:"MasterInstanceId"`

	// 只读实例副本集合
	ReadOnlyInstanceSet []*ReadOnlyInstance `json:"ReadOnlyInstanceSet,omitempty" name:"ReadOnlyInstanceSet"`
}

type ReadOnlyInstance struct {
	// 只读副本ID，格式如：mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读副本名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 只读副本唯一UID
	Uid *string `json:"Uid,omitempty" name:"Uid"`

	// 只读副本所在项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 只读副本状态。1：申请中 2：运行中 3：被延迟剔除 4：已隔离 5：回收中 6：已回收 7：任务执行中 8：已下线 9：实例扩容中 10：实例迁移中  12：重启中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 只读副本创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 只读副本更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 只读副本内存大小，单位G
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 只读副本存储空间大小，单位G
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 只读副本cpu核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 只读副本版本代号
	Version *string `json:"Version,omitempty" name:"Version"`

	// 宿主机代号
	Type *string `json:"Type,omitempty" name:"Type"`

	// 只读副本模式，2-单机
	Model *int64 `json:"Model,omitempty" name:"Model"`

	// 只读副本计费模式，1-包年包月，0-按量计费
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 只读副本权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 只读副本延迟时间，单位秒
	DelayTime *string `json:"DelayTime,omitempty" name:"DelayTime"`

	// 只读副本与主实例的同步状态。
	// Init:初始化
	// DeployReadOnlyInPorgress:部署副本进行中
	// DeployReadOnlySuccess:部署副本成功
	// DeployReadOnlyFail:部署副本失败
	// DeployMasterDBInPorgress:主节点上加入副本数据库进行中
	// DeployMasterDBSuccess:主节点上加入副本数据库成功
	// DeployMasterDBFail:主节点上加入副本数据库进失败
	// DeployReadOnlyDBInPorgress:副本还原加入数据库开始
	// DeployReadOnlyDBSuccess:副本还原加入数据库成功
	// DeployReadOnlyDBFail:副本还原加入数据库失败
	// SyncDelay:同步延迟
	// SyncFail:同步故障
	// SyncExcluded:已剔除只读组
	// SyncNormal:正常
	SynStatus *string `json:"SynStatus,omitempty" name:"SynStatus"`

	// 只读副本与主实例没有同步的库
	DatabaseDifference *string `json:"DatabaseDifference,omitempty" name:"DatabaseDifference"`

	// 只读副本与主实例没有同步的账户
	AccountDifference *string `json:"AccountDifference,omitempty" name:"AccountDifference"`

	// 只读副本计费开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 只读副本计费结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 只读副本隔离时间
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 只读副本所在地域
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 只读副本所在可用区
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
}

type ReadOnlyInstanceWeightPair struct {
	// 只读实例ID，格式如：mssqlro-3l3fgqn7
	ReadOnlyInstanceId *string `json:"ReadOnlyInstanceId,omitempty" name:"ReadOnlyInstanceId"`

	// 只读实例权重 ，范围是0-100
	ReadOnlyWeight *int64 `json:"ReadOnlyWeight,omitempty" name:"ReadOnlyWeight"`
}

// Predefined struct for user
type RecycleDBInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type RecycleDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RecycleDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecycleDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecycleDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecycleDBInstanceResponseParams struct {
	// 流程id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecycleDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RecycleDBInstanceResponseParams `json:"Response"`
}

func (r *RecycleDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecycleDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecycleReadOnlyGroupRequestParams struct {
	// 主实例的ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读组的ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

type RecycleReadOnlyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 主实例的ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 只读组的ID
	ReadOnlyGroupId *string `json:"ReadOnlyGroupId,omitempty" name:"ReadOnlyGroupId"`
}

func (r *RecycleReadOnlyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecycleReadOnlyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ReadOnlyGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecycleReadOnlyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecycleReadOnlyGroupResponseParams struct {
	// 任务流ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecycleReadOnlyGroupResponse struct {
	*tchttp.BaseResponse
	Response *RecycleReadOnlyGroupResponseParams `json:"Response"`
}

func (r *RecycleReadOnlyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecycleReadOnlyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域英文ID，类似ap-guanghou
	Region *string `json:"Region,omitempty" name:"Region"`

	// 地域中文名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域数字ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 该地域目前是否可以售卖，UNAVAILABLE-不可售卖；AVAILABLE-可售卖
	RegionState *string `json:"RegionState,omitempty" name:"RegionState"`
}

// Predefined struct for user
type RemoveBackupsRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待删除的备份名称，备份名称可通过DescribeBackups接口的FileName字段获得。单次请求批量删除备份数不能超过10个。
	BackupNames []*string `json:"BackupNames,omitempty" name:"BackupNames"`

	// 批量删除手动备份起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 批量删除手动备份截止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type RemoveBackupsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 待删除的备份名称，备份名称可通过DescribeBackups接口的FileName字段获得。单次请求批量删除备份数不能超过10个。
	BackupNames []*string `json:"BackupNames,omitempty" name:"BackupNames"`

	// 批量删除手动备份起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 批量删除手动备份截止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *RemoveBackupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveBackupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupNames")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveBackupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveBackupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemoveBackupsResponse struct {
	*tchttp.BaseResponse
	Response *RemoveBackupsResponseParams `json:"Response"`
}

func (r *RemoveBackupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveBackupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenameRestoreDatabase struct {
	// 库的名字，如果oldName不存在则返回失败。
	// 在用于离线迁移任务时可不填。
	OldName *string `json:"OldName,omitempty" name:"OldName"`

	// 库的新名字，在用于离线迁移时，不填则按照OldName命名，OldName和NewName不能同时不填。在用于克隆数据库时，OldName和NewName都必须填写，且不能重复
	NewName *string `json:"NewName,omitempty" name:"NewName"`
}

// Predefined struct for user
type RenewDBInstanceRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费多少个月，取值范围为1-48，默认为1
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券，0-不使用；1-使用；默认不实用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前只支持使用1张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 续费标记 0:正常续费 1:自动续费：只用于按量计费转包年包月时有效。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

type RenewDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 续费多少个月，取值范围为1-48，默认为1
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 是否自动使用代金券，0-不使用；1-使用；默认不实用
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID数组，目前只支持使用1张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 续费标记 0:正常续费 1:自动续费：只用于按量计费转包年包月时有效。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
}

func (r *RenewDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "AutoRenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDBInstanceResponseParams struct {
	// 订单名称
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewDBInstanceResponseParams `json:"Response"`
}

func (r *RenewDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewPostpaidDBInstanceRequestParams struct {
	// 实例ID，格式如：mssql-3l3fgqn7 或 mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type RenewPostpaidDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，格式如：mssql-3l3fgqn7 或 mssqlro-3l3fgqn7
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RenewPostpaidDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewPostpaidDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewPostpaidDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewPostpaidDBInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewPostpaidDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RenewPostpaidDBInstanceResponseParams `json:"Response"`
}

func (r *RenewPostpaidDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewPostpaidDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordRequestParams struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 更新后的账户密码信息数组
	Accounts []*AccountPassword `json:"Accounts,omitempty" name:"Accounts"`
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 更新后的账户密码信息数组
	Accounts []*AccountPassword `json:"Accounts,omitempty" name:"Accounts"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Accounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAccountPasswordResponseParams struct {
	// 修改帐号密码的异步任务流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceTag struct {
	// 标签key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

// Predefined struct for user
type RestartDBInstanceRequestParams struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type RestartDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *RestartDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDBInstanceResponseParams struct {
	// 异步任务流程ID
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartDBInstanceResponseParams `json:"Response"`
}

func (r *RestartDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreInstanceRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件ID，该ID可以通过DescribeBackups接口返回数据中的Id字段获得
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备份恢复到的同一个APPID下的实例ID，不填则恢复到原实例ID
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`

	// 按照ReNameRestoreDatabase中的库进行恢复，并重命名，不填则按照默认方式命名恢复的库，且恢复所有的库。
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`

	// 备份任务组ID，在单库备份文件模式下，可通过[DescribeBackups](https://cloud.tencent.com/document/product/238/19943) 接口获得。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

type RestoreInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份文件ID，该ID可以通过DescribeBackups接口返回数据中的Id字段获得
	BackupId *int64 `json:"BackupId,omitempty" name:"BackupId"`

	// 备份恢复到的同一个APPID下的实例ID，不填则恢复到原实例ID
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`

	// 按照ReNameRestoreDatabase中的库进行恢复，并重命名，不填则按照默认方式命名恢复的库，且恢复所有的库。
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`

	// 备份任务组ID，在单库备份文件模式下，可通过[DescribeBackups](https://cloud.tencent.com/document/product/238/19943) 接口获得。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *RestoreInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupId")
	delete(f, "TargetInstanceId")
	delete(f, "RenameRestore")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestoreInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestoreInstanceResponseParams struct {
	// 异步流程任务ID，使用FlowId调用DescribeFlowStatus接口获取任务执行状态
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestoreInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestoreInstanceResponseParams `json:"Response"`
}

func (r *RestoreInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestoreInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 回档类型，0-回档的数据库覆盖原库；1-回档的数据库以重命名的形式生成，不覆盖原库
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 需要回档的数据库
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// 回档目标时间点
	Time *string `json:"Time,omitempty" name:"Time"`

	// 备份恢复到的同一个APPID下的实例ID，不填则恢复到原实例ID
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`

	// 按照ReNameRestoreDatabase中的库进行重命名，仅在Type = 1重命名回档方式有效；不填则按照默认方式命名库，DBs参数确定要恢复的库
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

type RollbackInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 回档类型，0-回档的数据库覆盖原库；1-回档的数据库以重命名的形式生成，不覆盖原库
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 需要回档的数据库
	DBs []*string `json:"DBs,omitempty" name:"DBs"`

	// 回档目标时间点
	Time *string `json:"Time,omitempty" name:"Time"`

	// 备份恢复到的同一个APPID下的实例ID，不填则恢复到原实例ID
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" name:"TargetInstanceId"`

	// 按照ReNameRestoreDatabase中的库进行重命名，仅在Type = 1重命名回档方式有效；不填则按照默认方式命名库，DBs参数确定要恢复的库
	RenameRestore []*RenameRestoreDatabase `json:"RenameRestore,omitempty" name:"RenameRestore"`
}

func (r *RollbackInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "DBs")
	delete(f, "Time")
	delete(f, "TargetInstanceId")
	delete(f, "RenameRestore")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceResponseParams struct {
	// 异步任务ID
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollbackInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RollbackInstanceResponseParams `json:"Response"`
}

func (r *RollbackInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunMigrationRequestParams struct {
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type RunMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *uint64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *RunMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunMigrationResponseParams struct {
	// 迁移流程启动后，返回流程ID
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunMigrationResponse struct {
	*tchttp.BaseResponse
	Response *RunMigrationResponseParams `json:"Response"`
}

func (r *RunMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroup struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 入站规则
	InboundSet []*SecurityGroupPolicy `json:"InboundSet,omitempty" name:"InboundSet"`

	// 出站规则
	OutboundSet []*SecurityGroupPolicy `json:"OutboundSet,omitempty" name:"OutboundSet"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`

	// 安全组备注
	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type SecurityGroupPolicy struct {
	// 策略，ACCEPT 或者 DROP
	Action *string `json:"Action,omitempty" name:"Action"`

	// 目的 IP 或 IP 段，例如172.16.0.0/12
	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`

	// 端口或者端口范围
	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`

	// 网络协议，支持 UDP、TCP等
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 规则限定的方向，OUTPUT-出战规则  INPUT-进站规则
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type SlaveZones struct {
	// 备可用区地域码
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 备可用区
	SlaveZoneName *string `json:"SlaveZoneName,omitempty" name:"SlaveZoneName"`
}

type SlowlogInfo struct {
	// 慢查询日志文件唯一标识
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 文件生成的开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 文件生成的结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 文件大小（KB）
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 文件中log条数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 内网下载地址
	InternalAddr *string `json:"InternalAddr,omitempty" name:"InternalAddr"`

	// 外网下载地址
	ExternalAddr *string `json:"ExternalAddr,omitempty" name:"ExternalAddr"`

	// 状态（1成功 2失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type SpecInfo struct {
	// 实例规格ID，利用DescribeZones返回的SpecId，结合DescribeProductConfig返回的可售卖规格信息，可获悉某个可用区下可购买什么规格的实例
	SpecId *int64 `json:"SpecId,omitempty" name:"SpecId"`

	// 机型ID
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机型中文名称
	MachineTypeName *string `json:"MachineTypeName,omitempty" name:"MachineTypeName"`

	// 数据库版本信息。取值为2008R2（表示SQL Server 2008 R2），2012SP3（表示SQL Server 2012），2016SP1（表示SQL Server 2016 SP1）
	Version *string `json:"Version,omitempty" name:"Version"`

	// Version字段对应的版本名称
	VersionName *string `json:"VersionName,omitempty" name:"VersionName"`

	// 内存大小，单位GB
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// CPU核数
	CPU *int64 `json:"CPU,omitempty" name:"CPU"`

	// 此规格下最小的磁盘大小，单位GB
	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`

	// 此规格下最大的磁盘大小，单位GB
	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`

	// 此规格对应的QPS大小
	QPS *int64 `json:"QPS,omitempty" name:"QPS"`

	// 此规格的中文描述信息
	SuitInfo *string `json:"SuitInfo,omitempty" name:"SuitInfo"`

	// 此规格对应的包年包月Pid
	Pid *int64 `json:"Pid,omitempty" name:"Pid"`

	// 此规格对应的按量计费Pid列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostPid []*int64 `json:"PostPid,omitempty" name:"PostPid"`

	// 此规格下支持的付费模式，POST-仅支持按量计费 PRE-仅支持包年包月 ALL-支持所有
	PayModeStatus *string `json:"PayModeStatus,omitempty" name:"PayModeStatus"`

	// 产品类型，HA-高可用型(包括双机高可用，alwaysOn集群)，RO-只读副本型，SI-基础版本型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 跨可用区类型，MultiZones-只支持跨可用区，SameZones-只支持同可用区，ALL-支持所有
	MultiZonesStatus *string `json:"MultiZonesStatus,omitempty" name:"MultiZonesStatus"`
}

// Predefined struct for user
type StartBackupMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

type StartBackupMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`
}

func (r *StartBackupMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartBackupMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartBackupMigrationResponseParams struct {
	// 流程ID
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartBackupMigrationResponse struct {
	*tchttp.BaseResponse
	Response *StartBackupMigrationResponseParams `json:"Response"`
}

func (r *StartBackupMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartBackupMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartIncrementalMigrationRequestParams struct {
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量备份导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

type StartIncrementalMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 导入目标实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 备份导入任务ID，由CreateBackupMigration接口返回
	BackupMigrationId *string `json:"BackupMigrationId,omitempty" name:"BackupMigrationId"`

	// 增量备份导入任务ID
	IncrementalMigrationId *string `json:"IncrementalMigrationId,omitempty" name:"IncrementalMigrationId"`
}

func (r *StartIncrementalMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartIncrementalMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "BackupMigrationId")
	delete(f, "IncrementalMigrationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartIncrementalMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartIncrementalMigrationResponseParams struct {
	// 流程ID
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartIncrementalMigrationResponse struct {
	*tchttp.BaseResponse
	Response *StartIncrementalMigrationResponseParams `json:"Response"`
}

func (r *StartIncrementalMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartIncrementalMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrationCheckRequestParams struct {
	// 迁移任务id
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type StartMigrationCheckRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务id
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *StartMigrationCheckRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrationCheckRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMigrationCheckRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrationCheckResponseParams struct {
	// 迁移检查流程发起后，返回的流程id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartMigrationCheckResponse struct {
	*tchttp.BaseResponse
	Response *StartMigrationCheckResponseParams `json:"Response"`
}

func (r *StartMigrationCheckResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrationCheckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StepDetail struct {
	// 具体步骤返回信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 当前步骤状态，0成功，-2未开始
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 步骤名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

// Predefined struct for user
type StopMigrationRequestParams struct {
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

type StopMigrationRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	MigrateId *int64 `json:"MigrateId,omitempty" name:"MigrateId"`
}

func (r *StopMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MigrateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrationResponseParams struct {
	// 中止迁移流程发起后，返回的流程id
	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopMigrationResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrationResponseParams `json:"Response"`
}

func (r *StopMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDBInstanceRequestParams struct {
	// 主动销毁的实例ID列表，格式如：[mssql-3l3fgqn7]。与云数据库控制台页面中显示的实例ID相同
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

type TerminateDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 主动销毁的实例ID列表，格式如：[mssql-3l3fgqn7]。与云数据库控制台页面中显示的实例ID相同
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
}

func (r *TerminateDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateDBInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TerminateDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateDBInstanceResponseParams `json:"Response"`
}

func (r *TerminateDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceRequestParams struct {
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例升级后内存大小，单位GB，其值不能小于当前实例内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例升级后磁盘大小，单位GB，其值不能小于当前实例磁盘大小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 是否自动使用代金券，0 - 不使用；1 - 默认使用。取值默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID，目前单个订单只能使用一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 实例升级后的CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 升级sqlserver的版本，目前支持：2008R2（SQL Server 2008 Enterprise），2012SP3（SQL Server 2012 Enterprise）版本等。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息，版本不支持降级，不填则不修改版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 升级sqlserver的高可用架构,从镜像容灾升级到always on集群容灾，仅支持2017及以上版本且支持always on高可用的实例，不支持降级到镜像方式容灾，CLUSTER-升级为always on容灾，不填则不修改高可用架构
	HAType *string `json:"HAType,omitempty" name:"HAType"`

	// 修改实例是否为跨可用区容灾，SameZones-修改为同可用区 MultiZones-修改为夸可用区
	MultiZones *string `json:"MultiZones,omitempty" name:"MultiZones"`

	// 执行变配的方式，默认为 1。支持值包括：0 - 立刻执行，1 - 维护时间窗执行
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

type UpgradeDBInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID，形如mssql-j8kv137v
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例升级后内存大小，单位GB，其值不能小于当前实例内存大小
	Memory *int64 `json:"Memory,omitempty" name:"Memory"`

	// 实例升级后磁盘大小，单位GB，其值不能小于当前实例磁盘大小
	Storage *int64 `json:"Storage,omitempty" name:"Storage"`

	// 是否自动使用代金券，0 - 不使用；1 - 默认使用。取值默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID，目前单个订单只能使用一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 实例升级后的CPU核心数
	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`

	// 升级sqlserver的版本，目前支持：2008R2（SQL Server 2008 Enterprise），2012SP3（SQL Server 2012 Enterprise）版本等。每个地域支持售卖的版本不同，可通过DescribeProductConfig接口来拉取每个地域可售卖的版本信息，版本不支持降级，不填则不修改版本
	DBVersion *string `json:"DBVersion,omitempty" name:"DBVersion"`

	// 升级sqlserver的高可用架构,从镜像容灾升级到always on集群容灾，仅支持2017及以上版本且支持always on高可用的实例，不支持降级到镜像方式容灾，CLUSTER-升级为always on容灾，不填则不修改高可用架构
	HAType *string `json:"HAType,omitempty" name:"HAType"`

	// 修改实例是否为跨可用区容灾，SameZones-修改为同可用区 MultiZones-修改为夸可用区
	MultiZones *string `json:"MultiZones,omitempty" name:"MultiZones"`

	// 执行变配的方式，默认为 1。支持值包括：0 - 立刻执行，1 - 维护时间窗执行
	WaitSwitch *int64 `json:"WaitSwitch,omitempty" name:"WaitSwitch"`
}

func (r *UpgradeDBInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Memory")
	delete(f, "Storage")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "Cpu")
	delete(f, "DBVersion")
	delete(f, "HAType")
	delete(f, "MultiZones")
	delete(f, "WaitSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDBInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDBInstanceResponseParams struct {
	// 订单名称
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpgradeDBInstanceResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDBInstanceResponseParams `json:"Response"`
}

func (r *UpgradeDBInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDBInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {
	// 可用区英文ID，形如ap-guangzhou-1，表示广州一区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区中文名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区数字ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 该可用区目前可售卖的规格ID，利用SpecId，结合接口DescribeProductConfig返回的数据，可获悉该可用区目前可售卖的规格大小
	SpecId *int64 `json:"SpecId,omitempty" name:"SpecId"`

	// 当前可用区与规格下，可售卖的数据库版本，形如2008R2（表示SQL Server 2008 R2）。其可选值有2008R2（表示SQL Server 2008 R2），2012SP3（表示SQL Server 2012），2016SP1（表示SQL Server 2016 SP1）
	Version *string `json:"Version,omitempty" name:"Version"`
}