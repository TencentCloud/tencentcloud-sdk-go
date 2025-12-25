// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccessKey struct {
	// 访问密钥标识
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 密钥状态，激活（Active）或未激活（Inactive）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 密钥描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AccessKeyDetail struct {
	// 访问密钥标识
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 访问密钥（密钥仅创建时可见，请妥善保存）
	SecretAccessKey *string `json:"SecretAccessKey,omitnil,omitempty" name:"SecretAccessKey"`

	// 密钥状态，激活（Active）或未激活（Inactive）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type AddUserRequestParams struct {
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户是否可以登录控制台。传0子用户无法登录控制台，传1子用户可以登录控制台。
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 是否生成子用户密钥。传0不生成子用户密钥，传1生成子用户密钥。
	UseApi *uint64 `json:"UseApi,omitnil,omitempty" name:"UseApi"`

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大小写字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大小写字母、数字和特殊字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 子用户是否要在下次登录时重置密码。传0子用户下次登录控制台不需重置密码，传1子用户下次登录控制台需要重置密码。
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type AddUserRequest struct {
	*tchttp.BaseRequest
	
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户是否可以登录控制台。传0子用户无法登录控制台，传1子用户可以登录控制台。
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 是否生成子用户密钥。传0不生成子用户密钥，传1生成子用户密钥。
	UseApi *uint64 `json:"UseApi,omitnil,omitempty" name:"UseApi"`

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大小写字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大小写字母、数字和特殊字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 子用户是否要在下次登录时重置密码。传0子用户下次登录控制台不需重置密码，传1子用户下次登录控制台需要重置密码。
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

func (r *AddUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "ConsoleLogin")
	delete(f, "UseApi")
	delete(f, "Password")
	delete(f, "NeedResetPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserResponseParams struct {
	// 子用户 UIN
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 如果输入参数组合为自动生成随机密码，则返回生成的密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 子用户密钥 ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 子用户密钥 Key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserResponse struct {
	*tchttp.BaseResponse
	Response *AddUserResponseParams `json:"Response"`
}

func (r *AddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupRequestParams struct {
	// 添加的子用户 UIN/UID 和用户组 ID 关联关系
	Info []*GroupIdOfUidInfo `json:"Info,omitnil,omitempty" name:"Info"`
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest
	
	// 添加的子用户 UIN/UID 和用户组 ID 关联关系
	Info []*GroupIdOfUidInfo `json:"Info,omitnil,omitempty" name:"Info"`
}

func (r *AddUserToGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserToGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserToGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddUserToGroupResponseParams `json:"Response"`
}

func (r *AddUserToGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserToGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachEntityOfPolicy struct {
	// 实体ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 实体名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实体Uin
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 关联类型。1 用户关联 ； 2 用户组关联 3 角色关联
	RelatedType *uint64 `json:"RelatedType,omitnil,omitempty" name:"RelatedType"`

	// 策略关联时间
	AttachmentTime *string `json:"AttachmentTime,omitnil,omitempty" name:"AttachmentTime"`
}

// Predefined struct for user
type AttachGroupPolicyRequestParams struct {
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 用户组 id
	AttachGroupId *uint64 `json:"AttachGroupId,omitnil,omitempty" name:"AttachGroupId"`
}

type AttachGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 用户组 id
	AttachGroupId *uint64 `json:"AttachGroupId,omitnil,omitempty" name:"AttachGroupId"`
}

func (r *AttachGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "AttachGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachGroupPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachGroupPolicyResponseParams `json:"Response"`
}

func (r *AttachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachPolicyInfo struct {
	// 策略id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 创建来源，1 通过控制台创建, 2 通过策略语法创建
	CreateMode *uint64 `json:"CreateMode,omitnil,omitempty" name:"CreateMode"`

	// 取值为User和QCS。User代表自定义策略，QCS代表系统策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 策略关联操作者主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateOwnerUin *string `json:"OperateOwnerUin,omitnil,omitempty" name:"OperateOwnerUin"`

	// 策略关联操作者ID，如果UinType为0表示子账号Uin，如果UinType为1表示角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 取值为0和1。OperateUinType为0表示OperateUin字段是子账号Uin。如果OperateUinType为1表示OperateUin字段是角色ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUinType *uint64 `json:"OperateUinType,omitnil,omitempty" name:"OperateUinType"`

	// 是否已下线，1代表已下线，0代表未下线
	Deactived *uint64 `json:"Deactived,omitnil,omitempty" name:"Deactived"`

	// 已下线的产品列表
	DeactivedDetail []*string `json:"DeactivedDetail,omitnil,omitempty" name:"DeactivedDetail"`
}

// Predefined struct for user
type AttachRolePolicyRequestParams struct {
	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleId *string `json:"AttachRoleId,omitnil,omitempty" name:"AttachRoleId"`

	// 角色名称，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleName *string `json:"AttachRoleName,omitnil,omitempty" name:"AttachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

type AttachRolePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleId *string `json:"AttachRoleId,omitnil,omitempty" name:"AttachRoleId"`

	// 角色名称，用于指定角色，入参 AttachRoleId 与 AttachRoleName 二选一
	AttachRoleName *string `json:"AttachRoleName,omitnil,omitempty" name:"AttachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

func (r *AttachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "AttachRoleId")
	delete(f, "AttachRoleName")
	delete(f, "PolicyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachRolePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachRolePolicyResponseParams `json:"Response"`
}

func (r *AttachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyRequestParams struct {
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 子账号 uin
	AttachUin *uint64 `json:"AttachUin,omitnil,omitempty" name:"AttachUin"`
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 子账号 uin
	AttachUin *uint64 `json:"AttachUin,omitnil,omitempty" name:"AttachUin"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "AttachUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachUserPolicyResponseParams `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttachedPolicyOfRole struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 绑定时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略类型，User表示自定义策略，QCS表示预设策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略创建方式，1表示按产品功能或项目权限创建，其他表示按策略语法创建
	CreateMode *uint64 `json:"CreateMode,omitnil,omitempty" name:"CreateMode"`

	// 是否已下线(0:否 1:是)
	Deactived *uint64 `json:"Deactived,omitnil,omitempty" name:"Deactived"`

	// 已下线的产品列表
	DeactivedDetail []*string `json:"DeactivedDetail,omitnil,omitempty" name:"DeactivedDetail"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AttachedUserPolicy struct {
	// 策略ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略类型(1表示自定义策略，2表示预设策略)
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 创建模式(1表示按产品或项目权限创建的策略，其他表示策略语法创建的策略)
	CreateMode *string `json:"CreateMode,omitnil,omitempty" name:"CreateMode"`

	// 随组关联信息
	Groups []*AttachedUserPolicyGroupInfo `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 是否已下线(0:否 1:是)
	Deactived *uint64 `json:"Deactived,omitnil,omitempty" name:"Deactived"`

	// 已下线的产品列表
	DeactivedDetail []*string `json:"DeactivedDetail,omitnil,omitempty" name:"DeactivedDetail"`
}

type AttachedUserPolicyGroupInfo struct {
	// 分组ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type AuthToken struct {
	// 认证Token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 服务器时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentTime *int64 `json:"CurrentTime,omitnil,omitempty" name:"CurrentTime"`

	// 毫秒时间戳，根据轮转周期准确计算得到
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextRotationTime *int64 `json:"NextRotationTime,omitnil,omitempty" name:"NextRotationTime"`

	// 毫秒，如果轮转失败则为 -1
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRotationTimeCost *int64 `json:"LastRotationTimeCost,omitnil,omitempty" name:"LastRotationTimeCost"`

	// 成功：success
	// 失败：failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	RotationStatus *string `json:"RotationStatus,omitnil,omitempty" name:"RotationStatus"`

	// 成功：success
	// 失败：失败信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RotationMessage *string `json:"RotationMessage,omitnil,omitempty" name:"RotationMessage"`
}

// Predefined struct for user
type BuildDataFlowAuthTokenRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源用户名
	ResourceAccount *string `json:"ResourceAccount,omitnil,omitempty" name:"ResourceAccount"`
}

type BuildDataFlowAuthTokenRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源用户名
	ResourceAccount *string `json:"ResourceAccount,omitnil,omitempty" name:"ResourceAccount"`
}

func (r *BuildDataFlowAuthTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildDataFlowAuthTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceRegion")
	delete(f, "ResourceAccount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BuildDataFlowAuthTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BuildDataFlowAuthTokenResponseParams struct {
	// 认证凭据AuthToken信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Credentials *AuthToken `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BuildDataFlowAuthTokenResponse struct {
	*tchttp.BaseResponse
	Response *BuildDataFlowAuthTokenResponseParams `json:"Response"`
}

func (r *BuildDataFlowAuthTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BuildDataFlowAuthTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConsumeCustomMFATokenRequestParams struct {
	// 自定义多因子验证Token
	MFAToken *string `json:"MFAToken,omitnil,omitempty" name:"MFAToken"`
}

type ConsumeCustomMFATokenRequest struct {
	*tchttp.BaseRequest
	
	// 自定义多因子验证Token
	MFAToken *string `json:"MFAToken,omitnil,omitempty" name:"MFAToken"`
}

func (r *ConsumeCustomMFATokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConsumeCustomMFATokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MFAToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConsumeCustomMFATokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConsumeCustomMFATokenResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ConsumeCustomMFATokenResponse struct {
	*tchttp.BaseResponse
	Response *ConsumeCustomMFATokenResponseParams `json:"Response"`
}

func (r *ConsumeCustomMFATokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConsumeCustomMFATokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeyRequestParams struct {
	// 指定用户Uin，不填默认为当前用户创建访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 密钥描述，长度在1到1024之间，可包含大小写字符，数字以及特殊字符：=,.@:/-。 正则为：[\w+=,.@:/-]*。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateAccessKeyRequest struct {
	*tchttp.BaseRequest
	
	// 指定用户Uin，不填默认为当前用户创建访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 密钥描述，长度在1到1024之间，可包含大小写字符，数字以及特殊字符：=,.@:/-。 正则为：[\w+=,.@:/-]*。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateAccessKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessKeyResponseParams struct {
	// 访问密钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *AccessKeyDetail `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAccessKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessKeyResponseParams `json:"Response"`
}

func (r *CreateAccessKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// 用户组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMessageReceiverRequestParams struct {
	// 消息接收人的用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号国际区号，国内为86
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码, 例如：132****2492
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，例如：57*****@qq.com
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 消息接收人的备注，选填
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateMessageReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 消息接收人的用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号国际区号，国内为86
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号码, 例如：132****2492
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 邮箱，例如：57*****@qq.com
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 消息接收人的备注，选填
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateMessageReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMessageReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "CountryCode")
	delete(f, "PhoneNumber")
	delete(f, "Email")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMessageReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMessageReceiverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMessageReceiverResponse struct {
	*tchttp.BaseResponse
	Response *CreateMessageReceiverResponseParams `json:"Response"`
}

func (r *CreateMessageReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMessageReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOIDCConfigRequestParams struct {
	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID
	ClientId []*string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签名公钥，需要base64
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID
	ClientId []*string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签名公钥，需要base64
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "Name")
	delete(f, "IdentityKey")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateOIDCConfigResponseParams `json:"Response"`
}

func (r *CreateOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyRequestParams struct {
	// 策略名称。长度为1~128个字符，可包含英文字母、数字和+=,.@-_。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略文档
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略关联的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略名称。长度为1~128个字符，可包含英文字母、数字和+=,.@-_。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略文档
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略关联的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyName")
	delete(f, "PolicyDocument")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyResponseParams struct {
	// 新增策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyResponseParams `json:"Response"`
}

func (r *CreatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyVersionRequestParams struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略文本信息
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 是否设置为当前策略的版本
	SetAsDefault *bool `json:"SetAsDefault,omitnil,omitempty" name:"SetAsDefault"`
}

type CreatePolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略文本信息
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 是否设置为当前策略的版本
	SetAsDefault *bool `json:"SetAsDefault,omitnil,omitempty" name:"SetAsDefault"`
}

func (r *CreatePolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "PolicyDocument")
	delete(f, "SetAsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyVersionResponseParams struct {
	// 策略版本号
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *CreatePolicyVersionResponseParams `json:"Response"`
}

func (r *CreatePolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleRequestParams struct {
	// 角色名称。长度为1~128个字符，可包含英文字母、数字和+=,.@-_。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否允许登录 1 为允许 0 为不允许
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 申请角色临时密钥的最长有效期限制(范围：0~43200)
	SessionDuration *uint64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 角色绑定标签
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称。长度为1~128个字符，可包含英文字母、数字和+=,.@-_。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否允许登录 1 为允许 0 为不允许
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 申请角色临时密钥的最长有效期限制(范围：0~43200)
	SessionDuration *uint64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 角色绑定标签
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "PolicyDocument")
	delete(f, "Description")
	delete(f, "ConsoleLogin")
	delete(f, "SessionDuration")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleResponseParams struct {
	// 角色ID
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleResponseParams `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSAMLProviderRequestParams struct {
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SAML身份提供商Base64编码的元数据文档
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`
}

type CreateSAMLProviderRequest struct {
	*tchttp.BaseRequest
	
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SAML身份提供商Base64编码的元数据文档
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`
}

func (r *CreateSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "SAMLMetadataDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSAMLProviderResponseParams struct {
	// SAML身份提供商资源描述符
	ProviderArn *string `json:"ProviderArn,omitnil,omitempty" name:"ProviderArn"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *CreateSAMLProviderResponseParams `json:"Response"`
}

func (r *CreateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceLinkedRoleRequestParams struct {
	// 填写此角色的腾讯云服务载体，具体可查询文档（角色载体）字段
	// https://cloud.tencent.com/document/product/598/85165
	QCSServiceName []*string `json:"QCSServiceName,omitnil,omitempty" name:"QCSServiceName"`

	// 自定义后缀，根据您提供的字符串，与服务提供的前缀组合在一起以形成完整的角色名称。
	CustomSuffix *string `json:"CustomSuffix,omitnil,omitempty" name:"CustomSuffix"`

	// 角色说明。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色绑定标签。
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateServiceLinkedRoleRequest struct {
	*tchttp.BaseRequest
	
	// 填写此角色的腾讯云服务载体，具体可查询文档（角色载体）字段
	// https://cloud.tencent.com/document/product/598/85165
	QCSServiceName []*string `json:"QCSServiceName,omitnil,omitempty" name:"QCSServiceName"`

	// 自定义后缀，根据您提供的字符串，与服务提供的前缀组合在一起以形成完整的角色名称。
	CustomSuffix *string `json:"CustomSuffix,omitnil,omitempty" name:"CustomSuffix"`

	// 角色说明。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色绑定标签。
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateServiceLinkedRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceLinkedRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QCSServiceName")
	delete(f, "CustomSuffix")
	delete(f, "Description")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceLinkedRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceLinkedRoleResponseParams struct {
	// 角色ID
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateServiceLinkedRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceLinkedRoleResponseParams `json:"Response"`
}

func (r *CreateServiceLinkedRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceLinkedRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubAccountLoginIpPolicyRequestParams struct {
	// IP策略集合
	IpPolicies []*IpPolicy `json:"IpPolicies,omitnil,omitempty" name:"IpPolicies"`

	// 审批人类型，目前支持的类型有：SubAccountLoginLimitApproval（子账号登录限制审批）
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 被添加为协助审批人的账号ID数组
	ApproverUin []*uint64 `json:"ApproverUin,omitnil,omitempty" name:"ApproverUin"`

	// 是否禁用策略：0：不禁用，1：禁用
	DisablePolicy *uint64 `json:"DisablePolicy,omitnil,omitempty" name:"DisablePolicy"`

	// 策略类型：1：更新IP策略，2：设置异地登录校验校验规则
	PolicyType *uint64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type CreateSubAccountLoginIpPolicyRequest struct {
	*tchttp.BaseRequest
	
	// IP策略集合
	IpPolicies []*IpPolicy `json:"IpPolicies,omitnil,omitempty" name:"IpPolicies"`

	// 审批人类型，目前支持的类型有：SubAccountLoginLimitApproval（子账号登录限制审批）
	ApproverType *string `json:"ApproverType,omitnil,omitempty" name:"ApproverType"`

	// 被添加为协助审批人的账号ID数组
	ApproverUin []*uint64 `json:"ApproverUin,omitnil,omitempty" name:"ApproverUin"`

	// 是否禁用策略：0：不禁用，1：禁用
	DisablePolicy *uint64 `json:"DisablePolicy,omitnil,omitempty" name:"DisablePolicy"`

	// 策略类型：1：更新IP策略，2：设置异地登录校验校验规则
	PolicyType *uint64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *CreateSubAccountLoginIpPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubAccountLoginIpPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IpPolicies")
	delete(f, "ApproverType")
	delete(f, "ApproverUin")
	delete(f, "DisablePolicy")
	delete(f, "PolicyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubAccountLoginIpPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubAccountLoginIpPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSubAccountLoginIpPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubAccountLoginIpPolicyResponseParams `json:"Response"`
}

func (r *CreateSubAccountLoginIpPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubAccountLoginIpPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserOIDCConfigRequestParams struct {
	// 身份提供商URL。OpenID Connect身份提供商标识。
	// 对应企业IdP提供的Openid-configuration中"issuer"字段的值。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，固定值id_token
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// 签名公钥，需要base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。openid; email;profile。授权请求信息范围。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述信息。由用户自行定义。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商URL。OpenID Connect身份提供商标识。
	// 对应企业IdP提供的Openid-configuration中"issuer"字段的值。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，固定值id_token
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// 签名公钥，需要base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。openid; email;profile。授权请求信息范围。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述信息。由用户自行定义。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "IdentityKey")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserOIDCConfigResponseParams `json:"Response"`
}

func (r *CreateUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSAMLConfigRequestParams struct {
	// SAML元数据文档，需要base64 encode
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`

	// 辅助域名
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitnil,omitempty" name:"AuxiliaryDomain"`
}

type CreateUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
	
	// SAML元数据文档，需要base64 encode
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`

	// 辅助域名
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitnil,omitempty" name:"AuxiliaryDomain"`
}

func (r *CreateUserSAMLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSAMLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SAMLMetadataDocument")
	delete(f, "AuxiliaryDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserSAMLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSAMLConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserSAMLConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserSAMLConfigResponseParams `json:"Response"`
}

func (r *CreateUserSAMLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessKeyRequestParams struct {
	// 指定需要删除的AccessKeyId
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 指定用户Uin，不填默认为当前用户删除访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type DeleteAccessKeyRequest struct {
	*tchttp.BaseRequest
	
	// 指定需要删除的AccessKeyId
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 指定用户Uin，不填默认为当前用户删除访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *DeleteAccessKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccessKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessKeyResponseParams `json:"Response"`
}

func (r *DeleteAccessKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMessageReceiverRequestParams struct {
	// 消息接受者的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteMessageReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 消息接受者的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteMessageReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMessageReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMessageReceiverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMessageReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMessageReceiverResponseParams `json:"Response"`
}

func (r *DeleteMessageReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMessageReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOIDCConfigRequestParams struct {
	// OIDC身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// OIDC身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOIDCConfigResponseParams `json:"Response"`
}

func (r *DeleteOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyRequestParams struct {
	// 数组，数组成员是策略 id，支持批量删除策略
	PolicyId []*uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 数组，数组成员是策略 id，支持批量删除策略
	PolicyId []*uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DeletePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyResponseParams `json:"Response"`
}

func (r *DeletePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyVersionRequestParams struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略版本号
	VersionId []*uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DeletePolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略版本号
	VersionId []*uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DeletePolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *DeletePolicyVersionResponseParams `json:"Response"`
}

func (r *DeletePolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRolePermissionsBoundaryRequestParams struct {
	// 角色ID（与角色名至少填一个）
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名（与角色ID至少填一个）
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type DeleteRolePermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// 角色ID（与角色名至少填一个）
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名（与角色ID至少填一个）
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *DeleteRolePermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolePermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRolePermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRolePermissionsBoundaryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRolePermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRolePermissionsBoundaryResponseParams `json:"Response"`
}

func (r *DeleteRolePermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolePermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleRequestParams struct {
	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleResponseParams `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSAMLProviderRequestParams struct {
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteSAMLProviderRequest struct {
	*tchttp.BaseRequest
	
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DeleteSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSAMLProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSAMLProviderResponseParams `json:"Response"`
}

func (r *DeleteSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceLinkedRoleRequestParams struct {
	// 要删除的服务相关角色的名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type DeleteServiceLinkedRoleRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的服务相关角色的名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *DeleteServiceLinkedRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceLinkedRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceLinkedRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceLinkedRoleResponseParams struct {
	// 删除任务ID，可用于检查删除服务相关角色状态。
	DeletionTaskId *string `json:"DeletionTaskId,omitnil,omitempty" name:"DeletionTaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteServiceLinkedRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceLinkedRoleResponseParams `json:"Response"`
}

func (r *DeleteServiceLinkedRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceLinkedRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserPermissionsBoundaryRequestParams struct {
	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type DeleteUserPermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *DeleteUserPermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserPermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserPermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserPermissionsBoundaryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserPermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserPermissionsBoundaryResponseParams `json:"Response"`
}

func (r *DeleteUserPermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserPermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否强制删除该子用户，默认入参为0。0：若该用户存在未删除API密钥，则不删除用户；1：若该用户存在未删除API密钥，则先删除密钥后删除用户。删除密钥需要您拥有cam:DeleteApiKey权限，您将可以删除该用户下启用或禁用状态的所有密钥，无权限则删除密钥和用户失败
	Force *uint64 `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否强制删除该子用户，默认入参为0。0：若该用户存在未删除API密钥，则不删除用户；1：若该用户存在未删除API密钥，则先删除密钥后删除用户。删除密钥需要您拥有cam:DeleteApiKey权限，您将可以删除该用户下启用或禁用状态的所有密钥，无权限则删除密钥和用户失败
	Force *uint64 `json:"Force,omitnil,omitempty" name:"Force"`
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
	delete(f, "Name")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeOIDCConfigRequestParams struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOIDCConfigResponseParams struct {
	// 身份提供商类型 11角色身份提供商
	ProviderType *uint64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 签名公钥
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 客户端id
	ClientId []*string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 状态：0:未设置，11:已开启，2:已禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOIDCConfigResponseParams `json:"Response"`
}

func (r *DescribeOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListRequestParams struct {
	// 页码，从1开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 标签筛选
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从1开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 标签筛选
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListResponseParams struct {
	// 角色详情列表。
	List []*RoleInfo `json:"List,omitnil,omitempty" name:"List"`

	// 角色总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoleListResponseParams `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagCollRequestParams struct {
	// 子账号
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type DescribeSafeAuthFlagCollRequest struct {
	*tchttp.BaseRequest
	
	// 子账号
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

func (r *DescribeSafeAuthFlagCollRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagCollRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeAuthFlagCollRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagCollResponseParams struct {
	// 登录保护设置
	LoginFlag *LoginActionFlag `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 敏感操作保护设置
	ActionFlag *LoginActionFlag `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// 异地登录保护设置
	OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitnil,omitempty" name:"OffsiteFlag"`

	// 是否提示信任设备1 ：提示 0: 不提示
	PromptTrust *int64 `json:"PromptTrust,omitnil,omitempty" name:"PromptTrust"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSafeAuthFlagCollResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeAuthFlagCollResponseParams `json:"Response"`
}

func (r *DescribeSafeAuthFlagCollResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagCollResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagIntlRequestParams struct {

}

type DescribeSafeAuthFlagIntlRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSafeAuthFlagIntlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagIntlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeAuthFlagIntlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagIntlResponseParams struct {
	// 登录保护设置
	LoginFlag *LoginActionFlagIntl `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 敏感操作保护设置
	ActionFlag *LoginActionFlagIntl `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// 异地登录保护设置
	OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitnil,omitempty" name:"OffsiteFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSafeAuthFlagIntlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeAuthFlagIntlResponseParams `json:"Response"`
}

func (r *DescribeSafeAuthFlagIntlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagIntlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagRequestParams struct {

}

type DescribeSafeAuthFlagRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSafeAuthFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeAuthFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeAuthFlagResponseParams struct {
	// 登录保护设置
	LoginFlag *LoginActionFlag `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 敏感操作保护设置
	ActionFlag *LoginActionFlag `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`

	// 异地登录保护设置
	OffsiteFlag *OffsiteFlag `json:"OffsiteFlag,omitnil,omitempty" name:"OffsiteFlag"`

	// 是否提示信任设备：1: 提示  0: 不提示
	PromptTrust *uint64 `json:"PromptTrust,omitnil,omitempty" name:"PromptTrust"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSafeAuthFlagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeAuthFlagResponseParams `json:"Response"`
}

func (r *DescribeSafeAuthFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeAuthFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAccountsRequestParams struct {
	// 子用户UIN列表，最多支持50个UIN
	FilterSubAccountUin []*uint64 `json:"FilterSubAccountUin,omitnil,omitempty" name:"FilterSubAccountUin"`
}

type DescribeSubAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 子用户UIN列表，最多支持50个UIN
	FilterSubAccountUin []*uint64 `json:"FilterSubAccountUin,omitnil,omitempty" name:"FilterSubAccountUin"`
}

func (r *DescribeSubAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FilterSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubAccountsResponseParams struct {
	// 子用户列表
	SubAccounts []*SubAccountUser `json:"SubAccounts,omitnil,omitempty" name:"SubAccounts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubAccountsResponseParams `json:"Response"`
}

func (r *DescribeSubAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserOIDCConfigRequestParams struct {

}

type DescribeUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserOIDCConfigResponseParams struct {
	// 身份提供商类型。 12：用户OIDC身份提供商
	ProviderType *uint64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 签名公钥
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 客户端id
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 状态：0:未设置，11:已开启，2:已禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 授权请求Endpoint
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Scope
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 授权请求Response type
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserOIDCConfigResponseParams `json:"Response"`
}

func (r *DescribeUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSAMLConfigRequestParams struct {

}

type DescribeUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserSAMLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSAMLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSAMLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSAMLConfigResponseParams struct {
	// SAML元数据文档
	SAMLMetadata *string `json:"SAMLMetadata,omitnil,omitempty" name:"SAMLMetadata"`

	// 状态：0:未设置，11:已开启，2:已禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 辅助域名
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitnil,omitempty" name:"AuxiliaryDomain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserSAMLConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSAMLConfigResponseParams `json:"Response"`
}

func (r *DescribeUserSAMLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachGroupPolicyRequestParams struct {
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 用户组 id
	DetachGroupId *uint64 `json:"DetachGroupId,omitnil,omitempty" name:"DetachGroupId"`
}

type DetachGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 用户组 id
	DetachGroupId *uint64 `json:"DetachGroupId,omitnil,omitempty" name:"DetachGroupId"`
}

func (r *DetachGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "DetachGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachGroupPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachGroupPolicyResponseParams `json:"Response"`
}

func (r *DetachGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachRolePolicyRequestParams struct {
	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 DetachRoleId 与 DetachRoleName 二选一
	DetachRoleId *string `json:"DetachRoleId,omitnil,omitempty" name:"DetachRoleId"`

	// 角色名称，用于指定角色，入参 DetachRoleId 与 DetachRoleName 二选一
	DetachRoleName *string `json:"DetachRoleName,omitnil,omitempty" name:"DetachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

type DetachRolePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID，入参PolicyId与PolicyName二选一
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 角色ID，用于指定角色，入参 DetachRoleId 与 DetachRoleName 二选一
	DetachRoleId *string `json:"DetachRoleId,omitnil,omitempty" name:"DetachRoleId"`

	// 角色名称，用于指定角色，入参 DetachRoleId 与 DetachRoleName 二选一
	DetachRoleName *string `json:"DetachRoleName,omitnil,omitempty" name:"DetachRoleName"`

	// 策略名，入参PolicyId与PolicyName二选一
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

func (r *DetachRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "DetachRoleId")
	delete(f, "DetachRoleName")
	delete(f, "PolicyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachRolePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachRolePolicyResponseParams `json:"Response"`
}

func (r *DetachRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyRequestParams struct {
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 子账号 uin
	DetachUin *uint64 `json:"DetachUin,omitnil,omitempty" name:"DetachUin"`
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 子账号 uin
	DetachUin *uint64 `json:"DetachUin,omitnil,omitempty" name:"DetachUin"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "DetachUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachUserPolicyResponseParams `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserSSORequestParams struct {

}

type DisableUserSSORequest struct {
	*tchttp.BaseRequest
	
}

func (r *DisableUserSSORequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserSSORequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableUserSSORequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableUserSSOResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableUserSSOResponse struct {
	*tchttp.BaseResponse
	Response *DisableUserSSOResponseParams `json:"Response"`
}

func (r *DisableUserSSOResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableUserSSOResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountSummaryRequestParams struct {

}

type GetAccountSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAccountSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAccountSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAccountSummaryResponseParams struct {
	// 策略数
	Policies *uint64 `json:"Policies,omitnil,omitempty" name:"Policies"`

	// 角色数
	Roles *uint64 `json:"Roles,omitnil,omitempty" name:"Roles"`

	// 身份提供商数
	//
	// Deprecated: Idps is deprecated.
	Idps *uint64 `json:"Idps,omitnil,omitempty" name:"Idps"`

	// 子账户数
	User *uint64 `json:"User,omitnil,omitempty" name:"User"`

	// 分组数
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`

	// 分组用户总数
	Member *uint64 `json:"Member,omitnil,omitempty" name:"Member"`

	// 身份提供商数。
	IdentityProviders *uint64 `json:"IdentityProviders,omitnil,omitempty" name:"IdentityProviders"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAccountSummaryResponse struct {
	*tchttp.BaseResponse
	Response *GetAccountSummaryResponseParams `json:"Response"`
}

func (r *GetAccountSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAccountSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomMFATokenInfoRequestParams struct {
	// 自定义多因子验证Token，针对用户自定义的安全校验方式而生成的，以供查询用户安全校验时使用。
	MFAToken *string `json:"MFAToken,omitnil,omitempty" name:"MFAToken"`
}

type GetCustomMFATokenInfoRequest struct {
	*tchttp.BaseRequest
	
	// 自定义多因子验证Token，针对用户自定义的安全校验方式而生成的，以供查询用户安全校验时使用。
	MFAToken *string `json:"MFAToken,omitnil,omitempty" name:"MFAToken"`
}

func (r *GetCustomMFATokenInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomMFATokenInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MFAToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCustomMFATokenInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomMFATokenInfoResponseParams struct {
	// 自定义多因子验证Token对应的账号Id
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCustomMFATokenInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetCustomMFATokenInfoResponseParams `json:"Response"`
}

func (r *GetCustomMFATokenInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomMFATokenInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupRequestParams struct {
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type GetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *GetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupResponseParams struct {
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组成员数量
	GroupNum *uint64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// 用户组描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 用户组创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户组成员信息
	UserInfo []*GroupMemberInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetGroupResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupResponseParams `json:"Response"`
}

func (r *GetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyRequestParams struct {
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type GetPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *GetPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyResponseParams struct {
	// 策略名。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 1 表示自定义策略，2 表示预设策略。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略创建时间。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略最近更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 策略文档。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresetAlias *string `json:"PresetAlias,omitnil,omitempty" name:"PresetAlias"`

	// 是否是服务相关策略，0代表不是服务相关策略，1代表是服务相关策略。
	IsServiceLinkedRolePolicy *uint64 `json:"IsServiceLinkedRolePolicy,omitnil,omitempty" name:"IsServiceLinkedRolePolicy"`

	// 策略关联的标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPolicyResponse struct {
	*tchttp.BaseResponse
	Response *GetPolicyResponseParams `json:"Response"`
}

func (r *GetPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyVersionRequestParams struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略版本号，可由ListPolicyVersions获取
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type GetPolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略版本号，可由ListPolicyVersions获取
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *GetPolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPolicyVersionResponseParams struct {
	// 策略版本详情
	PolicyVersion *PolicyVersionDetail `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetPolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *GetPolicyVersionResponseParams `json:"Response"`
}

func (r *GetPolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRolePermissionBoundaryRequestParams struct {
	// 角色ID
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

type GetRolePermissionBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// 角色ID
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

func (r *GetRolePermissionBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRolePermissionBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRolePermissionBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRolePermissionBoundaryResponseParams struct {
	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略语法
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 策略类型：1.自定义策略，2.预设策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 创建方式：1.按产品功能或项目权限创建，2.按策略语法创建，3.按策略生成器创建，4.按标签授权创建，5.按权限边界规则创建
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateMode *int64 `json:"CreateMode,omitnil,omitempty" name:"CreateMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRolePermissionBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *GetRolePermissionBoundaryResponseParams `json:"Response"`
}

func (r *GetRolePermissionBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRolePermissionBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleRequestParams struct {
	// 角色 ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type GetRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色 ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *GetRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleResponseParams struct {
	// 角色详情
	RoleInfo *RoleInfo `json:"RoleInfo,omitnil,omitempty" name:"RoleInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRoleResponse struct {
	*tchttp.BaseResponse
	Response *GetRoleResponseParams `json:"Response"`
}

func (r *GetRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSAMLProviderRequestParams struct {
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type GetSAMLProviderRequest struct {
	*tchttp.BaseRequest
	
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *GetSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSAMLProviderResponseParams struct {
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SAML身份提供商创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// SAML身份提供商上次修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// SAML身份提供商元数据文档
	SAMLMetadata *string `json:"SAMLMetadata,omitnil,omitempty" name:"SAMLMetadata"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *GetSAMLProviderResponseParams `json:"Response"`
}

func (r *GetSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSecurityLastUsedRequestParams struct {
	// 查询密钥ID列表。最多支持10个。
	SecretIdList []*string `json:"SecretIdList,omitnil,omitempty" name:"SecretIdList"`
}

type GetSecurityLastUsedRequest struct {
	*tchttp.BaseRequest
	
	// 查询密钥ID列表。最多支持10个。
	SecretIdList []*string `json:"SecretIdList,omitnil,omitempty" name:"SecretIdList"`
}

func (r *GetSecurityLastUsedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecurityLastUsedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecretIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSecurityLastUsedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSecurityLastUsedResponseParams struct {
	// 密钥ID最近访问列表
	SecretIdLastUsedRows []*SecretIdLastUsed `json:"SecretIdLastUsedRows,omitnil,omitempty" name:"SecretIdLastUsedRows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSecurityLastUsedResponse struct {
	*tchttp.BaseResponse
	Response *GetSecurityLastUsedResponseParams `json:"Response"`
}

func (r *GetSecurityLastUsedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSecurityLastUsedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceLinkedRoleDeletionStatusRequestParams struct {
	// 删除任务ID
	DeletionTaskId *string `json:"DeletionTaskId,omitnil,omitempty" name:"DeletionTaskId"`
}

type GetServiceLinkedRoleDeletionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 删除任务ID
	DeletionTaskId *string `json:"DeletionTaskId,omitnil,omitempty" name:"DeletionTaskId"`
}

func (r *GetServiceLinkedRoleDeletionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceLinkedRoleDeletionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeletionTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetServiceLinkedRoleDeletionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetServiceLinkedRoleDeletionStatusResponseParams struct {
	// 状态：NOT_STARTED，IN_PROGRESS，SUCCEEDED，FAILED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 服务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetServiceLinkedRoleDeletionStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetServiceLinkedRoleDeletionStatusResponseParams `json:"Response"`
}

func (r *GetServiceLinkedRoleDeletionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetServiceLinkedRoleDeletionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserAppIdRequestParams struct {

}

type GetUserAppIdRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetUserAppIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserAppIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserAppIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserAppIdResponseParams struct {
	// 当前账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 当前账号OwnerUin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 当前账号AppId
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUserAppIdResponse struct {
	*tchttp.BaseResponse
	Response *GetUserAppIdResponseParams `json:"Response"`
}

func (r *GetUserAppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserPermissionBoundaryRequestParams struct {
	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type GetUserPermissionBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *GetUserPermissionBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserPermissionBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserPermissionBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserPermissionBoundaryResponseParams struct {
	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略语法
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 策略类型：1.自定义策略，2.预设策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *int64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 创建方式：1.按产品功能或项目权限创建，2.按策略语法创建，3.按策略生成器创建，4.按标签授权创建，5.按权限边界规则创建
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateMode *int64 `json:"CreateMode,omitnil,omitempty" name:"CreateMode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUserPermissionBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *GetUserPermissionBoundaryResponseParams `json:"Response"`
}

func (r *GetUserPermissionBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserPermissionBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserRequestParams struct {
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type GetUserRequest struct {
	*tchttp.BaseRequest
	
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *GetUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserResponseParams struct {
	// 子用户用户 UIN
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 子用户备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户能否登录控制台 0-无法登录控制台，1-可以登录控制台
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 最近一次登录ip
	RecentlyLoginIP *string `json:"RecentlyLoginIP,omitnil,omitempty" name:"RecentlyLoginIP"`

	// 最近一次登录时间
	RecentlyLoginTime *string `json:"RecentlyLoginTime,omitnil,omitempty" name:"RecentlyLoginTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUserResponse struct {
	*tchttp.BaseResponse
	Response *GetUserResponseParams `json:"Response"`
}

func (r *GetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupIdOfUidInfo struct {
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 子用户 Uin，Uid和Uin至少有一个必填
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type GroupInfo struct {
	// 用户组 ID。
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户组描述。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type GroupMemberInfo struct {
	// 子用户 Uid。
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 子用户 Uin。
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子用户名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 手机号。
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 手机区域代码。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 是否已验证手机。0-未验证  1-验证
	PhoneFlag *uint64 `json:"PhoneFlag,omitnil,omitempty" name:"PhoneFlag"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 是否已验证邮箱。0-未验证  1-验证
	EmailFlag *uint64 `json:"EmailFlag,omitnil,omitempty" name:"EmailFlag"`

	// 用户类型。1-全局协作者 2-项目协作者 3-消息接收者
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否为主消息接收人。0-否 1-是
	IsReceiverOwner *uint64 `json:"IsReceiverOwner,omitnil,omitempty" name:"IsReceiverOwner"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type IpPolicy struct {
	// IP段
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 策略效力，Allow或Deny
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`
}

// Predefined struct for user
type ListAccessKeysRequestParams struct {
	// 指定用户Uin，不填默认列出当前用户访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type ListAccessKeysRequest struct {
	*tchttp.BaseRequest
	
	// 指定用户Uin，不填默认列出当前用户访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *ListAccessKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccessKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAccessKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAccessKeysResponseParams struct {
	// 访问密钥列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKeys []*AccessKey `json:"AccessKeys,omitnil,omitempty" name:"AccessKeys"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAccessKeysResponse struct {
	*tchttp.BaseResponse
	Response *ListAccessKeysResponseParams `json:"Response"`
}

func (r *ListAccessKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAccessKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedGroupPoliciesRequestParams struct {
	// 用户组ID
	TargetGroupId *uint64 `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListAttachedGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 用户组ID
	TargetGroupId *uint64 `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListAttachedGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedGroupPoliciesResponseParams struct {
	// 策略总数。取值范围大于等于0。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 策略列表
	List []*AttachPolicyInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttachedGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedGroupPoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedRolePoliciesRequestParams struct {
	// 页码，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 角色 ID。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 按策略类型过滤，User表示仅查询自定义策略，QCS表示仅查询预设策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListAttachedRolePoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页行数，不能大于200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 角色 ID。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名。用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 按策略类型过滤，User表示仅查询自定义策略，QCS表示仅查询预设策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListAttachedRolePoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedRolePoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "RoleId")
	delete(f, "RoleName")
	delete(f, "PolicyType")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedRolePoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedRolePoliciesResponseParams struct {
	// 角色关联的策略列表
	List []*AttachedPolicyOfRole `json:"List,omitnil,omitempty" name:"List"`

	// 角色关联的策略总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttachedRolePoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedRolePoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedRolePoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedRolePoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserAllPoliciesRequestParams struct {
	// 目标用户Uin
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 每页数量，必须大于 0 且小于等于 200。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码，从 1开始，不能大于 200。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 关联类型。0:返回直接关联和随组关联策略，1:只返回直接关联策略，2:只返回随组关联策略。
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`

	// 策略类型。1表示自定义策略，2表示预设策略。
	StrategyType *uint64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListAttachedUserAllPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 目标用户Uin
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 每页数量，必须大于 0 且小于等于 200。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码，从 1开始，不能大于 200。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 关联类型。0:返回直接关联和随组关联策略，1:只返回直接关联策略，2:只返回随组关联策略。
	AttachType *uint64 `json:"AttachType,omitnil,omitempty" name:"AttachType"`

	// 策略类型。1表示自定义策略，2表示预设策略。
	StrategyType *uint64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListAttachedUserAllPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserAllPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "AttachType")
	delete(f, "StrategyType")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedUserAllPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserAllPoliciesResponseParams struct {
	// 策略列表数据。
	PolicyList []*AttachedUserPolicy `json:"PolicyList,omitnil,omitempty" name:"PolicyList"`

	// 策略总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttachedUserAllPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedUserAllPoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedUserAllPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserAllPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserPoliciesRequestParams struct {
	// 子账号 uin
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`
}

type ListAttachedUserPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 子账号 uin
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`
}

func (r *ListAttachedUserPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "Page")
	delete(f, "Rp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttachedUserPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttachedUserPoliciesResponseParams struct {
	// 策略总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 策略列表
	List []*AttachPolicyInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttachedUserPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListAttachedUserPoliciesResponseParams `json:"Response"`
}

func (r *ListAttachedUserPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttachedUserPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCollaboratorsRequestParams struct {
	// 分页的条数，默认是20条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的起始值，默认从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListCollaboratorsRequest struct {
	*tchttp.BaseRequest
	
	// 分页的条数，默认是20条。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页的起始值，默认从0开始。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListCollaboratorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCollaboratorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCollaboratorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCollaboratorsResponseParams struct {
	// 总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 协作者信息
	Data []*SubAccountInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCollaboratorsResponse struct {
	*tchttp.BaseResponse
	Response *ListCollaboratorsResponseParams `json:"Response"`
}

func (r *ListCollaboratorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCollaboratorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEntitiesForPolicyRequestParams struct {
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 可取值 'All'、'User'、'Group' 和 'Role'，'All' 表示获取所有实体类型，'User' 表示只获取子账号，'Group' 表示只获取用户组，'Role' 表示只获取角色，默认取 'All'
	EntityFilter *string `json:"EntityFilter,omitnil,omitempty" name:"EntityFilter"`
}

type ListEntitiesForPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略 id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 页码，默认值是 1，从 1 开始
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小，默认值是 20
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 可取值 'All'、'User'、'Group' 和 'Role'，'All' 表示获取所有实体类型，'User' 表示只获取子账号，'Group' 表示只获取用户组，'Role' 表示只获取角色，默认取 'All'
	EntityFilter *string `json:"EntityFilter,omitnil,omitempty" name:"EntityFilter"`
}

func (r *ListEntitiesForPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEntitiesForPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "EntityFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListEntitiesForPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListEntitiesForPolicyResponseParams struct {
	// 实体总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 实体列表
	List []*AttachEntityOfPolicy `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListEntitiesForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListEntitiesForPolicyResponseParams `json:"Response"`
}

func (r *ListEntitiesForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListEntitiesForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListGrantServiceAccessActionNode struct {
	// 接口名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 接口描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ListGrantServiceAccessNode struct {
	// 服务
	Service *ListGrantServiceAccessService `json:"Service,omitnil,omitempty" name:"Service"`

	// 接口信息
	Action []*ListGrantServiceAccessActionNode `json:"Action,omitnil,omitempty" name:"Action"`

	// 授权的策略
	Policy []*ListGrantServiceAccessPolicy `json:"Policy,omitnil,omitempty" name:"Policy"`
}

type ListGrantServiceAccessPolicy struct {
	// 策略ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型: Custom自定义策略，Presetting预设策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略描述
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`
}

type ListGrantServiceAccessService struct {
	// 服务
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

// Predefined struct for user
type ListGroupsForUserRequestParams struct {
	// 子用户 UID，入参Uid和SubUin二选一
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 子账号UIN，入参Uid和SubUin二选一
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

type ListGroupsForUserRequest struct {
	*tchttp.BaseRequest
	
	// 子用户 UID，入参Uid和SubUin二选一
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 子账号UIN，入参Uid和SubUin二选一
	SubUin *uint64 `json:"SubUin,omitnil,omitempty" name:"SubUin"`
}

func (r *ListGroupsForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uid")
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "SubUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGroupsForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsForUserResponseParams struct {
	// 子用户加入的用户组总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 用户组信息
	GroupInfo []*GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListGroupsForUserResponse struct {
	*tchttp.BaseResponse
	Response *ListGroupsForUserResponseParams `json:"Response"`
}

func (r *ListGroupsForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsRequestParams struct {
	// 页码。默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 按用户组名称匹配。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 页码。默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 按用户组名称匹配。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "Rp")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsResponseParams struct {
	// 用户组总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 用户组数组信息。
	GroupInfo []*GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ListGroupsResponseParams `json:"Response"`
}

func (r *ListGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesGrantingServiceAccessRequestParams struct {
	// 子账号uin，与RoleId、GroupId三选一必传
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 角色ID，与TargetUin、GroupId三选一必传
	RoleId *uint64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 用户组ID，与TargetUin、RoleId三选一必传
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 服务名，查看服务授权接口详情时需传该字段
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

type ListPoliciesGrantingServiceAccessRequest struct {
	*tchttp.BaseRequest
	
	// 子账号uin，与RoleId、GroupId三选一必传
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 角色ID，与TargetUin、GroupId三选一必传
	RoleId *uint64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 用户组ID，与TargetUin、RoleId三选一必传
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 服务名，查看服务授权接口详情时需传该字段
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`
}

func (r *ListPoliciesGrantingServiceAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesGrantingServiceAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "RoleId")
	delete(f, "GroupId")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPoliciesGrantingServiceAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesGrantingServiceAccessResponseParams struct {
	// 列表
	List []*ListGrantServiceAccessNode `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPoliciesGrantingServiceAccessResponse struct {
	*tchttp.BaseResponse
	Response *ListPoliciesGrantingServiceAccessResponseParams `json:"Response"`
}

func (r *ListPoliciesGrantingServiceAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesGrantingServiceAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesRequestParams struct {
	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码，默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 按策略名匹配
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量，默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码，默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 可取值 'All'、'QCS' 和 'Local'，'All' 获取所有策略，'QCS' 只获取预设策略，'Local' 只获取自定义策略，默认取 'All'
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 按策略名匹配
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "Scope")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesResponseParams struct {
	// 策略总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

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
	List []*StrategyInfo `json:"List,omitnil,omitempty" name:"List"`

	// 保留字段
	ServiceTypeList []*string `json:"ServiceTypeList,omitnil,omitempty" name:"ServiceTypeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ListPoliciesResponseParams `json:"Response"`
}

func (r *ListPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPolicyVersionsRequestParams struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type ListPolicyVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *ListPolicyVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPolicyVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPolicyVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPolicyVersionsResponseParams struct {
	// 策略版本列表
	Versions []*PolicyVersionItem `json:"Versions,omitnil,omitempty" name:"Versions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPolicyVersionsResponse struct {
	*tchttp.BaseResponse
	Response *ListPolicyVersionsResponseParams `json:"Response"`
}

func (r *ListPolicyVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPolicyVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiverRequestParams struct {
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type ListReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *ListReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReceiverResponseParams struct {
	// 总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 联系人列表
	Receivers []*Receiver `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReceiverResponse struct {
	*tchttp.BaseResponse
	Response *ListReceiverResponseParams `json:"Response"`
}

func (r *ListReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSAMLProvidersRequestParams struct {

}

type ListSAMLProvidersRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListSAMLProvidersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSAMLProvidersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSAMLProvidersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSAMLProvidersResponseParams struct {
	// SAML身份提供商总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// SAML身份提供商列表
	SAMLProviderSet []*SAMLProviderInfo `json:"SAMLProviderSet,omitnil,omitempty" name:"SAMLProviderSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSAMLProvidersResponse struct {
	*tchttp.BaseResponse
	Response *ListSAMLProvidersResponseParams `json:"Response"`
}

func (r *ListSAMLProvidersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSAMLProvidersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersForGroupRequestParams struct {
	// 用户组 ID。
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`
}

type ListUsersForGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组 ID。
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 页码。默认为1。
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数量。默认为20。
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`
}

func (r *ListUsersForGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersForGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Page")
	delete(f, "Rp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersForGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersForGroupResponseParams struct {
	// 用户组关联的用户总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 子用户信息。
	UserInfo []*GroupMemberInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUsersForGroupResponse struct {
	*tchttp.BaseResponse
	Response *ListUsersForGroupResponseParams `json:"Response"`
}

func (r *ListUsersForGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsersForGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersRequestParams struct {

}

type ListUsersRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersResponseParams struct {
	// 子用户信息
	Data []*SubAccountInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUsersResponse struct {
	*tchttp.BaseResponse
	Response *ListUsersResponseParams `json:"Response"`
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

// Predefined struct for user
type ListWeChatWorkSubAccountsRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type ListWeChatWorkSubAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *ListWeChatWorkSubAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWeChatWorkSubAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWeChatWorkSubAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWeChatWorkSubAccountsResponseParams struct {
	// 企业微信子用户列表。
	Data []*WeChatWorkSubAccount `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数目。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWeChatWorkSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *ListWeChatWorkSubAccountsResponseParams `json:"Response"`
}

func (r *ListWeChatWorkSubAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWeChatWorkSubAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LoginActionFlag struct {
	// 0: 非安全手机校验 1: 安全手机校验。
	Phone *uint64 `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 0: 非硬token校验 1: 硬token校验。
	Token *uint64 `json:"Token,omitnil,omitempty" name:"Token"`

	// 0: 非软token校验 1: 软token校验
	Stoken *uint64 `json:"Stoken,omitnil,omitempty" name:"Stoken"`

	// 0: 非微信校验 1: 微信校验
	Wechat *uint64 `json:"Wechat,omitnil,omitempty" name:"Wechat"`

	// 0: 非自定义校验 1: 自定义校验
	Custom *uint64 `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 0: 非邮箱校验 1: 邮箱校验
	Mail *uint64 `json:"Mail,omitnil,omitempty" name:"Mail"`

	// 0: 非u2f硬件token 1: u2f硬件token
	U2FToken *uint64 `json:"U2FToken,omitnil,omitempty" name:"U2FToken"`
}

type LoginActionFlagIntl struct {
	// 手机
	Phone *uint64 `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 硬token
	Token *uint64 `json:"Token,omitnil,omitempty" name:"Token"`

	// 软token
	Stoken *uint64 `json:"Stoken,omitnil,omitempty" name:"Stoken"`

	// 微信
	Wechat *uint64 `json:"Wechat,omitnil,omitempty" name:"Wechat"`

	// 自定义
	Custom *uint64 `json:"Custom,omitnil,omitempty" name:"Custom"`

	// 邮件
	Mail *uint64 `json:"Mail,omitnil,omitempty" name:"Mail"`

	// u2f硬件token
	U2FToken *uint64 `json:"U2FToken,omitnil,omitempty" name:"U2FToken"`
}

type LoginActionMfaFlag struct {
	// 是否设置手机号为登录和敏感操作安全校验方式， 1: 设置，0: 不设置
	Phone *uint64 `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 是否设置软token为登录和敏感操作安全校验方式， 1: 设置，0: 不设置
	Stoken *uint64 `json:"Stoken,omitnil,omitempty" name:"Stoken"`

	// 是否设置微信为登录和敏感操作安全校验方式， 1: 设置，0: 不设置
	Wechat *uint64 `json:"Wechat,omitnil,omitempty" name:"Wechat"`
}

type OffsiteFlag struct {
	// 验证标识
	VerifyFlag *uint64 `json:"VerifyFlag,omitnil,omitempty" name:"VerifyFlag"`

	// 手机通知
	NotifyPhone *uint64 `json:"NotifyPhone,omitnil,omitempty" name:"NotifyPhone"`

	// 邮箱通知
	NotifyEmail *int64 `json:"NotifyEmail,omitnil,omitempty" name:"NotifyEmail"`

	// 微信通知
	NotifyWechat *uint64 `json:"NotifyWechat,omitnil,omitempty" name:"NotifyWechat"`

	// 提示
	Tips *uint64 `json:"Tips,omitnil,omitempty" name:"Tips"`
}

type PolicyVersionDetail struct {
	// 策略版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 策略版本创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateDate *string `json:"CreateDate,omitnil,omitempty" name:"CreateDate"`

	// 是否是正在生效的版本。0表示不是，1表示是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefaultVersion *uint64 `json:"IsDefaultVersion,omitnil,omitempty" name:"IsDefaultVersion"`

	// 策略语法文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Document *string `json:"Document,omitnil,omitempty" name:"Document"`
}

type PolicyVersionItem struct {
	// 策略版本号
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 策略版本创建时间
	CreateDate *string `json:"CreateDate,omitnil,omitempty" name:"CreateDate"`

	// 是否是正在生效的版本。0表示不是，1表示是
	IsDefaultVersion *int64 `json:"IsDefaultVersion,omitnil,omitempty" name:"IsDefaultVersion"`
}

// Predefined struct for user
type PutRolePermissionsBoundaryRequestParams struct {
	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 角色ID（与角色名至少填一个）
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名（与角色ID至少填一个）
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type PutRolePermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 角色ID（与角色名至少填一个）
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名（与角色ID至少填一个）
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *PutRolePermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutRolePermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutRolePermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutRolePermissionsBoundaryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutRolePermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *PutRolePermissionsBoundaryResponseParams `json:"Response"`
}

func (r *PutRolePermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutRolePermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutUserPermissionsBoundaryRequestParams struct {
	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type PutUserPermissionsBoundaryRequest struct {
	*tchttp.BaseRequest
	
	// 子账号Uin
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 策略ID
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *PutUserPermissionsBoundaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutUserPermissionsBoundaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutUserPermissionsBoundaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutUserPermissionsBoundaryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutUserPermissionsBoundaryResponse struct {
	*tchttp.BaseResponse
	Response *PutUserPermissionsBoundaryResponseParams `json:"Response"`
}

func (r *PutUserPermissionsBoundaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutUserPermissionsBoundaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Receiver struct {
	// id
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 手机号码
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 手机号码是否验证
	PhoneFlag *int64 `json:"PhoneFlag,omitnil,omitempty" name:"PhoneFlag"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 邮箱是否验证
	EmailFlag *int64 `json:"EmailFlag,omitnil,omitempty" name:"EmailFlag"`

	// 是否主联系人
	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitnil,omitempty" name:"IsReceiverOwner"`

	// 是否允许微信接收通知
	WechatFlag *int64 `json:"WechatFlag,omitnil,omitempty" name:"WechatFlag"`

	// 账号uin
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 国家代码
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`
}

// Predefined struct for user
type RemoveUserFromGroupRequestParams struct {
	// 要删除的用户 UIN/UID和用户组 ID对应数组
	Info []*GroupIdOfUidInfo `json:"Info,omitnil,omitempty" name:"Info"`
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的用户 UIN/UID和用户组 ID对应数组
	Info []*GroupIdOfUidInfo `json:"Info,omitnil,omitempty" name:"Info"`
}

func (r *RemoveUserFromGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveUserFromGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveUserFromGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveUserFromGroupResponseParams `json:"Response"`
}

func (r *RemoveUserFromGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveUserFromGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleInfo struct {
	// 角色ID
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色的策略文档
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 角色描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色的创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 角色的最近一次时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 角色是否允许登录
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 角色类型，取user、system或service_linked
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// 有效时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionDuration *uint64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 服务相关角色删除TaskId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletionTaskId *string `json:"DeletionTaskId,omitnil,omitempty" name:"DeletionTaskId"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 角色RoleArn信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleArn *string `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`
}

type RoleTags struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SAMLProviderInfo struct {
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SAML身份提供商创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// SAML身份提供商上次修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type SecretIdLastUsed struct {
	// 密钥ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 最后访问日期(有1天延迟)
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUsedDate *string `json:"LastUsedDate,omitnil,omitempty" name:"LastUsedDate"`

	// 最后密钥访问日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSecretUsedDate *uint64 `json:"LastSecretUsedDate,omitnil,omitempty" name:"LastSecretUsedDate"`
}

// Predefined struct for user
type SetDefaultPolicyVersionRequestParams struct {
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略版本号，可由ListPolicyVersions获取
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type SetDefaultPolicyVersionRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略版本号，可由ListPolicyVersions获取
	VersionId *uint64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *SetDefaultPolicyVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultPolicyVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetDefaultPolicyVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetDefaultPolicyVersionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetDefaultPolicyVersionResponse struct {
	*tchttp.BaseResponse
	Response *SetDefaultPolicyVersionResponseParams `json:"Response"`
}

func (r *SetDefaultPolicyVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetDefaultPolicyVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMfaFlagRequestParams struct {
	// 设置用户的uin
	OpUin *uint64 `json:"OpUin,omitnil,omitempty" name:"OpUin"`

	// 登录保护设置
	LoginFlag *LoginActionMfaFlag `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 操作保护设置
	ActionFlag *LoginActionMfaFlag `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`
}

type SetMfaFlagRequest struct {
	*tchttp.BaseRequest
	
	// 设置用户的uin
	OpUin *uint64 `json:"OpUin,omitnil,omitempty" name:"OpUin"`

	// 登录保护设置
	LoginFlag *LoginActionMfaFlag `json:"LoginFlag,omitnil,omitempty" name:"LoginFlag"`

	// 操作保护设置
	ActionFlag *LoginActionMfaFlag `json:"ActionFlag,omitnil,omitempty" name:"ActionFlag"`
}

func (r *SetMfaFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMfaFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpUin")
	delete(f, "LoginFlag")
	delete(f, "ActionFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetMfaFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetMfaFlagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetMfaFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetMfaFlagResponseParams `json:"Response"`
}

func (r *SetMfaFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetMfaFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyInfo struct {
	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略创建时间。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略类型。1 表示自定义策略，2 表示预设策略。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建来源，1 通过控制台创建, 2 通过策略语法创建。
	CreateMode *uint64 `json:"CreateMode,omitnil,omitempty" name:"CreateMode"`

	// 关联的用户数
	Attachments *uint64 `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// 策略关联的产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// 当需要查询标记实体是否已经关联策略时不为null。0表示未关联策略，1表示已关联策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAttached *uint64 `json:"IsAttached,omitnil,omitempty" name:"IsAttached"`

	// 是否已下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deactived *uint64 `json:"Deactived,omitnil,omitempty" name:"Deactived"`

	// 已下线产品列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeactivedDetail []*string `json:"DeactivedDetail,omitnil,omitempty" name:"DeactivedDetail"`

	// 是否是服务相关角色策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsServiceLinkedPolicy *uint64 `json:"IsServiceLinkedPolicy,omitnil,omitempty" name:"IsServiceLinkedPolicy"`

	// 关联策略实体数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachEntityCount *int64 `json:"AttachEntityCount,omitnil,omitempty" name:"AttachEntityCount"`

	// 关联权限边界实体数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachEntityBoundaryCount *int64 `json:"AttachEntityBoundaryCount,omitnil,omitempty" name:"AttachEntityBoundaryCount"`

	// 最后编辑时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SubAccountInfo struct {
	// 子用户用户 ID
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 子用户备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户能否登录控制台
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`
}

type SubAccountUser struct {
	// 子用户用户 ID
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户 UID，UID是用户作为消息接收人时的唯一标识，和 UIN 一样可以唯一标识一个用户，可通过接口https://cloud.tencent.com/document/api/598/53486 获取
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 子用户备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户类型(2:子用户;3:企业微信子用户;4:协作者;5:消息接收人)
	UserType *uint64 `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 最近登录IP
	LastLoginIp *string `json:"LastLoginIp,omitnil,omitempty" name:"LastLoginIp"`

	// 最近登录时间，回参为空，即为未登录过控制台
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLoginTime *string `json:"LastLoginTime,omitnil,omitempty" name:"LastLoginTime"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type TagRoleRequestParams struct {
	// 标签
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 角色名，与角色ID至少输入一个
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID，与角色名至少输入一个
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

type TagRoleRequest struct {
	*tchttp.BaseRequest
	
	// 标签
	Tags []*RoleTags `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 角色名，与角色ID至少输入一个
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID，与角色名至少输入一个
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

func (r *TagRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	delete(f, "RoleName")
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TagRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TagRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TagRoleResponse struct {
	*tchttp.BaseResponse
	Response *TagRoleResponseParams `json:"Response"`
}

func (r *TagRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TagRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntagRoleRequestParams struct {
	// 标签键
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 角色名，与角色ID至少输入一个
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID，与角色名至少输入一个
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

type UntagRoleRequest struct {
	*tchttp.BaseRequest
	
	// 标签键
	TagKeys []*string `json:"TagKeys,omitnil,omitempty" name:"TagKeys"`

	// 角色名，与角色ID至少输入一个
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID，与角色名至少输入一个
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

func (r *UntagRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntagRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKeys")
	delete(f, "RoleName")
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UntagRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UntagRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UntagRoleResponse struct {
	*tchttp.BaseResponse
	Response *UntagRoleResponseParams `json:"Response"`
}

func (r *UntagRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UntagRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyRequestParams struct {
	// 指定需要更新的AccessKeyId
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 密钥状态，激活（Active）或未激活（Inactive）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 指定用户Uin，不填默认为当前用户更新访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type UpdateAccessKeyRequest struct {
	*tchttp.BaseRequest
	
	// 指定需要更新的AccessKeyId
	AccessKeyId *string `json:"AccessKeyId,omitnil,omitempty" name:"AccessKeyId"`

	// 密钥状态，激活（Active）或未激活（Inactive）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 指定用户Uin，不填默认为当前用户更新访问密钥
	TargetUin *uint64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *UpdateAccessKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessKeyId")
	delete(f, "Status")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAccessKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAccessKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAccessKeyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAccessKeyResponseParams `json:"Response"`
}

func (r *UpdateAccessKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAccessKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssumeRolePolicyRequestParams struct {
	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type UpdateAssumeRolePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略文档，示例：{"version":"2.0","statement":[{"action":"name/sts:AssumeRole","effect":"allow","principal":{"service":["cloudaudit.cloud.tencent.com","cls.cloud.tencent.com"]}}]}，principal用于指定角色的授权对象。获取该参数可参阅 获取角色详情（https://cloud.tencent.com/document/product/598/36221） 输出参数RoleInfo
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *UpdateAssumeRolePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssumeRolePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyDocument")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAssumeRolePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAssumeRolePolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAssumeRolePolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAssumeRolePolicyResponseParams `json:"Response"`
}

func (r *UpdateAssumeRolePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAssumeRolePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupRequestParams struct {
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 用户组 ID
	GroupId *uint64 `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *UpdateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateGroupResponse struct {
	*tchttp.BaseResponse
	Response *UpdateGroupResponseParams `json:"Response"`
}

func (r *UpdateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOIDCConfigRequestParams struct {
	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID
	ClientId []*string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签名公钥，需要base64
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID
	ClientId []*string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 签名公钥，需要base64
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "Name")
	delete(f, "IdentityKey")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOIDCConfigResponseParams `json:"Response"`
}

func (r *UpdateOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyRequestParams struct {
	// 策略ID，与PolicyName二选一必填
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名，与PolicyId二选一必填
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略文档
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 预设策略备注
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID，与PolicyName二选一必填
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名，与PolicyId二选一必填
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略文档
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 预设策略备注
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

func (r *UpdatePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "PolicyName")
	delete(f, "Description")
	delete(f, "PolicyDocument")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyResponseParams struct {
	// 策略id，入参是PolicyName时，才会返回
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdatePolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdatePolicyResponseParams `json:"Response"`
}

func (r *UpdatePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdatePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConsoleLoginRequestParams struct {
	// 是否可登录，可登录：1，不可登录：0
	ConsoleLogin *int64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 角色ID，入参 RoleId 与 RoleName 二选一
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type UpdateRoleConsoleLoginRequest struct {
	*tchttp.BaseRequest
	
	// 是否可登录，可登录：1，不可登录：0
	ConsoleLogin *int64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 角色ID，入参 RoleId 与 RoleName 二选一
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *UpdateRoleConsoleLoginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConsoleLoginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConsoleLogin")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRoleConsoleLoginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConsoleLoginResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRoleConsoleLoginResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleConsoleLoginResponseParams `json:"Response"`
}

func (r *UpdateRoleConsoleLoginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConsoleLoginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleDescriptionRequestParams struct {
	// 角色描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

type UpdateRoleDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 角色描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色ID，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称，用于指定角色，入参 RoleId 与 RoleName 二选一
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`
}

func (r *UpdateRoleDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Description")
	delete(f, "RoleId")
	delete(f, "RoleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRoleDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleDescriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRoleDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleDescriptionResponseParams `json:"Response"`
}

func (r *UpdateRoleDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleSessionDurationRequestParams struct {
	// 时长(秒)
	SessionDuration *uint64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 角色名(与角色ID必填一个)
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID(与角色名必填一个)
	RoleId *uint64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

type UpdateRoleSessionDurationRequest struct {
	*tchttp.BaseRequest
	
	// 时长(秒)
	SessionDuration *uint64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 角色名(与角色ID必填一个)
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色ID(与角色名必填一个)
	RoleId *uint64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`
}

func (r *UpdateRoleSessionDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleSessionDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionDuration")
	delete(f, "RoleName")
	delete(f, "RoleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRoleSessionDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleSessionDurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRoleSessionDurationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleSessionDurationResponseParams `json:"Response"`
}

func (r *UpdateRoleSessionDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleSessionDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSAMLProviderRequestParams struct {
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SAML身份提供商Base64编码的元数据文档
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`
}

type UpdateSAMLProviderRequest struct {
	*tchttp.BaseRequest
	
	// SAML身份提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// SAML身份提供商描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// SAML身份提供商Base64编码的元数据文档
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`
}

func (r *UpdateSAMLProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSAMLProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "SAMLMetadataDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSAMLProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSAMLProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSAMLProviderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSAMLProviderResponseParams `json:"Response"`
}

func (r *UpdateSAMLProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSAMLProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserOIDCConfigRequestParams struct {
	// 身份提供商URL。OpenID Connect身份提供商标识。
	// 对应企业IdP提供的Openid-configuration中"issuer"字段的值，该URL必须以https开头，符合标准URL格式，不允许带有query参数（以?标识）、fragment片段（以#标识）和登录信息（以@标识）。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID，允许英文字母、数字、特殊字符.-_:/，不能以特殊字符.-_:/开头，单个客户端ID最大64个字符。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值，该URL必须以https开头，符合标准URL格式，不允许带有query参数（以?标识）、fragment片段（以#标识）和登录信息（以@标识）。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，有code，id_token，固定值id_token。
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，有form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段,仅支持英文字母、数字、汉字、符号@、＆_[]-的组合，1-255个中文或英文字符
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// RSA签名公钥，JWKS格式，需要进行base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。有openid; email;profile三种。代表授权请求信息范围openid表示请求访问用户的身份信息，email表示请求访问用户的电子邮件地址，profile表示请求访问用户的基本信息。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述，长度为1~255个英文或中文字符，默认值为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商URL。OpenID Connect身份提供商标识。
	// 对应企业IdP提供的Openid-configuration中"issuer"字段的值，该URL必须以https开头，符合标准URL格式，不允许带有query参数（以?标识）、fragment片段（以#标识）和登录信息（以@标识）。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID，允许英文字母、数字、特殊字符.-_:/，不能以特殊字符.-_:/开头，单个客户端ID最大64个字符。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值，该URL必须以https开头，符合标准URL格式，不允许带有query参数（以?标识）、fragment片段（以#标识）和登录信息（以@标识）。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，有code，id_token，固定值id_token。
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，有form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段,仅支持英文字母、数字、汉字、符号@、＆_[]-的组合，1-255个中文或英文字符
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// RSA签名公钥，JWKS格式，需要进行base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。有openid; email;profile三种。代表授权请求信息范围openid表示请求访问用户的身份信息，email表示请求访问用户的电子邮件地址，profile表示请求访问用户的基本信息。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述，长度为1~255个英文或中文字符，默认值为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "IdentityKey")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserOIDCConfigResponseParams `json:"Response"`
}

func (r *UpdateUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserRequestParams struct {
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户是否可以登录控制台。传0子用户无法登录控制台，传1子用户可以登录控制台。
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大小写字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大小写字母、数字和特殊字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 子用户是否要在下次登录时重置密码。传0子用户下次登录控制台不需重置密码，传1子用户下次登录控制台需要重置密码。
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest
	
	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户是否可以登录控制台。传0子用户无法登录控制台，传1子用户可以登录控制台。
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 子用户控制台登录密码，若未进行密码规则设置则默认密码规则为8位以上同时包含大小写字母、数字和特殊字符。只有可以登录控制台时才有效，如果传空并且上面指定允许登录控制台，则自动生成随机密码，随机密码规则为32位包含大小写字母、数字和特殊字符。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 子用户是否要在下次登录时重置密码。传0子用户下次登录控制台不需重置密码，传1子用户下次登录控制台需要重置密码。
	NeedResetPassword *uint64 `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
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
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "ConsoleLogin")
	delete(f, "Password")
	delete(f, "NeedResetPassword")
	delete(f, "PhoneNum")
	delete(f, "CountryCode")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserResponseParams struct {
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
type UpdateUserSAMLConfigRequestParams struct {
	// 修改的操作类型:enable:启用,disable:禁用,updateSAML:修改元数据文档
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 元数据文档，需要base64 encode，仅当Operate为updateSAML时需要此参数
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`

	// 辅助域名
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitnil,omitempty" name:"AuxiliaryDomain"`
}

type UpdateUserSAMLConfigRequest struct {
	*tchttp.BaseRequest
	
	// 修改的操作类型:enable:启用,disable:禁用,updateSAML:修改元数据文档
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 元数据文档，需要base64 encode，仅当Operate为updateSAML时需要此参数
	SAMLMetadataDocument *string `json:"SAMLMetadataDocument,omitnil,omitempty" name:"SAMLMetadataDocument"`

	// 辅助域名
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitnil,omitempty" name:"AuxiliaryDomain"`
}

func (r *UpdateUserSAMLConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSAMLConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Operate")
	delete(f, "SAMLMetadataDocument")
	delete(f, "AuxiliaryDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserSAMLConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserSAMLConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserSAMLConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserSAMLConfigResponseParams `json:"Response"`
}

func (r *UpdateUserSAMLConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSAMLConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WeChatWorkSubAccount struct {
	// 子用户用户 ID
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子用户用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 子用户 UID
	Uid *uint64 `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 子用户能否登录控制台
	ConsoleLogin *uint64 `json:"ConsoleLogin,omitnil,omitempty" name:"ConsoleLogin"`

	// 手机号
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 区号
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 企业微信UserId
	WeChatWorkUserId *string `json:"WeChatWorkUserId,omitnil,omitempty" name:"WeChatWorkUserId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}