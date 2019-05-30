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

package v20190116

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddUserRequest struct {
	*tchttp.BaseRequest

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子用户备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 子用户是否可以登录控制台。传0子用户无法登录控制台，传1子用户可以登录控制台。
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 是否生成子用户密钥。传0不生成子用户密钥，传1生成子用户密钥。
	UseApi *uint64 `json:"UseApi,omitempty" name:"UseApi"`

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大写小字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大写小字母、数字和特殊字符。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 子用户是否要在下次登录时重置密码。传0子用户下次登录控制台不需重置密码，传1子用户下次登录控制台需要重置密码。
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *AddUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子用户 UIN
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 子用户用户名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 如果输入参数组合为自动生成随机密码，则返回生成的密码
		Password *string `json:"Password,omitempty" name:"Password"`

		// 子用户密钥 ID
		SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

		// 子用户密钥 Key
		SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

		// 子用户 UID
		Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest

	// 添加的子用户 UID 和用户组 ID 关联关系
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info" list`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserToGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserToGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserToGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachEntityOfPolicy struct {

	// 实体ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 实体名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实体Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 关联类型。1 用户关联 ； 2 用户组关联
	RelatedType *uint64 `json:"RelatedType,omitempty" name:"RelatedType"`
}

type AttachGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户组 id
	AttachGroupId *int64 `json:"AttachGroupId,omitempty" name:"AttachGroupId"`
}

func (r *AttachGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachPolicyInfo struct {

	// 策略id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`

	// 取值为user和QCS
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 子账号 uin
	AttachUin *int64 `json:"AttachUin,omitempty" name:"AttachUin"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 用户组描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组 ID
		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 策略文档
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新增策略id
		PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML身份提供商Base64编码的元数据文档
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *CreateSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SAML身份提供商资源描述符
		ProviderArn *string `json:"ProviderArn,omitempty" name:"ProviderArn"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest

	// 数组，数组成员是策略 id，支持批量删除策略
	PolicyId []*uint64 `json:"PolicyId,omitempty" name:"PolicyId" list`
}

func (r *DeletePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserRequest) FromJsonString(s string) error {
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

func (r *DeleteUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户组 id
	DetachGroupId *int64 `json:"DetachGroupId,omitempty" name:"DetachGroupId"`
}

func (r *DetachGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachGroupPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachGroupPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 子账号 uin
	DetachUin *int64 `json:"DetachUin,omitempty" name:"DetachUin"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *GetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组 ID
		GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

		// 用户组名称
		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

		// 用户组成员数量
		GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`

		// 用户组描述
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 用户组创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 用户组成员信息
		UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略Id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *GetPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

		// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
		Description *string `json:"Description,omitempty" name:"Description"`

		// 1 表示自定义策略，2 表示预设策略
	// 注意：此字段可能返回 null，表示取不到有效值。
		Type *uint64 `json:"Type,omitempty" name:"Type"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

		// 最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

		// 策略文档
	// 注意：此字段可能返回 null，表示取不到有效值。
		PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SAML身份提供商名称
		Name *string `json:"Name,omitempty" name:"Name"`

		// SAML身份提供商描述
		Description *string `json:"Description,omitempty" name:"Description"`

		// SAML身份提供商创建时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// SAML身份提供商上次修改时间
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// SAML身份提供商元数据文档
		SAMLMetadata *string `json:"SAMLMetadata,omitempty" name:"SAMLMetadata"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserRequest struct {
	*tchttp.BaseRequest

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子用户用户 ID
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 子用户用户名
		Name *string `json:"Name,omitempty" name:"Name"`

		// 子用户 UID
		Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

		// 子用户备注
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 子用户能否登录控制台
		ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

		// 手机号
		PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

		// 区号
		CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

		// 邮箱
		Email *string `json:"Email,omitempty" name:"Email"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GroupIdOfUidInfo struct {

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`
}

type GroupInfo struct {

	// 用户组 ID。
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 用户组名称。
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 用户组创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 用户组描述。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type GroupMemberInfo struct {

	// 子用户 Uid。
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 子用户 Uin。
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 子用户名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 手机号。
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 手机区域代码。
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 是否已验证手机。
	PhoneFlag *uint64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`

	// 邮箱地址。
	Email *string `json:"Email,omitempty" name:"Email"`

	// 是否已验证邮箱。
	EmailFlag *uint64 `json:"EmailFlag,omitempty" name:"EmailFlag"`

	// 用户类型。
	UserType *uint64 `json:"UserType,omitempty" name:"UserType"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否为主消息接收人。
	IsReceiverOwner *uint64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
}

type ListAttachedGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 用户组 id
	TargetGroupId *uint64 `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListAttachedGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedGroupPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 策略列表
		List []*AttachPolicyInfo `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedGroupPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserPoliciesRequest struct {
	*tchttp.BaseRequest

	// 子账号 uin
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListAttachedUserPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedUserPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedUserPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 策略列表
		List []*AttachPolicyInfo `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedUserPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedUserPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 可取值 'All'、'User'、'Group' 和 'Role'，'All' 表示获取所有实体类型，'User' 表示只获取子账号，'Group' 表示只获取用户组，'Role' 表示只获取角色，默认取 'All'
	EntityFilter *string `json:"EntityFilter,omitempty" name:"EntityFilter"`
}

func (r *ListEntitiesForPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListEntitiesForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实体总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 实体列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*AttachEntityOfPolicy `json:"List,omitempty" name:"List" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsForUserRequest struct {
	*tchttp.BaseRequest

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitempty" name:"Page"`
}

func (r *ListGroupsForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsForUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsForUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子用户加入的用户组总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 用户组信息
		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupsForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsForUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsRequest struct {
	*tchttp.BaseRequest

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 按用户组名称匹配。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组总数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 用户组数组信息。
		GroupInfo []*GroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest

	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 页码，默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// 按策略名匹配
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 策略数组，数组每个成员包括 policyId、policyName、addTime、type、description、 createMode 字段。其中： 
	// policyId：策略 id 
	// policyName：策略名
	// addTime：策略创建时间
	// type：1 表示自定义策略，2 表示预设策略 
	// description：策略描述 
	// createMode：1 表示按业务权限创建的策略，其他值表示可以查看策略语法和通过策略语法更新策略
		List []*StrategyInfo `json:"List,omitempty" name:"List" list`

		// 保留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServiceTypeList []*string `json:"ServiceTypeList,omitempty" name:"ServiceTypeList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSAMLProvidersRequest struct {
	*tchttp.BaseRequest
}

func (r *ListSAMLProvidersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSAMLProvidersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListSAMLProvidersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SAML身份提供商总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// SAML身份提供商列表
		SAMLProviderSet []*SAMLProviderInfo `json:"SAMLProviderSet,omitempty" name:"SAMLProviderSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListSAMLProvidersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListSAMLProvidersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组 ID。
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *ListUsersForGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersForGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersForGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户组关联的用户总数。
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 子用户信息。
		UserInfo []*GroupMemberInfo `json:"UserInfo,omitempty" name:"UserInfo" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersForGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersForGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersRequest struct {
	*tchttp.BaseRequest
}

func (r *ListUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListUsersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子用户信息
		Data []*SubAccountInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListUsersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest

	// 要删除的用户 UID和用户组 ID对应数组
	Info []*GroupIdOfUidInfo `json:"Info,omitempty" name:"Info" list`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveUserFromGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SAMLProviderInfo struct {

	// SAML身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML身份提供商创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// SAML身份提供商上次修改时间
	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
}

type StrategyInfo struct {

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 策略创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 策略类型。1 表示自定义策略，2 表示预设策略。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 策略描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`

	// 关联的用户数
	Attachments *uint64 `json:"Attachments,omitempty" name:"Attachments"`

	// 策略关联的产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

type SubAccountInfo struct {

	// 子用户用户 ID
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 子用户备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 子用户能否登录控制台
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest

	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitempty" name:"GroupId"`

	// 用户组名
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 用户组描述
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 策略文档
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdatePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSAMLProviderRequest struct {
	*tchttp.BaseRequest

	// SAML身份提供商名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// SAML身份提供商Base64编码的元数据文档
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitempty" name:"SAMLMetadataDocument"`
}

func (r *UpdateSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSAMLProviderRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateSAMLProviderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子用户备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 子用户是否可以登录控制台。传0子用户无法登录控制台，传1子用户可以登录控制台。
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大写小字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大写小字母、数字和特殊字符。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 子用户是否要在下次登录时重置密码。传0子用户下次登录控制台不需重置密码，传1子用户下次登录控制台需要重置密码。
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitempty" name:"NeedResetPassword"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *UpdateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
