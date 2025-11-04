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


func NewAcceptJoinShareUnitInvitationRequest() (request *AcceptJoinShareUnitInvitationRequest) {
    request = &AcceptJoinShareUnitInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AcceptJoinShareUnitInvitation")
    
    
    return
}

func NewAcceptJoinShareUnitInvitationResponse() (response *AcceptJoinShareUnitInvitationResponse) {
    response = &AcceptJoinShareUnitInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AcceptJoinShareUnitInvitation
// 接受加入共享单元邀请。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AcceptJoinShareUnitInvitation(request *AcceptJoinShareUnitInvitationRequest) (response *AcceptJoinShareUnitInvitationResponse, err error) {
    return c.AcceptJoinShareUnitInvitationWithContext(context.Background(), request)
}

// AcceptJoinShareUnitInvitation
// 接受加入共享单元邀请。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AcceptJoinShareUnitInvitationWithContext(ctx context.Context, request *AcceptJoinShareUnitInvitationRequest) (response *AcceptJoinShareUnitInvitationResponse, err error) {
    if request == nil {
        request = NewAcceptJoinShareUnitInvitationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AcceptJoinShareUnitInvitation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AcceptJoinShareUnitInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewAcceptJoinShareUnitInvitationResponse()
    err = c.Send(request, response)
    return
}

func NewAddExternalSAMLIdPCertificateRequest() (request *AddExternalSAMLIdPCertificateRequest) {
    request = &AddExternalSAMLIdPCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddExternalSAMLIdPCertificate")
    
    
    return
}

func NewAddExternalSAMLIdPCertificateResponse() (response *AddExternalSAMLIdPCertificateResponse) {
    response = &AddExternalSAMLIdPCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddExternalSAMLIdPCertificate
// 添加SAML签名证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_X509CERTIFICATEALREADYEXIST = "FailedOperation.X509CertificateAlreadyExist"
//  FAILEDOPERATION_X509CERTIFICATELIMITEXCEEDED = "FailedOperation.X509CertificateLimitExceeded"
//  FAILEDOPERATION_X509CERTIFICATEMINIMUMREQUIRED = "FailedOperation.X509CertificateMinimumRequired"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) AddExternalSAMLIdPCertificate(request *AddExternalSAMLIdPCertificateRequest) (response *AddExternalSAMLIdPCertificateResponse, err error) {
    return c.AddExternalSAMLIdPCertificateWithContext(context.Background(), request)
}

// AddExternalSAMLIdPCertificate
// 添加SAML签名证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_X509CERTIFICATEALREADYEXIST = "FailedOperation.X509CertificateAlreadyExist"
//  FAILEDOPERATION_X509CERTIFICATELIMITEXCEEDED = "FailedOperation.X509CertificateLimitExceeded"
//  FAILEDOPERATION_X509CERTIFICATEMINIMUMREQUIRED = "FailedOperation.X509CertificateMinimumRequired"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) AddExternalSAMLIdPCertificateWithContext(ctx context.Context, request *AddExternalSAMLIdPCertificateRequest) (response *AddExternalSAMLIdPCertificateResponse, err error) {
    if request == nil {
        request = NewAddExternalSAMLIdPCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddExternalSAMLIdPCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddExternalSAMLIdPCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddExternalSAMLIdPCertificateResponse()
    err = c.Send(request, response)
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddOrganizationMemberEmail")
    
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_NODEDEPTHEXCEEDLIMIT = "LimitExceeded.NodeDepthExceedLimit"
//  LIMITEXCEEDED_NODEEXCEEDLIMIT = "LimitExceeded.NodeExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) AddOrganizationNodeWithContext(ctx context.Context, request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
    if request == nil {
        request = NewAddOrganizationNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddOrganizationNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddOrganizationNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddOrganizationNodeResponse()
    err = c.Send(request, response)
    return
}

func NewAddPermissionPolicyToRoleConfigurationRequest() (request *AddPermissionPolicyToRoleConfigurationRequest) {
    request = &AddPermissionPolicyToRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddPermissionPolicyToRoleConfiguration")
    
    
    return
}

func NewAddPermissionPolicyToRoleConfigurationResponse() (response *AddPermissionPolicyToRoleConfigurationResponse) {
    response = &AddPermissionPolicyToRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddPermissionPolicyToRoleConfiguration
// 为权限配置添加策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_CUSTOMPOLICYOVERUPPERLIMIT = "FailedOperation.CustomPolicyOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYSTEMPOLICYOVERUPPERLIMIT = "FailedOperation.SystemPolicyOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_BINDPOLICYNAMENOTALLOWED = "InvalidParameter.BindPolicyNameNotAllowed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYNAMEALREADYEXISTS = "InvalidParameter.PolicyNameAlreadyExists"
//  INVALIDPARAMETER_POLICYNAMESIZEOVERUPPERLIMIT = "InvalidParameter.PolicyNameSizeOverUpperLimit"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) AddPermissionPolicyToRoleConfiguration(request *AddPermissionPolicyToRoleConfigurationRequest) (response *AddPermissionPolicyToRoleConfigurationResponse, err error) {
    return c.AddPermissionPolicyToRoleConfigurationWithContext(context.Background(), request)
}

// AddPermissionPolicyToRoleConfiguration
// 为权限配置添加策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_CUSTOMPOLICYOVERUPPERLIMIT = "FailedOperation.CustomPolicyOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_SYSTEMPOLICYOVERUPPERLIMIT = "FailedOperation.SystemPolicyOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_BINDPOLICYNAMENOTALLOWED = "InvalidParameter.BindPolicyNameNotAllowed"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYNAMEALREADYEXISTS = "InvalidParameter.PolicyNameAlreadyExists"
//  INVALIDPARAMETER_POLICYNAMESIZEOVERUPPERLIMIT = "InvalidParameter.PolicyNameSizeOverUpperLimit"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) AddPermissionPolicyToRoleConfigurationWithContext(ctx context.Context, request *AddPermissionPolicyToRoleConfigurationRequest) (response *AddPermissionPolicyToRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewAddPermissionPolicyToRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddPermissionPolicyToRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddPermissionPolicyToRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddPermissionPolicyToRoleConfigurationResponse()
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
// 创建共享单元。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
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
// 创建共享单元。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddShareUnit")
    
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
//  UNSUPPORTEDOPERATION_SHARINGTOOTHERORGANIZATIONMEMBER = "UnsupportedOperation.SharingToOtherOrganizationMember"
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
//  UNSUPPORTEDOPERATION_SHARINGTOOTHERORGANIZATIONMEMBER = "UnsupportedOperation.SharingToOtherOrganizationMember"
func (c *Client) AddShareUnitMembersWithContext(ctx context.Context, request *AddShareUnitMembersRequest) (response *AddShareUnitMembersResponse, err error) {
    if request == nil {
        request = NewAddShareUnitMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddShareUnitMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddShareUnitResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddShareUnitResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
    request = &AddUserToGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "AddUserToGroup")
    
    
    return
}

func NewAddUserToGroupResponse() (response *AddUserToGroupResponse) {
    response = &AddUserToGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserToGroup
// 为用户组添加用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPTYPEUSERTYPENOTMATCH = "FailedOperation.GroupTypeUserTypeNotMatch"
//  FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_USERADDGROUPCOUNTOVERUPPERLIMIT = "FailedOperation.UserAddGroupCountOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERALREADYEXISTS = "InvalidParameter.GroupUserAlreadyExists"
//  LIMITEXCEEDED_ADDUSERTOGROUPLIMITEXCEEDED = "LimitExceeded.AddUserToGroupLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    return c.AddUserToGroupWithContext(context.Background(), request)
}

// AddUserToGroup
// 为用户组添加用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPTYPEUSERTYPENOTMATCH = "FailedOperation.GroupTypeUserTypeNotMatch"
//  FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_USERADDGROUPCOUNTOVERUPPERLIMIT = "FailedOperation.UserAddGroupCountOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERALREADYEXISTS = "InvalidParameter.GroupUserAlreadyExists"
//  LIMITEXCEEDED_ADDUSERTOGROUPLIMITEXCEEDED = "LimitExceeded.AddUserToGroupLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddUserToGroupWithContext(ctx context.Context, request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
    if request == nil {
        request = NewAddUserToGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AddUserToGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserToGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserToGroupResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "AttachPolicy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "BindOrganizationMemberAuthAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewBindOrganizationPolicySubAccountRequest() (request *BindOrganizationPolicySubAccountRequest) {
    request = &BindOrganizationPolicySubAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationPolicySubAccount")
    
    
    return
}

func NewBindOrganizationPolicySubAccountResponse() (response *BindOrganizationPolicySubAccountResponse) {
    response = &BindOrganizationPolicySubAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindOrganizationPolicySubAccount
// 绑定成员访问授权策略和组织管理员子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationPolicySubAccount(request *BindOrganizationPolicySubAccountRequest) (response *BindOrganizationPolicySubAccountResponse, err error) {
    return c.BindOrganizationPolicySubAccountWithContext(context.Background(), request)
}

// BindOrganizationPolicySubAccount
// 绑定成员访问授权策略和组织管理员子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) BindOrganizationPolicySubAccountWithContext(ctx context.Context, request *BindOrganizationPolicySubAccountRequest) (response *BindOrganizationPolicySubAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationPolicySubAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "BindOrganizationPolicySubAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindOrganizationPolicySubAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindOrganizationPolicySubAccountResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CancelOrganizationMemberAuthAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelOrganizationMemberAuthAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCancelOrganizationPolicySubAccountRequest() (request *CancelOrganizationPolicySubAccountRequest) {
    request = &CancelOrganizationPolicySubAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationPolicySubAccount")
    
    
    return
}

func NewCancelOrganizationPolicySubAccountResponse() (response *CancelOrganizationPolicySubAccountResponse) {
    response = &CancelOrganizationPolicySubAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelOrganizationPolicySubAccount
// 解绑成员访问授权策略和组织管理员子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelOrganizationPolicySubAccount(request *CancelOrganizationPolicySubAccountRequest) (response *CancelOrganizationPolicySubAccountResponse, err error) {
    return c.CancelOrganizationPolicySubAccountWithContext(context.Background(), request)
}

// CancelOrganizationPolicySubAccount
// 解绑成员访问授权策略和组织管理员子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelOrganizationPolicySubAccountWithContext(ctx context.Context, request *CancelOrganizationPolicySubAccountRequest) (response *CancelOrganizationPolicySubAccountResponse, err error) {
    if request == nil {
        request = NewCancelOrganizationPolicySubAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CancelOrganizationPolicySubAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelOrganizationPolicySubAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelOrganizationPolicySubAccountResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CheckAccountDelete")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAccountDelete require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAccountDeleteResponse()
    err = c.Send(request, response)
    return
}

func NewClearExternalSAMLIdentityProviderRequest() (request *ClearExternalSAMLIdentityProviderRequest) {
    request = &ClearExternalSAMLIdentityProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ClearExternalSAMLIdentityProvider")
    
    
    return
}

func NewClearExternalSAMLIdentityProviderResponse() (response *ClearExternalSAMLIdentityProviderResponse) {
    response = &ClearExternalSAMLIdentityProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearExternalSAMLIdentityProvider
// 清空SAML身份提供商配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SSOSTATUSENABLENOTCLEARIDENTITYPROVIDER = "FailedOperation.SSoStatusEnableNotClearIdentityProvider"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) ClearExternalSAMLIdentityProvider(request *ClearExternalSAMLIdentityProviderRequest) (response *ClearExternalSAMLIdentityProviderResponse, err error) {
    return c.ClearExternalSAMLIdentityProviderWithContext(context.Background(), request)
}

// ClearExternalSAMLIdentityProvider
// 清空SAML身份提供商配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SSOSTATUSENABLENOTCLEARIDENTITYPROVIDER = "FailedOperation.SSoStatusEnableNotClearIdentityProvider"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) ClearExternalSAMLIdentityProviderWithContext(ctx context.Context, request *ClearExternalSAMLIdentityProviderRequest) (response *ClearExternalSAMLIdentityProviderResponse, err error) {
    if request == nil {
        request = NewClearExternalSAMLIdentityProviderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ClearExternalSAMLIdentityProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateGroup")
    
    
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGroup
// 创建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPOVERUPPERLIMIT = "FailedOperation.GroupOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNAMEFORMATERROR = "InvalidParameter.GroupNameFormatError"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

// CreateGroup
// 创建用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_GROUPOVERUPPERLIMIT = "FailedOperation.GroupOverUpperLimit"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNAMEFORMATERROR = "InvalidParameter.GroupNameFormatError"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrgServiceAssignRequest() (request *CreateOrgServiceAssignRequest) {
    request = &CreateOrgServiceAssignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateOrgServiceAssign")
    
    
    return
}

func NewCreateOrgServiceAssignResponse() (response *CreateOrgServiceAssignResponse) {
    response = &CreateOrgServiceAssignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrgServiceAssign
// 添加集团服务委派管理员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrgServiceAssign(request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
    return c.CreateOrgServiceAssignWithContext(context.Background(), request)
}

// CreateOrgServiceAssign
// 添加集团服务委派管理员
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_CREATEORGSERVICEASSIGNOVERLIMIT = "LimitExceeded.CreateOrgServiceAssignOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrgServiceAssignWithContext(ctx context.Context, request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
    if request == nil {
        request = NewCreateOrgServiceAssignRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrgServiceAssign")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrgServiceAssignResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganization")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationIdentity")
    
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
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
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMemberAuthIdentity")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMemberPolicy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateOrganizationMembersPolicy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreatePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleAssignmentRequest() (request *CreateRoleAssignmentRequest) {
    request = &CreateRoleAssignmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateRoleAssignment")
    
    
    return
}

func NewCreateRoleAssignmentResponse() (response *CreateRoleAssignmentResponse) {
    response = &CreateRoleAssignmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoleAssignment
// 在成员账号上授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONALREADYEXIST = "FailedOperation.RoleConfigurationAuthorizationAlreadyExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONOVERLIMIT = "FailedOperation.RoleConfigurationAuthorizationOverLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  LIMITEXCEEDED_CREATEROLEASSIGNMENTLIMITEXCEEDED = "LimitExceeded.CreateRoleAssignmentLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateRoleAssignment(request *CreateRoleAssignmentRequest) (response *CreateRoleAssignmentResponse, err error) {
    return c.CreateRoleAssignmentWithContext(context.Background(), request)
}

// CreateRoleAssignment
// 在成员账号上授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONALREADYEXIST = "FailedOperation.RoleConfigurationAuthorizationAlreadyExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONOVERLIMIT = "FailedOperation.RoleConfigurationAuthorizationOverLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  LIMITEXCEEDED_CREATEROLEASSIGNMENTLIMITEXCEEDED = "LimitExceeded.CreateRoleAssignmentLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateRoleAssignmentWithContext(ctx context.Context, request *CreateRoleAssignmentRequest) (response *CreateRoleAssignmentResponse, err error) {
    if request == nil {
        request = NewCreateRoleAssignmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateRoleAssignment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoleAssignment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleAssignmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleConfigurationRequest() (request *CreateRoleConfigurationRequest) {
    request = &CreateRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateRoleConfiguration")
    
    
    return
}

func NewCreateRoleConfigurationResponse() (response *CreateRoleConfigurationResponse) {
    response = &CreateRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoleConfiguration
// 创建权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGURATIONOVERUPPERLIMIT = "FailedOperation.ConfigurationOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_CONFIGURATIONNAMEALREADYEXISTS = "InvalidParameter.ConfigurationNameAlreadyExists"
//  INVALIDPARAMETER_CONFIGURATIONNAMEFORMATERROR = "InvalidParameter.ConfigurationNameFormatError"
func (c *Client) CreateRoleConfiguration(request *CreateRoleConfigurationRequest) (response *CreateRoleConfigurationResponse, err error) {
    return c.CreateRoleConfigurationWithContext(context.Background(), request)
}

// CreateRoleConfiguration
// 创建权限配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_CONFIGURATIONOVERUPPERLIMIT = "FailedOperation.ConfigurationOverUpperLimit"
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_CONFIGURATIONNAMEALREADYEXISTS = "InvalidParameter.ConfigurationNameAlreadyExists"
//  INVALIDPARAMETER_CONFIGURATIONNAMEFORMATERROR = "InvalidParameter.ConfigurationNameFormatError"
func (c *Client) CreateRoleConfigurationWithContext(ctx context.Context, request *CreateRoleConfigurationRequest) (response *CreateRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewCreateRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSCIMCredentialRequest() (request *CreateSCIMCredentialRequest) {
    request = &CreateSCIMCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateSCIMCredential")
    
    
    return
}

func NewCreateSCIMCredentialResponse() (response *CreateSCIMCredentialResponse) {
    response = &CreateSCIMCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSCIMCredential
// 创建SCIM密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SCIMCREDENTIALGENERATEERROR = "FailedOperation.ScimCredentialGenerateError"
//  LIMITEXCEEDED_SCIMCREDENTIALLIMITEXCEEDED = "LimitExceeded.ScimCredentialLimitExceeded"
func (c *Client) CreateSCIMCredential(request *CreateSCIMCredentialRequest) (response *CreateSCIMCredentialResponse, err error) {
    return c.CreateSCIMCredentialWithContext(context.Background(), request)
}

// CreateSCIMCredential
// 创建SCIM密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SCIMCREDENTIALGENERATEERROR = "FailedOperation.ScimCredentialGenerateError"
//  LIMITEXCEEDED_SCIMCREDENTIALLIMITEXCEEDED = "LimitExceeded.ScimCredentialLimitExceeded"
func (c *Client) CreateSCIMCredentialWithContext(ctx context.Context, request *CreateSCIMCredentialRequest) (response *CreateSCIMCredentialResponse, err error) {
    if request == nil {
        request = NewCreateSCIMCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateSCIMCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSCIMCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSCIMCredentialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
//  LIMITEXCEEDED_CREATEUSERLIMITEXCEEDED = "LimitExceeded.CreateUserLimitExceeded"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
//  LIMITEXCEEDED_CREATEUSERLIMITEXCEEDED = "LimitExceeded.CreateUserLimitExceeded"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserSyncProvisioningRequest() (request *CreateUserSyncProvisioningRequest) {
    request = &CreateUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "CreateUserSyncProvisioning")
    
    
    return
}

func NewCreateUserSyncProvisioningResponse() (response *CreateUserSyncProvisioningResponse) {
    response = &CreateUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserSyncProvisioning
// 创建子用户同步任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USERPROVISIONINGALREADYEXISTS = "FailedOperation.UserProvisioningAlreadyExists"
//  FAILEDOPERATION_USERPROVISIONINGOVERLIMIT = "FailedOperation.UserProvisioningOverLimit"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  LIMITEXCEEDED_CREATEUSERSYNCPROVISIONINGLIMITEXCEEDED = "LimitExceeded.CreateUserSyncProvisioningLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateUserSyncProvisioning(request *CreateUserSyncProvisioningRequest) (response *CreateUserSyncProvisioningResponse, err error) {
    return c.CreateUserSyncProvisioningWithContext(context.Background(), request)
}

// CreateUserSyncProvisioning
// 创建子用户同步任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USERPROVISIONINGALREADYEXISTS = "FailedOperation.UserProvisioningAlreadyExists"
//  FAILEDOPERATION_USERPROVISIONINGOVERLIMIT = "FailedOperation.UserProvisioningOverLimit"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  LIMITEXCEEDED_CREATEUSERSYNCPROVISIONINGLIMITEXCEEDED = "LimitExceeded.CreateUserSyncProvisioningLimitExceeded"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateUserSyncProvisioningWithContext(ctx context.Context, request *CreateUserSyncProvisioningRequest) (response *CreateUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewCreateUserSyncProvisioningRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "CreateUserSyncProvisioning")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserSyncProvisioningResponse()
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
//  FAILEDOPERATION_MEMBERACCOUNTDEREGISTERPENDING = "FailedOperation.MemberAccountDeregisterPending"
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
//  FAILEDOPERATION_MEMBERACCOUNTDEREGISTERPENDING = "FailedOperation.MemberAccountDeregisterPending"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGroup
// 删除用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEGROUPNOTALLOWEXISTUSER = "FailedOperation.DeleteGroupNotAllowExistUser"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTDELETE = "FailedOperation.ManualGroupNotDelete"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// 删除用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEGROUPNOTALLOWEXISTUSER = "FailedOperation.DeleteGroupNotAllowExistUser"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTDELETE = "FailedOperation.ManualGroupNotDelete"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteOrgServiceAssignRequest() (request *DeleteOrgServiceAssignRequest) {
    request = &DeleteOrgServiceAssignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteOrgServiceAssign")
    
    
    return
}

func NewDeleteOrgServiceAssignResponse() (response *DeleteOrgServiceAssignResponse) {
    response = &DeleteOrgServiceAssignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteOrgServiceAssign
// 删除集团服务委派管理员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DeleteOrgServiceAssign(request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
    return c.DeleteOrgServiceAssignWithContext(context.Background(), request)
}

// DeleteOrgServiceAssign
// 删除集团服务委派管理员
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  FAILEDOPERATION_ORGANIZATIONSERVICEASSIGNISUSE = "FailedOperation.OrganizationServiceAssignIsUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
func (c *Client) DeleteOrgServiceAssignWithContext(ctx context.Context, request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
    if request == nil {
        request = NewDeleteOrgServiceAssignRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrgServiceAssign")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteOrgServiceAssign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteOrgServiceAssignResponse()
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
//  FAILEDOPERATION_ORGMEMBERPOLICYEXIST = "FailedOperation.OrgMemberPolicyExist"
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
//  FAILEDOPERATION_ORGMEMBERPOLICYEXIST = "FailedOperation.OrgMemberPolicyExist"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganization")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationIdentity")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationMemberAuthIdentity")
    
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
// 从组织中移除成员账号，不会删除账号。
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
// 从组织中移除成员账号，不会删除账号。
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationMembersPolicy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteOrganizationNodes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeletePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleAssignmentRequest() (request *DeleteRoleAssignmentRequest) {
    request = &DeleteRoleAssignmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteRoleAssignment")
    
    
    return
}

func NewDeleteRoleAssignmentResponse() (response *DeleteRoleAssignmentResponse) {
    response = &DeleteRoleAssignmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoleAssignment
// 移除成员账号上的授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONAUTHORIZATIONNOTFOUND = "ResourceNotFound.RoleConfigurationAuthorizationNotFound"
func (c *Client) DeleteRoleAssignment(request *DeleteRoleAssignmentRequest) (response *DeleteRoleAssignmentResponse, err error) {
    return c.DeleteRoleAssignmentWithContext(context.Background(), request)
}

// DeleteRoleAssignment
// 移除成员账号上的授权
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONAUTHORIZATIONNOTFOUND = "ResourceNotFound.RoleConfigurationAuthorizationNotFound"
func (c *Client) DeleteRoleAssignmentWithContext(ctx context.Context, request *DeleteRoleAssignmentRequest) (response *DeleteRoleAssignmentResponse, err error) {
    if request == nil {
        request = NewDeleteRoleAssignmentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteRoleAssignment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoleAssignment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleAssignmentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRoleConfigurationRequest() (request *DeleteRoleConfigurationRequest) {
    request = &DeleteRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteRoleConfiguration")
    
    
    return
}

func NewDeleteRoleConfigurationResponse() (response *DeleteRoleConfigurationResponse) {
    response = &DeleteRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRoleConfiguration
// 删除权限配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGALREADYDEPLOYED = "FailedOperation.RoleConfigurationProvisioningAlreadyDeployed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYALREADYEXIST = "InvalidParameter.RolePolicyAlreadyExist"
func (c *Client) DeleteRoleConfiguration(request *DeleteRoleConfigurationRequest) (response *DeleteRoleConfigurationResponse, err error) {
    return c.DeleteRoleConfigurationWithContext(context.Background(), request)
}

// DeleteRoleConfiguration
// 删除权限配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGALREADYDEPLOYED = "FailedOperation.RoleConfigurationProvisioningAlreadyDeployed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYALREADYEXIST = "InvalidParameter.RolePolicyAlreadyExist"
func (c *Client) DeleteRoleConfigurationWithContext(ctx context.Context, request *DeleteRoleConfigurationRequest) (response *DeleteRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewDeleteRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSCIMCredentialRequest() (request *DeleteSCIMCredentialRequest) {
    request = &DeleteSCIMCredentialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteSCIMCredential")
    
    
    return
}

func NewDeleteSCIMCredentialResponse() (response *DeleteSCIMCredentialResponse) {
    response = &DeleteSCIMCredentialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSCIMCredential
// 删除SCIM密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
func (c *Client) DeleteSCIMCredential(request *DeleteSCIMCredentialRequest) (response *DeleteSCIMCredentialResponse, err error) {
    return c.DeleteSCIMCredentialWithContext(context.Background(), request)
}

// DeleteSCIMCredential
// 删除SCIM密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
func (c *Client) DeleteSCIMCredentialWithContext(ctx context.Context, request *DeleteSCIMCredentialRequest) (response *DeleteSCIMCredentialResponse, err error) {
    if request == nil {
        request = NewDeleteSCIMCredentialRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteSCIMCredential")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSCIMCredential require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSCIMCredentialResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteShareUnit")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteShareUnitMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteShareUnitResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteShareUnitResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteShareUnitResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALUSERNOTDELETE = "FailedOperation.ManualUserNotDelete"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTDELETE = "FailedOperation.SynchronizedUserNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_USERALREADYEXISTSGROUP = "InvalidParameter.UserAlreadyExistsGroup"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALUSERNOTDELETE = "FailedOperation.ManualUserNotDelete"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTDELETE = "FailedOperation.SynchronizedUserNotDelete"
//  FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_USERALREADYEXISTSGROUP = "InvalidParameter.UserAlreadyExistsGroup"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserSyncProvisioningRequest() (request *DeleteUserSyncProvisioningRequest) {
    request = &DeleteUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DeleteUserSyncProvisioning")
    
    
    return
}

func NewDeleteUserSyncProvisioningResponse() (response *DeleteUserSyncProvisioningResponse) {
    response = &DeleteUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUserSyncProvisioning
// 删除子用户同步任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USERPROVISIONINGFAILED = "FailedOperation.UserProvisioningFailed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) DeleteUserSyncProvisioning(request *DeleteUserSyncProvisioningRequest) (response *DeleteUserSyncProvisioningResponse, err error) {
    return c.DeleteUserSyncProvisioningWithContext(context.Background(), request)
}

// DeleteUserSyncProvisioning
// 删除子用户同步任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_USERPROVISIONINGFAILED = "FailedOperation.UserProvisioningFailed"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) DeleteUserSyncProvisioningWithContext(ctx context.Context, request *DeleteUserSyncProvisioningRequest) (response *DeleteUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewDeleteUserSyncProvisioningRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DeleteUserSyncProvisioning")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserSyncProvisioningResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeEffectivePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEffectivePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEffectivePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIdentityCenterRequest() (request *DescribeIdentityCenterRequest) {
    request = &DescribeIdentityCenterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeIdentityCenter")
    
    
    return
}

func NewDescribeIdentityCenterResponse() (response *DescribeIdentityCenterResponse) {
    response = &DescribeIdentityCenterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIdentityCenter
// 获取集团账号身份中心服务信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
func (c *Client) DescribeIdentityCenter(request *DescribeIdentityCenterRequest) (response *DescribeIdentityCenterResponse, err error) {
    return c.DescribeIdentityCenterWithContext(context.Background(), request)
}

// DescribeIdentityCenter
// 获取集团账号身份中心服务信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
func (c *Client) DescribeIdentityCenterWithContext(ctx context.Context, request *DescribeIdentityCenterRequest) (response *DescribeIdentityCenterResponse, err error) {
    if request == nil {
        request = NewDescribeIdentityCenterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeIdentityCenter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIdentityCenter require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIdentityCenterResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganization")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationAuthNode")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationFinancialByMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationFinancialByMonth")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationFinancialByProduct")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberAuthIdentities")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberEmailBind")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMemberPolicies")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMembersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationMembersAuthPolicyRequest() (request *DescribeOrganizationMembersAuthPolicyRequest) {
    request = &DescribeOrganizationMembersAuthPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembersAuthPolicy")
    
    
    return
}

func NewDescribeOrganizationMembersAuthPolicyResponse() (response *DescribeOrganizationMembersAuthPolicyResponse) {
    response = &DescribeOrganizationMembersAuthPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationMembersAuthPolicy
// 查询组织成员访问策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMembersAuthPolicy(request *DescribeOrganizationMembersAuthPolicyRequest) (response *DescribeOrganizationMembersAuthPolicyResponse, err error) {
    return c.DescribeOrganizationMembersAuthPolicyWithContext(context.Background(), request)
}

// DescribeOrganizationMembersAuthPolicy
// 查询组织成员访问策略列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMembersAuthPolicyWithContext(ctx context.Context, request *DescribeOrganizationMembersAuthPolicyRequest) (response *DescribeOrganizationMembersAuthPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationMembersAuthPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationMembersAuthPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationMembersAuthPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationMembersAuthPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeOrganizationNodes")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribePolicy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribePolicyConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePolicyConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePolicyConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceToShareMemberRequest() (request *DescribeResourceToShareMemberRequest) {
    request = &DescribeResourceToShareMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DescribeResourceToShareMember")
    
    
    return
}

func NewDescribeResourceToShareMemberResponse() (response *DescribeResourceToShareMemberResponse) {
    response = &DescribeResourceToShareMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceToShareMember
// 获取与我共享的资源列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourceToShareMember(request *DescribeResourceToShareMemberRequest) (response *DescribeResourceToShareMemberResponse, err error) {
    return c.DescribeResourceToShareMemberWithContext(context.Background(), request)
}

// DescribeResourceToShareMember
// 获取与我共享的资源列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourceToShareMemberWithContext(ctx context.Context, request *DescribeResourceToShareMemberRequest) (response *DescribeResourceToShareMemberResponse, err error) {
    if request == nil {
        request = NewDescribeResourceToShareMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeResourceToShareMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceToShareMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceToShareMemberResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareAreas")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareUnitMembers")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareUnitResources")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DescribeShareUnits")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DetachPolicy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DisablePolicyType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisablePolicyType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisablePolicyTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDismantleRoleConfigurationRequest() (request *DismantleRoleConfigurationRequest) {
    request = &DismantleRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "DismantleRoleConfiguration")
    
    
    return
}

func NewDismantleRoleConfigurationResponse() (response *DismantleRoleConfigurationResponse) {
    response = &DismantleRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DismantleRoleConfiguration
// 解除权限配置在成员账号上的部署
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGSTATUSERROR = "FailedOperation.RoleConfigurationProvisioningStatusError"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONPROVISIONINGNOTFOUND = "ResourceNotFound.RoleConfigurationProvisioningNotFound"
func (c *Client) DismantleRoleConfiguration(request *DismantleRoleConfigurationRequest) (response *DismantleRoleConfigurationResponse, err error) {
    return c.DismantleRoleConfigurationWithContext(context.Background(), request)
}

// DismantleRoleConfiguration
// 解除权限配置在成员账号上的部署
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//  FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGSTATUSERROR = "FailedOperation.RoleConfigurationProvisioningStatusError"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONPROVISIONINGNOTFOUND = "ResourceNotFound.RoleConfigurationProvisioningNotFound"
func (c *Client) DismantleRoleConfigurationWithContext(ctx context.Context, request *DismantleRoleConfigurationRequest) (response *DismantleRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewDismantleRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "DismantleRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismantleRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismantleRoleConfigurationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "EnablePolicyType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnablePolicyType require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnablePolicyTypeResponse()
    err = c.Send(request, response)
    return
}

func NewGetExternalSAMLIdentityProviderRequest() (request *GetExternalSAMLIdentityProviderRequest) {
    request = &GetExternalSAMLIdentityProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetExternalSAMLIdentityProvider")
    
    
    return
}

func NewGetExternalSAMLIdentityProviderResponse() (response *GetExternalSAMLIdentityProviderResponse) {
    response = &GetExternalSAMLIdentityProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetExternalSAMLIdentityProvider
// 查询SAML身份提供商配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) GetExternalSAMLIdentityProvider(request *GetExternalSAMLIdentityProviderRequest) (response *GetExternalSAMLIdentityProviderResponse, err error) {
    return c.GetExternalSAMLIdentityProviderWithContext(context.Background(), request)
}

// GetExternalSAMLIdentityProvider
// 查询SAML身份提供商配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLIDENTITYPROVIDERNOTFOUND = "ResourceNotFound.SAMLIdentityProviderNotFound"
func (c *Client) GetExternalSAMLIdentityProviderWithContext(ctx context.Context, request *GetExternalSAMLIdentityProviderRequest) (response *GetExternalSAMLIdentityProviderResponse, err error) {
    if request == nil {
        request = NewGetExternalSAMLIdentityProviderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetExternalSAMLIdentityProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupRequest() (request *GetGroupRequest) {
    request = &GetGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetGroup")
    
    
    return
}

func NewGetGroupResponse() (response *GetGroupResponse) {
    response = &GetGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetGroup
// 查询用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
    return c.GetGroupWithContext(context.Background(), request)
}

// GetGroup
// 查询用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest) (response *GetGroupResponse, err error) {
    if request == nil {
        request = NewGetGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetGroupResponse()
    err = c.Send(request, response)
    return
}

func NewGetProvisioningTaskStatusRequest() (request *GetProvisioningTaskStatusRequest) {
    request = &GetProvisioningTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetProvisioningTaskStatus")
    
    
    return
}

func NewGetProvisioningTaskStatusResponse() (response *GetProvisioningTaskStatusResponse) {
    response = &GetProvisioningTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetProvisioningTaskStatus
// 查询用户同步异步任务的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGTASKNOTFOUND = "ResourceNotFound.UserProvisioningTaskNotFound"
func (c *Client) GetProvisioningTaskStatus(request *GetProvisioningTaskStatusRequest) (response *GetProvisioningTaskStatusResponse, err error) {
    return c.GetProvisioningTaskStatusWithContext(context.Background(), request)
}

// GetProvisioningTaskStatus
// 查询用户同步异步任务的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGTASKNOTFOUND = "ResourceNotFound.UserProvisioningTaskNotFound"
func (c *Client) GetProvisioningTaskStatusWithContext(ctx context.Context, request *GetProvisioningTaskStatusRequest) (response *GetProvisioningTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetProvisioningTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetProvisioningTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetProvisioningTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetProvisioningTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetRoleConfigurationRequest() (request *GetRoleConfigurationRequest) {
    request = &GetRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetRoleConfiguration")
    
    
    return
}

func NewGetRoleConfigurationResponse() (response *GetRoleConfigurationResponse) {
    response = &GetRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetRoleConfiguration
// 查询权限配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) GetRoleConfiguration(request *GetRoleConfigurationRequest) (response *GetRoleConfigurationResponse, err error) {
    return c.GetRoleConfigurationWithContext(context.Background(), request)
}

// GetRoleConfiguration
// 查询权限配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) GetRoleConfigurationWithContext(ctx context.Context, request *GetRoleConfigurationRequest) (response *GetRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewGetRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewGetSCIMSynchronizationStatusRequest() (request *GetSCIMSynchronizationStatusRequest) {
    request = &GetSCIMSynchronizationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetSCIMSynchronizationStatus")
    
    
    return
}

func NewGetSCIMSynchronizationStatusResponse() (response *GetSCIMSynchronizationStatusResponse) {
    response = &GetSCIMSynchronizationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSCIMSynchronizationStatus
// 获取SCIM同步状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetSCIMSynchronizationStatus(request *GetSCIMSynchronizationStatusRequest) (response *GetSCIMSynchronizationStatusResponse, err error) {
    return c.GetSCIMSynchronizationStatusWithContext(context.Background(), request)
}

// GetSCIMSynchronizationStatus
// 获取SCIM同步状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetSCIMSynchronizationStatusWithContext(ctx context.Context, request *GetSCIMSynchronizationStatusRequest) (response *GetSCIMSynchronizationStatusResponse, err error) {
    if request == nil {
        request = NewGetSCIMSynchronizationStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetSCIMSynchronizationStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSCIMSynchronizationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSCIMSynchronizationStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskStatusRequest() (request *GetTaskStatusRequest) {
    request = &GetTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetTaskStatus")
    
    
    return
}

func NewGetTaskStatusResponse() (response *GetTaskStatusResponse) {
    response = &GetTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskStatus
// 查询异步任务的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONTASKNOTFOUND = "ResourceNotFound.RoleConfigurationTaskNotFound"
func (c *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    return c.GetTaskStatusWithContext(context.Background(), request)
}

// GetTaskStatus
// 查询异步任务的状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_ROLECONFIGURATIONTASKNOTFOUND = "ResourceNotFound.RoleConfigurationTaskNotFound"
func (c *Client) GetTaskStatusWithContext(ctx context.Context, request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetTaskStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetTaskStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserRequest() (request *GetUserRequest) {
    request = &GetUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetUser")
    
    
    return
}

func NewGetUserResponse() (response *GetUserResponse) {
    response = &GetUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUser
// 查询用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
    return c.GetUserWithContext(context.Background(), request)
}

// GetUser
// 查询用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest) (response *GetUserResponse, err error) {
    if request == nil {
        request = NewGetUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetUserSyncProvisioningRequest() (request *GetUserSyncProvisioningRequest) {
    request = &GetUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetUserSyncProvisioning")
    
    
    return
}

func NewGetUserSyncProvisioningResponse() (response *GetUserSyncProvisioningResponse) {
    response = &GetUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUserSyncProvisioning
// 查询CAM用户同步
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) GetUserSyncProvisioning(request *GetUserSyncProvisioningRequest) (response *GetUserSyncProvisioningResponse, err error) {
    return c.GetUserSyncProvisioningWithContext(context.Background(), request)
}

// GetUserSyncProvisioning
// 查询CAM用户同步
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) GetUserSyncProvisioningWithContext(ctx context.Context, request *GetUserSyncProvisioningRequest) (response *GetUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewGetUserSyncProvisioningRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetUserSyncProvisioning")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUserSyncProvisioningResponse()
    err = c.Send(request, response)
    return
}

func NewGetZoneSAMLServiceProviderInfoRequest() (request *GetZoneSAMLServiceProviderInfoRequest) {
    request = &GetZoneSAMLServiceProviderInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetZoneSAMLServiceProviderInfo")
    
    
    return
}

func NewGetZoneSAMLServiceProviderInfoResponse() (response *GetZoneSAMLServiceProviderInfoResponse) {
    response = &GetZoneSAMLServiceProviderInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetZoneSAMLServiceProviderInfo
// 查询SAML服务提供商配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SAMLMETADATADOCUMENTGETFAILED = "FailedOperation.SAMLMetadataDocumentGetFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLSERVICEPROVIDERNOTFOUND = "ResourceNotFound.SAMLServiceProviderNotFound"
func (c *Client) GetZoneSAMLServiceProviderInfo(request *GetZoneSAMLServiceProviderInfoRequest) (response *GetZoneSAMLServiceProviderInfoResponse, err error) {
    return c.GetZoneSAMLServiceProviderInfoWithContext(context.Background(), request)
}

// GetZoneSAMLServiceProviderInfo
// 查询SAML服务提供商配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_SAMLMETADATADOCUMENTGETFAILED = "FailedOperation.SAMLMetadataDocumentGetFailed"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_SAMLSERVICEPROVIDERNOTFOUND = "ResourceNotFound.SAMLServiceProviderNotFound"
func (c *Client) GetZoneSAMLServiceProviderInfoWithContext(ctx context.Context, request *GetZoneSAMLServiceProviderInfoRequest) (response *GetZoneSAMLServiceProviderInfoResponse, err error) {
    if request == nil {
        request = NewGetZoneSAMLServiceProviderInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetZoneSAMLServiceProviderInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetZoneSAMLServiceProviderInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetZoneSAMLServiceProviderInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetZoneStatisticsRequest() (request *GetZoneStatisticsRequest) {
    request = &GetZoneStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "GetZoneStatistics")
    
    
    return
}

func NewGetZoneStatisticsResponse() (response *GetZoneStatisticsResponse) {
    response = &GetZoneStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetZoneStatistics
// 查询空间的统计信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetZoneStatistics(request *GetZoneStatisticsRequest) (response *GetZoneStatisticsResponse, err error) {
    return c.GetZoneStatisticsWithContext(context.Background(), request)
}

// GetZoneStatistics
// 查询空间的统计信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetZoneStatisticsWithContext(ctx context.Context, request *GetZoneStatisticsRequest) (response *GetZoneStatisticsResponse, err error) {
    if request == nil {
        request = NewGetZoneStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "GetZoneStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetZoneStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetZoneStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewInviteOrganizationMemberRequest() (request *InviteOrganizationMemberRequest) {
    request = &InviteOrganizationMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "InviteOrganizationMember")
    
    
    return
}

func NewInviteOrganizationMemberResponse() (response *InviteOrganizationMemberResponse) {
    response = &InviteOrganizationMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InviteOrganizationMember
// 邀请组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYEXIST = "FailedOperation.ApplyExist"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_EXISTOTHERORGANIZATIONMEMBERSHARED = "FailedOperation.ExistOtherOrganizationMemberShared"
//  FAILEDOPERATION_GETACCOUNTREGION = "FailedOperation.GetAccountRegion"
//  FAILEDOPERATION_IMPORTFILEILLEGAL = "FailedOperation.ImportFileIllegal"
//  FAILEDOPERATION_INVITATIONEXIST = "FailedOperation.InvitationExist"
//  FAILEDOPERATION_MEMBEREXISTINOTHERORGANIZATION = "FailedOperation.MemberExistInOtherOrganization"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHRELATIONEXIST = "FailedOperation.OrganizationAuthRelationExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBEREXIST = "FailedOperation.OrganizationMemberExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_INVITATIONOVERLIMIT = "LimitExceeded.InvitationOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFMEMBER = "UnsupportedOperation.AbnormalFinancialStatusOfMember"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_AGENTNOTSAME = "UnsupportedOperation.AgentNotSame"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERHASVOUCHER = "UnsupportedOperation.MemberHasVoucher"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERISNOTCLIENT = "UnsupportedOperation.MemberIsNotClient"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) InviteOrganizationMember(request *InviteOrganizationMemberRequest) (response *InviteOrganizationMemberResponse, err error) {
    return c.InviteOrganizationMemberWithContext(context.Background(), request)
}

// InviteOrganizationMember
// 邀请组织成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYEXIST = "FailedOperation.ApplyExist"
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHINFONOTSAME = "FailedOperation.AuthInfoNotSame"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_EXISTOTHERORGANIZATIONMEMBERSHARED = "FailedOperation.ExistOtherOrganizationMemberShared"
//  FAILEDOPERATION_GETACCOUNTREGION = "FailedOperation.GetAccountRegion"
//  FAILEDOPERATION_IMPORTFILEILLEGAL = "FailedOperation.ImportFileIllegal"
//  FAILEDOPERATION_INVITATIONEXIST = "FailedOperation.InvitationExist"
//  FAILEDOPERATION_MEMBEREXISTINOTHERORGANIZATION = "FailedOperation.MemberExistInOtherOrganization"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_NOTSAMEREGION = "FailedOperation.NotSameRegion"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHRELATIONEXIST = "FailedOperation.OrganizationAuthRelationExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBEREXIST = "FailedOperation.OrganizationMemberExist"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  FAILEDOPERATION_RESENTINVITATION = "FailedOperation.ReSentInvitation"
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_INVITATIONOVERLIMIT = "LimitExceeded.InvitationOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFMEMBER = "UnsupportedOperation.AbnormalFinancialStatusOfMember"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_AGENTNOTSAME = "UnsupportedOperation.AgentNotSame"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERHASVOUCHER = "UnsupportedOperation.MemberHasVoucher"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_MEMBERISNOTCLIENT = "UnsupportedOperation.MemberIsNotClient"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) InviteOrganizationMemberWithContext(ctx context.Context, request *InviteOrganizationMemberRequest) (response *InviteOrganizationMemberResponse, err error) {
    if request == nil {
        request = NewInviteOrganizationMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "InviteOrganizationMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InviteOrganizationMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewInviteOrganizationMemberResponse()
    err = c.Send(request, response)
    return
}

func NewListExternalSAMLIdPCertificatesRequest() (request *ListExternalSAMLIdPCertificatesRequest) {
    request = &ListExternalSAMLIdPCertificatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListExternalSAMLIdPCertificates")
    
    
    return
}

func NewListExternalSAMLIdPCertificatesResponse() (response *ListExternalSAMLIdPCertificatesResponse) {
    response = &ListExternalSAMLIdPCertificatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListExternalSAMLIdPCertificates
// 查询SAML签名证书列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
func (c *Client) ListExternalSAMLIdPCertificates(request *ListExternalSAMLIdPCertificatesRequest) (response *ListExternalSAMLIdPCertificatesResponse, err error) {
    return c.ListExternalSAMLIdPCertificatesWithContext(context.Background(), request)
}

// ListExternalSAMLIdPCertificates
// 查询SAML签名证书列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  INTERNALERROR = "InternalError"
func (c *Client) ListExternalSAMLIdPCertificatesWithContext(ctx context.Context, request *ListExternalSAMLIdPCertificatesRequest) (response *ListExternalSAMLIdPCertificatesResponse, err error) {
    if request == nil {
        request = NewListExternalSAMLIdPCertificatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListExternalSAMLIdPCertificates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListExternalSAMLIdPCertificates require credential")
    }

    request.SetContext(ctx)
    
    response = NewListExternalSAMLIdPCertificatesResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupMembersRequest() (request *ListGroupMembersRequest) {
    request = &ListGroupMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListGroupMembers")
    
    
    return
}

func NewListGroupMembersResponse() (response *ListGroupMembersResponse) {
    response = &ListGroupMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListGroupMembers
// 查询用户组中的用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) ListGroupMembers(request *ListGroupMembersRequest) (response *ListGroupMembersResponse, err error) {
    return c.ListGroupMembersWithContext(context.Background(), request)
}

// ListGroupMembers
// 查询用户组中的用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) ListGroupMembersWithContext(ctx context.Context, request *ListGroupMembersRequest) (response *ListGroupMembersResponse, err error) {
    if request == nil {
        request = NewListGroupMembersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListGroupMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGroupMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGroupMembersResponse()
    err = c.Send(request, response)
    return
}

func NewListGroupsRequest() (request *ListGroupsRequest) {
    request = &ListGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListGroups")
    
    
    return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
    response = &ListGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListGroups
// 查询用户组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    return c.ListGroupsWithContext(context.Background(), request)
}

// ListGroups
// 查询用户组列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
    if request == nil {
        request = NewListGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewListGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewListJoinedGroupsForUserRequest() (request *ListJoinedGroupsForUserRequest) {
    request = &ListJoinedGroupsForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListJoinedGroupsForUser")
    
    
    return
}

func NewListJoinedGroupsForUserResponse() (response *ListJoinedGroupsForUserResponse) {
    response = &ListJoinedGroupsForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListJoinedGroupsForUser
// 查询用户加入的用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListJoinedGroupsForUser(request *ListJoinedGroupsForUserRequest) (response *ListJoinedGroupsForUserResponse, err error) {
    return c.ListJoinedGroupsForUserWithContext(context.Background(), request)
}

// ListJoinedGroupsForUser
// 查询用户加入的用户组
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListJoinedGroupsForUserWithContext(ctx context.Context, request *ListJoinedGroupsForUserRequest) (response *ListJoinedGroupsForUserResponse, err error) {
    if request == nil {
        request = NewListJoinedGroupsForUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListJoinedGroupsForUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListJoinedGroupsForUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewListJoinedGroupsForUserResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListNonCompliantResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListNonCompliantResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewListNonCompliantResourceResponse()
    err = c.Send(request, response)
    return
}

func NewListOrgServiceAssignMemberRequest() (request *ListOrgServiceAssignMemberRequest) {
    request = &ListOrgServiceAssignMemberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrgServiceAssignMember")
    
    
    return
}

func NewListOrgServiceAssignMemberResponse() (response *ListOrgServiceAssignMemberResponse) {
    response = &ListOrgServiceAssignMemberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrgServiceAssignMember
// 获取集团服务委派管理员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrgServiceAssignMember(request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
    return c.ListOrgServiceAssignMemberWithContext(context.Background(), request)
}

// ListOrgServiceAssignMember
// 获取集团服务委派管理员列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEORGSERVICEUSAGESTATUSERR = "FailedOperation.DescribeOrgServiceUsageStatusErr"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrgServiceAssignMemberWithContext(ctx context.Context, request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
    if request == nil {
        request = NewListOrgServiceAssignMemberRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListOrgServiceAssignMember")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrgServiceAssignMember require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrgServiceAssignMemberResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListOrganizationIdentity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationIdentity require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewListOrganizationServiceRequest() (request *ListOrganizationServiceRequest) {
    request = &ListOrganizationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationService")
    
    
    return
}

func NewListOrganizationServiceResponse() (response *ListOrganizationServiceResponse) {
    response = &ListOrganizationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListOrganizationService
// 获取集团服务设置列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationService(request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
    return c.ListOrganizationServiceWithContext(context.Background(), request)
}

// ListOrganizationService
// 获取集团服务设置列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationServiceWithContext(ctx context.Context, request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
    if request == nil {
        request = NewListOrganizationServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListOrganizationService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListOrganizationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewListOrganizationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewListPermissionPoliciesInRoleConfigurationRequest() (request *ListPermissionPoliciesInRoleConfigurationRequest) {
    request = &ListPermissionPoliciesInRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListPermissionPoliciesInRoleConfiguration")
    
    
    return
}

func NewListPermissionPoliciesInRoleConfigurationResponse() (response *ListPermissionPoliciesInRoleConfigurationResponse) {
    response = &ListPermissionPoliciesInRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListPermissionPoliciesInRoleConfiguration
// 获取权限配置中的策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) ListPermissionPoliciesInRoleConfiguration(request *ListPermissionPoliciesInRoleConfigurationRequest) (response *ListPermissionPoliciesInRoleConfigurationResponse, err error) {
    return c.ListPermissionPoliciesInRoleConfigurationWithContext(context.Background(), request)
}

// ListPermissionPoliciesInRoleConfiguration
// 获取权限配置中的策略列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) ListPermissionPoliciesInRoleConfigurationWithContext(ctx context.Context, request *ListPermissionPoliciesInRoleConfigurationRequest) (response *ListPermissionPoliciesInRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewListPermissionPoliciesInRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListPermissionPoliciesInRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPermissionPoliciesInRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPermissionPoliciesInRoleConfigurationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListPolicies")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListPoliciesForTarget")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListPoliciesForTarget require credential")
    }

    request.SetContext(ctx)
    
    response = NewListPoliciesForTargetResponse()
    err = c.Send(request, response)
    return
}

func NewListRoleAssignmentsRequest() (request *ListRoleAssignmentsRequest) {
    request = &ListRoleAssignmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListRoleAssignments")
    
    
    return
}

func NewListRoleAssignmentsResponse() (response *ListRoleAssignmentsResponse) {
    response = &ListRoleAssignmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRoleAssignments
// 查询授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleAssignments(request *ListRoleAssignmentsRequest) (response *ListRoleAssignmentsResponse, err error) {
    return c.ListRoleAssignmentsWithContext(context.Background(), request)
}

// ListRoleAssignments
// 查询授权列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleAssignmentsWithContext(ctx context.Context, request *ListRoleAssignmentsRequest) (response *ListRoleAssignmentsResponse, err error) {
    if request == nil {
        request = NewListRoleAssignmentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListRoleAssignments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleAssignments require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleAssignmentsResponse()
    err = c.Send(request, response)
    return
}

func NewListRoleConfigurationProvisioningsRequest() (request *ListRoleConfigurationProvisioningsRequest) {
    request = &ListRoleConfigurationProvisioningsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListRoleConfigurationProvisionings")
    
    
    return
}

func NewListRoleConfigurationProvisioningsResponse() (response *ListRoleConfigurationProvisioningsResponse) {
    response = &ListRoleConfigurationProvisioningsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRoleConfigurationProvisionings
// 查询权限配置部署列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurationProvisionings(request *ListRoleConfigurationProvisioningsRequest) (response *ListRoleConfigurationProvisioningsResponse, err error) {
    return c.ListRoleConfigurationProvisioningsWithContext(context.Background(), request)
}

// ListRoleConfigurationProvisionings
// 查询权限配置部署列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurationProvisioningsWithContext(ctx context.Context, request *ListRoleConfigurationProvisioningsRequest) (response *ListRoleConfigurationProvisioningsResponse, err error) {
    if request == nil {
        request = NewListRoleConfigurationProvisioningsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListRoleConfigurationProvisionings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleConfigurationProvisionings require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleConfigurationProvisioningsResponse()
    err = c.Send(request, response)
    return
}

func NewListRoleConfigurationsRequest() (request *ListRoleConfigurationsRequest) {
    request = &ListRoleConfigurationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListRoleConfigurations")
    
    
    return
}

func NewListRoleConfigurationsResponse() (response *ListRoleConfigurationsResponse) {
    response = &ListRoleConfigurationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRoleConfigurations
// 查询权限配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurations(request *ListRoleConfigurationsRequest) (response *ListRoleConfigurationsResponse, err error) {
    return c.ListRoleConfigurationsWithContext(context.Background(), request)
}

// ListRoleConfigurations
// 查询权限配置列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleConfigurationsWithContext(ctx context.Context, request *ListRoleConfigurationsRequest) (response *ListRoleConfigurationsResponse, err error) {
    if request == nil {
        request = NewListRoleConfigurationsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListRoleConfigurations")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRoleConfigurations require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRoleConfigurationsResponse()
    err = c.Send(request, response)
    return
}

func NewListSCIMCredentialsRequest() (request *ListSCIMCredentialsRequest) {
    request = &ListSCIMCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListSCIMCredentials")
    
    
    return
}

func NewListSCIMCredentialsResponse() (response *ListSCIMCredentialsResponse) {
    response = &ListSCIMCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSCIMCredentials
// 查询用户SCIM密钥列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) ListSCIMCredentials(request *ListSCIMCredentialsRequest) (response *ListSCIMCredentialsResponse, err error) {
    return c.ListSCIMCredentialsWithContext(context.Background(), request)
}

// ListSCIMCredentials
// 查询用户SCIM密钥列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) ListSCIMCredentialsWithContext(ctx context.Context, request *ListSCIMCredentialsRequest) (response *ListSCIMCredentialsResponse, err error) {
    if request == nil {
        request = NewListSCIMCredentialsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListSCIMCredentials")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSCIMCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSCIMCredentialsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListTargetsForPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTargetsForPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTargetsForPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewListTasksRequest() (request *ListTasksRequest) {
    request = &ListTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListTasks")
    
    
    return
}

func NewListTasksResponse() (response *ListTasksResponse) {
    response = &ListTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTasks
// 查询异步任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListTasks(request *ListTasksRequest) (response *ListTasksResponse, err error) {
    return c.ListTasksWithContext(context.Background(), request)
}

// ListTasks
// 查询异步任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListTasksWithContext(ctx context.Context, request *ListTasksRequest) (response *ListTasksResponse, err error) {
    if request == nil {
        request = NewListTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTasksResponse()
    err = c.Send(request, response)
    return
}

func NewListUserSyncProvisioningsRequest() (request *ListUserSyncProvisioningsRequest) {
    request = &ListUserSyncProvisioningsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListUserSyncProvisionings")
    
    
    return
}

func NewListUserSyncProvisioningsResponse() (response *ListUserSyncProvisioningsResponse) {
    response = &ListUserSyncProvisioningsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUserSyncProvisionings
// 查询CAM用户同步列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListUserSyncProvisionings(request *ListUserSyncProvisioningsRequest) (response *ListUserSyncProvisioningsResponse, err error) {
    return c.ListUserSyncProvisioningsWithContext(context.Background(), request)
}

// ListUserSyncProvisionings
// 查询CAM用户同步列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
func (c *Client) ListUserSyncProvisioningsWithContext(ctx context.Context, request *ListUserSyncProvisioningsRequest) (response *ListUserSyncProvisioningsResponse, err error) {
    if request == nil {
        request = NewListUserSyncProvisioningsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListUserSyncProvisionings")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUserSyncProvisionings require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUserSyncProvisioningsResponse()
    err = c.Send(request, response)
    return
}

func NewListUsersRequest() (request *ListUsersRequest) {
    request = &ListUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ListUsers")
    
    
    return
}

func NewListUsersResponse() (response *ListUsersResponse) {
    response = &ListUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsers
// 查询用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
    return c.ListUsersWithContext(context.Background(), request)
}

// ListUsers
// 查询用户列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest) (response *ListUsersResponse, err error) {
    if request == nil {
        request = NewListUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ListUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsersResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "MoveOrganizationNodeMembers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("MoveOrganizationNodeMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewMoveOrganizationNodeMembersResponse()
    err = c.Send(request, response)
    return
}

func NewOpenIdentityCenterRequest() (request *OpenIdentityCenterRequest) {
    request = &OpenIdentityCenterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "OpenIdentityCenter")
    
    
    return
}

func NewOpenIdentityCenterResponse() (response *OpenIdentityCenterResponse) {
    response = &OpenIdentityCenterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenIdentityCenter
// 开通身份中心服务（CIC）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERALREADYOPEN = "FailedOperation.IdentityCenterAlreadyOpen"
//  FAILEDOPERATION_IDENTITYCENTERNOTAUTH = "FailedOperation.IdentityCenterNotAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTENTERPRISEAUTH = "FailedOperation.IdentityCenterNotEnterpriseAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTORGANIZATIONMANAGER = "FailedOperation.IdentityCenterNotOrganizationManager"
//  FAILEDOPERATION_IDENTITYCENTERORGANIZATIONNOTOPEN = "FailedOperation.IdentityCenterOrganizationNotOpen"
//  INVALIDPARAMETERVALUE_IDENTITYCENTERZONENAMEALREADYEXIST = "InvalidParameterValue.IdentityCenterZoneNameAlreadyExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
func (c *Client) OpenIdentityCenter(request *OpenIdentityCenterRequest) (response *OpenIdentityCenterResponse, err error) {
    return c.OpenIdentityCenterWithContext(context.Background(), request)
}

// OpenIdentityCenter
// 开通身份中心服务（CIC）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERALREADYOPEN = "FailedOperation.IdentityCenterAlreadyOpen"
//  FAILEDOPERATION_IDENTITYCENTERNOTAUTH = "FailedOperation.IdentityCenterNotAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTENTERPRISEAUTH = "FailedOperation.IdentityCenterNotEnterpriseAuth"
//  FAILEDOPERATION_IDENTITYCENTERNOTORGANIZATIONMANAGER = "FailedOperation.IdentityCenterNotOrganizationManager"
//  FAILEDOPERATION_IDENTITYCENTERORGANIZATIONNOTOPEN = "FailedOperation.IdentityCenterOrganizationNotOpen"
//  INVALIDPARAMETERVALUE_IDENTITYCENTERZONENAMEALREADYEXIST = "InvalidParameterValue.IdentityCenterZoneNameAlreadyExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
func (c *Client) OpenIdentityCenterWithContext(ctx context.Context, request *OpenIdentityCenterRequest) (response *OpenIdentityCenterResponse, err error) {
    if request == nil {
        request = NewOpenIdentityCenterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "OpenIdentityCenter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenIdentityCenter require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenIdentityCenterResponse()
    err = c.Send(request, response)
    return
}

func NewProvisionRoleConfigurationRequest() (request *ProvisionRoleConfigurationRequest) {
    request = &ProvisionRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "ProvisionRoleConfiguration")
    
    
    return
}

func NewProvisionRoleConfigurationResponse() (response *ProvisionRoleConfigurationResponse) {
    response = &ProvisionRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ProvisionRoleConfiguration
// 将权限配置部署到成员账号上
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
func (c *Client) ProvisionRoleConfiguration(request *ProvisionRoleConfigurationRequest) (response *ProvisionRoleConfigurationResponse, err error) {
    return c.ProvisionRoleConfigurationWithContext(context.Background(), request)
}

// ProvisionRoleConfiguration
// 将权限配置部署到成员账号上
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//  FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//  INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
func (c *Client) ProvisionRoleConfigurationWithContext(ctx context.Context, request *ProvisionRoleConfigurationRequest) (response *ProvisionRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewProvisionRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "ProvisionRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ProvisionRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewProvisionRoleConfigurationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "QuitOrganization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QuitOrganization require credential")
    }

    request.SetContext(ctx)
    
    response = NewQuitOrganizationResponse()
    err = c.Send(request, response)
    return
}

func NewRejectJoinShareUnitInvitationRequest() (request *RejectJoinShareUnitInvitationRequest) {
    request = &RejectJoinShareUnitInvitationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RejectJoinShareUnitInvitation")
    
    
    return
}

func NewRejectJoinShareUnitInvitationResponse() (response *RejectJoinShareUnitInvitationResponse) {
    response = &RejectJoinShareUnitInvitationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RejectJoinShareUnitInvitation
// 拒绝加入共享单元邀请。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectJoinShareUnitInvitation(request *RejectJoinShareUnitInvitationRequest) (response *RejectJoinShareUnitInvitationResponse, err error) {
    return c.RejectJoinShareUnitInvitationWithContext(context.Background(), request)
}

// RejectJoinShareUnitInvitation
// 拒绝加入共享单元邀请。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SHAREMEMBERNOTEXIST = "FailedOperation.ShareMemberNotExist"
//  FAILEDOPERATION_SHAREUNITNOTEXIST = "FailedOperation.ShareUnitNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RejectJoinShareUnitInvitationWithContext(ctx context.Context, request *RejectJoinShareUnitInvitationRequest) (response *RejectJoinShareUnitInvitationResponse, err error) {
    if request == nil {
        request = NewRejectJoinShareUnitInvitationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RejectJoinShareUnitInvitation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RejectJoinShareUnitInvitation require credential")
    }

    request.SetContext(ctx)
    
    response = NewRejectJoinShareUnitInvitationResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveExternalSAMLIdPCertificateRequest() (request *RemoveExternalSAMLIdPCertificateRequest) {
    request = &RemoveExternalSAMLIdPCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RemoveExternalSAMLIdPCertificate")
    
    
    return
}

func NewRemoveExternalSAMLIdPCertificateResponse() (response *RemoveExternalSAMLIdPCertificateResponse) {
    response = &RemoveExternalSAMLIdPCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveExternalSAMLIdPCertificate
// 移除SAML签名证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_X509CERTIFICATELIMITEXCEEDED = "FailedOperation.X509CertificateLimitExceeded"
//  FAILEDOPERATION_X509CERTIFICATEMINIMUMREQUIRED = "FailedOperation.X509CertificateMinimumRequired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_X509CERTIFICATENOTFOUND = "ResourceNotFound.X509CertificateNotFound"
func (c *Client) RemoveExternalSAMLIdPCertificate(request *RemoveExternalSAMLIdPCertificateRequest) (response *RemoveExternalSAMLIdPCertificateResponse, err error) {
    return c.RemoveExternalSAMLIdPCertificateWithContext(context.Background(), request)
}

// RemoveExternalSAMLIdPCertificate
// 移除SAML签名证书
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_X509CERTIFICATELIMITEXCEEDED = "FailedOperation.X509CertificateLimitExceeded"
//  FAILEDOPERATION_X509CERTIFICATEMINIMUMREQUIRED = "FailedOperation.X509CertificateMinimumRequired"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_X509CERTIFICATENOTFOUND = "ResourceNotFound.X509CertificateNotFound"
func (c *Client) RemoveExternalSAMLIdPCertificateWithContext(ctx context.Context, request *RemoveExternalSAMLIdPCertificateRequest) (response *RemoveExternalSAMLIdPCertificateResponse, err error) {
    if request == nil {
        request = NewRemoveExternalSAMLIdPCertificateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RemoveExternalSAMLIdPCertificate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveExternalSAMLIdPCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveExternalSAMLIdPCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewRemovePermissionPolicyFromRoleConfigurationRequest() (request *RemovePermissionPolicyFromRoleConfigurationRequest) {
    request = &RemovePermissionPolicyFromRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RemovePermissionPolicyFromRoleConfiguration")
    
    
    return
}

func NewRemovePermissionPolicyFromRoleConfigurationResponse() (response *RemovePermissionPolicyFromRoleConfigurationResponse) {
    response = &RemovePermissionPolicyFromRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemovePermissionPolicyFromRoleConfiguration
// 为权限配置移除策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) RemovePermissionPolicyFromRoleConfiguration(request *RemovePermissionPolicyFromRoleConfigurationRequest) (response *RemovePermissionPolicyFromRoleConfigurationResponse, err error) {
    return c.RemovePermissionPolicyFromRoleConfigurationWithContext(context.Background(), request)
}

// RemovePermissionPolicyFromRoleConfiguration
// 为权限配置移除策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) RemovePermissionPolicyFromRoleConfigurationWithContext(ctx context.Context, request *RemovePermissionPolicyFromRoleConfigurationRequest) (response *RemovePermissionPolicyFromRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewRemovePermissionPolicyFromRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RemovePermissionPolicyFromRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemovePermissionPolicyFromRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemovePermissionPolicyFromRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
    request = &RemoveUserFromGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "RemoveUserFromGroup")
    
    
    return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
    response = &RemoveUserFromGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveUserFromGroup
// 从用户组中移除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"
//  LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"
func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    return c.RemoveUserFromGroupWithContext(context.Background(), request)
}

// RemoveUserFromGroup
// 从用户组中移除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//  INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"
//  LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"
func (c *Client) RemoveUserFromGroupWithContext(ctx context.Context, request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
    if request == nil {
        request = NewRemoveUserFromGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "RemoveUserFromGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserFromGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserFromGroupResponse()
    err = c.Send(request, response)
    return
}

func NewSendOrgMemberAccountBindEmailRequest() (request *SendOrgMemberAccountBindEmailRequest) {
    request = &SendOrgMemberAccountBindEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "SendOrgMemberAccountBindEmail")
    
    
    return
}

func NewSendOrgMemberAccountBindEmailResponse() (response *SendOrgMemberAccountBindEmailResponse) {
    response = &SendOrgMemberAccountBindEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SendOrgMemberAccountBindEmail
// 重新发送成员绑定邮箱激活邮件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SENDEMAILLIMIT = "LimitExceeded.SendEmailLimit"
//  LIMITEXCEEDED_SENDEMAILWITHINONEHOURLIMIT = "LimitExceeded.SendEmailWithinOneHourLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) SendOrgMemberAccountBindEmail(request *SendOrgMemberAccountBindEmailRequest) (response *SendOrgMemberAccountBindEmailResponse, err error) {
    return c.SendOrgMemberAccountBindEmailWithContext(context.Background(), request)
}

// SendOrgMemberAccountBindEmail
// 重新发送成员绑定邮箱激活邮件
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_SENDEMAILLIMIT = "LimitExceeded.SendEmailLimit"
//  LIMITEXCEEDED_SENDEMAILWITHINONEHOURLIMIT = "LimitExceeded.SendEmailWithinOneHourLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) SendOrgMemberAccountBindEmailWithContext(ctx context.Context, request *SendOrgMemberAccountBindEmailRequest) (response *SendOrgMemberAccountBindEmailResponse, err error) {
    if request == nil {
        request = NewSendOrgMemberAccountBindEmailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "SendOrgMemberAccountBindEmail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendOrgMemberAccountBindEmail require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendOrgMemberAccountBindEmailResponse()
    err = c.Send(request, response)
    return
}

func NewSetExternalSAMLIdentityProviderRequest() (request *SetExternalSAMLIdentityProviderRequest) {
    request = &SetExternalSAMLIdentityProviderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "SetExternalSAMLIdentityProvider")
    
    
    return
}

func NewSetExternalSAMLIdentityProviderResponse() (response *SetExternalSAMLIdentityProviderResponse) {
    response = &SetExternalSAMLIdentityProviderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetExternalSAMLIdentityProvider
// 配置SAML身份提供商信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DECODEMETADATADOCUMENTFAILED = "FailedOperation.DecodeMetadataDocumentFailed"
//  FAILEDOPERATION_METADATACOSPARSINGFAILED = "FailedOperation.MetadataCosParsingFailed"
//  FAILEDOPERATION_SAMLSERVICEPROVIDERCREATEFAILED = "FailedOperation.SAMLServiceProviderCreateFailed"
//  FAILEDOPERATION_UPLOADMETADATAFAILED = "FailedOperation.UploadMetadataFailed"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  FAILEDOPERATION_XMLDATAUNMARSHALFAILED = "FailedOperation.XMLDataUnmarshalFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SSOSTATUSINVALID = "InvalidParameterValue.SSoStatusInvalid"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) SetExternalSAMLIdentityProvider(request *SetExternalSAMLIdentityProviderRequest) (response *SetExternalSAMLIdentityProviderResponse, err error) {
    return c.SetExternalSAMLIdentityProviderWithContext(context.Background(), request)
}

// SetExternalSAMLIdentityProvider
// 配置SAML身份提供商信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_DECODEMETADATADOCUMENTFAILED = "FailedOperation.DecodeMetadataDocumentFailed"
//  FAILEDOPERATION_METADATACOSPARSINGFAILED = "FailedOperation.MetadataCosParsingFailed"
//  FAILEDOPERATION_SAMLSERVICEPROVIDERCREATEFAILED = "FailedOperation.SAMLServiceProviderCreateFailed"
//  FAILEDOPERATION_UPLOADMETADATAFAILED = "FailedOperation.UploadMetadataFailed"
//  FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//  FAILEDOPERATION_XMLDATAUNMARSHALFAILED = "FailedOperation.XMLDataUnmarshalFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_SSOSTATUSINVALID = "InvalidParameterValue.SSoStatusInvalid"
//  INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) SetExternalSAMLIdentityProviderWithContext(ctx context.Context, request *SetExternalSAMLIdentityProviderRequest) (response *SetExternalSAMLIdentityProviderResponse, err error) {
    if request == nil {
        request = NewSetExternalSAMLIdentityProviderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "SetExternalSAMLIdentityProvider")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetExternalSAMLIdentityProvider require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetExternalSAMLIdentityProviderResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCustomPolicyForRoleConfigurationRequest() (request *UpdateCustomPolicyForRoleConfigurationRequest) {
    request = &UpdateCustomPolicyForRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateCustomPolicyForRoleConfiguration")
    
    
    return
}

func NewUpdateCustomPolicyForRoleConfigurationResponse() (response *UpdateCustomPolicyForRoleConfigurationResponse) {
    response = &UpdateCustomPolicyForRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateCustomPolicyForRoleConfiguration
// 为权限配置修改自定义策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYTYPEERROR = "InvalidParameter.PolicyTypeError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) UpdateCustomPolicyForRoleConfiguration(request *UpdateCustomPolicyForRoleConfigurationRequest) (response *UpdateCustomPolicyForRoleConfigurationResponse, err error) {
    return c.UpdateCustomPolicyForRoleConfigurationWithContext(context.Background(), request)
}

// UpdateCustomPolicyForRoleConfiguration
// 为权限配置修改自定义策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//  INVALIDPARAMETER_POLICYTYPEERROR = "InvalidParameter.PolicyTypeError"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//  INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) UpdateCustomPolicyForRoleConfigurationWithContext(ctx context.Context, request *UpdateCustomPolicyForRoleConfigurationRequest) (response *UpdateCustomPolicyForRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateCustomPolicyForRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateCustomPolicyForRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCustomPolicyForRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCustomPolicyForRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
    request = &UpdateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateGroup")
    
    
    return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
    response = &UpdateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateGroup
// 修改用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    return c.UpdateGroupWithContext(context.Background(), request)
}

// UpdateGroup
// 修改用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
    if request == nil {
        request = NewUpdateGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateGroupResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationIdentity")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationMember")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationMemberEmailBind")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMemberEmailBind require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMemberEmailBindResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateOrganizationMembersPolicyRequest() (request *UpdateOrganizationMembersPolicyRequest) {
    request = &UpdateOrganizationMembersPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMembersPolicy")
    
    
    return
}

func NewUpdateOrganizationMembersPolicyResponse() (response *UpdateOrganizationMembersPolicyResponse) {
    response = &UpdateOrganizationMembersPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateOrganizationMembersPolicy
// 修改组织成员访问策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMembersPolicy(request *UpdateOrganizationMembersPolicyRequest) (response *UpdateOrganizationMembersPolicyResponse, err error) {
    return c.UpdateOrganizationMembersPolicyWithContext(context.Background(), request)
}

// UpdateOrganizationMembersPolicy
// 修改组织成员访问策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMembersPolicyWithContext(ctx context.Context, request *UpdateOrganizationMembersPolicyRequest) (response *UpdateOrganizationMembersPolicyResponse, err error) {
    if request == nil {
        request = NewUpdateOrganizationMembersPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationMembersPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateOrganizationMembersPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateOrganizationMembersPolicyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateOrganizationNode")
    
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdatePolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePolicyResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRoleConfigurationRequest() (request *UpdateRoleConfigurationRequest) {
    request = &UpdateRoleConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateRoleConfiguration")
    
    
    return
}

func NewUpdateRoleConfigurationResponse() (response *UpdateRoleConfigurationResponse) {
    response = &UpdateRoleConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateRoleConfiguration
// 修改权限配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) UpdateRoleConfiguration(request *UpdateRoleConfigurationRequest) (response *UpdateRoleConfigurationResponse, err error) {
    return c.UpdateRoleConfigurationWithContext(context.Background(), request)
}

// UpdateRoleConfiguration
// 修改权限配置信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) UpdateRoleConfigurationWithContext(ctx context.Context, request *UpdateRoleConfigurationRequest) (response *UpdateRoleConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateRoleConfigurationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateRoleConfiguration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRoleConfiguration require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRoleConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSCIMCredentialStatusRequest() (request *UpdateSCIMCredentialStatusRequest) {
    request = &UpdateSCIMCredentialStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateSCIMCredentialStatus")
    
    
    return
}

func NewUpdateSCIMCredentialStatusResponse() (response *UpdateSCIMCredentialStatusResponse) {
    response = &UpdateSCIMCredentialStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSCIMCredentialStatus
// 启用/禁用SCIM密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
//  INVALIDPARAMETER_USERSCIMCREDENTIALSTATUSERROR = "InvalidParameter.UserScimCredentialStatusError"
func (c *Client) UpdateSCIMCredentialStatus(request *UpdateSCIMCredentialStatusRequest) (response *UpdateSCIMCredentialStatusResponse, err error) {
    return c.UpdateSCIMCredentialStatusWithContext(context.Background(), request)
}

// UpdateSCIMCredentialStatus
// 启用/禁用SCIM密钥
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
//  INVALIDPARAMETER_USERSCIMCREDENTIALSTATUSERROR = "InvalidParameter.UserScimCredentialStatusError"
func (c *Client) UpdateSCIMCredentialStatusWithContext(ctx context.Context, request *UpdateSCIMCredentialStatusRequest) (response *UpdateSCIMCredentialStatusResponse, err error) {
    if request == nil {
        request = NewUpdateSCIMCredentialStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateSCIMCredentialStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSCIMCredentialStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSCIMCredentialStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateSCIMSynchronizationStatusRequest() (request *UpdateSCIMSynchronizationStatusRequest) {
    request = &UpdateSCIMSynchronizationStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateSCIMSynchronizationStatus")
    
    
    return
}

func NewUpdateSCIMSynchronizationStatusResponse() (response *UpdateSCIMSynchronizationStatusResponse) {
    response = &UpdateSCIMSynchronizationStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateSCIMSynchronizationStatus
// 启用/禁用用户SCIM同步
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMSYNCSTATUSERROR = "InvalidParameter.ScimSyncStatusError"
func (c *Client) UpdateSCIMSynchronizationStatus(request *UpdateSCIMSynchronizationStatusRequest) (response *UpdateSCIMSynchronizationStatusResponse, err error) {
    return c.UpdateSCIMSynchronizationStatusWithContext(context.Background(), request)
}

// UpdateSCIMSynchronizationStatus
// 启用/禁用用户SCIM同步
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_SCIMSYNCSTATUSERROR = "InvalidParameter.ScimSyncStatusError"
func (c *Client) UpdateSCIMSynchronizationStatusWithContext(ctx context.Context, request *UpdateSCIMSynchronizationStatusRequest) (response *UpdateSCIMSynchronizationStatusResponse, err error) {
    if request == nil {
        request = NewUpdateSCIMSynchronizationStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateSCIMSynchronizationStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateSCIMSynchronizationStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateSCIMSynchronizationStatusResponse()
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
//  FAILEDOPERATION_EXISTSHAREMEMBERNOTINORGANIZATION = "FailedOperation.ExistShareMemberNotInOrganization"
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
//  FAILEDOPERATION_EXISTSHAREMEMBERNOTINORGANIZATION = "FailedOperation.ExistShareMemberNotInOrganization"
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
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateShareUnit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateShareUnit require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateShareUnitResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
    request = &UpdateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateUser")
    
    
    return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
    response = &UpdateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUser
// 修改用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERTYPEERROR = "InvalidParameter.UserTypeError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    return c.UpdateUserWithContext(context.Background(), request)
}

// UpdateUser
// 修改用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//  INVALIDPARAMETER_USERTYPEERROR = "InvalidParameter.UserTypeError"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
    if request == nil {
        request = NewUpdateUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserStatusRequest() (request *UpdateUserStatusRequest) {
    request = &UpdateUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateUserStatus")
    
    
    return
}

func NewUpdateUserStatusResponse() (response *UpdateUserStatusResponse) {
    response = &UpdateUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserStatus
// 修改用户状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    return c.UpdateUserStatusWithContext(context.Background(), request)
}

// UpdateUserStatus
// 修改用户状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserStatusWithContext(ctx context.Context, request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
    if request == nil {
        request = NewUpdateUserStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateUserStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserStatusResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateUserSyncProvisioningRequest() (request *UpdateUserSyncProvisioningRequest) {
    request = &UpdateUserSyncProvisioningRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateUserSyncProvisioning")
    
    
    return
}

func NewUpdateUserSyncProvisioningResponse() (response *UpdateUserSyncProvisioningResponse) {
    response = &UpdateUserSyncProvisioningResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateUserSyncProvisioning
// 创建子用户同步任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) UpdateUserSyncProvisioning(request *UpdateUserSyncProvisioningRequest) (response *UpdateUserSyncProvisioningResponse, err error) {
    return c.UpdateUserSyncProvisioningWithContext(context.Background(), request)
}

// UpdateUserSyncProvisioning
// 创建子用户同步任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) UpdateUserSyncProvisioningWithContext(ctx context.Context, request *UpdateUserSyncProvisioningRequest) (response *UpdateUserSyncProvisioningResponse, err error) {
    if request == nil {
        request = NewUpdateUserSyncProvisioningRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateUserSyncProvisioning")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateUserSyncProvisioning require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateUserSyncProvisioningResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateZoneRequest() (request *UpdateZoneRequest) {
    request = &UpdateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("organization", APIVersion, "UpdateZone")
    
    
    return
}

func NewUpdateZoneResponse() (response *UpdateZoneResponse) {
    response = &UpdateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateZone
// 更新用户空间名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETERVALUE_ZONENAMEFORMATERROR = "InvalidParameterValue.ZoneNameFormatError"
func (c *Client) UpdateZone(request *UpdateZoneRequest) (response *UpdateZoneResponse, err error) {
    return c.UpdateZoneWithContext(context.Background(), request)
}

// UpdateZone
// 更新用户空间名
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETERVALUE_ZONENAMEFORMATERROR = "InvalidParameterValue.ZoneNameFormatError"
func (c *Client) UpdateZoneWithContext(ctx context.Context, request *UpdateZoneRequest) (response *UpdateZoneResponse, err error) {
    if request == nil {
        request = NewUpdateZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "organization", APIVersion, "UpdateZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateZoneResponse()
    err = c.Send(request, response)
    return
}
