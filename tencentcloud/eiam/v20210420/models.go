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

type ApplicationInfoSearchCriteria struct {

	// 应用匹配搜索关键字，匹配范围包括：应用名称、应用ID。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 应用类型。ApplicationType的取值范围有：OAUTH2、JWT、CAS、SAML2、FORM、OIDC、APIGW。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`
}

type ApplicationInformation struct {

	// 应用ID，是应用的全局唯一标识。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 应用展示名称，长度限制：64个字符。 默认与应用名字相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 应用创建时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedDate *string `json:"CreatedDate,omitempty" name:"CreatedDate"`

	// 上次更新时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// 应用状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

	// 应用图标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitempty" name:"Icon"`

	// 应用类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicationType *string `json:"ApplicationType,omitempty" name:"ApplicationType"`

	// 客户端id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
}

type AuthorizationInfo struct {

	// 应用唯一ID。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 应用名称。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 类型名称。
	EntityName *string `json:"EntityName,omitempty" name:"EntityName"`

	// 类型唯一ID。
	EntityId *string `json:"EntityId,omitempty" name:"EntityId"`

	// 上次更新时间，符合 ISO8601 标准。
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`

	// 授权类型唯一ID。
	AuthorizationId *string `json:"AuthorizationId,omitempty" name:"AuthorizationId"`
}

type AuthorizationInfoSearchCriteria struct {

	// 名称匹配搜索，当查询类型为用户时，匹配范围包括：用户名称、应用名称；当查询类型为用户组时，匹配范围包括：用户组名称、应用名称；当查询类型为组织机构时，匹配范围包括：组织机构名称、应用名称。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
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

	// 用户手机号。例如：+86-1xxxxxxxxxx。
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

		// 应用图标图片访问地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

		// 安全等级。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

		// 应用状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

		// 描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

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

	// 是否读取其子节点信息。当其为空或false时，默认仅读取当前机构节点信息。当其为true时，读取本机构节点以及其第一层子节点信息。
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

type DescribePublicKeyRequest struct {
	*tchttp.BaseRequest

	// 应用ID，是应用的全局唯一标识。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *DescribePublicKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// jwt验证签名所用的公钥信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`

		// jwt的密钥id。
	// 注意：此字段可能返回 null，表示取不到有效值。
		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

		// 应用ID，是应用的全局唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublicKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicKeyResponse) FromJsonString(s string) error {
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

type ListApplicationAuthorizationsRequest struct {
	*tchttp.BaseRequest

	// 查询类型，包含用户（User）、用户组（UserGroup）、组织机构（OrgNode）。
	EntityType *string `json:"EntityType,omitempty" name:"EntityType"`

	// 查询条件，支持多搜索条件组合、多数据范围匹配的搜索。同时支持查询信息内容全匹配、部分匹配、范围匹配等多种查询方式，具体查询方式为：双引号（""）表示全匹配、以星号（* ) 结尾表示字段部分匹配。如果该字段为空，则默认查全量表。
	SearchCondition *AuthorizationInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// 排序条件集合。可排序的属性支持：上次修改时间（lastModifiedDate）。如果该字段为空，则默认按照应用名称正向排序。
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// 分页偏移量。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页读取数量。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListApplicationAuthorizationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationAuthorizationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EntityType")
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListApplicationAuthorizationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationAuthorizationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的应用授权信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AuthorizationInfoList []*AuthorizationInfo `json:"AuthorizationInfoList,omitempty" name:"AuthorizationInfoList"`

		// 返回的应用信息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationAuthorizationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationAuthorizationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationsRequest struct {
	*tchttp.BaseRequest

	// 查询条件，支持多搜索条件组合、多数据范围匹配的搜索。同时支持查询信息内容全匹配、部分匹配、范围匹配等多种查询方式，具体查询方式为：双引号（""）表示全匹配、以星号（* ) 结尾表示字段部分匹配。如果该字段为空，则默认查全量表。
	SearchCondition *ApplicationInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// 排序条件集合。可排序的属性支持：应用名字（displayName）、创建时间（createdDate）、上次修改时间（lastModifiedDate）。如果该字段为空，则默认按照应用名字正向排序。
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// 分页偏移量。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页读取数量。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListApplicationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListApplicationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListApplicationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的应用信息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 返回的应用信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApplicationInfoList []*ApplicationInformation `json:"ApplicationInfoList,omitempty" name:"ApplicationInfoList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListApplicationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListApplicationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// 查询范围是否包括用户关联的用户组、组织机构的应用访问权限。默认为不查询 。传false表示不查询该范围，传true表示应用查询该范围。
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

type ListUserGroupsRequest struct {
	*tchttp.BaseRequest

	// 查询条件，支持多搜索条件组合、多数据范围匹配的搜索。同时支持查询信息内容全匹配、部分匹配、范围匹配等多种查询方式，具体查询方式为：双引号（""）表示全匹配、以星号（* ) 结尾表示字段部分匹配。如果该字段为空，则默认查全量表。
	SearchCondition *UserGroupInfoSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// 排序条件集合。可排序的属性支持：用户组名称（DisplayName）、用户组ID（UserGroupId）、上次更新时间（LastModifiedDate）。如果该字段为空，则默认按照用户组名称正向排序。
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// 分页偏移量。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页读取数量。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	delete(f, "SearchCondition")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的用户组列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserGroupList []*UserGroupInformation `json:"UserGroupList,omitempty" name:"UserGroupList"`

		// 返回的用户组信息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type ListUsersInOrgNodeRequest struct {
	*tchttp.BaseRequest

	// 机构节点ID，是机构节点全局唯一标识，长度限制：64个字符。如果为空默认读取机构根节点下用户信息。
	OrgNodeId *string `json:"OrgNodeId,omitempty" name:"OrgNodeId"`

	// 是否读取其子节点信息。当其为空或false时，默认仅读取当前机构节点信息。当其为true时，读取本机构节点以及其第一层子节点信息。
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

type ListUsersRequest struct {
	*tchttp.BaseRequest

	// 用户属性搜索条件，可查询条件包括：用户名、手机号码，邮箱、用户锁定状态、用户冻结状态、创建时间、上次修改时间，支持多种属性组合作为查询条件。同时支持查询信息内容全匹配、部分匹配、范围匹配等多种查询方式，具体查询方式为：双引号（“”）表示全匹配、以星号（*）结尾表示字段部分匹配、中括号以逗号分隔（[Min，Max]）表示闭区间查询、大括号以逗号分隔（{Min，Max}）表示开区间查询，中括号与大括号可以配合使用（例如：{Min，Max]表示最小值开区间，最大值闭区间查询）。范围匹配支持使用星号（例如{20,*]表示查询范围为大于20的所有数据）。范围查询同时支持时间段查询，支持的属性包括创建时间 （CreationTime）、上次修改时间（LastUpdateTime），查询的时间格式遵循 ISO 8601 标准，例如：2021-01-13T09:44:07.182+0000。
	SearchCondition *UserSearchCriteria `json:"SearchCondition,omitempty" name:"SearchCondition"`

	// 指定期望返回的用户属性，默认返回所有用户内置属性。内置用户属性包括：用户UUID（UserId）、用户昵称（DisplayName）、用户名字（UserName）、手机号（Phone）、邮箱（Email）、用户状态（Status）、用户组（SubjectGroups）机构路径（OrgPath）、备注（Description）、创建时间 （CreationTime）、上次修改时间（LastUpdateTime）、上次登录时间（LastLoginTime）。
	ExpectedFields []*string `json:"ExpectedFields,omitempty" name:"ExpectedFields"`

	// 排序条件集合。可排序的属性支持：用户名字（UserName）、手机号（Phone）、邮箱（Email）、用户状态（Status）、创建时间 （CreationTime）、上次修改时间（LastUpdateTime）、上次登录时间（LastLoginTime）。
	Sort *SortCondition `json:"Sort,omitempty" name:"Sort"`

	// 分页偏移量，默认为0。Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询，即只返回最多1000个用户。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页读取数量，默认为50，最大值为100。 Offset 和 Limit 两个字段需配合使用，即其中一个指定了，另一个必须指定。 如果不指定以上参数，则表示不进行分页查询，即只返回最多1000个用户。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 是否查看搜索结果的总数，默认该选项为false不查看。
	IncludeTotal *bool `json:"IncludeTotal,omitempty" name:"IncludeTotal"`
}

func (r *ListUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchCondition")
	delete(f, "ExpectedFields")
	delete(f, "Sort")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IncludeTotal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询返回的相关用户列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		UserList []*UserInformation `json:"UserList,omitempty" name:"UserList"`

		// 返回查询用户的总数量，仅当入参IncludeTotal等于true时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest

	// 应用ID，是应用的全局唯一标识。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`

	// 安全级别。
	SecureLevel *string `json:"SecureLevel,omitempty" name:"SecureLevel"`

	// 应用展示名称，长度限制：32个字符。 默认与应用名字相同。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 应用状态
	AppStatus *bool `json:"AppStatus,omitempty" name:"AppStatus"`

	// 应用图标图片访问地址。
	IconUrl *string `json:"IconUrl,omitempty" name:"IconUrl"`

	// 描述。长度不超过128。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "SecureLevel")
	delete(f, "DisplayName")
	delete(f, "AppStatus")
	delete(f, "IconUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationResponse) FromJsonString(s string) error {
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

type SortCondition struct {

	// 排序属性。
	SortKey *string `json:"SortKey,omitempty" name:"SortKey"`

	// 排序顺序，ASC为正向排序，DESC为反向排序。
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
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

type UserGroupInfoSearchCriteria struct {

	// 名称匹配搜索，匹配范围包括：用户组名称、用户组ID。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type UserGroupInformation struct {

	// 用户组ID。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 用户组名称。
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`

	// 上次更新时间，符合 ISO8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifiedDate *string `json:"LastModifiedDate,omitempty" name:"LastModifiedDate"`
}

type UserInfo struct {

	// 用户ID，是用户全局唯一标识，长度限制：64个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 昵称，长度限制：64个字符。 默认与用户名相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

type UserInformation struct {

	// 用户名，长度限制：32个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 昵称，长度限制：64个字符。 默认与用户名相同。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 用户备注，长度限制：512个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用户上次更新时间，遵循 ISO 8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 用户创建时间，遵循 ISO 8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 用户所属组织机构路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPath *string `json:"OrgPath,omitempty" name:"OrgPath"`

	// 带国家号的用户手机号，例如+86-00000000000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户所属用户组ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectGroups []*string `json:"SubjectGroups,omitempty" name:"SubjectGroups"`

	// 用户邮箱。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户上次登录时间，遵循 ISO 8601 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLoginTime *string `json:"LastLoginTime,omitempty" name:"LastLoginTime"`

	// 用户ID，是用户全局唯一标识，长度限制：64个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type UserSearchCriteria struct {

	// 用户名，长度限制：64个字符。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 用户手机号。
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 用户邮箱。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 用户状态，取值 NORMAL （正常）、FREEZE （已冻结）、LOCKED （已锁定）或 NOT_ENABLED （未启用）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 用户创建时间，遵循 ISO 8601 标准。
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 用户上次更新时间区间。
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" name:"LastUpdateTime"`

	// 名称匹配搜索，匹配范围包括：用户名称、用户ID。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}
