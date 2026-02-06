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

package v20211122

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-11-22"

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


func NewCancelIsolateDBInstancesRequest() (request *CancelIsolateDBInstancesRequest) {
    request = &CancelIsolateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "CancelIsolateDBInstances")
    
    
    return
}

func NewCancelIsolateDBInstancesResponse() (response *CancelIsolateDBInstancesResponse) {
    response = &CancelIsolateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelIsolateDBInstances
// 本接口（CancelIsolateDBInstances）提供批量解除隔离实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
func (c *Client) CancelIsolateDBInstances(request *CancelIsolateDBInstancesRequest) (response *CancelIsolateDBInstancesResponse, err error) {
    return c.CancelIsolateDBInstancesWithContext(context.Background(), request)
}

// CancelIsolateDBInstances
// 本接口（CancelIsolateDBInstances）提供批量解除隔离实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
func (c *Client) CancelIsolateDBInstancesWithContext(ctx context.Context, request *CancelIsolateDBInstancesRequest) (response *CancelIsolateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCancelIsolateDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "CancelIsolateDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelIsolateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelIsolateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBSBackupRequest() (request *CreateDBSBackupRequest) {
    request = &CreateDBSBackupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "CreateDBSBackup")
    
    
    return
}

func NewCreateDBSBackupResponse() (response *CreateDBSBackupResponse) {
    response = &CreateDBSBackupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBSBackup
// 创建实例备份集
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED_CREATEBACKUPTASKTHRESHOLDERR = "OperationDenied.CreateBackupTaskThresholdErr"
//  OPERATIONDENIED_MANUALBACKUPQUOTAPERDAYEXCEEDEDERR = "OperationDenied.ManualBackupQuotaPerDayExceededErr"
//  OPERATIONDENIED_MANUALBACKUPSETQUOTAEXCEEDEDERR = "OperationDenied.ManualBackupSetQuotaExceededErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) CreateDBSBackup(request *CreateDBSBackupRequest) (response *CreateDBSBackupResponse, err error) {
    return c.CreateDBSBackupWithContext(context.Background(), request)
}

// CreateDBSBackup
// 创建实例备份集
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED_CREATEBACKUPTASKTHRESHOLDERR = "OperationDenied.CreateBackupTaskThresholdErr"
//  OPERATIONDENIED_MANUALBACKUPQUOTAPERDAYEXCEEDEDERR = "OperationDenied.ManualBackupQuotaPerDayExceededErr"
//  OPERATIONDENIED_MANUALBACKUPSETQUOTAEXCEEDEDERR = "OperationDenied.ManualBackupSetQuotaExceededErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) CreateDBSBackupWithContext(ctx context.Context, request *CreateDBSBackupRequest) (response *CreateDBSBackupResponse, err error) {
    if request == nil {
        request = NewCreateDBSBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "CreateDBSBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBSBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBSBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBSBackupSetsRequest() (request *DeleteDBSBackupSetsRequest) {
    request = &DeleteDBSBackupSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DeleteDBSBackupSets")
    
    
    return
}

func NewDeleteDBSBackupSetsResponse() (response *DeleteDBSBackupSetsResponse) {
    response = &DeleteDBSBackupSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBSBackupSets
// 删除实例备份集
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DeleteDBSBackupSets(request *DeleteDBSBackupSetsRequest) (response *DeleteDBSBackupSetsResponse, err error) {
    return c.DeleteDBSBackupSetsWithContext(context.Background(), request)
}

// DeleteDBSBackupSets
// 删除实例备份集
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DeleteDBSBackupSetsWithContext(ctx context.Context, request *DeleteDBSBackupSetsRequest) (response *DeleteDBSBackupSetsResponse, err error) {
    if request == nil {
        request = NewDeleteDBSBackupSetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DeleteDBSBackupSets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBSBackupSets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBSBackupSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingEnableRequest() (request *DescribeBillingEnableRequest) {
    request = &DescribeBillingEnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeBillingEnable")
    
    
    return
}

func NewDescribeBillingEnableResponse() (response *DescribeBillingEnableResponse) {
    response = &DescribeBillingEnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingEnable
// 已无地方调用
//
// 
//
// 本接口（DescribeBillingEnable）用于查询计费是否开启
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeBillingEnable(request *DescribeBillingEnableRequest) (response *DescribeBillingEnableResponse, err error) {
    return c.DescribeBillingEnableWithContext(context.Background(), request)
}

// DescribeBillingEnable
// 已无地方调用
//
// 
//
// 本接口（DescribeBillingEnable）用于查询计费是否开启
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeBillingEnableWithContext(ctx context.Context, request *DescribeBillingEnableRequest) (response *DescribeBillingEnableResponse, err error) {
    if request == nil {
        request = NewDescribeBillingEnableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeBillingEnable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingEnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingEnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
    request = &DescribeDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBParameters")
    
    
    return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
    response = &DescribeDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBParameters
// 本接口（DescribeDBParameters）用于获取实例的当前参数设置。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_OSSGETVARIABLESERROR = "FailedOperation.OssGetVariablesError"
//  INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    return c.DescribeDBParametersWithContext(context.Background(), request)
}

// DescribeDBParameters
// 本接口（DescribeDBParameters）用于获取实例的当前参数设置。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_OSSGETVARIABLESERROR = "FailedOperation.OssGetVariablesError"
//  INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"
func (c *Client) DescribeDBParametersWithContext(ctx context.Context, request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
    if request == nil {
        request = NewDescribeDBParametersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBParameters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSAvailableRecoveryTimeRequest() (request *DescribeDBSAvailableRecoveryTimeRequest) {
    request = &DescribeDBSAvailableRecoveryTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSAvailableRecoveryTime")
    
    
    return
}

func NewDescribeDBSAvailableRecoveryTimeResponse() (response *DescribeDBSAvailableRecoveryTimeResponse) {
    response = &DescribeDBSAvailableRecoveryTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSAvailableRecoveryTime
// 可恢复时间查询
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSAvailableRecoveryTime(request *DescribeDBSAvailableRecoveryTimeRequest) (response *DescribeDBSAvailableRecoveryTimeResponse, err error) {
    return c.DescribeDBSAvailableRecoveryTimeWithContext(context.Background(), request)
}

// DescribeDBSAvailableRecoveryTime
// 可恢复时间查询
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSAvailableRecoveryTimeWithContext(ctx context.Context, request *DescribeDBSAvailableRecoveryTimeRequest) (response *DescribeDBSAvailableRecoveryTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBSAvailableRecoveryTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSAvailableRecoveryTime")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSAvailableRecoveryTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSAvailableRecoveryTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSCloneInstancesRequest() (request *DescribeDBSCloneInstancesRequest) {
    request = &DescribeDBSCloneInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSCloneInstances")
    
    
    return
}

func NewDescribeDBSCloneInstancesResponse() (response *DescribeDBSCloneInstancesResponse) {
    response = &DescribeDBSCloneInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSCloneInstances
// 查询实例克隆列表
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSCloneInstances(request *DescribeDBSCloneInstancesRequest) (response *DescribeDBSCloneInstancesResponse, err error) {
    return c.DescribeDBSCloneInstancesWithContext(context.Background(), request)
}

// DescribeDBSCloneInstances
// 查询实例克隆列表
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSCloneInstancesWithContext(ctx context.Context, request *DescribeDBSCloneInstancesRequest) (response *DescribeDBSCloneInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBSCloneInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSCloneInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSCloneInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSCloneInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSecurityGroups")
    
    
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
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
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
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
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
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseObjectsRequest() (request *DescribeDatabaseObjectsRequest) {
    request = &DescribeDatabaseObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDatabaseObjects")
    
    
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
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeDatabaseObjects(request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    return c.DescribeDatabaseObjectsWithContext(context.Background(), request)
}

// DescribeDatabaseObjects
// 本接口（DescribeDatabaseObjects）用于查询云数据库实例的数据库中的对象列表，包含表、存储过程、视图和函数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeDatabaseObjectsWithContext(ctx context.Context, request *DescribeDatabaseObjectsRequest) (response *DescribeDatabaseObjectsResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseObjectsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDatabaseObjects")
    
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
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDatabaseTable")
    
    
    return
}

func NewDescribeDatabaseTableResponse() (response *DescribeDatabaseTableResponse) {
    response = &DescribeDatabaseTableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDatabaseTable
// 冗余接口，无人调用
//
// 
//
// 本接口（DescribeDatabaseTable）用于查询云数据库实例的表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeDatabaseTable(request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    return c.DescribeDatabaseTableWithContext(context.Background(), request)
}

// DescribeDatabaseTable
// 冗余接口，无人调用
//
// 
//
// 本接口（DescribeDatabaseTable）用于查询云数据库实例的表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeDatabaseTableWithContext(ctx context.Context, request *DescribeDatabaseTableRequest) (response *DescribeDatabaseTableResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseTableRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDatabaseTable")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabaseTable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabaseTableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowRequest() (request *DescribeFlowRequest) {
    request = &DescribeFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeFlow")
    
    
    return
}

func NewDescribeFlowResponse() (response *DescribeFlowResponse) {
    response = &DescribeFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlow
// 本接口（DescribeFlow）用于查询异步任务流程状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    return c.DescribeFlowWithContext(context.Background(), request)
}

// DescribeFlow
// 本接口（DescribeFlow）用于查询异步任务流程状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeFlowWithContext(ctx context.Context, request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    if request == nil {
        request = NewDescribeFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyInstancesRequest() (request *DestroyInstancesRequest) {
    request = &DestroyInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DestroyInstances")
    
    
    return
}

func NewDestroyInstancesResponse() (response *DestroyInstancesResponse) {
    response = &DestroyInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyInstances
// 本接口（DestroyInstances）提供批量销毁实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) DestroyInstances(request *DestroyInstancesRequest) (response *DestroyInstancesResponse, err error) {
    return c.DestroyInstancesWithContext(context.Background(), request)
}

// DestroyInstances
// 本接口（DestroyInstances）提供批量销毁实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) DestroyInstancesWithContext(ctx context.Context, request *DestroyInstancesRequest) (response *DestroyInstancesResponse, err error) {
    if request == nil {
        request = NewDestroyInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DestroyInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateDBInstanceRequest() (request *IsolateDBInstanceRequest) {
    request = &IsolateDBInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "IsolateDBInstance")
    
    
    return
}

func NewIsolateDBInstanceResponse() (response *IsolateDBInstanceResponse) {
    response = &IsolateDBInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateDBInstance
// 本接口（IsolateDBInstance）提供批量隔离实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// 本接口（IsolateDBInstance）提供批量隔离实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "IsolateDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyAutoRenewFlag")
    
    
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoRenewFlag
// 本接口（ModifyAutoRenewFlag）用于修改自动续费标志
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// 本接口（ModifyAutoRenewFlag）用于修改自动续费标志
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_ISOLATEINSTANCEERROR = "FailedOperation.IsolateInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyAutoRenewFlag")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBinlogStatusRequest() (request *ModifyBinlogStatusRequest) {
    request = &ModifyBinlogStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyBinlogStatus")
    
    
    return
}

func NewModifyBinlogStatusResponse() (response *ModifyBinlogStatusResponse) {
    response = &ModifyBinlogStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBinlogStatus
// 接口功能已被 ModifyInstanceCdc 完全覆盖
//
// 
//
// 修改binlog状态
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
func (c *Client) ModifyBinlogStatus(request *ModifyBinlogStatusRequest) (response *ModifyBinlogStatusResponse, err error) {
    return c.ModifyBinlogStatusWithContext(context.Background(), request)
}

// ModifyBinlogStatus
// 接口功能已被 ModifyInstanceCdc 完全覆盖
//
// 
//
// 修改binlog状态
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
func (c *Client) ModifyBinlogStatusWithContext(ctx context.Context, request *ModifyBinlogStatusRequest) (response *ModifyBinlogStatusResponse, err error) {
    if request == nil {
        request = NewModifyBinlogStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyBinlogStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBinlogStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBinlogStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
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
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceSecurityGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
    request = &ModifyDBParametersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyDBParameters")
    
    
    return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
    response = &ModifyDBParametersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBParameters
// 本接口（ModifyDBParameters）用于修改实例参数。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  FAILEDOPERATION_OSSMODIFYVARIABLESERROR = "FailedOperation.OssModifyVariablesError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    return c.ModifyDBParametersWithContext(context.Background(), request)
}

// ModifyDBParameters
// 本接口（ModifyDBParameters）用于修改实例参数。
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  FAILEDOPERATION_OSSMODIFYVARIABLESERROR = "FailedOperation.OssModifyVariablesError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyDBParametersWithContext(ctx context.Context, request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
    if request == nil {
        request = NewModifyDBParametersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyDBParameters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBParameters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBParametersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBSBackupPolicyRequest() (request *ModifyDBSBackupPolicyRequest) {
    request = &ModifyDBSBackupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyDBSBackupPolicy")
    
    
    return
}

func NewModifyDBSBackupPolicyResponse() (response *ModifyDBSBackupPolicyResponse) {
    response = &ModifyDBSBackupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBSBackupPolicy
// 修改实例备份策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYBACKUPPOLICYERR = "FailedOperation.ModifyBackupPolicyErr"
//  FAILEDOPERATION_MODIFYPOLICYERR = "FailedOperation.ModifyPolicyErr"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESCRIBEDBOBJECTSERROR = "InternalError.DescribeDBObjectsError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  INVALIDPARAMETER_ILLEGALBACKUPPOLICYPARAMSERR = "InvalidParameter.IllegalBackupPolicyParamsErr"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifyDBSBackupPolicy(request *ModifyDBSBackupPolicyRequest) (response *ModifyDBSBackupPolicyResponse, err error) {
    return c.ModifyDBSBackupPolicyWithContext(context.Background(), request)
}

// ModifyDBSBackupPolicy
// 修改实例备份策略
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODIFYBACKUPPOLICYERR = "FailedOperation.ModifyBackupPolicyErr"
//  FAILEDOPERATION_MODIFYPOLICYERR = "FailedOperation.ModifyPolicyErr"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESCRIBEDBOBJECTSERROR = "InternalError.DescribeDBObjectsError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  INVALIDPARAMETER_ILLEGALBACKUPPOLICYPARAMSERR = "InvalidParameter.IllegalBackupPolicyParamsErr"
//  INVALIDPARAMETER_INPUTILLEGAL = "InvalidParameter.InputIllegal"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifyDBSBackupPolicyWithContext(ctx context.Context, request *ModifyDBSBackupPolicyRequest) (response *ModifyDBSBackupPolicyResponse, err error) {
    if request == nil {
        request = NewModifyDBSBackupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyDBSBackupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBSBackupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBSBackupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBSBackupSetCommentRequest() (request *ModifyDBSBackupSetCommentRequest) {
    request = &ModifyDBSBackupSetCommentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyDBSBackupSetComment")
    
    
    return
}

func NewModifyDBSBackupSetCommentResponse() (response *ModifyDBSBackupSetCommentResponse) {
    response = &ModifyDBSBackupSetCommentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBSBackupSetComment
// 修改备份集备注
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifyDBSBackupSetComment(request *ModifyDBSBackupSetCommentRequest) (response *ModifyDBSBackupSetCommentResponse, err error) {
    return c.ModifyDBSBackupSetCommentWithContext(context.Background(), request)
}

// ModifyDBSBackupSetComment
// 修改备份集备注
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifyDBSBackupSetCommentWithContext(ctx context.Context, request *ModifyDBSBackupSetCommentRequest) (response *ModifyDBSBackupSetCommentResponse, err error) {
    if request == nil {
        request = NewModifyDBSBackupSetCommentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyDBSBackupSetComment")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBSBackupSetComment require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBSBackupSetCommentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
    request = &ModifyInstanceNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyInstanceName")
    
    
    return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
    response = &ModifyInstanceNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceName
// 本接口（ModifyInstanceName）提供修改实例名称功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBUPDATEINSTANCEERROR = "FailedOperation.DBUpdateInstanceError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    return c.ModifyInstanceNameWithContext(context.Background(), request)
}

// ModifyInstanceName
// 本接口（ModifyInstanceName）提供修改实例名称功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBUPDATEINSTANCEERROR = "FailedOperation.DBUpdateInstanceError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNameRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyInstanceName")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNameResponse()
    err = c.Send(request, response)
    return
}
