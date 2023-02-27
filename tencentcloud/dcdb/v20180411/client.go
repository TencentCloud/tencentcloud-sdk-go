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

package v20180411

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-11"

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


func NewActiveHourDCDBInstanceRequest() (request *ActiveHourDCDBInstanceRequest) {
    request = &ActiveHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ActiveHourDCDBInstance")
    
    
    return
}

func NewActiveHourDCDBInstanceResponse() (response *ActiveHourDCDBInstanceResponse) {
    response = &ActiveHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ActiveHourDCDBInstance
// 解隔离DCDB后付费实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ActiveHourDCDBInstance(request *ActiveHourDCDBInstanceRequest) (response *ActiveHourDCDBInstanceResponse, err error) {
    return c.ActiveHourDCDBInstanceWithContext(context.Background(), request)
}

// ActiveHourDCDBInstance
// 解隔离DCDB后付费实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ActiveHourDCDBInstanceWithContext(ctx context.Context, request *ActiveHourDCDBInstanceRequest) (response *ActiveHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewActiveHourDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActiveHourDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewActiveHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// 本接口 (AssociateSecurityGroups) 用于安全组批量绑定云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_SETSVCLOCATIONFAILED = "FailedOperation.SetSvcLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_DBROWSAFFECTEDERROR = "InternalError.DBRowsAffectedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSERSGCOUNTFAILED = "InternalError.GetUserSGCountFailed"
//  INTERNALERROR_GETUSGQUOTAERROR = "InternalError.GetUsgQuotaError"
//  INTERNALERROR_INSERTFAIL = "InternalError.InsertFail"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INTERNALERROR_UPDATEDATABASEFAILED = "InternalError.UpdateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_SGCHECKFAIL = "InvalidParameter.SGCheckFail"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// 本接口 (AssociateSecurityGroups) 用于安全组批量绑定云资源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_SETSVCLOCATIONFAILED = "FailedOperation.SetSvcLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_DBROWSAFFECTEDERROR = "InternalError.DBRowsAffectedError"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSERSGCOUNTFAILED = "InternalError.GetUserSGCountFailed"
//  INTERNALERROR_GETUSGQUOTAERROR = "InternalError.GetUsgQuotaError"
//  INTERNALERROR_INSERTFAIL = "InternalError.InsertFail"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INTERNALERROR_UPDATEDATABASEFAILED = "InternalError.UpdateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_SGCHECKFAIL = "InvalidParameter.SGCheckFail"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"
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

func NewCancelDcnJobRequest() (request *CancelDcnJobRequest) {
    request = &CancelDcnJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CancelDcnJob")
    
    
    return
}

func NewCancelDcnJobResponse() (response *CancelDcnJobResponse) {
    response = &CancelDcnJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelDcnJob
// 取消DCN同步
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CancelDcnJob(request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
    return c.CancelDcnJobWithContext(context.Background(), request)
}

// CancelDcnJob
// 取消DCN同步
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CancelDcnJobWithContext(ctx context.Context, request *CancelDcnJobRequest) (response *CancelDcnJobResponse, err error) {
    if request == nil {
        request = NewCancelDcnJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelDcnJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelDcnJobResponse()
    err = c.Send(request, response)
    return
}

func NewCloneAccountRequest() (request *CloneAccountRequest) {
    request = &CloneAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CloneAccount")
    
    
    return
}

func NewCloneAccountResponse() (response *CloneAccountResponse) {
    response = &CloneAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloneAccount
// 本接口（CloneAccount）用于克隆实例账户。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) CloneAccount(request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    return c.CloneAccountWithContext(context.Background(), request)
}

// CloneAccount
// 本接口（CloneAccount）用于克隆实例账户。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) CloneAccountWithContext(ctx context.Context, request *CloneAccountRequest) (response *CloneAccountResponse, err error) {
    if request == nil {
        request = NewCloneAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseDBExtranetAccessRequest() (request *CloseDBExtranetAccessRequest) {
    request = &CloseDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CloseDBExtranetAccess")
    
    
    return
}

func NewCloseDBExtranetAccessResponse() (response *CloseDBExtranetAccessResponse) {
    response = &CloseDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseDBExtranetAccess
// 本接口(CloseDBExtranetAccess)用于关闭云数据库实例的外网访问。关闭外网访问后，外网地址将不可访问，查询实例列表接口将不返回对应实例的外网域名和端口信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INTERNALERROR_WANSERVICEFAILED = "InternalError.WanServiceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseDBExtranetAccess(request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    return c.CloseDBExtranetAccessWithContext(context.Background(), request)
}

// CloseDBExtranetAccess
// 本接口(CloseDBExtranetAccess)用于关闭云数据库实例的外网访问。关闭外网访问后，外网地址将不可访问，查询实例列表接口将不返回对应实例的外网域名和端口信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INTERNALERROR_WANSERVICEFAILED = "InternalError.WanServiceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CloseDBExtranetAccessWithContext(ctx context.Context, request *CloseDBExtranetAccessRequest) (response *CloseDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewCloseDBExtranetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseDBExtranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewCopyAccountPrivilegesRequest() (request *CopyAccountPrivilegesRequest) {
    request = &CopyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CopyAccountPrivileges")
    
    
    return
}

func NewCopyAccountPrivilegesResponse() (response *CopyAccountPrivilegesResponse) {
    response = &CopyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyAccountPrivileges
// 本接口（CopyAccountPrivileges）用于复制云数据库账号的权限。
//
// 注意：相同用户名，不同Host是不同的账号，Readonly属性相同的账号之间才能复制权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COPYRIGHTERROR = "FailedOperation.CopyRightError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERTYPE = "InvalidParameterValue.BadUserType"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyAccountPrivileges(request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    return c.CopyAccountPrivilegesWithContext(context.Background(), request)
}

// CopyAccountPrivileges
// 本接口（CopyAccountPrivileges）用于复制云数据库账号的权限。
//
// 注意：相同用户名，不同Host是不同的账号，Readonly属性相同的账号之间才能复制权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COPYRIGHTERROR = "FailedOperation.CopyRightError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERTYPE = "InvalidParameterValue.BadUserType"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CopyAccountPrivilegesWithContext(ctx context.Context, request *CopyAccountPrivilegesRequest) (response *CopyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewCopyAccountPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountRequest() (request *CreateAccountRequest) {
    request = &CreateAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateAccount")
    
    
    return
}

func NewCreateAccountResponse() (response *CreateAccountResponse) {
    response = &CreateAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccount
// 本接口（CreateAccount）用于创建云数据库账号。一个实例可以创建多个不同的账号，相同的用户名+不同的host是不同的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEUSERFAILED = "FailedOperation.CreateUserFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_OLDPROXYVERSION = "UnsupportedOperation.OldProxyVersion"
func (c *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
    return c.CreateAccountWithContext(context.Background(), request)
}

// CreateAccount
// 本接口（CreateAccount）用于创建云数据库账号。一个实例可以创建多个不同的账号，相同的用户名+不同的host是不同的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEUSERFAILED = "FailedOperation.CreateUserFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ACCOUNTALREADYEXISTS = "InvalidParameterValue.AccountAlreadyExists"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_OLDPROXYVERSION = "UnsupportedOperation.OldProxyVersion"
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

func NewCreateDCDBInstanceRequest() (request *CreateDCDBInstanceRequest) {
    request = &CreateDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateDCDBInstance")
    
    
    return
}

func NewCreateDCDBInstanceResponse() (response *CreateDCDBInstanceResponse) {
    response = &CreateDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDCDBInstance
// 本接口（CreateDCDBInstance）用于创建包年包月的云数据库实例，可通过传入实例规格、数据库版本号、购买时长等信息创建云数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALQUANTITY = "InvalidParameterValue.IllegalQuantity"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDCDBInstance(request *CreateDCDBInstanceRequest) (response *CreateDCDBInstanceResponse, err error) {
    return c.CreateDCDBInstanceWithContext(context.Background(), request)
}

// CreateDCDBInstance
// 本接口（CreateDCDBInstance）用于创建包年包月的云数据库实例，可通过传入实例规格、数据库版本号、购买时长等信息创建云数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALQUANTITY = "InvalidParameterValue.IllegalQuantity"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateDCDBInstanceWithContext(ctx context.Context, request *CreateDCDBInstanceRequest) (response *CreateDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHourDCDBInstanceRequest() (request *CreateHourDCDBInstanceRequest) {
    request = &CreateHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "CreateHourDCDBInstance")
    
    
    return
}

func NewCreateHourDCDBInstanceResponse() (response *CreateHourDCDBInstanceResponse) {
    response = &CreateHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateHourDCDBInstance
// 创建DCDB后付费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INTERNALERROR_RETREATETIME = "InternalError.RetreateTime"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateHourDCDBInstance(request *CreateHourDCDBInstanceRequest) (response *CreateHourDCDBInstanceResponse, err error) {
    return c.CreateHourDCDBInstanceWithContext(context.Background(), request)
}

// CreateHourDCDBInstance
// 创建DCDB后付费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_TAGDRYRUNERROR = "FailedOperation.TagDryRunError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INTERNALERROR_RETREATETIME = "InternalError.RetreateTime"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CHECKPARAMNOTPASS = "InvalidParameter.CheckParamNotPass"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALZONE = "InvalidParameterValue.IllegalZone"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateHourDCDBInstanceWithContext(ctx context.Context, request *CreateHourDCDBInstanceRequest) (response *CreateHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateHourDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateHourDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountRequest() (request *DeleteAccountRequest) {
    request = &DeleteAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DeleteAccount")
    
    
    return
}

func NewDeleteAccountResponse() (response *DeleteAccountResponse) {
    response = &DeleteAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccount
// 本接口（DeleteAccount）用于删除云数据库账号。用户名+host唯一确定一个账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERFAILED = "FailedOperation.DeleteUserFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteAccount(request *DeleteAccountRequest) (response *DeleteAccountResponse, err error) {
    return c.DeleteAccountWithContext(context.Background(), request)
}

// DeleteAccount
// 本接口（DeleteAccount）用于删除云数据库账号。用户名+host唯一确定一个账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETEUSERFAILED = "FailedOperation.DeleteUserFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
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

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccountPrivileges")
    
    
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountPrivileges
// 本接口（DescribeAccountPrivileges）用于查询云数据库账号权限。
//
// 注意：注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// 本接口（DescribeAccountPrivileges）用于查询云数据库账号权限。
//
// 注意：注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    
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
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询指定云数据库实例的账号列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询指定云数据库实例的账号列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETUSERLISTFAILED = "InternalError.GetUserListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
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

func NewDescribeDBLogFilesRequest() (request *DescribeDBLogFilesRequest) {
    request = &DescribeDBLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBLogFiles")
    
    
    return
}

func NewDescribeDBLogFilesResponse() (response *DescribeDBLogFilesResponse) {
    response = &DescribeDBLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBLogFiles
// 本接口(DescribeDBLogFiles)用于获取数据库的各种日志列表，包括冷备、binlog、errlog和slowlog。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_COSCONFIGURATION = "InternalError.CosConfiguration"
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_COSAPIFAILED = "ResourceUnavailable.CosApiFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBLogFiles(request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    return c.DescribeDBLogFilesWithContext(context.Background(), request)
}

// DescribeDBLogFiles
// 本接口(DescribeDBLogFiles)用于获取数据库的各种日志列表，包括冷备、binlog、errlog和slowlog。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_COSCONFIGURATION = "InternalError.CosConfiguration"
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_COSAPIFAILED = "ResourceUnavailable.CosApiFailed"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBLogFilesWithContext(ctx context.Context, request *DescribeDBLogFilesRequest) (response *DescribeDBLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeDBLogFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBLogFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
    request = &DescribeDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBParameters")
    
    
    return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
    response = &DescribeDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBParameters
// 本接口(DescribeDBParameters)用于获取数据库的当前参数设置。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    return c.DescribeDBParametersWithContext(context.Background(), request)
}

// DescribeDBParameters
// 本接口(DescribeDBParameters)用于获取数据库的当前参数设置。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBParametersWithContext(ctx context.Context, request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBParametersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSecurityGroups
// 本接口（DescribeDBSecurityGroups）用于查询实例安全组信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 本接口（DescribeDBSecurityGroups）用于查询实例安全组信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
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

func NewDescribeDBSlowLogsRequest() (request *DescribeDBSlowLogsRequest) {
    request = &DescribeDBSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSlowLogs")
    
    
    return
}

func NewDescribeDBSlowLogsResponse() (response *DescribeDBSlowLogsResponse) {
    response = &DescribeDBSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSlowLogs
// 本接口(DescribeDBSlowLogs)用于查询慢查询日志列表。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETSLOWLOGFAILED = "InternalError.GetSlowLogFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_ILLEGALTIME = "InvalidParameter.IllegalTime"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSlowLogs(request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
    return c.DescribeDBSlowLogsWithContext(context.Background(), request)
}

// DescribeDBSlowLogs
// 本接口(DescribeDBSlowLogs)用于查询慢查询日志列表。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETSLOWLOGFAILED = "InternalError.GetSlowLogFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_ILLEGALTIME = "InvalidParameter.IllegalTime"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSlowLogsWithContext(ctx context.Context, request *DescribeDBSlowLogsRequest) (response *DescribeDBSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSyncModeRequest() (request *DescribeDBSyncModeRequest) {
    request = &DescribeDBSyncModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDBSyncMode")
    
    
    return
}

func NewDescribeDBSyncModeResponse() (response *DescribeDBSyncModeResponse) {
    response = &DescribeDBSyncModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSyncMode
// 本接口（DescribeDBSyncMode）用于查询云数据库实例的同步模式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSyncMode(request *DescribeDBSyncModeRequest) (response *DescribeDBSyncModeResponse, err error) {
    return c.DescribeDBSyncModeWithContext(context.Background(), request)
}

// DescribeDBSyncMode
// 本接口（DescribeDBSyncMode）用于查询云数据库实例的同步模式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDBSyncModeWithContext(ctx context.Context, request *DescribeDBSyncModeRequest) (response *DescribeDBSyncModeResponse, err error) {
    if request == nil {
        request = NewDescribeDBSyncModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSyncMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstanceNodeInfoRequest() (request *DescribeDCDBInstanceNodeInfoRequest) {
    request = &DescribeDCDBInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstanceNodeInfo")
    
    
    return
}

func NewDescribeDCDBInstanceNodeInfoResponse() (response *DescribeDCDBInstanceNodeInfoResponse) {
    response = &DescribeDCDBInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBInstanceNodeInfo
// 本接口（DescribeDCDBInstanceNodeInfo）用于获取实例节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDCDBInstanceNodeInfo(request *DescribeDCDBInstanceNodeInfoRequest) (response *DescribeDCDBInstanceNodeInfoResponse, err error) {
    return c.DescribeDCDBInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeDCDBInstanceNodeInfo
// 本接口（DescribeDCDBInstanceNodeInfo）用于获取实例节点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDCDBInstanceNodeInfoWithContext(ctx context.Context, request *DescribeDCDBInstanceNodeInfoRequest) (response *DescribeDCDBInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstanceNodeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBInstanceNodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBInstancesRequest() (request *DescribeDCDBInstancesRequest) {
    request = &DescribeDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBInstances")
    
    
    return
}

func NewDescribeDCDBInstancesResponse() (response *DescribeDCDBInstancesResponse) {
    response = &DescribeDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBInstances
// 查询云数据库实例列表，支持通过项目ID、实例ID、内网地址、实例名称等来筛选实例。
//
// 如果不指定任何筛选条件，则默认返回10条实例记录，单次请求最多支持返回100条实例记录。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBInstances(request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
    return c.DescribeDCDBInstancesWithContext(context.Background(), request)
}

// DescribeDCDBInstances
// 查询云数据库实例列表，支持通过项目ID、实例ID、内网地址、实例名称等来筛选实例。
//
// 如果不指定任何筛选条件，则默认返回10条实例记录，单次请求最多支持返回100条实例记录。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SUBNETNOTFOUND = "InvalidParameter.SubnetNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBInstancesWithContext(ctx context.Context, request *DescribeDCDBInstancesRequest) (response *DescribeDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBPriceRequest() (request *DescribeDCDBPriceRequest) {
    request = &DescribeDCDBPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBPrice")
    
    
    return
}

func NewDescribeDCDBPriceResponse() (response *DescribeDCDBPriceResponse) {
    response = &DescribeDCDBPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBPrice
// 本接口（DescribeDCDBPrice）用于在购买实例前，查询实例的价格。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBPrice(request *DescribeDCDBPriceRequest) (response *DescribeDCDBPriceResponse, err error) {
    return c.DescribeDCDBPriceWithContext(context.Background(), request)
}

// DescribeDCDBPrice
// 本接口（DescribeDCDBPrice）用于在购买实例前，查询实例的价格。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBPriceWithContext(ctx context.Context, request *DescribeDCDBPriceRequest) (response *DescribeDCDBPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBRenewalPriceRequest() (request *DescribeDCDBRenewalPriceRequest) {
    request = &DescribeDCDBRenewalPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBRenewalPrice")
    
    
    return
}

func NewDescribeDCDBRenewalPriceResponse() (response *DescribeDCDBRenewalPriceResponse) {
    response = &DescribeDCDBRenewalPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBRenewalPrice
// 本接口（DescribeDCDBRenewalPrice）用于在续费分布式数据库实例时，查询续费的价格。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBRenewalPrice(request *DescribeDCDBRenewalPriceRequest) (response *DescribeDCDBRenewalPriceResponse, err error) {
    return c.DescribeDCDBRenewalPriceWithContext(context.Background(), request)
}

// DescribeDCDBRenewalPrice
// 本接口（DescribeDCDBRenewalPrice）用于在续费分布式数据库实例时，查询续费的价格。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBRenewalPriceWithContext(ctx context.Context, request *DescribeDCDBRenewalPriceRequest) (response *DescribeDCDBRenewalPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBRenewalPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBRenewalPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBRenewalPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBSaleInfoRequest() (request *DescribeDCDBSaleInfoRequest) {
    request = &DescribeDCDBSaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBSaleInfo")
    
    
    return
}

func NewDescribeDCDBSaleInfoResponse() (response *DescribeDCDBSaleInfoResponse) {
    response = &DescribeDCDBSaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBSaleInfo
// 本接口(DescribeDCDBSaleInfo)用于查询分布式数据库可售卖的地域和可用区信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBSaleInfo(request *DescribeDCDBSaleInfoRequest) (response *DescribeDCDBSaleInfoResponse, err error) {
    return c.DescribeDCDBSaleInfoWithContext(context.Background(), request)
}

// DescribeDCDBSaleInfo
// 本接口(DescribeDCDBSaleInfo)用于查询分布式数据库可售卖的地域和可用区信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBSaleInfoWithContext(ctx context.Context, request *DescribeDCDBSaleInfoRequest) (response *DescribeDCDBSaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBSaleInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBSaleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBSaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBShardsRequest() (request *DescribeDCDBShardsRequest) {
    request = &DescribeDCDBShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBShards")
    
    
    return
}

func NewDescribeDCDBShardsResponse() (response *DescribeDCDBShardsResponse) {
    response = &DescribeDCDBShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBShards
// 本接口（DescribeDCDBShards）用于查询云数据库实例的分片信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBShards(request *DescribeDCDBShardsRequest) (response *DescribeDCDBShardsResponse, err error) {
    return c.DescribeDCDBShardsWithContext(context.Background(), request)
}

// DescribeDCDBShards
// 本接口（DescribeDCDBShards）用于查询云数据库实例的分片信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_FENCEERROR = "InternalError.FenceError"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALEXCLUSTERID = "InvalidParameterValue.IllegalExclusterID"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBShardsWithContext(ctx context.Context, request *DescribeDCDBShardsRequest) (response *DescribeDCDBShardsResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBShardsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBShards require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDCDBUpgradePriceRequest() (request *DescribeDCDBUpgradePriceRequest) {
    request = &DescribeDCDBUpgradePriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDCDBUpgradePrice")
    
    
    return
}

func NewDescribeDCDBUpgradePriceResponse() (response *DescribeDCDBUpgradePriceResponse) {
    response = &DescribeDCDBUpgradePriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDCDBUpgradePrice
// 本接口（DescribeDCDBUpgradePrice）用于查询变配分布式数据库实例价格。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SHARDNOTEXIST = "InvalidParameterValue.ShardNotExist"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBUpgradePrice(request *DescribeDCDBUpgradePriceRequest) (response *DescribeDCDBUpgradePriceResponse, err error) {
    return c.DescribeDCDBUpgradePriceWithContext(context.Background(), request)
}

// DescribeDCDBUpgradePrice
// 本接口（DescribeDCDBUpgradePrice）用于查询变配分布式数据库实例价格。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SHARDNOTEXIST = "InvalidParameterValue.ShardNotExist"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDCDBUpgradePriceWithContext(ctx context.Context, request *DescribeDCDBUpgradePriceRequest) (response *DescribeDCDBUpgradePriceResponse, err error) {
    if request == nil {
        request = NewDescribeDCDBUpgradePriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDCDBUpgradePrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDCDBUpgradePriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseObjects")
    
    
    return
}

func NewDescribeDatabaseObjectsResponse() (response *DescribeDatabaseObjectsResponse) {
    response = &DescribeDatabaseObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabaseObjects
// 本接口（DescribeDatabaseObjects）用于查询云数据库实例的数据库中的对象列表，包含表、存储过程、视图和函数。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBOBJECTFAILED = "InternalError.GetDbObjectFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    return c.DescribeDatabaseObjectsWithContext(context.Background(), request)
}

// DescribeDatabaseObjects
// 本接口（DescribeDatabaseObjects）用于查询云数据库实例的数据库中的对象列表，包含表、存储过程、视图和函数。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBOBJECTFAILED = "InternalError.GetDbObjectFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseObjectsWithContext(ctx context.Context, request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseTableRequest() (request *DescribeDatabaseTableRequest) {
    request = &DescribeDatabaseTableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabaseTable")
    
    
    return
}

func NewDescribeDatabaseTableResponse() (response *DescribeDatabaseTableResponse) {
    response = &DescribeDatabaseTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabaseTable
// 本接口（DescribeDatabaseTable）用于查询云数据库实例的表信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETTABLEINFOFAILED = "InternalError.GetTableInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseTable(request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    return c.DescribeDatabaseTableWithContext(context.Background(), request)
}

// DescribeDatabaseTable
// 本接口（DescribeDatabaseTable）用于查询云数据库实例的表信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETTABLEINFOFAILED = "InternalError.GetTableInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabaseTableWithContext(ctx context.Context, request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseTableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabases
// 本接口（DescribeDatabases）用于查询云数据库实例的数据库列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBLISTFAILED = "InternalError.GetDbListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// 本接口（DescribeDatabases）用于查询云数据库实例的数据库列表。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBLISTFAILED = "InternalError.GetDbListFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDcnDetailRequest() (request *DescribeDcnDetailRequest) {
    request = &DescribeDcnDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeDcnDetail")
    
    
    return
}

func NewDescribeDcnDetailResponse() (response *DescribeDcnDetailResponse) {
    response = &DescribeDcnDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDcnDetail
// 获取实例灾备详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDcnDetail(request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
    return c.DescribeDcnDetailWithContext(context.Background(), request)
}

// DescribeDcnDetail
// 获取实例灾备详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeDcnDetailWithContext(ctx context.Context, request *DescribeDcnDetailRequest) (response *DescribeDcnDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDcnDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDcnDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDcnDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileDownloadUrlRequest() (request *DescribeFileDownloadUrlRequest) {
    request = &DescribeFileDownloadUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeFileDownloadUrl")
    
    
    return
}

func NewDescribeFileDownloadUrlResponse() (response *DescribeFileDownloadUrlResponse) {
    response = &DescribeFileDownloadUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileDownloadUrl
// 本接口(DescribeFileDownloadUrl)用于获取数据库指定备份或日志文件的下载连接。
//
// 可能返回的错误码:
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) DescribeFileDownloadUrl(request *DescribeFileDownloadUrlRequest) (response *DescribeFileDownloadUrlResponse, err error) {
    return c.DescribeFileDownloadUrlWithContext(context.Background(), request)
}

// DescribeFileDownloadUrl
// 本接口(DescribeFileDownloadUrl)用于获取数据库指定备份或日志文件的下载连接。
//
// 可能返回的错误码:
//  INTERNALERROR_COSSIGNURL = "InternalError.CosSignUrl"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) DescribeFileDownloadUrlWithContext(ctx context.Context, request *DescribeFileDownloadUrlRequest) (response *DescribeFileDownloadUrlResponse, err error) {
    if request == nil {
        request = NewDescribeFileDownloadUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileDownloadUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileDownloadUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeFlow")
    
    
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlow
// 本接口（DescribeFlow）用于查询流程状态
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_FLOWNOTFOUND = "InvalidParameter.FlowNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    return c.DescribeFlowWithContext(context.Background(), request)
}

// DescribeFlow
// 本接口（DescribeFlow）用于查询流程状态
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_FLOWNOTFOUND = "InvalidParameter.FlowNotFound"
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

func NewDescribeOrdersRequest() (request *DescribeOrdersRequest) {
    request = &DescribeOrdersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeOrders")
    
    
    return
}

func NewDescribeOrdersResponse() (response *DescribeOrdersResponse) {
    response = &DescribeOrdersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOrders
// 本接口（DescribeOrders）用于查询分布式数据库订单信息。传入订单ID来查询订单关联的分布式数据库实例，和对应的任务流程ID。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYORDERFAILED = "InternalError.QueryOrderFailed"
//  INVALIDPARAMETER_DEALNAMENOTGIVEN = "InvalidParameter.DealNameNotGiven"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeOrders(request *DescribeOrdersRequest) (response *DescribeOrdersResponse, err error) {
    return c.DescribeOrdersWithContext(context.Background(), request)
}

// DescribeOrders
// 本接口（DescribeOrders）用于查询分布式数据库订单信息。传入订单ID来查询订单关联的分布式数据库实例，和对应的任务流程ID。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_QUERYORDERFAILED = "InternalError.QueryOrderFailed"
//  INVALIDPARAMETER_DEALNAMENOTGIVEN = "InvalidParameter.DealNameNotGiven"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
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

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjectSecurityGroups")
    
    
    return
}

func NewDescribeProjectSecurityGroupsResponse() (response *DescribeProjectSecurityGroupsResponse) {
    response = &DescribeProjectSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjectSecurityGroups
// 本接口（DescribeProjectSecurityGroups）用于查询项目安全组信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSERSGCOUNTFAILED = "InternalError.GetUserSGCountFailed"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// 本接口（DescribeProjectSecurityGroups）用于查询项目安全组信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_GETUSERSGCOUNTFAILED = "InternalError.GetUserSGCountFailed"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProjects
// 本接口（DescribeProjects）用于查询项目列表
//
// 可能返回的错误码:
//  INTERNALERROR_LISTPROJECTFAILED = "InternalError.ListProjectFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// 本接口（DescribeProjects）用于查询项目列表
//
// 可能返回的错误码:
//  INTERNALERROR_LISTPROJECTFAILED = "InternalError.ListProjectFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeShardSpecRequest() (request *DescribeShardSpecRequest) {
    request = &DescribeShardSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeShardSpec")
    
    
    return
}

func NewDescribeShardSpecResponse() (response *DescribeShardSpecResponse) {
    response = &DescribeShardSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeShardSpec
// 查询可创建的分布式数据库可售卖的分片规格配置。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeShardSpec(request *DescribeShardSpecRequest) (response *DescribeShardSpecResponse, err error) {
    return c.DescribeShardSpecWithContext(context.Background(), request)
}

// DescribeShardSpec
// 查询可创建的分布式数据库可售卖的分片规格配置。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeShardSpecWithContext(ctx context.Context, request *DescribeShardSpecRequest) (response *DescribeShardSpecResponse, err error) {
    if request == nil {
        request = NewDescribeShardSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeShardSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeShardSpecResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSqlLogsRequest() (request *DescribeSqlLogsRequest) {
    request = &DescribeSqlLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeSqlLogs")
    
    
    return
}

func NewDescribeSqlLogsResponse() (response *DescribeSqlLogsResponse) {
    response = &DescribeSqlLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSqlLogs
// 本接口（DescribeSqlLogs）用于获取实例SQL日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MSGQUEUEOPERATIONFAILED = "FailedOperation.MsgQueueOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSqlLogs(request *DescribeSqlLogsRequest) (response *DescribeSqlLogsResponse, err error) {
    return c.DescribeSqlLogsWithContext(context.Background(), request)
}

// DescribeSqlLogs
// 本接口（DescribeSqlLogs）用于获取实例SQL日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MSGQUEUEOPERATIONFAILED = "FailedOperation.MsgQueueOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeSqlLogsWithContext(ctx context.Context, request *DescribeSqlLogsRequest) (response *DescribeSqlLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSqlLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSqlLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSqlLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserTasksRequest() (request *DescribeUserTasksRequest) {
    request = &DescribeUserTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DescribeUserTasks")
    
    
    return
}

func NewDescribeUserTasksResponse() (response *DescribeUserTasksResponse) {
    response = &DescribeUserTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserTasks
// 本接口（DescribeUserTasks）用于拉取用户任务列表
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeUserTasks(request *DescribeUserTasksRequest) (response *DescribeUserTasksResponse, err error) {
    return c.DescribeUserTasksWithContext(context.Background(), request)
}

// DescribeUserTasks
// 本接口（DescribeUserTasks）用于拉取用户任务列表
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeUserTasksWithContext(ctx context.Context, request *DescribeUserTasksRequest) (response *DescribeUserTasksResponse, err error) {
    if request == nil {
        request = NewDescribeUserTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyDCDBInstanceRequest() (request *DestroyDCDBInstanceRequest) {
    request = &DestroyDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DestroyDCDBInstance")
    
    
    return
}

func NewDestroyDCDBInstanceResponse() (response *DestroyDCDBInstanceResponse) {
    response = &DestroyDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyDCDBInstance
// 本接口(DestroyDCDBInstance)用于销毁已隔离的包年包月实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyDCDBInstance(request *DestroyDCDBInstanceRequest) (response *DestroyDCDBInstanceResponse, err error) {
    return c.DestroyDCDBInstanceWithContext(context.Background(), request)
}

// DestroyDCDBInstance
// 本接口(DestroyDCDBInstance)用于销毁已隔离的包年包月实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyDCDBInstanceWithContext(ctx context.Context, request *DestroyDCDBInstanceRequest) (response *DestroyDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyHourDCDBInstanceRequest() (request *DestroyHourDCDBInstanceRequest) {
    request = &DestroyHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DestroyHourDCDBInstance")
    
    
    return
}

func NewDestroyHourDCDBInstanceResponse() (response *DestroyHourDCDBInstanceResponse) {
    response = &DestroyHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyHourDCDBInstance
// 本接口（DestroyHourDCDBInstance）用于销毁按量计费实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyHourDCDBInstance(request *DestroyHourDCDBInstanceRequest) (response *DestroyHourDCDBInstanceResponse, err error) {
    return c.DestroyHourDCDBInstanceWithContext(context.Background(), request)
}

// DestroyHourDCDBInstance
// 本接口（DestroyHourDCDBInstance）用于销毁按量计费实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DestroyHourDCDBInstanceWithContext(ctx context.Context, request *DestroyHourDCDBInstanceRequest) (response *DestroyHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyHourDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyHourDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "DisassociateSecurityGroups")
    
    
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
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_ROUTENOTFOUND = "InternalError.RouteNotFound"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
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

func NewFlushBinlogRequest() (request *FlushBinlogRequest) {
    request = &FlushBinlogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "FlushBinlog")
    
    
    return
}

func NewFlushBinlogResponse() (response *FlushBinlogResponse) {
    response = &FlushBinlogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FlushBinlog
// 相当于在所有分片的mysqld中执行flush logs，完成切分的binlog将展示在各个分片控制台binlog列表里。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOFREQUENTLYCALLED = "LimitExceeded.TooFrequentlyCalled"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) FlushBinlog(request *FlushBinlogRequest) (response *FlushBinlogResponse, err error) {
    return c.FlushBinlogWithContext(context.Background(), request)
}

// FlushBinlog
// 相当于在所有分片的mysqld中执行flush logs，完成切分的binlog将展示在各个分片控制台binlog列表里。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_TOOFREQUENTLYCALLED = "LimitExceeded.TooFrequentlyCalled"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) FlushBinlogWithContext(ctx context.Context, request *FlushBinlogRequest) (response *FlushBinlogResponse, err error) {
    if request == nil {
        request = NewFlushBinlogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FlushBinlog require credential")
    }

    request.SetContext(ctx)
    
    response = NewFlushBinlogResponse()
    err = c.Send(request, response)
    return
}

func NewGrantAccountPrivilegesRequest() (request *GrantAccountPrivilegesRequest) {
    request = &GrantAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "GrantAccountPrivileges")
    
    
    return
}

func NewGrantAccountPrivilegesResponse() (response *GrantAccountPrivilegesResponse) {
    response = &GrantAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GrantAccountPrivileges
// 本接口（GrantAccountPrivileges）用于给云数据库账号赋权。
//
// 注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYRIGHTFAILED = "FailedOperation.ModifyRightFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  INVALIDPARAMETERVALUE_ILLEGALRIGHTPARAM = "InvalidParameterValue.IllegalRightParam"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) GrantAccountPrivileges(request *GrantAccountPrivilegesRequest) (response *GrantAccountPrivilegesResponse, err error) {
    return c.GrantAccountPrivilegesWithContext(context.Background(), request)
}

// GrantAccountPrivileges
// 本接口（GrantAccountPrivileges）用于给云数据库账号赋权。
//
// 注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYRIGHTFAILED = "FailedOperation.ModifyRightFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETRIGHTFAILED = "InternalError.GetRightFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  INVALIDPARAMETERVALUE_ILLEGALRIGHTPARAM = "InvalidParameterValue.IllegalRightParam"
//  INVALIDPARAMETERVALUE_SUPERUSERFORBIDDEN = "InvalidParameterValue.SuperUserForbidden"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
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

func NewInitDCDBInstancesRequest() (request *InitDCDBInstancesRequest) {
    request = &InitDCDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "InitDCDBInstances")
    
    
    return
}

func NewInitDCDBInstancesResponse() (response *InitDCDBInstancesResponse) {
    response = &InitDCDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitDCDBInstances
// 本接口(InitDCDBInstances)用于初始化云数据库实例，包括设置默认字符集、表名大小写敏感等。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ILLEGALINITPARAM = "InvalidParameterValue.IllegalInitParam"
//  INVALIDPARAMETERVALUE_SYNCMODENOTALLOWED = "InvalidParameterValue.SyncModeNotAllowed"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InitDCDBInstances(request *InitDCDBInstancesRequest) (response *InitDCDBInstancesResponse, err error) {
    return c.InitDCDBInstancesWithContext(context.Background(), request)
}

// InitDCDBInstances
// 本接口(InitDCDBInstances)用于初始化云数据库实例，包括设置默认字符集、表名大小写敏感等。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETERVALUE_ILLEGALINITPARAM = "InvalidParameterValue.IllegalInitParam"
//  INVALIDPARAMETERVALUE_SYNCMODENOTALLOWED = "InvalidParameterValue.SyncModeNotAllowed"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) InitDCDBInstancesWithContext(ctx context.Context, request *InitDCDBInstancesRequest) (response *InitDCDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDCDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitDCDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitDCDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDedicatedDBInstanceRequest() (request *IsolateDedicatedDBInstanceRequest) {
    request = &IsolateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "IsolateDedicatedDBInstance")
    
    
    return
}

func NewIsolateDedicatedDBInstanceResponse() (response *IsolateDedicatedDBInstanceResponse) {
    response = &IsolateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateDedicatedDBInstance
// 本接口（IsolateDedicatedDBInstance）用于隔离独享云数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSISOLATEINSTANCEFAILED = "FailedOperation.OssIsolateInstanceFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) IsolateDedicatedDBInstance(request *IsolateDedicatedDBInstanceRequest) (response *IsolateDedicatedDBInstanceResponse, err error) {
    return c.IsolateDedicatedDBInstanceWithContext(context.Background(), request)
}

// IsolateDedicatedDBInstance
// 本接口（IsolateDedicatedDBInstance）用于隔离独享云数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSISOLATEINSTANCEFAILED = "FailedOperation.OssIsolateInstanceFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) IsolateDedicatedDBInstanceWithContext(ctx context.Context, request *IsolateDedicatedDBInstanceRequest) (response *IsolateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDedicatedDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDedicatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateHourDCDBInstanceRequest() (request *IsolateHourDCDBInstanceRequest) {
    request = &IsolateHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "IsolateHourDCDBInstance")
    
    
    return
}

func NewIsolateHourDCDBInstanceResponse() (response *IsolateHourDCDBInstanceResponse) {
    response = &IsolateHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateHourDCDBInstance
// 隔离DCDB后付费实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateHourDCDBInstance(request *IsolateHourDCDBInstanceRequest) (response *IsolateHourDCDBInstanceResponse, err error) {
    return c.IsolateHourDCDBInstanceWithContext(context.Background(), request)
}

// IsolateHourDCDBInstance
// 隔离DCDB后付费实例
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateHourDCDBInstanceWithContext(ctx context.Context, request *IsolateHourDCDBInstanceRequest) (response *IsolateHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateHourDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateHourDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewKillSessionRequest() (request *KillSessionRequest) {
    request = &KillSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "KillSession")
    
    
    return
}

func NewKillSessionResponse() (response *KillSessionResponse) {
    response = &KillSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// KillSession
// 本接口（KillSession）用于杀死指定会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) KillSession(request *KillSessionRequest) (response *KillSessionResponse, err error) {
    return c.KillSessionWithContext(context.Background(), request)
}

// KillSession
// 本接口（KillSession）用于杀死指定会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) KillSessionWithContext(ctx context.Context, request *KillSessionRequest) (response *KillSessionResponse, err error) {
    if request == nil {
        request = NewKillSessionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillSession require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillSessionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAccountDescription")
    
    
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountDescription
// 本接口（ModifyAccountDescription）用于修改云数据库账号备注。
//
// 注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    return c.ModifyAccountDescriptionWithContext(context.Background(), request)
}

// ModifyAccountDescription
// 本接口（ModifyAccountDescription）用于修改云数据库账号备注。
//
// 注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXIST = "ResourceNotFound.AccountDoesNotExist"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyAccountPrivileges")
    
    
    return
}

func NewModifyAccountPrivilegesResponse() (response *ModifyAccountPrivilegesResponse) {
    response = &ModifyAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountPrivileges
// 本接口(ModifyAccountPrivileges)用于修改云数据库的账户的权限信息。
//
// 
//
// **注意**
//
// - 系统保留库："mysql"，只开放["SELECT"]权限
//
// - 只读账号授予读写权限会报错
//
// - 不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// 本接口(ModifyAccountPrivileges)用于修改云数据库的账户的权限信息。
//
// 
//
// **注意**
//
// - 系统保留库："mysql"，只开放["SELECT"]权限
//
// - 只读账号授予读写权限会报错
//
// - 不传该参数表示保留现有权限，如需清除，请在复杂类型Privileges字段传空数组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_BADUSERRIGHT = "InvalidParameterValue.BadUserRight"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountPrivilegesWithContext(ctx context.Context, request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyAccountPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBEncryptAttributesRequest() (request *ModifyDBEncryptAttributesRequest) {
    request = &ModifyDBEncryptAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBEncryptAttributes")
    
    
    return
}

func NewModifyDBEncryptAttributesResponse() (response *ModifyDBEncryptAttributesResponse) {
    response = &ModifyDBEncryptAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBEncryptAttributes
// 本接口(ModifyDBEncryptAttributes)用于修改实例数据加密。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBEncryptAttributes(request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    return c.ModifyDBEncryptAttributesWithContext(context.Background(), request)
}

// ModifyDBEncryptAttributes
// 本接口(ModifyDBEncryptAttributes)用于修改实例数据加密。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OSSOPERATIONFAILED = "FailedOperation.OssOperationFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_BADINSTANCESTATUS = "ResourceUnavailable.BadInstanceStatus"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBEncryptAttributesWithContext(ctx context.Context, request *ModifyDBEncryptAttributesRequest) (response *ModifyDBEncryptAttributesResponse, err error) {
    if request == nil {
        request = NewModifyDBEncryptAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBEncryptAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBEncryptAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceName")
    
    
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceName
// 本接口（ModifyDBInstanceName）用于修改实例名字
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    return c.ModifyDBInstanceNameWithContext(context.Background(), request)
}

// ModifyDBInstanceName
// 本接口（ModifyDBInstanceName）用于修改实例名字
//
// 可能返回的错误码:
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
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

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
    response = &ModifyDBInstanceSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceSecurityGroups
// 本接口（ModifyDBInstanceSecurityGroups）用于修改云数据库安全组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// 本接口（ModifyDBInstanceSecurityGroups）用于修改云数据库安全组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_AUTHNOSTRATEGY = "FailedOperation.AuthNoStrategy"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_GETSECURITYGROUPDETAILFAILED = "InternalError.GetSecurityGroupDetailFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_INSTANCESGOVERLIMITERROR = "InternalError.InstanceSGOverLimitError"
//  INTERNALERROR_LISTINSTANCERESPRESOURCECOUNTNOTMATCHERROR = "InternalError.ListInstanceRespResourceCountNotMatchError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_QUERYDATABASEFAILED = "InternalError.QueryDatabaseFailed"
//  INTERNALERROR_READDATABASEFAILED = "InternalError.ReadDatabaseFailed"
//  INTERNALERROR_SETSVCLOCATIONFAILED = "InternalError.SetSvcLocationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_PRODUCTCONFIGNOTEXISTEDERROR = "ResourceNotFound.ProductConfigNotExistedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_SGCHECKFAIL = "ResourceUnavailable.SGCheckFail"
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

func NewModifyDBInstancesProjectRequest() (request *ModifyDBInstancesProjectRequest) {
    request = &ModifyDBInstancesProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBInstancesProject")
    
    
    return
}

func NewModifyDBInstancesProjectResponse() (response *ModifyDBInstancesProjectResponse) {
    response = &ModifyDBInstancesProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstancesProject
// 本接口（ModifyDBInstancesProject）用于修改云数据库实例所属项目。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstancesProject(request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    return c.ModifyDBInstancesProjectWithContext(context.Background(), request)
}

// ModifyDBInstancesProject
// 本接口（ModifyDBInstancesProject）用于修改云数据库实例所属项目。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBInstancesProjectWithContext(ctx context.Context, request *ModifyDBInstancesProjectRequest) (response *ModifyDBInstancesProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstancesProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstancesProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstancesProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
    request = &ModifyDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBParameters")
    
    
    return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
    response = &ModifyDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBParameters
// 本接口(ModifyDBParameters)用于修改数据库参数。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    return c.ModifyDBParametersWithContext(context.Background(), request)
}

// ModifyDBParameters
// 本接口(ModifyDBParameters)用于修改数据库参数。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETDBCONFIGFAILED = "InternalError.GetDbConfigFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyDBParametersWithContext(ctx context.Context, request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBSyncModeRequest() (request *ModifyDBSyncModeRequest) {
    request = &ModifyDBSyncModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyDBSyncMode")
    
    
    return
}

func NewModifyDBSyncModeResponse() (response *ModifyDBSyncModeResponse) {
    response = &ModifyDBSyncModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBSyncMode
// 本接口（ModifyDBSyncMode）用于修改云数据库实例的同步模式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_BADSYNCMODE = "InvalidParameterValue.BadSyncMode"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBSyncMode(request *ModifyDBSyncModeRequest) (response *ModifyDBSyncModeResponse, err error) {
    return c.ModifyDBSyncModeWithContext(context.Background(), request)
}

// ModifyDBSyncMode
// 本接口（ModifyDBSyncMode）用于修改云数据库实例的同步模式。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_GETINSTANCEINFOFAILED = "InternalError.GetInstanceInfoFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_BADSYNCMODE = "InvalidParameterValue.BadSyncMode"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) ModifyDBSyncModeWithContext(ctx context.Context, request *ModifyDBSyncModeRequest) (response *ModifyDBSyncModeResponse, err error) {
    if request == nil {
        request = NewModifyDBSyncModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBSyncMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBSyncModeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNetworkRequest() (request *ModifyInstanceNetworkRequest) {
    request = &ModifyInstanceNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceNetwork")
    
    
    return
}

func NewModifyInstanceNetworkResponse() (response *ModifyInstanceNetworkResponse) {
    response = &ModifyInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceNetwork
// 本接口（ModifyInstanceNetwork）用于修改实例所属网络。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_WANSTATUSABNORMAL = "FailedOperation.WanStatusAbnormal"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_VPCOPERATIONFAILED = "InternalError.VpcOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNetwork(request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    return c.ModifyInstanceNetworkWithContext(context.Background(), request)
}

// ModifyInstanceNetwork
// 本接口（ModifyInstanceNetwork）用于修改实例所属网络。
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_WANSTATUSABNORMAL = "FailedOperation.WanStatusAbnormal"
//  INTERNALERROR_CHECKVIPSTATUSFAILED = "InternalError.CheckVipStatusFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETSUBNETFAILED = "InternalError.GetSubnetFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_VPCOPERATIONFAILED = "InternalError.VpcOperationFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNetworkWithContext(ctx context.Context, request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNetworkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceVipRequest() (request *ModifyInstanceVipRequest) {
    request = &ModifyInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceVip")
    
    
    return
}

func NewModifyInstanceVipResponse() (response *ModifyInstanceVipResponse) {
    response = &ModifyInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceVip
// 本接口（ModifyInstanceVip）用于修改实例Vip
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_VIPNOTCHANGE = "FailedOperation.VipNotChange"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVip(request *ModifyInstanceVipRequest) (response *ModifyInstanceVipResponse, err error) {
    return c.ModifyInstanceVipWithContext(context.Background(), request)
}

// ModifyInstanceVip
// 本接口（ModifyInstanceVip）用于修改实例Vip
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_VIPNOTCHANGE = "FailedOperation.VipNotChange"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VIPNOTINSUBNET = "InvalidParameter.VipNotInSubnet"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVipWithContext(ctx context.Context, request *ModifyInstanceVipRequest) (response *ModifyInstanceVipResponse, err error) {
    if request == nil {
        request = NewModifyInstanceVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceVipResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceVportRequest() (request *ModifyInstanceVportRequest) {
    request = &ModifyInstanceVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyInstanceVport")
    
    
    return
}

func NewModifyInstanceVportResponse() (response *ModifyInstanceVportResponse) {
    response = &ModifyInstanceVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceVport
// 本接口（ModifyInstanceVport）用于修改实例VPORT
//
// 可能返回的错误码:
//  FAILEDOPERATION_SGCHANGEVIP = "FailedOperation.SGChangeVip"
//  FAILEDOPERATION_VPCADDSERVICEFAILED = "FailedOperation.VpcAddServiceFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  INVALIDPARAMETER_VPORTUSED = "InvalidParameter.VportUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVport(request *ModifyInstanceVportRequest) (response *ModifyInstanceVportResponse, err error) {
    return c.ModifyInstanceVportWithContext(context.Background(), request)
}

// ModifyInstanceVport
// 本接口（ModifyInstanceVport）用于修改实例VPORT
//
// 可能返回的错误码:
//  FAILEDOPERATION_SGCHANGEVIP = "FailedOperation.SGChangeVip"
//  FAILEDOPERATION_VPCADDSERVICEFAILED = "FailedOperation.VpcAddServiceFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INNERSYSTEMERROR = "InternalError.InnerSystemError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_VIPUSED = "InvalidParameter.VipUsed"
//  INVALIDPARAMETER_VPORTUSED = "InvalidParameter.VportUsed"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceVportWithContext(ctx context.Context, request *ModifyInstanceVportRequest) (response *ModifyInstanceVportResponse, err error) {
    if request == nil {
        request = NewModifyInstanceVportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRealServerAccessStrategyRequest() (request *ModifyRealServerAccessStrategyRequest) {
    request = &ModifyRealServerAccessStrategyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ModifyRealServerAccessStrategy")
    
    
    return
}

func NewModifyRealServerAccessStrategyResponse() (response *ModifyRealServerAccessStrategyResponse) {
    response = &ModifyRealServerAccessStrategyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRealServerAccessStrategy
// 本接口(ModifyRealServerAccessStrategy)用于修改云数据库的VPCGW到RS的访问策略。
//
// 
//
// **注意**
//
// - 修改策略后只对新建立的连接生效，老连接不受影响
//
// - 就近访问只针对实例是跨可用区部署有用，单可用区部署实例就近与否并无作用
//
// - DB每个Node对应一个proxy，如果开启就近访问，将会把连接集中到对应可用区的proxy上，可能造成热点问题，这种情况下如果是线上业务，请务必根据自己的业务请求量测试符合预期后再进行就近策略变更
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_VPCUPDATEROUTEFAILED = "FailedOperation.VpcUpdateRouteFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyRealServerAccessStrategy(request *ModifyRealServerAccessStrategyRequest) (response *ModifyRealServerAccessStrategyResponse, err error) {
    return c.ModifyRealServerAccessStrategyWithContext(context.Background(), request)
}

// ModifyRealServerAccessStrategy
// 本接口(ModifyRealServerAccessStrategy)用于修改云数据库的VPCGW到RS的访问策略。
//
// 
//
// **注意**
//
// - 修改策略后只对新建立的连接生效，老连接不受影响
//
// - 就近访问只针对实例是跨可用区部署有用，单可用区部署实例就近与否并无作用
//
// - DB每个Node对应一个proxy，如果开启就近访问，将会把连接集中到对应可用区的proxy上，可能造成热点问题，这种情况下如果是线上业务，请务必根据自己的业务请求量测试符合预期后再进行就近策略变更
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_VPCUPDATEROUTEFAILED = "FailedOperation.VpcUpdateRouteFailed"
//  INTERNALERROR_GETVPCFAILED = "InternalError.GetVpcFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_SUBNETUNAVAILABLE = "InvalidParameter.SubnetUnavailable"
//  INVALIDPARAMETER_VPCNOTFOUND = "InvalidParameter.VpcNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyRealServerAccessStrategyWithContext(ctx context.Context, request *ModifyRealServerAccessStrategyRequest) (response *ModifyRealServerAccessStrategyResponse, err error) {
    if request == nil {
        request = NewModifyRealServerAccessStrategyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRealServerAccessStrategy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRealServerAccessStrategyResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBExtranetAccessRequest() (request *OpenDBExtranetAccessRequest) {
    request = &OpenDBExtranetAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "OpenDBExtranetAccess")
    
    
    return
}

func NewOpenDBExtranetAccessResponse() (response *OpenDBExtranetAccessResponse) {
    response = &OpenDBExtranetAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenDBExtranetAccess
// 本接口（OpenDBExtranetAccess）用于开通云数据库实例的外网访问。开通外网访问后，您可通过外网域名和端口访问实例，可使用查询实例列表接口获取外网域名和端口信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) OpenDBExtranetAccess(request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    return c.OpenDBExtranetAccessWithContext(context.Background(), request)
}

// OpenDBExtranetAccess
// 本接口（OpenDBExtranetAccess）用于开通云数据库实例的外网访问。开通外网访问后，您可通过外网域名和端口访问实例，可使用查询实例列表接口获取外网域名和端口信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_CREATEFLOWFAILED = "InternalError.CreateFlowFailed"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
func (c *Client) OpenDBExtranetAccessWithContext(ctx context.Context, request *OpenDBExtranetAccessRequest) (response *OpenDBExtranetAccessResponse, err error) {
    if request == nil {
        request = NewOpenDBExtranetAccessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenDBExtranetAccess require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenDBExtranetAccessResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDCDBInstanceRequest() (request *RenewDCDBInstanceRequest) {
    request = &RenewDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "RenewDCDBInstance")
    
    
    return
}

func NewRenewDCDBInstanceResponse() (response *RenewDCDBInstanceResponse) {
    response = &RenewDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewDCDBInstance
// 本接口（RenewDCDBInstance）用于续费分布式数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_NOTSUPPORTEDPAYMODE = "InvalidParameter.NotSupportedPayMode"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewDCDBInstance(request *RenewDCDBInstanceRequest) (response *RenewDCDBInstanceResponse, err error) {
    return c.RenewDCDBInstanceWithContext(context.Background(), request)
}

// RenewDCDBInstance
// 本接口（RenewDCDBInstance）用于续费分布式数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_NOTSUPPORTEDPAYMODE = "InvalidParameter.NotSupportedPayMode"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RenewDCDBInstanceWithContext(ctx context.Context, request *RenewDCDBInstanceRequest) (response *RenewDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewRenewDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于重置云数据库账号的密码。
//
// 注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESETPASSWORDFAILED = "FailedOperation.ResetPasswordFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    return c.ResetAccountPasswordWithContext(context.Background(), request)
}

// ResetAccountPassword
// 本接口（ResetAccountPassword）用于重置云数据库账号的密码。
//
// 注意：相同用户名，不同Host是不同的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESETPASSWORDFAILED = "FailedOperation.ResetPasswordFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_CHARACTERERROR = "InvalidParameter.CharacterError"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
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

func NewSwitchDBInstanceHARequest() (request *SwitchDBInstanceHARequest) {
    request = &SwitchDBInstanceHARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "SwitchDBInstanceHA")
    
    
    return
}

func NewSwitchDBInstanceHAResponse() (response *SwitchDBInstanceHAResponse) {
    response = &SwitchDBInstanceHAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDBInstanceHA
// 本接口(SwitchDBInstanceHA)用于实例主备切换。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_ZONEIDILLEGAL = "InvalidParameter.ZoneIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchDBInstanceHA(request *SwitchDBInstanceHARequest) (response *SwitchDBInstanceHAResponse, err error) {
    return c.SwitchDBInstanceHAWithContext(context.Background(), request)
}

// SwitchDBInstanceHA
// 本接口(SwitchDBInstanceHA)用于实例主备切换。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_GETINSTANCEDETAILFAILED = "InternalError.GetInstanceDetailFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_ZONEIDILLEGAL = "InvalidParameter.ZoneIdIllegal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) SwitchDBInstanceHAWithContext(ctx context.Context, request *SwitchDBInstanceHARequest) (response *SwitchDBInstanceHAResponse, err error) {
    if request == nil {
        request = NewSwitchDBInstanceHARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDBInstanceHA require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDBInstanceHAResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDedicatedDBInstanceRequest() (request *TerminateDedicatedDBInstanceRequest) {
    request = &TerminateDedicatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "TerminateDedicatedDBInstance")
    
    
    return
}

func NewTerminateDedicatedDBInstanceResponse() (response *TerminateDedicatedDBInstanceResponse) {
    response = &TerminateDedicatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateDedicatedDBInstance
// 本接口（TerminateDedicatedDBInstance）用于销毁已隔离的独享分布式数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDedicatedDBInstance(request *TerminateDedicatedDBInstanceRequest) (response *TerminateDedicatedDBInstanceResponse, err error) {
    return c.TerminateDedicatedDBInstanceWithContext(context.Background(), request)
}

// TerminateDedicatedDBInstance
// 本接口（TerminateDedicatedDBInstance）用于销毁已隔离的独享分布式数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEFLOWFAILED = "FailedOperation.CreateFlowFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCEALREADYDELETED = "ResourceUnavailable.InstanceAlreadyDeleted"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TerminateDedicatedDBInstanceWithContext(ctx context.Context, request *TerminateDedicatedDBInstanceRequest) (response *TerminateDedicatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateDedicatedDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDedicatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDedicatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDCDBInstanceRequest() (request *UpgradeDCDBInstanceRequest) {
    request = &UpgradeDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeDCDBInstance")
    
    
    return
}

func NewUpgradeDCDBInstanceResponse() (response *UpgradeDCDBInstanceResponse) {
    response = &UpgradeDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeDCDBInstance
// 本接口（UpgradeDCDBInstance）用于升级分布式数据库实例。本接口完成下单和支付两个动作，如果发生支付失败的错误，调用用户账户相关接口中的支付订单接口（PayDeals）重新支付即可。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_NOTSUPPORTEDPAYMODE = "InvalidParameter.NotSupportedPayMode"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDCDBInstance(request *UpgradeDCDBInstanceRequest) (response *UpgradeDCDBInstanceResponse, err error) {
    return c.UpgradeDCDBInstanceWithContext(context.Background(), request)
}

// UpgradeDCDBInstance
// 本接口（UpgradeDCDBInstance）用于升级分布式数据库实例。本接口完成下单和支付两个动作，如果发生支付失败的错误，调用用户账户相关接口中的支付订单接口（PayDeals）重新支付即可。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  INTERNALERROR_CAMAUTHFAILED = "InternalError.CamAuthFailed"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INVALIDPARAMETER_GENERICPARAMETERERROR = "InvalidParameter.GenericParameterError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_NOTSUPPORTEDPAYMODE = "InvalidParameter.NotSupportedPayMode"
//  INVALIDPARAMETER_SPECNOTFOUND = "InvalidParameter.SpecNotFound"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCENOTFOUND_NOINSTANCEFOUND = "ResourceNotFound.NoInstanceFound"
//  RESOURCEUNAVAILABLE_INSTANCEHASBEENLOCKED = "ResourceUnavailable.InstanceHasBeenLocked"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeDCDBInstanceWithContext(ctx context.Context, request *UpgradeDCDBInstanceRequest) (response *UpgradeDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeHourDCDBInstanceRequest() (request *UpgradeHourDCDBInstanceRequest) {
    request = &UpgradeHourDCDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dcdb", APIVersion, "UpgradeHourDCDBInstance")
    
    
    return
}

func NewUpgradeHourDCDBInstanceResponse() (response *UpgradeHourDCDBInstanceResponse) {
    response = &UpgradeHourDCDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeHourDCDBInstance
// 本接口（UpgradeHourDCDBInstance）用于升级后付费分布式数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_CREATEFLOWERROR = "InternalError.CreateFlowError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_NOTSUPPORTEDPAYMODE = "InvalidParameter.NotSupportedPayMode"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeHourDCDBInstance(request *UpgradeHourDCDBInstanceRequest) (response *UpgradeHourDCDBInstanceResponse, err error) {
    return c.UpgradeHourDCDBInstanceWithContext(context.Background(), request)
}

// UpgradeHourDCDBInstance
// 本接口（UpgradeHourDCDBInstance）用于升级后付费分布式数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEORDERFAILED = "FailedOperation.CreateOrderFailed"
//  INTERNALERROR_CREATEFLOWERROR = "InternalError.CreateFlowError"
//  INTERNALERROR_OPERATEDATABASEFAILED = "InternalError.OperateDatabaseFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_NOTSUPPORTEDPAYMODE = "InvalidParameter.NotSupportedPayMode"
//  INVALIDPARAMETERVALUE_SPECIDILLEGAL = "InvalidParameterValue.SpecIdIllegal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) UpgradeHourDCDBInstanceWithContext(ctx context.Context, request *UpgradeHourDCDBInstanceRequest) (response *UpgradeHourDCDBInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeHourDCDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeHourDCDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeHourDCDBInstanceResponse()
    err = c.Send(request, response)
    return
}
