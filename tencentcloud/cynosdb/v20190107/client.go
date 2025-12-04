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

package v20190107

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-07"

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


func NewActivateInstanceRequest() (request *ActivateInstanceRequest) {
    request = &ActivateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ActivateInstance")
    
    
    return
}

func NewActivateInstanceResponse() (response *ActivateInstanceResponse) {
    response = &ActivateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ActivateInstance
// 本接口（ActivateInstance）用于恢复已隔离的实例访问。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PREPAYPAYMODEERROR = "InvalidParameterValue.PrePayPayModeError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ActivateInstance(request *ActivateInstanceRequest) (response *ActivateInstanceResponse, err error) {
    return c.ActivateInstanceWithContext(context.Background(), request)
}

// ActivateInstance
// 本接口（ActivateInstance）用于恢复已隔离的实例访问。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PREPAYPAYMODEERROR = "InvalidParameterValue.PrePayPayModeError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ActivateInstanceWithContext(ctx context.Context, request *ActivateInstanceRequest) (response *ActivateInstanceResponse, err error) {
    if request == nil {
        request = NewActivateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ActivateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAddClusterSlaveZoneRequest() (request *AddClusterSlaveZoneRequest) {
    request = &AddClusterSlaveZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "AddClusterSlaveZone")
    
    
    return
}

func NewAddClusterSlaveZoneResponse() (response *AddClusterSlaveZoneResponse) {
    response = &AddClusterSlaveZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddClusterSlaveZone
// 本接口（AddClusterSlaveZone）用于对集群开启多可用区部署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddClusterSlaveZone(request *AddClusterSlaveZoneRequest) (response *AddClusterSlaveZoneResponse, err error) {
    return c.AddClusterSlaveZoneWithContext(context.Background(), request)
}

// AddClusterSlaveZone
// 本接口（AddClusterSlaveZone）用于对集群开启多可用区部署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddClusterSlaveZoneWithContext(ctx context.Context, request *AddClusterSlaveZoneRequest) (response *AddClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewAddClusterSlaveZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "AddClusterSlaveZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddClusterSlaveZoneResponse()
    err = c.Send(request, response)
    return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
    request = &AddInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "AddInstances")
    
    
    return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
    response = &AddInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddInstances
// 本接口（AddInstances）用于集群添加实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    return c.AddInstancesWithContext(context.Background(), request)
}

// AddInstances
// 本接口（AddInstances）用于集群添加实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddInstancesWithContext(ctx context.Context, request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
    if request == nil {
        request = NewAddInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "AddInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssociateSecurityGroups
// 本接口（AssociateSecurityGroups）用于安全组批量绑定云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// 本接口（AssociateSecurityGroups）用于安全组批量绑定云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "AssociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewBindClusterResourcePackagesRequest() (request *BindClusterResourcePackagesRequest) {
    request = &BindClusterResourcePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "BindClusterResourcePackages")
    
    
    return
}

func NewBindClusterResourcePackagesResponse() (response *BindClusterResourcePackagesResponse) {
    response = &BindClusterResourcePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindClusterResourcePackages
// 本接口（BindClusterResourcePackages）用于为集群绑定资源包。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) BindClusterResourcePackages(request *BindClusterResourcePackagesRequest) (response *BindClusterResourcePackagesResponse, err error) {
    return c.BindClusterResourcePackagesWithContext(context.Background(), request)
}

// BindClusterResourcePackages
// 本接口（BindClusterResourcePackages）用于为集群绑定资源包。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) BindClusterResourcePackagesWithContext(ctx context.Context, request *BindClusterResourcePackagesRequest) (response *BindClusterResourcePackagesResponse, err error) {
    if request == nil {
        request = NewBindClusterResourcePackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "BindClusterResourcePackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindClusterResourcePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindClusterResourcePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewCloseAuditServiceRequest() (request *CloseAuditServiceRequest) {
    request = &CloseAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseAuditService")
    
    
    return
}

func NewCloseAuditServiceResponse() (response *CloseAuditServiceResponse) {
    response = &CloseAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseAuditService
// 本接口（CloseAuditService）用于关闭 TDSQL-C MySQL 实例的数据库审计服务。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) CloseAuditService(request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    return c.CloseAuditServiceWithContext(context.Background(), request)
}

// CloseAuditService
// 本接口（CloseAuditService）用于关闭 TDSQL-C MySQL 实例的数据库审计服务。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) CloseAuditServiceWithContext(ctx context.Context, request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    if request == nil {
        request = NewCloseAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCloseClusterPasswordComplexityRequest() (request *CloseClusterPasswordComplexityRequest) {
    request = &CloseClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseClusterPasswordComplexity")
    
    
    return
}

func NewCloseClusterPasswordComplexityResponse() (response *CloseClusterPasswordComplexityResponse) {
    response = &CloseClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseClusterPasswordComplexity
// 本接口（CloseClusterPasswordComplexity）用于关闭集群密码复杂度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseClusterPasswordComplexity(request *CloseClusterPasswordComplexityRequest) (response *CloseClusterPasswordComplexityResponse, err error) {
    return c.CloseClusterPasswordComplexityWithContext(context.Background(), request)
}

// CloseClusterPasswordComplexity
// 本接口（CloseClusterPasswordComplexity）用于关闭集群密码复杂度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseClusterPasswordComplexityWithContext(ctx context.Context, request *CloseClusterPasswordComplexityRequest) (response *CloseClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewCloseClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxyRequest() (request *CloseProxyRequest) {
    request = &CloseProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseProxy")
    
    
    return
}

func NewCloseProxyResponse() (response *CloseProxyResponse) {
    response = &CloseProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseProxy
// 本接口（CloseProxy）用于关闭集群的数据库代理服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxy(request *CloseProxyRequest) (response *CloseProxyResponse, err error) {
    return c.CloseProxyWithContext(context.Background(), request)
}

// CloseProxy
// 本接口（CloseProxy）用于关闭集群的数据库代理服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxyWithContext(ctx context.Context, request *CloseProxyRequest) (response *CloseProxyResponse, err error) {
    if request == nil {
        request = NewCloseProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCloseProxyEndPointRequest() (request *CloseProxyEndPointRequest) {
    request = &CloseProxyEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseProxyEndPoint")
    
    
    return
}

func NewCloseProxyEndPointResponse() (response *CloseProxyEndPointResponse) {
    response = &CloseProxyEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseProxyEndPoint
// 本接口（CloseProxyEndPoint）用于关闭数据库代理连接地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxyEndPoint(request *CloseProxyEndPointRequest) (response *CloseProxyEndPointResponse, err error) {
    return c.CloseProxyEndPointWithContext(context.Background(), request)
}

// CloseProxyEndPoint
// 本接口（CloseProxyEndPoint）用于关闭数据库代理连接地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseProxyEndPointWithContext(ctx context.Context, request *CloseProxyEndPointRequest) (response *CloseProxyEndPointResponse, err error) {
    if request == nil {
        request = NewCloseProxyEndPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseProxyEndPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseProxyEndPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseProxyEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSSLRequest() (request *CloseSSLRequest) {
    request = &CloseSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseSSL")
    
    
    return
}

func NewCloseSSLResponse() (response *CloseSSLResponse) {
    response = &CloseSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseSSL
// 关闭SSL加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseSSL(request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    return c.CloseSSLWithContext(context.Background(), request)
}

// CloseSSL
// 关闭SSL加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseSSLWithContext(ctx context.Context, request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    if request == nil {
        request = NewCloseSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseSSLResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWanRequest() (request *CloseWanRequest) {
    request = &CloseWanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CloseWan")
    
    
    return
}

func NewCloseWanResponse() (response *CloseWanResponse) {
    response = &CloseWanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseWan
// 本接口（CloseWan）用于关闭外网。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseWan(request *CloseWanRequest) (response *CloseWanResponse, err error) {
    return c.CloseWanWithContext(context.Background(), request)
}

// CloseWan
// 本接口（CloseWan）用于关闭外网。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseWanWithContext(ctx context.Context, request *CloseWanRequest) (response *CloseWanResponse, err error) {
    if request == nil {
        request = NewCloseWanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CloseWan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseWan require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseWanResponse()
    err = c.Send(request, response)
    return
}

func NewCopyClusterPasswordComplexityRequest() (request *CopyClusterPasswordComplexityRequest) {
    request = &CopyClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CopyClusterPasswordComplexity")
    
    
    return
}

func NewCopyClusterPasswordComplexityResponse() (response *CopyClusterPasswordComplexityResponse) {
    response = &CopyClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CopyClusterPasswordComplexity
// 本接口（CopyClusterPasswordComplexity）用于复制集群密码复杂度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyClusterPasswordComplexity(request *CopyClusterPasswordComplexityRequest) (response *CopyClusterPasswordComplexityResponse, err error) {
    return c.CopyClusterPasswordComplexityWithContext(context.Background(), request)
}

// CopyClusterPasswordComplexity
// 本接口（CopyClusterPasswordComplexity）用于复制集群密码复杂度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyClusterPasswordComplexityWithContext(ctx context.Context, request *CopyClusterPasswordComplexityRequest) (response *CopyClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewCopyClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CopyClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountsRequest() (request *CreateAccountsRequest) {
    request = &CreateAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateAccounts")
    
    
    return
}

func NewCreateAccountsResponse() (response *CreateAccountsResponse) {
    response = &CreateAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccounts
// 本接口（CreateAccounts）用于创建用户账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTERROR = "InvalidParameterValue.AccountAlreadyExistError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccounts(request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    return c.CreateAccountsWithContext(context.Background(), request)
}

// CreateAccounts
// 本接口（CreateAccounts）用于创建用户账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTERROR = "InvalidParameterValue.AccountAlreadyExistError"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccountsWithContext(ctx context.Context, request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    if request == nil {
        request = NewCreateAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditLogFileRequest() (request *CreateAuditLogFileRequest) {
    request = &CreateAuditLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateAuditLogFile")
    
    
    return
}

func NewCreateAuditLogFileResponse() (response *CreateAuditLogFileResponse) {
    response = &CreateAuditLogFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuditLogFile
// 本接口(CreateAuditLogFile)用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITFILEOVERQUOTAERROR = "OperationDenied.AuditFileOverQuotaError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
func (c *Client) CreateAuditLogFile(request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    return c.CreateAuditLogFileWithContext(context.Background(), request)
}

// CreateAuditLogFile
// 本接口(CreateAuditLogFile)用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITFILEOVERQUOTAERROR = "OperationDenied.AuditFileOverQuotaError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
func (c *Client) CreateAuditLogFileWithContext(ctx context.Context, request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    if request == nil {
        request = NewCreateAuditLogFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateAuditLogFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditRuleTemplateRequest() (request *CreateAuditRuleTemplateRequest) {
    request = &CreateAuditRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateAuditRuleTemplate")
    
    
    return
}

func NewCreateAuditRuleTemplateResponse() (response *CreateAuditRuleTemplateResponse) {
    response = &CreateAuditRuleTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuditRuleTemplate
// 本接口（CreateAuditRuleTemplate）用于创建审计规则模板。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) CreateAuditRuleTemplate(request *CreateAuditRuleTemplateRequest) (response *CreateAuditRuleTemplateResponse, err error) {
    return c.CreateAuditRuleTemplateWithContext(context.Background(), request)
}

// CreateAuditRuleTemplate
// 本接口（CreateAuditRuleTemplate）用于创建审计规则模板。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) CreateAuditRuleTemplateWithContext(ctx context.Context, request *CreateAuditRuleTemplateRequest) (response *CreateAuditRuleTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAuditRuleTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateAuditRuleTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditRuleTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditRuleTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateBackup")
    
    
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackup
// 本接口（CreateBackup）用于为集群创建手动备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// 本接口（CreateBackup）用于为集群创建手动备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCLSDeliveryRequest() (request *CreateCLSDeliveryRequest) {
    request = &CreateCLSDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateCLSDelivery")
    
    
    return
}

func NewCreateCLSDeliveryResponse() (response *CreateCLSDeliveryResponse) {
    response = &CreateCLSDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCLSDelivery
// 本接口（CreateCLSDelivery）用于创建日志投递。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateCLSDelivery(request *CreateCLSDeliveryRequest) (response *CreateCLSDeliveryResponse, err error) {
    return c.CreateCLSDeliveryWithContext(context.Background(), request)
}

// CreateCLSDelivery
// 本接口（CreateCLSDelivery）用于创建日志投递。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateCLSDeliveryWithContext(ctx context.Context, request *CreateCLSDeliveryRequest) (response *CreateCLSDeliveryResponse, err error) {
    if request == nil {
        request = NewCreateCLSDeliveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateCLSDelivery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCLSDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCLSDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterDatabaseRequest() (request *CreateClusterDatabaseRequest) {
    request = &CreateClusterDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateClusterDatabase")
    
    
    return
}

func NewCreateClusterDatabaseResponse() (response *CreateClusterDatabaseResponse) {
    response = &CreateClusterDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusterDatabase
// 本接口（CreateClusterDatabase）用于创建数据库。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusterDatabase(request *CreateClusterDatabaseRequest) (response *CreateClusterDatabaseResponse, err error) {
    return c.CreateClusterDatabaseWithContext(context.Background(), request)
}

// CreateClusterDatabase
// 本接口（CreateClusterDatabase）用于创建数据库。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusterDatabaseWithContext(ctx context.Context, request *CreateClusterDatabaseRequest) (response *CreateClusterDatabaseResponse, err error) {
    if request == nil {
        request = NewCreateClusterDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateClusterDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusterDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClusterDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClustersRequest() (request *CreateClustersRequest) {
    request = &CreateClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateClusters")
    
    
    return
}

func NewCreateClustersResponse() (response *CreateClustersResponse) {
    response = &CreateClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClusters
// 本接口（CreateClusters）用于新购集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusters(request *CreateClustersRequest) (response *CreateClustersResponse, err error) {
    return c.CreateClustersWithContext(context.Background(), request)
}

// CreateClusters
// 本接口（CreateClusters）用于新购集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClustersWithContext(ctx context.Context, request *CreateClustersRequest) (response *CreateClustersResponse, err error) {
    if request == nil {
        request = NewCreateClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClustersResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrateClusterRequest() (request *CreateIntegrateClusterRequest) {
    request = &CreateIntegrateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateIntegrateCluster")
    
    
    return
}

func NewCreateIntegrateClusterResponse() (response *CreateIntegrateClusterResponse) {
    response = &CreateIntegrateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrateCluster
// 本接口（CreateClusters）用于新购集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIntegrateCluster(request *CreateIntegrateClusterRequest) (response *CreateIntegrateClusterResponse, err error) {
    return c.CreateIntegrateClusterWithContext(context.Background(), request)
}

// CreateIntegrateCluster
// 本接口（CreateClusters）用于新购集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIntegrateClusterWithContext(ctx context.Context, request *CreateIntegrateClusterRequest) (response *CreateIntegrateClusterResponse, err error) {
    if request == nil {
        request = NewCreateIntegrateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateIntegrateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateParamTemplateRequest() (request *CreateParamTemplateRequest) {
    request = &CreateParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateParamTemplate")
    
    
    return
}

func NewCreateParamTemplateResponse() (response *CreateParamTemplateResponse) {
    response = &CreateParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateParamTemplate
// 本接口（CreateParamTemplate）用于创建参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    return c.CreateParamTemplateWithContext(context.Background(), request)
}

// CreateParamTemplate
// 本接口（CreateParamTemplate）用于创建参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateParamTemplateWithContext(ctx context.Context, request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    if request == nil {
        request = NewCreateParamTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyRequest() (request *CreateProxyRequest) {
    request = &CreateProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateProxy")
    
    
    return
}

func NewCreateProxyResponse() (response *CreateProxyResponse) {
    response = &CreateProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProxy
// 本接口（CreateProxy）用于开启集群的数据库代理。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxy(request *CreateProxyRequest) (response *CreateProxyResponse, err error) {
    return c.CreateProxyWithContext(context.Background(), request)
}

// CreateProxy
// 本接口（CreateProxy）用于开启集群的数据库代理。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxyWithContext(ctx context.Context, request *CreateProxyRequest) (response *CreateProxyResponse, err error) {
    if request == nil {
        request = NewCreateProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxyEndPointRequest() (request *CreateProxyEndPointRequest) {
    request = &CreateProxyEndPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateProxyEndPoint")
    
    
    return
}

func NewCreateProxyEndPointResponse() (response *CreateProxyEndPointResponse) {
    response = &CreateProxyEndPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProxyEndPoint
// 本接口（CreateProxyEndPoint）用于创建数据库代理连接点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxyEndPoint(request *CreateProxyEndPointRequest) (response *CreateProxyEndPointResponse, err error) {
    return c.CreateProxyEndPointWithContext(context.Background(), request)
}

// CreateProxyEndPoint
// 本接口（CreateProxyEndPoint）用于创建数据库代理连接点。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateProxyEndPointWithContext(ctx context.Context, request *CreateProxyEndPointRequest) (response *CreateProxyEndPointResponse, err error) {
    if request == nil {
        request = NewCreateProxyEndPointRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateProxyEndPoint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProxyEndPoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProxyEndPointResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourcePackageRequest() (request *CreateResourcePackageRequest) {
    request = &CreateResourcePackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "CreateResourcePackage")
    
    
    return
}

func NewCreateResourcePackageResponse() (response *CreateResourcePackageResponse) {
    response = &CreateResourcePackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateResourcePackage
// 本接口（CreateResourcePackage）用于新购资源包。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATESOURCEPACKAGEERROR = "FailedOperation.CreateSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateResourcePackage(request *CreateResourcePackageRequest) (response *CreateResourcePackageResponse, err error) {
    return c.CreateResourcePackageWithContext(context.Background(), request)
}

// CreateResourcePackage
// 本接口（CreateResourcePackage）用于新购资源包。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATESOURCEPACKAGEERROR = "FailedOperation.CreateSourcePackageError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateResourcePackageWithContext(ctx context.Context, request *CreateResourcePackageRequest) (response *CreateResourcePackageResponse, err error) {
    if request == nil {
        request = NewCreateResourcePackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "CreateResourcePackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourcePackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourcePackageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountsRequest() (request *DeleteAccountsRequest) {
    request = &DeleteAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteAccounts")
    
    
    return
}

func NewDeleteAccountsResponse() (response *DeleteAccountsResponse) {
    response = &DeleteAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccounts
// 本接口（DeleteAccounts）用于删除用户账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccounts(request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    return c.DeleteAccountsWithContext(context.Background(), request)
}

// DeleteAccounts
// 本接口（DeleteAccounts）用于删除用户账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccountsWithContext(ctx context.Context, request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditLogFileRequest() (request *DeleteAuditLogFileRequest) {
    request = &DeleteAuditLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteAuditLogFile")
    
    
    return
}

func NewDeleteAuditLogFileResponse() (response *DeleteAuditLogFileResponse) {
    response = &DeleteAuditLogFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuditLogFile
// 本接口（DeleteAuditLogFile）用于删除云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAuditLogFile(request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    return c.DeleteAuditLogFileWithContext(context.Background(), request)
}

// DeleteAuditLogFile
// 本接口（DeleteAuditLogFile）用于删除云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAuditLogFileWithContext(ctx context.Context, request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    if request == nil {
        request = NewDeleteAuditLogFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteAuditLogFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRuleTemplatesRequest() (request *DeleteAuditRuleTemplatesRequest) {
    request = &DeleteAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteAuditRuleTemplates")
    
    
    return
}

func NewDeleteAuditRuleTemplatesResponse() (response *DeleteAuditRuleTemplatesResponse) {
    response = &DeleteAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuditRuleTemplates
// 本接口（DeleteAuditRuleTemplates）用于删除审计规则模板。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DeleteAuditRuleTemplates(request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    return c.DeleteAuditRuleTemplatesWithContext(context.Background(), request)
}

// DeleteAuditRuleTemplates
// 本接口（DeleteAuditRuleTemplates）用于删除审计规则模板。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DeleteAuditRuleTemplatesWithContext(ctx context.Context, request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
    request = &DeleteBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteBackup")
    
    
    return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
    response = &DeleteBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBackup
// 本接口（DeleteBackup）用于为集群删除手动备份，无法删除自动备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    return c.DeleteBackupWithContext(context.Background(), request)
}

// DeleteBackup
// 本接口（DeleteBackup）用于为集群删除手动备份，无法删除自动备份。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackupWithContext(ctx context.Context, request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    if request == nil {
        request = NewDeleteBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCLSDeliveryRequest() (request *DeleteCLSDeliveryRequest) {
    request = &DeleteCLSDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteCLSDelivery")
    
    
    return
}

func NewDeleteCLSDeliveryResponse() (response *DeleteCLSDeliveryResponse) {
    response = &DeleteCLSDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCLSDelivery
// 本接口（DeleteCLSDelivery）用于删除日志投递。
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteCLSDelivery(request *DeleteCLSDeliveryRequest) (response *DeleteCLSDeliveryResponse, err error) {
    return c.DeleteCLSDeliveryWithContext(context.Background(), request)
}

// DeleteCLSDelivery
// 本接口（DeleteCLSDelivery）用于删除日志投递。
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteCLSDeliveryWithContext(ctx context.Context, request *DeleteCLSDeliveryRequest) (response *DeleteCLSDeliveryResponse, err error) {
    if request == nil {
        request = NewDeleteCLSDeliveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteCLSDelivery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCLSDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCLSDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterDatabaseRequest() (request *DeleteClusterDatabaseRequest) {
    request = &DeleteClusterDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteClusterDatabase")
    
    
    return
}

func NewDeleteClusterDatabaseResponse() (response *DeleteClusterDatabaseResponse) {
    response = &DeleteClusterDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClusterDatabase
// 本接口（DeleteClusterDatabase）用于删除数据库。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterDatabase(request *DeleteClusterDatabaseRequest) (response *DeleteClusterDatabaseResponse, err error) {
    return c.DeleteClusterDatabaseWithContext(context.Background(), request)
}

// DeleteClusterDatabase
// 本接口（DeleteClusterDatabase）用于删除数据库。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterDatabaseWithContext(ctx context.Context, request *DeleteClusterDatabaseRequest) (response *DeleteClusterDatabaseResponse, err error) {
    if request == nil {
        request = NewDeleteClusterDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteClusterDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClusterDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClusterDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParamTemplateRequest() (request *DeleteParamTemplateRequest) {
    request = &DeleteParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DeleteParamTemplate")
    
    
    return
}

func NewDeleteParamTemplateResponse() (response *DeleteParamTemplateResponse) {
    response = &DeleteParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteParamTemplate
// 本接口（DeleteParamTemplate）用于删除用户创建的参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    return c.DeleteParamTemplateWithContext(context.Background(), request)
}

// DeleteParamTemplate
// 本接口（DeleteParamTemplate）用于删除用户创建的参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) DeleteParamTemplateWithContext(ctx context.Context, request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteParamTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DeleteParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountAllGrantPrivilegesRequest() (request *DescribeAccountAllGrantPrivilegesRequest) {
    request = &DescribeAccountAllGrantPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAccountAllGrantPrivileges")
    
    
    return
}

func NewDescribeAccountAllGrantPrivilegesResponse() (response *DescribeAccountAllGrantPrivilegesResponse) {
    response = &DescribeAccountAllGrantPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountAllGrantPrivileges
// 本接口（DescribeAccountAllGrantPrivileges）用于查询账号所有可授予的权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountAllGrantPrivileges(request *DescribeAccountAllGrantPrivilegesRequest) (response *DescribeAccountAllGrantPrivilegesResponse, err error) {
    return c.DescribeAccountAllGrantPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountAllGrantPrivileges
// 本接口（DescribeAccountAllGrantPrivileges）用于查询账号所有可授予的权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountAllGrantPrivilegesWithContext(ctx context.Context, request *DescribeAccountAllGrantPrivilegesRequest) (response *DescribeAccountAllGrantPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountAllGrantPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAccountAllGrantPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountAllGrantPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountAllGrantPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAccountPrivileges")
    
    
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountPrivileges
// 本接口（DescribeAccountPrivileges）用于查询账号已有权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// 本接口（DescribeAccountPrivileges）用于查询账号已有权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询数据库账号列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询数据库账号列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAccounts")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditInstanceListRequest() (request *DescribeAuditInstanceListRequest) {
    request = &DescribeAuditInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditInstanceList")
    
    
    return
}

func NewDescribeAuditInstanceListResponse() (response *DescribeAuditInstanceListResponse) {
    response = &DescribeAuditInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditInstanceList
// 本接口（DescribeAuditInstanceList）用于获取数据库审计的实例列表。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
func (c *Client) DescribeAuditInstanceList(request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    return c.DescribeAuditInstanceListWithContext(context.Background(), request)
}

// DescribeAuditInstanceList
// 本接口（DescribeAuditInstanceList）用于获取数据库审计的实例列表。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
func (c *Client) DescribeAuditInstanceListWithContext(ctx context.Context, request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAuditInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogFilesRequest() (request *DescribeAuditLogFilesRequest) {
    request = &DescribeAuditLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditLogFiles")
    
    
    return
}

func NewDescribeAuditLogFilesResponse() (response *DescribeAuditLogFilesResponse) {
    response = &DescribeAuditLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogFiles
// 本接口(DescribeAuditLogFiles)用于查询云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
func (c *Client) DescribeAuditLogFiles(request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    return c.DescribeAuditLogFilesWithContext(context.Background(), request)
}

// DescribeAuditLogFiles
// 本接口(DescribeAuditLogFiles)用于查询云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
func (c *Client) DescribeAuditLogFilesWithContext(ctx context.Context, request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditLogFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogsRequest() (request *DescribeAuditLogsRequest) {
    request = &DescribeAuditLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditLogs")
    
    
    return
}

func NewDescribeAuditLogsResponse() (response *DescribeAuditLogsResponse) {
    response = &DescribeAuditLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogs
// 本接口（DescribeAuditLogs）用于查询数据库审计日志。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    return c.DescribeAuditLogsWithContext(context.Background(), request)
}

// DescribeAuditLogs
// 本接口（DescribeAuditLogs）用于查询数据库审计日志。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditLogsWithContext(ctx context.Context, request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleTemplatesRequest() (request *DescribeAuditRuleTemplatesRequest) {
    request = &DescribeAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditRuleTemplates")
    
    
    return
}

func NewDescribeAuditRuleTemplatesResponse() (response *DescribeAuditRuleTemplatesResponse) {
    response = &DescribeAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditRuleTemplates
// 本接口（DescribeAuditRuleTemplates）用于查询审计规则模板信息。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditRuleTemplates(request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    return c.DescribeAuditRuleTemplatesWithContext(context.Background(), request)
}

// DescribeAuditRuleTemplates
// 本接口（DescribeAuditRuleTemplates）用于查询审计规则模板信息。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditRuleTemplatesWithContext(ctx context.Context, request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleWithInstanceIdsRequest() (request *DescribeAuditRuleWithInstanceIdsRequest) {
    request = &DescribeAuditRuleWithInstanceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeAuditRuleWithInstanceIds")
    
    
    return
}

func NewDescribeAuditRuleWithInstanceIdsResponse() (response *DescribeAuditRuleWithInstanceIdsResponse) {
    response = &DescribeAuditRuleWithInstanceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditRuleWithInstanceIds
// 本接口（DescribeAuditRuleWithInstanceIds）用于获取实例的审计规则。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DescribeAuditRuleWithInstanceIds(request *DescribeAuditRuleWithInstanceIdsRequest) (response *DescribeAuditRuleWithInstanceIdsResponse, err error) {
    return c.DescribeAuditRuleWithInstanceIdsWithContext(context.Background(), request)
}

// DescribeAuditRuleWithInstanceIds
// 本接口（DescribeAuditRuleWithInstanceIds）用于获取实例的审计规则。
//
// 可能返回的错误码:
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DescribeAuditRuleWithInstanceIdsWithContext(ctx context.Context, request *DescribeAuditRuleWithInstanceIdsRequest) (response *DescribeAuditRuleWithInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleWithInstanceIdsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeAuditRuleWithInstanceIds")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRuleWithInstanceIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRuleWithInstanceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupConfigRequest() (request *DescribeBackupConfigRequest) {
    request = &DescribeBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupConfig")
    
    
    return
}

func NewDescribeBackupConfigResponse() (response *DescribeBackupConfigResponse) {
    response = &DescribeBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupConfig
// 本接口（DescribeBackupConfig）用于获取指定集群的备份配置信息，包括全量备份时间段、备份文件保留时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    return c.DescribeBackupConfigWithContext(context.Background(), request)
}

// DescribeBackupConfig
// 本接口（DescribeBackupConfig）用于获取指定集群的备份配置信息，包括全量备份时间段、备份文件保留时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupConfigWithContext(ctx context.Context, request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadRestriction
// 该接口用户查询当前地域用户设置的默认备份下载来源限制
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// 该接口用户查询当前地域用户设置的默认备份下载来源限制
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupDownloadRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadUrlRequest() (request *DescribeBackupDownloadUrlRequest) {
    request = &DescribeBackupDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupDownloadUrl")
    
    
    return
}

func NewDescribeBackupDownloadUrlResponse() (response *DescribeBackupDownloadUrlResponse) {
    response = &DescribeBackupDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadUrl
// 本接口（DescribeBackupDownloadUrl）用于查询集群备份文件下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUrl(request *DescribeBackupDownloadUrlRequest) (response *DescribeBackupDownloadUrlResponse, err error) {
    return c.DescribeBackupDownloadUrlWithContext(context.Background(), request)
}

// DescribeBackupDownloadUrl
// 本接口（DescribeBackupDownloadUrl）用于查询集群备份文件下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUrlWithContext(ctx context.Context, request *DescribeBackupDownloadUrlRequest) (response *DescribeBackupDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadUserRestrictionRequest() (request *DescribeBackupDownloadUserRestrictionRequest) {
    request = &DescribeBackupDownloadUserRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupDownloadUserRestriction")
    
    
    return
}

func NewDescribeBackupDownloadUserRestrictionResponse() (response *DescribeBackupDownloadUserRestrictionResponse) {
    response = &DescribeBackupDownloadUserRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadUserRestriction
// 该接口用户查询当前地域用户级别设置的默认备份下载来源限制
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUserRestriction(request *DescribeBackupDownloadUserRestrictionRequest) (response *DescribeBackupDownloadUserRestrictionResponse, err error) {
    return c.DescribeBackupDownloadUserRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadUserRestriction
// 该接口用户查询当前地域用户级别设置的默认备份下载来源限制
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUserRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadUserRestrictionRequest) (response *DescribeBackupDownloadUserRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadUserRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupDownloadUserRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadUserRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadUserRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupListRequest() (request *DescribeBackupListRequest) {
    request = &DescribeBackupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBackupList")
    
    
    return
}

func NewDescribeBackupListResponse() (response *DescribeBackupListResponse) {
    response = &DescribeBackupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupList
// 本接口（DescribeBackupList）用于查询集群的备份文件列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupList(request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    return c.DescribeBackupListWithContext(context.Background(), request)
}

// DescribeBackupList
// 本接口（DescribeBackupList）用于查询集群的备份文件列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupListWithContext(ctx context.Context, request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    if request == nil {
        request = NewDescribeBackupListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBackupList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogConfigRequest() (request *DescribeBinlogConfigRequest) {
    request = &DescribeBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogConfig")
    
    
    return
}

func NewDescribeBinlogConfigResponse() (response *DescribeBinlogConfigResponse) {
    response = &DescribeBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogConfig
// 该接口（DescribeBinlogConfig）用于查询binlog配置
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeBinlogConfig(request *DescribeBinlogConfigRequest) (response *DescribeBinlogConfigResponse, err error) {
    return c.DescribeBinlogConfigWithContext(context.Background(), request)
}

// DescribeBinlogConfig
// 该接口（DescribeBinlogConfig）用于查询binlog配置
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeBinlogConfigWithContext(ctx context.Context, request *DescribeBinlogConfigRequest) (response *DescribeBinlogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogDownloadUrlRequest() (request *DescribeBinlogDownloadUrlRequest) {
    request = &DescribeBinlogDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogDownloadUrl")
    
    
    return
}

func NewDescribeBinlogDownloadUrlResponse() (response *DescribeBinlogDownloadUrlResponse) {
    response = &DescribeBinlogDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogDownloadUrl
// 本接口（DescribeBinlogDownloadUrl）用于查询 Binlog 的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogDownloadUrl(request *DescribeBinlogDownloadUrlRequest) (response *DescribeBinlogDownloadUrlResponse, err error) {
    return c.DescribeBinlogDownloadUrlWithContext(context.Background(), request)
}

// DescribeBinlogDownloadUrl
// 本接口（DescribeBinlogDownloadUrl）用于查询 Binlog 的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogDownloadUrlWithContext(ctx context.Context, request *DescribeBinlogDownloadUrlRequest) (response *DescribeBinlogDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogDownloadUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogDownloadUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogSaveDaysRequest() (request *DescribeBinlogSaveDaysRequest) {
    request = &DescribeBinlogSaveDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogSaveDays")
    
    
    return
}

func NewDescribeBinlogSaveDaysResponse() (response *DescribeBinlogSaveDaysResponse) {
    response = &DescribeBinlogSaveDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogSaveDays
// 此接口（DescribeBinlogSaveDays）用于查询集群的Binlog保留天数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogSaveDays(request *DescribeBinlogSaveDaysRequest) (response *DescribeBinlogSaveDaysResponse, err error) {
    return c.DescribeBinlogSaveDaysWithContext(context.Background(), request)
}

// DescribeBinlogSaveDays
// 此接口（DescribeBinlogSaveDays）用于查询集群的Binlog保留天数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_GETBACKUPSTRATEGYERROR = "FailedOperation.GetBackupStrategyError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogSaveDaysWithContext(ctx context.Context, request *DescribeBinlogSaveDaysRequest) (response *DescribeBinlogSaveDaysResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogSaveDaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogSaveDays")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogSaveDays require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogSaveDaysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogsRequest() (request *DescribeBinlogsRequest) {
    request = &DescribeBinlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeBinlogs")
    
    
    return
}

func NewDescribeBinlogsResponse() (response *DescribeBinlogsResponse) {
    response = &DescribeBinlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBinlogs
// 本接口（DescribeBinlogs）用来查询集群 Binlog 日志列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    return c.DescribeBinlogsWithContext(context.Background(), request)
}

// DescribeBinlogs
// 本接口（DescribeBinlogs）用来查询集群 Binlog 日志列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogsWithContext(ctx context.Context, request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeBinlogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChangedParamsAfterUpgradeRequest() (request *DescribeChangedParamsAfterUpgradeRequest) {
    request = &DescribeChangedParamsAfterUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeChangedParamsAfterUpgrade")
    
    
    return
}

func NewDescribeChangedParamsAfterUpgradeResponse() (response *DescribeChangedParamsAfterUpgradeResponse) {
    response = &DescribeChangedParamsAfterUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeChangedParamsAfterUpgrade
// 本接口（DescribeChangedParamsAfterUpgrade）用于查询升降配运行参数对比。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeChangedParamsAfterUpgrade(request *DescribeChangedParamsAfterUpgradeRequest) (response *DescribeChangedParamsAfterUpgradeResponse, err error) {
    return c.DescribeChangedParamsAfterUpgradeWithContext(context.Background(), request)
}

// DescribeChangedParamsAfterUpgrade
// 本接口（DescribeChangedParamsAfterUpgrade）用于查询升降配运行参数对比。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeChangedParamsAfterUpgradeWithContext(ctx context.Context, request *DescribeChangedParamsAfterUpgradeRequest) (response *DescribeChangedParamsAfterUpgradeResponse, err error) {
    if request == nil {
        request = NewDescribeChangedParamsAfterUpgradeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeChangedParamsAfterUpgrade")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChangedParamsAfterUpgrade require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChangedParamsAfterUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDatabaseTablesRequest() (request *DescribeClusterDatabaseTablesRequest) {
    request = &DescribeClusterDatabaseTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDatabaseTables")
    
    
    return
}

func NewDescribeClusterDatabaseTablesResponse() (response *DescribeClusterDatabaseTablesResponse) {
    response = &DescribeClusterDatabaseTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDatabaseTables
// 获取table列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDatabaseTables(request *DescribeClusterDatabaseTablesRequest) (response *DescribeClusterDatabaseTablesResponse, err error) {
    return c.DescribeClusterDatabaseTablesWithContext(context.Background(), request)
}

// DescribeClusterDatabaseTables
// 获取table列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDatabaseTablesWithContext(ctx context.Context, request *DescribeClusterDatabaseTablesRequest) (response *DescribeClusterDatabaseTablesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDatabaseTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDatabaseTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDatabaseTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDatabaseTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDatabasesRequest() (request *DescribeClusterDatabasesRequest) {
    request = &DescribeClusterDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDatabases")
    
    
    return
}

func NewDescribeClusterDatabasesResponse() (response *DescribeClusterDatabasesResponse) {
    response = &DescribeClusterDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDatabases
// 本接口（DescribeClusterDatabases）用于获取集群数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDatabases(request *DescribeClusterDatabasesRequest) (response *DescribeClusterDatabasesResponse, err error) {
    return c.DescribeClusterDatabasesWithContext(context.Background(), request)
}

// DescribeClusterDatabases
// 本接口（DescribeClusterDatabases）用于获取集群数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDatabasesWithContext(ctx context.Context, request *DescribeClusterDatabasesRequest) (response *DescribeClusterDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
    request = &DescribeClusterDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDetail")
    
    
    return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
    response = &DescribeClusterDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDetail
// 该接口（DescribeClusterDetail）用于显示集群详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    return c.DescribeClusterDetailWithContext(context.Background(), request)
}

// DescribeClusterDetail
// 该接口（DescribeClusterDetail）用于显示集群详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  OPERATIONDENIED_VERSIONNOTSUPPORTERROR = "OperationDenied.VersionNotSupportError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetailWithContext(ctx context.Context, request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterDetailDatabasesRequest() (request *DescribeClusterDetailDatabasesRequest) {
    request = &DescribeClusterDetailDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterDetailDatabases")
    
    
    return
}

func NewDescribeClusterDetailDatabasesResponse() (response *DescribeClusterDetailDatabasesResponse) {
    response = &DescribeClusterDetailDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterDetailDatabases
// 本接口（DescribeClusterDetailDatabases）用于查询数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetailDatabases(request *DescribeClusterDetailDatabasesRequest) (response *DescribeClusterDetailDatabasesResponse, err error) {
    return c.DescribeClusterDetailDatabasesWithContext(context.Background(), request)
}

// DescribeClusterDetailDatabases
// 本接口（DescribeClusterDetailDatabases）用于查询数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterDetailDatabasesWithContext(ctx context.Context, request *DescribeClusterDetailDatabasesRequest) (response *DescribeClusterDetailDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterDetailDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterDetailDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetailDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstanceGroupsRequest() (request *DescribeClusterInstanceGroupsRequest) {
    request = &DescribeClusterInstanceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterInstanceGroups")
    
    
    return
}

func NewDescribeClusterInstanceGroupsResponse() (response *DescribeClusterInstanceGroupsResponse) {
    response = &DescribeClusterInstanceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInstanceGroups
// 本接口（DescribeClusterInstanceGrps）用于查询实例组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGroups(request *DescribeClusterInstanceGroupsRequest) (response *DescribeClusterInstanceGroupsResponse, err error) {
    return c.DescribeClusterInstanceGroupsWithContext(context.Background(), request)
}

// DescribeClusterInstanceGroups
// 本接口（DescribeClusterInstanceGrps）用于查询实例组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGroupsWithContext(ctx context.Context, request *DescribeClusterInstanceGroupsRequest) (response *DescribeClusterInstanceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterInstanceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstanceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstanceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstanceGrpsRequest() (request *DescribeClusterInstanceGrpsRequest) {
    request = &DescribeClusterInstanceGrpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterInstanceGrps")
    
    
    return
}

func NewDescribeClusterInstanceGrpsResponse() (response *DescribeClusterInstanceGrpsResponse) {
    response = &DescribeClusterInstanceGrpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInstanceGrps
// 本接口（DescribeClusterInstanceGrps）用于查询实例组信息。 该接口已废弃，推荐使用DescribeClusterInstanceGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGrps(request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    return c.DescribeClusterInstanceGrpsWithContext(context.Background(), request)
}

// DescribeClusterInstanceGrps
// 本接口（DescribeClusterInstanceGrps）用于查询实例组信息。 该接口已废弃，推荐使用DescribeClusterInstanceGroups
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGrpsWithContext(ctx context.Context, request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceGrpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterInstanceGrps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInstanceGrps require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInstanceGrpsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterParamLogsRequest() (request *DescribeClusterParamLogsRequest) {
    request = &DescribeClusterParamLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterParamLogs")
    
    
    return
}

func NewDescribeClusterParamLogsResponse() (response *DescribeClusterParamLogsResponse) {
    response = &DescribeClusterParamLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterParamLogs
// 本接口（DescribeClusterParamLogs）用于查询参数修改记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParamLogs(request *DescribeClusterParamLogsRequest) (response *DescribeClusterParamLogsResponse, err error) {
    return c.DescribeClusterParamLogsWithContext(context.Background(), request)
}

// DescribeClusterParamLogs
// 本接口（DescribeClusterParamLogs）用于查询参数修改记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParamLogsWithContext(ctx context.Context, request *DescribeClusterParamLogsRequest) (response *DescribeClusterParamLogsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterParamLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterParamLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterParamLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterParamLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterParamsRequest() (request *DescribeClusterParamsRequest) {
    request = &DescribeClusterParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterParams")
    
    
    return
}

func NewDescribeClusterParamsResponse() (response *DescribeClusterParamsResponse) {
    response = &DescribeClusterParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterParams
// 本接口（DescribeClusterParams）用于查询集群参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParams(request *DescribeClusterParamsRequest) (response *DescribeClusterParamsResponse, err error) {
    return c.DescribeClusterParamsWithContext(context.Background(), request)
}

// DescribeClusterParams
// 本接口（DescribeClusterParams）用于查询集群参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParamsWithContext(ctx context.Context, request *DescribeClusterParamsRequest) (response *DescribeClusterParamsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterPasswordComplexityRequest() (request *DescribeClusterPasswordComplexityRequest) {
    request = &DescribeClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterPasswordComplexity")
    
    
    return
}

func NewDescribeClusterPasswordComplexityResponse() (response *DescribeClusterPasswordComplexityResponse) {
    response = &DescribeClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterPasswordComplexity
// 本接口（DescribeClusterPasswordComplexity）用于查看集群密码复杂度详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETOSSINFOERROR = "FailedOperation.GetOssInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterPasswordComplexity(request *DescribeClusterPasswordComplexityRequest) (response *DescribeClusterPasswordComplexityResponse, err error) {
    return c.DescribeClusterPasswordComplexityWithContext(context.Background(), request)
}

// DescribeClusterPasswordComplexity
// 本接口（DescribeClusterPasswordComplexity）用于查看集群密码复杂度详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GETOSSINFOERROR = "FailedOperation.GetOssInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterPasswordComplexityWithContext(ctx context.Context, request *DescribeClusterPasswordComplexityRequest) (response *DescribeClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewDescribeClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterReadOnlyRequest() (request *DescribeClusterReadOnlyRequest) {
    request = &DescribeClusterReadOnlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterReadOnly")
    
    
    return
}

func NewDescribeClusterReadOnlyResponse() (response *DescribeClusterReadOnlyResponse) {
    response = &DescribeClusterReadOnlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterReadOnly
// 本接口（DescribeClusterReadOnly）用于查询集群只读开关。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterReadOnly(request *DescribeClusterReadOnlyRequest) (response *DescribeClusterReadOnlyResponse, err error) {
    return c.DescribeClusterReadOnlyWithContext(context.Background(), request)
}

// DescribeClusterReadOnly
// 本接口（DescribeClusterReadOnly）用于查询集群只读开关。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterReadOnlyWithContext(ctx context.Context, request *DescribeClusterReadOnlyRequest) (response *DescribeClusterReadOnlyResponse, err error) {
    if request == nil {
        request = NewDescribeClusterReadOnlyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterReadOnly")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterReadOnly require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterReadOnlyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterTransparentEncryptInfoRequest() (request *DescribeClusterTransparentEncryptInfoRequest) {
    request = &DescribeClusterTransparentEncryptInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusterTransparentEncryptInfo")
    
    
    return
}

func NewDescribeClusterTransparentEncryptInfoResponse() (response *DescribeClusterTransparentEncryptInfoResponse) {
    response = &DescribeClusterTransparentEncryptInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterTransparentEncryptInfo
// 查询集群透明加密信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterTransparentEncryptInfo(request *DescribeClusterTransparentEncryptInfoRequest) (response *DescribeClusterTransparentEncryptInfoResponse, err error) {
    return c.DescribeClusterTransparentEncryptInfoWithContext(context.Background(), request)
}

// DescribeClusterTransparentEncryptInfo
// 查询集群透明加密信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterTransparentEncryptInfoWithContext(ctx context.Context, request *DescribeClusterTransparentEncryptInfoRequest) (response *DescribeClusterTransparentEncryptInfoResponse, err error) {
    if request == nil {
        request = NewDescribeClusterTransparentEncryptInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusterTransparentEncryptInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterTransparentEncryptInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterTransparentEncryptInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusters
// 本接口（DescribeClusters）用于查询集群列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 本接口（DescribeClusters）用于查询集群列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSecurityGroups
// 本接口（DescribeDBSecurityGroups）用于查询实例安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 本接口（DescribeDBSecurityGroups）用于查询实例安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeDBSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeFlow")
    
    
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlow
// 本接口（DescribeFlow）用于查询任务流信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    return c.DescribeFlowWithContext(context.Background(), request)
}

// DescribeFlow
// 本接口（DescribeFlow）用于查询任务流信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlowWithContext(ctx context.Context, request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCLSLogDeliveryRequest() (request *DescribeInstanceCLSLogDeliveryRequest) {
    request = &DescribeInstanceCLSLogDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceCLSLogDelivery")
    
    
    return
}

func NewDescribeInstanceCLSLogDeliveryResponse() (response *DescribeInstanceCLSLogDeliveryResponse) {
    response = &DescribeInstanceCLSLogDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceCLSLogDelivery
// 本接口（DescribeInstanceCLSLogDelivery）用于查询实例日志投递信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceCLSLogDelivery(request *DescribeInstanceCLSLogDeliveryRequest) (response *DescribeInstanceCLSLogDeliveryResponse, err error) {
    return c.DescribeInstanceCLSLogDeliveryWithContext(context.Background(), request)
}

// DescribeInstanceCLSLogDelivery
// 本接口（DescribeInstanceCLSLogDelivery）用于查询实例日志投递信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_GETROLEERROR = "InternalError.GetRoleError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceCLSLogDeliveryWithContext(ctx context.Context, request *DescribeInstanceCLSLogDeliveryRequest) (response *DescribeInstanceCLSLogDeliveryResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCLSLogDeliveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceCLSLogDelivery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceCLSLogDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceCLSLogDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDetailRequest() (request *DescribeInstanceDetailRequest) {
    request = &DescribeInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceDetail")
    
    
    return
}

func NewDescribeInstanceDetailResponse() (response *DescribeInstanceDetailResponse) {
    response = &DescribeInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceDetail
// 本接口(DescribeInstanceDetail)用于查询实例详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceDetail(request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    return c.DescribeInstanceDetailWithContext(context.Background(), request)
}

// DescribeInstanceDetail
// 本接口(DescribeInstanceDetail)用于查询实例详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceDetailWithContext(ctx context.Context, request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceErrorLogsRequest() (request *DescribeInstanceErrorLogsRequest) {
    request = &DescribeInstanceErrorLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceErrorLogs")
    
    
    return
}

func NewDescribeInstanceErrorLogsResponse() (response *DescribeInstanceErrorLogsResponse) {
    response = &DescribeInstanceErrorLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceErrorLogs
// 本接口（DescribeInstanceErrorLogs）用于查询实例错误日志列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_VALUENOTFOUND = "InvalidParameterValue.ValueNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceErrorLogs(request *DescribeInstanceErrorLogsRequest) (response *DescribeInstanceErrorLogsResponse, err error) {
    return c.DescribeInstanceErrorLogsWithContext(context.Background(), request)
}

// DescribeInstanceErrorLogs
// 本接口（DescribeInstanceErrorLogs）用于查询实例错误日志列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_VALUENOTFOUND = "InvalidParameterValue.ValueNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceErrorLogsWithContext(ctx context.Context, request *DescribeInstanceErrorLogsRequest) (response *DescribeInstanceErrorLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceErrorLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceErrorLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceErrorLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceErrorLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询实例参数列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询实例参数列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSlowQueriesRequest() (request *DescribeInstanceSlowQueriesRequest) {
    request = &DescribeInstanceSlowQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceSlowQueries")
    
    
    return
}

func NewDescribeInstanceSlowQueriesResponse() (response *DescribeInstanceSlowQueriesResponse) {
    response = &DescribeInstanceSlowQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSlowQueries
// 此接口（DescribeInstanceSlowQueries）用于查询实例慢日志详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_LENGTHOVERLIMIT = "OperationDenied.LengthOverLimit"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSlowQueries(request *DescribeInstanceSlowQueriesRequest) (response *DescribeInstanceSlowQueriesResponse, err error) {
    return c.DescribeInstanceSlowQueriesWithContext(context.Background(), request)
}

// DescribeInstanceSlowQueries
// 此接口（DescribeInstanceSlowQueries）用于查询实例慢日志详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  OPERATIONDENIED_LENGTHOVERLIMIT = "OperationDenied.LengthOverLimit"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSlowQueriesWithContext(ctx context.Context, request *DescribeInstanceSlowQueriesRequest) (response *DescribeInstanceSlowQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSlowQueriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceSlowQueries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSlowQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSlowQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSpecsRequest() (request *DescribeInstanceSpecsRequest) {
    request = &DescribeInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstanceSpecs")
    
    
    return
}

func NewDescribeInstanceSpecsResponse() (response *DescribeInstanceSpecsResponse) {
    response = &DescribeInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSpecs
// 本接口（DescribeInstanceSpecs）用于查询购买页可购买的实例规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    return c.DescribeInstanceSpecsWithContext(context.Background(), request)
}

// DescribeInstanceSpecs
// 本接口（DescribeInstanceSpecs）用于查询购买页可购买的实例规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSpecsWithContext(ctx context.Context, request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstanceSpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 本接口(DescribeInstances)用于查询实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 本接口(DescribeInstances)用于查询实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesWithinSameClusterRequest() (request *DescribeInstancesWithinSameClusterRequest) {
    request = &DescribeInstancesWithinSameClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeInstancesWithinSameCluster")
    
    
    return
}

func NewDescribeInstancesWithinSameClusterResponse() (response *DescribeInstancesWithinSameClusterResponse) {
    response = &DescribeInstancesWithinSameClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstancesWithinSameCluster
// 本接口（DescribeInstancesWithinSameCluster）用于查询同一集群下实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstancesWithinSameCluster(request *DescribeInstancesWithinSameClusterRequest) (response *DescribeInstancesWithinSameClusterResponse, err error) {
    return c.DescribeInstancesWithinSameClusterWithContext(context.Background(), request)
}

// DescribeInstancesWithinSameCluster
// 本接口（DescribeInstancesWithinSameCluster）用于查询同一集群下实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstancesWithinSameClusterWithContext(ctx context.Context, request *DescribeInstancesWithinSameClusterRequest) (response *DescribeInstancesWithinSameClusterResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesWithinSameClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeInstancesWithinSameCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesWithinSameCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesWithinSameClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrateTaskRequest() (request *DescribeIntegrateTaskRequest) {
    request = &DescribeIntegrateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeIntegrateTask")
    
    
    return
}

func NewDescribeIntegrateTaskResponse() (response *DescribeIntegrateTaskResponse) {
    response = &DescribeIntegrateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrateTask
// 本接口（DescribeIntegrateTask）用于查询集群任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIntegrateTask(request *DescribeIntegrateTaskRequest) (response *DescribeIntegrateTaskResponse, err error) {
    return c.DescribeIntegrateTaskWithContext(context.Background(), request)
}

// DescribeIntegrateTask
// 本接口（DescribeIntegrateTask）用于查询集群任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIntegrateTaskWithContext(ctx context.Context, request *DescribeIntegrateTaskRequest) (response *DescribeIntegrateTaskResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrateTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeIntegrateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIsolatedInstancesRequest() (request *DescribeIsolatedInstancesRequest) {
    request = &DescribeIsolatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeIsolatedInstances")
    
    
    return
}

func NewDescribeIsolatedInstancesResponse() (response *DescribeIsolatedInstancesResponse) {
    response = &DescribeIsolatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIsolatedInstances
// 本接口（DescribeIsolatedInstances）用于查询回收站实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIsolatedInstances(request *DescribeIsolatedInstancesRequest) (response *DescribeIsolatedInstancesResponse, err error) {
    return c.DescribeIsolatedInstancesWithContext(context.Background(), request)
}

// DescribeIsolatedInstances
// 本接口（DescribeIsolatedInstances）用于查询回收站实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_CAMSIGANDAUTHERROR = "FailedOperation.CamSigAndAuthError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIsolatedInstancesWithContext(ctx context.Context, request *DescribeIsolatedInstancesRequest) (response *DescribeIsolatedInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeIsolatedInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeIsolatedInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIsolatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIsolatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintainPeriodRequest() (request *DescribeMaintainPeriodRequest) {
    request = &DescribeMaintainPeriodRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeMaintainPeriod")
    
    
    return
}

func NewDescribeMaintainPeriodResponse() (response *DescribeMaintainPeriodResponse) {
    response = &DescribeMaintainPeriodResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMaintainPeriod
// 本接口（DescribeMaintainPeriod）用于查询实例维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintainPeriod(request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    return c.DescribeMaintainPeriodWithContext(context.Background(), request)
}

// DescribeMaintainPeriod
// 本接口（DescribeMaintainPeriod）用于查询实例维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintainPeriodWithContext(ctx context.Context, request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeMaintainPeriodRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeMaintainPeriod")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintainPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintainPeriodResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplateDetailRequest() (request *DescribeParamTemplateDetailRequest) {
    request = &DescribeParamTemplateDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeParamTemplateDetail")
    
    
    return
}

func NewDescribeParamTemplateDetailResponse() (response *DescribeParamTemplateDetailResponse) {
    response = &DescribeParamTemplateDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamTemplateDetail
// 本接口（DescribeParamTemplateDetail）用于查询用户参数模板详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplateDetail(request *DescribeParamTemplateDetailRequest) (response *DescribeParamTemplateDetailResponse, err error) {
    return c.DescribeParamTemplateDetailWithContext(context.Background(), request)
}

// DescribeParamTemplateDetail
// 本接口（DescribeParamTemplateDetail）用于查询用户参数模板详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplateDetailWithContext(ctx context.Context, request *DescribeParamTemplateDetailRequest) (response *DescribeParamTemplateDetailResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplateDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeParamTemplateDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplateDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParamTemplateDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplatesRequest() (request *DescribeParamTemplatesRequest) {
    request = &DescribeParamTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeParamTemplates")
    
    
    return
}

func NewDescribeParamTemplatesResponse() (response *DescribeParamTemplatesResponse) {
    response = &DescribeParamTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamTemplates
// 本接口（DescribeParamTemplates）用于查询用户指定产品下的所有参数模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// 本接口（DescribeParamTemplates）用于查询用户指定产品下的所有参数模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplatesWithContext(ctx context.Context, request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeParamTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParamTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProjectSecurityGroups")
    
    
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectSecurityGroups
// 本接口（DescribeProjectSecurityGroups）用于查询项目安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// 本接口（DescribeProjectSecurityGroups）用于查询项目安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroupsWithContext(ctx context.Context, request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProjectSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxiesRequest() (request *DescribeProxiesRequest) {
    request = &DescribeProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProxies")
    
    
    return
}

func NewDescribeProxiesResponse() (response *DescribeProxiesResponse) {
    response = &DescribeProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxies
// 本接口（DescribeProxies）用于查询数据库代理列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxies(request *DescribeProxiesRequest) (response *DescribeProxiesResponse, err error) {
    return c.DescribeProxiesWithContext(context.Background(), request)
}

// DescribeProxies
// 本接口（DescribeProxies）用于查询数据库代理列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxiesWithContext(ctx context.Context, request *DescribeProxiesRequest) (response *DescribeProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeProxiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProxies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyNodesRequest() (request *DescribeProxyNodesRequest) {
    request = &DescribeProxyNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProxyNodes")
    
    
    return
}

func NewDescribeProxyNodesResponse() (response *DescribeProxyNodesResponse) {
    response = &DescribeProxyNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxyNodes
// 本接口（DescribeProxyNodes）用于查询代理节点列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxyNodes(request *DescribeProxyNodesRequest) (response *DescribeProxyNodesResponse, err error) {
    return c.DescribeProxyNodesWithContext(context.Background(), request)
}

// DescribeProxyNodes
// 本接口（DescribeProxyNodes）用于查询代理节点列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProxyNodesWithContext(ctx context.Context, request *DescribeProxyNodesRequest) (response *DescribeProxyNodesResponse, err error) {
    if request == nil {
        request = NewDescribeProxyNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProxyNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyNodesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxySpecsRequest() (request *DescribeProxySpecsRequest) {
    request = &DescribeProxySpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeProxySpecs")
    
    
    return
}

func NewDescribeProxySpecsResponse() (response *DescribeProxySpecsResponse) {
    response = &DescribeProxySpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxySpecs
// 本接口（DescribeProxySpecs）用于查询数据库代理规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
func (c *Client) DescribeProxySpecs(request *DescribeProxySpecsRequest) (response *DescribeProxySpecsResponse, err error) {
    return c.DescribeProxySpecsWithContext(context.Background(), request)
}

// DescribeProxySpecs
// 本接口（DescribeProxySpecs）用于查询数据库代理规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
func (c *Client) DescribeProxySpecsWithContext(ctx context.Context, request *DescribeProxySpecsRequest) (response *DescribeProxySpecsResponse, err error) {
    if request == nil {
        request = NewDescribeProxySpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeProxySpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePackageDetailRequest() (request *DescribeResourcePackageDetailRequest) {
    request = &DescribeResourcePackageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcePackageDetail")
    
    
    return
}

func NewDescribeResourcePackageDetailResponse() (response *DescribeResourcePackageDetailResponse) {
    response = &DescribeResourcePackageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcePackageDetail
// 本接口（DescribeResourcePackageDetail）用于查询资源包使用详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEDETAILERROR = "FailedOperation.QuerySourcePackageDetailError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageDetail(request *DescribeResourcePackageDetailRequest) (response *DescribeResourcePackageDetailResponse, err error) {
    return c.DescribeResourcePackageDetailWithContext(context.Background(), request)
}

// DescribeResourcePackageDetail
// 本接口（DescribeResourcePackageDetail）用于查询资源包使用详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEDETAILERROR = "FailedOperation.QuerySourcePackageDetailError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageDetailWithContext(ctx context.Context, request *DescribeResourcePackageDetailRequest) (response *DescribeResourcePackageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePackageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcePackageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcePackageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcePackageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePackageListRequest() (request *DescribeResourcePackageListRequest) {
    request = &DescribeResourcePackageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcePackageList")
    
    
    return
}

func NewDescribeResourcePackageListResponse() (response *DescribeResourcePackageListResponse) {
    response = &DescribeResourcePackageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcePackageList
// 本接口（DescribeResourcePackageList）用于查询资源包列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageList(request *DescribeResourcePackageListRequest) (response *DescribeResourcePackageListResponse, err error) {
    return c.DescribeResourcePackageListWithContext(context.Background(), request)
}

// DescribeResourcePackageList
// 本接口（DescribeResourcePackageList）用于查询资源包列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageListWithContext(ctx context.Context, request *DescribeResourcePackageListRequest) (response *DescribeResourcePackageListResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePackageListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcePackageList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcePackageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcePackageListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcePackageSaleSpecRequest() (request *DescribeResourcePackageSaleSpecRequest) {
    request = &DescribeResourcePackageSaleSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcePackageSaleSpec")
    
    
    return
}

func NewDescribeResourcePackageSaleSpecResponse() (response *DescribeResourcePackageSaleSpecResponse) {
    response = &DescribeResourcePackageSaleSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcePackageSaleSpec
// 本接口（DescribeResourcePackageSaleSpec）用于查询资源包规格。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_UNSUPPORTSALESPECERROR = "OperationDenied.UnSupportSaleSpecError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageSaleSpec(request *DescribeResourcePackageSaleSpecRequest) (response *DescribeResourcePackageSaleSpecResponse, err error) {
    return c.DescribeResourcePackageSaleSpecWithContext(context.Background(), request)
}

// DescribeResourcePackageSaleSpec
// 本接口（DescribeResourcePackageSaleSpec）用于查询资源包规格。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_UNSUPPORTSALESPECERROR = "OperationDenied.UnSupportSaleSpecError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcePackageSaleSpecWithContext(ctx context.Context, request *DescribeResourcePackageSaleSpecRequest) (response *DescribeResourcePackageSaleSpecResponse, err error) {
    if request == nil {
        request = NewDescribeResourcePackageSaleSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcePackageSaleSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcePackageSaleSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcePackageSaleSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByDealNameRequest() (request *DescribeResourcesByDealNameRequest) {
    request = &DescribeResourcesByDealNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeResourcesByDealName")
    
    
    return
}

func NewDescribeResourcesByDealNameResponse() (response *DescribeResourcesByDealNameResponse) {
    response = &DescribeResourcesByDealNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcesByDealName
// 本接口（DescribeResourcesByDealName）用于查询订单关联实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealName(request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    return c.DescribeResourcesByDealNameWithContext(context.Background(), request)
}

// DescribeResourcesByDealName
// 本接口（DescribeResourcesByDealName）用于查询订单关联实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  OPERATIONDENIED_INSTANCEACCESSDENIEDERROR = "OperationDenied.InstanceAccessDeniedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealNameWithContext(ctx context.Context, request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByDealNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeResourcesByDealName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcesByDealName require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesByDealNameResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeRangeRequest() (request *DescribeRollbackTimeRangeRequest) {
    request = &DescribeRollbackTimeRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeRollbackTimeRange")
    
    
    return
}

func NewDescribeRollbackTimeRangeResponse() (response *DescribeRollbackTimeRangeResponse) {
    response = &DescribeRollbackTimeRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRollbackTimeRange
// 本接口（DescribeRollbackTimeRange）用于查询回档时间范围。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeRange(request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    return c.DescribeRollbackTimeRangeWithContext(context.Background(), request)
}

// DescribeRollbackTimeRange
// 本接口（DescribeRollbackTimeRange）用于查询回档时间范围。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeRangeWithContext(ctx context.Context, request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRangeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeRollbackTimeRange")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTimeRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTimeRangeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSSLStatusRequest() (request *DescribeSSLStatusRequest) {
    request = &DescribeSSLStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeSSLStatus")
    
    
    return
}

func NewDescribeSSLStatusResponse() (response *DescribeSSLStatusResponse) {
    response = &DescribeSSLStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSSLStatus
// 查询实例SSL状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSSLStatus(request *DescribeSSLStatusRequest) (response *DescribeSSLStatusResponse, err error) {
    return c.DescribeSSLStatusWithContext(context.Background(), request)
}

// DescribeSSLStatus
// 查询实例SSL状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSSLStatusWithContext(ctx context.Context, request *DescribeSSLStatusRequest) (response *DescribeSSLStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSSLStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeSSLStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSSLStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSSLStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessInstanceSpecsRequest() (request *DescribeServerlessInstanceSpecsRequest) {
    request = &DescribeServerlessInstanceSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeServerlessInstanceSpecs")
    
    
    return
}

func NewDescribeServerlessInstanceSpecsResponse() (response *DescribeServerlessInstanceSpecsResponse) {
    response = &DescribeServerlessInstanceSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessInstanceSpecs
// 查询Serverless实例可选规格
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerlessInstanceSpecs(request *DescribeServerlessInstanceSpecsRequest) (response *DescribeServerlessInstanceSpecsResponse, err error) {
    return c.DescribeServerlessInstanceSpecsWithContext(context.Background(), request)
}

// DescribeServerlessInstanceSpecs
// 查询Serverless实例可选规格
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeServerlessInstanceSpecsWithContext(ctx context.Context, request *DescribeServerlessInstanceSpecsRequest) (response *DescribeServerlessInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessInstanceSpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeServerlessInstanceSpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessInstanceSpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessInstanceSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeServerlessStrategyRequest() (request *DescribeServerlessStrategyRequest) {
    request = &DescribeServerlessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeServerlessStrategy")
    
    
    return
}

func NewDescribeServerlessStrategyResponse() (response *DescribeServerlessStrategyResponse) {
    response = &DescribeServerlessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeServerlessStrategy
// 查询serverless策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) DescribeServerlessStrategy(request *DescribeServerlessStrategyRequest) (response *DescribeServerlessStrategyResponse, err error) {
    return c.DescribeServerlessStrategyWithContext(context.Background(), request)
}

// DescribeServerlessStrategy
// 查询serverless策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) DescribeServerlessStrategyWithContext(ctx context.Context, request *DescribeServerlessStrategyRequest) (response *DescribeServerlessStrategyResponse, err error) {
    if request == nil {
        request = NewDescribeServerlessStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeServerlessStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeServerlessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeServerlessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlaveZonesRequest() (request *DescribeSlaveZonesRequest) {
    request = &DescribeSlaveZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeSlaveZones")
    
    
    return
}

func NewDescribeSlaveZonesResponse() (response *DescribeSlaveZonesResponse) {
    response = &DescribeSlaveZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlaveZones
// 查询从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeSlaveZones(request *DescribeSlaveZonesRequest) (response *DescribeSlaveZonesResponse, err error) {
    return c.DescribeSlaveZonesWithContext(context.Background(), request)
}

// DescribeSlaveZones
// 查询从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) DescribeSlaveZonesWithContext(ctx context.Context, request *DescribeSlaveZonesRequest) (response *DescribeSlaveZonesResponse, err error) {
    if request == nil {
        request = NewDescribeSlaveZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeSlaveZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlaveZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlaveZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportProxyVersionRequest() (request *DescribeSupportProxyVersionRequest) {
    request = &DescribeSupportProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeSupportProxyVersion")
    
    
    return
}

func NewDescribeSupportProxyVersionResponse() (response *DescribeSupportProxyVersionResponse) {
    response = &DescribeSupportProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSupportProxyVersion
// 本接口（DescribeSupportProxyVersion）用于查询支持的数据库代理版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSupportProxyVersion(request *DescribeSupportProxyVersionRequest) (response *DescribeSupportProxyVersionResponse, err error) {
    return c.DescribeSupportProxyVersionWithContext(context.Background(), request)
}

// DescribeSupportProxyVersion
// 本接口（DescribeSupportProxyVersion）用于查询支持的数据库代理版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSupportProxyVersionWithContext(ctx context.Context, request *DescribeSupportProxyVersionRequest) (response *DescribeSupportProxyVersionResponse, err error) {
    if request == nil {
        request = NewDescribeSupportProxyVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeSupportProxyVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportProxyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTasks
// 本接口（DescribeTasks）用于查询任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// 本接口（DescribeTasks）用于查询任务列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeZones
// 本接口（DescribeZones）用于查询可售卖地域可用区信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 本接口（DescribeZones）用于查询可售卖地域可用区信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DescribeZones")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeZonesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisassociateSecurityGroups
// 本接口（DisassociateSecurityGroups）用于安全组批量解绑云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// 本接口（DisassociateSecurityGroups）用于安全组批量解绑云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "DisassociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewExportInstanceErrorLogsRequest() (request *ExportInstanceErrorLogsRequest) {
    request = &ExportInstanceErrorLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ExportInstanceErrorLogs")
    
    
    return
}

func NewExportInstanceErrorLogsResponse() (response *ExportInstanceErrorLogsResponse) {
    response = &ExportInstanceErrorLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportInstanceErrorLogs
// 此接口（ExportInstanceErrorLogs）用于导出实例错误日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceErrorLogs(request *ExportInstanceErrorLogsRequest) (response *ExportInstanceErrorLogsResponse, err error) {
    return c.ExportInstanceErrorLogsWithContext(context.Background(), request)
}

// ExportInstanceErrorLogs
// 此接口（ExportInstanceErrorLogs）用于导出实例错误日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceErrorLogsWithContext(ctx context.Context, request *ExportInstanceErrorLogsRequest) (response *ExportInstanceErrorLogsResponse, err error) {
    if request == nil {
        request = NewExportInstanceErrorLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ExportInstanceErrorLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportInstanceErrorLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportInstanceErrorLogsResponse()
    err = c.Send(request, response)
    return
}

func NewExportInstanceSlowQueriesRequest() (request *ExportInstanceSlowQueriesRequest) {
    request = &ExportInstanceSlowQueriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ExportInstanceSlowQueries")
    
    
    return
}

func NewExportInstanceSlowQueriesResponse() (response *ExportInstanceSlowQueriesResponse) {
    response = &ExportInstanceSlowQueriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportInstanceSlowQueries
// 本接口（ExportInstanceSlowQueries）用于导出实例慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceSlowQueries(request *ExportInstanceSlowQueriesRequest) (response *ExportInstanceSlowQueriesResponse, err error) {
    return c.ExportInstanceSlowQueriesWithContext(context.Background(), request)
}

// ExportInstanceSlowQueries
// 本接口（ExportInstanceSlowQueries）用于导出实例慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceSlowQueriesWithContext(ctx context.Context, request *ExportInstanceSlowQueriesRequest) (response *ExportInstanceSlowQueriesResponse, err error) {
    if request == nil {
        request = NewExportInstanceSlowQueriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ExportInstanceSlowQueries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportInstanceSlowQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportInstanceSlowQueriesResponse()
    err = c.Send(request, response)
    return
}

func NewExportResourcePackageDeductDetailsRequest() (request *ExportResourcePackageDeductDetailsRequest) {
    request = &ExportResourcePackageDeductDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ExportResourcePackageDeductDetails")
    
    
    return
}

func NewExportResourcePackageDeductDetailsResponse() (response *ExportResourcePackageDeductDetailsResponse) {
    response = &ExportResourcePackageDeductDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportResourcePackageDeductDetails
// 资源包使用明细导出
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportResourcePackageDeductDetails(request *ExportResourcePackageDeductDetailsRequest) (response *ExportResourcePackageDeductDetailsResponse, err error) {
    return c.ExportResourcePackageDeductDetailsWithContext(context.Background(), request)
}

// ExportResourcePackageDeductDetails
// 资源包使用明细导出
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportResourcePackageDeductDetailsWithContext(ctx context.Context, request *ExportResourcePackageDeductDetailsRequest) (response *ExportResourcePackageDeductDetailsResponse, err error) {
    if request == nil {
        request = NewExportResourcePackageDeductDetailsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ExportResourcePackageDeductDetails")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportResourcePackageDeductDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportResourcePackageDeductDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
    request = &GrantAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "GrantAccountPrivileges")
    
    
    return
}

func NewGrantAccountPrivilegesResponse() (response *GrantAccountPrivilegesResponse) {
    response = &GrantAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GrantAccountPrivileges
// 本接口（GrantAccountPrivileges）用于批量授权账号权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DESCRIBEACCOUNTPRIVILEGESERROR = "FailedOperation.DescribeAccountPrivilegesError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    return c.GrantAccountPrivilegesWithContext(context.Background(), request)
}

// GrantAccountPrivileges
// 本接口（GrantAccountPrivileges）用于批量授权账号权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_DESCRIBEACCOUNTPRIVILEGESERROR = "FailedOperation.DescribeAccountPrivilegesError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivilegesWithContext(ctx context.Context, request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewGrantAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "GrantAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GrantAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewGrantAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateRequest() (request *InquirePriceCreateRequest) {
    request = &InquirePriceCreateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceCreate")
    
    
    return
}

func NewInquirePriceCreateResponse() (response *InquirePriceCreateResponse) {
    response = &InquirePriceCreateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreate
// 本接口（InquirePriceCreate）用于新购集群的价格查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceCreate(request *InquirePriceCreateRequest) (response *InquirePriceCreateResponse, err error) {
    return c.InquirePriceCreateWithContext(context.Background(), request)
}

// InquirePriceCreate
// 本接口（InquirePriceCreate）用于新购集群的价格查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceCreateWithContext(ctx context.Context, request *InquirePriceCreateRequest) (response *InquirePriceCreateResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceCreate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreate require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyRequest() (request *InquirePriceModifyRequest) {
    request = &InquirePriceModifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceModify")
    
    
    return
}

func NewInquirePriceModifyResponse() (response *InquirePriceModifyResponse) {
    response = &InquirePriceModifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceModify
// 变配预付费集群询价
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceModify(request *InquirePriceModifyRequest) (response *InquirePriceModifyResponse, err error) {
    return c.InquirePriceModifyWithContext(context.Background(), request)
}

// InquirePriceModify
// 变配预付费集群询价
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceModifyWithContext(ctx context.Context, request *InquirePriceModifyRequest) (response *InquirePriceModifyResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceModify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModify require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceMultiSpecRequest() (request *InquirePriceMultiSpecRequest) {
    request = &InquirePriceMultiSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceMultiSpec")
    
    
    return
}

func NewInquirePriceMultiSpecResponse() (response *InquirePriceMultiSpecResponse) {
    response = &InquirePriceMultiSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceMultiSpec
// 此接口（InquirePriceMultiSpec）用于批量询价
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceMultiSpec(request *InquirePriceMultiSpecRequest) (response *InquirePriceMultiSpecResponse, err error) {
    return c.InquirePriceMultiSpecWithContext(context.Background(), request)
}

// InquirePriceMultiSpec
// 此接口（InquirePriceMultiSpec）用于批量询价
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceMultiSpecWithContext(ctx context.Context, request *InquirePriceMultiSpecRequest) (response *InquirePriceMultiSpecResponse, err error) {
    if request == nil {
        request = NewInquirePriceMultiSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceMultiSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceMultiSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceMultiSpecResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewRequest() (request *InquirePriceRenewRequest) {
    request = &InquirePriceRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "InquirePriceRenew")
    
    
    return
}

func NewInquirePriceRenewResponse() (response *InquirePriceRenewResponse) {
    response = &InquirePriceRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRenew
// 本接口（InquirePriceRenew）用于查询续费集群价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BATCHGETINSTANCEERROR = "FailedOperation.BatchGetInstanceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceRenew(request *InquirePriceRenewRequest) (response *InquirePriceRenewResponse, err error) {
    return c.InquirePriceRenewWithContext(context.Background(), request)
}

// InquirePriceRenew
// 本接口（InquirePriceRenew）用于查询续费集群价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BATCHGETINSTANCEERROR = "FailedOperation.BatchGetInstanceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquirePriceRenewWithContext(ctx context.Context, request *InquirePriceRenewRequest) (response *InquirePriceRenewResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "InquirePriceRenew")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenew require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateClusterRequest() (request *IsolateClusterRequest) {
    request = &IsolateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "IsolateCluster")
    
    
    return
}

func NewIsolateClusterResponse() (response *IsolateClusterResponse) {
    response = &IsolateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateCluster
// 本接口（IsolateCluster）用于隔离集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateCluster(request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    return c.IsolateClusterWithContext(context.Background(), request)
}

// IsolateCluster
// 本接口（IsolateCluster）用于隔离集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateClusterWithContext(ctx context.Context, request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    if request == nil {
        request = NewIsolateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "IsolateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateInstanceRequest() (request *IsolateInstanceRequest) {
    request = &IsolateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "IsolateInstance")
    
    
    return
}

func NewIsolateInstanceResponse() (response *IsolateInstanceResponse) {
    response = &IsolateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateInstance
// 本接口(IsolateInstance)用于隔离实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateInstance(request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
    return c.IsolateInstanceWithContext(context.Background(), request)
}

// IsolateInstance
// 本接口(IsolateInstance)用于隔离实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateInstanceWithContext(ctx context.Context, request *IsolateInstanceRequest) (response *IsolateInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "IsolateInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountDescription")
    
    
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountDescription
// 本接口(ModifyAccountDescription)用于修改数据库账号描述信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    return c.ModifyAccountDescriptionWithContext(context.Background(), request)
}

// ModifyAccountDescription
// 本接口(ModifyAccountDescription)用于修改数据库账号描述信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountDescription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountHostRequest() (request *ModifyAccountHostRequest) {
    request = &ModifyAccountHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountHost")
    
    
    return
}

func NewModifyAccountHostResponse() (response *ModifyAccountHostResponse) {
    response = &ModifyAccountHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountHost
// 本接口（ModifyAccountHost）用于修改账号主机。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountHost(request *ModifyAccountHostRequest) (response *ModifyAccountHostResponse, err error) {
    return c.ModifyAccountHostWithContext(context.Background(), request)
}

// ModifyAccountHost
// 本接口（ModifyAccountHost）用于修改账号主机。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountHostWithContext(ctx context.Context, request *ModifyAccountHostRequest) (response *ModifyAccountHostResponse, err error) {
    if request == nil {
        request = NewModifyAccountHostRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountHost")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountParamsRequest() (request *ModifyAccountParamsRequest) {
    request = &ModifyAccountParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountParams")
    
    
    return
}

func NewModifyAccountParamsResponse() (response *ModifyAccountParamsResponse) {
    response = &ModifyAccountParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountParams
// 本接口（ModifyAccountParams）用于修改账号配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountParams(request *ModifyAccountParamsRequest) (response *ModifyAccountParamsResponse, err error) {
    return c.ModifyAccountParamsWithContext(context.Background(), request)
}

// ModifyAccountParams
// 本接口（ModifyAccountParams）用于修改账号配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountParamsWithContext(ctx context.Context, request *ModifyAccountParamsRequest) (response *ModifyAccountParamsResponse, err error) {
    if request == nil {
        request = NewModifyAccountParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountParamsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAccountPrivileges")
    
    
    return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
    response = &ModifyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAccountPrivileges
// 本接口（ModifyAccountPrivileges）用于修改账号库表权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// 本接口（ModifyAccountPrivileges）用于修改账号库表权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditRuleTemplatesRequest() (request *ModifyAuditRuleTemplatesRequest) {
    request = &ModifyAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAuditRuleTemplates")
    
    
    return
}

func NewModifyAuditRuleTemplatesResponse() (response *ModifyAuditRuleTemplatesResponse) {
    response = &ModifyAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditRuleTemplates
// 本接口（ModifyAuditRuleTemplates）用于修改审计规则模板。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyAuditRuleTemplates(request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    return c.ModifyAuditRuleTemplatesWithContext(context.Background(), request)
}

// ModifyAuditRuleTemplates
// 本接口（ModifyAuditRuleTemplates）用于修改审计规则模板。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyAuditRuleTemplatesWithContext(ctx context.Context, request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewModifyAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditRuleTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditServiceRequest() (request *ModifyAuditServiceRequest) {
    request = &ModifyAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyAuditService")
    
    
    return
}

func NewModifyAuditServiceResponse() (response *ModifyAuditServiceResponse) {
    response = &ModifyAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计日志保存时长、审计规则等服务配置。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifyAuditService(request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    return c.ModifyAuditServiceWithContext(context.Background(), request)
}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计日志保存时长、审计规则等服务配置。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifyAuditServiceWithContext(ctx context.Context, request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    if request == nil {
        request = NewModifyAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupConfigRequest() (request *ModifyBackupConfigRequest) {
    request = &ModifyBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupConfig")
    
    
    return
}

func NewModifyBackupConfigResponse() (response *ModifyBackupConfigResponse) {
    response = &ModifyBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupConfig
// 本接口（ModifyBackupConfig）用于修改指定集群的备份配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CYNOSDBMYSQLSETBACKUPSTRATEGY = "FailedOperation.CynosdbMysqlSetBackupStrategy"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    return c.ModifyBackupConfigWithContext(context.Background(), request)
}

// ModifyBackupConfig
// 本接口（ModifyBackupConfig）用于修改指定集群的备份配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CYNOSDBMYSQLSETBACKUPSTRATEGY = "FailedOperation.CynosdbMysqlSetBackupStrategy"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupConfigWithContext(ctx context.Context, request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupDownloadRestriction")
    
    
    return
}

func NewModifyBackupDownloadRestrictionResponse() (response *ModifyBackupDownloadRestrictionResponse) {
    response = &ModifyBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupDownloadRestriction
// 该接口用于修改用户当前地域的备份文件限制下载来源，可以设置内外网均可下载、仅内网可下载，或内网指定的vpc、ip可以下载。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// 该接口用于修改用户当前地域的备份文件限制下载来源，可以设置内外网均可下载、仅内网可下载，或内网指定的vpc、ip可以下载。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupDownloadRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupDownloadUserRestrictionRequest() (request *ModifyBackupDownloadUserRestrictionRequest) {
    request = &ModifyBackupDownloadUserRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupDownloadUserRestriction")
    
    
    return
}

func NewModifyBackupDownloadUserRestrictionResponse() (response *ModifyBackupDownloadUserRestrictionResponse) {
    response = &ModifyBackupDownloadUserRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupDownloadUserRestriction
// 该接口用于修改用户当前地域的备份文件限制下载来源，可以设置内外网均可下载、仅内网可下载，或内网指定的vpc、ip可以下载。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadUserRestriction(request *ModifyBackupDownloadUserRestrictionRequest) (response *ModifyBackupDownloadUserRestrictionResponse, err error) {
    return c.ModifyBackupDownloadUserRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadUserRestriction
// 该接口用于修改用户当前地域的备份文件限制下载来源，可以设置内外网均可下载、仅内网可下载，或内网指定的vpc、ip可以下载。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupDownloadUserRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadUserRestrictionRequest) (response *ModifyBackupDownloadUserRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadUserRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupDownloadUserRestriction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadUserRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadUserRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupNameRequest() (request *ModifyBackupNameRequest) {
    request = &ModifyBackupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBackupName")
    
    
    return
}

func NewModifyBackupNameResponse() (response *ModifyBackupNameResponse) {
    response = &ModifyBackupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupName
// 此接口（ModifyBackupName）用于修改备份文件备注名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupName(request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    return c.ModifyBackupNameWithContext(context.Background(), request)
}

// ModifyBackupName
// 此接口（ModifyBackupName）用于修改备份文件备注名。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupNameWithContext(ctx context.Context, request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    if request == nil {
        request = NewModifyBackupNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBackupName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBinlogConfigRequest() (request *ModifyBinlogConfigRequest) {
    request = &ModifyBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBinlogConfig")
    
    
    return
}

func NewModifyBinlogConfigResponse() (response *ModifyBinlogConfigResponse) {
    response = &ModifyBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBinlogConfig
// 该接口（ModifyBinlogConfig）用于修改Binlog配置
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyBinlogConfig(request *ModifyBinlogConfigRequest) (response *ModifyBinlogConfigResponse, err error) {
    return c.ModifyBinlogConfigWithContext(context.Background(), request)
}

// ModifyBinlogConfig
// 该接口（ModifyBinlogConfig）用于修改Binlog配置
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyBinlogConfigWithContext(ctx context.Context, request *ModifyBinlogConfigRequest) (response *ModifyBinlogConfigResponse, err error) {
    if request == nil {
        request = NewModifyBinlogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBinlogConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBinlogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBinlogSaveDaysRequest() (request *ModifyBinlogSaveDaysRequest) {
    request = &ModifyBinlogSaveDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyBinlogSaveDays")
    
    
    return
}

func NewModifyBinlogSaveDaysResponse() (response *ModifyBinlogSaveDaysResponse) {
    response = &ModifyBinlogSaveDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBinlogSaveDays
// 此接口（ModifyBinlogSaveDays）用于修改集群Binlog保留天数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBinlogSaveDays(request *ModifyBinlogSaveDaysRequest) (response *ModifyBinlogSaveDaysResponse, err error) {
    return c.ModifyBinlogSaveDaysWithContext(context.Background(), request)
}

// ModifyBinlogSaveDays
// 此接口（ModifyBinlogSaveDays）用于修改集群Binlog保留天数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBinlogSaveDaysWithContext(ctx context.Context, request *ModifyBinlogSaveDaysRequest) (response *ModifyBinlogSaveDaysResponse, err error) {
    if request == nil {
        request = NewModifyBinlogSaveDaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyBinlogSaveDays")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBinlogSaveDays require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBinlogSaveDaysResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterDatabaseRequest() (request *ModifyClusterDatabaseRequest) {
    request = &ModifyClusterDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterDatabase")
    
    
    return
}

func NewModifyClusterDatabaseResponse() (response *ModifyClusterDatabaseResponse) {
    response = &ModifyClusterDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterDatabase
// 本接口（ModifyClusterDatabase）用于修改数据库的账号授权。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterDatabase(request *ModifyClusterDatabaseRequest) (response *ModifyClusterDatabaseResponse, err error) {
    return c.ModifyClusterDatabaseWithContext(context.Background(), request)
}

// ModifyClusterDatabase
// 本接口（ModifyClusterDatabase）用于修改数据库的账号授权。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterDatabaseWithContext(ctx context.Context, request *ModifyClusterDatabaseRequest) (response *ModifyClusterDatabaseResponse, err error) {
    if request == nil {
        request = NewModifyClusterDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterDatabaseResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterGlobalEncryptionRequest() (request *ModifyClusterGlobalEncryptionRequest) {
    request = &ModifyClusterGlobalEncryptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterGlobalEncryption")
    
    
    return
}

func NewModifyClusterGlobalEncryptionResponse() (response *ModifyClusterGlobalEncryptionResponse) {
    response = &ModifyClusterGlobalEncryptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterGlobalEncryption
// 开关全局加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterGlobalEncryption(request *ModifyClusterGlobalEncryptionRequest) (response *ModifyClusterGlobalEncryptionResponse, err error) {
    return c.ModifyClusterGlobalEncryptionWithContext(context.Background(), request)
}

// ModifyClusterGlobalEncryption
// 开关全局加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterGlobalEncryptionWithContext(ctx context.Context, request *ModifyClusterGlobalEncryptionRequest) (response *ModifyClusterGlobalEncryptionResponse, err error) {
    if request == nil {
        request = NewModifyClusterGlobalEncryptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterGlobalEncryption")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterGlobalEncryption require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterGlobalEncryptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
    request = &ModifyClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterName")
    
    
    return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
    response = &ModifyClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterName
// 本接口（ModifyClusterName）用于修改集群名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    return c.ModifyClusterNameWithContext(context.Background(), request)
}

// ModifyClusterName
// 本接口（ModifyClusterName）用于修改集群名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterNameWithContext(ctx context.Context, request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterParamRequest() (request *ModifyClusterParamRequest) {
    request = &ModifyClusterParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterParam")
    
    
    return
}

func NewModifyClusterParamResponse() (response *ModifyClusterParamResponse) {
    response = &ModifyClusterParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterParam
// 本接口（ModifyClusterParam）用于修改集群参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterParam(request *ModifyClusterParamRequest) (response *ModifyClusterParamResponse, err error) {
    return c.ModifyClusterParamWithContext(context.Background(), request)
}

// ModifyClusterParam
// 本接口（ModifyClusterParam）用于修改集群参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterParamWithContext(ctx context.Context, request *ModifyClusterParamRequest) (response *ModifyClusterParamResponse, err error) {
    if request == nil {
        request = NewModifyClusterParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterPasswordComplexityRequest() (request *ModifyClusterPasswordComplexityRequest) {
    request = &ModifyClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterPasswordComplexity")
    
    
    return
}

func NewModifyClusterPasswordComplexityResponse() (response *ModifyClusterPasswordComplexityResponse) {
    response = &ModifyClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterPasswordComplexity
// 本接口（ModifyClusterPasswordComplexity）用于修改/开启集群密码复杂度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterPasswordComplexity(request *ModifyClusterPasswordComplexityRequest) (response *ModifyClusterPasswordComplexityResponse, err error) {
    return c.ModifyClusterPasswordComplexityWithContext(context.Background(), request)
}

// ModifyClusterPasswordComplexity
// 本接口（ModifyClusterPasswordComplexity）用于修改/开启集群密码复杂度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterPasswordComplexityWithContext(ctx context.Context, request *ModifyClusterPasswordComplexityRequest) (response *ModifyClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewModifyClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterReadOnlyRequest() (request *ModifyClusterReadOnlyRequest) {
    request = &ModifyClusterReadOnlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterReadOnly")
    
    
    return
}

func NewModifyClusterReadOnlyResponse() (response *ModifyClusterReadOnlyResponse) {
    response = &ModifyClusterReadOnlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterReadOnly
// 本接口（ModifyClusterReadOnly）用于修改集群只读开关。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterReadOnly(request *ModifyClusterReadOnlyRequest) (response *ModifyClusterReadOnlyResponse, err error) {
    return c.ModifyClusterReadOnlyWithContext(context.Background(), request)
}

// ModifyClusterReadOnly
// 本接口（ModifyClusterReadOnly）用于修改集群只读开关。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  OPERATIONDENIED_TASKCONFLICTERROR = "OperationDenied.TaskConflictError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterReadOnlyWithContext(ctx context.Context, request *ModifyClusterReadOnlyRequest) (response *ModifyClusterReadOnlyResponse, err error) {
    if request == nil {
        request = NewModifyClusterReadOnlyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterReadOnly")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterReadOnly require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterReadOnlyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterSlaveZoneRequest() (request *ModifyClusterSlaveZoneRequest) {
    request = &ModifyClusterSlaveZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterSlaveZone")
    
    
    return
}

func NewModifyClusterSlaveZoneResponse() (response *ModifyClusterSlaveZoneResponse) {
    response = &ModifyClusterSlaveZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterSlaveZone
// 本接口（ModifyClusterSlaveZone）用于变更集群的备可用区。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_PAUSEDSLSNOTALLOWMODIFYSLAVE = "OperationDenied.PausedSlsNotAllowModifySlave"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterSlaveZone(request *ModifyClusterSlaveZoneRequest) (response *ModifyClusterSlaveZoneResponse, err error) {
    return c.ModifyClusterSlaveZoneWithContext(context.Background(), request)
}

// ModifyClusterSlaveZone
// 本接口（ModifyClusterSlaveZone）用于变更集群的备可用区。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_PAUSEDSLSNOTALLOWMODIFYSLAVE = "OperationDenied.PausedSlsNotAllowModifySlave"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterSlaveZoneWithContext(ctx context.Context, request *ModifyClusterSlaveZoneRequest) (response *ModifyClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewModifyClusterSlaveZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterSlaveZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterSlaveZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterStorageRequest() (request *ModifyClusterStorageRequest) {
    request = &ModifyClusterStorageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyClusterStorage")
    
    
    return
}

func NewModifyClusterStorageResponse() (response *ModifyClusterStorageResponse) {
    response = &ModifyClusterStorageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyClusterStorage
// 本接口（ModifyClusterStorage）用于调整包年包月存储容量。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterStorage(request *ModifyClusterStorageRequest) (response *ModifyClusterStorageResponse, err error) {
    return c.ModifyClusterStorageWithContext(context.Background(), request)
}

// ModifyClusterStorage
// 本接口（ModifyClusterStorage）用于调整包年包月存储容量。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterStorageWithContext(ctx context.Context, request *ModifyClusterStorageRequest) (response *ModifyClusterStorageResponse, err error) {
    if request == nil {
        request = NewModifyClusterStorageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyClusterStorage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterStorage require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterStorageResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSecurityGroups
// 本接口（ModifyDBInstanceSecurityGroups）用于修改实例绑定的安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// 本接口（ModifyDBInstanceSecurityGroups）用于修改实例绑定的安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyInstanceName")
    
    
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceName
// 本接口(ModifyInstanceName)用于修改实例名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    return c.ModifyInstanceNameWithContext(context.Background(), request)
}

// ModifyInstanceName
// 本接口(ModifyInstanceName)用于修改实例名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyInstanceName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyInstanceParam")
    
    
    return
}

func NewModifyInstanceParamResponse() (response *ModifyInstanceParamResponse) {
    response = &ModifyInstanceParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceParam
// 本接口（ModifyInstanceParam）用于修改实例参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    return c.ModifyInstanceParamWithContext(context.Background(), request)
}

// ModifyInstanceParam
// 本接口（ModifyInstanceParam）用于修改实例参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceParamWithContext(ctx context.Context, request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyInstanceParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceUpgradeLimitDaysRequest() (request *ModifyInstanceUpgradeLimitDaysRequest) {
    request = &ModifyInstanceUpgradeLimitDaysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyInstanceUpgradeLimitDays")
    
    
    return
}

func NewModifyInstanceUpgradeLimitDaysResponse() (response *ModifyInstanceUpgradeLimitDaysResponse) {
    response = &ModifyInstanceUpgradeLimitDaysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceUpgradeLimitDays
// 本接口（ModifyInstanceUpgradeLimitDays）用于修改实例内核小版本的升级限制时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyInstanceUpgradeLimitDays(request *ModifyInstanceUpgradeLimitDaysRequest) (response *ModifyInstanceUpgradeLimitDaysResponse, err error) {
    return c.ModifyInstanceUpgradeLimitDaysWithContext(context.Background(), request)
}

// ModifyInstanceUpgradeLimitDays
// 本接口（ModifyInstanceUpgradeLimitDays）用于修改实例内核小版本的升级限制时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyInstanceUpgradeLimitDaysWithContext(ctx context.Context, request *ModifyInstanceUpgradeLimitDaysRequest) (response *ModifyInstanceUpgradeLimitDaysResponse, err error) {
    if request == nil {
        request = NewModifyInstanceUpgradeLimitDaysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyInstanceUpgradeLimitDays")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceUpgradeLimitDays require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceUpgradeLimitDaysResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintainPeriodConfigRequest() (request *ModifyMaintainPeriodConfigRequest) {
    request = &ModifyMaintainPeriodConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyMaintainPeriodConfig")
    
    
    return
}

func NewModifyMaintainPeriodConfigResponse() (response *ModifyMaintainPeriodConfigResponse) {
    response = &ModifyMaintainPeriodConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMaintainPeriodConfig
// 本接口（ModifyMaintainPeriodConfig）用于修改维护时间配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintainPeriodConfig(request *ModifyMaintainPeriodConfigRequest) (response *ModifyMaintainPeriodConfigResponse, err error) {
    return c.ModifyMaintainPeriodConfigWithContext(context.Background(), request)
}

// ModifyMaintainPeriodConfig
// 本接口（ModifyMaintainPeriodConfig）用于修改维护时间配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintainPeriodConfigWithContext(ctx context.Context, request *ModifyMaintainPeriodConfigRequest) (response *ModifyMaintainPeriodConfigResponse, err error) {
    if request == nil {
        request = NewModifyMaintainPeriodConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyMaintainPeriodConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintainPeriodConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintainPeriodConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyParamTemplateRequest() (request *ModifyParamTemplateRequest) {
    request = &ModifyParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyParamTemplate")
    
    
    return
}

func NewModifyParamTemplateResponse() (response *ModifyParamTemplateResponse) {
    response = &ModifyParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyParamTemplate
// 本接口（ModifyParamTemplate）用于修改用户参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    return c.ModifyParamTemplateWithContext(context.Background(), request)
}

// ModifyParamTemplate
// 本接口（ModifyParamTemplate）用于修改用户参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CAMDENIEDERROR = "OperationDenied.CamDeniedError"
func (c *Client) ModifyParamTemplateWithContext(ctx context.Context, request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    if request == nil {
        request = NewModifyParamTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyDescRequest() (request *ModifyProxyDescRequest) {
    request = &ModifyProxyDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyProxyDesc")
    
    
    return
}

func NewModifyProxyDescResponse() (response *ModifyProxyDescResponse) {
    response = &ModifyProxyDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProxyDesc
// 本接口（ModifyProxyDesc）用于修改数据库代理描述。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyProxyDesc(request *ModifyProxyDescRequest) (response *ModifyProxyDescResponse, err error) {
    return c.ModifyProxyDescWithContext(context.Background(), request)
}

// ModifyProxyDesc
// 本接口（ModifyProxyDesc）用于修改数据库代理描述。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ModifyProxyDescWithContext(ctx context.Context, request *ModifyProxyDescRequest) (response *ModifyProxyDescResponse, err error) {
    if request == nil {
        request = NewModifyProxyDescRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyProxyDesc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProxyDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProxyDescResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProxyRwSplitRequest() (request *ModifyProxyRwSplitRequest) {
    request = &ModifyProxyRwSplitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyProxyRwSplit")
    
    
    return
}

func NewModifyProxyRwSplitResponse() (response *ModifyProxyRwSplitResponse) {
    response = &ModifyProxyRwSplitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProxyRwSplit
// 本接口（ModifyProxyRwSplit）用于配置数据库代理读写分离。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyProxyRwSplit(request *ModifyProxyRwSplitRequest) (response *ModifyProxyRwSplitResponse, err error) {
    return c.ModifyProxyRwSplitWithContext(context.Background(), request)
}

// ModifyProxyRwSplit
// 本接口（ModifyProxyRwSplit）用于配置数据库代理读写分离。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyProxyRwSplitWithContext(ctx context.Context, request *ModifyProxyRwSplitRequest) (response *ModifyProxyRwSplitResponse, err error) {
    if request == nil {
        request = NewModifyProxyRwSplitRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyProxyRwSplit")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProxyRwSplit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProxyRwSplitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePackageClustersRequest() (request *ModifyResourcePackageClustersRequest) {
    request = &ModifyResourcePackageClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyResourcePackageClusters")
    
    
    return
}

func NewModifyResourcePackageClustersResponse() (response *ModifyResourcePackageClustersResponse) {
    response = &ModifyResourcePackageClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePackageClusters
// 本接口（ModifyResourcePackageClusters）用于修改资源包与集群之间的绑定关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageClusters(request *ModifyResourcePackageClustersRequest) (response *ModifyResourcePackageClustersResponse, err error) {
    return c.ModifyResourcePackageClustersWithContext(context.Background(), request)
}

// ModifyResourcePackageClusters
// 本接口（ModifyResourcePackageClusters）用于修改资源包与集群之间的绑定关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageClustersWithContext(ctx context.Context, request *ModifyResourcePackageClustersRequest) (response *ModifyResourcePackageClustersResponse, err error) {
    if request == nil {
        request = NewModifyResourcePackageClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyResourcePackageClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePackageClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePackageClustersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePackageNameRequest() (request *ModifyResourcePackageNameRequest) {
    request = &ModifyResourcePackageNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyResourcePackageName")
    
    
    return
}

func NewModifyResourcePackageNameResponse() (response *ModifyResourcePackageNameResponse) {
    response = &ModifyResourcePackageNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePackageName
// 本接口（ModifyResourcePackageName）用于修改资源包名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageName(request *ModifyResourcePackageNameRequest) (response *ModifyResourcePackageNameResponse, err error) {
    return c.ModifyResourcePackageNameWithContext(context.Background(), request)
}

// ModifyResourcePackageName
// 本接口（ModifyResourcePackageName）用于修改资源包名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BINDSOURCEPACKAGEERROR = "FailedOperation.BindSourcePackageError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackageNameWithContext(ctx context.Context, request *ModifyResourcePackageNameRequest) (response *ModifyResourcePackageNameResponse, err error) {
    if request == nil {
        request = NewModifyResourcePackageNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyResourcePackageName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePackageName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePackageNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcePackagesDeductionPriorityRequest() (request *ModifyResourcePackagesDeductionPriorityRequest) {
    request = &ModifyResourcePackagesDeductionPriorityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyResourcePackagesDeductionPriority")
    
    
    return
}

func NewModifyResourcePackagesDeductionPriorityResponse() (response *ModifyResourcePackagesDeductionPriorityResponse) {
    response = &ModifyResourcePackagesDeductionPriorityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcePackagesDeductionPriority
// 修改已绑定资源包抵扣优先级
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_MODIFYDEDUCTIONPRIORITYERROR = "FailedOperation.ModifyDeductionPriorityError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackagesDeductionPriority(request *ModifyResourcePackagesDeductionPriorityRequest) (response *ModifyResourcePackagesDeductionPriorityResponse, err error) {
    return c.ModifyResourcePackagesDeductionPriorityWithContext(context.Background(), request)
}

// ModifyResourcePackagesDeductionPriority
// 修改已绑定资源包抵扣优先级
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_MODIFYDEDUCTIONPRIORITYERROR = "FailedOperation.ModifyDeductionPriorityError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyResourcePackagesDeductionPriorityWithContext(ctx context.Context, request *ModifyResourcePackagesDeductionPriorityRequest) (response *ModifyResourcePackagesDeductionPriorityResponse, err error) {
    if request == nil {
        request = NewModifyResourcePackagesDeductionPriorityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyResourcePackagesDeductionPriority")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcePackagesDeductionPriority require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcePackagesDeductionPriorityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServerlessStrategyRequest() (request *ModifyServerlessStrategyRequest) {
    request = &ModifyServerlessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyServerlessStrategy")
    
    
    return
}

func NewModifyServerlessStrategyResponse() (response *ModifyServerlessStrategyResponse) {
    response = &ModifyServerlessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyServerlessStrategy
// 修改serverless策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SERVERLESSSETSTRATEGYERROR = "FailedOperation.ServerlessSetStrategyError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyServerlessStrategy(request *ModifyServerlessStrategyRequest) (response *ModifyServerlessStrategyResponse, err error) {
    return c.ModifyServerlessStrategyWithContext(context.Background(), request)
}

// ModifyServerlessStrategy
// 修改serverless策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SERVERLESSSETSTRATEGYERROR = "FailedOperation.ServerlessSetStrategyError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyServerlessStrategyWithContext(ctx context.Context, request *ModifyServerlessStrategyRequest) (response *ModifyServerlessStrategyResponse, err error) {
    if request == nil {
        request = NewModifyServerlessStrategyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyServerlessStrategy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServerlessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServerlessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyVipVportRequest() (request *ModifyVipVportRequest) {
    request = &ModifyVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ModifyVipVport")
    
    
    return
}

func NewModifyVipVportResponse() (response *ModifyVipVportResponse) {
    response = &ModifyVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyVipVport
// 本接口（ModifyVipVport）用于修改实例组ip，端口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyVipVport(request *ModifyVipVportRequest) (response *ModifyVipVportResponse, err error) {
    return c.ModifyVipVportWithContext(context.Background(), request)
}

// ModifyVipVport
// 本接口（ModifyVipVport）用于修改实例组ip，端口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyVipVportWithContext(ctx context.Context, request *ModifyVipVportRequest) (response *ModifyVipVportResponse, err error) {
    if request == nil {
        request = NewModifyVipVportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ModifyVipVport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyVipVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineClusterRequest() (request *OfflineClusterRequest) {
    request = &OfflineClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OfflineCluster")
    
    
    return
}

func NewOfflineClusterResponse() (response *OfflineClusterResponse) {
    response = &OfflineClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OfflineCluster
// 本接口（OfflineCluster）用于销毁集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineCluster(request *OfflineClusterRequest) (response *OfflineClusterResponse, err error) {
    return c.OfflineClusterWithContext(context.Background(), request)
}

// OfflineCluster
// 本接口（OfflineCluster）用于销毁集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineClusterWithContext(ctx context.Context, request *OfflineClusterRequest) (response *OfflineClusterResponse, err error) {
    if request == nil {
        request = NewOfflineClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OfflineCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineClusterResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineInstanceRequest() (request *OfflineInstanceRequest) {
    request = &OfflineInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OfflineInstance")
    
    
    return
}

func NewOfflineInstanceResponse() (response *OfflineInstanceResponse) {
    response = &OfflineInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OfflineInstance
// 本接口（OfflineInstance）用于销毁实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineInstance(request *OfflineInstanceRequest) (response *OfflineInstanceResponse, err error) {
    return c.OfflineInstanceWithContext(context.Background(), request)
}

// OfflineInstance
// 本接口（OfflineInstance）用于销毁实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCENOTFOUND_RESOURCEERROR = "ResourceNotFound.ResourceError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OfflineInstanceWithContext(ctx context.Context, request *OfflineInstanceRequest) (response *OfflineInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OfflineInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAuditServiceRequest() (request *OpenAuditServiceRequest) {
    request = &OpenAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenAuditService")
    
    
    return
}

func NewOpenAuditServiceResponse() (response *OpenAuditServiceResponse) {
    response = &OpenAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenAuditService
// 本接口（OpenAuditService）用于为实例开通数据库审计服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// 本接口（OpenAuditService）用于为实例开通数据库审计服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClusterPasswordComplexityRequest() (request *OpenClusterPasswordComplexityRequest) {
    request = &OpenClusterPasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenClusterPasswordComplexity")
    
    
    return
}

func NewOpenClusterPasswordComplexityResponse() (response *OpenClusterPasswordComplexityResponse) {
    response = &OpenClusterPasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClusterPasswordComplexity
// 本接口（OpenClusterPasswordComplexity）用于开启自定义密码复杂度功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterPasswordComplexity(request *OpenClusterPasswordComplexityRequest) (response *OpenClusterPasswordComplexityResponse, err error) {
    return c.OpenClusterPasswordComplexityWithContext(context.Background(), request)
}

// OpenClusterPasswordComplexity
// 本接口（OpenClusterPasswordComplexity）用于开启自定义密码复杂度功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterPasswordComplexityWithContext(ctx context.Context, request *OpenClusterPasswordComplexityRequest) (response *OpenClusterPasswordComplexityResponse, err error) {
    if request == nil {
        request = NewOpenClusterPasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenClusterPasswordComplexity")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClusterPasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClusterPasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClusterReadOnlyInstanceGroupAccessRequest() (request *OpenClusterReadOnlyInstanceGroupAccessRequest) {
    request = &OpenClusterReadOnlyInstanceGroupAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenClusterReadOnlyInstanceGroupAccess")
    
    
    return
}

func NewOpenClusterReadOnlyInstanceGroupAccessResponse() (response *OpenClusterReadOnlyInstanceGroupAccessResponse) {
    response = &OpenClusterReadOnlyInstanceGroupAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClusterReadOnlyInstanceGroupAccess
// 本接口（OpenClusterReadOnlyInstanceGroupAccess）用于开启只读实例组接入。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterReadOnlyInstanceGroupAccess(request *OpenClusterReadOnlyInstanceGroupAccessRequest) (response *OpenClusterReadOnlyInstanceGroupAccessResponse, err error) {
    return c.OpenClusterReadOnlyInstanceGroupAccessWithContext(context.Background(), request)
}

// OpenClusterReadOnlyInstanceGroupAccess
// 本接口（OpenClusterReadOnlyInstanceGroupAccess）用于开启只读实例组接入。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterReadOnlyInstanceGroupAccessWithContext(ctx context.Context, request *OpenClusterReadOnlyInstanceGroupAccessRequest) (response *OpenClusterReadOnlyInstanceGroupAccessResponse, err error) {
    if request == nil {
        request = NewOpenClusterReadOnlyInstanceGroupAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenClusterReadOnlyInstanceGroupAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClusterReadOnlyInstanceGroupAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClusterReadOnlyInstanceGroupAccessResponse()
    err = c.Send(request, response)
    return
}

func NewOpenClusterTransparentEncryptRequest() (request *OpenClusterTransparentEncryptRequest) {
    request = &OpenClusterTransparentEncryptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenClusterTransparentEncrypt")
    
    
    return
}

func NewOpenClusterTransparentEncryptResponse() (response *OpenClusterTransparentEncryptResponse) {
    response = &OpenClusterTransparentEncryptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenClusterTransparentEncrypt
// 开通集群透明加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterTransparentEncrypt(request *OpenClusterTransparentEncryptRequest) (response *OpenClusterTransparentEncryptResponse, err error) {
    return c.OpenClusterTransparentEncryptWithContext(context.Background(), request)
}

// OpenClusterTransparentEncrypt
// 开通集群透明加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenClusterTransparentEncryptWithContext(ctx context.Context, request *OpenClusterTransparentEncryptRequest) (response *OpenClusterTransparentEncryptResponse, err error) {
    if request == nil {
        request = NewOpenClusterTransparentEncryptRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenClusterTransparentEncrypt")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenClusterTransparentEncrypt require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenClusterTransparentEncryptResponse()
    err = c.Send(request, response)
    return
}

func NewOpenReadOnlyInstanceExclusiveAccessRequest() (request *OpenReadOnlyInstanceExclusiveAccessRequest) {
    request = &OpenReadOnlyInstanceExclusiveAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenReadOnlyInstanceExclusiveAccess")
    
    
    return
}

func NewOpenReadOnlyInstanceExclusiveAccessResponse() (response *OpenReadOnlyInstanceExclusiveAccessResponse) {
    response = &OpenReadOnlyInstanceExclusiveAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenReadOnlyInstanceExclusiveAccess
// 本接口（OpenReadOnlyInstanceExclusiveAccess）用于开通只读实例独有访问接入组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenReadOnlyInstanceExclusiveAccess(request *OpenReadOnlyInstanceExclusiveAccessRequest) (response *OpenReadOnlyInstanceExclusiveAccessResponse, err error) {
    return c.OpenReadOnlyInstanceExclusiveAccessWithContext(context.Background(), request)
}

// OpenReadOnlyInstanceExclusiveAccess
// 本接口（OpenReadOnlyInstanceExclusiveAccess）用于开通只读实例独有访问接入组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenReadOnlyInstanceExclusiveAccessWithContext(ctx context.Context, request *OpenReadOnlyInstanceExclusiveAccessRequest) (response *OpenReadOnlyInstanceExclusiveAccessResponse, err error) {
    if request == nil {
        request = NewOpenReadOnlyInstanceExclusiveAccessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenReadOnlyInstanceExclusiveAccess")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenReadOnlyInstanceExclusiveAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenReadOnlyInstanceExclusiveAccessResponse()
    err = c.Send(request, response)
    return
}

func NewOpenSSLRequest() (request *OpenSSLRequest) {
    request = &OpenSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenSSL")
    
    
    return
}

func NewOpenSSLResponse() (response *OpenSSLResponse) {
    response = &OpenSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenSSL
// 开启SSL加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenSSL(request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    return c.OpenSSLWithContext(context.Background(), request)
}

// OpenSSL
// 开启SSL加密
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenSSLWithContext(ctx context.Context, request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    if request == nil {
        request = NewOpenSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenSSLResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWanRequest() (request *OpenWanRequest) {
    request = &OpenWanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "OpenWan")
    
    
    return
}

func NewOpenWanResponse() (response *OpenWanResponse) {
    response = &OpenWanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenWan
// 本接口（OpenWan）用于开通外网。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenWan(request *OpenWanRequest) (response *OpenWanResponse, err error) {
    return c.OpenWanWithContext(context.Background(), request)
}

// OpenWan
// 本接口（OpenWan）用于开通外网。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_GETNETSERVICEINFOERROR = "FailedOperation.GetNetServiceInfoError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenWanWithContext(ctx context.Context, request *OpenWanRequest) (response *OpenWanResponse, err error) {
    if request == nil {
        request = NewOpenWanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "OpenWan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWan require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWanResponse()
    err = c.Send(request, response)
    return
}

func NewPauseServerlessRequest() (request *PauseServerlessRequest) {
    request = &PauseServerlessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "PauseServerless")
    
    
    return
}

func NewPauseServerlessResponse() (response *PauseServerlessResponse) {
    response = &PauseServerlessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseServerless
// 本接口（PauseServerless）用于暂停 serverless 集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) PauseServerless(request *PauseServerlessRequest) (response *PauseServerlessResponse, err error) {
    return c.PauseServerlessWithContext(context.Background(), request)
}

// PauseServerless
// 本接口（PauseServerless）用于暂停 serverless 集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) PauseServerlessWithContext(ctx context.Context, request *PauseServerlessRequest) (response *PauseServerlessResponse, err error) {
    if request == nil {
        request = NewPauseServerlessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "PauseServerless")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseServerless require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseServerlessResponse()
    err = c.Send(request, response)
    return
}

func NewRefundResourcePackageRequest() (request *RefundResourcePackageRequest) {
    request = &RefundResourcePackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RefundResourcePackage")
    
    
    return
}

func NewRefundResourcePackageResponse() (response *RefundResourcePackageResponse) {
    response = &RefundResourcePackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefundResourcePackage
// 本接口（RefundResourcePackage）用于资源包退款。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REFUNDSOURCEPACKAGEERROR = "FailedOperation.RefundSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RefundResourcePackage(request *RefundResourcePackageRequest) (response *RefundResourcePackageResponse, err error) {
    return c.RefundResourcePackageWithContext(context.Background(), request)
}

// RefundResourcePackage
// 本接口（RefundResourcePackage）用于资源包退款。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REFUNDSOURCEPACKAGEERROR = "FailedOperation.RefundSourcePackageError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RefundResourcePackageWithContext(ctx context.Context, request *RefundResourcePackageRequest) (response *RefundResourcePackageResponse, err error) {
    if request == nil {
        request = NewRefundResourcePackageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RefundResourcePackage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefundResourcePackage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefundResourcePackageResponse()
    err = c.Send(request, response)
    return
}

func NewReloadBalanceProxyNodeRequest() (request *ReloadBalanceProxyNodeRequest) {
    request = &ReloadBalanceProxyNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ReloadBalanceProxyNode")
    
    
    return
}

func NewReloadBalanceProxyNodeResponse() (response *ReloadBalanceProxyNodeResponse) {
    response = &ReloadBalanceProxyNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReloadBalanceProxyNode
// 本接口（ReloadBalanceProxyNode）用于负载均衡数据库代理。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ReloadBalanceProxyNode(request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    return c.ReloadBalanceProxyNodeWithContext(context.Background(), request)
}

// ReloadBalanceProxyNode
// 本接口（ReloadBalanceProxyNode）用于负载均衡数据库代理。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ReloadBalanceProxyNodeWithContext(ctx context.Context, request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    if request == nil {
        request = NewReloadBalanceProxyNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ReloadBalanceProxyNode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReloadBalanceProxyNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewReloadBalanceProxyNodeResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveClusterSlaveZoneRequest() (request *RemoveClusterSlaveZoneRequest) {
    request = &RemoveClusterSlaveZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RemoveClusterSlaveZone")
    
    
    return
}

func NewRemoveClusterSlaveZoneResponse() (response *RemoveClusterSlaveZoneResponse) {
    response = &RemoveClusterSlaveZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveClusterSlaveZone
// 本接口（RemoveClusterSlaveZone）用于关闭集群多可用区部署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveClusterSlaveZone(request *RemoveClusterSlaveZoneRequest) (response *RemoveClusterSlaveZoneResponse, err error) {
    return c.RemoveClusterSlaveZoneWithContext(context.Background(), request)
}

// RemoveClusterSlaveZone
// 本接口（RemoveClusterSlaveZone）用于关闭集群多可用区部署。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveClusterSlaveZoneWithContext(ctx context.Context, request *RemoveClusterSlaveZoneRequest) (response *RemoveClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewRemoveClusterSlaveZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RemoveClusterSlaveZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveClusterSlaveZoneResponse()
    err = c.Send(request, response)
    return
}

func NewRenewClustersRequest() (request *RenewClustersRequest) {
    request = &RenewClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RenewClusters")
    
    
    return
}

func NewRenewClustersResponse() (response *RenewClustersResponse) {
    response = &RenewClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewClusters
// 续费集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewClusters(request *RenewClustersRequest) (response *RenewClustersResponse, err error) {
    return c.RenewClustersWithContext(context.Background(), request)
}

// RenewClusters
// 续费集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDREGIONIDERROR = "InvalidParameterValue.InvalidRegionIdError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewClustersWithContext(ctx context.Context, request *RenewClustersRequest) (response *RenewClustersResponse, err error) {
    if request == nil {
        request = NewRenewClustersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RenewClusters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewClustersResponse()
    err = c.Send(request, response)
    return
}

func NewReplayInstanceAuditLogRequest() (request *ReplayInstanceAuditLogRequest) {
    request = &ReplayInstanceAuditLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ReplayInstanceAuditLog")
    
    
    return
}

func NewReplayInstanceAuditLogResponse() (response *ReplayInstanceAuditLogResponse) {
    response = &ReplayInstanceAuditLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplayInstanceAuditLog
// 回放实例审计日志
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITLOGCLOSEDERROR = "OperationDenied.AuditLogClosedError"
//  OPERATIONDENIED_FEATURENOTSUPPORTERROR = "OperationDenied.FeatureNotSupportError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_ISNOTROLLBACKCLUSTERERROR = "OperationDenied.IsNotRollbackClusterError"
//  OPERATIONDENIED_LIMITDAYFORAUDITREPLAYERROR = "OperationDenied.LimitDayForAuditReplayError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ReplayInstanceAuditLog(request *ReplayInstanceAuditLogRequest) (response *ReplayInstanceAuditLogResponse, err error) {
    return c.ReplayInstanceAuditLogWithContext(context.Background(), request)
}

// ReplayInstanceAuditLog
// 回放实例审计日志
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITLOGCLOSEDERROR = "OperationDenied.AuditLogClosedError"
//  OPERATIONDENIED_FEATURENOTSUPPORTERROR = "OperationDenied.FeatureNotSupportError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_ISNOTROLLBACKCLUSTERERROR = "OperationDenied.IsNotRollbackClusterError"
//  OPERATIONDENIED_LIMITDAYFORAUDITREPLAYERROR = "OperationDenied.LimitDayForAuditReplayError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) ReplayInstanceAuditLogWithContext(ctx context.Context, request *ReplayInstanceAuditLogRequest) (response *ReplayInstanceAuditLogResponse, err error) {
    if request == nil {
        request = NewReplayInstanceAuditLogRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ReplayInstanceAuditLog")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplayInstanceAuditLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplayInstanceAuditLogResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于修改数据库账号密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于修改数据库账号密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPASSWORDVALUEERROR = "InvalidParameterValue.InvalidPasswordValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ResetAccountPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RestartInstance")
    
    
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartInstance
// 本接口（RestartInstance）用于重启实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// 本接口（RestartInstance）用于重启实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RestartInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResumeServerlessRequest() (request *ResumeServerlessRequest) {
    request = &ResumeServerlessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "ResumeServerless")
    
    
    return
}

func NewResumeServerlessResponse() (response *ResumeServerlessResponse) {
    response = &ResumeServerlessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeServerless
// 本接口（ResumeServerless）用于恢复 serverless 集群（启动暂停的集群）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResumeServerless(request *ResumeServerlessRequest) (response *ResumeServerlessResponse, err error) {
    return c.ResumeServerlessWithContext(context.Background(), request)
}

// ResumeServerless
// 本接口（ResumeServerless）用于恢复 serverless 集群（启动暂停的集群）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResumeServerlessWithContext(ctx context.Context, request *ResumeServerlessRequest) (response *ResumeServerlessResponse, err error) {
    if request == nil {
        request = NewResumeServerlessRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "ResumeServerless")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeServerless require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeServerlessResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeAccountPrivilegesRequest() (request *RevokeAccountPrivilegesRequest) {
    request = &RevokeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RevokeAccountPrivileges")
    
    
    return
}

func NewRevokeAccountPrivilegesResponse() (response *RevokeAccountPrivilegesResponse) {
    response = &RevokeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RevokeAccountPrivileges
// 本接口（RevokeAccountPrivileges）用于批量回收账号权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RevokeAccountPrivileges(request *RevokeAccountPrivilegesRequest) (response *RevokeAccountPrivilegesResponse, err error) {
    return c.RevokeAccountPrivilegesWithContext(context.Background(), request)
}

// RevokeAccountPrivileges
// 本接口（RevokeAccountPrivileges）用于批量回收账号权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RevokeAccountPrivilegesWithContext(ctx context.Context, request *RevokeAccountPrivilegesRequest) (response *RevokeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewRevokeAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RevokeAccountPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RevokeAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewRevokeAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewRollBackClusterRequest() (request *RollBackClusterRequest) {
    request = &RollBackClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RollBackCluster")
    
    
    return
}

func NewRollBackClusterResponse() (response *RollBackClusterResponse) {
    response = &RollBackClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollBackCluster
// 本接口（RollBackCluster）用于集群回档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollBackCluster(request *RollBackClusterRequest) (response *RollBackClusterResponse, err error) {
    return c.RollBackClusterWithContext(context.Background(), request)
}

// RollBackCluster
// 本接口（RollBackCluster）用于集群回档。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollBackClusterWithContext(ctx context.Context, request *RollBackClusterRequest) (response *RollBackClusterResponse, err error) {
    if request == nil {
        request = NewRollBackClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RollBackCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollBackCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollBackClusterResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackToNewClusterRequest() (request *RollbackToNewClusterRequest) {
    request = &RollbackToNewClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "RollbackToNewCluster")
    
    
    return
}

func NewRollbackToNewClusterResponse() (response *RollbackToNewClusterResponse) {
    response = &RollbackToNewClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RollbackToNewCluster
// 本接口（RollbackToNewCluster）用于回档到新集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollbackToNewCluster(request *RollbackToNewClusterRequest) (response *RollbackToNewClusterResponse, err error) {
    return c.RollbackToNewClusterWithContext(context.Background(), request)
}

// RollbackToNewCluster
// 本接口（RollbackToNewCluster）用于回档到新集群。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_QUERYSPECBYSPECCODEERROR = "FailedOperation.QuerySpecBySpecCodeError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_INVALIDZONEIDERROR = "InvalidParameterValue.InvalidZoneIdError"
//  INVALIDPARAMETERVALUE_PROJECTIDNOTFOUND = "InvalidParameterValue.ProjectIdNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollbackToNewClusterWithContext(ctx context.Context, request *RollbackToNewClusterRequest) (response *RollbackToNewClusterResponse, err error) {
    if request == nil {
        request = NewRollbackToNewClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "RollbackToNewCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackToNewCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackToNewClusterResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClusterDatabasesRequest() (request *SearchClusterDatabasesRequest) {
    request = &SearchClusterDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SearchClusterDatabases")
    
    
    return
}

func NewSearchClusterDatabasesResponse() (response *SearchClusterDatabasesResponse) {
    response = &SearchClusterDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClusterDatabases
// 本接口（SearchClusterDatabases）用于搜索集群数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterDatabases(request *SearchClusterDatabasesRequest) (response *SearchClusterDatabasesResponse, err error) {
    return c.SearchClusterDatabasesWithContext(context.Background(), request)
}

// SearchClusterDatabases
// 本接口（SearchClusterDatabases）用于搜索集群数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXISTERROR = "InvalidParameterValue.AccountNotExistError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterDatabasesWithContext(ctx context.Context, request *SearchClusterDatabasesRequest) (response *SearchClusterDatabasesResponse, err error) {
    if request == nil {
        request = NewSearchClusterDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SearchClusterDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClusterDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClusterDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClusterTablesRequest() (request *SearchClusterTablesRequest) {
    request = &SearchClusterTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SearchClusterTables")
    
    
    return
}

func NewSearchClusterTablesResponse() (response *SearchClusterTablesResponse) {
    response = &SearchClusterTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClusterTables
// 本接口（SearchClusterTables）用于搜索集群数据表列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterTables(request *SearchClusterTablesRequest) (response *SearchClusterTablesResponse, err error) {
    return c.SearchClusterTablesWithContext(context.Background(), request)
}

// SearchClusterTables
// 本接口（SearchClusterTables）用于搜索集群数据表列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_INSTANCESTATUSLIMITERROR = "OperationDenied.InstanceStatusLimitError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterTablesWithContext(ctx context.Context, request *SearchClusterTablesRequest) (response *SearchClusterTablesResponse, err error) {
    if request == nil {
        request = NewSearchClusterTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SearchClusterTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClusterTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClusterTablesResponse()
    err = c.Send(request, response)
    return
}

func NewSetRenewFlagRequest() (request *SetRenewFlagRequest) {
    request = &SetRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SetRenewFlag")
    
    
    return
}

func NewSetRenewFlagResponse() (response *SetRenewFlagResponse) {
    response = &SetRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetRenewFlag
// 本接口（SetRenewFlag）用于设置实例的自动续费功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SetRenewFlag(request *SetRenewFlagRequest) (response *SetRenewFlagResponse, err error) {
    return c.SetRenewFlagWithContext(context.Background(), request)
}

// SetRenewFlag
// 本接口（SetRenewFlag）用于设置实例的自动续费功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SetRenewFlagWithContext(ctx context.Context, request *SetRenewFlagRequest) (response *SetRenewFlagResponse, err error) {
    if request == nil {
        request = NewSetRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SetRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewStartCLSDeliveryRequest() (request *StartCLSDeliveryRequest) {
    request = &StartCLSDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "StartCLSDelivery")
    
    
    return
}

func NewStartCLSDeliveryResponse() (response *StartCLSDeliveryResponse) {
    response = &StartCLSDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCLSDelivery
// 本接口（StartCLSDelivery）用于开启日志投递功能。
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartCLSDelivery(request *StartCLSDeliveryRequest) (response *StartCLSDeliveryResponse, err error) {
    return c.StartCLSDeliveryWithContext(context.Background(), request)
}

// StartCLSDelivery
// 本接口（StartCLSDelivery）用于开启日志投递功能。
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartCLSDeliveryWithContext(ctx context.Context, request *StartCLSDeliveryRequest) (response *StartCLSDeliveryResponse, err error) {
    if request == nil {
        request = NewStartCLSDeliveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "StartCLSDelivery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCLSDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCLSDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewStopCLSDeliveryRequest() (request *StopCLSDeliveryRequest) {
    request = &StopCLSDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "StopCLSDelivery")
    
    
    return
}

func NewStopCLSDeliveryResponse() (response *StopCLSDeliveryResponse) {
    response = &StopCLSDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCLSDelivery
// 本接口（StopCLSDelivery）用于停止日志投递功能。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StopCLSDelivery(request *StopCLSDeliveryRequest) (response *StopCLSDeliveryResponse, err error) {
    return c.StopCLSDeliveryWithContext(context.Background(), request)
}

// StopCLSDelivery
// 本接口（StopCLSDelivery）用于停止日志投递功能。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StopCLSDeliveryWithContext(ctx context.Context, request *StopCLSDeliveryRequest) (response *StopCLSDeliveryResponse, err error) {
    if request == nil {
        request = NewStopCLSDeliveryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "StopCLSDelivery")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCLSDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCLSDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchClusterVpcRequest() (request *SwitchClusterVpcRequest) {
    request = &SwitchClusterVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SwitchClusterVpc")
    
    
    return
}

func NewSwitchClusterVpcResponse() (response *SwitchClusterVpcResponse) {
    response = &SwitchClusterVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchClusterVpc
// 本接口（SwitchClusterVpc）用于更换集群vpc。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterVpc(request *SwitchClusterVpcRequest) (response *SwitchClusterVpcResponse, err error) {
    return c.SwitchClusterVpcWithContext(context.Background(), request)
}

// SwitchClusterVpc
// 本接口（SwitchClusterVpc）用于更换集群vpc。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterVpcWithContext(ctx context.Context, request *SwitchClusterVpcRequest) (response *SwitchClusterVpcResponse, err error) {
    if request == nil {
        request = NewSwitchClusterVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SwitchClusterVpc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchClusterVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchClusterVpcResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchClusterZoneRequest() (request *SwitchClusterZoneRequest) {
    request = &SwitchClusterZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SwitchClusterZone")
    
    
    return
}

func NewSwitchClusterZoneResponse() (response *SwitchClusterZoneResponse) {
    response = &SwitchClusterZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchClusterZone
// 本接口（SwitchClusterZone）用于切换集群的主备可用区。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterZone(request *SwitchClusterZoneRequest) (response *SwitchClusterZoneResponse, err error) {
    return c.SwitchClusterZoneWithContext(context.Background(), request)
}

// SwitchClusterZone
// 本接口（SwitchClusterZone）用于切换集群的主备可用区。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchClusterZoneWithContext(ctx context.Context, request *SwitchClusterZoneRequest) (response *SwitchClusterZoneResponse, err error) {
    if request == nil {
        request = NewSwitchClusterZoneRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SwitchClusterZone")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchClusterZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchClusterZoneResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchProxyVpcRequest() (request *SwitchProxyVpcRequest) {
    request = &SwitchProxyVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "SwitchProxyVpc")
    
    
    return
}

func NewSwitchProxyVpcResponse() (response *SwitchProxyVpcResponse) {
    response = &SwitchProxyVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchProxyVpc
// 本接口（SwitchProxyVpc）用于更换数据库代理vpc。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchProxyVpc(request *SwitchProxyVpcRequest) (response *SwitchProxyVpcResponse, err error) {
    return c.SwitchProxyVpcWithContext(context.Background(), request)
}

// SwitchProxyVpc
// 本接口（SwitchProxyVpc）用于更换数据库代理vpc。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchProxyVpcWithContext(ctx context.Context, request *SwitchProxyVpcRequest) (response *SwitchProxyVpcResponse, err error) {
    if request == nil {
        request = NewSwitchProxyVpcRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "SwitchProxyVpc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchProxyVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchProxyVpcResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindClusterResourcePackagesRequest() (request *UnbindClusterResourcePackagesRequest) {
    request = &UnbindClusterResourcePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UnbindClusterResourcePackages")
    
    
    return
}

func NewUnbindClusterResourcePackagesResponse() (response *UnbindClusterResourcePackagesResponse) {
    response = &UnbindClusterResourcePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindClusterResourcePackages
// 本接口（UnbindClusterResourcePackages）用于解除资源包与集群之间的绑定关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UnbindClusterResourcePackages(request *UnbindClusterResourcePackagesRequest) (response *UnbindClusterResourcePackagesResponse, err error) {
    return c.UnbindClusterResourcePackagesWithContext(context.Background(), request)
}

// UnbindClusterResourcePackages
// 本接口（UnbindClusterResourcePackages）用于解除资源包与集群之间的绑定关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYSOURCEPACKAGEERROR = "FailedOperation.QuerySourcePackageError"
//  FAILEDOPERATION_UNBINDSOURCEPACKAGEERROR = "FailedOperation.UnBindSourcePackageError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UnbindClusterResourcePackagesWithContext(ctx context.Context, request *UnbindClusterResourcePackagesRequest) (response *UnbindClusterResourcePackagesResponse, err error) {
    if request == nil {
        request = NewUnbindClusterResourcePackagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UnbindClusterResourcePackages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindClusterResourcePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindClusterResourcePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeClusterVersionRequest() (request *UpgradeClusterVersionRequest) {
    request = &UpgradeClusterVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeClusterVersion")
    
    
    return
}

func NewUpgradeClusterVersionResponse() (response *UpgradeClusterVersionResponse) {
    response = &UpgradeClusterVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeClusterVersion
// 本接口（UpgradeClusterVersion）用于更新内核小版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeClusterVersion(request *UpgradeClusterVersionRequest) (response *UpgradeClusterVersionResponse, err error) {
    return c.UpgradeClusterVersionWithContext(context.Background(), request)
}

// UpgradeClusterVersion
// 本接口（UpgradeClusterVersion）用于更新内核小版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_OPERATEWANFAIL = "InternalError.OperateWanFail"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_ISOLATENOTALLOWED = "InvalidParameter.IsolateNotAllowed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INTERNALACCOUNT = "InvalidParameterValue.InternalAccount"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeClusterVersionWithContext(ctx context.Context, request *UpgradeClusterVersionRequest) (response *UpgradeClusterVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeClusterVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeClusterVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeClusterVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeClusterVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstance
// 本接口（UpgradeInstance）用于实例变配。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 本接口（UpgradeInstance）用于实例变配。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDER = "FailedOperation.CreateOrder"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_TRADECREATEORDERERROR = "FailedOperation.TradeCreateOrderError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DBTYPENOTFOUND = "InvalidParameterValue.DBTypeNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_CLUSTEROPNOTALLOWEDERROR = "OperationDenied.ClusterOpNotAllowedError"
//  OPERATIONDENIED_INSTANCESTATUSDENIEDERROR = "OperationDenied.InstanceStatusDeniedError"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeProxyRequest() (request *UpgradeProxyRequest) {
    request = &UpgradeProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeProxy")
    
    
    return
}

func NewUpgradeProxyResponse() (response *UpgradeProxyResponse) {
    response = &UpgradeProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeProxy
// 本接口（UpgradeProxy）用于升级数据库代理配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SPECNOTFOUNDERROR = "FailedOperation.SpecNotFoundError"
//  OPERATIONDENIED_GETPROXYGROUPFAILEDERROR = "OperationDenied.GetProxyGroupFailedError"
//  OPERATIONDENIED_PROXYCONNECTCOUNTCHECKERROR = "OperationDenied.ProxyConnectCountCheckError"
//  OPERATIONDENIED_PROXYNODECOUNTCHECKERROR = "OperationDenied.ProxyNodeCountCheckError"
//  OPERATIONDENIED_PROXYNOTRUNNINGERROR = "OperationDenied.ProxyNotRunningError"
//  OPERATIONDENIED_PROXYSALEZONECHECKERROR = "OperationDenied.ProxySaleZoneCheckError"
//  OPERATIONDENIED_PROXYVERSIONCHECKERROR = "OperationDenied.ProxyVersionCheckError"
//  OPERATIONDENIED_PROXYZONECHECKERROR = "OperationDenied.ProxyZoneCheckError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeProxy(request *UpgradeProxyRequest) (response *UpgradeProxyResponse, err error) {
    return c.UpgradeProxyWithContext(context.Background(), request)
}

// UpgradeProxy
// 本接口（UpgradeProxy）用于升级数据库代理配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  FAILEDOPERATION_SPECNOTFOUNDERROR = "FailedOperation.SpecNotFoundError"
//  OPERATIONDENIED_GETPROXYGROUPFAILEDERROR = "OperationDenied.GetProxyGroupFailedError"
//  OPERATIONDENIED_PROXYCONNECTCOUNTCHECKERROR = "OperationDenied.ProxyConnectCountCheckError"
//  OPERATIONDENIED_PROXYNODECOUNTCHECKERROR = "OperationDenied.ProxyNodeCountCheckError"
//  OPERATIONDENIED_PROXYNOTRUNNINGERROR = "OperationDenied.ProxyNotRunningError"
//  OPERATIONDENIED_PROXYSALEZONECHECKERROR = "OperationDenied.ProxySaleZoneCheckError"
//  OPERATIONDENIED_PROXYVERSIONCHECKERROR = "OperationDenied.ProxyVersionCheckError"
//  OPERATIONDENIED_PROXYZONECHECKERROR = "OperationDenied.ProxyZoneCheckError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeProxyWithContext(ctx context.Context, request *UpgradeProxyRequest) (response *UpgradeProxyResponse, err error) {
    if request == nil {
        request = NewUpgradeProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeProxyResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeProxyVersionRequest() (request *UpgradeProxyVersionRequest) {
    request = &UpgradeProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "UpgradeProxyVersion")
    
    
    return
}

func NewUpgradeProxyVersionResponse() (response *UpgradeProxyVersionResponse) {
    response = &UpgradeProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeProxyVersion
// 本接口（UpgradeProxyVersion）用于升级数据库代理版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_PROXYSTOCKCHECKERROR = "OperationDenied.ProxyStockCheckError"
func (c *Client) UpgradeProxyVersion(request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    return c.UpgradeProxyVersionWithContext(context.Background(), request)
}

// UpgradeProxyVersion
// 本接口（UpgradeProxyVersion）用于升级数据库代理版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_PROXYSTOCKCHECKERROR = "OperationDenied.ProxyStockCheckError"
func (c *Client) UpgradeProxyVersionWithContext(ctx context.Context, request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeProxyVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cynosdb", APIVersion, "UpgradeProxyVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeProxyVersionResponse()
    err = c.Send(request, response)
    return
}
