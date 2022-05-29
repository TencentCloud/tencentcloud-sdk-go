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

package v20210331

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

	// 组织子账号Uin。
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// 策略名。
	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`

	// 身份ID。
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// 描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateOrganizationMemberPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyName")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 策略ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PolicyId *int64 `json:"PolicyId,omitempty" name:"PolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationMemberPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 关系策略  取值：Financial
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 关系权限 取值：1-查看账单、2-查看余额、3-资金划拨、4-合并出账、5-开票 ，1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// 账号名
	AccountName *string `json:"AccountName,omitempty" name:"AccountName"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 重试创建传记录ID
	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`

	// 代付者Uin
	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

	// 管理身份
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitempty" name:"IdentityRoleID"`
}

func (r *CreateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "NodeId")
	delete(f, "AccountName")
	delete(f, "Remark")
	delete(f, "RecordId")
	delete(f, "PayUin")
	delete(f, "IdentityRoleID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成员Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
		Uin *int64 `json:"Uin,omitempty" name:"Uin"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitempty" name:"Lang"`

	// 成员名或者成员ID搜索
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Lang")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成员列表
		Items []*OrgMember `json:"Items,omitempty" name:"Items"`

		// 总数目
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest

	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *DescribeOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 企业组织ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgId *int64 `json:"OrgId,omitempty" name:"OrgId"`

		// 创建者UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
		HostUin *int64 `json:"HostUin,omitempty" name:"HostUin"`

		// 创建者昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
		NickName *string `json:"NickName,omitempty" name:"NickName"`

		// 企业组织类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgType *int64 `json:"OrgType,omitempty" name:"OrgType"`

		// 组织管理员：true，组织成员：false
	// 注意：此字段可能返回 null，表示取不到有效值。
		IsManager *bool `json:"IsManager,omitempty" name:"IsManager"`

		// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`

		// 策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`

		// 策略权限
	// 注意：此字段可能返回 null，表示取不到有效值。
		OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`

		// 根节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RootNodeId *int64 `json:"RootNodeId,omitempty" name:"RootNodeId"`

		// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 成员加入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
		JoinTime *string `json:"JoinTime,omitempty" name:"JoinTime"`

		// 是否允许退出。允许：Allow，不允许：Denied。
	// 注意：此字段可能返回 null，表示取不到有效值。
		IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`

		// 代付者Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
		PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

		// 代付者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		PayName *string `json:"PayName,omitempty" name:"PayName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberIdentity struct {

	// 身份ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitempty" name:"IdentityId"`

	// 身份名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityAliasName *string `json:"IdentityAliasName,omitempty" name:"IdentityAliasName"`
}

type OrgMember struct {

	// 成员Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitempty" name:"MemberUin"`

	// 成员名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 成员类型，邀请：Invite， 创建：Create
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberType *string `json:"MemberType,omitempty" name:"MemberType"`

	// 关系策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyType *string `json:"OrgPolicyType,omitempty" name:"OrgPolicyType"`

	// 关系策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPolicyName *string `json:"OrgPolicyName,omitempty" name:"OrgPolicyName"`

	// 关系策略权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgPermission []*OrgPermission `json:"OrgPermission,omitempty" name:"OrgPermission"`

	// 所属节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *int64 `json:"NodeId,omitempty" name:"NodeId"`

	// 所属节点名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 是否允许成员退出。允许：Allow，不允许：Denied。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowQuit *string `json:"IsAllowQuit,omitempty" name:"IsAllowQuit"`

	// 代付者Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayUin *string `json:"PayUin,omitempty" name:"PayUin"`

	// 代付者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayName *string `json:"PayName,omitempty" name:"PayName"`

	// 管理身份
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitempty" name:"OrgIdentity"`

	// 安全信息绑定状态  未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindStatus *string `json:"BindStatus,omitempty" name:"BindStatus"`
}

type OrgPermission struct {

	// 权限Id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 权限名
	Name *string `json:"Name,omitempty" name:"Name"`
}
