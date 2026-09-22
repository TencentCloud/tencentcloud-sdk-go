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

package v20260709

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2026-07-09"

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


func NewBindExternalAgentRequest() (request *BindExternalAgentRequest) {
    request = &BindExternalAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "BindExternalAgent")
    
    
    return
}

func NewBindExternalAgentResponse() (response *BindExternalAgentResponse) {
    response = &BindExternalAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindExternalAgent
// 把外部 agent 绑定到某 managed agent
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) BindExternalAgent(request *BindExternalAgentRequest) (response *BindExternalAgentResponse, err error) {
    return c.BindExternalAgentWithContext(context.Background(), request)
}

// BindExternalAgent
// 把外部 agent 绑定到某 managed agent
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) BindExternalAgentWithContext(ctx context.Context, request *BindExternalAgentRequest) (response *BindExternalAgentResponse, err error) {
    if request == nil {
        request = NewBindExternalAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "BindExternalAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindExternalAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindExternalAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentRequest() (request *CreateAgentRequest) {
    request = &CreateAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "CreateAgent")
    
    
    return
}

func NewCreateAgentResponse() (response *CreateAgentResponse) {
    response = &CreateAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgent
// 创建一个新的 Managed Agent，同时自动生成 default 版本。配置采用 Manifest v2.0。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTACCOUNTIDINVALID = "InvalidParameter.AgentAccountIDInvalid"
//  INVALIDPARAMETER_AGENTNAMELENGTH = "InvalidParameter.AgentNameLength"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCEINUSE_AGENTIDCONFLICT = "ResourceInUse.AgentIDConflict"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTAUTHORIZED = "UnauthorizedOperation.AccountNotAuthorized"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgent(request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    return c.CreateAgentWithContext(context.Background(), request)
}

// CreateAgent
// 创建一个新的 Managed Agent，同时自动生成 default 版本。配置采用 Manifest v2.0。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTACCOUNTIDINVALID = "InvalidParameter.AgentAccountIDInvalid"
//  INVALIDPARAMETER_AGENTNAMELENGTH = "InvalidParameter.AgentNameLength"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCEINUSE_AGENTIDCONFLICT = "ResourceInUse.AgentIDConflict"
//  UNAUTHORIZEDOPERATION_ACCOUNTNOTAUTHORIZED = "UnauthorizedOperation.AccountNotAuthorized"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentWithContext(ctx context.Context, request *CreateAgentRequest) (response *CreateAgentResponse, err error) {
    if request == nil {
        request = NewCreateAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "CreateAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentSessionRequest() (request *CreateAgentSessionRequest) {
    request = &CreateAgentSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "CreateAgentSession")
    
    
    return
}

func NewCreateAgentSessionResponse() (response *CreateAgentSessionResponse) {
    response = &CreateAgentSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentSession
// 为指定 Agent 创建新的会话，返回会话 ID 和聊天凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_AGENTNOROUTINGCONFIG = "FailedOperation.AgentNoRoutingConfig"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_RUNTIMECREATEFAILED = "InternalError.RuntimeCreateFailed"
//  INTERNALERROR_SESSIONOPERATIONFAILED = "InternalError.SessionOperationFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_AGENTNOPUBLISHEDVERSION = "InvalidParameter.AgentNoPublishedVersion"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCEUNAVAILABLE_CHATTOKENUNAVAILABLE = "ResourceUnavailable.ChatTokenUnavailable"
//  RESOURCEUNAVAILABLE_RUNTIMEAPIUNAVAILABLE = "ResourceUnavailable.RuntimeAPIUnavailable"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentSession(request *CreateAgentSessionRequest) (response *CreateAgentSessionResponse, err error) {
    return c.CreateAgentSessionWithContext(context.Background(), request)
}

// CreateAgentSession
// 为指定 Agent 创建新的会话，返回会话 ID 和聊天凭证。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_AGENTNOROUTINGCONFIG = "FailedOperation.AgentNoRoutingConfig"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_RUNTIMECREATEFAILED = "InternalError.RuntimeCreateFailed"
//  INTERNALERROR_SESSIONOPERATIONFAILED = "InternalError.SessionOperationFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_AGENTNOPUBLISHEDVERSION = "InvalidParameter.AgentNoPublishedVersion"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCEUNAVAILABLE_CHATTOKENUNAVAILABLE = "ResourceUnavailable.ChatTokenUnavailable"
//  RESOURCEUNAVAILABLE_RUNTIMEAPIUNAVAILABLE = "ResourceUnavailable.RuntimeAPIUnavailable"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentSessionWithContext(ctx context.Context, request *CreateAgentSessionRequest) (response *CreateAgentSessionResponse, err error) {
    if request == nil {
        request = NewCreateAgentSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "CreateAgentSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentVersionRequest() (request *CreateAgentVersionRequest) {
    request = &CreateAgentVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "CreateAgentVersion")
    
    
    return
}

func NewCreateAgentVersionResponse() (response *CreateAgentVersionResponse) {
    response = &CreateAgentVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentVersion
// 完全新建版本：外部准备完整 Manifest 后直接传入，不引用任何已有版本。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentVersion(request *CreateAgentVersionRequest) (response *CreateAgentVersionResponse, err error) {
    return c.CreateAgentVersionWithContext(context.Background(), request)
}

// CreateAgentVersion
// 完全新建版本：外部准备完整 Manifest 后直接传入，不引用任何已有版本。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentVersionWithContext(ctx context.Context, request *CreateAgentVersionRequest) (response *CreateAgentVersionResponse, err error) {
    if request == nil {
        request = NewCreateAgentVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "CreateAgentVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentVersionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentVersionFromSourceRequest() (request *CreateAgentVersionFromSourceRequest) {
    request = &CreateAgentVersionFromSourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "CreateAgentVersionFromSource")
    
    
    return
}

func NewCreateAgentVersionFromSourceResponse() (response *CreateAgentVersionFromSourceResponse) {
    response = &CreateAgentVersionFromSourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAgentVersionFromSource
// 基于源版本创建新版本：Manifest / Model / Description 传入即整体覆盖，未传则沿用源版本。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentVersionFromSource(request *CreateAgentVersionFromSourceRequest) (response *CreateAgentVersionFromSourceResponse, err error) {
    return c.CreateAgentVersionFromSourceWithContext(context.Background(), request)
}

// CreateAgentVersionFromSource
// 基于源版本创建新版本：Manifest / Model / Description 传入即整体覆盖，未传则沿用源版本。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) CreateAgentVersionFromSourceWithContext(ctx context.Context, request *CreateAgentVersionFromSourceRequest) (response *CreateAgentVersionFromSourceResponse, err error) {
    if request == nil {
        request = NewCreateAgentVersionFromSourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "CreateAgentVersionFromSource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAgentVersionFromSource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAgentVersionFromSourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentRequest() (request *DeleteAgentRequest) {
    request = &DeleteAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DeleteAgent")
    
    
    return
}

func NewDeleteAgentResponse() (response *DeleteAgentResponse) {
    response = &DeleteAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAgent
// 删除指定的 Agent 及其所有版本。删除后不可恢复。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DeleteAgent(request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
    return c.DeleteAgentWithContext(context.Background(), request)
}

// DeleteAgent
// 删除指定的 Agent 及其所有版本。删除后不可恢复。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DeleteAgentWithContext(ctx context.Context, request *DeleteAgentRequest) (response *DeleteAgentResponse, err error) {
    if request == nil {
        request = NewDeleteAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DeleteAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentRequest() (request *DescribeAgentRequest) {
    request = &DescribeAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeAgent")
    
    
    return
}

func NewDescribeAgentResponse() (response *DescribeAgentResponse) {
    response = &DescribeAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgent
// 查询单个 Agent 的详细信息，包括基础配置和路由配置。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgent(request *DescribeAgentRequest) (response *DescribeAgentResponse, err error) {
    return c.DescribeAgentWithContext(context.Background(), request)
}

// DescribeAgent
// 查询单个 Agent 的详细信息，包括基础配置和路由配置。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentWithContext(ctx context.Context, request *DescribeAgentRequest) (response *DescribeAgentResponse, err error) {
    if request == nil {
        request = NewDescribeAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentListRequest() (request *DescribeAgentListRequest) {
    request = &DescribeAgentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeAgentList")
    
    
    return
}

func NewDescribeAgentListResponse() (response *DescribeAgentListResponse) {
    response = &DescribeAgentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentList
// 查询当前企业的 Agent 列表，支持分页、过滤和排序。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentList(request *DescribeAgentListRequest) (response *DescribeAgentListResponse, err error) {
    return c.DescribeAgentListWithContext(context.Background(), request)
}

// DescribeAgentList
// 查询当前企业的 Agent 列表，支持分页、过滤和排序。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentListWithContext(ctx context.Context, request *DescribeAgentListRequest) (response *DescribeAgentListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeAgentList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentSessionRequest() (request *DescribeAgentSessionRequest) {
    request = &DescribeAgentSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeAgentSession")
    
    
    return
}

func NewDescribeAgentSessionResponse() (response *DescribeAgentSessionResponse) {
    response = &DescribeAgentSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentSession
// 查询单个 Agent 会话详情：返回会话基础信息（会话名称 / Agent / 版本 / 状态 / 来源 / 发起人）与可用的聊天接入点列表（EndpointSet）。数据面鉴权走 DescribeUserAccessToken 的用户级访问令牌。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentSession(request *DescribeAgentSessionRequest) (response *DescribeAgentSessionResponse, err error) {
    return c.DescribeAgentSessionWithContext(context.Background(), request)
}

// DescribeAgentSession
// 查询单个 Agent 会话详情：返回会话基础信息（会话名称 / Agent / 版本 / 状态 / 来源 / 发起人）与可用的聊天接入点列表（EndpointSet）。数据面鉴权走 DescribeUserAccessToken 的用户级访问令牌。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentSessionWithContext(ctx context.Context, request *DescribeAgentSessionRequest) (response *DescribeAgentSessionResponse, err error) {
    if request == nil {
        request = NewDescribeAgentSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeAgentSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentSessionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentSessionListRequest() (request *DescribeAgentSessionListRequest) {
    request = &DescribeAgentSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeAgentSessionList")
    
    
    return
}

func NewDescribeAgentSessionListResponse() (response *DescribeAgentSessionListResponse) {
    response = &DescribeAgentSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentSessionList
// 分页查询企业下所有会话（跨 Agent）：支持按 SessionId / Status / AgentId / UserId 过滤，按创建 / 更新时间排序，返回会话摘要列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentSessionList(request *DescribeAgentSessionListRequest) (response *DescribeAgentSessionListResponse, err error) {
    return c.DescribeAgentSessionListWithContext(context.Background(), request)
}

// DescribeAgentSessionList
// 分页查询企业下所有会话（跨 Agent）：支持按 SessionId / Status / AgentId / UserId 过滤，按创建 / 更新时间排序，返回会话摘要列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentSessionListWithContext(ctx context.Context, request *DescribeAgentSessionListRequest) (response *DescribeAgentSessionListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentSessionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeAgentSessionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentSessionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentVersionRequest() (request *DescribeAgentVersionRequest) {
    request = &DescribeAgentVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeAgentVersion")
    
    
    return
}

func NewDescribeAgentVersionResponse() (response *DescribeAgentVersionResponse) {
    response = &DescribeAgentVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentVersion
// 查询单个版本的详细信息，包括 Manifest、Model、状态等。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentVersion(request *DescribeAgentVersionRequest) (response *DescribeAgentVersionResponse, err error) {
    return c.DescribeAgentVersionWithContext(context.Background(), request)
}

// DescribeAgentVersion
// 查询单个版本的详细信息，包括 Manifest、Model、状态等。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentVersionWithContext(ctx context.Context, request *DescribeAgentVersionRequest) (response *DescribeAgentVersionResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeAgentVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgentVersionListRequest() (request *DescribeAgentVersionListRequest) {
    request = &DescribeAgentVersionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeAgentVersionList")
    
    
    return
}

func NewDescribeAgentVersionListResponse() (response *DescribeAgentVersionListResponse) {
    response = &DescribeAgentVersionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAgentVersionList
// 查询指定 Agent 下的版本列表，支持分页和版本类型过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentVersionList(request *DescribeAgentVersionListRequest) (response *DescribeAgentVersionListResponse, err error) {
    return c.DescribeAgentVersionListWithContext(context.Background(), request)
}

// DescribeAgentVersionList
// 查询指定 Agent 下的版本列表，支持分页和版本类型过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeAgentVersionListWithContext(ctx context.Context, request *DescribeAgentVersionListRequest) (response *DescribeAgentVersionListResponse, err error) {
    if request == nil {
        request = NewDescribeAgentVersionListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeAgentVersionList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentVersionList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentVersionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBuiltinModelListRequest() (request *DescribeBuiltinModelListRequest) {
    request = &DescribeBuiltinModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeBuiltinModelList")
    
    
    return
}

func NewDescribeBuiltinModelListResponse() (response *DescribeBuiltinModelListResponse) {
    response = &DescribeBuiltinModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBuiltinModelList
// 查询当前企业的内置模型列表，支持分页与过滤。内置模型由平台预置，企业可按需启用/停用。过滤字段支持：ModelId（模型ID，模糊）、Name（模型名称，模糊）、Vendor（供应商，模糊）、Status（状态，精确：enabled/disabled）。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_MODELMANAGEMENTFAILED = "InternalError.ModelManagementFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MODELMANAGEMENT = "InvalidParameter.ModelManagement"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
func (c *Client) DescribeBuiltinModelList(request *DescribeBuiltinModelListRequest) (response *DescribeBuiltinModelListResponse, err error) {
    return c.DescribeBuiltinModelListWithContext(context.Background(), request)
}

// DescribeBuiltinModelList
// 查询当前企业的内置模型列表，支持分页与过滤。内置模型由平台预置，企业可按需启用/停用。过滤字段支持：ModelId（模型ID，模糊）、Name（模型名称，模糊）、Vendor（供应商，模糊）、Status（状态，精确：enabled/disabled）。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_MODELMANAGEMENTFAILED = "InternalError.ModelManagementFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MODELMANAGEMENT = "InvalidParameter.ModelManagement"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
func (c *Client) DescribeBuiltinModelListWithContext(ctx context.Context, request *DescribeBuiltinModelListRequest) (response *DescribeBuiltinModelListResponse, err error) {
    if request == nil {
        request = NewDescribeBuiltinModelListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeBuiltinModelList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBuiltinModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBuiltinModelListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConnectorListRequest() (request *DescribeConnectorListRequest) {
    request = &DescribeConnectorListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeConnectorList")
    
    
    return
}

func NewDescribeConnectorListResponse() (response *DescribeConnectorListResponse) {
    response = &DescribeConnectorListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConnectorList
// 查询指定企业下的连接器列表（PageNumber/PageSize 分页，支持名称模糊与状态、来源过滤）。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_IDENTITYSERVICEERROR = "FailedOperation.IdentityServiceError"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_CONNECTORREQUEST = "InvalidParameter.ConnectorRequest"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_ONEIDACCOUNTREQUIRED = "InvalidParameter.OneIDAccountRequired"
//  INVALIDPARAMETERVALUE_INVALIDCONNECTORSOURCE = "InvalidParameterValue.InvalidConnectorSource"
//  INVALIDPARAMETERVALUE_INVALIDCONNECTORSTATUS = "InvalidParameterValue.InvalidConnectorStatus"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCEUNAVAILABLE_IDENTITYNOTCONFIGURED = "ResourceUnavailable.IdentityNotConfigured"
//  UNAUTHORIZEDOPERATION_NOTCONNECTOROWNER = "UnauthorizedOperation.NotConnectorOwner"
func (c *Client) DescribeConnectorList(request *DescribeConnectorListRequest) (response *DescribeConnectorListResponse, err error) {
    return c.DescribeConnectorListWithContext(context.Background(), request)
}

// DescribeConnectorList
// 查询指定企业下的连接器列表（PageNumber/PageSize 分页，支持名称模糊与状态、来源过滤）。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_IDENTITYSERVICEERROR = "FailedOperation.IdentityServiceError"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_CONNECTORREQUEST = "InvalidParameter.ConnectorRequest"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_ONEIDACCOUNTREQUIRED = "InvalidParameter.OneIDAccountRequired"
//  INVALIDPARAMETERVALUE_INVALIDCONNECTORSOURCE = "InvalidParameterValue.InvalidConnectorSource"
//  INVALIDPARAMETERVALUE_INVALIDCONNECTORSTATUS = "InvalidParameterValue.InvalidConnectorStatus"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCEUNAVAILABLE_IDENTITYNOTCONFIGURED = "ResourceUnavailable.IdentityNotConfigured"
//  UNAUTHORIZEDOPERATION_NOTCONNECTOROWNER = "UnauthorizedOperation.NotConnectorOwner"
func (c *Client) DescribeConnectorListWithContext(ctx context.Context, request *DescribeConnectorListRequest) (response *DescribeConnectorListResponse, err error) {
    if request == nil {
        request = NewDescribeConnectorListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeConnectorList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConnectorList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConnectorListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExpertListRequest() (request *DescribeExpertListRequest) {
    request = &DescribeExpertListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeExpertList")
    
    
    return
}

func NewDescribeExpertListResponse() (response *DescribeExpertListResponse) {
    response = &DescribeExpertListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExpertList
// 分页查询 Expert 列表，支持关键词、分类、发布状态过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENTERPRISEIDREQUIRED = "InvalidParameter.EnterpriseIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  INVALIDPARAMETER_SOURCEREQUIRED = "InvalidParameter.SourceRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
func (c *Client) DescribeExpertList(request *DescribeExpertListRequest) (response *DescribeExpertListResponse, err error) {
    return c.DescribeExpertListWithContext(context.Background(), request)
}

// DescribeExpertList
// 分页查询 Expert 列表，支持关键词、分类、发布状态过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENTERPRISEIDREQUIRED = "InvalidParameter.EnterpriseIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  INVALIDPARAMETER_SOURCEREQUIRED = "InvalidParameter.SourceRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
func (c *Client) DescribeExpertListWithContext(ctx context.Context, request *DescribeExpertListRequest) (response *DescribeExpertListResponse, err error) {
    if request == nil {
        request = NewDescribeExpertListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeExpertList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExpertList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExpertListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalAgentRequest() (request *DescribeExternalAgentRequest) {
    request = &DescribeExternalAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeExternalAgent")
    
    
    return
}

func NewDescribeExternalAgentResponse() (response *DescribeExternalAgentResponse) {
    response = &DescribeExternalAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExternalAgent
// 查询某 managed agent 绑定的单个外部 agent 详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeExternalAgent(request *DescribeExternalAgentRequest) (response *DescribeExternalAgentResponse, err error) {
    return c.DescribeExternalAgentWithContext(context.Background(), request)
}

// DescribeExternalAgent
// 查询某 managed agent 绑定的单个外部 agent 详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeExternalAgentWithContext(ctx context.Context, request *DescribeExternalAgentRequest) (response *DescribeExternalAgentResponse, err error) {
    if request == nil {
        request = NewDescribeExternalAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeExternalAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalAgentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExternalAgentListRequest() (request *DescribeExternalAgentListRequest) {
    request = &DescribeExternalAgentListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeExternalAgentList")
    
    
    return
}

func NewDescribeExternalAgentListResponse() (response *DescribeExternalAgentListResponse) {
    response = &DescribeExternalAgentListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExternalAgentList
// 列某 managed agent 绑定的外部 agent 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeExternalAgentList(request *DescribeExternalAgentListRequest) (response *DescribeExternalAgentListResponse, err error) {
    return c.DescribeExternalAgentListWithContext(context.Background(), request)
}

// DescribeExternalAgentList
// 列某 managed agent 绑定的外部 agent 列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) DescribeExternalAgentListWithContext(ctx context.Context, request *DescribeExternalAgentListRequest) (response *DescribeExternalAgentListResponse, err error) {
    if request == nil {
        request = NewDescribeExternalAgentListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeExternalAgentList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExternalAgentList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExternalAgentListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageEventListRequest() (request *DescribeMessageEventListRequest) {
    request = &DescribeMessageEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeMessageEventList")
    
    
    return
}

func NewDescribeMessageEventListResponse() (response *DescribeMessageEventListResponse) {
    response = &DescribeMessageEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMessageEventList
// 按 Session 分页查询消息事件
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_SESSIONHISTORYNOTREADY = "FailedOperation.SessionHistoryNotReady"
//  FAILEDOPERATION_SESSIONHISTORYUNAVAILABLE = "FailedOperation.SessionHistoryUnavailable"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_SESSIONHISTORYFAILED = "InternalError.SessionHistoryFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_TRACINGNOTCONFIGURED = "InternalError.TracingNotConfigured"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  INVALIDPARAMETER_SESSIONIDREQUIRED = "InvalidParameter.SessionIDRequired"
//  INVALIDPARAMETER_TRACINGREQUEST = "InvalidParameter.TracingRequest"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) DescribeMessageEventList(request *DescribeMessageEventListRequest) (response *DescribeMessageEventListResponse, err error) {
    return c.DescribeMessageEventListWithContext(context.Background(), request)
}

// DescribeMessageEventList
// 按 Session 分页查询消息事件
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_SESSIONHISTORYNOTREADY = "FailedOperation.SessionHistoryNotReady"
//  FAILEDOPERATION_SESSIONHISTORYUNAVAILABLE = "FailedOperation.SessionHistoryUnavailable"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_SESSIONHISTORYFAILED = "InternalError.SessionHistoryFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_TRACINGNOTCONFIGURED = "InternalError.TracingNotConfigured"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  INVALIDPARAMETER_SESSIONIDREQUIRED = "InvalidParameter.SessionIDRequired"
//  INVALIDPARAMETER_TRACINGREQUEST = "InvalidParameter.TracingRequest"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"
func (c *Client) DescribeMessageEventListWithContext(ctx context.Context, request *DescribeMessageEventListRequest) (response *DescribeMessageEventListResponse, err error) {
    if request == nil {
        request = NewDescribeMessageEventListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeMessageEventList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageEventListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillListRequest() (request *DescribeSkillListRequest) {
    request = &DescribeSkillListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "DescribeSkillList")
    
    
    return
}

func NewDescribeSkillListResponse() (response *DescribeSkillListResponse) {
    response = &DescribeSkillListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSkillList
// 分页查询 Skill 列表，支持关键词、分类、发布状态过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENTERPRISEIDREQUIRED = "InvalidParameter.EnterpriseIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  INVALIDPARAMETER_SOURCEREQUIRED = "InvalidParameter.SourceRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
func (c *Client) DescribeSkillList(request *DescribeSkillListRequest) (response *DescribeSkillListResponse, err error) {
    return c.DescribeSkillListWithContext(context.Background(), request)
}

// DescribeSkillList
// 分页查询 Skill 列表，支持关键词、分类、发布状态过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENTERPRISEIDREQUIRED = "InvalidParameter.EnterpriseIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"
//  INVALIDPARAMETER_SOURCEREQUIRED = "InvalidParameter.SourceRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
func (c *Client) DescribeSkillListWithContext(ctx context.Context, request *DescribeSkillListRequest) (response *DescribeSkillListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "DescribeSkillList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSkillList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSkillListResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateAgentSessionRequest() (request *MigrateAgentSessionRequest) {
    request = &MigrateAgentSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "MigrateAgentSession")
    
    
    return
}

func NewMigrateAgentSessionResponse() (response *MigrateAgentSessionResponse) {
    response = &MigrateAgentSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MigrateAgentSession
// 将指定会话迁移到目标版本。SessionID / RuntimeID 保持不变，通过 AgentOS UpdateSession 在原沙箱上更新 manifest 到新版本；AgentId 必须与原 Session 一致（禁止跨 Agent 迁移）；ChatToken 复用旧值不轮转。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_SESSIONCROSSAGENTMIGRATE = "FailedOperation.SessionCrossAgentMigrate"
//  FAILEDOPERATION_SESSIONNOTMIGRABLE = "FailedOperation.SessionNotMigrable"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_SESSIONOPERATIONFAILED = "InternalError.SessionOperationFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_SESSIONIDREQUIRED = "InvalidParameter.SessionIDRequired"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) MigrateAgentSession(request *MigrateAgentSessionRequest) (response *MigrateAgentSessionResponse, err error) {
    return c.MigrateAgentSessionWithContext(context.Background(), request)
}

// MigrateAgentSession
// 将指定会话迁移到目标版本。SessionID / RuntimeID 保持不变，通过 AgentOS UpdateSession 在原沙箱上更新 manifest 到新版本；AgentId 必须与原 Session 一致（禁止跨 Agent 迁移）；ChatToken 复用旧值不轮转。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_SESSIONCROSSAGENTMIGRATE = "FailedOperation.SessionCrossAgentMigrate"
//  FAILEDOPERATION_SESSIONNOTMIGRABLE = "FailedOperation.SessionNotMigrable"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_SESSIONOPERATIONFAILED = "InternalError.SessionOperationFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_SESSIONIDREQUIRED = "InvalidParameter.SessionIDRequired"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) MigrateAgentSessionWithContext(ctx context.Context, request *MigrateAgentSessionRequest) (response *MigrateAgentSessionResponse, err error) {
    if request == nil {
        request = NewMigrateAgentSessionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "MigrateAgentSession")
    
    if c.GetCredential() == nil {
        return nil, errors.New("MigrateAgentSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewMigrateAgentSessionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentRequest() (request *ModifyAgentRequest) {
    request = &ModifyAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "ModifyAgent")
    
    
    return
}

func NewModifyAgentResponse() (response *ModifyAgentResponse) {
    response = &ModifyAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgent
// 修改 Agent 基础信息（名称、描述、头像）。AgentName / Description / AvatarUrl 均为可选，仅传递需要更新的字段。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgent(request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    return c.ModifyAgentWithContext(context.Background(), request)
}

// ModifyAgent
// 修改 Agent 基础信息（名称、描述、头像）。AgentName / Description / AvatarUrl 均为可选，仅传递需要更新的字段。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentWithContext(ctx context.Context, request *ModifyAgentRequest) (response *ModifyAgentResponse, err error) {
    if request == nil {
        request = NewModifyAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "ModifyAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentA2AConfigRequest() (request *ModifyAgentA2AConfigRequest) {
    request = &ModifyAgentA2AConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "ModifyAgentA2AConfig")
    
    
    return
}

func NewModifyAgentA2AConfigResponse() (response *ModifyAgentA2AConfigResponse) {
    response = &ModifyAgentA2AConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentA2AConfig
// 修改 Agent 的 A2A 配置。A2AEnabled 是 Agent 级唯一开关，与具体版本和流量分发策略无关。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_A2AENABLEDREQUIRED = "InvalidParameter.A2AEnabledRequired"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentA2AConfig(request *ModifyAgentA2AConfigRequest) (response *ModifyAgentA2AConfigResponse, err error) {
    return c.ModifyAgentA2AConfigWithContext(context.Background(), request)
}

// ModifyAgentA2AConfig
// 修改 Agent 的 A2A 配置。A2AEnabled 是 Agent 级唯一开关，与具体版本和流量分发策略无关。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_A2AENABLEDREQUIRED = "InvalidParameter.A2AEnabledRequired"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentA2AConfigWithContext(ctx context.Context, request *ModifyAgentA2AConfigRequest) (response *ModifyAgentA2AConfigResponse, err error) {
    if request == nil {
        request = NewModifyAgentA2AConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "ModifyAgentA2AConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentA2AConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentA2AConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentRoutingRequest() (request *ModifyAgentRoutingRequest) {
    request = &ModifyAgentRoutingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "ModifyAgentRouting")
    
    
    return
}

func NewModifyAgentRoutingResponse() (response *ModifyAgentRoutingResponse) {
    response = &ModifyAgentRoutingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentRouting
// 覆盖式写入 Agent 路由配置（版本权重）。所有 VersionId 必须属于同一 Agent 且未弃用；允许空数组（下线 Agent 对外流量）；非空时权重总和须等于 1。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_ROUTINGDUPLICATEVERSION = "FailedOperation.RoutingDuplicateVersion"
//  FAILEDOPERATION_ROUTINGINVALIDVERSION = "FailedOperation.RoutingInvalidVersion"
//  FAILEDOPERATION_ROUTINGINVALIDWEIGHT = "FailedOperation.RoutingInvalidWeight"
//  FAILEDOPERATION_ROUTINGWEIGHTSUMINVALID = "FailedOperation.RoutingWeightSumInvalid"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentRouting(request *ModifyAgentRoutingRequest) (response *ModifyAgentRoutingResponse, err error) {
    return c.ModifyAgentRoutingWithContext(context.Background(), request)
}

// ModifyAgentRouting
// 覆盖式写入 Agent 路由配置（版本权重）。所有 VersionId 必须属于同一 Agent 且未弃用；允许空数组（下线 Agent 对外流量）；非空时权重总和须等于 1。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_ROUTINGDUPLICATEVERSION = "FailedOperation.RoutingDuplicateVersion"
//  FAILEDOPERATION_ROUTINGINVALIDVERSION = "FailedOperation.RoutingInvalidVersion"
//  FAILEDOPERATION_ROUTINGINVALIDWEIGHT = "FailedOperation.RoutingInvalidWeight"
//  FAILEDOPERATION_ROUTINGWEIGHTSUMINVALID = "FailedOperation.RoutingWeightSumInvalid"
//  INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentRoutingWithContext(ctx context.Context, request *ModifyAgentRoutingRequest) (response *ModifyAgentRoutingResponse, err error) {
    if request == nil {
        request = NewModifyAgentRoutingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "ModifyAgentRouting")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentRouting require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentRoutingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentVersionRequest() (request *ModifyAgentVersionRequest) {
    request = &ModifyAgentVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "ModifyAgentVersion")
    
    
    return
}

func NewModifyAgentVersionResponse() (response *ModifyAgentVersionResponse) {
    response = &ModifyAgentVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAgentVersion
// 原地更新 default 或 test 版本的 Manifest / Model / Description / SandboxTemplateId / ConnectorSet（prod 版本冻结不可修改），五个可选字段至少提供一个。ConnectorSet 为全量覆盖语义：缺省表示不改动连接器绑定；空数组表示解绑全部连接器。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_VERSIONNOTEDITABLE = "FailedOperation.VersionNotEditable"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"
//  INVALIDPARAMETER_NOUPDATABLEFIELD = "InvalidParameter.NoUpdatableField"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentVersion(request *ModifyAgentVersionRequest) (response *ModifyAgentVersionResponse, err error) {
    return c.ModifyAgentVersionWithContext(context.Background(), request)
}

// ModifyAgentVersion
// 原地更新 default 或 test 版本的 Manifest / Model / Description / SandboxTemplateId / ConnectorSet（prod 版本冻结不可修改），五个可选字段至少提供一个。ConnectorSet 为全量覆盖语义：缺省表示不改动连接器绑定；空数组表示解绑全部连接器。
//
// 可能返回的错误码:
//  AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"
//  AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_VERSIONNOTEDITABLE = "FailedOperation.VersionNotEditable"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_NORESOLVER = "InternalError.NoResolver"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"
//  INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"
//  INVALIDPARAMETER_NOUPDATABLEFIELD = "InvalidParameter.NoUpdatableField"
//  INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"
//  RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) ModifyAgentVersionWithContext(ctx context.Context, request *ModifyAgentVersionRequest) (response *ModifyAgentVersionResponse, err error) {
    if request == nil {
        request = NewModifyAgentVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "ModifyAgentVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAgentVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAgentVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindExternalAgentRequest() (request *UnbindExternalAgentRequest) {
    request = &UnbindExternalAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("workbuddyenterprise", APIVersion, "UnbindExternalAgent")
    
    
    return
}

func NewUnbindExternalAgentResponse() (response *UnbindExternalAgentResponse) {
    response = &UnbindExternalAgentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindExternalAgent
// 解除外部 agent 与 managed agent 的绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) UnbindExternalAgent(request *UnbindExternalAgentRequest) (response *UnbindExternalAgentResponse, err error) {
    return c.UnbindExternalAgentWithContext(context.Background(), request)
}

// UnbindExternalAgent
// 解除外部 agent 与 managed agent 的绑定
//
// 可能返回的错误码:
//  FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"
//  FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"
//  FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"
//  FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"
//  INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"
//  INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"
//  INTERNALERROR_PANIC = "InternalError.Panic"
//  INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"
//  INTERNALERROR_UNKNOWN = "InternalError.Unknown"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"
//  INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"
//  INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"
//  INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"
//  INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"
//  INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"
//  INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"
//  MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"
//  RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"
//  RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"
//  UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
func (c *Client) UnbindExternalAgentWithContext(ctx context.Context, request *UnbindExternalAgentRequest) (response *UnbindExternalAgentResponse, err error) {
    if request == nil {
        request = NewUnbindExternalAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "workbuddyenterprise", APIVersion, "UnbindExternalAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindExternalAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindExternalAgentResponse()
    err = c.Send(request, response)
    return
}
