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

	// 规则名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板text（目前含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板text（目前含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 是否允许任意账号登陆
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 关联的用户列表
	UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`

	// 关联的用户组列表
	UserGroupSet []*Group `json:"UserGroupSet,omitempty" name:"UserGroupSet"`

	// 关联的主机列表
	DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`

	// 关联的主机组列表
	DeviceGroupSet []*Group `json:"DeviceGroupSet,omitempty" name:"DeviceGroupSet"`

	// 关联的账号列表
	AccountSet []*string `json:"AccountSet,omitempty" name:"AccountSet"`

	// 关联的高危命令模板列表
	CmdTemplateSet []*CmdTemplate `json:"CmdTemplateSet,omitempty" name:"CmdTemplateSet"`

	// 是否开启rdp磁盘映射文件上传
	AllowDiskFileUp *bool `json:"AllowDiskFileUp,omitempty" name:"AllowDiskFileUp"`

	// 是否开启rdp磁盘映射文件下载
	AllowDiskFileDown *bool `json:"AllowDiskFileDown,omitempty" name:"AllowDiskFileDown"`

	// 是否开启rz sz文件上传
	AllowShellFileUp *bool `json:"AllowShellFileUp,omitempty" name:"AllowShellFileUp"`

	// 是否开启rz sz文件下载
	AllowShellFileDown *bool `json:"AllowShellFileDown,omitempty" name:"AllowShellFileDown"`

	// 是否开启SFTP文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 生效日期
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 失效日期
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 策略状态，1-已生效，2-未生效，3-已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type CmdTemplate struct {

	// 模板ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 模板名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命令列表，\n分隔
	CmdList *string `json:"CmdList,omitempty" name:"CmdList"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// 权限名称，最大32字符，不能为空，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登陆
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板text（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板text（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启SFTP文件上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启SFTP文件下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`

	// 关联的主机ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 关联的主机组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号，账号name
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

	// 是否开启SFTP文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 生效日期，如果为空，默认1970-01-01T08:00:01+08:00
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 失效日期，如果为空，默认1970-01-01T08:00:01+08:00
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 访问权限ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 用户名，最大长度32字符，不能为空
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户姓名，最大长度32字符，不能为空
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 生效起始时间,不设置则为1970-01-01T08:00:01+08:00
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 生效结束时间,不设置则为1970-01-01T08:00:01+08:00
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 认证方式，0-本地 1-ldap 2-oauth,不传则默认为0
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 生效时间段, 0、1组成的字符串，长度168(7*24), 代表该用户的生效时间. 0 - 未生效，1 - 生效
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新建成功后返回的记录ID
		Id *uint64 `json:"Id,omitempty" name:"Id"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteAclsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeAclsRequest struct {
	*tchttp.BaseRequest

	// 访问权限ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 访问权限名称，模糊查询，最长64字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页，偏移位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否根据Name进行精确查询,默认值false
	Exact *bool `json:"Exact,omitempty" name:"Exact"`

	// 有权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`

	// 有权限的主机ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`

	// 策略状态，0-不限，1-已生效，2-未生效，3-已过期
	Status *uint64 `json:"Status,omitempty" name:"Status"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAclsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 访问权限记录集合，当前分页
		AclSet []*Acl `json:"AclSet,omitempty" name:"AclSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDasbImageIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 基础镜像ID
		BaseImageId *string `json:"BaseImageId,omitempty" name:"BaseImageId"`

		// AI镜像ID
		AiImageId *string `json:"AiImageId,omitempty" name:"AiImageId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeDevicesRequest struct {
	*tchttp.BaseRequest

	// 主机ID集合，非必需
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 主机名或主机IP，模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 暂未使用
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 地域码集合
	ApCodeSet []*string `json:"ApCodeSet,omitempty" name:"ApCodeSet"`

	// 操作系统类型
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 分页，偏移位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 有该主机访问权限的用户ID集合
	AuthorizedUserIdSet []*uint64 `json:"AuthorizedUserIdSet,omitempty" name:"AuthorizedUserIdSet"`

	// 过滤条件，主机绑定的堡垒机服务ID集合
	ResourceIdSet []*string `json:"ResourceIdSet,omitempty" name:"ResourceIdSet"`

	// 可提供按照多种类型过滤, 1-Linux, 2-Windows, 3-MySQL
	KindSet []*uint64 `json:"KindSet,omitempty" name:"KindSet"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDevicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 主机信息列表
		DeviceSet []*Device `json:"DeviceSet,omitempty" name:"DeviceSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeUsersRequest struct {
	*tchttp.BaseRequest

	// 如果IdSet不为空，则忽略其他参数
	IdSet []*uint64 `json:"IdSet,omitempty" name:"IdSet"`

	// 模糊查询，IdSet、UserName、Phone为空时才生效，对用户名和姓名进行模糊查询
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分页，偏移位置
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页条目数量，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 精确查询，IdSet为空时才生效
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 精确查询，IdSet、UserName为空时才生效
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 有访问权限的主机ID集合
	AuthorizedDeviceIdSet []*uint64 `json:"AuthorizedDeviceIdSet,omitempty" name:"AuthorizedDeviceIdSet"`

	// 认证方式，0-本地，1-ldap, 2-oauth 不传为全部
	AuthTypeSet []*uint64 `json:"AuthTypeSet,omitempty" name:"AuthTypeSet"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 用户信息列表
		UserSet []*User `json:"UserSet,omitempty" name:"UserSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

	// 主机记录ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 主机ID，对应cvm实例id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主机名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 地域编码
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 操作系统名称
	OsName *string `json:"OsName,omitempty" name:"OsName"`

	// 主机类型，1-Linux, 2-Windows
	Kind *uint64 `json:"Kind,omitempty" name:"Kind"`

	// 管理端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 所属主机组信息列表
	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`

	// 主机绑定的账号数
	AccountCount *uint64 `json:"AccountCount,omitempty" name:"AccountCount"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 堡垒机服务信息，注意没有绑定服务时为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *Resource `json:"Resource,omitempty" name:"Resource"`
}

type Group struct {

	// 组ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 组名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyAclRequest struct {
	*tchttp.BaseRequest

	// 权限名称，最大32字符，不能包含空白字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否开启磁盘映射
	AllowDiskRedirect *bool `json:"AllowDiskRedirect,omitempty" name:"AllowDiskRedirect"`

	// 是否允许任意账号登陆
	AllowAnyAccount *bool `json:"AllowAnyAccount,omitempty" name:"AllowAnyAccount"`

	// 访问权限ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 是否开启剪贴板文件上行
	AllowClipFileUp *bool `json:"AllowClipFileUp,omitempty" name:"AllowClipFileUp"`

	// 是否开启剪贴板文件下行
	AllowClipFileDown *bool `json:"AllowClipFileDown,omitempty" name:"AllowClipFileDown"`

	// 是否开启剪贴板text（含图片）上行
	AllowClipTextUp *bool `json:"AllowClipTextUp,omitempty" name:"AllowClipTextUp"`

	// 是否开启剪贴板text（含图片）下行
	AllowClipTextDown *bool `json:"AllowClipTextDown,omitempty" name:"AllowClipTextDown"`

	// 是否开启文件传输上传
	AllowFileUp *bool `json:"AllowFileUp,omitempty" name:"AllowFileUp"`

	// 文件传输上传大小限制
	MaxFileUpSize *uint64 `json:"MaxFileUpSize,omitempty" name:"MaxFileUpSize"`

	// 是否开启文件传输下载
	AllowFileDown *bool `json:"AllowFileDown,omitempty" name:"AllowFileDown"`

	// 文件传输下载大小限制
	MaxFileDownSize *uint64 `json:"MaxFileDownSize,omitempty" name:"MaxFileDownSize"`

	// 关联的用户ID
	UserIdSet []*uint64 `json:"UserIdSet,omitempty" name:"UserIdSet"`

	// 关联的用户组ID
	UserGroupIdSet []*uint64 `json:"UserGroupIdSet,omitempty" name:"UserGroupIdSet"`

	// 关联的主机ID
	DeviceIdSet []*uint64 `json:"DeviceIdSet,omitempty" name:"DeviceIdSet"`

	// 关联的主机组ID
	DeviceGroupIdSet []*uint64 `json:"DeviceGroupIdSet,omitempty" name:"DeviceGroupIdSet"`

	// 关联的账号，账号name
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

	// 是否开启SFTP文件删除
	AllowFileDel *bool `json:"AllowFileDel,omitempty" name:"AllowFileDel"`

	// 生效日期，如果为空，默认1970-01-01T08:00:01+08:00
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 失效日期，如果为空，默认1970-01-01T08:00:01+08:00
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAclResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ModifyUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 用户姓名，最大长度32字符，不能为空
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 手机号
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 生效起始时间,不设置则为1970-01-01 08:00:01
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 生效结束时间,不设置则为1970-01-01 08:00:01
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组ID集合
	GroupIdSet []*uint64 `json:"GroupIdSet,omitempty" name:"GroupIdSet"`

	// 认证方式，0-本地 1-ldap, 2-oauth不传则默认为0
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 生效时间段, 0、1组成的字符串，长度168(7*24), 代表该用户的生效时间. 0 - 未生效，1 - 生效
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type Resource struct {

	// 资源实例id，如bh-saas-s3ed4r5e
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 地域编码
	ApCode *string `json:"ApCode,omitempty" name:"ApCode"`

	// 实例规格信息（询价参数）
	SvArgs *string `json:"SvArgs,omitempty" name:"SvArgs"`

	// vpc id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 堡垒机规格对应的资产数
	Nodes *uint64 `json:"Nodes,omitempty" name:"Nodes"`

	// 自动续费标记，0表示默认状态，1表示自动续费，2表示明确不自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 资源状态，0未初始化，1正常，2隔离，3销毁，4初始化失败，5初始化中
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 实例名，如T-Sec-堡垒机（SaaS型）
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

	// 开通服务的VPC名称
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 开通服务的VPC对应的网段
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

	// 资源开通的高级功能列表，如:[DB]
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

type User struct {

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户姓名
	RealName *string `json:"RealName,omitempty" name:"RealName"`

	// 手机号码
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 电子邮件
	Email *string `json:"Email,omitempty" name:"Email"`

	// 生效起始时间
	ValidateFrom *string `json:"ValidateFrom,omitempty" name:"ValidateFrom"`

	// 生效结束时间
	ValidateTo *string `json:"ValidateTo,omitempty" name:"ValidateTo"`

	// 所属用户组列表
	GroupSet []*Group `json:"GroupSet,omitempty" name:"GroupSet"`

	// 认证方式，0-本地 1-ldap
	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`

	// 生效时间段, 0、1组成的字符串，长度168(7*24), 代表该用户的生效时间. 0 - 未生效，1 - 生效
	ValidateTime *string `json:"ValidateTime,omitempty" name:"ValidateTime"`
}
