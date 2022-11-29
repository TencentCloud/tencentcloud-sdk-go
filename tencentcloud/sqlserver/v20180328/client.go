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

package v20180328

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-28"

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


func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// 本接口(AssociateSecurityGroups)用于安全组批量绑定实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// 本接口(AssociateSecurityGroups)用于安全组批量绑定实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewCloneDBRequest() (request *CloneDBRequest) {
    request = &CloneDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CloneDB")
    
    
    return
}

func NewCloneDBResponse() (response *CloneDBResponse) {
    response = &CloneDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloneDB
// 本接口（CloneDB）用于克隆数据库，只支持克隆到本实例，克隆时必须指定新库名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloneDB(request *CloneDBRequest) (response *CloneDBResponse, err error) {
    return c.CloneDBWithContext(context.Background(), request)
}

// CloneDB
// 本接口（CloneDB）用于克隆数据库，只支持克隆到本实例，克隆时必须指定新库名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloneDBWithContext(ctx context.Context, request *CloneDBRequest) (response *CloneDBResponse, err error) {
    if request == nil {
        request = NewCloneDBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneDB require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneDBResponse()
    err = c.Send(request, response)
    return
}

func NewCloseInterCommunicationRequest() (request *CloseInterCommunicationRequest) {
    request = &CloseInterCommunicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CloseInterCommunication")
    
    
    return
}

func NewCloseInterCommunicationResponse() (response *CloseInterCommunicationResponse) {
    response = &CloseInterCommunicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseInterCommunication
// 本接口（CloseInterCommunication）用于关闭实例互通。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseInterCommunication(request *CloseInterCommunicationRequest) (response *CloseInterCommunicationResponse, err error) {
    return c.CloseInterCommunicationWithContext(context.Background(), request)
}

// CloseInterCommunication
// 本接口（CloseInterCommunication）用于关闭实例互通。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseInterCommunicationWithContext(ctx context.Context, request *CloseInterCommunicationRequest) (response *CloseInterCommunicationResponse, err error) {
    if request == nil {
        request = NewCloseInterCommunicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseInterCommunication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseInterCommunicationResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteExpansionRequest() (request *CompleteExpansionRequest) {
    request = &CompleteExpansionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CompleteExpansion")
    
    
    return
}

func NewCompleteExpansionResponse() (response *CompleteExpansionResponse) {
    response = &CompleteExpansionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CompleteExpansion
// 本接口（CompleteExpansion）在实例发起扩容后，实例状态处于“升级待切换”时，可立即完成实例升级切换操作，无需等待可维护时间窗。本接口需要在实例低峰时调用，在完全切换成功前，存在部分库不可访问的风险。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CompleteExpansion(request *CompleteExpansionRequest) (response *CompleteExpansionResponse, err error) {
    return c.CompleteExpansionWithContext(context.Background(), request)
}

// CompleteExpansion
// 本接口（CompleteExpansion）在实例发起扩容后，实例状态处于“升级待切换”时，可立即完成实例升级切换操作，无需等待可维护时间窗。本接口需要在实例低峰时调用，在完全切换成功前，存在部分库不可访问的风险。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CompleteExpansionWithContext(ctx context.Context, request *CompleteExpansionRequest) (response *CompleteExpansionResponse, err error) {
    if request == nil {
        request = NewCompleteExpansionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteExpansion require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteExpansionResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteMigrationRequest() (request *CompleteMigrationRequest) {
    request = &CompleteMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CompleteMigration")
    
    
    return
}

func NewCompleteMigrationResponse() (response *CompleteMigrationResponse) {
    response = &CompleteMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CompleteMigration
// 本接口（CompleteMigration）作用是完成一个迁移任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CompleteMigration(request *CompleteMigrationRequest) (response *CompleteMigrationResponse, err error) {
    return c.CompleteMigrationWithContext(context.Background(), request)
}

// CompleteMigration
// 本接口（CompleteMigration）作用是完成一个迁移任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CompleteMigrationWithContext(ctx context.Context, request *CompleteMigrationRequest) (response *CompleteMigrationResponse, err error) {
    if request == nil {
        request = NewCompleteMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// 本接口（CreateAccount）用于创建实例账号
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMEISILLEGAL = "InvalidParameterValue.AccountNameIsIllegal"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMEISKEYWORDS = "InvalidParameterValue.AccountNameIsKeyWords"
//  INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// 本接口（CreateAccount）用于创建实例账号
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTEXIST = "InvalidParameterValue.AccountExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMEISILLEGAL = "InvalidParameterValue.AccountNameIsIllegal"
//  INVALIDPARAMETERVALUE_ACCOUNTNAMEISKEYWORDS = "InvalidParameterValue.AccountNameIsKeyWords"
//  INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    if request == nil {
        request = NewCreateAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBackup")
    
    
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBackup
// 本接口(CreateBackup)用于创建备份。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// 本接口(CreateBackup)用于创建备份。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
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

func NewCreateBackupMigrationRequest() (request *CreateBackupMigrationRequest) {
    request = &CreateBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBackupMigration")
    
    
    return
}

func NewCreateBackupMigrationResponse() (response *CreateBackupMigrationResponse) {
    response = &CreateBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBackupMigration
// 本接口（CreateBackupMigration）用于创建备份导入任务。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackupMigration(request *CreateBackupMigrationRequest) (response *CreateBackupMigrationResponse, err error) {
    return c.CreateBackupMigrationWithContext(context.Background(), request)
}

// CreateBackupMigration
// 本接口（CreateBackupMigration）用于创建备份导入任务。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBackupMigrationWithContext(ctx context.Context, request *CreateBackupMigrationRequest) (response *CreateBackupMigrationResponse, err error) {
    if request == nil {
        request = NewCreateBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBasicDBInstancesRequest() (request *CreateBasicDBInstancesRequest) {
    request = &CreateBasicDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBasicDBInstances")
    
    
    return
}

func NewCreateBasicDBInstancesResponse() (response *CreateBasicDBInstancesResponse) {
    response = &CreateBasicDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBasicDBInstances
// 本接口（CreateBasicDBInstances）用于创建SQL server基础版实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBasicDBInstances(request *CreateBasicDBInstancesRequest) (response *CreateBasicDBInstancesResponse, err error) {
    return c.CreateBasicDBInstancesWithContext(context.Background(), request)
}

// CreateBasicDBInstances
// 本接口（CreateBasicDBInstances）用于创建SQL server基础版实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBasicDBInstancesWithContext(ctx context.Context, request *CreateBasicDBInstancesRequest) (response *CreateBasicDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateBasicDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBasicDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBasicDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBusinessDBInstancesRequest() (request *CreateBusinessDBInstancesRequest) {
    request = &CreateBusinessDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBusinessDBInstances")
    
    
    return
}

func NewCreateBusinessDBInstancesResponse() (response *CreateBusinessDBInstancesResponse) {
    response = &CreateBusinessDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBusinessDBInstances
// 本接口（CreateBusinessDBInstances）用于创建商业智能服务实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessDBInstances(request *CreateBusinessDBInstancesRequest) (response *CreateBusinessDBInstancesResponse, err error) {
    return c.CreateBusinessDBInstancesWithContext(context.Background(), request)
}

// CreateBusinessDBInstances
// 本接口（CreateBusinessDBInstances）用于创建商业智能服务实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessDBInstancesWithContext(ctx context.Context, request *CreateBusinessDBInstancesRequest) (response *CreateBusinessDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateBusinessDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBusinessDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBusinessDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBusinessIntelligenceFileRequest() (request *CreateBusinessIntelligenceFileRequest) {
    request = &CreateBusinessIntelligenceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateBusinessIntelligenceFile")
    
    
    return
}

func NewCreateBusinessIntelligenceFileResponse() (response *CreateBusinessIntelligenceFileResponse) {
    response = &CreateBusinessIntelligenceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBusinessIntelligenceFile
// 本接口（CreateBusinessIntelligenceFile）用于添加商业智能服务文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessIntelligenceFile(request *CreateBusinessIntelligenceFileRequest) (response *CreateBusinessIntelligenceFileResponse, err error) {
    return c.CreateBusinessIntelligenceFileWithContext(context.Background(), request)
}

// CreateBusinessIntelligenceFile
// 本接口（CreateBusinessIntelligenceFile）用于添加商业智能服务文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateBusinessIntelligenceFileWithContext(ctx context.Context, request *CreateBusinessIntelligenceFileRequest) (response *CreateBusinessIntelligenceFileResponse, err error) {
    if request == nil {
        request = NewCreateBusinessIntelligenceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBusinessIntelligenceFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBRequest() (request *CreateDBRequest) {
    request = &CreateDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDB")
    
    
    return
}

func NewCreateDBResponse() (response *CreateDBResponse) {
    response = &CreateDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDB
// 本接口（CreateDB）用于创建数据库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_CHARSETISILLEGAL = "InvalidParameterValue.CharsetIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PRIVILEGEISILLEGAL = "InvalidParameterValue.PrivilegeIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDB(request *CreateDBRequest) (response *CreateDBResponse, err error) {
    return c.CreateDBWithContext(context.Background(), request)
}

// CreateDB
// 本接口（CreateDB）用于创建数据库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_CHARSETISILLEGAL = "InvalidParameterValue.CharsetIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  INVALIDPARAMETERVALUE_PRIVILEGEISILLEGAL = "InvalidParameterValue.PrivilegeIsIllegal"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDBWithContext(ctx context.Context, request *CreateDBRequest) (response *CreateDBResponse, err error) {
    if request == nil {
        request = NewCreateDBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDB require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstancesRequest() (request *CreateDBInstancesRequest) {
    request = &CreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateDBInstances")
    
    
    return
}

func NewCreateDBInstancesResponse() (response *CreateDBInstancesResponse) {
    response = &CreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDBInstances
// 本接口（CreateDBInstances）用于创建实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    return c.CreateDBInstancesWithContext(context.Background(), request)
}

// CreateDBInstances
// 本接口（CreateDBInstances）用于创建实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDBInstancesWithContext(ctx context.Context, request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIncrementalMigrationRequest() (request *CreateIncrementalMigrationRequest) {
    request = &CreateIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateIncrementalMigration")
    
    
    return
}

func NewCreateIncrementalMigrationResponse() (response *CreateIncrementalMigrationResponse) {
    response = &CreateIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIncrementalMigration
// 本接口（CreateIncrementalMigration）用于创建增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCEINUSE_INCREMENTALMIGRATIONEXIST = "ResourceInUse.IncrementalMigrationExist"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BACKUPMIGRATIONRECOVERYTYPEERR = "ResourceUnavailable.BackupMigrationRecoveryTypeErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIncrementalMigration(request *CreateIncrementalMigrationRequest) (response *CreateIncrementalMigrationResponse, err error) {
    return c.CreateIncrementalMigrationWithContext(context.Background(), request)
}

// CreateIncrementalMigration
// 本接口（CreateIncrementalMigration）用于创建增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCEINUSE_INCREMENTALMIGRATIONEXIST = "ResourceInUse.IncrementalMigrationExist"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_BACKUPMIGRATIONRECOVERYTYPEERR = "ResourceUnavailable.BackupMigrationRecoveryTypeErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateIncrementalMigrationWithContext(ctx context.Context, request *CreateIncrementalMigrationRequest) (response *CreateIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewCreateIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrationRequest() (request *CreateMigrationRequest) {
    request = &CreateMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateMigration")
    
    
    return
}

func NewCreateMigrationResponse() (response *CreateMigrationResponse) {
    response = &CreateMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMigration
// 本接口（CreateMigration）作用是创建一个迁移任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  INVALIDPARAMETERVALUE_ONCVMTYPENOTSUPPORTED = "InvalidParameterValue.OnCvmTypeNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateMigration(request *CreateMigrationRequest) (response *CreateMigrationResponse, err error) {
    return c.CreateMigrationWithContext(context.Background(), request)
}

// CreateMigration
// 本接口（CreateMigration）作用是创建一个迁移任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  INVALIDPARAMETERVALUE_ONCVMTYPENOTSUPPORTED = "InvalidParameterValue.OnCvmTypeNotSupported"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateMigrationWithContext(ctx context.Context, request *CreateMigrationRequest) (response *CreateMigrationResponse, err error) {
    if request == nil {
        request = NewCreateMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePublishSubscribeRequest() (request *CreatePublishSubscribeRequest) {
    request = &CreatePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreatePublishSubscribe")
    
    
    return
}

func NewCreatePublishSubscribeResponse() (response *CreatePublishSubscribeResponse) {
    response = &CreatePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePublishSubscribe
// 本接口（CreatePublishSubscribe）用于创建两个数据库之间的发布订阅关系。作为订阅者，不能再充当发布者，作为发布者可以有多个订阅者实例。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreatePublishSubscribe(request *CreatePublishSubscribeRequest) (response *CreatePublishSubscribeResponse, err error) {
    return c.CreatePublishSubscribeWithContext(context.Background(), request)
}

// CreatePublishSubscribe
// 本接口（CreatePublishSubscribe）用于创建两个数据库之间的发布订阅关系。作为订阅者，不能再充当发布者，作为发布者可以有多个订阅者实例。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreatePublishSubscribeWithContext(ctx context.Context, request *CreatePublishSubscribeRequest) (response *CreatePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewCreatePublishSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReadOnlyDBInstancesRequest() (request *CreateReadOnlyDBInstancesRequest) {
    request = &CreateReadOnlyDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "CreateReadOnlyDBInstances")
    
    
    return
}

func NewCreateReadOnlyDBInstancesResponse() (response *CreateReadOnlyDBInstancesResponse) {
    response = &CreateReadOnlyDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateReadOnlyDBInstances
// 本接口（CreateReadOnlyDBInstances）用于添加只读副本实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_ROGROUPNAMEISILLEGAL = "InvalidParameterValue.RoGroupNameIsIllegal"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateReadOnlyDBInstances(request *CreateReadOnlyDBInstancesRequest) (response *CreateReadOnlyDBInstancesResponse, err error) {
    return c.CreateReadOnlyDBInstancesWithContext(context.Background(), request)
}

// CreateReadOnlyDBInstances
// 本接口（CreateReadOnlyDBInstances）用于添加只读副本实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  INVALIDPARAMETERVALUE_ROGROUPNAMEISILLEGAL = "InvalidParameterValue.RoGroupNameIsIllegal"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateReadOnlyDBInstancesWithContext(ctx context.Context, request *CreateReadOnlyDBInstancesRequest) (response *CreateReadOnlyDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateReadOnlyDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReadOnlyDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReadOnlyDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteAccount")
    
    
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccount
// 本接口（DeleteAccount）用于删除实例账号。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    return c.DeleteAccountWithContext(context.Background(), request)
}

// DeleteAccount
// 本接口（DeleteAccount）用于删除实例账号。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
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

func NewDeleteBackupMigrationRequest() (request *DeleteBackupMigrationRequest) {
    request = &DeleteBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteBackupMigration")
    
    
    return
}

func NewDeleteBackupMigrationResponse() (response *DeleteBackupMigrationResponse) {
    response = &DeleteBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBackupMigration
// 本接口（DeleteBackupMigration）用于删除备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackupMigration(request *DeleteBackupMigrationRequest) (response *DeleteBackupMigrationResponse, err error) {
    return c.DeleteBackupMigrationWithContext(context.Background(), request)
}

// DeleteBackupMigration
// 本接口（DeleteBackupMigration）用于删除备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBackupMigrationWithContext(ctx context.Context, request *DeleteBackupMigrationRequest) (response *DeleteBackupMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBusinessIntelligenceFileRequest() (request *DeleteBusinessIntelligenceFileRequest) {
    request = &DeleteBusinessIntelligenceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteBusinessIntelligenceFile")
    
    
    return
}

func NewDeleteBusinessIntelligenceFileResponse() (response *DeleteBusinessIntelligenceFileResponse) {
    response = &DeleteBusinessIntelligenceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBusinessIntelligenceFile
// 本接口（DeleteBusinessIntelligenceFile）用于删除商业智能文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBusinessIntelligenceFile(request *DeleteBusinessIntelligenceFileRequest) (response *DeleteBusinessIntelligenceFileResponse, err error) {
    return c.DeleteBusinessIntelligenceFileWithContext(context.Background(), request)
}

// DeleteBusinessIntelligenceFile
// 本接口（DeleteBusinessIntelligenceFile）用于删除商业智能文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteBusinessIntelligenceFileWithContext(ctx context.Context, request *DeleteBusinessIntelligenceFileRequest) (response *DeleteBusinessIntelligenceFileResponse, err error) {
    if request == nil {
        request = NewDeleteBusinessIntelligenceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBusinessIntelligenceFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBRequest() (request *DeleteDBRequest) {
    request = &DeleteDBRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDB")
    
    
    return
}

func NewDeleteDBResponse() (response *DeleteDBResponse) {
    response = &DeleteDBResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDB
// 本接口(DeleteDB)用于删除数据库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDB(request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
    return c.DeleteDBWithContext(context.Background(), request)
}

// DeleteDB
// 本接口(DeleteDB)用于删除数据库。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDBWithContext(ctx context.Context, request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
    if request == nil {
        request = NewDeleteDBRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDB require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBInstanceRequest() (request *DeleteDBInstanceRequest) {
    request = &DeleteDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteDBInstance")
    
    
    return
}

func NewDeleteDBInstanceResponse() (response *DeleteDBInstanceResponse) {
    response = &DeleteDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDBInstance
// 本接口（DeleteDBInstance）用于释放回收站中的SQL server实例(立即下线)。释放后的实例将保存一段时间后物理销毁。其发布订阅将自动解除，其ro副本将自动释放。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (response *DeleteDBInstanceResponse, err error) {
    return c.DeleteDBInstanceWithContext(context.Background(), request)
}

// DeleteDBInstance
// 本接口（DeleteDBInstance）用于释放回收站中的SQL server实例(立即下线)。释放后的实例将保存一段时间后物理销毁。其发布订阅将自动解除，其ro副本将自动释放。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteDBInstanceWithContext(ctx context.Context, request *DeleteDBInstanceRequest) (response *DeleteDBInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIncrementalMigrationRequest() (request *DeleteIncrementalMigrationRequest) {
    request = &DeleteIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteIncrementalMigration")
    
    
    return
}

func NewDeleteIncrementalMigrationResponse() (response *DeleteIncrementalMigrationResponse) {
    response = &DeleteIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIncrementalMigration
// 本接口（DeleteIncrementalMigration）用于删除增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteIncrementalMigration(request *DeleteIncrementalMigrationRequest) (response *DeleteIncrementalMigrationResponse, err error) {
    return c.DeleteIncrementalMigrationWithContext(context.Background(), request)
}

// DeleteIncrementalMigration
// 本接口（DeleteIncrementalMigration）用于删除增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteIncrementalMigrationWithContext(ctx context.Context, request *DeleteIncrementalMigrationRequest) (response *DeleteIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrationRequest() (request *DeleteMigrationRequest) {
    request = &DeleteMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeleteMigration")
    
    
    return
}

func NewDeleteMigrationResponse() (response *DeleteMigrationResponse) {
    response = &DeleteMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMigration
// 本接口（DeleteMigration）用于删除迁移任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteMigration(request *DeleteMigrationRequest) (response *DeleteMigrationResponse, err error) {
    return c.DeleteMigrationWithContext(context.Background(), request)
}

// DeleteMigration
// 本接口（DeleteMigration）用于删除迁移任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteMigrationWithContext(ctx context.Context, request *DeleteMigrationRequest) (response *DeleteMigrationResponse, err error) {
    if request == nil {
        request = NewDeleteMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePublishSubscribeRequest() (request *DeletePublishSubscribeRequest) {
    request = &DeletePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DeletePublishSubscribe")
    
    
    return
}

func NewDeletePublishSubscribeResponse() (response *DeletePublishSubscribeResponse) {
    response = &DeletePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePublishSubscribe
// 本接口（DeletePublishSubscribe）用于删除两个数据库间的发布订阅关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeletePublishSubscribe(request *DeletePublishSubscribeRequest) (response *DeletePublishSubscribeResponse, err error) {
    return c.DeletePublishSubscribeWithContext(context.Background(), request)
}

// DeletePublishSubscribe
// 本接口（DeletePublishSubscribe）用于删除两个数据库间的发布订阅关系。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeletePublishSubscribeWithContext(ctx context.Context, request *DeletePublishSubscribeRequest) (response *DeletePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewDeletePublishSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于拉取实例账户列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于拉取实例账户列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewDescribeBackupByFlowIdRequest() (request *DescribeBackupByFlowIdRequest) {
    request = &DescribeBackupByFlowIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupByFlowId")
    
    
    return
}

func NewDescribeBackupByFlowIdResponse() (response *DescribeBackupByFlowIdResponse) {
    response = &DescribeBackupByFlowIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupByFlowId
// 本接口(DescribeBackupByFlowId)用于通过备份创建流程的ID查询创建的备份详情，流程ID可从接口CreateBackup中获得。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupByFlowId(request *DescribeBackupByFlowIdRequest) (response *DescribeBackupByFlowIdResponse, err error) {
    return c.DescribeBackupByFlowIdWithContext(context.Background(), request)
}

// DescribeBackupByFlowId
// 本接口(DescribeBackupByFlowId)用于通过备份创建流程的ID查询创建的备份详情，流程ID可从接口CreateBackup中获得。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupByFlowIdWithContext(ctx context.Context, request *DescribeBackupByFlowIdRequest) (response *DescribeBackupByFlowIdResponse, err error) {
    if request == nil {
        request = NewDescribeBackupByFlowIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupByFlowId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupByFlowIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupCommandRequest() (request *DescribeBackupCommandRequest) {
    request = &DescribeBackupCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupCommand")
    
    
    return
}

func NewDescribeBackupCommandResponse() (response *DescribeBackupCommandResponse) {
    response = &DescribeBackupCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupCommand
// 本接口（DescribeBackupCommand）用于查询以规范的格式创建备份的命令。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupCommand(request *DescribeBackupCommandRequest) (response *DescribeBackupCommandResponse, err error) {
    return c.DescribeBackupCommandWithContext(context.Background(), request)
}

// DescribeBackupCommand
// 本接口（DescribeBackupCommand）用于查询以规范的格式创建备份的命令。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupCommandWithContext(ctx context.Context, request *DescribeBackupCommandRequest) (response *DescribeBackupCommandResponse, err error) {
    if request == nil {
        request = NewDescribeBackupCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupCommandResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupFilesRequest() (request *DescribeBackupFilesRequest) {
    request = &DescribeBackupFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupFiles")
    
    
    return
}

func NewDescribeBackupFilesResponse() (response *DescribeBackupFilesResponse) {
    response = &DescribeBackupFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupFiles
// 本接口(DescribeBackupFiles)用于在非打包备份模式下单个库对应的备份文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupFiles(request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
    return c.DescribeBackupFilesWithContext(context.Background(), request)
}

// DescribeBackupFiles
// 本接口(DescribeBackupFiles)用于在非打包备份模式下单个库对应的备份文件
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupFilesWithContext(ctx context.Context, request *DescribeBackupFilesRequest) (response *DescribeBackupFilesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupMigrationRequest() (request *DescribeBackupMigrationRequest) {
    request = &DescribeBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupMigration")
    
    
    return
}

func NewDescribeBackupMigrationResponse() (response *DescribeBackupMigrationResponse) {
    response = &DescribeBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupMigration
// 本接口（DescribeBackupMigration）用于创建增量备份导入任务。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupMigration(request *DescribeBackupMigrationRequest) (response *DescribeBackupMigrationResponse, err error) {
    return c.DescribeBackupMigrationWithContext(context.Background(), request)
}

// DescribeBackupMigration
// 本接口（DescribeBackupMigration）用于创建增量备份导入任务。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupMigrationWithContext(ctx context.Context, request *DescribeBackupMigrationRequest) (response *DescribeBackupMigrationResponse, err error) {
    if request == nil {
        request = NewDescribeBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupUploadSizeRequest() (request *DescribeBackupUploadSizeRequest) {
    request = &DescribeBackupUploadSizeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackupUploadSize")
    
    
    return
}

func NewDescribeBackupUploadSizeResponse() (response *DescribeBackupUploadSizeResponse) {
    response = &DescribeBackupUploadSizeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupUploadSize
// 本接口（DescribeBackupUploadSize）用于查询上传的备份文件大小。在备份上传类型是COS_UPLOAD(备份放在业务的对象存储上)时有效。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupUploadSize(request *DescribeBackupUploadSizeRequest) (response *DescribeBackupUploadSizeResponse, err error) {
    return c.DescribeBackupUploadSizeWithContext(context.Background(), request)
}

// DescribeBackupUploadSize
// 本接口（DescribeBackupUploadSize）用于查询上传的备份文件大小。在备份上传类型是COS_UPLOAD(备份放在业务的对象存储上)时有效。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INCREBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.IncreBackupMigrationNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupUploadSizeWithContext(ctx context.Context, request *DescribeBackupUploadSizeRequest) (response *DescribeBackupUploadSizeResponse, err error) {
    if request == nil {
        request = NewDescribeBackupUploadSizeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupUploadSize require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupUploadSizeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
    request = &DescribeBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBackups")
    
    
    return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
    response = &DescribeBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackups
// 本接口(DescribeBackups)用于查询备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSERROR = "FailedOperation.CosError"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    return c.DescribeBackupsWithContext(context.Background(), request)
}

// DescribeBackups
// 本接口(DescribeBackups)用于查询备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSERROR = "FailedOperation.CosError"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBackupsWithContext(ctx context.Context, request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBusinessIntelligenceFileRequest() (request *DescribeBusinessIntelligenceFileRequest) {
    request = &DescribeBusinessIntelligenceFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeBusinessIntelligenceFile")
    
    
    return
}

func NewDescribeBusinessIntelligenceFileResponse() (response *DescribeBusinessIntelligenceFileResponse) {
    response = &DescribeBusinessIntelligenceFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBusinessIntelligenceFile
// 本接口（DescribeBusinessIntelligenceFile）用于查询商业智能服务需要的文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBusinessIntelligenceFile(request *DescribeBusinessIntelligenceFileRequest) (response *DescribeBusinessIntelligenceFileResponse, err error) {
    return c.DescribeBusinessIntelligenceFileWithContext(context.Background(), request)
}

// DescribeBusinessIntelligenceFile
// 本接口（DescribeBusinessIntelligenceFile）用于查询商业智能服务需要的文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeBusinessIntelligenceFileWithContext(ctx context.Context, request *DescribeBusinessIntelligenceFileRequest) (response *DescribeBusinessIntelligenceFileResponse, err error) {
    if request == nil {
        request = NewDescribeBusinessIntelligenceFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBusinessIntelligenceFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBusinessIntelligenceFileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCrossRegionZoneRequest() (request *DescribeCrossRegionZoneRequest) {
    request = &DescribeCrossRegionZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeCrossRegionZone")
    
    
    return
}

func NewDescribeCrossRegionZoneResponse() (response *DescribeCrossRegionZoneResponse) {
    response = &DescribeCrossRegionZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCrossRegionZone
// 本接口(DescribeCrossRegionZone)根据主实例查询备机的容灾地域和可用区。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossRegionZone(request *DescribeCrossRegionZoneRequest) (response *DescribeCrossRegionZoneResponse, err error) {
    return c.DescribeCrossRegionZoneWithContext(context.Background(), request)
}

// DescribeCrossRegionZone
// 本接口(DescribeCrossRegionZone)根据主实例查询备机的容灾地域和可用区。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeCrossRegionZoneWithContext(ctx context.Context, request *DescribeCrossRegionZoneRequest) (response *DescribeCrossRegionZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCrossRegionZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCrossRegionZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCrossRegionZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBCharsetsRequest() (request *DescribeDBCharsetsRequest) {
    request = &DescribeDBCharsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBCharsets")
    
    
    return
}

func NewDescribeDBCharsetsResponse() (response *DescribeDBCharsetsResponse) {
    response = &DescribeDBCharsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBCharsets
// 本接口（DescribeDBCharsets）用于查询实例支持的数据库字符集。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDBCharsets(request *DescribeDBCharsetsRequest) (response *DescribeDBCharsetsResponse, err error) {
    return c.DescribeDBCharsetsWithContext(context.Background(), request)
}

// DescribeDBCharsets
// 本接口（DescribeDBCharsets）用于查询实例支持的数据库字符集。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDBCharsetsWithContext(ctx context.Context, request *DescribeDBCharsetsRequest) (response *DescribeDBCharsetsResponse, err error) {
    if request == nil {
        request = NewDescribeDBCharsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBCharsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBCharsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceInterRequest() (request *DescribeDBInstanceInterRequest) {
    request = &DescribeDBInstanceInterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstanceInter")
    
    
    return
}

func NewDescribeDBInstanceInterResponse() (response *DescribeDBInstanceInterResponse) {
    response = &DescribeDBInstanceInterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceInter
// 本接口（DescribeDBInstanceInter）用于查询互通实例的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstanceInter(request *DescribeDBInstanceInterRequest) (response *DescribeDBInstanceInterResponse, err error) {
    return c.DescribeDBInstanceInterWithContext(context.Background(), request)
}

// DescribeDBInstanceInter
// 本接口（DescribeDBInstanceInter）用于查询互通实例的信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstanceInterWithContext(ctx context.Context, request *DescribeDBInstanceInterRequest) (response *DescribeDBInstanceInterResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceInterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceInter require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceInterResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstances
// 本接口(DescribeDBInstances)用于查询实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口(DescribeDBInstances)用于查询实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSecurityGroups
// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewDescribeDBsRequest() (request *DescribeDBsRequest) {
    request = &DescribeDBsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBs")
    
    
    return
}

func NewDescribeDBsResponse() (response *DescribeDBsResponse) {
    response = &DescribeDBsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBs
// 本接口（DescribeDBs）用于查询数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBs(request *DescribeDBsRequest) (response *DescribeDBsResponse, err error) {
    return c.DescribeDBsWithContext(context.Background(), request)
}

// DescribeDBs
// 本接口（DescribeDBs）用于查询数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBsWithContext(ctx context.Context, request *DescribeDBsRequest) (response *DescribeDBsResponse, err error) {
    if request == nil {
        request = NewDescribeDBsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBsNormalRequest() (request *DescribeDBsNormalRequest) {
    request = &DescribeDBsNormalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeDBsNormal")
    
    
    return
}

func NewDescribeDBsNormalResponse() (response *DescribeDBsNormalResponse) {
    response = &DescribeDBsNormalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBsNormal
// 本接口(DescribeDBsNormal)用于查询数据库配置信息，此接口不包含数据库的关联账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBsNormal(request *DescribeDBsNormalRequest) (response *DescribeDBsNormalResponse, err error) {
    return c.DescribeDBsNormalWithContext(context.Background(), request)
}

// DescribeDBsNormal
// 本接口(DescribeDBsNormal)用于查询数据库配置信息，此接口不包含数据库的关联账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INTERFACENAMENOTFOUND = "InvalidParameter.InterfaceNameNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBsNormalWithContext(ctx context.Context, request *DescribeDBsNormalRequest) (response *DescribeDBsNormalResponse, err error) {
    if request == nil {
        request = NewDescribeDBsNormalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBsNormal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBsNormalResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowStatusRequest() (request *DescribeFlowStatusRequest) {
    request = &DescribeFlowStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeFlowStatus")
    
    
    return
}

func NewDescribeFlowStatusResponse() (response *DescribeFlowStatusResponse) {
    response = &DescribeFlowStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlowStatus
// 本接口(DescribeFlowStatus)用于查询流程状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlowStatus(request *DescribeFlowStatusRequest) (response *DescribeFlowStatusResponse, err error) {
    return c.DescribeFlowStatusWithContext(context.Background(), request)
}

// DescribeFlowStatus
// 本接口(DescribeFlowStatus)用于查询流程状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlowStatusWithContext(ctx context.Context, request *DescribeFlowStatusRequest) (response *DescribeFlowStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFlowStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIncrementalMigrationRequest() (request *DescribeIncrementalMigrationRequest) {
    request = &DescribeIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeIncrementalMigration")
    
    
    return
}

func NewDescribeIncrementalMigrationResponse() (response *DescribeIncrementalMigrationResponse) {
    response = &DescribeIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIncrementalMigration
// 本接口（DescribeIncrementalMigration）用于查询增量备份导入任务。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIncrementalMigration(request *DescribeIncrementalMigrationRequest) (response *DescribeIncrementalMigrationResponse, err error) {
    return c.DescribeIncrementalMigrationWithContext(context.Background(), request)
}

// DescribeIncrementalMigration
// 本接口（DescribeIncrementalMigration）用于查询增量备份导入任务。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeIncrementalMigrationWithContext(ctx context.Context, request *DescribeIncrementalMigrationRequest) (response *DescribeIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewDescribeIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceParamRecords")
    
    
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParamRecords
// 该接口（DescribeInstanceParamRecords）用于查询实例参数修改历史。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// 该接口（DescribeInstanceParamRecords）用于查询实例参数修改历史。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamRecordsWithContext(ctx context.Context, request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParamRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParams
// 该接口（DescribeInstanceParams）用于查询实例的参数列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// 该接口（DescribeInstanceParams）用于查询实例的参数列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PARAMSNOTFOUND = "ResourceNotFound.ParamsNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintenanceSpanRequest() (request *DescribeMaintenanceSpanRequest) {
    request = &DescribeMaintenanceSpanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMaintenanceSpan")
    
    
    return
}

func NewDescribeMaintenanceSpanResponse() (response *DescribeMaintenanceSpanResponse) {
    response = &DescribeMaintenanceSpanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaintenanceSpan
// 本接口（DescribeMaintenanceSpan）根据实例ID查询该实例的可维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintenanceSpan(request *DescribeMaintenanceSpanRequest) (response *DescribeMaintenanceSpanResponse, err error) {
    return c.DescribeMaintenanceSpanWithContext(context.Background(), request)
}

// DescribeMaintenanceSpan
// 本接口（DescribeMaintenanceSpan）根据实例ID查询该实例的可维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMaintenanceSpanWithContext(ctx context.Context, request *DescribeMaintenanceSpanRequest) (response *DescribeMaintenanceSpanResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceSpanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintenanceSpan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintenanceSpanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDatabasesRequest() (request *DescribeMigrationDatabasesRequest) {
    request = &DescribeMigrationDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrationDatabases")
    
    
    return
}

func NewDescribeMigrationDatabasesResponse() (response *DescribeMigrationDatabasesResponse) {
    response = &DescribeMigrationDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrationDatabases
// 本接口（DescribeMigrationDatabases）的作用是查询待迁移数据库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationDatabases(request *DescribeMigrationDatabasesRequest) (response *DescribeMigrationDatabasesResponse, err error) {
    return c.DescribeMigrationDatabasesWithContext(context.Background(), request)
}

// DescribeMigrationDatabases
// 本接口（DescribeMigrationDatabases）的作用是查询待迁移数据库列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationDatabasesWithContext(ctx context.Context, request *DescribeMigrationDatabasesRequest) (response *DescribeMigrationDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDatabasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDetailRequest() (request *DescribeMigrationDetailRequest) {
    request = &DescribeMigrationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrationDetail")
    
    
    return
}

func NewDescribeMigrationDetailResponse() (response *DescribeMigrationDetailResponse) {
    response = &DescribeMigrationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrationDetail
// 本接口（DescribeMigrationDetail）用于查询迁移任务的详细情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationDetail(request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    return c.DescribeMigrationDetailWithContext(context.Background(), request)
}

// DescribeMigrationDetail
// 本接口（DescribeMigrationDetail）用于查询迁移任务的详细情况
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationDetailWithContext(ctx context.Context, request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationsRequest() (request *DescribeMigrationsRequest) {
    request = &DescribeMigrationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeMigrations")
    
    
    return
}

func NewDescribeMigrationsResponse() (response *DescribeMigrationsResponse) {
    response = &DescribeMigrationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrations
// 本接口（DescribeMigrations）根据输入的限定条件，查询符合条件的迁移任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrations(request *DescribeMigrationsRequest) (response *DescribeMigrationsResponse, err error) {
    return c.DescribeMigrationsWithContext(context.Background(), request)
}

// DescribeMigrations
// 本接口（DescribeMigrations）根据输入的限定条件，查询符合条件的迁移任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeMigrationsWithContext(ctx context.Context, request *DescribeMigrationsRequest) (response *DescribeMigrationsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeOrders")
    
    
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrders
// 本接口（DescribeOrders）用于查询订单信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYORDERFAILED = "FailedOperation.QueryOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    return c.DescribeOrdersWithContext(context.Background(), request)
}

// DescribeOrders
// 本接口（DescribeOrders）用于查询订单信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYORDERFAILED = "FailedOperation.QueryOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeOrdersWithContext(ctx context.Context, request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    if request == nil {
        request = NewDescribeOrdersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrders require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrdersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductConfigRequest() (request *DescribeProductConfigRequest) {
    request = &DescribeProductConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeProductConfig")
    
    
    return
}

func NewDescribeProductConfigResponse() (response *DescribeProductConfigResponse) {
    response = &DescribeProductConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductConfig
// 本接口 (DescribeProductConfig) 用于查询售卖规格配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProductConfig(request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    return c.DescribeProductConfigWithContext(context.Background(), request)
}

// DescribeProductConfig
// 本接口 (DescribeProductConfig) 用于查询售卖规格配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProductConfigWithContext(ctx context.Context, request *DescribeProductConfigRequest) (response *DescribeProductConfigResponse, err error) {
    if request == nil {
        request = NewDescribeProductConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeProjectSecurityGroups")
    
    
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectSecurityGroups
// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
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

func NewDescribePublishSubscribeRequest() (request *DescribePublishSubscribeRequest) {
    request = &DescribePublishSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribePublishSubscribe")
    
    
    return
}

func NewDescribePublishSubscribeResponse() (response *DescribePublishSubscribeResponse) {
    response = &DescribePublishSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePublishSubscribe
// 本接口（DescribePublishSubscribe）用于查询发布订阅关系列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribePublishSubscribe(request *DescribePublishSubscribeRequest) (response *DescribePublishSubscribeResponse, err error) {
    return c.DescribePublishSubscribeWithContext(context.Background(), request)
}

// DescribePublishSubscribe
// 本接口（DescribePublishSubscribe）用于查询发布订阅关系列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribePublishSubscribeWithContext(ctx context.Context, request *DescribePublishSubscribeRequest) (response *DescribePublishSubscribeResponse, err error) {
    if request == nil {
        request = NewDescribePublishSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePublishSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePublishSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupByReadOnlyInstanceRequest() (request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) {
    request = &DescribeReadOnlyGroupByReadOnlyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupByReadOnlyInstance")
    
    
    return
}

func NewDescribeReadOnlyGroupByReadOnlyInstanceResponse() (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse) {
    response = &DescribeReadOnlyGroupByReadOnlyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReadOnlyGroupByReadOnlyInstance
// 本接口（DescribeReadOnlyGroupByReadOnlyInstance）用于通过只读副本实例ID查询其所在的只读组。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupByReadOnlyInstance(request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse, err error) {
    return c.DescribeReadOnlyGroupByReadOnlyInstanceWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupByReadOnlyInstance
// 本接口（DescribeReadOnlyGroupByReadOnlyInstance）用于通过只读副本实例ID查询其所在的只读组。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupByReadOnlyInstanceWithContext(ctx context.Context, request *DescribeReadOnlyGroupByReadOnlyInstanceRequest) (response *DescribeReadOnlyGroupByReadOnlyInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupByReadOnlyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupByReadOnlyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupByReadOnlyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupDetailsRequest() (request *DescribeReadOnlyGroupDetailsRequest) {
    request = &DescribeReadOnlyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupDetails")
    
    
    return
}

func NewDescribeReadOnlyGroupDetailsResponse() (response *DescribeReadOnlyGroupDetailsResponse) {
    response = &DescribeReadOnlyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReadOnlyGroupDetails
// 本接口（DescribeReadOnlyGroupDetails）用于查询只读组详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupDetails(request *DescribeReadOnlyGroupDetailsRequest) (response *DescribeReadOnlyGroupDetailsResponse, err error) {
    return c.DescribeReadOnlyGroupDetailsWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupDetails
// 本接口（DescribeReadOnlyGroupDetails）用于查询只读组详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupDetailsWithContext(ctx context.Context, request *DescribeReadOnlyGroupDetailsRequest) (response *DescribeReadOnlyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReadOnlyGroupListRequest() (request *DescribeReadOnlyGroupListRequest) {
    request = &DescribeReadOnlyGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeReadOnlyGroupList")
    
    
    return
}

func NewDescribeReadOnlyGroupListResponse() (response *DescribeReadOnlyGroupListResponse) {
    response = &DescribeReadOnlyGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeReadOnlyGroupList
// 本接口（DescribeReadOnlyGroupList）用于查询只读组列表。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupList(request *DescribeReadOnlyGroupListRequest) (response *DescribeReadOnlyGroupListResponse, err error) {
    return c.DescribeReadOnlyGroupListWithContext(context.Background(), request)
}

// DescribeReadOnlyGroupList
// 本接口（DescribeReadOnlyGroupList）用于查询只读组列表。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeReadOnlyGroupListWithContext(ctx context.Context, request *DescribeReadOnlyGroupListRequest) (response *DescribeReadOnlyGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeReadOnlyGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReadOnlyGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReadOnlyGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
    request = &DescribeRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRegions")
    
    
    return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
    response = &DescribeRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegions
// 本接口 (DescribeRegions) 用于查询售卖地域信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    return c.DescribeRegionsWithContext(context.Background(), request)
}

// DescribeRegions
// 本接口 (DescribeRegions) 用于查询售卖地域信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTimeRequest() (request *DescribeRollbackTimeRequest) {
    request = &DescribeRollbackTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeRollbackTime")
    
    
    return
}

func NewDescribeRollbackTimeResponse() (response *DescribeRollbackTimeResponse) {
    response = &DescribeRollbackTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRollbackTime
// 本接口（DescribeRollbackTime）用于查询实例可回档时间范围
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTime(request *DescribeRollbackTimeRequest) (response *DescribeRollbackTimeResponse, err error) {
    return c.DescribeRollbackTimeWithContext(context.Background(), request)
}

// DescribeRollbackTime
// 本接口（DescribeRollbackTime）用于查询实例可回档时间范围
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeRollbackTimeWithContext(ctx context.Context, request *DescribeRollbackTimeRequest) (response *DescribeRollbackTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowlogsRequest() (request *DescribeSlowlogsRequest) {
    request = &DescribeSlowlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeSlowlogs")
    
    
    return
}

func NewDescribeSlowlogsResponse() (response *DescribeSlowlogsResponse) {
    response = &DescribeSlowlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowlogs
// 本接口（DescribeSlowlogs）用于获取慢查询日志文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSlowlogs(request *DescribeSlowlogsRequest) (response *DescribeSlowlogsResponse, err error) {
    return c.DescribeSlowlogsWithContext(context.Background(), request)
}

// DescribeSlowlogs
// 本接口（DescribeSlowlogs）用于获取慢查询日志文件信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSlowlogsWithContext(ctx context.Context, request *DescribeSlowlogsRequest) (response *DescribeSlowlogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadBackupInfoRequest() (request *DescribeUploadBackupInfoRequest) {
    request = &DescribeUploadBackupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeUploadBackupInfo")
    
    
    return
}

func NewDescribeUploadBackupInfoResponse() (response *DescribeUploadBackupInfoResponse) {
    response = &DescribeUploadBackupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUploadBackupInfo
// 本接口（DescribeUploadBackupInfo）用于查询备份上传权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSPROPERTIESERROR = "FailedOperation.CosPropertiesError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_STSERROR = "InternalError.StsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_UPLOADTYPEERROR = "UnsupportedOperation.UploadTypeError"
func (c *Client) DescribeUploadBackupInfo(request *DescribeUploadBackupInfoRequest) (response *DescribeUploadBackupInfoResponse, err error) {
    return c.DescribeUploadBackupInfoWithContext(context.Background(), request)
}

// DescribeUploadBackupInfo
// 本接口（DescribeUploadBackupInfo）用于查询备份上传权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSPROPERTIESERROR = "FailedOperation.CosPropertiesError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_STSERROR = "InternalError.StsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_UPLOADTYPEERROR = "UnsupportedOperation.UploadTypeError"
func (c *Client) DescribeUploadBackupInfoWithContext(ctx context.Context, request *DescribeUploadBackupInfoRequest) (response *DescribeUploadBackupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadBackupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadBackupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUploadBackupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadIncrementalInfoRequest() (request *DescribeUploadIncrementalInfoRequest) {
    request = &DescribeUploadIncrementalInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeUploadIncrementalInfo")
    
    
    return
}

func NewDescribeUploadIncrementalInfoResponse() (response *DescribeUploadIncrementalInfoResponse) {
    response = &DescribeUploadIncrementalInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUploadIncrementalInfo
// 本接口（DescribeUploadIncrementalInfo）用于查询增量备份上传权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSPROPERTIESERROR = "FailedOperation.CosPropertiesError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_STSERROR = "InternalError.StsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_UPLOADTYPEERROR = "UnsupportedOperation.UploadTypeError"
func (c *Client) DescribeUploadIncrementalInfo(request *DescribeUploadIncrementalInfoRequest) (response *DescribeUploadIncrementalInfoResponse, err error) {
    return c.DescribeUploadIncrementalInfoWithContext(context.Background(), request)
}

// DescribeUploadIncrementalInfo
// 本接口（DescribeUploadIncrementalInfo）用于查询增量备份上传权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSPROPERTIESERROR = "FailedOperation.CosPropertiesError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_STSERROR = "InternalError.StsError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSPATHERROR = "InvalidParameterValue.CosPathError"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_UPLOADTYPEERROR = "UnsupportedOperation.UploadTypeError"
func (c *Client) DescribeUploadIncrementalInfoWithContext(ctx context.Context, request *DescribeUploadIncrementalInfoRequest) (response *DescribeUploadIncrementalInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUploadIncrementalInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadIncrementalInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUploadIncrementalInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
    request = &DescribeZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DescribeZones")
    
    
    return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
    response = &DescribeZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeZones
// 本接口 (DescribeZones) 用于查询当前可售卖的可用区信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
    return c.DescribeZonesWithContext(context.Background(), request)
}

// DescribeZones
// 本接口 (DescribeZones) 用于查询当前可售卖的可用区信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
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
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "DisassociateSecurityGroups")
    
    
    return
}

func NewDisassociateSecurityGroupsResponse() (response *DisassociateSecurityGroupsResponse) {
    response = &DisassociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DisassociateSecurityGroups
// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewInquiryPriceCreateDBInstancesRequest() (request *InquiryPriceCreateDBInstancesRequest) {
    request = &InquiryPriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceCreateDBInstances")
    
    
    return
}

func NewInquiryPriceCreateDBInstancesResponse() (response *InquiryPriceCreateDBInstancesResponse) {
    response = &InquiryPriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateDBInstances
// 本接口（InquiryPriceCreateDBInstances）用于查询申请实例价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BADGOODSNUM = "InvalidParameterValue.BadGoodsNum"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquiryPriceCreateDBInstances(request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    return c.InquiryPriceCreateDBInstancesWithContext(context.Background(), request)
}

// InquiryPriceCreateDBInstances
// 本接口（InquiryPriceCreateDBInstances）用于查询申请实例价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BADGOODSNUM = "InvalidParameterValue.BadGoodsNum"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquiryPriceCreateDBInstancesWithContext(ctx context.Context, request *InquiryPriceCreateDBInstancesRequest) (response *InquiryPriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewDBInstanceRequest() (request *InquiryPriceRenewDBInstanceRequest) {
    request = &InquiryPriceRenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceRenewDBInstance")
    
    
    return
}

func NewInquiryPriceRenewDBInstanceResponse() (response *InquiryPriceRenewDBInstanceResponse) {
    response = &InquiryPriceRenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceRenewDBInstance
// 本接口（InquiryPriceRenewDBInstance）用于查询续费实例的价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSTTYPENOTSUPPORTED = "InvalidParameterValue.CostTypeNotSupported"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquiryPriceRenewDBInstance(request *InquiryPriceRenewDBInstanceRequest) (response *InquiryPriceRenewDBInstanceResponse, err error) {
    return c.InquiryPriceRenewDBInstanceWithContext(context.Background(), request)
}

// InquiryPriceRenewDBInstance
// 本接口（InquiryPriceRenewDBInstance）用于查询续费实例的价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSTTYPENOTSUPPORTED = "InvalidParameterValue.CostTypeNotSupported"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquiryPriceRenewDBInstanceWithContext(ctx context.Context, request *InquiryPriceRenewDBInstanceRequest) (response *InquiryPriceRenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeDBInstanceRequest() (request *InquiryPriceUpgradeDBInstanceRequest) {
    request = &InquiryPriceUpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "InquiryPriceUpgradeDBInstance")
    
    
    return
}

func NewInquiryPriceUpgradeDBInstanceResponse() (response *InquiryPriceUpgradeDBInstanceResponse) {
    response = &InquiryPriceUpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceUpgradeDBInstance
// 本接口（InquiryPriceUpgradeDBInstance）用于查询升级实例的价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSTTYPENOTSUPPORTED = "InvalidParameterValue.CostTypeNotSupported"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCEEXPANDVOLUMELOW = "InvalidParameterValue.InstanceExpandVolumeLow"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquiryPriceUpgradeDBInstance(request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    return c.InquiryPriceUpgradeDBInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpgradeDBInstance
// 本接口（InquiryPriceUpgradeDBInstance）用于查询升级实例的价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_COSTTYPENOTSUPPORTED = "InvalidParameterValue.CostTypeNotSupported"
//  INVALIDPARAMETERVALUE_ILLEGALREGION = "InvalidParameterValue.IllegalRegion"
//  INVALIDPARAMETERVALUE_INSTANCEEXPANDVOLUMELOW = "InvalidParameterValue.InstanceExpandVolumeLow"
//  INVALIDPARAMETERVALUE_PARAMETERTYPEERROR = "InvalidParameterValue.ParameterTypeError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InquiryPriceUpgradeDBInstanceWithContext(ctx context.Context, request *InquiryPriceUpgradeDBInstanceRequest) (response *InquiryPriceUpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegeRequest() (request *ModifyAccountPrivilegeRequest) {
    request = &ModifyAccountPrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyAccountPrivilege")
    
    
    return
}

func NewModifyAccountPrivilegeResponse() (response *ModifyAccountPrivilegeResponse) {
    response = &ModifyAccountPrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountPrivilege
// 本接口（ModifyAccountPrivilege）用于修改实例账户权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilege(request *ModifyAccountPrivilegeRequest) (response *ModifyAccountPrivilegeResponse, err error) {
    return c.ModifyAccountPrivilegeWithContext(context.Background(), request)
}

// ModifyAccountPrivilege
// 本接口（ModifyAccountPrivilege）用于修改实例账户权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ADMINACCOUNTNOTUNIQUE = "InvalidParameterValue.AdminAccountNotUnique"
//  INVALIDPARAMETERVALUE_GRANTISILLEGAL = "InvalidParameterValue.GrantIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_DBNOTEXIT = "ResourceNotFound.DBNotExit"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilegeWithContext(ctx context.Context, request *ModifyAccountPrivilegeRequest) (response *ModifyAccountPrivilegeResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivilege require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPrivilegeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountRemarkRequest() (request *ModifyAccountRemarkRequest) {
    request = &ModifyAccountRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyAccountRemark")
    
    
    return
}

func NewModifyAccountRemarkResponse() (response *ModifyAccountRemarkResponse) {
    response = &ModifyAccountRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountRemark
// 本接口（ModifyAccountRemark）用于修改账户备注。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountRemark(request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    return c.ModifyAccountRemarkWithContext(context.Background(), request)
}

// ModifyAccountRemark
// 本接口（ModifyAccountRemark）用于修改账户备注。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ACCOUNTREMARKISILLEGAL = "InvalidParameterValue.AccountRemarkIsIllegal"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountRemarkWithContext(ctx context.Context, request *ModifyAccountRemarkRequest) (response *ModifyAccountRemarkResponse, err error) {
    if request == nil {
        request = NewModifyAccountRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupMigrationRequest() (request *ModifyBackupMigrationRequest) {
    request = &ModifyBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupMigration")
    
    
    return
}

func NewModifyBackupMigrationResponse() (response *ModifyBackupMigrationResponse) {
    response = &ModifyBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupMigration
// 本接口（ModifyBackupMigration）用于修改备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupMigration(request *ModifyBackupMigrationRequest) (response *ModifyBackupMigrationResponse, err error) {
    return c.ModifyBackupMigrationWithContext(context.Background(), request)
}

// ModifyBackupMigration
// 本接口（ModifyBackupMigration）用于修改备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupMigrationWithContext(ctx context.Context, request *ModifyBackupMigrationRequest) (response *ModifyBackupMigrationResponse, err error) {
    if request == nil {
        request = NewModifyBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupNameRequest() (request *ModifyBackupNameRequest) {
    request = &ModifyBackupNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupName")
    
    
    return
}

func NewModifyBackupNameResponse() (response *ModifyBackupNameResponse) {
    response = &ModifyBackupNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupName
// 本接口(ModifyBackupName)用于修改备份任务名称。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupName(request *ModifyBackupNameRequest) (response *ModifyBackupNameResponse, err error) {
    return c.ModifyBackupNameWithContext(context.Background(), request)
}

// ModifyBackupName
// 本接口(ModifyBackupName)用于修改备份任务名称。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
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

func NewModifyBackupStrategyRequest() (request *ModifyBackupStrategyRequest) {
    request = &ModifyBackupStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyBackupStrategy")
    
    
    return
}

func NewModifyBackupStrategyResponse() (response *ModifyBackupStrategyResponse) {
    response = &ModifyBackupStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupStrategy
// 本接口（ModifyBackupStrategy）用于修改备份策略
//
// 可能返回的错误码:
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
    return c.ModifyBackupStrategyWithContext(context.Background(), request)
}

// ModifyBackupStrategy
// 本接口（ModifyBackupStrategy）用于修改备份策略
//
// 可能返回的错误码:
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyBackupStrategyWithContext(ctx context.Context, request *ModifyBackupStrategyRequest) (response *ModifyBackupStrategyResponse, err error) {
    if request == nil {
        request = NewModifyBackupStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceName")
    
    
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceName
// 本接口（ModifyDBInstanceName）用于修改实例名字。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_INSTANCENAMEISILLEGAL = "InvalidParameterValue.InstanceNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    return c.ModifyDBInstanceNameWithContext(context.Background(), request)
}

// ModifyDBInstanceName
// 本接口（ModifyDBInstanceName）用于修改实例名字。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_INSTANCENAMEISILLEGAL = "InvalidParameterValue.InstanceNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceNameWithContext(ctx context.Context, request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNetworkRequest() (request *ModifyDBInstanceNetworkRequest) {
    request = &ModifyDBInstanceNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceNetwork")
    
    
    return
}

func NewModifyDBInstanceNetworkResponse() (response *ModifyDBInstanceNetworkResponse) {
    response = &ModifyDBInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceNetwork
// 本接口（ModifyDBInstanceNetwork）用于修改运行中实例的网络，仅支持从VPC网络到VPC网络的转换
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_VPCERROR = "FailedOperation.VPCError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBInstanceNetwork(request *ModifyDBInstanceNetworkRequest) (response *ModifyDBInstanceNetworkResponse, err error) {
    return c.ModifyDBInstanceNetworkWithContext(context.Background(), request)
}

// ModifyDBInstanceNetwork
// 本接口（ModifyDBInstanceNetwork）用于修改运行中实例的网络，仅支持从VPC网络到VPC网络的转换
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_VPCERROR = "FailedOperation.VPCError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INTERNALERROR_VPCERROR = "InternalError.VPCError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_VPCNOTEXIST = "ResourceNotFound.VpcNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_VPCNOTEXIST = "ResourceUnavailable.VpcNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDBInstanceNetworkWithContext(ctx context.Context, request *ModifyDBInstanceNetworkRequest) (response *ModifyDBInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceProject")
    
    
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceProject
// 本接口（ModifyDBInstanceProject）用于修改数据库实例所属项目。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    return c.ModifyDBInstanceProjectWithContext(context.Background(), request)
}

// ModifyDBInstanceProject
// 本接口（ModifyDBInstanceProject）用于修改数据库实例所属项目。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceProjectWithContext(ctx context.Context, request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceRenewFlagRequest() (request *ModifyDBInstanceRenewFlagRequest) {
    request = &ModifyDBInstanceRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceRenewFlag")
    
    
    return
}

func NewModifyDBInstanceRenewFlagResponse() (response *ModifyDBInstanceRenewFlagResponse) {
    response = &ModifyDBInstanceRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceRenewFlag
// 本接口（ModifyDBInstanceRenewFlag）用于修改实例续费标记
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceRenewFlag(request *ModifyDBInstanceRenewFlagRequest) (response *ModifyDBInstanceRenewFlagResponse, err error) {
    return c.ModifyDBInstanceRenewFlagWithContext(context.Background(), request)
}

// ModifyDBInstanceRenewFlag
// 本接口（ModifyDBInstanceRenewFlag）用于修改实例续费标记
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceRenewFlagWithContext(ctx context.Context, request *ModifyDBInstanceRenewFlagRequest) (response *ModifyDBInstanceRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
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
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SECURITYGROUPOPERATIONERROR = "FailedOperation.SecurityGroupOperationError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDISILLEGAL = "InvalidParameterValue.SecurityGroupIdIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewModifyDBNameRequest() (request *ModifyDBNameRequest) {
    request = &ModifyDBNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBName")
    
    
    return
}

func NewModifyDBNameResponse() (response *ModifyDBNameResponse) {
    response = &ModifyDBNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBName
// 本接口（ModifyDBName）用于更新数据库名。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBName(request *ModifyDBNameRequest) (response *ModifyDBNameResponse, err error) {
    return c.ModifyDBNameWithContext(context.Background(), request)
}

// ModifyDBName
// 本接口（ModifyDBName）用于更新数据库名。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBCHARILLEGAL = "InvalidParameterValue.DBCharIllegal"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  INVALIDPARAMETERVALUE_DBNAMEISKEYWRODS = "InvalidParameterValue.DBNameIsKeyWrods"
//  INVALIDPARAMETERVALUE_DBNAMESAME = "InvalidParameterValue.DBNameSame"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBNameWithContext(ctx context.Context, request *ModifyDBNameRequest) (response *ModifyDBNameResponse, err error) {
    if request == nil {
        request = NewModifyDBNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBRemarkRequest() (request *ModifyDBRemarkRequest) {
    request = &ModifyDBRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDBRemark")
    
    
    return
}

func NewModifyDBRemarkResponse() (response *ModifyDBRemarkResponse) {
    response = &ModifyDBRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBRemark
// 本接口（ModifyDBRemark）用于修改数据库备注。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBRemark(request *ModifyDBRemarkRequest) (response *ModifyDBRemarkResponse, err error) {
    return c.ModifyDBRemarkWithContext(context.Background(), request)
}

// ModifyDBRemark
// 本接口（ModifyDBRemark）用于修改数据库备注。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETERVALUE_DATABASEREMARKISILLEGAL = "InvalidParameterValue.DataBaseRemarkIsIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBRemarkWithContext(ctx context.Context, request *ModifyDBRemarkRequest) (response *ModifyDBRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDBRemarkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBRemark require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseCDCRequest() (request *ModifyDatabaseCDCRequest) {
    request = &ModifyDatabaseCDCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseCDC")
    
    
    return
}

func NewModifyDatabaseCDCResponse() (response *ModifyDatabaseCDCResponse) {
    response = &ModifyDatabaseCDCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseCDC
// 本接口(ModifyDatabaseCDC)用于开启、关闭数据库数据变更捕获(CDC)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseCDC(request *ModifyDatabaseCDCRequest) (response *ModifyDatabaseCDCResponse, err error) {
    return c.ModifyDatabaseCDCWithContext(context.Background(), request)
}

// ModifyDatabaseCDC
// 本接口(ModifyDatabaseCDC)用于开启、关闭数据库数据变更捕获(CDC)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseCDCWithContext(ctx context.Context, request *ModifyDatabaseCDCRequest) (response *ModifyDatabaseCDCResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseCDCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseCDC require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseCDCResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseCTRequest() (request *ModifyDatabaseCTRequest) {
    request = &ModifyDatabaseCTRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseCT")
    
    
    return
}

func NewModifyDatabaseCTResponse() (response *ModifyDatabaseCTResponse) {
    response = &ModifyDatabaseCTResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseCT
// 本接口(ModifyDatabaseCT)用于启用、禁用数据库数据变更跟踪(CT)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDatabaseCT(request *ModifyDatabaseCTRequest) (response *ModifyDatabaseCTResponse, err error) {
    return c.ModifyDatabaseCTWithContext(context.Background(), request)
}

// ModifyDatabaseCT
// 本接口(ModifyDatabaseCT)用于启用、禁用数据库数据变更跟踪(CT)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_NOTSUPPORTREPEAT = "UnsupportedOperation.NotSupportRepeat"
func (c *Client) ModifyDatabaseCTWithContext(ctx context.Context, request *ModifyDatabaseCTRequest) (response *ModifyDatabaseCTResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseCTRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseCT require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseCTResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseMdfRequest() (request *ModifyDatabaseMdfRequest) {
    request = &ModifyDatabaseMdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyDatabaseMdf")
    
    
    return
}

func NewModifyDatabaseMdfResponse() (response *ModifyDatabaseMdfResponse) {
    response = &ModifyDatabaseMdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatabaseMdf
// 本接口(ModifyDatabaseMdf)用于收缩数据库mdf(Shrink mdf)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseMdf(request *ModifyDatabaseMdfRequest) (response *ModifyDatabaseMdfResponse, err error) {
    return c.ModifyDatabaseMdfWithContext(context.Background(), request)
}

// ModifyDatabaseMdf
// 本接口(ModifyDatabaseMdf)用于收缩数据库mdf(Shrink mdf)
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE_DBNAMENOTNULL = "InvalidParameterValue.DBNameNotNull"
//  INVALIDPARAMETERVALUE_MODIFYTYPEVALUEINVALID = "InvalidParameterValue.ModifyTypeValueInvalid"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  RESOURCEUNAVAILABLE_NOTSUPPORTROINSTANCE = "ResourceUnavailable.NotSupportRoInstance"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDatabaseMdfWithContext(ctx context.Context, request *ModifyDatabaseMdfRequest) (response *ModifyDatabaseMdfResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseMdfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatabaseMdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatabaseMdfResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIncrementalMigrationRequest() (request *ModifyIncrementalMigrationRequest) {
    request = &ModifyIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyIncrementalMigration")
    
    
    return
}

func NewModifyIncrementalMigrationResponse() (response *ModifyIncrementalMigrationResponse) {
    response = &ModifyIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyIncrementalMigration
// 本接口（ModifyIncrementalMigration）用于修改增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyIncrementalMigration(request *ModifyIncrementalMigrationRequest) (response *ModifyIncrementalMigrationResponse, err error) {
    return c.ModifyIncrementalMigrationWithContext(context.Background(), request)
}

// ModifyIncrementalMigration
// 本接口（ModifyIncrementalMigration）用于修改增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyIncrementalMigrationWithContext(ctx context.Context, request *ModifyIncrementalMigrationRequest) (response *ModifyIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewModifyIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyInstanceParam")
    
    
    return
}

func NewModifyInstanceParamResponse() (response *ModifyInstanceParamResponse) {
    response = &ModifyInstanceParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceParam
// 本接口(ModifyInstanceParam)用于修改云数据库实例的参数。
//
// <b>注意</b>：如果修改的参数是需要<b>重启实例</b>的，那么实例将会按照WaitSwitch参数的设置(可能是立即执行也可能在可维护时间窗内自动执行)在执行参数修改时<b>重启实例</b>。
//
// 您可以通过DescribeInstanceParams接口查询修改参数时是否会重启实例，以免导致您的实例不符合预期重启。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    return c.ModifyInstanceParamWithContext(context.Background(), request)
}

// ModifyInstanceParam
// 本接口(ModifyInstanceParam)用于修改云数据库实例的参数。
//
// <b>注意</b>：如果修改的参数是需要<b>重启实例</b>的，那么实例将会按照WaitSwitch参数的设置(可能是立即执行也可能在可维护时间窗内自动执行)在执行参数修改时<b>重启实例</b>。
//
// 您可以通过DescribeInstanceParams接口查询修改参数时是否会重启实例，以免导致您的实例不符合预期重启。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyInstanceParamWithContext(ctx context.Context, request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceSpanRequest() (request *ModifyMaintenanceSpanRequest) {
    request = &ModifyMaintenanceSpanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyMaintenanceSpan")
    
    
    return
}

func NewModifyMaintenanceSpanResponse() (response *ModifyMaintenanceSpanResponse) {
    response = &ModifyMaintenanceSpanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMaintenanceSpan
// 本接口（ModifyMaintenanceSpan）用于修改实例的可维护时间窗
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintenanceSpan(request *ModifyMaintenanceSpanRequest) (response *ModifyMaintenanceSpanResponse, err error) {
    return c.ModifyMaintenanceSpanWithContext(context.Background(), request)
}

// ModifyMaintenanceSpan
// 本接口（ModifyMaintenanceSpan）用于修改实例的可维护时间窗
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMaintenanceSpanWithContext(ctx context.Context, request *ModifyMaintenanceSpanRequest) (response *ModifyMaintenanceSpanResponse, err error) {
    if request == nil {
        request = NewModifyMaintenanceSpanRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintenanceSpan require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintenanceSpanResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationRequest() (request *ModifyMigrationRequest) {
    request = &ModifyMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyMigration")
    
    
    return
}

func NewModifyMigrationResponse() (response *ModifyMigrationResponse) {
    response = &ModifyMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMigration
// 本接口（ModifyMigration）可以修改已有的迁移任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMigration(request *ModifyMigrationRequest) (response *ModifyMigrationResponse, err error) {
    return c.ModifyMigrationWithContext(context.Background(), request)
}

// ModifyMigration
// 本接口（ModifyMigration）可以修改已有的迁移任务信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_MIGRATIONNAMEISILLEGAL = "InvalidParameterValue.MigrationNameIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATEREGIONILLEGAL = "ResourceUnavailable.InstanceMigrateRegionIllegal"
//  RESOURCEUNAVAILABLE_INSTANCEMIGRATESTATUSINVALID = "ResourceUnavailable.InstanceMigrateStatusInvalid"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyMigrationWithContext(ctx context.Context, request *ModifyMigrationRequest) (response *ModifyMigrationResponse, err error) {
    if request == nil {
        request = NewModifyMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPublishSubscribeNameRequest() (request *ModifyPublishSubscribeNameRequest) {
    request = &ModifyPublishSubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyPublishSubscribeName")
    
    
    return
}

func NewModifyPublishSubscribeNameResponse() (response *ModifyPublishSubscribeNameResponse) {
    response = &ModifyPublishSubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPublishSubscribeName
// 本接口（ModifyPublishSubscribeName）修改发布订阅的名称。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyPublishSubscribeName(request *ModifyPublishSubscribeNameRequest) (response *ModifyPublishSubscribeNameResponse, err error) {
    return c.ModifyPublishSubscribeNameWithContext(context.Background(), request)
}

// ModifyPublishSubscribeName
// 本接口（ModifyPublishSubscribeName）修改发布订阅的名称。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PUBSUBNAMEISILLEGAL = "InvalidParameterValue.PubSubNameIsIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyPublishSubscribeNameWithContext(ctx context.Context, request *ModifyPublishSubscribeNameRequest) (response *ModifyPublishSubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifyPublishSubscribeNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPublishSubscribeName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPublishSubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyReadOnlyGroupDetailsRequest() (request *ModifyReadOnlyGroupDetailsRequest) {
    request = &ModifyReadOnlyGroupDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ModifyReadOnlyGroupDetails")
    
    
    return
}

func NewModifyReadOnlyGroupDetailsResponse() (response *ModifyReadOnlyGroupDetailsResponse) {
    response = &ModifyReadOnlyGroupDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyReadOnlyGroupDetails
// 本接口（ModifyReadOnlyGroupDetails）用于修改只读组详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyReadOnlyGroupDetails(request *ModifyReadOnlyGroupDetailsRequest) (response *ModifyReadOnlyGroupDetailsResponse, err error) {
    return c.ModifyReadOnlyGroupDetailsWithContext(context.Background(), request)
}

// ModifyReadOnlyGroupDetails
// 本接口（ModifyReadOnlyGroupDetails）用于修改只读组详情。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyReadOnlyGroupDetailsWithContext(ctx context.Context, request *ModifyReadOnlyGroupDetailsRequest) (response *ModifyReadOnlyGroupDetailsResponse, err error) {
    if request == nil {
        request = NewModifyReadOnlyGroupDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReadOnlyGroupDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReadOnlyGroupDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewOpenInterCommunicationRequest() (request *OpenInterCommunicationRequest) {
    request = &OpenInterCommunicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "OpenInterCommunication")
    
    
    return
}

func NewOpenInterCommunicationResponse() (response *OpenInterCommunicationResponse) {
    response = &OpenInterCommunicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenInterCommunication
// 本接口（OpenInterCommunication）用于打开实例的互通，实例互通可以实现商业智能服务相互联通。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenInterCommunication(request *OpenInterCommunicationRequest) (response *OpenInterCommunicationResponse, err error) {
    return c.OpenInterCommunicationWithContext(context.Background(), request)
}

// OpenInterCommunication
// 本接口（OpenInterCommunication）用于打开实例的互通，实例互通可以实现商业智能服务相互联通。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) OpenInterCommunicationWithContext(ctx context.Context, request *OpenInterCommunicationRequest) (response *OpenInterCommunicationResponse, err error) {
    if request == nil {
        request = NewOpenInterCommunicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenInterCommunication require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenInterCommunicationResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMigrationCheckProcessRequest() (request *QueryMigrationCheckProcessRequest) {
    request = &QueryMigrationCheckProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "QueryMigrationCheckProcess")
    
    
    return
}

func NewQueryMigrationCheckProcessResponse() (response *QueryMigrationCheckProcessResponse) {
    response = &QueryMigrationCheckProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMigrationCheckProcess
// 本接口（QueryMigrationCheckProcess）的作用是查询迁移检查任务的进度，适用于迁移源的类型为TencentDB for SQLServer 的迁移方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) QueryMigrationCheckProcess(request *QueryMigrationCheckProcessRequest) (response *QueryMigrationCheckProcessResponse, err error) {
    return c.QueryMigrationCheckProcessWithContext(context.Background(), request)
}

// QueryMigrationCheckProcess
// 本接口（QueryMigrationCheckProcess）的作用是查询迁移检查任务的进度，适用于迁移源的类型为TencentDB for SQLServer 的迁移方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) QueryMigrationCheckProcessWithContext(ctx context.Context, request *QueryMigrationCheckProcessRequest) (response *QueryMigrationCheckProcessResponse, err error) {
    if request == nil {
        request = NewQueryMigrationCheckProcessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMigrationCheckProcess require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMigrationCheckProcessResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleDBInstanceRequest() (request *RecycleDBInstanceRequest) {
    request = &RecycleDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RecycleDBInstance")
    
    
    return
}

func NewRecycleDBInstanceResponse() (response *RecycleDBInstanceResponse) {
    response = &RecycleDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecycleDBInstance
// 本接口（RecycleDBInstance）用于主动回收已下线的SQLSERVER实例
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecycleDBInstance(request *RecycleDBInstanceRequest) (response *RecycleDBInstanceResponse, err error) {
    return c.RecycleDBInstanceWithContext(context.Background(), request)
}

// RecycleDBInstance
// 本接口（RecycleDBInstance）用于主动回收已下线的SQLSERVER实例
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecycleDBInstanceWithContext(ctx context.Context, request *RecycleDBInstanceRequest) (response *RecycleDBInstanceResponse, err error) {
    if request == nil {
        request = NewRecycleDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecycleDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecycleDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRecycleReadOnlyGroupRequest() (request *RecycleReadOnlyGroupRequest) {
    request = &RecycleReadOnlyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RecycleReadOnlyGroup")
    
    
    return
}

func NewRecycleReadOnlyGroupResponse() (response *RecycleReadOnlyGroupResponse) {
    response = &RecycleReadOnlyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecycleReadOnlyGroup
// 本接口（RecycleReadOnlyGroup）立即回收只读组的资源，只读组占用的vip等资源将立即释放且不可找回。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ROGROUPSTATUSISILLEGAL = "InvalidParameterValue.RoGroupStatusIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecycleReadOnlyGroup(request *RecycleReadOnlyGroupRequest) (response *RecycleReadOnlyGroupResponse, err error) {
    return c.RecycleReadOnlyGroupWithContext(context.Background(), request)
}

// RecycleReadOnlyGroup
// 本接口（RecycleReadOnlyGroup）立即回收只读组的资源，只读组占用的vip等资源将立即释放且不可找回。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_ROGROUPSTATUSISILLEGAL = "InvalidParameterValue.RoGroupStatusIsIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecycleReadOnlyGroupWithContext(ctx context.Context, request *RecycleReadOnlyGroupRequest) (response *RecycleReadOnlyGroupResponse, err error) {
    if request == nil {
        request = NewRecycleReadOnlyGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecycleReadOnlyGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecycleReadOnlyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveBackupsRequest() (request *RemoveBackupsRequest) {
    request = &RemoveBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RemoveBackups")
    
    
    return
}

func NewRemoveBackupsResponse() (response *RemoveBackupsResponse) {
    response = &RemoveBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RemoveBackups
// 本接口（RemoveBackups）可以删除用户手动创建的备份文件。待删除的备份策略可以是实例备份，也可以是多库备份。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveBackups(request *RemoveBackupsRequest) (response *RemoveBackupsResponse, err error) {
    return c.RemoveBackupsWithContext(context.Background(), request)
}

// RemoveBackups
// 本接口（RemoveBackups）可以删除用户手动创建的备份文件。待删除的备份策略可以是实例备份，也可以是多库备份。
//
// 可能返回的错误码:
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RemoveBackupsWithContext(ctx context.Context, request *RemoveBackupsRequest) (response *RemoveBackupsResponse, err error) {
    if request == nil {
        request = NewRemoveBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBInstanceRequest() (request *RenewDBInstanceRequest) {
    request = &RenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RenewDBInstance")
    
    
    return
}

func NewRenewDBInstanceResponse() (response *RenewDBInstanceResponse) {
    response = &RenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewDBInstance
// 本接口（RenewDBInstance）用于续费实例。当被续费实例是按量计费实例时，则按量计费实例转为包年包月计费方式。
//
// 按量计费实例转包年包月询价可通过(InquiryPriceRenewDBInstance)接口获得。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewDBInstance(request *RenewDBInstanceRequest) (response *RenewDBInstanceResponse, err error) {
    return c.RenewDBInstanceWithContext(context.Background(), request)
}

// RenewDBInstance
// 本接口（RenewDBInstance）用于续费实例。当被续费实例是按量计费实例时，则按量计费实例转为包年包月计费方式。
//
// 按量计费实例转包年包月询价可通过(InquiryPriceRenewDBInstance)接口获得。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewDBInstanceWithContext(ctx context.Context, request *RenewDBInstanceRequest) (response *RenewDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewPostpaidDBInstanceRequest() (request *RenewPostpaidDBInstanceRequest) {
    request = &RenewPostpaidDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RenewPostpaidDBInstance")
    
    
    return
}

func NewRenewPostpaidDBInstanceResponse() (response *RenewPostpaidDBInstanceResponse) {
    response = &RenewPostpaidDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewPostpaidDBInstance
// 本接口（RenewPostpaidDBInstance）用于将通过接口TerminateDBInstance手动隔离的按量计费实例从回收站中恢复。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewPostpaidDBInstance(request *RenewPostpaidDBInstanceRequest) (response *RenewPostpaidDBInstanceResponse, err error) {
    return c.RenewPostpaidDBInstanceWithContext(context.Background(), request)
}

// RenewPostpaidDBInstance
// 本接口（RenewPostpaidDBInstance）用于将通过接口TerminateDBInstance手动隔离的按量计费实例从回收站中恢复。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewPostpaidDBInstanceWithContext(ctx context.Context, request *RenewPostpaidDBInstanceRequest) (response *RenewPostpaidDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewPostpaidDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewPostpaidDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewPostpaidDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于重置实例的账户密码。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于重置实例的账户密码。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_PASSWORDISILLEGAL = "InvalidParameterValue.PasswordIsIllegal"
//  RESOURCENOTFOUND_ACCOUNTNOTEXIST = "ResourceNotFound.AccountNotExist"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_ACCOUNTINVALIDSTATUS = "ResourceUnavailable.AccountInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
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

func NewRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
    request = &RestartDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RestartDBInstance")
    
    
    return
}

func NewRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
    response = &RestartDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartDBInstance
// 本接口（RestartDBInstance）用于重启数据库实例。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    return c.RestartDBInstanceWithContext(context.Background(), request)
}

// RestartDBInstance
// 本接口（RestartDBInstance）用于重启数据库实例。
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartDBInstanceWithContext(ctx context.Context, request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
    if request == nil {
        request = NewRestartDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestoreInstanceRequest() (request *RestoreInstanceRequest) {
    request = &RestoreInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RestoreInstance")
    
    
    return
}

func NewRestoreInstanceResponse() (response *RestoreInstanceResponse) {
    response = &RestoreInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestoreInstance
// 本接口（RestoreInstance）用于根据备份文件恢复实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    return c.RestoreInstanceWithContext(context.Background(), request)
}

// RestoreInstance
// 本接口（RestoreInstance）用于根据备份文件恢复实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_GCSERROR = "FailedOperation.GcsError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_DBEXIST = "InvalidParameterValue.DBExist"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestoreInstanceWithContext(ctx context.Context, request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    if request == nil {
        request = NewRestoreInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestoreInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestoreInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRollbackInstanceRequest() (request *RollbackInstanceRequest) {
    request = &RollbackInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RollbackInstance")
    
    
    return
}

func NewRollbackInstanceResponse() (response *RollbackInstanceResponse) {
    response = &RollbackInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RollbackInstance
// 本接口（RollbackInstance）用于回档实例
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollbackInstance(request *RollbackInstanceRequest) (response *RollbackInstanceResponse, err error) {
    return c.RollbackInstanceWithContext(context.Background(), request)
}

// RollbackInstance
// 本接口（RollbackInstance）用于回档实例
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RollbackInstanceWithContext(ctx context.Context, request *RollbackInstanceRequest) (response *RollbackInstanceResponse, err error) {
    if request == nil {
        request = NewRollbackInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RollbackInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRollbackInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRunMigrationRequest() (request *RunMigrationRequest) {
    request = &RunMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "RunMigration")
    
    
    return
}

func NewRunMigrationResponse() (response *RunMigrationResponse) {
    response = &RunMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunMigration
// 本接口（RunMigration）用于启动迁移任务，开始迁移
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RunMigration(request *RunMigrationRequest) (response *RunMigrationResponse, err error) {
    return c.RunMigrationWithContext(context.Background(), request)
}

// RunMigration
// 本接口（RunMigration）用于启动迁移任务，开始迁移
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  LIMITEXCEEDED_TOOMANYDB = "LimitExceeded.TooManyDB"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RunMigrationWithContext(ctx context.Context, request *RunMigrationRequest) (response *RunMigrationResponse, err error) {
    if request == nil {
        request = NewRunMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartBackupMigrationRequest() (request *StartBackupMigrationRequest) {
    request = &StartBackupMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartBackupMigration")
    
    
    return
}

func NewStartBackupMigrationResponse() (response *StartBackupMigrationResponse) {
    response = &StartBackupMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartBackupMigration
// 本接口（StartBackupMigration）用于启动备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartBackupMigration(request *StartBackupMigrationRequest) (response *StartBackupMigrationResponse, err error) {
    return c.StartBackupMigrationWithContext(context.Background(), request)
}

// StartBackupMigration
// 本接口（StartBackupMigration）用于启动备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBERROR = "FailedOperation.DBError"
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartBackupMigrationWithContext(ctx context.Context, request *StartBackupMigrationRequest) (response *StartBackupMigrationResponse, err error) {
    if request == nil {
        request = NewStartBackupMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBackupMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBackupMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartIncrementalMigrationRequest() (request *StartIncrementalMigrationRequest) {
    request = &StartIncrementalMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartIncrementalMigration")
    
    
    return
}

func NewStartIncrementalMigrationResponse() (response *StartIncrementalMigrationResponse) {
    response = &StartIncrementalMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartIncrementalMigration
// 本接口（StartIncrementalMigration）用于启动增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartIncrementalMigration(request *StartIncrementalMigrationRequest) (response *StartIncrementalMigrationResponse, err error) {
    return c.StartIncrementalMigrationWithContext(context.Background(), request)
}

// StartIncrementalMigration
// 本接口（StartIncrementalMigration）用于启动增量备份导入任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIGRATIONLOCKERROR = "FailedOperation.MigrationLockError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETERVALUE_BACKUPNAMEISILLEGAL = "InvalidParameterValue.BackupNameIsIllegal"
//  RESOURCENOTFOUND_BACKUPNOTFOUND = "ResourceNotFound.BackupNotFound"
//  RESOURCENOTFOUND_FULLBACKUPMIGRATIONNOTEXIST = "ResourceNotFound.FullBackupMigrationNotExist"
//  RESOURCEUNAVAILABLE_COSSTATUSERR = "ResourceUnavailable.CosStatusErr"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartIncrementalMigrationWithContext(ctx context.Context, request *StartIncrementalMigrationRequest) (response *StartIncrementalMigrationResponse, err error) {
    if request == nil {
        request = NewStartIncrementalMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartIncrementalMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartIncrementalMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewStartMigrationCheckRequest() (request *StartMigrationCheckRequest) {
    request = &StartMigrationCheckRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "StartMigrationCheck")
    
    
    return
}

func NewStartMigrationCheckResponse() (response *StartMigrationCheckResponse) {
    response = &StartMigrationCheckResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMigrationCheck
// 本接口（StartMigrationCheck）的作用是启动一个迁移前的校验任务，适用于迁移源的类型为TencentDB for SQLServer 的迁移方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartMigrationCheck(request *StartMigrationCheckRequest) (response *StartMigrationCheckResponse, err error) {
    return c.StartMigrationCheckWithContext(context.Background(), request)
}

// StartMigrationCheck
// 本接口（StartMigrationCheck）的作用是启动一个迁移前的校验任务，适用于迁移源的类型为TencentDB for SQLServer 的迁移方式
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StartMigrationCheckWithContext(ctx context.Context, request *StartMigrationCheckRequest) (response *StartMigrationCheckResponse, err error) {
    if request == nil {
        request = NewStartMigrationCheckRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMigrationCheck require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMigrationCheckResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrationRequest() (request *StopMigrationRequest) {
    request = &StopMigrationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "StopMigration")
    
    
    return
}

func NewStopMigrationResponse() (response *StopMigrationResponse) {
    response = &StopMigrationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMigration
// 本接口（StopMigration）作用是中止一个迁移任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StopMigration(request *StopMigrationRequest) (response *StopMigrationResponse, err error) {
    return c.StopMigrationWithContext(context.Background(), request)
}

// StopMigration
// 本接口（StopMigration）作用是中止一个迁移任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_DBCONNECTERROR = "InternalError.DBConnectError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_GCSERROR = "InternalError.GcsError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCENOTFOUND_DBNOTFOUND = "ResourceNotFound.DBNotFound"
//  RESOURCEUNAVAILABLE_DBINVALIDSTATUS = "ResourceUnavailable.DBInvalidStatus"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) StopMigrationWithContext(ctx context.Context, request *StopMigrationRequest) (response *StopMigrationResponse, err error) {
    if request == nil {
        request = NewStopMigrationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMigration require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMigrationResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDBInstanceRequest() (request *TerminateDBInstanceRequest) {
    request = &TerminateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "TerminateDBInstance")
    
    
    return
}

func NewTerminateDBInstanceResponse() (response *TerminateDBInstanceResponse) {
    response = &TerminateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDBInstance
// 本接口(TerminateDBInstance)用于主动隔离实例，使得实例进入回收站。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDBInstance(request *TerminateDBInstanceRequest) (response *TerminateDBInstanceResponse, err error) {
    return c.TerminateDBInstanceWithContext(context.Background(), request)
}

// TerminateDBInstance
// 本接口(TerminateDBInstance)用于主动隔离实例，使得实例进入回收站。
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSINVALID = "ResourceUnavailable.InstanceStatusInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDBInstanceWithContext(ctx context.Context, request *TerminateDBInstanceRequest) (response *TerminateDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("sqlserver", APIVersion, "UpgradeDBInstance")
    
    
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeDBInstance
// 本接口（UpgradeDBInstance）用于升级实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    return c.UpgradeDBInstanceWithContext(context.Background(), request)
}

// UpgradeDBInstance
// 本接口（UpgradeDBInstance）用于升级实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_PARAMSASSERTFAILED = "InvalidParameter.ParamsAssertFailed"
//  INVALIDPARAMETER_PAYORDERFAILED = "InvalidParameter.PayOrderFailed"
//  INVALIDPARAMETERVALUE_ILLEGALSPEC = "InvalidParameterValue.IllegalSpec"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDBInstanceWithContext(ctx context.Context, request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceResponse()
    err = c.Send(request, response)
    return
}
