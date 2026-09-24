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

const (
	// 此产品的特有错误码

	// AuthFailure.IdentityNotFound
	AUTHFAILURE_IDENTITYNOTFOUND = "AuthFailure.IdentityNotFound"

	// AuthFailure.IdentityResolutionFailed
	AUTHFAILURE_IDENTITYRESOLUTIONFAILED = "AuthFailure.IdentityResolutionFailed"

	// FailedOperation.AgentNoRoutingConfig
	FAILEDOPERATION_AGENTNOROUTINGCONFIG = "FailedOperation.AgentNoRoutingConfig"

	// FailedOperation.EntitlementNotFound
	FAILEDOPERATION_ENTITLEMENTNOTFOUND = "FailedOperation.EntitlementNotFound"

	// FailedOperation.ExternalBindFailed
	FAILEDOPERATION_EXTERNALBINDFAILED = "FailedOperation.ExternalBindFailed"

	// FailedOperation.ExternalOperationFailed
	FAILEDOPERATION_EXTERNALOPERATIONFAILED = "FailedOperation.ExternalOperationFailed"

	// FailedOperation.ExternalRegisterFailed
	FAILEDOPERATION_EXTERNALREGISTERFAILED = "FailedOperation.ExternalRegisterFailed"

	// FailedOperation.IdentityServiceError
	FAILEDOPERATION_IDENTITYSERVICEERROR = "FailedOperation.IdentityServiceError"

	// FailedOperation.RoutingDuplicateVersion
	FAILEDOPERATION_ROUTINGDUPLICATEVERSION = "FailedOperation.RoutingDuplicateVersion"

	// FailedOperation.RoutingInvalidVersion
	FAILEDOPERATION_ROUTINGINVALIDVERSION = "FailedOperation.RoutingInvalidVersion"

	// FailedOperation.RoutingInvalidWeight
	FAILEDOPERATION_ROUTINGINVALIDWEIGHT = "FailedOperation.RoutingInvalidWeight"

	// FailedOperation.RoutingWeightSumInvalid
	FAILEDOPERATION_ROUTINGWEIGHTSUMINVALID = "FailedOperation.RoutingWeightSumInvalid"

	// FailedOperation.SessionCrossAgentMigrate
	FAILEDOPERATION_SESSIONCROSSAGENTMIGRATE = "FailedOperation.SessionCrossAgentMigrate"

	// FailedOperation.SessionHistoryNotReady
	FAILEDOPERATION_SESSIONHISTORYNOTREADY = "FailedOperation.SessionHistoryNotReady"

	// FailedOperation.SessionHistoryUnavailable
	FAILEDOPERATION_SESSIONHISTORYUNAVAILABLE = "FailedOperation.SessionHistoryUnavailable"

	// FailedOperation.SessionNotMigrable
	FAILEDOPERATION_SESSIONNOTMIGRABLE = "FailedOperation.SessionNotMigrable"

	// FailedOperation.VersionNotEditable
	FAILEDOPERATION_VERSIONNOTEDITABLE = "FailedOperation.VersionNotEditable"

	// InternalError.AgentOperationFailed
	INTERNALERROR_AGENTOPERATIONFAILED = "InternalError.AgentOperationFailed"

	// InternalError.EntitlementCheckFailed
	INTERNALERROR_ENTITLEMENTCHECKFAILED = "InternalError.EntitlementCheckFailed"

	// InternalError.MarshalFailed
	INTERNALERROR_MARSHALFAILED = "InternalError.MarshalFailed"

	// InternalError.ModelManagementFailed
	INTERNALERROR_MODELMANAGEMENTFAILED = "InternalError.ModelManagementFailed"

	// InternalError.NoResolver
	INTERNALERROR_NORESOLVER = "InternalError.NoResolver"

	// InternalError.Panic
	INTERNALERROR_PANIC = "InternalError.Panic"

	// InternalError.ReadBodyFailed
	INTERNALERROR_READBODYFAILED = "InternalError.ReadBodyFailed"

	// InternalError.ResolveEnterpriseFailed
	INTERNALERROR_RESOLVEENTERPRISEFAILED = "InternalError.ResolveEnterpriseFailed"

	// InternalError.RuntimeCreateFailed
	INTERNALERROR_RUNTIMECREATEFAILED = "InternalError.RuntimeCreateFailed"

	// InternalError.SessionHistoryFailed
	INTERNALERROR_SESSIONHISTORYFAILED = "InternalError.SessionHistoryFailed"

	// InternalError.SessionOperationFailed
	INTERNALERROR_SESSIONOPERATIONFAILED = "InternalError.SessionOperationFailed"

	// InternalError.TenantResolveFailed
	INTERNALERROR_TENANTRESOLVEFAILED = "InternalError.TenantResolveFailed"

	// InternalError.TracingNotConfigured
	INTERNALERROR_TRACINGNOTCONFIGURED = "InternalError.TracingNotConfigured"

	// InternalError.Unknown
	INTERNALERROR_UNKNOWN = "InternalError.Unknown"

	// InternalError.UserAccessTokenExchangeFailed
	INTERNALERROR_USERACCESSTOKENEXCHANGEFAILED = "InternalError.UserAccessTokenExchangeFailed"

	// InternalError.UserAccessTokenNotConfigured
	INTERNALERROR_USERACCESSTOKENNOTCONFIGURED = "InternalError.UserAccessTokenNotConfigured"

	// InternalError.VersionOperationFailed
	INTERNALERROR_VERSIONOPERATIONFAILED = "InternalError.VersionOperationFailed"

	// InvalidParameter.A2AEnabledRequired
	INVALIDPARAMETER_A2AENABLEDREQUIRED = "InvalidParameter.A2AEnabledRequired"

	// InvalidParameter.ActionNotFound
	INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"

	// InvalidParameter.ActionRequired
	INVALIDPARAMETER_ACTIONREQUIRED = "InvalidParameter.ActionRequired"

	// InvalidParameter.AgentAccountIDInvalid
	INVALIDPARAMETER_AGENTACCOUNTIDINVALID = "InvalidParameter.AgentAccountIDInvalid"

	// InvalidParameter.AgentIDRequired
	INVALIDPARAMETER_AGENTIDREQUIRED = "InvalidParameter.AgentIDRequired"

	// Agent 名称长度需要在 2-32 位
	INVALIDPARAMETER_AGENTNAMELENGTH = "InvalidParameter.AgentNameLength"

	// InvalidParameter.AgentNoPublishedVersion
	INVALIDPARAMETER_AGENTNOPUBLISHEDVERSION = "InvalidParameter.AgentNoPublishedVersion"

	// InvalidParameter.ConnectorRequest
	INVALIDPARAMETER_CONNECTORREQUEST = "InvalidParameter.ConnectorRequest"

	// InvalidParameter.EmptyBody
	INVALIDPARAMETER_EMPTYBODY = "InvalidParameter.EmptyBody"

	// InvalidParameter.EndpointInvalid
	INVALIDPARAMETER_ENDPOINTINVALID = "InvalidParameter.EndpointInvalid"

	// InvalidParameter.EndpointRequired
	INVALIDPARAMETER_ENDPOINTREQUIRED = "InvalidParameter.EndpointRequired"

	// InvalidParameter.EnterpriseIDRequired
	INVALIDPARAMETER_ENTERPRISEIDREQUIRED = "InvalidParameter.EnterpriseIDRequired"

	// InvalidParameter.ExternalAgentIDRequired
	INVALIDPARAMETER_EXTERNALAGENTIDREQUIRED = "InvalidParameter.ExternalAgentIDRequired"

	// InvalidParameter.InvalidJSON
	INVALIDPARAMETER_INVALIDJSON = "InvalidParameter.InvalidJSON"

	// InvalidParameter.InvalidRequestBody
	INVALIDPARAMETER_INVALIDREQUESTBODY = "InvalidParameter.InvalidRequestBody"

	// InvalidParameter.ManifestRequired
	INVALIDPARAMETER_MANIFESTREQUIRED = "InvalidParameter.ManifestRequired"

	// InvalidParameter.ModelManagement
	INVALIDPARAMETER_MODELMANAGEMENT = "InvalidParameter.ModelManagement"

	// InvalidParameter.NoUpdatableField
	INVALIDPARAMETER_NOUPDATABLEFIELD = "InvalidParameter.NoUpdatableField"

	// InvalidParameter.OneIDAccountRequired
	INVALIDPARAMETER_ONEIDACCOUNTREQUIRED = "InvalidParameter.OneIDAccountRequired"

	// InvalidParameter.Pagination
	INVALIDPARAMETER_PAGINATION = "InvalidParameter.Pagination"

	// InvalidParameter.SessionIDRequired
	INVALIDPARAMETER_SESSIONIDREQUIRED = "InvalidParameter.SessionIDRequired"

	// InvalidParameter.SourceRequired
	INVALIDPARAMETER_SOURCEREQUIRED = "InvalidParameter.SourceRequired"

	// InvalidParameter.TracingRequest
	INVALIDPARAMETER_TRACINGREQUEST = "InvalidParameter.TracingRequest"

	// InvalidParameter.VersionIDRequired
	INVALIDPARAMETER_VERSIONIDREQUIRED = "InvalidParameter.VersionIDRequired"

	// InvalidParameterValue.InvalidConnectorSource
	INVALIDPARAMETERVALUE_INVALIDCONNECTORSOURCE = "InvalidParameterValue.InvalidConnectorSource"

	// InvalidParameterValue.InvalidConnectorStatus
	INVALIDPARAMETERVALUE_INVALIDCONNECTORSTATUS = "InvalidParameterValue.InvalidConnectorStatus"

	// MissingParameter.SubAccountUinRequired
	MISSINGPARAMETER_SUBACCOUNTUINREQUIRED = "MissingParameter.SubAccountUinRequired"

	// MissingParameter.UinRequired
	MISSINGPARAMETER_UINREQUIRED = "MissingParameter.UinRequired"

	// ResourceInUse.AgentIDConflict
	RESOURCEINUSE_AGENTIDCONFLICT = "ResourceInUse.AgentIDConflict"

	// ResourceNotFound.Agent
	RESOURCENOTFOUND_AGENT = "ResourceNotFound.Agent"

	// ResourceNotFound.ExternalAgent
	RESOURCENOTFOUND_EXTERNALAGENT = "ResourceNotFound.ExternalAgent"

	// ResourceNotFound.Session
	RESOURCENOTFOUND_SESSION = "ResourceNotFound.Session"

	// ResourceNotFound.SessionNotFound
	RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"

	// ResourceNotFound.Version
	RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"

	// ResourceUnavailable.A2ANotConfigured
	RESOURCEUNAVAILABLE_A2ANOTCONFIGURED = "ResourceUnavailable.A2ANotConfigured"

	// ResourceUnavailable.ChatTokenUnavailable
	RESOURCEUNAVAILABLE_CHATTOKENUNAVAILABLE = "ResourceUnavailable.ChatTokenUnavailable"

	// ResourceUnavailable.IdentityNotConfigured
	RESOURCEUNAVAILABLE_IDENTITYNOTCONFIGURED = "ResourceUnavailable.IdentityNotConfigured"

	// ResourceUnavailable.RuntimeAPIUnavailable
	RESOURCEUNAVAILABLE_RUNTIMEAPIUNAVAILABLE = "ResourceUnavailable.RuntimeAPIUnavailable"

	// UnauthorizedOperation.AccountNotAuthorized
	UNAUTHORIZEDOPERATION_ACCOUNTNOTAUTHORIZED = "UnauthorizedOperation.AccountNotAuthorized"

	// UnauthorizedOperation.NotConnectorOwner
	UNAUTHORIZEDOPERATION_NOTCONNECTOROWNER = "UnauthorizedOperation.NotConnectorOwner"

	// UnauthorizedOperation.UnprovisionedTenant
	UNAUTHORIZEDOPERATION_UNPROVISIONEDTENANT = "UnauthorizedOperation.UnprovisionedTenant"
)
