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
    "errors"
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


func NewAddOrganizationMemberEmailRequest() (request *AddOrganizationMemberEmailRequest) {
    request = &AddOrganizationMemberEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationMemberEmail")
    
    
    return
}

func NewAddOrganizationMemberEmailResponse() (response *AddOrganizationMemberEmailResponse) {
    response = &AddOrganizationMemberEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddOrganizationMemberEmail
// 添加组织成员邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"
//  LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddOrganizationMemberEmail(request *AddOrganizationMemberEmailRequest) (response *AddOrganizationMemberEmailResponse, err error) {
    return c.AddOrganizationMemberEmailWithContext(context.Background(), request)
}

// AddOrganizationMemberEmail
// 添加组织成员邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"
//  LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddOrganizationMemberEmailWithContext(ctx context.Context, request *AddOrganizationMemberEmailRequest) (response *AddOrganizationMemberEmailResponse, err error) {
    if request == nil {
        request = NewAddOrganizationMemberEmailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganizationMemberEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationMemberEmailResponse()
    err = c.Send(request, response)
    return
}

func NewAddOrganizationNodeRequest() (request *AddOrganizationNodeRequest) {
    request = &AddOrganizationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationNode")
    
    
    return
}

func NewAddOrganizationNodeResponse() (response *AddOrganizationNodeResponse) {
    response = &AddOrganizationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddOrganizationNode
// 添加企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNode(request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    return c.AddOrganizationNodeWithContext(context.Background(), request)
}

// AddOrganizationNode
// 添加企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNodeWithContext(ctx context.Context, request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewAddOrganizationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewAddShareUnitRequest() (request *AddShareUnitRequest) {
    request = &AddShareUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddShareUnit")
    
    
    return
}

func NewAddShareUnitResponse() (response *AddShareUnitResponse) {
    response = &AddShareUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddShareUnit
// 创建共享单元，只有企业组织管理员可创建。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEOVERLIMIT = "FailedOperation.ResourceOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddShareUnit(request *AddShareUnitRequest) (response *AddShareUnitResponse, err error) {
    return c.AddShareUnitWithContext(context.Background(), request)
}

// AddShareUnit
// 创建共享单元，只有企业组织管理员可创建。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEOVERLIMIT = "FailedOperation.ResourceOverLimit"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddShareUnitWithContext(ctx context.Context, request *AddShareUnitRequest) (response *AddShareUnitResponse, err error) {
    if request == nil {
        request = NewAddShareUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitResponse()
    err = c.Send(request, response)
    return
}

func NewAddShareUnitMembersRequest() (request *AddShareUnitMembersRequest) {
    request = &AddShareUnitMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddShareUnitMembers")
    
    
    return
}

func NewAddShareUnitMembersResponse() (response *AddShareUnitMembersResponse) {
    response = &AddShareUnitMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddShareUnitMembers
// 添加共享单元成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddShareUnitMembers(request *AddShareUnitMembersRequest) (response *AddShareUnitMembersResponse, err error) {
    return c.AddShareUnitMembersWithContext(context.Background(), request)
}

// AddShareUnitMembers
// 添加共享单元成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddShareUnitMembersWithContext(ctx context.Context, request *AddShareUnitMembersRequest) (response *AddShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewAddShareUnitMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnitMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitMembersResponse()
    err = c.Send(request, response)
    return
}

func NewAddShareUnitResourcesRequest() (request *AddShareUnitResourcesRequest) {
    request = &AddShareUnitResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddShareUnitResources")
    
    
    return
}

func NewAddShareUnitResourcesResponse() (response *AddShareUnitResourcesResponse) {
    response = &AddShareUnitResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddShareUnitResources
// 添加共享单元资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASDIFFERENTRESOURCETYPE = "FailedOperation.HasDifferentResourceType"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MEMBERUNSUPPORTEDOPERATION = "UnsupportedOperation.MemberUnsupportedOperation"
func (c *Client) AddShareUnitResources(request *AddShareUnitResourcesRequest) (response *AddShareUnitResourcesResponse, err error) {
    return c.AddShareUnitResourcesWithContext(context.Background(), request)
}

// AddShareUnitResources
// 添加共享单元资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_HASDIFFERENTRESOURCETYPE = "FailedOperation.HasDifferentResourceType"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_MEMBERUNSUPPORTEDOPERATION = "UnsupportedOperation.MemberUnsupportedOperation"
func (c *Client) AddShareUnitResourcesWithContext(ctx context.Context, request *AddShareUnitResourcesRequest) (response *AddShareUnitResourcesResponse, err error) {
    if request == nil {
        request = NewAddShareUnitResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewAttachPolicyRequest() (request *AttachPolicyRequest) {
    request = &AttachPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AttachPolicy")
    
    
    return
}

func NewAttachPolicyResponse() (response *AttachPolicyResponse) {
    response = &AttachPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachPolicy
// 绑定策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AttachPolicy(request *AttachPolicyRequest) (response *AttachPolicyResponse, err error) {
    return c.AttachPolicyWithContext(context.Background(), request)
}

// AttachPolicy
// 绑定策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AttachPolicyWithContext(ctx context.Context, request *AttachPolicyRequest) (response *AttachPolicyResponse, err error) {
    if request == nil {
        request = NewAttachPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachPolicyResponse()
    err = c.Send(request, response)
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
// 绑定组织成员和组织管理员子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    return c.BindOrganizationMemberAuthAccountWithContext(context.Background(), request)
}

// BindOrganizationMemberAuthAccount
// 绑定组织成员和组织管理员子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTIDENTITYEXIST = "FailedOperation.SubAccountIdentityExist"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationMemberAuthAccountWithContext(ctx context.Context, request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationMemberAuthAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCancelOrganizationMemberAuthAccountRequest() (request *CancelOrganizationMemberAuthAccountRequest) {
    request = &CancelOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccount")
    
    
    return
}

func NewCancelOrganizationMemberAuthAccountResponse() (response *CancelOrganizationMemberAuthAccountResponse) {
    response = &CancelOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelOrganizationMemberAuthAccount
// 取消组织成员和组织管理员子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CancelOrganizationMemberAuthAccount(request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
    return c.CancelOrganizationMemberAuthAccountWithContext(context.Background(), request)
}

// CancelOrganizationMemberAuthAccount
// 取消组织成员和组织管理员子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CancelOrganizationMemberAuthAccountWithContext(ctx context.Context, request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewCancelOrganizationMemberAuthAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAccountDeleteRequest() (request *CheckAccountDeleteRequest) {
    request = &CheckAccountDeleteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CheckAccountDelete")
    
    
    return
}

func NewCheckAccountDeleteResponse() (response *CheckAccountDeleteResponse) {
    response = &CheckAccountDeleteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAccountDelete
// 成员账号删除检查
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CheckAccountDelete(request *CheckAccountDeleteRequest) (response *CheckAccountDeleteResponse, err error) {
    return c.CheckAccountDeleteWithContext(context.Background(), request)
}

// CheckAccountDelete
// 成员账号删除检查
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CheckAccountDeleteWithContext(ctx context.Context, request *CheckAccountDeleteRequest) (response *CheckAccountDeleteResponse, err error) {
    if request == nil {
        request = NewCheckAccountDeleteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAccountDelete require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAccountDeleteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationRequest() (request *CreateOrganizationRequest) {
    request = &CreateOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganization")
    
    
    return
}

func NewCreateOrganizationResponse() (response *CreateOrganizationResponse) {
    response = &CreateOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganization
// 创建企业组织
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"
func (c *Client) CreateOrganization(request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
    return c.CreateOrganizationWithContext(context.Background(), request)
}

// CreateOrganization
// 创建企业组织
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"
func (c *Client) CreateOrganizationWithContext(ctx context.Context, request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationIdentityRequest() (request *CreateOrganizationIdentityRequest) {
    request = &CreateOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationIdentity")
    
    
    return
}

func NewCreateOrganizationIdentityResponse() (response *CreateOrganizationIdentityResponse) {
    response = &CreateOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationIdentity
// 添加组织身份
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) CreateOrganizationIdentity(request *CreateOrganizationIdentityRequest) (response *CreateOrganizationIdentityResponse, err error) {
    return c.CreateOrganizationIdentityWithContext(context.Background(), request)
}

// CreateOrganizationIdentity
// 添加组织身份
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) CreateOrganizationIdentityWithContext(ctx context.Context, request *CreateOrganizationIdentityRequest) (response *CreateOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationIdentityResponse()
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
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) CreateOrganizationMember(request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    return c.CreateOrganizationMemberWithContext(context.Background(), request)
}

// CreateOrganizationMember
// 创建组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) CreateOrganizationMemberWithContext(ctx context.Context, request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberAuthIdentityRequest() (request *CreateOrganizationMemberAuthIdentityRequest) {
    request = &CreateOrganizationMemberAuthIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberAuthIdentity")
    
    
    return
}

func NewCreateOrganizationMemberAuthIdentityResponse() (response *CreateOrganizationMemberAuthIdentityResponse) {
    response = &CreateOrganizationMemberAuthIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMemberAuthIdentity
// 添加组织成员访问授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMemberAuthIdentity(request *CreateOrganizationMemberAuthIdentityRequest) (response *CreateOrganizationMemberAuthIdentityResponse, err error) {
    return c.CreateOrganizationMemberAuthIdentityWithContext(context.Background(), request)
}

// CreateOrganizationMemberAuthIdentity
// 添加组织成员访问授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMemberAuthIdentityWithContext(ctx context.Context, request *CreateOrganizationMemberAuthIdentityRequest) (response *CreateOrganizationMemberAuthIdentityResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberAuthIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMemberAuthIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberAuthIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMemberPolicyRequest() (request *CreateOrganizationMemberPolicyRequest) {
    request = &CreateOrganizationMemberPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberPolicy")
    
    
    return
}

func NewCreateOrganizationMemberPolicyResponse() (response *CreateOrganizationMemberPolicyResponse) {
    response = &CreateOrganizationMemberPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMemberPolicy
// 创建组织成员访问授权策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberPolicy(request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
    return c.CreateOrganizationMemberPolicyWithContext(context.Background(), request)
}

// CreateOrganizationMemberPolicy
// 创建组织成员访问授权策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) CreateOrganizationMemberPolicyWithContext(ctx context.Context, request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMemberPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMemberPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMemberPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationMembersPolicyRequest() (request *CreateOrganizationMembersPolicyRequest) {
    request = &CreateOrganizationMembersPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMembersPolicy")
    
    
    return
}

func NewCreateOrganizationMembersPolicyResponse() (response *CreateOrganizationMembersPolicyResponse) {
    response = &CreateOrganizationMembersPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationMembersPolicy
// 创建组织成员访问策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMembersPolicy(request *CreateOrganizationMembersPolicyRequest) (response *CreateOrganizationMembersPolicyResponse, err error) {
    return c.CreateOrganizationMembersPolicyWithContext(context.Background(), request)
}

// CreateOrganizationMembersPolicy
// 创建组织成员访问策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMembersPolicyWithContext(ctx context.Context, request *CreateOrganizationMembersPolicyRequest) (response *CreateOrganizationMembersPolicyResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationMembersPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationMembersPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationMembersPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
    request = &CreatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreatePolicy")
    
    
    return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
    response = &CreatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePolicy
// 创建一个特殊类型的策略，您可以关联到企业组织Root节点、企业部门节点或者企业的成员账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//  LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    return c.CreatePolicyWithContext(context.Background(), request)
}

// CreatePolicy
// 创建一个特殊类型的策略，您可以关联到企业组织Root节点、企业部门节点或者企业的成员账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//  LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePolicyWithContext(ctx context.Context, request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
    if request == nil {
        request = NewCreatePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteAccount")
    
    
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccount
// 删除成员账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_DELETEACCOUNTDISABLED = "UnsupportedOperation.DeleteAccountDisabled"
//  UNSUPPORTEDOPERATION_INVITEACCOUNTNOTALLOWDELETE = "UnsupportedOperation.InviteAccountNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTEXISTRESOURCE = "UnsupportedOperation.MemberAccountExistResource"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    return c.DeleteAccountWithContext(context.Background(), request)
}

// DeleteAccount
// 删除成员账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_DELETEACCOUNTDISABLED = "UnsupportedOperation.DeleteAccountDisabled"
//  UNSUPPORTEDOPERATION_INVITEACCOUNTNOTALLOWDELETE = "UnsupportedOperation.InviteAccountNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTEXISTRESOURCE = "UnsupportedOperation.MemberAccountExistResource"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
func (c *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    if request == nil {
        request = NewDeleteAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
    request = &DeleteOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganization")
    
    
    return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
    response = &DeleteOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganization
// 删除企业组织
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    return c.DeleteOrganizationWithContext(context.Background(), request)
}

// DeleteOrganization
// 删除企业组织
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationWithContext(ctx context.Context, request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationIdentityRequest() (request *DeleteOrganizationIdentityRequest) {
    request = &DeleteOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationIdentity")
    
    
    return
}

func NewDeleteOrganizationIdentityResponse() (response *DeleteOrganizationIdentityResponse) {
    response = &DeleteOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationIdentity
// 删除组织身份
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationIdentity(request *DeleteOrganizationIdentityRequest) (response *DeleteOrganizationIdentityResponse, err error) {
    return c.DeleteOrganizationIdentityWithContext(context.Background(), request)
}

// DeleteOrganizationIdentity
// 删除组织身份
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationIdentityWithContext(ctx context.Context, request *DeleteOrganizationIdentityRequest) (response *DeleteOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMemberAuthIdentityRequest() (request *DeleteOrganizationMemberAuthIdentityRequest) {
    request = &DeleteOrganizationMemberAuthIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMemberAuthIdentity")
    
    
    return
}

func NewDeleteOrganizationMemberAuthIdentityResponse() (response *DeleteOrganizationMemberAuthIdentityResponse) {
    response = &DeleteOrganizationMemberAuthIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMemberAuthIdentity
// 删除组织成员访问授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberAuthIdentity(request *DeleteOrganizationMemberAuthIdentityRequest) (response *DeleteOrganizationMemberAuthIdentityResponse, err error) {
    return c.DeleteOrganizationMemberAuthIdentityWithContext(context.Background(), request)
}

// DeleteOrganizationMemberAuthIdentity
// 删除组织成员访问授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberAuthIdentityWithContext(ctx context.Context, request *DeleteOrganizationMemberAuthIdentityRequest) (response *DeleteOrganizationMemberAuthIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMemberAuthIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMemberAuthIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMemberAuthIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMembersRequest() (request *DeleteOrganizationMembersRequest) {
    request = &DeleteOrganizationMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembers")
    
    
    return
}

func NewDeleteOrganizationMembersResponse() (response *DeleteOrganizationMembersResponse) {
    response = &DeleteOrganizationMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMembers
// 批量删除企业组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    return c.DeleteOrganizationMembersWithContext(context.Background(), request)
}

// DeleteOrganizationMembers
// 批量删除企业组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembersWithContext(ctx context.Context, request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationMembersPolicyRequest() (request *DeleteOrganizationMembersPolicyRequest) {
    request = &DeleteOrganizationMembersPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembersPolicy")
    
    
    return
}

func NewDeleteOrganizationMembersPolicyResponse() (response *DeleteOrganizationMembersPolicyResponse) {
    response = &DeleteOrganizationMembersPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationMembersPolicy
// 删除组织成员访问策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMembersPolicy(request *DeleteOrganizationMembersPolicyRequest) (response *DeleteOrganizationMembersPolicyResponse, err error) {
    return c.DeleteOrganizationMembersPolicyWithContext(context.Background(), request)
}

// DeleteOrganizationMembersPolicy
// 删除组织成员访问策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMembersPolicyWithContext(ctx context.Context, request *DeleteOrganizationMembersPolicyRequest) (response *DeleteOrganizationMembersPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationMembersPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationMembersPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationMembersPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrganizationNodesRequest() (request *DeleteOrganizationNodesRequest) {
    request = &DeleteOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodes")
    
    
    return
}

func NewDeleteOrganizationNodesResponse() (response *DeleteOrganizationNodesResponse) {
    response = &DeleteOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrganizationNodes
// 批量删除企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodes(request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    return c.DeleteOrganizationNodesWithContext(context.Background(), request)
}

// DeleteOrganizationNodes
// 批量删除企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_NODENOTEMPTY = "FailedOperation.NodeNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONNODEDELETEOVERLIMIT = "FailedOperation.OrganizationNodeDeleteOverLimit"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEMPTY = "FailedOperation.OrganizationNodeNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationNodesWithContext(ctx context.Context, request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewDeleteOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
    request = &DeletePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeletePolicy")
    
    
    return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
    response = &DeletePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePolicy
// 删除策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORGANIZATIONPOLICYINUSED = "FailedOperation.OrganizationPolicyInUsed"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    return c.DeletePolicyWithContext(context.Background(), request)
}

// DeletePolicy
// 删除策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORGANIZATIONPOLICYINUSED = "FailedOperation.OrganizationPolicyInUsed"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
    if request == nil {
        request = NewDeletePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareUnitRequest() (request *DeleteShareUnitRequest) {
    request = &DeleteShareUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteShareUnit")
    
    
    return
}

func NewDeleteShareUnitResponse() (response *DeleteShareUnitResponse) {
    response = &DeleteShareUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShareUnit
// 删除共享单元。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnit(request *DeleteShareUnitRequest) (response *DeleteShareUnitResponse, err error) {
    return c.DeleteShareUnitWithContext(context.Background(), request)
}

// DeleteShareUnit
// 删除共享单元。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitWithContext(ctx context.Context, request *DeleteShareUnitRequest) (response *DeleteShareUnitResponse, err error) {
    if request == nil {
        request = NewDeleteShareUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareUnitMembersRequest() (request *DeleteShareUnitMembersRequest) {
    request = &DeleteShareUnitMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteShareUnitMembers")
    
    
    return
}

func NewDeleteShareUnitMembersResponse() (response *DeleteShareUnitMembersResponse) {
    response = &DeleteShareUnitMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShareUnitMembers
// 删除共享单元成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitMembers(request *DeleteShareUnitMembersRequest) (response *DeleteShareUnitMembersResponse, err error) {
    return c.DeleteShareUnitMembersWithContext(context.Background(), request)
}

// DeleteShareUnitMembers
// 删除共享单元成员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITMEMBEROVERLIMIT = "LimitExceeded.ShareUnitMemberOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitMembersWithContext(ctx context.Context, request *DeleteShareUnitMembersRequest) (response *DeleteShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewDeleteShareUnitMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnitMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareUnitResourcesRequest() (request *DeleteShareUnitResourcesRequest) {
    request = &DeleteShareUnitResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteShareUnitResources")
    
    
    return
}

func NewDeleteShareUnitResourcesResponse() (response *DeleteShareUnitResourcesResponse) {
    response = &DeleteShareUnitResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteShareUnitResources
// 删除共享单元资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitResources(request *DeleteShareUnitResourcesRequest) (response *DeleteShareUnitResourcesResponse, err error) {
    return c.DeleteShareUnitResourcesWithContext(context.Background(), request)
}

// DeleteShareUnitResources
// 删除共享单元资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SHAREAREANOTEXIST = "FailedOperation.ShareAreaNotExist"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  FAILEDOPERATION_SHARERESOURCENOTEXIST = "FailedOperation.ShareResourceNotExist"
//  FAILEDOPERATION_SHARERESOURCETYPENOTEXIST = "FailedOperation.ShareResourceTypeNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SHAREUNITRESOURCEOVERLIMIT = "LimitExceeded.ShareUnitResourceOverLimit"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteShareUnitResourcesWithContext(ctx context.Context, request *DeleteShareUnitResourcesRequest) (response *DeleteShareUnitResourcesResponse, err error) {
    if request == nil {
        request = NewDeleteShareUnitResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEffectivePolicyRequest() (request *DescribeEffectivePolicyRequest) {
    request = &DescribeEffectivePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeEffectivePolicy")
    
    
    return
}

func NewDescribeEffectivePolicyResponse() (response *DescribeEffectivePolicyResponse) {
    response = &DescribeEffectivePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEffectivePolicy
// 查询目标关联的有效策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_EFFECTIVEPOLICYNOTFOUND = "ResourceNotFound.EffectivePolicyNotFound"
func (c *Client) DescribeEffectivePolicy(request *DescribeEffectivePolicyRequest) (response *DescribeEffectivePolicyResponse, err error) {
    return c.DescribeEffectivePolicyWithContext(context.Background(), request)
}

// DescribeEffectivePolicy
// 查询目标关联的有效策略
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_EFFECTIVEPOLICYNOTFOUND = "ResourceNotFound.EffectivePolicyNotFound"
func (c *Client) DescribeEffectivePolicyWithContext(ctx context.Context, request *DescribeEffectivePolicyRequest) (response *DescribeEffectivePolicyResponse, err error) {
    if request == nil {
        request = NewDescribeEffectivePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEffectivePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEffectivePolicyResponse()
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
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    return c.DescribeOrganizationWithContext(context.Background(), request)
}

// DescribeOrganization
// 获取企业组织信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationWithContext(ctx context.Context, request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationAuthNodeRequest() (request *DescribeOrganizationAuthNodeRequest) {
    request = &DescribeOrganizationAuthNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationAuthNode")
    
    
    return
}

func NewDescribeOrganizationAuthNodeResponse() (response *DescribeOrganizationAuthNodeResponse) {
    response = &DescribeOrganizationAuthNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationAuthNode
// 获取已设置管理员的互信主体关系列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationAuthNode(request *DescribeOrganizationAuthNodeRequest) (response *DescribeOrganizationAuthNodeResponse, err error) {
    return c.DescribeOrganizationAuthNodeWithContext(context.Background(), request)
}

// DescribeOrganizationAuthNode
// 获取已设置管理员的互信主体关系列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationAuthNodeWithContext(ctx context.Context, request *DescribeOrganizationAuthNodeRequest) (response *DescribeOrganizationAuthNodeResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationAuthNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationAuthNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationAuthNodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationFinancialByMemberRequest() (request *DescribeOrganizationFinancialByMemberRequest) {
    request = &DescribeOrganizationFinancialByMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialByMember")
    
    
    return
}

func NewDescribeOrganizationFinancialByMemberResponse() (response *DescribeOrganizationFinancialByMemberResponse) {
    response = &DescribeOrganizationFinancialByMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationFinancialByMember
// 以成员维度获取组织财务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationFinancialByMember(request *DescribeOrganizationFinancialByMemberRequest) (response *DescribeOrganizationFinancialByMemberResponse, err error) {
    return c.DescribeOrganizationFinancialByMemberWithContext(context.Background(), request)
}

// DescribeOrganizationFinancialByMember
// 以成员维度获取组织财务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationFinancialByMemberWithContext(ctx context.Context, request *DescribeOrganizationFinancialByMemberRequest) (response *DescribeOrganizationFinancialByMemberResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationFinancialByMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationFinancialByMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationFinancialByMemberResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationFinancialByMonthRequest() (request *DescribeOrganizationFinancialByMonthRequest) {
    request = &DescribeOrganizationFinancialByMonthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialByMonth")
    
    
    return
}

func NewDescribeOrganizationFinancialByMonthResponse() (response *DescribeOrganizationFinancialByMonthResponse) {
    response = &DescribeOrganizationFinancialByMonthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationFinancialByMonth
// 以月维度获取组织财务信息趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationFinancialByMonth(request *DescribeOrganizationFinancialByMonthRequest) (response *DescribeOrganizationFinancialByMonthResponse, err error) {
    return c.DescribeOrganizationFinancialByMonthWithContext(context.Background(), request)
}

// DescribeOrganizationFinancialByMonth
// 以月维度获取组织财务信息趋势
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationFinancialByMonthWithContext(ctx context.Context, request *DescribeOrganizationFinancialByMonthRequest) (response *DescribeOrganizationFinancialByMonthResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationFinancialByMonthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationFinancialByMonth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationFinancialByMonthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationFinancialByProductRequest() (request *DescribeOrganizationFinancialByProductRequest) {
    request = &DescribeOrganizationFinancialByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialByProduct")
    
    
    return
}

func NewDescribeOrganizationFinancialByProductResponse() (response *DescribeOrganizationFinancialByProductResponse) {
    response = &DescribeOrganizationFinancialByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationFinancialByProduct
// 以产品维度获取组织财务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationFinancialByProduct(request *DescribeOrganizationFinancialByProductRequest) (response *DescribeOrganizationFinancialByProductResponse, err error) {
    return c.DescribeOrganizationFinancialByProductWithContext(context.Background(), request)
}

// DescribeOrganizationFinancialByProduct
// 以产品维度获取组织财务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeOrganizationFinancialByProductWithContext(ctx context.Context, request *DescribeOrganizationFinancialByProductRequest) (response *DescribeOrganizationFinancialByProductResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationFinancialByProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationFinancialByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationFinancialByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberAuthAccountsRequest() (request *DescribeOrganizationMemberAuthAccountsRequest) {
    request = &DescribeOrganizationMemberAuthAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
    
    
    return
}

func NewDescribeOrganizationMemberAuthAccountsResponse() (response *DescribeOrganizationMemberAuthAccountsResponse) {
    response = &DescribeOrganizationMemberAuthAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberAuthAccounts
// 获取组织成员被绑定授权关系的子账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthAccounts(request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
    return c.DescribeOrganizationMemberAuthAccountsWithContext(context.Background(), request)
}

// DescribeOrganizationMemberAuthAccounts
// 获取组织成员被绑定授权关系的子账号列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthAccountsWithContext(ctx context.Context, request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberAuthAccountsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberAuthIdentitiesRequest() (request *DescribeOrganizationMemberAuthIdentitiesRequest) {
    request = &DescribeOrganizationMemberAuthIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthIdentities")
    
    
    return
}

func NewDescribeOrganizationMemberAuthIdentitiesResponse() (response *DescribeOrganizationMemberAuthIdentitiesResponse) {
    response = &DescribeOrganizationMemberAuthIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberAuthIdentities
// 获取组织成员访问授权列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentities(request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
    return c.DescribeOrganizationMemberAuthIdentitiesWithContext(context.Background(), request)
}

// DescribeOrganizationMemberAuthIdentities
// 获取组织成员访问授权列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentitiesWithContext(ctx context.Context, request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberAuthIdentitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberAuthIdentities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberAuthIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberEmailBindRequest() (request *DescribeOrganizationMemberEmailBindRequest) {
    request = &DescribeOrganizationMemberEmailBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberEmailBind")
    
    
    return
}

func NewDescribeOrganizationMemberEmailBindResponse() (response *DescribeOrganizationMemberEmailBindResponse) {
    response = &DescribeOrganizationMemberEmailBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberEmailBind
// 查询成员邮箱绑定详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"
//  FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"
//  FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"
//  FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"
//  INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"
//  INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"
//  INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeOrganizationMemberEmailBind(request *DescribeOrganizationMemberEmailBindRequest) (response *DescribeOrganizationMemberEmailBindResponse, err error) {
    return c.DescribeOrganizationMemberEmailBindWithContext(context.Background(), request)
}

// DescribeOrganizationMemberEmailBind
// 查询成员邮箱绑定详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"
//  FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"
//  FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"
//  FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"
//  INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"
//  INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"
//  INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeOrganizationMemberEmailBindWithContext(ctx context.Context, request *DescribeOrganizationMemberEmailBindRequest) (response *DescribeOrganizationMemberEmailBindResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberEmailBindRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberEmailBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberEmailBindResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMemberPoliciesRequest() (request *DescribeOrganizationMemberPoliciesRequest) {
    request = &DescribeOrganizationMemberPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberPolicies")
    
    
    return
}

func NewDescribeOrganizationMemberPoliciesResponse() (response *DescribeOrganizationMemberPoliciesResponse) {
    response = &DescribeOrganizationMemberPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMemberPolicies
// 获取组织成员的授权策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberPolicies(request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
    return c.DescribeOrganizationMemberPoliciesWithContext(context.Background(), request)
}

// DescribeOrganizationMemberPolicies
// 获取组织成员的授权策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberPoliciesWithContext(ctx context.Context, request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMemberPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMemberPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMemberPoliciesResponse()
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
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    return c.DescribeOrganizationMembersWithContext(context.Background(), request)
}

// DescribeOrganizationMembers
// 获取企业组织成员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DescribeOrganizationMembersWithContext(ctx context.Context, request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationNodesRequest() (request *DescribeOrganizationNodesRequest) {
    request = &DescribeOrganizationNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodes")
    
    
    return
}

func NewDescribeOrganizationNodesResponse() (response *DescribeOrganizationNodesResponse) {
    response = &DescribeOrganizationNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationNodes
// 获取组织节点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationNodes(request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
    return c.DescribeOrganizationNodesWithContext(context.Background(), request)
}

// DescribeOrganizationNodes
// 获取组织节点列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationNodesWithContext(ctx context.Context, request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyRequest() (request *DescribePolicyRequest) {
    request = &DescribePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribePolicy")
    
    
    return
}

func NewDescribePolicyResponse() (response *DescribePolicyResponse) {
    response = &DescribePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicy
// 本接口（DescribePolicy）可用于查询查看策略详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) DescribePolicy(request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
    return c.DescribePolicyWithContext(context.Background(), request)
}

// DescribePolicy
// 本接口（DescribePolicy）可用于查询查看策略详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) DescribePolicyWithContext(ctx context.Context, request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
    if request == nil {
        request = NewDescribePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePolicyConfigRequest() (request *DescribePolicyConfigRequest) {
    request = &DescribePolicyConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribePolicyConfig")
    
    
    return
}

func NewDescribePolicyConfigResponse() (response *DescribePolicyConfigResponse) {
    response = &DescribePolicyConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePolicyConfig
// 本接口（DescribePolicyConfig）可用于查询企业组织策略配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribePolicyConfig(request *DescribePolicyConfigRequest) (response *DescribePolicyConfigResponse, err error) {
    return c.DescribePolicyConfigWithContext(context.Background(), request)
}

// DescribePolicyConfig
// 本接口（DescribePolicyConfig）可用于查询企业组织策略配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribePolicyConfigWithContext(ctx context.Context, request *DescribePolicyConfigRequest) (response *DescribePolicyConfigResponse, err error) {
    if request == nil {
        request = NewDescribePolicyConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareAreasRequest() (request *DescribeShareAreasRequest) {
    request = &DescribeShareAreasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareAreas")
    
    
    return
}

func NewDescribeShareAreasResponse() (response *DescribeShareAreasResponse) {
    response = &DescribeShareAreasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareAreas
// 获取可共享地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareAreas(request *DescribeShareAreasRequest) (response *DescribeShareAreasResponse, err error) {
    return c.DescribeShareAreasWithContext(context.Background(), request)
}

// DescribeShareAreas
// 获取可共享地域列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareAreasWithContext(ctx context.Context, request *DescribeShareAreasRequest) (response *DescribeShareAreasResponse, err error) {
    if request == nil {
        request = NewDescribeShareAreasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareAreas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareAreasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareUnitMembersRequest() (request *DescribeShareUnitMembersRequest) {
    request = &DescribeShareUnitMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareUnitMembers")
    
    
    return
}

func NewDescribeShareUnitMembersResponse() (response *DescribeShareUnitMembersResponse) {
    response = &DescribeShareUnitMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareUnitMembers
// 获取共享单元成员列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitMembers(request *DescribeShareUnitMembersRequest) (response *DescribeShareUnitMembersResponse, err error) {
    return c.DescribeShareUnitMembersWithContext(context.Background(), request)
}

// DescribeShareUnitMembers
// 获取共享单元成员列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitMembersWithContext(ctx context.Context, request *DescribeShareUnitMembersRequest) (response *DescribeShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewDescribeShareUnitMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareUnitMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareUnitMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareUnitResourcesRequest() (request *DescribeShareUnitResourcesRequest) {
    request = &DescribeShareUnitResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareUnitResources")
    
    
    return
}

func NewDescribeShareUnitResourcesResponse() (response *DescribeShareUnitResourcesResponse) {
    response = &DescribeShareUnitResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareUnitResources
// 获取共享单元资源列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitResources(request *DescribeShareUnitResourcesRequest) (response *DescribeShareUnitResourcesResponse, err error) {
    return c.DescribeShareUnitResourcesWithContext(context.Background(), request)
}

// DescribeShareUnitResources
// 获取共享单元资源列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitResourcesWithContext(ctx context.Context, request *DescribeShareUnitResourcesRequest) (response *DescribeShareUnitResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeShareUnitResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareUnitResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShareUnitsRequest() (request *DescribeShareUnitsRequest) {
    request = &DescribeShareUnitsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeShareUnits")
    
    
    return
}

func NewDescribeShareUnitsResponse() (response *DescribeShareUnitsResponse) {
    response = &DescribeShareUnitsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeShareUnits
// 获取共享单元列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnits(request *DescribeShareUnitsRequest) (response *DescribeShareUnitsResponse, err error) {
    return c.DescribeShareUnitsWithContext(context.Background(), request)
}

// DescribeShareUnits
// 获取共享单元列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeShareUnitsWithContext(ctx context.Context, request *DescribeShareUnitsRequest) (response *DescribeShareUnitsResponse, err error) {
    if request == nil {
        request = NewDescribeShareUnitsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShareUnits require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShareUnitsResponse()
    err = c.Send(request, response)
    return
}

func NewDetachPolicyRequest() (request *DetachPolicyRequest) {
    request = &DetachPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DetachPolicy")
    
    
    return
}

func NewDetachPolicyResponse() (response *DetachPolicyResponse) {
    response = &DetachPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachPolicy
// 解绑策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONDETACHLASTPOLICYERROR = "FailedOperation.OrganizationDetachLastPolicyError"
//  FAILEDOPERATION_ORGANIZATIONDETACHPOLICYERROR = "FailedOperation.OrganizationDetachPolicyError"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DetachPolicy(request *DetachPolicyRequest) (response *DetachPolicyResponse, err error) {
    return c.DetachPolicyWithContext(context.Background(), request)
}

// DetachPolicy
// 解绑策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONDETACHLASTPOLICYERROR = "FailedOperation.OrganizationDetachLastPolicyError"
//  FAILEDOPERATION_ORGANIZATIONDETACHPOLICYERROR = "FailedOperation.OrganizationDetachPolicyError"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DetachPolicyWithContext(ctx context.Context, request *DetachPolicyRequest) (response *DetachPolicyResponse, err error) {
    if request == nil {
        request = NewDetachPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDisablePolicyTypeRequest() (request *DisablePolicyTypeRequest) {
    request = &DisablePolicyTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DisablePolicyType")
    
    
    return
}

func NewDisablePolicyTypeResponse() (response *DisablePolicyTypeResponse) {
    response = &DisablePolicyTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisablePolicyType
// 禁用策略类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisablePolicyType(request *DisablePolicyTypeRequest) (response *DisablePolicyTypeResponse, err error) {
    return c.DisablePolicyTypeWithContext(context.Background(), request)
}

// DisablePolicyType
// 禁用策略类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisablePolicyTypeWithContext(ctx context.Context, request *DisablePolicyTypeRequest) (response *DisablePolicyTypeResponse, err error) {
    if request == nil {
        request = NewDisablePolicyTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisablePolicyType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisablePolicyTypeResponse()
    err = c.Send(request, response)
    return
}

func NewEnablePolicyTypeRequest() (request *EnablePolicyTypeRequest) {
    request = &EnablePolicyTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "EnablePolicyType")
    
    
    return
}

func NewEnablePolicyTypeResponse() (response *EnablePolicyTypeResponse) {
    response = &EnablePolicyTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnablePolicyType
// 启用策略类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_POLICYENABLEINVALID = "FailedOperation.PolicyEnableInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnablePolicyType(request *EnablePolicyTypeRequest) (response *EnablePolicyTypeResponse, err error) {
    return c.EnablePolicyTypeWithContext(context.Background(), request)
}

// EnablePolicyType
// 启用策略类型
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_POLICYENABLEINVALID = "FailedOperation.PolicyEnableInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnablePolicyTypeWithContext(ctx context.Context, request *EnablePolicyTypeRequest) (response *EnablePolicyTypeResponse, err error) {
    if request == nil {
        request = NewEnablePolicyTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnablePolicyType require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnablePolicyTypeResponse()
    err = c.Send(request, response)
    return
}

func NewListNonCompliantResourceRequest() (request *ListNonCompliantResourceRequest) {
    request = &ListNonCompliantResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListNonCompliantResource")
    
    
    return
}

func NewListNonCompliantResourceResponse() (response *ListNonCompliantResourceResponse) {
    response = &ListNonCompliantResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListNonCompliantResource
// 获取成员标签检测不合规资源列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListNonCompliantResource(request *ListNonCompliantResourceRequest) (response *ListNonCompliantResourceResponse, err error) {
    return c.ListNonCompliantResourceWithContext(context.Background(), request)
}

// ListNonCompliantResource
// 获取成员标签检测不合规资源列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListNonCompliantResourceWithContext(ctx context.Context, request *ListNonCompliantResourceRequest) (response *ListNonCompliantResourceResponse, err error) {
    if request == nil {
        request = NewListNonCompliantResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListNonCompliantResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewListNonCompliantResourceResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationIdentityRequest() (request *ListOrganizationIdentityRequest) {
    request = &ListOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationIdentity")
    
    
    return
}

func NewListOrganizationIdentityResponse() (response *ListOrganizationIdentityResponse) {
    response = &ListOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationIdentity
// 获取组织成员访问身份列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentity(request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
    return c.ListOrganizationIdentityWithContext(context.Background(), request)
}

// ListOrganizationIdentity
// 获取组织成员访问身份列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentityWithContext(ctx context.Context, request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewListOrganizationIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
    request = &ListPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListPolicies")
    
    
    return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
    response = &ListPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListPolicies
// 本接口（ListPolicies）可用于查询查看策略列表数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    return c.ListPoliciesWithContext(context.Background(), request)
}

// ListPolicies
// 本接口（ListPolicies）可用于查询查看策略列表数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
    if request == nil {
        request = NewListPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewListPoliciesForTargetRequest() (request *ListPoliciesForTargetRequest) {
    request = &ListPoliciesForTargetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListPoliciesForTarget")
    
    
    return
}

func NewListPoliciesForTargetResponse() (response *ListPoliciesForTargetResponse) {
    response = &ListPoliciesForTargetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListPoliciesForTarget
// 本接口（ListPoliciesForTarget）查询目标关联的策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListPoliciesForTarget(request *ListPoliciesForTargetRequest) (response *ListPoliciesForTargetResponse, err error) {
    return c.ListPoliciesForTargetWithContext(context.Background(), request)
}

// ListPoliciesForTarget
// 本接口（ListPoliciesForTarget）查询目标关联的策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListPoliciesForTargetWithContext(ctx context.Context, request *ListPoliciesForTargetRequest) (response *ListPoliciesForTargetResponse, err error) {
    if request == nil {
        request = NewListPoliciesForTargetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPoliciesForTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPoliciesForTargetResponse()
    err = c.Send(request, response)
    return
}

func NewListTargetsForPolicyRequest() (request *ListTargetsForPolicyRequest) {
    request = &ListTargetsForPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListTargetsForPolicy")
    
    
    return
}

func NewListTargetsForPolicyResponse() (response *ListTargetsForPolicyResponse) {
    response = &ListTargetsForPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTargetsForPolicy
// 本接口（ListTargetsForPolicy）查询某个指定策略关联的目标列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
func (c *Client) ListTargetsForPolicy(request *ListTargetsForPolicyRequest) (response *ListTargetsForPolicyResponse, err error) {
    return c.ListTargetsForPolicyWithContext(context.Background(), request)
}

// ListTargetsForPolicy
// 本接口（ListTargetsForPolicy）查询某个指定策略关联的目标列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
func (c *Client) ListTargetsForPolicyWithContext(ctx context.Context, request *ListTargetsForPolicyRequest) (response *ListTargetsForPolicyResponse, err error) {
    if request == nil {
        request = NewListTargetsForPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTargetsForPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTargetsForPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewMoveOrganizationNodeMembersRequest() (request *MoveOrganizationNodeMembersRequest) {
    request = &MoveOrganizationNodeMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNodeMembers")
    
    
    return
}

func NewMoveOrganizationNodeMembersResponse() (response *MoveOrganizationNodeMembersResponse) {
    response = &MoveOrganizationNodeMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MoveOrganizationNodeMembers
// 移动成员到指定企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationNodeMembers(request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
    return c.MoveOrganizationNodeMembersWithContext(context.Background(), request)
}

// MoveOrganizationNodeMembers
// 移动成员到指定企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_SOMEUINSNOTINORGANIZATION = "FailedOperation.SomeUinsNotInOrganization"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) MoveOrganizationNodeMembersWithContext(ctx context.Context, request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
    if request == nil {
        request = NewMoveOrganizationNodeMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveOrganizationNodeMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveOrganizationNodeMembersResponse()
    err = c.Send(request, response)
    return
}

func NewQuitOrganizationRequest() (request *QuitOrganizationRequest) {
    request = &QuitOrganizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "QuitOrganization")
    
    
    return
}

func NewQuitOrganizationResponse() (response *QuitOrganizationResponse) {
    response = &QuitOrganizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QuitOrganization
// 退出企业组织
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
func (c *Client) QuitOrganization(request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
    return c.QuitOrganizationWithContext(context.Background(), request)
}

// QuitOrganization
// 退出企业组织
//
// 可能返回的错误码:
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
func (c *Client) QuitOrganizationWithContext(ctx context.Context, request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
    if request == nil {
        request = NewQuitOrganizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuitOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewQuitOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationIdentityRequest() (request *UpdateOrganizationIdentityRequest) {
    request = &UpdateOrganizationIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationIdentity")
    
    
    return
}

func NewUpdateOrganizationIdentityResponse() (response *UpdateOrganizationIdentityResponse) {
    response = &UpdateOrganizationIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationIdentity
// 更新组织身份
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) UpdateOrganizationIdentity(request *UpdateOrganizationIdentityRequest) (response *UpdateOrganizationIdentityResponse, err error) {
    return c.UpdateOrganizationIdentityWithContext(context.Background(), request)
}

// UpdateOrganizationIdentity
// 更新组织身份
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) UpdateOrganizationIdentityWithContext(ctx context.Context, request *UpdateOrganizationIdentityRequest) (response *UpdateOrganizationIdentityResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationIdentityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationMemberRequest() (request *UpdateOrganizationMemberRequest) {
    request = &UpdateOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMember")
    
    
    return
}

func NewUpdateOrganizationMemberResponse() (response *UpdateOrganizationMemberResponse) {
    response = &UpdateOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationMember
// 更新组织成员信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CHANGEPERMISSIONRECORDEXIST = "FailedOperation.ChangePermissionRecordExist"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INVALIDPARAMETER_ALLOWQUITILLEGAL = "InvalidParameter.AllowQuitIllegal"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_DELETEDELEGATEPAYERNOTALLOW = "UnsupportedOperation.DeleteDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
func (c *Client) UpdateOrganizationMember(request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
    return c.UpdateOrganizationMemberWithContext(context.Background(), request)
}

// UpdateOrganizationMember
// 更新组织成员信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CHANGEPERMISSIONRECORDEXIST = "FailedOperation.ChangePermissionRecordExist"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  INVALIDPARAMETER_ALLOWQUITILLEGAL = "InvalidParameter.AllowQuitIllegal"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_DELETEDELEGATEPAYERNOTALLOW = "UnsupportedOperation.DeleteDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
func (c *Client) UpdateOrganizationMemberWithContext(ctx context.Context, request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationMemberRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationMemberEmailBindRequest() (request *UpdateOrganizationMemberEmailBindRequest) {
    request = &UpdateOrganizationMemberEmailBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMemberEmailBind")
    
    
    return
}

func NewUpdateOrganizationMemberEmailBindResponse() (response *UpdateOrganizationMemberEmailBindResponse) {
    response = &UpdateOrganizationMemberEmailBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationMemberEmailBind
// 修改绑定成员邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMemberEmailBind(request *UpdateOrganizationMemberEmailBindRequest) (response *UpdateOrganizationMemberEmailBindResponse, err error) {
    return c.UpdateOrganizationMemberEmailBindWithContext(context.Background(), request)
}

// UpdateOrganizationMemberEmailBind
// 修改绑定成员邮箱
//
// 可能返回的错误码:
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMemberEmailBindWithContext(ctx context.Context, request *UpdateOrganizationMemberEmailBindRequest) (response *UpdateOrganizationMemberEmailBindResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationMemberEmailBindRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMemberEmailBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMemberEmailBindResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationNodeRequest() (request *UpdateOrganizationNodeRequest) {
    request = &UpdateOrganizationNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationNode")
    
    
    return
}

func NewUpdateOrganizationNodeResponse() (response *UpdateOrganizationNodeResponse) {
    response = &UpdateOrganizationNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationNode
// 更新企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNode(request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    return c.UpdateOrganizationNodeWithContext(context.Background(), request)
}

// UpdateOrganizationNode
// 更新企业组织节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONNODENAMEUSED = "FailedOperation.OrganizationNodeNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) UpdateOrganizationNodeWithContext(ctx context.Context, request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
    request = &UpdatePolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdatePolicy")
    
    
    return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
    response = &UpdatePolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePolicy
// 编辑策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//  LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    return c.UpdatePolicyWithContext(context.Background(), request)
}

// UpdatePolicy
// 编辑策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//  LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdatePolicyWithContext(ctx context.Context, request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
    if request == nil {
        request = NewUpdatePolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateShareUnitRequest() (request *UpdateShareUnitRequest) {
    request = &UpdateShareUnitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateShareUnit")
    
    
    return
}

func NewUpdateShareUnitResponse() (response *UpdateShareUnitResponse) {
    response = &UpdateShareUnitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateShareUnit
// 更新共享单元。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateShareUnit(request *UpdateShareUnitRequest) (response *UpdateShareUnitResponse, err error) {
    return c.UpdateShareUnitWithContext(context.Background(), request)
}

// UpdateShareUnit
// 更新共享单元。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateShareUnitWithContext(ctx context.Context, request *UpdateShareUnitRequest) (response *UpdateShareUnitResponse, err error) {
    if request == nil {
        request = NewUpdateShareUnitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateShareUnitResponse()
    err = c.Send(request, response)
    return
}
