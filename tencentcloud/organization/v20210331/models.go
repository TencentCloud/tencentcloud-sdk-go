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

package v20210331

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AcceptJoinShareUnitInvitationRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

type AcceptJoinShareUnitInvitationRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

func (r *AcceptJoinShareUnitInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptJoinShareUnitInvitationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptJoinShareUnitInvitationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptJoinShareUnitInvitationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcceptJoinShareUnitInvitationResponse struct {
	*tchttp.BaseResponse
	Response *AcceptJoinShareUnitInvitationResponseParams `json:"Response"`
}

func (r *AcceptJoinShareUnitInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptJoinShareUnitInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddExternalSAMLIdPCertificateRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// PEM 格式的 X509 证书。  由 SAML 身份提供商提供。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

type AddExternalSAMLIdPCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// PEM 格式的 X509 证书。  由 SAML 身份提供商提供。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

func (r *AddExternalSAMLIdPCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddExternalSAMLIdPCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "X509Certificate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddExternalSAMLIdPCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddExternalSAMLIdPCertificateResponseParams struct {
	// SAML 签名证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddExternalSAMLIdPCertificateResponse struct {
	*tchttp.BaseResponse
	Response *AddExternalSAMLIdPCertificateResponseParams `json:"Response"`
}

func (r *AddExternalSAMLIdPCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddExternalSAMLIdPCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationMemberEmailRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type AddOrganizationMemberEmailRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

func (r *AddOrganizationMemberEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationMemberEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "Email")
	delete(f, "CountryCode")
	delete(f, "Phone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationMemberEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationMemberEmailResponseParams struct {
	// 绑定Id
	BindId *uint64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationMemberEmailResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationMemberEmailResponseParams `json:"Response"`
}

func (r *AddOrganizationMemberEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationMemberEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeRequestParams struct {
	// 父节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	ParentNodeId *uint64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 部门标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AddOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 父节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	ParentNodeId *uint64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 部门标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *AddOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParentNodeId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOrganizationNodeResponseParams struct {
	// 节点ID。
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *AddOrganizationNodeResponseParams `json:"Response"`
}

func (r *AddOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddPermissionPolicyToRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。 
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 权限策略名称，长度最大为 20策略，每个策略长度最大32个字符。如果要添加系统策略，建议使用RolePolicies参数。自定义策略时，数组长度最大为1。
	RolePolicyNames []*string `json:"RolePolicyNames,omitnil,omitempty" name:"RolePolicyNames"`

	// 添加的系统策略详情。
	RolePolicies []*PolicyDetail `json:"RolePolicies,omitnil,omitempty" name:"RolePolicies"`

	// 自定义策略内容。长度：最大 4096 个字符。当RolePolicyType为Inline时，该参数必须配置。关于权限策略的语法和结构，请参见权限策略语法和结构。
	CustomPolicyDocument *string `json:"CustomPolicyDocument,omitnil,omitempty" name:"CustomPolicyDocument"`

	// 自定义策略内容列表（跟RolePolicyNames一一对应）
	CustomPolicyDocuments []*string `json:"CustomPolicyDocuments,omitnil,omitempty" name:"CustomPolicyDocuments"`
}

type AddPermissionPolicyToRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。 
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 权限策略名称，长度最大为 20策略，每个策略长度最大32个字符。如果要添加系统策略，建议使用RolePolicies参数。自定义策略时，数组长度最大为1。
	RolePolicyNames []*string `json:"RolePolicyNames,omitnil,omitempty" name:"RolePolicyNames"`

	// 添加的系统策略详情。
	RolePolicies []*PolicyDetail `json:"RolePolicies,omitnil,omitempty" name:"RolePolicies"`

	// 自定义策略内容。长度：最大 4096 个字符。当RolePolicyType为Inline时，该参数必须配置。关于权限策略的语法和结构，请参见权限策略语法和结构。
	CustomPolicyDocument *string `json:"CustomPolicyDocument,omitnil,omitempty" name:"CustomPolicyDocument"`

	// 自定义策略内容列表（跟RolePolicyNames一一对应）
	CustomPolicyDocuments []*string `json:"CustomPolicyDocuments,omitnil,omitempty" name:"CustomPolicyDocuments"`
}

func (r *AddPermissionPolicyToRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddPermissionPolicyToRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "RolePolicyType")
	delete(f, "RolePolicyNames")
	delete(f, "RolePolicies")
	delete(f, "CustomPolicyDocument")
	delete(f, "CustomPolicyDocuments")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddPermissionPolicyToRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddPermissionPolicyToRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddPermissionPolicyToRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *AddPermissionPolicyToRoleConfigurationResponseParams `json:"Response"`
}

func (r *AddPermissionPolicyToRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddPermissionPolicyToRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitMembersRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享成员列表。最大10个。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

type AddShareUnitMembersRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享成员列表。最大10个。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

func (r *AddShareUnitMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Members")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShareUnitMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddShareUnitMembersResponse struct {
	*tchttp.BaseResponse
	Response *AddShareUnitMembersResponseParams `json:"Response"`
}

func (r *AddShareUnitMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitRequestParams struct {
	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元地域。可通过接口[DescribeShareAreas](https://cloud.tencent.com/document/product/850/103050)获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号，默认值：1
	ShareScope *uint64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`
}

type AddShareUnitRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元地域。可通过接口[DescribeShareAreas](https://cloud.tencent.com/document/product/850/103050)获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号，默认值：1
	ShareScope *uint64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`
}

func (r *AddShareUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Area")
	delete(f, "Description")
	delete(f, "ShareScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShareUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitResourcesRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ProductResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type AddShareUnitResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ProductResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *AddShareUnitResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Type")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddShareUnitResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitResourcesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddShareUnitResourcesResponse struct {
	*tchttp.BaseResponse
	Response *AddShareUnitResourcesResponseParams `json:"Response"`
}

func (r *AddShareUnitResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddShareUnitResponseParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddShareUnitResponse struct {
	*tchttp.BaseResponse
	Response *AddShareUnitResponseParams `json:"Response"`
}

func (r *AddShareUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddShareUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserToGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AddUserToGroupRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "UserId")
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

// Predefined struct for user
type AttachPolicyRequestParams struct {
	// 绑定策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type AttachPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 绑定策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *AttachPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetId")
	delete(f, "TargetType")
	delete(f, "PolicyId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachPolicyResponseParams `json:"Response"`
}

func (r *AttachPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthNode struct {
	// 互信主体关系ID
	RelationId *int64 `json:"RelationId,omitnil,omitempty" name:"RelationId"`

	// 互信主体名称
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// 主体管理员
	Manager *MemberMainInfo `json:"Manager,omitnil,omitempty" name:"Manager"`
}

type AuthRelationFile struct {
	// 文件名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件路径。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

// Predefined struct for user
type BindOrganizationMemberAuthAccountRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以调用[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

type BindOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以调用[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
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

// Predefined struct for user
type BindOrganizationMemberAuthAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindOrganizationMemberAuthAccountResponseParams `json:"Response"`
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

// Predefined struct for user
type BindOrganizationPolicySubAccountRequestParams struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

type BindOrganizationPolicySubAccountRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

func (r *BindOrganizationPolicySubAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationPolicySubAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindOrganizationPolicySubAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindOrganizationPolicySubAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindOrganizationPolicySubAccountResponse struct {
	*tchttp.BaseResponse
	Response *BindOrganizationPolicySubAccountResponseParams `json:"Response"`
}

func (r *BindOrganizationPolicySubAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindOrganizationPolicySubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织子账号Uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`
}

type CancelOrganizationMemberAuthAccountRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织子账号Uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`
}

func (r *CancelOrganizationMemberAuthAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelOrganizationMemberAuthAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationMemberAuthAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelOrganizationMemberAuthAccountResponse struct {
	*tchttp.BaseResponse
	Response *CancelOrganizationMemberAuthAccountResponseParams `json:"Response"`
}

func (r *CancelOrganizationMemberAuthAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationMemberAuthAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationPolicySubAccountRequestParams struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

type CancelOrganizationPolicySubAccountRequest struct {
	*tchttp.BaseRequest
	
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 组织管理员子账号Uin列表。最大5个
	OrgSubAccountUins []*int64 `json:"OrgSubAccountUins,omitnil,omitempty" name:"OrgSubAccountUins"`
}

func (r *CancelOrganizationPolicySubAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationPolicySubAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "OrgSubAccountUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelOrganizationPolicySubAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelOrganizationPolicySubAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelOrganizationPolicySubAccountResponse struct {
	*tchttp.BaseResponse
	Response *CancelOrganizationPolicySubAccountResponseParams `json:"Response"`
}

func (r *CancelOrganizationPolicySubAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelOrganizationPolicySubAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAccountDeleteRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type CheckAccountDeleteRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *CheckAccountDeleteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAccountDeleteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAccountDeleteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAccountDeleteResponseParams struct {
	// 成员是否允许删除。 true-是、false-否
	AllowDelete *bool `json:"AllowDelete,omitnil,omitempty" name:"AllowDelete"`

	// 不允许删除原因。
	NotAllowReason *NotAllowReason `json:"NotAllowReason,omitnil,omitempty" name:"NotAllowReason"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAccountDeleteResponse struct {
	*tchttp.BaseResponse
	Response *CheckAccountDeleteResponseParams `json:"Response"`
}

func (r *CheckAccountDeleteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAccountDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearExternalSAMLIdentityProviderRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type ClearExternalSAMLIdentityProviderRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *ClearExternalSAMLIdentityProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearExternalSAMLIdentityProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearExternalSAMLIdentityProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearExternalSAMLIdentityProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearExternalSAMLIdentityProviderResponse struct {
	*tchttp.BaseResponse
	Response *ClearExternalSAMLIdentityProviderResponseParams `json:"Response"`
}

func (r *ClearExternalSAMLIdentityProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearExternalSAMLIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的名称。  格式：允许英文字母、数字和特殊字符-。 长度：最大 128 个字符。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组类型  Manual：手动创建，Synchronized：外部导入
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的名称。  格式：允许英文字母、数字和特殊字符-。 长度：最大 128 个字符。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组类型  Manual：手动创建，Synchronized：外部导入
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`
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
	delete(f, "ZoneId")
	delete(f, "GroupName")
	delete(f, "Description")
	delete(f, "GroupType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// 用户组信息。
	GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

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
type CreateOrgServiceAssignRequestParams struct {
	// 委派管理员Uin列表。 最大长度20个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 集团服务ID。和集团服务产品标识二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品标识。和集团服务ID二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 委派管理员管理范围。 取值：1-全部成员 2-部分成员，默认值1
	ManagementScope *uint64 `json:"ManagementScope,omitnil,omitempty" name:"ManagementScope"`

	// 管理的成员Uin列表。ManagementScope为2时该参数有效
	ManagementScopeUins []*int64 `json:"ManagementScopeUins,omitnil,omitempty" name:"ManagementScopeUins"`

	// 管理的部门ID列表。ManagementScope为2时该参数有效
	ManagementScopeNodeIds []*int64 `json:"ManagementScopeNodeIds,omitnil,omitempty" name:"ManagementScopeNodeIds"`
}

type CreateOrgServiceAssignRequest struct {
	*tchttp.BaseRequest
	
	// 委派管理员Uin列表。 最大长度20个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 集团服务ID。和集团服务产品标识二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品标识。和集团服务ID二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 委派管理员管理范围。 取值：1-全部成员 2-部分成员，默认值1
	ManagementScope *uint64 `json:"ManagementScope,omitnil,omitempty" name:"ManagementScope"`

	// 管理的成员Uin列表。ManagementScope为2时该参数有效
	ManagementScopeUins []*int64 `json:"ManagementScopeUins,omitnil,omitempty" name:"ManagementScopeUins"`

	// 管理的部门ID列表。ManagementScope为2时该参数有效
	ManagementScopeNodeIds []*int64 `json:"ManagementScopeNodeIds,omitnil,omitempty" name:"ManagementScopeNodeIds"`
}

func (r *CreateOrgServiceAssignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgServiceAssignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUins")
	delete(f, "ServiceId")
	delete(f, "Product")
	delete(f, "ManagementScope")
	delete(f, "ManagementScopeUins")
	delete(f, "ManagementScopeNodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrgServiceAssignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrgServiceAssignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrgServiceAssignResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrgServiceAssignResponseParams `json:"Response"`
}

func (r *CreateOrgServiceAssignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrgServiceAssignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationIdentityRequestParams struct {
	// 身份名称
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份名称
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 身份策略
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityAliasName")
	delete(f, "IdentityPolicy")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationIdentityResponseParams struct {
	// 身份ID
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationIdentityResponseParams `json:"Response"`
}

func (r *CreateOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberAuthIdentityRequestParams struct {
	// 成员Uin列表。最多10个
	MemberUins []*uint64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 身份Id列表。最多5个，可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityIds []*uint64 `json:"IdentityIds,omitnil,omitempty" name:"IdentityIds"`
}

type CreateOrganizationMemberAuthIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin列表。最多10个
	MemberUins []*uint64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 身份Id列表。最多5个，可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityIds []*uint64 `json:"IdentityIds,omitnil,omitempty" name:"IdentityIds"`
}

func (r *CreateOrganizationMemberAuthIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberAuthIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUins")
	delete(f, "IdentityIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberAuthIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberAuthIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberAuthIdentityResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberAuthIdentityResponseParams `json:"Response"`
}

func (r *CreateOrganizationMemberAuthIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMemberAuthIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberPolicyRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略名。最大长度为128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以调用[DescribeOrganizationMemberAuthIdentities](https://cloud.tencent.com/document/product/850/82936)获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationMemberPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略名。最大长度为128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以调用[DescribeOrganizationMemberAuthIdentities](https://cloud.tencent.com/document/product/850/82936)获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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

// Predefined struct for user
type CreateOrganizationMemberPolicyResponseParams struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberPolicyResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateOrganizationMemberRequestParams struct {
	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨（若需要开启资金划拨权限，请联系您的商务经理内部开通。）、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析、9-预算管理、10-信用额度设置（若需要开启信用额度设置权限，请联系您的商务经理内部开通。），1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 账号名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 成员创建记录ID。创建异常重试时需要
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 成员访问身份ID列表。可以调用ListOrganizationIdentity获取，1默认支持
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// 认证主体关系ID。给不同主体创建成员时需要，可以调用DescribeOrganizationAuthNode获取
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`

	// 成员标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨（若需要开启资金划拨权限，请联系您的商务经理内部开通。）、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析、9-预算管理、10-信用额度设置（若需要开启信用额度设置权限，请联系您的商务经理内部开通。），1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 账号名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 成员创建记录ID。创建异常重试时需要
	RecordId *int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 成员访问身份ID列表。可以调用ListOrganizationIdentity获取，1默认支持
	IdentityRoleID []*uint64 `json:"IdentityRoleID,omitnil,omitempty" name:"IdentityRoleID"`

	// 认证主体关系ID。给不同主体创建成员时需要，可以调用DescribeOrganizationAuthNode获取
	AuthRelationId *int64 `json:"AuthRelationId,omitnil,omitempty" name:"AuthRelationId"`

	// 成员标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "AuthRelationId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMemberResponseParams struct {
	// 成员Uin。
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMemberResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateOrganizationMembersPolicyRequestParams struct {
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 策略名。长度1～128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 策略名。长度1～128个字符，支持英文字母、数字、符号+=,.@_-
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员访问身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateOrganizationMembersPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUins")
	delete(f, "PolicyName")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationMembersPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationMembersPolicyResponseParams struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *CreateOrganizationMembersPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationRequestParams struct {

}

type CreateOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CreateOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrganizationResponseParams struct {
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 创建者昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrganizationResponseParams `json:"Response"`
}

func (r *CreateOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyRequestParams struct {
	// 策略名。
	// 长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略名。
	// 长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	delete(f, "Name")
	delete(f, "Content")
	delete(f, "Type")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePolicyResponseParams struct {
	// 策略ID
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
type CreateRoleAssignmentRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 授权成员账号信息，最多授权50条。
	RoleAssignmentInfo []*RoleAssignmentInfo `json:"RoleAssignmentInfo,omitnil,omitempty" name:"RoleAssignmentInfo"`
}

type CreateRoleAssignmentRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 授权成员账号信息，最多授权50条。
	RoleAssignmentInfo []*RoleAssignmentInfo `json:"RoleAssignmentInfo,omitnil,omitempty" name:"RoleAssignmentInfo"`
}

func (r *CreateRoleAssignmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleAssignmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleAssignmentInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleAssignmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleAssignmentResponseParams struct {
	// 任务详情。
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleAssignmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleAssignmentResponseParams `json:"Response"`
}

func (r *CreateRoleAssignmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleAssignmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleConfigurationRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置名称。格式：包含英文字母、数字或短划线（-）。 长度：最大 128 个字符。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 权限配置的描述。 长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 会话持续时间。 CIC用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。 单位：秒。 取值范围：900 ~ 43200（15 分钟~12 小时）。 默认值：3600（1 小时）。
	SessionDuration *int64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 初始访问页面。 CIC用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。 该页面必须是腾讯云控制台页面。默认为空，表示跳转到腾讯云控制台首页。
	RelayState *string `json:"RelayState,omitnil,omitempty" name:"RelayState"`
}

type CreateRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置名称。格式：包含英文字母、数字或短划线（-）。 长度：最大 128 个字符。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 权限配置的描述。 长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 会话持续时间。 CIC用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。 单位：秒。 取值范围：900 ~ 43200（15 分钟~12 小时）。 默认值：3600（1 小时）。
	SessionDuration *int64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 初始访问页面。 CIC用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。 该页面必须是腾讯云控制台页面。默认为空，表示跳转到腾讯云控制台首页。
	RelayState *string `json:"RelayState,omitnil,omitempty" name:"RelayState"`
}

func (r *CreateRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationName")
	delete(f, "Description")
	delete(f, "SessionDuration")
	delete(f, "RelayState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleConfigurationResponseParams struct {
	// 配置访问详情
	RoleConfigurationInfo *RoleConfiguration `json:"RoleConfigurationInfo,omitnil,omitempty" name:"RoleConfigurationInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleConfigurationResponseParams `json:"Response"`
}

func (r *CreateRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSCIMCredentialRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过期时间（秒），最小1小时，最大99年。如果不传则默认一年过期
	ExpireDuration *uint64 `json:"ExpireDuration,omitnil,omitempty" name:"ExpireDuration"`
}

type CreateSCIMCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 过期时间（秒），最小1小时，最大99年。如果不传则默认一年过期
	ExpireDuration *uint64 `json:"ExpireDuration,omitnil,omitempty" name:"ExpireDuration"`
}

func (r *CreateSCIMCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSCIMCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ExpireDuration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSCIMCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSCIMCredentialResponseParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// SCIM密钥类型。
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// SCIM 密钥的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// SCIM 密钥的过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// SCIM密钥状态，Enabled已开启，Disabled已关闭。
	CredentialStatus *string `json:"CredentialStatus,omitnil,omitempty" name:"CredentialStatus"`

	// SCIM密钥。
	CredentialSecret *string `json:"CredentialSecret,omitnil,omitempty" name:"CredentialSecret"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSCIMCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateSCIMCredentialResponseParams `json:"Response"`
}

func (r *CreateSCIMCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSCIMCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户名称。空间内必须唯一。不支持修改。  格式：包含数字、英文字母和特殊符号+ = , . @ - _ 。  长度：最大 64 个字符
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的姓。  长度：最大 64 个字符。
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// 用户的名。  长度：最大 64 个字符。
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// 用户的显示名称。  长度：最大 256 个字符。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。  长度：最大 128 个字符。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户的状态。取值：  Enabled（默认值）：启用。 Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 是否需要重置密码： true: 需要重置  false: 不需要重置密码。 默认false
	NeedResetPassword *bool `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户名称。空间内必须唯一。不支持修改。  格式：包含数字、英文字母和特殊符号+ = , . @ - _ 。  长度：最大 64 个字符
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的姓。  长度：最大 64 个字符。
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// 用户的名。  长度：最大 64 个字符。
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// 用户的显示名称。  长度：最大 256 个字符。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。  长度：最大 1024 个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。  长度：最大 128 个字符。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户的状态。取值：  Enabled（默认值）：启用。 Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 是否需要重置密码： true: 需要重置  false: 不需要重置密码。 默认false
	NeedResetPassword *bool `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`
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
	delete(f, "ZoneId")
	delete(f, "UserName")
	delete(f, "FirstName")
	delete(f, "LastName")
	delete(f, "DisplayName")
	delete(f, "Description")
	delete(f, "Email")
	delete(f, "UserStatus")
	delete(f, "UserType")
	delete(f, "NeedResetPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 用户详情
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

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
type CreateUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM用户同步信息。
	UserSyncProvisionings []*UserSyncProvisioning `json:"UserSyncProvisionings,omitnil,omitempty" name:"UserSyncProvisionings"`
}

type CreateUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM用户同步信息。
	UserSyncProvisionings []*UserSyncProvisioning `json:"UserSyncProvisionings,omitnil,omitempty" name:"UserSyncProvisionings"`
}

func (r *CreateUserSyncProvisioningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSyncProvisioningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "UserSyncProvisionings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserSyncProvisioningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserSyncProvisioningResponseParams struct {
	// 任务详细。
	Tasks []*UserProvisioningsTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *CreateUserSyncProvisioningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	delete(f, "ZoneId")
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
type DeleteOrgServiceAssignRequestParams struct {
	// 委派管理员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 集团服务ID。和集团服务产品标识二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品标识。和集团服务ID二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteOrgServiceAssignRequest struct {
	*tchttp.BaseRequest
	
	// 委派管理员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 集团服务ID。和集团服务产品标识二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品标识。和集团服务ID二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteOrgServiceAssignRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgServiceAssignRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "ServiceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrgServiceAssignRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrgServiceAssignResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrgServiceAssignResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrgServiceAssignResponseParams `json:"Response"`
}

func (r *DeleteOrgServiceAssignResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrgServiceAssignResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationIdentityRequestParams struct {
	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DeleteOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DeleteOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationIdentityResponseParams `json:"Response"`
}

func (r *DeleteOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberAuthIdentityRequestParams struct {
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DeleteOrganizationMemberAuthIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DeleteOrganizationMemberAuthIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberAuthIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMemberAuthIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMemberAuthIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMemberAuthIdentityResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMemberAuthIdentityResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMemberAuthIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMemberAuthIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersPolicyRequestParams struct {
	// 访问策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DeleteOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 访问策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DeleteOrganizationMembersPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersRequestParams struct {
	// 被删除成员的Uin列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DeleteOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// 被删除成员的Uin列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DeleteOrganizationMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationMembersResponseParams `json:"Response"`
}

func (r *DeleteOrganizationMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesRequestParams struct {
	// 节点ID列表。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId []*int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

type DeleteOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// 节点ID列表。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId []*int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`
}

func (r *DeleteOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationNodesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationNodesResponseParams `json:"Response"`
}

func (r *DeleteOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationRequestParams struct {

}

type DeleteOrganizationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOrganizationResponseParams `json:"Response"`
}

func (r *DeleteOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePolicyRequestParams struct {
	// 需要删除的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeletePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "Type")
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
type DeleteRoleAssignmentRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-********）， 当PrincipalType取值为User时，该值为用户 ID（u-********）。 	
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 当您移除一个集团账号目标账号上使用某权限配置的最后一个授权时，是否同时解除权限配置部署。取值： DeprovisionForLastRoleAssignmentOnAccount：解除权限配置部署。 None（默认值）：不解除权限配置部署。
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitnil,omitempty" name:"DeprovisionStrategy"`
}

type DeleteRoleAssignmentRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-********）， 当PrincipalType取值为User时，该值为用户 ID（u-********）。 	
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 当您移除一个集团账号目标账号上使用某权限配置的最后一个授权时，是否同时解除权限配置部署。取值： DeprovisionForLastRoleAssignmentOnAccount：解除权限配置部署。 None（默认值）：不解除权限配置部署。
	DeprovisionStrategy *string `json:"DeprovisionStrategy,omitnil,omitempty" name:"DeprovisionStrategy"`
}

func (r *DeleteRoleAssignmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleAssignmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "TargetType")
	delete(f, "TargetUin")
	delete(f, "PrincipalType")
	delete(f, "PrincipalId")
	delete(f, "DeprovisionStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleAssignmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleAssignmentResponseParams struct {
	// 任务详情
	Task *TaskInfo `json:"Task,omitnil,omitempty" name:"Task"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleAssignmentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleAssignmentResponseParams `json:"Response"`
}

func (r *DeleteRoleAssignmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleAssignmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

type DeleteRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

func (r *DeleteRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleConfigurationResponseParams `json:"Response"`
}

func (r *DeleteRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSCIMCredentialRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

type DeleteSCIMCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

func (r *DeleteSCIMCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSCIMCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CredentialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSCIMCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSCIMCredentialResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSCIMCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSCIMCredentialResponseParams `json:"Response"`
}

func (r *DeleteSCIMCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSCIMCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitMembersRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 成员列表。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

type DeleteShareUnitMembersRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 成员列表。
	Members []*ShareMember `json:"Members,omitnil,omitempty" name:"Members"`
}

func (r *DeleteShareUnitMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Members")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareUnitMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareUnitMembersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareUnitMembersResponseParams `json:"Response"`
}

func (r *DeleteShareUnitMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

type DeleteShareUnitRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

func (r *DeleteShareUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitResourcesRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ShareResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type DeleteShareUnitResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享资源列表。最大10个。
	Resources []*ShareResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *DeleteShareUnitResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Type")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteShareUnitResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitResourcesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareUnitResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareUnitResourcesResponseParams `json:"Response"`
}

func (r *DeleteShareUnitResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteShareUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteShareUnitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteShareUnitResponseParams `json:"Response"`
}

func (r *DeleteShareUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteShareUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	delete(f, "ZoneId")
	delete(f, "UserId")
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
type DeleteUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

type DeleteUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

func (r *DeleteUserSyncProvisioningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserSyncProvisioningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "UserProvisioningId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserSyncProvisioningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserSyncProvisioningResponseParams struct {
	// 任务详情。
	Tasks *UserProvisioningsTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *DeleteUserSyncProvisioningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEffectivePolicyRequestParams struct {
	// 账号uin或者节点id。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`
}

type DescribeEffectivePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 账号uin或者节点id。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`
}

func (r *DescribeEffectivePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEffectivePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEffectivePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEffectivePolicyResponseParams struct {
	// 有效策略。
	EffectivePolicy *EffectivePolicy `json:"EffectivePolicy,omitnil,omitempty" name:"EffectivePolicy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEffectivePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEffectivePolicyResponseParams `json:"Response"`
}

func (r *DescribeEffectivePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEffectivePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentityCenterRequestParams struct {

}

type DescribeIdentityCenterRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIdentityCenterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentityCenterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdentityCenterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentityCenterResponseParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 服务开启状态，Disabled代表未开通，Enabled代表已开通
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// SCIM 同步状态。Enabled：启用。 Disabled：禁用。
	ScimSyncStatus *string `json:"ScimSyncStatus,omitnil,omitempty" name:"ScimSyncStatus"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIdentityCenterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdentityCenterResponseParams `json:"Response"`
}

func (r *DescribeIdentityCenterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentityCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationAuthNodeRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 互信主体名称。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`
}

type DescribeOrganizationAuthNodeRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 互信主体名称。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`
}

func (r *DescribeOrganizationAuthNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationAuthNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "AuthName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationAuthNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationAuthNodeResponseParams struct {
	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 条目详情。
	Items []*AuthNode `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationAuthNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationAuthNodeResponseParams `json:"Response"`
}

func (r *DescribeOrganizationAuthNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationAuthNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMemberRequestParams struct {
	// 查询开始月份。格式：yyyy-mm，例如：2021-01。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份。
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

type DescribeOrganizationFinancialByMemberRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始月份。格式：yyyy-mm，例如：2021-01。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份。
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EndMonth")
	delete(f, "MemberUins")
	delete(f, "ProductCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationFinancialByMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMemberResponseParams struct {
	// 当月总消耗。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 成员消耗详情。
	Items []*OrgMemberFinancial `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationFinancialByMemberResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationFinancialByMemberResponseParams `json:"Response"`
}

func (r *DescribeOrganizationFinancialByMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMonthRequestParams struct {
	// 查询月数。取值范围：1~6，默认值：6
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

type DescribeOrganizationFinancialByMonthRequest struct {
	*tchttp.BaseRequest
	
	// 查询月数。取值范围：1~6，默认值：6
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByMonthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMonthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "EndMonth")
	delete(f, "MemberUins")
	delete(f, "ProductCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationFinancialByMonthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByMonthResponseParams struct {
	// 产品消耗详情。
	Items []*OrgFinancialByMonth `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationFinancialByMonthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationFinancialByMonthResponseParams `json:"Response"`
}

func (r *DescribeOrganizationFinancialByMonthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByProductRequestParams struct {
	// 查询开始月份。格式：yyyy-mm，例如：2021-01
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

type DescribeOrganizationFinancialByProductRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始月份。格式：yyyy-mm，例如：2021-01
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 限制数目。取值范围：1~50，默认值：10	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结束月份。格式：yyyy-mm，例如：2021-01,默认值为查询开始月份
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 查询成员列表。 最大100个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 查询产品列表。 最大100个
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`
}

func (r *DescribeOrganizationFinancialByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EndMonth")
	delete(f, "MemberUins")
	delete(f, "ProductCodes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationFinancialByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationFinancialByProductResponseParams struct {
	// 当月总消耗。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 产品消耗详情。
	Items []*OrgProductFinancial `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationFinancialByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationFinancialByProductResponseParams `json:"Response"`
}

func (r *DescribeOrganizationFinancialByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationFinancialByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeOrganizationMemberAuthAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 策略ID。可以通过[DescribeOrganizationMemberPolicies](https://cloud.tencent.com/document/product/850/82935)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeOrganizationMemberAuthAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthAccountsResponseParams struct {
	// 列表
	Items []*OrgMemberAuthAccount `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthAccountsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 组织成员Uin。入参MemberUin与IdentityId至少填写一个
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。入参MemberUin与IdentityId至少填写一个, 可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

type DescribeOrganizationMemberAuthIdentitiesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 组织成员Uin。入参MemberUin与IdentityId至少填写一个
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 身份ID。入参MemberUin与IdentityId至少填写一个, 可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "IdentityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberAuthIdentitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberAuthIdentitiesResponseParams struct {
	// 授权身份列表。
	Items []*OrgMemberAuthIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberAuthIdentitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberAuthIdentitiesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberAuthIdentitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberAuthIdentitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberEmailBindRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type DescribeOrganizationMemberEmailBindRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *DescribeOrganizationMemberEmailBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberEmailBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberEmailBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberEmailBindResponseParams struct {
	// 绑定ID。
	BindId *uint64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 申请时间。
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 安全手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 绑定状态。    未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// 绑定时间。
	BindTime *string `json:"BindTime,omitnil,omitempty" name:"BindTime"`

	// 失败说明。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 安全手机绑定状态 。 未绑定：0，已绑定：1
	PhoneBind *uint64 `json:"PhoneBind,omitnil,omitempty" name:"PhoneBind"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberEmailBindResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberEmailBindResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberEmailBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberEmailBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 搜索关键字。可用于策略名或描述搜索
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeOrganizationMemberPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 搜索关键字。可用于策略名或描述搜索
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeOrganizationMemberPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMemberPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMemberPoliciesResponseParams struct {
	// 列表。
	Items []*OrgMemberPolicy `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMemberPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMemberPoliciesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMemberPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMemberPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersAuthPolicyRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 集团管理员子账号uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`

	// 成员访问策略Id。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeOrganizationMembersAuthPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 成员uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 集团管理员子账号uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`

	// 成员访问策略Id。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeOrganizationMembersAuthPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersAuthPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MemberUin")
	delete(f, "OrgSubAccountUin")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMembersAuthPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersAuthPolicyResponseParams struct {
	// 访问授权策略列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*OrgMembersAuthPolicy `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMembersAuthPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMembersAuthPolicyResponseParams `json:"Response"`
}

func (r *DescribeOrganizationMembersAuthPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationMembersAuthPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 成员名称或者成员ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 主体名称搜索。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// 可信服务产品简称。可信服务管理员查询时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 成员标签搜索列表，最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 组织单元ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 组织单元名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type DescribeOrganizationMembersRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 成员名称或者成员ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 主体名称搜索。
	AuthName *string `json:"AuthName,omitnil,omitempty" name:"AuthName"`

	// 可信服务产品简称。可信服务管理员查询时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 成员标签搜索列表，最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 组织单元ID
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 组织单元名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
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
	delete(f, "AuthName")
	delete(f, "Product")
	delete(f, "Tags")
	delete(f, "NodeId")
	delete(f, "NodeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationMembersResponseParams struct {
	// 成员列表。
	Items []*OrgMember `json:"Items,omitnil,omitempty" name:"Items"`

	// 总数目。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationMembersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeOrganizationNodesRequestParams struct {
	// 限制数目。最大50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 部门标签搜索列表，最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeOrganizationNodesRequest struct {
	*tchttp.BaseRequest
	
	// 限制数目。最大50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 部门标签搜索列表，最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeOrganizationNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationNodesResponseParams struct {
	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表详情。
	Items []*OrgNode `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationNodesResponseParams `json:"Response"`
}

func (r *DescribeOrganizationNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationRequestParams struct {
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 可信服务产品简称。查询是否该可信服务管理员时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 可信服务产品简称。查询是否该可信服务管理员时必须指定
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
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
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationResponseParams struct {
	// 企业组织ID。
	OrgId *int64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`

	// 创建者UIN。
	HostUin *int64 `json:"HostUin,omitnil,omitempty" name:"HostUin"`

	// 创建者昵称。
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 企业组织类型。
	OrgType *int64 `json:"OrgType,omitnil,omitempty" name:"OrgType"`

	// 是否组织管理员。是：true ，否：false
	IsManager *bool `json:"IsManager,omitnil,omitempty" name:"IsManager"`

	// 策略类型。财务管理：Financial
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// 策略名。
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// 成员财务权限列表。
	OrgPermission []*OrgPermission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// 组织根节点ID。
	RootNodeId *int64 `json:"RootNodeId,omitnil,omitempty" name:"RootNodeId"`

	// 组织创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 成员加入时间。
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`

	// 成员是否允许退出。允许：Allow，不允许：Denied
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 代付者名称。
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// 是否可信服务管理员。是：true，否：false
	IsAssignManager *bool `json:"IsAssignManager,omitnil,omitempty" name:"IsAssignManager"`

	// 是否实名主体管理员。是：true，否：false
	IsAuthManager *bool `json:"IsAuthManager,omitnil,omitempty" name:"IsAuthManager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePolicyConfigRequestParams struct {
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值0，取值范围：0-服务控制策略、1-标签策略
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribePolicyConfigRequest struct {
	*tchttp.BaseRequest
	
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值0，取值范围：0-服务控制策略、1-标签策略
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribePolicyConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyConfigResponseParams struct {
	// 开启状态。0-未开启、1-开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 策略类型。SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyConfigResponseParams `json:"Response"`
}

func (r *DescribePolicyConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyRequestParams struct {
	// 策略Id。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type DescribePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *DescribePolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "PolicyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyResponseParams struct {
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型。1-自定义 2-预设策略
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略文档。
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 策略更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 策略创建时间。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyResponseParams `json:"Response"`
}

func (r *DescribePolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceToShareMemberRequestParams struct {
	// 共享地域。可通过接口[DescribeShareAreas](https://cloud.tencent.com/document/product/850/103050)获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字，支持业务资源ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 业务资源ID。最大50个
	ProductResourceIds []*string `json:"ProductResourceIds,omitnil,omitempty" name:"ProductResourceIds"`
}

type DescribeResourceToShareMemberRequest struct {
	*tchttp.BaseRequest
	
	// 共享地域。可通过接口[DescribeShareAreas](https://cloud.tencent.com/document/product/850/103050)获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字，支持业务资源ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 业务资源ID。最大50个
	ProductResourceIds []*string `json:"ProductResourceIds,omitnil,omitempty" name:"ProductResourceIds"`
}

func (r *DescribeResourceToShareMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceToShareMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "Type")
	delete(f, "ProductResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceToShareMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceToShareMemberResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*ShareResourceToMember `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceToShareMemberResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceToShareMemberResponseParams `json:"Response"`
}

func (r *DescribeResourceToShareMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceToShareMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareAreasRequestParams struct {
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

type DescribeShareAreasRequest struct {
	*tchttp.BaseRequest
	
	// 国际站：en，国内站：zh
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`
}

func (r *DescribeShareAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareAreasResponseParams struct {
	// 详情
	Items []*ShareArea `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareAreasResponseParams `json:"Response"`
}

func (r *DescribeShareAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitMembersRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持成员Uin搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeShareUnitMembersRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持成员Uin搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeShareUnitMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareUnitMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitMembersResponseParams struct {
	// 总数目。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 共享单元成员列表。
	Items []*ShareUnitMember `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareUnitMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareUnitMembersResponseParams `json:"Response"`
}

func (r *DescribeShareUnitMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitResourcesRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持产品资源ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeShareUnitResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持产品资源ID搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 共享资源类型。支持共享的资源类型,请参见[资源共享概述](https://cloud.tencent.com/document/product/850/59489)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeShareUnitResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareUnitResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitResourcesResponseParams struct {
	// 总数目。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 共享单元资源列表。
	Items []*ShareUnitResource `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareUnitResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareUnitResourcesResponseParams `json:"Response"`
}

func (r *DescribeShareUnitResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitsRequestParams struct {
	// 共享单元地域。可通过接口[DescribeShareAreas](https://cloud.tencent.com/document/product/850/103050)获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持UnitId和Name搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeShareUnitsRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元地域。可通过接口[DescribeShareAreas](https://cloud.tencent.com/document/product/850/103050)获取支持共享的地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键字。支持UnitId和Name搜索。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeShareUnitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Area")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShareUnitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShareUnitsResponseParams struct {
	// 总数目。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 共享单元列表。
	Items []*ManagerShareUnit `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeShareUnitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShareUnitsResponseParams `json:"Response"`
}

func (r *DescribeShareUnitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShareUnitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPolicyRequestParams struct {
	// 解绑策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DetachPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 解绑策略目标ID。成员Uin或部门ID
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目标类型。取值范围：NODE-部门、MEMBER-成员
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略ID。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DetachPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetId")
	delete(f, "TargetType")
	delete(f, "PolicyId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachPolicyResponseParams `json:"Response"`
}

func (r *DetachPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisablePolicyTypeRequestParams struct {
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type DisablePolicyTypeRequest struct {
	*tchttp.BaseRequest
	
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *DisablePolicyTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisablePolicyTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "PolicyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisablePolicyTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisablePolicyTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisablePolicyTypeResponse struct {
	*tchttp.BaseResponse
	Response *DisablePolicyTypeResponseParams `json:"Response"`
}

func (r *DisablePolicyTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisablePolicyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismantleRoleConfigurationRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type DismantleRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *DismantleRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismantleRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "TargetType")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DismantleRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DismantleRoleConfigurationResponseParams struct {
	// 任务详情。
	Task *RoleProvisioningsTask `json:"Task,omitnil,omitempty" name:"Task"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DismantleRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DismantleRoleConfigurationResponseParams `json:"Response"`
}

func (r *DismantleRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DismantleRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EffectivePolicy struct {
	// 目标ID。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 有效策略内容。
	PolicyContent *string `json:"PolicyContent,omitnil,omitempty" name:"PolicyContent"`

	// 有效策略更新时间。
	LastUpdatedTimestamp *uint64 `json:"LastUpdatedTimestamp,omitnil,omitempty" name:"LastUpdatedTimestamp"`
}

// Predefined struct for user
type EnablePolicyTypeRequestParams struct {
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type EnablePolicyTypeRequest struct {
	*tchttp.BaseRequest
	
	// 企业组织Id。可以调用[DescribeOrganization](https://cloud.tencent.com/document/product/850/67059)获取
	OrganizationId *uint64 `json:"OrganizationId,omitnil,omitempty" name:"OrganizationId"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

func (r *EnablePolicyTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnablePolicyTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrganizationId")
	delete(f, "PolicyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnablePolicyTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnablePolicyTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnablePolicyTypeResponse struct {
	*tchttp.BaseResponse
	Response *EnablePolicyTypeResponseParams `json:"Response"`
}

func (r *EnablePolicyTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnablePolicyTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetExternalSAMLIdentityProviderRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetExternalSAMLIdentityProviderRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetExternalSAMLIdentityProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetExternalSAMLIdentityProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetExternalSAMLIdentityProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetExternalSAMLIdentityProviderResponseParams struct {
	// saml 身份提供商配置信息。
	SAMLIdentityProviderConfiguration *SAMLIdentityProviderConfiguration `json:"SAMLIdentityProviderConfiguration,omitnil,omitempty" name:"SAMLIdentityProviderConfiguration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetExternalSAMLIdentityProviderResponse struct {
	*tchttp.BaseResponse
	Response *GetExternalSAMLIdentityProviderResponseParams `json:"Response"`
}

func (r *GetExternalSAMLIdentityProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetExternalSAMLIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type GetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	delete(f, "ZoneId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupResponseParams struct {
	// 用户组信息
	GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

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
type GetProvisioningTaskStatusRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetProvisioningTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetProvisioningTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProvisioningTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProvisioningTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProvisioningTaskStatusResponseParams struct {
	// 任务状态信息。
	TaskStatus *TaskStatus `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetProvisioningTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetProvisioningTaskStatusResponseParams `json:"Response"`
}

func (r *GetProvisioningTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProvisioningTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

type GetRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

func (r *GetRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRoleConfigurationResponseParams struct {
	// 权限配置详情
	RoleConfigurationInfo *RoleConfiguration `json:"RoleConfigurationInfo,omitnil,omitempty" name:"RoleConfigurationInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *GetRoleConfigurationResponseParams `json:"Response"`
}

func (r *GetRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSCIMSynchronizationStatusRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetSCIMSynchronizationStatusRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetSCIMSynchronizationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSCIMSynchronizationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSCIMSynchronizationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSCIMSynchronizationStatusResponseParams struct {
	// SCIM 同步状态。Enabled：启用。 Disabled：禁用。
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitnil,omitempty" name:"SCIMSynchronizationStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetSCIMSynchronizationStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetSCIMSynchronizationStatusResponseParams `json:"Response"`
}

func (r *GetSCIMSynchronizationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSCIMSynchronizationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type GetTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *GetTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusResponseParams struct {
	// 任务状态信息。
	TaskStatus *TaskStatus `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskStatusResponseParams `json:"Response"`
}

func (r *GetTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserRequestParams struct {
	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
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
	delete(f, "UserId")
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserResponseParams struct {
	// 用户信息。
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

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

// Predefined struct for user
type GetUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM 用户同步的 ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

type GetUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// CAM 用户同步的 ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`
}

func (r *GetUserSyncProvisioningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserSyncProvisioningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "UserProvisioningId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUserSyncProvisioningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetUserSyncProvisioningResponseParams struct {
	// CAM 用户同步信息。
	UserProvisioning *UserProvisioning `json:"UserProvisioning,omitnil,omitempty" name:"UserProvisioning"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *GetUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *GetUserSyncProvisioningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetZoneSAMLServiceProviderInfoRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetZoneSAMLServiceProviderInfoRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetZoneSAMLServiceProviderInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetZoneSAMLServiceProviderInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetZoneSAMLServiceProviderInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetZoneSAMLServiceProviderInfoResponseParams struct {
	// saml服务提供商配置信息
	SAMLServiceProvider *SAMLServiceProvider `json:"SAMLServiceProvider,omitnil,omitempty" name:"SAMLServiceProvider"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetZoneSAMLServiceProviderInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetZoneSAMLServiceProviderInfoResponseParams `json:"Response"`
}

func (r *GetZoneSAMLServiceProviderInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetZoneSAMLServiceProviderInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetZoneStatisticsRequestParams struct {
	// 空间ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type GetZoneStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *GetZoneStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetZoneStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetZoneStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetZoneStatisticsResponseParams struct {
	// 空间的统计信息。
	ZoneStatistics *ZoneStatistics `json:"ZoneStatistics,omitnil,omitempty" name:"ZoneStatistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetZoneStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetZoneStatisticsResponseParams `json:"Response"`
}

func (r *GetZoneStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetZoneStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 用户组的名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户组的类型  Manual：手动创建，Synchronized：外部导入。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 用户组的修改时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 用户组的 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 组员数量。
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// 如果有入参FilterUsers，用户在用户组返回true，否则返回false
	IsSelected *bool `json:"IsSelected,omitnil,omitempty" name:"IsSelected"`
}

type GroupMembers struct {
	// 查询username。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的显示名称。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户状态 Enabled：启用， Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户加入用户组的时间
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`
}

type IdentityPolicy struct {
	// CAM预设策略ID。PolicyType 为预设策略时有效且必选
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// CAM预设策略名称。PolicyType 为预设策略时有效且必选
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略类型。取值 1-自定义策略  2-预设策略；默认值2
	PolicyType *uint64 `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 自定义策略内容，遵循CAM策略语法。PolicyType 为自定义策略时有效且必选
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`
}

// Predefined struct for user
type InviteOrganizationMemberRequestParams struct {
	// 被邀请账号Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨（若需要开启资金划拨权限，请联系您的商务经理内部开通。）、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析、9-预算管理、10-信用额度设置（若需要开启信用额度设置权限，请联系您的商务经理内部开通。），1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否允许成员退出。允许：Allow，不允许：Denied。
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 互信实名主体名称。
	RelationAuthName *string `json:"RelationAuthName,omitnil,omitempty" name:"RelationAuthName"`

	// 互信主体证明文件列表。
	AuthFile []*AuthRelationFile `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 成员标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type InviteOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// 被邀请账号Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 关系策略。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。取值：1-查看账单、2-查看余额、3-资金划拨（若需要开启资金划拨权限，请联系您的商务经理内部开通。）、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析、9-预算管理、10-信用额度设置（若需要开启信用额度设置权限，请联系您的商务经理内部开通。），1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 成员所属部门的节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否允许成员退出。允许：Allow，不允许：Denied。
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。成员代付费时需要
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 互信实名主体名称。
	RelationAuthName *string `json:"RelationAuthName,omitnil,omitempty" name:"RelationAuthName"`

	// 互信主体证明文件列表。
	AuthFile []*AuthRelationFile `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 成员标签列表。最大10个
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *InviteOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InviteOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "Name")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "NodeId")
	delete(f, "Remark")
	delete(f, "IsAllowQuit")
	delete(f, "PayUin")
	delete(f, "RelationAuthName")
	delete(f, "AuthFile")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InviteOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InviteOrganizationMemberResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InviteOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *InviteOrganizationMemberResponseParams `json:"Response"`
}

func (r *InviteOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InviteOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JoinedGroups struct {
	// 用户组的名称。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 用户组的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户组 ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户组的类型。取值：
	// 
	// Manual：手动创建。
	// Synchronized：外部同步。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 加入用户组的时间
	JoinTime *string `json:"JoinTime,omitnil,omitempty" name:"JoinTime"`
}

// Predefined struct for user
type ListExternalSAMLIdPCertificatesRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type ListExternalSAMLIdPCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

func (r *ListExternalSAMLIdPCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListExternalSAMLIdPCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListExternalSAMLIdPCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListExternalSAMLIdPCertificatesResponseParams struct {
	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// SAML 签名证书列表
	SAMLIdPCertificates []*SAMLIdPCertificate `json:"SAMLIdPCertificates,omitnil,omitempty" name:"SAMLIdPCertificates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListExternalSAMLIdPCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *ListExternalSAMLIdPCertificatesResponseParams `json:"Response"`
}

func (r *ListExternalSAMLIdPCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListExternalSAMLIdPCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupMembersRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type ListGroupMembersRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

func (r *ListGroupMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGroupMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupMembersResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 用户组的用户列表
	GroupMembers []*GroupMembers `json:"GroupMembers,omitnil,omitempty" name:"GroupMembers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListGroupMembersResponse struct {
	*tchttp.BaseResponse
	Response *ListGroupMembersResponseParams `json:"Response"`
}

func (r *ListGroupMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListGroupMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 过滤条件。  格式：<Attribute> <Operator> <Value>，不区分大小写。目前，<Attribute>只支持GroupName，<Operator>只支持eq（Equals）和sw（Start With）。  示例：Filter = "GroupName sw test"，表示查询名称以 test 开头的全部用户组。Filter = "GroupName eq testgroup"，表示查询名称为 testgroup 的用户组。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 用户组的类型  Manual：手动创建，Synchronized：外部导入。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 筛选的用户，该用户关联的用户组会返回IsSelected=1
	FilterUsers []*string `json:"FilterUsers,omitnil,omitempty" name:"FilterUsers"`

	// 排序的字段，目前只支持CreateTime，默认是CreateTime字段
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型：Desc 倒序 Asc  正序，需要您和SortField一起设置
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 翻页offset. 不要与NextToken同时使用，优先使用NextToken
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 过滤条件。  格式：<Attribute> <Operator> <Value>，不区分大小写。目前，<Attribute>只支持GroupName，<Operator>只支持eq（Equals）和sw（Start With）。  示例：Filter = "GroupName sw test"，表示查询名称以 test 开头的全部用户组。Filter = "GroupName eq testgroup"，表示查询名称为 testgroup 的用户组。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 用户组的类型  Manual：手动创建，Synchronized：外部导入。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 筛选的用户，该用户关联的用户组会返回IsSelected=1
	FilterUsers []*string `json:"FilterUsers,omitnil,omitempty" name:"FilterUsers"`

	// 排序的字段，目前只支持CreateTime，默认是CreateTime字段
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型：Desc 倒序 Asc  正序，需要您和SortField一起设置
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 翻页offset. 不要与NextToken同时使用，优先使用NextToken
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "ZoneId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "Filter")
	delete(f, "GroupType")
	delete(f, "FilterUsers")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListGroupsResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 用户组列表。
	Groups []*GroupInfo `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

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
type ListJoinedGroupsForUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

type ListJoinedGroupsForUserRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

func (r *ListJoinedGroupsForUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJoinedGroupsForUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "UserId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListJoinedGroupsForUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListJoinedGroupsForUserResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 用户加入的用户组列表
	JoinedGroups []*JoinedGroups `json:"JoinedGroups,omitnil,omitempty" name:"JoinedGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListJoinedGroupsForUserResponse struct {
	*tchttp.BaseResponse
	Response *ListJoinedGroupsForUserResponseParams `json:"Response"`
}

func (r *ListJoinedGroupsForUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListJoinedGroupsForUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListNonCompliantResourceRequestParams struct {
	// 限制数目。取值范围：1~50。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitnil,omitempty" name:"PaginationToken"`

	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type ListNonCompliantResourceRequest struct {
	*tchttp.BaseRequest
	
	// 限制数目。取值范围：1~50。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 从上一页的响应中获取的下一页的Token值。
	// 如果是第一次请求，设置为空。
	PaginationToken *string `json:"PaginationToken,omitnil,omitempty" name:"PaginationToken"`

	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

func (r *ListNonCompliantResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListNonCompliantResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxResults")
	delete(f, "MemberUin")
	delete(f, "PaginationToken")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListNonCompliantResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListNonCompliantResourceResponseParams struct {
	// 资源及标签合规信息。
	Items []*ResourceTagMapping `json:"Items,omitnil,omitempty" name:"Items"`

	// 获取的下一页的Token值。
	PaginationToken *string `json:"PaginationToken,omitnil,omitempty" name:"PaginationToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListNonCompliantResourceResponse struct {
	*tchttp.BaseResponse
	Response *ListNonCompliantResourceResponseParams `json:"Response"`
}

func (r *ListNonCompliantResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListNonCompliantResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrgServiceAssignMemberRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集团服务ID。和集团服务产品标识二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品标识。和集团服务ID二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type ListOrgServiceAssignMemberRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集团服务ID。和集团服务产品标识二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品标识。和集团服务ID二选一必填，可以通过[ListOrganizationService](https://cloud.tencent.com/document/product/850/109561)获取
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *ListOrgServiceAssignMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrgServiceAssignMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ServiceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrgServiceAssignMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrgServiceAssignMemberResponseParams struct {
	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 委派管理员列表。
	Items []*OrganizationServiceAssignMember `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrgServiceAssignMemberResponse struct {
	*tchttp.BaseResponse
	Response *ListOrgServiceAssignMemberResponseParams `json:"Response"`
}

func (r *ListOrgServiceAssignMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrgServiceAssignMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityRequestParams struct {
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 身份ID。可以通过身份ID搜索
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份类型。取值范围 1-预设, 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

type ListOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍。默认值 : 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50。默认值：10。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 身份ID。可以通过身份ID搜索
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份类型。取值范围 1-预设, 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`
}

func (r *ListOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	delete(f, "IdentityId")
	delete(f, "IdentityType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationIdentityResponseParams struct {
	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 条目详情。
	Items []*OrgIdentity `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationIdentityResponseParams `json:"Response"`
}

func (r *ListOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationServiceRequestParams struct {
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type ListOrganizationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。取值是limit的整数倍，默认值 : 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目。取值范围：1~50，默认值：10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 名称搜索关键字。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *ListOrganizationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListOrganizationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListOrganizationServiceResponseParams struct {
	// 总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 集团服务列表。
	Items []*OrganizationServiceAssign `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListOrganizationServiceResponse struct {
	*tchttp.BaseResponse
	Response *ListOrganizationServiceResponseParams `json:"Response"`
}

func (r *ListOrganizationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListOrganizationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPermissionPoliciesInRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 按策略名称搜索
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type ListPermissionPoliciesInRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 按策略名称搜索
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *ListPermissionPoliciesInRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPermissionPoliciesInRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "RolePolicyType")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPermissionPoliciesInRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPermissionPoliciesInRoleConfigurationResponseParams struct {
	// 权限策略总个数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 权限策略列表。
	RolePolicies []*RolePolicie `json:"RolePolicies,omitnil,omitempty" name:"RolePolicies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPermissionPoliciesInRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ListPermissionPoliciesInRoleConfigurationResponseParams `json:"Response"`
}

func (r *ListPermissionPoliciesInRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPermissionPoliciesInRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListPoliciesForTarget struct {
	// 策略Id
	StrategyId *uint64 `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 关联的账号或节点
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 关联类型 1-节点 2-用户
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 策略创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 部门名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略绑定时间
	AttachTime *string `json:"AttachTime,omitnil,omitempty" name:"AttachTime"`
}

// Predefined struct for user
type ListPoliciesForTargetRequestParams struct {
	// 账号uin或者节点id。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 搜索关键字。按照策略名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListPoliciesForTargetRequest struct {
	*tchttp.BaseRequest
	
	// 账号uin或者节点id。
	TargetId *uint64 `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 搜索关键字。按照策略名称搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListPoliciesForTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesForTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetId")
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "PolicyType")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPoliciesForTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesForTargetResponseParams struct {
	// 总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 目标关联的策略列表。
	List []*ListPoliciesForTarget `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPoliciesForTargetResponse struct {
	*tchttp.BaseResponse
	Response *ListPoliciesForTargetResponseParams `json:"Response"`
}

func (r *ListPoliciesForTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPoliciesForTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesRequestParams struct {
	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询范围。取值范围： All-获取所有策略、QCS-只获取预设策略、Local-只获取自定义策略，默认值：All
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 搜索关键字。按照策略名搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
}

type ListPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 查询范围。取值范围： All-获取所有策略、QCS-只获取预设策略、Local-只获取自定义策略，默认值：All
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 搜索关键字。按照策略名搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`
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
	delete(f, "PolicyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPoliciesResponseParams struct {
	// 策略总数
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 策略列表数据
	List []*ListPolicyNode `json:"List,omitnil,omitempty" name:"List"`

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

type ListPolicyNode struct {
	// 策略创建时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 策略绑定次数
	AttachedTimes *uint64 `json:"AttachedTimes,omitnil,omitempty" name:"AttachedTimes"`

	// 策略描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略Id
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 策略类型 1-自定义 2-预设
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ListRoleAssignmentsRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****)，当PrincipalType取值为User时，该值为用户 ID （u-****）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询条件，目前只支持权限配置名称查询。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type ListRoleAssignmentsRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****)，当PrincipalType取值为User时，该值为用户 ID （u-****）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询条件，目前只支持权限配置名称查询。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *ListRoleAssignmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleAssignmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "TargetType")
	delete(f, "TargetUin")
	delete(f, "PrincipalType")
	delete(f, "PrincipalId")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRoleAssignmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRoleAssignmentsResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 集团账号目标账号的授权列表。
	RoleAssignments []*RoleAssignments `json:"RoleAssignments,omitnil,omitempty" name:"RoleAssignments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRoleAssignmentsResponse struct {
	*tchttp.BaseResponse
	Response *ListRoleAssignmentsResponseParams `json:"Response"`
}

func (r *ListRoleAssignmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleAssignmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRoleConfigurationProvisioningsRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// Deployed: 部署成功 DeployedRequired：需要重新部署 DeployFailed：部署失败
	DeploymentStatus *string `json:"DeploymentStatus,omitnil,omitempty" name:"DeploymentStatus"`

	// 支持配置名称搜索。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type ListRoleConfigurationProvisioningsRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// Deployed: 部署成功 DeployedRequired：需要重新部署 DeployFailed：部署失败
	DeploymentStatus *string `json:"DeploymentStatus,omitnil,omitempty" name:"DeploymentStatus"`

	// 支持配置名称搜索。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *ListRoleConfigurationProvisioningsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleConfigurationProvisioningsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "TargetType")
	delete(f, "TargetUin")
	delete(f, "DeploymentStatus")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRoleConfigurationProvisioningsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRoleConfigurationProvisioningsResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 部成员账号列表。
	RoleConfigurationProvisionings []*RoleConfigurationProvisionings `json:"RoleConfigurationProvisionings,omitnil,omitempty" name:"RoleConfigurationProvisionings"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRoleConfigurationProvisioningsResponse struct {
	*tchttp.BaseResponse
	Response *ListRoleConfigurationProvisioningsResponseParams `json:"Response"`
}

func (r *ListRoleConfigurationProvisioningsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleConfigurationProvisioningsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRoleConfigurationsRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 过滤文本。不区分大小写。目前，支持 RoleConfigurationName和Description. 示例：Filter = "test"，表示查询名称或描述里包含 test 的权限配置。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 检索成员账号是否配置过权限，如果配置过返回IsSelected: true, 否则返回false。
	FilterTargets []*int64 `json:"FilterTargets,omitnil,omitempty" name:"FilterTargets"`

	// 授权的用户UserId或者用户组的GroupId，必须和入参数FilterTargets一起设置
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`
}

type ListRoleConfigurationsRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 过滤文本。不区分大小写。目前，支持 RoleConfigurationName和Description. 示例：Filter = "test"，表示查询名称或描述里包含 test 的权限配置。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 检索成员账号是否配置过权限，如果配置过返回IsSelected: true, 否则返回false。
	FilterTargets []*int64 `json:"FilterTargets,omitnil,omitempty" name:"FilterTargets"`

	// 授权的用户UserId或者用户组的GroupId，必须和入参数FilterTargets一起设置
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`
}

func (r *ListRoleConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleConfigurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "Filter")
	delete(f, "FilterTargets")
	delete(f, "PrincipalId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRoleConfigurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRoleConfigurationsResponseParams struct {
	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 权限配置列表。
	RoleConfigurations []*RoleConfiguration `json:"RoleConfigurations,omitnil,omitempty" name:"RoleConfigurations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRoleConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *ListRoleConfigurationsResponseParams `json:"Response"`
}

func (r *ListRoleConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRoleConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSCIMCredentialsRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

type ListSCIMCredentialsRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`
}

func (r *ListSCIMCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSCIMCredentialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CredentialId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSCIMCredentialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSCIMCredentialsResponseParams struct {
	// SCIM密钥数量。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// SCIM 密钥信息。
	SCIMCredentials []*SCIMCredential `json:"SCIMCredentials,omitnil,omitempty" name:"SCIMCredentials"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSCIMCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *ListSCIMCredentialsResponseParams `json:"Response"`
}

func (r *ListSCIMCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSCIMCredentialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListTargetsForPolicyNode struct {
	// scp账号uin或节点Id
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 关联类型 1-节点关联 2-用户关联
	RelatedType *uint64 `json:"RelatedType,omitnil,omitempty" name:"RelatedType"`

	// 账号或者节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 绑定时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`
}

// Predefined struct for user
type ListTargetsForPolicyRequestParams struct {
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。取值范围：All-全部、User-用户、Node-节点
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 按照多个策略id搜索，空格隔开。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListTargetsForPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id。
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 每页数量。默认值是 20，必须大于 0 且小于或等于 200
	Rp *uint64 `json:"Rp,omitnil,omitempty" name:"Rp"`

	// 页码。默认值是 1，从 1开始，不能大于 200
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 策略类型。取值范围：All-全部、User-用户、Node-节点
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 按照多个策略id搜索，空格隔开。
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListTargetsForPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsForPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Rp")
	delete(f, "Page")
	delete(f, "TargetType")
	delete(f, "PolicyType")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTargetsForPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTargetsForPolicyResponseParams struct {
	// 总数。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 指定SCP策略关联目标列表。
	List []*ListTargetsForPolicyNode `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTargetsForPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListTargetsForPolicyResponseParams `json:"Response"`
}

func (r *ListTargetsForPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTargetsForPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****）， 当PrincipalType取值为User时，该值为用户 ID（u-****）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type ListTasksRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步 ID。取值： 当PrincipalType取值为Group时，该值为用户组 ID（g-****）， 当PrincipalType取值为User时，该值为用户 ID（u-****）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

func (r *ListTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PrincipalId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "PrincipalType")
	delete(f, "TargetUin")
	delete(f, "TargetType")
	delete(f, "RoleConfigurationId")
	delete(f, "Status")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTasksResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// 任务详情
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListTasksResponseParams `json:"Response"`
}

func (r *ListTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserSyncProvisioningsRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步 ID。取值：  当PrincipalType取值为Group时，该值为用户组 ID（g-********）。 当PrincipalType取值为User时，该值为用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 检测条件。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type ListUserSyncProvisioningsRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步 ID。取值：  当PrincipalType取值为Group时，该值为用户组 ID（g-********）。 当PrincipalType取值为User时，该值为用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法多次查询，直到IsTruncated为false，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// CAM 用户同步的身份类型。取值： User：表示同步的身份是用户。 Group：表示同步的身份是用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 检测条件。
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *ListUserSyncProvisioningsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserSyncProvisioningsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PrincipalId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "PrincipalType")
	delete(f, "TargetUin")
	delete(f, "TargetType")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUserSyncProvisioningsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUserSyncProvisioningsResponseParams struct {
	// 查询返回结果下一页的令牌。  说明 只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

	// CAM同步的用户列表。
	UserProvisionings []*UserProvisioning `json:"UserProvisionings,omitnil,omitempty" name:"UserProvisionings"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUserSyncProvisioningsResponse struct {
	*tchttp.BaseResponse
	Response *ListUserSyncProvisioningsResponseParams `json:"Response"`
}

func (r *ListUserSyncProvisioningsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUserSyncProvisioningsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户状态 Enabled：启用， Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 过滤条件。  目前仅支持用户名，邮箱，用户userId，描述
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法经过多次查询，直到IsTruncated为false时，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 筛选的用户组，该用户组关联的子用户会返回IsSelected=1
	FilterGroups []*string `json:"FilterGroups,omitnil,omitempty" name:"FilterGroups"`

	// 排序的字段，目前只支持CreateTime，默认是CreateTime字段
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型：Desc 倒序 Asc  正序，需要您和SortField一起设置
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 翻页offset. 不要与NextToken同时使用，优先使用NextToken
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListUsersRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户状态 Enabled：启用， Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 过滤条件。  目前仅支持用户名，邮箱，用户userId，描述
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 每页的最大数据条数。  取值范围：1~100。  默认值：10。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 查询返回结果下一页的令牌。首次调用 API 不需要NextToken。  当您首次调用 API 时，如果返回数据总条数超过MaxResults限制，数据会被截断，只返回MaxResults条数据，同时，返回参数IsTruncated为true，返回一个NextToken。您可以使用上一次返回的NextToken继续调用 API，其他请求参数保持不变，查询被截断的数据。您可以按此方法经过多次查询，直到IsTruncated为false时，表示全部数据查询完毕。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 筛选的用户组，该用户组关联的子用户会返回IsSelected=1
	FilterGroups []*string `json:"FilterGroups,omitnil,omitempty" name:"FilterGroups"`

	// 排序的字段，目前只支持CreateTime，默认是CreateTime字段
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 排序类型：Desc 倒序 Asc  正序，需要您和SortField一起设置
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 翻页offset. 不要与NextToken同时使用，优先使用NextToken
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "ZoneId")
	delete(f, "UserStatus")
	delete(f, "UserType")
	delete(f, "Filter")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "FilterGroups")
	delete(f, "SortField")
	delete(f, "SortType")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsersResponseParams struct {
	// 符合请求参数条件的数据总条数。
	TotalCounts *int64 `json:"TotalCounts,omitnil,omitempty" name:"TotalCounts"`

	// 每页的最大数据条数。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 用户列表。
	Users []*UserInfo `json:"Users,omitnil,omitempty" name:"Users"`

	// 查询返回结果下一页的令牌。只有IsTruncated为true时，才显示该参数。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 返回结果是否被截断。取值：  true：已截断。 false：未截断。
	IsTruncated *bool `json:"IsTruncated,omitnil,omitempty" name:"IsTruncated"`

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

type ManagerShareUnit struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元管理员Uin。
	Uin *int64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 共享单元管理员OwnerUin。
	OwnerUin *int64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 共享单元地域。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 共享单元资源数。
	ShareResourceNum *int64 `json:"ShareResourceNum,omitnil,omitempty" name:"ShareResourceNum"`

	// 共享单元成员数。
	ShareMemberNum *int64 `json:"ShareMemberNum,omitnil,omitempty" name:"ShareMemberNum"`

	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号
	ShareScope *uint64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`
}

type MemberIdentity struct {
	// 身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份名称。
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`
}

type MemberMainInfo struct {
	// 成员uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

// Predefined struct for user
type MoveOrganizationNodeMembersRequestParams struct {
	// 组织节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 成员Uin列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

type MoveOrganizationNodeMembersRequest struct {
	*tchttp.BaseRequest
	
	// 组织节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 成员Uin列表。
	MemberUin []*int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`
}

func (r *MoveOrganizationNodeMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "MemberUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveOrganizationNodeMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveOrganizationNodeMembersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MoveOrganizationNodeMembersResponse struct {
	*tchttp.BaseResponse
	Response *MoveOrganizationNodeMembersResponseParams `json:"Response"`
}

func (r *MoveOrganizationNodeMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveOrganizationNodeMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeMainInfo struct {
	// 部门ID
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 部门名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type NotAllowReason struct {
	// 是否创建的成员。true-是、false-否；成员不是创建的成员不允许删除
	IsCreateMember *bool `json:"IsCreateMember,omitnil,omitempty" name:"IsCreateMember"`

	// 成员删除许可。true-开启、false-关闭；成员删除许可关闭时不允许删除
	DeletionPermission *bool `json:"DeletionPermission,omitnil,omitempty" name:"DeletionPermission"`

	// 是否可信服务委派管理员。true-是、false-否；成员是可信服务委派管理员不允许删除
	IsAssignManager *bool `json:"IsAssignManager,omitnil,omitempty" name:"IsAssignManager"`

	// 是否主体管理员。true-是、false-否；成员是主体管理员不允许删除
	IsAuthManager *bool `json:"IsAuthManager,omitnil,omitempty" name:"IsAuthManager"`

	// 是否共享资源管理员。true-是、false-否；成员是共享资源管理员不允许删除
	IsShareManager *bool `json:"IsShareManager,omitnil,omitempty" name:"IsShareManager"`

	// 成员是否设置了操作审批。true-是、false-否；成员设置了操作审批时不允许删除
	OperateProcess *bool `json:"OperateProcess,omitnil,omitempty" name:"OperateProcess"`

	// 是否允许解除成员财务权限。true-是、false-否；成员不能解除财务权限时不允许删除
	BillingPermission *bool `json:"BillingPermission,omitnil,omitempty" name:"BillingPermission"`

	// 存在的资源列表。账号存在资源时不允许删除
	ExistResources []*string `json:"ExistResources,omitnil,omitempty" name:"ExistResources"`

	// 检测失败的资源列表。账号有资源检测失败时不允许删除。
	DetectFailedResources []*string `json:"DetectFailedResources,omitnil,omitempty" name:"DetectFailedResources"`
}

// Predefined struct for user
type OpenIdentityCenterRequestParams struct {
	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type OpenIdentityCenterRequest struct {
	*tchttp.BaseRequest
	
	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

func (r *OpenIdentityCenterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenIdentityCenterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenIdentityCenterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenIdentityCenterResponseParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenIdentityCenterResponse struct {
	*tchttp.BaseResponse
	Response *OpenIdentityCenterResponseParams `json:"Response"`
}

func (r *OpenIdentityCenterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenIdentityCenterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrgFinancialByMonth struct {
	// 记录ID。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 月份，格式：yyyy-mm，示例：2021-01。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 消耗金额，单元：元。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 比上月增长率%。正数增长，负数下降，空值无法统计。
	GrowthRate *string `json:"GrowthRate,omitnil,omitempty" name:"GrowthRate"`
}

type OrgIdentity struct {
	// 身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份名称。
	IdentityAliasName *string `json:"IdentityAliasName,omitnil,omitempty" name:"IdentityAliasName"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`

	// 身份类型。 1-预设、 2-自定义
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgMember struct {
	// 成员Uin
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 成员类型，邀请：Invite， 创建：Create
	MemberType *string `json:"MemberType,omitnil,omitempty" name:"MemberType"`

	// 关系策略类型
	OrgPolicyType *string `json:"OrgPolicyType,omitnil,omitempty" name:"OrgPolicyType"`

	// 关系策略名
	OrgPolicyName *string `json:"OrgPolicyName,omitnil,omitempty" name:"OrgPolicyName"`

	// 关系策略权限
	OrgPermission []*OrgPermission `json:"OrgPermission,omitnil,omitempty" name:"OrgPermission"`

	// 所属节点ID
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 所属节点名
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否允许成员退出。允许：Allow，不允许：Denied。
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 代付者名称
	PayName *string `json:"PayName,omitnil,omitempty" name:"PayName"`

	// 管理身份
	OrgIdentity []*MemberIdentity `json:"OrgIdentity,omitnil,omitempty" name:"OrgIdentity"`

	// 安全信息绑定状态  未绑定：Unbound，待激活：Valid，绑定成功：Success，绑定失败：Failed
	BindStatus *string `json:"BindStatus,omitnil,omitempty" name:"BindStatus"`

	// 成员权限状态 已确认：Confirmed ，待确认：UnConfirmed
	PermissionStatus *string `json:"PermissionStatus,omitnil,omitempty" name:"PermissionStatus"`

	// 成员标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 腾讯云昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`
}

type OrgMemberAuthAccount struct {
	// 组织子账号Uin。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份角色名。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份角色别名。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 子账号名称
	OrgSubAccountName *string `json:"OrgSubAccountName,omitnil,omitempty" name:"OrgSubAccountName"`
}

type OrgMemberAuthIdentity struct {
	// 身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份的角色名。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份的角色别名。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 身份描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 首次配置成功的时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次配置成功的时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 身份类型。取值： 1-预设身份  2-自定义身份
	IdentityType *uint64 `json:"IdentityType,omitnil,omitempty" name:"IdentityType"`

	// 配置状态。取值：1-配置完成 2-需重新配置
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

type OrgMemberFinancial struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`

	// 消耗金额，单位：元。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 占比%。
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type OrgMemberPolicy struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 身份ID。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份角色名。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份角色别名。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type OrgMembersAuthPolicy struct {
	// 身份Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份的角色名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleName *string `json:"IdentityRoleName,omitnil,omitempty" name:"IdentityRoleName"`

	// 身份的角色别名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentityRoleAliasName *string `json:"IdentityRoleAliasName,omitnil,omitempty" name:"IdentityRoleAliasName"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 成员访问策略Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 成员访问策略名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 成员uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`

	// 子账号uin或者用户组Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgSubAccountUin *int64 `json:"OrgSubAccountUin,omitnil,omitempty" name:"OrgSubAccountUin"`

	// 子账号名称或者用户组名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgSubAccountName *string `json:"OrgSubAccountName,omitnil,omitempty" name:"OrgSubAccountName"`

	// 绑定类型。1-子账号、2-用户组
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindType *uint64 `json:"BindType,omitnil,omitempty" name:"BindType"`

	// 成员信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Members []*MemberMainInfo `json:"Members,omitnil,omitempty" name:"Members"`
}

type OrgNode struct {
	// 组织节点ID
	NodeId *int64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 父节点ID
	ParentNodeId *int64 `json:"ParentNodeId,omitnil,omitempty" name:"ParentNodeId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 成员标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type OrgPermission struct {
	// 权限Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 权限名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type OrgProductFinancial struct {
	// 产品Code。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 产品名。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 产品消耗，单位：元。
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 占比%。
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type OrganizationServiceAssign struct {
	// 集团服务ID。
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品名称。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 是否支持委派。取值: 1-是  2-否
	IsAssign *uint64 `json:"IsAssign,omitnil,omitempty" name:"IsAssign"`

	// 集团服务描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 当前委派管理员数。
	MemberNum *string `json:"MemberNum,omitnil,omitempty" name:"MemberNum"`

	// 帮助文档。
	Document *string `json:"Document,omitnil,omitempty" name:"Document"`

	// 集团服务产品控制台路径。
	ConsoleUrl *string `json:"ConsoleUrl,omitnil,omitempty" name:"ConsoleUrl"`

	// 是否接入使用状态。取值: 1-是 
	//  2-否
	IsUsageStatus *uint64 `json:"IsUsageStatus,omitnil,omitempty" name:"IsUsageStatus"`

	// 委派管理员数量限制。
	CanAssignCount *uint64 `json:"CanAssignCount,omitnil,omitempty" name:"CanAssignCount"`

	// 集团服务产品标识。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 是否支持集团服务授权。取值 1-是、2-否
	ServiceGrant *uint64 `json:"ServiceGrant,omitnil,omitempty" name:"ServiceGrant"`

	// 集团服务授权启用状态。ServiceGrant值为1时该字段有效 ，取值：Enabled-开启  Disabled-关闭 
	GrantStatus *string `json:"GrantStatus,omitnil,omitempty" name:"GrantStatus"`

	// 是否支持设置委派管理范围。取值: 1-是  2-否
	IsSetManagementScope *uint64 `json:"IsSetManagementScope,omitnil,omitempty" name:"IsSetManagementScope"`
}

type OrganizationServiceAssignMember struct {
	// 集团服务ID。
	ServiceId *uint64 `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 集团服务产品名称。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 委派管理员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 委派管理员名称。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`

	// 启用状态 。取值：0-服务无启用状态  1-已启用  2-未启用
	UsageStatus *uint64 `json:"UsageStatus,omitnil,omitempty" name:"UsageStatus"`

	// 委派时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 委派管理员管理范围。取值: 1-全部成员  2-部分成员
	ManagementScope *uint64 `json:"ManagementScope,omitnil,omitempty" name:"ManagementScope"`

	// 管理的成员Uin列表。ManagementScope值为2时该参数有效
	ManagementScopeMembers []*MemberMainInfo `json:"ManagementScopeMembers,omitnil,omitempty" name:"ManagementScopeMembers"`

	// 管理的部门ID列表。ManagementScope值为2时该参数有效
	ManagementScopeNodes []*NodeMainInfo `json:"ManagementScopeNodes,omitnil,omitempty" name:"ManagementScopeNodes"`
}

type PolicyDetail struct {
	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略名称。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`
}

type ProductResource struct {
	// 产品资源ID。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`

	// 资源六段式最后一节
	//
	// Deprecated: ResourceGrantLast is deprecated.
	ResourceGrantLast *string `json:"ResourceGrantLast,omitnil,omitempty" name:"ResourceGrantLast"`
}

// Predefined struct for user
type ProvisionRoleConfigurationRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

type ProvisionRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`
}

func (r *ProvisionRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProvisionRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "TargetType")
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ProvisionRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ProvisionRoleConfigurationResponseParams struct {
	// 任务详情。
	Task *RoleProvisioningsTask `json:"Task,omitnil,omitempty" name:"Task"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ProvisionRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ProvisionRoleConfigurationResponseParams `json:"Response"`
}

func (r *ProvisionRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ProvisionRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuitOrganizationRequestParams struct {
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

type QuitOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 企业组织ID
	OrgId *uint64 `json:"OrgId,omitnil,omitempty" name:"OrgId"`
}

func (r *QuitOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QuitOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QuitOrganizationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QuitOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *QuitOrganizationResponseParams `json:"Response"`
}

func (r *QuitOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QuitOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectJoinShareUnitInvitationRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

type RejectJoinShareUnitInvitationRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`
}

func (r *RejectJoinShareUnitInvitationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectJoinShareUnitInvitationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectJoinShareUnitInvitationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectJoinShareUnitInvitationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RejectJoinShareUnitInvitationResponse struct {
	*tchttp.BaseResponse
	Response *RejectJoinShareUnitInvitationResponseParams `json:"Response"`
}

func (r *RejectJoinShareUnitInvitationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectJoinShareUnitInvitationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveExternalSAMLIdPCertificateRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// PEM 格式的 X509 证书。  由 SAML 身份提供商提供。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

type RemoveExternalSAMLIdPCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// PEM 格式的 X509 证书。  由 SAML 身份提供商提供。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`
}

func (r *RemoveExternalSAMLIdPCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveExternalSAMLIdPCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CertificateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveExternalSAMLIdPCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveExternalSAMLIdPCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveExternalSAMLIdPCertificateResponse struct {
	*tchttp.BaseResponse
	Response *RemoveExternalSAMLIdPCertificateResponseParams `json:"Response"`
}

func (r *RemoveExternalSAMLIdPCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveExternalSAMLIdPCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePermissionPolicyFromRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 权限策略名称，长度最大为 32 个字符。
	RolePolicyName *string `json:"RolePolicyName,omitnil,omitempty" name:"RolePolicyName"`

	// 策略ID。
	RolePolicyId *int64 `json:"RolePolicyId,omitnil,omitempty" name:"RolePolicyId"`
}

type RemovePermissionPolicyFromRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略类型。取值：  System：系统策略。复用 CAM 的系统策略。 Custom: 自定义策略。按照 CAM 权限策略语法和结构编写的自定义策略。
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 权限策略名称，长度最大为 32 个字符。
	RolePolicyName *string `json:"RolePolicyName,omitnil,omitempty" name:"RolePolicyName"`

	// 策略ID。
	RolePolicyId *int64 `json:"RolePolicyId,omitnil,omitempty" name:"RolePolicyId"`
}

func (r *RemovePermissionPolicyFromRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePermissionPolicyFromRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "RolePolicyType")
	delete(f, "RolePolicyName")
	delete(f, "RolePolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemovePermissionPolicyFromRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePermissionPolicyFromRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemovePermissionPolicyFromRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *RemovePermissionPolicyFromRoleConfigurationResponseParams `json:"Response"`
}

func (r *RemovePermissionPolicyFromRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePermissionPolicyFromRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveUserFromGroupRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RemoveUserFromGroupRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 用户ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
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
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "UserId")
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

type ResourceTagMapping struct {
	// 资源六段式。腾讯云使用资源六段式描述一个资源。
	// 例如：qcs::${ServiceType}:${Region}:${Account}:${ResourcePreifx}/${ResourceId}。
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 合规详情。
	ComplianceDetails *TagComplianceDetails `json:"ComplianceDetails,omitnil,omitempty" name:"ComplianceDetails"`

	// 资源标签。
	Tags []*Tags `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RoleAssignmentInfo struct {
	// CAM 用户同步的身份 ID。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。
	// 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值：
	// 
	// User：表示该 CAM 用户同步的身份是CIC用户。
	// Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 同步集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`
}

type RoleAssignments struct {
	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CAM 用户同步的身份 ID。取值： 当PrincipalType取值为Group时，该值为CIC 用户组 ID（g-********）。 当PrincipalType取值为User时，该值为CIC 用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值： User：表示该 CAM 用户同步的身份是CIC用户。 Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户名称或者用户组名称
	PrincipalName *string `json:"PrincipalName,omitnil,omitempty" name:"PrincipalName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 集团账号目标账号的名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`
}

type RoleConfiguration struct {
	// 权限配置配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置配名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 权限配置的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 会话持续时间。CIC 用户使用访问配置访问成员账号时，会话最多保持的时间。
	// 单位：秒。
	SessionDuration *int64 `json:"SessionDuration,omitnil,omitempty" name:"SessionDuration"`

	// 初始访问页面。CIC 用户使用访问配置访问成员账号时，初始访问的页面地址。
	RelayState *string `json:"RelayState,omitnil,omitempty" name:"RelayState"`

	// 权限配置的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 权限配置的更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 如果有入参FilterTargets查询成员账号是否配置过权限，配置了返回true，否则返回false。
	IsSelected *bool `json:"IsSelected,omitnil,omitempty" name:"IsSelected"`
}

type RoleConfigurationProvisionings struct {
	// Deployed: 部署成功 DeployedRequired：需要重新部署 DeployFailed：部署失败
	DeploymentStatus *string `json:"DeploymentStatus,omitnil,omitempty" name:"DeploymentStatus"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 集团账号目标账号的名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 创建时间，
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间，
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

type RolePolicie struct {
	// 策略ID。
	RolePolicyId *int64 `json:"RolePolicyId,omitnil,omitempty" name:"RolePolicyId"`

	// 权限策略名称
	RolePolicyName *string `json:"RolePolicyName,omitnil,omitempty" name:"RolePolicyName"`

	// 权限策略类型
	RolePolicyType *string `json:"RolePolicyType,omitnil,omitempty" name:"RolePolicyType"`

	// 自定义策略内容。仅自定义策略返回该参数。
	RolePolicyDocument *string `json:"RolePolicyDocument,omitnil,omitempty" name:"RolePolicyDocument"`

	// 权限策略被添加到权限配置的时间。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`
}

type RoleProvisioningsTask struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 授权的集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态：InProgress: 进行中，Failed: 失败 3:Success: 成功
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type SAMLIdPCertificate struct {
	// 证书序列号。
	SerialNumber *string `json:"SerialNumber,omitnil,omitempty" name:"SerialNumber"`

	// 证书颁发者。
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// 证书版本。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 证书ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// PEM 格式的公钥证书（Base64 编码）。
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 证书的签名算法。
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitnil,omitempty" name:"SignatureAlgorithm"`

	// 证书的过期时间。
	NotAfter *string `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// 证书的创建时间。
	NotBefore *string `json:"NotBefore,omitnil,omitempty" name:"NotBefore"`

	// 证书的主体。
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// PEM 格式的 X509 证书。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

type SAMLIdentityProviderConfiguration struct {
	// IdP 标识。
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// SSO 登录的启用状态。取值：  Enabled：启用。 Disabled（默认值）：禁用。
	SSOStatus *string `json:"SSOStatus,omitnil,omitempty" name:"SSOStatus"`

	// IdP 元数据文档（Base64 编码）。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// X509证书ID。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// IdP 的登录地址。
	LoginUrl *string `json:"LoginUrl,omitnil,omitempty" name:"LoginUrl"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SAMLServiceProvider struct {
	// https://tencentcloudsso.com/saml/sp/z-sjw8ensa**
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SP 元数据文档（Base64 编码）。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// SP 的 ACS URL。
	AcsUrl *string `json:"AcsUrl,omitnil,omitempty" name:"AcsUrl"`
}

type SCIMCredential struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥状态，Enabled已开启，Disabled已关闭。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// SCIM密钥类型。
	CredentialType *string `json:"CredentialType,omitnil,omitempty" name:"CredentialType"`

	// SCIM 密钥的创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// SCIM 密钥的过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type SendOrgMemberAccountBindEmailRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`
}

type SendOrgMemberAccountBindEmailRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`
}

func (r *SendOrgMemberAccountBindEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendOrgMemberAccountBindEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "BindId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendOrgMemberAccountBindEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendOrgMemberAccountBindEmailResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendOrgMemberAccountBindEmailResponse struct {
	*tchttp.BaseResponse
	Response *SendOrgMemberAccountBindEmailResponseParams `json:"Response"`
}

func (r *SendOrgMemberAccountBindEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendOrgMemberAccountBindEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetExternalSAMLIdentityProviderRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IdP 元数据文档（Base64 编码）。  由支持 SAML 2.0 协议的 IdP 提供。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// SSO 登录的启用状态。取值：  Enabled：启用。 Disabled（默认值）：禁用。
	SSOStatus *string `json:"SSOStatus,omitnil,omitempty" name:"SSOStatus"`

	// IdP 标识。
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// IdP 的登录地址。
	LoginUrl *string `json:"LoginUrl,omitnil,omitempty" name:"LoginUrl"`

	// PEM 格式的 X509 证书。指定该参数会替换所有已经存在的证书。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

type SetExternalSAMLIdentityProviderRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// IdP 元数据文档（Base64 编码）。  由支持 SAML 2.0 协议的 IdP 提供。
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitnil,omitempty" name:"EncodedMetadataDocument"`

	// SSO 登录的启用状态。取值：  Enabled：启用。 Disabled（默认值）：禁用。
	SSOStatus *string `json:"SSOStatus,omitnil,omitempty" name:"SSOStatus"`

	// IdP 标识。
	EntityId *string `json:"EntityId,omitnil,omitempty" name:"EntityId"`

	// IdP 的登录地址。
	LoginUrl *string `json:"LoginUrl,omitnil,omitempty" name:"LoginUrl"`

	// PEM 格式的 X509 证书。指定该参数会替换所有已经存在的证书。
	X509Certificate *string `json:"X509Certificate,omitnil,omitempty" name:"X509Certificate"`
}

func (r *SetExternalSAMLIdentityProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetExternalSAMLIdentityProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "EncodedMetadataDocument")
	delete(f, "SSOStatus")
	delete(f, "EntityId")
	delete(f, "LoginUrl")
	delete(f, "X509Certificate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetExternalSAMLIdentityProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetExternalSAMLIdentityProviderResponseParams struct {
	// 证书ID。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetExternalSAMLIdentityProviderResponse struct {
	*tchttp.BaseResponse
	Response *SetExternalSAMLIdentityProviderResponseParams `json:"Response"`
}

func (r *SetExternalSAMLIdentityProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetExternalSAMLIdentityProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShareArea struct {
	// 地域名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 地域标识。
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 地域ID。
	AreaId *int64 `json:"AreaId,omitnil,omitempty" name:"AreaId"`
}

type ShareMember struct {
	// 共享成员Uin。
	ShareMemberUin *int64 `json:"ShareMemberUin,omitnil,omitempty" name:"ShareMemberUin"`
}

type ShareResource struct {
	// 共享资源ID。
	//
	// Deprecated: ResourceId is deprecated.
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 产品资源ID。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`
}

type ShareResourceToMember struct {
	// 共享单元资源ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。
	UnitName *string `json:"UnitName,omitnil,omitempty" name:"UnitName"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 业务资源ID。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`

	// 共享账号Uin。
	ShareManagerUin *int64 `json:"ShareManagerUin,omitnil,omitempty" name:"ShareManagerUin"`
}

type ShareUnitMember struct {
	// 共享成员Uin。
	ShareMemberUin *int64 `json:"ShareMemberUin,omitnil,omitempty" name:"ShareMemberUin"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ShareUnitResource struct {
	// 共享资源ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 共享资源类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 产品资源ID。
	ProductResourceId *string `json:"ProductResourceId,omitnil,omitempty" name:"ProductResourceId"`

	// 共享单元成员数。
	SharedMemberNum *uint64 `json:"SharedMemberNum,omitnil,omitempty" name:"SharedMemberNum"`

	// 使用中共享单元成员数。
	SharedMemberUseNum *uint64 `json:"SharedMemberUseNum,omitnil,omitempty" name:"SharedMemberUseNum"`

	// 共享管理员OwnerUin。
	ShareManagerUin *int64 `json:"ShareManagerUin,omitnil,omitempty" name:"ShareManagerUin"`
}

type Tag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagComplianceDetails struct {
	// 合规状态。true-合规，false-不合规
	ComplianceStatus *bool `json:"ComplianceStatus,omitnil,omitempty" name:"ComplianceStatus"`

	// 值不合规的标签键列表。
	KeysWithNonCompliantValues []*string `json:"KeysWithNonCompliantValues,omitnil,omitempty" name:"KeysWithNonCompliantValues"`

	// 键不合规的标签键列表。
	NonCompliantKeys []*string `json:"NonCompliantKeys,omitnil,omitempty" name:"NonCompliantKeys"`
}

type Tags struct {
	// 标签键。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TaskInfo struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 权限配置ID。
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限配置名称。
	RoleConfigurationName *string `json:"RoleConfigurationName,omitnil,omitempty" name:"RoleConfigurationName"`

	// 授权的目标成员账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 用户授权的身份ID,如果是身份类型是CIC用户,则为用户ID; 如果是用户组，则为用户组ID;
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// 用户授权的身份类型, User代表CIC用户, Group代表CIC用户组
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 任务类型。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`
}

type TaskStatus struct {
	// 任务状态。取值：  InProgress：任务执行中。 Success：任务执行成功。 Failed：任务执行失败。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务 ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型。取值：
	// ProvisionRoleConfiguration：部署权限配置。
	// DeprovisionRoleConfiguration：解除权限配置部署。
	// CreateRoleAssignment：在成员 账号上授权。
	// DeleteRoleAssignment：移除 成员 账号上的授权。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务失败原因。
	// 说明
	// 只有Status为Failed，才会显示该参数。
	FailureReason *string `json:"FailureReason,omitnil,omitempty" name:"FailureReason"`
}

// Predefined struct for user
type UpdateCustomPolicyForRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略名称，长度最大为 32 个字符。
	CustomPolicyName *string `json:"CustomPolicyName,omitnil,omitempty" name:"CustomPolicyName"`

	// 自定义策略内容。长度：最大 4096 个字符。当RolePolicyType为Inline时，该参数必须配置。关于权限策略的语法和结构，请参见权限策略语法和结构。
	NewCustomPolicyDocument *string `json:"NewCustomPolicyDocument,omitnil,omitempty" name:"NewCustomPolicyDocument"`
}

type UpdateCustomPolicyForRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 权限策略名称，长度最大为 32 个字符。
	CustomPolicyName *string `json:"CustomPolicyName,omitnil,omitempty" name:"CustomPolicyName"`

	// 自定义策略内容。长度：最大 4096 个字符。当RolePolicyType为Inline时，该参数必须配置。关于权限策略的语法和结构，请参见权限策略语法和结构。
	NewCustomPolicyDocument *string `json:"NewCustomPolicyDocument,omitnil,omitempty" name:"NewCustomPolicyDocument"`
}

func (r *UpdateCustomPolicyForRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomPolicyForRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "CustomPolicyName")
	delete(f, "NewCustomPolicyDocument")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCustomPolicyForRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCustomPolicyForRoleConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCustomPolicyForRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCustomPolicyForRoleConfigurationResponseParams `json:"Response"`
}

func (r *UpdateCustomPolicyForRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCustomPolicyForRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 新的用户组名称。
	NewGroupName *string `json:"NewGroupName,omitnil,omitempty" name:"NewGroupName"`

	// 新的用户组描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`
}

type UpdateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 新的用户组名称。
	NewGroupName *string `json:"NewGroupName,omitnil,omitempty" name:"NewGroupName"`

	// 新的用户组描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`
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
	delete(f, "ZoneId")
	delete(f, "GroupId")
	delete(f, "NewGroupName")
	delete(f, "NewDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateGroupResponseParams struct {
	// 用户组信息。
	GroupInfo *GroupInfo `json:"GroupInfo,omitnil,omitempty" name:"GroupInfo"`

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
type UpdateOrganizationIdentityRequestParams struct {
	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`
}

type UpdateOrganizationIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份ID。可以通过[ListOrganizationIdentity](https://cloud.tencent.com/document/product/850/82934)获取
	IdentityId *uint64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 身份描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 身份策略。
	IdentityPolicy []*IdentityPolicy `json:"IdentityPolicy,omitnil,omitempty" name:"IdentityPolicy"`
}

func (r *UpdateOrganizationIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityId")
	delete(f, "Description")
	delete(f, "IdentityPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationIdentityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationIdentityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationIdentityResponseParams `json:"Response"`
}

func (r *UpdateOrganizationIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberEmailBindRequestParams struct {
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

type UpdateOrganizationMemberEmailBindRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *int64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 绑定ID。可以通过[DescribeOrganizationMemberEmailBind](https://cloud.tencent.com/document/product/850/93332)获取
	BindId *int64 `json:"BindId,omitnil,omitempty" name:"BindId"`

	// 邮箱地址。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 国际区号。
	CountryCode *string `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 手机号。
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`
}

func (r *UpdateOrganizationMemberEmailBindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberEmailBindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "BindId")
	delete(f, "Email")
	delete(f, "CountryCode")
	delete(f, "Phone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationMemberEmailBindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberEmailBindResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationMemberEmailBindResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMemberEmailBindResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMemberEmailBindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberEmailBindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberRequestParams struct {
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。最大长度为40个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 关系策略类型。PolicyType不为空，PermissionIds不能为空。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。PermissionIds不为空，PolicyType不能为空。
	// 取值：1-查看账单、2-查看余额、3-资金划拨（若需要开启资金划拨权限，请联系您的商务经理内部开通。）、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析、9-预算管理、10-信用额度设置（若需要开启信用额度设置权限，请联系您的商务经理内部开通。），1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 是否允许成员退出组织。取值：Allow-允许、Denied-不允许
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。成员财务权限有代付费时需要，取值为成员对应主体的主体管理员Uin
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 是否同步组织成员名称到成员账号昵称。取值： 1-同步 0-不同步
	IsModifyNickName *uint64 `json:"IsModifyNickName,omitnil,omitempty" name:"IsModifyNickName"`
}

type UpdateOrganizationMemberRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin。
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称。最大长度为25个字符，支持英文字母、数字、汉字、符号+@、&._[]-:,
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。最大长度为40个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 关系策略类型。PolicyType不为空，PermissionIds不能为空。取值：Financial
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 成员财务权限ID列表。PermissionIds不为空，PolicyType不能为空。
	// 取值：1-查看账单、2-查看余额、3-资金划拨（若需要开启资金划拨权限，请联系您的商务经理内部开通。）、4-合并出账、5-开票、6-优惠继承、7-代付费、8-成本分析、9-预算管理、10-信用额度设置（若需要开启信用额度设置权限，请联系您的商务经理内部开通。），1、2 默认必须
	PermissionIds []*uint64 `json:"PermissionIds,omitnil,omitempty" name:"PermissionIds"`

	// 是否允许成员退出组织。取值：Allow-允许、Denied-不允许
	IsAllowQuit *string `json:"IsAllowQuit,omitnil,omitempty" name:"IsAllowQuit"`

	// 代付者Uin。成员财务权限有代付费时需要，取值为成员对应主体的主体管理员Uin
	PayUin *string `json:"PayUin,omitnil,omitempty" name:"PayUin"`

	// 是否同步组织成员名称到成员账号昵称。取值： 1-同步 0-不同步
	IsModifyNickName *uint64 `json:"IsModifyNickName,omitnil,omitempty" name:"IsModifyNickName"`
}

func (r *UpdateOrganizationMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUin")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "PolicyType")
	delete(f, "PermissionIds")
	delete(f, "IsAllowQuit")
	delete(f, "PayUin")
	delete(f, "IsModifyNickName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMemberResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationMemberResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMemberResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMembersPolicyRequestParams struct {
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 成员访问策略Id。可通过DescribeOrganizationMemberPolicies获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 成员访问身份ID。可通过ListOrganizationIdentity获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateOrganizationMembersPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 成员Uin列表。最多10个
	MemberUins []*int64 `json:"MemberUins,omitnil,omitempty" name:"MemberUins"`

	// 成员访问策略Id。可通过DescribeOrganizationMemberPolicies获取
	PolicyId *uint64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 成员访问身份ID。可通过ListOrganizationIdentity获取
	IdentityId *int64 `json:"IdentityId,omitnil,omitempty" name:"IdentityId"`

	// 策略描述。最大长度为128个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateOrganizationMembersPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMembersPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MemberUins")
	delete(f, "PolicyId")
	delete(f, "IdentityId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationMembersPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationMembersPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationMembersPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationMembersPolicyResponseParams `json:"Response"`
}

func (r *UpdateOrganizationMembersPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationMembersPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeRequestParams struct {
	// 节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type UpdateOrganizationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 节点ID。可以通过[DescribeOrganizationNodes](https://cloud.tencent.com/document/product/850/82926)获取
	NodeId *uint64 `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称。最大长度为40个字符，支持英文字母、数字、汉字、符号+@、&._[]-
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *UpdateOrganizationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateOrganizationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateOrganizationNodeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateOrganizationNodeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateOrganizationNodeResponseParams `json:"Response"`
}

func (r *UpdateOrganizationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateOrganizationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyRequestParams struct {
	// 需要编辑的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略名。长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type UpdatePolicyRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的策略ID。可以调用[ListPolicies](https://cloud.tencent.com/document/product/850/105311)获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 策略描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 策略内容。参考CAM策略语法
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 策略名。长度为1~128个字符，可以包含汉字、英文字母、数字和下划线（_）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略类型。默认值SERVICE_CONTROL_POLICY，取值范围：SERVICE_CONTROL_POLICY-服务控制策略、TAG_POLICY-标签策略
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	delete(f, "Description")
	delete(f, "Content")
	delete(f, "Name")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdatePolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdatePolicyResponseParams struct {
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
type UpdateRoleConfigurationRequestParams struct {
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 新的权限配置描述。  长度：最大 1024 个字符。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 新的会话持续时间。  CIC 用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。  单位：秒。  取值范围：900-43200（15 分钟-12 小时）。
	NewSessionDuration *int64 `json:"NewSessionDuration,omitnil,omitempty" name:"NewSessionDuration"`

	// 新的初始访问页面。  CIC 用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。  该页面必须是腾讯云控制台页面。
	NewRelayState *string `json:"NewRelayState,omitnil,omitempty" name:"NewRelayState"`
}

type UpdateRoleConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 权限配置 ID
	RoleConfigurationId *string `json:"RoleConfigurationId,omitnil,omitempty" name:"RoleConfigurationId"`

	// 新的权限配置描述。  长度：最大 1024 个字符。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 新的会话持续时间。  CIC 用户使用权限配置访问集团账号目标账号时，会话最多保持的时间。  单位：秒。  取值范围：900-43200（15 分钟-12 小时）。
	NewSessionDuration *int64 `json:"NewSessionDuration,omitnil,omitempty" name:"NewSessionDuration"`

	// 新的初始访问页面。  CIC 用户使用权限配置访问集团账号目标账号时，初始访问的页面地址。  该页面必须是腾讯云控制台页面。
	NewRelayState *string `json:"NewRelayState,omitnil,omitempty" name:"NewRelayState"`
}

func (r *UpdateRoleConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RoleConfigurationId")
	delete(f, "NewDescription")
	delete(f, "NewSessionDuration")
	delete(f, "NewRelayState")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRoleConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRoleConfigurationResponseParams struct {
	// 权限配置详情
	RoleConfigurationInfo *RoleConfiguration `json:"RoleConfigurationInfo,omitnil,omitempty" name:"RoleConfigurationInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRoleConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRoleConfigurationResponseParams `json:"Response"`
}

func (r *UpdateRoleConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRoleConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSCIMCredentialStatusRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// SCIM密钥状态。Enabled：启用。 Disabled：禁用。
	NewStatus *string `json:"NewStatus,omitnil,omitempty" name:"NewStatus"`
}

type UpdateSCIMCredentialStatusRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM密钥ID。scimcred-前缀开头，后面是12位随机数字/小写字母。
	CredentialId *string `json:"CredentialId,omitnil,omitempty" name:"CredentialId"`

	// SCIM密钥状态。Enabled：启用。 Disabled：禁用。
	NewStatus *string `json:"NewStatus,omitnil,omitempty" name:"NewStatus"`
}

func (r *UpdateSCIMCredentialStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSCIMCredentialStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CredentialId")
	delete(f, "NewStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSCIMCredentialStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSCIMCredentialStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSCIMCredentialStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSCIMCredentialStatusResponseParams `json:"Response"`
}

func (r *UpdateSCIMCredentialStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSCIMCredentialStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSCIMSynchronizationStatusRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM 同步状态。Enabled：启用。Disabled：禁用。
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitnil,omitempty" name:"SCIMSynchronizationStatus"`
}

type UpdateSCIMSynchronizationStatusRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// SCIM 同步状态。Enabled：启用。Disabled：禁用。
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitnil,omitempty" name:"SCIMSynchronizationStatus"`
}

func (r *UpdateSCIMSynchronizationStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSCIMSynchronizationStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SCIMSynchronizationStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSCIMSynchronizationStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSCIMSynchronizationStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSCIMSynchronizationStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSCIMSynchronizationStatusResponseParams `json:"Response"`
}

func (r *UpdateSCIMSynchronizationStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSCIMSynchronizationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateShareUnitRequestParams struct {
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号，默认值：1
	ShareScope *uint64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`
}

type UpdateShareUnitRequest struct {
	*tchttp.BaseRequest
	
	// 共享单元ID。
	UnitId *string `json:"UnitId,omitnil,omitempty" name:"UnitId"`

	// 共享单元名称。仅支持大小写字母、数字、-、以及_的组合，3-128个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 共享单元描述。最大128个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 共享范围。取值：1-仅允许集团组织内共享 2-允许共享给任意账号，默认值：1
	ShareScope *uint64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`
}

func (r *UpdateShareUnitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateShareUnitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UnitId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "ShareScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateShareUnitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateShareUnitResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateShareUnitResponse struct {
	*tchttp.BaseResponse
	Response *UpdateShareUnitResponseParams `json:"Response"`
}

func (r *UpdateShareUnitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateShareUnitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserRequestParams struct {
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的名。
	NewFirstName *string `json:"NewFirstName,omitnil,omitempty" name:"NewFirstName"`

	// 用户的姓。
	NewLastName *string `json:"NewLastName,omitnil,omitempty" name:"NewLastName"`

	// 用户的显示名称。
	NewDisplayName *string `json:"NewDisplayName,omitnil,omitempty" name:"NewDisplayName"`

	// 用户的描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 用户的电子邮箱。
	NewEmail *string `json:"NewEmail,omitnil,omitempty" name:"NewEmail"`

	// 是否需要重置密码
	NeedResetPassword *bool `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`
}

type UpdateUserRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的名。
	NewFirstName *string `json:"NewFirstName,omitnil,omitempty" name:"NewFirstName"`

	// 用户的姓。
	NewLastName *string `json:"NewLastName,omitnil,omitempty" name:"NewLastName"`

	// 用户的显示名称。
	NewDisplayName *string `json:"NewDisplayName,omitnil,omitempty" name:"NewDisplayName"`

	// 用户的描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 用户的电子邮箱。
	NewEmail *string `json:"NewEmail,omitnil,omitempty" name:"NewEmail"`

	// 是否需要重置密码
	NeedResetPassword *bool `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`
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
	delete(f, "ZoneId")
	delete(f, "UserId")
	delete(f, "NewFirstName")
	delete(f, "NewLastName")
	delete(f, "NewDisplayName")
	delete(f, "NewDescription")
	delete(f, "NewEmail")
	delete(f, "NeedResetPassword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserResponseParams struct {
	// 用户信息
	UserInfo *UserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

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
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的状态。取值：  Enabled：启用。 Disabled：禁用。
	NewUserStatus *string `json:"NewUserStatus,omitnil,omitempty" name:"NewUserStatus"`
}

type UpdateUserStatusRequest struct {
	*tchttp.BaseRequest
	
	// 空间 ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户 ID。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的状态。取值：  Enabled：启用。 Disabled：禁用。
	NewUserStatus *string `json:"NewUserStatus,omitnil,omitempty" name:"NewUserStatus"`
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
	delete(f, "ZoneId")
	delete(f, "UserId")
	delete(f, "NewUserStatus")
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
type UpdateUserSyncProvisioningRequestParams struct {
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的iD
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	// 用户同步描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。 
	NewDuplicationStateful *string `json:"NewDuplicationStateful,omitnil,omitempty" name:"NewDuplicationStateful"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。 
	NewDeletionStrategy *string `json:"NewDeletionStrategy,omitnil,omitempty" name:"NewDeletionStrategy"`
}

type UpdateUserSyncProvisioningRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 用户同步的iD
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	// 用户同步描述。
	NewDescription *string `json:"NewDescription,omitnil,omitempty" name:"NewDescription"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。 
	NewDuplicationStateful *string `json:"NewDuplicationStateful,omitnil,omitempty" name:"NewDuplicationStateful"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。 
	NewDeletionStrategy *string `json:"NewDeletionStrategy,omitnil,omitempty" name:"NewDeletionStrategy"`
}

func (r *UpdateUserSyncProvisioningRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSyncProvisioningRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "UserProvisioningId")
	delete(f, "NewDescription")
	delete(f, "NewDuplicationStateful")
	delete(f, "NewDeletionStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserSyncProvisioningRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserSyncProvisioningResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserSyncProvisioningResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserSyncProvisioningResponseParams `json:"Response"`
}

func (r *UpdateUserSyncProvisioningResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserSyncProvisioningResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateZoneRequestParams struct {
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。
	NewZoneName *string `json:"NewZoneName,omitnil,omitempty" name:"NewZoneName"`
}

type UpdateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID。z-前缀开头，后面是12位随机数字/小写字母
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 空间名，必须全局唯一。包含小写字母、数字和短划线（-）。不能以短划线（-）开头或结尾，且不能有两个连续的短划线（-）。长度：2~64 个字符。
	NewZoneName *string `json:"NewZoneName,omitnil,omitempty" name:"NewZoneName"`
}

func (r *UpdateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "NewZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateZoneResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateZoneResponse struct {
	*tchttp.BaseResponse
	Response *UpdateZoneResponseParams `json:"Response"`
}

func (r *UpdateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserInfo struct {
	// 查询username。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户的名。
	FirstName *string `json:"FirstName,omitnil,omitempty" name:"FirstName"`

	// 用户的姓。
	LastName *string `json:"LastName,omitnil,omitempty" name:"LastName"`

	// 用户的显示名称。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 用户的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户的电子邮箱。目录内必须唯一。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户状态 Enabled：启用， Disabled：禁用。
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户类型  Manual：手动创建，Synchronized：外部导入。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户的修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否选中
	IsSelected *bool `json:"IsSelected,omitnil,omitempty" name:"IsSelected"`

	// 用户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 下次登录是否需要重置密码， true: 需要重置密码， false：不需要重置密码
	NeedResetPassword *bool `json:"NeedResetPassword,omitnil,omitempty" name:"NeedResetPassword"`
}

type UserProvisioning struct {
	// CAM 用户同步的状态。取值：
	// 
	// Enabled：CAM 用户同步已启用。
	// Disabled：CAM 用户同步未启用。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CAM 用户同步的状态。取值：
	// Enabled：CAM 用户同步已启用。
	// Disabled：CAM 用户同步未启用。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// CAM 用户同步的身份 ID。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。
	// 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份名称。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组名称。
	// 当PrincipalType取值为User时，该值为CIC用户名称。
	PrincipalName *string `json:"PrincipalName,omitnil,omitempty" name:"PrincipalName"`

	// CAM 用户同步的身份类型。取值：
	// 
	// User：表示该 CAM 用户同步的身份是CIC用户。
	// Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 集团账号目标账号的名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。
	DuplicationStrategy *string `json:"DuplicationStrategy,omitnil,omitempty" name:"DuplicationStrategy"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。
	DeletionStrategy *string `json:"DeletionStrategy,omitnil,omitempty" name:"DeletionStrategy"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

type UserProvisioningsTask struct {
	// 任务ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 授权的集团账号目标账号的UIN
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 任务类型。StartProvisioning：用户同步，DeleteProvisioning：删除用户同步
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态：InProgress: 进行中，Failed: 失败 3:Success: 成功
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 用户同步ID。
	UserProvisioningId *string `json:"UserProvisioningId,omitnil,omitempty" name:"UserProvisioningId"`

	//  CAM 用户同步的身份 ID。取值： 当PrincipalType取值为Group时，该值为CIC 用户组 ID（g-********）。 当PrincipalType取值为User时，该值为CIC 用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值： User：表示该 CAM 用户同步的身份是CIC 用户。 Group：表示该 CAM 用户同步的身份是CIC 用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 用户或者用户组名称。
	PrincipalName *string `json:"PrincipalName,omitnil,omitempty" name:"PrincipalName"`

	// 冲突策略。KeepBoth:两者都保留;TakeOver:替换
	DuplicationStrategy *string `json:"DuplicationStrategy,omitnil,omitempty" name:"DuplicationStrategy"`

	// 删除策略。Delete:删除;Keep:保留
	DeletionStrategy *string `json:"DeletionStrategy,omitnil,omitempty" name:"DeletionStrategy"`
}

type UserSyncProvisioning struct {
	// 描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// CAM 用户同步的身份 ID。取值：
	// 当PrincipalType取值为Group时，该值为CIC用户组 ID（g-********）。
	// 当PrincipalType取值为User时，该值为CIC用户 ID（u-********）。
	PrincipalId *string `json:"PrincipalId,omitnil,omitempty" name:"PrincipalId"`

	// CAM 用户同步的身份类型。取值：
	// 
	// User：表示该 CAM 用户同步的身份是CIC用户。
	// Group：表示该 CAM 用户同步的身份是CIC用户组。
	PrincipalType *string `json:"PrincipalType,omitnil,omitempty" name:"PrincipalType"`

	// 同步的集团账号目标账号的UIN。
	TargetUin *int64 `json:"TargetUin,omitnil,omitempty" name:"TargetUin"`

	// 冲突策略。当CIC 用户同步到 CAM 时，如果 CAM 中存在同名用户时的处理策略。取值： KeepBoth：两者都保留。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则对CIC 用户的用户名添加后缀_cic后尝试创建该用户名的 CAM 用户。 TakeOver：替换。当CIC 用户被同步到 CAM 时，如果 CAM 已经存在同名用户，则直接将已经存在的 CAM 用户替换为CIC 同步用户。
	DuplicationStrategy *string `json:"DuplicationStrategy,omitnil,omitempty" name:"DuplicationStrategy"`

	// 删除策略。删除 CAM 用户同步时，对已同步的 CAM 用户的处理策略。取值： Delete：删除。删除 CAM 用户同步时，会删除从CIC 已经同步到 CAM 中的 CAM 用户。 Keep：保留。删除 RAM 用户同步时，会保留从CIC 已经同步到 CAM 中的 CAM 用户。
	DeletionStrategy *string `json:"DeletionStrategy,omitnil,omitempty" name:"DeletionStrategy"`

	// 同步的集团账号目标账号的类型，ManagerUin管理账号;MemberUin成员账号
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

type ZoneStatistics struct {
	// 用户配额。
	UserQuota *int64 `json:"UserQuota,omitnil,omitempty" name:"UserQuota"`

	// 用户组配额。
	GroupQuota *int64 `json:"GroupQuota,omitnil,omitempty" name:"GroupQuota"`

	// 权限配置配额。
	RoleConfigurationQuota *int64 `json:"RoleConfigurationQuota,omitnil,omitempty" name:"RoleConfigurationQuota"`

	// 权限配置绑定的系统策略配额。
	SystemPolicyPerRoleConfigurationQuota *int64 `json:"SystemPolicyPerRoleConfigurationQuota,omitnil,omitempty" name:"SystemPolicyPerRoleConfigurationQuota"`

	// 用户数。
	UserCount *int64 `json:"UserCount,omitnil,omitempty" name:"UserCount"`

	// 用户组数。
	GroupCount *int64 `json:"GroupCount,omitnil,omitempty" name:"GroupCount"`

	// 权限配置数
	RoleConfigurationCount *int64 `json:"RoleConfigurationCount,omitnil,omitempty" name:"RoleConfigurationCount"`

	// 同步用户数。
	UserProvisioningCount *int64 `json:"UserProvisioningCount,omitnil,omitempty" name:"UserProvisioningCount"`

	// 同步角色数。
	RoleConfigurationSyncCount *int64 `json:"RoleConfigurationSyncCount,omitnil,omitempty" name:"RoleConfigurationSyncCount"`
}