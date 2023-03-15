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

package v20220308

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-03-08"

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
    
    request.Init().WithApiInfo("keewidb", APIVersion, "AssociateSecurityGroups")
    
    
    return
}

func NewAssociateSecurityGroupsResponse() (response *AssociateSecurityGroupsResponse) {
    response = &AssociateSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssociateSecurityGroups
// 本接口 (AssociateSecurityGroups) 用于安全组批量绑定多个指定实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
func (c *Client) AssociateSecurityGroups(request *AssociateSecurityGroupsRequest) (response *AssociateSecurityGroupsResponse, err error) {
    return c.AssociateSecurityGroupsWithContext(context.Background(), request)
}

// AssociateSecurityGroups
// 本接口 (AssociateSecurityGroups) 用于安全组批量绑定多个指定实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
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

func NewChangeInstanceMasterRequest() (request *ChangeInstanceMasterRequest) {
    request = &ChangeInstanceMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ChangeInstanceMaster")
    
    
    return
}

func NewChangeInstanceMasterResponse() (response *ChangeInstanceMasterResponse) {
    response = &ChangeInstanceMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeInstanceMaster
// 本接口（ChangeInstanceMaster）用于将副本节点提升为主节点。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ChangeInstanceMaster(request *ChangeInstanceMasterRequest) (response *ChangeInstanceMasterResponse, err error) {
    return c.ChangeInstanceMasterWithContext(context.Background(), request)
}

// ChangeInstanceMaster
// 本接口（ChangeInstanceMaster）用于将副本节点提升为主节点。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ChangeInstanceMasterWithContext(ctx context.Context, request *ChangeInstanceMasterRequest) (response *ChangeInstanceMasterResponse, err error) {
    if request == nil {
        request = NewChangeInstanceMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeInstanceMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeInstanceMasterResponse()
    err = c.Send(request, response)
    return
}

func NewCleanUpInstanceRequest() (request *CleanUpInstanceRequest) {
    request = &CleanUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "CleanUpInstance")
    
    
    return
}

func NewCleanUpInstanceResponse() (response *CleanUpInstanceResponse) {
    response = &CleanUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CleanUpInstance
// 本接口（CleanUpInstance）用于立即下线回收站已隔离的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CleanUpInstance(request *CleanUpInstanceRequest) (response *CleanUpInstanceResponse, err error) {
    return c.CleanUpInstanceWithContext(context.Background(), request)
}

// CleanUpInstance
// 本接口（CleanUpInstance）用于立即下线回收站已隔离的实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CleanUpInstanceWithContext(ctx context.Context, request *CleanUpInstanceRequest) (response *CleanUpInstanceResponse, err error) {
    if request == nil {
        request = NewCleanUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CleanUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCleanUpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewClearInstanceRequest() (request *ClearInstanceRequest) {
    request = &ClearInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ClearInstance")
    
    
    return
}

func NewClearInstanceResponse() (response *ClearInstanceResponse) {
    response = &ClearInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ClearInstance
// 本接口（ClearInstance）用于清空实例数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) ClearInstance(request *ClearInstanceRequest) (response *ClearInstanceResponse, err error) {
    return c.ClearInstanceWithContext(context.Background(), request)
}

// ClearInstance
// 本接口（ClearInstance）用于清空实例数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) ClearInstanceWithContext(ctx context.Context, request *ClearInstanceRequest) (response *ClearInstanceResponse, err error) {
    if request == nil {
        request = NewClearInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ClearInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewClearInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBackupManuallyRequest() (request *CreateBackupManuallyRequest) {
    request = &CreateBackupManuallyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "CreateBackupManually")
    
    
    return
}

func NewCreateBackupManuallyResponse() (response *CreateBackupManuallyResponse) {
    response = &CreateBackupManuallyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBackupManually
// 手动发起备份
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CreateBackupManually(request *CreateBackupManuallyRequest) (response *CreateBackupManuallyResponse, err error) {
    return c.CreateBackupManuallyWithContext(context.Background(), request)
}

// CreateBackupManually
// 手动发起备份
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) CreateBackupManuallyWithContext(ctx context.Context, request *CreateBackupManuallyRequest) (response *CreateBackupManuallyResponse, err error) {
    if request == nil {
        request = NewCreateBackupManuallyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBackupManually require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBackupManuallyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstances
// 创建数据库实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_INVALIDSIZE = "LimitExceeded.InvalidSize"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// 创建数据库实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_INVALIDSIZE = "LimitExceeded.InvalidSize"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) CreateInstancesWithContext(ctx context.Context, request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    if request == nil {
        request = NewCreateInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoBackupConfigRequest() (request *DescribeAutoBackupConfigRequest) {
    request = &DescribeAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeAutoBackupConfig")
    
    
    return
}

func NewDescribeAutoBackupConfigResponse() (response *DescribeAutoBackupConfigResponse) {
    response = &DescribeAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoBackupConfig
// 本接口（DescribeAutoBackupConfig）用于获取自动备份配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoBackupConfig(request *DescribeAutoBackupConfigRequest) (response *DescribeAutoBackupConfigResponse, err error) {
    return c.DescribeAutoBackupConfigWithContext(context.Background(), request)
}

// DescribeAutoBackupConfig
// 本接口（DescribeAutoBackupConfig）用于获取自动备份配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeAutoBackupConfigWithContext(ctx context.Context, request *DescribeAutoBackupConfigRequest) (response *DescribeAutoBackupConfigResponse, err error) {
    if request == nil {
        request = NewDescribeAutoBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAutoBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAutoBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeDBSecurityGroups")
    
    
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
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewDescribeInstanceBackupsRequest() (request *DescribeInstanceBackupsRequest) {
    request = &DescribeInstanceBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceBackups")
    
    
    return
}

func NewDescribeInstanceBackupsResponse() (response *DescribeInstanceBackupsResponse) {
    response = &DescribeInstanceBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceBackups
// 本接口（DescribeInstanceBackups）用于查询实例全量备份列表。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceBackups(request *DescribeInstanceBackupsRequest) (response *DescribeInstanceBackupsResponse, err error) {
    return c.DescribeInstanceBackupsWithContext(context.Background(), request)
}

// DescribeInstanceBackups
// 本接口（DescribeInstanceBackups）用于查询实例全量备份列表。
//
// 可能返回的错误码:
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceBackupsWithContext(ctx context.Context, request *DescribeInstanceBackupsRequest) (response *DescribeInstanceBackupsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceBackupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceBackups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceBackupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceBinlogsRequest() (request *DescribeInstanceBinlogsRequest) {
    request = &DescribeInstanceBinlogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceBinlogs")
    
    
    return
}

func NewDescribeInstanceBinlogsResponse() (response *DescribeInstanceBinlogsResponse) {
    response = &DescribeInstanceBinlogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceBinlogs
// 本接口（DescribeInstanceBinlogs）用于查询增量备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceBinlogs(request *DescribeInstanceBinlogsRequest) (response *DescribeInstanceBinlogsResponse, err error) {
    return c.DescribeInstanceBinlogsWithContext(context.Background(), request)
}

// DescribeInstanceBinlogs
// 本接口（DescribeInstanceBinlogs）用于查询增量备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceBinlogsWithContext(ctx context.Context, request *DescribeInstanceBinlogsRequest) (response *DescribeInstanceBinlogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceBinlogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceBinlogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceBinlogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDealDetailRequest() (request *DescribeInstanceDealDetailRequest) {
    request = &DescribeInstanceDealDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceDealDetail")
    
    
    return
}

func NewDescribeInstanceDealDetailResponse() (response *DescribeInstanceDealDetailResponse) {
    response = &DescribeInstanceDealDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceDealDetail
// 本接口（DescribeInstanceDealDetail）用于查询预付费订单信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeInstanceDealDetail(request *DescribeInstanceDealDetailRequest) (response *DescribeInstanceDealDetailResponse, err error) {
    return c.DescribeInstanceDealDetailWithContext(context.Background(), request)
}

// DescribeInstanceDealDetail
// 本接口（DescribeInstanceDealDetail）用于查询预付费订单信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeInstanceDealDetailWithContext(ctx context.Context, request *DescribeInstanceDealDetailRequest) (response *DescribeInstanceDealDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDealDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDealDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDealDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodeInfoRequest() (request *DescribeInstanceNodeInfoRequest) {
    request = &DescribeInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceNodeInfo")
    
    
    return
}

func NewDescribeInstanceNodeInfoResponse() (response *DescribeInstanceNodeInfoResponse) {
    response = &DescribeInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceNodeInfo
// 本接口（DescribeInstanceNodeInfo）查询实例节点信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfo(request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    return c.DescribeInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodeInfo
// 本接口（DescribeInstanceNodeInfo）查询实例节点信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfoWithContext(ctx context.Context, request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceNodeInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceNodeInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceNodeInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceParamRecordsRequest() (request *DescribeInstanceParamRecordsRequest) {
    request = &DescribeInstanceParamRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceParamRecords")
    
    
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParamRecords
// 本接口（DescribeInstanceParamRecords）查询参数配置修改历史列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// 本接口（DescribeInstanceParamRecords）查询参数配置修改历史列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
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
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询实例的参数列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询实例的参数列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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

func NewDescribeInstanceReplicasRequest() (request *DescribeInstanceReplicasRequest) {
    request = &DescribeInstanceReplicasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstanceReplicas")
    
    
    return
}

func NewDescribeInstanceReplicasResponse() (response *DescribeInstanceReplicasResponse) {
    response = &DescribeInstanceReplicasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceReplicas
// 本接口（DescribeInstanceReplicas）用于获取实例副本节点信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeInstanceReplicas(request *DescribeInstanceReplicasRequest) (response *DescribeInstanceReplicasResponse, err error) {
    return c.DescribeInstanceReplicasWithContext(context.Background(), request)
}

// DescribeInstanceReplicas
// 本接口（DescribeInstanceReplicas）用于获取实例副本节点信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeInstanceReplicasWithContext(ctx context.Context, request *DescribeInstanceReplicasRequest) (response *DescribeInstanceReplicasResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceReplicasRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceReplicas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceReplicasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 本接口（DescribeInstances）可以根据地域、网络、实例id、标签、计费方式等条件，搜索查询实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 本接口（DescribeInstances）可以根据地域、网络、实例id、标签、计费方式等条件，搜索查询实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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

func NewDescribeMaintenanceWindowRequest() (request *DescribeMaintenanceWindowRequest) {
    request = &DescribeMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeMaintenanceWindow")
    
    
    return
}

func NewDescribeMaintenanceWindowResponse() (response *DescribeMaintenanceWindowResponse) {
    response = &DescribeMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMaintenanceWindow
// 本接口（DescribeMaintenanceWindow）用于查询实例维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeMaintenanceWindow(request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    return c.DescribeMaintenanceWindowWithContext(context.Background(), request)
}

// DescribeMaintenanceWindow
// 本接口（DescribeMaintenanceWindow）用于查询实例维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeMaintenanceWindowWithContext(ctx context.Context, request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewDescribeMaintenanceWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
    request = &DescribeProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeProductInfo")
    
    
    return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
    response = &DescribeProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProductInfo
// 本接口查询指定可用区和实例类型下keewidb 的售卖规格， 如果用户不在购买白名单中，将不能查询该可用区或该类型的售卖规格详情。申请购买某地域白名单可以提交工单
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProductInfo(request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    return c.DescribeProductInfoWithContext(context.Background(), request)
}

// DescribeProductInfo
// 本接口查询指定可用区和实例类型下keewidb 的售卖规格， 如果用户不在购买白名单中，将不能查询该可用区或该类型的售卖规格详情。申请购买某地域白名单可以提交工单
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeProductInfoWithContext(ctx context.Context, request *DescribeProductInfoRequest) (response *DescribeProductInfoResponse, err error) {
    if request == nil {
        request = NewDescribeProductInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProductInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProductInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeProjectSecurityGroups")
    
    
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
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
func (c *Client) DescribeProjectSecurityGroups(request *DescribeProjectSecurityGroupsRequest) (response *DescribeProjectSecurityGroupsResponse, err error) {
    return c.DescribeProjectSecurityGroupsWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroups
// 本接口(DescribeProjectSecurityGroups)用于查询项目的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
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

func NewDescribeProxySlowLogRequest() (request *DescribeProxySlowLogRequest) {
    request = &DescribeProxySlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeProxySlowLog")
    
    
    return
}

func NewDescribeProxySlowLogResponse() (response *DescribeProxySlowLogResponse) {
    response = &DescribeProxySlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeProxySlowLog
// 本接口（DescribeProxySlowLog）用于查询代理（Proxy）慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeProxySlowLog(request *DescribeProxySlowLogRequest) (response *DescribeProxySlowLogResponse, err error) {
    return c.DescribeProxySlowLogWithContext(context.Background(), request)
}

// DescribeProxySlowLog
// 本接口（DescribeProxySlowLog）用于查询代理（Proxy）慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeProxySlowLogWithContext(ctx context.Context, request *DescribeProxySlowLogRequest) (response *DescribeProxySlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeProxySlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeTaskInfo")
    
    
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskInfo
// 本接口（DescribeTaskInfo）用于查询异步任务结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    return c.DescribeTaskInfoWithContext(context.Background(), request)
}

// DescribeTaskInfo
// 本接口（DescribeTaskInfo）用于查询异步任务结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskInfoWithContext(ctx context.Context, request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTaskInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskListRequest() (request *DescribeTaskListRequest) {
    request = &DescribeTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeTaskList")
    
    
    return
}

func NewDescribeTaskListResponse() (response *DescribeTaskListResponse) {
    response = &DescribeTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskList
// 本接口（DescribeTaskList）用于查询任务列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskList(request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    return c.DescribeTaskListWithContext(context.Background(), request)
}

// DescribeTaskList
// 本接口（DescribeTaskList）用于查询任务列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskListWithContext(ctx context.Context, request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTendisSlowLogRequest() (request *DescribeTendisSlowLogRequest) {
    request = &DescribeTendisSlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DescribeTendisSlowLog")
    
    
    return
}

func NewDescribeTendisSlowLogResponse() (response *DescribeTendisSlowLogResponse) {
    response = &DescribeTendisSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTendisSlowLog
// 本接口（DescribeTendisSlowLog）用于查询实例慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeTendisSlowLog(request *DescribeTendisSlowLogRequest) (response *DescribeTendisSlowLogResponse, err error) {
    return c.DescribeTendisSlowLogWithContext(context.Background(), request)
}

// DescribeTendisSlowLog
// 本接口（DescribeTendisSlowLog）用于查询实例慢日志。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeTendisSlowLogWithContext(ctx context.Context, request *DescribeTendisSlowLogRequest) (response *DescribeTendisSlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeTendisSlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTendisSlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTendisSlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPostpaidInstanceRequest() (request *DestroyPostpaidInstanceRequest) {
    request = &DestroyPostpaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DestroyPostpaidInstance")
    
    
    return
}

func NewDestroyPostpaidInstanceResponse() (response *DestroyPostpaidInstanceResponse) {
    response = &DestroyPostpaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyPostpaidInstance
// 本接口（DestroyPostpaidInstance）用于退还按量计费实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyPostpaidInstance(request *DestroyPostpaidInstanceRequest) (response *DestroyPostpaidInstanceResponse, err error) {
    return c.DestroyPostpaidInstanceWithContext(context.Background(), request)
}

// DestroyPostpaidInstance
// 本接口（DestroyPostpaidInstance）用于退还按量计费实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyPostpaidInstanceWithContext(ctx context.Context, request *DestroyPostpaidInstanceRequest) (response *DestroyPostpaidInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPostpaidInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPostpaidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPostpaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyPrepaidInstanceRequest() (request *DestroyPrepaidInstanceRequest) {
    request = &DestroyPrepaidInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DestroyPrepaidInstance")
    
    
    return
}

func NewDestroyPrepaidInstanceResponse() (response *DestroyPrepaidInstanceResponse) {
    response = &DestroyPrepaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyPrepaidInstance
// 本接口（DestroyPrepaidInstance）用于退还包年包月计费实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyPrepaidInstance(request *DestroyPrepaidInstanceRequest) (response *DestroyPrepaidInstanceResponse, err error) {
    return c.DestroyPrepaidInstanceWithContext(context.Background(), request)
}

// DestroyPrepaidInstance
// 本接口（DestroyPrepaidInstance）用于退还包年包月计费实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DestroyPrepaidInstanceWithContext(ctx context.Context, request *DestroyPrepaidInstanceRequest) (response *DestroyPrepaidInstanceResponse, err error) {
    if request == nil {
        request = NewDestroyPrepaidInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyPrepaidInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyPrepaidInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "DisassociateSecurityGroups")
    
    
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
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
func (c *Client) DisassociateSecurityGroups(request *DisassociateSecurityGroupsRequest) (response *DisassociateSecurityGroupsResponse, err error) {
    return c.DisassociateSecurityGroupsWithContext(context.Background(), request)
}

// DisassociateSecurityGroups
// 本接口(DisassociateSecurityGroups)用于安全组批量解绑实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
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

func NewModifyAutoBackupConfigRequest() (request *ModifyAutoBackupConfigRequest) {
    request = &ModifyAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyAutoBackupConfig")
    
    
    return
}

func NewModifyAutoBackupConfigResponse() (response *ModifyAutoBackupConfigResponse) {
    response = &ModifyAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAutoBackupConfig
// 本接口（ModifyAutoBackupConfig）用于修改自动备份配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_WEEKDAYSISINVALID = "InvalidParameterValue.WeekDaysIsInvalid"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyAutoBackupConfig(request *ModifyAutoBackupConfigRequest) (response *ModifyAutoBackupConfigResponse, err error) {
    return c.ModifyAutoBackupConfigWithContext(context.Background(), request)
}

// ModifyAutoBackupConfig
// 本接口（ModifyAutoBackupConfig）用于修改自动备份配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_WEEKDAYSISINVALID = "InvalidParameterValue.WeekDaysIsInvalid"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyAutoBackupConfigWithContext(ctx context.Context, request *ModifyAutoBackupConfigRequest) (response *ModifyAutoBackupConfigResponse, err error) {
    if request == nil {
        request = NewModifyAutoBackupConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoBackupConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoBackupConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConnectionConfigRequest() (request *ModifyConnectionConfigRequest) {
    request = &ModifyConnectionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyConnectionConfig")
    
    
    return
}

func NewModifyConnectionConfigResponse() (response *ModifyConnectionConfigResponse) {
    response = &ModifyConnectionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyConnectionConfig
// 本接口（ModifyConnectionConfig）用于修改实例的连接配置，包括带宽和最大连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyConnectionConfig(request *ModifyConnectionConfigRequest) (response *ModifyConnectionConfigResponse, err error) {
    return c.ModifyConnectionConfigWithContext(context.Background(), request)
}

// ModifyConnectionConfig
// 本接口（ModifyConnectionConfig）用于修改实例的连接配置，包括带宽和最大连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyConnectionConfigWithContext(ctx context.Context, request *ModifyConnectionConfigRequest) (response *ModifyConnectionConfigResponse, err error) {
    if request == nil {
        request = NewModifyConnectionConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConnectionConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConnectionConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
    request = &ModifyDBInstanceSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
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
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
    return c.ModifyDBInstanceSecurityGroupsWithContext(context.Background(), request)
}

// ModifyDBInstanceSecurityGroups
// 本接口(ModifyDBInstanceSecurityGroups)用于修改实例绑定的安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"
//  FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
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

func NewModifyInstanceRequest() (request *ModifyInstanceRequest) {
    request = &ModifyInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstance
// 本接口（ModifyInstance）用于修改实例相关信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// 本接口（ModifyInstance）用于修改实例相关信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
func (c *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    if request == nil {
        request = NewModifyInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamsRequest() (request *ModifyInstanceParamsRequest) {
    request = &ModifyInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyInstanceParams")
    
    
    return
}

func NewModifyInstanceParamsResponse() (response *ModifyInstanceParamsResponse) {
    response = &ModifyInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceParams
// 本接口（ModifyInstanceParams）用于修改实例参数配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyInstanceParams(request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    return c.ModifyInstanceParamsWithContext(context.Background(), request)
}

// ModifyInstanceParams
// 本接口（ModifyInstanceParams）用于修改实例参数配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ModifyInstanceParamsWithContext(ctx context.Context, request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    if request == nil {
        request = NewModifyInstanceParamsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceParams require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceParamsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceWindowRequest() (request *ModifyMaintenanceWindowRequest) {
    request = &ModifyMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyMaintenanceWindow")
    
    
    return
}

func NewModifyMaintenanceWindowResponse() (response *ModifyMaintenanceWindowResponse) {
    response = &ModifyMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMaintenanceWindow
// 本接口（ModifyMaintenanceWindow）修改实例维护时间窗时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyMaintenanceWindow(request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    return c.ModifyMaintenanceWindowWithContext(context.Background(), request)
}

// ModifyMaintenanceWindow
// 本接口（ModifyMaintenanceWindow）修改实例维护时间窗时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyMaintenanceWindowWithContext(ctx context.Context, request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    if request == nil {
        request = NewModifyMaintenanceWindowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMaintenanceWindow require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMaintenanceWindowResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNetworkConfigRequest() (request *ModifyNetworkConfigRequest) {
    request = &ModifyNetworkConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ModifyNetworkConfig")
    
    
    return
}

func NewModifyNetworkConfigResponse() (response *ModifyNetworkConfigResponse) {
    response = &ModifyNetworkConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNetworkConfig
// 本接口（ModifyNetworkConfig）用于修改实例网络配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyNetworkConfig(request *ModifyNetworkConfigRequest) (response *ModifyNetworkConfigResponse, err error) {
    return c.ModifyNetworkConfigWithContext(context.Background(), request)
}

// ModifyNetworkConfig
// 本接口（ModifyNetworkConfig）用于修改实例网络配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyNetworkConfigWithContext(ctx context.Context, request *ModifyNetworkConfigRequest) (response *ModifyNetworkConfigResponse, err error) {
    if request == nil {
        request = NewModifyNetworkConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNetworkConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNetworkConfigResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "RenewInstance")
    
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RenewInstance
// 本接口（RenewInstance）用于为包年包月计费实例续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    return c.RenewInstanceWithContext(context.Background(), request)
}

// RenewInstance
// 本接口（RenewInstance）用于为包年包月计费实例续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    if request == nil {
        request = NewRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewResetPasswordRequest() (request *ResetPasswordRequest) {
    request = &ResetPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "ResetPassword")
    
    
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetPassword
// 本接口（ResetPassword）用于重置数据库访问密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

// ResetPassword
// 本接口（ResetPassword）用于重置数据库访问密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) ResetPasswordWithContext(ctx context.Context, request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    if request == nil {
        request = NewResetPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewStartUpInstanceRequest() (request *StartUpInstanceRequest) {
    request = &StartUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "StartUpInstance")
    
    
    return
}

func NewStartUpInstanceResponse() (response *StartUpInstanceResponse) {
    response = &StartUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartUpInstance
// 本接口（StartUpInstance）用于按量计费实例解隔离
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StartUpInstance(request *StartUpInstanceRequest) (response *StartUpInstanceResponse, err error) {
    return c.StartUpInstanceWithContext(context.Background(), request)
}

// StartUpInstance
// 本接口（StartUpInstance）用于按量计费实例解隔离
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) StartUpInstanceWithContext(ctx context.Context, request *StartUpInstanceRequest) (response *StartUpInstanceResponse, err error) {
    if request == nil {
        request = NewStartUpInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartUpInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartUpInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("keewidb", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeInstance
// 本接口（UpgradeInstance）用于对实例进行配置变更。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_SPECNOTSUPPORTED = "InvalidParameter.SpecNotSupported"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 本接口（UpgradeInstance）用于对实例进行配置变更。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETER_SPECNOTSUPPORTED = "InvalidParameter.SpecNotSupported"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
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
