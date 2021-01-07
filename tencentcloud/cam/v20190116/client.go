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
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-16"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddUserRequest() (request *AddUserRequest) {
    request = &AddUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AddUser")
    return
}

func NewAddUserResponse() (response *AddUserResponse) {
    response = &AddUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加子用户
func (c *Client) AddUser(request *AddUserRequest) (response *AddUserResponse, err error) {
    if request == nil {
        request = NewAddUserRequest()
    }
    response = NewAddUserResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
    request = &AddUserToGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AddUserToGroup")
    return
}

func NewAddUserToGroupResponse() (response *AddUserToGroupResponse) {
    response = &AddUserToGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用户加入到用户组
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToGroupRequest()
    }
    response = NewAddUserToGroupResponse()
    err = c.Send(request, response)
    return
}

func NewAttachGroupPolicyRequest() (request *AttachGroupPolicyRequest) {
    request = &AttachGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AttachGroupPolicy")
    return
}

func NewAttachGroupPolicyResponse() (response *AttachGroupPolicyResponse) {
    response = &AttachGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachGroupPolicy）可用于绑定策略到用户组。
func (c *Client) AttachGroupPolicy(request *AttachGroupPolicyRequest) (response *AttachGroupPolicyResponse, err error) {
    if request == nil {
        request = NewAttachGroupPolicyRequest()
    }
    response = NewAttachGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAttachRolePolicyRequest() (request *AttachRolePolicyRequest) {
    request = &AttachRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AttachRolePolicy")
    return
}

func NewAttachRolePolicyResponse() (response *AttachRolePolicyResponse) {
    response = &AttachRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachRolePolicy）用于绑定策略到角色。
func (c *Client) AttachRolePolicy(request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
    if request == nil {
        request = NewAttachRolePolicyRequest()
    }
    response = NewAttachRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewAttachUserPolicyRequest() (request *AttachUserPolicyRequest) {
    request = &AttachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "AttachUserPolicy")
    return
}

func NewAttachUserPolicyResponse() (response *AttachUserPolicyResponse) {
    response = &AttachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（AttachUserPolicy）可用于绑定到用户的策略。
func (c *Client) AttachUserPolicy(request *AttachUserPolicyRequest) (response *AttachUserPolicyResponse, err error) {
    if request == nil {
        request = NewAttachUserPolicyRequest()
    }
    response = NewAttachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewConsumeCustomMFATokenRequest() (request *ConsumeCustomMFATokenRequest) {
    request = &ConsumeCustomMFATokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ConsumeCustomMFAToken")
    return
}

func NewConsumeCustomMFATokenResponse() (response *ConsumeCustomMFATokenResponse) {
    response = &ConsumeCustomMFATokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证自定义多因子Token
func (c *Client) ConsumeCustomMFAToken(request *ConsumeCustomMFATokenRequest) (response *ConsumeCustomMFATokenResponse, err error) {
    if request == nil {
        request = NewConsumeCustomMFATokenRequest()
    }
    response = NewConsumeCustomMFATokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建用户组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
    request = &CreatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreatePolicy")
    return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
    response = &CreatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreatePolicy）可用于创建策略。
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyVersionRequest() (request *CreatePolicyVersionRequest) {
    request = &CreatePolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreatePolicyVersion")
    return
}

func NewCreatePolicyVersionResponse() (response *CreatePolicyVersionResponse) {
    response = &CreatePolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（CreatePolicyVersion）用于新增策略版本，用户创建了一个策略版本之后可以方便的通过变更策略版本的方式来变更策略。
func (c *Client) CreatePolicyVersion(request *CreatePolicyVersionRequest) (response *CreatePolicyVersionResponse, err error) {
    if request == nil {
        request = NewCreatePolicyVersionRequest()
    }
    response = NewCreatePolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
    request = &CreateRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateRole")
    return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
    response = &CreateRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateRole）用于创建角色。
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
    if request == nil {
        request = NewCreateRoleRequest()
    }
    response = NewCreateRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSAMLProviderRequest() (request *CreateSAMLProviderRequest) {
    request = &CreateSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateSAMLProvider")
    return
}

func NewCreateSAMLProviderResponse() (response *CreateSAMLProviderResponse) {
    response = &CreateSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建SAML身份提供商
func (c *Client) CreateSAMLProvider(request *CreateSAMLProviderRequest) (response *CreateSAMLProviderResponse, err error) {
    if request == nil {
        request = NewCreateSAMLProviderRequest()
    }
    response = NewCreateSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateServiceLinkedRoleRequest() (request *CreateServiceLinkedRoleRequest) {
    request = &CreateServiceLinkedRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "CreateServiceLinkedRole")
    return
}

func NewCreateServiceLinkedRoleResponse() (response *CreateServiceLinkedRoleResponse) {
    response = &CreateServiceLinkedRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建服务相关角色
func (c *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (response *CreateServiceLinkedRoleResponse, err error) {
    if request == nil {
        request = NewCreateServiceLinkedRoleRequest()
    }
    response = NewCreateServiceLinkedRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
    request = &DeletePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeletePolicy")
    return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
    response = &DeletePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeletePolicy）可用于删除策略。
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    if request == nil {
        request = NewDeletePolicyRequest()
    }
    response = NewDeletePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyVersionRequest() (request *DeletePolicyVersionRequest) {
    request = &DeletePolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeletePolicyVersion")
    return
}

func NewDeletePolicyVersionResponse() (response *DeletePolicyVersionResponse) {
    response = &DeletePolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeletePolicyVersion）可用于删除一个策略的策略版本。
func (c *Client) DeletePolicyVersion(request *DeletePolicyVersionRequest) (response *DeletePolicyVersionResponse, err error) {
    if request == nil {
        request = NewDeletePolicyVersionRequest()
    }
    response = NewDeletePolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
    request = &DeleteRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteRole")
    return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
    response = &DeleteRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteRole）用于删除指定角色。
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
    if request == nil {
        request = NewDeleteRoleRequest()
    }
    response = NewDeleteRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRolePermissionsBoundaryRequest() (request *DeleteRolePermissionsBoundaryRequest) {
    request = &DeleteRolePermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteRolePermissionsBoundary")
    return
}

func NewDeleteRolePermissionsBoundaryResponse() (response *DeleteRolePermissionsBoundaryResponse) {
    response = &DeleteRolePermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除角色权限边界
func (c *Client) DeleteRolePermissionsBoundary(request *DeleteRolePermissionsBoundaryRequest) (response *DeleteRolePermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewDeleteRolePermissionsBoundaryRequest()
    }
    response = NewDeleteRolePermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSAMLProviderRequest() (request *DeleteSAMLProviderRequest) {
    request = &DeleteSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteSAMLProvider")
    return
}

func NewDeleteSAMLProviderResponse() (response *DeleteSAMLProviderResponse) {
    response = &DeleteSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除SAML身份提供商
func (c *Client) DeleteSAMLProvider(request *DeleteSAMLProviderRequest) (response *DeleteSAMLProviderResponse, err error) {
    if request == nil {
        request = NewDeleteSAMLProviderRequest()
    }
    response = NewDeleteSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteServiceLinkedRoleRequest() (request *DeleteServiceLinkedRoleRequest) {
    request = &DeleteServiceLinkedRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteServiceLinkedRole")
    return
}

func NewDeleteServiceLinkedRoleResponse() (response *DeleteServiceLinkedRoleResponse) {
    response = &DeleteServiceLinkedRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除服务相关角色
func (c *Client) DeleteServiceLinkedRole(request *DeleteServiceLinkedRoleRequest) (response *DeleteServiceLinkedRoleResponse, err error) {
    if request == nil {
        request = NewDeleteServiceLinkedRoleRequest()
    }
    response = NewDeleteServiceLinkedRoleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteUser")
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除子用户
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserPermissionsBoundaryRequest() (request *DeleteUserPermissionsBoundaryRequest) {
    request = &DeleteUserPermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DeleteUserPermissionsBoundary")
    return
}

func NewDeleteUserPermissionsBoundaryResponse() (response *DeleteUserPermissionsBoundaryResponse) {
    response = &DeleteUserPermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除用户权限边界
func (c *Client) DeleteUserPermissionsBoundary(request *DeleteUserPermissionsBoundaryRequest) (response *DeleteUserPermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewDeleteUserPermissionsBoundaryRequest()
    }
    response = NewDeleteUserPermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
    request = &DescribeRoleListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DescribeRoleList")
    return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
    response = &DescribeRoleListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeRoleList）用于获取账号下的角色列表。
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
    if request == nil {
        request = NewDescribeRoleListRequest()
    }
    response = NewDescribeRoleListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeAuthFlagRequest() (request *DescribeSafeAuthFlagRequest) {
    request = &DescribeSafeAuthFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DescribeSafeAuthFlag")
    return
}

func NewDescribeSafeAuthFlagResponse() (response *DescribeSafeAuthFlagResponse) {
    response = &DescribeSafeAuthFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询安全设置
func (c *Client) DescribeSafeAuthFlag(request *DescribeSafeAuthFlagRequest) (response *DescribeSafeAuthFlagResponse, err error) {
    if request == nil {
        request = NewDescribeSafeAuthFlagRequest()
    }
    response = NewDescribeSafeAuthFlagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSafeAuthFlagCollRequest() (request *DescribeSafeAuthFlagCollRequest) {
    request = &DescribeSafeAuthFlagCollRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DescribeSafeAuthFlagColl")
    return
}

func NewDescribeSafeAuthFlagCollResponse() (response *DescribeSafeAuthFlagCollResponse) {
    response = &DescribeSafeAuthFlagCollResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询安全设置
func (c *Client) DescribeSafeAuthFlagColl(request *DescribeSafeAuthFlagCollRequest) (response *DescribeSafeAuthFlagCollResponse, err error) {
    if request == nil {
        request = NewDescribeSafeAuthFlagCollRequest()
    }
    response = NewDescribeSafeAuthFlagCollResponse()
    err = c.Send(request, response)
    return
}

func NewDetachGroupPolicyRequest() (request *DetachGroupPolicyRequest) {
    request = &DetachGroupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DetachGroupPolicy")
    return
}

func NewDetachGroupPolicyResponse() (response *DetachGroupPolicyResponse) {
    response = &DetachGroupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DetachGroupPolicy）可用于解除绑定到用户组的策略。
func (c *Client) DetachGroupPolicy(request *DetachGroupPolicyRequest) (response *DetachGroupPolicyResponse, err error) {
    if request == nil {
        request = NewDetachGroupPolicyRequest()
    }
    response = NewDetachGroupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachRolePolicyRequest() (request *DetachRolePolicyRequest) {
    request = &DetachRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DetachRolePolicy")
    return
}

func NewDetachRolePolicyResponse() (response *DetachRolePolicyResponse) {
    response = &DetachRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DetachRolePolicy）用于解除绑定角色的策略。
func (c *Client) DetachRolePolicy(request *DetachRolePolicyRequest) (response *DetachRolePolicyResponse, err error) {
    if request == nil {
        request = NewDetachRolePolicyRequest()
    }
    response = NewDetachRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDetachUserPolicyRequest() (request *DetachUserPolicyRequest) {
    request = &DetachUserPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "DetachUserPolicy")
    return
}

func NewDetachUserPolicyResponse() (response *DetachUserPolicyResponse) {
    response = &DetachUserPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DetachUserPolicy）可用于解除绑定到用户的策略。
func (c *Client) DetachUserPolicy(request *DetachUserPolicyRequest) (response *DetachUserPolicyResponse, err error) {
    if request == nil {
        request = NewDetachUserPolicyRequest()
    }
    response = NewDetachUserPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewGetCustomMFATokenInfoRequest() (request *GetCustomMFATokenInfoRequest) {
    request = &GetCustomMFATokenInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetCustomMFATokenInfo")
    return
}

func NewGetCustomMFATokenInfoResponse() (response *GetCustomMFATokenInfoResponse) {
    response = &GetCustomMFATokenInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取自定义多因子Token关联信息
func (c *Client) GetCustomMFATokenInfo(request *GetCustomMFATokenInfoRequest) (response *GetCustomMFATokenInfoResponse, err error) {
    if request == nil {
        request = NewGetCustomMFATokenInfoRequest()
    }
    response = NewGetCustomMFATokenInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupRequest() (request *GetGroupRequest) {
    request = &GetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetGroup")
    return
}

func NewGetGroupResponse() (response *GetGroupResponse) {
    response = &GetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户组详情
func (c *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
    if request == nil {
        request = NewGetGroupRequest()
    }
    response = NewGetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewGetPolicyRequest() (request *GetPolicyRequest) {
    request = &GetPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetPolicy")
    return
}

func NewGetPolicyResponse() (response *GetPolicyResponse) {
    response = &GetPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetPolicy）可用于查询查看策略详情。
func (c *Client) GetPolicy(request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
    if request == nil {
        request = NewGetPolicyRequest()
    }
    response = NewGetPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewGetPolicyVersionRequest() (request *GetPolicyVersionRequest) {
    request = &GetPolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetPolicyVersion")
    return
}

func NewGetPolicyVersionResponse() (response *GetPolicyVersionResponse) {
    response = &GetPolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（GetPolicyVersion）用于查询策略版本详情
func (c *Client) GetPolicyVersion(request *GetPolicyVersionRequest) (response *GetPolicyVersionResponse, err error) {
    if request == nil {
        request = NewGetPolicyVersionRequest()
    }
    response = NewGetPolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoleRequest() (request *GetRoleRequest) {
    request = &GetRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetRole")
    return
}

func NewGetRoleResponse() (response *GetRoleResponse) {
    response = &GetRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（GetRole）用于获取指定角色的详细信息。
func (c *Client) GetRole(request *GetRoleRequest) (response *GetRoleResponse, err error) {
    if request == nil {
        request = NewGetRoleRequest()
    }
    response = NewGetRoleResponse()
    err = c.Send(request, response)
    return
}

func NewGetSAMLProviderRequest() (request *GetSAMLProviderRequest) {
    request = &GetSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetSAMLProvider")
    return
}

func NewGetSAMLProviderResponse() (response *GetSAMLProviderResponse) {
    response = &GetSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询SAML身份提供商详情
func (c *Client) GetSAMLProvider(request *GetSAMLProviderRequest) (response *GetSAMLProviderResponse, err error) {
    if request == nil {
        request = NewGetSAMLProviderRequest()
    }
    response = NewGetSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewGetServiceLinkedRoleDeletionStatusRequest() (request *GetServiceLinkedRoleDeletionStatusRequest) {
    request = &GetServiceLinkedRoleDeletionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetServiceLinkedRoleDeletionStatus")
    return
}

func NewGetServiceLinkedRoleDeletionStatusResponse() (response *GetServiceLinkedRoleDeletionStatusResponse) {
    response = &GetServiceLinkedRoleDeletionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据删除TaskId获取服务相关角色删除状态
func (c *Client) GetServiceLinkedRoleDeletionStatus(request *GetServiceLinkedRoleDeletionStatusRequest) (response *GetServiceLinkedRoleDeletionStatusResponse, err error) {
    if request == nil {
        request = NewGetServiceLinkedRoleDeletionStatusRequest()
    }
    response = NewGetServiceLinkedRoleDeletionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserRequest() (request *GetUserRequest) {
    request = &GetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "GetUser")
    return
}

func NewGetUserResponse() (response *GetUserResponse) {
    response = &GetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询子用户
func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
    if request == nil {
        request = NewGetUserRequest()
    }
    response = NewGetUserResponse()
    err = c.Send(request, response)
    return
}

func NewListAccessKeysRequest() (request *ListAccessKeysRequest) {
    request = &ListAccessKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListAccessKeys")
    return
}

func NewListAccessKeysResponse() (response *ListAccessKeysResponse) {
    response = &ListAccessKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 列出指定CAM用户的访问密钥
func (c *Client) ListAccessKeys(request *ListAccessKeysRequest) (response *ListAccessKeysResponse, err error) {
    if request == nil {
        request = NewListAccessKeysRequest()
    }
    response = NewListAccessKeysResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedGroupPoliciesRequest() (request *ListAttachedGroupPoliciesRequest) {
    request = &ListAttachedGroupPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedGroupPolicies")
    return
}

func NewListAttachedGroupPoliciesResponse() (response *ListAttachedGroupPoliciesResponse) {
    response = &ListAttachedGroupPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListAttachedGroupPolicies）可用于查询用户组关联的策略列表。
func (c *Client) ListAttachedGroupPolicies(request *ListAttachedGroupPoliciesRequest) (response *ListAttachedGroupPoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedGroupPoliciesRequest()
    }
    response = NewListAttachedGroupPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedRolePoliciesRequest() (request *ListAttachedRolePoliciesRequest) {
    request = &ListAttachedRolePoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedRolePolicies")
    return
}

func NewListAttachedRolePoliciesResponse() (response *ListAttachedRolePoliciesResponse) {
    response = &ListAttachedRolePoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListAttachedRolePolicies）用于获取角色绑定的策略列表。
func (c *Client) ListAttachedRolePolicies(request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedRolePoliciesRequest()
    }
    response = NewListAttachedRolePoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListAttachedUserPoliciesRequest() (request *ListAttachedUserPoliciesRequest) {
    request = &ListAttachedUserPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListAttachedUserPolicies")
    return
}

func NewListAttachedUserPoliciesResponse() (response *ListAttachedUserPoliciesResponse) {
    response = &ListAttachedUserPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListAttachedUserPolicies）可用于查询子账号关联的策略列表。
func (c *Client) ListAttachedUserPolicies(request *ListAttachedUserPoliciesRequest) (response *ListAttachedUserPoliciesResponse, err error) {
    if request == nil {
        request = NewListAttachedUserPoliciesRequest()
    }
    response = NewListAttachedUserPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListCollaboratorsRequest() (request *ListCollaboratorsRequest) {
    request = &ListCollaboratorsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListCollaborators")
    return
}

func NewListCollaboratorsResponse() (response *ListCollaboratorsResponse) {
    response = &ListCollaboratorsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取协作者列表
func (c *Client) ListCollaborators(request *ListCollaboratorsRequest) (response *ListCollaboratorsResponse, err error) {
    if request == nil {
        request = NewListCollaboratorsRequest()
    }
    response = NewListCollaboratorsResponse()
    err = c.Send(request, response)
    return
}

func NewListEntitiesForPolicyRequest() (request *ListEntitiesForPolicyRequest) {
    request = &ListEntitiesForPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListEntitiesForPolicy")
    return
}

func NewListEntitiesForPolicyResponse() (response *ListEntitiesForPolicyResponse) {
    response = &ListEntitiesForPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListEntitiesForPolicy）可用于查询策略关联的实体列表。
func (c *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
    if request == nil {
        request = NewListEntitiesForPolicyRequest()
    }
    response = NewListEntitiesForPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupsRequest() (request *ListGroupsRequest) {
    request = &ListGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListGroups")
    return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
    response = &ListGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户组列表
func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    if request == nil {
        request = NewListGroupsRequest()
    }
    response = NewListGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupsForUserRequest() (request *ListGroupsForUserRequest) {
    request = &ListGroupsForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListGroupsForUser")
    return
}

func NewListGroupsForUserResponse() (response *ListGroupsForUserResponse) {
    response = &ListGroupsForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 列出用户关联的用户组
func (c *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (response *ListGroupsForUserResponse, err error) {
    if request == nil {
        request = NewListGroupsForUserRequest()
    }
    response = NewListGroupsForUserResponse()
    err = c.Send(request, response)
    return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
    request = &ListPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListPolicies")
    return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
    response = &ListPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ListPolicies）可用于查询策略列表。
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    if request == nil {
        request = NewListPoliciesRequest()
    }
    response = NewListPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListPolicyVersionsRequest() (request *ListPolicyVersionsRequest) {
    request = &ListPolicyVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListPolicyVersions")
    return
}

func NewListPolicyVersionsResponse() (response *ListPolicyVersionsResponse) {
    response = &ListPolicyVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 该接口（ListPolicyVersions）用于获取策略版本列表
func (c *Client) ListPolicyVersions(request *ListPolicyVersionsRequest) (response *ListPolicyVersionsResponse, err error) {
    if request == nil {
        request = NewListPolicyVersionsRequest()
    }
    response = NewListPolicyVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewListSAMLProvidersRequest() (request *ListSAMLProvidersRequest) {
    request = &ListSAMLProvidersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListSAMLProviders")
    return
}

func NewListSAMLProvidersResponse() (response *ListSAMLProvidersResponse) {
    response = &ListSAMLProvidersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询SAML身份提供商列表
func (c *Client) ListSAMLProviders(request *ListSAMLProvidersRequest) (response *ListSAMLProvidersResponse, err error) {
    if request == nil {
        request = NewListSAMLProvidersRequest()
    }
    response = NewListSAMLProvidersResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListUsers")
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 拉取子用户
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    if request == nil {
        request = NewListUsersRequest()
    }
    response = NewListUsersResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersForGroupRequest() (request *ListUsersForGroupRequest) {
    request = &ListUsersForGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListUsersForGroup")
    return
}

func NewListUsersForGroupResponse() (response *ListUsersForGroupResponse) {
    response = &ListUsersForGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户组关联的用户列表
func (c *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (response *ListUsersForGroupResponse, err error) {
    if request == nil {
        request = NewListUsersForGroupRequest()
    }
    response = NewListUsersForGroupResponse()
    err = c.Send(request, response)
    return
}

func NewListWeChatWorkSubAccountsRequest() (request *ListWeChatWorkSubAccountsRequest) {
    request = &ListWeChatWorkSubAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "ListWeChatWorkSubAccounts")
    return
}

func NewListWeChatWorkSubAccountsResponse() (response *ListWeChatWorkSubAccountsResponse) {
    response = &ListWeChatWorkSubAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取企业微信子用户列表
func (c *Client) ListWeChatWorkSubAccounts(request *ListWeChatWorkSubAccountsRequest) (response *ListWeChatWorkSubAccountsResponse, err error) {
    if request == nil {
        request = NewListWeChatWorkSubAccountsRequest()
    }
    response = NewListWeChatWorkSubAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewPutRolePermissionsBoundaryRequest() (request *PutRolePermissionsBoundaryRequest) {
    request = &PutRolePermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "PutRolePermissionsBoundary")
    return
}

func NewPutRolePermissionsBoundaryResponse() (response *PutRolePermissionsBoundaryResponse) {
    response = &PutRolePermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置角色权限边界
func (c *Client) PutRolePermissionsBoundary(request *PutRolePermissionsBoundaryRequest) (response *PutRolePermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewPutRolePermissionsBoundaryRequest()
    }
    response = NewPutRolePermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewPutUserPermissionsBoundaryRequest() (request *PutUserPermissionsBoundaryRequest) {
    request = &PutUserPermissionsBoundaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "PutUserPermissionsBoundary")
    return
}

func NewPutUserPermissionsBoundaryResponse() (response *PutUserPermissionsBoundaryResponse) {
    response = &PutUserPermissionsBoundaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置用户权限边界
func (c *Client) PutUserPermissionsBoundary(request *PutUserPermissionsBoundaryRequest) (response *PutUserPermissionsBoundaryResponse, err error) {
    if request == nil {
        request = NewPutUserPermissionsBoundaryRequest()
    }
    response = NewPutUserPermissionsBoundaryResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
    request = &RemoveUserFromGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "RemoveUserFromGroup")
    return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
    response = &RemoveUserFromGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 从用户组删除用户
func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromGroupRequest()
    }
    response = NewRemoveUserFromGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSetDefaultPolicyVersionRequest() (request *SetDefaultPolicyVersionRequest) {
    request = &SetDefaultPolicyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "SetDefaultPolicyVersion")
    return
}

func NewSetDefaultPolicyVersionResponse() (response *SetDefaultPolicyVersionResponse) {
    response = &SetDefaultPolicyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（SetDefaultPolicyVersion）可用于设置生效的策略版本。
func (c *Client) SetDefaultPolicyVersion(request *SetDefaultPolicyVersionRequest) (response *SetDefaultPolicyVersionResponse, err error) {
    if request == nil {
        request = NewSetDefaultPolicyVersionRequest()
    }
    response = NewSetDefaultPolicyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewSetMfaFlagRequest() (request *SetMfaFlagRequest) {
    request = &SetMfaFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "SetMfaFlag")
    return
}

func NewSetMfaFlagResponse() (response *SetMfaFlagResponse) {
    response = &SetMfaFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置子用户的登录保护和敏感操作校验方式
func (c *Client) SetMfaFlag(request *SetMfaFlagRequest) (response *SetMfaFlagResponse, err error) {
    if request == nil {
        request = NewSetMfaFlagRequest()
    }
    response = NewSetMfaFlagResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssumeRolePolicyRequest() (request *UpdateAssumeRolePolicyRequest) {
    request = &UpdateAssumeRolePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateAssumeRolePolicy")
    return
}

func NewUpdateAssumeRolePolicyResponse() (response *UpdateAssumeRolePolicyResponse) {
    response = &UpdateAssumeRolePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateAssumeRolePolicy）用于修改角色信任策略的策略文档。
func (c *Client) UpdateAssumeRolePolicy(request *UpdateAssumeRolePolicyRequest) (response *UpdateAssumeRolePolicyResponse, err error) {
    if request == nil {
        request = NewUpdateAssumeRolePolicyRequest()
    }
    response = NewUpdateAssumeRolePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
    request = &UpdateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateGroup")
    return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
    response = &UpdateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新用户组
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    if request == nil {
        request = NewUpdateGroupRequest()
    }
    response = NewUpdateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdatePolicy")
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdatePolicy ）可用于更新策略。
// 如果已存在策略版本，本接口会直接更新策略的默认版本，不会创建新版本，如果不存在任何策略版本，则直接创建一个默认版本。
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleConsoleLoginRequest() (request *UpdateRoleConsoleLoginRequest) {
    request = &UpdateRoleConsoleLoginRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleConsoleLogin")
    return
}

func NewUpdateRoleConsoleLoginResponse() (response *UpdateRoleConsoleLoginResponse) {
    response = &UpdateRoleConsoleLoginResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateRoleConsoleLogin）用于修改角色是否可登录。
func (c *Client) UpdateRoleConsoleLogin(request *UpdateRoleConsoleLoginRequest) (response *UpdateRoleConsoleLoginResponse, err error) {
    if request == nil {
        request = NewUpdateRoleConsoleLoginRequest()
    }
    response = NewUpdateRoleConsoleLoginResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleDescriptionRequest() (request *UpdateRoleDescriptionRequest) {
    request = &UpdateRoleDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleDescription")
    return
}

func NewUpdateRoleDescriptionResponse() (response *UpdateRoleDescriptionResponse) {
    response = &UpdateRoleDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（UpdateRoleDescription）用于修改角色的描述信息。
func (c *Client) UpdateRoleDescription(request *UpdateRoleDescriptionRequest) (response *UpdateRoleDescriptionResponse, err error) {
    if request == nil {
        request = NewUpdateRoleDescriptionRequest()
    }
    response = NewUpdateRoleDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSAMLProviderRequest() (request *UpdateSAMLProviderRequest) {
    request = &UpdateSAMLProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateSAMLProvider")
    return
}

func NewUpdateSAMLProviderResponse() (response *UpdateSAMLProviderResponse) {
    response = &UpdateSAMLProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新SAML身份提供商信息
func (c *Client) UpdateSAMLProvider(request *UpdateSAMLProviderRequest) (response *UpdateSAMLProviderResponse, err error) {
    if request == nil {
        request = NewUpdateSAMLProviderRequest()
    }
    response = NewUpdateSAMLProviderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cam", APIVersion, "UpdateUser")
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新子用户
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    response = NewUpdateUserResponse()
    err = c.Send(request, response)
    return
}
