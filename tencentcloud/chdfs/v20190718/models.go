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

package v20190718

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessGroup struct {

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// 权限组描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AccessRule struct {

	// 权限规则ID
	AccessRuleId *uint64 `json:"AccessRuleId,omitempty" name:"AccessRuleId"`

	// 权限规则地址（网段或IP）
	Address *string `json:"Address,omitempty" name:"Address"`

	// 权限规则访问模式（1：只读；2：读写）
	AccessMode *uint64 `json:"AccessMode,omitempty" name:"AccessMode"`

	// 优先级（取值范围1~100，值越小优先级越高）
	Priority *uint64 `json:"Priority,omitempty" name:"Priority"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type CreateAccessGroupRequest struct {
	*tchttp.BaseRequest

	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// 权限组描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限组
		AccessGroup *AccessGroup `json:"AccessGroup,omitempty" name:"AccessGroup"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccessRulesRequest struct {
	*tchttp.BaseRequest

	// 多个权限规则，上限为10
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules" list`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *CreateAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccessRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAccessRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFileSystemRequest struct {
	*tchttp.BaseRequest

	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// 文件系统容量（byte），下限为1M，上限为1P，且必须是1M的整数倍
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// 文件系统描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFileSystemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件系统
		FileSystem *FileSystem `json:"FileSystem,omitempty" name:"FileSystem"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFileSystemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMountPointRequest struct {
	*tchttp.BaseRequest

	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 挂载点状态（1：打开；2：关闭）
	MountPointStatus *uint64 `json:"MountPointStatus,omitempty" name:"MountPointStatus"`
}

func (r *CreateMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMountPointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMountPointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 挂载点
		MountPoint *MountPoint `json:"MountPoint,omitempty" name:"MountPoint"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMountPointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessGroupRequest struct {
	*tchttp.BaseRequest

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *DeleteAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessRulesRequest struct {
	*tchttp.BaseRequest

	// 多个权限规则ID，上限为10
	AccessRuleIds []*uint64 `json:"AccessRuleIds,omitempty" name:"AccessRuleIds" list`
}

func (r *DeleteAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccessRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAccessRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFileSystemRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DeleteFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFileSystemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFileSystemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMountPointRequest struct {
	*tchttp.BaseRequest

	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`
}

func (r *DeleteMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMountPointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMountPointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMountPointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，Name可选“AccessGroupId“和“AccessGroupName”，Values上限为10
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为所有
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限组列表
		AccessGroups []*AccessGroup `json:"AccessGroups,omitempty" name:"AccessGroups" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRulesRequest struct {
	*tchttp.BaseRequest

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为所有
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限规则列表
		AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccessRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`
}

func (r *DescribeFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileSystemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件系统
		FileSystem *FileSystem `json:"FileSystem,omitempty" name:"FileSystem"`

		// 文件系统已使用大小（byte）
		FileSystemCapacityUsed *uint64 `json:"FileSystemCapacityUsed,omitempty" name:"FileSystemCapacityUsed"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileSystemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemsRequest struct {
	*tchttp.BaseRequest

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为所有
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFileSystemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileSystemsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFileSystemsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文件系统列表
		FileSystems []*FileSystem `json:"FileSystems,omitempty" name:"FileSystems" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFileSystemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFileSystemsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMountPointRequest struct {
	*tchttp.BaseRequest

	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`
}

func (r *DescribeMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMountPointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMountPointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 挂载点
		MountPoint *MountPoint `json:"MountPoint,omitempty" name:"MountPoint"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMountPointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMountPointsRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID
	// 注意：若根据AccessGroupId查看挂载点列表，则无需设置FileSystemId
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 权限组ID
	// 注意：若根据FileSystemId查看挂载点列表，则无需设置AccessGroupId
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为所有
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMountPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMountPointsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMountPointsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 挂载点列表
		MountPoints []*MountPoint `json:"MountPoints,omitempty" name:"MountPoints" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMountPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMountPointsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FileSystem struct {

	// appid
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// 文件系统描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 文件系统块大小（byte）
	BlockSize *uint64 `json:"BlockSize,omitempty" name:"BlockSize"`

	// 文件系统容量（byte）
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`

	// 文件系统状态（1：创建中；2：创建成功；3：创建失败）
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type Filter struct {

	// 过滤字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ModifyAccessGroupRequest struct {
	*tchttp.BaseRequest

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// 权限组名称
	AccessGroupName *string `json:"AccessGroupName,omitempty" name:"AccessGroupName"`

	// 权限组描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAccessGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccessGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccessGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessRulesRequest struct {
	*tchttp.BaseRequest

	// 多个权限规则，上限为10
	AccessRules []*AccessRule `json:"AccessRules,omitempty" name:"AccessRules" list`
}

func (r *ModifyAccessRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccessRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAccessRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAccessRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAccessRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFileSystemRequest struct {
	*tchttp.BaseRequest

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 文件系统名称
	FileSystemName *string `json:"FileSystemName,omitempty" name:"FileSystemName"`

	// 文件系统描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 文件系统容量（byte），下限为1M，上限为1P，且必须是1M的整数倍
	CapacityQuota *uint64 `json:"CapacityQuota,omitempty" name:"CapacityQuota"`
}

func (r *ModifyFileSystemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFileSystemRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyFileSystemResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFileSystemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyFileSystemResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMountPointRequest struct {
	*tchttp.BaseRequest

	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// 挂载点状态
	MountPointStatus *uint64 `json:"MountPointStatus,omitempty" name:"MountPointStatus"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`
}

func (r *ModifyMountPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMountPointRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMountPointResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMountPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMountPointResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MountPoint struct {

	// 挂载点ID
	MountPointId *string `json:"MountPointId,omitempty" name:"MountPointId"`

	// 挂载点名称
	MountPointName *string `json:"MountPointName,omitempty" name:"MountPointName"`

	// 文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" name:"FileSystemId"`

	// 权限组ID
	AccessGroupId *string `json:"AccessGroupId,omitempty" name:"AccessGroupId"`

	// VPC网络ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 挂载点状态（1：打开；2：关闭）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}
