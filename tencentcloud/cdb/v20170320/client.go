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

package v20170320

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-20"

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


func NewAddTimeWindowRequest() (request *AddTimeWindowRequest) {
    request = &AddTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "AddTimeWindow")
    
    
    return
}

func NewAddTimeWindowResponse() (response *AddTimeWindowResponse) {
    response = &AddTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddTimeWindow
// 本接口(AddTimeWindow)用于添加云数据库实例的维护时间窗口，以指定实例在哪些时间段可以自动执行切换访问操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) AddTimeWindow(request *AddTimeWindowRequest) (response *AddTimeWindowResponse, err error) {
    return c.AddTimeWindowWithContext(context.Background(), request)
}

// AddTimeWindow
// 本接口(AddTimeWindow)用于添加云数据库实例的维护时间窗口，以指定实例在哪些时间段可以自动执行切换访问操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) AddTimeWindowWithContext(ctx context.Context, request *AddTimeWindowRequest) (response *AddTimeWindowResponse, err error) {
    if request == nil {
        request = NewAddTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTimeWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "AssociateSecurityGroups")
    
    
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
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// 本接口(AssociateSecurityGroups)用于安全组批量绑定实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
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

func NewBalanceRoGroupLoadRequest() (request *BalanceRoGroupLoadRequest) {
    request = &BalanceRoGroupLoadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "BalanceRoGroupLoad")
    
    
    return
}

func NewBalanceRoGroupLoadResponse() (response *BalanceRoGroupLoadResponse) {
    response = &BalanceRoGroupLoadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BalanceRoGroupLoad
// 本接口(BalanceRoGroupLoad)用于重新均衡 RO 组内实例的负载。注意，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库，谨慎操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) BalanceRoGroupLoad(request *BalanceRoGroupLoadRequest) (response *BalanceRoGroupLoadResponse, err error) {
    return c.BalanceRoGroupLoadWithContext(context.Background(), request)
}

// BalanceRoGroupLoad
// 本接口(BalanceRoGroupLoad)用于重新均衡 RO 组内实例的负载。注意，RO 组内 RO 实例会有一次数据库连接瞬断，请确保应用程序能重连数据库，谨慎操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) BalanceRoGroupLoadWithContext(ctx context.Context, request *BalanceRoGroupLoadRequest) (response *BalanceRoGroupLoadResponse, err error) {
    if request == nil {
        request = NewBalanceRoGroupLoadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BalanceRoGroupLoad require credential")
    }

    request.SetContext(ctx)
    
    response = NewBalanceRoGroupLoadResponse()
    err = c.Send(request, response)
    return
}

func NewCloseCDBProxyRequest() (request *CloseCDBProxyRequest) {
    request = &CloseCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CloseCDBProxy")
    
    
    return
}

func NewCloseCDBProxyResponse() (response *CloseCDBProxyResponse) {
    response = &CloseCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseCDBProxy
// 关闭数据库代理
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CloseCDBProxy(request *CloseCDBProxyRequest) (response *CloseCDBProxyResponse, err error) {
    return c.CloseCDBProxyWithContext(context.Background(), request)
}

// CloseCDBProxy
// 关闭数据库代理
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CloseCDBProxyWithContext(ctx context.Context, request *CloseCDBProxyRequest) (response *CloseCDBProxyResponse, err error) {
    if request == nil {
        request = NewCloseCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCloseWanServiceRequest() (request *CloseWanServiceRequest) {
    request = &CloseWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CloseWanService")
    
    
    return
}

func NewCloseWanServiceResponse() (response *CloseWanServiceResponse) {
    response = &CloseWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseWanService
// 本接口(CloseWanService)用于关闭云数据库实例的外网访问。关闭外网访问后，外网地址将不可访问。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CloseWanService(request *CloseWanServiceRequest) (response *CloseWanServiceResponse, err error) {
    return c.CloseWanServiceWithContext(context.Background(), request)
}

// CloseWanService
// 本接口(CloseWanService)用于关闭云数据库实例的外网访问。关闭外网访问后，外网地址将不可访问。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CloseWanServiceWithContext(ctx context.Context, request *CloseWanServiceRequest) (response *CloseWanServiceResponse, err error) {
    if request == nil {
        request = NewCloseWanServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseWanService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountsRequest() (request *CreateAccountsRequest) {
    request = &CreateAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAccounts")
    
    
    return
}

func NewCreateAccountsResponse() (response *CreateAccountsResponse) {
    response = &CreateAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAccounts
// 本接口(CreateAccounts)用于创建云数据库的账户，需要指定新的账户名和域名，以及所对应的密码，同时可以设置账号的备注信息以及最大可用连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) CreateAccounts(request *CreateAccountsRequest) (response *CreateAccountsResponse, err error) {
    return c.CreateAccountsWithContext(context.Background(), request)
}

// CreateAccounts
// 本接口(CreateAccounts)用于创建云数据库的账户，需要指定新的账户名和域名，以及所对应的密码，同时可以设置账号的备注信息以及最大可用连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAuditLogFile")
    
    
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
//  FAILEDOPERATION_CREATEAUDITLOGFILEERROR = "FailedOperation.CreateAuditLogFileError"
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITFILEOVERQUOTAERROR = "OperationDenied.AuditFileOverQuotaError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_OVERQUOTA = "OperationDenied.OverQuota"
//  OVERQUOTA = "OverQuota"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateAuditLogFile(request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    return c.CreateAuditLogFileWithContext(context.Background(), request)
}

// CreateAuditLogFile
// 本接口(CreateAuditLogFile)用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITLOGFILEERROR = "FailedOperation.CreateAuditLogFileError"
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITFILEOVERQUOTAERROR = "OperationDenied.AuditFileOverQuotaError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_OVERQUOTA = "OperationDenied.OverQuota"
//  OVERQUOTA = "OverQuota"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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

func NewCreateAuditPolicyRequest() (request *CreateAuditPolicyRequest) {
    request = &CreateAuditPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAuditPolicy")
    
    
    return
}

func NewCreateAuditPolicyResponse() (response *CreateAuditPolicyResponse) {
    response = &CreateAuditPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAuditPolicy
// 本接口(CreateAuditPolicy)用于创建云数据库实例的审计策略，即将审计规则绑定到具体的云数据库实例上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  FAILEDOPERATION_TYPEINCONFLICT = "FailedOperation.TypeInConflict"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITOSSLOGICERROR = "InternalError.AuditOssLogicError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITRULEHASBIND = "OperationDenied.AuditRuleHasBind"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_DBBRAINPOLICYCONFLICT = "OperationDenied.DBBrainPolicyConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAuditPolicy(request *CreateAuditPolicyRequest) (response *CreateAuditPolicyResponse, err error) {
    return c.CreateAuditPolicyWithContext(context.Background(), request)
}

// CreateAuditPolicy
// 本接口(CreateAuditPolicy)用于创建云数据库实例的审计策略，即将审计规则绑定到具体的云数据库实例上。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  FAILEDOPERATION_TYPEINCONFLICT = "FailedOperation.TypeInConflict"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITOSSLOGICERROR = "InternalError.AuditOssLogicError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITRULEHASBIND = "OperationDenied.AuditRuleHasBind"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_DBBRAINPOLICYCONFLICT = "OperationDenied.DBBrainPolicyConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAuditPolicyWithContext(ctx context.Context, request *CreateAuditPolicyRequest) (response *CreateAuditPolicyResponse, err error) {
    if request == nil {
        request = NewCreateAuditPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditRuleRequest() (request *CreateAuditRuleRequest) {
    request = &CreateAuditRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAuditRule")
    
    
    return
}

func NewCreateAuditRuleResponse() (response *CreateAuditRuleResponse) {
    response = &CreateAuditRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAuditRule
// 本接口(CreateAuditRule)用于创建用户在当前地域的审计规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITRULEEXISTERROR = "OperationDenied.AuditRuleExistError"
func (c *Client) CreateAuditRule(request *CreateAuditRuleRequest) (response *CreateAuditRuleResponse, err error) {
    return c.CreateAuditRuleWithContext(context.Background(), request)
}

// CreateAuditRule
// 本接口(CreateAuditRule)用于创建用户在当前地域的审计规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_SERVICEERROR = "InternalError.ServiceError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITRULEEXISTERROR = "OperationDenied.AuditRuleExistError"
func (c *Client) CreateAuditRuleWithContext(ctx context.Context, request *CreateAuditRuleRequest) (response *CreateAuditRuleResponse, err error) {
    if request == nil {
        request = NewCreateAuditRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
    request = &CreateBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateBackup")
    
    
    return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
    response = &CreateBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBackup
// 本接口(CreateBackup)用于创建数据库备份。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OVERQUOTA = "OverQuota"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// 本接口(CreateBackup)用于创建数据库备份。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OVERQUOTA = "OverQuota"
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

func NewCreateCloneInstanceRequest() (request *CreateCloneInstanceRequest) {
    request = &CreateCloneInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateCloneInstance")
    
    
    return
}

func NewCreateCloneInstanceResponse() (response *CreateCloneInstanceResponse) {
    response = &CreateCloneInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCloneInstance
// 本接口(CreateCloneInstance) 用于从目标源实例创建一个克隆实例，可以指定克隆实例回档到源实例的指定物理备份文件或者指定的回档时间点。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateCloneInstance(request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    return c.CreateCloneInstanceWithContext(context.Background(), request)
}

// CreateCloneInstance
// 本接口(CreateCloneInstance) 用于从目标源实例创建一个克隆实例，可以指定克隆实例回档到源实例的指定物理备份文件或者指定的回档时间点。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateCloneInstanceWithContext(ctx context.Context, request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    if request == nil {
        request = NewCreateCloneInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloneInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloneInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBImportJobRequest() (request *CreateDBImportJobRequest) {
    request = &CreateDBImportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBImportJob")
    
    
    return
}

func NewCreateDBImportJobResponse() (response *CreateDBImportJobResponse) {
    response = &CreateDBImportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDBImportJob
// 本接口(CreateDBImportJob)用于创建云数据库数据导入任务。
//
// 
//
// 注意，用户进行数据导入任务的文件，必须提前上传到腾讯云。用户须在控制台进行文件导入。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_EXECUTESQLERROR = "InternalError.ExecuteSQLError"
//  INTERNALERROR_IMPORTERROR = "InternalError.ImportError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CreateDBImportJob(request *CreateDBImportJobRequest) (response *CreateDBImportJobResponse, err error) {
    return c.CreateDBImportJobWithContext(context.Background(), request)
}

// CreateDBImportJob
// 本接口(CreateDBImportJob)用于创建云数据库数据导入任务。
//
// 
//
// 注意，用户进行数据导入任务的文件，必须提前上传到腾讯云。用户须在控制台进行文件导入。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_EXECUTESQLERROR = "InternalError.ExecuteSQLError"
//  INTERNALERROR_IMPORTERROR = "InternalError.ImportError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) CreateDBImportJobWithContext(ctx context.Context, request *CreateDBImportJobRequest) (response *CreateDBImportJobResponse, err error) {
    if request == nil {
        request = NewCreateDBImportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBImportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBInstance")
    
    
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDBInstance
// 本接口(CreateDBInstance)用于创建包年包月的云数据库实例（包括主实例、灾备实例和只读实例），可通过传入实例规格、MySQL 版本号、购买时长和数量等信息创建云数据库实例。
//
// 
//
// 该接口为异步接口，您还可以使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询该实例的详细信息。当该实例的 Status 为1，且 TaskStatus 为0，表示实例已经发货成功。
//
// 
//
// 1. 首先请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口查询可创建的实例规格信息，然后请使用 [查询数据库价格](https://cloud.tencent.com/document/api/236/18566) 接口查询可创建实例的售卖价格；
//
// 2. 单次创建实例最大支持 100 个，实例时长最大支持 36 个月；
//
// 3. 支持创建 MySQL 5.5 、 MySQL 5.6 、 MySQL 5.7 、 MySQL 8.0 版本；
//
// 4. 支持创建主实例、只读实例、灾备实例；
//
// 5. 当入参指定 Port，ParamList 或 Password 时，该实例会进行初始化操作（不支持基础版实例）；
//
// 6. 当入参指定 ParamTemplateId 或 AlarmPolicyList 时，需将SDK提升至最新版本方可支持；
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    return c.CreateDBInstanceWithContext(context.Background(), request)
}

// CreateDBInstance
// 本接口(CreateDBInstance)用于创建包年包月的云数据库实例（包括主实例、灾备实例和只读实例），可通过传入实例规格、MySQL 版本号、购买时长和数量等信息创建云数据库实例。
//
// 
//
// 该接口为异步接口，您还可以使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询该实例的详细信息。当该实例的 Status 为1，且 TaskStatus 为0，表示实例已经发货成功。
//
// 
//
// 1. 首先请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口查询可创建的实例规格信息，然后请使用 [查询数据库价格](https://cloud.tencent.com/document/api/236/18566) 接口查询可创建实例的售卖价格；
//
// 2. 单次创建实例最大支持 100 个，实例时长最大支持 36 个月；
//
// 3. 支持创建 MySQL 5.5 、 MySQL 5.6 、 MySQL 5.7 、 MySQL 8.0 版本；
//
// 4. 支持创建主实例、只读实例、灾备实例；
//
// 5. 当入参指定 Port，ParamList 或 Password 时，该实例会进行初始化操作（不支持基础版实例）；
//
// 6. 当入参指定 ParamTemplateId 或 AlarmPolicyList 时，需将SDK提升至最新版本方可支持；
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceHourRequest() (request *CreateDBInstanceHourRequest) {
    request = &CreateDBInstanceHourRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDBInstanceHour")
    
    
    return
}

func NewCreateDBInstanceHourResponse() (response *CreateDBInstanceHourResponse) {
    response = &CreateDBInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDBInstanceHour
// 本接口(CreateDBInstanceHour)用于创建按量计费的实例，可通过传入实例规格、MySQL 版本号和数量等信息创建云数据库实例，支持主实例、灾备实例和只读实例的创建。
//
// 
//
// 该接口为异步接口，您还可以使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询该实例的详细信息。当该实例的 Status 为 1，且 TaskStatus 为 0，表示实例已经发货成功。
//
// 
//
// 1. 首先请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口查询可创建的实例规格信息，然后请使用 [查询数据库价格](https://cloud.tencent.com/document/api/236/18566) 接口查询可创建实例的售卖价格；
//
// 2. 单次创建实例最大支持 100 个，实例时长最大支持 36 个月；
//
// 3. 支持创建 MySQL 5.5、MySQL 5.6 、MySQL 5.7 和 MySQL 8.0 版本；
//
// 4. 支持创建主实例、灾备实例和只读实例；
//
// 5. 当入参指定 Port，ParamList 或 Password 时，该实例会进行初始化操作；
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    return c.CreateDBInstanceHourWithContext(context.Background(), request)
}

// CreateDBInstanceHour
// 本接口(CreateDBInstanceHour)用于创建按量计费的实例，可通过传入实例规格、MySQL 版本号和数量等信息创建云数据库实例，支持主实例、灾备实例和只读实例的创建。
//
// 
//
// 该接口为异步接口，您还可以使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询该实例的详细信息。当该实例的 Status 为 1，且 TaskStatus 为 0，表示实例已经发货成功。
//
// 
//
// 1. 首先请使用 [获取云数据库可售卖规格](https://cloud.tencent.com/document/api/236/17229) 接口查询可创建的实例规格信息，然后请使用 [查询数据库价格](https://cloud.tencent.com/document/api/236/18566) 接口查询可创建实例的售卖价格；
//
// 2. 单次创建实例最大支持 100 个，实例时长最大支持 36 个月；
//
// 3. 支持创建 MySQL 5.5、MySQL 5.6 、MySQL 5.7 和 MySQL 8.0 版本；
//
// 4. 支持创建主实例、灾备实例和只读实例；
//
// 5. 当入参指定 Port，ParamList 或 Password 时，该实例会进行初始化操作；
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceHourWithContext(ctx context.Context, request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceHour require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDeployGroupRequest() (request *CreateDeployGroupRequest) {
    request = &CreateDeployGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDeployGroup")
    
    
    return
}

func NewCreateDeployGroupResponse() (response *CreateDeployGroupResponse) {
    response = &CreateDeployGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDeployGroup
// 本接口(CreateDeployGroup)用于创建放置实例的置放群组
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) CreateDeployGroup(request *CreateDeployGroupRequest) (response *CreateDeployGroupResponse, err error) {
    return c.CreateDeployGroupWithContext(context.Background(), request)
}

// CreateDeployGroup
// 本接口(CreateDeployGroup)用于创建放置实例的置放群组
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) CreateDeployGroupWithContext(ctx context.Context, request *CreateDeployGroupRequest) (response *CreateDeployGroupResponse, err error) {
    if request == nil {
        request = NewCreateDeployGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDeployGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDeployGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateParamTemplateRequest() (request *CreateParamTemplateRequest) {
    request = &CreateParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateParamTemplate")
    
    
    return
}

func NewCreateParamTemplateResponse() (response *CreateParamTemplateResponse) {
    response = &CreateParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateParamTemplate
// 该接口（CreateParamTemplate）用于创建参数模板，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    return c.CreateParamTemplateWithContext(context.Background(), request)
}

// CreateParamTemplate
// 该接口（CreateParamTemplate）用于创建参数模板，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) CreateParamTemplateWithContext(ctx context.Context, request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    if request == nil {
        request = NewCreateParamTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoInstanceIpRequest() (request *CreateRoInstanceIpRequest) {
    request = &CreateRoInstanceIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateRoInstanceIp")
    
    
    return
}

func NewCreateRoInstanceIpResponse() (response *CreateRoInstanceIpResponse) {
    response = &CreateRoInstanceIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoInstanceIp
// 本接口(CreateRoInstanceIp)用于创建云数据库只读实例的独立VIP。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEROVIPERROR = "FailedOperation.CreateRoVipError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) CreateRoInstanceIp(request *CreateRoInstanceIpRequest) (response *CreateRoInstanceIpResponse, err error) {
    return c.CreateRoInstanceIpWithContext(context.Background(), request)
}

// CreateRoInstanceIp
// 本接口(CreateRoInstanceIp)用于创建云数据库只读实例的独立VIP。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEROVIPERROR = "FailedOperation.CreateRoVipError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) CreateRoInstanceIpWithContext(ctx context.Context, request *CreateRoInstanceIpRequest) (response *CreateRoInstanceIpResponse, err error) {
    if request == nil {
        request = NewCreateRoInstanceIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoInstanceIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoInstanceIpResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountsRequest() (request *DeleteAccountsRequest) {
    request = &DeleteAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteAccounts")
    
    
    return
}

func NewDeleteAccountsResponse() (response *DeleteAccountsResponse) {
    response = &DeleteAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAccounts
// 本接口(DeleteAccounts)用于删除云数据库的账户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DeleteAccounts(request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    return c.DeleteAccountsWithContext(context.Background(), request)
}

// DeleteAccounts
// 本接口(DeleteAccounts)用于删除云数据库的账户。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DeleteAccountsWithContext(ctx context.Context, request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteAccountsRequest()
    }
    
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteAuditLogFile")
    
    
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
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteAuditLogFile(request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    return c.DeleteAuditLogFileWithContext(context.Background(), request)
}

// DeleteAuditLogFile
// 本接口(DeleteAuditLogFile)用于删除云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewDeleteAuditPolicyRequest() (request *DeleteAuditPolicyRequest) {
    request = &DeleteAuditPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteAuditPolicy")
    
    
    return
}

func NewDeleteAuditPolicyResponse() (response *DeleteAuditPolicyResponse) {
    response = &DeleteAuditPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAuditPolicy
// 本接口(DeleteAuditPolicy)用于删除用户的审计策略。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITDELETEPOLICYERROR = "InternalError.AuditDeletePolicyError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_ATLEASTALLRULEAUDITPOLICYERROR = "OperationDenied.AtLeastAllRuleAuditPolicyError"
//  OPERATIONDENIED_ATLEASTAUDITPOLICYERROR = "OperationDenied.AtLeastAuditPolicyError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteAuditPolicy(request *DeleteAuditPolicyRequest) (response *DeleteAuditPolicyResponse, err error) {
    return c.DeleteAuditPolicyWithContext(context.Background(), request)
}

// DeleteAuditPolicy
// 本接口(DeleteAuditPolicy)用于删除用户的审计策略。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITDELETEPOLICYERROR = "InternalError.AuditDeletePolicyError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_ATLEASTALLRULEAUDITPOLICYERROR = "OperationDenied.AtLeastAllRuleAuditPolicyError"
//  OPERATIONDENIED_ATLEASTAUDITPOLICYERROR = "OperationDenied.AtLeastAuditPolicyError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteAuditPolicyWithContext(ctx context.Context, request *DeleteAuditPolicyRequest) (response *DeleteAuditPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteAuditPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRuleRequest() (request *DeleteAuditRuleRequest) {
    request = &DeleteAuditRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteAuditRule")
    
    
    return
}

func NewDeleteAuditRuleResponse() (response *DeleteAuditRuleResponse) {
    response = &DeleteAuditRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAuditRule
// 本接口(DeleteAuditRule)用于删除用户的审计规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_DELETEAUDITFAILERROR = "FailedOperation.DeleteAuditFailError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  OPERATIONDENIED_AUDITRULEDELETEERROR = "OperationDenied.AuditRuleDeleteError"
//  OPERATIONDENIED_AUDITRULEHASBIND = "OperationDenied.AuditRuleHasBind"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAuditRule(request *DeleteAuditRuleRequest) (response *DeleteAuditRuleResponse, err error) {
    return c.DeleteAuditRuleWithContext(context.Background(), request)
}

// DeleteAuditRule
// 本接口(DeleteAuditRule)用于删除用户的审计规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_DELETEAUDITFAILERROR = "FailedOperation.DeleteAuditFailError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  OPERATIONDENIED_AUDITRULEDELETEERROR = "OperationDenied.AuditRuleDeleteError"
//  OPERATIONDENIED_AUDITRULEHASBIND = "OperationDenied.AuditRuleHasBind"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteAuditRuleWithContext(ctx context.Context, request *DeleteAuditRuleRequest) (response *DeleteAuditRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBackupRequest() (request *DeleteBackupRequest) {
    request = &DeleteBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteBackup")
    
    
    return
}

func NewDeleteBackupResponse() (response *DeleteBackupResponse) {
    response = &DeleteBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBackup
// 本接口(DeleteBackup)用于删除数据库备份。本接口只支持删除手动发起的备份。
//
// 可能返回的错误码:
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DeleteBackup(request *DeleteBackupRequest) (response *DeleteBackupResponse, err error) {
    return c.DeleteBackupWithContext(context.Background(), request)
}

// DeleteBackup
// 本接口(DeleteBackup)用于删除数据库备份。本接口只支持删除手动发起的备份。
//
// 可能返回的错误码:
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
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

func NewDeleteDeployGroupsRequest() (request *DeleteDeployGroupsRequest) {
    request = &DeleteDeployGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteDeployGroups")
    
    
    return
}

func NewDeleteDeployGroupsResponse() (response *DeleteDeployGroupsResponse) {
    response = &DeleteDeployGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDeployGroups
// 根据置放群组ID删除置放群组（置放群组中有资源存在时不能删除该置放群组）
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeleteDeployGroups(request *DeleteDeployGroupsRequest) (response *DeleteDeployGroupsResponse, err error) {
    return c.DeleteDeployGroupsWithContext(context.Background(), request)
}

// DeleteDeployGroups
// 根据置放群组ID删除置放群组（置放群组中有资源存在时不能删除该置放群组）
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) DeleteDeployGroupsWithContext(ctx context.Context, request *DeleteDeployGroupsRequest) (response *DeleteDeployGroupsResponse, err error) {
    if request == nil {
        request = NewDeleteDeployGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDeployGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDeployGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParamTemplateRequest() (request *DeleteParamTemplateRequest) {
    request = &DeleteParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteParamTemplate")
    
    
    return
}

func NewDeleteParamTemplateResponse() (response *DeleteParamTemplateResponse) {
    response = &DeleteParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteParamTemplate
// 该接口（DeleteParamTemplate）用于删除参数模板，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    return c.DeleteParamTemplateWithContext(context.Background(), request)
}

// DeleteParamTemplate
// 该接口（DeleteParamTemplate）用于删除参数模板，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteParamTemplateWithContext(ctx context.Context, request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteParamTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTimeWindowRequest() (request *DeleteTimeWindowRequest) {
    request = &DeleteTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteTimeWindow")
    
    
    return
}

func NewDeleteTimeWindowResponse() (response *DeleteTimeWindowResponse) {
    response = &DeleteTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTimeWindow
// 本接口(DeleteTimeWindow)用于删除云数据库实例的维护时间窗口。删除实例维护时间窗口之后，默认的维护时间窗为 03:00-04:00，即当选择在维护时间窗口内切换访问新实例时，默认会在 03:00-04:00 点进行切换访问新实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DeleteTimeWindow(request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    return c.DeleteTimeWindowWithContext(context.Background(), request)
}

// DeleteTimeWindow
// 本接口(DeleteTimeWindow)用于删除云数据库实例的维护时间窗口。删除实例维护时间窗口之后，默认的维护时间窗为 03:00-04:00，即当选择在维护时间窗口内切换访问新实例时，默认会在 03:00-04:00 点进行切换访问新实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DeleteTimeWindowWithContext(ctx context.Context, request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    if request == nil {
        request = NewDeleteTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTimeWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountPrivilegesRequest() (request *DescribeAccountPrivilegesRequest) {
    request = &DescribeAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAccountPrivileges")
    
    
    return
}

func NewDescribeAccountPrivilegesResponse() (response *DescribeAccountPrivilegesResponse) {
    response = &DescribeAccountPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountPrivileges
// 本接口(DescribeAccountPrivileges)用于查询云数据库账户支持的权限信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// 本接口(DescribeAccountPrivileges)用于查询云数据库账户支持的权限信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// 本接口(DescribeAccounts)用于查询云数据库的所有账户信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口(DescribeAccounts)用于查询云数据库的所有账户信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
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

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAsyncRequestInfo")
    
    
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAsyncRequestInfo
// 本接口(DescribeAsyncRequestInfo)用于查询云数据库实例异步任务的执行结果。
//
// 可能返回的错误码:
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// 本接口(DescribeAsyncRequestInfo)用于查询云数据库实例异步任务的执行结果。
//
// 可能返回的错误码:
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditConfigRequest() (request *DescribeAuditConfigRequest) {
    request = &DescribeAuditConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditConfig")
    
    
    return
}

func NewDescribeAuditConfigResponse() (response *DescribeAuditConfigResponse) {
    response = &DescribeAuditConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditConfig
// 本接口(DescribeAuditConfig)用于查询云数据库审计策略的服务配置，包括审计日志保存时长等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAuditConfig(request *DescribeAuditConfigRequest) (response *DescribeAuditConfigResponse, err error) {
    return c.DescribeAuditConfigWithContext(context.Background(), request)
}

// DescribeAuditConfig
// 本接口(DescribeAuditConfig)用于查询云数据库审计策略的服务配置，包括审计日志保存时长等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAuditConfigWithContext(ctx context.Context, request *DescribeAuditConfigRequest) (response *DescribeAuditConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAuditConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogFilesRequest() (request *DescribeAuditLogFilesRequest) {
    request = &DescribeAuditLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditLogFiles")
    
    
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
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditLogFiles(request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    return c.DescribeAuditLogFilesWithContext(context.Background(), request)
}

// DescribeAuditLogFiles
// 本接口(DescribeAuditLogFiles)用于查询云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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

func NewDescribeAuditPoliciesRequest() (request *DescribeAuditPoliciesRequest) {
    request = &DescribeAuditPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditPolicies")
    
    
    return
}

func NewDescribeAuditPoliciesResponse() (response *DescribeAuditPoliciesResponse) {
    response = &DescribeAuditPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditPolicies
// 本接口(DescribeAuditPolicies)用于查询云数据库实例的审计策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeAuditPolicies(request *DescribeAuditPoliciesRequest) (response *DescribeAuditPoliciesResponse, err error) {
    return c.DescribeAuditPoliciesWithContext(context.Background(), request)
}

// DescribeAuditPolicies
// 本接口(DescribeAuditPolicies)用于查询云数据库实例的审计策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeAuditPoliciesWithContext(ctx context.Context, request *DescribeAuditPoliciesRequest) (response *DescribeAuditPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRulesRequest() (request *DescribeAuditRulesRequest) {
    request = &DescribeAuditRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditRules")
    
    
    return
}

func NewDescribeAuditRulesResponse() (response *DescribeAuditRulesResponse) {
    response = &DescribeAuditRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuditRules
// 本接口(DescribeAuditRules)用于查询用户在当前地域的审计规则。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
func (c *Client) DescribeAuditRules(request *DescribeAuditRulesRequest) (response *DescribeAuditRulesResponse, err error) {
    return c.DescribeAuditRulesWithContext(context.Background(), request)
}

// DescribeAuditRules
// 本接口(DescribeAuditRules)用于查询用户在当前地域的审计规则。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
func (c *Client) DescribeAuditRulesWithContext(ctx context.Context, request *DescribeAuditRulesRequest) (response *DescribeAuditRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRulesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupConfigRequest() (request *DescribeBackupConfigRequest) {
    request = &DescribeBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupConfig")
    
    
    return
}

func NewDescribeBackupConfigResponse() (response *DescribeBackupConfigResponse) {
    response = &DescribeBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupConfig
// 本接口(DescribeBackupConfig)用于查询数据库备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    return c.DescribeBackupConfigWithContext(context.Background(), request)
}

// DescribeBackupConfig
// 本接口(DescribeBackupConfig)用于查询数据库备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
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

func NewDescribeBackupDatabasesRequest() (request *DescribeBackupDatabasesRequest) {
    request = &DescribeBackupDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDatabases")
    
    
    return
}

func NewDescribeBackupDatabasesResponse() (response *DescribeBackupDatabasesResponse) {
    response = &DescribeBackupDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupDatabases
// 本接口(DescribeBackupDatabases)用于查询备份文件包含的库 (已废弃)。
//
// 旧版本支持全量备份后，用户如果分库表下载逻辑备份文件，需要用到此接口。
//
// 新版本支持(CreateBackup)创建逻辑备份的时候，直接发起指定库表备份，用户直接下载该备份文件即可。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupDatabases(request *DescribeBackupDatabasesRequest) (response *DescribeBackupDatabasesResponse, err error) {
    return c.DescribeBackupDatabasesWithContext(context.Background(), request)
}

// DescribeBackupDatabases
// 本接口(DescribeBackupDatabases)用于查询备份文件包含的库 (已废弃)。
//
// 旧版本支持全量备份后，用户如果分库表下载逻辑备份文件，需要用到此接口。
//
// 新版本支持(CreateBackup)创建逻辑备份的时候，直接发起指定库表备份，用户直接下载该备份文件即可。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupDatabasesWithContext(ctx context.Context, request *DescribeBackupDatabasesRequest) (response *DescribeBackupDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDatabasesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDatabasesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupDownloadRestriction
// 该接口用户查询当前地域用户设置的默认备份下载来源限制。
//
// 可能返回的错误码:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// 该接口用户查询当前地域用户设置的默认备份下载来源限制。
//
// 可能返回的错误码:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBackupDownloadRestrictionWithContext(ctx context.Context, request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadRestrictionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupEncryptionStatusRequest() (request *DescribeBackupEncryptionStatusRequest) {
    request = &DescribeBackupEncryptionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupEncryptionStatus")
    
    
    return
}

func NewDescribeBackupEncryptionStatusResponse() (response *DescribeBackupEncryptionStatusResponse) {
    response = &DescribeBackupEncryptionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupEncryptionStatus
// 本接口(DescribeBackupEncryptionStatus)用于查询实例默认备份加密状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBackupEncryptionStatus(request *DescribeBackupEncryptionStatusRequest) (response *DescribeBackupEncryptionStatusResponse, err error) {
    return c.DescribeBackupEncryptionStatusWithContext(context.Background(), request)
}

// DescribeBackupEncryptionStatus
// 本接口(DescribeBackupEncryptionStatus)用于查询实例默认备份加密状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeBackupEncryptionStatusWithContext(ctx context.Context, request *DescribeBackupEncryptionStatusRequest) (response *DescribeBackupEncryptionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeBackupEncryptionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupEncryptionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupEncryptionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupOverviewRequest() (request *DescribeBackupOverviewRequest) {
    request = &DescribeBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupOverview")
    
    
    return
}

func NewDescribeBackupOverviewResponse() (response *DescribeBackupOverviewResponse) {
    response = &DescribeBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupOverview
// 本接口(DescribeBackupOverview)用于查询用户的备份概览。返回用户当前备份总个数、备份总的占用容量、赠送的免费容量、计费容量（容量单位为字节）。
//
// 可能返回的错误码:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeBackupOverview(request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    return c.DescribeBackupOverviewWithContext(context.Background(), request)
}

// DescribeBackupOverview
// 本接口(DescribeBackupOverview)用于查询用户的备份概览。返回用户当前备份总个数、备份总的占用容量、赠送的免费容量、计费容量（容量单位为字节）。
//
// 可能返回的错误码:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeBackupOverviewWithContext(ctx context.Context, request *DescribeBackupOverviewRequest) (response *DescribeBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupSummariesRequest() (request *DescribeBackupSummariesRequest) {
    request = &DescribeBackupSummariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupSummaries")
    
    
    return
}

func NewDescribeBackupSummariesResponse() (response *DescribeBackupSummariesResponse) {
    response = &DescribeBackupSummariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupSummaries
// 本接口(DescribeBackupSummaries)用于查询备份的统计情况，返回以实例为维度的备份占用容量，以及每个实例的数据备份和日志备份的个数和容量（容量单位为字节）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBackupSummaries(request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    return c.DescribeBackupSummariesWithContext(context.Background(), request)
}

// DescribeBackupSummaries
// 本接口(DescribeBackupSummaries)用于查询备份的统计情况，返回以实例为维度的备份占用容量，以及每个实例的数据备份和日志备份的个数和容量（容量单位为字节）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBackupSummariesWithContext(ctx context.Context, request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupSummariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupSummaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupSummariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupTablesRequest() (request *DescribeBackupTablesRequest) {
    request = &DescribeBackupTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupTables")
    
    
    return
}

func NewDescribeBackupTablesResponse() (response *DescribeBackupTablesResponse) {
    response = &DescribeBackupTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackupTables
// 本接口(DescribeBackupTables)用于查询指定的数据库的备份数据表名 (已废弃)。
//
// 旧版本支持全量备份后，用户如果分库表下载逻辑备份文件，需要用到此接口。
//
// 新版本支持(CreateBackup)创建逻辑备份的时候，直接发起指定库表备份，用户直接下载该备份文件即可。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupTables(request *DescribeBackupTablesRequest) (response *DescribeBackupTablesResponse, err error) {
    return c.DescribeBackupTablesWithContext(context.Background(), request)
}

// DescribeBackupTables
// 本接口(DescribeBackupTables)用于查询指定的数据库的备份数据表名 (已废弃)。
//
// 旧版本支持全量备份后，用户如果分库表下载逻辑备份文件，需要用到此接口。
//
// 新版本支持(CreateBackup)创建逻辑备份的时候，直接发起指定库表备份，用户直接下载该备份文件即可。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupTablesWithContext(ctx context.Context, request *DescribeBackupTablesRequest) (response *DescribeBackupTablesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupsRequest() (request *DescribeBackupsRequest) {
    request = &DescribeBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackups")
    
    
    return
}

func NewDescribeBackupsResponse() (response *DescribeBackupsResponse) {
    response = &DescribeBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBackups
// 本接口(DescribeBackups)用于查询云数据库实例的备份数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBackups(request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    return c.DescribeBackupsWithContext(context.Background(), request)
}

// DescribeBackups
// 本接口(DescribeBackups)用于查询云数据库实例的备份数据。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeBinlogBackupOverviewRequest() (request *DescribeBinlogBackupOverviewRequest) {
    request = &DescribeBinlogBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBinlogBackupOverview")
    
    
    return
}

func NewDescribeBinlogBackupOverviewResponse() (response *DescribeBinlogBackupOverviewResponse) {
    response = &DescribeBinlogBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBinlogBackupOverview
// 本接口(DescribeBinlogBackupOverview)用于查询用户在当前地域总的日志备份概览。
//
// 可能返回的错误码:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBinlogBackupOverview(request *DescribeBinlogBackupOverviewRequest) (response *DescribeBinlogBackupOverviewResponse, err error) {
    return c.DescribeBinlogBackupOverviewWithContext(context.Background(), request)
}

// DescribeBinlogBackupOverview
// 本接口(DescribeBinlogBackupOverview)用于查询用户在当前地域总的日志备份概览。
//
// 可能返回的错误码:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBinlogBackupOverviewWithContext(ctx context.Context, request *DescribeBinlogBackupOverviewRequest) (response *DescribeBinlogBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeBinlogBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogBackupOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBinlogsRequest() (request *DescribeBinlogsRequest) {
    request = &DescribeBinlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBinlogs")
    
    
    return
}

func NewDescribeBinlogsResponse() (response *DescribeBinlogsResponse) {
    response = &DescribeBinlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBinlogs
// 本接口(DescribeBinlogs)用于查询云数据库实例的 binlog 文件列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeBinlogs(request *DescribeBinlogsRequest) (response *DescribeBinlogsResponse, err error) {
    return c.DescribeBinlogsWithContext(context.Background(), request)
}

// DescribeBinlogs
// 本接口(DescribeBinlogs)用于查询云数据库实例的 binlog 文件列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
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

func NewDescribeCDBProxyRequest() (request *DescribeCDBProxyRequest) {
    request = &DescribeCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCDBProxy")
    
    
    return
}

func NewDescribeCDBProxyResponse() (response *DescribeCDBProxyResponse) {
    response = &DescribeCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCDBProxy
// 查询数据库代理（待下线，替换接口QueryCDBProxy）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCDBProxy(request *DescribeCDBProxyRequest) (response *DescribeCDBProxyResponse, err error) {
    return c.DescribeCDBProxyWithContext(context.Background(), request)
}

// DescribeCDBProxy
// 查询数据库代理（待下线，替换接口QueryCDBProxy）
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCDBProxyWithContext(ctx context.Context, request *DescribeCDBProxyRequest) (response *DescribeCDBProxyResponse, err error) {
    if request == nil {
        request = NewDescribeCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdbZoneConfigRequest() (request *DescribeCdbZoneConfigRequest) {
    request = &DescribeCdbZoneConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCdbZoneConfig")
    
    
    return
}

func NewDescribeCdbZoneConfigResponse() (response *DescribeCdbZoneConfigResponse) {
    response = &DescribeCdbZoneConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCdbZoneConfig
// 本接口(DescribeCdbZoneConfig)用于查询云数据库各地域可售卖的规格配置。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
func (c *Client) DescribeCdbZoneConfig(request *DescribeCdbZoneConfigRequest) (response *DescribeCdbZoneConfigResponse, err error) {
    return c.DescribeCdbZoneConfigWithContext(context.Background(), request)
}

// DescribeCdbZoneConfig
// 本接口(DescribeCdbZoneConfig)用于查询云数据库各地域可售卖的规格配置。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
func (c *Client) DescribeCdbZoneConfigWithContext(ctx context.Context, request *DescribeCdbZoneConfigRequest) (response *DescribeCdbZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribeCdbZoneConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdbZoneConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdbZoneConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloneListRequest() (request *DescribeCloneListRequest) {
    request = &DescribeCloneListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCloneList")
    
    
    return
}

func NewDescribeCloneListResponse() (response *DescribeCloneListResponse) {
    response = &DescribeCloneListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloneList
// 本接口(DescribeCloneList) 用于查询用户实例的克隆任务列表。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCloneList(request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    return c.DescribeCloneListWithContext(context.Background(), request)
}

// DescribeCloneList
// 本接口(DescribeCloneList) 用于查询用户实例的克隆任务列表。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCloneListWithContext(ctx context.Context, request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    if request == nil {
        request = NewDescribeCloneListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloneList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloneListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBFeaturesRequest() (request *DescribeDBFeaturesRequest) {
    request = &DescribeDBFeaturesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBFeatures")
    
    
    return
}

func NewDescribeDBFeaturesResponse() (response *DescribeDBFeaturesResponse) {
    response = &DescribeDBFeaturesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBFeatures
// 本接口(DescribeDBFeatures)用于查询云数据库版本属性，包括是否支持数据库加密、数据库审计等功能。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBFeatures(request *DescribeDBFeaturesRequest) (response *DescribeDBFeaturesResponse, err error) {
    return c.DescribeDBFeaturesWithContext(context.Background(), request)
}

// DescribeDBFeatures
// 本接口(DescribeDBFeatures)用于查询云数据库版本属性，包括是否支持数据库加密、数据库审计等功能。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBFeaturesWithContext(ctx context.Context, request *DescribeDBFeaturesRequest) (response *DescribeDBFeaturesResponse, err error) {
    if request == nil {
        request = NewDescribeDBFeaturesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBFeatures require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBFeaturesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBImportRecordsRequest() (request *DescribeDBImportRecordsRequest) {
    request = &DescribeDBImportRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBImportRecords")
    
    
    return
}

func NewDescribeDBImportRecordsResponse() (response *DescribeDBImportRecordsResponse) {
    response = &DescribeDBImportRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBImportRecords
// 本接口(DescribeDBImportRecords)用于查询云数据库导入任务操作日志。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBImportRecords(request *DescribeDBImportRecordsRequest) (response *DescribeDBImportRecordsResponse, err error) {
    return c.DescribeDBImportRecordsWithContext(context.Background(), request)
}

// DescribeDBImportRecords
// 本接口(DescribeDBImportRecords)用于查询云数据库导入任务操作日志。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBImportRecordsWithContext(ctx context.Context, request *DescribeDBImportRecordsRequest) (response *DescribeDBImportRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBImportRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBImportRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBImportRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceCharsetRequest() (request *DescribeDBInstanceCharsetRequest) {
    request = &DescribeDBInstanceCharsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceCharset")
    
    
    return
}

func NewDescribeDBInstanceCharsetResponse() (response *DescribeDBInstanceCharsetResponse) {
    response = &DescribeDBInstanceCharsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceCharset
// 本接口(DescribeDBInstanceCharset)用于查询云数据库实例的字符集，获取字符集的名称。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceCharset(request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    return c.DescribeDBInstanceCharsetWithContext(context.Background(), request)
}

// DescribeDBInstanceCharset
// 本接口(DescribeDBInstanceCharset)用于查询云数据库实例的字符集，获取字符集的名称。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceCharsetWithContext(ctx context.Context, request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceCharsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceCharset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceCharsetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceConfigRequest() (request *DescribeDBInstanceConfigRequest) {
    request = &DescribeDBInstanceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceConfig")
    
    
    return
}

func NewDescribeDBInstanceConfigResponse() (response *DescribeDBInstanceConfigResponse) {
    response = &DescribeDBInstanceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceConfig
// 本接口(DescribeDBInstanceConfig)用于云数据库实例的配置信息，包括同步模式，部署模式等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceConfig(request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    return c.DescribeDBInstanceConfigWithContext(context.Background(), request)
}

// DescribeDBInstanceConfig
// 本接口(DescribeDBInstanceConfig)用于云数据库实例的配置信息，包括同步模式，部署模式等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceConfigWithContext(ctx context.Context, request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceGTIDRequest() (request *DescribeDBInstanceGTIDRequest) {
    request = &DescribeDBInstanceGTIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceGTID")
    
    
    return
}

func NewDescribeDBInstanceGTIDResponse() (response *DescribeDBInstanceGTIDResponse) {
    response = &DescribeDBInstanceGTIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceGTID
// 本接口(DescribeDBInstanceGTID)用于查询云数据库实例是否开通了 GTID，不支持版本为 5.5 以及以下的实例。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceGTID(request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    return c.DescribeDBInstanceGTIDWithContext(context.Background(), request)
}

// DescribeDBInstanceGTID
// 本接口(DescribeDBInstanceGTID)用于查询云数据库实例是否开通了 GTID，不支持版本为 5.5 以及以下的实例。
//
// 可能返回的错误码:
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceGTIDWithContext(ctx context.Context, request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceGTIDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceGTID require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceGTIDResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceInfoRequest() (request *DescribeDBInstanceInfoRequest) {
    request = &DescribeDBInstanceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceInfo")
    
    
    return
}

func NewDescribeDBInstanceInfoResponse() (response *DescribeDBInstanceInfoResponse) {
    response = &DescribeDBInstanceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceInfo
// 查询实例基本信息（实例 ID ，实例名称，是否开通加密 ）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) DescribeDBInstanceInfo(request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    return c.DescribeDBInstanceInfoWithContext(context.Background(), request)
}

// DescribeDBInstanceInfo
// 查询实例基本信息（实例 ID ，实例名称，是否开通加密 ）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) DescribeDBInstanceInfoWithContext(ctx context.Context, request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceRebootTimeRequest() (request *DescribeDBInstanceRebootTimeRequest) {
    request = &DescribeDBInstanceRebootTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceRebootTime")
    
    
    return
}

func NewDescribeDBInstanceRebootTimeResponse() (response *DescribeDBInstanceRebootTimeResponse) {
    response = &DescribeDBInstanceRebootTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstanceRebootTime
// 本接口(DescribeDBInstanceRebootTime)用于查询云数据库实例重启预计所需的时间。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceRebootTime(request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    return c.DescribeDBInstanceRebootTimeWithContext(context.Background(), request)
}

// DescribeDBInstanceRebootTime
// 本接口(DescribeDBInstanceRebootTime)用于查询云数据库实例重启预计所需的时间。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceRebootTimeWithContext(ctx context.Context, request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceRebootTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceRebootTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceRebootTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBInstances
// 本接口(DescribeDBInstances)用于查询云数据库实例列表，支持通过项目 ID、实例 ID、访问地址、实例状态等过滤条件来筛选实例。支持查询主实例、灾备实例和只读实例信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口(DescribeDBInstances)用于查询云数据库实例列表，支持通过项目 ID、实例 ID、访问地址、实例状态等过滤条件来筛选实例。支持查询主实例、灾备实例和只读实例信息列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeDBPriceRequest() (request *DescribeDBPriceRequest) {
    request = &DescribeDBPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBPrice")
    
    
    return
}

func NewDescribeDBPriceResponse() (response *DescribeDBPriceResponse) {
    response = &DescribeDBPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBPrice
// 本接口(DescribeDBPrice)用于查询购买或续费云数据库实例的价格，支持查询按量计费或者包年包月的价格。可传入实例类型、购买时长、购买数量、内存大小、硬盘大小和可用区信息等来查询实例价格。可传入实例名称来查询实例续费价格。
//
// 
//
// 注意：对某个地域进行询价，请使用对应地域的接入点，接入点信息请参照 <a href="https://cloud.tencent.com/document/api/236/15832">服务地址</a> 文档。例如：对广州地域进行询价，请把请求发到：cdb.ap-guangzhou.tencentcloudapi.com。同理对上海地域询价，把请求发到：cdb.ap-shanghai.tencentcloudapi.com。
//
// 可能返回的错误码:
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDBPrice(request *DescribeDBPriceRequest) (response *DescribeDBPriceResponse, err error) {
    return c.DescribeDBPriceWithContext(context.Background(), request)
}

// DescribeDBPrice
// 本接口(DescribeDBPrice)用于查询购买或续费云数据库实例的价格，支持查询按量计费或者包年包月的价格。可传入实例类型、购买时长、购买数量、内存大小、硬盘大小和可用区信息等来查询实例价格。可传入实例名称来查询实例续费价格。
//
// 
//
// 注意：对某个地域进行询价，请使用对应地域的接入点，接入点信息请参照 <a href="https://cloud.tencent.com/document/api/236/15832">服务地址</a> 文档。例如：对广州地域进行询价，请把请求发到：cdb.ap-guangzhou.tencentcloudapi.com。同理对上海地域询价，把请求发到：cdb.ap-shanghai.tencentcloudapi.com。
//
// 可能返回的错误码:
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeDBPriceWithContext(ctx context.Context, request *DescribeDBPriceRequest) (response *DescribeDBPriceResponse, err error) {
    if request == nil {
        request = NewDescribeDBPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBSecurityGroups")
    
    
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
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewDescribeDBSwitchRecordsRequest() (request *DescribeDBSwitchRecordsRequest) {
    request = &DescribeDBSwitchRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBSwitchRecords")
    
    
    return
}

func NewDescribeDBSwitchRecordsResponse() (response *DescribeDBSwitchRecordsResponse) {
    response = &DescribeDBSwitchRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDBSwitchRecords
// 本接口(DescribeDBSwitchRecords)用于查询云数据库实例切换记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBSwitchRecords(request *DescribeDBSwitchRecordsRequest) (response *DescribeDBSwitchRecordsResponse, err error) {
    return c.DescribeDBSwitchRecordsWithContext(context.Background(), request)
}

// DescribeDBSwitchRecords
// 本接口(DescribeDBSwitchRecords)用于查询云数据库实例切换记录。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBSwitchRecordsWithContext(ctx context.Context, request *DescribeDBSwitchRecordsRequest) (response *DescribeDBSwitchRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSwitchRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSwitchRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSwitchRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDataBackupOverviewRequest() (request *DescribeDataBackupOverviewRequest) {
    request = &DescribeDataBackupOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDataBackupOverview")
    
    
    return
}

func NewDescribeDataBackupOverviewResponse() (response *DescribeDataBackupOverviewResponse) {
    response = &DescribeDataBackupOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDataBackupOverview
// 本接口(DescribeDataBackupOverview)用于查询用户在当前地域总的数据备份概览。
//
// 可能返回的错误码:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDataBackupOverview(request *DescribeDataBackupOverviewRequest) (response *DescribeDataBackupOverviewResponse, err error) {
    return c.DescribeDataBackupOverviewWithContext(context.Background(), request)
}

// DescribeDataBackupOverview
// 本接口(DescribeDataBackupOverview)用于查询用户在当前地域总的数据备份概览。
//
// 可能返回的错误码:
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDataBackupOverviewWithContext(ctx context.Context, request *DescribeDataBackupOverviewRequest) (response *DescribeDataBackupOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeDataBackupOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDataBackupOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDataBackupOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDatabases")
    
    
    return
}

func NewDescribeDatabasesResponse() (response *DescribeDatabasesResponse) {
    response = &DescribeDatabasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatabases
// 本接口(DescribeDatabases)用于查询云数据库实例的数据库信息，仅支持主实例和灾备实例，不支持只读实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// 本接口(DescribeDatabases)用于查询云数据库实例的数据库信息，仅支持主实例和灾备实例，不支持只读实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
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

func NewDescribeDefaultParamsRequest() (request *DescribeDefaultParamsRequest) {
    request = &DescribeDefaultParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDefaultParams")
    
    
    return
}

func NewDescribeDefaultParamsResponse() (response *DescribeDefaultParamsResponse) {
    response = &DescribeDefaultParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDefaultParams
// 该接口（DescribeDefaultParams）用于查询默认的可设置参数列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDefaultParams(request *DescribeDefaultParamsRequest) (response *DescribeDefaultParamsResponse, err error) {
    return c.DescribeDefaultParamsWithContext(context.Background(), request)
}

// DescribeDefaultParams
// 该接口（DescribeDefaultParams）用于查询默认的可设置参数列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDefaultParamsWithContext(ctx context.Context, request *DescribeDefaultParamsRequest) (response *DescribeDefaultParamsResponse, err error) {
    if request == nil {
        request = NewDescribeDefaultParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDefaultParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDefaultParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeployGroupListRequest() (request *DescribeDeployGroupListRequest) {
    request = &DescribeDeployGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDeployGroupList")
    
    
    return
}

func NewDescribeDeployGroupListResponse() (response *DescribeDeployGroupListResponse) {
    response = &DescribeDeployGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeployGroupList
// 本接口(DescribeDeployGroupList)用于查询用户的置放群组列表，可以指定置放群组 ID 或置放群组名称。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeDeployGroupList(request *DescribeDeployGroupListRequest) (response *DescribeDeployGroupListResponse, err error) {
    return c.DescribeDeployGroupListWithContext(context.Background(), request)
}

// DescribeDeployGroupList
// 本接口(DescribeDeployGroupList)用于查询用户的置放群组列表，可以指定置放群组 ID 或置放群组名称。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeDeployGroupListWithContext(ctx context.Context, request *DescribeDeployGroupListRequest) (response *DescribeDeployGroupListResponse, err error) {
    if request == nil {
        request = NewDescribeDeployGroupListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeployGroupList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeployGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceMonitorInfoRequest() (request *DescribeDeviceMonitorInfoRequest) {
    request = &DescribeDeviceMonitorInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDeviceMonitorInfo")
    
    
    return
}

func NewDescribeDeviceMonitorInfoResponse() (response *DescribeDeviceMonitorInfoResponse) {
    response = &DescribeDeviceMonitorInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceMonitorInfo
// 本接口（DescribeDeviceMonitorInfo）用于查询云数据库物理机当天的监控信息，暂只支持内存488G、硬盘6T的实例查询。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_FUNCTIONDENIED = "OperationDenied.FunctionDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeDeviceMonitorInfo(request *DescribeDeviceMonitorInfoRequest) (response *DescribeDeviceMonitorInfoResponse, err error) {
    return c.DescribeDeviceMonitorInfoWithContext(context.Background(), request)
}

// DescribeDeviceMonitorInfo
// 本接口（DescribeDeviceMonitorInfo）用于查询云数据库物理机当天的监控信息，暂只支持内存488G、硬盘6T的实例查询。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_FUNCTIONDENIED = "OperationDenied.FunctionDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeDeviceMonitorInfoWithContext(ctx context.Context, request *DescribeDeviceMonitorInfoRequest) (response *DescribeDeviceMonitorInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceMonitorInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceMonitorInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceMonitorInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeErrorLogDataRequest() (request *DescribeErrorLogDataRequest) {
    request = &DescribeErrorLogDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeErrorLogData")
    
    
    return
}

func NewDescribeErrorLogDataResponse() (response *DescribeErrorLogDataResponse) {
    response = &DescribeErrorLogDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeErrorLogData
// 根据检索条件查询实例错误日志详情。只能查询一个月之内的错误日志。
//
// 使用时需要注意：可能存在单条错误日志太大，导致整个http请求的回包太大，进而引发接口超时。一旦发生超时，建议您缩小查询时的Limit参数值，从而降低包的大小，让接口能够及时返回内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeErrorLogData(request *DescribeErrorLogDataRequest) (response *DescribeErrorLogDataResponse, err error) {
    return c.DescribeErrorLogDataWithContext(context.Background(), request)
}

// DescribeErrorLogData
// 根据检索条件查询实例错误日志详情。只能查询一个月之内的错误日志。
//
// 使用时需要注意：可能存在单条错误日志太大，导致整个http请求的回包太大，进而引发接口超时。一旦发生超时，建议您缩小查询时的Limit参数值，从而降低包的大小，让接口能够及时返回内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeErrorLogDataWithContext(ctx context.Context, request *DescribeErrorLogDataRequest) (response *DescribeErrorLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeErrorLogDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeErrorLogData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceParamRecords")
    
    
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
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// 该接口（DescribeInstanceParamRecords）用于查询实例参数修改历史。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceParams")
    
    
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
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// 该接口（DescribeInstanceParams）用于查询实例的参数列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
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

func NewDescribeLocalBinlogConfigRequest() (request *DescribeLocalBinlogConfigRequest) {
    request = &DescribeLocalBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeLocalBinlogConfig")
    
    
    return
}

func NewDescribeLocalBinlogConfigResponse() (response *DescribeLocalBinlogConfigResponse) {
    response = &DescribeLocalBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLocalBinlogConfig
// 该接口用于查询实例本地binlog保留策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeLocalBinlogConfig(request *DescribeLocalBinlogConfigRequest) (response *DescribeLocalBinlogConfigResponse, err error) {
    return c.DescribeLocalBinlogConfigWithContext(context.Background(), request)
}

// DescribeLocalBinlogConfig
// 该接口用于查询实例本地binlog保留策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeLocalBinlogConfigWithContext(ctx context.Context, request *DescribeLocalBinlogConfigRequest) (response *DescribeLocalBinlogConfigResponse, err error) {
    if request == nil {
        request = NewDescribeLocalBinlogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLocalBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLocalBinlogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplateInfoRequest() (request *DescribeParamTemplateInfoRequest) {
    request = &DescribeParamTemplateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeParamTemplateInfo")
    
    
    return
}

func NewDescribeParamTemplateInfoResponse() (response *DescribeParamTemplateInfoResponse) {
    response = &DescribeParamTemplateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeParamTemplateInfo
// 该接口（DescribeParamTemplateInfo）用于查询参数模板详情，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplateInfo(request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    return c.DescribeParamTemplateInfoWithContext(context.Background(), request)
}

// DescribeParamTemplateInfo
// 该接口（DescribeParamTemplateInfo）用于查询参数模板详情，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplateInfoWithContext(ctx context.Context, request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeParamTemplateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeParamTemplateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeParamTemplatesRequest() (request *DescribeParamTemplatesRequest) {
    request = &DescribeParamTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeParamTemplates")
    
    
    return
}

func NewDescribeParamTemplatesResponse() (response *DescribeParamTemplatesResponse) {
    response = &DescribeParamTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeParamTemplates
// 该接口（DescribeParamTemplates）查询参数模板列表，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// 该接口（DescribeParamTemplates）查询参数模板列表，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProjectSecurityGroups")
    
    
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
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewDescribeProxyConnectionPoolConfRequest() (request *DescribeProxyConnectionPoolConfRequest) {
    request = &DescribeProxyConnectionPoolConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProxyConnectionPoolConf")
    
    
    return
}

func NewDescribeProxyConnectionPoolConfResponse() (response *DescribeProxyConnectionPoolConfResponse) {
    response = &DescribeProxyConnectionPoolConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyConnectionPoolConf
// 获取数据库代理连接池相关规格配置
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeProxyConnectionPoolConf(request *DescribeProxyConnectionPoolConfRequest) (response *DescribeProxyConnectionPoolConfResponse, err error) {
    return c.DescribeProxyConnectionPoolConfWithContext(context.Background(), request)
}

// DescribeProxyConnectionPoolConf
// 获取数据库代理连接池相关规格配置
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeProxyConnectionPoolConfWithContext(ctx context.Context, request *DescribeProxyConnectionPoolConfRequest) (response *DescribeProxyConnectionPoolConfResponse, err error) {
    if request == nil {
        request = NewDescribeProxyConnectionPoolConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyConnectionPoolConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyConnectionPoolConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyCustomConfRequest() (request *DescribeProxyCustomConfRequest) {
    request = &DescribeProxyCustomConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProxyCustomConf")
    
    
    return
}

func NewDescribeProxyCustomConfResponse() (response *DescribeProxyCustomConfResponse) {
    response = &DescribeProxyCustomConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxyCustomConf
// 查询代理规格配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeProxyCustomConf(request *DescribeProxyCustomConfRequest) (response *DescribeProxyCustomConfResponse, err error) {
    return c.DescribeProxyCustomConfWithContext(context.Background(), request)
}

// DescribeProxyCustomConf
// 查询代理规格配置
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeProxyCustomConfWithContext(ctx context.Context, request *DescribeProxyCustomConfRequest) (response *DescribeProxyCustomConfResponse, err error) {
    if request == nil {
        request = NewDescribeProxyCustomConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyCustomConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyCustomConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRemoteBackupConfigRequest() (request *DescribeRemoteBackupConfigRequest) {
    request = &DescribeRemoteBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRemoteBackupConfig")
    
    
    return
}

func NewDescribeRemoteBackupConfigResponse() (response *DescribeRemoteBackupConfigResponse) {
    response = &DescribeRemoteBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRemoteBackupConfig
// 本接口(DescribeRemoteBackupConfig)用于查询数据库异地备份配置信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeRemoteBackupConfig(request *DescribeRemoteBackupConfigRequest) (response *DescribeRemoteBackupConfigResponse, err error) {
    return c.DescribeRemoteBackupConfigWithContext(context.Background(), request)
}

// DescribeRemoteBackupConfig
// 本接口(DescribeRemoteBackupConfig)用于查询数据库异地备份配置信息。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeRemoteBackupConfigWithContext(ctx context.Context, request *DescribeRemoteBackupConfigRequest) (response *DescribeRemoteBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeRemoteBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRemoteBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRemoteBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoGroupsRequest() (request *DescribeRoGroupsRequest) {
    request = &DescribeRoGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRoGroups")
    
    
    return
}

func NewDescribeRoGroupsResponse() (response *DescribeRoGroupsResponse) {
    response = &DescribeRoGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoGroups
// 本接口(DescribeRoGroups)用于查询云数据库实例的所有的RO组的信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeRoGroups(request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    return c.DescribeRoGroupsWithContext(context.Background(), request)
}

// DescribeRoGroups
// 本接口(DescribeRoGroups)用于查询云数据库实例的所有的RO组的信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeRoGroupsWithContext(ctx context.Context, request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRoGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoMinScaleRequest() (request *DescribeRoMinScaleRequest) {
    request = &DescribeRoMinScaleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRoMinScale")
    
    
    return
}

func NewDescribeRoMinScaleResponse() (response *DescribeRoMinScaleResponse) {
    response = &DescribeRoMinScaleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoMinScale
// 本接口(DescribeRoMinScale)用于获取只读实例购买、升级时的最小规格。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeRoMinScale(request *DescribeRoMinScaleRequest) (response *DescribeRoMinScaleResponse, err error) {
    return c.DescribeRoMinScaleWithContext(context.Background(), request)
}

// DescribeRoMinScale
// 本接口(DescribeRoMinScale)用于获取只读实例购买、升级时的最小规格。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) DescribeRoMinScaleWithContext(ctx context.Context, request *DescribeRoMinScaleRequest) (response *DescribeRoMinScaleResponse, err error) {
    if request == nil {
        request = NewDescribeRoMinScaleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoMinScale require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoMinScaleResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackRangeTimeRequest() (request *DescribeRollbackRangeTimeRequest) {
    request = &DescribeRollbackRangeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRollbackRangeTime")
    
    
    return
}

func NewDescribeRollbackRangeTimeResponse() (response *DescribeRollbackRangeTimeResponse) {
    response = &DescribeRollbackRangeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRollbackRangeTime
// 本接口(DescribeRollbackRangeTime)用于查询云数据库实例可回档的时间范围。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRollbackRangeTime(request *DescribeRollbackRangeTimeRequest) (response *DescribeRollbackRangeTimeResponse, err error) {
    return c.DescribeRollbackRangeTimeWithContext(context.Background(), request)
}

// DescribeRollbackRangeTime
// 本接口(DescribeRollbackRangeTime)用于查询云数据库实例可回档的时间范围。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRollbackRangeTimeWithContext(ctx context.Context, request *DescribeRollbackRangeTimeRequest) (response *DescribeRollbackRangeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackRangeTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackRangeTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackRangeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRollbackTaskDetailRequest() (request *DescribeRollbackTaskDetailRequest) {
    request = &DescribeRollbackTaskDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeRollbackTaskDetail")
    
    
    return
}

func NewDescribeRollbackTaskDetailResponse() (response *DescribeRollbackTaskDetailResponse) {
    response = &DescribeRollbackTaskDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRollbackTaskDetail
// 本接口(DescribeRollbackTaskDetail)用于查询云数据库实例回档任务详情。
//
// 可能返回的错误码:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeRollbackTaskDetail(request *DescribeRollbackTaskDetailRequest) (response *DescribeRollbackTaskDetailResponse, err error) {
    return c.DescribeRollbackTaskDetailWithContext(context.Background(), request)
}

// DescribeRollbackTaskDetail
// 本接口(DescribeRollbackTaskDetail)用于查询云数据库实例回档任务详情。
//
// 可能返回的错误码:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeRollbackTaskDetailWithContext(ctx context.Context, request *DescribeRollbackTaskDetailRequest) (response *DescribeRollbackTaskDetailResponse, err error) {
    if request == nil {
        request = NewDescribeRollbackTaskDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogDataRequest() (request *DescribeSlowLogDataRequest) {
    request = &DescribeSlowLogDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSlowLogData")
    
    
    return
}

func NewDescribeSlowLogDataResponse() (response *DescribeSlowLogDataResponse) {
    response = &DescribeSlowLogDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowLogData
// 条件检索实例的慢日志。只允许查看一个月之内的慢日志。
//
// 使用时需要注意：可能存在单条慢日志太大，导致整个http请求的回包太大，进而引发接口超时。一旦发生超时，建议您缩小查询时的Limit参数值，从而降低包的大小，让接口能够及时返回内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeSlowLogData(request *DescribeSlowLogDataRequest) (response *DescribeSlowLogDataResponse, err error) {
    return c.DescribeSlowLogDataWithContext(context.Background(), request)
}

// DescribeSlowLogData
// 条件检索实例的慢日志。只允许查看一个月之内的慢日志。
//
// 使用时需要注意：可能存在单条慢日志太大，导致整个http请求的回包太大，进而引发接口超时。一旦发生超时，建议您缩小查询时的Limit参数值，从而降低包的大小，让接口能够及时返回内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeSlowLogDataWithContext(ctx context.Context, request *DescribeSlowLogDataRequest) (response *DescribeSlowLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSlowLogs")
    
    
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSlowLogs
// 本接口(DescribeSlowLogs)用于获取云数据库实例的慢查询日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    return c.DescribeSlowLogsWithContext(context.Background(), request)
}

// DescribeSlowLogs
// 本接口(DescribeSlowLogs)用于获取云数据库实例的慢查询日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSupportedPrivilegesRequest() (request *DescribeSupportedPrivilegesRequest) {
    request = &DescribeSupportedPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSupportedPrivileges")
    
    
    return
}

func NewDescribeSupportedPrivilegesResponse() (response *DescribeSupportedPrivilegesResponse) {
    response = &DescribeSupportedPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSupportedPrivileges
// 本接口(DescribeSupportedPrivileges)用于查询云数据库的支持的权限信息，包括全局权限，数据库权限，表权限以及列权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeSupportedPrivileges(request *DescribeSupportedPrivilegesRequest) (response *DescribeSupportedPrivilegesResponse, err error) {
    return c.DescribeSupportedPrivilegesWithContext(context.Background(), request)
}

// DescribeSupportedPrivileges
// 本接口(DescribeSupportedPrivileges)用于查询云数据库的支持的权限信息，包括全局权限，数据库权限，表权限以及列权限。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeSupportedPrivilegesWithContext(ctx context.Context, request *DescribeSupportedPrivilegesRequest) (response *DescribeSupportedPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeSupportedPrivilegesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportedPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportedPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTablesRequest() (request *DescribeTablesRequest) {
    request = &DescribeTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTables")
    
    
    return
}

func NewDescribeTablesResponse() (response *DescribeTablesResponse) {
    response = &DescribeTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTables
// 本接口(DescribeTables)用于查询云数据库实例的数据库表信息，仅支持主实例和灾备实例，不支持只读实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    return c.DescribeTablesWithContext(context.Background(), request)
}

// DescribeTables
// 本接口(DescribeTables)用于查询云数据库实例的数据库表信息，仅支持主实例和灾备实例，不支持只读实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsOfInstanceIdsRequest() (request *DescribeTagsOfInstanceIdsRequest) {
    request = &DescribeTagsOfInstanceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTagsOfInstanceIds")
    
    
    return
}

func NewDescribeTagsOfInstanceIdsResponse() (response *DescribeTagsOfInstanceIdsResponse) {
    response = &DescribeTagsOfInstanceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTagsOfInstanceIds
// 本接口(DescribeTagsOfInstanceIds)用于获取云数据库实例的标签信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_TASKERROR = "InternalError.TaskError"
//  INTERNALERROR_TIMEWINDOWERROR = "InternalError.TimeWindowError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTagsOfInstanceIds(request *DescribeTagsOfInstanceIdsRequest) (response *DescribeTagsOfInstanceIdsResponse, err error) {
    return c.DescribeTagsOfInstanceIdsWithContext(context.Background(), request)
}

// DescribeTagsOfInstanceIds
// 本接口(DescribeTagsOfInstanceIds)用于获取云数据库实例的标签信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_TASKERROR = "InternalError.TaskError"
//  INTERNALERROR_TIMEWINDOWERROR = "InternalError.TimeWindowError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTagsOfInstanceIdsWithContext(ctx context.Context, request *DescribeTagsOfInstanceIdsRequest) (response *DescribeTagsOfInstanceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsOfInstanceIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagsOfInstanceIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagsOfInstanceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTasksRequest() (request *DescribeTasksRequest) {
    request = &DescribeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTasks")
    
    
    return
}

func NewDescribeTasksResponse() (response *DescribeTasksResponse) {
    response = &DescribeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTasks
// 本接口(DescribeTasks)用于查询云数据库实例任务列表。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeTasks(request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    return c.DescribeTasksWithContext(context.Background(), request)
}

// DescribeTasks
// 本接口(DescribeTasks)用于查询云数据库实例任务列表。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimeWindowRequest() (request *DescribeTimeWindowRequest) {
    request = &DescribeTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTimeWindow")
    
    
    return
}

func NewDescribeTimeWindowResponse() (response *DescribeTimeWindowResponse) {
    response = &DescribeTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimeWindow
// 本接口(DescribeTimeWindow)用于查询云数据库实例的维护时间窗口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeTimeWindow(request *DescribeTimeWindowRequest) (response *DescribeTimeWindowResponse, err error) {
    return c.DescribeTimeWindowWithContext(context.Background(), request)
}

// DescribeTimeWindow
// 本接口(DescribeTimeWindow)用于查询云数据库实例的维护时间窗口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeTimeWindowWithContext(ctx context.Context, request *DescribeTimeWindowRequest) (response *DescribeTimeWindowResponse, err error) {
    if request == nil {
        request = NewDescribeTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTimeWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUploadedFilesRequest() (request *DescribeUploadedFilesRequest) {
    request = &DescribeUploadedFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeUploadedFiles")
    
    
    return
}

func NewDescribeUploadedFilesResponse() (response *DescribeUploadedFilesResponse) {
    response = &DescribeUploadedFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUploadedFiles
// 本接口(DescribeUploadedFiles)用于查询用户导入的SQL文件列表，全地域公共参数Region均为ap-shanghai。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUploadedFiles(request *DescribeUploadedFilesRequest) (response *DescribeUploadedFilesResponse, err error) {
    return c.DescribeUploadedFilesWithContext(context.Background(), request)
}

// DescribeUploadedFiles
// 本接口(DescribeUploadedFiles)用于查询用户导入的SQL文件列表，全地域公共参数Region均为ap-shanghai。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUploadedFilesWithContext(ctx context.Context, request *DescribeUploadedFilesRequest) (response *DescribeUploadedFilesResponse, err error) {
    if request == nil {
        request = NewDescribeUploadedFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUploadedFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUploadedFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DisassociateSecurityGroups")
    
    
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
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTMATCH = "InternalError.ResourceNotMatch"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
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

func NewInitDBInstancesRequest() (request *InitDBInstancesRequest) {
    request = &InitDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "InitDBInstances")
    
    
    return
}

func NewInitDBInstancesResponse() (response *InitDBInstancesResponse) {
    response = &InitDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitDBInstances
// 本接口(InitDBInstances)用于初始化云数据库实例，包括初始化密码、默认字符集、实例端口号等。该接口已经废弃，在发货接口CreateDBInstance、CreateDBInstanceHour可以直接使用参数Password设置密码，使用参数ParamList设置字符集，使用参数Port设置端口号。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) InitDBInstances(request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    return c.InitDBInstancesWithContext(context.Background(), request)
}

// InitDBInstances
// 本接口(InitDBInstances)用于初始化云数据库实例，包括初始化密码、默认字符集、实例端口号等。该接口已经废弃，在发货接口CreateDBInstance、CreateDBInstanceHour可以直接使用参数Password设置密码，使用参数ParamList设置字符集，使用参数Port设置端口号。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) InitDBInstancesWithContext(ctx context.Context, request *InitDBInstancesRequest) (response *InitDBInstancesResponse, err error) {
    if request == nil {
        request = NewInitDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeInstancesRequest() (request *InquiryPriceUpgradeInstancesRequest) {
    request = &InquiryPriceUpgradeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "InquiryPriceUpgradeInstances")
    
    
    return
}

func NewInquiryPriceUpgradeInstancesResponse() (response *InquiryPriceUpgradeInstancesResponse) {
    response = &InquiryPriceUpgradeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceUpgradeInstances
// 本接口(InquiryPriceUpgradeInstances)用于查询云数据库实例升级的价格，支持查询按量计费或者包年包月实例的升级价格，实例类型支持主实例、灾备实例和只读实例。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) InquiryPriceUpgradeInstances(request *InquiryPriceUpgradeInstancesRequest) (response *InquiryPriceUpgradeInstancesResponse, err error) {
    return c.InquiryPriceUpgradeInstancesWithContext(context.Background(), request)
}

// InquiryPriceUpgradeInstances
// 本接口(InquiryPriceUpgradeInstances)用于查询云数据库实例升级的价格，支持查询按量计费或者包年包月实例的升级价格，实例类型支持主实例、灾备实例和只读实例。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) InquiryPriceUpgradeInstancesWithContext(ctx context.Context, request *InquiryPriceUpgradeInstancesRequest) (response *InquiryPriceUpgradeInstancesResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpgradeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "IsolateDBInstance")
    
    
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateDBInstance
// 本接口(IsolateDBInstance)用于隔离云数据库实例，隔离后不能通过IP和端口访问数据库。隔离的实例可在回收站中进行开机。若为欠费隔离，请尽快进行充值。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDBINSTANCELOCKFAILERROR = "FailedOperation.CdbInstanceLockFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTREFUNDERROR = "OperationDenied.UnSupportRefundError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// 本接口(IsolateDBInstance)用于隔离云数据库实例，隔离后不能通过IP和端口访问数据库。隔离的实例可在回收站中进行开机。若为欠费隔离，请尽快进行充值。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDBINSTANCELOCKFAILERROR = "FailedOperation.CdbInstanceLockFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTREFUNDERROR = "OperationDenied.UnSupportRefundError"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountDescription")
    
    
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountDescription
// 本接口(ModifyAccountDescription)用于修改云数据库账户的备注信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    return c.ModifyAccountDescriptionWithContext(context.Background(), request)
}

// ModifyAccountDescription
// 本接口(ModifyAccountDescription)用于修改云数据库账户的备注信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
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

func NewModifyAccountHostRequest() (request *ModifyAccountHostRequest) {
    request = &ModifyAccountHostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountHost")
    
    
    return
}

func NewModifyAccountHostResponse() (response *ModifyAccountHostResponse) {
    response = &ModifyAccountHostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountHost
// 本接口(ModifyAccountHost)用于修改云数据库账户的主机。
//
// 可能返回的错误码:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_ASYNCTASKSTATUSERROR = "FailedOperation.AsyncTaskStatusError"
//  FAILEDOPERATION_JSONMARSHALERROR = "FailedOperation.JsonMarshalError"
//  FAILEDOPERATION_JSONUNMARSHALERROR = "FailedOperation.JsonUnmarshalError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountHost(request *ModifyAccountHostRequest) (response *ModifyAccountHostResponse, err error) {
    return c.ModifyAccountHostWithContext(context.Background(), request)
}

// ModifyAccountHost
// 本接口(ModifyAccountHost)用于修改云数据库账户的主机。
//
// 可能返回的错误码:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_ASYNCTASKSTATUSERROR = "FailedOperation.AsyncTaskStatusError"
//  FAILEDOPERATION_JSONMARSHALERROR = "FailedOperation.JsonMarshalError"
//  FAILEDOPERATION_JSONUNMARSHALERROR = "FailedOperation.JsonUnmarshalError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountHostWithContext(ctx context.Context, request *ModifyAccountHostRequest) (response *ModifyAccountHostResponse, err error) {
    if request == nil {
        request = NewModifyAccountHostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountHost require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountHostResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountMaxUserConnectionsRequest() (request *ModifyAccountMaxUserConnectionsRequest) {
    request = &ModifyAccountMaxUserConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountMaxUserConnections")
    
    
    return
}

func NewModifyAccountMaxUserConnectionsResponse() (response *ModifyAccountMaxUserConnectionsResponse) {
    response = &ModifyAccountMaxUserConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountMaxUserConnections
// 本接口(ModifyAccountMaxUserConnections)用于修改云数据库账户最大可用连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCTASKSTATUSERROR = "FailedOperation.AsyncTaskStatusError"
//  FAILEDOPERATION_JSONUNMARSHALERROR = "FailedOperation.JsonUnmarshalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ModifyAccountMaxUserConnections(request *ModifyAccountMaxUserConnectionsRequest) (response *ModifyAccountMaxUserConnectionsResponse, err error) {
    return c.ModifyAccountMaxUserConnectionsWithContext(context.Background(), request)
}

// ModifyAccountMaxUserConnections
// 本接口(ModifyAccountMaxUserConnections)用于修改云数据库账户最大可用连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCTASKSTATUSERROR = "FailedOperation.AsyncTaskStatusError"
//  FAILEDOPERATION_JSONUNMARSHALERROR = "FailedOperation.JsonUnmarshalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ModifyAccountMaxUserConnectionsWithContext(ctx context.Context, request *ModifyAccountMaxUserConnectionsRequest) (response *ModifyAccountMaxUserConnectionsResponse, err error) {
    if request == nil {
        request = NewModifyAccountMaxUserConnectionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountMaxUserConnections require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountMaxUserConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPasswordRequest() (request *ModifyAccountPasswordRequest) {
    request = &ModifyAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountPassword")
    
    
    return
}

func NewModifyAccountPasswordResponse() (response *ModifyAccountPasswordResponse) {
    response = &ModifyAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountPassword
// 本接口(ModifyAccountPassword)用于修改云数据库账户的密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (response *ModifyAccountPasswordResponse, err error) {
    return c.ModifyAccountPasswordWithContext(context.Background(), request)
}

// ModifyAccountPassword
// 本接口(ModifyAccountPassword)用于修改云数据库账户的密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPasswordWithContext(ctx context.Context, request *ModifyAccountPasswordRequest) (response *ModifyAccountPasswordResponse, err error) {
    if request == nil {
        request = NewModifyAccountPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAccountPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountPrivilegesRequest() (request *ModifyAccountPrivilegesRequest) {
    request = &ModifyAccountPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAccountPrivileges")
    
    
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
// 注意，修改账号权限时，需要传入该账号下的全量权限信息。用户可以先通过 [查询云数据库账户的权限信息
//
// ](https://cloud.tencent.com/document/api/236/17500) 查询该账号下的全量权限信息，然后进行权限修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) ModifyAccountPrivileges(request *ModifyAccountPrivilegesRequest) (response *ModifyAccountPrivilegesResponse, err error) {
    return c.ModifyAccountPrivilegesWithContext(context.Background(), request)
}

// ModifyAccountPrivileges
// 本接口(ModifyAccountPrivileges)用于修改云数据库的账户的权限信息。
//
// 
//
// 注意，修改账号权限时，需要传入该账号下的全量权限信息。用户可以先通过 [查询云数据库账户的权限信息
//
// ](https://cloud.tencent.com/document/api/236/17500) 查询该账号下的全量权限信息，然后进行权限修改。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
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

func NewModifyAuditConfigRequest() (request *ModifyAuditConfigRequest) {
    request = &ModifyAuditConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAuditConfig")
    
    
    return
}

func NewModifyAuditConfigResponse() (response *ModifyAuditConfigResponse) {
    response = &ModifyAuditConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAuditConfig
// 本接口(ModifyAuditConfig)用于修改云数据库审计策略的服务配置，包括审计日志保存时长等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_MODIFYAUDITSTATUSERROR = "OperationDenied.ModifyAuditStatusError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyAuditConfig(request *ModifyAuditConfigRequest) (response *ModifyAuditConfigResponse, err error) {
    return c.ModifyAuditConfigWithContext(context.Background(), request)
}

// ModifyAuditConfig
// 本接口(ModifyAuditConfig)用于修改云数据库审计策略的服务配置，包括审计日志保存时长等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_MODIFYAUDITSTATUSERROR = "OperationDenied.ModifyAuditStatusError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyAuditConfigWithContext(ctx context.Context, request *ModifyAuditConfigRequest) (response *ModifyAuditConfigResponse, err error) {
    if request == nil {
        request = NewModifyAuditConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditRuleRequest() (request *ModifyAuditRuleRequest) {
    request = &ModifyAuditRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAuditRule")
    
    
    return
}

func NewModifyAuditRuleResponse() (response *ModifyAuditRuleResponse) {
    response = &ModifyAuditRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAuditRule
// 本接口(ModifyAuditRule)用于修改用户的审计规则。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITRULEEXISTERROR = "OperationDenied.AuditRuleExistError"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAuditRule(request *ModifyAuditRuleRequest) (response *ModifyAuditRuleResponse, err error) {
    return c.ModifyAuditRuleWithContext(context.Background(), request)
}

// ModifyAuditRule
// 本接口(ModifyAuditRule)用于修改用户的审计规则。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_AUDITRULEEXISTERROR = "OperationDenied.AuditRuleExistError"
//  OPERATIONDENIED_AUDITRULENOTEXISTERROR = "OperationDenied.AuditRuleNotExistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyAuditRuleWithContext(ctx context.Context, request *ModifyAuditRuleRequest) (response *ModifyAuditRuleResponse, err error) {
    if request == nil {
        request = NewModifyAuditRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAutoRenewFlag")
    
    
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoRenewFlag
// 本接口(ModifyAutoRenewFlag)用于修改云数据库实例的自动续费标记。仅支持包年包月的实例设置自动续费标记。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// 本接口(ModifyAutoRenewFlag)用于修改云数据库实例的自动续费标记。仅支持包年包月的实例设置自动续费标记。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupConfigRequest() (request *ModifyBackupConfigRequest) {
    request = &ModifyBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyBackupConfig")
    
    
    return
}

func NewModifyBackupConfigResponse() (response *ModifyBackupConfigResponse) {
    response = &ModifyBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupConfig
// 本接口(ModifyBackupConfig)用于修改数据库备份配置信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    return c.ModifyBackupConfigWithContext(context.Background(), request)
}

// ModifyBackupConfig
// 本接口(ModifyBackupConfig)用于修改数据库备份配置信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
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

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyBackupDownloadRestriction")
    
    
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
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// 该接口用于修改用户当前地域的备份文件限制下载来源，可以设置内外网均可下载、仅内网可下载，或内网指定的vpc、ip可以下载。
//
// 可能返回的错误码:
//  INTERNALERROR_AUTHERROR = "InternalError.AuthError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyBackupDownloadRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadRestrictionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupDownloadRestriction require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupDownloadRestrictionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBackupEncryptionStatusRequest() (request *ModifyBackupEncryptionStatusRequest) {
    request = &ModifyBackupEncryptionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyBackupEncryptionStatus")
    
    
    return
}

func NewModifyBackupEncryptionStatusResponse() (response *ModifyBackupEncryptionStatusResponse) {
    response = &ModifyBackupEncryptionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyBackupEncryptionStatus
// 本接口(ModifyBackupEncryptionStatus)用于设置实例备份文件是否加密。 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTOPERATIONDENIED = "OperationDenied.AccountOperationDenied"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
func (c *Client) ModifyBackupEncryptionStatus(request *ModifyBackupEncryptionStatusRequest) (response *ModifyBackupEncryptionStatusResponse, err error) {
    return c.ModifyBackupEncryptionStatusWithContext(context.Background(), request)
}

// ModifyBackupEncryptionStatus
// 本接口(ModifyBackupEncryptionStatus)用于设置实例备份文件是否加密。 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ACCOUNTOPERATIONDENIED = "OperationDenied.AccountOperationDenied"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
func (c *Client) ModifyBackupEncryptionStatusWithContext(ctx context.Context, request *ModifyBackupEncryptionStatusRequest) (response *ModifyBackupEncryptionStatusResponse, err error) {
    if request == nil {
        request = NewModifyBackupEncryptionStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupEncryptionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupEncryptionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyConnectionPoolRequest() (request *ModifyCDBProxyConnectionPoolRequest) {
    request = &ModifyCDBProxyConnectionPoolRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxyConnectionPool")
    
    
    return
}

func NewModifyCDBProxyConnectionPoolResponse() (response *ModifyCDBProxyConnectionPoolResponse) {
    response = &ModifyCDBProxyConnectionPoolResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxyConnectionPool
// 请求该接口配置数据库连接池；支持的连接池配置请求DescribeProxyConnectionPoolConf接口获取。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyConnectionPool(request *ModifyCDBProxyConnectionPoolRequest) (response *ModifyCDBProxyConnectionPoolResponse, err error) {
    return c.ModifyCDBProxyConnectionPoolWithContext(context.Background(), request)
}

// ModifyCDBProxyConnectionPool
// 请求该接口配置数据库连接池；支持的连接池配置请求DescribeProxyConnectionPoolConf接口获取。
//
// 可能返回的错误码:
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyConnectionPoolWithContext(ctx context.Context, request *ModifyCDBProxyConnectionPoolRequest) (response *ModifyCDBProxyConnectionPoolResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyConnectionPoolRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxyConnectionPool require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyConnectionPoolResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyDescRequest() (request *ModifyCDBProxyDescRequest) {
    request = &ModifyCDBProxyDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxyDesc")
    
    
    return
}

func NewModifyCDBProxyDescResponse() (response *ModifyCDBProxyDescResponse) {
    response = &ModifyCDBProxyDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxyDesc
// 修改数据库代理描述
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyDesc(request *ModifyCDBProxyDescRequest) (response *ModifyCDBProxyDescResponse, err error) {
    return c.ModifyCDBProxyDescWithContext(context.Background(), request)
}

// ModifyCDBProxyDesc
// 修改数据库代理描述
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCDBProxyDescWithContext(ctx context.Context, request *ModifyCDBProxyDescRequest) (response *ModifyCDBProxyDescResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyDescRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxyDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyDescResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCDBProxyVipVPortRequest() (request *ModifyCDBProxyVipVPortRequest) {
    request = &ModifyCDBProxyVipVPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCDBProxyVipVPort")
    
    
    return
}

func NewModifyCDBProxyVipVPortResponse() (response *ModifyCDBProxyVipVPortResponse) {
    response = &ModifyCDBProxyVipVPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCDBProxyVipVPort
// 修改数据库代理VIP或端口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCDBProxyVipVPort(request *ModifyCDBProxyVipVPortRequest) (response *ModifyCDBProxyVipVPortResponse, err error) {
    return c.ModifyCDBProxyVipVPortWithContext(context.Background(), request)
}

// ModifyCDBProxyVipVPort
// 修改数据库代理VIP或端口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCDBProxyVipVPortWithContext(ctx context.Context, request *ModifyCDBProxyVipVPortRequest) (response *ModifyCDBProxyVipVPortResponse, err error) {
    if request == nil {
        request = NewModifyCDBProxyVipVPortRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCDBProxyVipVPort require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCDBProxyVipVPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNameRequest() (request *ModifyDBInstanceNameRequest) {
    request = &ModifyDBInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceName")
    
    
    return
}

func NewModifyDBInstanceNameResponse() (response *ModifyDBInstanceNameResponse) {
    response = &ModifyDBInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceName
// 本接口(ModifyDBInstanceName)用于修改云数据库实例的名称。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyDBInstanceName(request *ModifyDBInstanceNameRequest) (response *ModifyDBInstanceNameResponse, err error) {
    return c.ModifyDBInstanceNameWithContext(context.Background(), request)
}

// ModifyDBInstanceName
// 本接口(ModifyDBInstanceName)用于修改云数据库实例的名称。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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

func NewModifyDBInstanceProjectRequest() (request *ModifyDBInstanceProjectRequest) {
    request = &ModifyDBInstanceProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceProject")
    
    
    return
}

func NewModifyDBInstanceProjectResponse() (response *ModifyDBInstanceProjectResponse) {
    response = &ModifyDBInstanceProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceProject
// 本接口(ModifyDBInstanceProject)用于修改云数据库实例的所属项目。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyDBInstanceProject(request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    return c.ModifyDBInstanceProjectWithContext(context.Background(), request)
}

// ModifyDBInstanceProject
// 本接口(ModifyDBInstanceProject)用于修改云数据库实例的所属项目。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
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
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DFWERROR = "InternalError.DfwError"
//  INTERNALERROR_RESOURCENOTUNIQUE = "InternalError.ResourceNotUnique"
//  INTERNALERROR_SECURITYGROUPERROR = "InternalError.SecurityGroupError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
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

func NewModifyDBInstanceVipVportRequest() (request *ModifyDBInstanceVipVportRequest) {
    request = &ModifyDBInstanceVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceVipVport")
    
    
    return
}

func NewModifyDBInstanceVipVportResponse() (response *ModifyDBInstanceVipVportResponse) {
    response = &ModifyDBInstanceVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDBInstanceVipVport
// 本接口(ModifyDBInstanceVipVport)用于修改云数据库实例的IP和端口号，也可进行基础网络转 VPC 网络和 VPC 网络下的子网变更。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceVipVport(request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    return c.ModifyDBInstanceVipVportWithContext(context.Background(), request)
}

// ModifyDBInstanceVipVport
// 本接口(ModifyDBInstanceVipVport)用于修改云数据库实例的IP和端口号，也可进行基础网络转 VPC 网络和 VPC 网络下的子网变更。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceVipVportWithContext(ctx context.Context, request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceVipVportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceVipVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamRequest() (request *ModifyInstanceParamRequest) {
    request = &ModifyInstanceParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstanceParam")
    
    
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
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
func (c *Client) ModifyInstanceParam(request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    return c.ModifyInstanceParamWithContext(context.Background(), request)
}

// ModifyInstanceParam
// 本接口(ModifyInstanceParam)用于修改云数据库实例的参数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
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

func NewModifyInstancePasswordComplexityRequest() (request *ModifyInstancePasswordComplexityRequest) {
    request = &ModifyInstancePasswordComplexityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstancePasswordComplexity")
    
    
    return
}

func NewModifyInstancePasswordComplexityResponse() (response *ModifyInstancePasswordComplexityResponse) {
    response = &ModifyInstancePasswordComplexityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancePasswordComplexity
// 本接口(ModifyInstancePasswordComplexity)用于修改云数据库实例的密码复杂度。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstancePasswordComplexity(request *ModifyInstancePasswordComplexityRequest) (response *ModifyInstancePasswordComplexityResponse, err error) {
    return c.ModifyInstancePasswordComplexityWithContext(context.Background(), request)
}

// ModifyInstancePasswordComplexity
// 本接口(ModifyInstancePasswordComplexity)用于修改云数据库实例的密码复杂度。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstancePasswordComplexityWithContext(ctx context.Context, request *ModifyInstancePasswordComplexityRequest) (response *ModifyInstancePasswordComplexityResponse, err error) {
    if request == nil {
        request = NewModifyInstancePasswordComplexityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancePasswordComplexity require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancePasswordComplexityResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceTagRequest() (request *ModifyInstanceTagRequest) {
    request = &ModifyInstanceTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyInstanceTag")
    
    
    return
}

func NewModifyInstanceTagResponse() (response *ModifyInstanceTagResponse) {
    response = &ModifyInstanceTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceTag
// 本接口(ModifyInstanceTag)用于对实例标签进行添加、修改或者删除。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyInstanceTag(request *ModifyInstanceTagRequest) (response *ModifyInstanceTagResponse, err error) {
    return c.ModifyInstanceTagWithContext(context.Background(), request)
}

// ModifyInstanceTag
// 本接口(ModifyInstanceTag)用于对实例标签进行添加、修改或者删除。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_BACKUPERROR = "CdbError.BackupError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TAGERROR = "InternalError.TagError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyInstanceTagWithContext(ctx context.Context, request *ModifyInstanceTagRequest) (response *ModifyInstanceTagResponse, err error) {
    if request == nil {
        request = NewModifyInstanceTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceTagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyLocalBinlogConfigRequest() (request *ModifyLocalBinlogConfigRequest) {
    request = &ModifyLocalBinlogConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyLocalBinlogConfig")
    
    
    return
}

func NewModifyLocalBinlogConfigResponse() (response *ModifyLocalBinlogConfigResponse) {
    response = &ModifyLocalBinlogConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyLocalBinlogConfig
// 该接口用于修改实例本地binlog保留策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
func (c *Client) ModifyLocalBinlogConfig(request *ModifyLocalBinlogConfigRequest) (response *ModifyLocalBinlogConfigResponse, err error) {
    return c.ModifyLocalBinlogConfigWithContext(context.Background(), request)
}

// ModifyLocalBinlogConfig
// 该接口用于修改实例本地binlog保留策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
func (c *Client) ModifyLocalBinlogConfigWithContext(ctx context.Context, request *ModifyLocalBinlogConfigRequest) (response *ModifyLocalBinlogConfigResponse, err error) {
    if request == nil {
        request = NewModifyLocalBinlogConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyLocalBinlogConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyLocalBinlogConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNameOrDescByDpIdRequest() (request *ModifyNameOrDescByDpIdRequest) {
    request = &ModifyNameOrDescByDpIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyNameOrDescByDpId")
    
    
    return
}

func NewModifyNameOrDescByDpIdResponse() (response *ModifyNameOrDescByDpIdResponse) {
    response = &ModifyNameOrDescByDpIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNameOrDescByDpId
// 修改置放群组的名称或者描述
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) ModifyNameOrDescByDpId(request *ModifyNameOrDescByDpIdRequest) (response *ModifyNameOrDescByDpIdResponse, err error) {
    return c.ModifyNameOrDescByDpIdWithContext(context.Background(), request)
}

// ModifyNameOrDescByDpId
// 修改置放群组的名称或者描述
//
// 可能返回的错误码:
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DEPLOYGROUPNOTEMPTY = "InvalidParameter.DeployGroupNotEmpty"
//  INVALIDPARAMETER_OVERDEPLOYGROUPQUOTA = "InvalidParameter.OverDeployGroupQuota"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
//  INVALIDPARAMETER_RESOURCENOTFOUND = "InvalidParameter.ResourceNotFound"
func (c *Client) ModifyNameOrDescByDpIdWithContext(ctx context.Context, request *ModifyNameOrDescByDpIdRequest) (response *ModifyNameOrDescByDpIdResponse, err error) {
    if request == nil {
        request = NewModifyNameOrDescByDpIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNameOrDescByDpId require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNameOrDescByDpIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyParamTemplateRequest() (request *ModifyParamTemplateRequest) {
    request = &ModifyParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyParamTemplate")
    
    
    return
}

func NewModifyParamTemplateResponse() (response *ModifyParamTemplateResponse) {
    response = &ModifyParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyParamTemplate
// 该接口（ModifyParamTemplate）用于修改参数模板，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    return c.ModifyParamTemplateWithContext(context.Background(), request)
}

// ModifyParamTemplate
// 该接口（ModifyParamTemplate）用于修改参数模板，全地域公共参数Region均为ap-guangzhou。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESOURCEEXISTS = "InvalidParameter.ResourceExists"
func (c *Client) ModifyParamTemplateWithContext(ctx context.Context, request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    if request == nil {
        request = NewModifyParamTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRemoteBackupConfigRequest() (request *ModifyRemoteBackupConfigRequest) {
    request = &ModifyRemoteBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyRemoteBackupConfig")
    
    
    return
}

func NewModifyRemoteBackupConfigResponse() (response *ModifyRemoteBackupConfigResponse) {
    response = &ModifyRemoteBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRemoteBackupConfig
// 本接口(ModifyRemoteBackupConfig)用于修改数据库异地备份配置信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyRemoteBackupConfig(request *ModifyRemoteBackupConfigRequest) (response *ModifyRemoteBackupConfigResponse, err error) {
    return c.ModifyRemoteBackupConfigWithContext(context.Background(), request)
}

// ModifyRemoteBackupConfig
// 本接口(ModifyRemoteBackupConfig)用于修改数据库异地备份配置信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyRemoteBackupConfigWithContext(ctx context.Context, request *ModifyRemoteBackupConfigRequest) (response *ModifyRemoteBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyRemoteBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRemoteBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRemoteBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoGroupInfoRequest() (request *ModifyRoGroupInfoRequest) {
    request = &ModifyRoGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyRoGroupInfo")
    
    
    return
}

func NewModifyRoGroupInfoResponse() (response *ModifyRoGroupInfoResponse) {
    response = &ModifyRoGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRoGroupInfo
// 本接口（ModifyRoGroupInfo）用于更新云数据库只读组的信息。包括设置实例延迟超限剔除策略，设置只读实例读权重，设置复制延迟时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_CONFLICTROSTATUS = "OperationDenied.ConflictRoStatus"
//  OPERATIONDENIED_CONFLICTSTATUS = "OperationDenied.ConflictStatus"
//  OPERATIONDENIED_DELAYREPLICATIONRUNNING = "OperationDenied.DelayReplicationRunning"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) ModifyRoGroupInfo(request *ModifyRoGroupInfoRequest) (response *ModifyRoGroupInfoResponse, err error) {
    return c.ModifyRoGroupInfoWithContext(context.Background(), request)
}

// ModifyRoGroupInfo
// 本接口（ModifyRoGroupInfo）用于更新云数据库只读组的信息。包括设置实例延迟超限剔除策略，设置只读实例读权重，设置复制延迟时间等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_CONFLICTROSTATUS = "OperationDenied.ConflictRoStatus"
//  OPERATIONDENIED_CONFLICTSTATUS = "OperationDenied.ConflictStatus"
//  OPERATIONDENIED_DELAYREPLICATIONRUNNING = "OperationDenied.DelayReplicationRunning"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) ModifyRoGroupInfoWithContext(ctx context.Context, request *ModifyRoGroupInfoRequest) (response *ModifyRoGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyRoGroupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTimeWindowRequest() (request *ModifyTimeWindowRequest) {
    request = &ModifyTimeWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyTimeWindow")
    
    
    return
}

func NewModifyTimeWindowResponse() (response *ModifyTimeWindowResponse) {
    response = &ModifyTimeWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTimeWindow
// 本接口(ModifyTimeWindow)用于更新云数据库实例的维护时间窗口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyTimeWindow(request *ModifyTimeWindowRequest) (response *ModifyTimeWindowResponse, err error) {
    return c.ModifyTimeWindowWithContext(context.Background(), request)
}

// ModifyTimeWindow
// 本接口(ModifyTimeWindow)用于更新云数据库实例的维护时间窗口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyTimeWindowWithContext(ctx context.Context, request *ModifyTimeWindowRequest) (response *ModifyTimeWindowResponse, err error) {
    if request == nil {
        request = NewModifyTimeWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTimeWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedInstancesRequest() (request *OfflineIsolatedInstancesRequest) {
    request = &OfflineIsolatedInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OfflineIsolatedInstances")
    
    
    return
}

func NewOfflineIsolatedInstancesResponse() (response *OfflineIsolatedInstancesResponse) {
    response = &OfflineIsolatedInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OfflineIsolatedInstances
// 本接口(OfflineIsolatedInstances)用于立即下线隔离状态的云数据库实例。进行操作的实例状态必须为隔离状态，即通过 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询到 Status 值为 5 的实例。
//
// 
//
// 该接口为异步操作，部分资源的回收可能存在延迟。您可以通过使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口，指定实例 InstanceId 和状态 Status 为 [5,6,7] 进行查询，若返回实例为空，则实例资源已全部释放。
//
// 
//
// 注意，实例下线后，相关资源和数据将无法找回，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OfflineIsolatedInstances(request *OfflineIsolatedInstancesRequest) (response *OfflineIsolatedInstancesResponse, err error) {
    return c.OfflineIsolatedInstancesWithContext(context.Background(), request)
}

// OfflineIsolatedInstances
// 本接口(OfflineIsolatedInstances)用于立即下线隔离状态的云数据库实例。进行操作的实例状态必须为隔离状态，即通过 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口查询到 Status 值为 5 的实例。
//
// 
//
// 该接口为异步操作，部分资源的回收可能存在延迟。您可以通过使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口，指定实例 InstanceId 和状态 Status 为 [5,6,7] 进行查询，若返回实例为空，则实例资源已全部释放。
//
// 
//
// 注意，实例下线后，相关资源和数据将无法找回，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OfflineIsolatedInstancesWithContext(ctx context.Context, request *OfflineIsolatedInstancesRequest) (response *OfflineIsolatedInstancesResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAuditServiceRequest() (request *OpenAuditServiceRequest) {
    request = &OpenAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OpenAuditService")
    
    
    return
}

func NewOpenAuditServiceResponse() (response *OpenAuditServiceResponse) {
    response = &OpenAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenAuditService
// CDB实例开通审计服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// CDB实例开通审计服务
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
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

func NewOpenDBInstanceEncryptionRequest() (request *OpenDBInstanceEncryptionRequest) {
    request = &OpenDBInstanceEncryptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OpenDBInstanceEncryption")
    
    
    return
}

func NewOpenDBInstanceEncryptionResponse() (response *OpenDBInstanceEncryptionResponse) {
    response = &OpenDBInstanceEncryptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenDBInstanceEncryption
// 本接口(OpenDBInstanceEncryption)用于启用实例数据存储加密功能，支持用户指定自定义密钥。
//
// 
//
// 注意，启用实例数据存储加密之前，需要进行以下操作：
//
// 
//
// 1、进行 [实例初始化](https://cloud.tencent.com/document/api/236/15873) 操作；
//
// 
//
// 2、开启 [KMS服务](https://console.cloud.tencent.com/kms2)；
//
// 
//
// 3、对云数据库(MySQL)[授予访问KMS密钥的权限](https://console.cloud.tencent.com/cam/role)，角色名为MySQL_QCSRole，预设策略名为QcloudAccessForMySQLRole；
//
// 
//
// 该 API 耗时可能到10s，客户端可能超时，如果调用 API 返回 InternalError ，请您调用DescribeDBInstanceInfo 确认后端加密是否开通成功。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) OpenDBInstanceEncryption(request *OpenDBInstanceEncryptionRequest) (response *OpenDBInstanceEncryptionResponse, err error) {
    return c.OpenDBInstanceEncryptionWithContext(context.Background(), request)
}

// OpenDBInstanceEncryption
// 本接口(OpenDBInstanceEncryption)用于启用实例数据存储加密功能，支持用户指定自定义密钥。
//
// 
//
// 注意，启用实例数据存储加密之前，需要进行以下操作：
//
// 
//
// 1、进行 [实例初始化](https://cloud.tencent.com/document/api/236/15873) 操作；
//
// 
//
// 2、开启 [KMS服务](https://console.cloud.tencent.com/kms2)；
//
// 
//
// 3、对云数据库(MySQL)[授予访问KMS密钥的权限](https://console.cloud.tencent.com/cam/role)，角色名为MySQL_QCSRole，预设策略名为QcloudAccessForMySQLRole；
//
// 
//
// 该 API 耗时可能到10s，客户端可能超时，如果调用 API 返回 InternalError ，请您调用DescribeDBInstanceInfo 确认后端加密是否开通成功。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) OpenDBInstanceEncryptionWithContext(ctx context.Context, request *OpenDBInstanceEncryptionRequest) (response *OpenDBInstanceEncryptionResponse, err error) {
    if request == nil {
        request = NewOpenDBInstanceEncryptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenDBInstanceEncryption require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenDBInstanceEncryptionResponse()
    err = c.Send(request, response)
    return
}

func NewOpenDBInstanceGTIDRequest() (request *OpenDBInstanceGTIDRequest) {
    request = &OpenDBInstanceGTIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OpenDBInstanceGTID")
    
    
    return
}

func NewOpenDBInstanceGTIDResponse() (response *OpenDBInstanceGTIDResponse) {
    response = &OpenDBInstanceGTIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenDBInstanceGTID
// 本接口(OpenDBInstanceGTID)用于开启云数据库实例的 GTID，只支持版本为 5.6 以及以上的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OpenDBInstanceGTID(request *OpenDBInstanceGTIDRequest) (response *OpenDBInstanceGTIDResponse, err error) {
    return c.OpenDBInstanceGTIDWithContext(context.Background(), request)
}

// OpenDBInstanceGTID
// 本接口(OpenDBInstanceGTID)用于开启云数据库实例的 GTID，只支持版本为 5.6 以及以上的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OpenDBInstanceGTIDWithContext(ctx context.Context, request *OpenDBInstanceGTIDRequest) (response *OpenDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewOpenDBInstanceGTIDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenDBInstanceGTID require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenDBInstanceGTIDResponse()
    err = c.Send(request, response)
    return
}

func NewOpenWanServiceRequest() (request *OpenWanServiceRequest) {
    request = &OpenWanServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OpenWanService")
    
    
    return
}

func NewOpenWanServiceResponse() (response *OpenWanServiceResponse) {
    response = &OpenWanServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OpenWanService
// 本接口(OpenWanService)用于开通实例外网访问。
//
// 
//
// 注意，实例开通外网访问之前，需要先将实例进行 [实例初始化](https://cloud.tencent.com/document/api/236/15873) 操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) OpenWanService(request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    return c.OpenWanServiceWithContext(context.Background(), request)
}

// OpenWanService
// 本接口(OpenWanService)用于开通实例外网访问。
//
// 
//
// 注意，实例开通外网访问之前，需要先将实例进行 [实例初始化](https://cloud.tencent.com/document/api/236/15873) 操作。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) OpenWanServiceWithContext(ctx context.Context, request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    if request == nil {
        request = NewOpenWanServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWanService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWanServiceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCDBProxyRequest() (request *QueryCDBProxyRequest) {
    request = &QueryCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "QueryCDBProxy")
    
    
    return
}

func NewQueryCDBProxyResponse() (response *QueryCDBProxyResponse) {
    response = &QueryCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCDBProxy
// 查询代理详情
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) QueryCDBProxy(request *QueryCDBProxyRequest) (response *QueryCDBProxyResponse, err error) {
    return c.QueryCDBProxyWithContext(context.Background(), request)
}

// QueryCDBProxy
// 查询代理详情
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) QueryCDBProxyWithContext(ctx context.Context, request *QueryCDBProxyRequest) (response *QueryCDBProxyResponse, err error) {
    if request == nil {
        request = NewQueryCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseIsolatedDBInstancesRequest() (request *ReleaseIsolatedDBInstancesRequest) {
    request = &ReleaseIsolatedDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ReleaseIsolatedDBInstances")
    
    
    return
}

func NewReleaseIsolatedDBInstancesResponse() (response *ReleaseIsolatedDBInstancesResponse) {
    response = &ReleaseIsolatedDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReleaseIsolatedDBInstances
// 本接口（ReleaseIsolatedDBInstances）用于恢复已隔离云数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONINCONFLICTERR = "FailedOperation.OperationInConflictErr"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ReleaseIsolatedDBInstances(request *ReleaseIsolatedDBInstancesRequest) (response *ReleaseIsolatedDBInstancesResponse, err error) {
    return c.ReleaseIsolatedDBInstancesWithContext(context.Background(), request)
}

// ReleaseIsolatedDBInstances
// 本接口（ReleaseIsolatedDBInstances）用于恢复已隔离云数据库实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONINCONFLICTERR = "FailedOperation.OperationInConflictErr"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ReleaseIsolatedDBInstancesWithContext(ctx context.Context, request *ReleaseIsolatedDBInstancesRequest) (response *ReleaseIsolatedDBInstancesResponse, err error) {
    if request == nil {
        request = NewReleaseIsolatedDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseIsolatedDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseIsolatedDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewReloadBalanceProxyNodeRequest() (request *ReloadBalanceProxyNodeRequest) {
    request = &ReloadBalanceProxyNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ReloadBalanceProxyNode")
    
    
    return
}

func NewReloadBalanceProxyNodeResponse() (response *ReloadBalanceProxyNodeResponse) {
    response = &ReloadBalanceProxyNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReloadBalanceProxyNode
// 重新负载均衡数据库代理
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
func (c *Client) ReloadBalanceProxyNode(request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    return c.ReloadBalanceProxyNodeWithContext(context.Background(), request)
}

// ReloadBalanceProxyNode
// 重新负载均衡数据库代理
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
func (c *Client) ReloadBalanceProxyNodeWithContext(ctx context.Context, request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    if request == nil {
        request = NewReloadBalanceProxyNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReloadBalanceProxyNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewReloadBalanceProxyNodeResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBInstanceRequest() (request *RenewDBInstanceRequest) {
    request = &RenewDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "RenewDBInstance")
    
    
    return
}

func NewRenewDBInstanceResponse() (response *RenewDBInstanceResponse) {
    response = &RenewDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewDBInstance
// 本接口(RenewDBInstance)用于续费云数据库实例，支持付费模式为包年包月的实例。按量计费实例可通过该接口续费为包年包月的实例。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) RenewDBInstance(request *RenewDBInstanceRequest) (response *RenewDBInstanceResponse, err error) {
    return c.RenewDBInstanceWithContext(context.Background(), request)
}

// RenewDBInstance
// 本接口(RenewDBInstance)用于续费云数据库实例，支持付费模式为包年包月的实例。按量计费实例可通过该接口续费为包年包月的实例。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
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

func NewResetRootAccountRequest() (request *ResetRootAccountRequest) {
    request = &ResetRootAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ResetRootAccount")
    
    
    return
}

func NewResetRootAccountResponse() (response *ResetRootAccountResponse) {
    response = &ResetRootAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetRootAccount
// 重置实例ROOT账号，初始化账号权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ResetRootAccount(request *ResetRootAccountRequest) (response *ResetRootAccountResponse, err error) {
    return c.ResetRootAccountWithContext(context.Background(), request)
}

// ResetRootAccount
// 重置实例ROOT账号，初始化账号权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ResetRootAccountWithContext(ctx context.Context, request *ResetRootAccountRequest) (response *ResetRootAccountResponse, err error) {
    if request == nil {
        request = NewResetRootAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetRootAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetRootAccountResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstancesRequest() (request *RestartDBInstancesRequest) {
    request = &RestartDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "RestartDBInstances")
    
    
    return
}

func NewRestartDBInstancesResponse() (response *RestartDBInstancesResponse) {
    response = &RestartDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartDBInstances
// 本接口(RestartDBInstances)用于重启云数据库实例。
//
// 
//
// 注意：
//
// 1、本接口只支持主实例进行重启操作；
//
// 2、实例状态必须为正常，并且没有其他异步任务在执行中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) RestartDBInstances(request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    return c.RestartDBInstancesWithContext(context.Background(), request)
}

// RestartDBInstances
// 本接口(RestartDBInstances)用于重启云数据库实例。
//
// 
//
// 注意：
//
// 1、本接口只支持主实例进行重启操作；
//
// 2、实例状态必须为正常，并且没有其他异步任务在执行中。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) RestartDBInstancesWithContext(ctx context.Context, request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    if request == nil {
        request = NewRestartDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStartBatchRollbackRequest() (request *StartBatchRollbackRequest) {
    request = &StartBatchRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StartBatchRollback")
    
    
    return
}

func NewStartBatchRollbackResponse() (response *StartBatchRollbackResponse) {
    response = &StartBatchRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartBatchRollback
// 该接口（StartBatchRollback）用于批量回档云数据库实例的库表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  OVERQUOTA = "OverQuota"
func (c *Client) StartBatchRollback(request *StartBatchRollbackRequest) (response *StartBatchRollbackResponse, err error) {
    return c.StartBatchRollbackWithContext(context.Background(), request)
}

// StartBatchRollback
// 该接口（StartBatchRollback）用于批量回档云数据库实例的库表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  OVERQUOTA = "OverQuota"
func (c *Client) StartBatchRollbackWithContext(ctx context.Context, request *StartBatchRollbackRequest) (response *StartBatchRollbackResponse, err error) {
    if request == nil {
        request = NewStartBatchRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBatchRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBatchRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewStartReplicationRequest() (request *StartReplicationRequest) {
    request = &StartReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StartReplication")
    
    
    return
}

func NewStartReplicationResponse() (response *StartReplicationResponse) {
    response = &StartReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartReplication
// 开启 RO 复制，从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StartReplication(request *StartReplicationRequest) (response *StartReplicationResponse, err error) {
    return c.StartReplicationWithContext(context.Background(), request)
}

// StartReplication
// 开启 RO 复制，从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StartReplicationWithContext(ctx context.Context, request *StartReplicationRequest) (response *StartReplicationResponse, err error) {
    if request == nil {
        request = NewStartReplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartReplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartReplicationResponse()
    err = c.Send(request, response)
    return
}

func NewStopDBImportJobRequest() (request *StopDBImportJobRequest) {
    request = &StopDBImportJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StopDBImportJob")
    
    
    return
}

func NewStopDBImportJobResponse() (response *StopDBImportJobResponse) {
    response = &StopDBImportJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopDBImportJob
// 本接口(StopDBImportJob)用于终止数据导入任务。
//
// 可能返回的错误码:
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  INTERNALERROR_IMPORTERROR = "InternalError.ImportError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) StopDBImportJob(request *StopDBImportJobRequest) (response *StopDBImportJobResponse, err error) {
    return c.StopDBImportJobWithContext(context.Background(), request)
}

// StopDBImportJob
// 本接口(StopDBImportJob)用于终止数据导入任务。
//
// 可能返回的错误码:
//  CDBERROR_IMPORTERROR = "CdbError.ImportError"
//  INTERNALERROR_IMPORTERROR = "InternalError.ImportError"
//  INTERNALERROR_TASKFRAMEERROR = "InternalError.TaskFrameError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDASYNCREQUESTID = "InvalidParameter.InvalidAsyncRequestId"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) StopDBImportJobWithContext(ctx context.Context, request *StopDBImportJobRequest) (response *StopDBImportJobResponse, err error) {
    if request == nil {
        request = NewStopDBImportJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDBImportJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDBImportJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopReplicationRequest() (request *StopReplicationRequest) {
    request = &StopReplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StopReplication")
    
    
    return
}

func NewStopReplicationResponse() (response *StopReplicationResponse) {
    response = &StopReplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopReplication
// 停止 RO 复制，中断从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDELAYRO = "FailedOperation.NotDelayRo"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StopReplication(request *StopReplicationRequest) (response *StopReplicationResponse, err error) {
    return c.StopReplicationWithContext(context.Background(), request)
}

// StopReplication
// 停止 RO 复制，中断从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDELAYRO = "FailedOperation.NotDelayRo"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) StopReplicationWithContext(ctx context.Context, request *StopReplicationRequest) (response *StopReplicationResponse, err error) {
    if request == nil {
        request = NewStopReplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopReplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopReplicationResponse()
    err = c.Send(request, response)
    return
}

func NewStopRollbackRequest() (request *StopRollbackRequest) {
    request = &StopRollbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StopRollback")
    
    
    return
}

func NewStopRollbackResponse() (response *StopRollbackResponse) {
    response = &StopRollbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopRollback
// 本接口(StopRollback) 用于撤销实例正在进行的回档任务，该接口返回一个异步任务id。 撤销结果可以通过 DescribeAsyncRequestInfo 查询任务的执行情况。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_ASYNCREQUESTERROR = "InternalError.AsyncRequestError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StopRollback(request *StopRollbackRequest) (response *StopRollbackResponse, err error) {
    return c.StopRollbackWithContext(context.Background(), request)
}

// StopRollback
// 本接口(StopRollback) 用于撤销实例正在进行的回档任务，该接口返回一个异步任务id。 撤销结果可以通过 DescribeAsyncRequestInfo 查询任务的执行情况。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  CDBERROR_TASKERROR = "CdbError.TaskError"
//  INTERNALERROR_ASYNCREQUESTERROR = "InternalError.AsyncRequestError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) StopRollbackWithContext(ctx context.Context, request *StopRollbackRequest) (response *StopRollbackResponse, err error) {
    if request == nil {
        request = NewStopRollbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchCDBProxyRequest() (request *SwitchCDBProxyRequest) {
    request = &SwitchCDBProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchCDBProxy")
    
    
    return
}

func NewSwitchCDBProxyResponse() (response *SwitchCDBProxyResponse) {
    response = &SwitchCDBProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchCDBProxy
// 数据库代理配置变更或则升级版本后手动发起立即切换
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) SwitchCDBProxy(request *SwitchCDBProxyRequest) (response *SwitchCDBProxyResponse, err error) {
    return c.SwitchCDBProxyWithContext(context.Background(), request)
}

// SwitchCDBProxy
// 数据库代理配置变更或则升级版本后手动发起立即切换
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) SwitchCDBProxyWithContext(ctx context.Context, request *SwitchCDBProxyRequest) (response *SwitchCDBProxyResponse, err error) {
    if request == nil {
        request = NewSwitchCDBProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDBInstanceMasterSlaveRequest() (request *SwitchDBInstanceMasterSlaveRequest) {
    request = &SwitchDBInstanceMasterSlaveRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchDBInstanceMasterSlave")
    
    
    return
}

func NewSwitchDBInstanceMasterSlaveResponse() (response *SwitchDBInstanceMasterSlaveResponse) {
    response = &SwitchDBInstanceMasterSlaveResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDBInstanceMasterSlave
// 该接口 (SwitchDBInstanceMasterSlave) 支持用户主动切换实例主从角色。
//
// 可能返回的错误码:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) SwitchDBInstanceMasterSlave(request *SwitchDBInstanceMasterSlaveRequest) (response *SwitchDBInstanceMasterSlaveResponse, err error) {
    return c.SwitchDBInstanceMasterSlaveWithContext(context.Background(), request)
}

// SwitchDBInstanceMasterSlave
// 该接口 (SwitchDBInstanceMasterSlave) 支持用户主动切换实例主从角色。
//
// 可能返回的错误码:
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) SwitchDBInstanceMasterSlaveWithContext(ctx context.Context, request *SwitchDBInstanceMasterSlaveRequest) (response *SwitchDBInstanceMasterSlaveResponse, err error) {
    if request == nil {
        request = NewSwitchDBInstanceMasterSlaveRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDBInstanceMasterSlave require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDBInstanceMasterSlaveResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchDrInstanceToMasterRequest() (request *SwitchDrInstanceToMasterRequest) {
    request = &SwitchDrInstanceToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchDrInstanceToMaster")
    
    
    return
}

func NewSwitchDrInstanceToMasterResponse() (response *SwitchDrInstanceToMasterResponse) {
    response = &SwitchDrInstanceToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchDrInstanceToMaster
// 本接口(SwitchDrInstanceToMaster)用于将云数据库灾备实例切换为主实例，注意请求必须发到灾备实例所在的地域。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) SwitchDrInstanceToMaster(request *SwitchDrInstanceToMasterRequest) (response *SwitchDrInstanceToMasterResponse, err error) {
    return c.SwitchDrInstanceToMasterWithContext(context.Background(), request)
}

// SwitchDrInstanceToMaster
// 本接口(SwitchDrInstanceToMaster)用于将云数据库灾备实例切换为主实例，注意请求必须发到灾备实例所在的地域。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) SwitchDrInstanceToMasterWithContext(ctx context.Context, request *SwitchDrInstanceToMasterRequest) (response *SwitchDrInstanceToMasterResponse, err error) {
    if request == nil {
        request = NewSwitchDrInstanceToMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchDrInstanceToMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchDrInstanceToMasterResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchForUpgradeRequest() (request *SwitchForUpgradeRequest) {
    request = &SwitchForUpgradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SwitchForUpgrade")
    
    
    return
}

func NewSwitchForUpgradeResponse() (response *SwitchForUpgradeResponse) {
    response = &SwitchForUpgradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwitchForUpgrade
// 本接口(SwitchForUpgrade)用于切换访问新实例，针对主升级中的实例处于待切换状态时，用户可主动发起该流程。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) SwitchForUpgrade(request *SwitchForUpgradeRequest) (response *SwitchForUpgradeResponse, err error) {
    return c.SwitchForUpgradeWithContext(context.Background(), request)
}

// SwitchForUpgrade
// 本接口(SwitchForUpgrade)用于切换访问新实例，针对主升级中的实例处于待切换状态时，用户可主动发起该流程。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) SwitchForUpgradeWithContext(ctx context.Context, request *SwitchForUpgradeRequest) (response *SwitchForUpgradeResponse, err error) {
    if request == nil {
        request = NewSwitchForUpgradeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchForUpgrade require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchForUpgradeResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeCDBProxyVersionRequest() (request *UpgradeCDBProxyVersionRequest) {
    request = &UpgradeCDBProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeCDBProxyVersion")
    
    
    return
}

func NewUpgradeCDBProxyVersionResponse() (response *UpgradeCDBProxyVersionResponse) {
    response = &UpgradeCDBProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeCDBProxyVersion
// 升级数据库代理版本
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) UpgradeCDBProxyVersion(request *UpgradeCDBProxyVersionRequest) (response *UpgradeCDBProxyVersionResponse, err error) {
    return c.UpgradeCDBProxyVersionWithContext(context.Background(), request)
}

// UpgradeCDBProxyVersion
// 升级数据库代理版本
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) UpgradeCDBProxyVersionWithContext(ctx context.Context, request *UpgradeCDBProxyVersionRequest) (response *UpgradeCDBProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeCDBProxyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeCDBProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeCDBProxyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceRequest() (request *UpgradeDBInstanceRequest) {
    request = &UpgradeDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeDBInstance")
    
    
    return
}

func NewUpgradeDBInstanceResponse() (response *UpgradeDBInstanceResponse) {
    response = &UpgradeDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeDBInstance
// 本接口(UpgradeDBInstance)用于升级或降级云数据库实例的配置，实例类型支持主实例、灾备实例和只读实例。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (response *UpgradeDBInstanceResponse, err error) {
    return c.UpgradeDBInstanceWithContext(context.Background(), request)
}

// UpgradeDBInstance
// 本接口(UpgradeDBInstance)用于升级或降级云数据库实例的配置，实例类型支持主实例、灾备实例和只读实例。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
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

func NewUpgradeDBInstanceEngineVersionRequest() (request *UpgradeDBInstanceEngineVersionRequest) {
    request = &UpgradeDBInstanceEngineVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "UpgradeDBInstanceEngineVersion")
    
    
    return
}

func NewUpgradeDBInstanceEngineVersionResponse() (response *UpgradeDBInstanceEngineVersionResponse) {
    response = &UpgradeDBInstanceEngineVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeDBInstanceEngineVersion
// 本接口(UpgradeDBInstanceEngineVersion)用于升级云数据库实例版本，实例类型支持主实例、灾备实例和只读实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    return c.UpgradeDBInstanceEngineVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceEngineVersion
// 本接口(UpgradeDBInstanceEngineVersion)用于升级云数据库实例版本，实例类型支持主实例、灾备实例和只读实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstanceEngineVersionWithContext(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceEngineVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstanceEngineVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceEngineVersionResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyRootAccountRequest() (request *VerifyRootAccountRequest) {
    request = &VerifyRootAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "VerifyRootAccount")
    
    
    return
}

func NewVerifyRootAccountResponse() (response *VerifyRootAccountResponse) {
    response = &VerifyRootAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VerifyRootAccount
// 本接口(VerifyRootAccount)用于校验云数据库实例的 ROOT 账号是否有足够的权限进行授权操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) VerifyRootAccount(request *VerifyRootAccountRequest) (response *VerifyRootAccountResponse, err error) {
    return c.VerifyRootAccountWithContext(context.Background(), request)
}

// VerifyRootAccount
// 本接口(VerifyRootAccount)用于校验云数据库实例的 ROOT 账号是否有足够的权限进行授权操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONCHARACTERERROR = "InvalidParameterValue.AccountDescriptionCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTDESCRIPTIONLENGTHERROR = "InvalidParameterValue.AccountDescriptionLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTHOSTRULEERROR = "InvalidParameterValue.AccountHostRuleError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDCHARACTERERROR = "InvalidParameterValue.AccountPasswordCharacterError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDLENGTHERROR = "InvalidParameterValue.AccountPasswordLengthError"
//  INVALIDPARAMETERVALUE_ACCOUNTPASSWORDRULEERROR = "InvalidParameterValue.AccountPasswordRuleError"
//  INVALIDPARAMETERVALUE_USERNAMERULEERROR = "InvalidParameterValue.UserNameRuleError"
//  INVALIDPARAMETERVALUE_USERNOTEXISTERROR = "InvalidParameterValue.UserNotExistError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTNOROOTERROR = "InvalidParameterValue.VerifyAccountNoRootError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPASSWORDERROR = "InvalidParameterValue.VerifyAccountPasswordError"
//  INVALIDPARAMETERVALUE_VERIFYACCOUNTPRIVERROR = "InvalidParameterValue.VerifyAccountPrivError"
//  MISSINGPARAMETER_ACCOUNTMISSINGPARAMETERERROR = "MissingParameter.AccountMissingParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETEROOTACCOUNTERROR = "OperationDenied.DeleteRootAccountError"
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) VerifyRootAccountWithContext(ctx context.Context, request *VerifyRootAccountRequest) (response *VerifyRootAccountResponse, err error) {
    if request == nil {
        request = NewVerifyRootAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyRootAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyRootAccountResponse()
    err = c.Send(request, response)
    return
}
