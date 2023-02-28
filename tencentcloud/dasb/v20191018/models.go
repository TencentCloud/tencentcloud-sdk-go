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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Acl struct {
	// 访问权限ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 访问权限名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（目前含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（目前含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，暂未启用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，暂未启用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 关联的用户列表
	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`

	// 关联的用户组列表
	UserGroupSet []*Group `json:"UserGroupSet,omitempty" name:"UserGroupSet"`

	// 关联的资产列表
	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`

	// 关联的资产组列表
	DeviceGroupSet []*Group `json:"DeviceGroupSet,omitempty" name:"DeviceGroupSet"`

	// 关联的账号列表
	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`

	// 关联的高危命令模板列表
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`

	// 是否开启 rz sz 文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`

	// 是否开启 rz sz 文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 所属部门的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitempty" name:"Department"`
}

// Predefined struct for user
type AddDeviceGroupMembersRequestParams struct {
	// 资产组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需要添加到资产组的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

type AddDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需要添加到资产组的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

type AddUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 上一次同步的结果。 0 - 从未进行, 1 - 成功， 2 - 失败
	LastStatus *uint64 `json:"LastStatus,omitempty" name:"LastStatus"`

	// 同步任务是否正在进行中
	InProcess *bool `json:"InProcess,omitempty" name:"InProcess"`
}

// Predefined struct for user
type BindDeviceAccountPasswordRequestParams struct {
	// 主机账号ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机账号密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type BindDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机账号密码
	Password *string `json:"Password,omitempty" name:"Password"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机账号私钥，最新长度128字节，最大长度8192字节
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 主机账号私钥口令，最大长度256字节
	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
}

type BindDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机账号私钥，最新长度128字节，最大长度8192字节
	PrivateKey *string `json:"PrivateKey,omitempty" name:"PrivateKey"`

	// 主机账号私钥口令，最大长度256字节
	PrivateKeyPassword *string `json:"PrivateKeyPassword,omitempty" name:"PrivateKeyPassword"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 堡垒机服务ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type BindDeviceResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 堡垒机服务ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 高危命令模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命令列表，命令之间用换行符（"\n"）分隔
	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
}

// Predefined struct for user
type CreateAclRequestParams struct {
	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启 SFTP 文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启 SFTP 文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID集合
	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 访问权限所属部门的ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest
	
	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启 SFTP 文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启 SFTP 文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID集合
	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID集合
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 访问权限所属部门的ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	delete(f, "AllowDiskFileUp")
	delete(f, "AllowDiskFileDown")
	delete(f, "AllowShellFileUp")
	delete(f, "AllowShellFileDown")
	delete(f, "AllowFileDel")
	delete(f, "ValidateFrom")
	delete(f, "ValidateTo")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAclResponseParams struct {
	// 新建成功的访问权限ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

type CreateAssetSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步资产类别，1 - 主机资产, 2 - 数据库资产
	Category *uint64 `json:"Category,omitempty" name:"Category"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命令列表，\n分隔，最大长度32768字节
	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
}

type CreateCmdTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名，最大长度32字符，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命令列表，\n分隔，最大长度32768字节
	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmdTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCmdTemplateResponseParams struct {
	// 新建成功后返回的记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitempty" name:"Account"`
}

type CreateDeviceAccountRequest struct {
	*tchttp.BaseRequest
	
	// 主机记录ID
	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitempty" name:"Account"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CreateDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组名，最大长度32字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateUserGroupRequestParams struct {
	// 用户组名，最大长度32字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户组所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组名，最大长度32字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 用户组所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户姓名，最大长度20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码， 按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户名, 3-20个字符, 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户姓名，最大长度20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码， 按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地， 1 - LDAP， 2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`

	// 所属部门ID，如：“1.2.3”
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteAclsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的权限ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteCmdTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteDeviceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需要删除的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

type DeleteDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需要删除的资产ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需删除的成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
}

type DeleteUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 需删除的成员用户ID集合
	MemberIdSet []*uint64 `json:"MemberIdSet,omitempty" name:"MemberIdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的用户ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *string `json:"Id,omitempty" name:"Id"`

	// 部门名称，1 - 256个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 部门管理员账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Managers []*string `json:"Managers,omitempty" name:"Managers"`
}

// Predefined struct for user
type DeployResourceRequestParams struct {
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 需要开通服务的地域
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 需要开通服务的VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 需要开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 需要开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 需要开通服务的VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 需要开通服务的VPC对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 需要开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
}

type DeployResourceRequest struct {
	*tchttp.BaseRequest
	
	// 需要开通服务的资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 需要开通服务的地域
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 子网所在可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 需要开通服务的VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 需要开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 需要开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 需要开通服务的VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 需要开通服务的VPC对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 需要开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否根据Name进行精确查询，默认值false
	Exact *bool `json:"Exact,omitempty" name:"Exact"`

	// 有访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`

	// 有访问权限的资产ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 部门ID，用于过滤属于某个部门的访问权限
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type DescribeAclsRequest struct {
	*tchttp.BaseRequest
	
	// 访问权限ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20，最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否根据Name进行精确查询，默认值false
	Exact *bool `json:"Exact,omitempty" name:"Exact"`

	// 有访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`

	// 有访问权限的资产ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`

	// 访问权限状态，1 - 已生效，2 - 未生效，3 - 已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 部门ID，用于过滤属于某个部门的访问权限
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 访问权限列表
	AclSet []*Acl `json:"AclSet,omitempty" name:"AclSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Category *uint64 `json:"Category,omitempty" name:"Category"`
}

type DescribeAssetSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// 查询的资产同步类型。1 -主机资产， 2 - 数据库资产
	Category *uint64 `json:"Category,omitempty" name:"Category"`
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
	Status *AssetSyncStatus `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 命令模板名，模糊查询，最大长度64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCmdTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 命令模板ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 命令模板名，模糊查询，最大长度64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BaseImageId *string `json:"BaseImageId,omitempty" name:"BaseImageId"`

	// AI镜像ID
	AiImageId *string `json:"AiImageId,omitempty" name:"AiImageId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交
	Account *string `json:"Account,omitempty" name:"Account"`

	// 主机ID，未使用IdSet时必须携带
	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeDeviceAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 主机账号ID集合，非必需，如果使用IdSet则忽略其他过滤参数
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 主机账号名，模糊查询，不能单独出现，必须于DeviceId一起提交
	Account *string `json:"Account,omitempty" name:"Account"`

	// 主机ID，未使用IdSet时必须携带
	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 账号信息列表
	DeviceAccountSet []*DeviceAccount `json:"DeviceAccountSet,omitempty" name:"DeviceAccountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数，默认20, 最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeDeviceGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// true - 查询已在该资产组的资产，false - 查询未在该资产组的资产
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数，默认20, 最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 资产类型，1 - Linux，2 - Windows，3 - MySQL，4 - SQLServer
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资产组成员列表
	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 资产组名，最长64个字符，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的资产组
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type DescribeDeviceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 资产组ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 资产组名，最长64个字符，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的资产组
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资产组列表
	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`

	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 有该资产访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`

	// 过滤条件，资产绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`

	// 过滤条件，可按照部门ID进行过滤
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 过滤数组。支持的Name：
	// BindingStatus 绑定状态
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest
	
	// 资产ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 资产名或资产IP，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`

	// 操作系统类型, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 有该资产访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`

	// 过滤条件，资产绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`

	// 过滤条件，可按照部门ID进行过滤
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`

	// 过滤条件，可按照标签键、标签进行过滤。如果同时指定标签键和标签过滤条件，它们之间为“AND”的关系
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`

	// 过滤数组。支持的Name：
	// BindingStatus 绑定状态
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 资产信息列表
	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeResourcesRequestParams struct {
	// 地域码, 如: ap-guangzhou
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 按照堡垒机开通的 VPC 实例ID查询
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 地域码, 如: ap-guangzhou
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 按照堡垒机开通的 VPC 实例ID查询
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 资源ID集合，当传入ID集合时忽略 ApCode 和 VpcId
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
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
	ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 用户名或用户姓名，最长64个字符，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type DescribeUserGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// true - 查询已添加到该用户组的用户，false - 查询未添加到该用户组的用户
	Bound *bool `json:"Bound,omitempty" name:"Bound"`

	// 用户名或用户姓名，最长64个字符，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 所属部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户组成员列表
	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 用户组名，模糊查询,长度：0-64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的用户组
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type DescribeUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 用户组名，模糊查询,长度：0-64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，缺省20，最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 部门ID，用于过滤属于某个部门的用户组
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户组列表
	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 查询具有指定资产ID访问权限的用户
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitempty" name:"AuthTypeSet"`

	// 部门ID，用于过滤属于某个部门的用户
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 如果IdSet不为空，则忽略其他参数
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页偏移位置，默认值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20, 最大500
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效。
	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 查询具有指定资产ID访问权限的用户
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0 - 本地, 1 - LDAP, 2 - OAuth, 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitempty" name:"AuthTypeSet"`

	// 部门ID，用于过滤属于某个部门的用户
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	delete(f, "AuthorizedDeviceIdSet")
	delete(f, "AuthTypeSet")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// 用户总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户列表
	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 实例ID，对应CVM、CDB等实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资产名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 地域编码
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 资产类型 1 - Linux, 2 - Windows, 3 - MySQL, 4 - SQLServer
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 管理端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 所属资产组列表
	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`

	// 资产绑定的账号数
	AccountCount *uint64 `json:"AccountCount,omitempty" name:"AccountCount"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 堡垒机服务信息，注意没有绑定服务时为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`

	// 资产所属部门
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitempty" name:"Department"`
}

type DeviceAccount struct {
	// 账号ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机ID
	DeviceId *uint64 `json:"DeviceId,omitempty" name:"DeviceId"`

	// 账号名
	Account *string `json:"Account,omitempty" name:"Account"`

	// true-已托管密码，false-未托管密码
	BoundPassword *bool `json:"BoundPassword,omitempty" name:"BoundPassword"`

	// true-已托管私钥，false-未托管私钥
	BoundPrivateKey *bool `json:"BoundPrivateKey,omitempty" name:"BoundPrivateKey"`
}

type ExternalDevice struct {
	// 操作系统名称，只能是Linux、Windows或MySQL
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// IP地址
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 管理端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 主机名，可为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产所属的部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Group struct {
	// 组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 所属部门信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitempty" name:"Department"`

	// 个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

// Predefined struct for user
type ImportExternalDeviceRequestParams struct {
	// 资产参数列表
	DeviceSet []*ExternalDevice `json:"DeviceSet,omitempty" name:"DeviceSet"`
}

type ImportExternalDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 资产参数列表
	DeviceSet []*ExternalDevice `json:"DeviceSet,omitempty" name:"DeviceSet"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyAclRequestParams struct {
	// 访问权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 权限所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type ModifyAclRequest struct {
	*tchttp.BaseRequest
	
	// 访问权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登录
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板文本（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板文本（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制（预留参数，目前暂未使用）
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制（预留参数，目前暂未使用）
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`

	// 关联的资产ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 关联的资产组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号
	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`

	// 关联的高危命令模板ID
	CmdTemplateIdSet []*uint64 `json:"CmdTemplateIdSet,omitempty" name:"CmdTemplateIdSet"`

	// 是否开启 RDP 磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`

	// 是否开启 RDP 磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`

	// 是否开启 SFTP 文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 访问权限生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 访问权限失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则访问权限长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 权限所属部门的ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	delete(f, "AllowDiskFileUp")
	delete(f, "AllowDiskFileDown")
	delete(f, "AllowShellFileUp")
	delete(f, "AllowShellFileDown")
	delete(f, "AllowFileDel")
	delete(f, "ValidateFrom")
	delete(f, "ValidateTo")
	delete(f, "DepartmentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAclResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDeviceGroupRequestParams struct {
	// 资产组名，最大长度32字符，不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type ModifyDeviceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组名，最大长度32字符，不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资产组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 资产组所属部门ID，如：1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 管理端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 资产所属组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 资产所属部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type ModifyDeviceRequest struct {
	*tchttp.BaseRequest
	
	// 资产记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 管理端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 资产所属组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 资产所属部门ID
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyUserRequestParams struct {
	// 用户ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 用户姓名，最大长度20个字符，不能包含空格
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`

	// 用户所属部门的ID，如1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 用户姓名，最大长度20个字符，不能包含空格
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-23T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth 不传则默认为0
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`

	// 用户所属部门的ID，如1.2.3
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ResetDeviceAccountPasswordRequestParams struct {
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type ResetDeviceAccountPasswordRequest struct {
	*tchttp.BaseRequest
	
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type ResetDeviceAccountPrivateKeyRequest struct {
	*tchttp.BaseRequest
	
	// ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
}

type ResetUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID集合
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 地域编码
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 服务实例规格信息
	SvArgs *string `json:"SvArgs,omitempty" name:"SvArgs"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 服务规格对应的资产数
	Nodes *uint64 `json:"Nodes,omitempty" name:"Nodes"`

	// 自动续费标记，0 - 表示默认状态，1 - 表示自动续费，2 - 表示明确不自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 资源状态，0 - 未初始化，1 - 正常，2 - 隔离，3 - 销毁，4 - 初始化失败，5 - 初始化中
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 服务实例名，如T-Sec-堡垒机（SaaS型）
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 定价模型ID
	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`

	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 商品码, p_cds_dasb
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子商品码, sp_cds_dasb_bh_saas
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 是否过期，true-过期，false-未过期
	Expired *bool `json:"Expired,omitempty" name:"Expired"`

	// 是否开通，true-开通，false-未开通
	Deployed *bool `json:"Deployed,omitempty" name:"Deployed"`

	// 开通服务的 VPC 名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 开通服务的 VPC 对应的网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitempty" name:"VpcCidrBlock"`

	// 开通服务的子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 开通服务的子网名称
	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`

	// 开通服务的子网网段
	CidrBlock *string `json:"CidrBlock,omitempty" name:"CidrBlock"`

	// 外部IP
	PublicIpSet []*string `json:"PublicIpSet,omitempty" name:"PublicIpSet"`

	// 内部IP
	PrivateIpSet []*string `json:"PrivateIpSet,omitempty" name:"PrivateIpSet"`

	// 服务开通的高级功能列表，如:[DB]
	ModuleSet []*string `json:"ModuleSet,omitempty" name:"ModuleSet"`

	// 已使用的授权点数
	UsedNodes *uint64 `json:"UsedNodes,omitempty" name:"UsedNodes"`

	// 扩展点数
	ExtendPoints *uint64 `json:"ExtendPoints,omitempty" name:"ExtendPoints"`

	// 带宽扩展包个数(4M)
	PackageBandwidth *uint64 `json:"PackageBandwidth,omitempty" name:"PackageBandwidth"`

	// 授权点数扩展包个数(50点)
	PackageNode *uint64 `json:"PackageNode,omitempty" name:"PackageNode"`
}

type TagFilter struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type User struct {
	// 用户名, 3-20个字符 必须以英文字母开头，且不能包含字母、数字、.、_、-以外的字符
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户姓名， 最大20个字符，不能包含空白字符
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 手机号码， 大陆手机号直接填写，如果是其他国家、地区号码,按照"国家地区代码|手机号"的格式输入。如: "+852|xxxxxxxx"
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户生效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 用户失效时间，如:"2021-09-22T00:00:00+00:00"
	// 生效、失效时间不填则用户长期有效
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组列表
	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`

	// 认证方式，0 - 本地，1 - LDAP，2 - OAuth
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 访问时间段限制， 由0、1组成的字符串，长度168(7 × 24)，代表该用户在一周中允许访问的时间段。字符串中第N个字符代表在一周中的第N个小时， 0 - 代表不允许访问，1 - 代表允许访问
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`

	// 用户所属部门（用于出参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Department *Department `json:"Department,omitempty" name:"Department"`

	// 用户所属部门（用于入参）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DepartmentId *string `json:"DepartmentId,omitempty" name:"DepartmentId"`
}