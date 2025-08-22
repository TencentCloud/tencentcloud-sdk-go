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
// 本接口（AddTimeWindow）用于添加云数据库实例的维护时间窗口，以指定实例在哪些时间段可以自动执行切换访问操作。
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
// 本接口（AddTimeWindow）用于添加云数据库实例的维护时间窗口，以指定实例在哪些时间段可以自动执行切换访问操作。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "AddTimeWindow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddTimeWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddTimeWindowResponse()
    err = c.Send(request, response)
    return
}

func NewAdjustCdbProxyRequest() (request *AdjustCdbProxyRequest) {
    request = &AdjustCdbProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "AdjustCdbProxy")
    
    
    return
}

func NewAdjustCdbProxyResponse() (response *AdjustCdbProxyResponse) {
    response = &AdjustCdbProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AdjustCdbProxy
// 本接口（AdjustCdbProxy）用于调整数据库代理配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYNODECOUNTLIMITERROR = "OperationDenied.ProxyNodeCountLimitError"
func (c *Client) AdjustCdbProxy(request *AdjustCdbProxyRequest) (response *AdjustCdbProxyResponse, err error) {
    return c.AdjustCdbProxyWithContext(context.Background(), request)
}

// AdjustCdbProxy
// 本接口（AdjustCdbProxy）用于调整数据库代理配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_PROXYGROUPSTATUSERROR = "FailedOperation.ProxyGroupStatusError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYNODECOUNTLIMITERROR = "OperationDenied.ProxyNodeCountLimitError"
func (c *Client) AdjustCdbProxyWithContext(ctx context.Context, request *AdjustCdbProxyRequest) (response *AdjustCdbProxyResponse, err error) {
    if request == nil {
        request = NewAdjustCdbProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "AdjustCdbProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AdjustCdbProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewAdjustCdbProxyResponse()
    err = c.Send(request, response)
    return
}

func NewAdjustCdbProxyAddressRequest() (request *AdjustCdbProxyAddressRequest) {
    request = &AdjustCdbProxyAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "AdjustCdbProxyAddress")
    
    
    return
}

func NewAdjustCdbProxyAddressResponse() (response *AdjustCdbProxyAddressResponse) {
    response = &AdjustCdbProxyAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AdjustCdbProxyAddress
// 本接口（AdjustCdbProxyAddress）用于调整数据库代理地址配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
//  OPERATIONDENIED_UNSUPPORTCREATEADDRESSERROR = "OperationDenied.UnsupportCreateAddressError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) AdjustCdbProxyAddress(request *AdjustCdbProxyAddressRequest) (response *AdjustCdbProxyAddressResponse, err error) {
    return c.AdjustCdbProxyAddressWithContext(context.Background(), request)
}

// AdjustCdbProxyAddress
// 本接口（AdjustCdbProxyAddress）用于调整数据库代理地址配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
//  OPERATIONDENIED_UNSUPPORTCREATEADDRESSERROR = "OperationDenied.UnsupportCreateAddressError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) AdjustCdbProxyAddressWithContext(ctx context.Context, request *AdjustCdbProxyAddressRequest) (response *AdjustCdbProxyAddressResponse, err error) {
    if request == nil {
        request = NewAdjustCdbProxyAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "AdjustCdbProxyAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AdjustCdbProxyAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewAdjustCdbProxyAddressResponse()
    err = c.Send(request, response)
    return
}

func NewAnalyzeAuditLogsRequest() (request *AnalyzeAuditLogsRequest) {
    request = &AnalyzeAuditLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "AnalyzeAuditLogs")
    
    
    return
}

func NewAnalyzeAuditLogsResponse() (response *AnalyzeAuditLogsResponse) {
    response = &AnalyzeAuditLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AnalyzeAuditLogs
// 本接口（AnalyzeAuditLogs）用于在不同过滤条件下的审计日志结果集中，选定特定的数据列进行聚合统计。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALINNERSERVICEERROR = "InternalError.InternalInnerServiceError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AnalyzeAuditLogs(request *AnalyzeAuditLogsRequest) (response *AnalyzeAuditLogsResponse, err error) {
    return c.AnalyzeAuditLogsWithContext(context.Background(), request)
}

// AnalyzeAuditLogs
// 本接口（AnalyzeAuditLogs）用于在不同过滤条件下的审计日志结果集中，选定特定的数据列进行聚合统计。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALINNERSERVICEERROR = "InternalError.InternalInnerServiceError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AnalyzeAuditLogsWithContext(ctx context.Context, request *AnalyzeAuditLogsRequest) (response *AnalyzeAuditLogsResponse, err error) {
    if request == nil {
        request = NewAnalyzeAuditLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "AnalyzeAuditLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AnalyzeAuditLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewAnalyzeAuditLogsResponse()
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) AssociateSecurityGroupsWithContext(ctx context.Context, request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewAssociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "AssociateSecurityGroups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "BalanceRoGroupLoad")
    
    if c.GetCredential() == nil {
        return nil, errors.New("BalanceRoGroupLoad require credential")
    }

    request.SetContext(ctx)
    
    response = NewBalanceRoGroupLoadResponse()
    err = c.Send(request, response)
    return
}

func NewCheckMigrateClusterRequest() (request *CheckMigrateClusterRequest) {
    request = &CheckMigrateClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CheckMigrateCluster")
    
    
    return
}

func NewCheckMigrateClusterResponse() (response *CheckMigrateClusterResponse) {
    response = &CheckMigrateClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckMigrateCluster
// 本接口（CheckMigrateCluster）用于高可用实例一键迁移到云盘版校验。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) CheckMigrateCluster(request *CheckMigrateClusterRequest) (response *CheckMigrateClusterResponse, err error) {
    return c.CheckMigrateClusterWithContext(context.Background(), request)
}

// CheckMigrateCluster
// 本接口（CheckMigrateCluster）用于高可用实例一键迁移到云盘版校验。
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) CheckMigrateClusterWithContext(ctx context.Context, request *CheckMigrateClusterRequest) (response *CheckMigrateClusterResponse, err error) {
    if request == nil {
        request = NewCheckMigrateClusterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CheckMigrateCluster")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckMigrateCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckMigrateClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCloseAuditServiceRequest() (request *CloseAuditServiceRequest) {
    request = &CloseAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CloseAuditService")
    
    
    return
}

func NewCloseAuditServiceResponse() (response *CloseAuditServiceResponse) {
    response = &CloseAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseAuditService
// 实例关闭审计服务
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) CloseAuditService(request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    return c.CloseAuditServiceWithContext(context.Background(), request)
}

// CloseAuditService
// 实例关闭审计服务
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) CloseAuditServiceWithContext(ctx context.Context, request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    if request == nil {
        request = NewCloseAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CloseAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAuditServiceResponse()
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
// 本接口（CloseCDBProxy）用于关闭数据库代理。
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
// 本接口（CloseCDBProxy）用于关闭数据库代理。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CloseCDBProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseCDBProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseCDBProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCloseCdbProxyAddressRequest() (request *CloseCdbProxyAddressRequest) {
    request = &CloseCdbProxyAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CloseCdbProxyAddress")
    
    
    return
}

func NewCloseCdbProxyAddressResponse() (response *CloseCdbProxyAddressResponse) {
    response = &CloseCdbProxyAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseCdbProxyAddress
// 本接口（CloseCdbProxyAddress）用于请求关闭数据库代理地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYADDRESSLIMITERROR = "OperationDenied.ProxyAddressLimitError"
//  OPERATIONDENIED_PROXYADDRESSNOTFUND = "OperationDenied.ProxyAddressNotFund"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CloseCdbProxyAddress(request *CloseCdbProxyAddressRequest) (response *CloseCdbProxyAddressResponse, err error) {
    return c.CloseCdbProxyAddressWithContext(context.Background(), request)
}

// CloseCdbProxyAddress
// 本接口（CloseCdbProxyAddress）用于请求关闭数据库代理地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYADDRESSLIMITERROR = "OperationDenied.ProxyAddressLimitError"
//  OPERATIONDENIED_PROXYADDRESSNOTFUND = "OperationDenied.ProxyAddressNotFund"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CloseCdbProxyAddressWithContext(ctx context.Context, request *CloseCdbProxyAddressRequest) (response *CloseCdbProxyAddressResponse, err error) {
    if request == nil {
        request = NewCloseCdbProxyAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CloseCdbProxyAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseCdbProxyAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseCdbProxyAddressResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSSLRequest() (request *CloseSSLRequest) {
    request = &CloseSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CloseSSL")
    
    
    return
}

func NewCloseSSLResponse() (response *CloseSSLResponse) {
    response = &CloseSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseSSL
// 本接口（CloseSSL）用于关闭 SSL 连接功能。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) CloseSSL(request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    return c.CloseSSLWithContext(context.Background(), request)
}

// CloseSSL
// 本接口（CloseSSL）用于关闭 SSL 连接功能。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) CloseSSLWithContext(ctx context.Context, request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    if request == nil {
        request = NewCloseSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CloseSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseSSLResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CloseWanService")
    
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
// 本接口（CreateAccounts）用于创建云数据库的账户，需要指定新的账户名和域名，以及所对应的密码，同时可以设置账号的备注信息以及最大可用连接数。
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
// 本接口（CreateAccounts）用于创建云数据库的账户，需要指定新的账户名和域名，以及所对应的密码，同时可以设置账号的备注信息以及最大可用连接数。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateAccounts")
    
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
// 本接口（CreateAuditLogFile）用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITLOGFILEERROR = "FailedOperation.CreateAuditLogFileError"
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
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
// 本接口（CreateAuditLogFile）用于创建云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITLOGFILEERROR = "FailedOperation.CreateAuditLogFileError"
//  INTERNALERROR_AUDITCREATELOGFILEERROR = "InternalError.AuditCreateLogFileError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateAuditLogFile")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateAuditPolicy")
    
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
// 不再支持审计规则创建
//
// 
//
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
// 不再支持审计规则创建
//
// 
//
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateAuditRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditRuleTemplateRequest() (request *CreateAuditRuleTemplateRequest) {
    request = &CreateAuditRuleTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateAuditRuleTemplate")
    
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateAuditRuleTemplate")
    
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
// 本接口（CreateBackup）用于创建数据库备份。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_TASKCONFLICTERROR = "FailedOperation.TaskConflictError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OVERQUOTA = "OverQuota"
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    return c.CreateBackupWithContext(context.Background(), request)
}

// CreateBackup
// 本接口（CreateBackup）用于创建数据库备份。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  FAILEDOPERATION_TASKCONFLICTERROR = "FailedOperation.TaskConflictError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OVERQUOTA = "OverQuota"
func (c *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
    if request == nil {
        request = NewCreateBackupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCdbProxyRequest() (request *CreateCdbProxyRequest) {
    request = &CreateCdbProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateCdbProxy")
    
    
    return
}

func NewCreateCdbProxyResponse() (response *CreateCdbProxyResponse) {
    response = &CreateCdbProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCdbProxy
// 本接口（CreateCdbProxy）用于主实例创建数据库代理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REPEATCREATEPROXYERROR = "FailedOperation.RepeatCreateProxyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) CreateCdbProxy(request *CreateCdbProxyRequest) (response *CreateCdbProxyResponse, err error) {
    return c.CreateCdbProxyWithContext(context.Background(), request)
}

// CreateCdbProxy
// 本接口（CreateCdbProxy）用于主实例创建数据库代理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REPEATCREATEPROXYERROR = "FailedOperation.RepeatCreateProxyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) CreateCdbProxyWithContext(ctx context.Context, request *CreateCdbProxyRequest) (response *CreateCdbProxyResponse, err error) {
    if request == nil {
        request = NewCreateCdbProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateCdbProxy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCdbProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCdbProxyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCdbProxyAddressRequest() (request *CreateCdbProxyAddressRequest) {
    request = &CreateCdbProxyAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateCdbProxyAddress")
    
    
    return
}

func NewCreateCdbProxyAddressResponse() (response *CreateCdbProxyAddressResponse) {
    response = &CreateCdbProxyAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCdbProxyAddress
// 本接口（CreateCdbProxyAddress）用于数据库代理增加代理地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_VPCIPINUSEERROR = "FailedOperation.VpcIpInUseError"
//  FAILEDOPERATION_VPCIPINVALIDERROR = "FailedOperation.VpcIpInvalidError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_PROXYADDRESSLIMITERROR = "OperationDenied.ProxyAddressLimitError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CreateCdbProxyAddress(request *CreateCdbProxyAddressRequest) (response *CreateCdbProxyAddressResponse, err error) {
    return c.CreateCdbProxyAddressWithContext(context.Background(), request)
}

// CreateCdbProxyAddress
// 本接口（CreateCdbProxyAddress）用于数据库代理增加代理地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_VPCIPINUSEERROR = "FailedOperation.VpcIpInUseError"
//  FAILEDOPERATION_VPCIPINVALIDERROR = "FailedOperation.VpcIpInvalidError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED_PROXYADDRESSLIMITERROR = "OperationDenied.ProxyAddressLimitError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) CreateCdbProxyAddressWithContext(ctx context.Context, request *CreateCdbProxyAddressRequest) (response *CreateCdbProxyAddressResponse, err error) {
    if request == nil {
        request = NewCreateCdbProxyAddressRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateCdbProxyAddress")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCdbProxyAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCdbProxyAddressResponse()
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
// 本接口（CreateCloneInstance）用于从目标源实例创建一个克隆实例，可以指定克隆实例回档到源实例的指定物理备份文件或者指定的回档时间点。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
// 本接口（CreateCloneInstance）用于从目标源实例创建一个克隆实例，可以指定克隆实例回档到源实例的指定物理备份文件或者指定的回档时间点。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateCloneInstance")
    
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
// 本接口（CreateDBImportJob）用于创建云数据库数据导入任务。
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
// 本接口（CreateDBImportJob）用于创建云数据库数据导入任务。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateDBImportJob")
    
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
// 本接口（CreateDBInstance）用于创建包年包月的云数据库实例（包括主实例、灾备实例和只读实例），可通过传入实例规格、MySQL 版本号、购买时长和数量等信息创建云数据库实例。
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
// 5. 当入参指定 ParamTemplateId 或 AlarmPolicyList 时，需将SDK提升至最新版本方可支持；
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
// 本接口（CreateDBInstance）用于创建包年包月的云数据库实例（包括主实例、灾备实例和只读实例），可通过传入实例规格、MySQL 版本号、购买时长和数量等信息创建云数据库实例。
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
// 5. 当入参指定 ParamTemplateId 或 AlarmPolicyList 时，需将SDK提升至最新版本方可支持；
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateDBInstance")
    
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
// 本接口（CreateDBInstanceHour）用于创建按量计费的实例，可通过传入实例规格、MySQL 版本号和数量等信息创建云数据库实例，支持主实例、灾备实例和只读实例的创建。
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
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceHour(request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    return c.CreateDBInstanceHourWithContext(context.Background(), request)
}

// CreateDBInstanceHour
// 本接口（CreateDBInstanceHour）用于创建按量计费的实例，可通过传入实例规格、MySQL 版本号和数量等信息创建云数据库实例，支持主实例、灾备实例和只读实例的创建。
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
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_OTHERODERINPROCESS = "OperationDenied.OtherOderInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGPASSWORD = "OperationDenied.WrongPassword"
func (c *Client) CreateDBInstanceHourWithContext(ctx context.Context, request *CreateDBInstanceHourRequest) (response *CreateDBInstanceHourResponse, err error) {
    if request == nil {
        request = NewCreateDBInstanceHourRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateDBInstanceHour")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBInstanceHour require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBInstanceHourResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatabaseRequest() (request *CreateDatabaseRequest) {
    request = &CreateDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateDatabase")
    
    
    return
}

func NewCreateDatabaseResponse() (response *CreateDatabaseResponse) {
    response = &CreateDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDatabase
// 本接口(CreateDatabase)用于在云数据库实例中创建数据库。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) CreateDatabase(request *CreateDatabaseRequest) (response *CreateDatabaseResponse, err error) {
    return c.CreateDatabaseWithContext(context.Background(), request)
}

// CreateDatabase
// 本接口(CreateDatabase)用于在云数据库实例中创建数据库。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
func (c *Client) CreateDatabaseWithContext(ctx context.Context, request *CreateDatabaseRequest) (response *CreateDatabaseResponse, err error) {
    if request == nil {
        request = NewCreateDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatabaseResponse()
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
// 本接口（CreateDeployGroup）用于创建放置实例的置放群组。
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
// 本接口（CreateDeployGroup）用于创建放置实例的置放群组。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateDeployGroup")
    
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
// 该接口（CreateParamTemplate）用于创建参数模板。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
// 该接口（CreateParamTemplate）用于创建参数模板。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateParamTemplate")
    
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateRoInstanceIp")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoInstanceIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoInstanceIpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRotationPasswordRequest() (request *CreateRotationPasswordRequest) {
    request = &CreateRotationPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "CreateRotationPassword")
    
    
    return
}

func NewCreateRotationPasswordResponse() (response *CreateRotationPasswordResponse) {
    response = &CreateRotationPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRotationPassword
// 本接口（CreateRotationPassword）用于开启密码轮转。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEROVIPERROR = "FailedOperation.CreateRoVipError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
func (c *Client) CreateRotationPassword(request *CreateRotationPasswordRequest) (response *CreateRotationPasswordResponse, err error) {
    return c.CreateRotationPasswordWithContext(context.Background(), request)
}

// CreateRotationPassword
// 本接口（CreateRotationPassword）用于开启密码轮转。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEROVIPERROR = "FailedOperation.CreateRoVipError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
func (c *Client) CreateRotationPasswordWithContext(ctx context.Context, request *CreateRotationPasswordRequest) (response *CreateRotationPasswordResponse, err error) {
    if request == nil {
        request = NewCreateRotationPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "CreateRotationPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRotationPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRotationPasswordResponse()
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
// 本接口（DeleteAccounts）用于删除云数据库的账户。
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
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DeleteAccounts(request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    return c.DeleteAccountsWithContext(context.Background(), request)
}

// DeleteAccounts
// 本接口（DeleteAccounts）用于删除云数据库的账户。
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
//  OPERATIONDENIED_INSTTYPENOTSUPPORT = "OperationDenied.InstTypeNotSupport"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DeleteAccountsWithContext(ctx context.Context, request *DeleteAccountsRequest) (response *DeleteAccountsResponse, err error) {
    if request == nil {
        request = NewDeleteAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteAccounts")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteAuditLogFile")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteAuditPolicy")
    
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
// 不再支持审计规则创建
//
// 
//
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
// 不再支持审计规则创建
//
// 
//
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteAuditRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRuleTemplatesRequest() (request *DeleteAuditRuleTemplatesRequest) {
    request = &DeleteAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteAuditRuleTemplates")
    
    
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
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
func (c *Client) DeleteAuditRuleTemplates(request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    return c.DeleteAuditRuleTemplatesWithContext(context.Background(), request)
}

// DeleteAuditRuleTemplates
// 本接口（DeleteAuditRuleTemplates）用于删除审计规则模板。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
func (c *Client) DeleteAuditRuleTemplatesWithContext(ctx context.Context, request *DeleteAuditRuleTemplatesRequest) (response *DeleteAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteAuditRuleTemplates")
    
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
// 本接口（DeleteBackup）用于删除数据库备份。本接口只支持删除手动发起的备份。
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
// 本接口（DeleteBackup）用于删除数据库备份。本接口只支持删除手动发起的备份。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteBackup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBackup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBackupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDatabaseRequest() (request *DeleteDatabaseRequest) {
    request = &DeleteDatabaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteDatabase")
    
    
    return
}

func NewDeleteDatabaseResponse() (response *DeleteDatabaseResponse) {
    response = &DeleteDatabaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDatabase
// 本接口(DeleteDatabase)用于在云数据库实例中删除数据库。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DeleteDatabase(request *DeleteDatabaseRequest) (response *DeleteDatabaseResponse, err error) {
    return c.DeleteDatabaseWithContext(context.Background(), request)
}

// DeleteDatabase
// 本接口(DeleteDatabase)用于在云数据库实例中删除数据库。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DeleteDatabaseWithContext(ctx context.Context, request *DeleteDatabaseRequest) (response *DeleteDatabaseResponse, err error) {
    if request == nil {
        request = NewDeleteDatabaseRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteDatabase")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDatabase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatabaseResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteDeployGroups")
    
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
// 该接口（DeleteParamTemplate）用于删除参数模板。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
// 该接口（DeleteParamTemplate）用于删除参数模板。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRotationPasswordRequest() (request *DeleteRotationPasswordRequest) {
    request = &DeleteRotationPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DeleteRotationPassword")
    
    
    return
}

func NewDeleteRotationPasswordResponse() (response *DeleteRotationPasswordResponse) {
    response = &DeleteRotationPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRotationPassword
// 本接口（DeleteRotationPassword）用于关闭实例账户密码轮转。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRotationPassword(request *DeleteRotationPasswordRequest) (response *DeleteRotationPasswordResponse, err error) {
    return c.DeleteRotationPasswordWithContext(context.Background(), request)
}

// DeleteRotationPassword
// 本接口（DeleteRotationPassword）用于关闭实例账户密码轮转。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRotationPasswordWithContext(ctx context.Context, request *DeleteRotationPasswordRequest) (response *DeleteRotationPasswordResponse, err error) {
    if request == nil {
        request = NewDeleteRotationPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteRotationPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRotationPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRotationPasswordResponse()
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
// 本接口（DeleteTimeWindow）用于删除云数据库实例的维护时间窗口。删除实例维护时间窗口之后，默认的维护时间窗为每天的03:00-04:00，数据校验延迟阈值为10秒，即当选择在维护时间窗口内切换访问新实例时，默认会在03:00-04:00点进行切换访问新实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DeleteTimeWindow(request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    return c.DeleteTimeWindowWithContext(context.Background(), request)
}

// DeleteTimeWindow
// 本接口（DeleteTimeWindow）用于删除云数据库实例的维护时间窗口。删除实例维护时间窗口之后，默认的维护时间窗为每天的03:00-04:00，数据校验延迟阈值为10秒，即当选择在维护时间窗口内切换访问新实例时，默认会在03:00-04:00点进行切换访问新实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DeleteTimeWindowWithContext(ctx context.Context, request *DeleteTimeWindowRequest) (response *DeleteTimeWindowResponse, err error) {
    if request == nil {
        request = NewDeleteTimeWindowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DeleteTimeWindow")
    
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
// 本接口（DescribeAccountPrivileges）用于查询云数据库账户支持的权限信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccountPrivileges(request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    return c.DescribeAccountPrivilegesWithContext(context.Background(), request)
}

// DescribeAccountPrivileges
// 本接口（DescribeAccountPrivileges）用于查询云数据库账户支持的权限信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_CREATEACCOUNTERROR = "FailedOperation.CreateAccountError"
//  FAILEDOPERATION_GETPRIVILEGEERROR = "FailedOperation.GetPrivilegeError"
//  FAILEDOPERATION_PRIVILEGEDATAILLEGAL = "FailedOperation.PrivilegeDataIllegal"
//  FAILEDOPERATION_RESPONSEVALUEERROR = "FailedOperation.ResponseValueError"
//  FAILEDOPERATION_STARTFLOWERROR = "FailedOperation.StartFlowError"
//  FAILEDOPERATION_SUBMITASYNCTASKERROR = "FailedOperation.SubmitAsyncTaskError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_NOTSUPPORTMODIFYLOCALROOTHOSTERROR = "OperationDenied.NotSupportModifyLocalRootHostError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccountPrivilegesWithContext(ctx context.Context, request *DescribeAccountPrivilegesRequest) (response *DescribeAccountPrivilegesResponse, err error) {
    if request == nil {
        request = NewDescribeAccountPrivilegesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAccountPrivileges")
    
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
// 本接口（DescribeAccounts）用于查询云数据库的所有账户信息。
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    return c.DescribeAccountsWithContext(context.Background(), request)
}

// DescribeAccounts
// 本接口（DescribeAccounts）用于查询云数据库的所有账户信息。
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALASSERTERROR = "InternalError.InternalAssertError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALREQUESTERROR = "InternalError.InternalRequestError"
//  INTERNALERROR_REGEXPCOMPILEERROR = "InternalError.RegexpCompileError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_PRIVILEGESUNSUPPORTEDERROR = "UnsupportedOperation.PrivilegesUnsupportedError"
func (c *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
    if request == nil {
        request = NewDescribeAccountsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAccounts")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAsyncRequestInfo")
    
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
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
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
//  FAILEDOPERATION_INSTANCEQUERYERROR = "FailedOperation.InstanceQueryError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAuditConfigWithContext(ctx context.Context, request *DescribeAuditConfigRequest) (response *DescribeAuditConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAuditConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditInstanceListRequest() (request *DescribeAuditInstanceListRequest) {
    request = &DescribeAuditInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditInstanceList")
    
    
    return
}

func NewDescribeAuditInstanceListResponse() (response *DescribeAuditInstanceListResponse) {
    response = &DescribeAuditInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditInstanceList
// 本接口（DescribeAuditInstanceList）用于获取审计实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
func (c *Client) DescribeAuditInstanceList(request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    return c.DescribeAuditInstanceListWithContext(context.Background(), request)
}

// DescribeAuditInstanceList
// 本接口（DescribeAuditInstanceList）用于获取审计实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
func (c *Client) DescribeAuditInstanceListWithContext(ctx context.Context, request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAuditInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditInstanceList")
    
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
// 本接口（DescribeAuditLogFiles）用于查询云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditLogFiles(request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    return c.DescribeAuditLogFilesWithContext(context.Background(), request)
}

// DescribeAuditLogFiles
// 本接口（DescribeAuditLogFiles）用于查询云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_TRANSACTIOBEGINERROR = "InternalError.TransactioBeginError"
//  INTERNALERROR_TRANSACTIONBEGINERROR = "InternalError.TransactionBeginError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeAuditLogFilesWithContext(ctx context.Context, request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditLogFiles")
    
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditLogs")
    
    
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
//  FAILEDOPERATION_QUERYAUDITTASKFAILERROR = "FailedOperation.QueryAuditTaskFailError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_QUERYAUDITLOGSERROR = "OperationDenied.QueryAuditLogsError"
//  OPERATIONDENIED_RESOURCENOTFOUNDERROR = "OperationDenied.ResourceNotFoundError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    return c.DescribeAuditLogsWithContext(context.Background(), request)
}

// DescribeAuditLogs
// 本接口（DescribeAuditLogs）用于查询数据库审计日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYAUDITTASKFAILERROR = "FailedOperation.QueryAuditTaskFailError"
//  INTERNALERROR_AUDITDESCRIBELOGERROR = "InternalError.AuditDescribeLogError"
//  INTERNALERROR_AUDITERROR = "InternalError.AuditError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_QUERYAUDITLOGSERROR = "OperationDenied.QueryAuditLogsError"
//  OPERATIONDENIED_RESOURCENOTFOUNDERROR = "OperationDenied.ResourceNotFoundError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAuditLogsWithContext(ctx context.Context, request *DescribeAuditLogsRequest) (response *DescribeAuditLogsResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogsResponse()
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
// 本接口（DescribeAuditPolicies）用于查询云数据库实例的审计策略。
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
// 本接口（DescribeAuditPolicies）用于查询云数据库实例的审计策略。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditPolicies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleTemplateModifyHistoryRequest() (request *DescribeAuditRuleTemplateModifyHistoryRequest) {
    request = &DescribeAuditRuleTemplateModifyHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditRuleTemplateModifyHistory")
    
    
    return
}

func NewDescribeAuditRuleTemplateModifyHistoryResponse() (response *DescribeAuditRuleTemplateModifyHistoryResponse) {
    response = &DescribeAuditRuleTemplateModifyHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditRuleTemplateModifyHistory
// 本接口（DescribeAuditRuleTemplateModifyHistory）用于查询规则模板变更记录。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DescribeAuditRuleTemplateModifyHistory(request *DescribeAuditRuleTemplateModifyHistoryRequest) (response *DescribeAuditRuleTemplateModifyHistoryResponse, err error) {
    return c.DescribeAuditRuleTemplateModifyHistoryWithContext(context.Background(), request)
}

// DescribeAuditRuleTemplateModifyHistory
// 本接口（DescribeAuditRuleTemplateModifyHistory）用于查询规则模板变更记录。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DescribeAuditRuleTemplateModifyHistoryWithContext(ctx context.Context, request *DescribeAuditRuleTemplateModifyHistoryRequest) (response *DescribeAuditRuleTemplateModifyHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleTemplateModifyHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditRuleTemplateModifyHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRuleTemplateModifyHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRuleTemplateModifyHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRuleTemplatesRequest() (request *DescribeAuditRuleTemplatesRequest) {
    request = &DescribeAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeAuditRuleTemplates")
    
    
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
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DescribeAuditRuleTemplates(request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    return c.DescribeAuditRuleTemplatesWithContext(context.Background(), request)
}

// DescribeAuditRuleTemplates
// 本接口（DescribeAuditRuleTemplates）用于查询审计规则模板信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DescribeAuditRuleTemplatesWithContext(ctx context.Context, request *DescribeAuditRuleTemplatesRequest) (response *DescribeAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditRuleTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditRuleTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditRuleTemplatesResponse()
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
// 不再支持审计规则创建
//
// 
//
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
// 不再支持审计规则创建
//
// 
//
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeAuditRules")
    
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
// 本接口（DescribeBackupConfig）用于查询数据库备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupConfig(request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    return c.DescribeBackupConfigWithContext(context.Background(), request)
}

// DescribeBackupConfig
// 本接口（DescribeBackupConfig）用于查询数据库备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeBackupConfigWithContext(ctx context.Context, request *DescribeBackupConfigRequest) (response *DescribeBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackupConfig")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDecryptionKeyRequest() (request *DescribeBackupDecryptionKeyRequest) {
    request = &DescribeBackupDecryptionKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeBackupDecryptionKey")
    
    
    return
}

func NewDescribeBackupDecryptionKeyResponse() (response *DescribeBackupDecryptionKeyResponse) {
    response = &DescribeBackupDecryptionKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDecryptionKey
// 本接口（DescribeBackupDecryptionKey）用于查询备份文件解密密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
func (c *Client) DescribeBackupDecryptionKey(request *DescribeBackupDecryptionKeyRequest) (response *DescribeBackupDecryptionKeyResponse, err error) {
    return c.DescribeBackupDecryptionKeyWithContext(context.Background(), request)
}

// DescribeBackupDecryptionKey
// 本接口（DescribeBackupDecryptionKey）用于查询备份文件解密密钥。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_NOTSUPPORTBASIC = "OperationDenied.NotSupportBasic"
func (c *Client) DescribeBackupDecryptionKeyWithContext(ctx context.Context, request *DescribeBackupDecryptionKeyRequest) (response *DescribeBackupDecryptionKeyResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDecryptionKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackupDecryptionKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDecryptionKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDecryptionKeyResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackupDownloadRestriction")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackupEncryptionStatus")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackupOverview")
    
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
// 本接口（DescribeBackupSummaries）用于查询备份的统计情况，返回以实例为维度的备份占用容量，以及每个实例的数据备份和日志备份的个数和容量（容量单位为字节）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBackupSummaries(request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    return c.DescribeBackupSummariesWithContext(context.Background(), request)
}

// DescribeBackupSummaries
// 本接口（DescribeBackupSummaries）用于查询备份的统计情况，返回以实例为维度的备份占用容量，以及每个实例的数据备份和日志备份的个数和容量（容量单位为字节）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_SUBACCOUNTDENIED = "AuthFailure.SubAccountDenied"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_SUBACCOUNTOPERATIONDENIED = "OperationDenied.SubAccountOperationDenied"
func (c *Client) DescribeBackupSummariesWithContext(ctx context.Context, request *DescribeBackupSummariesRequest) (response *DescribeBackupSummariesResponse, err error) {
    if request == nil {
        request = NewDescribeBackupSummariesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackupSummaries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupSummaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupSummariesResponse()
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBackupsWithContext(ctx context.Context, request *DescribeBackupsRequest) (response *DescribeBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeBackupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBackups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBinlogBackupOverview")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeBinlogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBinlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCPUExpandStrategyInfoRequest() (request *DescribeCPUExpandStrategyInfoRequest) {
    request = &DescribeCPUExpandStrategyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCPUExpandStrategyInfo")
    
    
    return
}

func NewDescribeCPUExpandStrategyInfoResponse() (response *DescribeCPUExpandStrategyInfoResponse) {
    response = &DescribeCPUExpandStrategyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCPUExpandStrategyInfo
// 通过该 API 可以查询实例的 CPU 弹性扩容信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_REMOTECALLUNMARSHALERROR = "FailedOperation.RemoteCallUnmarshalError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
func (c *Client) DescribeCPUExpandStrategyInfo(request *DescribeCPUExpandStrategyInfoRequest) (response *DescribeCPUExpandStrategyInfoResponse, err error) {
    return c.DescribeCPUExpandStrategyInfoWithContext(context.Background(), request)
}

// DescribeCPUExpandStrategyInfo
// 通过该 API 可以查询实例的 CPU 弹性扩容信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_REMOTECALLUNMARSHALERROR = "FailedOperation.RemoteCallUnmarshalError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
func (c *Client) DescribeCPUExpandStrategyInfoWithContext(ctx context.Context, request *DescribeCPUExpandStrategyInfoRequest) (response *DescribeCPUExpandStrategyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCPUExpandStrategyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeCPUExpandStrategyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCPUExpandStrategyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCPUExpandStrategyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdbProxyInfoRequest() (request *DescribeCdbProxyInfoRequest) {
    request = &DescribeCdbProxyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCdbProxyInfo")
    
    
    return
}

func NewDescribeCdbProxyInfoResponse() (response *DescribeCdbProxyInfoResponse) {
    response = &DescribeCdbProxyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCdbProxyInfo
// 本接口（DescribeCdbProxyInfo）用于查询数据库代理详情信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DescribeCdbProxyInfo(request *DescribeCdbProxyInfoRequest) (response *DescribeCdbProxyInfoResponse, err error) {
    return c.DescribeCdbProxyInfoWithContext(context.Background(), request)
}

// DescribeCdbProxyInfo
// 本接口（DescribeCdbProxyInfo）用于查询数据库代理详情信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DescribeCdbProxyInfoWithContext(ctx context.Context, request *DescribeCdbProxyInfoRequest) (response *DescribeCdbProxyInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCdbProxyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeCdbProxyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdbProxyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdbProxyInfoResponse()
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
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
func (c *Client) DescribeCdbZoneConfig(request *DescribeCdbZoneConfigRequest) (response *DescribeCdbZoneConfigResponse, err error) {
    return c.DescribeCdbZoneConfigWithContext(context.Background(), request)
}

// DescribeCdbZoneConfig
// 本接口(DescribeCdbZoneConfig)用于查询云数据库各地域可售卖的规格配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
func (c *Client) DescribeCdbZoneConfigWithContext(ctx context.Context, request *DescribeCdbZoneConfigRequest) (response *DescribeCdbZoneConfigResponse, err error) {
    if request == nil {
        request = NewDescribeCdbZoneConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeCdbZoneConfig")
    
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
// 本接口（DescribeCloneList）用于查询用户实例的克隆任务列表。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCloneList(request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    return c.DescribeCloneListWithContext(context.Background(), request)
}

// DescribeCloneList
// 本接口（DescribeCloneList）用于查询用户实例的克隆任务列表。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DescribeCloneListWithContext(ctx context.Context, request *DescribeCloneListRequest) (response *DescribeCloneListResponse, err error) {
    if request == nil {
        request = NewDescribeCloneListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeCloneList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloneList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloneListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClusterInfoRequest() (request *DescribeClusterInfoRequest) {
    request = &DescribeClusterInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeClusterInfo")
    
    
    return
}

func NewDescribeClusterInfoResponse() (response *DescribeClusterInfoResponse) {
    response = &DescribeClusterInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeClusterInfo
// 本接口（DescribeClusterInfo）用于查询云盘版实例信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterInfo(request *DescribeClusterInfoRequest) (response *DescribeClusterInfoResponse, err error) {
    return c.DescribeClusterInfoWithContext(context.Background(), request)
}

// DescribeClusterInfo
// 本接口（DescribeClusterInfo）用于查询云盘版实例信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeClusterInfoWithContext(ctx context.Context, request *DescribeClusterInfoRequest) (response *DescribeClusterInfoResponse, err error) {
    if request == nil {
        request = NewDescribeClusterInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeClusterInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusterInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClusterInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCpuExpandHistoryRequest() (request *DescribeCpuExpandHistoryRequest) {
    request = &DescribeCpuExpandHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeCpuExpandHistory")
    
    
    return
}

func NewDescribeCpuExpandHistoryResponse() (response *DescribeCpuExpandHistoryResponse) {
    response = &DescribeCpuExpandHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCpuExpandHistory
// 本接口（DescribeCpuExpandHistory）用于查询扩容历史。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCpuExpandHistory(request *DescribeCpuExpandHistoryRequest) (response *DescribeCpuExpandHistoryResponse, err error) {
    return c.DescribeCpuExpandHistoryWithContext(context.Background(), request)
}

// DescribeCpuExpandHistory
// 本接口（DescribeCpuExpandHistory）用于查询扩容历史。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeCpuExpandHistoryWithContext(ctx context.Context, request *DescribeCpuExpandHistoryRequest) (response *DescribeCpuExpandHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeCpuExpandHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeCpuExpandHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCpuExpandHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCpuExpandHistoryResponse()
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
// 本接口（DescribeDBFeatures）用于查询云数据库版本属性，包括是否支持数据库加密、数据库审计等功能。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBFeatures(request *DescribeDBFeaturesRequest) (response *DescribeDBFeaturesResponse, err error) {
    return c.DescribeDBFeaturesWithContext(context.Background(), request)
}

// DescribeDBFeatures
// 本接口（DescribeDBFeatures）用于查询云数据库版本属性，包括是否支持数据库加密、数据库审计等功能。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBFeaturesWithContext(ctx context.Context, request *DescribeDBFeaturesRequest) (response *DescribeDBFeaturesResponse, err error) {
    if request == nil {
        request = NewDescribeDBFeaturesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBFeatures")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBImportRecords")
    
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceCharsetWithContext(ctx context.Context, request *DescribeDBInstanceCharsetRequest) (response *DescribeDBInstanceCharsetResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceCharsetRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstanceCharset")
    
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
// 本接口（DescribeDBInstanceConfig）用于查询云数据库实例的配置信息，包括同步模式，部署模式等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceConfig(request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    return c.DescribeDBInstanceConfigWithContext(context.Background(), request)
}

// DescribeDBInstanceConfig
// 本接口（DescribeDBInstanceConfig）用于查询云数据库实例的配置信息，包括同步模式，部署模式等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceConfigWithContext(ctx context.Context, request *DescribeDBInstanceConfigRequest) (response *DescribeDBInstanceConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstanceConfig")
    
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDBInstanceGTIDWithContext(ctx context.Context, request *DescribeDBInstanceGTIDRequest) (response *DescribeDBInstanceGTIDResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceGTIDRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstanceGTID")
    
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
// 本接口（DescribeDBInstanceInfo）用于查询实例基本信息（实例 ID，实例名称，是否开通加密），只读实例不支持查询。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) DescribeDBInstanceInfo(request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    return c.DescribeDBInstanceInfoWithContext(context.Background(), request)
}

// DescribeDBInstanceInfo
// 本接口（DescribeDBInstanceInfo）用于查询实例基本信息（实例 ID，实例名称，是否开通加密），只读实例不支持查询。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_RESOURCENOTEXISTS = "InvalidParameter.ResourceNotExists"
func (c *Client) DescribeDBInstanceInfoWithContext(ctx context.Context, request *DescribeDBInstanceInfoRequest) (response *DescribeDBInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstanceInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBInstanceLogToCLSRequest() (request *DescribeDBInstanceLogToCLSRequest) {
    request = &DescribeDBInstanceLogToCLSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeDBInstanceLogToCLS")
    
    
    return
}

func NewDescribeDBInstanceLogToCLSResponse() (response *DescribeDBInstanceLogToCLSResponse) {
    response = &DescribeDBInstanceLogToCLSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBInstanceLogToCLS
// 本接口（DescribeDBInstanceLogToCLS）用于查询实例慢日志、错误日志投递CLS的配置，通过AppId、Region以及实例ID过滤出当前实例日志投递CLS的配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) DescribeDBInstanceLogToCLS(request *DescribeDBInstanceLogToCLSRequest) (response *DescribeDBInstanceLogToCLSResponse, err error) {
    return c.DescribeDBInstanceLogToCLSWithContext(context.Background(), request)
}

// DescribeDBInstanceLogToCLS
// 本接口（DescribeDBInstanceLogToCLS）用于查询实例慢日志、错误日志投递CLS的配置，通过AppId、Region以及实例ID过滤出当前实例日志投递CLS的配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TIMEOUTERROR = "FailedOperation.TimeoutError"
//  INTERNALERROR_CAUTHERROR = "InternalError.CauthError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  OPERATIONDENIED_RESOURCENOTFUNDERROR = "OperationDenied.ResourceNotFundError"
func (c *Client) DescribeDBInstanceLogToCLSWithContext(ctx context.Context, request *DescribeDBInstanceLogToCLSRequest) (response *DescribeDBInstanceLogToCLSResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceLogToCLSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstanceLogToCLS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBInstanceLogToCLS require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBInstanceLogToCLSResponse()
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
// 本接口（DescribeDBInstanceRebootTime）用于查询云数据库实例重启预计所需的时间。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceRebootTime(request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    return c.DescribeDBInstanceRebootTimeWithContext(context.Background(), request)
}

// DescribeDBInstanceRebootTime
// 本接口（DescribeDBInstanceRebootTime）用于查询云数据库实例重启预计所需的时间。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_OSSERROR = "InternalError.OssError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeDBInstanceRebootTimeWithContext(ctx context.Context, request *DescribeDBInstanceRebootTimeRequest) (response *DescribeDBInstanceRebootTimeResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstanceRebootTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstanceRebootTime")
    
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
// 本接口（DescribeDBInstances）用于查询云数据库实例列表，支持通过项目 ID、实例 ID、访问地址、实例状态等过滤条件来筛选实例。支持查询主实例、灾备实例和只读实例信息列表。
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    return c.DescribeDBInstancesWithContext(context.Background(), request)
}

// DescribeDBInstances
// 本接口（DescribeDBInstances）用于查询云数据库实例列表，支持通过项目 ID、实例 ID、访问地址、实例状态等过滤条件来筛选实例。支持查询主实例、灾备实例和只读实例信息列表。
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBInstances")
    
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
// 本接口（DescribeDBPrice）用于查询购买或续费云数据库实例的价格，支持查询按量计费或者包年包月的价格。可传入实例类型、购买时长、购买数量、内存大小、硬盘大小和可用区信息等来查询实例价格。可传入实例名称来查询实例续费价格。
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
// 本接口（DescribeDBPrice）用于查询购买或续费云数据库实例的价格，支持查询按量计费或者包年包月的价格。可传入实例类型、购买时长、购买数量、内存大小、硬盘大小和可用区信息等来查询实例价格。可传入实例名称来查询实例续费价格。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBPrice")
    
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED = "OperationDenied"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBSecurityGroupsWithContext(ctx context.Context, request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBSecurityGroups")
    
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeDBSwitchRecordsWithContext(ctx context.Context, request *DescribeDBSwitchRecordsRequest) (response *DescribeDBSwitchRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDBSwitchRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDBSwitchRecords")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDataBackupOverview")
    
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DBOPERATIONERROR = "InternalError.DBOperationError"
//  INTERNALERROR_EXESQLERROR = "InternalError.ExeSqlError"
//  INTERNALERROR_INTERNALSERVICEERRORERR = "InternalError.InternalServiceErrorErr"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest) (response *DescribeDatabasesResponse, err error) {
    if request == nil {
        request = NewDescribeDatabasesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDatabases")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDefaultParams")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDeployGroupList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeDeviceMonitorInfo")
    
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeErrorLogDataWithContext(ctx context.Context, request *DescribeErrorLogDataRequest) (response *DescribeErrorLogDataResponse, err error) {
    if request == nil {
        request = NewDescribeErrorLogDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeErrorLogData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeErrorLogData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeErrorLogDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAlarmEventsRequest() (request *DescribeInstanceAlarmEventsRequest) {
    request = &DescribeInstanceAlarmEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceAlarmEvents")
    
    
    return
}

func NewDescribeInstanceAlarmEventsResponse() (response *DescribeInstanceAlarmEventsResponse) {
    response = &DescribeInstanceAlarmEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceAlarmEvents
// 本接口（DescribeInstanceAlarmEvents）用于查询实例发生的事件信息。
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeInstanceAlarmEvents(request *DescribeInstanceAlarmEventsRequest) (response *DescribeInstanceAlarmEventsResponse, err error) {
    return c.DescribeInstanceAlarmEventsWithContext(context.Background(), request)
}

// DescribeInstanceAlarmEvents
// 本接口（DescribeInstanceAlarmEvents）用于查询实例发生的事件信息。
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESULTOVERLIMIT = "OperationDenied.ResultOverLimit"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeInstanceAlarmEventsWithContext(ctx context.Context, request *DescribeInstanceAlarmEventsRequest) (response *DescribeInstanceAlarmEventsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAlarmEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeInstanceAlarmEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAlarmEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAlarmEventsResponse()
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENAMENOTFOUND = "InvalidParameter.InstanceNameNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeInstanceParamRecordsWithContext(ctx context.Context, request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeInstanceParamRecords")
    
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeInstanceParamsWithContext(ctx context.Context, request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceParamsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeInstanceParams")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceUpgradeCheckJobRequest() (request *DescribeInstanceUpgradeCheckJobRequest) {
    request = &DescribeInstanceUpgradeCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceUpgradeCheckJob")
    
    
    return
}

func NewDescribeInstanceUpgradeCheckJobResponse() (response *DescribeInstanceUpgradeCheckJobResponse) {
    response = &DescribeInstanceUpgradeCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceUpgradeCheckJob
// 该接口（DescribeInstanceUpgradeCheckJob）查询实例版本升级校验任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeInstanceUpgradeCheckJob(request *DescribeInstanceUpgradeCheckJobRequest) (response *DescribeInstanceUpgradeCheckJobResponse, err error) {
    return c.DescribeInstanceUpgradeCheckJobWithContext(context.Background(), request)
}

// DescribeInstanceUpgradeCheckJob
// 该接口（DescribeInstanceUpgradeCheckJob）查询实例版本升级校验任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeInstanceUpgradeCheckJobWithContext(ctx context.Context, request *DescribeInstanceUpgradeCheckJobRequest) (response *DescribeInstanceUpgradeCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceUpgradeCheckJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeInstanceUpgradeCheckJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceUpgradeCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceUpgradeCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceUpgradeTypeRequest() (request *DescribeInstanceUpgradeTypeRequest) {
    request = &DescribeInstanceUpgradeTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeInstanceUpgradeType")
    
    
    return
}

func NewDescribeInstanceUpgradeTypeResponse() (response *DescribeInstanceUpgradeTypeResponse) {
    response = &DescribeInstanceUpgradeTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceUpgradeType
// 本接口（DescribeInstanceUpgradeType）用于查询数据库实例升级类型。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeInstanceUpgradeType(request *DescribeInstanceUpgradeTypeRequest) (response *DescribeInstanceUpgradeTypeResponse, err error) {
    return c.DescribeInstanceUpgradeTypeWithContext(context.Background(), request)
}

// DescribeInstanceUpgradeType
// 本接口（DescribeInstanceUpgradeType）用于查询数据库实例升级类型。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeInstanceUpgradeTypeWithContext(ctx context.Context, request *DescribeInstanceUpgradeTypeRequest) (response *DescribeInstanceUpgradeTypeResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceUpgradeTypeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeInstanceUpgradeType")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceUpgradeType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceUpgradeTypeResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeLocalBinlogConfig")
    
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
// 该接口（DescribeParamTemplateInfo）用于查询参数模板详情。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
// 该接口（DescribeParamTemplateInfo）用于查询参数模板详情。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeParamTemplateInfo")
    
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
// 该接口（DescribeParamTemplates）查询参数模板列表。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// 该接口（DescribeParamTemplates）查询参数模板列表。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR_DATABASEERROR = "CdbError.DatabaseError"
//  INTERNALERROR_PARAMERROR = "InternalError.ParamError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
func (c *Client) DescribeParamTemplatesWithContext(ctx context.Context, request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeParamTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeParamTemplates")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeProjectSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupsResponse()
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeProxyCustomConf(request *DescribeProxyCustomConfRequest) (response *DescribeProxyCustomConfResponse, err error) {
    return c.DescribeProxyCustomConfWithContext(context.Background(), request)
}

// DescribeProxyCustomConf
// 查询代理规格配置
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) DescribeProxyCustomConfWithContext(ctx context.Context, request *DescribeProxyCustomConfRequest) (response *DescribeProxyCustomConfResponse, err error) {
    if request == nil {
        request = NewDescribeProxyCustomConfRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeProxyCustomConf")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyCustomConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyCustomConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxySupportParamRequest() (request *DescribeProxySupportParamRequest) {
    request = &DescribeProxySupportParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeProxySupportParam")
    
    
    return
}

func NewDescribeProxySupportParamResponse() (response *DescribeProxySupportParamResponse) {
    response = &DescribeProxySupportParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxySupportParam
// 本接口（DescribeProxySupportParam）用于查询实例支持代理版本和参数。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeProxySupportParam(request *DescribeProxySupportParamRequest) (response *DescribeProxySupportParamResponse, err error) {
    return c.DescribeProxySupportParamWithContext(context.Background(), request)
}

// DescribeProxySupportParam
// 本接口（DescribeProxySupportParam）用于查询实例支持代理版本和参数。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeProxySupportParamWithContext(ctx context.Context, request *DescribeProxySupportParamRequest) (response *DescribeProxySupportParamResponse, err error) {
    if request == nil {
        request = NewDescribeProxySupportParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeProxySupportParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySupportParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySupportParamResponse()
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeRemoteBackupConfig(request *DescribeRemoteBackupConfigRequest) (response *DescribeRemoteBackupConfigResponse, err error) {
    return c.DescribeRemoteBackupConfigWithContext(context.Background(), request)
}

// DescribeRemoteBackupConfig
// 本接口(DescribeRemoteBackupConfig)用于查询数据库异地备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) DescribeRemoteBackupConfigWithContext(ctx context.Context, request *DescribeRemoteBackupConfigRequest) (response *DescribeRemoteBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeRemoteBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeRemoteBackupConfig")
    
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
// 本接口（DescribeRoGroups）用于查询云数据库实例的所有的 RO 组的信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeRoGroups(request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    return c.DescribeRoGroupsWithContext(context.Background(), request)
}

// DescribeRoGroups
// 本接口（DescribeRoGroups）用于查询云数据库实例的所有的 RO 组的信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
func (c *Client) DescribeRoGroupsWithContext(ctx context.Context, request *DescribeRoGroupsRequest) (response *DescribeRoGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRoGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeRoGroups")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeRoMinScale")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeRollbackRangeTime")
    
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
// 本接口（DescribeRollbackTaskDetail）用于查询云数据库实例回档任务详情。
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
// 本接口（DescribeRollbackTaskDetail）用于查询云数据库实例回档任务详情。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeRollbackTaskDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRollbackTaskDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRollbackTaskDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSSLStatusRequest() (request *DescribeSSLStatusRequest) {
    request = &DescribeSSLStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeSSLStatus")
    
    
    return
}

func NewDescribeSSLStatusResponse() (response *DescribeSSLStatusResponse) {
    response = &DescribeSSLStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSSLStatus
// 本接口（DescribeSSLStatus）用于查询 SSL 开通情况。如果已经开通 SSL ，会同步返回证书下载链接。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSSLStatus(request *DescribeSSLStatusRequest) (response *DescribeSSLStatusResponse, err error) {
    return c.DescribeSSLStatusWithContext(context.Background(), request)
}

// DescribeSSLStatus
// 本接口（DescribeSSLStatus）用于查询 SSL 开通情况。如果已经开通 SSL ，会同步返回证书下载链接。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) DescribeSSLStatusWithContext(ctx context.Context, request *DescribeSSLStatusRequest) (response *DescribeSSLStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSSLStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeSSLStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSSLStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSSLStatusResponse()
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
// 本接口（DescribeSlowLogData）用于使用条件检索实例的慢日志。只允许查看一个月之内的慢日志。
//
// 使用时需要注意：可能存在单条慢日志太大，导致整个http请求的回包太大，进而引发接口超时。一旦发生超时，建议您缩小查询时的Limit参数值，从而降低包的大小，让接口能够及时返回内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGCONTENTOVERLIMIT = "FailedOperation.LogContentOverLimit"
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_RESULTSETOVERLIMIT = "FailedOperation.ResultSetOverLimit"
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
// 本接口（DescribeSlowLogData）用于使用条件检索实例的慢日志。只允许查看一个月之内的慢日志。
//
// 使用时需要注意：可能存在单条慢日志太大，导致整个http请求的回包太大，进而引发接口超时。一旦发生超时，建议您缩小查询时的Limit参数值，从而降低包的大小，让接口能够及时返回内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_LOGCONTENTOVERLIMIT = "FailedOperation.LogContentOverLimit"
//  FAILEDOPERATION_QUERYLOGERROR = "FailedOperation.QueryLogError"
//  FAILEDOPERATION_RESULTSETOVERLIMIT = "FailedOperation.ResultSetOverLimit"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeSlowLogData")
    
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
// 本接口（DescribeSlowLogs）用于获取云数据库实例的慢查询日志。
//
// 说明：若单次查询数据量过大，则有可能响应超时，建议缩短单次查询时间范围，如一小时，避免导致超时。
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
// 本接口（DescribeSlowLogs）用于获取云数据库实例的慢查询日志。
//
// 说明：若单次查询数据量过大，则有可能响应超时，建议缩短单次查询时间范围，如一小时，避免导致超时。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeSlowLogs")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeSupportedPrivileges")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSupportedPrivileges require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSupportedPrivilegesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTableColumnsRequest() (request *DescribeTableColumnsRequest) {
    request = &DescribeTableColumnsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "DescribeTableColumns")
    
    
    return
}

func NewDescribeTableColumnsResponse() (response *DescribeTableColumnsResponse) {
    response = &DescribeTableColumnsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTableColumns
// 本接口（DescribeTableColumns）用于查询云数据库实例的指定数据库表的列信息，仅支持主实例和灾备实例。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeTableColumns(request *DescribeTableColumnsRequest) (response *DescribeTableColumnsResponse, err error) {
    return c.DescribeTableColumnsWithContext(context.Background(), request)
}

// DescribeTableColumns
// 本接口（DescribeTableColumns）用于查询云数据库实例的指定数据库表的列信息，仅支持主实例和灾备实例。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeTableColumnsWithContext(ctx context.Context, request *DescribeTableColumnsRequest) (response *DescribeTableColumnsResponse, err error) {
    if request == nil {
        request = NewDescribeTableColumnsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeTableColumns")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTableColumns require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTableColumnsResponse()
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
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
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
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeTables")
    
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
// 本接口（DescribeTagsOfInstanceIds）用于获取云数据库实例的标签信息。
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
// 本接口（DescribeTagsOfInstanceIds）用于获取云数据库实例的标签信息。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeTagsOfInstanceIds")
    
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
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
//  INTERNALERROR_CDBERROR = "InternalError.CdbError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_DESERROR = "InternalError.DesError"
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest) (response *DescribeTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeTasks")
    
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) DescribeTimeWindowWithContext(ctx context.Context, request *DescribeTimeWindowRequest) (response *DescribeTimeWindowResponse, err error) {
    if request == nil {
        request = NewDescribeTimeWindowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeTimeWindow")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DescribeUploadedFiles")
    
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
func (c *Client) DisassociateSecurityGroupsWithContext(ctx context.Context, request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewDisassociateSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "DisassociateSecurityGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisassociateSecurityGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisassociateSecurityGroupsResponse()
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
// 本接口（InquiryPriceUpgradeInstances）用于查询云数据库实例升级的价格，支持查询按量计费或者包年包月实例的升级价格，实例类型支持主实例、灾备实例和只读实例。
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
// 本接口（InquiryPriceUpgradeInstances）用于查询云数据库实例升级的价格，支持查询按量计费或者包年包月实例的升级价格，实例类型支持主实例、灾备实例和只读实例。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "InquiryPriceUpgradeInstances")
    
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
// 本接口（IsolateDBInstance）用于隔离云数据库实例，隔离后不能通过IP和端口访问数据库。隔离的实例可在回收站中进行开机。若为欠费隔离，请尽快进行充值。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDBINSTANCELOCKFAILERROR = "FailedOperation.CdbInstanceLockFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTREFUNDERROR = "OperationDenied.UnSupportRefundError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) IsolateDBInstance(request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    return c.IsolateDBInstanceWithContext(context.Background(), request)
}

// IsolateDBInstance
// 本接口（IsolateDBInstance）用于隔离云数据库实例，隔离后不能通过IP和端口访问数据库。隔离的实例可在回收站中进行开机。若为欠费隔离，请尽快进行充值。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CDBINSTANCELOCKFAILERROR = "FailedOperation.CdbInstanceLockFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTREFUNDERROR = "OperationDenied.UnSupportRefundError"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) IsolateDBInstanceWithContext(ctx context.Context, request *IsolateDBInstanceRequest) (response *IsolateDBInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateDBInstanceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "IsolateDBInstance")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAccountDescription")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAccountHost")
    
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
// 本接口（ModifyAccountMaxUserConnections）用于修改云数据库账户最大可用连接数。
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
// 本接口（ModifyAccountMaxUserConnections）用于修改云数据库账户最大可用连接数。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAccountMaxUserConnections")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAccountPassword")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAccountPrivileges")
    
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
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
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
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAuditConfig")
    
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
// 不再支持审计规则创建
//
// 
//
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
// 不再支持审计规则创建
//
// 
//
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAuditRule")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditRuleResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditRuleTemplatesRequest() (request *ModifyAuditRuleTemplatesRequest) {
    request = &ModifyAuditRuleTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAuditRuleTemplates")
    
    
    return
}

func NewModifyAuditRuleTemplatesResponse() (response *ModifyAuditRuleTemplatesResponse) {
    response = &ModifyAuditRuleTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditRuleTemplates
// 修改审计规则模板
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyAuditRuleTemplates(request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    return c.ModifyAuditRuleTemplatesWithContext(context.Background(), request)
}

// ModifyAuditRuleTemplates
// 修改审计规则模板
//
// 可能返回的错误码:
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) ModifyAuditRuleTemplatesWithContext(ctx context.Context, request *ModifyAuditRuleTemplatesRequest) (response *ModifyAuditRuleTemplatesResponse, err error) {
    if request == nil {
        request = NewModifyAuditRuleTemplatesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAuditRuleTemplates")
    
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
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyAuditService")
    
    
    return
}

func NewModifyAuditServiceResponse() (response *ModifyAuditServiceResponse) {
    response = &ModifyAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计日志保存时长、审计规则等服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_MODIFYAUDITSTATUSERROR = "OperationDenied.ModifyAuditStatusError"
func (c *Client) ModifyAuditService(request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    return c.ModifyAuditServiceWithContext(context.Background(), request)
}

// ModifyAuditService
// 本接口(ModifyAuditService)用于修改云数据库审计日志保存时长、审计规则等服务配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDITCONFIGNOTEXIST = "FailedOperation.AuditConfigNotExist"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_AUDITPOLICYNOTEXISTERROR = "OperationDenied.AuditPolicyNotExistError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_MODIFYAUDITSTATUSERROR = "OperationDenied.ModifyAuditStatusError"
func (c *Client) ModifyAuditServiceWithContext(ctx context.Context, request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    if request == nil {
        request = NewModifyAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditServiceResponse()
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
// 本接口（ModifyAutoRenewFlag）用于修改云数据库实例的自动续费标记。仅支持包年包月的实例设置自动续费标记。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// 本接口（ModifyAutoRenewFlag）用于修改云数据库实例的自动续费标记。仅支持包年包月的实例设置自动续费标记。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyAutoRenewFlag")
    
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
// 本接口（ModifyBackupConfig）用于修改数据库备份配置信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyBackupConfig(request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    return c.ModifyBackupConfigWithContext(context.Background(), request)
}

// ModifyBackupConfig
// 本接口（ModifyBackupConfig）用于修改数据库备份配置信息。
//
// 可能返回的错误码:
//  CDBERROR = "CdbError"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyBackupConfigWithContext(ctx context.Context, request *ModifyBackupConfigRequest) (response *ModifyBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyBackupConfig")
    
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
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
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
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyBackupDownloadRestrictionWithContext(ctx context.Context, request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    if request == nil {
        request = NewModifyBackupDownloadRestrictionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyBackupDownloadRestriction")
    
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
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
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
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyBackupEncryptionStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyBackupEncryptionStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyBackupEncryptionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCdbProxyAddressDescRequest() (request *ModifyCdbProxyAddressDescRequest) {
    request = &ModifyCdbProxyAddressDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCdbProxyAddressDesc")
    
    
    return
}

func NewModifyCdbProxyAddressDescResponse() (response *ModifyCdbProxyAddressDescResponse) {
    response = &ModifyCdbProxyAddressDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCdbProxyAddressDesc
// 本接口（ModifyCdbProxyAddressDesc）用于修改代理地址描述信息。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCdbProxyAddressDesc(request *ModifyCdbProxyAddressDescRequest) (response *ModifyCdbProxyAddressDescResponse, err error) {
    return c.ModifyCdbProxyAddressDescWithContext(context.Background(), request)
}

// ModifyCdbProxyAddressDesc
// 本接口（ModifyCdbProxyAddressDesc）用于修改代理地址描述信息。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ModifyCdbProxyAddressDescWithContext(ctx context.Context, request *ModifyCdbProxyAddressDescRequest) (response *ModifyCdbProxyAddressDescResponse, err error) {
    if request == nil {
        request = NewModifyCdbProxyAddressDescRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyCdbProxyAddressDesc")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCdbProxyAddressDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCdbProxyAddressDescResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCdbProxyAddressVipAndVPortRequest() (request *ModifyCdbProxyAddressVipAndVPortRequest) {
    request = &ModifyCdbProxyAddressVipAndVPortRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCdbProxyAddressVipAndVPort")
    
    
    return
}

func NewModifyCdbProxyAddressVipAndVPortResponse() (response *ModifyCdbProxyAddressVipAndVPortResponse) {
    response = &ModifyCdbProxyAddressVipAndVPortResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCdbProxyAddressVipAndVPort
// 本接口（ModifyCdbProxyAddressVipAndVPort）用于修改数据库代理地址VPC信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_VPCIPINUSEERROR = "FailedOperation.VpcIpInUseError"
//  FAILEDOPERATION_VPCIPNOTINSUBNETERROR = "FailedOperation.VpcIpNotInSubnetError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCdbProxyAddressVipAndVPort(request *ModifyCdbProxyAddressVipAndVPortRequest) (response *ModifyCdbProxyAddressVipAndVPortResponse, err error) {
    return c.ModifyCdbProxyAddressVipAndVPortWithContext(context.Background(), request)
}

// ModifyCdbProxyAddressVipAndVPort
// 本接口（ModifyCdbProxyAddressVipAndVPort）用于修改数据库代理地址VPC信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  FAILEDOPERATION_VPCIPINUSEERROR = "FailedOperation.VpcIpInUseError"
//  FAILEDOPERATION_VPCIPNOTINSUBNETERROR = "FailedOperation.VpcIpNotInSubnetError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCdbProxyAddressVipAndVPortWithContext(ctx context.Context, request *ModifyCdbProxyAddressVipAndVPortRequest) (response *ModifyCdbProxyAddressVipAndVPortResponse, err error) {
    if request == nil {
        request = NewModifyCdbProxyAddressVipAndVPortRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyCdbProxyAddressVipAndVPort")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCdbProxyAddressVipAndVPort require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCdbProxyAddressVipAndVPortResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCdbProxyParamRequest() (request *ModifyCdbProxyParamRequest) {
    request = &ModifyCdbProxyParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyCdbProxyParam")
    
    
    return
}

func NewModifyCdbProxyParamResponse() (response *ModifyCdbProxyParamResponse) {
    response = &ModifyCdbProxyParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCdbProxyParam
// 本接口（ModifyCdbProxyParam）用于配置数据库代理参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCdbProxyParam(request *ModifyCdbProxyParamRequest) (response *ModifyCdbProxyParamResponse, err error) {
    return c.ModifyCdbProxyParamWithContext(context.Background(), request)
}

// ModifyCdbProxyParam
// 本接口（ModifyCdbProxyParam）用于配置数据库代理参数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyCdbProxyParamWithContext(ctx context.Context, request *ModifyCdbProxyParamRequest) (response *ModifyCdbProxyParamResponse, err error) {
    if request == nil {
        request = NewModifyCdbProxyParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyCdbProxyParam")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCdbProxyParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCdbProxyParamResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceLogToCLSRequest() (request *ModifyDBInstanceLogToCLSRequest) {
    request = &ModifyDBInstanceLogToCLSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceLogToCLS")
    
    
    return
}

func NewModifyDBInstanceLogToCLSResponse() (response *ModifyDBInstanceLogToCLSResponse) {
    response = &ModifyDBInstanceLogToCLSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceLogToCLS
// 开启/关闭CDB慢日志、错误日志投递CLS
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) ModifyDBInstanceLogToCLS(request *ModifyDBInstanceLogToCLSRequest) (response *ModifyDBInstanceLogToCLSResponse, err error) {
    return c.ModifyDBInstanceLogToCLSWithContext(context.Background(), request)
}

// ModifyDBInstanceLogToCLS
// 开启/关闭CDB慢日志、错误日志投递CLS
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETER_JSONUNMARSHALERROR = "InvalidParameter.JsonUnmarshalError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
func (c *Client) ModifyDBInstanceLogToCLSWithContext(ctx context.Context, request *ModifyDBInstanceLogToCLSRequest) (response *ModifyDBInstanceLogToCLSResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceLogToCLSRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyDBInstanceLogToCLS")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceLogToCLS require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceLogToCLSResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyDBInstanceName")
    
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyDBInstanceProjectWithContext(ctx context.Context, request *ModifyDBInstanceProjectRequest) (response *ModifyDBInstanceProjectResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceProjectRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyDBInstanceProject")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceProjectResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceReadOnlyStatusRequest() (request *ModifyDBInstanceReadOnlyStatusRequest) {
    request = &ModifyDBInstanceReadOnlyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyDBInstanceReadOnlyStatus")
    
    
    return
}

func NewModifyDBInstanceReadOnlyStatusResponse() (response *ModifyDBInstanceReadOnlyStatusResponse) {
    response = &ModifyDBInstanceReadOnlyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDBInstanceReadOnlyStatus
// 本接口（ModifyDBInstanceReadOnlyStatus）用户设置MySQL云数据库实例为只读
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceReadOnlyStatus(request *ModifyDBInstanceReadOnlyStatusRequest) (response *ModifyDBInstanceReadOnlyStatusResponse, err error) {
    return c.ModifyDBInstanceReadOnlyStatusWithContext(context.Background(), request)
}

// ModifyDBInstanceReadOnlyStatus
// 本接口（ModifyDBInstanceReadOnlyStatus）用户设置MySQL云数据库实例为只读
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceReadOnlyStatusWithContext(ctx context.Context, request *ModifyDBInstanceReadOnlyStatusRequest) (response *ModifyDBInstanceReadOnlyStatusResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceReadOnlyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyDBInstanceReadOnlyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDBInstanceReadOnlyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDBInstanceReadOnlyStatusResponse()
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
func (c *Client) ModifyDBInstanceSecurityGroupsWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceSecurityGroupsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
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
// 本接口（ModifyDBInstanceVipVport）用于修改云数据库实例的IP和端口号，也可进行基础网络转 VPC 网络和 VPC 网络下的子网变更。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceVipVport(request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    return c.ModifyDBInstanceVipVportWithContext(context.Background(), request)
}

// ModifyDBInstanceVipVport
// 本接口（ModifyDBInstanceVipVport）用于修改云数据库实例的IP和端口号，也可进行基础网络转 VPC 网络和 VPC 网络下的子网变更。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_STATUSCONFLICT = "FailedOperation.StatusConflict"
//  INTERNALERROR_JSONERROR = "InternalError.JSONError"
//  INTERNALERROR_NETWORKERROR = "InternalError.NetworkError"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyDBInstanceVipVportWithContext(ctx context.Context, request *ModifyDBInstanceVipVportRequest) (response *ModifyDBInstanceVipVportResponse, err error) {
    if request == nil {
        request = NewModifyDBInstanceVipVportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyDBInstanceVipVport")
    
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
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
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
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
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
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
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
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) ModifyInstanceParamWithContext(ctx context.Context, request *ModifyInstanceParamRequest) (response *ModifyInstanceParamResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyInstanceParam")
    
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
// 本接口（ModifyInstancePasswordComplexity）用于修改云数据库实例的密码复杂度。
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
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstancePasswordComplexity(request *ModifyInstancePasswordComplexityRequest) (response *ModifyInstancePasswordComplexityResponse, err error) {
    return c.ModifyInstancePasswordComplexityWithContext(context.Background(), request)
}

// ModifyInstancePasswordComplexity
// 本接口（ModifyInstancePasswordComplexity）用于修改云数据库实例的密码复杂度。
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
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyInstancePasswordComplexityWithContext(ctx context.Context, request *ModifyInstancePasswordComplexityRequest) (response *ModifyInstancePasswordComplexityResponse, err error) {
    if request == nil {
        request = NewModifyInstancePasswordComplexityRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyInstancePasswordComplexity")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyInstanceTag")
    
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
// 本接口（ModifyLocalBinlogConfig）用于修改实例本地 binlog 保留策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
func (c *Client) ModifyLocalBinlogConfig(request *ModifyLocalBinlogConfigRequest) (response *ModifyLocalBinlogConfigResponse, err error) {
    return c.ModifyLocalBinlogConfigWithContext(context.Background(), request)
}

// ModifyLocalBinlogConfig
// 本接口（ModifyLocalBinlogConfig）用于修改实例本地 binlog 保留策略。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_INSTANCEUNSUPPORTEDOPERATEERROR = "OperationDenied.InstanceUnsupportedOperateError"
func (c *Client) ModifyLocalBinlogConfigWithContext(ctx context.Context, request *ModifyLocalBinlogConfigRequest) (response *ModifyLocalBinlogConfigResponse, err error) {
    if request == nil {
        request = NewModifyLocalBinlogConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyLocalBinlogConfig")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyNameOrDescByDpId")
    
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
// 该接口（ModifyParamTemplate）用于修改参数模板。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
// 该接口（ModifyParamTemplate）用于修改参数模板。
//
// 说明：参数模板为公共组件，配置完成后全地域生效。接口调用配置地域可选择广州、新加坡。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyParamTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyParamTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyParamTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyProtectModeRequest() (request *ModifyProtectModeRequest) {
    request = &ModifyProtectModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyProtectMode")
    
    
    return
}

func NewModifyProtectModeResponse() (response *ModifyProtectModeResponse) {
    response = &ModifyProtectModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyProtectMode
// 该接口（ModifyProtectMode）用于修改实例的同步方式。
//
// 说明：仅专属集群可调用，该接口即将下线。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyProtectMode(request *ModifyProtectModeRequest) (response *ModifyProtectModeResponse, err error) {
    return c.ModifyProtectModeWithContext(context.Background(), request)
}

// ModifyProtectMode
// 该接口（ModifyProtectMode）用于修改实例的同步方式。
//
// 说明：仅专属集群可调用，该接口即将下线。
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONINPROCESS = "OperationDenied.ActionInProcess"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ModifyProtectModeWithContext(ctx context.Context, request *ModifyProtectModeRequest) (response *ModifyProtectModeResponse, err error) {
    if request == nil {
        request = NewModifyProtectModeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyProtectMode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyProtectMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyProtectModeResponse()
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
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyRemoteBackupConfig(request *ModifyRemoteBackupConfigRequest) (response *ModifyRemoteBackupConfigResponse, err error) {
    return c.ModifyRemoteBackupConfigWithContext(context.Background(), request)
}

// ModifyRemoteBackupConfig
// 本接口(ModifyRemoteBackupConfig)用于修改数据库异地备份配置信息。
//
// 可能返回的错误码:
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) ModifyRemoteBackupConfigWithContext(ctx context.Context, request *ModifyRemoteBackupConfigRequest) (response *ModifyRemoteBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyRemoteBackupConfigRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyRemoteBackupConfig")
    
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
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
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
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) ModifyRoGroupInfoWithContext(ctx context.Context, request *ModifyRoGroupInfoRequest) (response *ModifyRoGroupInfoResponse, err error) {
    if request == nil {
        request = NewModifyRoGroupInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyRoGroupInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoGroupVipVportRequest() (request *ModifyRoGroupVipVportRequest) {
    request = &ModifyRoGroupVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ModifyRoGroupVipVport")
    
    
    return
}

func NewModifyRoGroupVipVportResponse() (response *ModifyRoGroupVipVportResponse) {
    response = &ModifyRoGroupVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRoGroupVipVport
// 该接口（ModifyRoGroupVipVport）用于修改Ro组的vip和vport。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyRoGroupVipVport(request *ModifyRoGroupVipVportRequest) (response *ModifyRoGroupVipVportResponse, err error) {
    return c.ModifyRoGroupVipVportWithContext(context.Background(), request)
}

// ModifyRoGroupVipVport
// 该接口（ModifyRoGroupVipVport）用于修改Ro组的vip和vport。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR_VPCERROR = "InternalError.VpcError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyRoGroupVipVportWithContext(ctx context.Context, request *ModifyRoGroupVipVportRequest) (response *ModifyRoGroupVipVportResponse, err error) {
    if request == nil {
        request = NewModifyRoGroupVipVportRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyRoGroupVipVport")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRoGroupVipVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRoGroupVipVportResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ModifyTimeWindow")
    
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
// 该接口为异步操作，部分资源的回收可能存在延迟。您可以通过使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口，指定实例 InstanceId 和状态 Status 为 [5,6,7] 进行查询，其中5代表已隔离，6代表下线中，7代表已下线。若返回实例为空，则实例资源已全部释放。
//
// 
//
// 注意，实例下线后，相关资源和数据将无法找回，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
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
// 该接口为异步操作，部分资源的回收可能存在延迟。您可以通过使用 [查询实例列表](https://cloud.tencent.com/document/api/236/15872) 接口，指定实例 InstanceId 和状态 Status 为 [5,6,7] 进行查询，其中5代表已隔离，6代表下线中，7代表已下线。若返回实例为空，则实例资源已全部释放。
//
// 
//
// 注意，实例下线后，相关资源和数据将无法找回，请谨慎操作。
//
// 可能返回的错误码:
//  INTERNALERROR_FTPERROR = "InternalError.FtpError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) OfflineIsolatedInstancesWithContext(ctx context.Context, request *OfflineIsolatedInstancesRequest) (response *OfflineIsolatedInstancesResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "OfflineIsolatedInstances")
    
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
// 本接口（OpenAuditService）用 CDB 实例开通审计服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// 本接口（OpenAuditService）用 CDB 实例开通审计服务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEAUDITFAILERROR = "FailedOperation.CreateAuditFailError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_AUDITSTATUSERROR = "OperationDenied.AuditStatusError"
//  OPERATIONDENIED_AUDITTASKCONFLICTERROR = "OperationDenied.AuditTaskConflictError"
//  OPERATIONDENIED_INSTANCESTATUSERROR = "OperationDenied.InstanceStatusError"
//  OPERATIONDENIED_UNSUPPORTOPENAUDITERROR = "OperationDenied.UnsupportOpenAuditError"
//  RESOURCENOTFOUND_INSTANCENOTFUNDERROR = "ResourceNotFound.InstanceNotFundError"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "OpenAuditService")
    
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
// 本接口（OpenDBInstanceEncryption）用于启用实例数据存储加密功能，支持用户指定自定义密钥。
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
// 4、开启加密后不允许关闭。
//
// 
//
// 该 API 耗时可能到10s，客户端可能超时，如果调用 API 返回 InternalError ，请您调用 [DescribeDBInstanceInfo](https://cloud.tencent.com/document/product/236/44160) 确认后端加密是否开通成功，调用后参数 Encryption 为 YES 表示已开通成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_HTTPREQUESTERROR = "InternalError.HttpRequestError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) OpenDBInstanceEncryption(request *OpenDBInstanceEncryptionRequest) (response *OpenDBInstanceEncryptionResponse, err error) {
    return c.OpenDBInstanceEncryptionWithContext(context.Background(), request)
}

// OpenDBInstanceEncryption
// 本接口（OpenDBInstanceEncryption）用于启用实例数据存储加密功能，支持用户指定自定义密钥。
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
// 4、开启加密后不允许关闭。
//
// 
//
// 该 API 耗时可能到10s，客户端可能超时，如果调用 API 返回 InternalError ，请您调用 [DescribeDBInstanceInfo](https://cloud.tencent.com/document/product/236/44160) 确认后端加密是否开通成功，调用后参数 Encryption 为 YES 表示已开通成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  CDBERROR = "CdbError"
//  INTERNALERROR_HTTPREQUESTERROR = "InternalError.HttpRequestError"
//  INTERNALERROR_KMSERROR = "InternalError.KmsError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) OpenDBInstanceEncryptionWithContext(ctx context.Context, request *OpenDBInstanceEncryptionRequest) (response *OpenDBInstanceEncryptionResponse, err error) {
    if request == nil {
        request = NewOpenDBInstanceEncryptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "OpenDBInstanceEncryption")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "OpenDBInstanceGTID")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenDBInstanceGTID require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenDBInstanceGTIDResponse()
    err = c.Send(request, response)
    return
}

func NewOpenSSLRequest() (request *OpenSSLRequest) {
    request = &OpenSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "OpenSSL")
    
    
    return
}

func NewOpenSSLResponse() (response *OpenSSLResponse) {
    response = &OpenSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenSSL
// 本接口（OpenSSL）用于开启 SSL 连接功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_OPERATIONINCONFLICTERROR = "FailedOperation.OperationInConflictError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) OpenSSL(request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    return c.OpenSSLWithContext(context.Background(), request)
}

// OpenSSL
// 本接口（OpenSSL）用于开启 SSL 连接功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONACTIONERROR = "FailedOperation.DBOperationActionError"
//  FAILEDOPERATION_OPERATIONINCONFLICTERROR = "FailedOperation.OperationInConflictError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  OPERATIONDENIED_INSTANCELOCKERCONFLICT = "OperationDenied.InstanceLockerConflict"
func (c *Client) OpenSSLWithContext(ctx context.Context, request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    if request == nil {
        request = NewOpenSSLRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "OpenSSL")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenSSLResponse()
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
func (c *Client) OpenWanServiceWithContext(ctx context.Context, request *OpenWanServiceRequest) (response *OpenWanServiceResponse, err error) {
    if request == nil {
        request = NewOpenWanServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "OpenWanService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenWanService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenWanServiceResponse()
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
// 本接口（ReleaseIsolatedDBInstances）用于恢复已隔离云数据库实例。仅用于按量计费实例的解隔离，包年包月实例的解隔离请使用 RenewDBInstance 。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONINCONFLICTERROR = "FailedOperation.OperationInConflictError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
// 本接口（ReleaseIsolatedDBInstances）用于恢复已隔离云数据库实例。仅用于按量计费实例的解隔离，包年包月实例的解隔离请使用 RenewDBInstance 。
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATIONINCONFLICTERROR = "FailedOperation.OperationInConflictError"
//  INTERNALERROR_DATABASEACCESSERROR = "InternalError.DatabaseAccessError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ReleaseIsolatedDBInstances")
    
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ReloadBalanceProxyNode(request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    return c.ReloadBalanceProxyNodeWithContext(context.Background(), request)
}

// ReloadBalanceProxyNode
// 重新负载均衡数据库代理
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_HTTPERROR = "InternalError.HttpError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ReloadBalanceProxyNodeWithContext(ctx context.Context, request *ReloadBalanceProxyNodeRequest) (response *ReloadBalanceProxyNodeResponse, err error) {
    if request == nil {
        request = NewReloadBalanceProxyNodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ReloadBalanceProxyNode")
    
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
// 本接口（RenewDBInstance）用于续费云数据库实例，支持付费模式为包年包月的实例。按量计费实例可通过该接口续费为包年包月的实例。
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
// 本接口（RenewDBInstance）用于续费云数据库实例，支持付费模式为包年包月的实例。按量计费实例可通过该接口续费为包年包月的实例。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "RenewDBInstance")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewDBInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewDBInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "ResetPassword")
    
    
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetPassword
// 手动刷新轮转密码
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

// ResetPassword
// 手动刷新轮转密码
//
// 可能返回的错误码:
//  INTERNALERROR_CDBCGWERROR = "InternalError.CdbCgwError"
//  INTERNALERROR_TRADEERROR = "InternalError.TradeError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ResetPassword")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetPasswordResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "ResetRootAccount")
    
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
// 本接口（RestartDBInstances）用于重启云数据库实例。
//
// 
//
// 注意：
//
// 1、本接口支持主实例、只读实例、灾备实例进行重启操作。
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
// 本接口（RestartDBInstances）用于重启云数据库实例。
//
// 
//
// 注意：
//
// 1、本接口支持主实例、只读实例、灾备实例进行重启操作。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "RestartDBInstances")
    
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
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
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_WRONGSTATUS = "OperationDenied.WrongStatus"
//  OVERQUOTA = "OverQuota"
func (c *Client) StartBatchRollbackWithContext(ctx context.Context, request *StartBatchRollbackRequest) (response *StartBatchRollbackResponse, err error) {
    if request == nil {
        request = NewStartBatchRollbackRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StartBatchRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartBatchRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartBatchRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewStartCpuExpandRequest() (request *StartCpuExpandRequest) {
    request = &StartCpuExpandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StartCpuExpand")
    
    
    return
}

func NewStartCpuExpandResponse() (response *StartCpuExpandResponse) {
    response = &StartCpuExpandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCpuExpand
// 通过该 API，可以开启 CPU 弹性扩容，包括一次性的手动扩容以及自动弹性扩容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCETASKCONFLICTERROR = "FailedOperation.InstanceTaskConflictError"
//  FAILEDOPERATION_NOTCHANGESTRATEGY = "FailedOperation.NotChangeStrategy"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  UNSUPPORTEDOPERATION_NOTSUPPORTNORMALINSTANCE = "UnsupportedOperation.NotSupportNormalInstance"
func (c *Client) StartCpuExpand(request *StartCpuExpandRequest) (response *StartCpuExpandResponse, err error) {
    return c.StartCpuExpandWithContext(context.Background(), request)
}

// StartCpuExpand
// 通过该 API，可以开启 CPU 弹性扩容，包括一次性的手动扩容以及自动弹性扩容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCETASKCONFLICTERROR = "FailedOperation.InstanceTaskConflictError"
//  FAILEDOPERATION_NOTCHANGESTRATEGY = "FailedOperation.NotChangeStrategy"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  MISSINGPARAMETER_MISSINGPARAMERROR = "MissingParameter.MissingParamError"
//  UNSUPPORTEDOPERATION_NOTSUPPORTNORMALINSTANCE = "UnsupportedOperation.NotSupportNormalInstance"
func (c *Client) StartCpuExpandWithContext(ctx context.Context, request *StartCpuExpandRequest) (response *StartCpuExpandResponse, err error) {
    if request == nil {
        request = NewStartCpuExpandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StartCpuExpand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCpuExpand require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCpuExpandResponse()
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
// 本接口（StartReplication）用于开启 RO 复制，从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) StartReplication(request *StartReplicationRequest) (response *StartReplicationResponse, err error) {
    return c.StartReplicationWithContext(context.Background(), request)
}

// StartReplication
// 本接口（StartReplication）用于开启 RO 复制，从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) StartReplicationWithContext(ctx context.Context, request *StartReplicationRequest) (response *StartReplicationResponse, err error) {
    if request == nil {
        request = NewStartReplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StartReplication")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartReplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartReplicationResponse()
    err = c.Send(request, response)
    return
}

func NewStopCpuExpandRequest() (request *StopCpuExpandRequest) {
    request = &StopCpuExpandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "StopCpuExpand")
    
    
    return
}

func NewStopCpuExpandResponse() (response *StopCpuExpandResponse) {
    response = &StopCpuExpandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCpuExpand
// 通过该API，可以关闭 CPU 弹性扩容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCETASKCONFLICTERROR = "FailedOperation.InstanceTaskConflictError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) StopCpuExpand(request *StopCpuExpandRequest) (response *StopCpuExpandResponse, err error) {
    return c.StopCpuExpandWithContext(context.Background(), request)
}

// StopCpuExpand
// 通过该API，可以关闭 CPU 弹性扩容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INSTANCETASKCONFLICTERROR = "FailedOperation.InstanceTaskConflictError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
func (c *Client) StopCpuExpandWithContext(ctx context.Context, request *StopCpuExpandRequest) (response *StopCpuExpandResponse, err error) {
    if request == nil {
        request = NewStopCpuExpandRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StopCpuExpand")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCpuExpand require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCpuExpandResponse()
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
// 本接口（StopDBImportJob）用于终止数据导入任务。
//
// 说明：只有未完成的导入任务支持被终止，且终止后已执行的 SQL 部分会被保留。
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
// 本接口（StopDBImportJob）用于终止数据导入任务。
//
// 说明：只有未完成的导入任务支持被终止，且终止后已执行的 SQL 部分会被保留。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StopDBImportJob")
    
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
// 本接口（StopReplication）用于停止 RO 复制，中断从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDELAYRO = "FailedOperation.NotDelayRo"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) StopReplication(request *StopReplicationRequest) (response *StopReplicationResponse, err error) {
    return c.StopReplicationWithContext(context.Background(), request)
}

// StopReplication
// 本接口（StopReplication）用于停止 RO 复制，中断从主实例同步数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTDELAYRO = "FailedOperation.NotDelayRo"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_INSTANCETASKRUNNING = "OperationDenied.InstanceTaskRunning"
//  RESOURCENOTFOUND_INSTANCENOTFOUNDERROR = "ResourceNotFound.InstanceNotFoundError"
func (c *Client) StopReplicationWithContext(ctx context.Context, request *StopReplicationRequest) (response *StopReplicationResponse, err error) {
    if request == nil {
        request = NewStopReplicationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StopReplication")
    
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
// 本接口（StopRollback）用于撤销实例正在进行的回档任务，该接口返回一个异步任务 ID。撤销结果可以通过 [DescribeAsyncRequestInfo](https://cloud.tencent.com/document/api/236/20410) 查询任务的执行情况。
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
// 本接口（StopRollback）用于撤销实例正在进行的回档任务，该接口返回一个异步任务 ID。撤销结果可以通过 [DescribeAsyncRequestInfo](https://cloud.tencent.com/document/api/236/20410) 查询任务的执行情况。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "StopRollback")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopRollback require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopRollbackResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitInstanceUpgradeCheckJobRequest() (request *SubmitInstanceUpgradeCheckJobRequest) {
    request = &SubmitInstanceUpgradeCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdb", APIVersion, "SubmitInstanceUpgradeCheckJob")
    
    
    return
}

func NewSubmitInstanceUpgradeCheckJobResponse() (response *SubmitInstanceUpgradeCheckJobResponse) {
    response = &SubmitInstanceUpgradeCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitInstanceUpgradeCheckJob
// 该接口（SubmitInstanceUpgradeCheckJob）提交实例版本升级校验任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) SubmitInstanceUpgradeCheckJob(request *SubmitInstanceUpgradeCheckJobRequest) (response *SubmitInstanceUpgradeCheckJobResponse, err error) {
    return c.SubmitInstanceUpgradeCheckJobWithContext(context.Background(), request)
}

// SubmitInstanceUpgradeCheckJob
// 该接口（SubmitInstanceUpgradeCheckJob）提交实例版本升级校验任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_EXCEPTIONPARAM = "InvalidParameter.ExceptionParam"
//  OPERATIONDENIED_INSTANCETASKSTATUSERROR = "OperationDenied.InstanceTaskStatusError"
//  RESOURCENOTFOUND_CDBINSTANCENOTFOUNDERROR = "ResourceNotFound.CdbInstanceNotFoundError"
func (c *Client) SubmitInstanceUpgradeCheckJobWithContext(ctx context.Context, request *SubmitInstanceUpgradeCheckJobRequest) (response *SubmitInstanceUpgradeCheckJobResponse, err error) {
    if request == nil {
        request = NewSubmitInstanceUpgradeCheckJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "SubmitInstanceUpgradeCheckJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitInstanceUpgradeCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitInstanceUpgradeCheckJobResponse()
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
// 本接口（SwitchCDBProxy）用于数据库代理配置变更或者升级版本后手动发起立即切换。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) SwitchCDBProxy(request *SwitchCDBProxyRequest) (response *SwitchCDBProxyResponse, err error) {
    return c.SwitchCDBProxyWithContext(context.Background(), request)
}

// SwitchCDBProxy
// 本接口（SwitchCDBProxy）用于数据库代理配置变更或者升级版本后手动发起立即切换。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_PROXYUPGRADETASKSTATUSERROR = "OperationDenied.ProxyUpgradeTaskStatusError"
func (c *Client) SwitchCDBProxyWithContext(ctx context.Context, request *SwitchCDBProxyRequest) (response *SwitchCDBProxyResponse, err error) {
    if request == nil {
        request = NewSwitchCDBProxyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "SwitchCDBProxy")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "SwitchDBInstanceMasterSlave")
    
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
// 本接口（SwitchDrInstanceToMaster）用于将云数据库灾备实例切换为主实例，注意请求必须发到灾备实例所在的地域。
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
// 本接口（SwitchDrInstanceToMaster）用于将云数据库灾备实例切换为主实例，注意请求必须发到灾备实例所在的地域。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "SwitchDrInstanceToMaster")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "SwitchForUpgrade")
    
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
// 本接口（UpgradeCDBProxyVersion）用于升级数据库代理版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) UpgradeCDBProxyVersion(request *UpgradeCDBProxyVersionRequest) (response *UpgradeCDBProxyVersionResponse, err error) {
    return c.UpgradeCDBProxyVersionWithContext(context.Background(), request)
}

// UpgradeCDBProxyVersion
// 本接口（UpgradeCDBProxyVersion）用于升级数据库代理版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DESCRIBEPROXYGROUPERROR = "FailedOperation.DescribeProxyGroupError"
//  INTERNALERROR_DBRECORDNOTEXISTERROR = "InternalError.DBRecordNotExistError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) UpgradeCDBProxyVersionWithContext(ctx context.Context, request *UpgradeCDBProxyVersionRequest) (response *UpgradeCDBProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeCDBProxyVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "UpgradeCDBProxyVersion")
    
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
// 本接口（UpgradeDBInstance）用于升级或降级云数据库实例的配置，实例类型支持主实例、灾备实例和只读实例。如果进行迁移业务，请一定填写实例规格（CPU、内存），不然系统会默认以最小允许规格传参。
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
// 本接口（UpgradeDBInstance）用于升级或降级云数据库实例的配置，实例类型支持主实例、灾备实例和只读实例。如果进行迁移业务，请一定填写实例规格（CPU、内存），不然系统会默认以最小允许规格传参。
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "UpgradeDBInstance")
    
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
// 本接口（UpgradeDBInstanceEngineVersion）用于升级云数据库实例版本，实例类型支持主实例、灾备实例和只读实例等。升级前请通过 [SubmitInstanceUpgradeCheckJob](https://cloud.tencent.com/document/product/236/110468) 提交升级检查任务，通过后才能升级。
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
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    return c.UpgradeDBInstanceEngineVersionWithContext(context.Background(), request)
}

// UpgradeDBInstanceEngineVersion
// 本接口（UpgradeDBInstanceEngineVersion）用于升级云数据库实例版本，实例类型支持主实例、灾备实例和只读实例等。升级前请通过 [SubmitInstanceUpgradeCheckJob](https://cloud.tencent.com/document/product/236/110468) 提交升级检查任务，通过后才能升级。
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
//  OPERATIONDENIED_ACTIONNOTSUPPORT = "OperationDenied.ActionNotSupport"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) UpgradeDBInstanceEngineVersionWithContext(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest) (response *UpgradeDBInstanceEngineVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeDBInstanceEngineVersionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "UpgradeDBInstanceEngineVersion")
    
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
    c.InitBaseRequest(&request.BaseRequest, "cdb", APIVersion, "VerifyRootAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyRootAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyRootAccountResponse()
    err = c.Send(request, response)
    return
}
