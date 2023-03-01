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
// 本接口(CreateAccountUser)用于创建mongodb实例账号。
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
// 本接口(CreateAccountUser)用于创建mongodb实例账号。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAccountUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAccountUserResponse()
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
// 备份实例接口
//
// 可能返回的错误码:
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
// 备份实例接口
//
// 可能返回的错误码:
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
// 本接口(CreateDBInstance)用于创建包年包月的MongoDB云数据库实例。接口支持的售卖规格，可从查询云数据库的售卖规格（DescribeSpecInfo）获取。
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
// 本接口(CreateDBInstance)用于创建包年包月的MongoDB云数据库实例。接口支持的售卖规格，可从查询云数据库的售卖规格（DescribeSpecInfo）获取。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceHour require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceHourResponse()
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
// 本接口(DescribeAccountUsers)用于获取当前实例的全部账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
func (c *Client) DescribeAccountUsers(request *DescribeAccountUsersRequest) (response *DescribeAccountUsersResponse, err error) {
    return c.DescribeAccountUsersWithContext(context.Background(), request)
}

// DescribeAccountUsers
// 本接口(DescribeAccountUsers)用于获取当前实例的全部账号。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
func (c *Client) DescribeAccountUsersWithContext(ctx context.Context, request *DescribeAccountUsersRequest) (response *DescribeAccountUsersResponse, err error) {
    if request == nil {
        request = NewDescribeAccountUsersRequest()
    }
    
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
// 查询异步任务状态接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// 查询异步任务状态接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
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
// 查询备份下载任务信息
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
// 查询备份下载任务信息
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDownloadTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDownloadTaskResponse()
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
// 本接口(DescribeClientConnections)用于查询实例客户端连接信息，包括连接IP和连接数量。
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
// 本接口(DescribeClientConnections)用于查询实例客户端连接信息，包括连接IP和连接数量。
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
// 本接口(DescribeCurrentOp)用于查询MongoDB云数据库实例的当前正在执行的操作。
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
// 本接口(DescribeCurrentOp)用于查询MongoDB云数据库实例的当前正在执行的操作。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceDeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceDealResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceNodeProperty require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceNodePropertyResponse()
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
// 本接口(DescribeDBInstances)用于查询云数据库实例列表，支持通过项目ID、实例ID、实例状态等过滤条件来筛选实例。支持查询主实例、灾备实例和只读实例信息列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口(DescribeDBInstances)用于查询云数据库实例列表，支持通过项目ID、实例ID、实例状态等过滤条件来筛选实例。支持查询主实例、灾备实例和只读实例信息列表。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCESTATUS = "InvalidParameterValue.IllegalInstanceStatus"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_VPCIDORSUBNETIDNOTFOUND = "InvalidParameterValue.VpcIdOrSubnetIdNotFound"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
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
// 本接口(DescribeInstanceParams)用于查询当前实例可修改的参数列表。
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
// 本接口(DescribeInstanceParams)用于查询当前实例可修改的参数列表。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
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
// 查询实例绑定的安全组
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
// 查询实例绑定的安全组
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
// 本接口(DescribeSpecInfo)用于查询实例的售卖规格。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) DescribeSpecInfo(request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    return c.DescribeSpecInfoWithContext(context.Background(), request)
}

// DescribeSpecInfo
// 本接口(DescribeSpecInfo)用于查询实例的售卖规格。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REGIONERROR = "InvalidParameterValue.RegionError"
//  INVALIDPARAMETERVALUE_ZONEERROR = "InvalidParameterValue.ZoneError"
func (c *Client) DescribeSpecInfoWithContext(ctx context.Context, request *DescribeSpecInfoRequest) (response *DescribeSpecInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSpecInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpecInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpecInfoResponse()
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
// 本接口 (InquirePriceModifyDBInstanceSpec) 用于调整实例的配置询价。
//
// 可能返回的错误码:
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceModifyDBInstanceSpec(request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    return c.InquirePriceModifyDBInstanceSpecWithContext(context.Background(), request)
}

// InquirePriceModifyDBInstanceSpec
// 本接口 (InquirePriceModifyDBInstanceSpec) 用于调整实例的配置询价。
//
// 可能返回的错误码:
//  INTERNALERROR_FINDINSTANCEFAILED = "InternalError.FindInstanceFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceModifyDBInstanceSpecWithContext(ctx context.Context, request *InquirePriceModifyDBInstanceSpecRequest) (response *InquirePriceModifyDBInstanceSpecResponse, err error) {
    if request == nil {
        request = NewInquirePriceModifyDBInstanceSpecRequest()
    }
    
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
//  INVALIDPARAMETERVALUE_ZONECLOSED = "InvalidParameterValue.ZoneClosed"
func (c *Client) InquirePriceRenewDBInstancesWithContext(ctx context.Context, request *InquirePriceRenewDBInstancesRequest) (response *InquirePriceRenewDBInstancesResponse, err error) {
    if request == nil {
        request = NewInquirePriceRenewDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquirePriceRenewDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquirePriceRenewDBInstancesResponse()
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
// 本接口(IsolateDBInstance)用于隔离MongoDB云数据库按量计费实例。隔离后实例保留在回收站中，不能再写入数据。隔离一定时间后，实例会彻底删除，回收站保存时间请参考按量计费的服务条款。在隔离中的按量计费实例无法恢复，请谨慎操作。
//
// 可能返回的错误码:
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
// 本接口(IsolateDBInstance)用于隔离MongoDB云数据库按量计费实例。隔离后实例保留在回收站中，不能再写入数据。隔离一定时间后，实例会彻底删除，回收站保存时间请参考按量计费的服务条款。在隔离中的按量计费实例无法恢复，请谨慎操作。
//
// 可能返回的错误码:
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
// 本接口(KillOps)用于终止MongoDB云数据库实例上执行的特定操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_NODENOTFOUNDINREPLICA = "InvalidParameterValue.NodeNotFoundInReplica"
func (c *Client) KillOps(request *KillOpsRequest) (response *KillOpsResponse, err error) {
    return c.KillOpsWithContext(context.Background(), request)
}

// KillOps
// 本接口(KillOps)用于终止MongoDB云数据库实例上执行的特定操作。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_NODENOTFOUNDINREPLICA = "InvalidParameterValue.NodeNotFoundInReplica"
func (c *Client) KillOpsWithContext(ctx context.Context, request *KillOpsRequest) (response *KillOpsResponse, err error) {
    if request == nil {
        request = NewKillOpsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillOps require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillOpsResponse()
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
// 本接口(ModifyDBInstanceNetworkAddress)用于修改云数据库实例的网络信息，可进行基础网络转VPC网络和VPC网络之间的变换。
//
// 可能返回的错误码:
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
// 本接口(ModifyDBInstanceNetworkAddress)用于修改云数据库实例的网络信息，可进行基础网络转VPC网络和VPC网络之间的变换。
//
// 可能返回的错误码:
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceNetworkAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceNetworkAddressResponse()
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
// 修改实例绑定的安全组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSecurityGroup(request *ModifyDBInstanceSecurityGroupRequest) (response *ModifyDBInstanceSecurityGroupResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroup
// 修改实例绑定的安全组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_NOTFOUNDINSTANCE = "InvalidParameterValue.NotFoundInstance"
//  INVALIDPARAMETERVALUE_SECURITYGROUPID = "InvalidParameterValue.SecurityGroupId"
//  INVALIDPARAMETERVALUE_STATUSABNORMAL = "InvalidParameterValue.StatusAbnormal"
func (c *Client) ModifyDBInstanceSecurityGroupWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupRequest) (response *ModifyDBInstanceSecurityGroupResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupRequest()
    }
    
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
// 本接口(ModifyDBInstanceSpec)用于调整MongoDB云数据库实例配置。接口支持的售卖规格，可从查询云数据库的售卖规格（DescribeSpecInfo）获取。
//
// 可能返回的错误码:
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
// 本接口(ModifyDBInstanceSpec)用于调整MongoDB云数据库实例配置。接口支持的售卖规格，可从查询云数据库的售卖规格（DescribeSpecInfo）获取。
//
// 可能返回的错误码:
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSpecResponse()
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
func (c *Client) OfflineIsolatedDBInstanceWithContext(ctx context.Context, request *OfflineIsolatedDBInstanceRequest) (response *OfflineIsolatedDBInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedDBInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedDBInstanceResponse()
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
// 修改实例用户的密码
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
// 修改实例用户的密码
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetDBInstancePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetDBInstancePasswordResponse()
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
// 本接口(SetAccountUserPrivilege)用于设置mongodb实例的账号权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetAccountUserPrivilege(request *SetAccountUserPrivilegeRequest) (response *SetAccountUserPrivilegeResponse, err error) {
    return c.SetAccountUserPrivilegeWithContext(context.Background(), request)
}

// SetAccountUserPrivilege
// 本接口(SetAccountUserPrivilege)用于设置mongodb实例的账号权限。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CHECKAPPIDFAILED = "InternalError.CheckAppIdFailed"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SetAccountUserPrivilegeWithContext(ctx context.Context, request *SetAccountUserPrivilegeRequest) (response *SetAccountUserPrivilegeResponse, err error) {
    if request == nil {
        request = NewSetAccountUserPrivilegeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetAccountUserPrivilege require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetAccountUserPrivilegeResponse()
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
// 本接口（TerminateDBInstances）可将包年包月实例退还隔离。
//
// 可能返回的错误码:
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
// 本接口（TerminateDBInstances）可将包年包月实例退还隔离。
//
// 可能返回的错误码:
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateDBInstancesResponse()
    err = c.Send(request, response)
    return
}
