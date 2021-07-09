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

package v20210420

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddUserToUserGroupRequest struct {
	*tchttp.BaseRequest

	// 加入用户组的用户ID列表。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 用户组ID，是用户组的全局唯一标识。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *AddUserToUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserToUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddUserToUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserToUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplicationAuthorizationInfo struct {

	// 用户在被授权应用下对应的账号列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationAccounts []*string `json:"ApplicationAccounts,omitempty" name:"ApplicationAccounts"`

	// 应用ID，是应用的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 展示用户所在的用户组、机构节点拥有该应用的访问权限的ID信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InheritedForm *InheritedForm `json:"InheritedForm,omitempty" name:"InheritedForm"`
}

type CreateOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点名称，长度限制：64个字符。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 父机构节点ID，如果为空则默认创建在机构根节点下。
	ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

	// 机构节点描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户自定义可选填的机构节点对外ID，如果非空则校验此ID的唯一性。
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`
}

func (r *CreateOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DisplayName")
	delete(f, "ParentOrgNodeId")
	delete(f, "Description")
	delete(f, "CustomizedOrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机构节点ID，是机构节点的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest

	// 昵称，长度限制：64个字符。 DisplayName是唯一的。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 用户备注，长度限制：512个字符。
	Description *string `json:"Description,omitempty" name:"Description"`
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
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组ID，是用户组的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 用户名，长度限制：64个字符。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户密码， 需要符合密码策略的配置。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 昵称，长度限制：64个字符。 默认与用户名相同。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 用户备注，长度限制：512个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户所属用户组ID列表。
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// 用户手机号。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户所属组织机构唯一ID。如果为空，默认为在根节点下创建用户。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 用户过期时间，遵循 ISO 8601 标准。
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// 用户邮箱。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 密码是否需要重置，为空默认为false不需要重置密码。
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`
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
	delete(f, "Password")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "UserGroupIds")
	delete(f, "Phone")
	delete(f, "OrgNodeId")
	delete(f, "ExpirationTime")
	delete(f, "Email")
	delete(f, "PwdNeedReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的新创建的用户ID，是该用户的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserId *string `json:"UserId,omitempty" name:"UserId"`

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

type DeleteOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点ID，是机构节点的全局唯一标识。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`
}

func (r *DeleteOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组ID，是用户组的全局唯一标识。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *DeleteUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 用户名，长度限制：32个字符。 Username 和 UserId 需选择一个作为搜索条件；如两个条件同时使用则默认使用Username作为搜索条件。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户 id。 Username 和 UserId 需选择一个作为搜索条件；如两个条件同时使用则默认使用Username作为搜索条件。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用id，是应用的全局唯一标识，与ClientId参数不能同时为空。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 客户端id，与ApplicationId参数不能同时为空。
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
}

func (r *DescribeApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ClientId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥id。
	// 注意：此字段可能返回 null，表示取不到有效值。
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// 应用展示名称，长度限制：64个字符。 默认与应用名字相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

		// 应用最后修改时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
		LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

		// 客户端id。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientId *string `json:"ClientId,omitempty" name:"ClientId"`

		// 应用类型，即创建应用时所选择的应用模版类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

		// 应用创建时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

		// 应用id，是应用的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 令牌有效时间，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TokenExpired *int64 `json:"TokenExpired,omitempty" name:"TokenExpired"`

		// 客户端secret。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClientSecret *string `json:"ClientSecret,omitempty" name:"ClientSecret"`

		// 公钥信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

		// 授权地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AuthorizeUrl *string `json:"AuthorizeUrl,omitempty" name:"AuthorizeUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点ID，是机构节点全局唯一标识，长度限制：64个字符。如果为空默认读取机构根节点信息。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 是否读取其子节点信息。当读取层数为空或0时，默认仅读取当前机构节点信息。当读取层数为1时，读取本机构节点以及其第一层子节点信息。
	IncludeOrgNodeChildInfo *bool `json:"IncludeOrgNodeChildInfo,omitempty" name:"IncludeOrgNodeChildInfo"`
}

func (r *DescribeOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	delete(f, "IncludeOrgNodeChildInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机构节点展示名称，长度限制：64个字符。 默认与机构名相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

		// 机构节点最后修改时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
		LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

		// 机构节点外部ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`

		// 当前机构节点的父节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

		// 机构节点ID，是机构节点的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

		// 数据来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

		// 机构节点创建时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

		// 当前机构节点下的子节点列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgNodeChildInfo []*OrgNodeChildInfo `json:"OrgNodeChildInfo,omitempty" name:"OrgNodeChildInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组ID，是用户组的全局唯一标识。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *DescribeUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 昵称，长度限制：64个字符。 DisplayName不唯一。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

		// 用户备注，长度限制：512个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 用户组ID，是用户组的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest

	// 用户名，长度限制：64个字符。 Username 和 UserId 需至少一个不为空；都不为空时优先使用 Username。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户 id，长度限制：64个字符。 Username 和 UserId 需至少一个不为空；都不为空时优先使用 Username。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户名。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserName *string `json:"UserName,omitempty" name:"UserName"`

		// 用户状态，取值 NORMAL （正常）、FREEZE （已冻结）、LOCKED （已锁定）或 NOT_ENABLED （未启用）。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *string `json:"Status,omitempty" name:"Status"`

		// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
		DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

		// 用户备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 用户所属用户组 id 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

		// 用户 id，长度限制：64个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserId *string `json:"UserId,omitempty" name:"UserId"`

		// 用户邮箱。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Email *string `json:"Email,omitempty" name:"Email"`

		// 用户手机号。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Phone *string `json:"Phone,omitempty" name:"Phone"`

		// 用户所属组织机构 Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

		// 数据来源
	// 注意：此字段可能返回 null，表示取不到有效值。
		DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

		// 用户过期时间，遵循 ISO 8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

		// 用户激活时间，遵循 ISO 8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ActivationTime *string `json:"ActivationTime,omitempty" name:"ActivationTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InheritedForm struct {

	// 用户所在的用户组ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// 用户所在的机构节点ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgNodeIds []*string `json:"OrgNodeIds,omitempty" name:"OrgNodeIds"`
}

type ListAuthorizedApplicationsToOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点 Id 。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`
}

func (r *ListAuthorizedApplicationsToOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuthorizedApplicationsToOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAuthorizedApplicationsToOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机构节点拥有访问权限的应用 id 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAuthorizedApplicationsToOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAuthorizedApplicationsToUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组 Id 。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *ListAuthorizedApplicationsToUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuthorizedApplicationsToUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAuthorizedApplicationsToUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组拥有访问权限的应用 id 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationIds []*string `json:"ApplicationIds,omitempty" name:"ApplicationIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAuthorizedApplicationsToUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAuthorizedApplicationsToUserRequest struct {
	*tchttp.BaseRequest

	// 用户 ID。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 查询范围是否包括用户关联的用户组、组织机构的应用访问权限。默认为不查询 。传0表示不查询该范围，传1表示应用查询该范围。
	IncludeInheritedAuthorizations *bool `json:"IncludeInheritedAuthorizations,omitempty" name:"IncludeInheritedAuthorizations"`
}

func (r *ListAuthorizedApplicationsToUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "IncludeInheritedAuthorizations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuthorizedApplicationsToUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAuthorizedApplicationsToUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户拥有访问权限的应用信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationAuthorizationInfo []*ApplicationAuthorizationInfo `json:"ApplicationAuthorizationInfo,omitempty" name:"ApplicationAuthorizationInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAuthorizedApplicationsToUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuthorizedApplicationsToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUserGroupsOfUserRequest struct {
	*tchttp.BaseRequest

	// 用户ID，是用户的全局唯一标识。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *ListUserGroupsOfUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsOfUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserGroupsOfUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUserGroupsOfUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户所属的用户组ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

		// 用户ID，是用户的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserId *string `json:"UserId,omitempty" name:"UserId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUserGroupsOfUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserGroupsOfUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersInOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点ID，是机构节点全局唯一标识，长度限制：64个字符。如果为空默认读取机构根节点下用户信息。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 限制读取子节点信息层数。当读取层数为空或0时，默认仅读取当前机构节点信息。当读取层数为1时，读取本机构节点以及其第一层子节点信息。
	IncludeOrgNodeChildInfo *bool `json:"IncludeOrgNodeChildInfo,omitempty" name:"IncludeOrgNodeChildInfo"`
}

func (r *ListUsersInOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	delete(f, "IncludeOrgNodeChildInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersInOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersInOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机构子节点下的用户信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgNodeChildUserInfo []*OrgNodeChildUserInfo `json:"OrgNodeChildUserInfo,omitempty" name:"OrgNodeChildUserInfo"`

		// 机构ID，是机构节点全局唯一标识，长度限制：64个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

		// 用户信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

		// 当前机构节点下的用户总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalUserNum *int64 `json:"TotalUserNum,omitempty" name:"TotalUserNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersInOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersInUserGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组ID，是用户组的全局唯一标识。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *ListUsersInUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersInUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersInUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组ID，是用户组的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

		// 返回的用户信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

		// 返回的用户信息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersInUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersInUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserInfoRequest struct {
	*tchttp.BaseRequest

	// 用户名，长度限制：32个字符。 Username 和 UserId 需选择一个作为搜索条件；如两个条件同时使用则默认使用Username作为搜索条件。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 昵称，长度限制：64个字符。 默认与用户名相同。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 用户备注，长度限制：512个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户所属用户组ID列表。
	UserGroupIds []*string `json:"UserGroupIds,omitempty" name:"UserGroupIds"`

	// 用户 id。 Username 和 UserId 需选择一个作为搜索条件；如两个条件同时使用则默认使用Username作为搜索条件。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户手机号。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户过期时间，遵循 ISO 8601 标准。
	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`

	// 用户密码， 需要符合密码策略的配置。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 用户邮箱。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 密码是否需要重置，为空默认为false不需要重置密码。
	PwdNeedReset *bool `json:"PwdNeedReset,omitempty" name:"PwdNeedReset"`
}

func (r *ModifyUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserName")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "UserGroupIds")
	delete(f, "UserId")
	delete(f, "Phone")
	delete(f, "ExpirationTime")
	delete(f, "Password")
	delete(f, "Email")
	delete(f, "PwdNeedReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgNodeChildInfo struct {

	// 机构节点展示名称，长度限制：64个字符。 默认与机构名相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 机构节点最后修改时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// 用户自定义可选填的机构节点对外ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`

	// 当前机构节点的父节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentOrgNodeId *string `json:"ParentOrgNodeId,omitempty" name:"ParentOrgNodeId"`

	// 机构节点ID，是机构节点的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 数据来源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 机构节点创建时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`
}

type OrgNodeChildUserInfo struct {

	// 机构ID，是机构节点全局唯一标识，长度限制：64个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 用户信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserInfo []*UserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// 当前机构节点下的用户总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUserNum *int64 `json:"TotalUserNum,omitempty" name:"TotalUserNum"`
}

type RemoveUserFromUserGroupRequest struct {
	*tchttp.BaseRequest

	// 要加入用户组的用户ID列表。
	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`

	// 用户组ID，是用户组的全局唯一标识。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`
}

func (r *RemoveUserFromUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserFromUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserFromUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点ID，是机构节点的全局唯一标识。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 机构节点名称，长度限制：64个字符。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 机构节点描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户自定义可选填的机构节点对外ID，如果非空则校验此ID的唯一性。
	CustomizedOrgNodeId *string `json:"CustomizedOrgNodeId,omitempty" name:"CustomizedOrgNodeId"`
}

func (r *UpdateOrgNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrgNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgNodeId")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "CustomizedOrgNodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrgNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateOrgNodeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateOrgNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrgNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {

	// 用户ID，是用户全局唯一标识，长度限制：64个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 昵称，长度限制：64个字符。 默认与用户名相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}
