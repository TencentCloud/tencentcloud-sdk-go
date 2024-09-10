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

package v20180412

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-12"

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


func NewAddReplicationInstanceRequest() (request *AddReplicationInstanceRequest) {
    request = &AddReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "AddReplicationInstance")
    
    
    return
}

func NewAddReplicationInstanceResponse() (response *AddReplicationInstanceResponse) {
    response = &AddReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddReplicationInstance
// 添加复制组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_INSTANCENAMERULEERROR = "InvalidParameterValue.InstanceNameRuleError"
//  INVALIDPARAMETERVALUE_NOTREPEATBIND = "InvalidParameterValue.NotRepeatBind"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_EXCEEDUPPERLIMIT = "LimitExceeded.ExceedUpperLimit"
//  LIMITEXCEEDED_INSTANCENOTEMPTY = "LimitExceeded.InstanceNotEmpty"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) AddReplicationInstance(request *AddReplicationInstanceRequest) (response *AddReplicationInstanceResponse, err error) {
    return c.AddReplicationInstanceWithContext(context.Background(), request)
}

// AddReplicationInstance
// 添加复制组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_INSTANCENAMERULEERROR = "InvalidParameterValue.InstanceNameRuleError"
//  INVALIDPARAMETERVALUE_NOTREPEATBIND = "InvalidParameterValue.NotRepeatBind"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_EXCEEDUPPERLIMIT = "LimitExceeded.ExceedUpperLimit"
//  LIMITEXCEEDED_INSTANCENOTEMPTY = "LimitExceeded.InstanceNotEmpty"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) AddReplicationInstanceWithContext(ctx context.Context, request *AddReplicationInstanceRequest) (response *AddReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewAddReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddReplicationInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewAllocateWanAddressRequest() (request *AllocateWanAddressRequest) {
    request = &AllocateWanAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "AllocateWanAddress")
    
    
    return
}

func NewAllocateWanAddressResponse() (response *AllocateWanAddressResponse) {
    response = &AllocateWanAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AllocateWanAddress
// 开通外网
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_ISNOTVPCINSTANCE = "InvalidParameter.IsNotVpcInstance"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
func (c *Client) AllocateWanAddress(request *AllocateWanAddressRequest) (response *AllocateWanAddressResponse, err error) {
    return c.AllocateWanAddressWithContext(context.Background(), request)
}

// AllocateWanAddress
// 开通外网
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_ISNOTVPCINSTANCE = "InvalidParameter.IsNotVpcInstance"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
func (c *Client) AllocateWanAddressWithContext(ctx context.Context, request *AllocateWanAddressRequest) (response *AllocateWanAddressResponse, err error) {
    if request == nil {
        request = NewAllocateWanAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AllocateWanAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewAllocateWanAddressResponse()
    err = c.Send(request, response)
    return
}

func NewApplyParamsTemplateRequest() (request *ApplyParamsTemplateRequest) {
    request = &ApplyParamsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ApplyParamsTemplate")
    
    
    return
}

func NewApplyParamsTemplateResponse() (response *ApplyParamsTemplateResponse) {
    response = &ApplyParamsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyParamsTemplate
// 应用参数模板到实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ApplyParamsTemplate(request *ApplyParamsTemplateRequest) (response *ApplyParamsTemplateResponse, err error) {
    return c.ApplyParamsTemplateWithContext(context.Background(), request)
}

// ApplyParamsTemplate
// 应用参数模板到实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ApplyParamsTemplateWithContext(ctx context.Context, request *ApplyParamsTemplateRequest) (response *ApplyParamsTemplateResponse, err error) {
    if request == nil {
        request = NewApplyParamsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyParamsTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyParamsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewAssociateSecurityGroupsRequest() (request *AssociateSecurityGroupsRequest) {
    request = &AssociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "AssociateSecurityGroups")
    
    
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
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_INSTANCESGOVERLIMITERROR = "InvalidParameter.InstanceSGOverLimitError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_INSTANCESGOVERLIMITERROR = "InvalidParameter.InstanceSGOverLimitError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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

func NewChangeInstanceRoleRequest() (request *ChangeInstanceRoleRequest) {
    request = &ChangeInstanceRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ChangeInstanceRole")
    
    
    return
}

func NewChangeInstanceRoleResponse() (response *ChangeInstanceRoleResponse) {
    response = &ChangeInstanceRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeInstanceRole
// 复制组实例更换角色
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeInstanceRole(request *ChangeInstanceRoleRequest) (response *ChangeInstanceRoleResponse, err error) {
    return c.ChangeInstanceRoleWithContext(context.Background(), request)
}

// ChangeInstanceRole
// 复制组实例更换角色
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeInstanceRoleWithContext(ctx context.Context, request *ChangeInstanceRoleRequest) (response *ChangeInstanceRoleResponse, err error) {
    if request == nil {
        request = NewChangeInstanceRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeInstanceRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeInstanceRoleResponse()
    err = c.Send(request, response)
    return
}

func NewChangeMasterInstanceRequest() (request *ChangeMasterInstanceRequest) {
    request = &ChangeMasterInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ChangeMasterInstance")
    
    
    return
}

func NewChangeMasterInstanceResponse() (response *ChangeMasterInstanceResponse) {
    response = &ChangeMasterInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeMasterInstance
// 该接口（ChangeMasterInstance）用于将复制组内只读实例设置为主实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeMasterInstance(request *ChangeMasterInstanceRequest) (response *ChangeMasterInstanceResponse, err error) {
    return c.ChangeMasterInstanceWithContext(context.Background(), request)
}

// ChangeMasterInstance
// 该接口（ChangeMasterInstance）用于将复制组内只读实例设置为主实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ChangeMasterInstanceWithContext(ctx context.Context, request *ChangeMasterInstanceRequest) (response *ChangeMasterInstanceResponse, err error) {
    if request == nil {
        request = NewChangeMasterInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeMasterInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeMasterInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewChangeReplicaToMasterRequest() (request *ChangeReplicaToMasterRequest) {
    request = &ChangeReplicaToMasterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ChangeReplicaToMaster")
    
    
    return
}

func NewChangeReplicaToMasterResponse() (response *ChangeReplicaToMasterResponse) {
    response = &ChangeReplicaToMasterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeReplicaToMaster
// 本接口（ChangeReplicaToMaster）适用于实例副本组提主或副本提主。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) ChangeReplicaToMaster(request *ChangeReplicaToMasterRequest) (response *ChangeReplicaToMasterResponse, err error) {
    return c.ChangeReplicaToMasterWithContext(context.Background(), request)
}

// ChangeReplicaToMaster
// 本接口（ChangeReplicaToMaster）适用于实例副本组提主或副本提主。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) ChangeReplicaToMasterWithContext(ctx context.Context, request *ChangeReplicaToMasterRequest) (response *ChangeReplicaToMasterResponse, err error) {
    if request == nil {
        request = NewChangeReplicaToMasterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeReplicaToMaster require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeReplicaToMasterResponse()
    err = c.Send(request, response)
    return
}

func NewCleanUpInstanceRequest() (request *CleanUpInstanceRequest) {
    request = &CleanUpInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CleanUpInstance")
    
    
    return
}

func NewCleanUpInstanceResponse() (response *CleanUpInstanceResponse) {
    response = &CleanUpInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CleanUpInstance
// 回收站实例立即下线
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
// 回收站实例立即下线
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
    
    request.Init().WithApiInfo("redis", APIVersion, "ClearInstance")
    
    
    return
}

func NewClearInstanceResponse() (response *ClearInstanceResponse) {
    response = &ClearInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ClearInstance
// 清空Redis实例的实例数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ClearInstance(request *ClearInstanceRequest) (response *ClearInstanceResponse, err error) {
    return c.ClearInstanceWithContext(context.Background(), request)
}

// ClearInstance
// 清空Redis实例的实例数据。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewCloneInstancesRequest() (request *CloneInstancesRequest) {
    request = &CloneInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CloneInstances")
    
    
    return
}

func NewCloneInstancesResponse() (response *CloneInstancesResponse) {
    response = &CloneInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloneInstances
// 本接口（CloneInstances）用于基于当前实例的备份文件克隆一个完整的新实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"
//  INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  LIMITEXCEEDED_REACHTHEAMOUNTLIMIT = "LimitExceeded.ReachTheAmountLimit"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"
//  RESOURCEUNAVAILABLE_NOSERVICEAVAILABLEFORTHISZONEID = "ResourceUnavailable.NoServiceAvailableForThisZoneId"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CloneInstances(request *CloneInstancesRequest) (response *CloneInstancesResponse, err error) {
    return c.CloneInstancesWithContext(context.Background(), request)
}

// CloneInstances
// 本接口（CloneInstances）用于基于当前实例的备份文件克隆一个完整的新实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"
//  INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  LIMITEXCEEDED_REACHTHEAMOUNTLIMIT = "LimitExceeded.ReachTheAmountLimit"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"
//  RESOURCEUNAVAILABLE_NOSERVICEAVAILABLEFORTHISZONEID = "ResourceUnavailable.NoServiceAvailableForThisZoneId"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CloneInstancesWithContext(ctx context.Context, request *CloneInstancesRequest) (response *CloneInstancesResponse, err error) {
    if request == nil {
        request = NewCloneInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloneInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloneInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCloseSSLRequest() (request *CloseSSLRequest) {
    request = &CloseSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CloseSSL")
    
    
    return
}

func NewCloseSSLResponse() (response *CloseSSLResponse) {
    response = &CloseSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseSSL
// 关闭SSL
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) CloseSSL(request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    return c.CloseSSLWithContext(context.Background(), request)
}

// CloseSSL
// 关闭SSL
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) CloseSSLWithContext(ctx context.Context, request *CloseSSLRequest) (response *CloseSSLResponse, err error) {
    if request == nil {
        request = NewCloseSSLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseSSLResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceAccountRequest() (request *CreateInstanceAccountRequest) {
    request = &CreateInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateInstanceAccount")
    
    
    return
}

func NewCreateInstanceAccountResponse() (response *CreateInstanceAccountResponse) {
    response = &CreateInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstanceAccount
// 该接口（CreateInstanceAccount）用于自定义访问实例的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) CreateInstanceAccount(request *CreateInstanceAccountRequest) (response *CreateInstanceAccountResponse, err error) {
    return c.CreateInstanceAccountWithContext(context.Background(), request)
}

// CreateInstanceAccount
// 该接口（CreateInstanceAccount）用于自定义访问实例的账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) CreateInstanceAccountWithContext(ctx context.Context, request *CreateInstanceAccountRequest) (response *CreateInstanceAccountResponse, err error) {
    if request == nil {
        request = NewCreateInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancesRequest() (request *CreateInstancesRequest) {
    request = &CreateInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateInstances")
    
    
    return
}

func NewCreateInstancesResponse() (response *CreateInstancesResponse) {
    response = &CreateInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateInstances
// 本接口（CreateInstances）用于创建 Redis 实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"
//  INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"
//  RESOURCEUNAVAILABLE_NOTYPEIDREDISSERVICE = "ResourceUnavailable.NoTypeIdRedisService"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
    return c.CreateInstancesWithContext(context.Background(), request)
}

// CreateInstances
// 本接口（CreateInstances）用于创建 Redis 实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"
//  INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"
//  INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"
//  RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"
//  RESOURCEUNAVAILABLE_NOTYPEIDREDISSERVICE = "ResourceUnavailable.NoTypeIdRedisService"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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

func NewCreateParamTemplateRequest() (request *CreateParamTemplateRequest) {
    request = &CreateParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateParamTemplate")
    
    
    return
}

func NewCreateParamTemplateResponse() (response *CreateParamTemplateResponse) {
    response = &CreateParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateParamTemplate
// 创建参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) CreateParamTemplate(request *CreateParamTemplateRequest) (response *CreateParamTemplateResponse, err error) {
    return c.CreateParamTemplateWithContext(context.Background(), request)
}

// CreateParamTemplate
// 创建参数模板。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewCreateReplicationGroupRequest() (request *CreateReplicationGroupRequest) {
    request = &CreateReplicationGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "CreateReplicationGroup")
    
    
    return
}

func NewCreateReplicationGroupResponse() (response *CreateReplicationGroupResponse) {
    response = &CreateReplicationGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReplicationGroup
// 该接口（CreateReplicationGroup）用于创建复制组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_INSTANCENAMERULEERROR = "InvalidParameterValue.InstanceNameRuleError"
//  INVALIDPARAMETERVALUE_NOTREPEATBIND = "InvalidParameterValue.NotRepeatBind"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_INSTANCENOTEMPTY = "LimitExceeded.InstanceNotEmpty"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) CreateReplicationGroup(request *CreateReplicationGroupRequest) (response *CreateReplicationGroupResponse, err error) {
    return c.CreateReplicationGroupWithContext(context.Background(), request)
}

// CreateReplicationGroup
// 该接口（CreateReplicationGroup）用于创建复制组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_INSTANCENAMERULEERROR = "InvalidParameterValue.InstanceNameRuleError"
//  INVALIDPARAMETERVALUE_NOTREPEATBIND = "InvalidParameterValue.NotRepeatBind"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_INSTANCENOTEMPTY = "LimitExceeded.InstanceNotEmpty"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) CreateReplicationGroupWithContext(ctx context.Context, request *CreateReplicationGroupRequest) (response *CreateReplicationGroupResponse, err error) {
    if request == nil {
        request = NewCreateReplicationGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReplicationGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReplicationGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceAccountRequest() (request *DeleteInstanceAccountRequest) {
    request = &DeleteInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DeleteInstanceAccount")
    
    
    return
}

func NewDeleteInstanceAccountResponse() (response *DeleteInstanceAccountResponse) {
    response = &DeleteInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteInstanceAccount
// 删除实例子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DeleteInstanceAccount(request *DeleteInstanceAccountRequest) (response *DeleteInstanceAccountResponse, err error) {
    return c.DeleteInstanceAccountWithContext(context.Background(), request)
}

// DeleteInstanceAccount
// 删除实例子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DeleteInstanceAccountWithContext(ctx context.Context, request *DeleteInstanceAccountRequest) (response *DeleteInstanceAccountResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteParamTemplateRequest() (request *DeleteParamTemplateRequest) {
    request = &DeleteParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DeleteParamTemplate")
    
    
    return
}

func NewDeleteParamTemplateResponse() (response *DeleteParamTemplateResponse) {
    response = &DeleteParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteParamTemplate
// 删除参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DeleteParamTemplate(request *DeleteParamTemplateRequest) (response *DeleteParamTemplateResponse, err error) {
    return c.DeleteParamTemplateWithContext(context.Background(), request)
}

// DeleteParamTemplate
// 删除参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewDeleteReplicationInstanceRequest() (request *DeleteReplicationInstanceRequest) {
    request = &DeleteReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DeleteReplicationInstance")
    
    
    return
}

func NewDeleteReplicationInstanceResponse() (response *DeleteReplicationInstanceResponse) {
    response = &DeleteReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteReplicationInstance
// 移除复制组成员 注：接口下线中，请使用 RemoveReplicationInstance
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) DeleteReplicationInstance(request *DeleteReplicationInstanceRequest) (response *DeleteReplicationInstanceResponse, err error) {
    return c.DeleteReplicationInstanceWithContext(context.Background(), request)
}

// DeleteReplicationInstance
// 移除复制组成员 注：接口下线中，请使用 RemoveReplicationInstance
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) DeleteReplicationInstanceWithContext(ctx context.Context, request *DeleteReplicationInstanceRequest) (response *DeleteReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteReplicationInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoBackupConfigRequest() (request *DescribeAutoBackupConfigRequest) {
    request = &DescribeAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeAutoBackupConfig")
    
    
    return
}

func NewDescribeAutoBackupConfigResponse() (response *DescribeAutoBackupConfigResponse) {
    response = &DescribeAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAutoBackupConfig
// 本接口（DescribeAutoBackupConfig）用于获取自动备份配置规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeAutoBackupConfig(request *DescribeAutoBackupConfigRequest) (response *DescribeAutoBackupConfigResponse, err error) {
    return c.DescribeAutoBackupConfigWithContext(context.Background(), request)
}

// DescribeAutoBackupConfig
// 本接口（DescribeAutoBackupConfig）用于获取自动备份配置规则。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewDescribeBackupDetailRequest() (request *DescribeBackupDetailRequest) {
    request = &DescribeBackupDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBackupDetail")
    
    
    return
}

func NewDescribeBackupDetailResponse() (response *DescribeBackupDetailResponse) {
    response = &DescribeBackupDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDetail
// 本接口（DescribeBackupDetail）用于查询实例的备份信息详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeBackupDetail(request *DescribeBackupDetailRequest) (response *DescribeBackupDetailResponse, err error) {
    return c.DescribeBackupDetailWithContext(context.Background(), request)
}

// DescribeBackupDetail
// 本接口（DescribeBackupDetail）用于查询实例的备份信息详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeBackupDetailWithContext(ctx context.Context, request *DescribeBackupDetailRequest) (response *DescribeBackupDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBackupDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBackupDownloadRestrictionRequest() (request *DescribeBackupDownloadRestrictionRequest) {
    request = &DescribeBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBackupDownloadRestriction")
    
    
    return
}

func NewDescribeBackupDownloadRestrictionResponse() (response *DescribeBackupDownloadRestrictionResponse) {
    response = &DescribeBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupDownloadRestriction
// 本接口（DescribeBackupDownloadRestriction）用于查询当前地域数据库备份文件的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeBackupDownloadRestriction(request *DescribeBackupDownloadRestrictionRequest) (response *DescribeBackupDownloadRestrictionResponse, err error) {
    return c.DescribeBackupDownloadRestrictionWithContext(context.Background(), request)
}

// DescribeBackupDownloadRestriction
// 本接口（DescribeBackupDownloadRestriction）用于查询当前地域数据库备份文件的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewDescribeBackupUrlRequest() (request *DescribeBackupUrlRequest) {
    request = &DescribeBackupUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBackupUrl")
    
    
    return
}

func NewDescribeBackupUrlResponse() (response *DescribeBackupUrlResponse) {
    response = &DescribeBackupUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBackupUrl
// 本接口（DescribeBackupUrl）用于查询备份 Rdb 文件的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSINVALID = "ResourceUnavailable.BackupStatusInvalid"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) DescribeBackupUrl(request *DescribeBackupUrlRequest) (response *DescribeBackupUrlResponse, err error) {
    return c.DescribeBackupUrlWithContext(context.Background(), request)
}

// DescribeBackupUrl
// 本接口（DescribeBackupUrl）用于查询备份 Rdb 文件的下载地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSINVALID = "ResourceUnavailable.BackupStatusInvalid"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) DescribeBackupUrlWithContext(ctx context.Context, request *DescribeBackupUrlRequest) (response *DescribeBackupUrlResponse, err error) {
    if request == nil {
        request = NewDescribeBackupUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBackupUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBackupUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBandwidthRangeRequest() (request *DescribeBandwidthRangeRequest) {
    request = &DescribeBandwidthRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeBandwidthRange")
    
    
    return
}

func NewDescribeBandwidthRangeResponse() (response *DescribeBandwidthRangeResponse) {
    response = &DescribeBandwidthRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBandwidthRange
// 本接口（DescribeBandwidthRange）用于查询实例带宽信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeBandwidthRange(request *DescribeBandwidthRangeRequest) (response *DescribeBandwidthRangeResponse, err error) {
    return c.DescribeBandwidthRangeWithContext(context.Background(), request)
}

// DescribeBandwidthRange
// 本接口（DescribeBandwidthRange）用于查询实例带宽信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeBandwidthRangeWithContext(ctx context.Context, request *DescribeBandwidthRangeRequest) (response *DescribeBandwidthRangeResponse, err error) {
    if request == nil {
        request = NewDescribeBandwidthRangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBandwidthRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBandwidthRangeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCommonDBInstancesRequest() (request *DescribeCommonDBInstancesRequest) {
    request = &DescribeCommonDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeCommonDBInstances")
    
    
    return
}

func NewDescribeCommonDBInstancesResponse() (response *DescribeCommonDBInstancesResponse) {
    response = &DescribeCommonDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCommonDBInstances
// 查询Redis实例列表信息。该接口已废弃。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeCommonDBInstances(request *DescribeCommonDBInstancesRequest) (response *DescribeCommonDBInstancesResponse, err error) {
    return c.DescribeCommonDBInstancesWithContext(context.Background(), request)
}

// DescribeCommonDBInstances
// 查询Redis实例列表信息。该接口已废弃。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeCommonDBInstancesWithContext(ctx context.Context, request *DescribeCommonDBInstancesRequest) (response *DescribeCommonDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCommonDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCommonDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCommonDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
    request = &DescribeDBSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeDBSecurityGroups")
    
    
    return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
    response = &DescribeDBSecurityGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSecurityGroups
// 本接口（DescribeDBSecurityGroups）用于查询实例的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
    return c.DescribeDBSecurityGroupsWithContext(context.Background(), request)
}

// DescribeDBSecurityGroups
// 本接口（DescribeDBSecurityGroups）用于查询实例的安全组详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"
//  FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"
//  FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"
//  FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
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

func NewDescribeGlobalReplicationAreaRequest() (request *DescribeGlobalReplicationAreaRequest) {
    request = &DescribeGlobalReplicationAreaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeGlobalReplicationArea")
    
    
    return
}

func NewDescribeGlobalReplicationAreaResponse() (response *DescribeGlobalReplicationAreaResponse) {
    response = &DescribeGlobalReplicationAreaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlobalReplicationArea
// 查询全球复制支持地域信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeGlobalReplicationArea(request *DescribeGlobalReplicationAreaRequest) (response *DescribeGlobalReplicationAreaResponse, err error) {
    return c.DescribeGlobalReplicationAreaWithContext(context.Background(), request)
}

// DescribeGlobalReplicationArea
// 查询全球复制支持地域信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeGlobalReplicationAreaWithContext(ctx context.Context, request *DescribeGlobalReplicationAreaRequest) (response *DescribeGlobalReplicationAreaResponse, err error) {
    if request == nil {
        request = NewDescribeGlobalReplicationAreaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlobalReplicationArea require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlobalReplicationAreaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAccountRequest() (request *DescribeInstanceAccountRequest) {
    request = &DescribeInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceAccount")
    
    
    return
}

func NewDescribeInstanceAccountResponse() (response *DescribeInstanceAccountResponse) {
    response = &DescribeInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceAccount
// 本接口（DescribeInstanceAccount）用于查看实例子账号信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeInstanceAccount(request *DescribeInstanceAccountRequest) (response *DescribeInstanceAccountResponse, err error) {
    return c.DescribeInstanceAccountWithContext(context.Background(), request)
}

// DescribeInstanceAccount
// 本接口（DescribeInstanceAccount）用于查看实例子账号信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeInstanceAccountWithContext(ctx context.Context, request *DescribeInstanceAccountRequest) (response *DescribeInstanceAccountResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceBackupsRequest() (request *DescribeInstanceBackupsRequest) {
    request = &DescribeInstanceBackupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceBackups")
    
    
    return
}

func NewDescribeInstanceBackupsResponse() (response *DescribeInstanceBackupsResponse) {
    response = &DescribeInstanceBackupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceBackups
// 本接口（DescribeInstanceBackups）用于查询实例备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceBackups(request *DescribeInstanceBackupsRequest) (response *DescribeInstanceBackupsResponse, err error) {
    return c.DescribeInstanceBackupsWithContext(context.Background(), request)
}

// DescribeInstanceBackups
// 本接口（DescribeInstanceBackups）用于查询实例备份列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewDescribeInstanceDTSInfoRequest() (request *DescribeInstanceDTSInfoRequest) {
    request = &DescribeInstanceDTSInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceDTSInfo")
    
    
    return
}

func NewDescribeInstanceDTSInfoResponse() (response *DescribeInstanceDTSInfoResponse) {
    response = &DescribeInstanceDTSInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceDTSInfo
// 查询实例DTS信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceDTSInfo(request *DescribeInstanceDTSInfoRequest) (response *DescribeInstanceDTSInfoResponse, err error) {
    return c.DescribeInstanceDTSInfoWithContext(context.Background(), request)
}

// DescribeInstanceDTSInfo
// 查询实例DTS信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceDTSInfoWithContext(ctx context.Context, request *DescribeInstanceDTSInfoRequest) (response *DescribeInstanceDTSInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceDTSInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceDTSInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceDTSInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceDealDetailRequest() (request *DescribeInstanceDealDetailRequest) {
    request = &DescribeInstanceDealDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceDealDetail")
    
    
    return
}

func NewDescribeInstanceDealDetailResponse() (response *DescribeInstanceDealDetailResponse) {
    response = &DescribeInstanceDealDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceDealDetail
// 本接口（DescribeInstanceDealDetail）用于查询订单信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) DescribeInstanceDealDetail(request *DescribeInstanceDealDetailRequest) (response *DescribeInstanceDealDetailResponse, err error) {
    return c.DescribeInstanceDealDetailWithContext(context.Background(), request)
}

// DescribeInstanceDealDetail
// 本接口（DescribeInstanceDealDetail）用于查询订单信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
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

func NewDescribeInstanceEventsRequest() (request *DescribeInstanceEventsRequest) {
    request = &DescribeInstanceEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceEvents")
    
    
    return
}

func NewDescribeInstanceEventsResponse() (response *DescribeInstanceEventsResponse) {
    response = &DescribeInstanceEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceEvents
// 本接口（DescribeInstanceEvents）用于查询 Redis 实例事件信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) DescribeInstanceEvents(request *DescribeInstanceEventsRequest) (response *DescribeInstanceEventsResponse, err error) {
    return c.DescribeInstanceEventsWithContext(context.Background(), request)
}

// DescribeInstanceEvents
// 本接口（DescribeInstanceEvents）用于查询 Redis 实例事件信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) DescribeInstanceEventsWithContext(ctx context.Context, request *DescribeInstanceEventsRequest) (response *DescribeInstanceEventsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogDeliveryRequest() (request *DescribeInstanceLogDeliveryRequest) {
    request = &DescribeInstanceLogDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceLogDelivery")
    
    
    return
}

func NewDescribeInstanceLogDeliveryResponse() (response *DescribeInstanceLogDeliveryResponse) {
    response = &DescribeInstanceLogDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceLogDelivery
// 本接口（DescribeInstanceLogDelivery）用于查询实例的日志投递配置。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceLogDelivery(request *DescribeInstanceLogDeliveryRequest) (response *DescribeInstanceLogDeliveryResponse, err error) {
    return c.DescribeInstanceLogDeliveryWithContext(context.Background(), request)
}

// DescribeInstanceLogDelivery
// 本接口（DescribeInstanceLogDelivery）用于查询实例的日志投递配置。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceLogDeliveryWithContext(ctx context.Context, request *DescribeInstanceLogDeliveryRequest) (response *DescribeInstanceLogDeliveryResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogDeliveryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeyRequest() (request *DescribeInstanceMonitorBigKeyRequest) {
    request = &DescribeInstanceMonitorBigKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKey")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeyResponse() (response *DescribeInstanceMonitorBigKeyResponse) {
    response = &DescribeInstanceMonitorBigKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorBigKey
// 腾讯云数据库 Redis 已经于2022年10月31日下线查询实例大 Key 接口。具体公告，请参见[查询实例大 Key 接口下线公告](https://cloud.tencent.com/document/product/239/81005)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKey(request *DescribeInstanceMonitorBigKeyRequest) (response *DescribeInstanceMonitorBigKeyResponse, err error) {
    return c.DescribeInstanceMonitorBigKeyWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKey
// 腾讯云数据库 Redis 已经于2022年10月31日下线查询实例大 Key 接口。具体公告，请参见[查询实例大 Key 接口下线公告](https://cloud.tencent.com/document/product/239/81005)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeyRequest) (response *DescribeInstanceMonitorBigKeyResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeySizeDistRequest() (request *DescribeInstanceMonitorBigKeySizeDistRequest) {
    request = &DescribeInstanceMonitorBigKeySizeDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKeySizeDist")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeySizeDistResponse() (response *DescribeInstanceMonitorBigKeySizeDistResponse) {
    response = &DescribeInstanceMonitorBigKeySizeDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorBigKeySizeDist
// 腾讯云数据库 Redis 已经于2022年10月31日下线查询实例大 Key 接口。具体公告，请参见 [查询实例大 Key 接口下线公告](https://cloud.tencent.com/document/product/239/81005)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeySizeDist(request *DescribeInstanceMonitorBigKeySizeDistRequest) (response *DescribeInstanceMonitorBigKeySizeDistResponse, err error) {
    return c.DescribeInstanceMonitorBigKeySizeDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKeySizeDist
// 腾讯云数据库 Redis 已经于2022年10月31日下线查询实例大 Key 接口。具体公告，请参见 [查询实例大 Key 接口下线公告](https://cloud.tencent.com/document/product/239/81005)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeySizeDistWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeySizeDistRequest) (response *DescribeInstanceMonitorBigKeySizeDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeySizeDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKeySizeDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeySizeDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorBigKeyTypeDistRequest() (request *DescribeInstanceMonitorBigKeyTypeDistRequest) {
    request = &DescribeInstanceMonitorBigKeyTypeDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorBigKeyTypeDist")
    
    
    return
}

func NewDescribeInstanceMonitorBigKeyTypeDistResponse() (response *DescribeInstanceMonitorBigKeyTypeDistResponse) {
    response = &DescribeInstanceMonitorBigKeyTypeDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorBigKeyTypeDist
// 腾讯云数据库 Redis 已经于2022年10月31日下线查询实例大 Key 接口。具体公告，请参见 [查询实例大 Key 接口下线公告](https://cloud.tencent.com/document/product/239/81005)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyTypeDist(request *DescribeInstanceMonitorBigKeyTypeDistRequest) (response *DescribeInstanceMonitorBigKeyTypeDistResponse, err error) {
    return c.DescribeInstanceMonitorBigKeyTypeDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorBigKeyTypeDist
// 腾讯云数据库 Redis 已经于2022年10月31日下线查询实例大 Key 接口。具体公告，请参见 [查询实例大 Key 接口下线公告](https://cloud.tencent.com/document/product/239/81005)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorBigKeyTypeDistWithContext(ctx context.Context, request *DescribeInstanceMonitorBigKeyTypeDistRequest) (response *DescribeInstanceMonitorBigKeyTypeDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorBigKeyTypeDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorBigKeyTypeDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorBigKeyTypeDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorHotKeyRequest() (request *DescribeInstanceMonitorHotKeyRequest) {
    request = &DescribeInstanceMonitorHotKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorHotKey")
    
    
    return
}

func NewDescribeInstanceMonitorHotKeyResponse() (response *DescribeInstanceMonitorHotKeyResponse) {
    response = &DescribeInstanceMonitorHotKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorHotKey
// 本接口（DescribeInstanceMonitorHotKey）用于查询实例热Key。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorHotKey(request *DescribeInstanceMonitorHotKeyRequest) (response *DescribeInstanceMonitorHotKeyResponse, err error) {
    return c.DescribeInstanceMonitorHotKeyWithContext(context.Background(), request)
}

// DescribeInstanceMonitorHotKey
// 本接口（DescribeInstanceMonitorHotKey）用于查询实例热Key。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorHotKeyWithContext(ctx context.Context, request *DescribeInstanceMonitorHotKeyRequest) (response *DescribeInstanceMonitorHotKeyResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorHotKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorHotKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorHotKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorSIPRequest() (request *DescribeInstanceMonitorSIPRequest) {
    request = &DescribeInstanceMonitorSIPRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorSIP")
    
    
    return
}

func NewDescribeInstanceMonitorSIPResponse() (response *DescribeInstanceMonitorSIPResponse) {
    response = &DescribeInstanceMonitorSIPResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorSIP
// 该接口已下线，请使用数据库智能管家 DBbrain 接口 [DescribeProxyProcessStatistics] (https://cloud.tencent.com/document/product/1130/84544) 获取实例访问来源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorSIP(request *DescribeInstanceMonitorSIPRequest) (response *DescribeInstanceMonitorSIPResponse, err error) {
    return c.DescribeInstanceMonitorSIPWithContext(context.Background(), request)
}

// DescribeInstanceMonitorSIP
// 该接口已下线，请使用数据库智能管家 DBbrain 接口 [DescribeProxyProcessStatistics] (https://cloud.tencent.com/document/product/1130/84544) 获取实例访问来源。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorSIPWithContext(ctx context.Context, request *DescribeInstanceMonitorSIPRequest) (response *DescribeInstanceMonitorSIPResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorSIPRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorSIP require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorSIPResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTookDistRequest() (request *DescribeInstanceMonitorTookDistRequest) {
    request = &DescribeInstanceMonitorTookDistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTookDist")
    
    
    return
}

func NewDescribeInstanceMonitorTookDistResponse() (response *DescribeInstanceMonitorTookDistResponse) {
    response = &DescribeInstanceMonitorTookDistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorTookDist
// 查询实例访问的耗时分布
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTookDist(request *DescribeInstanceMonitorTookDistRequest) (response *DescribeInstanceMonitorTookDistResponse, err error) {
    return c.DescribeInstanceMonitorTookDistWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTookDist
// 查询实例访问的耗时分布
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTookDistWithContext(ctx context.Context, request *DescribeInstanceMonitorTookDistRequest) (response *DescribeInstanceMonitorTookDistResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTookDistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTookDist require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTookDistResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTopNCmdRequest() (request *DescribeInstanceMonitorTopNCmdRequest) {
    request = &DescribeInstanceMonitorTopNCmdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTopNCmd")
    
    
    return
}

func NewDescribeInstanceMonitorTopNCmdResponse() (response *DescribeInstanceMonitorTopNCmdResponse) {
    response = &DescribeInstanceMonitorTopNCmdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorTopNCmd
// 查询实例访问命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmd(request *DescribeInstanceMonitorTopNCmdRequest) (response *DescribeInstanceMonitorTopNCmdResponse, err error) {
    return c.DescribeInstanceMonitorTopNCmdWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTopNCmd
// 查询实例访问命令
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdWithContext(ctx context.Context, request *DescribeInstanceMonitorTopNCmdRequest) (response *DescribeInstanceMonitorTopNCmdResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTopNCmdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTopNCmd require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTopNCmdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceMonitorTopNCmdTookRequest() (request *DescribeInstanceMonitorTopNCmdTookRequest) {
    request = &DescribeInstanceMonitorTopNCmdTookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceMonitorTopNCmdTook")
    
    
    return
}

func NewDescribeInstanceMonitorTopNCmdTookResponse() (response *DescribeInstanceMonitorTopNCmdTookResponse) {
    response = &DescribeInstanceMonitorTopNCmdTookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceMonitorTopNCmdTook
// 查询实例CPU耗时
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdTook(request *DescribeInstanceMonitorTopNCmdTookRequest) (response *DescribeInstanceMonitorTopNCmdTookResponse, err error) {
    return c.DescribeInstanceMonitorTopNCmdTookWithContext(context.Background(), request)
}

// DescribeInstanceMonitorTopNCmdTook
// 查询实例CPU耗时
//
// 可能返回的错误码:
//  FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"
func (c *Client) DescribeInstanceMonitorTopNCmdTookWithContext(ctx context.Context, request *DescribeInstanceMonitorTopNCmdTookRequest) (response *DescribeInstanceMonitorTopNCmdTookResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceMonitorTopNCmdTookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceMonitorTopNCmdTook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceMonitorTopNCmdTookResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceNodeInfoRequest() (request *DescribeInstanceNodeInfoRequest) {
    request = &DescribeInstanceNodeInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceNodeInfo")
    
    
    return
}

func NewDescribeInstanceNodeInfoResponse() (response *DescribeInstanceNodeInfoResponse) {
    response = &DescribeInstanceNodeInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceNodeInfo
// 本接口（DescribeInstanceNodeInfo）用于查询实例节点信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceNodeInfo(request *DescribeInstanceNodeInfoRequest) (response *DescribeInstanceNodeInfoResponse, err error) {
    return c.DescribeInstanceNodeInfoWithContext(context.Background(), request)
}

// DescribeInstanceNodeInfo
// 本接口（DescribeInstanceNodeInfo）用于查询实例节点信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceParamRecords")
    
    
    return
}

func NewDescribeInstanceParamRecordsResponse() (response *DescribeInstanceParamRecordsResponse) {
    response = &DescribeInstanceParamRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceParamRecords
// 查询参数修改历史列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeInstanceParamRecords(request *DescribeInstanceParamRecordsRequest) (response *DescribeInstanceParamRecordsResponse, err error) {
    return c.DescribeInstanceParamRecordsWithContext(context.Background(), request)
}

// DescribeInstanceParamRecords
// 查询参数修改历史列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceParams")
    
    
    return
}

func NewDescribeInstanceParamsResponse() (response *DescribeInstanceParamsResponse) {
    response = &DescribeInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询实例参数列表。
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeInstanceParams(request *DescribeInstanceParamsRequest) (response *DescribeInstanceParamsResponse, err error) {
    return c.DescribeInstanceParamsWithContext(context.Background(), request)
}

// DescribeInstanceParams
// 本接口（DescribeInstanceParams）用于查询实例参数列表。
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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

func NewDescribeInstanceSecurityGroupRequest() (request *DescribeInstanceSecurityGroupRequest) {
    request = &DescribeInstanceSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceSecurityGroup")
    
    
    return
}

func NewDescribeInstanceSecurityGroupResponse() (response *DescribeInstanceSecurityGroupResponse) {
    response = &DescribeInstanceSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSecurityGroup
// 本接口（DescribeInstanceSecurityGroup）用于查询实例安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceSecurityGroup(request *DescribeInstanceSecurityGroupRequest) (response *DescribeInstanceSecurityGroupResponse, err error) {
    return c.DescribeInstanceSecurityGroupWithContext(context.Background(), request)
}

// DescribeInstanceSecurityGroup
// 本接口（DescribeInstanceSecurityGroup）用于查询实例安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceSecurityGroupWithContext(ctx context.Context, request *DescribeInstanceSecurityGroupRequest) (response *DescribeInstanceSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceShardsRequest() (request *DescribeInstanceShardsRequest) {
    request = &DescribeInstanceShardsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceShards")
    
    
    return
}

func NewDescribeInstanceShardsResponse() (response *DescribeInstanceShardsResponse) {
    response = &DescribeInstanceShardsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceShards
// 本接口（DescribeInstanceShards）用于获取集群架构实例的分片信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
func (c *Client) DescribeInstanceShards(request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    return c.DescribeInstanceShardsWithContext(context.Background(), request)
}

// DescribeInstanceShards
// 本接口（DescribeInstanceShards）用于获取集群架构实例的分片信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
func (c *Client) DescribeInstanceShardsWithContext(ctx context.Context, request *DescribeInstanceShardsRequest) (response *DescribeInstanceShardsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceShardsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceShards require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceShardsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSpecBandwidthRequest() (request *DescribeInstanceSpecBandwidthRequest) {
    request = &DescribeInstanceSpecBandwidthRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceSpecBandwidth")
    
    
    return
}

func NewDescribeInstanceSpecBandwidthResponse() (response *DescribeInstanceSpecBandwidthResponse) {
    response = &DescribeInstanceSpecBandwidthResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSpecBandwidth
// 本接口（DescribeInstanceSpecBandwidth）用于查询或计算带宽规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceSpecBandwidth(request *DescribeInstanceSpecBandwidthRequest) (response *DescribeInstanceSpecBandwidthResponse, err error) {
    return c.DescribeInstanceSpecBandwidthWithContext(context.Background(), request)
}

// DescribeInstanceSpecBandwidth
// 本接口（DescribeInstanceSpecBandwidth）用于查询或计算带宽规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceSpecBandwidthWithContext(ctx context.Context, request *DescribeInstanceSpecBandwidthRequest) (response *DescribeInstanceSpecBandwidthResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSpecBandwidthRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSpecBandwidth require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSpecBandwidthResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceSupportFeatureRequest() (request *DescribeInstanceSupportFeatureRequest) {
    request = &DescribeInstanceSupportFeatureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceSupportFeature")
    
    
    return
}

func NewDescribeInstanceSupportFeatureResponse() (response *DescribeInstanceSupportFeatureResponse) {
    response = &DescribeInstanceSupportFeatureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceSupportFeature
// 本接口（DescribeInstanceSupportFeature）用于查询实例支持的功能特性。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceSupportFeature(request *DescribeInstanceSupportFeatureRequest) (response *DescribeInstanceSupportFeatureResponse, err error) {
    return c.DescribeInstanceSupportFeatureWithContext(context.Background(), request)
}

// DescribeInstanceSupportFeature
// 本接口（DescribeInstanceSupportFeature）用于查询实例支持的功能特性。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) DescribeInstanceSupportFeatureWithContext(ctx context.Context, request *DescribeInstanceSupportFeatureRequest) (response *DescribeInstanceSupportFeatureResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceSupportFeatureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceSupportFeature require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceSupportFeatureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceZoneInfoRequest() (request *DescribeInstanceZoneInfoRequest) {
    request = &DescribeInstanceZoneInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstanceZoneInfo")
    
    
    return
}

func NewDescribeInstanceZoneInfoResponse() (response *DescribeInstanceZoneInfoResponse) {
    response = &DescribeInstanceZoneInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstanceZoneInfo
// 本接口（DescribeInstanceZoneInfo）用于查询 Redis 节点详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceZoneInfo(request *DescribeInstanceZoneInfoRequest) (response *DescribeInstanceZoneInfoResponse, err error) {
    return c.DescribeInstanceZoneInfoWithContext(context.Background(), request)
}

// DescribeInstanceZoneInfo
// 本接口（DescribeInstanceZoneInfo）用于查询 Redis 节点详细信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeInstanceZoneInfoWithContext(ctx context.Context, request *DescribeInstanceZoneInfoRequest) (response *DescribeInstanceZoneInfoResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceZoneInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceZoneInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceZoneInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeInstances
// 本接口（DescribeInstances）用于查询Redis实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 本接口（DescribeInstances）用于查询Redis实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeMaintenanceWindow")
    
    
    return
}

func NewDescribeMaintenanceWindowResponse() (response *DescribeMaintenanceWindowResponse) {
    response = &DescribeMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMaintenanceWindow
// 本接口（DescribeMaintenanceWindow）用于查询实例维护时间窗。在实例需要进行版本升级或者架构升级的时候，会在维护时间窗时间内进行切换
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeMaintenanceWindow(request *DescribeMaintenanceWindowRequest) (response *DescribeMaintenanceWindowResponse, err error) {
    return c.DescribeMaintenanceWindowWithContext(context.Background(), request)
}

// DescribeMaintenanceWindow
// 本接口（DescribeMaintenanceWindow）用于查询实例维护时间窗。在实例需要进行版本升级或者架构升级的时候，会在维护时间窗时间内进行切换
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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

func NewDescribeParamTemplateInfoRequest() (request *DescribeParamTemplateInfoRequest) {
    request = &DescribeParamTemplateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeParamTemplateInfo")
    
    
    return
}

func NewDescribeParamTemplateInfoResponse() (response *DescribeParamTemplateInfoResponse) {
    response = &DescribeParamTemplateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamTemplateInfo
// 本接口（DescribeParamTemplateInfo）用于查询参数模板详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
func (c *Client) DescribeParamTemplateInfo(request *DescribeParamTemplateInfoRequest) (response *DescribeParamTemplateInfoResponse, err error) {
    return c.DescribeParamTemplateInfoWithContext(context.Background(), request)
}

// DescribeParamTemplateInfo
// 本接口（DescribeParamTemplateInfo）用于查询参数模板详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeParamTemplates")
    
    
    return
}

func NewDescribeParamTemplatesResponse() (response *DescribeParamTemplatesResponse) {
    response = &DescribeParamTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeParamTemplates
// 查询参数模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeParamTemplates(request *DescribeParamTemplatesRequest) (response *DescribeParamTemplatesResponse, err error) {
    return c.DescribeParamTemplatesWithContext(context.Background(), request)
}

// DescribeParamTemplates
// 查询参数模板列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewDescribeProductInfoRequest() (request *DescribeProductInfoRequest) {
    request = &DescribeProductInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProductInfo")
    
    
    return
}

func NewDescribeProductInfoResponse() (response *DescribeProductInfoResponse) {
    response = &DescribeProductInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProductInfo
// 本接口（DescribeProductInfo）用于查询全地域 Redis 的售卖规格。
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
// 本接口（DescribeProductInfo）用于查询全地域 Redis 的售卖规格。
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

func NewDescribeProjectSecurityGroupRequest() (request *DescribeProjectSecurityGroupRequest) {
    request = &DescribeProjectSecurityGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProjectSecurityGroup")
    
    
    return
}

func NewDescribeProjectSecurityGroupResponse() (response *DescribeProjectSecurityGroupResponse) {
    response = &DescribeProjectSecurityGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjectSecurityGroup
// 查询项目安全组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  RESOURCEUNAVAILABLE_SECURITYGROUPNOTSUPPORTED = "ResourceUnavailable.SecurityGroupNotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeProjectSecurityGroup(request *DescribeProjectSecurityGroupRequest) (response *DescribeProjectSecurityGroupResponse, err error) {
    return c.DescribeProjectSecurityGroupWithContext(context.Background(), request)
}

// DescribeProjectSecurityGroup
// 查询项目安全组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  RESOURCEUNAVAILABLE_SECURITYGROUPNOTSUPPORTED = "ResourceUnavailable.SecurityGroupNotSupported"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DescribeProjectSecurityGroupWithContext(ctx context.Context, request *DescribeProjectSecurityGroupRequest) (response *DescribeProjectSecurityGroupResponse, err error) {
    if request == nil {
        request = NewDescribeProjectSecurityGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjectSecurityGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectSecurityGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectSecurityGroupsRequest() (request *DescribeProjectSecurityGroupsRequest) {
    request = &DescribeProjectSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProjectSecurityGroups")
    
    
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
//  FAILEDOPERATION_FLOWNOTEXISTS = "FailedOperation.FlowNotExists"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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
//  FAILEDOPERATION_FLOWNOTEXISTS = "FailedOperation.FlowNotExists"
//  FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
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

func NewDescribeProxySlowLogRequest() (request *DescribeProxySlowLogRequest) {
    request = &DescribeProxySlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeProxySlowLog")
    
    
    return
}

func NewDescribeProxySlowLogResponse() (response *DescribeProxySlowLogResponse) {
    response = &DescribeProxySlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxySlowLog
// 本接口（DescribeProxySlowLog）用于查询代理慢查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeProxySlowLog(request *DescribeProxySlowLogRequest) (response *DescribeProxySlowLogResponse, err error) {
    return c.DescribeProxySlowLogWithContext(context.Background(), request)
}

// DescribeProxySlowLog
// 本接口（DescribeProxySlowLog）用于查询代理慢查询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
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

func NewDescribeRedisClusterOverviewRequest() (request *DescribeRedisClusterOverviewRequest) {
    request = &DescribeRedisClusterOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeRedisClusterOverview")
    
    
    return
}

func NewDescribeRedisClusterOverviewResponse() (response *DescribeRedisClusterOverviewResponse) {
    response = &DescribeRedisClusterOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisClusterOverview
// 查询Redis独享集群概览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) DescribeRedisClusterOverview(request *DescribeRedisClusterOverviewRequest) (response *DescribeRedisClusterOverviewResponse, err error) {
    return c.DescribeRedisClusterOverviewWithContext(context.Background(), request)
}

// DescribeRedisClusterOverview
// 查询Redis独享集群概览信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
func (c *Client) DescribeRedisClusterOverviewWithContext(ctx context.Context, request *DescribeRedisClusterOverviewRequest) (response *DescribeRedisClusterOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeRedisClusterOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisClusterOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisClusterOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisClustersRequest() (request *DescribeRedisClustersRequest) {
    request = &DescribeRedisClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeRedisClusters")
    
    
    return
}

func NewDescribeRedisClustersResponse() (response *DescribeRedisClustersResponse) {
    response = &DescribeRedisClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisClusters
// 查询Redis独享集群列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRedisClusters(request *DescribeRedisClustersRequest) (response *DescribeRedisClustersResponse, err error) {
    return c.DescribeRedisClustersWithContext(context.Background(), request)
}

// DescribeRedisClusters
// 查询Redis独享集群列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRedisClustersWithContext(ctx context.Context, request *DescribeRedisClustersRequest) (response *DescribeRedisClustersResponse, err error) {
    if request == nil {
        request = NewDescribeRedisClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisClustersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationGroupRequest() (request *DescribeReplicationGroupRequest) {
    request = &DescribeReplicationGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeReplicationGroup")
    
    
    return
}

func NewDescribeReplicationGroupResponse() (response *DescribeReplicationGroupResponse) {
    response = &DescribeReplicationGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplicationGroup
// 本接口（DescribeReplicationGroup）用于查询复制组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
func (c *Client) DescribeReplicationGroup(request *DescribeReplicationGroupRequest) (response *DescribeReplicationGroupResponse, err error) {
    return c.DescribeReplicationGroupWithContext(context.Background(), request)
}

// DescribeReplicationGroup
// 本接口（DescribeReplicationGroup）用于查询复制组。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
func (c *Client) DescribeReplicationGroupWithContext(ctx context.Context, request *DescribeReplicationGroupRequest) (response *DescribeReplicationGroupResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReplicationGroupInstanceRequest() (request *DescribeReplicationGroupInstanceRequest) {
    request = &DescribeReplicationGroupInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeReplicationGroupInstance")
    
    
    return
}

func NewDescribeReplicationGroupInstanceResponse() (response *DescribeReplicationGroupInstanceResponse) {
    response = &DescribeReplicationGroupInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReplicationGroupInstance
// 查询复制组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeReplicationGroupInstance(request *DescribeReplicationGroupInstanceRequest) (response *DescribeReplicationGroupInstanceResponse, err error) {
    return c.DescribeReplicationGroupInstanceWithContext(context.Background(), request)
}

// DescribeReplicationGroupInstance
// 查询复制组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DescribeReplicationGroupInstanceWithContext(ctx context.Context, request *DescribeReplicationGroupInstanceRequest) (response *DescribeReplicationGroupInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeReplicationGroupInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReplicationGroupInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReplicationGroupInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSSLStatusRequest() (request *DescribeSSLStatusRequest) {
    request = &DescribeSSLStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeSSLStatus")
    
    
    return
}

func NewDescribeSSLStatusResponse() (response *DescribeSSLStatusResponse) {
    response = &DescribeSSLStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSSLStatus
// 本接口（DescribeSSLStatus）用于查询实例 SSL 认证相关信息，包括开启状态、配置状态、证书地址等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSSLStatus(request *DescribeSSLStatusRequest) (response *DescribeSSLStatusResponse, err error) {
    return c.DescribeSSLStatusWithContext(context.Background(), request)
}

// DescribeSSLStatus
// 本接口（DescribeSSLStatus）用于查询实例 SSL 认证相关信息，包括开启状态、配置状态、证书地址等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ACTIONNOTFOUND = "InvalidParameter.ActionNotFound"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSSLStatusWithContext(ctx context.Context, request *DescribeSSLStatusRequest) (response *DescribeSSLStatusResponse, err error) {
    if request == nil {
        request = NewDescribeSSLStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSSLStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSSLStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogRequest() (request *DescribeSlowLogRequest) {
    request = &DescribeSlowLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeSlowLog")
    
    
    return
}

func NewDescribeSlowLogResponse() (response *DescribeSlowLogResponse) {
    response = &DescribeSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLog
// 本接口（DescribeSlowLog）查询实例慢查询记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSlowLog(request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    return c.DescribeSlowLogWithContext(context.Background(), request)
}

// DescribeSlowLog
// 本接口（DescribeSlowLog）查询实例慢查询记录。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeSlowLogWithContext(ctx context.Context, request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
    request = &DescribeTaskInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTaskInfo")
    
    
    return
}

func NewDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
    response = &DescribeTaskInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskInfo
// 本接口（DescribeTaskInfo）用于获取指定任务的执行情况。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
    return c.DescribeTaskInfoWithContext(context.Background(), request)
}

// DescribeTaskInfo
// 本接口（DescribeTaskInfo）用于获取指定任务的执行情况。
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTaskList")
    
    
    return
}

func NewDescribeTaskListResponse() (response *DescribeTaskListResponse) {
    response = &DescribeTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTaskList
// 本接口（DescribeTaskList）用于查询指定实例的任务列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) DescribeTaskList(request *DescribeTaskListRequest) (response *DescribeTaskListResponse, err error) {
    return c.DescribeTaskListWithContext(context.Background(), request)
}

// DescribeTaskList
// 本接口（DescribeTaskList）用于查询指定实例的任务列表信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DescribeTendisSlowLog")
    
    
    return
}

func NewDescribeTendisSlowLogResponse() (response *DescribeTendisSlowLogResponse) {
    response = &DescribeTendisSlowLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTendisSlowLog
// 查询Tendis慢查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) DescribeTendisSlowLog(request *DescribeTendisSlowLogRequest) (response *DescribeTendisSlowLogResponse, err error) {
    return c.DescribeTendisSlowLogWithContext(context.Background(), request)
}

// DescribeTendisSlowLog
// 查询Tendis慢查询
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DestroyPostpaidInstance")
    
    
    return
}

func NewDestroyPostpaidInstanceResponse() (response *DestroyPostpaidInstanceResponse) {
    response = &DestroyPostpaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyPostpaidInstance
// 按量计费实例销毁
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) DestroyPostpaidInstance(request *DestroyPostpaidInstanceRequest) (response *DestroyPostpaidInstanceResponse, err error) {
    return c.DestroyPostpaidInstanceWithContext(context.Background(), request)
}

// DestroyPostpaidInstance
// 按量计费实例销毁
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "DestroyPrepaidInstance")
    
    
    return
}

func NewDestroyPrepaidInstanceResponse() (response *DestroyPrepaidInstanceResponse) {
    response = &DestroyPrepaidInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyPrepaidInstance
// 包年包月实例退还
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEISOLATED = "ResourceUnavailable.InstanceIsolated"
//  RESOURCEUNAVAILABLE_INSTANCENODEAL = "ResourceUnavailable.InstanceNoDeal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) DestroyPrepaidInstance(request *DestroyPrepaidInstanceRequest) (response *DestroyPrepaidInstanceResponse, err error) {
    return c.DestroyPrepaidInstanceWithContext(context.Background(), request)
}

// DestroyPrepaidInstance
// 包年包月实例退还
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
//  RESOURCEUNAVAILABLE_INSTANCEISOLATED = "ResourceUnavailable.InstanceIsolated"
//  RESOURCEUNAVAILABLE_INSTANCENODEAL = "ResourceUnavailable.InstanceNoDeal"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewDisableReplicaReadonlyRequest() (request *DisableReplicaReadonlyRequest) {
    request = &DisableReplicaReadonlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DisableReplicaReadonly")
    
    
    return
}

func NewDisableReplicaReadonlyResponse() (response *DisableReplicaReadonlyResponse) {
    response = &DisableReplicaReadonlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableReplicaReadonly
// 禁用读写分离
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
func (c *Client) DisableReplicaReadonly(request *DisableReplicaReadonlyRequest) (response *DisableReplicaReadonlyResponse, err error) {
    return c.DisableReplicaReadonlyWithContext(context.Background(), request)
}

// DisableReplicaReadonly
// 禁用读写分离
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
func (c *Client) DisableReplicaReadonlyWithContext(ctx context.Context, request *DisableReplicaReadonlyRequest) (response *DisableReplicaReadonlyResponse, err error) {
    if request == nil {
        request = NewDisableReplicaReadonlyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableReplicaReadonly require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableReplicaReadonlyResponse()
    err = c.Send(request, response)
    return
}

func NewDisassociateSecurityGroupsRequest() (request *DisassociateSecurityGroupsRequest) {
    request = &DisassociateSecurityGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "DisassociateSecurityGroups")
    
    
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
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"
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

func NewEnableReplicaReadonlyRequest() (request *EnableReplicaReadonlyRequest) {
    request = &EnableReplicaReadonlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "EnableReplicaReadonly")
    
    
    return
}

func NewEnableReplicaReadonlyResponse() (response *EnableReplicaReadonlyResponse) {
    response = &EnableReplicaReadonlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableReplicaReadonly
// 启用读写分离
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) EnableReplicaReadonly(request *EnableReplicaReadonlyRequest) (response *EnableReplicaReadonlyResponse, err error) {
    return c.EnableReplicaReadonlyWithContext(context.Background(), request)
}

// EnableReplicaReadonly
// 启用读写分离
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"
//  RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"
func (c *Client) EnableReplicaReadonlyWithContext(ctx context.Context, request *EnableReplicaReadonlyRequest) (response *EnableReplicaReadonlyResponse, err error) {
    if request == nil {
        request = NewEnableReplicaReadonlyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableReplicaReadonly require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableReplicaReadonlyResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateInstanceRequest() (request *InquiryPriceCreateInstanceRequest) {
    request = &InquiryPriceCreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceCreateInstance")
    
    
    return
}

func NewInquiryPriceCreateInstanceResponse() (response *InquiryPriceCreateInstanceResponse) {
    response = &InquiryPriceCreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceCreateInstance
// 查询新购实例价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) InquiryPriceCreateInstance(request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    return c.InquiryPriceCreateInstanceWithContext(context.Background(), request)
}

// InquiryPriceCreateInstance
// 查询新购实例价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) InquiryPriceCreateInstanceWithContext(ctx context.Context, request *InquiryPriceCreateInstanceRequest) (response *InquiryPriceCreateInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceRenewInstanceRequest() (request *InquiryPriceRenewInstanceRequest) {
    request = &InquiryPriceRenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceRenewInstance")
    
    
    return
}

func NewInquiryPriceRenewInstanceResponse() (response *InquiryPriceRenewInstanceResponse) {
    response = &InquiryPriceRenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceRenewInstance
// 本接口（InquiryPriceRenewInstance）用于查询包年包月计费实例的续费价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) InquiryPriceRenewInstance(request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    return c.InquiryPriceRenewInstanceWithContext(context.Background(), request)
}

// InquiryPriceRenewInstance
// 本接口（InquiryPriceRenewInstance）用于查询包年包月计费实例的续费价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) InquiryPriceRenewInstanceWithContext(ctx context.Context, request *InquiryPriceRenewInstanceRequest) (response *InquiryPriceRenewInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceRenewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceRenewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceRenewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceUpgradeInstanceRequest() (request *InquiryPriceUpgradeInstanceRequest) {
    request = &InquiryPriceUpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "InquiryPriceUpgradeInstance")
    
    
    return
}

func NewInquiryPriceUpgradeInstanceResponse() (response *InquiryPriceUpgradeInstanceResponse) {
    response = &InquiryPriceUpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InquiryPriceUpgradeInstance
// 查询实例扩容价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) InquiryPriceUpgradeInstance(request *InquiryPriceUpgradeInstanceRequest) (response *InquiryPriceUpgradeInstanceResponse, err error) {
    return c.InquiryPriceUpgradeInstanceWithContext(context.Background(), request)
}

// InquiryPriceUpgradeInstance
// 查询实例扩容价格
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) InquiryPriceUpgradeInstanceWithContext(ctx context.Context, request *InquiryPriceUpgradeInstanceRequest) (response *InquiryPriceUpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceUpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewKillMasterGroupRequest() (request *KillMasterGroupRequest) {
    request = &KillMasterGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "KillMasterGroup")
    
    
    return
}

func NewKillMasterGroupResponse() (response *KillMasterGroupResponse) {
    response = &KillMasterGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillMasterGroup
// 本接口（KillMasterGroup）模拟故障。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTOPERATION = "ResourceUnavailable.InstanceNotSupportOperation"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) KillMasterGroup(request *KillMasterGroupRequest) (response *KillMasterGroupResponse, err error) {
    return c.KillMasterGroupWithContext(context.Background(), request)
}

// KillMasterGroup
// 本接口（KillMasterGroup）模拟故障。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTOPERATION = "ResourceUnavailable.InstanceNotSupportOperation"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) KillMasterGroupWithContext(ctx context.Context, request *KillMasterGroupRequest) (response *KillMasterGroupResponse, err error) {
    if request == nil {
        request = NewKillMasterGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillMasterGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillMasterGroupResponse()
    err = c.Send(request, response)
    return
}

func NewManualBackupInstanceRequest() (request *ManualBackupInstanceRequest) {
    request = &ManualBackupInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ManualBackupInstance")
    
    
    return
}

func NewManualBackupInstanceResponse() (response *ManualBackupInstanceResponse) {
    response = &ManualBackupInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManualBackupInstance
// 本接口（ManualBackupInstance）用于手动备份Redis实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMMITFLOWERROR = "FailedOperation.CommitFlowError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ManualBackupInstance(request *ManualBackupInstanceRequest) (response *ManualBackupInstanceResponse, err error) {
    return c.ManualBackupInstanceWithContext(context.Background(), request)
}

// ManualBackupInstance
// 本接口（ManualBackupInstance）用于手动备份Redis实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION_COMMITFLOWERROR = "FailedOperation.CommitFlowError"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ManualBackupInstanceWithContext(ctx context.Context, request *ManualBackupInstanceRequest) (response *ManualBackupInstanceResponse, err error) {
    if request == nil {
        request = NewManualBackupInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManualBackupInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewManualBackupInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModfiyInstancePasswordRequest() (request *ModfiyInstancePasswordRequest) {
    request = &ModfiyInstancePasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModfiyInstancePassword")
    
    
    return
}

func NewModfiyInstancePasswordResponse() (response *ModfiyInstancePasswordResponse) {
    response = &ModfiyInstancePasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModfiyInstancePassword
// 本接口（ModfiyInstancePassword）用于修改实例访问密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModfiyInstancePassword(request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    return c.ModfiyInstancePasswordWithContext(context.Background(), request)
}

// ModfiyInstancePassword
// 本接口（ModfiyInstancePassword）用于修改实例访问密码。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModfiyInstancePasswordWithContext(ctx context.Context, request *ModfiyInstancePasswordRequest) (response *ModfiyInstancePasswordResponse, err error) {
    if request == nil {
        request = NewModfiyInstancePasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModfiyInstancePassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModfiyInstancePasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoBackupConfigRequest() (request *ModifyAutoBackupConfigRequest) {
    request = &ModifyAutoBackupConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyAutoBackupConfig")
    
    
    return
}

func NewModifyAutoBackupConfigResponse() (response *ModifyAutoBackupConfigResponse) {
    response = &ModifyAutoBackupConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoBackupConfig
// 本接口（ModifyAutoBackupConfig）用于设置自动备份的配置。
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyAutoBackupConfig(request *ModifyAutoBackupConfigRequest) (response *ModifyAutoBackupConfigResponse, err error) {
    return c.ModifyAutoBackupConfigWithContext(context.Background(), request)
}

// ModifyAutoBackupConfig
// 本接口（ModifyAutoBackupConfig）用于设置自动备份的配置。
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewModifyBackupDownloadRestrictionRequest() (request *ModifyBackupDownloadRestrictionRequest) {
    request = &ModifyBackupDownloadRestrictionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyBackupDownloadRestriction")
    
    
    return
}

func NewModifyBackupDownloadRestrictionResponse() (response *ModifyBackupDownloadRestrictionResponse) {
    response = &ModifyBackupDownloadRestrictionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyBackupDownloadRestriction
// 本接口（ModifyBackupDownloadRestriction）用于修改备份文件下载的网络信息与地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyBackupDownloadRestriction(request *ModifyBackupDownloadRestrictionRequest) (response *ModifyBackupDownloadRestrictionResponse, err error) {
    return c.ModifyBackupDownloadRestrictionWithContext(context.Background(), request)
}

// ModifyBackupDownloadRestriction
// 本接口（ModifyBackupDownloadRestriction）用于修改备份文件下载的网络信息与地址。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewModifyConnectionConfigRequest() (request *ModifyConnectionConfigRequest) {
    request = &ModifyConnectionConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyConnectionConfig")
    
    
    return
}

func NewModifyConnectionConfigResponse() (response *ModifyConnectionConfigResponse) {
    response = &ModifyConnectionConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConnectionConfig
// 修改实例的连接配置，包括带宽和最大连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ModifyConnectionConfig(request *ModifyConnectionConfigRequest) (response *ModifyConnectionConfigResponse, err error) {
    return c.ModifyConnectionConfigWithContext(context.Background(), request)
}

// ModifyConnectionConfig
// 修改实例的连接配置，包括带宽和最大连接数。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyDBInstanceSecurityGroups")
    
    
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
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_INSTANCESGOVERLIMITERROR = "InvalidParameter.InstanceSGOverLimitError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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
//  FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"
//  INTERNALERROR_INSTANCEOPERATEPERMISSIONERROR = "InternalError.InstanceOperatePermissionError"
//  INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"
//  INVALIDPARAMETER_INSTANCESGOVERLIMITERROR = "InvalidParameter.InstanceSGOverLimitError"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstance")
    
    
    return
}

func NewModifyInstanceResponse() (response *ModifyInstanceResponse) {
    response = &ModifyInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstance
// 修改实例相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
func (c *Client) ModifyInstance(request *ModifyInstanceRequest) (response *ModifyInstanceResponse, err error) {
    return c.ModifyInstanceWithContext(context.Background(), request)
}

// ModifyInstance
// 修改实例相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
//  UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
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

func NewModifyInstanceAccountRequest() (request *ModifyInstanceAccountRequest) {
    request = &ModifyInstanceAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceAccount")
    
    
    return
}

func NewModifyInstanceAccountResponse() (response *ModifyInstanceAccountResponse) {
    response = &ModifyInstanceAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceAccount
// 修改实例子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceAccount(request *ModifyInstanceAccountRequest) (response *ModifyInstanceAccountResponse, err error) {
    return c.ModifyInstanceAccountWithContext(context.Background(), request)
}

// ModifyInstanceAccount
// 修改实例子账号
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceAccountWithContext(ctx context.Context, request *ModifyInstanceAccountRequest) (response *ModifyInstanceAccountResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAccountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAccountResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAvailabilityZonesRequest() (request *ModifyInstanceAvailabilityZonesRequest) {
    request = &ModifyInstanceAvailabilityZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceAvailabilityZones")
    
    
    return
}

func NewModifyInstanceAvailabilityZonesResponse() (response *ModifyInstanceAvailabilityZonesResponse) {
    response = &ModifyInstanceAvailabilityZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceAvailabilityZones
// 本接口（ModifyInstanceAvailabilityZones）用于变更实例可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ModifyInstanceAvailabilityZones(request *ModifyInstanceAvailabilityZonesRequest) (response *ModifyInstanceAvailabilityZonesResponse, err error) {
    return c.ModifyInstanceAvailabilityZonesWithContext(context.Background(), request)
}

// ModifyInstanceAvailabilityZones
// 本接口（ModifyInstanceAvailabilityZones）用于变更实例可用区
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ModifyInstanceAvailabilityZonesWithContext(ctx context.Context, request *ModifyInstanceAvailabilityZonesRequest) (response *ModifyInstanceAvailabilityZonesResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAvailabilityZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAvailabilityZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAvailabilityZonesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceEventRequest() (request *ModifyInstanceEventRequest) {
    request = &ModifyInstanceEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceEvent")
    
    
    return
}

func NewModifyInstanceEventResponse() (response *ModifyInstanceEventResponse) {
    response = &ModifyInstanceEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceEvent
// 本接口（ModifyInstanceEvent）用于修改实例的运维事件的执行计划。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ModifyInstanceEvent(request *ModifyInstanceEventRequest) (response *ModifyInstanceEventResponse, err error) {
    return c.ModifyInstanceEventWithContext(context.Background(), request)
}

// ModifyInstanceEvent
// 本接口（ModifyInstanceEvent）用于修改实例的运维事件的执行计划。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) ModifyInstanceEventWithContext(ctx context.Context, request *ModifyInstanceEventRequest) (response *ModifyInstanceEventResponse, err error) {
    if request == nil {
        request = NewModifyInstanceEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceEventResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceLogDeliveryRequest() (request *ModifyInstanceLogDeliveryRequest) {
    request = &ModifyInstanceLogDeliveryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceLogDelivery")
    
    
    return
}

func NewModifyInstanceLogDeliveryResponse() (response *ModifyInstanceLogDeliveryResponse) {
    response = &ModifyInstanceLogDeliveryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceLogDelivery
// 本接口（ModifyInstanceLogDelivery）用于开启或关闭投递实例日志到CLS。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyInstanceLogDelivery(request *ModifyInstanceLogDeliveryRequest) (response *ModifyInstanceLogDeliveryResponse, err error) {
    return c.ModifyInstanceLogDeliveryWithContext(context.Background(), request)
}

// ModifyInstanceLogDelivery
// 本接口（ModifyInstanceLogDelivery）用于开启或关闭投递实例日志到CLS。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyInstanceLogDeliveryWithContext(ctx context.Context, request *ModifyInstanceLogDeliveryRequest) (response *ModifyInstanceLogDeliveryResponse, err error) {
    if request == nil {
        request = NewModifyInstanceLogDeliveryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceLogDelivery require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceLogDeliveryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceParamsRequest() (request *ModifyInstanceParamsRequest) {
    request = &ModifyInstanceParamsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceParams")
    
    
    return
}

func NewModifyInstanceParamsResponse() (response *ModifyInstanceParamsResponse) {
    response = &ModifyInstanceParamsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceParams
// 本接口(ModifyInstanceParams)用于修改Redis实例的参数配置。
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyInstanceParams(request *ModifyInstanceParamsRequest) (response *ModifyInstanceParamsResponse, err error) {
    return c.ModifyInstanceParamsWithContext(context.Background(), request)
}

// ModifyInstanceParams
// 本接口(ModifyInstanceParams)用于修改Redis实例的参数配置。
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
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewModifyInstanceReadOnlyRequest() (request *ModifyInstanceReadOnlyRequest) {
    request = &ModifyInstanceReadOnlyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyInstanceReadOnly")
    
    
    return
}

func NewModifyInstanceReadOnlyResponse() (response *ModifyInstanceReadOnlyResponse) {
    response = &ModifyInstanceReadOnlyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyInstanceReadOnly
// 设置实例输入模式
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyInstanceReadOnly(request *ModifyInstanceReadOnlyRequest) (response *ModifyInstanceReadOnlyResponse, err error) {
    return c.ModifyInstanceReadOnlyWithContext(context.Background(), request)
}

// ModifyInstanceReadOnly
// 设置实例输入模式
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) ModifyInstanceReadOnlyWithContext(ctx context.Context, request *ModifyInstanceReadOnlyRequest) (response *ModifyInstanceReadOnlyResponse, err error) {
    if request == nil {
        request = NewModifyInstanceReadOnlyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceReadOnly require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceReadOnlyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMaintenanceWindowRequest() (request *ModifyMaintenanceWindowRequest) {
    request = &ModifyMaintenanceWindowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyMaintenanceWindow")
    
    
    return
}

func NewModifyMaintenanceWindowResponse() (response *ModifyMaintenanceWindowResponse) {
    response = &ModifyMaintenanceWindowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMaintenanceWindow
// 修改实例维护时间窗时间，需要进行版本升级或者架构升级的实例，会在维护时间窗内进行时间切换。注意：已经发起版本升级或者架构升级的实例，无法修改维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
func (c *Client) ModifyMaintenanceWindow(request *ModifyMaintenanceWindowRequest) (response *ModifyMaintenanceWindowResponse, err error) {
    return c.ModifyMaintenanceWindowWithContext(context.Background(), request)
}

// ModifyMaintenanceWindow
// 修改实例维护时间窗时间，需要进行版本升级或者架构升级的实例，会在维护时间窗内进行时间切换。注意：已经发起版本升级或者架构升级的实例，无法修改维护时间窗。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
//  UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"
//  UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyNetworkConfig")
    
    
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
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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
//  INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  RESOURCEUNAVAILABLE_INSTANCEUNLOCKEDERROR = "ResourceUnavailable.InstanceUnLockedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewModifyParamTemplateRequest() (request *ModifyParamTemplateRequest) {
    request = &ModifyParamTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyParamTemplate")
    
    
    return
}

func NewModifyParamTemplateResponse() (response *ModifyParamTemplateResponse) {
    response = &ModifyParamTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyParamTemplate
// 修改参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) ModifyParamTemplate(request *ModifyParamTemplateRequest) (response *ModifyParamTemplateResponse, err error) {
    return c.ModifyParamTemplateWithContext(context.Background(), request)
}

// ModifyParamTemplate
// 修改参数模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
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

func NewModifyReplicationGroupRequest() (request *ModifyReplicationGroupRequest) {
    request = &ModifyReplicationGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ModifyReplicationGroup")
    
    
    return
}

func NewModifyReplicationGroupResponse() (response *ModifyReplicationGroupResponse) {
    response = &ModifyReplicationGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyReplicationGroup
// 修改复制组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyReplicationGroup(request *ModifyReplicationGroupRequest) (response *ModifyReplicationGroupResponse, err error) {
    return c.ModifyReplicationGroupWithContext(context.Background(), request)
}

// ModifyReplicationGroup
// 修改复制组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ModifyReplicationGroupWithContext(ctx context.Context, request *ModifyReplicationGroupRequest) (response *ModifyReplicationGroupResponse, err error) {
    if request == nil {
        request = NewModifyReplicationGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyReplicationGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyReplicationGroupResponse()
    err = c.Send(request, response)
    return
}

func NewOpenSSLRequest() (request *OpenSSLRequest) {
    request = &OpenSSLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "OpenSSL")
    
    
    return
}

func NewOpenSSLResponse() (response *OpenSSLResponse) {
    response = &OpenSSLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenSSL
// 开启SSL
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) OpenSSL(request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    return c.OpenSSLWithContext(context.Background(), request)
}

// OpenSSL
// 开启SSL
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) OpenSSLWithContext(ctx context.Context, request *OpenSSLRequest) (response *OpenSSLResponse, err error) {
    if request == nil {
        request = NewOpenSSLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenSSL require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenSSLResponse()
    err = c.Send(request, response)
    return
}

func NewReleaseWanAddressRequest() (request *ReleaseWanAddressRequest) {
    request = &ReleaseWanAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "ReleaseWanAddress")
    
    
    return
}

func NewReleaseWanAddressResponse() (response *ReleaseWanAddressResponse) {
    response = &ReleaseWanAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReleaseWanAddress
// 关闭外网
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ReleaseWanAddress(request *ReleaseWanAddressRequest) (response *ReleaseWanAddressResponse, err error) {
    return c.ReleaseWanAddressWithContext(context.Background(), request)
}

// ReleaseWanAddress
// 关闭外网
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
func (c *Client) ReleaseWanAddressWithContext(ctx context.Context, request *ReleaseWanAddressRequest) (response *ReleaseWanAddressResponse, err error) {
    if request == nil {
        request = NewReleaseWanAddressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReleaseWanAddress require credential")
    }

    request.SetContext(ctx)
    
    response = NewReleaseWanAddressResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveReplicationInstanceRequest() (request *RemoveReplicationInstanceRequest) {
    request = &RemoveReplicationInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "RemoveReplicationInstance")
    
    
    return
}

func NewRemoveReplicationInstanceResponse() (response *RemoveReplicationInstanceResponse) {
    response = &RemoveReplicationInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveReplicationInstance
// 移除复制组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) RemoveReplicationInstance(request *RemoveReplicationInstanceRequest) (response *RemoveReplicationInstanceResponse, err error) {
    return c.RemoveReplicationInstanceWithContext(context.Background(), request)
}

// RemoveReplicationInstance
// 移除复制组成员
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
//  INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"
//  LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) RemoveReplicationInstanceWithContext(ctx context.Context, request *RemoveReplicationInstanceRequest) (response *RemoveReplicationInstanceResponse, err error) {
    if request == nil {
        request = NewRemoveReplicationInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveReplicationInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveReplicationInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRenewInstanceRequest() (request *RenewInstanceRequest) {
    request = &RenewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "RenewInstance")
    
    
    return
}

func NewRenewInstanceResponse() (response *RenewInstanceResponse) {
    response = &RenewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewInstance
// 本接口（RenewInstance）可用于为实例续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
func (c *Client) RenewInstance(request *RenewInstanceRequest) (response *RenewInstanceResponse, err error) {
    return c.RenewInstanceWithContext(context.Background(), request)
}

// RenewInstance
// 本接口（RenewInstance）可用于为实例续费。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"
//  LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"
//  RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"
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
    
    request.Init().WithApiInfo("redis", APIVersion, "ResetPassword")
    
    
    return
}

func NewResetPasswordResponse() (response *ResetPasswordResponse) {
    response = &ResetPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetPassword
// 重置密码
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) ResetPassword(request *ResetPasswordRequest) (response *ResetPasswordResponse, err error) {
    return c.ResetPasswordWithContext(context.Background(), request)
}

// ResetPassword
// 重置密码
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"
//  INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewRestoreInstanceRequest() (request *RestoreInstanceRequest) {
    request = &RestoreInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "RestoreInstance")
    
    
    return
}

func NewRestoreInstanceResponse() (response *RestoreInstanceResponse) {
    response = &RestoreInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RestoreInstance
// 恢复 CRS 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPLOCKEDERROR = "ResourceUnavailable.BackupLockedError"
//  RESOURCEUNAVAILABLE_BACKUPSPECERROR = "ResourceUnavailable.BackupSpecError"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSABNORMAL = "ResourceUnavailable.BackupStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) RestoreInstance(request *RestoreInstanceRequest) (response *RestoreInstanceResponse, err error) {
    return c.RestoreInstanceWithContext(context.Background(), request)
}

// RestoreInstance
// 恢复 CRS 实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"
//  INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_BACKUPLOCKEDERROR = "ResourceUnavailable.BackupLockedError"
//  RESOURCEUNAVAILABLE_BACKUPSPECERROR = "ResourceUnavailable.BackupSpecError"
//  RESOURCEUNAVAILABLE_BACKUPSTATUSABNORMAL = "ResourceUnavailable.BackupStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewStartupInstanceRequest() (request *StartupInstanceRequest) {
    request = &StartupInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "StartupInstance")
    
    
    return
}

func NewStartupInstanceResponse() (response *StartupInstanceResponse) {
    response = &StartupInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartupInstance
// 实例解隔离
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) StartupInstance(request *StartupInstanceRequest) (response *StartupInstanceResponse, err error) {
    return c.StartupInstanceWithContext(context.Background(), request)
}

// StartupInstance
// 实例解隔离
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) StartupInstanceWithContext(ctx context.Context, request *StartupInstanceRequest) (response *StartupInstanceResponse, err error) {
    if request == nil {
        request = NewStartupInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartupInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartupInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchAccessNewInstanceRequest() (request *SwitchAccessNewInstanceRequest) {
    request = &SwitchAccessNewInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "SwitchAccessNewInstance")
    
    
    return
}

func NewSwitchAccessNewInstanceResponse() (response *SwitchAccessNewInstanceResponse) {
    response = &SwitchAccessNewInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchAccessNewInstance
// 本接口（SwitchAccessNewInstance）针对处于时间窗口中待切换操作的实例，用户可主动发起该操作。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) SwitchAccessNewInstance(request *SwitchAccessNewInstanceRequest) (response *SwitchAccessNewInstanceResponse, err error) {
    return c.SwitchAccessNewInstanceWithContext(context.Background(), request)
}

// SwitchAccessNewInstance
// 本接口（SwitchAccessNewInstance）针对处于时间窗口中待切换操作的实例，用户可主动发起该操作。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
func (c *Client) SwitchAccessNewInstanceWithContext(ctx context.Context, request *SwitchAccessNewInstanceRequest) (response *SwitchAccessNewInstanceResponse, err error) {
    if request == nil {
        request = NewSwitchAccessNewInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchAccessNewInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchAccessNewInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchInstanceVipRequest() (request *SwitchInstanceVipRequest) {
    request = &SwitchInstanceVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "SwitchInstanceVip")
    
    
    return
}

func NewSwitchInstanceVipResponse() (response *SwitchInstanceVipResponse) {
    response = &SwitchInstanceVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchInstanceVip
// 在通过DTS支持跨可用区灾备的场景中，通过该接口交换实例VIP完成实例灾备切换。交换VIP后目标实例可写，源和目标实例VIP互换，同时源与目标实例间DTS同步任务断开
//
// 可能返回的错误码:
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) SwitchInstanceVip(request *SwitchInstanceVipRequest) (response *SwitchInstanceVipResponse, err error) {
    return c.SwitchInstanceVipWithContext(context.Background(), request)
}

// SwitchInstanceVip
// 在通过DTS支持跨可用区灾备的场景中，通过该接口交换实例VIP完成实例灾备切换。交换VIP后目标实例可写，源和目标实例VIP互换，同时源与目标实例间DTS同步任务断开
//
// 可能返回的错误码:
//  FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"
//  RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
func (c *Client) SwitchInstanceVipWithContext(ctx context.Context, request *SwitchInstanceVipRequest) (response *SwitchInstanceVipResponse, err error) {
    if request == nil {
        request = NewSwitchInstanceVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchInstanceVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchInstanceVipResponse()
    err = c.Send(request, response)
    return
}

func NewSwitchProxyRequest() (request *SwitchProxyRequest) {
    request = &SwitchProxyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "SwitchProxy")
    
    
    return
}

func NewSwitchProxyResponse() (response *SwitchProxyResponse) {
    response = &SwitchProxyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SwitchProxy
// Proxy模拟故障接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) SwitchProxy(request *SwitchProxyRequest) (response *SwitchProxyResponse, err error) {
    return c.SwitchProxyWithContext(context.Background(), request)
}

// SwitchProxy
// Proxy模拟故障接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
func (c *Client) SwitchProxyWithContext(ctx context.Context, request *SwitchProxyRequest) (response *SwitchProxyResponse, err error) {
    if request == nil {
        request = NewSwitchProxyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SwitchProxy require credential")
    }

    request.SetContext(ctx)
    
    response = NewSwitchProxyResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstance
// 本接口（UpgradeInstance）用于变更实例的配置规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MEMSIZENOTINRANGE = "InvalidParameterValue.MemSizeNotInRange"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 本接口（UpgradeInstance）用于变更实例的配置规格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_MEMSIZENOTINRANGE = "InvalidParameterValue.MemSizeNotInRange"
//  INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"
//  INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"
//  RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"
//  RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"
//  RESOURCEUNAVAILABLE_SALEOUT = "ResourceUnavailable.SaleOut"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
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

func NewUpgradeInstanceVersionRequest() (request *UpgradeInstanceVersionRequest) {
    request = &UpgradeInstanceVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeInstanceVersion")
    
    
    return
}

func NewUpgradeInstanceVersionResponse() (response *UpgradeInstanceVersionResponse) {
    response = &UpgradeInstanceVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeInstanceVersion
// 将当前实例升级到更高版本，或者将当前标准架构升级至集群架构。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) UpgradeInstanceVersion(request *UpgradeInstanceVersionRequest) (response *UpgradeInstanceVersionResponse, err error) {
    return c.UpgradeInstanceVersionWithContext(context.Background(), request)
}

// UpgradeInstanceVersion
// 将当前实例升级到更高版本，或者将当前标准架构升级至集群架构。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"
//  UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"
func (c *Client) UpgradeInstanceVersionWithContext(ctx context.Context, request *UpgradeInstanceVersionRequest) (response *UpgradeInstanceVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstanceVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeProxyVersionRequest() (request *UpgradeProxyVersionRequest) {
    request = &UpgradeProxyVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeProxyVersion")
    
    
    return
}

func NewUpgradeProxyVersionResponse() (response *UpgradeProxyVersionResponse) {
    response = &UpgradeProxyVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeProxyVersion
// 实例proxy版本升级
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
func (c *Client) UpgradeProxyVersion(request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    return c.UpgradeProxyVersionWithContext(context.Background(), request)
}

// UpgradeProxyVersion
// 实例proxy版本升级
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"
func (c *Client) UpgradeProxyVersionWithContext(ctx context.Context, request *UpgradeProxyVersionRequest) (response *UpgradeProxyVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeProxyVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeProxyVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeProxyVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeSmallVersionRequest() (request *UpgradeSmallVersionRequest) {
    request = &UpgradeSmallVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeSmallVersion")
    
    
    return
}

func NewUpgradeSmallVersionResponse() (response *UpgradeSmallVersionResponse) {
    response = &UpgradeSmallVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeSmallVersion
// 实例小版本升级
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) UpgradeSmallVersion(request *UpgradeSmallVersionRequest) (response *UpgradeSmallVersionResponse, err error) {
    return c.UpgradeSmallVersionWithContext(context.Background(), request)
}

// UpgradeSmallVersion
// 实例小版本升级
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
func (c *Client) UpgradeSmallVersionWithContext(ctx context.Context, request *UpgradeSmallVersionRequest) (response *UpgradeSmallVersionResponse, err error) {
    if request == nil {
        request = NewUpgradeSmallVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeSmallVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeSmallVersionResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeVersionToMultiAvailabilityZonesRequest() (request *UpgradeVersionToMultiAvailabilityZonesRequest) {
    request = &UpgradeVersionToMultiAvailabilityZonesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("redis", APIVersion, "UpgradeVersionToMultiAvailabilityZones")
    
    
    return
}

func NewUpgradeVersionToMultiAvailabilityZonesResponse() (response *UpgradeVersionToMultiAvailabilityZonesResponse) {
    response = &UpgradeVersionToMultiAvailabilityZonesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeVersionToMultiAvailabilityZones
// 升级实例支持多AZ
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) UpgradeVersionToMultiAvailabilityZones(request *UpgradeVersionToMultiAvailabilityZonesRequest) (response *UpgradeVersionToMultiAvailabilityZonesResponse, err error) {
    return c.UpgradeVersionToMultiAvailabilityZonesWithContext(context.Background(), request)
}

// UpgradeVersionToMultiAvailabilityZones
// 升级实例支持多AZ
//
// 可能返回的错误码:
//  FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"
//  UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"
//  UNSUPPORTEDOPERATION_LIMITPROXYVERSION = "UnsupportedOperation.LimitProxyVersion"
func (c *Client) UpgradeVersionToMultiAvailabilityZonesWithContext(ctx context.Context, request *UpgradeVersionToMultiAvailabilityZonesRequest) (response *UpgradeVersionToMultiAvailabilityZonesResponse, err error) {
    if request == nil {
        request = NewUpgradeVersionToMultiAvailabilityZonesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeVersionToMultiAvailabilityZones require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeVersionToMultiAvailabilityZonesResponse()
    err = c.Send(request, response)
    return
}
