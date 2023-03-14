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
// 本接口(ActivateInstance)用于恢复已隔离的实例访问。
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
// 本接口(ActivateInstance)用于恢复已隔离的实例访问。
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
// 增加从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddClusterSlaveZone(request *AddClusterSlaveZoneRequest) (response *AddClusterSlaveZoneResponse, err error) {
    return c.AddClusterSlaveZoneWithContext(context.Background(), request)
}

// AddClusterSlaveZone
// 增加从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AddClusterSlaveZoneWithContext(ctx context.Context, request *AddClusterSlaveZoneRequest) (response *AddClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewAddClusterSlaveZoneRequest()
    }
    
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
// 本接口（AddInstances）用于集群添加实例
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
// 本接口（AddInstances）用于集群添加实例
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
// 安全组批量绑定云资源
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
// 安全组批量绑定云资源
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssociateSecurityGroupsResponse()
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
// TDSQL-C for MySQL实例关闭审计服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseAuditService(request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    return c.CloseAuditServiceWithContext(context.Background(), request)
}

// CloseAuditService
// TDSQL-C for MySQL实例关闭审计服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_LISTINSTANCEFAILED = "InternalError.ListInstanceFailed"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseAuditServiceWithContext(ctx context.Context, request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    if request == nil {
        request = NewCloseAuditServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAuditServiceResponse()
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
// 创建账号
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
// 创建账号
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
func (c *Client) CreateAuditLogFile(request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    return c.CreateAuditLogFileWithContext(context.Background(), request)
}

// CreateAuditLogFile
// 本接口(CreateAuditLogFile)用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
func (c *Client) CreateAuditLogFileWithContext(ctx context.Context, request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    if request == nil {
        request = NewCreateAuditLogFileRequest()
    }
    
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
// 创建审计规则模版
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) CreateAuditRuleTemplate(request *CreateAuditRuleTemplateRequest) (response *CreateAuditRuleTemplateResponse, err error) {
    return c.CreateAuditRuleTemplateWithContext(context.Background(), request)
}

// CreateAuditRuleTemplate
// 创建审计规则模版
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) CreateAuditRuleTemplateWithContext(ctx context.Context, request *CreateAuditRuleTemplateRequest) (response *CreateAuditRuleTemplateResponse, err error) {
    if request == nil {
        request = NewCreateAuditRuleTemplateRequest()
    }
    
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
// 为集群创建手动备份
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// 为集群创建手动备份
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupResponse()
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
// 创建集群
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
// 创建集群
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClustersResponse()
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
// 本接口(DeleteAuditLogFile)用于删除云数据库实例的审计日志文件。
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
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAuditLogFile(request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    return c.DeleteAuditLogFileWithContext(context.Background(), request)
}

// DeleteAuditLogFile
// 本接口(DeleteAuditLogFile)用于删除云数据库实例的审计日志文件。
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
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAuditLogFileWithContext(ctx context.Context, request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    if request == nil {
        request = NewDeleteAuditLogFileRequest()
    }
    
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
// 删除审计规则模版
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
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAuditRuleTemplates(request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    return c.DeleteAuditRuleTemplatesWithContext(context.Background(), request)
}

// DeleteAuditRuleTemplates
// 删除审计规则模版
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
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  LIMITEXCEEDED_USERINSTANCELIMIT = "LimitExceeded.UserInstanceLimit"
//  OPERATIONDENIED_INSUFFICIENTBALANCEERROR = "OperationDenied.InsufficientBalanceError"
//  OPERATIONDENIED_USERNOTAUTHENTICATEDERROR = "OperationDenied.UserNotAuthenticatedError"
//  UNAUTHORIZEDOPERATION_NOTREALNAMEACCOUNT = "UnauthorizedOperation.NotRealNameAccount"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAuditRuleTemplatesWithContext(ctx context.Context, request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRuleTemplatesRequest()
    }
    
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
// 为集群删除手动备份，无法删除自动备份
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
// 为集群删除手动备份，无法删除自动备份
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupResponse()
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
// 账号所有权限
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
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountAllGrantPrivileges(request *DescribeAccountAllGrantPrivilegesRequest) (response *DescribeAccountAllGrantPrivilegesResponse, err error) {
    return c.DescribeAccountAllGrantPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountAllGrantPrivileges
// 账号所有权限
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
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountAllGrantPrivilegesWithContext(ctx context.Context, request *DescribeAccountAllGrantPrivilegesRequest) (response *DescribeAccountAllGrantPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountAllGrantPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountAllGrantPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountAllGrantPrivilegesResponse()
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
// 本接口(DescribeAccounts)用于查询数据库管理账号。
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
// 本接口(DescribeAccounts)用于查询数据库管理账号。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccounts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
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
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAuditLogFiles(request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    return c.DescribeAuditLogFilesWithContext(context.Background(), request)
}

// DescribeAuditLogFiles
// 本接口(DescribeAuditLogFiles)用于查询云数据库实例的审计日志文件。
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
func (c *Client) DescribeAuditLogFilesWithContext(ctx context.Context, request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogFilesRequest()
    }
    
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
// 本接口(DescribeAuditLogs)用于查询数据库审计日志。
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
func (c *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    return c.DescribeAuditLogsWithContext(context.Background(), request)
}

// DescribeAuditLogs
// 本接口(DescribeAuditLogs)用于查询数据库审计日志。
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
func (c *Client) DescribeAuditLogsWithContext(ctx context.Context, request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogsRequest()
    }
    
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
// 查询审计规则模版信息
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
func (c *Client) DescribeAuditRuleTemplates(request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    return c.DescribeAuditRuleTemplatesWithContext(context.Background(), request)
}

// DescribeAuditRuleTemplates
// 查询审计规则模版信息
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
func (c *Client) DescribeAuditRuleTemplatesWithContext(ctx context.Context, request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleTemplatesRequest()
    }
    
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
// 获取实例的审计规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DescribeAuditRuleWithInstanceIds(request *DescribeAuditRuleWithInstanceIdsRequest) (response *DescribeAuditRuleWithInstanceIdsResponse, err error) {
    return c.DescribeAuditRuleWithInstanceIdsWithContext(context.Background(), request)
}

// DescribeAuditRuleWithInstanceIds
// 获取实例的审计规则
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) DescribeAuditRuleWithInstanceIdsWithContext(ctx context.Context, request *DescribeAuditRuleWithInstanceIdsRequest) (response *DescribeAuditRuleWithInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleWithInstanceIdsRequest()
    }
    
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
// 获取指定集群的备份配置信息，包括全量备份时间段，备份文件保留时间
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
// 获取指定集群的备份配置信息，包括全量备份时间段，备份文件保留时间
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupConfigResponse()
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
// 此接口（DescribeBackupDownloadUrl）用于查询集群备份文件下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUrl(request *DescribeBackupDownloadUrlRequest) (response *DescribeBackupDownloadUrlResponse, err error) {
    return c.DescribeBackupDownloadUrlWithContext(context.Background(), request)
}

// DescribeBackupDownloadUrl
// 此接口（DescribeBackupDownloadUrl）用于查询集群备份文件下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupDownloadUrlWithContext(ctx context.Context, request *DescribeBackupDownloadUrlRequest) (response *DescribeBackupDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadUrlResponse()
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
// 查询备份文件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupList(request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    return c.DescribeBackupListWithContext(context.Background(), request)
}

// DescribeBackupList
// 查询备份文件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupListWithContext(ctx context.Context, request *DescribeBackupListRequest) (response *DescribeBackupListResponse, err error) {
    if request == nil {
        request = NewDescribeBackupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupListResponse()
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
// 此接口（DescribeBinlogDownloadUrl）用于查询Binlog的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogDownloadUrl(request *DescribeBinlogDownloadUrlRequest) (response *DescribeBinlogDownloadUrlResponse, err error) {
    return c.DescribeBinlogDownloadUrlWithContext(context.Background(), request)
}

// DescribeBinlogDownloadUrl
// 此接口（DescribeBinlogDownloadUrl）用于查询Binlog的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogDownloadUrlWithContext(ctx context.Context, request *DescribeBinlogDownloadUrlRequest) (response *DescribeBinlogDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogDownloadUrlRequest()
    }
    
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
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
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
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogSaveDaysWithContext(ctx context.Context, request *DescribeBinlogSaveDaysRequest) (response *DescribeBinlogSaveDaysResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogSaveDaysRequest()
    }
    
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
// 此接口（DescribeBinlogs）用来查询集群Binlog日志列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    return c.DescribeBinlogsWithContext(context.Background(), request)
}

// DescribeBinlogs
// 此接口（DescribeBinlogs）用来查询集群Binlog日志列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBinlogsWithContext(ctx context.Context, request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogsResponse()
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
// 该接口（DescribeClusterDetail）显示集群详情
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
// 该接口（DescribeClusterDetail）显示集群详情
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterDetailResponse()
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
// 本接口（DescribeClusterInstanceGrps）用于查询实例组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGrps(request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    return c.DescribeClusterInstanceGrpsWithContext(context.Background(), request)
}

// DescribeClusterInstanceGrps
// 本接口（DescribeClusterInstanceGrps）用于查询实例组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstanceGrpsWithContext(ctx context.Context, request *DescribeClusterInstanceGrpsRequest) (response *DescribeClusterInstanceGrpsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstanceGrpsRequest()
    }
    
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
// 本接口（DescribeClusterParamLogs）查询参数修改日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParamLogs(request *DescribeClusterParamLogsRequest) (response *DescribeClusterParamLogsResponse, err error) {
    return c.DescribeClusterParamLogsWithContext(context.Background(), request)
}

// DescribeClusterParamLogs
// 本接口（DescribeClusterParamLogs）查询参数修改日志
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWNOTFOUNDERROR = "FailedOperation.FlowNotFoundError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_FLOWNOTFOUND = "InvalidParameterValue.FlowNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterParamLogsWithContext(ctx context.Context, request *DescribeClusterParamLogsRequest) (response *DescribeClusterParamLogsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterParamLogsRequest()
    }
    
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
// 本接口（DescribeClusterParams）用于查询集群参数
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
// 本接口（DescribeClusterParams）用于查询集群参数
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterParamsResponse()
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
// 查询集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 查询集群列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_CAMCHECKRESOURCEERROR = "FailedOperation.CamCheckResourceError"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_ILLEGALORDERBY = "InvalidParameterValue.IllegalOrderBy"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMBOTHSETERROR = "InvalidParameterValue.ParamBothSetError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
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
// 查询实例安全组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 查询实例安全组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    
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
// 本接口（DescribeFlow）用于查询任务流信息
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
// 本接口（DescribeFlow）用于查询任务流信息
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDetailResponse()
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
// 此接口（DescribeInstanceSlowQueries）用于查询实例慢查询日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSlowQueries(request *DescribeInstanceSlowQueriesRequest) (response *DescribeInstanceSlowQueriesResponse, err error) {
    return c.DescribeInstanceSlowQueriesWithContext(context.Background(), request)
}

// DescribeInstanceSlowQueries
// 此接口（DescribeInstanceSlowQueries）用于查询实例慢查询日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSlowQueriesWithContext(ctx context.Context, request *DescribeInstanceSlowQueriesRequest) (response *DescribeInstanceSlowQueriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSlowQueriesRequest()
    }
    
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
// 本接口（DescribeInstanceSpecs）用于查询实例规格
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    return c.DescribeInstanceSpecsWithContext(context.Background(), request)
}

// DescribeInstanceSpecs
// 本接口（DescribeInstanceSpecs）用于查询实例规格
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceSpecsWithContext(ctx context.Context, request *DescribeInstanceSpecsRequest) (response *DescribeInstanceSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSpecsRequest()
    }
    
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
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
// 查询实例维护时间窗
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintainPeriod(request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    return c.DescribeMaintainPeriodWithContext(context.Background(), request)
}

// DescribeMaintainPeriod
// 查询实例维护时间窗
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintainPeriodWithContext(ctx context.Context, request *DescribeMaintainPeriodRequest) (response *DescribeMaintainPeriodResponse, err error) {
    if request == nil {
        request = NewDescribeMaintainPeriodRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintainPeriod require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintainPeriodResponse()
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
// 查询用户指定产品下的所有参数模板信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// 查询用户指定产品下的所有参数模板信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeParamTemplatesWithContext(ctx context.Context, request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplatesRequest()
    }
    
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
// 查询项目安全组信息
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
// 查询项目安全组信息
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
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
// 根据计费订单id查询资源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealName(request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    return c.DescribeResourcesByDealNameWithContext(context.Background(), request)
}

// DescribeResourcesByDealName
// 根据计费订单id查询资源列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealNameWithContext(ctx context.Context, request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByDealNameRequest()
    }
    
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
// 查询指定集群有效回滚时间范围
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeRange(request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    return c.DescribeRollbackTimeRangeWithContext(context.Background(), request)
}

// DescribeRollbackTimeRange
// 查询指定集群有效回滚时间范围
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeRangeWithContext(ctx context.Context, request *DescribeRollbackTimeRangeRequest) (response *DescribeRollbackTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTimeRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTimeRangeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeValidityRequest() (request *DescribeRollbackTimeValidityRequest) {
    request = &DescribeRollbackTimeValidityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cynosdb", APIVersion, "DescribeRollbackTimeValidity")
    
    
    return
}

func NewDescribeRollbackTimeValidityResponse() (response *DescribeRollbackTimeValidityResponse) {
    response = &DescribeRollbackTimeValidityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRollbackTimeValidity
// 指定时间和集群查询是否可回滚
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeValidity(request *DescribeRollbackTimeValidityRequest) (response *DescribeRollbackTimeValidityResponse, err error) {
    return c.DescribeRollbackTimeValidityWithContext(context.Background(), request)
}

// DescribeRollbackTimeValidity
// 指定时间和集群查询是否可回滚
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeValidityWithContext(ctx context.Context, request *DescribeRollbackTimeValidityRequest) (response *DescribeRollbackTimeValidityResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeValidityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTimeValidity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTimeValidityResponse()
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
// 本接口(DescribeZones)用于查询可售卖地域可用区信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 本接口(DescribeZones)用于查询可售卖地域可用区信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    if request == nil {
        request = NewDescribeZonesRequest()
    }
    
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
// 安全组批量解绑云资源
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
// 安全组批量解绑云资源
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
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
// 此接口（ExportInstanceSlowQueries）用于导出实例慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceSlowQueries(request *ExportInstanceSlowQueriesRequest) (response *ExportInstanceSlowQueriesResponse, err error) {
    return c.ExportInstanceSlowQueriesWithContext(context.Background(), request)
}

// ExportInstanceSlowQueries
// 此接口（ExportInstanceSlowQueries）用于导出实例慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ExportInstanceSlowQueriesWithContext(ctx context.Context, request *ExportInstanceSlowQueriesRequest) (response *ExportInstanceSlowQueriesResponse, err error) {
    if request == nil {
        request = NewExportInstanceSlowQueriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportInstanceSlowQueries require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportInstanceSlowQueriesResponse()
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
// 批量授权账号权限
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
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    return c.GrantAccountPrivilegesWithContext(context.Background(), request)
}

// GrantAccountPrivileges
// 批量授权账号权限
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
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivilegesWithContext(ctx context.Context, request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewGrantAccountPrivilegesRequest()
    }
    
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
// 查询新购集群价格
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
// 查询新购集群价格
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreate require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateResponse()
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
// 查询续费集群价格
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
// 查询续费集群价格
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
// 隔离集群
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
// 隔离集群
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateInstanceResponse()
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
// 修改账号参数
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
// 修改账号参数
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountParamsResponse()
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
// 修改审计规则模版
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
func (c *Client) ModifyAuditRuleTemplates(request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    return c.ModifyAuditRuleTemplatesWithContext(context.Background(), request)
}

// ModifyAuditRuleTemplates
// 修改审计规则模版
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
func (c *Client) ModifyAuditRuleTemplatesWithContext(ctx context.Context, request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewModifyAuditRuleTemplatesRequest()
    }
    
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
func (c *Client) ModifyAuditService(request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    return c.ModifyAuditServiceWithContext(context.Background(), request)
}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计日志保存时长、审计规则等服务配置。
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
func (c *Client) ModifyAuditServiceWithContext(ctx context.Context, request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    if request == nil {
        request = NewModifyAuditServiceRequest()
    }
    
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
// 修改指定集群的备份配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    return c.ModifyBackupConfigWithContext(context.Background(), request)
}

// ModifyBackupConfig
// 修改指定集群的备份配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupConfigWithContext(ctx context.Context, request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupConfigResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupNameResponse()
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
// 修改集群名称
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
// 修改集群名称
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
// 修改集群参数
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
// 修改集群参数
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyClusterParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyClusterParamResponse()
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
// 修改从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterSlaveZone(request *ModifyClusterSlaveZoneRequest) (response *ModifyClusterSlaveZoneResponse, err error) {
    return c.ModifyClusterSlaveZoneWithContext(context.Background(), request)
}

// ModifyClusterSlaveZone
// 修改从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterSlaveZoneWithContext(ctx context.Context, request *ModifyClusterSlaveZoneRequest) (response *ModifyClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewModifyClusterSlaveZoneRequest()
    }
    
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
// 升级预付费存储
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
// 升级预付费存储
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
// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
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
// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
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
// 修改维护时间配置
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
// 修改维护时间配置
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintainPeriodConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintainPeriodConfigResponse()
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
// 修改实例组ip，端口
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
// 修改实例组ip，端口
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
// 下线集群
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
// 下线集群
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
// 下线实例
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
// 下线实例
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
// TDSQL-C for MySQL实例开通审计服务
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// TDSQL-C for MySQL实例开通审计服务
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAuditServiceResponse()
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
// 暂停serverless集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) PauseServerless(request *PauseServerlessRequest) (response *PauseServerlessResponse, err error) {
    return c.PauseServerlessWithContext(context.Background(), request)
}

// PauseServerless
// 暂停serverless集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWCREATEERROR = "FailedOperation.FlowCreateError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_DBMODENOTSERVERLESSERROR = "InvalidParameterValue.DbModeNotServerlessError"
//  INVALIDPARAMETERVALUE_PARAMERROR = "InvalidParameterValue.ParamError"
//  OPERATIONDENIED_SERVERLESSCLUSTERSTATUSDENIED = "OperationDenied.ServerlessClusterStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) PauseServerlessWithContext(ctx context.Context, request *PauseServerlessRequest) (response *PauseServerlessResponse, err error) {
    if request == nil {
        request = NewPauseServerlessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseServerless require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseServerlessResponse()
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
// 删除从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveClusterSlaveZone(request *RemoveClusterSlaveZoneRequest) (response *RemoveClusterSlaveZoneResponse, err error) {
    return c.RemoveClusterSlaveZoneWithContext(context.Background(), request)
}

// RemoveClusterSlaveZone
// 删除从可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  OPERATIONDENIED_CLUSTERSTATUSDENIEDERROR = "OperationDenied.ClusterStatusDeniedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveClusterSlaveZoneWithContext(ctx context.Context, request *RemoveClusterSlaveZoneRequest) (response *RemoveClusterSlaveZoneResponse, err error) {
    if request == nil {
        request = NewRemoveClusterSlaveZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveClusterSlaveZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveClusterSlaveZoneResponse()
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
// 本接口(ResetAccountPassword)用于重置实例的数据库账号密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
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
// 本接口(ResetAccountPassword)用于重置实例的数据库账号密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_OPERATIONNOTSUPPORT = "InternalError.OperationNotSupport"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
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
// 重启实例
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
// 重启实例
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
// 恢复serverless集群
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
// 恢复serverless集群
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
// 批量回收账号权限
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
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RevokeAccountPrivileges(request *RevokeAccountPrivilegesRequest) (response *RevokeAccountPrivilegesResponse, err error) {
    return c.RevokeAccountPrivilegesWithContext(context.Background(), request)
}

// RevokeAccountPrivileges
// 批量回收账号权限
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
//  OPERATIONDENIED_SERVERLESSINSTANCESTATUSDENIED = "OperationDenied.ServerlessInstanceStatusDenied"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKFAIL = "ResourceUnavailable.InstanceLockFail"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RevokeAccountPrivilegesWithContext(ctx context.Context, request *RevokeAccountPrivilegesRequest) (response *RevokeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewRevokeAccountPrivilegesRequest()
    }
    
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
// 本接口（RollBackCluster）用于回档集群
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
// 本接口（RollBackCluster）用于回档集群
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollBackCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollBackClusterResponse()
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
// 本接口(SearchClusterDatabases)搜索集群database列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterDatabases(request *SearchClusterDatabasesRequest) (response *SearchClusterDatabasesResponse, err error) {
    return c.SearchClusterDatabasesWithContext(context.Background(), request)
}

// SearchClusterDatabases
// 本接口(SearchClusterDatabases)搜索集群database列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_OPERATIONFAILEDERROR = "FailedOperation.OperationFailedError"
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
func (c *Client) SearchClusterDatabasesWithContext(ctx context.Context, request *SearchClusterDatabasesRequest) (response *SearchClusterDatabasesResponse, err error) {
    if request == nil {
        request = NewSearchClusterDatabasesRequest()
    }
    
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
// 本接口(SearchClusterTables)搜索集群数据表列表
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
// 本接口(SearchClusterTables)搜索集群数据表列表
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
// SetRenewFlag设置实例的自动续费功能
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
// SetRenewFlag设置实例的自动续费功能
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetRenewFlagResponse()
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
// 更换集群vpc
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
// 更换集群vpc
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
// 切换到从可用区
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
// 切换到从可用区
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
// 本接口(SwitchProxyVpc)更换数据库代理vpc
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
// 本接口(SwitchProxyVpc)更换数据库代理vpc
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchProxyVpc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchProxyVpcResponse()
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
// 升级实例
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
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 升级实例
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
//  RESOURCENOTFOUND_CLUSTERNOTFOUNDERROR = "ResourceNotFound.ClusterNotFoundError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}
