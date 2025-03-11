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

package v20201112

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccessGroup struct {
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`

	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitnil,omitempty" name:"AccessGroupName"`

	// 权限组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// VPC网络类型（1：CVM）
	VpcType *uint64 `json:"VpcType,omitnil,omitempty" name:"VpcType"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type AccessRule struct {
	// 权限规则ID
	AccessRuleId *uint64 `json:"AccessRuleId,omitnil,omitempty" name:"AccessRuleId"`

	// 权限规则地址（网段或IP）
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 权限规则访问模式（1：只读；2：读写）
	AccessMode *uint64 `json:"AccessMode,omitnil,omitempty" name:"AccessMode"`

	// 优先级（取值范围1~100，值越小优先级越高）
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type AssociateAccessGroupsRequestParams struct {
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 权限组ID列表
	AccessGroupIds []*string `json:"AccessGroupIds,omitnil,omitempty" name:"AccessGroupIds"`
}

type AssociateAccessGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 权限组ID列表
	AccessGroupIds []*string `json:"AccessGroupIds,omitnil,omitempty" name:"AccessGroupIds"`
}

func (r *AssociateAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAccessGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	delete(f, "AccessGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateAccessGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateAccessGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateAccessGroupsResponseParams `json:"Response"`
}

func (r *AssociateAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAccessGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessGroupRequestParams struct {
	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitnil,omitempty" name:"AccessGroupName"`

	// VPC网络类型（1：CVM）
	VpcType *uint64 `json:"VpcType,omitnil,omitempty" name:"VpcType"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 权限组描述，默认为空字符串
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitnil,omitempty" name:"AccessGroupName"`

	// VPC网络类型（1：CVM）
	VpcType *uint64 `json:"VpcType,omitnil,omitempty" name:"VpcType"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 权限组描述，默认为空字符串
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupName")
	delete(f, "VpcType")
	delete(f, "VpcId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessGroupResponseParams struct {
	// 权限组
	AccessGroup *AccessGroup `json:"AccessGroup,omitnil,omitempty" name:"AccessGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessGroupResponseParams `json:"Response"`
}

func (r *CreateAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessRulesRequestParams struct {
	// 多个权限规则，上限为10
	AccessRules []*AccessRule `json:"AccessRules,omitnil,omitempty" name:"AccessRules"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

type CreateAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// 多个权限规则，上限为10
	AccessRules []*AccessRule `json:"AccessRules,omitnil,omitempty" name:"AccessRules"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

func (r *CreateAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRules")
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessRulesResponseParams struct {
	// 权限规则列表
	AccessRules []*AccessRule `json:"AccessRules,omitnil,omitempty" name:"AccessRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessRulesResponseParams `json:"Response"`
}

func (r *CreateAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemRequestParams struct {
	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitnil,omitempty" name:"FileSystemName"`

	// 是否校验POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitnil,omitempty" name:"PosixAcl"`

	// 文件系统描述，默认为空字符串
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 文件系统容量（byte），下限为1GB，上限为1PB，且必须是1GB的整数倍
	CapacityQuota *uint64 `json:"CapacityQuota,omitnil,omitempty" name:"CapacityQuota"`

	// 超级用户名列表，默认为空数组
	SuperUsers []*string `json:"SuperUsers,omitnil,omitempty" name:"SuperUsers"`

	// 根目录Inode用户名，默认为hadoop
	RootInodeUser *string `json:"RootInodeUser,omitnil,omitempty" name:"RootInodeUser"`

	// 根目录Inode组名，默认为supergroup
	RootInodeGroup *string `json:"RootInodeGroup,omitnil,omitempty" name:"RootInodeGroup"`

	// 是否打开Ranger地址校验
	EnableRanger *bool `json:"EnableRanger,omitnil,omitempty" name:"EnableRanger"`

	// Ranger地址列表，默认为空数组
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitnil,omitempty" name:"RangerServiceAddresses"`

	// 多个资源标签，可以为空数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitnil,omitempty" name:"FileSystemName"`

	// 是否校验POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitnil,omitempty" name:"PosixAcl"`

	// 文件系统描述，默认为空字符串
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 文件系统容量（byte），下限为1GB，上限为1PB，且必须是1GB的整数倍
	CapacityQuota *uint64 `json:"CapacityQuota,omitnil,omitempty" name:"CapacityQuota"`

	// 超级用户名列表，默认为空数组
	SuperUsers []*string `json:"SuperUsers,omitnil,omitempty" name:"SuperUsers"`

	// 根目录Inode用户名，默认为hadoop
	RootInodeUser *string `json:"RootInodeUser,omitnil,omitempty" name:"RootInodeUser"`

	// 根目录Inode组名，默认为supergroup
	RootInodeGroup *string `json:"RootInodeGroup,omitnil,omitempty" name:"RootInodeGroup"`

	// 是否打开Ranger地址校验
	EnableRanger *bool `json:"EnableRanger,omitnil,omitempty" name:"EnableRanger"`

	// Ranger地址列表，默认为空数组
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitnil,omitempty" name:"RangerServiceAddresses"`

	// 多个资源标签，可以为空数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemName")
	delete(f, "PosixAcl")
	delete(f, "Description")
	delete(f, "CapacityQuota")
	delete(f, "SuperUsers")
	delete(f, "RootInodeUser")
	delete(f, "RootInodeGroup")
	delete(f, "EnableRanger")
	delete(f, "RangerServiceAddresses")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileSystemResponseParams struct {
	// 文件系统
	FileSystem *FileSystem `json:"FileSystem,omitnil,omitempty" name:"FileSystem"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileSystemResponseParams `json:"Response"`
}

func (r *CreateFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifeCycleRulesRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 多个生命周期规则，上限为10
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitnil,omitempty" name:"LifeCycleRules"`
}

type CreateLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 多个生命周期规则，上限为10
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitnil,omitempty" name:"LifeCycleRules"`
}

func (r *CreateLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "LifeCycleRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifeCycleRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifeCycleRulesResponseParams `json:"Response"`
}

func (r *CreateLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMountPointRequestParams struct {
	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitnil,omitempty" name:"MountPointName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 挂载点状态（1：打开；2：关闭）
	MountPointStatus *uint64 `json:"MountPointStatus,omitnil,omitempty" name:"MountPointStatus"`
}

type CreateMountPointRequest struct {
	*tchttp.BaseRequest
	
	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitnil,omitempty" name:"MountPointName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 挂载点状态（1：打开；2：关闭）
	MountPointStatus *uint64 `json:"MountPointStatus,omitnil,omitempty" name:"MountPointStatus"`
}

func (r *CreateMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointName")
	delete(f, "FileSystemId")
	delete(f, "MountPointStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMountPointResponseParams struct {
	// 挂载点
	MountPoint *MountPoint `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMountPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateMountPointResponseParams `json:"Response"`
}

func (r *CreateMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRestoreTasksRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 多个回热任务，上限为10
	RestoreTasks []*RestoreTask `json:"RestoreTasks,omitnil,omitempty" name:"RestoreTasks"`
}

type CreateRestoreTasksRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 多个回热任务，上限为10
	RestoreTasks []*RestoreTask `json:"RestoreTasks,omitnil,omitempty" name:"RestoreTasks"`
}

func (r *CreateRestoreTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRestoreTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "RestoreTasks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRestoreTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRestoreTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRestoreTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateRestoreTasksResponseParams `json:"Response"`
}

func (r *CreateRestoreTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRestoreTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessGroupRequestParams struct {
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

type DeleteAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

func (r *DeleteAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessGroupResponseParams `json:"Response"`
}

func (r *DeleteAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessRulesRequestParams struct {
	// 多个权限规则ID，上限为10
	AccessRuleIds []*uint64 `json:"AccessRuleIds,omitnil,omitempty" name:"AccessRuleIds"`
}

type DeleteAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// 多个权限规则ID，上限为10
	AccessRuleIds []*uint64 `json:"AccessRuleIds,omitnil,omitempty" name:"AccessRuleIds"`
}

func (r *DeleteAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessRulesResponseParams `json:"Response"`
}

func (r *DeleteAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSystemRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DeleteFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DeleteFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileSystemResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileSystemResponseParams `json:"Response"`
}

func (r *DeleteFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifeCycleRulesRequestParams struct {
	// 多个生命周期规则ID，上限为10
	LifeCycleRuleIds []*uint64 `json:"LifeCycleRuleIds,omitnil,omitempty" name:"LifeCycleRuleIds"`
}

type DeleteLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// 多个生命周期规则ID，上限为10
	LifeCycleRuleIds []*uint64 `json:"LifeCycleRuleIds,omitnil,omitempty" name:"LifeCycleRuleIds"`
}

func (r *DeleteLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifeCycleRuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifeCycleRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLifeCycleRulesResponseParams `json:"Response"`
}

func (r *DeleteLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountPointRequestParams struct {
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`
}

type DeleteMountPointRequest struct {
	*tchttp.BaseRequest
	
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`
}

func (r *DeleteMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMountPointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMountPointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMountPointResponseParams `json:"Response"`
}

func (r *DeleteMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupRequestParams struct {
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

type DescribeAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

func (r *DescribeAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupResponseParams struct {
	// 权限组
	AccessGroup *AccessGroup `json:"AccessGroup,omitnil,omitempty" name:"AccessGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessGroupResponseParams `json:"Response"`
}

func (r *DescribeAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupsRequestParams struct {
	// VPC网络ID
	// 备注：入参只能指定VpcId和OwnerUin的其中一个
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 资源所属者Uin
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type DescribeAccessGroupsRequest struct {
	*tchttp.BaseRequest
	
	// VPC网络ID
	// 备注：入参只能指定VpcId和OwnerUin的其中一个
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 资源所属者Uin
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

func (r *DescribeAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessGroupsResponseParams struct {
	// 权限组列表
	AccessGroups []*AccessGroup `json:"AccessGroups,omitnil,omitempty" name:"AccessGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessGroupsResponseParams `json:"Response"`
}

func (r *DescribeAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessRulesRequestParams struct {
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

type DescribeAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

func (r *DescribeAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessRulesResponseParams struct {
	// 权限规则列表
	AccessRules []*AccessRule `json:"AccessRules,omitnil,omitempty" name:"AccessRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessRulesResponseParams `json:"Response"`
}

func (r *DescribeAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemResponseParams struct {
	// 文件系统
	FileSystem *FileSystem `json:"FileSystem,omitnil,omitempty" name:"FileSystem"`

	// 文件系统已使用容量（byte）
	CapacityUsed *uint64 `json:"CapacityUsed,omitnil,omitempty" name:"CapacityUsed"`

	// 已使用COS归档存储容量（byte）
	ArchiveCapacityUsed *uint64 `json:"ArchiveCapacityUsed,omitnil,omitempty" name:"ArchiveCapacityUsed"`

	// 已使用COS标准存储容量（byte）
	StandardCapacityUsed *uint64 `json:"StandardCapacityUsed,omitnil,omitempty" name:"StandardCapacityUsed"`

	// 已使用COS低频存储容量（byte）
	DegradeCapacityUsed *uint64 `json:"DegradeCapacityUsed,omitnil,omitempty" name:"DegradeCapacityUsed"`

	// 已使用COS深度归档存储容量（byte）
	DeepArchiveCapacityUsed *uint64 `json:"DeepArchiveCapacityUsed,omitnil,omitempty" name:"DeepArchiveCapacityUsed"`

	// 已使用COS智能分层存储容量（byte）
	IntelligentCapacityUsed *uint64 `json:"IntelligentCapacityUsed,omitnil,omitempty" name:"IntelligentCapacityUsed"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSystemResponseParams `json:"Response"`
}

func (r *DescribeFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemsRequestParams struct {

}

type DescribeFileSystemsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileSystemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileSystemsResponseParams struct {
	// 文件系统列表
	FileSystems []*FileSystem `json:"FileSystems,omitnil,omitempty" name:"FileSystems"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileSystemsResponseParams `json:"Response"`
}

func (r *DescribeFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileSystemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifeCycleRulesRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifeCycleRulesResponseParams struct {
	// 生命周期规则列表
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitnil,omitempty" name:"LifeCycleRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLifeCycleRulesResponseParams `json:"Response"`
}

func (r *DescribeLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointRequestParams struct {
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`
}

type DescribeMountPointRequest struct {
	*tchttp.BaseRequest
	
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`
}

func (r *DescribeMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointResponseParams struct {
	// 挂载点
	MountPoint *MountPoint `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountPointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountPointResponseParams `json:"Response"`
}

func (r *DescribeMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointsRequestParams struct {
	// 文件系统ID
	// 备注：入参只能指定AccessGroupId、FileSystemId和OwnerUin的其中一个
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`

	// 资源所属者Uin
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type DescribeMountPointsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	// 备注：入参只能指定AccessGroupId、FileSystemId和OwnerUin的其中一个
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`

	// 资源所属者Uin
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

func (r *DescribeMountPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "AccessGroupId")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMountPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMountPointsResponseParams struct {
	// 挂载点列表
	MountPoints []*MountPoint `json:"MountPoints,omitnil,omitempty" name:"MountPoints"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMountPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMountPointsResponseParams `json:"Response"`
}

func (r *DescribeMountPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMountPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceTagsResponseParams struct {
	// 资源标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceTagsResponseParams `json:"Response"`
}

func (r *DescribeResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTasksRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

type DescribeRestoreTasksRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`
}

func (r *DescribeRestoreTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRestoreTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRestoreTasksResponseParams struct {
	// 回热任务列表
	RestoreTasks []*RestoreTask `json:"RestoreTasks,omitnil,omitempty" name:"RestoreTasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRestoreTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRestoreTasksResponseParams `json:"Response"`
}

func (r *DescribeRestoreTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRestoreTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAccessGroupsRequestParams struct {
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 权限组ID列表
	AccessGroupIds []*string `json:"AccessGroupIds,omitnil,omitempty" name:"AccessGroupIds"`
}

type DisassociateAccessGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 权限组ID列表
	AccessGroupIds []*string `json:"AccessGroupIds,omitnil,omitempty" name:"AccessGroupIds"`
}

func (r *DisassociateAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAccessGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	delete(f, "AccessGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateAccessGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAccessGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateAccessGroupsResponseParams `json:"Response"`
}

func (r *DisassociateAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAccessGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileSystem struct {
	// 资源所属用户AppId
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitnil,omitempty" name:"FileSystemName"`

	// 文件系统描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件系统块大小（byte）
	BlockSize *uint64 `json:"BlockSize,omitnil,omitempty" name:"BlockSize"`

	// 文件系统容量（byte）
	CapacityQuota *uint64 `json:"CapacityQuota,omitnil,omitempty" name:"CapacityQuota"`

	// 文件系统状态（1：创建中；2：创建成功；3：创建失败）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 超级用户名列表
	SuperUsers []*string `json:"SuperUsers,omitnil,omitempty" name:"SuperUsers"`

	// POSIX权限控制
	PosixAcl *bool `json:"PosixAcl,omitnil,omitempty" name:"PosixAcl"`

	// 是否打开Ranger地址校验
	EnableRanger *bool `json:"EnableRanger,omitnil,omitempty" name:"EnableRanger"`

	// Ranger地址列表
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitnil,omitempty" name:"RangerServiceAddresses"`
}

type LifeCycleRule struct {
	// 生命周期规则ID
	LifeCycleRuleId *uint64 `json:"LifeCycleRuleId,omitnil,omitempty" name:"LifeCycleRuleId"`

	// 生命周期规则名称
	LifeCycleRuleName *string `json:"LifeCycleRuleName,omitnil,omitempty" name:"LifeCycleRuleName"`

	// 生命周期规则路径（目录或文件）
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 生命周期规则转换列表
	Transitions []*Transition `json:"Transitions,omitnil,omitempty" name:"Transitions"`

	// 生命周期规则状态（1：打开；2：关闭）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 生命周期规则当前路径具体存储量
	Summary *Summary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// Summary更新时间
	LastSummaryTime *string `json:"LastSummaryTime,omitnil,omitempty" name:"LastSummaryTime"`
}

// Predefined struct for user
type ModifyAccessGroupRequestParams struct {
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`

	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitnil,omitempty" name:"AccessGroupName"`

	// 权限组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyAccessGroupRequest struct {
	*tchttp.BaseRequest
	
	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`

	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitnil,omitempty" name:"AccessGroupName"`

	// 权限组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessGroupId")
	delete(f, "AccessGroupName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessGroupResponseParams `json:"Response"`
}

func (r *ModifyAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessRulesRequestParams struct {
	// 多个权限规则，上限为10
	AccessRules []*AccessRule `json:"AccessRules,omitnil,omitempty" name:"AccessRules"`
}

type ModifyAccessRulesRequest struct {
	*tchttp.BaseRequest
	
	// 多个权限规则，上限为10
	AccessRules []*AccessRule `json:"AccessRules,omitnil,omitempty" name:"AccessRules"`
}

func (r *ModifyAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessRulesResponseParams `json:"Response"`
}

func (r *ModifyAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitnil,omitempty" name:"FileSystemName"`

	// 文件系统描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 文件系统容量（byte），下限为1GB，上限为1PB，且必须是1GB的整数倍
	// 注意：修改的文件系统容量不能小于当前使用量
	CapacityQuota *uint64 `json:"CapacityQuota,omitnil,omitempty" name:"CapacityQuota"`

	// 超级用户名列表，可以为空数组
	SuperUsers []*string `json:"SuperUsers,omitnil,omitempty" name:"SuperUsers"`

	// 是否校验POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitnil,omitempty" name:"PosixAcl"`

	// 是否打开Ranger地址校验
	EnableRanger *bool `json:"EnableRanger,omitnil,omitempty" name:"EnableRanger"`

	// Ranger地址列表，可以为空数组
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitnil,omitempty" name:"RangerServiceAddresses"`
}

type ModifyFileSystemRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitnil,omitempty" name:"FileSystemName"`

	// 文件系统描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 文件系统容量（byte），下限为1GB，上限为1PB，且必须是1GB的整数倍
	// 注意：修改的文件系统容量不能小于当前使用量
	CapacityQuota *uint64 `json:"CapacityQuota,omitnil,omitempty" name:"CapacityQuota"`

	// 超级用户名列表，可以为空数组
	SuperUsers []*string `json:"SuperUsers,omitnil,omitempty" name:"SuperUsers"`

	// 是否校验POSIX ACL
	PosixAcl *bool `json:"PosixAcl,omitnil,omitempty" name:"PosixAcl"`

	// 是否打开Ranger地址校验
	EnableRanger *bool `json:"EnableRanger,omitnil,omitempty" name:"EnableRanger"`

	// Ranger地址列表，可以为空数组
	RangerServiceAddresses []*string `json:"RangerServiceAddresses,omitnil,omitempty" name:"RangerServiceAddresses"`
}

func (r *ModifyFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "FileSystemName")
	delete(f, "Description")
	delete(f, "CapacityQuota")
	delete(f, "SuperUsers")
	delete(f, "PosixAcl")
	delete(f, "EnableRanger")
	delete(f, "RangerServiceAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFileSystemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFileSystemResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFileSystemResponseParams `json:"Response"`
}

func (r *ModifyFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFileSystemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifeCycleRulesRequestParams struct {
	// 多个生命周期规则，上限为10
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitnil,omitempty" name:"LifeCycleRules"`
}

type ModifyLifeCycleRulesRequest struct {
	*tchttp.BaseRequest
	
	// 多个生命周期规则，上限为10
	LifeCycleRules []*LifeCycleRule `json:"LifeCycleRules,omitnil,omitempty" name:"LifeCycleRules"`
}

func (r *ModifyLifeCycleRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifeCycleRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifeCycleRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLifeCycleRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifeCycleRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLifeCycleRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLifeCycleRulesResponseParams `json:"Response"`
}

func (r *ModifyLifeCycleRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifeCycleRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMountPointRequestParams struct {
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitnil,omitempty" name:"MountPointName"`

	// 挂载点状态
	MountPointStatus *uint64 `json:"MountPointStatus,omitnil,omitempty" name:"MountPointStatus"`
}

type ModifyMountPointRequest struct {
	*tchttp.BaseRequest
	
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitnil,omitempty" name:"MountPointName"`

	// 挂载点状态
	MountPointStatus *uint64 `json:"MountPointStatus,omitnil,omitempty" name:"MountPointStatus"`
}

func (r *ModifyMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMountPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPointId")
	delete(f, "MountPointName")
	delete(f, "MountPointStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMountPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMountPointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMountPointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMountPointResponseParams `json:"Response"`
}

func (r *ModifyMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMountPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsRequestParams struct {
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 多个资源标签，可以为空数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyResourceTagsRequest struct {
	*tchttp.BaseRequest
	
	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 多个资源标签，可以为空数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyResourceTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileSystemId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceTagsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceTagsResponseParams `json:"Response"`
}

func (r *ModifyResourceTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountPoint struct {
	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitnil,omitempty" name:"MountPointId"`

	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitnil,omitempty" name:"MountPointName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitnil,omitempty" name:"FileSystemId"`

	// 挂载点状态（1：打开；2：关闭）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 绑定的权限组ID列表
	AccessGroupIds []*string `json:"AccessGroupIds,omitnil,omitempty" name:"AccessGroupIds"`
}

type RestoreTask struct {
	// 回热任务ID
	RestoreTaskId *uint64 `json:"RestoreTaskId,omitnil,omitempty" name:"RestoreTaskId"`

	// 回热任务文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 回热任务类型（1：标准；2：极速；3：批量，暂时仅支持极速）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 指定恢复出的临时副本的有效时长（单位天）
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 回热任务状态（1：绑定文件中；2：绑定文件完成；3：文件回热中；4：文件回热完成）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type Summary struct {
	// 已使用容量（byte）
	CapacityUsed *uint64 `json:"CapacityUsed,omitnil,omitempty" name:"CapacityUsed"`

	// 已使用COS标准存储容量（byte）
	StandardCapacityUsed *uint64 `json:"StandardCapacityUsed,omitnil,omitempty" name:"StandardCapacityUsed"`

	// 已使用COS低频存储容量（byte）
	DegradeCapacityUsed *uint64 `json:"DegradeCapacityUsed,omitnil,omitempty" name:"DegradeCapacityUsed"`

	// 已使用COS归档存储容量（byte）
	ArchiveCapacityUsed *uint64 `json:"ArchiveCapacityUsed,omitnil,omitempty" name:"ArchiveCapacityUsed"`

	// 已使用COS深度归档存储容量（byte）
	DeepArchiveCapacityUsed *uint64 `json:"DeepArchiveCapacityUsed,omitnil,omitempty" name:"DeepArchiveCapacityUsed"`

	// 已使用COS智能分层存储容量（byte）
	IntelligentCapacityUsed *uint64 `json:"IntelligentCapacityUsed,omitnil,omitempty" name:"IntelligentCapacityUsed"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Transition struct {
	// 触发时间（单位天）
	Days *uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 转换类型（1：归档；2：删除；3：低频；4：深度归档；5：智能分层）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}