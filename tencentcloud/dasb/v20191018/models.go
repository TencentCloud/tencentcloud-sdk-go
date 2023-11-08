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

package v20191018

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ACTemplate struct {
	// 模版id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`

	// 模版描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type Acl struct {
	// 访问权限ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 访问权限名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil" name:"AllowDiskRedirect"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（目前含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（目前含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，暂未启用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，暂未启用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil" name:"MaxFileDownSize"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil" name:"AllowAnyAccount"`

	// 关联的用户列表
	UserSet []*User `json:"UserSet,omitnil" name:"UserSet"`

	// 关联的用户组列表
	UserGroupSet []*Group `json:"UserGroupSet,omitnil" name:"UserGroupSet"`

	// 关联的资产列表
	DeviceSet []*Device `json:"DeviceSet,omitnil" name:"DeviceSet"`

	// 关联的资产组列表
	DeviceGroupSet []*Group `json:"DeviceGroupSet,omitnil" name:"DeviceGroupSet"`

	// 关联的账号列表
	AccountSet []*string `json:"AccountSet,omitnil" name:"AccountSet"`

	// 关联的高危命令模板列表
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitnil" name:"CmdTemplateSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil" name:"AllowDiskFileDown"`

	// 是否开启 rz sz 文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil" name:"AllowShellFileUp"`

	// 是否开启 rz sz 文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 所属部门的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil" name:"Department"`

	// 是否允许使用访问串，默认允许
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil" name:"AllowAccessCredential"`

	// 关联的数据库高危命令列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ACTemplateSet []*ACTemplate `json:"ACTemplateSet,omitnil" name:"ACTemplateSet"`
}

// Predefined struct for user
type AddDeviceGroupMembersRequestParams struct {
	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 需要添加到资产组的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

type AddDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 需要添加到资产组的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

func (r *AddDeviceGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDeviceGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDeviceGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDeviceGroupMembersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *AddDeviceGroupMembersResponseParams `json:"Response"`
}

func (r *AddDeviceGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserGroupMembersRequestParams struct {
	// 用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

type AddUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

func (r *AddUserGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserGroupMembersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddUserGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *AddUserGroupMembersResponseParams `json:"Response"`
}

func (r *AddUserGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetSyncStatus struct {
	// 上一次同步完成的时间
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 上一次同步的结果。 0 - 从未进行, 1 - 成功， 2 - 失败
	LastStatus *uint64 `json:"LastStatus,omitnil" name:"LastStatus"`

	// 同步任务是否正在进行中
	InProcess *bool `json:"InProcess,omitnil" name:"InProcess"`
}

type AuditLogResult struct {
	// 被审计会话的Sid
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 审计者的编号
	Uin *string `json:"Uin,omitnil" name:"Uin"`

	// 审计动作发生的时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 审计者的Ip
	ClientIp *string `json:"ClientIp,omitnil" name:"ClientIp"`

	// 审计动作类型，1--回放、2--中断、3--监控
	Operation *int64 `json:"Operation,omitnil" name:"Operation"`

	// 被审计主机的Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 被审计主机的主机名
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 被审计会话所属的类型，如字符会话
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 被审计主机的内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 被审计主机的外部Ip
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 审计者的子账号
	SubAccountUin *string `json:"SubAccountUin,omitnil" name:"SubAccountUin"`
}

// Predefined struct for user
type BindDeviceAccountPasswordRequestParams struct {
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机账号密码
	Password *string `json:"Password,omitnil" name:"Password"`
}

type BindDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机账号密码
	Password *string `json:"Password,omitnil" name:"Password"`
}

func (r *BindDeviceAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDeviceAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDeviceAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDeviceAccountPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindDeviceAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *BindDeviceAccountPasswordResponseParams `json:"Response"`
}

func (r *BindDeviceAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDeviceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDeviceAccountPrivateKeyRequestParams struct {
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机账号私钥，最新长度128字节，最大长度8192字节
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// 主机账号私钥口令，最大长度256字节
	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitnil" name:"PrivateKeyPassword"`
}

type BindDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机账号私钥，最新长度128字节，最大长度8192字节
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// 主机账号私钥口令，最大长度256字节
	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitnil" name:"PrivateKeyPassword"`
}

func (r *BindDeviceAccountPrivateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDeviceAccountPrivateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "PrivateKey")
	delete(f, "PrivateKeyPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDeviceAccountPrivateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDeviceAccountPrivateKeyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindDeviceAccountPrivateKeyResponse struct {
	*tchttp.BaseResponse
	Response *BindDeviceAccountPrivateKeyResponseParams `json:"Response"`
}

func (r *BindDeviceAccountPrivateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDeviceAccountPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDeviceResourceRequestParams struct {
	// 资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 堡垒机服务ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type BindDeviceResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 堡垒机服务ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

func (r *BindDeviceResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDeviceResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceIdSet")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDeviceResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDeviceResourceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindDeviceResourceResponse struct {
	*tchttp.BaseResponse
	Response *BindDeviceResourceResponseParams `json:"Response"`
}

func (r *BindDeviceResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindDeviceResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmdTemplate struct {
	// 高危命令模板ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 高危命令模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命令列表，命令之间用换行符（"\n"）分隔
	CmdList *string `json:"CmdList,omitnil" name:"CmdList"`
}

type Command struct {
	// 命令
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// 命令输入的时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms
	TimeOffset *uint64 `json:"TimeOffset,omitnil" name:"TimeOffset"`

	// 命令执行情况，1--允许，2--拒绝，3--确认
	Action *int64 `json:"Action,omitnil" name:"Action"`

	// 会话id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 设备account
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil" name:"Account"`

	// 设备ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// source ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromIp *string `json:"FromIp,omitnil" name:"FromIp"`

	// 该命令所属会话的会话开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionTime *string `json:"SessionTime,omitnil" name:"SessionTime"`

	// 该命令所属会话的会话开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SessTime is deprecated.
	SessTime *string `json:"SessTime,omitnil" name:"SessTime"`

	// 复核时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfirmTime *string `json:"ConfirmTime,omitnil" name:"ConfirmTime"`

	// 用户部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentId *string `json:"UserDepartmentId,omitnil" name:"UserDepartmentId"`

	// 用户部门name
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentName *string `json:"UserDepartmentName,omitnil" name:"UserDepartmentName"`

	// 设备部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitnil" name:"DeviceDepartmentId"`

	// 设备部门name
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitnil" name:"DeviceDepartmentName"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil" name:"AllowClipTextDown"`

	// 是否开启 SFTP 文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil" name:"MaxFileUpSize"`

	// 是否开启 SFTP 文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil" name:"MaxFileDownSize"`

	// 关联的用户ID集合
	UserIdSet []*uint64 `json:"UserIdSet,omitnil" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil" name:"UserGroupIdSet"`

	// 关联的资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil" name:"ACTemplateIdSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 访问权限所属部门的ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 是否允许使用访问串，默认允许
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil" name:"AllowAccessCredential"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil" name:"AllowClipTextDown"`

	// 是否开启 SFTP 文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil" name:"MaxFileUpSize"`

	// 是否开启 SFTP 文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil" name:"MaxFileDownSize"`

	// 关联的用户ID集合
	UserIdSet []*uint64 `json:"UserIdSet,omitnil" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil" name:"UserGroupIdSet"`

	// 关联的资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil" name:"ACTemplateIdSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 访问权限所属部门的ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 是否允许使用访问串，默认允许
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil" name:"AllowAccessCredential"`
}

func (r *CreateAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AllowDiskRedirect")
	delete(f, "AllowAnyAccount")
	delete(f, "AllowClipFileUp")
	delete(f, "AllowClipFileDown")
	delete(f, "AllowClipTextUp")
	delete(f, "AllowClipTextDown")
	delete(f, "AllowFileUp")
	delete(f, "MaxFileUpSize")
	delete(f, "AllowFileDown")
	delete(f, "MaxFileDownSize")
	delete(f, "UserIdSet")
	delete(f, "UserGroupIdSet")
	delete(f, "DeviceIdSet")
	delete(f, "DeviceGroupIdSet")
	delete(f, "AccountSet")
	delete(f, "CmdTemplateIdSet")
	delete(f, "ACTemplateIdSet")
	delete(f, "AllowDiskFileUp")
	delete(f, "AllowDiskFileDown")
	delete(f, "AllowShellFileUp")
	delete(f, "AllowShellFileDown")
	delete(f, "AllowFileDel")
	delete(f, "ValidateFrom")
	delete(f, "ValidateTo")
	delete(f, "DepartmentId")
	delete(f, "AllowAccessCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAclResponseParams struct {
	// 新建成功的访问权限ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAclResponse struct {
	*tchttp.BaseResponse
	Response *CreateAclResponseParams `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetSyncJobRequestParams struct {
	// 同步资产类别，1 - 主机资产, 2 - 数据库资产
	Category *uint64 `json:"Category,omitnil" name:"Category"`
}

type CreateAssetSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步资产类别，1 - 主机资产, 2 - 数据库资产
	Category *uint64 `json:"Category,omitnil" name:"Category"`
}

func (r *CreateAssetSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Category")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAssetSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetSyncJobResponseParams `json:"Response"`
}

func (r *CreateAssetSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCmdTemplateRequestParams struct {
	// 模板名，最大长度32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命令列表，\n分隔，最大长度32768字节
	CmdList *string `json:"CmdList,omitnil" name:"CmdList"`

	// 标识cmdlist字段前端是否为base64加密传值.
	// 0:表示非base64加密
	// 1:表示是base64加密
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`
}

type CreateCmdTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名，最大长度32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命令列表，\n分隔，最大长度32768字节
	CmdList *string `json:"CmdList,omitnil" name:"CmdList"`

	// 标识cmdlist字段前端是否为base64加密传值.
	// 0:表示非base64加密
	// 1:表示是base64加密
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`
}

func (r *CreateCmdTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmdTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "CmdList")
	delete(f, "Encoding")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmdTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCmdTemplateResponseParams struct {
	// 新建成功后返回的记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCmdTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmdTemplateResponseParams `json:"Response"`
}

func (r *CreateCmdTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmdTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceAccountRequestParams struct {
	// 主机记录ID
	DeviceId *uint64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitnil" name:"Account"`
}

type CreateDeviceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 主机记录ID
	DeviceId *uint64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitnil" name:"Account"`
}

func (r *CreateDeviceAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceId")
	delete(f, "Account")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceAccountResponseParams struct {
	// 新建成功后返回的记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDeviceAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceAccountResponseParams `json:"Response"`
}

func (r *CreateDeviceAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceGroupRequestParams struct {
	// 资产组名，最大长度32字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type CreateDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组名，最大长度32字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *CreateDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDeviceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDeviceGroupResponseParams struct {
	// 新建成功的资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDeviceGroupResponseParams `json:"Response"`
}

func (r *CreateDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceRequestParams struct {
	// 部署region
	DeployRegion *string `json:"DeployRegion,omitnil" name:"DeployRegion"`

	// 部署堡垒机的VpcId
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 部署堡垒机的SubnetId
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 资源类型。取值:standard/pro
	ResourceEdition *string `json:"ResourceEdition,omitnil" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil" name:"ResourceNode"`

	// 计费周期
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 计费时长
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 计费模式 1预付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 部署zone
	DeployZone *string `json:"DeployZone,omitnil" name:"DeployZone"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 部署region
	DeployRegion *string `json:"DeployRegion,omitnil" name:"DeployRegion"`

	// 部署堡垒机的VpcId
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 部署堡垒机的SubnetId
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 资源类型。取值:standard/pro
	ResourceEdition *string `json:"ResourceEdition,omitnil" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil" name:"ResourceNode"`

	// 计费周期
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 计费时长
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 计费模式 1预付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 部署zone
	DeployZone *string `json:"DeployZone,omitnil" name:"DeployZone"`
}

func (r *CreateResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeployRegion")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ResourceEdition")
	delete(f, "ResourceNode")
	delete(f, "TimeUnit")
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	delete(f, "AutoRenewFlag")
	delete(f, "DeployZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResourceResponseParams struct {
	// 实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateResourceResponseParams `json:"Response"`
}

func (r *CreateResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupRequestParams struct {
	// 用户组名，最大长度32字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户组所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组名，最大长度32字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 用户组所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *CreateUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupResponseParams struct {
	// 新建成功的用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserGroupResponseParams `json:"Response"`
}

func (r *CreateUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 用户名, 3-20个字符, 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户姓名，最大长度20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码， 按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil" name:"GroupIdSet"`

	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil" name:"ValidateTime"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户名, 3-20个字符, 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户姓名，最大长度20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码， 按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil" name:"GroupIdSet"`

	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil" name:"ValidateTime"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "RealName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "ValidateFrom")
	delete(f, "ValidateTo")
	delete(f, "GroupIdSet")
	delete(f, "AuthType")
	delete(f, "ValidateTime")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 新建用户的ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAclsRequestParams struct {
	// 待删除的权限ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteAclsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的权限ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAclsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAclsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAclsResponseParams `json:"Response"`
}

func (r *DeleteAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCmdTemplatesRequestParams struct {
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteCmdTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteCmdTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmdTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmdTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCmdTemplatesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCmdTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmdTemplatesResponseParams `json:"Response"`
}

func (r *DeleteCmdTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmdTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceAccountsRequestParams struct {
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteDeviceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteDeviceAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceAccountsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDeviceAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceAccountsResponseParams `json:"Response"`
}

func (r *DeleteDeviceAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceGroupMembersRequestParams struct {
	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 需要删除的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

type DeleteDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 需要删除的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

func (r *DeleteDeviceGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceGroupMembersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceGroupMembersResponseParams `json:"Response"`
}

func (r *DeleteDeviceGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceGroupsRequestParams struct {
	// 待删除的资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteDeviceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDeviceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceGroupsResponseParams `json:"Response"`
}

func (r *DeleteDeviceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDevicesRequestParams struct {
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDevicesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDevicesResponseParams `json:"Response"`
}

func (r *DeleteDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupMembersRequestParams struct {
	// 用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 需删除的成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

type DeleteUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 需删除的成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil" name:"MemberIdSet"`
}

func (r *DeleteUserGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "MemberIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupMembersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserGroupMembersResponseParams `json:"Response"`
}

func (r *DeleteUserGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupsRequestParams struct {
	// 待删除的用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserGroupsResponseParams `json:"Response"`
}

func (r *DeleteUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// 待删除的用户ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的用户ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersResponseParams `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Department struct {
	// 部门ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 部门名称，1 - 256个字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 部门管理员账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Managers []*string `json:"Managers,omitnil" name:"Managers"`

	// 管理员用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerUsers []*DepartmentManagerUser `json:"ManagerUsers,omitnil" name:"ManagerUsers"`
}

type DepartmentManagerUser struct {
	// 管理员Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerId *string `json:"ManagerId,omitnil" name:"ManagerId"`

	// 管理员姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerName *string `json:"ManagerName,omitnil" name:"ManagerName"`
}

// Predefined struct for user
type DeployResourceRequestParams struct {
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 需要开通服务的地域
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 需要开通服务的VPC
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 需要开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 需要开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 需要开通服务的VPC名称
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 需要开通服务的VPC对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil" name:"VpcCidrBlock"`

	// 需要开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`
}

type DeployResourceRequest struct {
	*tchttp.BaseRequest
	
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 需要开通服务的地域
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 需要开通服务的VPC
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 需要开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 需要开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 需要开通服务的VPC名称
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 需要开通服务的VPC对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil" name:"VpcCidrBlock"`

	// 需要开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`
}

func (r *DeployResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ApCode")
	delete(f, "Zone")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "CidrBlock")
	delete(f, "VpcName")
	delete(f, "VpcCidrBlock")
	delete(f, "SubnetName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployResourceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeployResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeployResourceResponseParams `json:"Response"`
}

func (r *DeployResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeployResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclsRequestParams struct {
	// 访问权限ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 是否根据Name进行精确查询，默认值false
	Exact *bool `json:"Exact,omitnil" name:"Exact"`

	// 有访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil" name:"AuthorizedUserIdSet"`

	// 有访问权限的资产ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil" name:"AuthorizedDeviceIdSet"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 部门ID，用于过滤属于某个部门的访问权限
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type DescribeAclsRequest struct {
	*tchttp.BaseRequest
	
	// 访问权限ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 是否根据Name进行精确查询，默认值false
	Exact *bool `json:"Exact,omitnil" name:"Exact"`

	// 有访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil" name:"AuthorizedUserIdSet"`

	// 有访问权限的资产ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil" name:"AuthorizedDeviceIdSet"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 部门ID，用于过滤属于某个部门的访问权限
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *DescribeAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Exact")
	delete(f, "AuthorizedUserIdSet")
	delete(f, "AuthorizedDeviceIdSet")
	delete(f, "Status")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclsResponseParams struct {
	// 访问权限总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 访问权限列表
	AclSet []*Acl `json:"AclSet,omitnil" name:"AclSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAclsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAclsResponseParams `json:"Response"`
}

func (r *DescribeAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSyncStatusRequestParams struct {
	// 查询的资产同步类型。1 -主机资产， 2 - 数据库资产
	Category *uint64 `json:"Category,omitnil" name:"Category"`
}

type DescribeAssetSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// 查询的资产同步类型。1 -主机资产， 2 - 数据库资产
	Category *uint64 `json:"Category,omitnil" name:"Category"`
}

func (r *DescribeAssetSyncStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSyncStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Category")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetSyncStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSyncStatusResponseParams struct {
	// 资产同步结果
	Status *AssetSyncStatus `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssetSyncStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetSyncStatusResponseParams `json:"Response"`
}

func (r *DescribeAssetSyncStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmdTemplatesRequestParams struct {
	// 命令模板ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 命令模板名，模糊查询，最大长度64字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeCmdTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 命令模板ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 命令模板名，模糊查询，最大长度64字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeCmdTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmdTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmdTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmdTemplatesResponseParams struct {
	// 命令模板列表
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitnil" name:"CmdTemplateSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmdTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmdTemplatesResponseParams `json:"Response"`
}

func (r *DescribeCmdTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmdTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDasbImageIdsRequestParams struct {

}

type DescribeDasbImageIdsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDasbImageIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDasbImageIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDasbImageIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDasbImageIdsResponseParams struct {
	// 基础镜像ID
	BaseImageId *string `json:"BaseImageId,omitnil" name:"BaseImageId"`

	// AI镜像ID
	AiImageId *string `json:"AiImageId,omitnil" name:"AiImageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDasbImageIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDasbImageIdsResponseParams `json:"Response"`
}

func (r *DescribeDasbImageIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDasbImageIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceAccountsRequestParams struct {
	// 主机账号ID集合，非必需，如果使用IdSet则忽略其他过滤参数
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交
	Account *string `json:"Account,omitnil" name:"Account"`

	// 主机ID，未使用IdSet时必须携带
	DeviceId *uint64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDeviceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID集合，非必需，如果使用IdSet则忽略其他过滤参数
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交
	Account *string `json:"Account,omitnil" name:"Account"`

	// 主机ID，未使用IdSet时必须携带
	DeviceId *uint64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDeviceAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Account")
	delete(f, "DeviceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceAccountsResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 账号信息列表
	DeviceAccountSet []*DeviceAccount `json:"DeviceAccountSet,omitnil" name:"DeviceAccountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceAccountsResponseParams `json:"Response"`
}

func (r *DescribeDeviceAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceGroupMembersRequestParams struct {
	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产
	Bound *bool `json:"Bound,omitnil" name:"Bound"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

type DescribeDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产
	Bound *bool `json:"Bound,omitnil" name:"Bound"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

func (r *DescribeDeviceGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Bound")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Kind")
	delete(f, "DepartmentId")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceGroupMembersResponseParams struct {
	// 资产组成员总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资产组成员列表
	DeviceSet []*Device `json:"DeviceSet,omitnil" name:"DeviceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceGroupMembersResponseParams `json:"Response"`
}

func (r *DescribeDeviceGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceGroupsRequestParams struct {
	// 资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 资产组名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的资产组
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type DescribeDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 资产组名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的资产组
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *DescribeDeviceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceGroupsResponseParams struct {
	// 资产组总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资产组列表
	GroupSet []*Group `json:"GroupSet,omitnil" name:"GroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDeviceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceGroupsResponseParams `json:"Response"`
}

func (r *DescribeDeviceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesRequestParams struct {
	// 资产ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitnil" name:"ApCodeSet"`

	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 有该资产访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil" name:"AuthorizedUserIdSet"`

	// 过滤条件，资产绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitnil" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	KindSet []*uint64 `json:"KindSet,omitnil" name:"KindSet"`

	// 资产是否包含托管账号。1，包含；0，不包含
	ManagedAccount *string `json:"ManagedAccount,omitnil" name:"ManagedAccount"`

	// 过滤条件，可按照部门ID进行过滤
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`

	// 过滤数组。支持的Name：
	// BindingStatus 绑定状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitnil" name:"ApCodeSet"`

	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 有该资产访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil" name:"AuthorizedUserIdSet"`

	// 过滤条件，资产绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitnil" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	KindSet []*uint64 `json:"KindSet,omitnil" name:"KindSet"`

	// 资产是否包含托管账号。1，包含；0，不包含
	ManagedAccount *string `json:"ManagedAccount,omitnil" name:"ManagedAccount"`

	// 过滤条件，可按照部门ID进行过滤
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`

	// 过滤数组。支持的Name：
	// BindingStatus 绑定状态
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Name")
	delete(f, "Ip")
	delete(f, "ApCodeSet")
	delete(f, "Kind")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AuthorizedUserIdSet")
	delete(f, "ResourceIdSet")
	delete(f, "KindSet")
	delete(f, "ManagedAccount")
	delete(f, "DepartmentId")
	delete(f, "TagFilters")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDevicesResponseParams struct {
	// 资产总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 资产信息列表
	DeviceSet []*Device `json:"DeviceSet,omitnil" name:"DeviceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDevicesResponseParams `json:"Response"`
}

func (r *DescribeDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDevicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginEventRequestParams struct {
	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API
	Entry *uint64 `json:"Entry,omitnil" name:"Entry"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeLoginEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API
	Entry *uint64 `json:"Entry,omitnil" name:"Entry"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeLoginEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "RealName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SourceIp")
	delete(f, "Entry")
	delete(f, "Result")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginEventResponseParams struct {
	// 登录日志列表
	LoginEventSet []*LoginEvent `json:"LoginEventSet,omitnil" name:"LoginEventSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLoginEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginEventResponseParams `json:"Response"`
}

func (r *DescribeLoginEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOperationEventRequestParams struct {
	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 操作类型，参考DescribeOperationType返回结果
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeOperationEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 操作类型，参考DescribeOperationType返回结果
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeOperationEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOperationEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "RealName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SourceIp")
	delete(f, "Kind")
	delete(f, "Result")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOperationEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOperationEventResponseParams struct {
	// 操作日志列表
	OperationEventSet []*OperationEvent `json:"OperationEventSet,omitnil" name:"OperationEventSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOperationEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOperationEventResponseParams `json:"Response"`
}

func (r *DescribeOperationEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOperationEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesRequestParams struct {
	// 地域码, 如: ap-guangzhou
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 按照堡垒机开通的 VPC 实例ID查询
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 地域码, 如: ap-guangzhou
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 按照堡垒机开通的 VPC 实例ID查询
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApCode")
	delete(f, "VpcId")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesResponseParams struct {
	// 堡垒机资源列表
	ResourceSet []*Resource `json:"ResourceSet,omitnil" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesResponseParams `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupMembersRequestParams struct {
	// 用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户
	Bound *bool `json:"Bound,omitnil" name:"Bound"`

	// 用户名或用户姓名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type DescribeUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户
	Bound *bool `json:"Bound,omitnil" name:"Bound"`

	// 用户名或用户姓名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *DescribeUserGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Bound")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupMembersResponseParams struct {
	// 用户组成员总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 用户组成员列表
	UserSet []*User `json:"UserSet,omitnil" name:"UserSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupMembersResponseParams `json:"Response"`
}

func (r *DescribeUserGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupsRequestParams struct {
	// 用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 用户组名，模糊查询,长度：0-64字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的用户组
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type DescribeUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 用户组名，模糊查询,长度：0-64字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的用户组
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *DescribeUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupsResponseParams struct {
	// 用户组总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 用户组列表
	GroupSet []*Group `json:"GroupSet,omitnil" name:"GroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupsResponseParams `json:"Response"`
}

func (r *DescribeUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// 如果IdSet不为空，则忽略其他参数
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 邮箱，精确查询
	Email *string `json:"Email,omitnil" name:"Email"`

	// 查询具有指定资产ID访问权限的用户
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitnil" name:"AuthTypeSet"`

	// 部门ID，用于过滤属于某个部门的用户
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 参数过滤数组
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 如果IdSet不为空，则忽略其他参数
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 邮箱，精确查询
	Email *string `json:"Email,omitnil" name:"Email"`

	// 查询具有指定资产ID访问权限的用户
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitnil" name:"AuthTypeSet"`

	// 部门ID，用于过滤属于某个部门的用户
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 参数过滤数组
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	delete(f, "Name")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UserName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "AuthorizedDeviceIdSet")
	delete(f, "AuthTypeSet")
	delete(f, "DepartmentId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// 用户总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 用户列表
	UserSet []*User `json:"UserSet,omitnil" name:"UserSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 资产ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 实例ID，对应CVM、CDB等实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 地域编码
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 操作系统名称
	OsName *string `json:"OsName,omitnil" name:"OsName"`

	// 资产类型 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 所属资产组列表
	GroupSet []*Group `json:"GroupSet,omitnil" name:"GroupSet"`

	// 资产绑定的账号数
	AccountCount *uint64 `json:"AccountCount,omitnil" name:"AccountCount"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 堡垒机服务信息，注意没有绑定服务时为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitnil" name:"Resource"`

	// 资产所属部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil" name:"Department"`

	// 数据库资产的多节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpPortSet []*string `json:"IpPortSet,omitnil" name:"IpPortSet"`
}

type DeviceAccount struct {
	// 账号ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 主机ID
	DeviceId *uint64 `json:"DeviceId,omitnil" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitnil" name:"Account"`

	// true-已托管密码，false-未托管密码
	BoundPassword *bool `json:"BoundPassword,omitnil" name:"BoundPassword"`

	// true-已托管私钥，false-未托管私钥
	BoundPrivateKey *bool `json:"BoundPrivateKey,omitnil" name:"BoundPrivateKey"`
}

type ExternalDevice struct {
	// 操作系统名称，只能是Linux、Windows或MySQL
	OsName *string `json:"OsName,omitnil" name:"OsName"`

	// IP地址
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 主机名，可为空
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资产所属的部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 资产多节点：字段ip和端口
	IpPortSet []*string `json:"IpPortSet,omitnil" name:"IpPortSet"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段的过滤值。
	// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type Group struct {
	// 组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 组名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 所属部门信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil" name:"Department"`

	// 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

// Predefined struct for user
type ImportExternalDeviceRequestParams struct {
	// 资产参数列表
	DeviceSet []*ExternalDevice `json:"DeviceSet,omitnil" name:"DeviceSet"`
}

type ImportExternalDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 资产参数列表
	DeviceSet []*ExternalDevice `json:"DeviceSet,omitnil" name:"DeviceSet"`
}

func (r *ImportExternalDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportExternalDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportExternalDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportExternalDeviceResponseParams struct {
	// 资产ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImportExternalDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ImportExternalDeviceResponseParams `json:"Response"`
}

func (r *ImportExternalDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportExternalDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginEvent struct {
	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 操作时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 来源IP
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API
	Entry *uint64 `json:"Entry,omitnil" name:"Entry"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil" name:"Result"`
}

// Predefined struct for user
type ModifyAclRequestParams struct {
	// 访问权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitnil" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil" name:"UserGroupIdSet"`

	// 关联的资产ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil" name:"ACTemplateIdSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 权限所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 是否允许使用访问串
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil" name:"AllowAccessCredential"`
}

type ModifyAclRequest struct {
	*tchttp.BaseRequest
	
	// 访问权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitnil" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil" name:"UserGroupIdSet"`

	// 关联的资产ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil" name:"ACTemplateIdSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 权限所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 是否允许使用访问串
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil" name:"AllowAccessCredential"`
}

func (r *ModifyAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AllowDiskRedirect")
	delete(f, "AllowAnyAccount")
	delete(f, "Id")
	delete(f, "AllowClipFileUp")
	delete(f, "AllowClipFileDown")
	delete(f, "AllowClipTextUp")
	delete(f, "AllowClipTextDown")
	delete(f, "AllowFileUp")
	delete(f, "MaxFileUpSize")
	delete(f, "AllowFileDown")
	delete(f, "MaxFileDownSize")
	delete(f, "UserIdSet")
	delete(f, "UserGroupIdSet")
	delete(f, "DeviceIdSet")
	delete(f, "DeviceGroupIdSet")
	delete(f, "AccountSet")
	delete(f, "CmdTemplateIdSet")
	delete(f, "ACTemplateIdSet")
	delete(f, "AllowDiskFileUp")
	delete(f, "AllowDiskFileDown")
	delete(f, "AllowShellFileUp")
	delete(f, "AllowShellFileDown")
	delete(f, "AllowFileDel")
	delete(f, "ValidateFrom")
	delete(f, "ValidateTo")
	delete(f, "DepartmentId")
	delete(f, "AllowAccessCredential")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAclResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAclResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAclResponseParams `json:"Response"`
}

func (r *ModifyAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCmdTemplateRequestParams struct {
	// 模板名，最长32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命令列表，\n分隔，最长32768字节
	CmdList *string `json:"CmdList,omitnil" name:"CmdList"`

	// 命令模板ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CmdList字段前端是否base64传值。
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`
}

type ModifyCmdTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名，最长32字符，不能包含空白字符
	Name *string `json:"Name,omitnil" name:"Name"`

	// 命令列表，\n分隔，最长32768字节
	CmdList *string `json:"CmdList,omitnil" name:"CmdList"`

	// 命令模板ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// CmdList字段前端是否base64传值。
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`
}

func (r *ModifyCmdTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmdTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "CmdList")
	delete(f, "Id")
	delete(f, "Encoding")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmdTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCmdTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCmdTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmdTemplateResponseParams `json:"Response"`
}

func (r *ModifyCmdTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmdTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceGroupRequestParams struct {
	// 资产组名，最大长度32字符，不能为空
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type ModifyDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组名，最大长度32字符，不能为空
	Name *string `json:"Name,omitnil" name:"Name"`

	// 资产组ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *ModifyDeviceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Id")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDeviceGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceGroupResponseParams `json:"Response"`
}

func (r *ModifyDeviceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceRequestParams struct {
	// 资产记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 资产所属组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil" name:"GroupIdSet"`

	// 资产所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 资产记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 资产所属组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil" name:"GroupIdSet"`

	// 资产所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *ModifyDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Port")
	delete(f, "GroupIdSet")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDeviceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDeviceResponseParams `json:"Response"`
}

func (r *ModifyDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDeviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceRequestParams struct {
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 已废弃
	Status *string `json:"Status,omitnil" name:"Status"`

	// 已废弃
	ModuleSet []*string `json:"ModuleSet,omitnil" name:"ModuleSet"`

	// 实例版本
	ResourceEdition *string `json:"ResourceEdition,omitnil" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil" name:"ResourceNode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *int64 `json:"PackageBandwidth,omitnil" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *int64 `json:"PackageNode,omitnil" name:"PackageNode"`

	// 日志投递
	LogDelivery *int64 `json:"LogDelivery,omitnil" name:"LogDelivery"`
}

type ModifyResourceRequest struct {
	*tchttp.BaseRequest
	
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 已废弃
	Status *string `json:"Status,omitnil" name:"Status"`

	// 已废弃
	ModuleSet []*string `json:"ModuleSet,omitnil" name:"ModuleSet"`

	// 实例版本
	ResourceEdition *string `json:"ResourceEdition,omitnil" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil" name:"ResourceNode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *int64 `json:"PackageBandwidth,omitnil" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *int64 `json:"PackageNode,omitnil" name:"PackageNode"`

	// 日志投递
	LogDelivery *int64 `json:"LogDelivery,omitnil" name:"LogDelivery"`
}

func (r *ModifyResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "Status")
	delete(f, "ModuleSet")
	delete(f, "ResourceEdition")
	delete(f, "ResourceNode")
	delete(f, "AutoRenewFlag")
	delete(f, "PackageBandwidth")
	delete(f, "PackageNode")
	delete(f, "LogDelivery")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyResourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceResponseParams `json:"Response"`
}

func (r *ModifyResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 用户ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 用户姓名，最大长度20个字符，不能包含空格
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil" name:"GroupIdSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil" name:"ValidateTime"`

	// 用户所属部门的ID，如1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 用户姓名，最大长度20个字符，不能包含空格
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil" name:"GroupIdSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil" name:"ValidateTime"`

	// 用户所属部门的ID，如1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "RealName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "ValidateFrom")
	delete(f, "ValidateTo")
	delete(f, "GroupIdSet")
	delete(f, "AuthType")
	delete(f, "ValidateTime")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperationEvent struct {
	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 操作时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 来源IP
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 操作类型
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 具体操作内容
	Operation *string `json:"Operation,omitnil" name:"Operation"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil" name:"Result"`
}

// Predefined struct for user
type ResetDeviceAccountPasswordRequestParams struct {
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type ResetDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *ResetDeviceAccountPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceAccountPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDeviceAccountPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDeviceAccountPasswordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetDeviceAccountPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetDeviceAccountPasswordResponseParams `json:"Response"`
}

func (r *ResetDeviceAccountPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDeviceAccountPrivateKeyRequestParams struct {
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type ResetDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *ResetDeviceAccountPrivateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceAccountPrivateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetDeviceAccountPrivateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetDeviceAccountPrivateKeyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetDeviceAccountPrivateKeyResponse struct {
	*tchttp.BaseResponse
	Response *ResetDeviceAccountPrivateKeyResponseParams `json:"Response"`
}

func (r *ResetDeviceAccountPrivateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetDeviceAccountPrivateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetUserRequestParams struct {
	// 用户ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

type ResetUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID集合
	IdSet []*uint64 `json:"IdSet,omitnil" name:"IdSet"`
}

func (r *ResetUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetUserResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetUserResponse struct {
	*tchttp.BaseResponse
	Response *ResetUserResponseParams `json:"Response"`
}

func (r *ResetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 服务实例ID，如bh-saas-s3ed4r5e
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 地域编码
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 服务实例规格信息
	SvArgs *string `json:"SvArgs,omitnil" name:"SvArgs"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 服务规格对应的资产数
	Nodes *uint64 `json:"Nodes,omitnil" name:"Nodes"`

	// 自动续费标记，0 - 表示默认状态，1 - 表示自动续费，2 - 表示明确不自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 资源状态，0 - 未初始化，1 - 正常，2 - 隔离，3 - 销毁，4 - 初始化失败，5 - 初始化中
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 服务实例名，如T-Sec-堡垒机（SaaS型）
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 定价模型ID
	Pid *uint64 `json:"Pid,omitnil" name:"Pid"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 商品码, p_cds_dasb
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 子商品码, sp_cds_dasb_bh_saas
	SubProductCode *string `json:"SubProductCode,omitnil" name:"SubProductCode"`

	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 是否过期，true-过期，false-未过期
	Expired *bool `json:"Expired,omitnil" name:"Expired"`

	// 是否开通，true-开通，false-未开通
	Deployed *bool `json:"Deployed,omitnil" name:"Deployed"`

	// 开通服务的 VPC 名称
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 开通服务的 VPC 对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil" name:"VpcCidrBlock"`

	// 开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 外部IP
	PublicIpSet []*string `json:"PublicIpSet,omitnil" name:"PublicIpSet"`

	// 内部IP
	PrivateIpSet []*string `json:"PrivateIpSet,omitnil" name:"PrivateIpSet"`

	// 服务开通的高级功能列表，如:[DB]
	ModuleSet []*string `json:"ModuleSet,omitnil" name:"ModuleSet"`

	// 已使用的授权点数
	UsedNodes *uint64 `json:"UsedNodes,omitnil" name:"UsedNodes"`

	// 扩展点数
	ExtendPoints *uint64 `json:"ExtendPoints,omitnil" name:"ExtendPoints"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *uint64 `json:"PackageBandwidth,omitnil" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *uint64 `json:"PackageNode,omitnil" name:"PackageNode"`

	// 日志投递规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogDeliveryArgs *string `json:"LogDeliveryArgs,omitnil" name:"LogDeliveryArgs"`
}

// Predefined struct for user
type SearchAuditLogRequestParams struct {
	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type SearchAuditLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *SearchAuditLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAuditLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchAuditLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAuditLogResponseParams struct {
	// 审计日志
	AuditLogSet []*AuditLogResult `json:"AuditLogSet,omitnil" name:"AuditLogSet"`

	// 日志总数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchAuditLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchAuditLogResponseParams `json:"Response"`
}

func (r *SearchAuditLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAuditLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchCommandBySidRequestParams struct {
	// 会话Id
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 命令，可模糊搜索
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 根据拦截状态进行过滤
	AuditAction []*uint64 `json:"AuditAction,omitnil" name:"AuditAction"`
}

type SearchCommandBySidRequest struct {
	*tchttp.BaseRequest
	
	// 会话Id
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 命令，可模糊搜索
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 根据拦截状态进行过滤
	AuditAction []*uint64 `json:"AuditAction,omitnil" name:"AuditAction"`
}

func (r *SearchCommandBySidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchCommandBySidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sid")
	delete(f, "Cmd")
	delete(f, "Encoding")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AuditAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchCommandBySidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchCommandBySidResponseParams struct {
	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 命令列表
	CommandSet []*Command `json:"CommandSet,omitnil" name:"CommandSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchCommandBySidResponse struct {
	*tchttp.BaseResponse
	Response *SearchCommandBySidResponseParams `json:"Response"`
}

func (r *SearchCommandBySidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchCommandBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchCommandRequestParams struct {
	// 搜索区间的开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 搜索区间的结束时间，不填默认为开始时间到现在为止
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 资产实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 资产的公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 资产的内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 执行的命令
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`

	// 根据拦截状态进行过滤：1 - 已执行，2 - 被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil" name:"AuditAction"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type SearchCommandRequest struct {
	*tchttp.BaseRequest
	
	// 搜索区间的开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 搜索区间的结束时间，不填默认为开始时间到现在为止
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 资产实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 资产的公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 资产的内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 执行的命令
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`

	// 根据拦截状态进行过滤：1 - 已执行，2 - 被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil" name:"AuditAction"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *SearchCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserName")
	delete(f, "RealName")
	delete(f, "InstanceId")
	delete(f, "DeviceName")
	delete(f, "PublicIp")
	delete(f, "PrivateIp")
	delete(f, "Cmd")
	delete(f, "Encoding")
	delete(f, "AuditAction")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchCommandResponseParams struct {
	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 命令列表
	Commands []*SearchCommandResult `json:"Commands,omitnil" name:"Commands"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchCommandResponse struct {
	*tchttp.BaseResponse
	Response *SearchCommandResponseParams `json:"Response"`
}

func (r *SearchCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCommandResult struct {
	// 命令输入的时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 命令
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// 命令执行情况，1--允许，2--拒绝
	Action *uint64 `json:"Action,omitnil" name:"Action"`

	// 命令所属的会话ID
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms
	TimeOffset *uint64 `json:"TimeOffset,omitnil" name:"TimeOffset"`

	// 账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil" name:"Account"`

	// source ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromIp *string `json:"FromIp,omitnil" name:"FromIp"`

	// 该命令所属会话的会话开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionTime *string `json:"SessionTime,omitnil" name:"SessionTime"`

	// 该命令所属会话的会话开始时间（废弃，使用SessionTime）
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SessTime is deprecated.
	SessTime *string `json:"SessTime,omitnil" name:"SessTime"`

	// 复核时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfirmTime *string `json:"ConfirmTime,omitnil" name:"ConfirmTime"`

	// 部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentId *string `json:"UserDepartmentId,omitnil" name:"UserDepartmentId"`

	// 用户部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentName *string `json:"UserDepartmentName,omitnil" name:"UserDepartmentName"`

	// 设备部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitnil" name:"DeviceDepartmentId"`

	// 设备部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitnil" name:"DeviceDepartmentName"`
}

// Predefined struct for user
type SearchFileBySidRequestParams struct {
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 是否创建审计日志,通过查看按钮调用时为true,其他为false
	AuditLog *bool `json:"AuditLog,omitnil" name:"AuditLog"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 可填写路径名或文件名
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 分页用偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 1-已执行，  2-被阻断
	AuditAction *int64 `json:"AuditAction,omitnil" name:"AuditAction"`

	// 以Protocol和Method为条件查询
	TypeFilters []*SearchFileTypeFilter `json:"TypeFilters,omitnil" name:"TypeFilters"`
}

type SearchFileBySidRequest struct {
	*tchttp.BaseRequest
	
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Sid *string `json:"Sid,omitnil" name:"Sid"`

	// 是否创建审计日志,通过查看按钮调用时为true,其他为false
	AuditLog *bool `json:"AuditLog,omitnil" name:"AuditLog"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 可填写路径名或文件名
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 分页用偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 1-已执行，  2-被阻断
	AuditAction *int64 `json:"AuditAction,omitnil" name:"AuditAction"`

	// 以Protocol和Method为条件查询
	TypeFilters []*SearchFileTypeFilter `json:"TypeFilters,omitnil" name:"TypeFilters"`
}

func (r *SearchFileBySidRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFileBySidRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sid")
	delete(f, "AuditLog")
	delete(f, "Limit")
	delete(f, "FileName")
	delete(f, "Offset")
	delete(f, "AuditAction")
	delete(f, "TypeFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFileBySidRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFileBySidResponseParams struct {
	// 记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 某会话的文件操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchFileBySidResult []*SearchFileBySidResult `json:"SearchFileBySidResult,omitnil" name:"SearchFileBySidResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchFileBySidResponse struct {
	*tchttp.BaseResponse
	Response *SearchFileBySidResponseParams `json:"Response"`
}

func (r *SearchFileBySidResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFileBySidResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileBySidResult struct {
	// 文件操作时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 1-上传文件 2-下载文件 3-删除文件 4-移动文件 5-重命名文件 6-新建文件夹 7-移动文件夹 8-重命名文件夹 9-删除文件夹
	Method *int64 `json:"Method,omitnil" name:"Method"`

	// 文件传输协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// method为上传、下载、删除时文件在服务器上的位置, 或重命名、移动文件前文件的位置
	FileCurr *string `json:"FileCurr,omitnil" name:"FileCurr"`

	// method为重命名、移动文件时代表移动后的新位置.其他情况为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNew *string `json:"FileNew,omitnil" name:"FileNew"`

	// method为上传文件、下载文件、删除文件时显示文件大小。其他情况为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil" name:"Size"`

	// 堡垒机拦截情况, 1-已执行，  2-被阻断
	Action *int64 `json:"Action,omitnil" name:"Action"`
}

// Predefined struct for user
type SearchFileRequestParams struct {
	// 查询开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹
	Method []*uint64 `json:"Method,omitnil" name:"Method"`

	// 可填写路径名或文件（夹）名
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 1-已执行，  2-被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil" name:"AuditAction"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type SearchFileRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹
	Method []*uint64 `json:"Method,omitnil" name:"Method"`

	// 可填写路径名或文件（夹）名
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 1-已执行，  2-被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil" name:"AuditAction"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *SearchFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UserName")
	delete(f, "RealName")
	delete(f, "InstanceId")
	delete(f, "DeviceName")
	delete(f, "PublicIp")
	delete(f, "PrivateIp")
	delete(f, "Method")
	delete(f, "FileName")
	delete(f, "AuditAction")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchFileResponseParams struct {
	// 记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 文件操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Files []*SearchFileResult `json:"Files,omitnil" name:"Files"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchFileResponse struct {
	*tchttp.BaseResponse
	Response *SearchFileResponseParams `json:"Response"`
}

func (r *SearchFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchFileResult struct {
	// 文件传输的时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 操作结果：1 - 已执行，2 - 已阻断
	Action *uint64 `json:"Action,omitnil" name:"Action"`

	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹
	Method *uint64 `json:"Method,omitnil" name:"Method"`

	// 下载的文件（夹）路径及名称
	FileCurr *string `json:"FileCurr,omitnil" name:"FileCurr"`

	// 上传或新建文件（夹）路径及名称
	FileNew *string `json:"FileNew,omitnil" name:"FileNew"`
}

type SearchFileTypeFilter struct {
	// 需要查询的文件传输类型，如SFTP/CLIP/RDP/RZSZ
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 在当前指定的protocol下进一步过滤具体操作类型,如剪贴板文件上传，剪贴板文件下载等
	Method []*int64 `json:"Method,omitnil" name:"Method"`
}

// Predefined struct for user
type SearchSessionCommandRequestParams struct {
	// 检索的目标命令，为模糊搜索
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认值为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Cmd字段前端是否做base64加密
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type SearchSessionCommandRequest struct {
	*tchttp.BaseRequest
	
	// 检索的目标命令，为模糊搜索
	Cmd *string `json:"Cmd,omitnil" name:"Cmd"`

	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 默认值为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Cmd字段前端是否做base64加密
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil" name:"Encoding"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *SearchSessionCommandRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchSessionCommandRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Cmd")
	delete(f, "StartTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Encoding")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchSessionCommandRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchSessionCommandResponseParams struct {
	// 记录总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 命令和所属会话
	CommandSessionSet []*SessionCommand `json:"CommandSessionSet,omitnil" name:"CommandSessionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchSessionCommandResponse struct {
	*tchttp.BaseResponse
	Response *SearchSessionCommandResponseParams `json:"Response"`
}

func (r *SearchSessionCommandResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchSessionCommandResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchSessionRequestParams struct {
	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 用户名，长度不超过20
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 账号，长度不超过64
	Account *string `json:"Account,omitnil" name:"Account"`

	// 来源Ip
	FromIp *string `json:"FromIp,omitnil" name:"FromIp"`

	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 搜索区间的结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 会话协议类型，只能是1、2、3或4 对应关系为1-tui 2-gui 3-file 4-数据库。若入参是Id，则非必传，否则为必传。
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 姓名，长度不超过20
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 主机名，长度不超过64
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 状态，1为活跃，2为结束，3为强制离线，4为其他错误
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Id *string `json:"Id,omitnil" name:"Id"`
}

type SearchSessionRequest struct {
	*tchttp.BaseRequest
	
	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 用户名，长度不超过20
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 账号，长度不超过64
	Account *string `json:"Account,omitnil" name:"Account"`

	// 来源Ip
	FromIp *string `json:"FromIp,omitnil" name:"FromIp"`

	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 搜索区间的结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 会话协议类型，只能是1、2、3或4 对应关系为1-tui 2-gui 3-file 4-数据库。若入参是Id，则非必传，否则为必传。
	Kind *uint64 `json:"Kind,omitnil" name:"Kind"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 姓名，长度不超过20
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 主机名，长度不超过64
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 状态，1为活跃，2为结束，3为强制离线，4为其他错误
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *SearchSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PrivateIp")
	delete(f, "PublicIp")
	delete(f, "UserName")
	delete(f, "Account")
	delete(f, "FromIp")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Kind")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RealName")
	delete(f, "DeviceName")
	delete(f, "Status")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchSessionResponseParams struct {
	// 记录数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 会话信息列表
	SessionSet []*SessionResult `json:"SessionSet,omitnil" name:"SessionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchSessionResponse struct {
	*tchttp.BaseResponse
	Response *SearchSessionResponseParams `json:"Response"`
}

func (r *SearchSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SessionCommand struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 账号
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 命令列表
	Commands []*Command `json:"Commands,omitnil" name:"Commands"`

	// 记录数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 会话Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 设备id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 设备所属的地域
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`
}

type SessionResult struct {
	// 用户名
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 主机账号
	Account *string `json:"Account,omitnil" name:"Account"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 会话大小
	Size *uint64 `json:"Size,omitnil" name:"Size"`

	// 设备ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil" name:"DeviceName"`

	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 来源Ip
	FromIp *string `json:"FromIp,omitnil" name:"FromIp"`

	// 会话持续时长
	Duration *float64 `json:"Duration,omitnil" name:"Duration"`

	// 该会话内命令数量 ，搜索图形会话时该字段无意义
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 该会话内高危命令数，搜索图形时该字段无意义
	DangerCount *uint64 `json:"DangerCount,omitnil" name:"DangerCount"`

	// 会话状态，如1会话活跃  2会话结束  3强制离线  4其他错误
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 会话Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 设备所属的地域
	ApCode *string `json:"ApCode,omitnil" name:"ApCode"`

	// 会话协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitnil" name:"TagValue"`
}

type User struct {
	// 用户名, 3-20个字符 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitnil" name:"UserName"`

	// 用户姓名， 最大20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitnil" name:"RealName"`

	// 用户ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 手机号码， 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil" name:"ValidateTo"`

	// 所属用户组列表
	GroupSet []*Group `json:"GroupSet,omitnil" name:"GroupSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth
	AuthType *uint64 `json:"AuthType,omitnil" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil" name:"ValidateTime"`

	// 用户所属部门（用于出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil" name:"Department"`

	// 用户所属部门（用于入参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil" name:"DepartmentId"`

	// 激活状态 0 - 未激活 1 - 激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveStatus *uint64 `json:"ActiveStatus,omitnil" name:"ActiveStatus"`

	// 锁定状态 0 - 未锁定 1 - 锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockStatus *uint64 `json:"LockStatus,omitnil" name:"LockStatus"`

	// 状态 与Filter中一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}