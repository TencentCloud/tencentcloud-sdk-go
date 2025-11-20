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

package v20190725

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-25"

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


func NewAssignProjectRequest() (request *AssignProjectRequest) {
    request = &AssignProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "AssignProject")
    
    
    return
}

func NewAssignProjectResponse() (response *AssignProjectResponse) {
    response = &AssignProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AssignProject
// 本接口(AssignProject)用于指定云数据库实例的所属项目。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
func (c *Client) AssignProject(request *AssignProjectRequest) (response *AssignProjectResponse, err error) {
    return c.AssignProjectWithContext(context.Background(), request)
}

// AssignProject
// 本接口(AssignProject)用于指定云数据库实例的所属项目。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
func (c *Client) AssignProjectWithContext(ctx context.Context, request *AssignProjectRequest) (response *AssignProjectResponse, err error) {
    if request == nil {
        request = NewAssignProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "AssignProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssignProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewAssignProjectResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccountUserRequest() (request *CreateAccountUserRequest) {
    request = &CreateAccountUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateAccountUser")
    
    
    return
}

func NewCreateAccountUserResponse() (response *CreateAccountUserResponse) {
    response = &CreateAccountUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAccountUser
// 本接口（CreateAccountUser）用于自定义实例访问账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_PASSWORDERROR = "InternalError.PasswordError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
func (c *Client) CreateAccountUser(request *CreateAccountUserRequest) (response *CreateAccountUserResponse, err error) {
    return c.CreateAccountUserWithContext(context.Background(), request)
}

// CreateAccountUser
// 本接口（CreateAccountUser）用于自定义实例访问账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_PASSWORDERROR = "InternalError.PasswordError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
func (c *Client) CreateAccountUserWithContext(ctx context.Context, request *CreateAccountUserRequest) (response *CreateAccountUserResponse, err error) {
    if request == nil {
        request = NewCreateAccountUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateAccountUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccountUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountUserResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditLogFileRequest() (request *CreateAuditLogFileRequest) {
    request = &CreateAuditLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateAuditLogFile")
    
    
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
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) CreateAuditLogFile(request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    return c.CreateAuditLogFileWithContext(context.Background(), request)
}

// CreateAuditLogFile
// 本接口(CreateAuditLogFile)用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) CreateAuditLogFileWithContext(ctx context.Context, request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    if request == nil {
        request = NewCreateAuditLogFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateAuditLogFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupDBInstanceRequest() (request *CreateBackupDBInstanceRequest) {
    request = &CreateBackupDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateBackupDBInstance")
    
    
    return
}

func NewCreateBackupDBInstanceResponse() (response *CreateBackupDBInstanceResponse) {
    response = &CreateBackupDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupDBInstance
// 本接口（CreateBackupDBInstance）用于备份实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRANSPARENTDATAENCRYPTIONALREADYOPEN = "FailedOperation.TransparentDataEncryptionAlreadyOpen"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) CreateBackupDBInstance(request *CreateBackupDBInstanceRequest) (response *CreateBackupDBInstanceResponse, err error) {
    return c.CreateBackupDBInstanceWithContext(context.Background(), request)
}

// CreateBackupDBInstance
// 本接口（CreateBackupDBInstance）用于备份实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TRANSPARENTDATAENCRYPTIONALREADYOPEN = "FailedOperation.TransparentDataEncryptionAlreadyOpen"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) CreateBackupDBInstanceWithContext(ctx context.Context, request *CreateBackupDBInstanceRequest) (response *CreateBackupDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateBackupDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateBackupDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupDownloadTaskRequest() (request *CreateBackupDownloadTaskRequest) {
    request = &CreateBackupDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateBackupDownloadTask")
    
    
    return
}

func NewCreateBackupDownloadTaskResponse() (response *CreateBackupDownloadTaskResponse) {
    response = &CreateBackupDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBackupDownloadTask
// 本接口用来创建某个备份文件的下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateBackupDownloadTask(request *CreateBackupDownloadTaskRequest) (response *CreateBackupDownloadTaskResponse, err error) {
    return c.CreateBackupDownloadTaskWithContext(context.Background(), request)
}

// CreateBackupDownloadTask
// 本接口用来创建某个备份文件的下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateBackupDownloadTaskWithContext(ctx context.Context, request *CreateBackupDownloadTaskRequest) (response *CreateBackupDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateBackupDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateBackupDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
    request = &CreateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstance")
    
    
    return
}

func NewCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
    response = &CreateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstance
// 本接口(CreateDBInstance)用于创建包年包月的MongoDB云数据库实例。接口支持的售卖规格，可通过接口查询 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/35767) 获取。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SECONDARYNUMERROR = "InvalidParameterValue.SecondaryNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    return c.CreateDBInstanceWithContext(context.Background(), request)
}

// CreateDBInstance
// 本接口(CreateDBInstance)用于创建包年包月的MongoDB云数据库实例。接口支持的售卖规格，可通过接口查询 [DescribeSpecInfo](https://cloud.tencent.com/document/product/240/35767) 获取。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SECONDARYNUMERROR = "InvalidParameterValue.SecondaryNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateDBInstance")
    
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
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstanceHour")
    
    
    return
}

func NewCreateDBInstanceHourResponse() (response *CreateDBInstanceHourResponse) {
    response = &CreateDBInstanceHourResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstanceHour
// 本接口（CreateDBInstanceHour）用于创建按量计费的MongoDB云数据库实例。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEERROR = "InvalidParameterValue.ClusterTypeError"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_POSTPAIDINSTANCEBEYONDLIMIT = "InvalidParameterValue.PostPaidInstanceBeyondLimit"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    return c.CreateDBInstanceHourWithContext(context.Background(), request)
}

// CreateDBInstanceHour
// 本接口（CreateDBInstanceHour）用于创建按量计费的MongoDB云数据库实例。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEERROR = "InvalidParameterValue.ClusterTypeError"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_POSTPAIDINSTANCEBEYONDLIMIT = "InvalidParameterValue.PostPaidInstanceBeyondLimit"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceHourWithContext(ctx context.Context, request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateDBInstanceHour")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceHour require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstanceParamTplRequest() (request *CreateDBInstanceParamTplRequest) {
    request = &CreateDBInstanceParamTplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateDBInstanceParamTpl")
    
    
    return
}

func NewCreateDBInstanceParamTplResponse() (response *CreateDBInstanceParamTplResponse) {
    response = &CreateDBInstanceParamTplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstanceParamTpl
// 本接口(CreateDBInstanceParamTpl)用于创建云数据库MongoDB实例的参数模板
//
// **说明：CreateDBInstanceParamTpl API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEERROR = "InvalidParameterValue.ClusterTypeError"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_POSTPAIDINSTANCEBEYONDLIMIT = "InvalidParameterValue.PostPaidInstanceBeyondLimit"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceParamTpl(request *CreateDBInstanceParamTplRequest) (response *CreateDBInstanceParamTplResponse, err error) {
    return c.CreateDBInstanceParamTplWithContext(context.Background(), request)
}

// CreateDBInstanceParamTpl
// 本接口(CreateDBInstanceParamTpl)用于创建云数据库MongoDB实例的参数模板
//
// **说明：CreateDBInstanceParamTpl API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CLUSTERTYPEERROR = "InvalidParameterValue.ClusterTypeError"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_POSTPAIDINSTANCEBEYONDLIMIT = "InvalidParameterValue.PostPaidInstanceBeyondLimit"
//  INVALIDPARAMETERVALUE_PROJECTNOTFOUND = "InvalidParameterValue.ProjectNotFound"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_TAGNOTFOUND = "InvalidParameterValue.TagNotFound"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) CreateDBInstanceParamTplWithContext(ctx context.Context, request *CreateDBInstanceParamTplRequest) (response *CreateDBInstanceParamTplResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceParamTplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateDBInstanceParamTpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceParamTpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceParamTplResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogDownloadTaskRequest() (request *CreateLogDownloadTaskRequest) {
    request = &CreateLogDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "CreateLogDownloadTask")
    
    
    return
}

func NewCreateLogDownloadTaskResponse() (response *CreateLogDownloadTaskResponse) {
    response = &CreateLogDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateLogDownloadTask
// 创建日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateLogDownloadTask(request *CreateLogDownloadTaskRequest) (response *CreateLogDownloadTaskResponse, err error) {
    return c.CreateLogDownloadTaskWithContext(context.Background(), request)
}

// CreateLogDownloadTask
// 创建日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) CreateLogDownloadTaskWithContext(ctx context.Context, request *CreateLogDownloadTaskRequest) (response *CreateLogDownloadTaskResponse, err error) {
    if request == nil {
        request = NewCreateLogDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "CreateLogDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccountUserRequest() (request *DeleteAccountUserRequest) {
    request = &DeleteAccountUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteAccountUser")
    
    
    return
}

func NewDeleteAccountUserResponse() (response *DeleteAccountUserResponse) {
    response = &DeleteAccountUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAccountUser
// 本接口（DeleteAccountUser）用于删除实例的自定义账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteAccountUser(request *DeleteAccountUserRequest) (response *DeleteAccountUserResponse, err error) {
    return c.DeleteAccountUserWithContext(context.Background(), request)
}

// DeleteAccountUser
// 本接口（DeleteAccountUser）用于删除实例的自定义账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteAccountUserWithContext(ctx context.Context, request *DeleteAccountUserRequest) (response *DeleteAccountUserResponse, err error) {
    if request == nil {
        request = NewDeleteAccountUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DeleteAccountUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAccountUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAccountUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditLogFileRequest() (request *DeleteAuditLogFileRequest) {
    request = &DeleteAuditLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteAuditLogFile")
    
    
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
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DeleteAuditLogFile(request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    return c.DeleteAuditLogFileWithContext(context.Background(), request)
}

// DeleteAuditLogFile
// 本接口(DeleteAuditLogFile)用于删除云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITDELETELOGFILEERROR = "InternalError.AuditDeleteLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DeleteAuditLogFileWithContext(ctx context.Context, request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    if request == nil {
        request = NewDeleteAuditLogFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DeleteAuditLogFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogDownloadTaskRequest() (request *DeleteLogDownloadTaskRequest) {
    request = &DeleteLogDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DeleteLogDownloadTask")
    
    
    return
}

func NewDeleteLogDownloadTaskResponse() (response *DeleteLogDownloadTaskResponse) {
    response = &DeleteLogDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteLogDownloadTask
// 删除日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteLogDownloadTask(request *DeleteLogDownloadTaskRequest) (response *DeleteLogDownloadTaskResponse, err error) {
    return c.DeleteLogDownloadTaskWithContext(context.Background(), request)
}

// DeleteLogDownloadTask
// 删除日志下载任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DeleteLogDownloadTaskWithContext(ctx context.Context, request *DeleteLogDownloadTaskRequest) (response *DeleteLogDownloadTaskResponse, err error) {
    if request == nil {
        request = NewDeleteLogDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DeleteLogDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountUsersRequest() (request *DescribeAccountUsersRequest) {
    request = &DescribeAccountUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAccountUsers")
    
    
    return
}

func NewDescribeAccountUsersResponse() (response *DescribeAccountUsersResponse) {
    response = &DescribeAccountUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAccountUsers
// 本接口（DescribeAccountUsers）用于获取当前实例的全部账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeAccountUsers(request *DescribeAccountUsersRequest) (response *DescribeAccountUsersResponse, err error) {
    return c.DescribeAccountUsersWithContext(context.Background(), request)
}

// DescribeAccountUsers
// 本接口（DescribeAccountUsers）用于获取当前实例的全部账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeAccountUsersWithContext(ctx context.Context, request *DescribeAccountUsersRequest) (response *DescribeAccountUsersResponse, err error) {
    if request == nil {
        request = NewDescribeAccountUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAccountUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAsyncRequestInfo")
    
    
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAsyncRequestInfo
// 本接口（DescribeAsyncRequestInfo）用于查询异步任务状态接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// 本接口（DescribeAsyncRequestInfo）用于查询异步任务状态接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAsyncRequestInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditInstanceListRequest() (request *DescribeAuditInstanceListRequest) {
    request = &DescribeAuditInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeAuditInstanceList")
    
    
    return
}

func NewDescribeAuditInstanceListResponse() (response *DescribeAuditInstanceListResponse) {
    response = &DescribeAuditInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditInstanceList
// 本接口（DescribeAuditInstanceList）用于查询开通或未开通数据库审计的实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYAUDITTASKFAILERROR = "FailedOperation.QueryAuditTaskFailError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTOPERATIONDENIED = "OperationDenied.AccountOperationDenied"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DescribeAuditInstanceList(request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    return c.DescribeAuditInstanceListWithContext(context.Background(), request)
}

// DescribeAuditInstanceList
// 本接口（DescribeAuditInstanceList）用于查询开通或未开通数据库审计的实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYAUDITTASKFAILERROR = "FailedOperation.QueryAuditTaskFailError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACCOUNTOPERATIONDENIED = "OperationDenied.AccountOperationDenied"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) DescribeAuditInstanceListWithContext(ctx context.Context, request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAuditInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeAuditInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadTaskRequest() (request *DescribeBackupDownloadTaskRequest) {
    request = &DescribeBackupDownloadTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeBackupDownloadTask")
    
    
    return
}

func NewDescribeBackupDownloadTaskResponse() (response *DescribeBackupDownloadTaskResponse) {
    response = &DescribeBackupDownloadTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadTask
// 本接口（DescribeBackupDownloadTask）用于查询备份下载任务信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeBackupDownloadTask(request *DescribeBackupDownloadTaskRequest) (response *DescribeBackupDownloadTaskResponse, err error) {
    return c.DescribeBackupDownloadTaskWithContext(context.Background(), request)
}

// DescribeBackupDownloadTask
// 本接口（DescribeBackupDownloadTask）用于查询备份下载任务信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPFILENOTFOUND = "InvalidParameterValue.BackupFileNotFound"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeBackupDownloadTaskWithContext(ctx context.Context, request *DescribeBackupDownloadTaskRequest) (response *DescribeBackupDownloadTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDownloadTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeBackupDownloadTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupRulesRequest() (request *DescribeBackupRulesRequest) {
    request = &DescribeBackupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeBackupRules")
    
    
    return
}

func NewDescribeBackupRulesResponse() (response *DescribeBackupRulesResponse) {
    response = &DescribeBackupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupRules
// 本接口（DescribeBackupRules）用于获取实例自动备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
func (c *Client) DescribeBackupRules(request *DescribeBackupRulesRequest) (response *DescribeBackupRulesResponse, err error) {
    return c.DescribeBackupRulesWithContext(context.Background(), request)
}

// DescribeBackupRules
// 本接口（DescribeBackupRules）用于获取实例自动备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
func (c *Client) DescribeBackupRulesWithContext(ctx context.Context, request *DescribeBackupRulesRequest) (response *DescribeBackupRulesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeBackupRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClientConnectionsRequest() (request *DescribeClientConnectionsRequest) {
    request = &DescribeClientConnectionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeClientConnections")
    
    
    return
}

func NewDescribeClientConnectionsResponse() (response *DescribeClientConnectionsResponse) {
    response = &DescribeClientConnectionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClientConnections
// 本接口（DescribeClientConnections）用于查询实例客户端连接信息，包括连接 IP 和连接数量。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MONGOVERSIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.MongoVersionNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROXYNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.ProxyNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.RegionNotSupportQueryClient"
func (c *Client) DescribeClientConnections(request *DescribeClientConnectionsRequest) (response *DescribeClientConnectionsResponse, err error) {
    return c.DescribeClientConnectionsWithContext(context.Background(), request)
}

// DescribeClientConnections
// 本接口（DescribeClientConnections）用于查询实例客户端连接信息，包括连接 IP 和连接数量。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MONGOVERSIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.MongoVersionNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PROXYNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.ProxyNotSupportQueryClient"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTQUERYCLIENT = "InvalidParameterValue.RegionNotSupportQueryClient"
func (c *Client) DescribeClientConnectionsWithContext(ctx context.Context, request *DescribeClientConnectionsRequest) (response *DescribeClientConnectionsResponse, err error) {
    if request == nil {
        request = NewDescribeClientConnectionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeClientConnections")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClientConnections require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClientConnectionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCurrentOpRequest() (request *DescribeCurrentOpRequest) {
    request = &DescribeCurrentOpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeCurrentOp")
    
    
    return
}

func NewDescribeCurrentOpResponse() (response *DescribeCurrentOpResponse) {
    response = &DescribeCurrentOpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCurrentOp
// 本接口（DescribeCurrentOp）用于查询云数据库实例的当前正在执行的操作。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYOUTOFRANGE = "InvalidParameterValue.QueryOutOfRange"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTOPERATION = "InvalidParameterValue.RegionNotSupportOperation"
//  INVALIDPARAMETERVALUE_REPLICANOTFOUND = "InvalidParameterValue.ReplicaNotFound"
func (c *Client) DescribeCurrentOp(request *DescribeCurrentOpRequest) (response *DescribeCurrentOpResponse, err error) {
    return c.DescribeCurrentOpWithContext(context.Background(), request)
}

// DescribeCurrentOp
// 本接口（DescribeCurrentOp）用于查询云数据库实例的当前正在执行的操作。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYOUTOFRANGE = "InvalidParameterValue.QueryOutOfRange"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTOPERATION = "InvalidParameterValue.RegionNotSupportOperation"
//  INVALIDPARAMETERVALUE_REPLICANOTFOUND = "InvalidParameterValue.ReplicaNotFound"
func (c *Client) DescribeCurrentOpWithContext(ctx context.Context, request *DescribeCurrentOpRequest) (response *DescribeCurrentOpResponse, err error) {
    if request == nil {
        request = NewDescribeCurrentOpRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeCurrentOp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCurrentOp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCurrentOpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBBackupsRequest() (request *DescribeDBBackupsRequest) {
    request = &DescribeDBBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBBackups")
    
    
    return
}

func NewDescribeDBBackupsResponse() (response *DescribeDBBackupsResponse) {
    response = &DescribeDBBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBBackups
// 本接口（DescribeDBBackups）用于查询实例备份列表，目前只支持查询7天内的备份记录。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeDBBackups(request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    return c.DescribeDBBackupsWithContext(context.Background(), request)
}

// DescribeDBBackups
// 本接口（DescribeDBBackups）用于查询实例备份列表，目前只支持查询7天内的备份记录。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeDBBackupsWithContext(ctx context.Context, request *DescribeDBBackupsRequest) (response *DescribeDBBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBBackups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceDealRequest() (request *DescribeDBInstanceDealRequest) {
    request = &DescribeDBInstanceDealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceDeal")
    
    
    return
}

func NewDescribeDBInstanceDealResponse() (response *DescribeDBInstanceDealResponse) {
    response = &DescribeDBInstanceDealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceDeal
// 本接口（DescribeDBInstanceDeal）用于获取MongoDB购买、续费及变配订单详细。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceDeal(request *DescribeDBInstanceDealRequest) (response *DescribeDBInstanceDealResponse, err error) {
    return c.DescribeDBInstanceDealWithContext(context.Background(), request)
}

// DescribeDBInstanceDeal
// 本接口（DescribeDBInstanceDeal）用于获取MongoDB购买、续费及变配订单详细。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceDealWithContext(ctx context.Context, request *DescribeDBInstanceDealRequest) (response *DescribeDBInstanceDealResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceDealRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceDeal")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceDeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceDealResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceNamespaceRequest() (request *DescribeDBInstanceNamespaceRequest) {
    request = &DescribeDBInstanceNamespaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceNamespace")
    
    
    return
}

func NewDescribeDBInstanceNamespaceResponse() (response *DescribeDBInstanceNamespaceResponse) {
    response = &DescribeDBInstanceNamespaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceNamespace
// 本接口（DescribeDBInstanceNamespace）用于查询数据库的表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceNamespace(request *DescribeDBInstanceNamespaceRequest) (response *DescribeDBInstanceNamespaceResponse, err error) {
    return c.DescribeDBInstanceNamespaceWithContext(context.Background(), request)
}

// DescribeDBInstanceNamespace
// 本接口（DescribeDBInstanceNamespace）用于查询数据库的表信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceNamespaceWithContext(ctx context.Context, request *DescribeDBInstanceNamespaceRequest) (response *DescribeDBInstanceNamespaceResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceNamespaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceNamespace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceNamespace require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceNamespaceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceNodePropertyRequest() (request *DescribeDBInstanceNodePropertyRequest) {
    request = &DescribeDBInstanceNodePropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceNodeProperty")
    
    
    return
}

func NewDescribeDBInstanceNodePropertyResponse() (response *DescribeDBInstanceNodePropertyResponse) {
    response = &DescribeDBInstanceNodePropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceNodeProperty
// 本接口用于查询节点的属性，包括节点所在可用区、节点名称、地址、角色、状态、主从延迟、优先级、投票权、标签等属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_KERNELRESPONSETIMEOUT = "FailedOperation.KernelResponseTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceNodeProperty(request *DescribeDBInstanceNodePropertyRequest) (response *DescribeDBInstanceNodePropertyResponse, err error) {
    return c.DescribeDBInstanceNodePropertyWithContext(context.Background(), request)
}

// DescribeDBInstanceNodeProperty
// 本接口用于查询节点的属性，包括节点所在可用区、节点名称、地址、角色、状态、主从延迟、优先级、投票权、标签等属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION_KERNELRESPONSETIMEOUT = "FailedOperation.KernelResponseTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceNodePropertyWithContext(ctx context.Context, request *DescribeDBInstanceNodePropertyRequest) (response *DescribeDBInstanceNodePropertyResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceNodePropertyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceNodeProperty")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceNodeProperty require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceNodePropertyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceParamTplRequest() (request *DescribeDBInstanceParamTplRequest) {
    request = &DescribeDBInstanceParamTplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceParamTpl")
    
    
    return
}

func NewDescribeDBInstanceParamTplResponse() (response *DescribeDBInstanceParamTplResponse) {
    response = &DescribeDBInstanceParamTplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceParamTpl
// 本接口(DescribeDBInstanceParamTpl )用于查询当前账号下所有MongoDB数据库参数模板
//
// **说明：DescribeDBInstanceParamTpl  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTpl(request *DescribeDBInstanceParamTplRequest) (response *DescribeDBInstanceParamTplResponse, err error) {
    return c.DescribeDBInstanceParamTplWithContext(context.Background(), request)
}

// DescribeDBInstanceParamTpl
// 本接口(DescribeDBInstanceParamTpl )用于查询当前账号下所有MongoDB数据库参数模板
//
// **说明：DescribeDBInstanceParamTpl  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTplWithContext(ctx context.Context, request *DescribeDBInstanceParamTplRequest) (response *DescribeDBInstanceParamTplResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceParamTplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceParamTpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceParamTpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceParamTplResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceParamTplDetailRequest() (request *DescribeDBInstanceParamTplDetailRequest) {
    request = &DescribeDBInstanceParamTplDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceParamTplDetail")
    
    
    return
}

func NewDescribeDBInstanceParamTplDetailResponse() (response *DescribeDBInstanceParamTplDetailResponse) {
    response = &DescribeDBInstanceParamTplDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceParamTplDetail
// 本接口(DescribeDBInstanceParamTplDetail )用于查询MongoDB云数据库实例的参数模板详情。
//
// **说明：DescribeDBInstanceParamTplDetail  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTplDetail(request *DescribeDBInstanceParamTplDetailRequest) (response *DescribeDBInstanceParamTplDetailResponse, err error) {
    return c.DescribeDBInstanceParamTplDetailWithContext(context.Background(), request)
}

// DescribeDBInstanceParamTplDetail
// 本接口(DescribeDBInstanceParamTplDetail )用于查询MongoDB云数据库实例的参数模板详情。
//
// **说明：DescribeDBInstanceParamTplDetail  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceParamTplDetailWithContext(ctx context.Context, request *DescribeDBInstanceParamTplDetailRequest) (response *DescribeDBInstanceParamTplDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceParamTplDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceParamTplDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceParamTplDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceParamTplDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceURLRequest() (request *DescribeDBInstanceURLRequest) {
    request = &DescribeDBInstanceURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstanceURL")
    
    
    return
}

func NewDescribeDBInstanceURLResponse() (response *DescribeDBInstanceURLResponse) {
    response = &DescribeDBInstanceURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceURL
// 本接口（DescribeDBInstanceURL）用于获取指定实例的 URI 形式的连接串访问地址示例。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceURL(request *DescribeDBInstanceURLRequest) (response *DescribeDBInstanceURLResponse, err error) {
    return c.DescribeDBInstanceURLWithContext(context.Background(), request)
}

// DescribeDBInstanceURL
// 本接口（DescribeDBInstanceURL）用于获取指定实例的 URI 形式的连接串访问地址示例。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeDBInstanceURLWithContext(ctx context.Context, request *DescribeDBInstanceURLRequest) (response *DescribeDBInstanceURLResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceURLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstanceURL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceURL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceURLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstances
// 本接口（DescribeDBInstances）用于查询云数据库实例列表，支持通过项目ID、实例ID、实例状态等过滤条件来筛选主实例、灾备实例和只读实例信息列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  LIMITEXCEEDED_TOOMANYREQUESTS = "LimitExceeded.TooManyRequests"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口（DescribeDBInstances）用于查询云数据库实例列表，支持通过项目ID、实例ID、实例状态等过滤条件来筛选主实例、灾备实例和只读实例信息列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  LIMITEXCEEDED_TOOMANYREQUESTS = "LimitExceeded.TooManyRequests"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDetailedSlowLogsRequest() (request *DescribeDetailedSlowLogsRequest) {
    request = &DescribeDetailedSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeDetailedSlowLogs")
    
    
    return
}

func NewDescribeDetailedSlowLogsResponse() (response *DescribeDetailedSlowLogsResponse) {
    response = &DescribeDetailedSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDetailedSlowLogs
// 本接口（DescribeDetailedSlowLogs）用于查询实例慢日志详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITPARAOUTOFRANGE = "InvalidParameterValue.LimitParaOutOfRange"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OFFSETPARAOUTOFRANGE = "InvalidParameterValue.OffsetParaOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeDetailedSlowLogs(request *DescribeDetailedSlowLogsRequest) (response *DescribeDetailedSlowLogsResponse, err error) {
    return c.DescribeDetailedSlowLogsWithContext(context.Background(), request)
}

// DescribeDetailedSlowLogs
// 本接口（DescribeDetailedSlowLogs）用于查询实例慢日志详情。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LIMITPARAOUTOFRANGE = "InvalidParameterValue.LimitParaOutOfRange"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OFFSETPARAOUTOFRANGE = "InvalidParameterValue.OffsetParaOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeDetailedSlowLogsWithContext(ctx context.Context, request *DescribeDetailedSlowLogsRequest) (response *DescribeDetailedSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDetailedSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeDetailedSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDetailedSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDetailedSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamsRequest() (request *DescribeInstanceParamsRequest) {
    request = &DescribeInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询当前实例可修改的参数列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CURRENTINSTANCENOTSUPPORTMODIFYPARAMS = "InvalidParameter.CurrentInstanceNotSupportModifyParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询当前实例可修改的参数列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CURRENTINSTANCENOTSUPPORTMODIFYPARAMS = "InvalidParameter.CurrentInstanceNotSupportModifyParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSSLRequest() (request *DescribeInstanceSSLRequest) {
    request = &DescribeInstanceSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeInstanceSSL")
    
    
    return
}

func NewDescribeInstanceSSLResponse() (response *DescribeInstanceSSLResponse) {
    response = &DescribeInstanceSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSSL
// 查看实例SSL开启状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) DescribeInstanceSSL(request *DescribeInstanceSSLRequest) (response *DescribeInstanceSSLResponse, err error) {
    return c.DescribeInstanceSSLWithContext(context.Background(), request)
}

// DescribeInstanceSSL
// 查看实例SSL开启状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) DescribeInstanceSSLWithContext(ctx context.Context, request *DescribeInstanceSSLRequest) (response *DescribeInstanceSSLResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeInstanceSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSSLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogDownloadTasksRequest() (request *DescribeLogDownloadTasksRequest) {
    request = &DescribeLogDownloadTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeLogDownloadTasks")
    
    
    return
}

func NewDescribeLogDownloadTasksResponse() (response *DescribeLogDownloadTasksResponse) {
    response = &DescribeLogDownloadTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeLogDownloadTasks
// 日志下载任务查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeLogDownloadTasks(request *DescribeLogDownloadTasksRequest) (response *DescribeLogDownloadTasksResponse, err error) {
    return c.DescribeLogDownloadTasksWithContext(context.Background(), request)
}

// DescribeLogDownloadTasks
// 日志下载任务查询
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeLogDownloadTasksWithContext(ctx context.Context, request *DescribeLogDownloadTasksRequest) (response *DescribeLogDownloadTasksResponse, err error) {
    if request == nil {
        request = NewDescribeLogDownloadTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeLogDownloadTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogDownloadTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogDownloadTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMongodbLogsRequest() (request *DescribeMongodbLogsRequest) {
    request = &DescribeMongodbLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeMongodbLogs")
    
    
    return
}

func NewDescribeMongodbLogsResponse() (response *DescribeMongodbLogsResponse) {
    response = &DescribeMongodbLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMongodbLogs
// 该接口（DescribeMongodbLogs）用于查询运行日志。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeMongodbLogs(request *DescribeMongodbLogsRequest) (response *DescribeMongodbLogsResponse, err error) {
    return c.DescribeMongodbLogsWithContext(context.Background(), request)
}

// DescribeMongodbLogs
// 该接口（DescribeMongodbLogs）用于查询运行日志。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_TIMEFORMATERR = "InvalidParameterValue.TimeFormatErr"
func (c *Client) DescribeMongodbLogsWithContext(ctx context.Context, request *DescribeMongodbLogsRequest) (response *DescribeMongodbLogsResponse, err error) {
    if request == nil {
        request = NewDescribeMongodbLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeMongodbLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMongodbLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMongodbLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
    request = &DescribeSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSecurityGroup")
    
    
    return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
    response = &DescribeSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityGroup
// 本接口（DescribeSecurityGroup）用于查询实例绑定的安全组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
    return c.DescribeSecurityGroupWithContext(context.Background(), request)
}

// DescribeSecurityGroup
// 本接口（DescribeSecurityGroup）用于查询实例绑定的安全组。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeSecurityGroupWithContext(ctx context.Context, request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogPatternsRequest() (request *DescribeSlowLogPatternsRequest) {
    request = &DescribeSlowLogPatternsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogPatterns")
    
    
    return
}

func NewDescribeSlowLogPatternsResponse() (response *DescribeSlowLogPatternsResponse) {
    response = &DescribeSlowLogPatternsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogPatterns
// 本接口（DescribeSlowLogPatterns）用于获取数据库实例慢日志的统计信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogPatterns(request *DescribeSlowLogPatternsRequest) (response *DescribeSlowLogPatternsResponse, err error) {
    return c.DescribeSlowLogPatternsWithContext(context.Background(), request)
}

// DescribeSlowLogPatterns
// 本接口（DescribeSlowLogPatterns）用于获取数据库实例慢日志的统计信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogPatternsWithContext(ctx context.Context, request *DescribeSlowLogPatternsRequest) (response *DescribeSlowLogPatternsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogPatternsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSlowLogPatterns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogPatterns require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogPatternsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSlowLogs")
    
    
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogs
// 本接口（DescribeSlowLogs）用于获取云数据库慢日志信息。接口只支持查询最近7天内慢日志。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    return c.DescribeSlowLogsWithContext(context.Background(), request)
}

// DescribeSlowLogs
// 本接口（DescribeSlowLogs）用于获取云数据库慢日志信息。接口只支持查询最近7天内慢日志。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_QUERYTIMEOUTOFRANGE = "InvalidParameterValue.QueryTimeOutOfRange"
//  INVALIDPARAMETERVALUE_QUERYTIMERANGEBEYONDLIMIT = "InvalidParameterValue.QueryTimeRangeBeyondLimit"
//  INVALIDPARAMETERVALUE_SLOWMSBELOWLIMIT = "InvalidParameterValue.SlowMSBelowLimit"
//  INVALIDPARAMETERVALUE_STARTTIMENOTBEFORETHANENDTIME = "InvalidParameterValue.StartTimeNotBeforeThanEndTime"
func (c *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecInfoRequest() (request *DescribeSpecInfoRequest) {
    request = &DescribeSpecInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeSpecInfo")
    
    
    return
}

func NewDescribeSpecInfoResponse() (response *DescribeSpecInfoResponse) {
    response = &DescribeSpecInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpecInfo
// 本接口（DescribeSpecInfo）用于查询实例的售卖规格。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) DescribeSpecInfo(request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    return c.DescribeSpecInfoWithContext(context.Background(), request)
}

// DescribeSpecInfo
// 本接口（DescribeSpecInfo）用于查询实例的售卖规格。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) DescribeSpecInfoWithContext(ctx context.Context, request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSpecInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeSpecInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpecInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpecInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTransparentDataEncryptionStatusRequest() (request *DescribeTransparentDataEncryptionStatusRequest) {
    request = &DescribeTransparentDataEncryptionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DescribeTransparentDataEncryptionStatus")
    
    
    return
}

func NewDescribeTransparentDataEncryptionStatusResponse() (response *DescribeTransparentDataEncryptionStatusResponse) {
    response = &DescribeTransparentDataEncryptionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTransparentDataEncryptionStatus
// 获取实例透明加密的开启状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeTransparentDataEncryptionStatus(request *DescribeTransparentDataEncryptionStatusRequest) (response *DescribeTransparentDataEncryptionStatusResponse, err error) {
    return c.DescribeTransparentDataEncryptionStatusWithContext(context.Background(), request)
}

// DescribeTransparentDataEncryptionStatus
// 获取实例透明加密的开启状态
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) DescribeTransparentDataEncryptionStatusWithContext(ctx context.Context, request *DescribeTransparentDataEncryptionStatusRequest) (response *DescribeTransparentDataEncryptionStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTransparentDataEncryptionStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DescribeTransparentDataEncryptionStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTransparentDataEncryptionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTransparentDataEncryptionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDropDBInstanceParamTplRequest() (request *DropDBInstanceParamTplRequest) {
    request = &DropDBInstanceParamTplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "DropDBInstanceParamTpl")
    
    
    return
}

func NewDropDBInstanceParamTplResponse() (response *DropDBInstanceParamTplResponse) {
    response = &DropDBInstanceParamTplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DropDBInstanceParamTpl
// 本接口(DropDBInstanceParamTpl )用于删除云数据库MongoDB实例的参数模板
//
// **说明：DropDBInstanceParamTpl  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DropDBInstanceParamTpl(request *DropDBInstanceParamTplRequest) (response *DropDBInstanceParamTplResponse, err error) {
    return c.DropDBInstanceParamTplWithContext(context.Background(), request)
}

// DropDBInstanceParamTpl
// 本接口(DropDBInstanceParamTpl )用于删除云数据库MongoDB实例的参数模板
//
// **说明：DropDBInstanceParamTpl  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DropDBInstanceParamTplWithContext(ctx context.Context, request *DropDBInstanceParamTplRequest) (response *DropDBInstanceParamTplResponse, err error) {
    if request == nil {
        request = NewDropDBInstanceParamTplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "DropDBInstanceParamTpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DropDBInstanceParamTpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDropDBInstanceParamTplResponse()
    err = c.Send(request, response)
    return
}

func NewEnableTransparentDataEncryptionRequest() (request *EnableTransparentDataEncryptionRequest) {
    request = &EnableTransparentDataEncryptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "EnableTransparentDataEncryption")
    
    
    return
}

func NewEnableTransparentDataEncryptionResponse() (response *EnableTransparentDataEncryptionResponse) {
    response = &EnableTransparentDataEncryptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableTransparentDataEncryption
// 本接口（EnableTransparentDataEncryption）用于开启云数据库 MongoDB 的透明加密能力。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) EnableTransparentDataEncryption(request *EnableTransparentDataEncryptionRequest) (response *EnableTransparentDataEncryptionResponse, err error) {
    return c.EnableTransparentDataEncryptionWithContext(context.Background(), request)
}

// EnableTransparentDataEncryption
// 本接口（EnableTransparentDataEncryption）用于开启云数据库 MongoDB 的透明加密能力。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) EnableTransparentDataEncryptionWithContext(ctx context.Context, request *EnableTransparentDataEncryptionRequest) (response *EnableTransparentDataEncryptionResponse, err error) {
    if request == nil {
        request = NewEnableTransparentDataEncryptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "EnableTransparentDataEncryption")
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableTransparentDataEncryption require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableTransparentDataEncryptionResponse()
    err = c.Send(request, response)
    return
}

func NewFlashBackDBInstanceRequest() (request *FlashBackDBInstanceRequest) {
    request = &FlashBackDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "FlashBackDBInstance")
    
    
    return
}

func NewFlashBackDBInstanceResponse() (response *FlashBackDBInstanceResponse) {
    response = &FlashBackDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FlashBackDBInstance
// 该接口用于发起按 Key 闪回任务，依据数据的闪回 Key（默认为 id）对数据进行极速回档，快速恢复业务。
//
// **说明：按 Key 闪回于2023年09月11日正式进行公测，在此期间，该接口仅对公测用户开放。**
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLASHBACKBYKEYNOTOPEN = "FailedOperation.FlashbackByKeyNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERSNIL = "InvalidParameter.ParametersNil"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNAUTHORIZEDOPERATION_NOACCESS = "UnauthorizedOperation.NoAccess"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) FlashBackDBInstance(request *FlashBackDBInstanceRequest) (response *FlashBackDBInstanceResponse, err error) {
    return c.FlashBackDBInstanceWithContext(context.Background(), request)
}

// FlashBackDBInstance
// 该接口用于发起按 Key 闪回任务，依据数据的闪回 Key（默认为 id）对数据进行极速回档，快速恢复业务。
//
// **说明：按 Key 闪回于2023年09月11日正式进行公测，在此期间，该接口仅对公测用户开放。**
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLASHBACKBYKEYNOTOPEN = "FailedOperation.FlashbackByKeyNotOpen"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMETERSNIL = "InvalidParameter.ParametersNil"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  UNAUTHORIZEDOPERATION_NOACCESS = "UnauthorizedOperation.NoAccess"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) FlashBackDBInstanceWithContext(ctx context.Context, request *FlashBackDBInstanceRequest) (response *FlashBackDBInstanceResponse, err error) {
    if request == nil {
        request = NewFlashBackDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "FlashBackDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FlashBackDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewFlashBackDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewFlushInstanceRouterConfigRequest() (request *FlushInstanceRouterConfigRequest) {
    request = &FlushInstanceRouterConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "FlushInstanceRouterConfig")
    
    
    return
}

func NewFlushInstanceRouterConfigResponse() (response *FlushInstanceRouterConfigResponse) {
    response = &FlushInstanceRouterConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FlushInstanceRouterConfig
// 在所有mongos上执行FlushRouterConfig命令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FlushInstanceRouterConfig(request *FlushInstanceRouterConfigRequest) (response *FlushInstanceRouterConfigResponse, err error) {
    return c.FlushInstanceRouterConfigWithContext(context.Background(), request)
}

// FlushInstanceRouterConfig
// 在所有mongos上执行FlushRouterConfig命令
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) FlushInstanceRouterConfigWithContext(ctx context.Context, request *FlushInstanceRouterConfigRequest) (response *FlushInstanceRouterConfigResponse, err error) {
    if request == nil {
        request = NewFlushInstanceRouterConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "FlushInstanceRouterConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FlushInstanceRouterConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewFlushInstanceRouterConfigResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceCreateDBInstancesRequest() (request *InquirePriceCreateDBInstancesRequest) {
    request = &InquirePriceCreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceCreateDBInstances")
    
    
    return
}

func NewInquirePriceCreateDBInstancesResponse() (response *InquirePriceCreateDBInstancesResponse) {
    response = &InquirePriceCreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceCreateDBInstances
// 本接口（InquirePriceCreateDBInstances）用于创建数据库实例询价。本接口参数中必须传入region参数，否则无法通过校验。本接口仅允许针对购买限制范围内的实例配置进行询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) InquirePriceCreateDBInstances(request *InquirePriceCreateDBInstancesRequest) (response *InquirePriceCreateDBInstancesResponse, err error) {
    return c.InquirePriceCreateDBInstancesWithContext(context.Background(), request)
}

// InquirePriceCreateDBInstances
// 本接口（InquirePriceCreateDBInstances）用于创建数据库实例询价。本接口参数中必须传入region参数，否则无法通过校验。本接口仅允许针对购买限制范围内的实例配置进行询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MACHINETYPEERROR = "InvalidParameterValue.MachineTypeError"
//  INVALIDPARAMETERVALUE_MONGOVERSIONERROR = "InvalidParameterValue.MongoVersionError"
//  INVALIDPARAMETERVALUE_REPLICASETNUMERROR = "InvalidParameterValue.ReplicaSetNumError"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) InquirePriceCreateDBInstancesWithContext(ctx context.Context, request *InquirePriceCreateDBInstancesRequest) (response *InquirePriceCreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceCreateDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InquirePriceCreateDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceCreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceCreateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceModifyDBInstanceSpecRequest() (request *InquirePriceModifyDBInstanceSpecRequest) {
    request = &InquirePriceModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceModifyDBInstanceSpec")
    
    
    return
}

func NewInquirePriceModifyDBInstanceSpecResponse() (response *InquirePriceModifyDBInstanceSpecResponse) {
    response = &InquirePriceModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceModifyDBInstanceSpec
// 本接口 (InquirePriceModifyDBInstanceSpec) 用于查询实例配置变更后的价格。
//
// 可能返回的错误码:
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceModifyDBInstanceSpec(request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    return c.InquirePriceModifyDBInstanceSpecWithContext(context.Background(), request)
}

// InquirePriceModifyDBInstanceSpec
// 本接口 (InquirePriceModifyDBInstanceSpec) 用于查询实例配置变更后的价格。
//
// 可能返回的错误码:
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceModifyDBInstanceSpecWithContext(ctx context.Context, request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDBInstanceSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InquirePriceModifyDBInstanceSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceModifyDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewInquirePriceRenewDBInstancesRequest() (request *InquirePriceRenewDBInstancesRequest) {
    request = &InquirePriceRenewDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InquirePriceRenewDBInstances")
    
    
    return
}

func NewInquirePriceRenewDBInstancesResponse() (response *InquirePriceRenewDBInstancesResponse) {
    response = &InquirePriceRenewDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquirePriceRenewDBInstances
// 本接口 (InquiryPriceRenewDBInstances) 用于续费包年包月实例询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceRenewDBInstances(request *InquirePriceRenewDBInstancesRequest) (response *InquirePriceRenewDBInstancesResponse, err error) {
    return c.InquirePriceRenewDBInstancesWithContext(context.Background(), request)
}

// InquirePriceRenewDBInstances
// 本接口 (InquiryPriceRenewDBInstances) 用于续费包年包月实例询价。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceRenewDBInstancesWithContext(ctx context.Context, request *InquirePriceRenewDBInstancesRequest) (response *InquirePriceRenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InquirePriceRenewDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewInstanceEnableSSLRequest() (request *InstanceEnableSSLRequest) {
    request = &InstanceEnableSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "InstanceEnableSSL")
    
    
    return
}

func NewInstanceEnableSSLResponse() (response *InstanceEnableSSLResponse) {
    response = &InstanceEnableSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InstanceEnableSSL
// 本接口（InstanceEnableSSL）用于设置实例SSL状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) InstanceEnableSSL(request *InstanceEnableSSLRequest) (response *InstanceEnableSSLResponse, err error) {
    return c.InstanceEnableSSLWithContext(context.Background(), request)
}

// InstanceEnableSSL
// 本接口（InstanceEnableSSL）用于设置实例SSL状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_KERNELVERSIONNOTSUPPORT = "UnsupportedOperation.KernelVersionNotSupport"
//  UNSUPPORTEDOPERATION_SECONDARYVERSIONNOTSUPPORTAUDIT = "UnsupportedOperation.SecondaryVersionNotSupportAudit"
//  UNSUPPORTEDOPERATION_VERSIONNOTSUPPORT = "UnsupportedOperation.VersionNotSupport"
func (c *Client) InstanceEnableSSLWithContext(ctx context.Context, request *InstanceEnableSSLRequest) (response *InstanceEnableSSLResponse, err error) {
    if request == nil {
        request = NewInstanceEnableSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "InstanceEnableSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("InstanceEnableSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewInstanceEnableSSLResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "IsolateDBInstance")
    
    
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDBInstance
// 本接口（IsolateDBInstance）用于隔离 MongoDB 云数据库按量计费实例。隔离后实例保留在回收站中，不能再写入数据。隔离一定时间后，实例会彻底删除，回收站保存时间请参考按量计费的服务条款。已删除的按量计费实例无法恢复，请谨慎操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PREPAIDINSTANCEUNABLETOISOLATE = "InvalidParameterValue.PrePaidInstanceUnableToIsolate"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// 本接口（IsolateDBInstance）用于隔离 MongoDB 云数据库按量计费实例。隔离后实例保留在回收站中，不能再写入数据。隔离一定时间后，实例会彻底删除，回收站保存时间请参考按量计费的服务条款。已删除的按量计费实例无法恢复，请谨慎操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PREPAIDINSTANCEUNABLETOISOLATE = "InvalidParameterValue.PrePaidInstanceUnableToIsolate"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "IsolateDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewKillOpsRequest() (request *KillOpsRequest) {
    request = &KillOpsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "KillOps")
    
    
    return
}

func NewKillOpsResponse() (response *KillOpsResponse) {
    response = &KillOpsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillOps
// 本接口（KillOps）用于终止 MongoDB 云数据库实例上执行的特定操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_NODENOTFOUNDINREPLICA = "InvalidParameterValue.NodeNotFoundInReplica"
func (c *Client) KillOps(request *KillOpsRequest) (response *KillOpsResponse, err error) {
    return c.KillOpsWithContext(context.Background(), request)
}

// KillOps
// 本接口（KillOps）用于终止 MongoDB 云数据库实例上执行的特定操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_NODENOTFOUNDINREPLICA = "InvalidParameterValue.NodeNotFoundInReplica"
func (c *Client) KillOpsWithContext(ctx context.Context, request *KillOpsRequest) (response *KillOpsResponse, err error) {
    if request == nil {
        request = NewKillOpsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "KillOps")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillOps require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillOpsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditServiceRequest() (request *ModifyAuditServiceRequest) {
    request = &ModifyAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyAuditService")
    
    
    return
}

func NewModifyAuditServiceResponse() (response *ModifyAuditServiceResponse) {
    response = &ModifyAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计策略的服务配置，包括审计日志保存时长等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyAuditService(request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    return c.ModifyAuditServiceWithContext(context.Background(), request)
}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计策略的服务配置，包括审计日志保存时长等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ModifyAuditServiceWithContext(ctx context.Context, request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    if request == nil {
        request = NewModifyAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceNetworkAddressRequest() (request *ModifyDBInstanceNetworkAddressRequest) {
    request = &ModifyDBInstanceNetworkAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceNetworkAddress")
    
    
    return
}

func NewModifyDBInstanceNetworkAddressResponse() (response *ModifyDBInstanceNetworkAddressResponse) {
    response = &ModifyDBInstanceNetworkAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceNetworkAddress
// 本接口（ModifyDBInstanceNetworkAddress）用于修改云数据库实例的网络信息，支持基础网络切换为私有网络、私有网络切换私有网络。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWMODIFYADDRAFTEROPENWANSERVICE = "FailedOperation.NotAllowModifyAddrAfterOpenWanService"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVIP = "InvalidParameter.InvalidVip"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
func (c *Client) ModifyDBInstanceNetworkAddress(request *ModifyDBInstanceNetworkAddressRequest) (response *ModifyDBInstanceNetworkAddressResponse, err error) {
    return c.ModifyDBInstanceNetworkAddressWithContext(context.Background(), request)
}

// ModifyDBInstanceNetworkAddress
// 本接口（ModifyDBInstanceNetworkAddress）用于修改云数据库实例的网络信息，支持基础网络切换为私有网络、私有网络切换私有网络。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWMODIFYADDRAFTEROPENWANSERVICE = "FailedOperation.NotAllowModifyAddrAfterOpenWanService"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDVIP = "InvalidParameter.InvalidVip"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
func (c *Client) ModifyDBInstanceNetworkAddressWithContext(ctx context.Context, request *ModifyDBInstanceNetworkAddressRequest) (response *ModifyDBInstanceNetworkAddressResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceNetworkAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceNetworkAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNetworkAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNetworkAddressResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceParamTplRequest() (request *ModifyDBInstanceParamTplRequest) {
    request = &ModifyDBInstanceParamTplRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceParamTpl")
    
    
    return
}

func NewModifyDBInstanceParamTplResponse() (response *ModifyDBInstanceParamTplResponse) {
    response = &ModifyDBInstanceParamTplResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceParamTpl
// 本接口(ModifyDBInstanceParamTpl )用于修改MongoDB云数据库实例的参数模板。
//
// **说明：ModifyDBInstanceParamTpl  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ModifyDBInstanceParamTpl(request *ModifyDBInstanceParamTplRequest) (response *ModifyDBInstanceParamTplResponse, err error) {
    return c.ModifyDBInstanceParamTplWithContext(context.Background(), request)
}

// ModifyDBInstanceParamTpl
// 本接口(ModifyDBInstanceParamTpl )用于修改MongoDB云数据库实例的参数模板。
//
// **说明：ModifyDBInstanceParamTpl  API正在公测中，在此期间，该接口仅对公测用户开放**
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ModifyDBInstanceParamTplWithContext(ctx context.Context, request *ModifyDBInstanceParamTplRequest) (response *ModifyDBInstanceParamTplResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceParamTplRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceParamTpl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceParamTpl require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceParamTplResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupRequest() (request *ModifyDBInstanceSecurityGroupRequest) {
    request = &ModifyDBInstanceSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceSecurityGroup")
    
    
    return
}

func NewModifyDBInstanceSecurityGroupResponse() (response *ModifyDBInstanceSecurityGroupResponse) {
    response = &ModifyDBInstanceSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSecurityGroup
// 本接口（ModifyDBInstanceSecurityGroup）用于修改实例绑定的安全组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSecurityGroup(request *ModifyDBInstanceSecurityGroupRequest) (response *ModifyDBInstanceSecurityGroupResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroup
// 本接口（ModifyDBInstanceSecurityGroup）用于修改实例绑定的安全组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSecurityGroupWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupRequest) (response *ModifyDBInstanceSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceSecurityGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSpecRequest() (request *ModifyDBInstanceSpecRequest) {
    request = &ModifyDBInstanceSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyDBInstanceSpec")
    
    
    return
}

func NewModifyDBInstanceSpecResponse() (response *ModifyDBInstanceSpecResponse) {
    response = &ModifyDBInstanceSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceSpec
// 本接口（ModifyDBInstanceSpec）用于调整MongoDB云数据库实例配置。接口支持的售卖规格，可从查询云数据库的售卖规格（[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)）获取。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_ZONECLOSED = "InvalidParameter.ZoneClosed"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MODIFYMODEERROR = "InvalidParameterValue.ModifyModeError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OPLOGSIZEOUTOFRANGE = "InvalidParameterValue.OplogSizeOutOfRange"
//  INVALIDPARAMETERVALUE_SETDISKLESSTHANUSED = "InvalidParameterValue.SetDiskLessThanUsed"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSpec(request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    return c.ModifyDBInstanceSpecWithContext(context.Background(), request)
}

// ModifyDBInstanceSpec
// 本接口（ModifyDBInstanceSpec）用于调整MongoDB云数据库实例配置。接口支持的售卖规格，可从查询云数据库的售卖规格（[DescribeSpecInfo](https://cloud.tencent.com/document/product/240/38567)）获取。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_ZONECLOSED = "InvalidParameter.ZoneClosed"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_MODIFYMODEERROR = "InvalidParameterValue.ModifyModeError"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_OPLOGSIZEOUTOFRANGE = "InvalidParameterValue.OplogSizeOutOfRange"
//  INVALIDPARAMETERVALUE_SETDISKLESSTHANUSED = "InvalidParameterValue.SetDiskLessThanUsed"
//  INVALIDPARAMETERVALUE_SPECNOTONSALE = "InvalidParameterValue.SpecNotOnSale"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest) (response *ModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSpecRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyDBInstanceSpec")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamsRequest() (request *ModifyInstanceParamsRequest) {
    request = &ModifyInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ModifyInstanceParams")
    
    
    return
}

func NewModifyInstanceParamsResponse() (response *ModifyInstanceParamsResponse) {
    response = &ModifyInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceParams
// 本接口（ModifyInstanceParams）用于修改mongoDB实例的参数配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATIONNOTALLOWEDININSTANCELOCKING = "FailedOperation.OperationNotAllowedInInstanceLocking"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MODIFYMONGODBPARAMS = "InvalidParameter.ModifyMongodbParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MODIFYMONGODBPARAMS = "InvalidParameterValue.ModifyMongodbParams"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyInstanceParams(request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    return c.ModifyInstanceParamsWithContext(context.Background(), request)
}

// ModifyInstanceParams
// 本接口（ModifyInstanceParams）用于修改mongoDB实例的参数配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPERATIONNOTALLOWEDININSTANCELOCKING = "FailedOperation.OperationNotAllowedInInstanceLocking"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_MODIFYMONGODBPARAMS = "InvalidParameter.ModifyMongodbParams"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MODIFYMONGODBPARAMS = "InvalidParameterValue.ModifyMongodbParams"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyInstanceParamsWithContext(ctx context.Context, request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ModifyInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedDBInstanceRequest() (request *OfflineIsolatedDBInstanceRequest) {
    request = &OfflineIsolatedDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "OfflineIsolatedDBInstance")
    
    
    return
}

func NewOfflineIsolatedDBInstanceResponse() (response *OfflineIsolatedDBInstanceResponse) {
    response = &OfflineIsolatedDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OfflineIsolatedDBInstance
// 本接口(OfflineIsolatedDBInstance)用于立即下线隔离状态的云数据库实例。进行操作的实例状态必须为隔离状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALSTATUSTOOFFLINE = "InvalidParameterValue.IllegalStatusToOffline"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) OfflineIsolatedDBInstance(request *OfflineIsolatedDBInstanceRequest) (response *OfflineIsolatedDBInstanceResponse, err error) {
    return c.OfflineIsolatedDBInstanceWithContext(context.Background(), request)
}

// OfflineIsolatedDBInstance
// 本接口(OfflineIsolatedDBInstance)用于立即下线隔离状态的云数据库实例。进行操作的实例状态必须为隔离状态。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALSTATUSTOOFFLINE = "InvalidParameterValue.IllegalStatusToOffline"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) OfflineIsolatedDBInstanceWithContext(ctx context.Context, request *OfflineIsolatedDBInstanceRequest) (response *OfflineIsolatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "OfflineIsolatedDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAuditServiceRequest() (request *OpenAuditServiceRequest) {
    request = &OpenAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "OpenAuditService")
    
    
    return
}

func NewOpenAuditServiceResponse() (response *OpenAuditServiceResponse) {
    response = &OpenAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenAuditService
// 本接口(OpenAuditService)用于开通云数据库实例的审计。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// 本接口(OpenAuditService)用于开通云数据库实例的审计。
//
// 可能返回的错误码:
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_AUDITMODIFYSTATUSERROR = "InternalError.AuditModifyStatusError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INTERNALERROR_COSERROR = "InternalError.CosError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INNERCOMMONERROR = "InternalError.InnerCommonError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_JSONUNMARSHALERROR = "InternalError.JsonUnmarshalError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_AUDITPOLICYCONFLICTERROR = "OperationDenied.AuditPolicyConflictError"
//  OPERATIONDENIED_AUDITPOLICYEXISTERROR = "OperationDenied.AuditPolicyExistError"
//  OPERATIONDENIED_AUDITPOLICYOVERQUOTAERROR = "OperationDenied.AuditPolicyOverQuotaError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "OpenAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewRenameInstanceRequest() (request *RenameInstanceRequest) {
    request = &RenameInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "RenameInstance")
    
    
    return
}

func NewRenameInstanceResponse() (response *RenameInstanceResponse) {
    response = &RenameInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenameInstance
// 本接口(RenameInstance)用于修改云数据库实例的名称。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenameInstance(request *RenameInstanceRequest) (response *RenameInstanceResponse, err error) {
    return c.RenameInstanceWithContext(context.Background(), request)
}

// RenameInstance
// 本接口(RenameInstance)用于修改云数据库实例的名称。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenameInstanceWithContext(ctx context.Context, request *RenameInstanceRequest) (response *RenameInstanceResponse, err error) {
    if request == nil {
        request = NewRenameInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "RenameInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenameInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenameInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDBInstancesRequest() (request *RenewDBInstancesRequest) {
    request = &RenewDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "RenewDBInstances")
    
    
    return
}

func NewRenewDBInstancesResponse() (response *RenewDBInstancesResponse) {
    response = &RenewDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewDBInstances
// 本接口(RenewDBInstance)用于续费云数据库实例，仅支持付费模式为包年包月的实例。按量计费实例不需要续费。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenewDBInstances(request *RenewDBInstancesRequest) (response *RenewDBInstancesResponse, err error) {
    return c.RenewDBInstancesWithContext(context.Background(), request)
}

// RenewDBInstances
// 本接口(RenewDBInstance)用于续费云数据库实例，仅支持付费模式为包年包月的实例。按量计费实例不需要续费。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INVALIDTRADEOPERATION = "InvalidParameterValue.InvalidTradeOperation"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) RenewDBInstancesWithContext(ctx context.Context, request *RenewDBInstancesRequest) (response *RenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewRenewDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "RenewDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetDBInstancePasswordRequest() (request *ResetDBInstancePasswordRequest) {
    request = &ResetDBInstancePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "ResetDBInstancePassword")
    
    
    return
}

func NewResetDBInstancePasswordResponse() (response *ResetDBInstancePasswordResponse) {
    response = &ResetDBInstancePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetDBInstancePassword
// 本接口（ResetDBInstancePassword）用于重置实例访问密码。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_USERNOTFOUND = "InvalidParameterValue.UserNotFound"
func (c *Client) ResetDBInstancePassword(request *ResetDBInstancePasswordRequest) (response *ResetDBInstancePasswordResponse, err error) {
    return c.ResetDBInstancePasswordWithContext(context.Background(), request)
}

// ResetDBInstancePassword
// 本接口（ResetDBInstancePassword）用于重置实例访问密码。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_PASSWORDRULEFAILED = "InvalidParameterValue.PasswordRuleFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
//  INVALIDPARAMETERVALUE_USERNOTFOUND = "InvalidParameterValue.UserNotFound"
func (c *Client) ResetDBInstancePasswordWithContext(ctx context.Context, request *ResetDBInstancePasswordRequest) (response *ResetDBInstancePasswordResponse, err error) {
    if request == nil {
        request = NewResetDBInstancePasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "ResetDBInstancePassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDBInstancePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDBInstancePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartNodesRequest() (request *RestartNodesRequest) {
    request = &RestartNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "RestartNodes")
    
    
    return
}

func NewRestartNodesResponse() (response *RestartNodesResponse) {
    response = &RestartNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartNodes
// 本接口（RestartNodes）用于批量重启数据库节点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) RestartNodes(request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    return c.RestartNodesWithContext(context.Background(), request)
}

// RestartNodes
// 本接口（RestartNodes）用于批量重启数据库节点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) RestartNodesWithContext(ctx context.Context, request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    if request == nil {
        request = NewRestartNodesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "RestartNodes")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartNodesResponse()
    err = c.Send(request, response)
    return
}

func NewSetAccountUserPrivilegeRequest() (request *SetAccountUserPrivilegeRequest) {
    request = &SetAccountUserPrivilegeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetAccountUserPrivilege")
    
    
    return
}

func NewSetAccountUserPrivilegeResponse() (response *SetAccountUserPrivilegeResponse) {
    response = &SetAccountUserPrivilegeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetAccountUserPrivilege
// 本接口（SetAccountUserPrivilege）用于设置实例的账号权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetAccountUserPrivilege(request *SetAccountUserPrivilegeRequest) (response *SetAccountUserPrivilegeResponse, err error) {
    return c.SetAccountUserPrivilegeWithContext(context.Background(), request)
}

// SetAccountUserPrivilege
// 本接口（SetAccountUserPrivilege）用于设置实例的账号权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetAccountUserPrivilegeWithContext(ctx context.Context, request *SetAccountUserPrivilegeRequest) (response *SetAccountUserPrivilegeResponse, err error) {
    if request == nil {
        request = NewSetAccountUserPrivilegeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetAccountUserPrivilege")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAccountUserPrivilege require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAccountUserPrivilegeResponse()
    err = c.Send(request, response)
    return
}

func NewSetBackupRulesRequest() (request *SetBackupRulesRequest) {
    request = &SetBackupRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetBackupRules")
    
    
    return
}

func NewSetBackupRulesResponse() (response *SetBackupRulesResponse) {
    response = &SetBackupRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetBackupRules
// 本接口(SetBackupRules)用于设置 MongoDB 云数据库的自动备份规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
func (c *Client) SetBackupRules(request *SetBackupRulesRequest) (response *SetBackupRulesResponse, err error) {
    return c.SetBackupRulesWithContext(context.Background(), request)
}

// SetBackupRules
// 本接口(SetBackupRules)用于设置 MongoDB 云数据库的自动备份规则。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
func (c *Client) SetBackupRulesWithContext(ctx context.Context, request *SetBackupRulesRequest) (response *SetBackupRulesResponse, err error) {
    if request == nil {
        request = NewSetBackupRulesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetBackupRules")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetBackupRules require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetBackupRulesResponse()
    err = c.Send(request, response)
    return
}

func NewSetDBInstanceDeletionProtectionRequest() (request *SetDBInstanceDeletionProtectionRequest) {
    request = &SetDBInstanceDeletionProtectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetDBInstanceDeletionProtection")
    
    
    return
}

func NewSetDBInstanceDeletionProtectionResponse() (response *SetDBInstanceDeletionProtectionResponse) {
    response = &SetDBInstanceDeletionProtectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetDBInstanceDeletionProtection
// 本接口（SetDBInstanceDeletionProtection）用于设置实例销毁保护
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) SetDBInstanceDeletionProtection(request *SetDBInstanceDeletionProtectionRequest) (response *SetDBInstanceDeletionProtectionResponse, err error) {
    return c.SetDBInstanceDeletionProtectionWithContext(context.Background(), request)
}

// SetDBInstanceDeletionProtection
// 本接口（SetDBInstanceDeletionProtection）用于设置实例销毁保护
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENDELETED = "InvalidParameterValue.InstanceHasBeenDeleted"
//  INVALIDPARAMETERVALUE_INSTANCEHASBEENISOLATED = "InvalidParameterValue.InstanceHasBeenIsolated"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
func (c *Client) SetDBInstanceDeletionProtectionWithContext(ctx context.Context, request *SetDBInstanceDeletionProtectionRequest) (response *SetDBInstanceDeletionProtectionResponse, err error) {
    if request == nil {
        request = NewSetDBInstanceDeletionProtectionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetDBInstanceDeletionProtection")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetDBInstanceDeletionProtection require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetDBInstanceDeletionProtectionResponse()
    err = c.Send(request, response)
    return
}

func NewSetInstanceMaintenanceRequest() (request *SetInstanceMaintenanceRequest) {
    request = &SetInstanceMaintenanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "SetInstanceMaintenance")
    
    
    return
}

func NewSetInstanceMaintenanceResponse() (response *SetInstanceMaintenanceResponse) {
    response = &SetInstanceMaintenanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetInstanceMaintenance
// 本接口（SetInstanceMaintenance ） 用于设置实例维护时间窗。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetInstanceMaintenance(request *SetInstanceMaintenanceRequest) (response *SetInstanceMaintenanceResponse, err error) {
    return c.SetInstanceMaintenanceWithContext(context.Background(), request)
}

// SetInstanceMaintenance
// 本接口（SetInstanceMaintenance ） 用于设置实例维护时间窗。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetInstanceMaintenanceWithContext(ctx context.Context, request *SetInstanceMaintenanceRequest) (response *SetInstanceMaintenanceResponse, err error) {
    if request == nil {
        request = NewSetInstanceMaintenanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "SetInstanceMaintenance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetInstanceMaintenance require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetInstanceMaintenanceResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateDBInstancesRequest() (request *TerminateDBInstancesRequest) {
    request = &TerminateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "TerminateDBInstances")
    
    
    return
}

func NewTerminateDBInstancesResponse() (response *TerminateDBInstancesResponse) {
    response = &TerminateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateDBInstances
// 本接口（TerminateDBInstances）用于退还包年包月计费实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) TerminateDBInstances(request *TerminateDBInstancesRequest) (response *TerminateDBInstancesResponse, err error) {
    return c.TerminateDBInstancesWithContext(context.Background(), request)
}

// TerminateDBInstances
// 本接口（TerminateDBInstances）用于退还包年包月计费实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DELETIONPROTECTIONENABLED = "FailedOperation.DeletionProtectionEnabled"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKAPPIDFAILED = "InvalidParameterValue.CheckAppIdFailed"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) TerminateDBInstancesWithContext(ctx context.Context, request *TerminateDBInstancesRequest) (response *TerminateDBInstancesResponse, err error) {
    if request == nil {
        request = NewTerminateDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "TerminateDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDBInstanceKernelVersionRequest() (request *UpgradeDBInstanceKernelVersionRequest) {
    request = &UpgradeDBInstanceKernelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "UpgradeDBInstanceKernelVersion")
    
    
    return
}

func NewUpgradeDBInstanceKernelVersionResponse() (response *UpgradeDBInstanceKernelVersionResponse) {
    response = &UpgradeDBInstanceKernelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDBInstanceKernelVersion
// 本接口(UpgradeDBInstanceKernelVersion)用于升级数据库实例内核版本。
//
// 可能返回的错误码:
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) UpgradeDBInstanceKernelVersion(request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
    return c.UpgradeDBInstanceKernelVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceKernelVersion
// 本接口(UpgradeDBInstanceKernelVersion)用于升级数据库实例内核版本。
//
// 可能返回的错误码:
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_LOCKFAILED = "InvalidParameterValue.LockFailed"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) UpgradeDBInstanceKernelVersionWithContext(ctx context.Context, request *UpgradeDBInstanceKernelVersionRequest) (response *UpgradeDBInstanceKernelVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceKernelVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "UpgradeDBInstanceKernelVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDBInstanceKernelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDBInstanceKernelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeDbInstanceVersionRequest() (request *UpgradeDbInstanceVersionRequest) {
    request = &UpgradeDbInstanceVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mongodb", APIVersion, "UpgradeDbInstanceVersion")
    
    
    return
}

func NewUpgradeDbInstanceVersionResponse() (response *UpgradeDbInstanceVersionResponse) {
    response = &UpgradeDbInstanceVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeDbInstanceVersion
// 本接口用于跨版本升级数据库内核。当前仅支持3.6版本升级为4.0版本、4.0版本升级为4.2版本、4.2版本升级为4.4版本及4.4版本升级为5.0版本。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATEERROR = "InternalError.DBOperateError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) UpgradeDbInstanceVersion(request *UpgradeDbInstanceVersionRequest) (response *UpgradeDbInstanceVersionResponse, err error) {
    return c.UpgradeDbInstanceVersionWithContext(context.Background(), request)
}

// UpgradeDbInstanceVersion
// 本接口用于跨版本升级数据库内核。当前仅支持3.6版本升级为4.0版本、4.0版本升级为4.2版本、4.2版本升级为4.4版本及4.4版本升级为5.0版本。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATEERROR = "InternalError.DBOperateError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) UpgradeDbInstanceVersionWithContext(ctx context.Context, request *UpgradeDbInstanceVersionRequest) (response *UpgradeDbInstanceVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDbInstanceVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "mongodb", APIVersion, "UpgradeDbInstanceVersion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeDbInstanceVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeDbInstanceVersionResponse()
    err = c.Send(request, response)
    return
}
