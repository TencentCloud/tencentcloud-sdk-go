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

package v20220331

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AppAssociatedUserGroupIds struct {
	// 用户组id
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 应用id
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`
}

// Predefined struct for user
type CreateApiImportUserJobRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 导入的用户数据
	DataFlowUserCreateList []*ImportUser `json:"DataFlowUserCreateList,omitnil,omitempty" name:"DataFlowUserCreateList"`
}

type CreateApiImportUserJobRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 导入的用户数据
	DataFlowUserCreateList []*ImportUser `json:"DataFlowUserCreateList,omitnil,omitempty" name:"DataFlowUserCreateList"`
}

func (r *CreateApiImportUserJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiImportUserJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "DataFlowUserCreateList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiImportUserJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiImportUserJobResponseParams struct {
	// 数据流任务
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApiImportUserJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiImportUserJobResponseParams `json:"Response"`
}

func (r *CreateApiImportUserJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiImportUserJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileExportUserJobRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 导出的数据类型
	// 
	// <li> **NDJSON** </li>  New-line Delimited JSON
	// <li> **CSV** </li>  Comma-Separated Values
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Key可选值为condition、userGroupId
	// 
	// <li> **condition** </li>	Values = 查询条件，用户ID，用户名称，手机或邮箱
	// <li> **userGroupId** </li>	Values = 用户组ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 导出用户包含的属性和映射名称，为空时包含所有的属性
	ExportPropertyMaps []*ExportPropertyMap `json:"ExportPropertyMaps,omitnil,omitempty" name:"ExportPropertyMaps"`
}

type CreateFileExportUserJobRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 导出的数据类型
	// 
	// <li> **NDJSON** </li>  New-line Delimited JSON
	// <li> **CSV** </li>  Comma-Separated Values
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Key可选值为condition、userGroupId
	// 
	// <li> **condition** </li>	Values = 查询条件，用户ID，用户名称，手机或邮箱
	// <li> **userGroupId** </li>	Values = 用户组ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 导出用户包含的属性和映射名称，为空时包含所有的属性
	ExportPropertyMaps []*ExportPropertyMap `json:"ExportPropertyMaps,omitnil,omitempty" name:"ExportPropertyMaps"`
}

func (r *CreateFileExportUserJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileExportUserJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Format")
	delete(f, "Filters")
	delete(f, "ExportPropertyMaps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFileExportUserJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFileExportUserJobResponseParams struct {
	// 数据流任务
	Job *Job `json:"Job,omitnil,omitempty" name:"Job"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateFileExportUserJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateFileExportUserJobResponseParams `json:"Response"`
}

func (r *CreateFileExportUserJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFileExportUserJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupRequestParams struct {
	// 用户组名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	delete(f, "DisplayName")
	delete(f, "UserStoreId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupResponseParams struct {
	// 用户组ID
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

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
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 用户组ID
	UserGroup []*string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 生日
	Birthdate *int64 `json:"Birthdate,omitnil,omitempty" name:"Birthdate"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitnil,omitempty" name:"CustomizationAttributes"`

	// 索引字段1
	IndexedAttribute1 *string `json:"IndexedAttribute1,omitnil,omitempty" name:"IndexedAttribute1"`

	// 索引字段2
	IndexedAttribute2 *string `json:"IndexedAttribute2,omitnil,omitempty" name:"IndexedAttribute2"`

	// 索引字段3
	IndexedAttribute3 *string `json:"IndexedAttribute3,omitnil,omitempty" name:"IndexedAttribute3"`

	// 索引字段4
	IndexedAttribute4 *string `json:"IndexedAttribute4,omitnil,omitempty" name:"IndexedAttribute4"`

	// 索引字段5
	IndexedAttribute5 *string `json:"IndexedAttribute5,omitnil,omitempty" name:"IndexedAttribute5"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 用户组ID
	UserGroup []*string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 生日
	Birthdate *int64 `json:"Birthdate,omitnil,omitempty" name:"Birthdate"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitnil,omitempty" name:"CustomizationAttributes"`

	// 索引字段1
	IndexedAttribute1 *string `json:"IndexedAttribute1,omitnil,omitempty" name:"IndexedAttribute1"`

	// 索引字段2
	IndexedAttribute2 *string `json:"IndexedAttribute2,omitnil,omitempty" name:"IndexedAttribute2"`

	// 索引字段3
	IndexedAttribute3 *string `json:"IndexedAttribute3,omitnil,omitempty" name:"IndexedAttribute3"`

	// 索引字段4
	IndexedAttribute4 *string `json:"IndexedAttribute4,omitnil,omitempty" name:"IndexedAttribute4"`

	// 索引字段5
	IndexedAttribute5 *string `json:"IndexedAttribute5,omitnil,omitempty" name:"IndexedAttribute5"`
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
	delete(f, "UserStoreId")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Password")
	delete(f, "UserName")
	delete(f, "Nickname")
	delete(f, "Address")
	delete(f, "UserGroup")
	delete(f, "Birthdate")
	delete(f, "CustomizationAttributes")
	delete(f, "IndexedAttribute1")
	delete(f, "IndexedAttribute2")
	delete(f, "IndexedAttribute3")
	delete(f, "IndexedAttribute4")
	delete(f, "IndexedAttribute5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 创建的用户信息
	User *User `json:"User,omitnil,omitempty" name:"User"`

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
type CreateUserStoreRequestParams struct {
	// 用户池名字
	UserPoolName *string `json:"UserPoolName,omitnil,omitempty" name:"UserPoolName"`

	// 用户池描述
	UserPoolDesc *string `json:"UserPoolDesc,omitnil,omitempty" name:"UserPoolDesc"`

	// 用户池logo
	UserPoolLogo *string `json:"UserPoolLogo,omitnil,omitempty" name:"UserPoolLogo"`
}

type CreateUserStoreRequest struct {
	*tchttp.BaseRequest
	
	// 用户池名字
	UserPoolName *string `json:"UserPoolName,omitnil,omitempty" name:"UserPoolName"`

	// 用户池描述
	UserPoolDesc *string `json:"UserPoolDesc,omitnil,omitempty" name:"UserPoolDesc"`

	// 用户池logo
	UserPoolLogo *string `json:"UserPoolLogo,omitnil,omitempty" name:"UserPoolLogo"`
}

func (r *CreateUserStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserPoolName")
	delete(f, "UserPoolDesc")
	delete(f, "UserPoolLogo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserStoreResponseParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserStoreResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserStoreResponseParams `json:"Response"`
}

func (r *CreateUserStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupsRequestParams struct {
	// 用户组ID数组
	UserGroupIds []*string `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID数组
	UserGroupIds []*string `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`
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
	delete(f, "UserGroupIds")
	delete(f, "UserStoreId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupsResponseParams struct {
	// 删除的用户组关联的应用信息
	UserGroupDeletedInfo *UserGroupDeleteResp `json:"UserGroupDeletedInfo,omitnil,omitempty" name:"UserGroupDeletedInfo"`

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
type DeleteUserStoreRequestParams struct {
	// 用户池ID
	UserPoolId *string `json:"UserPoolId,omitnil,omitempty" name:"UserPoolId"`
}

type DeleteUserStoreRequest struct {
	*tchttp.BaseRequest
	
	// 用户池ID
	UserPoolId *string `json:"UserPoolId,omitnil,omitempty" name:"UserPoolId"`
}

func (r *DeleteUserStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserPoolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserStoreResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserStoreResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserStoreResponseParams `json:"Response"`
}

func (r *DeleteUserStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID数组
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID数组
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
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
	delete(f, "UserStoreId")
	delete(f, "UserIds")
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

// Predefined struct for user
type DescribeUserByIdRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 返回信息是否为原文
	// 
	// <li> **false** </li>	默认，返回信息为脱敏信息
	// <li> **true** </li>	返回用户信息原文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`
}

type DescribeUserByIdRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 返回信息是否为原文
	// 
	// <li> **false** </li>	默认，返回信息为脱敏信息
	// <li> **true** </li>	返回用户信息原文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`
}

func (r *DescribeUserByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Original")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserByIdResponseParams struct {
	// 用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserByIdResponseParams `json:"Response"`
}

func (r *DescribeUserByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 查询条件，根据propertycode和propertykey
	Filters []*QueryUserFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否返回明文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`

	// 排序设置
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 查询条件，根据propertycode和propertykey
	Filters []*QueryUserFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否返回明文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`

	// 排序设置
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "Filters")
	delete(f, "Original")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserResponseParams struct {
	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 用户列表
	Content []*User `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserResponseParams `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorDetails struct {
	// 用户信息
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 失败原因
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`
}

type ExportPropertyMap struct {
	// 用户属性code
	UserPropertyCode *string `json:"UserPropertyCode,omitnil,omitempty" name:"UserPropertyCode"`

	// 用户属性映射名称
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`
}

type FailedUsers struct {
	// 失败用户标识
	FailedUserIdentification *string `json:"FailedUserIdentification,omitnil,omitempty" name:"FailedUserIdentification"`

	// 导入的用户失败原因
	FailedReason *string `json:"FailedReason,omitnil,omitempty" name:"FailedReason"`
}

type Filter struct {
	// key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 逻辑值
	Logic *bool `json:"Logic,omitnil,omitempty" name:"Logic"`
}

type ImportUser struct {
	// 用户名
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 身份证号
	ResidentIdentityCard *string `json:"ResidentIdentityCard,omitnil,omitempty" name:"ResidentIdentityCard"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 用户组ID
	UserGroup []*string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// QQ qqOpenId
	QqOpenId *string `json:"QqOpenId,omitnil,omitempty" name:"QqOpenId"`

	// QQ qqUnionId
	QqUnionId *string `json:"QqUnionId,omitnil,omitempty" name:"QqUnionId"`

	// 微信wechatOpenId
	WechatOpenId *string `json:"WechatOpenId,omitnil,omitempty" name:"WechatOpenId"`

	// 微信wechatUnionId
	WechatUnionId *string `json:"WechatUnionId,omitnil,omitempty" name:"WechatUnionId"`

	// 支付宝alipayUserId
	AlipayUserId *string `json:"AlipayUserId,omitnil,omitempty" name:"AlipayUserId"`

	// 企业微信weComUserId
	WeComUserId *string `json:"WeComUserId,omitnil,omitempty" name:"WeComUserId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 生日
	Birthdate *string `json:"Birthdate,omitnil,omitempty" name:"Birthdate"`

	// 姓名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 坐标
	Locale *string `json:"Locale,omitnil,omitempty" name:"Locale"`

	// 性别（MALE;FEMALE;UNKNOWN）
	Gender *string `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 实名核验方式
	IdentityVerificationMethod *string `json:"IdentityVerificationMethod,omitnil,omitempty" name:"IdentityVerificationMethod"`

	// 是否已实名核验
	IdentityVerified *bool `json:"IdentityVerified,omitnil,omitempty" name:"IdentityVerified"`

	// 工作
	Job *string `json:"Job,omitnil,omitempty" name:"Job"`

	// 国家
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// 时区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 密码密文
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitnil,omitempty" name:"CustomizationAttributes"`

	// 密码盐
	Salt *Salt `json:"Salt,omitnil,omitempty" name:"Salt"`

	// 密码加密方式（SHA1;BCRYPT）
	PasswordEncryptTypeEnum *string `json:"PasswordEncryptTypeEnum,omitnil,omitempty" name:"PasswordEncryptTypeEnum"`

	// 索引字段1
	IndexedAttribute1 *string `json:"IndexedAttribute1,omitnil,omitempty" name:"IndexedAttribute1"`

	// 索引字段2
	IndexedAttribute2 *string `json:"IndexedAttribute2,omitnil,omitempty" name:"IndexedAttribute2"`

	// 索引字段3
	IndexedAttribute3 *string `json:"IndexedAttribute3,omitnil,omitempty" name:"IndexedAttribute3"`

	// 索引字段4
	IndexedAttribute4 *string `json:"IndexedAttribute4,omitnil,omitempty" name:"IndexedAttribute4"`

	// 索引字段5
	IndexedAttribute5 *string `json:"IndexedAttribute5,omitnil,omitempty" name:"IndexedAttribute5"`
}

type Job struct {
	// 任务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务状态
	// 
	// <li> **PENDING** </li>  待执行
	// <li> **PROCESSING** </li>  执行中
	// <li> **COMPLETED** </li>  完成
	// <li> **FAILED** </li>  失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型
	// 
	// <li> **IMPORT_USER** </li>  用户导入
	// <li> **EXPORT_USER** </li>  用户导出
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 任务创建时间
	CreatedDate *int64 `json:"CreatedDate,omitnil,omitempty" name:"CreatedDate"`

	// 任务的数据类型
	// 
	// <li> **NDJSON** </li>  New-line Delimited JSON
	// <li> **CSV** </li>  Comma-Separated Values
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 任务结果下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 失败详情
	ErrorDetails []*ErrorDetails `json:"ErrorDetails,omitnil,omitempty" name:"ErrorDetails"`

	// 失败的用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedUsers []*FailedUsers `json:"FailedUsers,omitnil,omitempty" name:"FailedUsers"`
}

// Predefined struct for user
type LinkAccountRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 主用户ID
	PrimaryUserId *string `json:"PrimaryUserId,omitnil,omitempty" name:"PrimaryUserId"`

	// 从用户ID
	SecondaryUserId *string `json:"SecondaryUserId,omitnil,omitempty" name:"SecondaryUserId"`

	// 融合属性
	// 
	// <li> **PHONENUMBER** </li>	  手机号码
	// <li> **EMAIL** </li>  邮箱
	UserLinkedOnAttribute *string `json:"UserLinkedOnAttribute,omitnil,omitempty" name:"UserLinkedOnAttribute"`
}

type LinkAccountRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 主用户ID
	PrimaryUserId *string `json:"PrimaryUserId,omitnil,omitempty" name:"PrimaryUserId"`

	// 从用户ID
	SecondaryUserId *string `json:"SecondaryUserId,omitnil,omitempty" name:"SecondaryUserId"`

	// 融合属性
	// 
	// <li> **PHONENUMBER** </li>	  手机号码
	// <li> **EMAIL** </li>  邮箱
	UserLinkedOnAttribute *string `json:"UserLinkedOnAttribute,omitnil,omitempty" name:"UserLinkedOnAttribute"`
}

func (r *LinkAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PrimaryUserId")
	delete(f, "SecondaryUserId")
	delete(f, "UserLinkedOnAttribute")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LinkAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LinkAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LinkAccountResponse struct {
	*tchttp.BaseResponse
	Response *LinkAccountResponseParams `json:"Response"`
}

func (r *LinkAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LinkAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJobsRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 任务ID列表，为空时返回全部任务
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`
}

type ListJobsRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 任务ID列表，为空时返回全部任务
	JobIds []*string `json:"JobIds,omitnil,omitempty" name:"JobIds"`
}

func (r *ListJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "JobIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJobsResponseParams struct {
	// 任务列表
	JobSet []*Job `json:"JobSet,omitnil,omitempty" name:"JobSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListJobsResponse struct {
	*tchttp.BaseResponse
	Response *ListJobsResponseParams `json:"Response"`
}

func (r *ListJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogMessageByConditionRequestParams struct {
	// 用户池ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 开始时间，时间戳精确到毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Key可选值为events
	// 
	// <li> **events** </li>	Values为["SIGNUP", "USER_UPDATE", "USER_DELETE", "USER_CREATE", "ACCOUNT_LINKING"] 中的一个或多个
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ListLogMessageByConditionRequest struct {
	*tchttp.BaseRequest
	
	// 用户池ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 开始时间，时间戳精确到毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Key可选值为events
	// 
	// <li> **events** </li>	Values为["SIGNUP", "USER_UPDATE", "USER_DELETE", "USER_CREATE", "ACCOUNT_LINKING"] 中的一个或多个
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ListLogMessageByConditionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogMessageByConditionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "StartTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListLogMessageByConditionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListLogMessageByConditionResponseParams struct {
	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分页对象
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 日志列表
	Content []*LogMessage `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListLogMessageByConditionResponse struct {
	*tchttp.BaseResponse
	Response *ListLogMessageByConditionResponseParams `json:"Response"`
}

func (r *ListLogMessageByConditionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListLogMessageByConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserByPropertyRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 查询的属性
	// 
	// <li> **phoneNumber** </li>	  手机号码
	// <li> **email** </li>  邮箱
	PropertyCode *string `json:"PropertyCode,omitnil,omitempty" name:"PropertyCode"`

	// 属性值
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`

	// 返回信息是否为原文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`
}

type ListUserByPropertyRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 查询的属性
	// 
	// <li> **phoneNumber** </li>	  手机号码
	// <li> **email** </li>  邮箱
	PropertyCode *string `json:"PropertyCode,omitnil,omitempty" name:"PropertyCode"`

	// 属性值
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`

	// 返回信息是否为原文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`
}

func (r *ListUserByPropertyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserByPropertyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "PropertyCode")
	delete(f, "PropertyValue")
	delete(f, "Original")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserByPropertyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserByPropertyResponseParams struct {
	// 用户列表
	Users []*User `json:"Users,omitnil,omitempty" name:"Users"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUserByPropertyResponse struct {
	*tchttp.BaseResponse
	Response *ListUserByPropertyResponseParams `json:"Response"`
}

func (r *ListUserByPropertyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserByPropertyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// Key可选值为condition
	// 
	// <li> **condition** </li>	Values = 查询条件，用户组ID或用户组名称
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ListUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// Key可选值为condition
	// 
	// <li> **condition** </li>	Values = 查询条件，用户组ID或用户组名称
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ListUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserGroupsResponseParams struct {
	// 用户组列表
	Content []*UserGroup `json:"Content,omitnil,omitempty" name:"Content"`

	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分页
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListUserGroupsResponseParams `json:"Response"`
}

func (r *ListUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// Key可选值为condition、userGroupId
	// 
	// <li> **condition** </li>	Values = 查询条件，用户ID，用户名称，手机或邮箱
	// <li> **userGroupId** </li>	Values = 用户组ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回信息是否为原文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`
}

type ListUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 分页数据
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// Key可选值为condition、userGroupId
	// 
	// <li> **condition** </li>	Values = 查询条件，用户ID，用户名称，手机或邮箱
	// <li> **userGroupId** </li>	Values = 用户组ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回信息是否为原文
	Original *bool `json:"Original,omitnil,omitempty" name:"Original"`
}

func (r *ListUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "Pageable")
	delete(f, "Filters")
	delete(f, "Original")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserResponseParams struct {
	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分页对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pageable *Pageable `json:"Pageable,omitnil,omitempty" name:"Pageable"`

	// 用户列表
	Content []*User `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUserResponse struct {
	*tchttp.BaseResponse
	Response *ListUserResponseParams `json:"Response"`
}

func (r *ListUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserStoreRequestParams struct {

}

type ListUserStoreRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListUserStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserStoreResponseParams struct {
	// 用户目录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserStoreSet []*UserStore `json:"UserStoreSet,omitnil,omitempty" name:"UserStoreSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUserStoreResponse struct {
	*tchttp.BaseResponse
	Response *ListUserStoreResponseParams `json:"Response"`
}

func (r *ListUserStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogMessage struct {
	// 日志标识
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 用户池ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 事件编码
	EventCode *string `json:"EventCode,omitnil,omitempty" name:"EventCode"`

	// 事件发生时间戳，单位：毫秒
	EventDate *int64 `json:"EventDate,omitnil,omitempty" name:"EventDate"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 事件参与者
	// 
	// <li> **TENANT** </li>  租户
	// <li> **USER** </li>  用户
	Participant *string `json:"Participant,omitnil,omitempty" name:"Participant"`

	// 应用clientId
	ApplicationClientId *string `json:"ApplicationClientId,omitnil,omitempty" name:"ApplicationClientId"`

	// 应用名称
	ApplicationName *string `json:"ApplicationName,omitnil,omitempty" name:"ApplicationName"`

	// 认证源ID
	AuthSourceId *string `json:"AuthSourceId,omitnil,omitempty" name:"AuthSourceId"`

	// 认证源名称
	AuthSourceName *string `json:"AuthSourceName,omitnil,omitempty" name:"AuthSourceName"`

	// 认证源类型
	AuthSourceType *string `json:"AuthSourceType,omitnil,omitempty" name:"AuthSourceType"`

	// 认证源类别
	AuthSourceCategory *string `json:"AuthSourceCategory,omitnil,omitempty" name:"AuthSourceCategory"`

	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 用户代理
	UserAgent *string `json:"UserAgent,omitnil,omitempty" name:"UserAgent"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 日志结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionResult *string `json:"ActionResult,omitnil,omitempty" name:"ActionResult"`
}

type MemberMap struct {
	// 健
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Pageable struct {
	// 每页数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 当前页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type QueryUserFilter struct {
	// 属性key
	PropertyKey *string `json:"PropertyKey,omitnil,omitempty" name:"PropertyKey"`

	// 属性value
	PropertyValue *string `json:"PropertyValue,omitnil,omitempty" name:"PropertyValue"`

	// 逻辑值，等于true，不等于false
	Logic *bool `json:"Logic,omitnil,omitempty" name:"Logic"`

	// 操作逻辑符（支持> < = >= <=  != between）
	OperateLogic *string `json:"OperateLogic,omitnil,omitempty" name:"OperateLogic"`
}

// Predefined struct for user
type ResetPasswordRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`
}

type ResetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`
}

func (r *ResetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserStoreId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetPasswordResponseParams struct {
	// 重置后的用户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *ResetPasswordResponseParams `json:"Response"`
}

func (r *ResetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Salt struct {
	// 盐值
	SaltValue *string `json:"SaltValue,omitnil,omitempty" name:"SaltValue"`

	// 盐值位置
	SaltLocation *SaltLocation `json:"SaltLocation,omitnil,omitempty" name:"SaltLocation"`
}

type SaltLocation struct {
	// 密码加盐的类型（HEAD，TAIL，OTHER）
	SaltLocationTypeEnum *string `json:"SaltLocationTypeEnum,omitnil,omitempty" name:"SaltLocationTypeEnum"`

	// 加盐规则
	SaltLocationRule *SaltLocationRule `json:"SaltLocationRule,omitnil,omitempty" name:"SaltLocationRule"`
}

type SaltLocationRule struct {
	// 表达式
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

// Predefined struct for user
type SetPasswordRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type SetPasswordRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *SetPasswordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetPasswordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetPasswordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetPasswordResponse struct {
	*tchttp.BaseResponse
	Response *SetPasswordResponseParams `json:"Response"`
}

func (r *SetPasswordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {
	// 排序字段的key，参考自定义属性
	PropertyKey *string `json:"PropertyKey,omitnil,omitempty" name:"PropertyKey"`

	// 升序或者降序，ASC/DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

// Predefined struct for user
type UpdateUserGroupRequestParams struct {
	// 用户组ID
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 用户组名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 用户组名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	delete(f, "DisplayName")
	delete(f, "UserStoreId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserGroupResponseParams `json:"Response"`
}

func (r *UpdateUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 用户组
	UserGroup []*string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 生日
	Birthdate *int64 `json:"Birthdate,omitnil,omitempty" name:"Birthdate"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitnil,omitempty" name:"CustomizationAttributes"`

	// 索引字段1
	IndexedAttribute1 *string `json:"IndexedAttribute1,omitnil,omitempty" name:"IndexedAttribute1"`

	// 索引字段2
	IndexedAttribute2 *string `json:"IndexedAttribute2,omitnil,omitempty" name:"IndexedAttribute2"`

	// 索引字段3
	IndexedAttribute3 *string `json:"IndexedAttribute3,omitnil,omitempty" name:"IndexedAttribute3"`

	// 索引字段4
	IndexedAttribute4 *string `json:"IndexedAttribute4,omitnil,omitempty" name:"IndexedAttribute4"`

	// 索引字段5
	IndexedAttribute5 *string `json:"IndexedAttribute5,omitnil,omitempty" name:"IndexedAttribute5"`
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户名称
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 用户组
	UserGroup []*string `json:"UserGroup,omitnil,omitempty" name:"UserGroup"`

	// 生日
	Birthdate *int64 `json:"Birthdate,omitnil,omitempty" name:"Birthdate"`

	// 自定义属性
	CustomizationAttributes []*MemberMap `json:"CustomizationAttributes,omitnil,omitempty" name:"CustomizationAttributes"`

	// 索引字段1
	IndexedAttribute1 *string `json:"IndexedAttribute1,omitnil,omitempty" name:"IndexedAttribute1"`

	// 索引字段2
	IndexedAttribute2 *string `json:"IndexedAttribute2,omitnil,omitempty" name:"IndexedAttribute2"`

	// 索引字段3
	IndexedAttribute3 *string `json:"IndexedAttribute3,omitnil,omitempty" name:"IndexedAttribute3"`

	// 索引字段4
	IndexedAttribute4 *string `json:"IndexedAttribute4,omitnil,omitempty" name:"IndexedAttribute4"`

	// 索引字段5
	IndexedAttribute5 *string `json:"IndexedAttribute5,omitnil,omitempty" name:"IndexedAttribute5"`
}

func (r *UpdateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserStoreId")
	delete(f, "UserName")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Nickname")
	delete(f, "Address")
	delete(f, "UserGroup")
	delete(f, "Birthdate")
	delete(f, "CustomizationAttributes")
	delete(f, "IndexedAttribute1")
	delete(f, "IndexedAttribute2")
	delete(f, "IndexedAttribute3")
	delete(f, "IndexedAttribute4")
	delete(f, "IndexedAttribute5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserResponseParams struct {
	// 更新之后的用户信息
	User *User `json:"User,omitnil,omitempty" name:"User"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserResponseParams `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStatusRequestParams struct {
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户状态
	// 
	// <li> **NORMAL** </li>	  正常
	// <li> **LOCK** </li>  锁定
	// <li> **FREEZE** </li>	  冻结
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateUserStatusRequest struct {
	*tchttp.BaseRequest
	
	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户状态
	// 
	// <li> **NORMAL** </li>	  正常
	// <li> **LOCK** </li>  锁定
	// <li> **FREEZE** </li>	  冻结
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdateUserStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserStoreId")
	delete(f, "UserId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserStatusResponseParams `json:"Response"`
}

func (r *UpdateUserStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStoreRequestParams struct {
	// 用户池ID
	UserPoolId *string `json:"UserPoolId,omitnil,omitempty" name:"UserPoolId"`

	// 用户池名字
	UserPoolName *string `json:"UserPoolName,omitnil,omitempty" name:"UserPoolName"`

	// 用户池描述
	UserPoolDesc *string `json:"UserPoolDesc,omitnil,omitempty" name:"UserPoolDesc"`

	// 用户池logo
	UserPoolLogo *string `json:"UserPoolLogo,omitnil,omitempty" name:"UserPoolLogo"`
}

type UpdateUserStoreRequest struct {
	*tchttp.BaseRequest
	
	// 用户池ID
	UserPoolId *string `json:"UserPoolId,omitnil,omitempty" name:"UserPoolId"`

	// 用户池名字
	UserPoolName *string `json:"UserPoolName,omitnil,omitempty" name:"UserPoolName"`

	// 用户池描述
	UserPoolDesc *string `json:"UserPoolDesc,omitnil,omitempty" name:"UserPoolDesc"`

	// 用户池logo
	UserPoolLogo *string `json:"UserPoolLogo,omitnil,omitempty" name:"UserPoolLogo"`
}

func (r *UpdateUserStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserPoolId")
	delete(f, "UserPoolName")
	delete(f, "UserPoolDesc")
	delete(f, "UserPoolLogo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserStoreResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserStoreResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserStoreResponseParams `json:"Response"`
}

func (r *UpdateUserStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 手机号
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 上次登录时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSignOn *int64 `json:"LastSignOn,omitnil,omitempty" name:"LastSignOn"`

	// 创建时间
	CreatedDate *int64 `json:"CreatedDate,omitnil,omitempty" name:"CreatedDate"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 用户来源
	UserDataSourceEnum *string `json:"UserDataSourceEnum,omitnil,omitempty" name:"UserDataSourceEnum"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 生日
	// 注意：此字段可能返回 null，表示取不到有效值。
	Birthdate *int64 `json:"Birthdate,omitnil,omitempty" name:"Birthdate"`

	// 用户组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroups []*string `json:"UserGroups,omitnil,omitempty" name:"UserGroups"`

	// 上次修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedDate *int64 `json:"LastModifiedDate,omitnil,omitempty" name:"LastModifiedDate"`

	// 自定义属性
	CustomAttributes []*MemberMap `json:"CustomAttributes,omitnil,omitempty" name:"CustomAttributes"`

	// 身份证号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResidentIdentityCard *string `json:"ResidentIdentityCard,omitnil,omitempty" name:"ResidentIdentityCard"`

	// QQ的OpenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	QqOpenId *string `json:"QqOpenId,omitnil,omitempty" name:"QqOpenId"`

	// QQ的UnionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	QqUnionId *string `json:"QqUnionId,omitnil,omitempty" name:"QqUnionId"`

	// 微信的WechatOpenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatOpenId *string `json:"WechatOpenId,omitnil,omitempty" name:"WechatOpenId"`

	// 微信的WechatUnionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatUnionId *string `json:"WechatUnionId,omitnil,omitempty" name:"WechatUnionId"`

	// 支付宝的AlipayUserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlipayUserId *string `json:"AlipayUserId,omitnil,omitempty" name:"AlipayUserId"`

	// 企业微信的WeComUserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeComUserId *string `json:"WeComUserId,omitnil,omitempty" name:"WeComUserId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locale *string `json:"Locale,omitnil,omitempty" name:"Locale"`

	// 性别
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gender *string `json:"Gender,omitnil,omitempty" name:"Gender"`

	// 实名核验方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityVerificationMethod *string `json:"IdentityVerificationMethod,omitnil,omitempty" name:"IdentityVerificationMethod"`

	// 是否已经实名核验
	IdentityVerified *bool `json:"IdentityVerified,omitnil,omitempty" name:"IdentityVerified"`

	// 工作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *string `json:"Job,omitnil,omitempty" name:"Job"`

	// 国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nationality *string `json:"Nationality,omitnil,omitempty" name:"Nationality"`

	// 是否主账号（进行过账号融合后，主账号为true，从账号为false）
	Primary *bool `json:"Primary,omitnil,omitempty" name:"Primary"`

	// 时区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 是否已经首次登录
	AlreadyFirstLogin *bool `json:"AlreadyFirstLogin,omitnil,omitempty" name:"AlreadyFirstLogin"`

	// 租户id
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 用户目录id
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 锁定类型（分为管理员锁定，和登录策略锁定）
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockType *string `json:"LockType,omitnil,omitempty" name:"LockType"`

	// 锁定时间点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockTime *int64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// 索引字段1
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexedAttribute1 *string `json:"IndexedAttribute1,omitnil,omitempty" name:"IndexedAttribute1"`

	// 索引字段2
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexedAttribute2 *string `json:"IndexedAttribute2,omitnil,omitempty" name:"IndexedAttribute2"`

	// 索引字段3
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexedAttribute3 *string `json:"IndexedAttribute3,omitnil,omitempty" name:"IndexedAttribute3"`

	// 索引字段4
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexedAttribute4 *string `json:"IndexedAttribute4,omitnil,omitempty" name:"IndexedAttribute4"`

	// 索引字段5
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexedAttribute5 *string `json:"IndexedAttribute5,omitnil,omitempty" name:"IndexedAttribute5"`
}

type UserGroup struct {
	// 用户组ID
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 用户组名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户组描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户目录ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 创建时间
	CreatedDate *int64 `json:"CreatedDate,omitnil,omitempty" name:"CreatedDate"`

	// 最近更新时间
	LastModifyDate *int64 `json:"LastModifyDate,omitnil,omitempty" name:"LastModifyDate"`
}

type UserGroupDeleteResp struct {
	// 错误详情
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 用户组关联的应用信息
	AppAssociatedUserGroupIds []*AppAssociatedUserGroupIds `json:"AppAssociatedUserGroupIds,omitnil,omitempty" name:"AppAssociatedUserGroupIds"`
}

type UserStore struct {
	// 租户ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 用户池logo
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserStoreLogo *string `json:"UserStoreLogo,omitnil,omitempty" name:"UserStoreLogo"`

	// 用户池描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserStoreDesc *string `json:"UserStoreDesc,omitnil,omitempty" name:"UserStoreDesc"`

	// 用户池名称
	UserStoreName *string `json:"UserStoreName,omitnil,omitempty" name:"UserStoreName"`

	// 用户数量
	UserNum *int64 `json:"UserNum,omitnil,omitempty" name:"UserNum"`

	// 用户池ID
	UserStoreId *string `json:"UserStoreId,omitnil,omitempty" name:"UserStoreId"`

	// 应用数量
	AppNum *int64 `json:"AppNum,omitnil,omitempty" name:"AppNum"`

	// 上次切换的用户池
	LastStatus *bool `json:"LastStatus,omitnil,omitempty" name:"LastStatus"`

	// 默认用户池
	DefaultStatus *bool `json:"DefaultStatus,omitnil,omitempty" name:"DefaultStatus"`

	// 创建时间
	CreateDate *int64 `json:"CreateDate,omitnil,omitempty" name:"CreateDate"`

	// 上次切换时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastStatusTime *int64 `json:"LastStatusTime,omitnil,omitempty" name:"LastStatusTime"`

	// 用户目录域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserStoreProtocolHost *string `json:"UserStoreProtocolHost,omitnil,omitempty" name:"UserStoreProtocolHost"`
}