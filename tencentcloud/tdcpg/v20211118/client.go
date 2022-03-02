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

package v20211118

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-11-18"

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


func NewCloneClusterToPointInTimeRequest() (request *CloneClusterToPointInTimeRequest) {
    request = &CloneClusterToPointInTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "CloneClusterToPointInTime")
    
    
    return
}

func NewCloneClusterToPointInTimeResponse() (response *CloneClusterToPointInTimeResponse) {
    response = &CloneClusterToPointInTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloneClusterToPointInTime
// 使用指定时间点的备份克隆一个新的集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_SPECSTORAGELACK = "FailedOperation.SpecStorageLack"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BACKUPDATAPOINTINVALID = "InvalidParameterValue.BackupDataPointInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SOURCEBACKUPCLUSTERIDINVALID = "InvalidParameterValue.SourceBackupClusterIdInvalid"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_VPCDENIEDERROR = "InvalidParameterValue.VpcDeniedError"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  INVALIDPARAMETERVALUE_VPCSUBNETIPLACK = "InvalidParameterValue.VpcSubnetIpLack"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERCLUSTERLIMIT = "LimitExceeded.UserClusterLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloneClusterToPointInTime(request *CloneClusterToPointInTimeRequest) (response *CloneClusterToPointInTimeResponse, err error) {
    if request == nil {
        request = NewCloneClusterToPointInTimeRequest()
    }
    
    response = NewCloneClusterToPointInTimeResponse()
    err = c.Send(request, response)
    return
}

// CloneClusterToPointInTime
// 使用指定时间点的备份克隆一个新的集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_SPECSTORAGELACK = "FailedOperation.SpecStorageLack"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BACKUPDATAPOINTINVALID = "InvalidParameterValue.BackupDataPointInvalid"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_SOURCEBACKUPCLUSTERIDINVALID = "InvalidParameterValue.SourceBackupClusterIdInvalid"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  INVALIDPARAMETERVALUE_VPCDENIEDERROR = "InvalidParameterValue.VpcDeniedError"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  INVALIDPARAMETERVALUE_VPCSUBNETIPLACK = "InvalidParameterValue.VpcSubnetIpLack"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERCLUSTERLIMIT = "LimitExceeded.UserClusterLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloneClusterToPointInTimeWithContext(ctx context.Context, request *CloneClusterToPointInTimeRequest) (response *CloneClusterToPointInTimeResponse, err error) {
    if request == nil {
        request = NewCloneClusterToPointInTimeRequest()
    }
    request.SetContext(ctx)
    
    response = NewCloneClusterToPointInTimeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
    request = &CreateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "CreateCluster")
    
    
    return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
    response = &CreateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_VPCDENIEDERROR = "InvalidParameterValue.VpcDeniedError"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  INVALIDPARAMETERVALUE_VPCSUBNETIPLACK = "InvalidParameterValue.VpcSubnetIpLack"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERCLUSTERLIMIT = "LimitExceeded.UserClusterLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

// CreateCluster
// 创建集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_GETVPCFAILED = "FailedOperation.GetVpcFailed"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INVALIDDBVERSION = "InvalidParameterValue.InvalidDBVersion"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  INVALIDPARAMETERVALUE_VPCDENIEDERROR = "InvalidParameterValue.VpcDeniedError"
//  INVALIDPARAMETERVALUE_VPCNOTFOUND = "InvalidParameterValue.VpcNotFound"
//  INVALIDPARAMETERVALUE_VPCSUBNETIPLACK = "InvalidParameterValue.VpcSubnetIpLack"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  LIMITEXCEEDED_USERCLUSTERLIMIT = "LimitExceeded.UserClusterLimit"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
    if request == nil {
        request = NewCreateClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClusterInstancesRequest() (request *CreateClusterInstancesRequest) {
    request = &CreateClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "CreateClusterInstances")
    
    
    return
}

func NewCreateClusterInstancesResponse() (response *CreateClusterInstancesResponse) {
    response = &CreateClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateClusterInstances
// 在集群中新建实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusterInstances(request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewCreateClusterInstancesRequest()
    }
    
    response = NewCreateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// CreateClusterInstances
// 在集群中新建实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  LIMITEXCEEDED_CLUSTERINSTANCELIMIT = "LimitExceeded.ClusterInstanceLimit"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) CreateClusterInstancesWithContext(ctx context.Context, request *CreateClusterInstancesRequest) (response *CreateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewCreateClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
    request = &DeleteClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DeleteCluster")
    
    
    return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
    response = &DeleteClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCluster
// 删除集群，集群中的实例和数据都将被删除，且无法恢复。只有当集群状态处于isolated(已隔离)时才生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

// DeleteCluster
// 删除集群，集群中的实例和数据都将被删除，且无法恢复。只有当集群状态处于isolated(已隔离)时才生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterWithContext(ctx context.Context, request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
    if request == nil {
        request = NewDeleteClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteClusterResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClusterInstancesRequest() (request *DeleteClusterInstancesRequest) {
    request = &DeleteClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DeleteClusterInstances")
    
    
    return
}

func NewDeleteClusterInstancesResponse() (response *DeleteClusterInstancesResponse) {
    response = &DeleteClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteClusterInstances
// 删除实例。只有当实例状态处于isolated(已隔离)时才生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterInstances(request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteClusterInstancesRequest()
    }
    
    response = NewDeleteClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// DeleteClusterInstances
// 删除实例。只有当实例状态处于isolated(已隔离)时才生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DeleteClusterInstancesWithContext(ctx context.Context, request *DeleteClusterInstancesRequest) (response *DeleteClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDeleteClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
    request = &DescribeAccountsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeAccounts")
    
    
    return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
    response = &DescribeAccountsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccounts
// 查询数据库账号信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAccounts
// 查询数据库账号信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAccountsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterBackupsRequest() (request *DescribeClusterBackupsRequest) {
    request = &DescribeClusterBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeClusterBackups")
    
    
    return
}

func NewDescribeClusterBackupsResponse() (response *DescribeClusterBackupsResponse) {
    response = &DescribeClusterBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterBackups
// 查询集群的备份集
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterBackups(request *DescribeClusterBackupsRequest) (response *DescribeClusterBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterBackupsRequest()
    }
    
    response = NewDescribeClusterBackupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeClusterBackups
// 查询集群的备份集
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterBackupsWithContext(ctx context.Context, request *DescribeClusterBackupsRequest) (response *DescribeClusterBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterBackupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeClusterBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterEndpointsRequest() (request *DescribeClusterEndpointsRequest) {
    request = &DescribeClusterEndpointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeClusterEndpoints")
    
    
    return
}

func NewDescribeClusterEndpointsResponse() (response *DescribeClusterEndpointsResponse) {
    response = &DescribeClusterEndpointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterEndpoints
// 查询集群接入点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterEndpoints(request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointsRequest()
    }
    
    response = NewDescribeClusterEndpointsResponse()
    err = c.Send(request, response)
    return
}

// DescribeClusterEndpoints
// 查询集群接入点信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterEndpointsWithContext(ctx context.Context, request *DescribeClusterEndpointsRequest) (response *DescribeClusterEndpointsResponse, err error) {
    if request == nil {
        request = NewDescribeClusterEndpointsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeClusterEndpointsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
    request = &DescribeClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeClusterInstances")
    
    
    return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
    response = &DescribeClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterInstances
// 查询实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// DescribeClusterInstances
// 查询实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusterInstancesWithContext(ctx context.Context, request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterRecoveryTimeRangeRequest() (request *DescribeClusterRecoveryTimeRangeRequest) {
    request = &DescribeClusterRecoveryTimeRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeClusterRecoveryTimeRange")
    
    
    return
}

func NewDescribeClusterRecoveryTimeRangeResponse() (response *DescribeClusterRecoveryTimeRangeResponse) {
    response = &DescribeClusterRecoveryTimeRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusterRecoveryTimeRange
// 查询集群可回档时间范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BACKUPDATAPOINTINVALID = "InvalidParameterValue.BackupDataPointInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterRecoveryTimeRange(request *DescribeClusterRecoveryTimeRangeRequest) (response *DescribeClusterRecoveryTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRecoveryTimeRangeRequest()
    }
    
    response = NewDescribeClusterRecoveryTimeRangeResponse()
    err = c.Send(request, response)
    return
}

// DescribeClusterRecoveryTimeRange
// 查询集群可回档时间范围
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BACKUPDATAPOINTINVALID = "InvalidParameterValue.BackupDataPointInvalid"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_STORAGEPOOLNOTFOUND = "InvalidParameterValue.StoragePoolNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeClusterRecoveryTimeRangeWithContext(ctx context.Context, request *DescribeClusterRecoveryTimeRangeRequest) (response *DescribeClusterRecoveryTimeRangeResponse, err error) {
    if request == nil {
        request = NewDescribeClusterRecoveryTimeRangeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeClusterRecoveryTimeRangeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusters
// 查询集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

// DescribeClusters
// 查询集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByDealNameRequest() (request *DescribeResourcesByDealNameRequest) {
    request = &DescribeResourcesByDealNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "DescribeResourcesByDealName")
    
    
    return
}

func NewDescribeResourcesByDealNameResponse() (response *DescribeResourcesByDealNameResponse) {
    response = &DescribeResourcesByDealNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourcesByDealName
// 根据订单号获取资源信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealName(request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByDealNameRequest()
    }
    
    response = NewDescribeResourcesByDealNameResponse()
    err = c.Send(request, response)
    return
}

// DescribeResourcesByDealName
// 根据订单号获取资源信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DEALNAMENOTFOUND = "InvalidParameterValue.DealNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) DescribeResourcesByDealNameWithContext(ctx context.Context, request *DescribeResourcesByDealNameRequest) (response *DescribeResourcesByDealNameResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByDealNameRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeResourcesByDealNameResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateClusterRequest() (request *IsolateClusterRequest) {
    request = &IsolateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "IsolateCluster")
    
    
    return
}

func NewIsolateClusterResponse() (response *IsolateClusterResponse) {
    response = &IsolateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateCluster
// 隔离集群，集群的接入点网络将会断掉无法连接使用数据库。只有当集群状态处于running(运行中)时才生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IsolateCluster(request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    if request == nil {
        request = NewIsolateClusterRequest()
    }
    
    response = NewIsolateClusterResponse()
    err = c.Send(request, response)
    return
}

// IsolateCluster
// 隔离集群，集群的接入点网络将会断掉无法连接使用数据库。只有当集群状态处于running(运行中)时才生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IsolateClusterWithContext(ctx context.Context, request *IsolateClusterRequest) (response *IsolateClusterResponse, err error) {
    if request == nil {
        request = NewIsolateClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewIsolateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateClusterInstancesRequest() (request *IsolateClusterInstancesRequest) {
    request = &IsolateClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "IsolateClusterInstances")
    
    
    return
}

func NewIsolateClusterInstancesResponse() (response *IsolateClusterInstancesResponse) {
    response = &IsolateClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateClusterInstances
// 隔离实例。此接口只针对状态为running的实例生效，使用场景包括：
//
//  - 批量隔离集群内所有的实例
//
//  - 在读写实例为running(运行中)时，单个/批量隔离只读实例
//
//  - 集群内所有只读实例为isolated(已隔离)时，单独隔离读写实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateClusterInstances(request *IsolateClusterInstancesRequest) (response *IsolateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewIsolateClusterInstancesRequest()
    }
    
    response = NewIsolateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// IsolateClusterInstances
// 隔离实例。此接口只针对状态为running的实例生效，使用场景包括：
//
//  - 批量隔离集群内所有的实例
//
//  - 在读写实例为running(运行中)时，单个/批量隔离只读实例
//
//  - 集群内所有只读实例为isolated(已隔离)时，单独隔离读写实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) IsolateClusterInstancesWithContext(ctx context.Context, request *IsolateClusterInstancesRequest) (response *IsolateClusterInstancesResponse, err error) {
    if request == nil {
        request = NewIsolateClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewIsolateClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
    request = &ModifyAccountDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "ModifyAccountDescription")
    
    
    return
}

func NewModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
    response = &ModifyAccountDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAccountDescription
// 修改数据库账号描述
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

// ModifyAccountDescription
// 修改数据库账号描述
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyAccountDescriptionRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAccountDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterEndpointWanStatusRequest() (request *ModifyClusterEndpointWanStatusRequest) {
    request = &ModifyClusterEndpointWanStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "ModifyClusterEndpointWanStatus")
    
    
    return
}

func NewModifyClusterEndpointWanStatusResponse() (response *ModifyClusterEndpointWanStatusResponse) {
    response = &ModifyClusterEndpointWanStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterEndpointWanStatus
// 开启或者关闭接入点外网
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ENDPOINTNOTFOUND = "InvalidParameterValue.EndpointNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterEndpointWanStatus(request *ModifyClusterEndpointWanStatusRequest) (response *ModifyClusterEndpointWanStatusResponse, err error) {
    if request == nil {
        request = NewModifyClusterEndpointWanStatusRequest()
    }
    
    response = NewModifyClusterEndpointWanStatusResponse()
    err = c.Send(request, response)
    return
}

// ModifyClusterEndpointWanStatus
// 开启或者关闭接入点外网
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ENDPOINTNOTFOUND = "InvalidParameterValue.EndpointNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterEndpointWanStatusWithContext(ctx context.Context, request *ModifyClusterEndpointWanStatusRequest) (response *ModifyClusterEndpointWanStatusResponse, err error) {
    if request == nil {
        request = NewModifyClusterEndpointWanStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyClusterEndpointWanStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterInstancesSpecRequest() (request *ModifyClusterInstancesSpecRequest) {
    request = &ModifyClusterInstancesSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "ModifyClusterInstancesSpec")
    
    
    return
}

func NewModifyClusterInstancesSpecResponse() (response *ModifyClusterInstancesSpecResponse) {
    response = &ModifyClusterInstancesSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterInstancesSpec
// 修改实例规格，此接口只针对状态为running(运行中)的实例生效
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_SPECNOTCHANGE = "FailedOperation.SpecNotChange"
//  FAILEDOPERATION_SPECSTORAGELACK = "FailedOperation.SpecStorageLack"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterInstancesSpec(request *ModifyClusterInstancesSpecRequest) (response *ModifyClusterInstancesSpecResponse, err error) {
    if request == nil {
        request = NewModifyClusterInstancesSpecRequest()
    }
    
    response = NewModifyClusterInstancesSpecResponse()
    err = c.Send(request, response)
    return
}

// ModifyClusterInstancesSpec
// 修改实例规格，此接口只针对状态为running(运行中)的实例生效
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  FAILEDOPERATION_SPECNOTCHANGE = "FailedOperation.SpecNotChange"
//  FAILEDOPERATION_SPECSTORAGELACK = "FailedOperation.SpecStorageLack"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDSPEC = "InvalidParameterValue.InvalidSpec"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterInstancesSpecWithContext(ctx context.Context, request *ModifyClusterInstancesSpecRequest) (response *ModifyClusterInstancesSpecResponse, err error) {
    if request == nil {
        request = NewModifyClusterInstancesSpecRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyClusterInstancesSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
    request = &ModifyClusterNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "ModifyClusterName")
    
    
    return
}

func NewModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
    response = &ModifyClusterNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClusterName
// 修改集群名字
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

// ModifyClusterName
// 修改集群名字
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALINSTANCENAME = "InvalidParameterValue.IllegalInstanceName"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClusterNameWithContext(ctx context.Context, request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
    if request == nil {
        request = NewModifyClusterNameRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyClusterNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyClustersAutoRenewFlagRequest() (request *ModifyClustersAutoRenewFlagRequest) {
    request = &ModifyClustersAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "ModifyClustersAutoRenewFlag")
    
    
    return
}

func NewModifyClustersAutoRenewFlagResponse() (response *ModifyClustersAutoRenewFlagResponse) {
    response = &ModifyClustersAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyClustersAutoRenewFlag
// 修改集群自动续费，只对预付费集群生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClustersAutoRenewFlag(request *ModifyClustersAutoRenewFlagRequest) (response *ModifyClustersAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyClustersAutoRenewFlagRequest()
    }
    
    response = NewModifyClustersAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

// ModifyClustersAutoRenewFlag
// 修改集群自动续费，只对预付费集群生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyClustersAutoRenewFlagWithContext(ctx context.Context, request *ModifyClustersAutoRenewFlagRequest) (response *ModifyClustersAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyClustersAutoRenewFlagRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyClustersAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverClusterRequest() (request *RecoverClusterRequest) {
    request = &RecoverClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "RecoverCluster")
    
    
    return
}

func NewRecoverClusterResponse() (response *RecoverClusterResponse) {
    response = &RecoverClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecoverCluster
// 恢复集群，恢复集群的接入点网络，恢复后继续连接使用数据库。只有当集群状态处于isolated(已隔离)时才生效。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecoverCluster(request *RecoverClusterRequest) (response *RecoverClusterResponse, err error) {
    if request == nil {
        request = NewRecoverClusterRequest()
    }
    
    response = NewRecoverClusterResponse()
    err = c.Send(request, response)
    return
}

// RecoverCluster
// 恢复集群，恢复集群的接入点网络，恢复后继续连接使用数据库。只有当集群状态处于isolated(已隔离)时才生效。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecoverClusterWithContext(ctx context.Context, request *RecoverClusterRequest) (response *RecoverClusterResponse, err error) {
    if request == nil {
        request = NewRecoverClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewRecoverClusterResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverClusterInstancesRequest() (request *RecoverClusterInstancesRequest) {
    request = &RecoverClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "RecoverClusterInstances")
    
    
    return
}

func NewRecoverClusterInstancesResponse() (response *RecoverClusterInstancesResponse) {
    response = &RecoverClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecoverClusterInstances
// 恢复实例。此接口的使用场景包括：
//
//  - 读写实例状态为running(运行中)时，批量恢复状态为isolated(已隔离)的只读实例
//
//  - 读写实例状态为isolated(已隔离)时，恢复读写实例
//
//  - 读写实例状态为isolated(已隔离)时，批量恢复读写实例以及状态为isolated(已隔离)的只读实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecoverClusterInstances(request *RecoverClusterInstancesRequest) (response *RecoverClusterInstancesResponse, err error) {
    if request == nil {
        request = NewRecoverClusterInstancesRequest()
    }
    
    response = NewRecoverClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// RecoverClusterInstances
// 恢复实例。此接口的使用场景包括：
//
//  - 读写实例状态为running(运行中)时，批量恢复状态为isolated(已隔离)的只读实例
//
//  - 读写实例状态为isolated(已隔离)时，恢复读写实例
//
//  - 读写实例状态为isolated(已隔离)时，批量恢复读写实例以及状态为isolated(已隔离)的只读实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_STATUSERROR = "FailedOperation.StatusError"
//  FAILEDOPERATION_TASKCONFLICT = "FailedOperation.TaskConflict"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RecoverClusterInstancesWithContext(ctx context.Context, request *RecoverClusterInstancesRequest) (response *RecoverClusterInstancesResponse, err error) {
    if request == nil {
        request = NewRecoverClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRecoverClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewRenewClusterRequest() (request *RenewClusterRequest) {
    request = &RenewClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "RenewCluster")
    
    
    return
}

func NewRenewClusterResponse() (response *RenewClusterResponse) {
    response = &RenewClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewCluster
// 续费集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) RenewCluster(request *RenewClusterRequest) (response *RenewClusterResponse, err error) {
    if request == nil {
        request = NewRenewClusterRequest()
    }
    
    response = NewRenewClusterResponse()
    err = c.Send(request, response)
    return
}

// RenewCluster
// 续费集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) RenewClusterWithContext(ctx context.Context, request *RenewClusterRequest) (response *RenewClusterResponse, err error) {
    if request == nil {
        request = NewRenewClusterRequest()
    }
    request.SetContext(ctx)
    
    response = NewRenewClusterResponse()
    err = c.Send(request, response)
    return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
    request = &ResetAccountPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "ResetAccountPassword")
    
    
    return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
    response = &ResetAccountPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetAccountPassword
// 重置数据库账号密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

// ResetAccountPassword
// 重置数据库账号密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTFOUND = "InvalidParameterValue.AccountNotFound"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_ILLEGALPASSWORD = "InvalidParameterValue.IllegalPassword"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
    if request == nil {
        request = NewResetAccountPasswordRequest()
    }
    request.SetContext(ctx)
    
    response = NewResetAccountPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartClusterInstancesRequest() (request *RestartClusterInstancesRequest) {
    request = &RestartClusterInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "RestartClusterInstances")
    
    
    return
}

func NewRestartClusterInstancesResponse() (response *RestartClusterInstancesResponse) {
    response = &RestartClusterInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartClusterInstances
// 重启实例，此接口只针对状态为running(运行中)的实例生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartClusterInstances(request *RestartClusterInstancesRequest) (response *RestartClusterInstancesResponse, err error) {
    if request == nil {
        request = NewRestartClusterInstancesRequest()
    }
    
    response = NewRestartClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

// RestartClusterInstances
// 重启实例，此接口只针对状态为running(运行中)的实例生效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_FLOWERROR = "FailedOperation.FlowError"
//  FAILEDOPERATION_INTERNALSERVICEACCESSERROR = "FailedOperation.InternalServiceAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_PARAMETEROUTRANGEERROR = "InvalidParameterValue.ParameterOutRangeError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) RestartClusterInstancesWithContext(ctx context.Context, request *RestartClusterInstancesRequest) (response *RestartClusterInstancesResponse, err error) {
    if request == nil {
        request = NewRestartClusterInstancesRequest()
    }
    request.SetContext(ctx)
    
    response = NewRestartClusterInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewTransformClusterPayModeRequest() (request *TransformClusterPayModeRequest) {
    request = &TransformClusterPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tdcpg", APIVersion, "TransformClusterPayMode")
    
    
    return
}

func NewTransformClusterPayModeResponse() (response *TransformClusterPayModeResponse) {
    response = &TransformClusterPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransformClusterPayMode
// 转换集群付费模式，目前只支持从 后付费 转换成 与预付费。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TransformClusterPayMode(request *TransformClusterPayModeRequest) (response *TransformClusterPayModeResponse, err error) {
    if request == nil {
        request = NewTransformClusterPayModeRequest()
    }
    
    response = NewTransformClusterPayModeResponse()
    err = c.Send(request, response)
    return
}

// TransformClusterPayMode
// 转换集群付费模式，目前只支持从 后付费 转换成 与预付费。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATABASEACCESSERROR = "FailedOperation.DatabaseAccessError"
//  FAILEDOPERATION_PAYMODEINVALID = "FailedOperation.PayModeInvalid"
//  FAILEDOPERATION_TRADEACCESSERROR = "FailedOperation.TradeAccessError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLUSTERNOTFOUND = "InvalidParameterValue.ClusterNotFound"
//  INVALIDPARAMETERVALUE_INSTANCENOTFOUND = "InvalidParameterValue.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_REGIONZONEUNAVAILABLE = "InvalidParameterValue.RegionZoneUnavailable"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) TransformClusterPayModeWithContext(ctx context.Context, request *TransformClusterPayModeRequest) (response *TransformClusterPayModeResponse, err error) {
    if request == nil {
        request = NewTransformClusterPayModeRequest()
    }
    request.SetContext(ctx)
    
    response = NewTransformClusterPayModeResponse()
    err = c.Send(request, response)
    return
}
