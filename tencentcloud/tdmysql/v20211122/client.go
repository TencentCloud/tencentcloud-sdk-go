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

func NewCreateCloneInstanceRequest() (request *CreateCloneInstanceRequest) {
    request = &CreateCloneInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "CreateCloneInstance")
    
    
    return
}

func NewCreateCloneInstanceResponse() (response *CreateCloneInstanceResponse) {
    response = &CreateCloneInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloneInstance
// 本接口（CreateCloneInstance）提供创建克隆实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION_ADDTAGDRYRUNERROR = "DryRunOperation.AddTagDryrunError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBCREATEINSTANCEERROR = "FailedOperation.DBCreateInstanceError"
//  FAILEDOPERATION_DESCRIBEFULLBACKUPLISTERROR = "FailedOperation.DescribeFullBackupListError"
//  INVALIDPARAMETER_CREATEINSTANCEPARAMERROR = "InvalidParameter.CreateInstanceParamError"
//  INVALIDPARAMETERVALUE_CHECKCLONEINSTANCEDISKERROR = "InvalidParameterValue.CheckCloneInstanceDiskError"
//  INVALIDPARAMETERVALUE_CHECKDISKERROR = "InvalidParameterValue.CheckDiskError"
//  INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
//  INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"
//  INVALIDPARAMETERVALUE_NODEREPLICASINVALIDERROR = "InvalidParameterValue.NodeReplicasInvalidError"
//  LIMITEXCEEDED_ADDTAGCOUNTERROR = "LimitExceeded.AddTagCountError"
//  LIMITEXCEEDED_OUTOFINSTANCECNTLIMITERROR = "LimitExceeded.OutOfInstanceCntLimitError"
//  LIMITEXCEEDED_OUTOFNODEREPLICASLIMITERROR = "LimitExceeded.OutOfNodeReplicasLimitError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
func (c *Client) CreateCloneInstance(request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    return c.CreateCloneInstanceWithContext(context.Background(), request)
}

// CreateCloneInstance
// 本接口（CreateCloneInstance）提供创建克隆实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION_ADDTAGDRYRUNERROR = "DryRunOperation.AddTagDryrunError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBCREATEINSTANCEERROR = "FailedOperation.DBCreateInstanceError"
//  FAILEDOPERATION_DESCRIBEFULLBACKUPLISTERROR = "FailedOperation.DescribeFullBackupListError"
//  INVALIDPARAMETER_CREATEINSTANCEPARAMERROR = "InvalidParameter.CreateInstanceParamError"
//  INVALIDPARAMETERVALUE_CHECKCLONEINSTANCEDISKERROR = "InvalidParameterValue.CheckCloneInstanceDiskError"
//  INVALIDPARAMETERVALUE_CHECKDISKERROR = "InvalidParameterValue.CheckDiskError"
//  INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
//  INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"
//  INVALIDPARAMETERVALUE_NODEREPLICASINVALIDERROR = "InvalidParameterValue.NodeReplicasInvalidError"
//  LIMITEXCEEDED_ADDTAGCOUNTERROR = "LimitExceeded.AddTagCountError"
//  LIMITEXCEEDED_OUTOFINSTANCECNTLIMITERROR = "LimitExceeded.OutOfInstanceCntLimitError"
//  LIMITEXCEEDED_OUTOFNODEREPLICASLIMITERROR = "LimitExceeded.OutOfNodeReplicasLimitError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
func (c *Client) CreateCloneInstanceWithContext(ctx context.Context, request *CreateCloneInstanceRequest) (response *CreateCloneInstanceResponse, err error) {
    if request == nil {
        request = NewCreateCloneInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "CreateCloneInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloneInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloneInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBInstancesRequest() (request *CreateDBInstancesRequest) {
    request = &CreateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "CreateDBInstances")
    
    
    return
}

func NewCreateDBInstancesResponse() (response *CreateDBInstancesResponse) {
    response = &CreateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBInstances
// 本接口（CreateDBInstances）提供批量创建实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION_ADDTAGDRYRUNERROR = "DryRunOperation.AddTagDryrunError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBCREATEINSTANCEERROR = "FailedOperation.DBCreateInstanceError"
//  INVALIDPARAMETERVALUE_CHECKDISKERROR = "InvalidParameterValue.CheckDiskError"
//  INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
//  INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_NODEREPLICASINVALIDERROR = "InvalidParameterValue.NodeReplicasInvalidError"
//  LIMITEXCEEDED_ADDTAGCOUNTERROR = "LimitExceeded.AddTagCountError"
//  LIMITEXCEEDED_OUTOFINSTANCECNTLIMITERROR = "LimitExceeded.OutOfInstanceCntLimitError"
//  LIMITEXCEEDED_OUTOFNODEREPLICASLIMITERROR = "LimitExceeded.OutOfNodeReplicasLimitError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
func (c *Client) CreateDBInstances(request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    return c.CreateDBInstancesWithContext(context.Background(), request)
}

// CreateDBInstances
// 本接口（CreateDBInstances）提供批量创建实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION_ADDTAGDRYRUNERROR = "DryRunOperation.AddTagDryrunError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBCREATEINSTANCEERROR = "FailedOperation.DBCreateInstanceError"
//  INVALIDPARAMETERVALUE_CHECKDISKERROR = "InvalidParameterValue.CheckDiskError"
//  INVALIDPARAMETERVALUE_CHECKINSTANCEVERSIONERROR = "InvalidParameterValue.CheckInstanceVersionError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
//  INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  INVALIDPARAMETERVALUE_NODEREPLICASINVALIDERROR = "InvalidParameterValue.NodeReplicasInvalidError"
//  LIMITEXCEEDED_ADDTAGCOUNTERROR = "LimitExceeded.AddTagCountError"
//  LIMITEXCEEDED_OUTOFINSTANCECNTLIMITERROR = "LimitExceeded.OutOfInstanceCntLimitError"
//  LIMITEXCEEDED_OUTOFNODEREPLICASLIMITERROR = "LimitExceeded.OutOfNodeReplicasLimitError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
func (c *Client) CreateDBInstancesWithContext(ctx context.Context, request *CreateDBInstancesRequest) (response *CreateDBInstancesResponse, err error) {
    if request == nil {
        request = NewCreateDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "CreateDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstancesResponse()
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
// 创建实例手工备份  CreateDBSBackup
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
// 创建实例手工备份  CreateDBSBackup
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

func NewCreateUsersRequest() (request *CreateUsersRequest) {
    request = &CreateUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "CreateUsers")
    
    
    return
}

func NewCreateUsersResponse() (response *CreateUsersResponse) {
    response = &CreateUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUsers
// 本接口（CreateUsers）用于批量创建用户
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED_CREATEBACKUPTASKTHRESHOLDERR = "OperationDenied.CreateBackupTaskThresholdErr"
//  OPERATIONDENIED_MANUALBACKUPQUOTAPERDAYEXCEEDEDERR = "OperationDenied.ManualBackupQuotaPerDayExceededErr"
//  OPERATIONDENIED_MANUALBACKUPSETQUOTAEXCEEDEDERR = "OperationDenied.ManualBackupSetQuotaExceededErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) CreateUsers(request *CreateUsersRequest) (response *CreateUsersResponse, err error) {
    return c.CreateUsersWithContext(context.Background(), request)
}

// CreateUsers
// 本接口（CreateUsers）用于批量创建用户
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED_CREATEBACKUPTASKTHRESHOLDERR = "OperationDenied.CreateBackupTaskThresholdErr"
//  OPERATIONDENIED_MANUALBACKUPQUOTAPERDAYEXCEEDEDERR = "OperationDenied.ManualBackupQuotaPerDayExceededErr"
//  OPERATIONDENIED_MANUALBACKUPSETQUOTAEXCEEDEDERR = "OperationDenied.ManualBackupSetQuotaExceededErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) CreateUsersWithContext(ctx context.Context, request *CreateUsersRequest) (response *CreateUsersResponse, err error) {
    if request == nil {
        request = NewCreateUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "CreateUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUsersResponse()
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
// 删除实例手工备份 DeleteDBSBackupSets
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
// 删除实例手工备份 DeleteDBSBackupSets
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

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
    request = &DeleteUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DeleteUsers")
    
    
    return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
    response = &DeleteUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteUsers
// 本接口（DeleteUsers）用于批量删除用户
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    return c.DeleteUsersWithContext(context.Background(), request)
}

// DeleteUsers
// 本接口（DeleteUsers）用于批量删除用户
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DELETERUNNINGBACKUPTASKERR = "OperationDenied.DeleteRunningBackupTaskErr"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
    if request == nil {
        request = NewDeleteUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DeleteUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceDetailRequest() (request *DescribeDBInstanceDetailRequest) {
    request = &DescribeDBInstanceDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBInstanceDetail")
    
    
    return
}

func NewDescribeDBInstanceDetailResponse() (response *DescribeDBInstanceDetailResponse) {
    response = &DescribeDBInstanceDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceDetail
// 本接口（DescribeDBInstanceDetail）提供查询实例详情功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDBInstanceDetail(request *DescribeDBInstanceDetailRequest) (response *DescribeDBInstanceDetailResponse, err error) {
    return c.DescribeDBInstanceDetailWithContext(context.Background(), request)
}

// DescribeDBInstanceDetail
// 本接口（DescribeDBInstanceDetail）提供查询实例详情功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDBInstanceDetailWithContext(ctx context.Context, request *DescribeDBInstanceDetailRequest) (response *DescribeDBInstanceDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBInstanceDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
    request = &DescribeDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBInstances")
    
    
    return
}

func NewDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
    response = &DescribeDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstances
// 本接口（DescribeDBInstances）提供查询实例列表功能
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBQUERYTAGERROR = "FailedOperation.DBQueryTagError"
//  INVALIDPARAMETERVALUE_INSTANCEFILTERKEYERROR = "InvalidParameterValue.InstanceFilterKeyError"
//  LIMITEXCEEDED_OUTOFINSTANCECOUNTLIMITERROR = "LimitExceeded.OutOfInstanceCountLimitError"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口（DescribeDBInstances）提供查询实例列表功能
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBQUERYTAGERROR = "FailedOperation.DBQueryTagError"
//  INVALIDPARAMETERVALUE_INSTANCEFILTERKEYERROR = "InvalidParameterValue.InstanceFilterKeyError"
//  LIMITEXCEEDED_OUTOFINSTANCECOUNTLIMITERROR = "LimitExceeded.OutOfInstanceCountLimitError"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstancesResponse()
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

func NewDescribeDBSArchiveLogsRequest() (request *DescribeDBSArchiveLogsRequest) {
    request = &DescribeDBSArchiveLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSArchiveLogs")
    
    
    return
}

func NewDescribeDBSArchiveLogsResponse() (response *DescribeDBSArchiveLogsResponse) {
    response = &DescribeDBSArchiveLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSArchiveLogs
// 查询实例归档日志列表 DescribeDBSArchiveLogs
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESCRIBEDBOBJECTSERROR = "InternalError.DescribeDBObjectsError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSArchiveLogs(request *DescribeDBSArchiveLogsRequest) (response *DescribeDBSArchiveLogsResponse, err error) {
    return c.DescribeDBSArchiveLogsWithContext(context.Background(), request)
}

// DescribeDBSArchiveLogs
// 查询实例归档日志列表 DescribeDBSArchiveLogs
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESCRIBEDBOBJECTSERROR = "InternalError.DescribeDBObjectsError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSArchiveLogsWithContext(ctx context.Context, request *DescribeDBSArchiveLogsRequest) (response *DescribeDBSArchiveLogsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSArchiveLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSArchiveLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSArchiveLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSArchiveLogsResponse()
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
// 获取可恢复时间 DescribeDBSAvailableRecoveryTime
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSAvailableRecoveryTime(request *DescribeDBSAvailableRecoveryTimeRequest) (response *DescribeDBSAvailableRecoveryTimeResponse, err error) {
    return c.DescribeDBSAvailableRecoveryTimeWithContext(context.Background(), request)
}

// DescribeDBSAvailableRecoveryTime
// 获取可恢复时间 DescribeDBSAvailableRecoveryTime
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

func NewDescribeDBSBackupPolicyRequest() (request *DescribeDBSBackupPolicyRequest) {
    request = &DescribeDBSBackupPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSBackupPolicy")
    
    
    return
}

func NewDescribeDBSBackupPolicyResponse() (response *DescribeDBSBackupPolicyResponse) {
    response = &DescribeDBSBackupPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSBackupPolicy
// 查询实例备份策略 DescribeDBSBackupPolicy
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSBackupPolicy(request *DescribeDBSBackupPolicyRequest) (response *DescribeDBSBackupPolicyResponse, err error) {
    return c.DescribeDBSBackupPolicyWithContext(context.Background(), request)
}

// DescribeDBSBackupPolicy
// 查询实例备份策略 DescribeDBSBackupPolicy
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSBackupPolicyWithContext(ctx context.Context, request *DescribeDBSBackupPolicyRequest) (response *DescribeDBSBackupPolicyResponse, err error) {
    if request == nil {
        request = NewDescribeDBSBackupPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSBackupPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSBackupPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSBackupPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSBackupSetsRequest() (request *DescribeDBSBackupSetsRequest) {
    request = &DescribeDBSBackupSetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSBackupSets")
    
    
    return
}

func NewDescribeDBSBackupSetsResponse() (response *DescribeDBSBackupSetsResponse) {
    response = &DescribeDBSBackupSetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSBackupSets
// 查询实例备份集信息 DescribeDBSBackupSets
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSBackupSets(request *DescribeDBSBackupSetsRequest) (response *DescribeDBSBackupSetsResponse, err error) {
    return c.DescribeDBSBackupSetsWithContext(context.Background(), request)
}

// DescribeDBSBackupSets
// 查询实例备份集信息 DescribeDBSBackupSets
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSBackupSetsWithContext(ctx context.Context, request *DescribeDBSBackupSetsRequest) (response *DescribeDBSBackupSetsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSBackupSetsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSBackupSets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSBackupSets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSBackupSetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSBackupStatisticsRequest() (request *DescribeDBSBackupStatisticsRequest) {
    request = &DescribeDBSBackupStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSBackupStatistics")
    
    
    return
}

func NewDescribeDBSBackupStatisticsResponse() (response *DescribeDBSBackupStatisticsResponse) {
    response = &DescribeDBSBackupStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSBackupStatistics
// 查询实例备份空间概览 DescribeDBSBackupStatistics
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSBackupStatistics(request *DescribeDBSBackupStatisticsRequest) (response *DescribeDBSBackupStatisticsResponse, err error) {
    return c.DescribeDBSBackupStatisticsWithContext(context.Background(), request)
}

// DescribeDBSBackupStatistics
// 查询实例备份空间概览 DescribeDBSBackupStatistics
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSBackupStatisticsWithContext(ctx context.Context, request *DescribeDBSBackupStatisticsRequest) (response *DescribeDBSBackupStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSBackupStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSBackupStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSBackupStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSBackupStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSBackupStatisticsDetailRequest() (request *DescribeDBSBackupStatisticsDetailRequest) {
    request = &DescribeDBSBackupStatisticsDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDBSBackupStatisticsDetail")
    
    
    return
}

func NewDescribeDBSBackupStatisticsDetailResponse() (response *DescribeDBSBackupStatisticsDetailResponse) {
    response = &DescribeDBSBackupStatisticsDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSBackupStatisticsDetail
// 查询备份集统计详情 DescribeDBSBackupStatisticsDetail
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeDBSBackupStatisticsDetail(request *DescribeDBSBackupStatisticsDetailRequest) (response *DescribeDBSBackupStatisticsDetailResponse, err error) {
    return c.DescribeDBSBackupStatisticsDetailWithContext(context.Background(), request)
}

// DescribeDBSBackupStatisticsDetail
// 查询备份集统计详情 DescribeDBSBackupStatisticsDetail
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeDBSBackupStatisticsDetailWithContext(ctx context.Context, request *DescribeDBSBackupStatisticsDetailRequest) (response *DescribeDBSBackupStatisticsDetailResponse, err error) {
    if request == nil {
        request = NewDescribeDBSBackupStatisticsDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDBSBackupStatisticsDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSBackupStatisticsDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSBackupStatisticsDetailResponse()
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
// 查询实例克隆列表 DescribeDBSCloneInstances
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) DescribeDBSCloneInstances(request *DescribeDBSCloneInstancesRequest) (response *DescribeDBSCloneInstancesResponse, err error) {
    return c.DescribeDBSCloneInstancesWithContext(context.Background(), request)
}

// DescribeDBSCloneInstances
// 查询实例克隆列表 DescribeDBSCloneInstances
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

func NewDescribeDatabasesRequest() (request *DescribeDatabasesRequest) {
    request = &DescribeDatabasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeDatabases")
    
    
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
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeDatabases(request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    return c.DescribeDatabasesWithContext(context.Background(), request)
}

// DescribeDatabases
// 本接口（DescribeDatabases）用于查询云数据库实例的数据库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeDatabases")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatabases require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatabasesResponse()
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
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeFlow(request *DescribeFlowRequest) (response *DescribeFlowResponse, err error) {
    return c.DescribeFlowWithContext(context.Background(), request)
}

// DescribeFlow
// 本接口（DescribeFlow）用于查询异步任务流程状态
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
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

func NewDescribeInstanceSSLStatusRequest() (request *DescribeInstanceSSLStatusRequest) {
    request = &DescribeInstanceSSLStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeInstanceSSLStatus")
    
    
    return
}

func NewDescribeInstanceSSLStatusResponse() (response *DescribeInstanceSSLStatusResponse) {
    response = &DescribeInstanceSSLStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSSLStatus
// 本接口（DescribeInstanceSSLStatus）提供实例SSL状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeInstanceSSLStatus(request *DescribeInstanceSSLStatusRequest) (response *DescribeInstanceSSLStatusResponse, err error) {
    return c.DescribeInstanceSSLStatusWithContext(context.Background(), request)
}

// DescribeInstanceSSLStatus
// 本接口（DescribeInstanceSSLStatus）提供实例SSL状态查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
func (c *Client) DescribeInstanceSSLStatusWithContext(ctx context.Context, request *DescribeInstanceSSLStatusRequest) (response *DescribeInstanceSSLStatusResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSSLStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeInstanceSSLStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSSLStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSSLStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMaintenanceWindowRequest() (request *DescribeMaintenanceWindowRequest) {
    request = &DescribeMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeMaintenanceWindow")
    
    
    return
}

func NewDescribeMaintenanceWindowResponse() (response *DescribeMaintenanceWindowResponse) {
    response = &DescribeMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMaintenanceWindow
// 查询维护时间窗口配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) DescribeMaintenanceWindow(request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    return c.DescribeMaintenanceWindowWithContext(context.Background(), request)
}

// DescribeMaintenanceWindow
// 查询维护时间窗口配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) DescribeMaintenanceWindowWithContext(ctx context.Context, request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceWindowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeMaintenanceWindow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSaleInfoRequest() (request *DescribeSaleInfoRequest) {
    request = &DescribeSaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeSaleInfo")
    
    
    return
}

func NewDescribeSaleInfoResponse() (response *DescribeSaleInfoResponse) {
    response = &DescribeSaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSaleInfo
// 本接口（DescribeSaleInfo）提供查询可用售卖地域功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"
func (c *Client) DescribeSaleInfo(request *DescribeSaleInfoRequest) (response *DescribeSaleInfoResponse, err error) {
    return c.DescribeSaleInfoWithContext(context.Background(), request)
}

// DescribeSaleInfo
// 本接口（DescribeSaleInfo）提供查询可用售卖地域功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"
func (c *Client) DescribeSaleInfoWithContext(ctx context.Context, request *DescribeSaleInfoRequest) (response *DescribeSaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeSaleInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeSaleInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSaleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeSlowLogs")
    
    
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogs
// 本接口提供查询慢日志功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    return c.DescribeSlowLogsWithContext(context.Background(), request)
}

// DescribeSlowLogs
// 本接口提供查询慢日志功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"
func (c *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSpecsRequest() (request *DescribeSpecsRequest) {
    request = &DescribeSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeSpecs")
    
    
    return
}

func NewDescribeSpecsResponse() (response *DescribeSpecsResponse) {
    response = &DescribeSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSpecs
// 本接口（DescribeSpecs）提供查询售卖规格功能
//
// 可能返回的错误码:
//  AUTHFAILURE_UINWHITELISTCHECKERROR = "AuthFailure.UinWhiteListCheckError"
//  FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"
//  UNSUPPORTEDOPERATION_GETMAXNODENUMERROR = "UnsupportedOperation.GetMaxNodeNumError"
func (c *Client) DescribeSpecs(request *DescribeSpecsRequest) (response *DescribeSpecsResponse, err error) {
    return c.DescribeSpecsWithContext(context.Background(), request)
}

// DescribeSpecs
// 本接口（DescribeSpecs）提供查询售卖规格功能
//
// 可能返回的错误码:
//  AUTHFAILURE_UINWHITELISTCHECKERROR = "AuthFailure.UinWhiteListCheckError"
//  FAILEDOPERATION_DBQUERYZONEERROR = "FailedOperation.DBQueryZoneError"
//  UNSUPPORTEDOPERATION_GETMAXNODENUMERROR = "UnsupportedOperation.GetMaxNodeNumError"
func (c *Client) DescribeSpecsWithContext(ctx context.Context, request *DescribeSpecsRequest) (response *DescribeSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeSpecsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeSpecs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserPrivilegesRequest() (request *DescribeUserPrivilegesRequest) {
    request = &DescribeUserPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeUserPrivileges")
    
    
    return
}

func NewDescribeUserPrivilegesResponse() (response *DescribeUserPrivilegesResponse) {
    response = &DescribeUserPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserPrivileges
// 本接口（DescribeUserPrivileges）提供查询用户的权限功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYPRIVILEGESERROR = "FailedOperation.DBQueryPrivilegesError"
//  INVALIDPARAMETERVALUE_CHECKHOSTERROR = "InvalidParameterValue.CheckHostError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) DescribeUserPrivileges(request *DescribeUserPrivilegesRequest) (response *DescribeUserPrivilegesResponse, err error) {
    return c.DescribeUserPrivilegesWithContext(context.Background(), request)
}

// DescribeUserPrivileges
// 本接口（DescribeUserPrivileges）提供查询用户的权限功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYPRIVILEGESERROR = "FailedOperation.DBQueryPrivilegesError"
//  INVALIDPARAMETERVALUE_CHECKHOSTERROR = "InvalidParameterValue.CheckHostError"
//  INVALIDPARAMETERVALUE_CHECKNAMEERROR = "InvalidParameterValue.CheckNameError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) DescribeUserPrivilegesWithContext(ctx context.Context, request *DescribeUserPrivilegesRequest) (response *DescribeUserPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeUserPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeUserPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
    request = &DescribeUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "DescribeUsers")
    
    
    return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
    response = &DescribeUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsers
// 本接口（DescribeUsers）提供查询用户列表功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBQUERYUSERERROR = "FailedOperation.DBQueryUserError"
//  FAILEDOPERATION_QUERYUSERERROR = "FailedOperation.QueryUserError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    return c.DescribeUsersWithContext(context.Background(), request)
}

// DescribeUsers
// 本接口（DescribeUsers）提供查询用户列表功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBQUERYUSERERROR = "FailedOperation.DBQueryUserError"
//  FAILEDOPERATION_QUERYUSERERROR = "FailedOperation.QueryUserError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeUsersWithContext(ctx context.Context, request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
    if request == nil {
        request = NewDescribeUsersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "DescribeUsers")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsersResponse()
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

func NewExpandInstanceRequest() (request *ExpandInstanceRequest) {
    request = &ExpandInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ExpandInstance")
    
    
    return
}

func NewExpandInstanceResponse() (response *ExpandInstanceResponse) {
    response = &ExpandInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExpandInstance
// 本接口（ExpandInstance）提供横向扩容实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_EXPANDINSTANCEERROR = "FailedOperation.ExpandInstanceError"
//  LIMITEXCEEDED_OUTOFLIMITERROR = "LimitExceeded.OutOfLimitError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
func (c *Client) ExpandInstance(request *ExpandInstanceRequest) (response *ExpandInstanceResponse, err error) {
    return c.ExpandInstanceWithContext(context.Background(), request)
}

// ExpandInstance
// 本接口（ExpandInstance）提供横向扩容实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_EXPANDINSTANCEERROR = "FailedOperation.ExpandInstanceError"
//  LIMITEXCEEDED_OUTOFLIMITERROR = "LimitExceeded.OutOfLimitError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
func (c *Client) ExpandInstanceWithContext(ctx context.Context, request *ExpandInstanceRequest) (response *ExpandInstanceResponse, err error) {
    if request == nil {
        request = NewExpandInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ExpandInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExpandInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewExpandInstanceResponse()
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

func NewModifyDBInstanceVPortRequest() (request *ModifyDBInstanceVPortRequest) {
    request = &ModifyDBInstanceVPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyDBInstanceVPort")
    
    
    return
}

func NewModifyDBInstanceVPortResponse() (response *ModifyDBInstanceVPortResponse) {
    response = &ModifyDBInstanceVPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceVPort
// 本接口(ModifyDBInstanceVPort)修改实例VPC端口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBUPDATEINSTANCEERROR = "FailedOperation.DBUpdateInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"
//  INVALIDPARAMETERVALUE_VPORTRANGEERROR = "InvalidParameterValue.VportRangeError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyDBInstanceVPort(request *ModifyDBInstanceVPortRequest) (response *ModifyDBInstanceVPortResponse, err error) {
    return c.ModifyDBInstanceVPortWithContext(context.Background(), request)
}

// ModifyDBInstanceVPort
// 本接口(ModifyDBInstanceVPort)修改实例VPC端口
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBUPDATEINSTANCEERROR = "FailedOperation.DBUpdateInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  INVALIDPARAMETERVALUE_CHECKVPCERROR = "InvalidParameterValue.CheckVpcError"
//  INVALIDPARAMETERVALUE_VPORTRANGEERROR = "InvalidParameterValue.VportRangeError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyDBInstanceVPortWithContext(ctx context.Context, request *ModifyDBInstanceVPortRequest) (response *ModifyDBInstanceVPortResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceVPortRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyDBInstanceVPort")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceVPort require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceVPortResponse()
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
// 修改实例备份策略 ModifyDBSBackupPolicy
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
// 修改实例备份策略 ModifyDBSBackupPolicy
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
// 修改实例备份备注 ModifyDBSBackupSetComment
//
// 可能返回的错误码:
//  INTERNALERROR_ROUTERNOTFOUND = "InternalError.RouterNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifyDBSBackupSetComment(request *ModifyDBSBackupSetCommentRequest) (response *ModifyDBSBackupSetCommentResponse, err error) {
    return c.ModifyDBSBackupSetCommentWithContext(context.Background(), request)
}

// ModifyDBSBackupSetComment
// 修改实例备份备注 ModifyDBSBackupSetComment
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

func NewModifyInstanceNetworkRequest() (request *ModifyInstanceNetworkRequest) {
    request = &ModifyInstanceNetworkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyInstanceNetwork")
    
    
    return
}

func NewModifyInstanceNetworkResponse() (response *ModifyInstanceNetworkResponse) {
    response = &ModifyInstanceNetworkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceNetwork
// 本接口（ModifyInstanceNetwork）用于修改实例所属网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
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
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNetwork(request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    return c.ModifyInstanceNetworkWithContext(context.Background(), request)
}

// ModifyInstanceNetwork
// 本接口（ModifyInstanceNetwork）用于修改实例所属网络
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
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
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceNetworkWithContext(ctx context.Context, request *ModifyInstanceNetworkRequest) (response *ModifyInstanceNetworkResponse, err error) {
    if request == nil {
        request = NewModifyInstanceNetworkRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyInstanceNetwork")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceNetwork require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceNetworkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceSSLStatusRequest() (request *ModifyInstanceSSLStatusRequest) {
    request = &ModifyInstanceSSLStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyInstanceSSLStatus")
    
    
    return
}

func NewModifyInstanceSSLStatusResponse() (response *ModifyInstanceSSLStatusResponse) {
    response = &ModifyInstanceSSLStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceSSLStatus
// 本接口（ModifyInstanceSSLStatus）提供开关实例SSL的功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
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
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceSSLStatus(request *ModifyInstanceSSLStatusRequest) (response *ModifyInstanceSSLStatusResponse, err error) {
    return c.ModifyInstanceSSLStatusWithContext(context.Background(), request)
}

// ModifyInstanceSSLStatus
// 本接口（ModifyInstanceSSLStatus）提供开关实例SSL的功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_APPLYVIPFAILED = "FailedOperation.ApplyVipFailed"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
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
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
func (c *Client) ModifyInstanceSSLStatusWithContext(ctx context.Context, request *ModifyInstanceSSLStatusRequest) (response *ModifyInstanceSSLStatusResponse, err error) {
    if request == nil {
        request = NewModifyInstanceSSLStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyInstanceSSLStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceSSLStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceSSLStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceWindowRequest() (request *ModifyMaintenanceWindowRequest) {
    request = &ModifyMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyMaintenanceWindow")
    
    
    return
}

func NewModifyMaintenanceWindowResponse() (response *ModifyMaintenanceWindowResponse) {
    response = &ModifyMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMaintenanceWindow
// 新增/修改实例维护时间窗口配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBUPSERTMAINTENANCEWINDOWERROR = "FailedOperation.DBUpsertMaintenanceWindowError"
//  INVALIDPARAMETERVALUE_MAINTENANCEWINDOWPARAMERROR = "InvalidParameterValue.MaintenanceWindowParamError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyMaintenanceWindow(request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    return c.ModifyMaintenanceWindowWithContext(context.Background(), request)
}

// ModifyMaintenanceWindow
// 新增/修改实例维护时间窗口配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_DBUPSERTMAINTENANCEWINDOWERROR = "FailedOperation.DBUpsertMaintenanceWindowError"
//  INVALIDPARAMETERVALUE_MAINTENANCEWINDOWPARAMERROR = "InvalidParameterValue.MaintenanceWindowParamError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) ModifyMaintenanceWindowWithContext(ctx context.Context, request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewModifyMaintenanceWindowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyMaintenanceWindow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserPrivilegesRequest() (request *ModifyUserPrivilegesRequest) {
    request = &ModifyUserPrivilegesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ModifyUserPrivileges")
    
    
    return
}

func NewModifyUserPrivilegesResponse() (response *ModifyUserPrivilegesResponse) {
    response = &ModifyUserPrivilegesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserPrivileges
// 本接口(ModifyPrivileges)修改用户权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBCOUNTLIMITERROR = "FailedOperation.DbCountLimitError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
//  FAILEDOPERATION_TABLECOUNTLIMITERROR = "FailedOperation.TableCountLimitError"
//  FAILEDOPERATION_UPDATEPRIVILEGESERROR = "FailedOperation.UpdatePrivilegesError"
func (c *Client) ModifyUserPrivileges(request *ModifyUserPrivilegesRequest) (response *ModifyUserPrivilegesResponse, err error) {
    return c.ModifyUserPrivilegesWithContext(context.Background(), request)
}

// ModifyUserPrivileges
// 本接口(ModifyPrivileges)修改用户权限
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBCOUNTLIMITERROR = "FailedOperation.DbCountLimitError"
//  FAILEDOPERATION_QUERYDBERROR = "FailedOperation.QueryDBError"
//  FAILEDOPERATION_TABLECOUNTLIMITERROR = "FailedOperation.TableCountLimitError"
//  FAILEDOPERATION_UPDATEPRIVILEGESERROR = "FailedOperation.UpdatePrivilegesError"
func (c *Client) ModifyUserPrivilegesWithContext(ctx context.Context, request *ModifyUserPrivilegesRequest) (response *ModifyUserPrivilegesResponse, err error) {
    if request == nil {
        request = NewModifyUserPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ModifyUserPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewResetUserPasswordRequest() (request *ResetUserPasswordRequest) {
    request = &ResetUserPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "ResetUserPassword")
    
    
    return
}

func NewResetUserPasswordResponse() (response *ResetUserPasswordResponse) {
    response = &ResetUserPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetUserPassword
// 本接口（ResetUserPassword）提供重置用户密码功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYUSERERROR = "FailedOperation.DBQueryUserError"
//  FAILEDOPERATION_DBUPDATEUSERERROR = "FailedOperation.DBUpdateUserError"
//  FAILEDOPERATION_QUERYUSERERROR = "FailedOperation.QueryUserError"
//  INVALIDPARAMETERVALUE_CHECKPASSWDERROR = "InvalidParameterValue.CheckPasswdError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_RESETUSERPASSWORDERROR = "OperationDenied.ResetUserPasswordError"
func (c *Client) ResetUserPassword(request *ResetUserPasswordRequest) (response *ResetUserPasswordResponse, err error) {
    return c.ResetUserPasswordWithContext(context.Background(), request)
}

// ResetUserPassword
// 本接口（ResetUserPassword）提供重置用户密码功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DBQUERYUSERERROR = "FailedOperation.DBQueryUserError"
//  FAILEDOPERATION_DBUPDATEUSERERROR = "FailedOperation.DBUpdateUserError"
//  FAILEDOPERATION_QUERYUSERERROR = "FailedOperation.QueryUserError"
//  INVALIDPARAMETERVALUE_CHECKPASSWDERROR = "InvalidParameterValue.CheckPasswdError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_RESETUSERPASSWORDERROR = "OperationDenied.ResetUserPasswordError"
func (c *Client) ResetUserPasswordWithContext(ctx context.Context, request *ResetUserPasswordRequest) (response *ResetUserPasswordResponse, err error) {
    if request == nil {
        request = NewResetUserPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "ResetUserPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetUserPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetUserPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewRestartDBInstancesRequest() (request *RestartDBInstancesRequest) {
    request = &RestartDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "RestartDBInstances")
    
    
    return
}

func NewRestartDBInstancesResponse() (response *RestartDBInstancesResponse) {
    response = &RestartDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestartDBInstances
// 本接口（RestartDBInstances）用于重启数据库实例
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CHECKSUPPORTACTIONERROR = "FailedOperation.CheckSupportActionError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) RestartDBInstances(request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    return c.RestartDBInstancesWithContext(context.Background(), request)
}

// RestartDBInstances
// 本接口（RestartDBInstances）用于重启数据库实例
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CHECKSUPPORTACTIONERROR = "FailedOperation.CheckSupportActionError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) RestartDBInstancesWithContext(ctx context.Context, request *RestartDBInstancesRequest) (response *RestartDBInstancesResponse, err error) {
    if request == nil {
        request = NewRestartDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "RestartDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tdmysql", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstance
// 本接口（UpgradeInstance）提供纵向扩容实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CHECKSUPPORTACTIONERROR = "FailedOperation.CheckSupportActionError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
//  OPERATIONDENIED_CHECKDECREASESPECERROR = "OperationDenied.CheckDecreaseSpecError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 本接口（UpgradeInstance）提供纵向扩容实例功能
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMAUTHERROR = "AuthFailure.CamAuthError"
//  AUTHFAILURE_CHECKCAMAUTHERROR = "AuthFailure.CheckCamAuthError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CHECKSUPPORTACTIONERROR = "FailedOperation.CheckSupportActionError"
//  FAILEDOPERATION_CREATEFLOWERROR = "FailedOperation.CreateFlowError"
//  FAILEDOPERATION_DBQUERYINSTANCEERROR = "FailedOperation.DBQueryInstanceError"
//  FAILEDOPERATION_LOCKINSTANCEERROR = "FailedOperation.LockInstanceError"
//  INVALIDPARAMETERVALUE_CHECKSPECERROR = "InvalidParameterValue.CheckSpecError"
//  LIMITEXCEEDED_OUTOFSPECLIMITERROR = "LimitExceeded.OutOfSpecLimitError"
//  OPERATIONDENIED_CHECKDECREASESPECERROR = "OperationDenied.CheckDecreaseSpecError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tdmysql", APIVersion, "UpgradeInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}
