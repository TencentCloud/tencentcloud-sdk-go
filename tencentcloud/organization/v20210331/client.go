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
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-31"

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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBindOrganizationMemberAuthAccountRequest() (request *BindOrganizationMemberAuthAccountRequest) {
    request = &BindOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberAuthAccount")
    
    
    return
}

func NewBindOrganizationMemberAuthAccountResponse() (response *BindOrganizationMemberAuthAccountResponse) {
    response = &BindOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindOrganizationMemberAuthAccount
// 绑定组织成员和子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationMemberAuthAccountRequest()
    }
    
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

// BindOrganizationMemberAuthAccount
// 绑定组织成员和子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) BindOrganizationMemberAuthAccountWithContext(ctx context.Context, request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationMemberAuthAccountRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberRequest() (request *CreateOrganizationMemberRequest) {
    request = &CreateOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMember")
    
    
    return
}

func NewCreateOrganizationMemberResponse() (response *CreateOrganizationMemberResponse) {
    response = &CreateOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOrganizationMember
// 创建组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMember(request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberRequest()
    }
    
    response = NewCreateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

// CreateOrganizationMember
// 创建组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberWithContext(ctx context.Context, request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
    request = &DescribeOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganization")
    
    
    return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
    response = &DescribeOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrganization
// 获取企业组织信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

// DescribeOrganization
// 获取企业组织信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationWithContext(ctx context.Context, request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
    request = &DescribeOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembers")
    
    
    return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
    response = &DescribeOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrganizationMembers
// 获取企业组织成员列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMembersRequest()
    }
    
    response = NewDescribeOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

// DescribeOrganizationMembers
// 获取企业组织成员列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMembersWithContext(ctx context.Context, request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMembersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}
