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

type AccessKey struct {

	// 访问密钥标识
	AccessKeyId *string `json:"AccessKeyId,omitempty" name:"AccessKeyId"`

	// 密钥状态，激活（Active）或未激活（Inactive）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

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

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大小写字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大小写字母、数字和特殊字符。
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

	// 策略关联时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachmentTime *string `json:"AttachmentTime,omitempty" name:"AttachmentTime"`
}

type AttachGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户组 id
	AttachGroupId *uint64 `json:"AttachGroupId,omitempty" name:"AttachGroupId"`
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

	// 策略备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 策略关联操作者主帐号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOwnerUin *string `json:"OperateOwnerUin,omitempty" name:"OperateOwnerUin"`

	// 策略关联操作者ID，如果UinType为0表示子帐号Uin，如果UinType为1表示角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// UinType为0表示OperateUin字段是子帐号Uin，如果UinType为1表示OperateUin字段是角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUinType *uint64 `json:"OperateUinType,omitempty" name:"OperateUinType"`

	// 是否已下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// 已下线的产品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail" list`
}

type AttachRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleId *string `json:"AttachRoleId,omitempty" name:"AttachRoleId"`

	// 角色名称，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleName *string `json:"AttachRoleName,omitempty" name:"AttachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *AttachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 子账号 uin
	AttachUin *uint64 `json:"AttachUin,omitempty" name:"AttachUin"`
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

type AttachedPolicyOfRole struct {

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 绑定时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 策略类型，User表示自定义策略，QCS表示预设策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 策略创建方式，1表示按产品功能或项目权限创建，其他表示按策略语法创建
	CreateMode *uint64 `json:"CreateMode,omitempty" name:"CreateMode"`

	// 是否已下线(0:否 1:是)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// 已下线的产品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail" list`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type ConsumeCustomMFATokenRequest struct {
	*tchttp.BaseRequest

	// 自定义多因子验证Token
	MFAToken *string `json:"MFAToken,omitempty" name:"MFAToken"`
}

func (r *ConsumeCustomMFATokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConsumeCustomMFATokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConsumeCustomMFATokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConsumeCustomMFATokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConsumeCustomMFATokenResponse) FromJsonString(s string) error {
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

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
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

		// 新增策略ID
		PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

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

type CreatePolicyVersionRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略文本信息
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 是否设置为当前策略的版本
	SetAsDefault *bool `json:"SetAsDefault,omitempty" name:"SetAsDefault"`
}

func (r *CreatePolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
		VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePolicyVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 是否允许登录 1 为允许 0 为不允许
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 申请角色临时密钥的最长有效期限制(范围：0~43200)
	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
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

type CreateServiceLinkedRoleRequest struct {
	*tchttp.BaseRequest

	// 授权服务，附加了此角色的腾讯云服务主体。
	QCSServiceName []*string `json:"QCSServiceName,omitempty" name:"QCSServiceName" list`

	// 自定义后缀，根据您提供的字符串，与服务提供的前缀组合在一起以形成完整的角色名称。
	CustomSuffix *string `json:"CustomSuffix,omitempty" name:"CustomSuffix"`

	// 角色说明。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateServiceLinkedRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceLinkedRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateServiceLinkedRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色ID
		RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceLinkedRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateServiceLinkedRoleResponse) FromJsonString(s string) error {
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

type DeletePolicyVersionRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略版本号
	VersionId []*uint64 `json:"VersionId,omitempty" name:"VersionId" list`
}

func (r *DeletePolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePolicyVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRolePermissionsBoundaryRequest struct {
	*tchttp.BaseRequest

	// 角色ID（与角色名至少填一个）
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名（与角色ID至少填一个）
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRolePermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRolePermissionsBoundaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRolePermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRolePermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRolePermissionsBoundaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRoleResponse) FromJsonString(s string) error {
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

type DeleteServiceLinkedRoleRequest struct {
	*tchttp.BaseRequest

	// 要删除的服务相关角色的名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *DeleteServiceLinkedRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceLinkedRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteServiceLinkedRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除任务ID，可用于检查删除服务相关角色状态。
		DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteServiceLinkedRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteServiceLinkedRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserPermissionsBoundaryRequest struct {
	*tchttp.BaseRequest

	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *DeleteUserPermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserPermissionsBoundaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserPermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserPermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteUserPermissionsBoundaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否强制删除该子用户，默认入参为0。0：若该用户存在未删除API密钥，则不删除用户；1：若该用户存在未删除API密钥，则先删除密钥后删除用户。删除密钥需要您拥有cam:DeleteApiKey权限，您将可以删除该用户下启用或禁用状态的所有密钥，无权限则删除密钥和用户失败
	Force *uint64 `json:"Force,omitempty" name:"Force"`
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

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest

	// 页码，从1开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		List []*RoleInfo `json:"List,omitempty" name:"List" list`

		// 角色总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRoleListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeAuthFlagCollRequest struct {
	*tchttp.BaseRequest

	// 子账号
	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
}

func (r *DescribeSafeAuthFlagCollRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSafeAuthFlagCollRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeAuthFlagCollResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 登录保护设置
		LoginFlag *LoginActionFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

		// 敏感操作保护设置
		ActionFlag *LoginActionFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`

		// 异地登录保护设置
		OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSafeAuthFlagCollResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSafeAuthFlagCollResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeAuthFlagRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSafeAuthFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSafeAuthFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSafeAuthFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 登录保护设置
		LoginFlag *LoginActionFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

		// 敏感操作保护设置
		ActionFlag *LoginActionFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`

		// 异地登录保护设置
		OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitempty" name:"OffsiteFlag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSafeAuthFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSafeAuthFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachGroupPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 用户组 id
	DetachGroupId *uint64 `json:"DetachGroupId,omitempty" name:"DetachGroupId"`
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

type DetachRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	DetachRoleId *string `json:"DetachRoleId,omitempty" name:"DetachRoleId"`

	// 角色名称，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	DetachRoleName *string `json:"DetachRoleName,omitempty" name:"DetachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

func (r *DetachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest

	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 子账号 uin
	DetachUin *uint64 `json:"DetachUin,omitempty" name:"DetachUin"`
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

type GetCustomMFATokenInfoRequest struct {
	*tchttp.BaseRequest

	// 自定义多因子验证Token
	MFAToken *string `json:"MFAToken,omitempty" name:"MFAToken"`
}

func (r *GetCustomMFATokenInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomMFATokenInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetCustomMFATokenInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 自定义多因子验证Token对应的帐号Id
		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetCustomMFATokenInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetCustomMFATokenInfoResponse) FromJsonString(s string) error {
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

		// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
		PresetAlias *string `json:"PresetAlias,omitempty" name:"PresetAlias"`

		// 是否服务相关策略
	// 注意：此字段可能返回 null，表示取不到有效值。
		IsServiceLinkedRolePolicy *uint64 `json:"IsServiceLinkedRolePolicy,omitempty" name:"IsServiceLinkedRolePolicy"`

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

type GetPolicyVersionRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略版本号
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *GetPolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetPolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略版本详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		PolicyVersion *PolicyVersionDetail `json:"PolicyVersion,omitempty" name:"PolicyVersion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetPolicyVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRoleRequest struct {
	*tchttp.BaseRequest

	// 角色 ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *GetRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色详情
		RoleInfo *RoleInfo `json:"RoleInfo,omitempty" name:"RoleInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRoleResponse) FromJsonString(s string) error {
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

type GetServiceLinkedRoleDeletionStatusRequest struct {
	*tchttp.BaseRequest

	// 删除任务ID
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`
}

func (r *GetServiceLinkedRoleDeletionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceLinkedRoleDeletionStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetServiceLinkedRoleDeletionStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态：NOT_STARTED，IN_PROGRESS，SUCCEEDED，FAILED
		Status *string `json:"Status,omitempty" name:"Status"`

		// 失败原因
		Reason *string `json:"Reason,omitempty" name:"Reason"`

		// 服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`

		// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceLinkedRoleDeletionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetServiceLinkedRoleDeletionStatusResponse) FromJsonString(s string) error {
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

		// 子用户用户 UIN
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

type ListAccessKeysRequest struct {
	*tchttp.BaseRequest

	// 指定用户Uin，不填默认列出当前用户访问密钥
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *ListAccessKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAccessKeysRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAccessKeysResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 访问密钥列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		AccessKeys []*AccessKey `json:"AccessKeys,omitempty" name:"AccessKeys" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAccessKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAccessKeysResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedGroupPoliciesRequest struct {
	*tchttp.BaseRequest

	// 用户组ID
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

type ListAttachedRolePoliciesRequest struct {
	*tchttp.BaseRequest

	// 页码，从 1 开始
	Page *uint64 `json:"Page,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitempty" name:"Rp"`

	// 角色 ID。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 按策略类型过滤，User表示仅查询自定义策略，QCS表示仅查询预设策略
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAttachedRolePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色关联的策略列表
		List []*AttachedPolicyOfRole `json:"List,omitempty" name:"List" list`

		// 角色关联的策略总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
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

type ListCollaboratorsRequest struct {
	*tchttp.BaseRequest

	// 分页条数，缺省为20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页起始值，缺省为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListCollaboratorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCollaboratorsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListCollaboratorsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 协作者信息
		Data []*SubAccountInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCollaboratorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListCollaboratorsResponse) FromJsonString(s string) error {
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

	// 子账号UIN
	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
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
	// Attachments: 关联的用户数
	// ServiceType: 策略关联的产品
	// IsAttached: 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略
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

type ListPolicyVersionsRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *ListPolicyVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPolicyVersionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListPolicyVersionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略版本列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Versions []*PolicyVersionItem `json:"Versions,omitempty" name:"Versions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListPolicyVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListPolicyVersionsResponse) FromJsonString(s string) error {
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

type ListWeChatWorkSubAccountsRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *ListWeChatWorkSubAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListWeChatWorkSubAccountsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListWeChatWorkSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业微信子用户列表。
		Data []*WeChatWorkSubAccount `json:"Data,omitempty" name:"Data" list`

		// 总数目。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListWeChatWorkSubAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListWeChatWorkSubAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LoginActionFlag struct {

	// 手机
	Phone *uint64 `json:"Phone,omitempty" name:"Phone"`

	// 硬token
	Token *uint64 `json:"Token,omitempty" name:"Token"`

	// 软token
	Stoken *uint64 `json:"Stoken,omitempty" name:"Stoken"`

	// 微信
	Wechat *uint64 `json:"Wechat,omitempty" name:"Wechat"`

	// 自定义
	Custom *uint64 `json:"Custom,omitempty" name:"Custom"`
}

type LoginActionMfaFlag struct {

	// 手机
	Phone *uint64 `json:"Phone,omitempty" name:"Phone"`

	// 软token
	Stoken *uint64 `json:"Stoken,omitempty" name:"Stoken"`

	// 微信
	Wechat *uint64 `json:"Wechat,omitempty" name:"Wechat"`
}

type OffsiteFlag struct {

	// 验证标识
	VerifyFlag *uint64 `json:"VerifyFlag,omitempty" name:"VerifyFlag"`

	// 手机通知
	NotifyPhone *uint64 `json:"NotifyPhone,omitempty" name:"NotifyPhone"`

	// 邮箱通知
	NotifyEmail *int64 `json:"NotifyEmail,omitempty" name:"NotifyEmail"`

	// 微信通知
	NotifyWechat *uint64 `json:"NotifyWechat,omitempty" name:"NotifyWechat"`

	// 提示
	Tips *uint64 `json:"Tips,omitempty" name:"Tips"`
}

type PolicyVersionDetail struct {

	// 策略版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`

	// 策略版本创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 是否是正在生效的版本。0表示不是，1表示是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultVersion *uint64 `json:"IsDefaultVersion,omitempty" name:"IsDefaultVersion"`

	// 策略语法文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Document *string `json:"Document,omitempty" name:"Document"`
}

type PolicyVersionItem struct {

	// 策略版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`

	// 策略版本创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitempty" name:"CreateDate"`

	// 是否是正在生效的版本。0表示不是，1表示是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultVersion *int64 `json:"IsDefaultVersion,omitempty" name:"IsDefaultVersion"`
}

type PutRolePermissionsBoundaryRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 角色ID（与角色名至少填一个）
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名（与角色ID至少填一个）
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *PutRolePermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutRolePermissionsBoundaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutRolePermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutRolePermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutRolePermissionsBoundaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutUserPermissionsBoundaryRequest struct {
	*tchttp.BaseRequest

	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitempty" name:"TargetUin"`

	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *PutUserPermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutUserPermissionsBoundaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutUserPermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutUserPermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutUserPermissionsBoundaryResponse) FromJsonString(s string) error {
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

type RoleInfo struct {

	// 角色ID
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色的策略文档
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 角色的创建时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 角色的最近一次时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 角色是否允许登录
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 角色类型，取user、system或service_linked
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleType *string `json:"RoleType,omitempty" name:"RoleType"`

	// 有效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionDuration *uint64 `json:"SessionDuration,omitempty" name:"SessionDuration"`

	// 服务相关角色删除TaskId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" name:"DeletionTaskId"`
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

type SetDefaultPolicyVersionRequest struct {
	*tchttp.BaseRequest

	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略版本号
	VersionId *uint64 `json:"VersionId,omitempty" name:"VersionId"`
}

func (r *SetDefaultPolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDefaultPolicyVersionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultPolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetDefaultPolicyVersionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetMfaFlagRequest struct {
	*tchttp.BaseRequest

	// 设置用户的uin
	OpUin *uint64 `json:"OpUin,omitempty" name:"OpUin"`

	// 登录保护设置
	LoginFlag *LoginActionMfaFlag `json:"LoginFlag,omitempty" name:"LoginFlag"`

	// 操作保护设置
	ActionFlag *LoginActionMfaFlag `json:"ActionFlag,omitempty" name:"ActionFlag"`
}

func (r *SetMfaFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetMfaFlagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetMfaFlagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMfaFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetMfaFlagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAttached *uint64 `json:"IsAttached,omitempty" name:"IsAttached"`

	// 是否已下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deactived *uint64 `json:"Deactived,omitempty" name:"Deactived"`

	// 已下线产品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeactivedDetail []*string `json:"DeactivedDetail,omitempty" name:"DeactivedDetail" list`

	// 是否是服务相关角色策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsServiceLinkedPolicy *uint64 `json:"IsServiceLinkedPolicy,omitempty" name:"IsServiceLinkedPolicy"`
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

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type UpdateAssumeRolePolicyRequest struct {
	*tchttp.BaseRequest

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateAssumeRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssumeRolePolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAssumeRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssumeRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssumeRolePolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 策略ID，与PolicyName二选一必填
	PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 策略名，与PolicyId二选一必填
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 策略描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitempty" name:"PolicyDocument"`

	// 预设策略备注
	Alias *string `json:"Alias,omitempty" name:"Alias"`
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

		// 策略id，入参是PolicyName时，才会返回
	// 注意：此字段可能返回 null，表示取不到有效值。
		PolicyId *uint64 `json:"PolicyId,omitempty" name:"PolicyId"`

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

type UpdateRoleConsoleLoginRequest struct {
	*tchttp.BaseRequest

	// 是否可登录，可登录：1，不可登录：0
	ConsoleLogin *int64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 角色ID
	RoleId *int64 `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateRoleConsoleLoginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleConsoleLoginRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleConsoleLoginResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleConsoleLoginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleConsoleLoginResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionRequest struct {
	*tchttp.BaseRequest

	// 角色描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

func (r *UpdateRoleDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleDescriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRoleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRoleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRoleDescriptionResponse) FromJsonString(s string) error {
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

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大小写字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大小写字母、数字和特殊字符。
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

type WeChatWorkSubAccount struct {

	// 子用户用户 ID
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitempty" name:"Uid"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 子用户能否登录控制台
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitempty" name:"ConsoleLogin"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 企业微信UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeChatWorkUserId *string `json:"WeChatWorkUserId,omitempty" name:"WeChatWorkUserId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}
