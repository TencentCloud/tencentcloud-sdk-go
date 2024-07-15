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
	// 模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 模板描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type Acl struct {
	// 访问权限ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 访问权限名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil,omitempty" name:"AllowDiskRedirect"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（目前含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（目前含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，暂未启用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，暂未启用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil,omitempty" name:"MaxFileDownSize"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil,omitempty" name:"AllowAnyAccount"`

	// 关联的用户列表
	UserSet []*User `json:"UserSet,omitnil,omitempty" name:"UserSet"`

	// 关联的用户组列表
	UserGroupSet []*Group `json:"UserGroupSet,omitnil,omitempty" name:"UserGroupSet"`

	// 关联的资产列表
	DeviceSet []*Device `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 关联的资产组列表
	DeviceGroupSet []*Group `json:"DeviceGroupSet,omitnil,omitempty" name:"DeviceGroupSet"`

	// 关联的账号列表
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 关联的高危命令模板列表
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitnil,omitempty" name:"CmdTemplateSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil,omitempty" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil,omitempty" name:"AllowDiskFileDown"`

	// 是否开启 rz sz 文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil,omitempty" name:"AllowShellFileUp"`

	// 是否开启 rz sz 文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属部门的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 是否允许使用访问串，默认允许
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil,omitempty" name:"AllowAccessCredential"`

	// 关联的数据库高危命令列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ACTemplateSet []*ACTemplate `json:"ACTemplateSet,omitnil,omitempty" name:"ACTemplateSet"`

	// 关联的白命令命令
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhiteCmds []*string `json:"WhiteCmds,omitnil,omitempty" name:"WhiteCmds"`
}

// Predefined struct for user
type AddDeviceGroupMembersRequestParams struct {
	// 资产组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 需要添加到资产组的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
}

type AddDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 需要添加到资产组的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
}

type AddUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 上一次同步的结果。 0 - 从未进行, 1 - 成功， 2 - 失败
	LastStatus *uint64 `json:"LastStatus,omitnil,omitempty" name:"LastStatus"`

	// 同步任务是否正在进行中
	InProcess *bool `json:"InProcess,omitnil,omitempty" name:"InProcess"`
}

type AuditLogResult struct {
	// 被审计会话的Sid
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 审计者的编号
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 审计动作发生的时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 审计者的Ip
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 审计动作类型，1--回放、2--中断、3--监控
	Operation *int64 `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 被审计主机的Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 被审计主机的主机名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 被审计会话所属的类型，如字符会话
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 被审计主机的内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 被审计主机的外部Ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 审计者的子账号
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`
}

// Predefined struct for user
type BindDeviceAccountPasswordRequestParams struct {
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 主机账号密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type BindDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 主机账号密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 主机账号私钥，最新长度128字节，最大长度8192字节
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 主机账号私钥口令，最大长度256字节
	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitnil,omitempty" name:"PrivateKeyPassword"`
}

type BindDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 主机账号私钥，最新长度128字节，最大长度8192字节
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 主机账号私钥口令，最大长度256字节
	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitnil,omitempty" name:"PrivateKeyPassword"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 堡垒机服务ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 网络域ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type BindDeviceResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 堡垒机服务ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 网络域ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindDeviceResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindDeviceResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ChangePwdTaskDetail struct {
	// 资产信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// 资产账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 上次改密结果。0-未改密  1-改密成功 2-改密失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastChangeStatus *uint64 `json:"LastChangeStatus,omitnil,omitempty" name:"LastChangeStatus"`
}

type ChangePwdTaskInfo struct {
	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属部门信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 改密方式。1：使用执行账号。2：修改自身密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChangeMethod *uint64 `json:"ChangeMethod,omitnil,omitempty" name:"ChangeMethod"`

	// 执行账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunAccount *string `json:"RunAccount,omitnil,omitempty" name:"RunAccount"`

	// 密码生成策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthGenerationStrategy *uint64 `json:"AuthGenerationStrategy,omitnil,omitempty" name:"AuthGenerationStrategy"`

	// 密码长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PasswordLength *uint64 `json:"PasswordLength,omitnil,omitempty" name:"PasswordLength"`

	// 包含小写字母
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmallLetter *uint64 `json:"SmallLetter,omitnil,omitempty" name:"SmallLetter"`

	// 包含大写字母
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigLetter *uint64 `json:"BigLetter,omitnil,omitempty" name:"BigLetter"`

	// 包含数字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Digit *uint64 `json:"Digit,omitnil,omitempty" name:"Digit"`

	// 包含的特殊字符，base64
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// 改密完成通知。0-通知，1-不通知
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompleteNotify *uint64 `json:"CompleteNotify,omitnil,omitempty" name:"CompleteNotify"`

	// 通知人邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyEmails []*string `json:"NotifyEmails,omitnil,omitempty" name:"NotifyEmails"`

	// 加密附件密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilePassword *string `json:"FilePassword,omitnil,omitempty" name:"FilePassword"`

	// 需要改密的账户
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 需要改密的主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceSet []*Device `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 任务类型：4手动，5自动
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 首次执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`

	// 下次执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextTime *string `json:"NextTime,omitnil,omitempty" name:"NextTime"`
}

type Clb struct {
	// 负载均衡IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbIp *string `json:"ClbIp,omitnil,omitempty" name:"ClbIp"`
}

type CmdTemplate struct {
	// 高危命令模板ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 高危命令模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命令列表，命令之间用换行符（"\n"）分隔
	CmdList *string `json:"CmdList,omitnil,omitempty" name:"CmdList"`
}

type Command struct {
	// 命令
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 命令输入的时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms
	TimeOffset *uint64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 命令执行情况，1--允许，2--拒绝，3--确认
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 会话id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 设备account
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 设备ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// source ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromIp *string `json:"FromIp,omitnil,omitempty" name:"FromIp"`

	// 该命令所属会话的会话开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionTime *string `json:"SessionTime,omitnil,omitempty" name:"SessionTime"`

	// 该命令所属会话的会话开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SessTime is deprecated.
	SessTime *string `json:"SessTime,omitnil,omitempty" name:"SessTime"`

	// 复核时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfirmTime *string `json:"ConfirmTime,omitnil,omitempty" name:"ConfirmTime"`

	// 用户部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentId *string `json:"UserDepartmentId,omitnil,omitempty" name:"UserDepartmentId"`

	// 用户部门name
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentName *string `json:"UserDepartmentName,omitnil,omitempty" name:"UserDepartmentName"`

	// 设备部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitnil,omitempty" name:"DeviceDepartmentId"`

	// 设备部门name
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitnil,omitempty" name:"DeviceDepartmentName"`

	// 会话大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil,omitempty" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil,omitempty" name:"AllowClipTextDown"`

	// 是否开启 SFTP 文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil,omitempty" name:"MaxFileUpSize"`

	// 是否开启 SFTP 文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID集合
	UserIdSet []*uint64 `json:"UserIdSet,omitnil,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil,omitempty" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil,omitempty" name:"ACTemplateIdSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil,omitempty" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 访问权限所属部门的ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 是否允许使用访问串，默认允许
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil,omitempty" name:"AllowAccessCredential"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil,omitempty" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil,omitempty" name:"AllowClipTextDown"`

	// 是否开启 SFTP 文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil,omitempty" name:"MaxFileUpSize"`

	// 是否开启 SFTP 文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID集合
	UserIdSet []*uint64 `json:"UserIdSet,omitnil,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil,omitempty" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil,omitempty" name:"ACTemplateIdSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil,omitempty" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 访问权限所属部门的ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 是否允许使用访问串，默认允许
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil,omitempty" name:"AllowAccessCredential"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Category *uint64 `json:"Category,omitnil,omitempty" name:"Category"`
}

type CreateAssetSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步资产类别，1 - 主机资产, 2 - 数据库资产
	Category *uint64 `json:"Category,omitnil,omitempty" name:"Category"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateChangePwdTaskRequestParams struct {
	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 资产id数组
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 修改的账户数组
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 改密方式。1：使用执行账号修改密码；2：修改自身密码
	ChangeMethod *int64 `json:"ChangeMethod,omitnil,omitempty" name:"ChangeMethod"`

	// 认证生成方式。 1:自动生成相同密码 2:自动生成不同密码 3:手动指定相同密码
	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitnil,omitempty" name:"AuthGenerationStrategy"`

	// 执行账号
	RunAccount *string `json:"RunAccount,omitnil,omitempty" name:"RunAccount"`

	// 手动指定密码时必传
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密码限制长度，长度大于 12 位
	PasswordLength *int64 `json:"PasswordLength,omitnil,omitempty" name:"PasswordLength"`

	// 密码包含小写字母。0：否，1：是
	SmallLetter *int64 `json:"SmallLetter,omitnil,omitempty" name:"SmallLetter"`

	// 密码包含大写字母。0：否，1：是
	BigLetter *int64 `json:"BigLetter,omitnil,omitempty" name:"BigLetter"`

	// 密码包含数字。0：否，1：是
	Digit *int64 `json:"Digit,omitnil,omitempty" name:"Digit"`

	// 密码包含的特殊字符（base64编码），包含：^[-_#();%~!+=]*$
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// 改密完成通知。0：不通知 
	//   1：通知
	CompleteNotify *int64 `json:"CompleteNotify,omitnil,omitempty" name:"CompleteNotify"`

	// 通知邮箱
	NotifyEmails []*string `json:"NotifyEmails,omitnil,omitempty" name:"NotifyEmails"`

	// 加密压缩文件密码
	FilePassword *string `json:"FilePassword,omitnil,omitempty" name:"FilePassword"`

	// 所属部门id。“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 任务类型  4-手工执行  5-周期自动执行
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 执行周期，单位天（大于等于 1，小于等于 365）
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 周期任务首次执行时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`
}

type CreateChangePwdTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 资产id数组
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 修改的账户数组
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 改密方式。1：使用执行账号修改密码；2：修改自身密码
	ChangeMethod *int64 `json:"ChangeMethod,omitnil,omitempty" name:"ChangeMethod"`

	// 认证生成方式。 1:自动生成相同密码 2:自动生成不同密码 3:手动指定相同密码
	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitnil,omitempty" name:"AuthGenerationStrategy"`

	// 执行账号
	RunAccount *string `json:"RunAccount,omitnil,omitempty" name:"RunAccount"`

	// 手动指定密码时必传
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密码限制长度，长度大于 12 位
	PasswordLength *int64 `json:"PasswordLength,omitnil,omitempty" name:"PasswordLength"`

	// 密码包含小写字母。0：否，1：是
	SmallLetter *int64 `json:"SmallLetter,omitnil,omitempty" name:"SmallLetter"`

	// 密码包含大写字母。0：否，1：是
	BigLetter *int64 `json:"BigLetter,omitnil,omitempty" name:"BigLetter"`

	// 密码包含数字。0：否，1：是
	Digit *int64 `json:"Digit,omitnil,omitempty" name:"Digit"`

	// 密码包含的特殊字符（base64编码），包含：^[-_#();%~!+=]*$
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// 改密完成通知。0：不通知 
	//   1：通知
	CompleteNotify *int64 `json:"CompleteNotify,omitnil,omitempty" name:"CompleteNotify"`

	// 通知邮箱
	NotifyEmails []*string `json:"NotifyEmails,omitnil,omitempty" name:"NotifyEmails"`

	// 加密压缩文件密码
	FilePassword *string `json:"FilePassword,omitnil,omitempty" name:"FilePassword"`

	// 所属部门id。“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 任务类型  4-手工执行  5-周期自动执行
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 执行周期，单位天（大于等于 1，小于等于 365）
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 周期任务首次执行时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`
}

func (r *CreateChangePwdTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChangePwdTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "DeviceIdSet")
	delete(f, "AccountSet")
	delete(f, "ChangeMethod")
	delete(f, "AuthGenerationStrategy")
	delete(f, "RunAccount")
	delete(f, "Password")
	delete(f, "PasswordLength")
	delete(f, "SmallLetter")
	delete(f, "BigLetter")
	delete(f, "Digit")
	delete(f, "Symbol")
	delete(f, "CompleteNotify")
	delete(f, "NotifyEmails")
	delete(f, "FilePassword")
	delete(f, "DepartmentId")
	delete(f, "Type")
	delete(f, "Period")
	delete(f, "FirstTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChangePwdTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChangePwdTaskResponseParams struct {
	// 任务id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChangePwdTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateChangePwdTaskResponseParams `json:"Response"`
}

func (r *CreateChangePwdTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCmdTemplateRequestParams struct {
	// 模板名，最大长度32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命令列表，\n分隔，最大长度32768字节
	CmdList *string `json:"CmdList,omitnil,omitempty" name:"CmdList"`

	// 标识cmdlist字段前端是否为base64加密传值.
	// 0:表示非base64加密
	// 1:表示是base64加密
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`
}

type CreateCmdTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名，最大长度32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命令列表，\n分隔，最大长度32768字节
	CmdList *string `json:"CmdList,omitnil,omitempty" name:"CmdList"`

	// 标识cmdlist字段前端是否为base64加密传值.
	// 0:表示非base64加密
	// 1:表示是base64加密
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeviceId *uint64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`
}

type CreateDeviceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 主机记录ID
	DeviceId *uint64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type CreateDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组名，最大长度32字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeployRegion *string `json:"DeployRegion,omitnil,omitempty" name:"DeployRegion"`

	// 部署堡垒机的VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 部署堡垒机的SubnetId
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 资源类型。取值:standard/pro
	ResourceEdition *string `json:"ResourceEdition,omitnil,omitempty" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil,omitempty" name:"ResourceNode"`

	// 计费周期
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 计费时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 计费模式 1预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 部署zone
	DeployZone *string `json:"DeployZone,omitnil,omitempty" name:"DeployZone"`
}

type CreateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 部署region
	DeployRegion *string `json:"DeployRegion,omitnil,omitempty" name:"DeployRegion"`

	// 部署堡垒机的VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 部署堡垒机的SubnetId
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 资源类型。取值:standard/pro
	ResourceEdition *string `json:"ResourceEdition,omitnil,omitempty" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil,omitempty" name:"ResourceNode"`

	// 计费周期
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 计费时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 计费模式 1预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 部署zone
	DeployZone *string `json:"DeployZone,omitnil,omitempty" name:"DeployZone"`
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
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户组所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组名，最大长度32字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户组所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户姓名，最大长度20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 按照"国家地区代码|手机号"的格式输入。如: "+86|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil,omitempty" name:"ValidateTime"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户名, 3-20个字符, 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户姓名，最大长度20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 按照"国家地区代码|手机号"的格式输入。如: "+86|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil,omitempty" name:"ValidateTime"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteAclsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的权限ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteChangePwdTaskRequestParams struct {
	// 改密任务id列表
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteChangePwdTaskRequest struct {
	*tchttp.BaseRequest
	
	// 改密任务id列表
	IdSet []*int64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

func (r *DeleteChangePwdTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChangePwdTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteChangePwdTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteChangePwdTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteChangePwdTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteChangePwdTaskResponseParams `json:"Response"`
}

func (r *DeleteChangePwdTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCmdTemplatesRequestParams struct {
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteCmdTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteDeviceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 需要删除的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
}

type DeleteDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 需要删除的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 需删除的成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
}

type DeleteUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 需删除的成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的用户ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 部门名称，1 - 256个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 部门管理员账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Managers []*string `json:"Managers,omitnil,omitempty" name:"Managers"`

	// 管理员用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerUsers []*DepartmentManagerUser `json:"ManagerUsers,omitnil,omitempty" name:"ManagerUsers"`
}

type DepartmentManagerUser struct {
	// 管理员Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerId *string `json:"ManagerId,omitnil,omitempty" name:"ManagerId"`

	// 管理员姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagerName *string `json:"ManagerName,omitnil,omitempty" name:"ManagerName"`
}

// Predefined struct for user
type DeployResourceRequestParams struct {
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 需要开通服务的地域
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 需要开通服务的VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 需要开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 需要开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 需要开通服务的VPC名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 需要开通服务的VPC对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 需要开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 需要开通实例所属的CDC集群ID
	CdcClusterId *string `json:"CdcClusterId,omitnil,omitempty" name:"CdcClusterId"`
}

type DeployResourceRequest struct {
	*tchttp.BaseRequest
	
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 需要开通服务的地域
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 需要开通服务的VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 需要开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 需要开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 需要开通服务的VPC名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 需要开通服务的VPC对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 需要开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 需要开通实例所属的CDC集群ID
	CdcClusterId *string `json:"CdcClusterId,omitnil,omitempty" name:"CdcClusterId"`
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
	delete(f, "CdcClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeployResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeployResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否根据Name进行精确查询，默认值false
	Exact *bool `json:"Exact,omitnil,omitempty" name:"Exact"`

	// 有访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil,omitempty" name:"AuthorizedUserIdSet"`

	// 有访问权限的资产ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil,omitempty" name:"AuthorizedDeviceIdSet"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 部门ID，用于过滤属于某个部门的访问权限
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type DescribeAclsRequest struct {
	*tchttp.BaseRequest
	
	// 访问权限ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否根据Name进行精确查询，默认值false
	Exact *bool `json:"Exact,omitnil,omitempty" name:"Exact"`

	// 有访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil,omitempty" name:"AuthorizedUserIdSet"`

	// 有访问权限的资产ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil,omitempty" name:"AuthorizedDeviceIdSet"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 部门ID，用于过滤属于某个部门的访问权限
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 访问权限列表
	AclSet []*Acl `json:"AclSet,omitnil,omitempty" name:"AclSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Category *uint64 `json:"Category,omitnil,omitempty" name:"Category"`
}

type DescribeAssetSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// 查询的资产同步类型。1 -主机资产， 2 - 数据库资产
	Category *uint64 `json:"Category,omitnil,omitempty" name:"Category"`
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
	Status *AssetSyncStatus `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeChangePwdTaskDetailRequestParams struct {
	// 改密任务Id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 过滤数组，支持：InstanceId 资产ID，DeviceName 资产名称，Ip 内外IP，Account 资产账号，LastChangeStatus 上次改密状态。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移位置，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目。默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeChangePwdTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 改密任务Id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 过滤数组，支持：InstanceId 资产ID，DeviceName 资产名称，Ip 内外IP，Account 资产账号，LastChangeStatus 上次改密状态。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移位置，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目。默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeChangePwdTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangePwdTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationId")
	delete(f, "DepartmentId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChangePwdTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChangePwdTaskDetailResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Details []*ChangePwdTaskDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChangePwdTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChangePwdTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeChangePwdTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangePwdTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChangePwdTaskRequestParams struct {
	// 过滤数组。过滤数组。Name支持以下值: OperationId 任务ID TaskName 任务名
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 分页偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeChangePwdTaskRequest struct {
	*tchttp.BaseRequest
	
	// 过滤数组。过滤数组。Name支持以下值: OperationId 任务ID TaskName 任务名
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 分页偏移量，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeChangePwdTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangePwdTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "DepartmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChangePwdTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChangePwdTaskResponseParams struct {
	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ChangePwdTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeChangePwdTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChangePwdTaskResponseParams `json:"Response"`
}

func (r *DescribeChangePwdTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmdTemplatesRequestParams struct {
	// 命令模板ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 命令模板名，模糊查询，最大长度64字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeCmdTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 命令模板ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 命令模板名，模糊查询，最大长度64字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitnil,omitempty" name:"CmdTemplateSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BaseImageId *string `json:"BaseImageId,omitnil,omitempty" name:"BaseImageId"`

	// AI镜像ID
	AiImageId *string `json:"AiImageId,omitnil,omitempty" name:"AiImageId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机ID，未使用IdSet时必须携带
	DeviceId *uint64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDeviceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID集合，非必需，如果使用IdSet则忽略其他过滤参数
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 主机ID，未使用IdSet时必须携带
	DeviceId *uint64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 账号信息列表
	DeviceAccountSet []*DeviceAccount `json:"DeviceAccountSet,omitnil,omitempty" name:"DeviceAccountSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产
	Bound *bool `json:"Bound,omitnil,omitempty" name:"Bound"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产
	Bound *bool `json:"Bound,omitnil,omitempty" name:"Bound"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产组成员列表
	DeviceSet []*Device `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 资产组名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的资产组
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type DescribeDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 资产组名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的资产组
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产组列表
	GroupSet []*Group `json:"GroupSet,omitnil,omitempty" name:"GroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitnil,omitempty" name:"ApCodeSet"`

	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 有该资产访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil,omitempty" name:"AuthorizedUserIdSet"`

	// 过滤条件，资产绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitnil,omitempty" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	KindSet []*uint64 `json:"KindSet,omitnil,omitempty" name:"KindSet"`

	// 资产是否包含托管账号。1，包含；0，不包含
	ManagedAccount *string `json:"ManagedAccount,omitnil,omitempty" name:"ManagedAccount"`

	// 过滤条件，可按照部门ID进行过滤
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 过滤数组。支持的Name：
	// BindingStatus 绑定状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitnil,omitempty" name:"ApCodeSet"`

	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 有该资产访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitnil,omitempty" name:"AuthorizedUserIdSet"`

	// 过滤条件，资产绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitnil,omitempty" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	KindSet []*uint64 `json:"KindSet,omitnil,omitempty" name:"KindSet"`

	// 资产是否包含托管账号。1，包含；0，不包含
	ManagedAccount *string `json:"ManagedAccount,omitnil,omitempty" name:"ManagedAccount"`

	// 过滤条件，可按照部门ID进行过滤
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 过滤数组。支持的Name：
	// BindingStatus 绑定状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资产信息列表
	DeviceSet []*Device `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API
	Entry *uint64 `json:"Entry,omitnil,omitempty" name:"Entry"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeLoginEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API
	Entry *uint64 `json:"Entry,omitnil,omitempty" name:"Entry"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	LoginEventSet []*LoginEvent `json:"LoginEventSet,omitnil,omitempty" name:"LoginEventSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 操作类型，参考DescribeOperationType返回结果
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeOperationEventRequest struct {
	*tchttp.BaseRequest
	
	// 用户名，如果不包含其他条件时对user_name or real_name两个字段模糊查询
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名，模糊查询
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 查询时间范围，起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询时间范围，结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 来源IP，模糊查询
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 操作类型，参考DescribeOperationType返回结果
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页每页记录数，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	OperationEventSet []*OperationEvent `json:"OperationEventSet,omitnil,omitempty" name:"OperationEventSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 按照堡垒机开通的 VPC 实例ID查询
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 每页条目数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 地域码, 如: ap-guangzhou
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 按照堡垒机开通的 VPC 实例ID查询
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 每页条目数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesResponseParams struct {
	// 堡垒机资源列表
	ResourceSet []*Resource `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// 堡垒机资源数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户
	Bound *bool `json:"Bound,omitnil,omitempty" name:"Bound"`

	// 用户名或用户姓名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type DescribeUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户
	Bound *bool `json:"Bound,omitnil,omitempty" name:"Bound"`

	// 用户名或用户姓名，最长64个字符，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 用户组成员列表
	UserSet []*User `json:"UserSet,omitnil,omitempty" name:"UserSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 用户组名，模糊查询,长度：0-64字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的用户组
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type DescribeUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 用户组名，模糊查询,长度：0-64字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的用户组
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 用户组列表
	GroupSet []*Group `json:"GroupSet,omitnil,omitempty" name:"GroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱，精确查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 查询具有指定资产ID访问权限的用户
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil,omitempty" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitnil,omitempty" name:"AuthTypeSet"`

	// 部门ID，用于过滤属于某个部门的用户
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 参数过滤数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 如果IdSet不为空，则忽略其他参数
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱，精确查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 查询具有指定资产ID访问权限的用户
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitnil,omitempty" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitnil,omitempty" name:"AuthTypeSet"`

	// 部门ID，用于过滤属于某个部门的用户
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 参数过滤数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 用户列表
	UserSet []*User `json:"UserSet,omitnil,omitempty" name:"UserSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 实例ID，对应CVM、CDB等实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 地域编码
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 操作系统名称
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// 资产类型 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 所属资产组列表
	GroupSet []*Group `json:"GroupSet,omitnil,omitempty" name:"GroupSet"`

	// 资产绑定的账号数
	AccountCount *uint64 `json:"AccountCount,omitnil,omitempty" name:"AccountCount"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 堡垒机服务信息，注意没有绑定服务时为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 资产所属部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 数据库资产的多节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpPortSet []*string `json:"IpPortSet,omitnil,omitempty" name:"IpPortSet"`

	// 网络域Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 网络域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`
}

type DeviceAccount struct {
	// 账号ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 主机ID
	DeviceId *uint64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// true-已托管密码，false-未托管密码
	BoundPassword *bool `json:"BoundPassword,omitnil,omitempty" name:"BoundPassword"`

	// true-已托管私钥，false-未托管私钥
	BoundPrivateKey *bool `json:"BoundPrivateKey,omitnil,omitempty" name:"BoundPrivateKey"`
}

type ExternalDevice struct {
	// 操作系统名称，只能是Linux、Windows或MySQL
	OsName *string `json:"OsName,omitnil,omitempty" name:"OsName"`

	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 主机名，可为空
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产所属的部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 资产多节点：字段ip和端口
	IpPortSet []*string `json:"IpPortSet,omitnil,omitempty" name:"IpPortSet"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Group struct {
	// 组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属部门信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

// Predefined struct for user
type ImportExternalDeviceRequestParams struct {
	// 资产参数列表
	DeviceSet []*ExternalDevice `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`
}

type ImportExternalDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 资产参数列表
	DeviceSet []*ExternalDevice `json:"DeviceSet,omitnil,omitempty" name:"DeviceSet"`
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
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 操作时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 来源IP
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 登录入口：1-字符界面,2-图形界面，3-web页面, 4-API
	Entry *uint64 `json:"Entry,omitnil,omitempty" name:"Entry"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type ModifyAclRequestParams struct {
	// 访问权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil,omitempty" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitnil,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil,omitempty" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil,omitempty" name:"ACTemplateIdSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil,omitempty" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 权限所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 是否允许使用访问串
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil,omitempty" name:"AllowAccessCredential"`
}

type ModifyAclRequest struct {
	*tchttp.BaseRequest
	
	// 访问权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitnil,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitnil,omitempty" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitnil,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitnil,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitnil,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitnil,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitnil,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitnil,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitnil,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitnil,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitnil,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitnil,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitnil,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitnil,omitempty" name:"CmdTemplateIdSet"`

	// 关联高危DB模板ID
	ACTemplateIdSet []*string `json:"ACTemplateIdSet,omitnil,omitempty" name:"ACTemplateIdSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitnil,omitempty" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitnil,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitnil,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitnil,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitnil,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 权限所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 是否允许使用访问串
	AllowAccessCredential *bool `json:"AllowAccessCredential,omitnil,omitempty" name:"AllowAccessCredential"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyChangePwdTaskRequestParams struct {
	// 改密任务id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 改密资产id列表
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 改密资产的账号列表
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 修改类型：1：修改任务信息  2：关联任务资产账号
	ModifyType *uint64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 改密方式。1：使用执行账号修改密码；2：修改自身密码
	ChangeMethod *int64 `json:"ChangeMethod,omitnil,omitempty" name:"ChangeMethod"`

	// 密码生成方式。 1:自动生成相同密码 2:自动生成不同密码 3:手动指定相同密码
	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitnil,omitempty" name:"AuthGenerationStrategy"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属部门ID，"1,2,3"
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 任务的执行账号	
	RunAccount *string `json:"RunAccount,omitnil,omitempty" name:"RunAccount"`

	// 密码，手动指定密码时必传。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密码限制长度，自动生成密码必传。	
	PasswordLength *int64 `json:"PasswordLength,omitnil,omitempty" name:"PasswordLength"`

	// 密码包含小写字母，0：否，1：是。
	SmallLetter *int64 `json:"SmallLetter,omitnil,omitempty" name:"SmallLetter"`

	// 密码包含大写字母，0：否，1：是。
	BigLetter *int64 `json:"BigLetter,omitnil,omitempty" name:"BigLetter"`

	// 密码包含数字，0：否，1：是。
	Digit *int64 `json:"Digit,omitnil,omitempty" name:"Digit"`

	// 密码包含的特殊字符（base64编码），包含：^[-_#();%~!+=]*$
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// 改密完成通知。0：不通知 1：通知
	CompleteNotify *int64 `json:"CompleteNotify,omitnil,omitempty" name:"CompleteNotify"`

	// 通知邮箱
	NotifyEmails []*string `json:"NotifyEmails,omitnil,omitempty" name:"NotifyEmails"`

	// 加密压缩文件密码
	FilePassword *string `json:"FilePassword,omitnil,omitempty" name:"FilePassword"`

	// 任务类型， 4：手工执行  5：周期自动执行
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 周期任务周期，单位天（大于等于 1，小于等于 365）
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 周期任务首次执行时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`
}

type ModifyChangePwdTaskRequest struct {
	*tchttp.BaseRequest
	
	// 改密任务id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 改密资产id列表
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitnil,omitempty" name:"DeviceIdSet"`

	// 改密资产的账号列表
	AccountSet []*string `json:"AccountSet,omitnil,omitempty" name:"AccountSet"`

	// 修改类型：1：修改任务信息  2：关联任务资产账号
	ModifyType *uint64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 改密方式。1：使用执行账号修改密码；2：修改自身密码
	ChangeMethod *int64 `json:"ChangeMethod,omitnil,omitempty" name:"ChangeMethod"`

	// 密码生成方式。 1:自动生成相同密码 2:自动生成不同密码 3:手动指定相同密码
	AuthGenerationStrategy *int64 `json:"AuthGenerationStrategy,omitnil,omitempty" name:"AuthGenerationStrategy"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属部门ID，"1,2,3"
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 任务的执行账号	
	RunAccount *string `json:"RunAccount,omitnil,omitempty" name:"RunAccount"`

	// 密码，手动指定密码时必传。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密码限制长度，自动生成密码必传。	
	PasswordLength *int64 `json:"PasswordLength,omitnil,omitempty" name:"PasswordLength"`

	// 密码包含小写字母，0：否，1：是。
	SmallLetter *int64 `json:"SmallLetter,omitnil,omitempty" name:"SmallLetter"`

	// 密码包含大写字母，0：否，1：是。
	BigLetter *int64 `json:"BigLetter,omitnil,omitempty" name:"BigLetter"`

	// 密码包含数字，0：否，1：是。
	Digit *int64 `json:"Digit,omitnil,omitempty" name:"Digit"`

	// 密码包含的特殊字符（base64编码），包含：^[-_#();%~!+=]*$
	Symbol *string `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// 改密完成通知。0：不通知 1：通知
	CompleteNotify *int64 `json:"CompleteNotify,omitnil,omitempty" name:"CompleteNotify"`

	// 通知邮箱
	NotifyEmails []*string `json:"NotifyEmails,omitnil,omitempty" name:"NotifyEmails"`

	// 加密压缩文件密码
	FilePassword *string `json:"FilePassword,omitnil,omitempty" name:"FilePassword"`

	// 任务类型， 4：手工执行  5：周期自动执行
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 周期任务周期，单位天（大于等于 1，小于等于 365）
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 周期任务首次执行时间
	FirstTime *string `json:"FirstTime,omitnil,omitempty" name:"FirstTime"`
}

func (r *ModifyChangePwdTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChangePwdTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationId")
	delete(f, "DeviceIdSet")
	delete(f, "AccountSet")
	delete(f, "ModifyType")
	delete(f, "ChangeMethod")
	delete(f, "AuthGenerationStrategy")
	delete(f, "TaskName")
	delete(f, "DepartmentId")
	delete(f, "RunAccount")
	delete(f, "Password")
	delete(f, "PasswordLength")
	delete(f, "SmallLetter")
	delete(f, "BigLetter")
	delete(f, "Digit")
	delete(f, "Symbol")
	delete(f, "CompleteNotify")
	delete(f, "NotifyEmails")
	delete(f, "FilePassword")
	delete(f, "Type")
	delete(f, "Period")
	delete(f, "FirstTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyChangePwdTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyChangePwdTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyChangePwdTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyChangePwdTaskResponseParams `json:"Response"`
}

func (r *ModifyChangePwdTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCmdTemplateRequestParams struct {
	// 模板名，最长32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命令列表，\n分隔，最长32768字节
	CmdList *string `json:"CmdList,omitnil,omitempty" name:"CmdList"`

	// 命令模板ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// CmdList字段前端是否base64传值。
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`
}

type ModifyCmdTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名，最长32字符，不能包含空白字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命令列表，\n分隔，最长32768字节
	CmdList *string `json:"CmdList,omitnil,omitempty" name:"CmdList"`

	// 命令模板ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// CmdList字段前端是否base64传值。
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type ModifyDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组名，最大长度32字符，不能为空
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 资产所属组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 资产所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 网络域Id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 资产记录ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 管理端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 资产所属组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 资产所属部门ID
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 网络域Id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`
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
	delete(f, "DomainId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDeviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDeviceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyOAuthSettingRequestParams struct {
	// 是否开启OAuth认证，false-不开启，true-开启。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// OAuth认证方式。
	AuthMethod *string `json:"AuthMethod,omitnil,omitempty" name:"AuthMethod"`

	// OAuth认证客户端Id
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OAuth认证客户端密钥
	ClientSecret *string `json:"ClientSecret,omitnil,omitempty" name:"ClientSecret"`

	// 获取OAuth认证授权码URL
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`

	// 获取OAuth令牌URL
	TokenUrl *string `json:"TokenUrl,omitnil,omitempty" name:"TokenUrl"`

	// 获取OAuth用户信息URL
	UserInfoUrl *string `json:"UserInfoUrl,omitnil,omitempty" name:"UserInfoUrl"`

	// 使用Okta认证时指定范围。为空时默认使用 openid、profile、email。
	Scopes []*string `json:"Scopes,omitnil,omitempty" name:"Scopes"`
}

type ModifyOAuthSettingRequest struct {
	*tchttp.BaseRequest
	
	// 是否开启OAuth认证，false-不开启，true-开启。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// OAuth认证方式。
	AuthMethod *string `json:"AuthMethod,omitnil,omitempty" name:"AuthMethod"`

	// OAuth认证客户端Id
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OAuth认证客户端密钥
	ClientSecret *string `json:"ClientSecret,omitnil,omitempty" name:"ClientSecret"`

	// 获取OAuth认证授权码URL
	CodeUrl *string `json:"CodeUrl,omitnil,omitempty" name:"CodeUrl"`

	// 获取OAuth令牌URL
	TokenUrl *string `json:"TokenUrl,omitnil,omitempty" name:"TokenUrl"`

	// 获取OAuth用户信息URL
	UserInfoUrl *string `json:"UserInfoUrl,omitnil,omitempty" name:"UserInfoUrl"`

	// 使用Okta认证时指定范围。为空时默认使用 openid、profile、email。
	Scopes []*string `json:"Scopes,omitnil,omitempty" name:"Scopes"`
}

func (r *ModifyOAuthSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOAuthSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "AuthMethod")
	delete(f, "ClientId")
	delete(f, "ClientSecret")
	delete(f, "CodeUrl")
	delete(f, "TokenUrl")
	delete(f, "UserInfoUrl")
	delete(f, "Scopes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOAuthSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOAuthSettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyOAuthSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOAuthSettingResponseParams `json:"Response"`
}

func (r *ModifyOAuthSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOAuthSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceRequestParams struct {
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 已废弃
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 已废弃
	ModuleSet []*string `json:"ModuleSet,omitnil,omitempty" name:"ModuleSet"`

	// 实例版本
	ResourceEdition *string `json:"ResourceEdition,omitnil,omitempty" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil,omitempty" name:"ResourceNode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *int64 `json:"PackageBandwidth,omitnil,omitempty" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *int64 `json:"PackageNode,omitnil,omitempty" name:"PackageNode"`

	// 日志投递
	LogDelivery *int64 `json:"LogDelivery,omitnil,omitempty" name:"LogDelivery"`
}

type ModifyResourceRequest struct {
	*tchttp.BaseRequest
	
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 已废弃
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 已废弃
	ModuleSet []*string `json:"ModuleSet,omitnil,omitempty" name:"ModuleSet"`

	// 实例版本
	ResourceEdition *string `json:"ResourceEdition,omitnil,omitempty" name:"ResourceEdition"`

	// 资源节点数
	ResourceNode *int64 `json:"ResourceNode,omitnil,omitempty" name:"ResourceNode"`

	// 自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *int64 `json:"PackageBandwidth,omitnil,omitempty" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *int64 `json:"PackageNode,omitnil,omitempty" name:"PackageNode"`

	// 日志投递
	LogDelivery *int64 `json:"LogDelivery,omitnil,omitempty" name:"LogDelivery"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyUserGroupRequestParams struct {
	// 用户组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户组名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户组所属的部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type ModifyUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户组名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户组所属的部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

func (r *ModifyUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserGroupResponseParams `json:"Response"`
}

func (r *ModifyUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 用户ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户姓名，最大长度20个字符，不能包含空格
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 按照"国家地区代码|手机号"的格式输入。如: "+86|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil,omitempty" name:"ValidateTime"`

	// 用户所属部门的ID，如1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户姓名，最大长度20个字符，不能包含空格
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 按照"国家地区代码|手机号"的格式输入。如: "+86|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitnil,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil,omitempty" name:"ValidateTime"`

	// 用户所属部门的ID，如1.2.3
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 操作时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 来源IP
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 操作类型
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 具体操作内容
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 操作结果，1-成功，2-失败
	Result *uint64 `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type ResetDeviceAccountPasswordRequestParams struct {
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type ResetDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type ResetDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
}

type ResetUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID集合
	IdSet []*uint64 `json:"IdSet,omitnil,omitempty" name:"IdSet"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 地域编码
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 服务实例规格信息
	SvArgs *string `json:"SvArgs,omitnil,omitempty" name:"SvArgs"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 服务规格对应的资产数
	Nodes *uint64 `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 自动续费标记，0 - 表示默认状态，1 - 表示自动续费，2 - 表示明确不自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 资源状态，0 - 未初始化，1 - 正常，2 - 隔离，3 - 销毁，4 - 初始化失败，5 - 初始化中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 服务实例名，如T-Sec-堡垒机（SaaS型）
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 定价模型ID
	Pid *uint64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 商品码, p_cds_dasb
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子商品码, sp_cds_dasb_bh_saas
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 是否过期，true-过期，false-未过期
	Expired *bool `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 是否开通，true-开通，false-未开通
	Deployed *bool `json:"Deployed,omitnil,omitempty" name:"Deployed"`

	// 开通服务的 VPC 名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 开通服务的 VPC 对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// 开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 外部IP
	PublicIpSet []*string `json:"PublicIpSet,omitnil,omitempty" name:"PublicIpSet"`

	// 内部IP
	PrivateIpSet []*string `json:"PrivateIpSet,omitnil,omitempty" name:"PrivateIpSet"`

	// 服务开通的高级功能列表，如:[DB]
	ModuleSet []*string `json:"ModuleSet,omitnil,omitempty" name:"ModuleSet"`

	// 已使用的授权点数
	UsedNodes *uint64 `json:"UsedNodes,omitnil,omitempty" name:"UsedNodes"`

	// 扩展点数
	ExtendPoints *uint64 `json:"ExtendPoints,omitnil,omitempty" name:"ExtendPoints"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *uint64 `json:"PackageBandwidth,omitnil,omitempty" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *uint64 `json:"PackageNode,omitnil,omitempty" name:"PackageNode"`

	// 日志投递规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogDeliveryArgs *string `json:"LogDeliveryArgs,omitnil,omitempty" name:"LogDeliveryArgs"`

	// 堡垒机资源LB
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClbSet []*Clb `json:"ClbSet,omitnil,omitempty" name:"ClbSet"`

	// 网络域个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainCount *int64 `json:"DomainCount,omitnil,omitempty" name:"DomainCount"`

	// 已使用网络域个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedDomainCount *uint64 `json:"UsedDomainCount,omitnil,omitempty" name:"UsedDomainCount"`
}

type RunChangePwdTaskDetail struct {
	// 资产id
	DeviceId *uint64 `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// 资产账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`
}

// Predefined struct for user
type RunChangePwdTaskRequestParams struct {
	// 任务Id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 部门id
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 改密任务详情
	Details []*RunChangePwdTaskDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

type RunChangePwdTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	OperationId *string `json:"OperationId,omitnil,omitempty" name:"OperationId"`

	// 部门id
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 改密任务详情
	Details []*RunChangePwdTaskDetail `json:"Details,omitnil,omitempty" name:"Details"`
}

func (r *RunChangePwdTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunChangePwdTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationId")
	delete(f, "DepartmentId")
	delete(f, "Details")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunChangePwdTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunChangePwdTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunChangePwdTaskResponse struct {
	*tchttp.BaseResponse
	Response *RunChangePwdTaskResponseParams `json:"Response"`
}

func (r *RunChangePwdTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunChangePwdTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAuditLogRequestParams struct {
	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type SearchAuditLogRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	AuditLogSet []*AuditLogResult `json:"AuditLogSet,omitnil,omitempty" name:"AuditLogSet"`

	// 日志总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 命令，可模糊搜索
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据拦截状态进行过滤
	AuditAction []*uint64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`
}

type SearchCommandBySidRequest struct {
	*tchttp.BaseRequest
	
	// 会话Id
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 命令，可模糊搜索
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据拦截状态进行过滤
	AuditAction []*uint64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 命令列表
	CommandSet []*Command `json:"CommandSet,omitnil,omitempty" name:"CommandSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 搜索区间的结束时间，不填默认为开始时间到现在为止
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 资产实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 资产的公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 资产的内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 执行的命令
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 根据拦截状态进行过滤：1 - 已执行，2 - 被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type SearchCommandRequest struct {
	*tchttp.BaseRequest
	
	// 搜索区间的开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 搜索区间的结束时间，不填默认为开始时间到现在为止
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 资产实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 资产的公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 资产的内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 执行的命令
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// Cmd字段是前端传值是否进行base64.
	// 0:否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 根据拦截状态进行过滤：1 - 已执行，2 - 被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`

	// 每页容量，默认20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 命令列表
	Commands []*SearchCommandResult `json:"Commands,omitnil,omitempty" name:"Commands"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 命令
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 命令执行情况，1--允许，2--拒绝
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 命令所属的会话ID
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 命令执行时间相对于所属会话开始时间的偏移量，单位ms
	TimeOffset *uint64 `json:"TimeOffset,omitnil,omitempty" name:"TimeOffset"`

	// 账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// source ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	FromIp *string `json:"FromIp,omitnil,omitempty" name:"FromIp"`

	// 该命令所属会话的会话开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionTime *string `json:"SessionTime,omitnil,omitempty" name:"SessionTime"`

	// 该命令所属会话的会话开始时间（废弃，使用SessionTime）
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SessTime is deprecated.
	SessTime *string `json:"SessTime,omitnil,omitempty" name:"SessTime"`

	// 复核时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfirmTime *string `json:"ConfirmTime,omitnil,omitempty" name:"ConfirmTime"`

	// 部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentId *string `json:"UserDepartmentId,omitnil,omitempty" name:"UserDepartmentId"`

	// 用户部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDepartmentName *string `json:"UserDepartmentName,omitnil,omitempty" name:"UserDepartmentName"`

	// 设备部门id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentId *string `json:"DeviceDepartmentId,omitnil,omitempty" name:"DeviceDepartmentId"`

	// 设备部门名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeviceDepartmentName *string `json:"DeviceDepartmentName,omitnil,omitempty" name:"DeviceDepartmentName"`

	// 会话大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`
}

// Predefined struct for user
type SearchFileBySidRequestParams struct {
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 是否创建审计日志,通过查看按钮调用时为true,其他为false
	AuditLog *bool `json:"AuditLog,omitnil,omitempty" name:"AuditLog"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 可填写路径名或文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 分页用偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1-已执行，  2-被阻断
	AuditAction *int64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`

	// 以Protocol和Method为条件查询
	TypeFilters []*SearchFileTypeFilter `json:"TypeFilters,omitnil,omitempty" name:"TypeFilters"`
}

type SearchFileBySidRequest struct {
	*tchttp.BaseRequest
	
	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Sid *string `json:"Sid,omitnil,omitempty" name:"Sid"`

	// 是否创建审计日志,通过查看按钮调用时为true,其他为false
	AuditLog *bool `json:"AuditLog,omitnil,omitempty" name:"AuditLog"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 可填写路径名或文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 分页用偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 1-已执行，  2-被阻断
	AuditAction *int64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`

	// 以Protocol和Method为条件查询
	TypeFilters []*SearchFileTypeFilter `json:"TypeFilters,omitnil,omitempty" name:"TypeFilters"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 某会话的文件操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchFileBySidResult []*SearchFileBySidResult `json:"SearchFileBySidResult,omitnil,omitempty" name:"SearchFileBySidResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 1-上传文件 2-下载文件 3-删除文件 4-移动文件 5-重命名文件 6-新建文件夹 7-移动文件夹 8-重命名文件夹 9-删除文件夹
	Method *int64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 文件传输协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// method为上传、下载、删除时文件在服务器上的位置, 或重命名、移动文件前文件的位置
	FileCurr *string `json:"FileCurr,omitnil,omitempty" name:"FileCurr"`

	// method为重命名、移动文件时代表移动后的新位置.其他情况为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileNew *string `json:"FileNew,omitnil,omitempty" name:"FileNew"`

	// method为上传文件、下载文件、删除文件时显示文件大小。其他情况为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 堡垒机拦截情况, 1-已执行，  2-被阻断
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`
}

// Predefined struct for user
type SearchFileRequestParams struct {
	// 查询开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹
	Method []*uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 可填写路径名或文件（夹）名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 1-已执行，  2-被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type SearchFileRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹
	Method []*uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 可填写路径名或文件（夹）名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 1-已执行，  2-被阻断
	AuditAction []*uint64 `json:"AuditAction,omitnil,omitempty" name:"AuditAction"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 文件操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Files []*SearchFileResult `json:"Files,omitnil,omitempty" name:"Files"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 资产ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名称
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 资产公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 资产内网IP
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 操作结果：1 - 已执行，2 - 已阻断
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 操作类型：1 - 文件上传，2 - 文件下载，3 - 文件删除，4 - 文件(夹)移动，5 - 文件(夹)重命名，6 - 新建文件夹，9 - 删除文件夹
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 下载的文件（夹）路径及名称
	FileCurr *string `json:"FileCurr,omitnil,omitempty" name:"FileCurr"`

	// 上传或新建文件（夹）路径及名称
	FileNew *string `json:"FileNew,omitnil,omitempty" name:"FileNew"`
}

type SearchFileTypeFilter struct {
	// 需要查询的文件传输类型，如SFTP/CLIP/RDP/RZSZ
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 在当前指定的protocol下进一步过滤具体操作类型,如剪贴板文件上传，剪贴板文件下载等
	Method []*int64 `json:"Method,omitnil,omitempty" name:"Method"`
}

// Predefined struct for user
type SearchSessionCommandRequestParams struct {
	// 检索的目标命令，为模糊搜索
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 默认值为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Cmd字段前端是否做base64加密
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type SearchSessionCommandRequest struct {
	*tchttp.BaseRequest
	
	// 检索的目标命令，为模糊搜索
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 开始时间，不得早于当前时间的180天前
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 默认值为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Cmd字段前端是否做base64加密
	// 0：否，1：是
	Encoding *uint64 `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 命令和所属会话
	CommandSessionSet []*SessionCommand `json:"CommandSessionSet,omitnil,omitempty" name:"CommandSessionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 用户名，长度不超过20
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号，长度不超过64
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 来源Ip
	FromIp *string `json:"FromIp,omitnil,omitempty" name:"FromIp"`

	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 搜索区间的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 会话协议类型，只能是1、2、3或4 对应关系为1-tui 2-gui 3-file 4-数据库。若入参是Id，则非必传，否则为必传。
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 姓名，长度不超过20
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 主机名，长度不超过64
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 状态，1为活跃，2为结束，3为强制离线，4为其他错误
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type SearchSessionRequest struct {
	*tchttp.BaseRequest
	
	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 用户名，长度不超过20
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号，长度不超过64
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 来源Ip
	FromIp *string `json:"FromIp,omitnil,omitempty" name:"FromIp"`

	// 搜索区间的开始时间。若入参是Id，则非必传，否则为必传。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 搜索区间的结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 会话协议类型，只能是1、2、3或4 对应关系为1-tui 2-gui 3-file 4-数据库。若入参是Id，则非必传，否则为必传。
	Kind *uint64 `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页的页内记录数，默认为20，最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 姓名，长度不超过20
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 主机名，长度不超过64
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 状态，1为活跃，2为结束，3为强制离线，4为其他错误
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 若入参为Id，则其他入参字段不作为搜索依据，仅按照Id来搜索会话
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 会话信息列表
	SessionSet []*SessionResult `json:"SessionSet,omitnil,omitempty" name:"SessionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 账号
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 命令列表
	Commands []*Command `json:"Commands,omitnil,omitempty" name:"Commands"`

	// 记录数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 会话Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备所属的地域
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`
}

type SessionResult struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 姓名
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 主机账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 会话大小
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 设备ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备名
	DeviceName *string `json:"DeviceName,omitnil,omitempty" name:"DeviceName"`

	// 内部Ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 外部Ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 来源Ip
	FromIp *string `json:"FromIp,omitnil,omitempty" name:"FromIp"`

	// 会话持续时长
	Duration *float64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 该会话内命令数量 ，搜索图形会话时该字段无意义
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 该会话内高危命令数，搜索图形时该字段无意义
	DangerCount *uint64 `json:"DangerCount,omitnil,omitempty" name:"DangerCount"`

	// 会话状态，如1会话活跃  2会话结束  3强制离线  4其他错误
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 会话Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 设备所属的地域
	ApCode *string `json:"ApCode,omitnil,omitempty" name:"ApCode"`

	// 会话协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type User struct {
	// 用户名, 3-20个字符 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户姓名， 最大20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitnil,omitempty" name:"RealName"`

	// 用户ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 手机号码， 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitnil,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitnil,omitempty" name:"ValidateTo"`

	// 所属用户组列表
	GroupSet []*Group `json:"GroupSet,omitnil,omitempty" name:"GroupSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitnil,omitempty" name:"ValidateTime"`

	// 用户所属部门（用于出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitnil,omitempty" name:"Department"`

	// 用户所属部门（用于入参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitnil,omitempty" name:"DepartmentId"`

	// 激活状态 0 - 未激活 1 - 激活
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveStatus *uint64 `json:"ActiveStatus,omitnil,omitempty" name:"ActiveStatus"`

	// 锁定状态 0 - 未锁定 1 - 锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockStatus *uint64 `json:"LockStatus,omitnil,omitempty" name:"LockStatus"`

	// 状态 与Filter中一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 权限版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclVersion *uint64 `json:"AclVersion,omitnil,omitempty" name:"AclVersion"`
}