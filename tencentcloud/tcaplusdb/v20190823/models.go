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

package v20190823

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Application struct {
	// 审批单号
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 申请类型
	ApplicationType *int64 `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 表格组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 申请人
	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`

	// 建单时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 处理状态 -1 撤回 0-待审核 1-已经审核并提交任务 2-已驳回
	ApplicationStatus *int64 `json:"ApplicationStatus,omitempty" name:"ApplicationStatus"`

	// 表格组Id
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 已提交的任务Id，未提交申请为0
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 腾讯云上table的唯一键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 审批人
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteUser *string `json:"ExecuteUser,omitempty" name:"ExecuteUser"`

	// 执行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteStatus *string `json:"ExecuteStatus,omitempty" name:"ExecuteStatus"`

	// 该申请单是否可以被当前用户审批
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanCensor *bool `json:"CanCensor,omitempty" name:"CanCensor"`

	// 该申请单是否可以被当前用户撤回
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanWithdrawal *bool `json:"CanWithdrawal,omitempty" name:"CanWithdrawal"`
}

type ApplyResult struct {
	// 申请单id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 申请类型
	ApplicationType *int64 `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 处理状态 0-待审核 1-已经审核并提交任务 2-已驳回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationStatus *int64 `json:"ApplicationStatus,omitempty" name:"ApplicationStatus"`

	// 已提交的任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type ApplyStatus struct {
	// 集群id-申请单id
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 处理状态-1-撤回 1-通过 2-驳回，非0状态的申请单不可改变状态。
	ApplicationStatus *int64 `json:"ApplicationStatus,omitempty" name:"ApplicationStatus"`

	// 申请单类型
	ApplicationType *int64 `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

// Predefined struct for user
type ClearTablesRequestParams struct {
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待清理表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type ClearTablesRequest struct {
	*tchttp.BaseRequest
	
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待清理表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *ClearTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearTablesResponseParams struct {
	// 清除表结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 清除表结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ClearTablesResponse struct {
	*tchttp.BaseResponse
	Response *ClearTablesResponseParams `json:"Response"`
}

func (r *ClearTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {
	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群数据描述语言类型，如：`PROTO`,`TDR`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 集群关联的用户私有网络实例ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群关联的用户子网实例ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 集群密码
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密码状态
	PasswordStatus *string `json:"PasswordStatus,omitempty" name:"PasswordStatus"`

	// TcaplusDB SDK连接参数，接入ID
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// TcaplusDB SDK连接参数，接入地址
	ApiAccessIp *string `json:"ApiAccessIp,omitempty" name:"ApiAccessIp"`

	// TcaplusDB SDK连接参数，接入端口
	ApiAccessPort *int64 `json:"ApiAccessPort,omitempty" name:"ApiAccessPort"`

	// 如果PasswordStatus是unmodifiable说明有旧密码还未过期，此字段将显示旧密码过期的时间，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// TcaplusDB SDK连接参数，接入ipv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAccessIpv6 *string `json:"ApiAccessIpv6,omitempty" name:"ApiAccessIpv6"`

	// 集群类型，0,1:共享集群; 2:独立集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群状态, 0：表示正常运行中，1：表示冻结隔离一般欠费进入此状态，2：表示待回收，一般用户主动触发删除进入这个状态，3：待释放，进入这个状态，表示可以释放此表占用的资源了，4：变更中
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 读CU
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadCapacityUnit *int64 `json:"ReadCapacityUnit,omitempty" name:"ReadCapacityUnit"`

	// 写CU
	// 注意：此字段可能返回 null，表示取不到有效值。
	WriteCapacityUnit *int64 `json:"WriteCapacityUnit,omitempty" name:"WriteCapacityUnit"`

	// 磁盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskVolume *int64 `json:"DiskVolume,omitempty" name:"DiskVolume"`

	// 独占server机器信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerList []*ServerDetailInfo `json:"ServerList,omitempty" name:"ServerList"`

	// 独占proxy机器信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyList []*ProxyDetailInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// 是否开启审核 0-不开启 1-开启
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// 审批人uin列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbaUins []*string `json:"DbaUins,omitempty" name:"DbaUins"`

	// 是否开启了数据订阅
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataFlowStatus *int64 `json:"DataFlowStatus,omitempty" name:"DataFlowStatus"`

	// 数据订阅的kafka信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaInfo *KafkaInfo `json:"KafkaInfo,omitempty" name:"KafkaInfo"`

	// 集群Txh备份文件多少天后过期删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxhBackupExpireDay *uint64 `json:"TxhBackupExpireDay,omitempty" name:"TxhBackupExpireDay"`

	// 集群Ulog备份文件多少天后过期删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	UlogBackupExpireDay *uint64 `json:"UlogBackupExpireDay,omitempty" name:"UlogBackupExpireDay"`

	// 集群Ulog备份文件过期策略是否为只读， 0： UlogBackupExpire是只读，不可修改， 1： UlogBackupExpire可以修改（当前业务存在Svrid第二段等于clusterid的机器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsReadOnlyUlogBackupExpireDay *uint64 `json:"IsReadOnlyUlogBackupExpireDay,omitempty" name:"IsReadOnlyUlogBackupExpireDay"`

	// restproxy状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestProxyStatus *int64 `json:"RestProxyStatus,omitempty" name:"RestProxyStatus"`
}

// Predefined struct for user
type CompareIdlFilesRequestParams struct {
	// 待修改表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 选中的已上传IDL文件列表，与NewIdlFiles必选其一
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// 本次上传IDL文件列表，与ExistingIdlFiles必选其一
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

type CompareIdlFilesRequest struct {
	*tchttp.BaseRequest
	
	// 待修改表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 选中的已上传IDL文件列表，与NewIdlFiles必选其一
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// 本次上传IDL文件列表，与ExistingIdlFiles必选其一
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

func (r *CompareIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareIdlFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "ExistingIdlFiles")
	delete(f, "NewIdlFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompareIdlFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompareIdlFilesResponseParams struct {
	// 本次上传校验所有的IDL文件信息列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 本次校验合法的表格数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 读取IDL描述文件后,根据用户指示的所选中表格解析校验结果
	TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CompareIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *CompareIdlFilesResponseParams `json:"Response"`
}

func (r *CompareIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompareIdlFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompareTablesInfo struct {
	// 源表格的集群id
	SrcTableClusterId *string `json:"SrcTableClusterId,omitempty" name:"SrcTableClusterId"`

	// 源表格的表格组id
	SrcTableGroupId *string `json:"SrcTableGroupId,omitempty" name:"SrcTableGroupId"`

	// 源表格的表名
	SrcTableName *string `json:"SrcTableName,omitempty" name:"SrcTableName"`

	// 目标表格的集群id
	DstTableClusterId *string `json:"DstTableClusterId,omitempty" name:"DstTableClusterId"`

	// 目标表格的表格组id
	DstTableGroupId *string `json:"DstTableGroupId,omitempty" name:"DstTableGroupId"`

	// 目标表格的表名
	DstTableName *string `json:"DstTableName,omitempty" name:"DstTableName"`

	// 源表格的实例id
	SrcTableInstanceId *string `json:"SrcTableInstanceId,omitempty" name:"SrcTableInstanceId"`

	// 目标表格的实例id
	DstTableInstanceId *string `json:"DstTableInstanceId,omitempty" name:"DstTableInstanceId"`
}

// Predefined struct for user
type CreateBackupRequestParams struct {
	// 待创建备份表所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建备份表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest
	
	// 待创建备份表所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建备份表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
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
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBackupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBackupResponseParams struct {
	// 创建的备份任务ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 创建的备份申请ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds"`

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
type CreateClusterRequestParams struct {
	// 集群数据描述语言类型，如：`PROTO`，`TDR`或`MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 集群名称，可使用中文或英文字符，最大长度32个字符
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群所绑定的私有网络实例ID，形如：vpc-f49l6u0z
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群所绑定的子网实例ID，形如：subnet-pxir56ns
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群访问密码，必须是a-zA-Z0-9的字符,且必须包含数字和大小写字母
	Password *string `json:"Password,omitempty" name:"Password"`

	// 集群标签列表
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 集群是否开启IPv6功能
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`

	// 独占集群占用的svr机器
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// 独占集群占用的proxy机器
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// 集群类型1共享2独占
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 密码认证类型，0 静态认证， 1 签名认证
	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群数据描述语言类型，如：`PROTO`，`TDR`或`MIX`
	IdlType *string `json:"IdlType,omitempty" name:"IdlType"`

	// 集群名称，可使用中文或英文字符，最大长度32个字符
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群所绑定的私有网络实例ID，形如：vpc-f49l6u0z
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集群所绑定的子网实例ID，形如：subnet-pxir56ns
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 集群访问密码，必须是a-zA-Z0-9的字符,且必须包含数字和大小写字母
	Password *string `json:"Password,omitempty" name:"Password"`

	// 集群标签列表
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`

	// 集群是否开启IPv6功能
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`

	// 独占集群占用的svr机器
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// 独占集群占用的proxy机器
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// 集群类型1共享2独占
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`

	// 密码认证类型，0 静态认证， 1 签名认证
	AuthType *int64 `json:"AuthType,omitempty" name:"AuthType"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdlType")
	delete(f, "ClusterName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Password")
	delete(f, "ResourceTags")
	delete(f, "Ipv6Enable")
	delete(f, "ServerList")
	delete(f, "ProxyList")
	delete(f, "ClusterType")
	delete(f, "AuthType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotsRequestParams struct {
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 快照列表
	SelectedTables []*SnapshotInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type CreateSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 快照列表
	SelectedTables []*SnapshotInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *CreateSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotsResponseParams struct {
	// 批量创建的快照数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 批量创建的快照结果列表
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotsResponseParams `json:"Response"`
}

func (r *CreateSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableGroupRequestParams struct {
	// 表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组名称，可以采用中文、英文或数字字符，最大长度32个字符
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格组ID，可以由用户指定，但在同一个集群内不能重复，如果不指定则采用自增的模式
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格组标签列表
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type CreateTableGroupRequest struct {
	*tchttp.BaseRequest
	
	// 表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组名称，可以采用中文、英文或数字字符，最大长度32个字符
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格组ID，可以由用户指定，但在同一个集群内不能重复，如果不指定则采用自增的模式
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格组标签列表
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

func (r *CreateTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupName")
	delete(f, "TableGroupId")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableGroupResponseParams struct {
	// 创建成功的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableGroupResponseParams `json:"Response"`
}

func (r *CreateTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTablesRequestParams struct {
	// 待创建表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用户选定的建表格IDL文件列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 待创建表格信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 表格标签列表
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

type CreateTablesRequest struct {
	*tchttp.BaseRequest
	
	// 待创建表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用户选定的建表格IDL文件列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 待创建表格信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 表格标签列表
	ResourceTags []*TagInfoUnit `json:"ResourceTags,omitempty" name:"ResourceTags"`
}

func (r *CreateTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IdlFiles")
	delete(f, "SelectedTables")
	delete(f, "ResourceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTablesResponseParams struct {
	// 批量创建表格结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 批量创建表格结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTablesResponse struct {
	*tchttp.BaseResponse
	Response *CreateTablesResponseParams `json:"Response"`
}

func (r *CreateTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// 待删除的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// 删除集群生成的任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIdlFilesRequestParams struct {
	// IDL所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除的IDL文件信息列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`
}

type DeleteIdlFilesRequest struct {
	*tchttp.BaseRequest
	
	// IDL所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除的IDL文件信息列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`
}

func (r *DeleteIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIdlFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IdlFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIdlFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIdlFilesResponseParams struct {
	// 结果记录数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 删除结果
	IdlFileInfos []*IdlFileInfoWithoutContent `json:"IdlFileInfos,omitempty" name:"IdlFileInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIdlFilesResponseParams `json:"Response"`
}

func (r *DeleteIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIdlFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsRequestParams struct {
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 删除的快照列表
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 删除的快照列表
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotsResponseParams struct {
	// 批量删除的快照数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 批量删除的快照结果
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotsResponseParams `json:"Response"`
}

func (r *DeleteSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableDataFlowRequestParams struct {
	// 表格所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除分布式索引的表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteTableDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除分布式索引的表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteTableDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableDataFlowResponseParams struct {
	// 删除表格分布式索引结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 删除表格分布式索引结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTableDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableDataFlowResponseParams `json:"Response"`
}

func (r *DeleteTableDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableGroupRequestParams struct {
	// 表格组所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

type DeleteTableGroupRequest struct {
	*tchttp.BaseRequest
	
	// 表格组所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`
}

func (r *DeleteTableGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableGroupResponseParams struct {
	// 删除表格组所创建的任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTableGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableGroupResponseParams `json:"Response"`
}

func (r *DeleteTableGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableIndexRequestParams struct {
	// 表格所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除分布式索引的表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteTableIndexRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除分布式索引的表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteTableIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableIndexResponseParams struct {
	// 删除表格分布式索引结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 删除表格分布式索引结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTableIndexResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableIndexResponseParams `json:"Response"`
}

func (r *DeleteTableIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTablesRequestParams struct {
	// 待删除表所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DeleteTablesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除表所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待删除表信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DeleteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTablesResponseParams struct {
	// 删除表结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 删除表结果详情列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTablesResponseParams `json:"Response"`
}

func (r *DeleteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsRequestParams struct {
	// 集群ID，用于获取指定集群的单据
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 分页，限制当前返回多少条记录，大于等于10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，从多少条数据开始返回
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 申请单状态，用于过滤，0-待审核 1-已经审核并提交任务 2-已驳回
	CensorStatus *int64 `json:"CensorStatus,omitempty" name:"CensorStatus"`

	// 表格组id，用于过滤
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名，用于过滤
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 申请人uin，用于过滤
	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`

	// 申请类型，用于过滤，0加表 1删除表 2清理表 3修改表 4表重建 5存储层扩缩容 6接入层扩缩容 7复制表数据 8key回档
	ApplyType *int64 `json:"ApplyType,omitempty" name:"ApplyType"`
}

type DescribeApplicationsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，用于获取指定集群的单据
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 分页，限制当前返回多少条记录，大于等于10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页，从多少条数据开始返回
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 申请单状态，用于过滤，0-待审核 1-已经审核并提交任务 2-已驳回
	CensorStatus *int64 `json:"CensorStatus,omitempty" name:"CensorStatus"`

	// 表格组id，用于过滤
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名，用于过滤
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 申请人uin，用于过滤
	Applicant *string `json:"Applicant,omitempty" name:"Applicant"`

	// 申请类型，用于过滤，0加表 1删除表 2清理表 3修改表 4表重建 5存储层扩缩容 6接入层扩缩容 7复制表数据 8key回档
	ApplyType *int64 `json:"ApplyType,omitempty" name:"ApplyType"`
}

func (r *DescribeApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CensorStatus")
	delete(f, "TableGroupId")
	delete(f, "TableName")
	delete(f, "Applicant")
	delete(f, "ApplyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationsResponseParams struct {
	// 申请单列表
	Applications []*Application `json:"Applications,omitempty" name:"Applications"`

	// 申请单个数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationsResponseParams `json:"Response"`
}

func (r *DescribeApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTagsRequestParams struct {
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

type DescribeClusterTagsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterTagsResponseParams struct {
	// 集群标签信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*TagsInfoOfCluster `json:"Rows,omitempty" name:"Rows"`

	// 返回结果个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterTagsResponseParams `json:"Response"`
}

func (r *DescribeClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 指定查询的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 查询过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是否启用Ipv6
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 查询过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 是否启用Ipv6
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Ipv6Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群实例数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 集群实例列表
	Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdlFileInfosRequestParams struct {
	// 文件所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 文件所属表格组ID
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 指定文件ID列表
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIdlFileInfosRequest struct {
	*tchttp.BaseRequest
	
	// 文件所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 文件所属表格组ID
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 指定文件ID列表
	IdlFileIds []*string `json:"IdlFileIds,omitempty" name:"IdlFileIds"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIdlFileInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdlFileInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "IdlFileIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdlFileInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdlFileInfosResponseParams struct {
	// 文件数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 文件详情列表
	IdlFileInfos []*IdlFileInfo `json:"IdlFileInfos,omitempty" name:"IdlFileInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIdlFileInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdlFileInfosResponseParams `json:"Response"`
}

func (r *DescribeIdlFileInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdlFileInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineRequestParams struct {
	// 不为0，表示查询支持ipv6的机器
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

type DescribeMachineRequest struct {
	*tchttp.BaseRequest
	
	// 不为0，表示查询支持ipv6的机器
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

func (r *DescribeMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ipv6Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMachineResponseParams struct {
	// 独占机器资源列表
	PoolList []*PoolInfo `json:"PoolList,omitempty" name:"PoolList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeMachineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMachineResponseParams `json:"Response"`
}

func (r *DescribeMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineResponse) FromJsonString(s string) error {
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
	// 可用区详情结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 可用区详情结果列表
	RegionInfos []*RegionInfo `json:"RegionInfos,omitempty" name:"RegionInfos"`

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
type DescribeSnapshotsRequestParams struct {
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 批量拉取快照的表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DescribeSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 批量拉取快照的表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DescribeSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "TableName")
	delete(f, "SnapshotName")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotsResponseParams struct {
	// 快照数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 快照结果列表
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotsResponseParams `json:"Response"`
}

func (r *DescribeSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupTagsRequestParams struct {
	// 待查询标签表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询标签表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`
}

type DescribeTableGroupTagsRequest struct {
	*tchttp.BaseRequest
	
	// 待查询标签表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询标签表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`
}

func (r *DescribeTableGroupTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableGroupTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupTagsResponseParams struct {
	// 表格组标签信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*TagsInfoOfTableGroup `json:"Rows,omitempty" name:"Rows"`

	// 返回结果个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableGroupTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableGroupTagsResponseParams `json:"Response"`
}

func (r *DescribeTableGroupTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupsRequestParams struct {
	// 表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 过滤条件，本接口支持：TableGroupName，TableGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTableGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 过滤条件，本接口支持：TableGroupName，TableGroupId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTableGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableGroupsResponseParams struct {
	// 表格组数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格组信息列表
	TableGroups []*TableGroupInfo `json:"TableGroups,omitempty" name:"TableGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableGroupsResponseParams `json:"Response"`
}

func (r *DescribeTableGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableTagsRequestParams struct {
	// 表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type DescribeTableTagsRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *DescribeTableTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableTagsResponseParams struct {
	// 返回结果总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格标签信息列表
	Rows []*TagsInfoOfTable `json:"Rows,omitempty" name:"Rows"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableTagsResponseParams `json:"Response"`
}

func (r *DescribeTableTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesInRecycleRequestParams struct {
	// 待查询表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询表格所属表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询结果偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果返回记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTablesInRecycleRequest struct {
	*tchttp.BaseRequest
	
	// 待查询表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询表格所属表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询结果偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果返回记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesInRecycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesInRecycleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesInRecycleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesInRecycleResponseParams struct {
	// 表格数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格详情结果列表
	TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTablesInRecycleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesInRecycleResponseParams `json:"Response"`
}

func (r *DescribeTablesInRecycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesInRecycleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 待查询表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询表格所属表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 待查询表格信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询结果偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果返回记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 待查询表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待查询表格所属表格组ID列表
	TableGroupIds []*string `json:"TableGroupIds,omitempty" name:"TableGroupIds"`

	// 待查询表格信息列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 过滤条件，本接口支持：TableName，TableInstanceId
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询结果偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果返回记录数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupIds")
	delete(f, "SelectedTables")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 表格数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格详情结果列表
	TableInfos []*TableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 需要查询任务所属的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 需要查询的任务ID列表
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 过滤条件，本接口支持：Content，TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询任务所属的集群ID列表
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 需要查询的任务ID列表
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 过滤条件，本接口支持：Content，TaskType, Operator, Time
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询列表偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询列表返回记录数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "TaskIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 查询到的任务详情列表
	TaskInfos []*TaskInfoNew `json:"TaskInfos,omitempty" name:"TaskInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUinInWhitelistRequestParams struct {

}

type DescribeUinInWhitelistRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUinInWhitelistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUinInWhitelistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUinInWhitelistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUinInWhitelistResponseParams struct {
	// 查询结果：`FALSE` 否；`TRUE` 是
	Result *string `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUinInWhitelistResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUinInWhitelistResponseParams `json:"Response"`
}

func (r *DescribeUinInWhitelistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUinInWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRestProxyRequestParams struct {
	// 对应appid
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DisableRestProxyRequest struct {
	*tchttp.BaseRequest
	
	// 对应appid
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DisableRestProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRestProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRestProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRestProxyResponseParams struct {
	// RestProxy的状态，0为关闭，1为开启中，2为开启，3为关闭中
	RestProxyStatus *uint64 `json:"RestProxyStatus,omitempty" name:"RestProxyStatus"`

	// TaskId由 AppInstanceId-taskId 组成，以区分不同集群的任务
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DisableRestProxyResponse struct {
	*tchttp.BaseResponse
	Response *DisableRestProxyResponseParams `json:"Response"`
}

func (r *DisableRestProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRestProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRestProxyRequestParams struct {
	// 对应于appid
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type EnableRestProxyRequest struct {
	*tchttp.BaseRequest
	
	// 对应于appid
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *EnableRestProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRestProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableRestProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRestProxyResponseParams struct {
	// RestProxy的状态，0为关闭，1为开启中，2为开启，3为关闭中
	RestProxyStatus *uint64 `json:"RestProxyStatus,omitempty" name:"RestProxyStatus"`

	// TaskId由 AppInstanceId-taskId 组成，以区分不同集群的任务
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type EnableRestProxyResponse struct {
	*tchttp.BaseResponse
	Response *EnableRestProxyResponseParams `json:"Response"`
}

func (r *EnableRestProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRestProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorInfo struct {
	// 错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误信息
	Message *string `json:"Message,omitempty" name:"Message"`
}

type FieldInfo struct {
	// 表格字段名称
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 字段是否是主键字段
	IsPrimaryKey *string `json:"IsPrimaryKey,omitempty" name:"IsPrimaryKey"`

	// 字段类型
	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`

	// 字段长度
	FieldSize *int64 `json:"FieldSize,omitempty" name:"FieldSize"`
}

type Filter struct {
	// 过滤字段名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 过滤字段值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type IdlFileInfo struct {
	// 文件名称，不包含扩展名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 数据描述语言（IDL）类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件扩展名
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// 文件大小（Bytes）
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件ID，对于已上传的文件有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// 文件内容，对于本次新上传的文件有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type IdlFileInfoWithoutContent struct {
	// 文件名称，不包含扩展名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 数据描述语言（IDL）类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件扩展名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// 文件大小（Bytes）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 文件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *int64 `json:"FileId,omitempty" name:"FileId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

// Predefined struct for user
type ImportSnapshotsRequestParams struct {
	// 表格所属的集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用于导入的快照信息
	Snapshots *SnapshotInfo `json:"Snapshots,omitempty" name:"Snapshots"`

	// 是否导入部分记录，TRUE表示导入部分记录，FALSE表示全表导入
	ImportSpecialKey *string `json:"ImportSpecialKey,omitempty" name:"ImportSpecialKey"`

	// 是否导入到当前表，TRUE表示导入到当前表，FALSE表示导入到新表
	ImportOriginTable *string `json:"ImportOriginTable,omitempty" name:"ImportOriginTable"`

	// 部分记录的key文件
	KeyFile *KeyFile `json:"KeyFile,omitempty" name:"KeyFile"`

	// 如果导入到新表，此为新表所属的表格组id
	NewTableGroupId *string `json:"NewTableGroupId,omitempty" name:"NewTableGroupId"`

	// 如果导入到新表，此为新表的表名，系统会以该名称自动创建一张结构相同的空表
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

type ImportSnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属的集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 用于导入的快照信息
	Snapshots *SnapshotInfo `json:"Snapshots,omitempty" name:"Snapshots"`

	// 是否导入部分记录，TRUE表示导入部分记录，FALSE表示全表导入
	ImportSpecialKey *string `json:"ImportSpecialKey,omitempty" name:"ImportSpecialKey"`

	// 是否导入到当前表，TRUE表示导入到当前表，FALSE表示导入到新表
	ImportOriginTable *string `json:"ImportOriginTable,omitempty" name:"ImportOriginTable"`

	// 部分记录的key文件
	KeyFile *KeyFile `json:"KeyFile,omitempty" name:"KeyFile"`

	// 如果导入到新表，此为新表所属的表格组id
	NewTableGroupId *string `json:"NewTableGroupId,omitempty" name:"NewTableGroupId"`

	// 如果导入到新表，此为新表的表名，系统会以该名称自动创建一张结构相同的空表
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`
}

func (r *ImportSnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Snapshots")
	delete(f, "ImportSpecialKey")
	delete(f, "ImportOriginTable")
	delete(f, "KeyFile")
	delete(f, "NewTableGroupId")
	delete(f, "NewTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportSnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSnapshotsResponseParams struct {
	// TaskId由 AppInstanceId-taskId 组成，以区分不同集群的任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImportSnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *ImportSnapshotsResponseParams `json:"Response"`
}

func (r *ImportSnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KafkaInfo struct {
	// kafaka address
	Address *string `json:"Address,omitempty" name:"Address"`

	// kafaka topic
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// kafka username
	User *string `json:"User,omitempty" name:"User"`

	// kafka password
	Password *string `json:"Password,omitempty" name:"Password"`

	// ckafka实例
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 是否走VPC
	IsVpc *int64 `json:"IsVpc,omitempty" name:"IsVpc"`
}

type KeyFile struct {
	// key文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// key文件扩展名
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// key文件内容
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`

	// key文件大小
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`
}

type MachineInfo struct {
	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 机器数量
	MachineNum *int64 `json:"MachineNum,omitempty" name:"MachineNum"`
}

type MergeTableResult struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 成功时此字段返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 对比的表格信息
	Table *CompareTablesInfo `json:"Table,omitempty" name:"Table"`

	// 申请单Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

// Predefined struct for user
type MergeTablesDataRequestParams struct {
	// 选取的表格
	SelectedTables []*MergeTablesInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// true只做对比，false既对比又执行
	IsOnlyCompare *bool `json:"IsOnlyCompare,omitempty" name:"IsOnlyCompare"`
}

type MergeTablesDataRequest struct {
	*tchttp.BaseRequest
	
	// 选取的表格
	SelectedTables []*MergeTablesInfo `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// true只做对比，false既对比又执行
	IsOnlyCompare *bool `json:"IsOnlyCompare,omitempty" name:"IsOnlyCompare"`
}

func (r *MergeTablesDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergeTablesDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SelectedTables")
	delete(f, "IsOnlyCompare")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergeTablesDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MergeTablesDataResponseParams struct {
	// 合服结果集
	Results []*MergeTableResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MergeTablesDataResponse struct {
	*tchttp.BaseResponse
	Response *MergeTablesDataResponseParams `json:"Response"`
}

func (r *MergeTablesDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergeTablesDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergeTablesInfo struct {
	// 合服的表格信息
	MergeTables *CompareTablesInfo `json:"MergeTables,omitempty" name:"MergeTables"`

	// 是否检查索引
	CheckIndex *bool `json:"CheckIndex,omitempty" name:"CheckIndex"`
}

// Predefined struct for user
type ModifyCensorshipRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群是否开启审核 0-关闭 1-开启
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// 审批人uin列表
	Uins []*string `json:"Uins,omitempty" name:"Uins"`
}

type ModifyCensorshipRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群是否开启审核 0-关闭 1-开启
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// 审批人uin列表
	Uins []*string `json:"Uins,omitempty" name:"Uins"`
}

func (r *ModifyCensorshipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCensorshipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Censorship")
	delete(f, "Uins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCensorshipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCensorshipResponseParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 已加入审批人的uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uins []*string `json:"Uins,omitempty" name:"Uins"`

	// 集群是否开启审核 0-关闭 1-开启
	Censorship *int64 `json:"Censorship,omitempty" name:"Censorship"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCensorshipResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCensorshipResponseParams `json:"Response"`
}

func (r *ModifyCensorshipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCensorshipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterMachineRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// svr占用的机器
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// proxy占用的机器
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// 集群类型1共享集群2独占集群
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`
}

type ModifyClusterMachineRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// svr占用的机器
	ServerList []*MachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// proxy占用的机器
	ProxyList []*MachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`

	// 集群类型1共享集群2独占集群
	ClusterType *int64 `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *ModifyClusterMachineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterMachineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ServerList")
	delete(f, "ProxyList")
	delete(f, "ClusterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterMachineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterMachineResponseParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterMachineResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterMachineResponseParams `json:"Response"`
}

func (r *ModifyClusterMachineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameRequestParams struct {
	// 需要修改名称的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要修改的集群名称，可使用中文或英文字符，最大长度32个字符
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type ModifyClusterNameRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改名称的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 需要修改的集群名称，可使用中文或英文字符，最大长度32个字符
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyClusterNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterNameResponseParams `json:"Response"`
}

func (r *ModifyClusterNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordRequestParams struct {
	// 需要修改密码的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 集群旧密码预期失效时间
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// 集群新密码，密码必须是a-zA-Z0-9的字符,且必须包含数字和大小写字母
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 更新模式： `1` 更新密码；`2` 更新旧密码失效时间，默认为`1` 模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type ModifyClusterPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改密码的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群旧密码
	OldPassword *string `json:"OldPassword,omitempty" name:"OldPassword"`

	// 集群旧密码预期失效时间
	OldPasswordExpireTime *string `json:"OldPasswordExpireTime,omitempty" name:"OldPasswordExpireTime"`

	// 集群新密码，密码必须是a-zA-Z0-9的字符,且必须包含数字和大小写字母
	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`

	// 更新模式： `1` 更新密码；`2` 更新旧密码失效时间，默认为`1` 模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyClusterPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "OldPassword")
	delete(f, "OldPasswordExpireTime")
	delete(f, "NewPassword")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterPasswordResponseParams `json:"Response"`
}

func (r *ModifyClusterPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterTagsRequestParams struct {
	// 待修改标签的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待增加或修改的标签列表
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 待删除的标签
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyClusterTagsRequest struct {
	*tchttp.BaseRequest
	
	// 待修改标签的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待增加或修改的标签列表
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 待删除的标签
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyClusterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterTagsResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClusterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterTagsResponseParams `json:"Response"`
}

func (r *ModifyClusterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsRequestParams struct {
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 快照列表
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type ModifySnapshotsRequest struct {
	*tchttp.BaseRequest
	
	// 表格所属集群id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 快照列表
	SelectedTables []*SnapshotInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *ModifySnapshotsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotsResponseParams struct {
	// 批量修改的快照数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 批量修改的快照结果列表
	TableResults []*SnapshotResult `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySnapshotsResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotsResponseParams `json:"Response"`
}

func (r *ModifySnapshotsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupNameRequestParams struct {
	// 表格组所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改名称的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 新的表格组名称，可以使用中英文字符和符号
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`
}

type ModifyTableGroupNameRequest struct {
	*tchttp.BaseRequest
	
	// 表格组所属的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改名称的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 新的表格组名称，可以使用中英文字符和符号
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`
}

func (r *ModifyTableGroupNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "TableGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableGroupNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableGroupNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableGroupNameResponseParams `json:"Response"`
}

func (r *ModifyTableGroupNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupTagsRequestParams struct {
	// 待修改标签表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改标签表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 待增加或修改的标签列表
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 待删除的标签
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyTableGroupTagsRequest struct {
	*tchttp.BaseRequest
	
	// 待修改标签表格组所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改标签表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 待增加或修改的标签列表
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 待删除的标签
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyTableGroupTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableGroupTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableGroupTagsResponseParams struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableGroupTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableGroupTagsResponseParams `json:"Response"`
}

func (r *ModifyTableGroupTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableGroupTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableMemosRequestParams struct {
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 选定表详情列表
	TableMemos []*SelectedTableInfoNew `json:"TableMemos,omitempty" name:"TableMemos"`
}

type ModifyTableMemosRequest struct {
	*tchttp.BaseRequest
	
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 选定表详情列表
	TableMemos []*SelectedTableInfoNew `json:"TableMemos,omitempty" name:"TableMemos"`
}

func (r *ModifyTableMemosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableMemosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableMemos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableMemosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableMemosResponseParams struct {
	// 表备注修改结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表备注修改结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableMemosResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableMemosResponseParams `json:"Response"`
}

func (r *ModifyTableMemosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableMemosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableQuotasRequestParams struct {
	// 带扩缩容表所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 已选中待修改的表配额列表
	TableQuotas []*SelectedTableInfoNew `json:"TableQuotas,omitempty" name:"TableQuotas"`
}

type ModifyTableQuotasRequest struct {
	*tchttp.BaseRequest
	
	// 带扩缩容表所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 已选中待修改的表配额列表
	TableQuotas []*SelectedTableInfoNew `json:"TableQuotas,omitempty" name:"TableQuotas"`
}

func (r *ModifyTableQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableQuotasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableQuotas")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableQuotasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableQuotasResponseParams struct {
	// 扩缩容结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 扩缩容结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableQuotasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableQuotasResponseParams `json:"Response"`
}

func (r *ModifyTableQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableQuotasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableTagsRequestParams struct {
	// 待修改标签表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改标签表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 待增加或修改的标签列表
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 待删除的标签列表
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

type ModifyTableTagsRequest struct {
	*tchttp.BaseRequest
	
	// 待修改标签表格所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待修改标签表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 待增加或修改的标签列表
	ReplaceTags []*TagInfoUnit `json:"ReplaceTags,omitempty" name:"ReplaceTags"`

	// 待删除的标签列表
	DeleteTags []*TagInfoUnit `json:"DeleteTags,omitempty" name:"DeleteTags"`
}

func (r *ModifyTableTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "ReplaceTags")
	delete(f, "DeleteTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableTagsResponseParams struct {
	// 返回结果总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回结果
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTableTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableTagsResponseParams `json:"Response"`
}

func (r *ModifyTableTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTablesRequestParams struct {
	// 待修改表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 选中的改表IDL文件
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 待改表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type ModifyTablesRequest struct {
	*tchttp.BaseRequest
	
	// 待修改表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 选中的改表IDL文件
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 待改表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *ModifyTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "IdlFiles")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTablesResponseParams struct {
	// 修改表结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 修改表结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTablesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTablesResponseParams `json:"Response"`
}

func (r *ModifyTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParsedTableInfoNew struct {
	// 表格描述语言类型：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格数据结构类型：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 主键字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyFields *string `json:"KeyFields,omitempty" name:"KeyFields"`

	// 原主键字段信息，改表校验时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldKeyFields *string `json:"OldKeyFields,omitempty" name:"OldKeyFields"`

	// 非主键字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueFields *string `json:"ValueFields,omitempty" name:"ValueFields"`

	// 原非主键字段信息，改表校验时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldValueFields *string `json:"OldValueFields,omitempty" name:"OldValueFields"`

	// 所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 主键字段总大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SumKeyFieldSize *int64 `json:"SumKeyFieldSize,omitempty" name:"SumKeyFieldSize"`

	// 非主键字段总大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SumValueFieldSize *int64 `json:"SumValueFieldSize,omitempty" name:"SumValueFieldSize"`

	// 索引键集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexKeySet *string `json:"IndexKeySet,omitempty" name:"IndexKeySet"`

	// 分表因子集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// TDR版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TdrVersion *int64 `json:"TdrVersion,omitempty" name:"TdrVersion"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// LIST类型表格元素个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// SORTLIST类型表格排序字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// SORTLIST类型表格排序顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`
}

type PoolInfo struct {
	// 唯一id
	PoolUid *int64 `json:"PoolUid,omitempty" name:"PoolUid"`

	// 是否支持ipv6
	Ipv6Enable *int64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`

	// 剩余可用app
	AvailableAppCount *int64 `json:"AvailableAppCount,omitempty" name:"AvailableAppCount"`

	// svr机器列表
	ServerList []*ServerMachineInfo `json:"ServerList,omitempty" name:"ServerList"`

	// proxy机器列表
	ProxyList []*ProxyMachineInfo `json:"ProxyList,omitempty" name:"ProxyList"`
}

type ProxyDetailInfo struct {
	// proxy的唯一id
	ProxyUid *string `json:"ProxyUid,omitempty" name:"ProxyUid"`

	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 请求包速度
	ProcessSpeed *int64 `json:"ProcessSpeed,omitempty" name:"ProcessSpeed"`

	// 请求包时延
	AverageProcessDelay *int64 `json:"AverageProcessDelay,omitempty" name:"AverageProcessDelay"`

	// 慢处理包速度
	SlowProcessSpeed *int64 `json:"SlowProcessSpeed,omitempty" name:"SlowProcessSpeed"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type ProxyMachineInfo struct {
	// 唯一id
	ProxyUid *string `json:"ProxyUid,omitempty" name:"ProxyUid"`

	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 可分配proxy资源数
	AvailableCount *int64 `json:"AvailableCount,omitempty" name:"AvailableCount"`
}

// Predefined struct for user
type RecoverRecycleTablesRequestParams struct {
	// 表所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待恢复表信息
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type RecoverRecycleTablesRequest struct {
	*tchttp.BaseRequest
	
	// 表所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待恢复表信息
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *RecoverRecycleTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverRecycleTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverRecycleTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverRecycleTablesResponseParams struct {
	// 恢复表结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 恢复表信息列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecoverRecycleTablesResponse struct {
	*tchttp.BaseResponse
	Response *RecoverRecycleTablesResponseParams `json:"Response"`
}

func (r *RecoverRecycleTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverRecycleTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {
	// 地域Ap-Code
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地域缩写
	RegionAbbr *string `json:"RegionAbbr,omitempty" name:"RegionAbbr"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`

	// 是否支持ipv6，0:不支持，1:支持
	Ipv6Enable *uint64 `json:"Ipv6Enable,omitempty" name:"Ipv6Enable"`
}

// Predefined struct for user
type RollbackTablesRequestParams struct {
	// 待回档表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待回档表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 待回档时间
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// 回档模式，支持：`KEYS`
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type RollbackTablesRequest struct {
	*tchttp.BaseRequest
	
	// 待回档表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待回档表格列表
	SelectedTables []*SelectedTableInfoNew `json:"SelectedTables,omitempty" name:"SelectedTables"`

	// 待回档时间
	RollbackTime *string `json:"RollbackTime,omitempty" name:"RollbackTime"`

	// 回档模式，支持：`KEYS`
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *RollbackTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	delete(f, "RollbackTime")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackTablesResponseParams struct {
	// 表格回档任务结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格回档任务结果列表
	TableResults []*TableRollbackResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RollbackTablesResponse struct {
	*tchttp.BaseResponse
	Response *RollbackTablesResponseParams `json:"Response"`
}

func (r *RollbackTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SelectedTableInfoNew struct {
	// 表所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表实例ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格描述语言类型：`PROTO`或`TDR`
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格数据结构类型：`GENERIC`或`LIST`
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// LIST表元素个数
	ListElementNum *int64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// 表格预留容量（GB）
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// 表格预留读CU
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// 表格预留写CU
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// 表格备注信息
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// Key回档文件名，回档专用
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// Key回档文件扩展名，回档专用
	FileExtType *string `json:"FileExtType,omitempty" name:"FileExtType"`

	// Key回档文件大小，回档专用
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// Key回档文件内容，回档专用
	FileContent *string `json:"FileContent,omitempty" name:"FileContent"`
}

type SelectedTableWithField struct {
	// 表所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表实例ID
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格描述语言类型：`PROTO`或`TDR`
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格数据结构类型：`GENERIC`或`LIST`
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 待创建索引、缓写、数据订阅的字段列表
	SelectedFields []*FieldInfo `json:"SelectedFields,omitempty" name:"SelectedFields"`

	// 索引分片数
	ShardNum *uint64 `json:"ShardNum,omitempty" name:"ShardNum"`

	// ckafka实例信息
	KafkaInfo *KafkaInfo `json:"KafkaInfo,omitempty" name:"KafkaInfo"`
}

type ServerDetailInfo struct {
	// svr唯一id
	ServerUid *string `json:"ServerUid,omitempty" name:"ServerUid"`

	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`

	// 内存占用量
	MemoryRate *int64 `json:"MemoryRate,omitempty" name:"MemoryRate"`

	// 磁盘占用量
	DiskRate *int64 `json:"DiskRate,omitempty" name:"DiskRate"`

	// 读次数
	ReadNum *int64 `json:"ReadNum,omitempty" name:"ReadNum"`

	// 写次数
	WriteNum *int64 `json:"WriteNum,omitempty" name:"WriteNum"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`
}

type ServerMachineInfo struct {
	// 机器唯一id
	ServerUid *string `json:"ServerUid,omitempty" name:"ServerUid"`

	// 机器类型
	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
}

// Predefined struct for user
type SetTableDataFlowRequestParams struct {
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建分布式索引表格列表
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type SetTableDataFlowRequest struct {
	*tchttp.BaseRequest
	
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建分布式索引表格列表
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *SetTableDataFlowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableDataFlowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTableDataFlowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTableDataFlowResponseParams struct {
	// 表格数据订阅创建结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格数据订阅创建结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTableDataFlowResponse struct {
	*tchttp.BaseResponse
	Response *SetTableDataFlowResponseParams `json:"Response"`
}

func (r *SetTableDataFlowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableDataFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTableIndexRequestParams struct {
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建分布式索引表格列表
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

type SetTableIndexRequest struct {
	*tchttp.BaseRequest
	
	// 表所属集群实例ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建分布式索引表格列表
	SelectedTables []*SelectedTableWithField `json:"SelectedTables,omitempty" name:"SelectedTables"`
}

func (r *SetTableIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "SelectedTables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTableIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTableIndexResponseParams struct {
	// 表格分布式索引创建结果数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 表格分布式索引创建结果列表
	TableResults []*TableResultNew `json:"TableResults,omitempty" name:"TableResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTableIndexResponse struct {
	*tchttp.BaseResponse
	Response *SetTableIndexResponseParams `json:"Response"`
}

func (r *SetTableIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTableIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotInfo struct {
	// 所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照时间点
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`

	// 快照过期时间点
	SnapshotDeadTime *string `json:"SnapshotDeadTime,omitempty" name:"SnapshotDeadTime"`
}

type SnapshotInfoNew struct {
	// 所属表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 快照名称
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照过期时间点
	SnapshotDeadTime *string `json:"SnapshotDeadTime,omitempty" name:"SnapshotDeadTime"`
}

type SnapshotResult struct {
	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 快照名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotName *string `json:"SnapshotName,omitempty" name:"SnapshotName"`

	// 快照的时间点
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotTime *string `json:"SnapshotTime,omitempty" name:"SnapshotTime"`

	// 快照的过期时间点
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotDeadTime *string `json:"SnapshotDeadTime,omitempty" name:"SnapshotDeadTime"`

	// 快照创建时间点
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotCreateTime *string `json:"SnapshotCreateTime,omitempty" name:"SnapshotCreateTime"`

	// 快照大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotSize *uint64 `json:"SnapshotSize,omitempty" name:"SnapshotSize"`

	// 快照状态，0 生成中 1 正常 2 删除中 3 已失效 4 回档使用中
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotStatus *uint64 `json:"SnapshotStatus,omitempty" name:"SnapshotStatus"`
}

type SyncTableField struct {
	// TcaplusDB表字段名称
	SourceName *string `json:"SourceName,omitempty" name:"SourceName"`

	// 目标缓写表的字段名称
	TargetName *string `json:"TargetName,omitempty" name:"TargetName"`
}

type SyncTableInfo struct {
	// 目标缓写表的分表数目
	TargetTableSplitNum *uint64 `json:"TargetTableSplitNum,omitempty" name:"TargetTableSplitNum"`

	// 目标缓写表名前缀
	TargetTableNamePrefix []*string `json:"TargetTableNamePrefix,omitempty" name:"TargetTableNamePrefix"`

	// 缓写数据库实例ID
	TargetSyncDBInstanceId *string `json:"TargetSyncDBInstanceId,omitempty" name:"TargetSyncDBInstanceId"`

	// 缓写表所在数据库名称
	TargetDatabaseName *string `json:"TargetDatabaseName,omitempty" name:"TargetDatabaseName"`

	// 缓写状态，0：创建中，1：进行中，2：关闭，-1：被删除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 表格所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格所在表格组ID
	TableGroupId *uint64 `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格ID
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// TcaplusDB表主键字段到目标缓写表字段的映射
	KeyFieldMapping []*SyncTableField `json:"KeyFieldMapping,omitempty" name:"KeyFieldMapping"`

	// TcaplusDB表字段到目标缓写表字段的映射
	ValueFieldMapping []*SyncTableField `json:"ValueFieldMapping,omitempty" name:"ValueFieldMapping"`
}

type TableGroupInfo struct {
	// 表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格组名称
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格组创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 表格组包含的表格数量
	TableCount *uint64 `json:"TableCount,omitempty" name:"TableCount"`

	// 表格组包含的表格存储总量（MB）
	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`

	// 表格Txh备份文件多少天后过期删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxhBackupExpireDay *uint64 `json:"TxhBackupExpireDay,omitempty" name:"TxhBackupExpireDay"`

	// 是否开启mysql负载均衡,0未开启 1开启中 2已开启
	EnableMysql *uint64 `json:"EnableMysql,omitempty" name:"EnableMysql"`

	// mysql负载均衡vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	MysqlConnIp *string `json:"MysqlConnIp,omitempty" name:"MysqlConnIp"`

	// mysql负载均衡vport
	// 注意：此字段可能返回 null，表示取不到有效值。
	MysqlConnPort *uint64 `json:"MysqlConnPort,omitempty" name:"MysqlConnPort"`
}

type TableInfoNew struct {
	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表格数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格所属集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格所属集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 表格所属表格组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupName *string `json:"TableGroupName,omitempty" name:"TableGroupName"`

	// 表格主键字段结构json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyStruct *string `json:"KeyStruct,omitempty" name:"KeyStruct"`

	// 表格非主键字段结构json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStruct *string `json:"ValueStruct,omitempty" name:"ValueStruct"`

	// 表格分表因子集合，对PROTO类型表格有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardingKeySet *string `json:"ShardingKeySet,omitempty" name:"ShardingKeySet"`

	// 表格索引键字段集合，对PROTO类型表格有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStruct *string `json:"IndexStruct,omitempty" name:"IndexStruct"`

	// LIST类型表格元素个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListElementNum *uint64 `json:"ListElementNum,omitempty" name:"ListElementNum"`

	// 表格所关联IDL文件信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 表格预留容量（GB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedVolume *int64 `json:"ReservedVolume,omitempty" name:"ReservedVolume"`

	// 表格预留读CU
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedReadQps *int64 `json:"ReservedReadQps,omitempty" name:"ReservedReadQps"`

	// 表格预留写CU
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReservedWriteQps *int64 `json:"ReservedWriteQps,omitempty" name:"ReservedWriteQps"`

	// 表格实际数据量大小（MB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableSize *int64 `json:"TableSize,omitempty" name:"TableSize"`

	// 表格状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 表格创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 表格最后一次修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 表格备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memo *string `json:"Memo,omitempty" name:"Memo"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// TcaplusDB SDK数据访问接入ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiAccessId *string `json:"ApiAccessId,omitempty" name:"ApiAccessId"`

	// SORTLIST类型表格排序字段个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortFieldNum *int64 `json:"SortFieldNum,omitempty" name:"SortFieldNum"`

	// SORTLIST类型表格排序顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortRule *int64 `json:"SortRule,omitempty" name:"SortRule"`

	// 表格分布式索引/缓写、kafka数据订阅信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbClusterInfoStruct *string `json:"DbClusterInfoStruct,omitempty" name:"DbClusterInfoStruct"`

	// 表格Txh备份文件多少天后过期删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	TxhBackupExpireDay *uint64 `json:"TxhBackupExpireDay,omitempty" name:"TxhBackupExpireDay"`

	// 表格的缓写信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncTableInfo *SyncTableInfo `json:"SyncTableInfo,omitempty" name:"SyncTableInfo"`
}

type TableResultNew struct {
	// 表格实例ID，形如：tcaplus-3be64cbb
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 任务ID列表，对于创建多任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 腾讯云申请审核单Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type TableRollbackResultNew struct {
	// 表格实例ID，形如：tcaplus-3be64cbb
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 任务ID，对于创建单任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格数据结构类型，如：`GENERIC`或`LIST`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitempty" name:"TableType"`

	// 表格数据描述语言（IDL）类型，如：`PROTO`或`TDR`
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableIdlType *string `json:"TableIdlType,omitempty" name:"TableIdlType"`

	// 表格所属表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`

	// 任务ID列表，对于创建多任务的接口有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 上传的key文件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 校验成功Key数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccKeyNum *uint64 `json:"SuccKeyNum,omitempty" name:"SuccKeyNum"`

	// Key文件中包含总的Key数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalKeyNum *uint64 `json:"TotalKeyNum,omitempty" name:"TotalKeyNum"`
}

type TagInfoUnit struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagsInfoOfCluster struct {
	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TagsInfoOfTable struct {
	// 表格实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInstanceId *string `json:"TableInstanceId,omitempty" name:"TableInstanceId"`

	// 表格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TagsInfoOfTableGroup struct {
	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 表格组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfoUnit `json:"Tags,omitempty" name:"Tags"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *ErrorInfo `json:"Error,omitempty" name:"Error"`
}

type TaskInfoNew struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务所关联的TcaplusDB内部事务ID
	TransId *string `json:"TransId,omitempty" name:"TransId"`

	// 任务所属集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 任务所属集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 任务进度
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 任务创建时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 操作者
	Operator *string `json:"Operator,omitempty" name:"Operator"`

	// 任务详情
	Content *string `json:"Content,omitempty" name:"Content"`
}

// Predefined struct for user
type UpdateApplyRequestParams struct {
	// 申请单状态
	ApplyStatus []*ApplyStatus `json:"ApplyStatus,omitempty" name:"ApplyStatus"`
}

type UpdateApplyRequest struct {
	*tchttp.BaseRequest
	
	// 申请单状态
	ApplyStatus []*ApplyStatus `json:"ApplyStatus,omitempty" name:"ApplyStatus"`
}

func (r *UpdateApplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateApplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateApplyResponseParams struct {
	// 已更新的申请单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyResults []*ApplyResult `json:"ApplyResults,omitempty" name:"ApplyResults"`

	// 更新数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateApplyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateApplyResponseParams `json:"Response"`
}

func (r *UpdateApplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateApplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyIdlFilesRequestParams struct {
	// 待创建表格的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建表格的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 曾经上传过的IDL文件信息列表，与NewIdlFiles至少有一者
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// 待上传的IDL文件信息列表，与ExistingIdlFiles至少有一者
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

type VerifyIdlFilesRequest struct {
	*tchttp.BaseRequest
	
	// 待创建表格的集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 待创建表格的表格组ID
	TableGroupId *string `json:"TableGroupId,omitempty" name:"TableGroupId"`

	// 曾经上传过的IDL文件信息列表，与NewIdlFiles至少有一者
	ExistingIdlFiles []*IdlFileInfo `json:"ExistingIdlFiles,omitempty" name:"ExistingIdlFiles"`

	// 待上传的IDL文件信息列表，与ExistingIdlFiles至少有一者
	NewIdlFiles []*IdlFileInfo `json:"NewIdlFiles,omitempty" name:"NewIdlFiles"`
}

func (r *VerifyIdlFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyIdlFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "TableGroupId")
	delete(f, "ExistingIdlFiles")
	delete(f, "NewIdlFiles")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyIdlFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyIdlFilesResponseParams struct {
	// 本次上传校验所有的IDL文件信息列表
	IdlFiles []*IdlFileInfo `json:"IdlFiles,omitempty" name:"IdlFiles"`

	// 读取IDL描述文件后解析出的合法表数量，不包含已经创建的表
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 读取IDL描述文件后解析出的合法表列表，不包含已经创建的表
	TableInfos []*ParsedTableInfoNew `json:"TableInfos,omitempty" name:"TableInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type VerifyIdlFilesResponse struct {
	*tchttp.BaseResponse
	Response *VerifyIdlFilesResponseParams `json:"Response"`
}

func (r *VerifyIdlFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyIdlFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}